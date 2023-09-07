package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// ClusterService owns the /cluster methods.
type ClusterService struct {
	*client.Client
}

// DeleteCluster deletes the cluster with the given ID.
func (s ClusterService) DeleteCluster(ctx context.Context, id int) error {
	p := path(cloud.ClusterPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteClusterDatastore deletes the datastore from the cluster.
func (s ClusterService) DeleteClusterDatastore(ctx context.Context, id, datastore int) error {
	p := path(cloud.ClusterPath, id, "datastore", datastore)

	return s.Delete(ctx, p, nil)
}

// DeleteClusterHost deletes the host from the cluster.
func (s ClusterService) DeleteClusterHost(ctx context.Context, id, host int) error {
	p := path(cloud.ClusterPath, id, "host", host)

	return s.Delete(ctx, p, nil)
}

// DeleteClusterNetwork deletes the network from the cluster.
func (s ClusterService) DeleteClusterNetwork(ctx context.Context, id, network int) error {
	p := path(cloud.ClusterPath, id, "network", network)

	return s.Delete(ctx, p, nil)
}

// Cluster returns information about a cluster.
func (s ClusterService) Cluster(ctx context.Context, id int) (*cloud.ListClusterResponse, error) {
	var resp cloud.ListClusterResponse

	p := path(cloud.ClusterPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Clusters returns information about all clusters.
func (s ClusterService) Clusters(ctx context.Context) (*cloud.ListClustersResponse, error) {
	var resp cloud.ListClustersResponse

	p := path(cloud.ClusterPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateCluster updates the cluster.
func (s ClusterService) UpdateCluster(ctx context.Context, id int, req cloud.UpdateClusterRequest) (*cloud.UpdateClusterResponse, error) {
	var resp cloud.UpdateClusterResponse

	p := path(cloud.ClusterPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddClusterDatastore adds the datastore to the cluster.
func (s ClusterService) AddClusterDatastore(ctx context.Context, id, datastore int) (*cloud.AddClusterDatastoreResponse, error) {
	var resp cloud.AddClusterDatastoreResponse

	p := path(cloud.ClusterPath, id, "datastore", datastore)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddClusterHost adds the host to the cluster.
func (s ClusterService) AddClusterHost(ctx context.Context, id, host int) (*cloud.AddClusterHostResponse, error) {
	var resp cloud.AddClusterHostResponse

	p := path(cloud.ClusterPath, id, "host", host)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameCluster updates the name of the cluster.
func (s ClusterService) RenameCluster(ctx context.Context, id int, req cloud.RenameClusterRequest) (*cloud.RenameClusterResponse, error) {
	var resp cloud.RenameClusterResponse

	p := path(cloud.ClusterPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddClusterNetwork adds the network to the cluster.
func (s ClusterService) AddClusterNetwork(ctx context.Context, id, network int) (*cloud.AddClusterNetworkResponse, error) {
	var resp cloud.AddClusterNetworkResponse

	p := path(cloud.ClusterPath, id, "network", network)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateCluster creates a cluster.
func (s ClusterService) CreateCluster(ctx context.Context, req cloud.CreateClusterRequest) (*cloud.CreateClusterResponse, error) {
	var resp cloud.CreateClusterResponse

	p := path(cloud.ClusterPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
