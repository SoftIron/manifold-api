package snapshot

import (
	"context"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/internal/api"
)

// Service owns the /snapper methods.
type Service struct {
	*client.Client
	root string
}

// NewService returns a new Service for snapshot daemon operations.
func NewService(c *client.Client, path string) *Service {
	return &Service{Client: c, root: path}
}

// path returns a URL path with the correct prefix appended.
func (s Service) path(dirs ...interface{}) string {
	return api.Path(s.root, PathPrefix, dirs...)
}

// Status returns the status of the snapshot daemon.
func (s Service) Status(ctx context.Context) (*ListStatusResponse, error) {
	var resp ListStatusResponse

	p := s.path("status")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Stack returns the runtime stack for all the goroutines of the snapshot
// daemon.
func (s Service) Stack(ctx context.Context) (*ListStackResponse, error) {
	var resp ListStackResponse

	p := s.path("stack")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ManualSnapshots returns the list of manual snapshots.
func (s Service) ManualSnapshots(ctx context.Context) (*ListManualResponse, error) {
	var resp ListManualResponse

	p := s.path("snapshots", "manual")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RemoteSnapshots returns the list of remote snapshots.
func (s Service) RemoteSnapshots(ctx context.Context) (*ListRemoteResponse, error) {
	var resp ListRemoteResponse

	p := s.path("snapshots", "remote")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ArchiveSnapshots returns the list of archive snapshots.
func (s Service) ArchiveSnapshots(ctx context.Context) (*ListArchiveResponse, error) {
	var resp ListArchiveResponse

	p := s.path("snapshots", "archive")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
