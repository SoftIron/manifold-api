package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// VirtualDataCenterService owns the /vdc methods.
type VirtualDataCenterService struct {
	*client.Client
}

// DeleteVirtualDataCenter deletes the VDC with the given ID.
func (s VirtualDataCenterService) DeleteVirtualDataCenter(ctx context.Context, id int) error {
	p := path(cloud.VirtualDataCenterPath, id)

	return s.Delete(ctx, p, nil)
}

// RemoveVDCGroup removes a group from the VDC with the given ID.
func (s VirtualDataCenterService) RemoveVDCGroup(ctx context.Context, id, group int) error {
	p := path(cloud.VirtualDataCenterPath, id, "group", group)

	return s.Delete(ctx, p, nil)
}

// RemoveVDCCluster removes a cluster from the VDC with the given ID.
func (s VirtualDataCenterService) RemoveVDCCluster(ctx context.Context, id, zone, cluster int) error {
	p := path(cloud.VirtualDataCenterPath, id, "zone", zone, "cluster", cluster)

	return s.Delete(ctx, p, nil)
}

// RemoveVDCDatastore removes a datastore from the VDC with the given ID.
func (s VirtualDataCenterService) RemoveVDCDatastore(ctx context.Context, id, zone, datastore int) error {
	p := path(cloud.VirtualDataCenterPath, id, "zone", zone, "datastore", datastore)

	return s.Delete(ctx, p, nil)
}

// RemoveVDCHost removes a host from the VDC with the given ID.
func (s VirtualDataCenterService) RemoveVDCHost(ctx context.Context, id, zone, host int) error {
	p := path(cloud.VirtualDataCenterPath, id, "zone", zone, "host", host)

	return s.Delete(ctx, p, nil)
}

// RemoveVDCNetwork removes a network from the VDC with the given ID.
func (s VirtualDataCenterService) RemoveVDCNetwork(ctx context.Context, id, zone, network int) error {
	p := path(cloud.VirtualDataCenterPath, id, "zone", zone, "network", network)

	return s.Delete(ctx, p, nil)
}

// VirtualDataCenter returns information about a vdc.
func (s VirtualDataCenterService) VirtualDataCenter(ctx context.Context, id int) (*cloud.ListVDCResponse, error) {
	var resp cloud.ListVDCResponse

	p := path(cloud.VirtualDataCenterPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// VirtualDataCenters returns information about all vdcs.
func (s VirtualDataCenterService) VirtualDataCenters(ctx context.Context) (*cloud.ListVDCsResponse, error) {
	var resp cloud.ListVDCsResponse

	p := path(cloud.VirtualDataCenterPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateVirtualDataCenter updates the VDC with the given ID.
func (s VirtualDataCenterService) UpdateVirtualDataCenter(ctx context.Context, id int, req cloud.UpdateVDCRequest) (*cloud.UpdateVDCResponse, error) {
	var resp cloud.UpdateVDCResponse

	p := path(cloud.VirtualDataCenterPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddVirtualDataCenterGroup adds a group to the VDC with the given ID.
func (s VirtualDataCenterService) AddVirtualDataCenterGroup(ctx context.Context, id int) (*cloud.AddVDCGroupResponse, error) {
	var resp cloud.AddVDCGroupResponse

	p := path(cloud.VirtualDataCenterPath, id, "group")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameVirtualDataCenter renames the VDC with the given ID.
func (s VirtualDataCenterService) RenameVirtualDataCenter(ctx context.Context, id int, req cloud.RenameVDCRequest) (*cloud.RenameVDCResponse, error) {
	var resp cloud.RenameVDCResponse

	p := path(cloud.VirtualDataCenterPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddVirtualDataCenterCluster adds a cluster to the VDC with the given ID.
func (s VirtualDataCenterService) AddVirtualDataCenterCluster(ctx context.Context, id, zone, cluster int) (*cloud.AddVDCClusterResponse, error) {
	var resp cloud.AddVDCClusterResponse

	p := path(cloud.VirtualDataCenterPath, id, "zone", zone, "cluster", cluster)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddVirtualDataCenterDatastore adds a datastore to the VDC with the given ID.
func (s VirtualDataCenterService) AddVirtualDataCenterDatastore(ctx context.Context, id, zone, datastore int) (*cloud.AddVDCDatastoreResponse, error) {
	var resp cloud.AddVDCDatastoreResponse

	p := path(cloud.VirtualDataCenterPath, id, "zone", zone, "datastore", datastore)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddVirtualDataCenterHost adds a host to the VDC with the given ID.
func (s VirtualDataCenterService) AddVirtualDataCenterHost(ctx context.Context, id, zone, host int) (*cloud.AddVDCHostResponse, error) {
	var resp cloud.AddVDCHostResponse

	p := path(cloud.VirtualDataCenterPath, id, "zone", zone, "host", host)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddVirtualDataCenterNetwork adds a network to the VDC with the given ID.
func (s VirtualDataCenterService) AddVirtualDataCenterNetwork(ctx context.Context, id, zone, network int) (*cloud.AddVDCNetworkResponse, error) {
	var resp cloud.AddVDCNetworkResponse

	p := path(cloud.VirtualDataCenterPath, id, "zone", zone, "network", network)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateVirtualDataCenter creates a new VDC.
func (s VirtualDataCenterService) CreateVirtualDataCenter(ctx context.Context, req cloud.CreateVDCRequest) (*cloud.CreateVDCResponse, error) {
	var resp cloud.CreateVDCResponse

	p := path(cloud.VirtualDataCenterPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
