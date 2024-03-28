package sifi

import (
	"context"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/deprecated/v1/hc"
)

// ClusterService owns the /cluster methods.
type ClusterService struct {
	*client.Client
}

// DeleteCluster deletes the cluster with the given ID.
func (s ClusterService) DeleteCluster(ctx context.Context, id int) error {
	p := path(hc.ClusterPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteClusterDatastore deletes the datastore from the cluster.
func (s ClusterService) DeleteClusterDatastore(ctx context.Context, id, datastore int) error {
	p := path(hc.ClusterPath, id, "datastore", datastore)

	return s.Delete(ctx, p, nil)
}

// DeleteClusterHost deletes the host from the cluster.
func (s ClusterService) DeleteClusterHost(ctx context.Context, id, host int) error {
	p := path(hc.ClusterPath, id, "host", host)

	return s.Delete(ctx, p, nil)
}

// DeleteClusterNetwork deletes the network from the cluster.
func (s ClusterService) DeleteClusterNetwork(ctx context.Context, id, network int) error {
	p := path(hc.ClusterPath, id, "network", network)

	return s.Delete(ctx, p, nil)
}

// Cluster returns information about a cluster.
func (s ClusterService) Cluster(ctx context.Context, id int) (*hc.ListClusterResponse, error) {
	var resp hc.ListClusterResponse

	p := path(hc.ClusterPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Clusters returns information about all clusters.
func (s ClusterService) Clusters(ctx context.Context) (*hc.ListClustersResponse, error) {
	var resp hc.ListClustersResponse

	p := path(hc.ClusterPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateCluster updates the cluster.
func (s ClusterService) UpdateCluster(ctx context.Context, id int, req hc.UpdateClusterRequest) (*hc.UpdateClusterResponse, error) {
	var resp hc.UpdateClusterResponse

	p := path(hc.ClusterPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddClusterDatastore adds the datastore to the cluster.
func (s ClusterService) AddClusterDatastore(ctx context.Context, id, datastore int) (*hc.AddClusterDatastoreResponse, error) {
	var resp hc.AddClusterDatastoreResponse

	p := path(hc.ClusterPath, id, "datastore", datastore)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddClusterHost adds the host to the cluster.
func (s ClusterService) AddClusterHost(ctx context.Context, id, host int) (*hc.AddClusterHostResponse, error) {
	var resp hc.AddClusterHostResponse

	p := path(hc.ClusterPath, id, "host", host)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameCluster updates the name of the cluster.
func (s ClusterService) RenameCluster(ctx context.Context, id int, req hc.RenameClusterRequest) (*hc.RenameClusterResponse, error) {
	var resp hc.RenameClusterResponse

	p := path(hc.ClusterPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddClusterNetwork adds the network to the cluster.
func (s ClusterService) AddClusterNetwork(ctx context.Context, id, network int) (*hc.AddClusterNetworkResponse, error) {
	var resp hc.AddClusterNetworkResponse

	p := path(hc.ClusterPath, id, "network", network)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateCluster creates a cluster.
func (s ClusterService) CreateCluster(ctx context.Context, req hc.CreateClusterRequest) (*hc.CreateClusterResponse, error) {
	var resp hc.CreateClusterResponse

	p := path(hc.ClusterPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
