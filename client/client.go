// Package client is a generic REST client for BasicAuth / JWT services.
package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// LoginResponse is the response payload for /login.
type LoginResponse struct {
	Token string `json:"token"`
}

// Options are the options for the client.
type Options struct {
	Username string
	Password string
	Address  string
}

// Client is a connection to a REST service.
type Client struct {
	BaseURL  string
	Password string
	Username string
	Logger   logger
	Debug    bool
	Token    string
	http.Client
}

type logger interface {
	Error(string, ...interface{})
	Warn(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
}

type flagSet interface {
	StringVar(*string, string, string, string)
}

type nullLogger struct{}

func (n nullLogger) Error(string, ...interface{}) {}
func (n nullLogger) Warn(string, ...interface{})  {}
func (n nullLogger) Info(string, ...interface{})  {}
func (n nullLogger) Debug(string, ...interface{}) {}

// NewOptions returns the Options after binding to the flagset.
func NewOptions(f flagSet, serviceName string, port int) *Options {
	var options Options

	username := "username"
	password := "password"
	address := "address"
	service := "service"

	if serviceName != "" { // optional prefix for flag names
		username = serviceName + "-" + username
		password = serviceName + "-" + password
		address = serviceName + "-" + address
		service = serviceName + " " + service
	}

	f.StringVar(&options.Username, username, "", "username for "+service)
	f.StringVar(&options.Password, password, "", "password for "+service)
	f.StringVar(&options.Address, address, fmt.Sprint("localhost:", port), "address for "+service)

	return &options
}

// New returns a new client.
func New(o *Options) *Client {
	return &Client{
		Client: http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // nolint: gosec
			},
			Timeout: 5 * time.Second,
		},
		BaseURL:  fmt.Sprintf("https://%s", o.Address),
		Username: o.Username,
		Password: o.Password,
		Logger:   nullLogger{},
	}
}

// Login logs into the service. It is optional to call this method as any
// Request will automatically call Login if the Client is missing its Token.
// Call this when you want feedback right away on the acceptance of the
// Username/Password credentials.
func (c *Client) Login(ctx context.Context) error {
	c.Logger.Debug("request", "method", http.MethodGet, "url", c.URL("login"))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.URL("login"), http.NoBody)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	r, err := c.Do(req)
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed: %s (%s)", r.Status, r.Request.URL.Redacted())
	}

	defer r.Body.Close()

	var resp LoginResponse

	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return err
	}

	c.Token = resp.Token

	return nil
}

// URL returns the URL for the given path.
func (c *Client) URL(path string) string {
	return c.BaseURL + "/" + path
}

// Error is an error returned by the HTTP client.
type Error struct {
	Code int
	Text string
}

// Error implements the error interface for h.
func (e Error) Error() string {
	return fmt.Sprintf("sifi %s (%d) - %s", strings.ToLower(http.StatusText(e.Code)), e.Code, e.Text)
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

// requestWrapper sends a method request to the given URL. If the Client does not
// have a token the Login method is called first. If a 401 response is received
// it assumes the token has expired and will re-Login and retry the request.
func (c *Client) requestWrapper(ctx context.Context, method, path string, in, out interface{}) error {
	if c == nil {
		return errors.New("no service")
	}

	c.Logger.Debug("request", "method", method, "path", path)

	if c.Token == "" {
		if err := c.Login(ctx); err != nil {
			return err
		}
	}

	if err := c.request(ctx, method, path, in, out); err != nil {
		var e Error // 401 means the token expired so automatically re-login and try again.

		if errors.As(err, &e) && e.Code == http.StatusUnauthorized {
			c.Logger.Info("token expired")

			if err := c.Login(ctx); err != nil {
				return err
			}

			return c.request(ctx, method, path, in, out)
		}

		return err
	}

	return nil
}

func (c *Client) request(ctx context.Context, method, path string, in, out interface{}) error {
	var r io.Reader = http.NoBody

	if in != nil { // encode body
		b, err := json.Marshal(in)
		if err != nil {
			return err
		}

		r = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.URL(path), r)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if c.Debug {
		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return err
		}

		fmt.Printf("RESPONSE:\n%s", string(b))
	}

	if resp.StatusCode != http.StatusOK {
		var b bytes.Buffer

		_, err := b.ReadFrom(resp.Body)
		if err != nil {
			return Error{Code: resp.StatusCode, Text: err.Error()}
		}

		return Error{Code: resp.StatusCode, Text: b.String()}
	}

	if out != nil { // parse response
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return err
		}
	}

	return nil
}
