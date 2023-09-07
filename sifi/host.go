package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// HostService owns the /host methods.
type HostService struct {
	*client.Client
}

// Hosts returns information about all physical hosts.
func (s HostService) Hosts(ctx context.Context) (*cloud.ListHostsResponse, error) {
	var resp cloud.ListHostsResponse

	p := path(cloud.HostPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
