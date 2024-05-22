// Package client is a generic REST client for BasicAuth / JWT services.
package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// LoginResponse is the response payload for /login.
type LoginResponse struct { // copied from sifi/auth/auth.go
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// Options are the options for the client.
type Options struct {
	Username string
	Password string
	Address  string
	Timeout  time.Duration
}

// Client is a connection to a REST service.
type Client struct {
	BaseURL      string
	Password     string
	Username     string
	Logger       *slog.Logger
	Debug        bool
	AccessToken  string
	RefreshToken string
	UserAgent    string
	NewError     func(code int, r io.Reader) error
	LoginPath    string
	http.Client
}

// Response is the response headers and body from a HTTP request.
type Response struct {
	Header http.Header
	Body   io.Reader
}

type flagSet interface {
	StringVar(*string, string, string, string)
	DurationVar(*time.Duration, string, time.Duration, string)
}

// NewOptions returns the Options after binding to the flagset.
func NewOptions(f flagSet, serviceName string, port int) *Options {
	var options Options

	username := "username"
	password := "password"
	address := "address"
	service := "service"
	timeout := "timeout"

	if serviceName != "" { // optional prefix for flag names
		username = serviceName + "-" + username
		password = serviceName + "-" + password
		address = serviceName + "-" + address
		service = serviceName + " " + service
		timeout = serviceName + "-" + timeout
	}

	f.StringVar(&options.Username, username, "", "username for "+service)
	f.StringVar(&options.Password, password, "", "password for "+service)
	f.StringVar(&options.Address, address, fmt.Sprint(":", port), "address for "+service)
	f.DurationVar(&options.Timeout, timeout, 15*time.Second, "timeout for "+service)

	return &options
}

// New returns a new client.
func New(o *Options) *Client {
	return &Client{
		Client: http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // nolint: gosec
			},
			Timeout: o.Timeout,
		},
		BaseURL:   fmt.Sprintf("https://%s", o.Address),
		Username:  o.Username,
		Password:  o.Password,
		Logger:    slog.New(discardHandler{}),
		NewError:  newDefaultError,
		LoginPath: "/login",
	}
}

//go:generate go run "github.com/dmarkham/enumer" -type authType -linecomment -text
type authType int

// Authentication types.
const (
	NoAuth           authType = iota // none
	AccessTokenAuth                  // access-token
	RefreshTokenAuth                 // refresh-token
	BasicAuth                        // basic-auth
)

// Login logs into the service. It is optional to call this method as any
// request will automatically call Login if the Client is missing its Token.
// Call this when you want feedback right away on the acceptance of the
// Username/Password credentials.
//
// If the Client has a RefreshToken it will attempt to use that first, and then
// fall back to BasicAuth if the RefreshToken is invalid.
func (c *Client) Login(ctx context.Context) error {
	if c.RefreshToken != "" {
		if err := c.login(ctx, RefreshTokenAuth); err == nil { // success
			return nil
		}
	}

	return c.login(ctx, BasicAuth)
}

// Post sends a Post request to the given URL.
func (c *Client) Post(ctx context.Context, url string, body, resp interface{}) error {
	return c.requestWrapper(ctx, http.MethodPost, url, body, resp)
}

// Put sends a Put request to the given URL.
func (c *Client) Put(ctx context.Context, url string, body, resp interface{}) error {
	return c.requestWrapper(ctx, http.MethodPut, url, body, resp)
}

// Patch sends a Patch request to the given URL.
func (c *Client) Patch(ctx context.Context, url string, body, resp interface{}) error {
	return c.requestWrapper(ctx, http.MethodPatch, url, body, resp)
}

// Get sends a Get request to the given URL.
func (c *Client) Get(ctx context.Context, url string, resp interface{}) error {
	return c.requestWrapper(ctx, http.MethodGet, url, nil, resp)
}

// Delete sends a Delete request to the given URL.
func (c *Client) Delete(ctx context.Context, url string, resp interface{}) error {
	return c.requestWrapper(ctx, http.MethodDelete, url, nil, resp)
}

// Options sends an Options request to the given URL.
func (c *Client) Options(ctx context.Context, url string, resp interface{}) error {
	return c.requestWrapper(ctx, http.MethodOptions, url, nil, resp)
}

// Request sends a HTTP request to the Server. The optional request body is read
// from r, and any response body is read and returned. Clients can use this
// low-level method to make arbitrary requests and avoid marshaling and
// unmarshaling JSON.
func (c *Client) Request(ctx context.Context, auth authType, method, path string, header map[string]string, body io.Reader) (*Response, error) {
	if body == nil {
		body = http.NoBody
	}

	req, err := http.NewRequestWithContext(ctx, method, c.url(path), body)
	if err != nil {
		return nil, err
	}

	log := c.Logger.With(
		slog.String("method", method),
		slog.String("path", path))

	switch auth {
	case AccessTokenAuth:
		req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	case RefreshTokenAuth:
		req.Header.Set("Authorization", "Bearer "+c.RefreshToken)
	case BasicAuth:
		req.SetBasicAuth(c.Username, c.Password)
	}

	log.Debug("request", slog.Any("authorization", auth))

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	if c.Debug {
		b, err := httputil.DumpRequest(req, true)
		if err != nil {
			return nil, err
		}

		fmt.Printf("REQUEST:\n%s", hex.Dump(b))
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if c.Debug {
		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}

		fmt.Printf("RESPONSE:\n%s", hex.Dump(b))
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, errUnauthorized // special error to trigger re-Login on stale token
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
		return nil, c.NewError(resp.StatusCode, resp.Body)
	}

	var buf bytes.Buffer

	if _, err = buf.ReadFrom(resp.Body); err != nil {
		return nil, err
	}

	return &Response{Header: resp.Header.Clone(), Body: &buf}, nil
}

func (c *Client) login(ctx context.Context, auth authType) error {
	resp, err := c.Request(ctx, auth, http.MethodGet, c.LoginPath, nil, http.NoBody)
	if err != nil {
		return err
	}

	var login LoginResponse

	if err := json.NewDecoder(resp.Body).Decode(&login); err != nil {
		return err
	}

	c.AccessToken = login.AccessToken
	c.RefreshToken = login.RefreshToken

	return nil
}

// url returns the url for the given path.
func (c *Client) url(path string) string {
	path = strings.TrimPrefix(path, "/") // prevent accidental double slash

	return c.BaseURL + "/" + path
}

// requestWrapper sends a method request to the given URL. If the Client does not
// have a token the Login method is called first. If a 401 response is received
// it assumes the token has expired and will re-Login and retry the request.
func (c *Client) requestWrapper(ctx context.Context, method, path string, in, out interface{}) error {
	if c == nil {
		return errors.New("no service")
	}

	if c.AccessToken == "" {
		c.Logger.Info("not logged in")
		if err := c.Login(ctx); err != nil {
			return err
		}
	}

	request := func() error { return c.request(ctx, method, path, in, out) }

	if err := request(); err != nil {
		if !errors.Is(err, errUnauthorized) {
			return err
		}

		c.Logger.Info("token rejected")
		if err := c.Login(ctx); err != nil {
			return err
		}

		return request()
	}

	return nil
}

func (c *Client) request(ctx context.Context, method, path string, in, out interface{}) error {
	var body io.Reader = http.NoBody

	if in != nil { // encode body
		b, err := json.Marshal(in)
		if err != nil {
			return err
		}

		body = bytes.NewReader(b)
	}

	resp, err := c.Request(ctx, AccessTokenAuth, method, path, nil, body)
	if err != nil {
		return err
	}

	if out != nil { // parse response
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return err
		}
	}

	return nil
}

var errUnauthorized = errors.New("unauthorized")

type defaultError struct {
	Code int
	Body string
}

func (e defaultError) Error() string {
	return fmt.Sprintf("%s (%d) - %s", strings.ToLower(http.StatusText(e.Code)), e.Code, e.Body)
}

func newDefaultError(code int, r io.Reader) error {
	var b bytes.Buffer

	if _, err := b.ReadFrom(r); err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	return &defaultError{Code: code, Body: b.String()}
}

// TODO: something very much like this will eventually be in the slog package (go 1.23 or later).

// discardHandler is a slog.Handler that discards all log records.
type discardHandler struct{}

func (discardHandler) Enabled(context.Context, slog.Level) bool  { return false }
func (discardHandler) Handle(context.Context, slog.Record) error { return nil }
func (d discardHandler) WithAttrs([]slog.Attr) slog.Handler      { return d }
func (d discardHandler) WithGroup(string) slog.Handler           { return d }
