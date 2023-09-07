package sifi

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// ComputeService owns the /compute methods.
type ComputeService struct {
	*client.Client
}

// DeleteComputeHost deletes the compute host with the given ID.
func (s ComputeService) DeleteComputeHost(ctx context.Context, id int) error {
	p := path(cloud.ComputePath, id)

	return s.Delete(ctx, p, nil)
}

// ComputeHost returns information about a compute host.
func (s ComputeService) ComputeHost(ctx context.Context, id int) (*cloud.ListComputeHostResponse, error) {
	var resp cloud.ListComputeHostResponse

	p := path(cloud.ComputePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ComputeHosts returns information about all compute hosts.
func (s ComputeService) ComputeHosts(ctx context.Context) (*cloud.ListComputeHostsResponse, error) {
	var resp cloud.ListComputeHostsResponse

	p := path(cloud.ComputePath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ComputeHostMonitoring returns monitoring data for a compute host.
func (s ComputeService) ComputeHostMonitoring(ctx context.Context, id int) (*cloud.ListComputeHostMonitoringResponse, error) {
	var resp cloud.ListComputeHostMonitoringResponse

	p := path(cloud.ComputePath, id, "monitoring")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ComputeHostsMonitoring returns monitoring data for all compute hosts.
func (s ComputeService) ComputeHostsMonitoring(ctx context.Context, interval time.Duration) (*cloud.ListComputeHostsMonitoringResponse, error) {
	var resp cloud.ListComputeHostsMonitoringResponse

	p := path(cloud.ComputePath, "monitoring")

	q := make(url.Values)
	q.Add("seconds", strconv.Itoa(int(interval.Seconds())))

	if err := s.Get(ctx, p+"?"+q.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateComputeHost updates the compute host with the given ID.
func (s ComputeService) UpdateComputeHost(ctx context.Context, id int, req cloud.UpdateComputeHostRequest) (*cloud.UpdateComputeHostResponse, error) {
	var resp cloud.UpdateComputeHostResponse

	p := path(cloud.ComputePath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameComputeHost renames the compute host with the given ID.
func (s ComputeService) RenameComputeHost(ctx context.Context, id int, req cloud.RenameComputeHostRequest) (*cloud.RenameComputeHostResponse, error) {
	var resp cloud.RenameComputeHostResponse

	p := path(cloud.ComputePath, id, "name")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetComputeHostStatus sets the status of the compute host with the given ID.
func (s ComputeService) SetComputeHostStatus(ctx context.Context, id int, req cloud.SetComputeHostStatusRequest) (*cloud.SetComputeHostStatusResponse, error) {
	var resp cloud.SetComputeHostStatusResponse

	p := path(cloud.ComputePath, id, "status")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateComputeHost creates a new compute host.
func (s ComputeService) CreateComputeHost(ctx context.Context, req cloud.CreateComputeHostRequest) (*cloud.CreateComputeHostResponse, error) {
	var resp cloud.CreateComputeHostResponse

	p := path(cloud.ComputePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
