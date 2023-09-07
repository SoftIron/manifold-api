package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/deprecated/v1/hc"
)

// SystemService owns the /system methods.
type SystemService struct {
	*client.Client
}

// SystemVersion returns the version of HC.
func (s SystemService) SystemVersion(ctx context.Context) (*hc.ListSystemVersionResponse, error) {
	var resp hc.ListSystemVersionResponse

	p := path(hc.SystemPath, "version")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SystemConfig returns HC config.
func (s SystemService) SystemConfig(ctx context.Context) (*hc.ListSystemConfigResponse, error) {
	var resp hc.ListSystemConfigResponse

	p := path(hc.SystemPath, "config")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
