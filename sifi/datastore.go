package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// DatastoreService owns the /datastore methods.
type DatastoreService struct {
	*client.Client
}

// DeleteDatastore deletes the datastore with the given ID.
func (s DatastoreService) DeleteDatastore(ctx context.Context, id string) error {
	p := path(cloud.DatastorePath, id)

	return s.Delete(ctx, p, nil)
}

// Datastore returns information about a datastore.
func (s DatastoreService) Datastore(ctx context.Context, id int) (*cloud.ListDatastoreResponse, error) {
	var resp cloud.ListDatastoreResponse

	p := path(cloud.DatastorePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Datastores returns information about all datastores.
func (s DatastoreService) Datastores(ctx context.Context) (*cloud.ListDatastoresResponse, error) {
	var resp cloud.ListDatastoresResponse

	p := path(cloud.DatastorePath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateDatastore updates the datastore with the given ID.
func (s DatastoreService) UpdateDatastore(ctx context.Context, id int, req cloud.UpdateDatastoreRequest) (*cloud.UpdateDatastoreResponse, error) {
	var resp cloud.UpdateDatastoreResponse

	p := path(cloud.DatastorePath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableDatastore enables the datastore with the given ID.
func (s DatastoreService) EnableDatastore(ctx context.Context, id int, req cloud.EnableDatastoreRequest) (*cloud.EnableDatastoreResponse, error) {
	var resp cloud.EnableDatastoreResponse

	p := path(cloud.DatastorePath, id, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameDatastore renames the datastore with the given ID.
func (s DatastoreService) RenameDatastore(ctx context.Context, id string, req cloud.RenameDatastoreRequest) (*cloud.RenameDatastoreResponse, error) {
	var resp cloud.RenameDatastoreResponse

	p := path(cloud.DatastorePath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeDatastoreOwnership changes the datastore ownership with the given ID.
func (s DatastoreService) ChangeDatastoreOwnership(ctx context.Context, id string, req cloud.ChangeDatastoreOwnershipRequest) (*cloud.ChangeDatastoreOwnershipResponse, error) {
	var resp cloud.ChangeDatastoreOwnershipResponse

	p := path(cloud.DatastorePath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeDatastorePermissions changes the datastore permissions with the given ID.
func (s DatastoreService) ChangeDatastorePermissions(ctx context.Context, id string, req cloud.ChangeDatastorePermissionsRequest) (*cloud.ChangeDatastorePermissionsResponse, error) {
	var resp cloud.ChangeDatastorePermissionsResponse

	p := path(cloud.DatastorePath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateDatastore creates a new datastore.
func (s DatastoreService) CreateDatastore(ctx context.Context, req cloud.CreateDatastoreRequest) (*cloud.CreateDatastoreResponse, error) {
	var resp cloud.CreateDatastoreResponse

	p := path(cloud.DatastorePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
