package cloud

import (
	"context"
)

// SystemService owns the /cloud/system methods.
type SystemService struct {
	*service
}

// SystemVersion returns the version of HC.
func (s SystemService) SystemVersion(ctx context.Context) (*SystemVersionResponse, error) {
	var resp SystemVersionResponse

	p := s.path(SystemPath, "version")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SystemConfig returns HC config.
func (s SystemService) SystemConfig(ctx context.Context) (*SystemConfigResponse, error) {
	var resp SystemConfigResponse

	p := s.path(SystemPath, "config")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
