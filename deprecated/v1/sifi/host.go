package sifi

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/deprecated/v1/hc"
)

// HostService owns the /host methods.
type HostService struct {
	*client.Client
}

// DeleteHost deletes the host with the given ID.
func (s HostService) DeleteHost(ctx context.Context, id int) error {
	p := path(hc.HostPath, id)

	return s.Delete(ctx, p, nil)
}

// Host returns information about a host.
func (s HostService) Host(ctx context.Context, id int) (*hc.ListHostResponse, error) {
	var resp hc.ListHostResponse

	p := path(hc.HostPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Hosts returns information about all hosts.
func (s HostService) Hosts(ctx context.Context) (*hc.ListHostsResponse, error) {
	var resp hc.ListHostsResponse

	p := path(hc.HostPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HostMonitoring returns monitoring data for a host.
func (s HostService) HostMonitoring(ctx context.Context, id int) (*hc.ListHostMonitoringResponse, error) {
	var resp hc.ListHostMonitoringResponse

	p := path(hc.HostPath, id, "monitoring")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HostsMonitoring returns monitoring data for all hosts.
func (s HostService) HostsMonitoring(ctx context.Context, interval time.Duration) (*hc.ListHostsMonitoringResponse, error) {
	var resp hc.ListHostsMonitoringResponse

	p := path(hc.HostPath, "monitoring")

	q := make(url.Values)
	q.Add("seconds", strconv.Itoa(int(interval.Seconds())))

	if err := s.Get(ctx, p+"?"+q.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateHost updates the host with the given ID.
func (s HostService) UpdateHost(ctx context.Context, id int, req hc.UpdateHostRequest) (*hc.UpdateHostResponse, error) {
	var resp hc.UpdateHostResponse

	p := path(hc.HostPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameHost renames the host with the given ID.
func (s HostService) RenameHost(ctx context.Context, id int, req hc.RenameHostRequest) (*hc.RenameHostResponse, error) {
	var resp hc.RenameHostResponse

	p := path(hc.HostPath, id, "name")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetHostStatus sets the status of the host with the given ID.
func (s HostService) SetHostStatus(ctx context.Context, id int, req hc.SetHostStatusRequest) (*hc.SetHostStatusResponse, error) {
	var resp hc.SetHostStatusResponse

	p := path(hc.HostPath, id, "status")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateHost creates a new host.
func (s HostService) CreateHost(ctx context.Context, req hc.CreateHostRequest) (*hc.CreateHostResponse, error) {
	var resp hc.CreateHostResponse

	p := path(hc.HostPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
