package cloud

import (
	"context"
)

// ClusterService owns the /cloud/cluster methods.
type ClusterService struct {
	*service
}

// DeleteCluster deletes the cluster with the given ID.
func (s ClusterService) DeleteCluster(ctx context.Context, id int) error {
	p := s.path(ClusterPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteClusterDatastore deletes the datastore from the cluster.
func (s ClusterService) DeleteClusterDatastore(ctx context.Context, id, datastore int) error {
	p := s.path(ClusterPath, id, "datastore", datastore)

	return s.Delete(ctx, p, nil)
}

// DeleteClusterHost deletes the host from the cluster.
func (s ClusterService) DeleteClusterHost(ctx context.Context, id, host int) error {
	p := s.path(ClusterPath, id, "host", host)

	return s.Delete(ctx, p, nil)
}

// DeleteClusterNetwork deletes the network from the cluster.
func (s ClusterService) DeleteClusterNetwork(ctx context.Context, id, network int) error {
	p := s.path(ClusterPath, id, "network", network)

	return s.Delete(ctx, p, nil)
}

// Cluster returns information about a cluster.
func (s ClusterService) Cluster(ctx context.Context, id int) (*ClusterResponse, error) {
	var resp ClusterResponse

	p := s.path(ClusterPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Clusters returns information about all clusters.
func (s ClusterService) Clusters(ctx context.Context) (*ClustersResponse, error) {
	var resp ClustersResponse

	p := s.path(ClusterPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateCluster updates the cluster.
func (s ClusterService) UpdateCluster(ctx context.Context, id int, req UpdateClusterRequest) (*UpdateClusterResponse, error) {
	var resp UpdateClusterResponse

	p := s.path(ClusterPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddClusterDatastore adds the datastore to the cluster.
func (s ClusterService) AddClusterDatastore(ctx context.Context, id, datastore int) (*AddClusterDatastoreResponse, error) {
	var resp AddClusterDatastoreResponse

	p := s.path(ClusterPath, id, "datastore", datastore)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddClusterHost adds the host to the cluster.
func (s ClusterService) AddClusterHost(ctx context.Context, id, host int) (*AddClusterHostResponse, error) {
	var resp AddClusterHostResponse

	p := s.path(ClusterPath, id, "host", host)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameCluster updates the name of the cluster.
func (s ClusterService) RenameCluster(ctx context.Context, id int, req RenameClusterRequest) (*RenameClusterResponse, error) {
	var resp RenameClusterResponse

	p := s.path(ClusterPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddClusterNetwork adds the network to the cluster.
func (s ClusterService) AddClusterNetwork(ctx context.Context, id, network int) (*AddClusterNetworkResponse, error) {
	var resp AddClusterNetworkResponse

	p := s.path(ClusterPath, id, "network", network)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateCluster creates a cluster.
func (s ClusterService) CreateCluster(ctx context.Context, req CreateClusterRequest) (*CreateClusterResponse, error) {
	var resp CreateClusterResponse

	p := s.path(ClusterPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
