package sifi

import (
	"context"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/deprecated/v1/hc"
)

// MarketService owns the /market methods.
type MarketService struct {
	*client.Client
}

// DeleteMarket deletes the market with the given ID.
func (s MarketService) DeleteMarket(ctx context.Context, id int) error {
	p := path(hc.MarketPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteMarketApp deletes the marketplace application with the given ID.
func (s MarketService) DeleteMarketApp(ctx context.Context, id, app int) error {
	p := path(hc.MarketPath, id, "app", app)

	return s.Delete(ctx, p, nil)
}

// Market returns information about a marketplace.
func (s MarketService) Market(ctx context.Context, id int) (*hc.ListMarketResponse, error) {
	var resp hc.ListMarketResponse

	p := path(hc.MarketPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Markets returns information about all marketplaces.
func (s MarketService) Markets(ctx context.Context) (*hc.ListMarketsResponse, error) {
	var resp hc.ListMarketsResponse

	p := path(hc.MarketPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MarketApp returns information about a marketplace application.
func (s MarketService) MarketApp(ctx context.Context, id int) (*hc.ListMarketAppResponse, error) {
	var resp hc.ListMarketAppResponse

	p := path(hc.MarketPath, "app", id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MarketApps returns information about all marketplace applications.
func (s MarketService) MarketApps(ctx context.Context) (*hc.ListMarketAppsResponse, error) {
	var resp hc.ListMarketAppsResponse

	p := path(hc.MarketPath, "app")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateMarket updates the marketplace with the given ID.
func (s MarketService) UpdateMarket(ctx context.Context, id int, req hc.UpdateMarketRequest) (*hc.UpdateMarketResponse, error) {
	var resp hc.UpdateMarketResponse

	p := path(hc.MarketPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableMarket enables the marketplace with the given ID.
func (s MarketService) EnableMarket(ctx context.Context, id int, req hc.EnableMarketRequest) (*hc.EnableMarketResponse, error) {
	var resp hc.EnableMarketResponse

	p := path(hc.MarketPath, id, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameMarket renames the marketplace with the given ID.
func (s MarketService) RenameMarket(ctx context.Context, id int, req hc.RenameMarketRequest) (*hc.RenameMarketResponse, error) {
	var resp hc.RenameMarketResponse

	p := path(hc.MarketPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeMarketOwnership changes the ownership of the marketplace with the given ID.
func (s MarketService) ChangeMarketOwnership(ctx context.Context, id int, req hc.ChangeMarketOwnershipRequest) (*hc.ChangeMarketOwnershipResponse, error) {
	var resp hc.ChangeMarketOwnershipResponse

	p := path(hc.MarketPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeMarketPermissions changes the permissions of the marketplace with the given ID.
func (s MarketService) ChangeMarketPermissions(ctx context.Context, id int, req hc.ChangeMarketPermissionsRequest) (*hc.ChangeMarketPermissionsResponse, error) {
	var resp hc.ChangeMarketPermissionsResponse

	p := path(hc.MarketPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateMarketApp updates the marketplace application with the given ID.
func (s MarketService) UpdateMarketApp(ctx context.Context, id, app int, req hc.UpdateMarketAppRequest) (*hc.UpdateMarketAppResponse, error) {
	var resp hc.UpdateMarketAppResponse

	p := path(hc.MarketPath, id, "app", app)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableMarketApp enables the marketplace application with the given ID.
func (s MarketService) EnableMarketApp(ctx context.Context, id, app int, req hc.EnableMarketAppRequest) (*hc.EnableMarketAppResponse, error) {
	var resp hc.EnableMarketAppResponse

	p := path(hc.MarketPath, id, "app", app, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockMarketApp locks the marketplace application with the given ID.
func (s MarketService) LockMarketApp(ctx context.Context, id, app int, req hc.LockMarketAppRequest) (*hc.LockMarketAppResponse, error) {
	var resp hc.LockMarketAppResponse

	p := path(hc.MarketPath, id, "app", app, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameMarketApp renames the marketplace application with the given ID.
func (s MarketService) RenameMarketApp(ctx context.Context, id, app int, req hc.RenameMarketAppRequest) (*hc.RenameMarketAppResponse, error) {
	var resp hc.RenameMarketAppResponse

	p := path(hc.MarketPath, id, "app", app, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeMarketAppOwnership changes the ownership of the marketplace application with the given ID.
func (s MarketService) ChangeMarketAppOwnership(ctx context.Context, id, app int, req hc.ChangeMarketAppOwnershipRequest) (*hc.ChangeMarketAppOwnershipResponse, error) {
	var resp hc.ChangeMarketAppOwnershipResponse

	p := path(hc.MarketPath, id, "app", app, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeMarketAppPermissions changes the permissions of the marketplace application with the given ID.
func (s MarketService) ChangeMarketAppPermissions(ctx context.Context, id, app int, req hc.ChangeMarketAppPermissionsRequest) (*hc.ChangeMarketAppPermissionsResponse, error) {
	var resp hc.ChangeMarketAppPermissionsResponse

	p := path(hc.MarketPath, id, "app", app, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockMarketApp unlocks the marketplace application with the given ID.
func (s MarketService) UnlockMarketApp(ctx context.Context, id, app int) (*hc.UnlockMarketAppResponse, error) {
	var resp hc.UnlockMarketAppResponse

	p := path(hc.MarketPath, id, "app", app, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateMarket creates a new marketplace.
func (s MarketService) CreateMarket(ctx context.Context, req hc.CreateMarketRequest) (*hc.CreateMarketResponse, error) {
	var resp hc.CreateMarketResponse

	p := path(hc.MarketPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateMarketApp creates a new marketplace application.
func (s MarketService) CreateMarketApp(ctx context.Context, id int, req hc.CreateMarketAppRequest) (*hc.CreateMarketAppResponse, error) {
	var resp hc.CreateMarketAppResponse

	p := path(hc.MarketPath, id, "app")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
