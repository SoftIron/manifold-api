package cloud

import (
	"context"
)

// MarketService owns the /cloud/market methods.
type MarketService struct {
	*service
}

// DeleteMarket deletes the market with the given ID.
func (s MarketService) DeleteMarket(ctx context.Context, id int) error {
	p := s.path(MarketPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteMarketApp deletes the marketplace application with the given ID.
func (s MarketService) DeleteMarketApp(ctx context.Context, id int) error {
	p := s.path(MarketAppPath, id)

	return s.Delete(ctx, p, nil)
}

// Market returns information about a marketplace.
func (s MarketService) Market(ctx context.Context, id int) (*MarketResponse, error) {
	var resp MarketResponse

	p := s.path(MarketPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Markets returns information about all marketplaces.
func (s MarketService) Markets(ctx context.Context) (*MarketsResponse, error) {
	var resp MarketsResponse

	p := s.path(MarketPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MarketApp returns information about a marketplace application.
func (s MarketService) MarketApp(ctx context.Context, id int) (*MarketAppResponse, error) {
	var resp MarketAppResponse

	p := s.path(MarketAppPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MarketApps returns information about all marketplace applications.
func (s MarketService) MarketApps(ctx context.Context) (*MarketAppsResponse, error) {
	var resp MarketAppsResponse

	p := s.path(MarketAppPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateMarket updates the marketplace with the given ID.
func (s MarketService) UpdateMarket(ctx context.Context, id int, req UpdateMarketRequest) (*UpdateMarketResponse, error) {
	var resp UpdateMarketResponse

	p := s.path(MarketPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableMarket enables the marketplace with the given ID.
func (s MarketService) EnableMarket(ctx context.Context, id int, req EnableMarketRequest) (*EnableMarketResponse, error) {
	var resp EnableMarketResponse

	p := s.path(MarketPath, id, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameMarket renames the marketplace with the given ID.
func (s MarketService) RenameMarket(ctx context.Context, id int, req RenameMarketRequest) (*RenameMarketResponse, error) {
	var resp RenameMarketResponse

	p := s.path(MarketPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeMarketOwnership changes the ownership of the marketplace with the given ID.
func (s MarketService) ChangeMarketOwnership(ctx context.Context, id int, req ChangeMarketOwnershipRequest) (*ChangeMarketOwnershipResponse, error) {
	var resp ChangeMarketOwnershipResponse

	p := s.path(MarketPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeMarketPermissions changes the permissions of the marketplace with the given ID.
func (s MarketService) ChangeMarketPermissions(ctx context.Context, id int, req ChangeMarketPermissionsRequest) (*ChangeMarketPermissionsResponse, error) {
	var resp ChangeMarketPermissionsResponse

	p := s.path(MarketPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateMarketApp updates the marketplace application with the given ID.
func (s MarketService) UpdateMarketApp(ctx context.Context, id int, req UpdateMarketAppRequest) (*UpdateMarketAppResponse, error) {
	var resp UpdateMarketAppResponse

	p := s.path(MarketAppPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableMarketApp enables the marketplace application with the given ID.
func (s MarketService) EnableMarketApp(ctx context.Context, id int, req EnableMarketAppRequest) (*EnableMarketAppResponse, error) {
	var resp EnableMarketAppResponse

	p := s.path(MarketAppPath, id, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockMarketApp locks the marketplace application with the given ID.
func (s MarketService) LockMarketApp(ctx context.Context, id int, req LockMarketAppRequest) (*LockMarketAppResponse, error) {
	var resp LockMarketAppResponse

	p := s.path(MarketAppPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameMarketApp renames the marketplace application with the given ID.
func (s MarketService) RenameMarketApp(ctx context.Context, id int, req RenameMarketAppRequest) (*RenameMarketAppResponse, error) {
	var resp RenameMarketAppResponse

	p := s.path(MarketAppPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeMarketAppOwnership changes the ownership of the marketplace application with the given ID.
func (s MarketService) ChangeMarketAppOwnership(ctx context.Context, id int, req ChangeMarketAppOwnershipRequest) (*ChangeMarketAppOwnershipResponse, error) {
	var resp ChangeMarketAppOwnershipResponse

	p := s.path(MarketAppPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeMarketAppPermissions changes the permissions of the marketplace application with the given ID.
func (s MarketService) ChangeMarketAppPermissions(ctx context.Context, id int, req ChangeMarketAppPermissionsRequest) (*ChangeMarketAppPermissionsResponse, error) {
	var resp ChangeMarketAppPermissionsResponse

	p := s.path(MarketAppPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockMarketApp unlocks the marketplace application with the given ID.
func (s MarketService) UnlockMarketApp(ctx context.Context, id int) (*UnlockMarketAppResponse, error) {
	var resp UnlockMarketAppResponse

	p := s.path(MarketAppPath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateMarket creates a new marketplace.
func (s MarketService) CreateMarket(ctx context.Context, req CreateMarketRequest) (*CreateMarketResponse, error) {
	var resp CreateMarketResponse

	p := s.path(MarketPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateMarketApp creates a new marketplace application.
func (s MarketService) CreateMarketApp(ctx context.Context, id int, req CreateMarketAppRequest) (*CreateMarketAppResponse, error) {
	var resp CreateMarketAppResponse

	p := s.path(MarketPath, id, "app")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
