package sifi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/cloud"
	"github.com/softiron/manifold-api/metal"
	"github.com/softiron/manifold-api/snapshot"
)

// Client is a connection to the SIFI service.
type Client struct {
	Cloud    *cloud.Service
	Metal    *metal.Service
	Snapshot *snapshot.Service

	*client.Client
}

// WithLogger is an option setting function for NewClient that sets the logger
// on the underlying generic http client.
func WithLogger(logger *slog.Logger) func(*client.Client) {
	return func(c *client.Client) {
		c.Logger = logger
	}
}

// WithUserAgent is an option setting function for NewClient that sets the
// user-agent on the underlying generic http client.
func WithUserAgent(ua string) func(*client.Client) {
	return func(c *client.Client) {
		c.UserAgent = ua
	}
}

// WithDebug is an option setting function for NewClient that sets the debug
// mode on the underlying generic http client.
func WithDebug(b bool) func(*client.Client) {
	return func(c *client.Client) {
		c.Debug = b
	}
}

// NewClient returns a new client.
func NewClient(o *client.Options, fn ...func(*client.Client)) *Client {
	c := client.New(o)
	c.NewError = newError

	for _, f := range fn {
		f(c)
	}

	prefix := APIPrefix + "/" + APIVersion

	return &Client{
		Cloud:    cloud.NewService(c, prefix),
		Metal:    metal.NewService(c, prefix),
		Snapshot: snapshot.NewService(c, prefix),
		Client:   c,
	}
}

func newError(code int, r io.Reader) error {
	var resp ResponseError

	decodeErr := json.NewDecoder(r).Decode(&resp)
	if decodeErr == nil {
		return &resp // error response was correct JSON format
	}

	var b bytes.Buffer

	_, err := b.ReadFrom(r)
	if err != nil { // something really bad happened
		return NewErrorResponse(code, err, "failed to read response body")
	}

	return NewErrorResponse(code, fmt.Errorf("response body unknown format: %w", decodeErr), b.String())
}
