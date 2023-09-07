package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// SystemService owns the /system methods.
type SystemService struct {
	*client.Client
}

// SystemVersion returns the version of HC.
func (s SystemService) SystemVersion(ctx context.Context) (*cloud.ListSystemVersionResponse, error) {
	var resp cloud.ListSystemVersionResponse

	p := path(cloud.SystemPath, "version")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SystemConfig returns HC config.
func (s SystemService) SystemConfig(ctx context.Context) (*cloud.ListSystemConfigResponse, error) {
	var resp cloud.ListSystemConfigResponse

	p := path(cloud.SystemPath, "config")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
