package sifi

import (
	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
	v1 "github.com/softiron/hypercloud-api/deprecated/v1/sifi"
	"github.com/softiron/hypercloud-api/metal"
	"github.com/softiron/hypercloud-api/snapshot"
)

// Client is a connection to the SIFI service.
type Client struct {
	Cloud    *cloud.Service
	Metal    *metal.Service
	Snapshot *snapshot.Service
	V1       v1.Client // deprecated API

	*client.Client
}

// NewClient returns a new client.
func NewClient(o *client.Options) *Client {
	c := client.New(o)

	return &Client{
		Cloud:    cloud.NewService(c, APIVersionPath),
		Metal:    metal.NewService(c, APIVersionPath),
		Snapshot: snapshot.NewService(c, APIVersionPath),
		V1:       *v1.NewClient(o),

		Client: c,
	}
}
