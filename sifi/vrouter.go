package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// VRouterService own the /vrouter methods.
type VRouterService struct {
	*client.Client
}

// DeleteVRouterNIC deletes a NIC from the Virtual Router with the given ID.
func (s VRouterService) DeleteVRouterNIC(ctx context.Context, id int) error {
	p := path(cloud.VirtualRouterPath, "nic", id)

	return s.Delete(ctx, p, nil)
}

// DeleteVRouter deletes a Virtual Router with the given ID.
func (s VRouterService) DeleteVRouter(ctx context.Context, id int) error {
	p := path(cloud.VirtualRouterPath, id)

	return s.Delete(ctx, p, nil)
}

// VRouter returns information about a vrouter.
func (s VRouterService) VRouter(ctx context.Context, id int) (*cloud.ListVRouterResponse, error) {
	var resp cloud.ListVRouterResponse

	p := path(cloud.VirtualRouterPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// VRouters returns information about all vrouters.
func (s VRouterService) VRouters(ctx context.Context) (*cloud.ListVRoutersResponse, error) {
	var resp cloud.ListVRoutersResponse

	p := path(cloud.VirtualRouterPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateVRouter updates a Virtual Router with the given ID.
func (s VRouterService) UpdateVRouter(ctx context.Context, id int, req cloud.UpdateVRouterRequest) (*cloud.UpdateVRouterResponse, error) {
	var resp cloud.UpdateVRouterResponse

	p := path(cloud.VirtualRouterPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstantiateVRouter instantiates a Virtual Router with the given ID.
func (s VRouterService) InstantiateVRouter(ctx context.Context, id int, req cloud.InstantiateVRouterRequest) (*cloud.InstantiateVRouterResponse, error) {
	var resp cloud.InstantiateVRouterResponse

	p := path(cloud.VirtualRouterPath, id, "instantiate")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockVRouter locks a Virtual Router with the given ID.
func (s VRouterService) LockVRouter(ctx context.Context, id int, req cloud.LockVRouterRequest) (*cloud.LockVRouterResponse, error) {
	var resp cloud.LockVRouterResponse

	p := path(cloud.VirtualRouterPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameVRouter renames a Virtual Router with the given ID.
func (s VRouterService) RenameVRouter(ctx context.Context, id int, req cloud.RenameVRouterRequest) (*cloud.RenameVRouterResponse, error) {
	var resp cloud.RenameVRouterResponse

	p := path(cloud.VirtualRouterPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeVRouterOwnership changes the ownership of a Virtual Router with the given ID.
func (s VRouterService) ChangeVRouterOwnership(ctx context.Context, id int, req cloud.ChangeVRouterOwnershipRequest) (*cloud.ChangeVRouterOwnershipResponse, error) {
	var resp cloud.ChangeVRouterOwnershipResponse

	p := path(cloud.VirtualRouterPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeVRouterPermissions changes the permissions of a Virtual Router with the given ID.
func (s VRouterService) ChangeVRouterPermissions(ctx context.Context, id int, req cloud.ChangeVRouterPermissionsRequest) (*cloud.ChangeVRouterPermissionsResponse, error) {
	var resp cloud.ChangeVRouterPermissionsResponse

	p := path(cloud.VirtualRouterPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockVRouter unlocks a Virtual Router with the given ID.
func (s VRouterService) UnlockVRouter(ctx context.Context, id int) (*cloud.UnlockVRouterResponse, error) {
	var resp cloud.UnlockVRouterResponse

	p := path(cloud.VirtualRouterPath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateVRouter creates a new Virtual Router.
func (s VRouterService) CreateVRouter(ctx context.Context, req cloud.CreateVRouterRequest) (*cloud.CreateVRouterResponse, error) {
	var resp cloud.CreateVRouterResponse

	p := path(cloud.VirtualRouterPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateVRouterNIC creates a new NIC for the Virtual Router with the given ID.
func (s VRouterService) CreateVRouterNIC(ctx context.Context, id int, req cloud.CreateVRouterNICRequest) (*cloud.CreateVRouterNICResponse, error) {
	var resp cloud.CreateVRouterNICResponse

	p := path(cloud.VirtualRouterPath, id, "nic")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
