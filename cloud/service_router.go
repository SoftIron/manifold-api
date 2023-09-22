package cloud

import (
	"context"
)

// RouterService own the /cloud/router methods.
type RouterService struct {
	*service
}

// DeleteRouterNIC deletes a NIC from the Router with the given ID.
func (s RouterService) DeleteRouterNIC(ctx context.Context, id int) error {
	p := s.path(RouterPath, "nic", id)

	return s.Delete(ctx, p, nil)
}

// DeleteRouter deletes a Router with the given ID.
func (s RouterService) DeleteRouter(ctx context.Context, id int) error {
	p := s.path(RouterPath, id)

	return s.Delete(ctx, p, nil)
}

// Router returns information about a router.
func (s RouterService) Router(ctx context.Context, id int) (*RouterResponse, error) {
	var resp RouterResponse

	p := s.path(RouterPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Routers returns information about all routers.
func (s RouterService) Routers(ctx context.Context) (*RoutersResponse, error) {
	var resp RoutersResponse

	p := s.path(RouterPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateRouter updates a Router with the given ID.
func (s RouterService) UpdateRouter(ctx context.Context, id int, req UpdateRouterRequest) (*UpdateRouterResponse, error) {
	var resp UpdateRouterResponse

	p := s.path(RouterPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstantiateRouter instantiates a Router with the given ID.
func (s RouterService) InstantiateRouter(ctx context.Context, id int, req InstantiateRouterRequest) (*InstantiateRouterResponse, error) {
	var resp InstantiateRouterResponse

	p := s.path(RouterPath, id, "instantiate")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockRouter locks a Router with the given ID.
func (s RouterService) LockRouter(ctx context.Context, id int, req LockRouterRequest) (*LockRouterResponse, error) {
	var resp LockRouterResponse

	p := s.path(RouterPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameRouter renames a Router with the given ID.
func (s RouterService) RenameRouter(ctx context.Context, id int, req RenameRouterRequest) (*RenameRouterResponse, error) {
	var resp RenameRouterResponse

	p := s.path(RouterPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeRouterOwnership changes the ownership of a Router with the given ID.
func (s RouterService) ChangeRouterOwnership(ctx context.Context, id int, req ChangeRouterOwnershipRequest) (*ChangeRouterOwnershipResponse, error) {
	var resp ChangeRouterOwnershipResponse

	p := s.path(RouterPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeRouterPermissions changes the permissions of a Router with the given ID.
func (s RouterService) ChangeRouterPermissions(ctx context.Context, id int, req ChangeRouterPermissionsRequest) (*ChangeRouterPermissionsResponse, error) {
	var resp ChangeRouterPermissionsResponse

	p := s.path(RouterPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockRouter unlocks a Router with the given ID.
func (s RouterService) UnlockRouter(ctx context.Context, id int) (*UnlockRouterResponse, error) {
	var resp UnlockRouterResponse

	p := s.path(RouterPath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateRouter creates a new Router.
func (s RouterService) CreateRouter(ctx context.Context, req CreateRouterRequest) (*CreateRouterResponse, error) {
	var resp CreateRouterResponse

	p := s.path(RouterPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateRouterNIC creates a new NIC for the Router with the given ID.
func (s RouterService) CreateRouterNIC(ctx context.Context, id int, req CreateRouterNICRequest) (*CreateRouterNICResponse, error) {
	var resp CreateRouterNICResponse

	p := s.path(RouterPath, id, "nic")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
