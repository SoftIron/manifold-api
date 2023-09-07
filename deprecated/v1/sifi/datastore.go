package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/deprecated/v1/hc"
)

// DatastoreService owns the /datastore methods.
type DatastoreService struct {
	*client.Client
}

// DeleteDatastore deletes the datastore with the given ID.
func (s DatastoreService) DeleteDatastore(ctx context.Context, id string) error {
	p := path(hc.DatastorePath, id)

	return s.Delete(ctx, p, nil)
}

// Datastore returns information about a datastore.
func (s DatastoreService) Datastore(ctx context.Context, id int) (*hc.ListDatastoreResponse, error) {
	var resp hc.ListDatastoreResponse

	p := path(hc.DatastorePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Datastores returns information about all datastores.
func (s DatastoreService) Datastores(ctx context.Context) (*hc.ListDatastoresResponse, error) {
	var resp hc.ListDatastoresResponse

	p := path(hc.DatastorePath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateDatastore updates the datastore with the given ID.
func (s DatastoreService) UpdateDatastore(ctx context.Context, id int, req hc.UpdateDatastoreRequest) (*hc.UpdateDatastoreResponse, error) {
	var resp hc.UpdateDatastoreResponse

	p := path(hc.DatastorePath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableDatastore enables the datastore with the given ID.
func (s DatastoreService) EnableDatastore(ctx context.Context, id int, req hc.EnableDatastoreRequest) (*hc.EnableDatastoreResponse, error) {
	var resp hc.EnableDatastoreResponse

	p := path(hc.DatastorePath, id, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameDatastore renames the datastore with the given ID.
func (s DatastoreService) RenameDatastore(ctx context.Context, id string, req hc.RenameDatastoreRequest) (*hc.RenameDatastoreResponse, error) {
	var resp hc.RenameDatastoreResponse

	p := path(hc.DatastorePath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeDatastoreOwnership changes the datastore ownership with the given ID.
func (s DatastoreService) ChangeDatastoreOwnership(ctx context.Context, id string, req hc.ChangeDatastoreOwnershipRequest) (*hc.ChangeDatastoreOwnershipResponse, error) {
	var resp hc.ChangeDatastoreOwnershipResponse

	p := path(hc.DatastorePath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeDatastorePermissions changes the datastore permissions with the given ID.
func (s DatastoreService) ChangeDatastorePermissions(ctx context.Context, id string, req hc.ChangeDatastorePermissionsRequest) (*hc.ChangeDatastorePermissionsResponse, error) {
	var resp hc.ChangeDatastorePermissionsResponse

	p := path(hc.DatastorePath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateDatastore creates a new datastore.
func (s DatastoreService) CreateDatastore(ctx context.Context, req hc.CreateDatastoreRequest) (*hc.CreateDatastoreResponse, error) {
	var resp hc.CreateDatastoreResponse

	p := path(hc.DatastorePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
