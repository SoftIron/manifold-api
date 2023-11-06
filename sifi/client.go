package sifi

import (
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
	prefix := APIPrefix + "/" + APIVersion

	return &Client{
		Cloud:    cloud.NewService(c, prefix),
		Metal:    metal.NewService(c, prefix),
		Snapshot: snapshot.NewService(c, prefix),

		Client: c,
	}
}
