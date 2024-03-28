package sifi

import (
	"context"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/deprecated/v1/hc"
)

// VirtualDataCenterService owns the /vdc methods.
type VirtualDataCenterService struct {
	*client.Client
}

// DeleteVirtualDataCenter deletes the VDC with the given ID.
func (s VirtualDataCenterService) DeleteVirtualDataCenter(ctx context.Context, id int) error {
	p := path(hc.VirtualDataCenterPath, id)

	return s.Delete(ctx, p, nil)
}

// RemoveVDCGroup removes a group from the VDC with the given ID.
func (s VirtualDataCenterService) RemoveVDCGroup(ctx context.Context, id, group int) error {
	p := path(hc.VirtualDataCenterPath, id, "group", group)

	return s.Delete(ctx, p, nil)
}

// RemoveVDCCluster removes a cluster from the VDC with the given ID.
func (s VirtualDataCenterService) RemoveVDCCluster(ctx context.Context, id, zone, cluster int) error {
	p := path(hc.VirtualDataCenterPath, id, "zone", zone, "cluster", cluster)

	return s.Delete(ctx, p, nil)
}

// RemoveVDCDatastore removes a datastore from the VDC with the given ID.
func (s VirtualDataCenterService) RemoveVDCDatastore(ctx context.Context, id, zone, datastore int) error {
	p := path(hc.VirtualDataCenterPath, id, "zone", zone, "datastore", datastore)

	return s.Delete(ctx, p, nil)
}

// RemoveVDCHost removes a host from the VDC with the given ID.
func (s VirtualDataCenterService) RemoveVDCHost(ctx context.Context, id, zone, host int) error {
	p := path(hc.VirtualDataCenterPath, id, "zone", zone, "host", host)

	return s.Delete(ctx, p, nil)
}

// RemoveVDCNetwork removes a network from the VDC with the given ID.
func (s VirtualDataCenterService) RemoveVDCNetwork(ctx context.Context, id, zone, network int) error {
	p := path(hc.VirtualDataCenterPath, id, "zone", zone, "network", network)

	return s.Delete(ctx, p, nil)
}

// VirtualDataCenter returns information about a vdc.
func (s VirtualDataCenterService) VirtualDataCenter(ctx context.Context, id int) (*hc.ListVDCResponse, error) {
	var resp hc.ListVDCResponse

	p := path(hc.VirtualDataCenterPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// VirtualDataCenters returns information about all vdcs.
func (s VirtualDataCenterService) VirtualDataCenters(ctx context.Context) (*hc.ListVDCsResponse, error) {
	var resp hc.ListVDCsResponse

	p := path(hc.VirtualDataCenterPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateVirtualDataCenter updates the VDC with the given ID.
func (s VirtualDataCenterService) UpdateVirtualDataCenter(ctx context.Context, id int, req hc.UpdateVDCRequest) (*hc.UpdateVDCResponse, error) {
	var resp hc.UpdateVDCResponse

	p := path(hc.VirtualDataCenterPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddVirtualDataCenterGroup adds a group to the VDC with the given ID.
func (s VirtualDataCenterService) AddVirtualDataCenterGroup(ctx context.Context, id int) (*hc.AddVDCGroupResponse, error) {
	var resp hc.AddVDCGroupResponse

	p := path(hc.VirtualDataCenterPath, id, "group")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameVirtualDataCenter renames the VDC with the given ID.
func (s VirtualDataCenterService) RenameVirtualDataCenter(ctx context.Context, id int, req hc.RenameVDCRequest) (*hc.RenameVDCResponse, error) {
	var resp hc.RenameVDCResponse

	p := path(hc.VirtualDataCenterPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddVirtualDataCenterCluster adds a cluster to the VDC with the given ID.
func (s VirtualDataCenterService) AddVirtualDataCenterCluster(ctx context.Context, id, zone, cluster int) (*hc.AddVDCClusterResponse, error) {
	var resp hc.AddVDCClusterResponse

	p := path(hc.VirtualDataCenterPath, id, "zone", zone, "cluster", cluster)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddVirtualDataCenterDatastore adds a datastore to the VDC with the given ID.
func (s VirtualDataCenterService) AddVirtualDataCenterDatastore(ctx context.Context, id, zone, datastore int) (*hc.AddVDCDatastoreResponse, error) {
	var resp hc.AddVDCDatastoreResponse

	p := path(hc.VirtualDataCenterPath, id, "zone", zone, "datastore", datastore)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddVirtualDataCenterHost adds a host to the VDC with the given ID.
func (s VirtualDataCenterService) AddVirtualDataCenterHost(ctx context.Context, id, zone, host int) (*hc.AddVDCHostResponse, error) {
	var resp hc.AddVDCHostResponse

	p := path(hc.VirtualDataCenterPath, id, "zone", zone, "host", host)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddVirtualDataCenterNetwork adds a network to the VDC with the given ID.
func (s VirtualDataCenterService) AddVirtualDataCenterNetwork(ctx context.Context, id, zone, network int) (*hc.AddVDCNetworkResponse, error) {
	var resp hc.AddVDCNetworkResponse

	p := path(hc.VirtualDataCenterPath, id, "zone", zone, "network", network)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateVirtualDataCenter creates a new VDC.
func (s VirtualDataCenterService) CreateVirtualDataCenter(ctx context.Context, req hc.CreateVDCRequest) (*hc.CreateVDCResponse, error) {
	var resp hc.CreateVDCResponse

	p := path(hc.VirtualDataCenterPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
