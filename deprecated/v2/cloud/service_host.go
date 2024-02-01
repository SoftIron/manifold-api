package cloud

import (
	"context"
	"net/url"
	"strconv"
	"time"
)

// HostService owns the /cloud/compute methods.
type HostService struct {
	*service
}

// DeleteHost deletes the compute host with the given ID.
func (s HostService) DeleteHost(ctx context.Context, id int) error {
	p := s.path(HostPath, id)

	return s.Delete(ctx, p, nil)
}

// Host returns information about a compute host.
func (s HostService) Host(ctx context.Context, id int) (*HostResponse, error) {
	var resp HostResponse

	p := s.path(HostPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Hosts returns information about all compute hosts.
func (s HostService) Hosts(ctx context.Context) (*HostsResponse, error) {
	var resp HostsResponse

	p := s.path(HostPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HostMonitoring returns monitoring data for a compute host.
func (s HostService) HostMonitoring(ctx context.Context, id int) (*HostMonitoringResponse, error) {
	var resp HostMonitoringResponse

	p := s.path(HostPath, id, "monitoring")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HostsMonitoring returns monitoring data for all compute hosts.
func (s HostService) HostsMonitoring(ctx context.Context, interval time.Duration) (*HostsMonitoringResponse, error) {
	var resp HostsMonitoringResponse

	p := s.path(HostPath, "monitoring")

	q := make(url.Values)
	q.Add("seconds", strconv.Itoa(int(interval.Seconds())))

	if err := s.Get(ctx, p+"?"+q.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateComputeHost updates the compute host with the given ID.
func (s HostService) UpdateComputeHost(ctx context.Context, id int, req UpdateHostRequest) (*UpdateHostResponse, error) {
	var resp UpdateHostResponse

	p := s.path(HostPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameHost renames the compute host with the given ID.
func (s HostService) RenameHost(ctx context.Context, id int, req RenameHostRequest) (*RenameHostResponse, error) {
	var resp RenameHostResponse

	p := s.path(HostPath, id, "name")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetHostStatus sets the status of the compute host with the given ID.
func (s HostService) SetHostStatus(ctx context.Context, id int, req SetHostStatusRequest) (*SetHostStatusResponse, error) {
	var resp SetHostStatusResponse

	p := s.path(HostPath, id, "status")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateHost creates a new compute host.
func (s HostService) CreateHost(ctx context.Context, req CreateHostRequest) (*CreateHostResponse, error) {
	var resp CreateHostResponse

	p := s.path(HostPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
