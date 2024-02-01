package cloud

import (
	"context"
)

// DataCenterService owns the /cloud/datacenter methods.
type DataCenterService struct {
	*service
}

// DeleteDataCenter deletes the Data Center with the given ID.
func (s DataCenterService) DeleteDataCenter(ctx context.Context, id int) error {
	p := s.path(DataCenterPath, id)

	return s.Delete(ctx, p, nil)
}

// RemoveDataCenterGroup removes a group from the Data Center with the given ID.
func (s DataCenterService) RemoveDataCenterGroup(ctx context.Context, id, group int) error {
	p := s.path(DataCenterPath, id, "group", group)

	return s.Delete(ctx, p, nil)
}

// RemoveDataCenterCluster removes a cluster from the Data Center with the given ID.
func (s DataCenterService) RemoveDataCenterCluster(ctx context.Context, id, zone, cluster int) error {
	p := s.path(DataCenterPath, id, "zone", zone, "cluster", cluster)

	return s.Delete(ctx, p, nil)
}

// RemoveDataCenterDatastore removes a datastore from the Data Center with the given ID.
func (s DataCenterService) RemoveDataCenterDatastore(ctx context.Context, id, zone, datastore int) error {
	p := s.path(DataCenterPath, id, "zone", zone, "datastore", datastore)

	return s.Delete(ctx, p, nil)
}

// RemoveDataCenterHost removes a host from the Data Center with the given ID.
func (s DataCenterService) RemoveDataCenterHost(ctx context.Context, id, zone, host int) error {
	p := s.path(DataCenterPath, id, "zone", zone, "host", host)

	return s.Delete(ctx, p, nil)
}

// RemoveDataCenterNetwork removes a network from the Data Center with the given ID.
func (s DataCenterService) RemoveDataCenterNetwork(ctx context.Context, id, zone, network int) error {
	p := s.path(DataCenterPath, id, "zone", zone, "network", network)

	return s.Delete(ctx, p, nil)
}

// DataCenter returns information about a vdc.
func (s DataCenterService) DataCenter(ctx context.Context, id int) (*DataCenterResponse, error) {
	var resp DataCenterResponse

	p := s.path(DataCenterPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// DataCenters returns information about all vdcs.
func (s DataCenterService) DataCenters(ctx context.Context) (*DataCentersResponse, error) {
	var resp DataCentersResponse

	p := s.path(DataCenterPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateDataCenter updates the Data Center with the given ID.
func (s DataCenterService) UpdateDataCenter(ctx context.Context, id int, req UpdateDataCenterRequest) (*UpdateDataCenterResponse, error) {
	var resp UpdateDataCenterResponse

	p := s.path(DataCenterPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddDataCenterGroup adds a group to the Data Center with the given ID.
func (s DataCenterService) AddDataCenterGroup(ctx context.Context, id int) (*AddDataCenterGroupResponse, error) {
	var resp AddDataCenterGroupResponse

	p := s.path(DataCenterPath, id, "group")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameDataCenter renames the Data Center with the given ID.
func (s DataCenterService) RenameDataCenter(ctx context.Context, id int, req RenameDataCenterRequest) (*RenameDataCenterResponse, error) {
	var resp RenameDataCenterResponse

	p := s.path(DataCenterPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddDataCenterCluster adds a cluster to the Data Center with the given ID.
func (s DataCenterService) AddDataCenterCluster(ctx context.Context, id, zone, cluster int) (*AddDataCenterClusterResponse, error) {
	var resp AddDataCenterClusterResponse

	p := s.path(DataCenterPath, id, "zone", zone, "cluster", cluster)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddDataCenterDatastore adds a datastore to the Data Center with the given ID.
func (s DataCenterService) AddDataCenterDatastore(ctx context.Context, id, zone, datastore int) (*AddDataCenterDatastoreResponse, error) {
	var resp AddDataCenterDatastoreResponse

	p := s.path(DataCenterPath, id, "zone", zone, "datastore", datastore)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddDataCenterHost adds a host to the Data Center with the given ID.
func (s DataCenterService) AddDataCenterHost(ctx context.Context, id, zone, host int) (*AddDataCenterHostResponse, error) {
	var resp AddDataCenterHostResponse

	p := s.path(DataCenterPath, id, "zone", zone, "host", host)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddDataCenterNetwork adds a network to the Data Center with the given ID.
func (s DataCenterService) AddDataCenterNetwork(ctx context.Context, id, zone, network int) (*AddDataCenterNetworkResponse, error) {
	var resp AddDataCenterNetworkResponse

	p := s.path(DataCenterPath, id, "zone", zone, "network", network)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateDataCenter creates a new Data Center.
func (s DataCenterService) CreateDataCenter(ctx context.Context, req CreateDataCenterRequest) (*CreateDataCenterResponse, error) {
	var resp CreateDataCenterResponse

	p := s.path(DataCenterPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
