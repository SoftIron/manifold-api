package cloud

import (
	"context"
)

// NetworkService owns the /cloud/network methods.
type NetworkService struct {
	*service
}

// FreeNetworkAddressRange free an address range from the network with the given ID.
func (s NetworkService) FreeNetworkAddressRange(ctx context.Context, id, ar int) error {
	p := s.path(NetworkPath, id, "reservation", ar)

	return s.Delete(ctx, p, nil)
}

// DeleteNetwork deletes the network with the given ID.
func (s NetworkService) DeleteNetwork(ctx context.Context, id int) error {
	p := s.path(NetworkPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteNetworkAddressRange deletes an address range from the network with the given ID.
func (s NetworkService) DeleteNetworkAddressRange(ctx context.Context, id, ar int) error {
	p := s.path(NetworkPath, id, "address-range", ar)

	return s.Delete(ctx, p, nil)
}

// DeleteNetworkTemplate deletes the network template with the given ID.
func (s NetworkService) DeleteNetworkTemplate(ctx context.Context, id int) error {
	p := s.path(NetworkPath, "template", id)

	return s.Delete(ctx, p, nil)
}

// Network returns information about a network.
func (s NetworkService) Network(ctx context.Context, id int) (*NetworkResponse, error) {
	var resp NetworkResponse

	p := s.path(NetworkPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Networks returns information about all networks.
func (s NetworkService) Networks(ctx context.Context) (*NetworksResponse, error) {
	var resp NetworksResponse

	p := s.path(NetworkPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// NetworkTemplate returns information about a networks.
func (s NetworkService) NetworkTemplate(ctx context.Context, id int) (*NetworkTemplateResponse, error) {
	var resp NetworkTemplateResponse

	p := s.path(NetworkPath, "template", id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// NetworkTemplates returns information about all networkss.
func (s NetworkService) NetworkTemplates(ctx context.Context) (*NetworkTemplatesResponse, error) {
	var resp NetworkTemplatesResponse

	p := s.path(NetworkPath, "template")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateNetwork updates the network with the given ID.
func (s NetworkService) UpdateNetwork(ctx context.Context, id int, req UpdateNetworkRequest) (*UpdateNetworkResponse, error) {
	var resp UpdateNetworkResponse

	p := s.path(NetworkPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateNetworkAddressRange updates an address range from the network with the given ID.
func (s NetworkService) UpdateNetworkAddressRange(ctx context.Context, id int, req UpdateNetworkAddressRangeRequest) (*UpdateNetworkAddressRangeResponse, error) {
	var resp UpdateNetworkAddressRangeResponse

	p := s.path(NetworkPath, id, "address-range")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HoldNetwork holds the network with the given ID.
func (s NetworkService) HoldNetwork(ctx context.Context, id int, req HoldNetworkRequest) (*HoldNetworkResponse, error) {
	var resp HoldNetworkResponse

	p := s.path(NetworkPath, id, "hold")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockNetwork locks the network with the given ID.
func (s NetworkService) LockNetwork(ctx context.Context, id int, req LockNetworkRequest) (*LockNetworkResponse, error) {
	var resp LockNetworkResponse

	p := s.path(NetworkPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameNetwork renames the network with the given ID.
func (s NetworkService) RenameNetwork(ctx context.Context, id int, req RenameNetworkRequest) (*RenameNetworkResponse, error) {
	var resp RenameNetworkResponse

	p := s.path(NetworkPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeNetworkOwnership changes the ownership of the network with the given ID.
func (s NetworkService) ChangeNetworkOwnership(ctx context.Context, id int, req ChangeNetworkOwnershipRequest) (*ChangeNetworkOwnershipResponse, error) {
	var resp ChangeNetworkOwnershipResponse

	p := s.path(NetworkPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeNetworkPermissions changes the permissions of the network with the given ID.
func (s NetworkService) ChangeNetworkPermissions(ctx context.Context, id int, req ChangeNetworkPermissionsRequest) (*ChangeNetworkPermissionsResponse, error) {
	var resp ChangeNetworkPermissionsResponse

	p := s.path(NetworkPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RecoverNetworks recovers the network with the given ID.
func (s NetworkService) RecoverNetworks(ctx context.Context, id int, req RecoverNetworkRequest) (*RecoverNetworkResponse, error) {
	var resp RecoverNetworkResponse

	p := s.path(NetworkPath, id, "recover")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ReleaseNetwork releases the network with the given ID.
func (s NetworkService) ReleaseNetwork(ctx context.Context, id int, req ReleaseNetworkRequest) (*ReleaseNetworkResponse, error) {
	var resp ReleaseNetworkResponse

	p := s.path(NetworkPath, id, "release")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ReserveNetwork reserves the network with the given ID.
func (s NetworkService) ReserveNetwork(ctx context.Context, id int, req ReserveNetworkRequest) (*ReserveNetworkResponse, error) {
	var resp ReserveNetworkResponse

	p := s.path(NetworkPath, id, "reserve")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateNetworkTemplate updates the network template with the given ID.
func (s NetworkService) UpdateNetworkTemplate(ctx context.Context, id int, req UpdateNetworkTemplateRequest) (*UpdateNetworkTemplateResponse, error) {
	var resp UpdateNetworkTemplateResponse

	p := s.path(NetworkPath, id, "template")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockNetwork unlocks the network with the given ID.
func (s NetworkService) UnlockNetwork(ctx context.Context, id int) (*UnlockNetworkResponse, error) {
	var resp UnlockNetworkResponse

	p := s.path(NetworkPath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// InstantiateNetworkTemplate instantiates the network template with the given ID.
func (s NetworkService) InstantiateNetworkTemplate(ctx context.Context, id int, req InstantiateNetworkTemplateRequest) (*InstantiateNetworkTemplateResponse, error) {
	var resp InstantiateNetworkTemplateResponse

	p := s.path(NetworkPath, "template", id, "instantiate")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// LockNetworkTemplate locks the network template with the given ID.
func (s NetworkService) LockNetworkTemplate(ctx context.Context, id int, req LockNetworkTemplateRequest) (*LockNetworkTemplateResponse, error) {
	var resp LockNetworkTemplateResponse

	p := s.path(NetworkPath, "template", id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// RenameNetworkTemplate renames the network template with the given ID.
func (s NetworkService) RenameNetworkTemplate(ctx context.Context, id int, req RenameNetworkTemplateRequest) (*RenameNetworkTemplateResponse, error) {
	var resp RenameNetworkTemplateResponse

	p := s.path(NetworkPath, "template", id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// ChangeNetworkTemplateOwnership changes the ownership of the network template with the given ID.
func (s NetworkService) ChangeNetworkTemplateOwnership(ctx context.Context, id int, req ChangeNetworkTemplateOwnershipRequest) (*ChangeNetworkTemplateOwnershipResponse, error) {
	var resp ChangeNetworkTemplateOwnershipResponse

	p := s.path(NetworkPath, "template", id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// ChangeNetworkTemplatePermissions changes the permissions of the network template with the given ID.
func (s NetworkService) ChangeNetworkTemplatePermissions(ctx context.Context, id int, req ChangeNetworkTemplatePermissionsRequest) (*ChangeNetworkTemplatePermissionsResponse, error) {
	var resp ChangeNetworkTemplatePermissionsResponse

	p := s.path(NetworkPath, "template", id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// UnlockNetworkTemplate unlocks the network template with the given ID.
func (s NetworkService) UnlockNetworkTemplate(ctx context.Context, id int) (*UnlockNetworkTemplateResponse, error) {
	var resp UnlockNetworkTemplateResponse

	p := s.path(NetworkPath, "template", id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// CreateNetwork creates a new network.
func (s NetworkService) CreateNetwork(ctx context.Context, req CreateNetworkRequest) (*CreateNetworkResponse, error) {
	var resp CreateNetworkResponse

	p := s.path(NetworkPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddNetworkAddressRange adds an address range to the network with the given ID.
func (s NetworkService) AddNetworkAddressRange(ctx context.Context, id int, req AddNetworkAddressRangeRequest) (*AddNetworkAddressRangeResponse, error) {
	var resp AddNetworkAddressRangeResponse

	p := s.path(NetworkPath, id, "address-range")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateNetworkTemplate creates a new network template.
func (s NetworkService) CreateNetworkTemplate(ctx context.Context, req CreateNetworkTemplateRequest) (*CreateNetworkTemplateResponse, error) {
	var resp CreateNetworkTemplateResponse

	p := s.path(NetworkPath, "template")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneNetworkTemplate clones the network template with the given ID.
func (s NetworkService) CloneNetworkTemplate(ctx context.Context, id int, req CloneNetworkTemplateRequest) (*CloneNetworkTemplateResponse, error) {
	var resp CloneNetworkTemplateResponse

	p := s.path(NetworkPath, "template", id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
