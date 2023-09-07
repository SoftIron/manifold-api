package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/deprecated/v1/hc"
)

// VRouterService own the /vrouter methods.
type VRouterService struct {
	*client.Client
}

// DeleteVRouterNIC deletes a NIC from the Virtual Router with the given ID.
func (s VRouterService) DeleteVRouterNIC(ctx context.Context, id int) error {
	p := path(hc.VirtualRouterPath, "nic", id)

	return s.Delete(ctx, p, nil)
}

// DeleteVRouter deletes a Virtual Router with the given ID.
func (s VRouterService) DeleteVRouter(ctx context.Context, id int) error {
	p := path(hc.VirtualRouterPath, id)

	return s.Delete(ctx, p, nil)
}

// VRouter returns information about a vrouter.
func (s VRouterService) VRouter(ctx context.Context, id int) (*hc.ListVRouterResponse, error) {
	var resp hc.ListVRouterResponse

	p := path(hc.VirtualRouterPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// VRouters returns information about all vrouters.
func (s VRouterService) VRouters(ctx context.Context) (*hc.ListVRoutersResponse, error) {
	var resp hc.ListVRoutersResponse

	p := path(hc.VirtualRouterPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateVRouter updates a Virtual Router with the given ID.
func (s VRouterService) UpdateVRouter(ctx context.Context, id int, req hc.UpdateVRouterRequest) (*hc.UpdateVRouterResponse, error) {
	var resp hc.UpdateVRouterResponse

	p := path(hc.VirtualRouterPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstantiateVRouter instantiates a Virtual Router with the given ID.
func (s VRouterService) InstantiateVRouter(ctx context.Context, id int, req hc.InstantiateVRouterRequest) (*hc.InstantiateVRouterResponse, error) {
	var resp hc.InstantiateVRouterResponse

	p := path(hc.VirtualRouterPath, id, "instantiate")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockVRouter locks a Virtual Router with the given ID.
func (s VRouterService) LockVRouter(ctx context.Context, id int, req hc.LockVRouterRequest) (*hc.LockVRouterResponse, error) {
	var resp hc.LockVRouterResponse

	p := path(hc.VirtualRouterPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameVRouter renames a Virtual Router with the given ID.
func (s VRouterService) RenameVRouter(ctx context.Context, id int, req hc.RenameVRouterRequest) (*hc.RenameVRouterResponse, error) {
	var resp hc.RenameVRouterResponse

	p := path(hc.VirtualRouterPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeVRouterOwnership changes the ownership of a Virtual Router with the given ID.
func (s VRouterService) ChangeVRouterOwnership(ctx context.Context, id int, req hc.ChangeVRouterOwnershipRequest) (*hc.ChangeVRouterOwnershipResponse, error) {
	var resp hc.ChangeVRouterOwnershipResponse

	p := path(hc.VirtualRouterPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeVRouterPermissions changes the permissions of a Virtual Router with the given ID.
func (s VRouterService) ChangeVRouterPermissions(ctx context.Context, id int, req hc.ChangeVRouterPermissionsRequest) (*hc.ChangeVRouterPermissionsResponse, error) {
	var resp hc.ChangeVRouterPermissionsResponse

	p := path(hc.VirtualRouterPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockVRouter unlocks a Virtual Router with the given ID.
func (s VRouterService) UnlockVRouter(ctx context.Context, id int) (*hc.UnlockVRouterResponse, error) {
	var resp hc.UnlockVRouterResponse

	p := path(hc.VirtualRouterPath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateVRouter creates a new Virtual Router.
func (s VRouterService) CreateVRouter(ctx context.Context, req hc.CreateVRouterRequest) (*hc.CreateVRouterResponse, error) {
	var resp hc.CreateVRouterResponse

	p := path(hc.VirtualRouterPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateVRouterNIC creates a new NIC for the Virtual Router with the given ID.
func (s VRouterService) CreateVRouterNIC(ctx context.Context, id int, req hc.CreateVRouterNICRequest) (*hc.CreateVRouterNICResponse, error) {
	var resp hc.CreateVRouterNICResponse

	p := path(hc.VirtualRouterPath, id, "nic")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
