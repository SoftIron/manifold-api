package cloud

import (
	"context"
)

// DatastoreService owns the /cloud/datastore methods.
type DatastoreService struct {
	*service
}

// DeleteDatastore deletes the datastore with the given ID.
func (s DatastoreService) DeleteDatastore(ctx context.Context, id string) error {
	p := s.path(DatastorePath, id)

	return s.Delete(ctx, p, nil)
}

// Datastore returns information about a datastore.
func (s DatastoreService) Datastore(ctx context.Context, id int) (*DatastoreResponse, error) {
	var resp DatastoreResponse

	p := s.path(DatastorePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Datastores returns information about all datastores.
func (s DatastoreService) Datastores(ctx context.Context) (*DatastoresResponse, error) {
	var resp DatastoresResponse

	p := s.path(DatastorePath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateDatastore updates the datastore with the given ID.
func (s DatastoreService) UpdateDatastore(ctx context.Context, id int, req UpdateDatastoreRequest) (*UpdateDatastoreResponse, error) {
	var resp UpdateDatastoreResponse

	p := s.path(DatastorePath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableDatastore enables the datastore with the given ID.
func (s DatastoreService) EnableDatastore(ctx context.Context, id int, req EnableDatastoreRequest) (*EnableDatastoreResponse, error) {
	var resp EnableDatastoreResponse

	p := s.path(DatastorePath, id, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameDatastore renames the datastore with the given ID.
func (s DatastoreService) RenameDatastore(ctx context.Context, id string, req RenameDatastoreRequest) (*RenameDatastoreResponse, error) {
	var resp RenameDatastoreResponse

	p := s.path(DatastorePath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeDatastoreOwnership changes the datastore ownership with the given ID.
func (s DatastoreService) ChangeDatastoreOwnership(ctx context.Context, id string, req ChangeDatastoreOwnershipRequest) (*ChangeDatastoreOwnershipResponse, error) {
	var resp ChangeDatastoreOwnershipResponse

	p := s.path(DatastorePath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeDatastorePermissions changes the datastore permissions with the given ID.
func (s DatastoreService) ChangeDatastorePermissions(ctx context.Context, id string, req ChangeDatastorePermissionsRequest) (*ChangeDatastorePermissionsResponse, error) {
	var resp ChangeDatastorePermissionsResponse

	p := s.path(DatastorePath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateDatastore creates a new datastore.
func (s DatastoreService) CreateDatastore(ctx context.Context, req CreateDatastoreRequest) (*CreateDatastoreResponse, error) {
	var resp CreateDatastoreResponse

	p := s.path(DatastorePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
