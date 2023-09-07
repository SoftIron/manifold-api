package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// MarketService owns the /market methods.
type MarketService struct {
	*client.Client
}

// DeleteMarket deletes the market with the given ID.
func (s MarketService) DeleteMarket(ctx context.Context, id int) error {
	p := path(cloud.MarketPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteMarketApp deletes the marketplace application with the given ID.
func (s MarketService) DeleteMarketApp(ctx context.Context, id, app int) error {
	p := path(cloud.MarketPath, id, "app", app)

	return s.Delete(ctx, p, nil)
}

// Market returns information about a marketplace.
func (s MarketService) Market(ctx context.Context, id int) (*cloud.ListMarketResponse, error) {
	var resp cloud.ListMarketResponse

	p := path(cloud.MarketPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Markets returns information about all marketplaces.
func (s MarketService) Markets(ctx context.Context) (*cloud.ListMarketsResponse, error) {
	var resp cloud.ListMarketsResponse

	p := path(cloud.MarketPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MarketApp returns information about a marketplace application.
func (s MarketService) MarketApp(ctx context.Context, id int) (*cloud.ListMarketAppResponse, error) {
	var resp cloud.ListMarketAppResponse

	p := path(cloud.MarketPath, "app", id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MarketApps returns information about all marketplace applications.
func (s MarketService) MarketApps(ctx context.Context) (*cloud.ListMarketAppsResponse, error) {
	var resp cloud.ListMarketAppsResponse

	p := path(cloud.MarketPath, "app")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateMarket updates the marketplace with the given ID.
func (s MarketService) UpdateMarket(ctx context.Context, id int, req cloud.UpdateMarketRequest) (*cloud.UpdateMarketResponse, error) {
	var resp cloud.UpdateMarketResponse

	p := path(cloud.MarketPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableMarket enables the marketplace with the given ID.
func (s MarketService) EnableMarket(ctx context.Context, id int, req cloud.EnableMarketRequest) (*cloud.EnableMarketResponse, error) {
	var resp cloud.EnableMarketResponse

	p := path(cloud.MarketPath, id, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameMarket renames the marketplace with the given ID.
func (s MarketService) RenameMarket(ctx context.Context, id int, req cloud.RenameMarketRequest) (*cloud.RenameMarketResponse, error) {
	var resp cloud.RenameMarketResponse

	p := path(cloud.MarketPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeMarketOwnership changes the ownership of the marketplace with the given ID.
func (s MarketService) ChangeMarketOwnership(ctx context.Context, id int, req cloud.ChangeMarketOwnershipRequest) (*cloud.ChangeMarketOwnershipResponse, error) {
	var resp cloud.ChangeMarketOwnershipResponse

	p := path(cloud.MarketPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeMarketPermissions changes the permissions of the marketplace with the given ID.
func (s MarketService) ChangeMarketPermissions(ctx context.Context, id int, req cloud.ChangeMarketPermissionsRequest) (*cloud.ChangeMarketPermissionsResponse, error) {
	var resp cloud.ChangeMarketPermissionsResponse

	p := path(cloud.MarketPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateMarketApp updates the marketplace application with the given ID.
func (s MarketService) UpdateMarketApp(ctx context.Context, id, app int, req cloud.UpdateMarketAppRequest) (*cloud.UpdateMarketAppResponse, error) {
	var resp cloud.UpdateMarketAppResponse

	p := path(cloud.MarketPath, id, "app", app)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableMarketApp enables the marketplace application with the given ID.
func (s MarketService) EnableMarketApp(ctx context.Context, id, app int, req cloud.EnableMarketAppRequest) (*cloud.EnableMarketAppResponse, error) {
	var resp cloud.EnableMarketAppResponse

	p := path(cloud.MarketPath, id, "app", app, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockMarketApp locks the marketplace application with the given ID.
func (s MarketService) LockMarketApp(ctx context.Context, id, app int, req cloud.LockMarketAppRequest) (*cloud.LockMarketAppResponse, error) {
	var resp cloud.LockMarketAppResponse

	p := path(cloud.MarketPath, id, "app", app, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameMarketApp renames the marketplace application with the given ID.
func (s MarketService) RenameMarketApp(ctx context.Context, id, app int, req cloud.RenameMarketAppRequest) (*cloud.RenameMarketAppResponse, error) {
	var resp cloud.RenameMarketAppResponse

	p := path(cloud.MarketPath, id, "app", app, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeMarketAppOwnership changes the ownership of the marketplace application with the given ID.
func (s MarketService) ChangeMarketAppOwnership(ctx context.Context, id, app int, req cloud.ChangeMarketAppOwnershipRequest) (*cloud.ChangeMarketAppOwnershipResponse, error) {
	var resp cloud.ChangeMarketAppOwnershipResponse

	p := path(cloud.MarketPath, id, "app", app, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeMarketAppPermissions changes the permissions of the marketplace application with the given ID.
func (s MarketService) ChangeMarketAppPermissions(ctx context.Context, id, app int, req cloud.ChangeMarketAppPermissionsRequest) (*cloud.ChangeMarketAppPermissionsResponse, error) {
	var resp cloud.ChangeMarketAppPermissionsResponse

	p := path(cloud.MarketPath, id, "app", app, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockMarketApp unlocks the marketplace application with the given ID.
func (s MarketService) UnlockMarketApp(ctx context.Context, id, app int) (*cloud.UnlockMarketAppResponse, error) {
	var resp cloud.UnlockMarketAppResponse

	p := path(cloud.MarketPath, id, "app", app, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateMarket creates a new marketplace.
func (s MarketService) CreateMarket(ctx context.Context, req cloud.CreateMarketRequest) (*cloud.CreateMarketResponse, error) {
	var resp cloud.CreateMarketResponse

	p := path(cloud.MarketPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateMarketApp creates a new marketplace application.
func (s MarketService) CreateMarketApp(ctx context.Context, id int, req cloud.CreateMarketAppRequest) (*cloud.CreateMarketAppResponse, error) {
	var resp cloud.CreateMarketAppResponse

	p := path(cloud.MarketPath, id, "app")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
