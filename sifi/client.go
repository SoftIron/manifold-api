package sifi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
	"github.com/softiron/hypercloud-api/metal"
	"github.com/softiron/hypercloud-api/snapshot"
)

// Client is a connection to the SIFI service.
type Client struct {
	Cloud    *cloud.Service
	Metal    *metal.Service
	Snapshot *snapshot.Service

	*client.Client
}

// NewClient returns a new client.
func NewClient(o *client.Options) *Client {
	c := client.New(o)
	c.NewError = newError

	prefix := APIPrefix + "/" + APIVersion

	return &Client{
		Cloud:    cloud.NewService(c, prefix),
		Metal:    metal.NewService(c, prefix),
		Snapshot: snapshot.NewService(c, prefix),

		Client: c,
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
