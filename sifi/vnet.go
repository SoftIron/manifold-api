package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// VNetService owns the /vnet methods.
type VNetService struct {
	*client.Client
}

// FreeVNetAddressRange free an address range from the virtual network with the given ID.
func (s VNetService) FreeVNetAddressRange(ctx context.Context, id, ar int) error {
	p := path(cloud.VirtualNetworkPath, id, "reservation", ar)

	return s.Delete(ctx, p, nil)
}

// DeleteVNet deletes the virtual network with the given ID.
func (s VNetService) DeleteVNet(ctx context.Context, id int) error {
	p := path(cloud.VirtualNetworkPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteVNetAddressRange deletes an address range from the virtual network with the given ID.
func (s VNetService) DeleteVNetAddressRange(ctx context.Context, id, ar int) error {
	p := path(cloud.VirtualNetworkPath, id, "address-range", ar)

	return s.Delete(ctx, p, nil)
}

// DeleteVNetTemplate deletes the virtual network template with the given ID.
func (s VNetService) DeleteVNetTemplate(ctx context.Context, id int) error {
	p := path(cloud.VirtualNetworkPath, "template", id)

	return s.Delete(ctx, p, nil)
}

// VNet returns information about a virtual network.
func (s VNetService) VNet(ctx context.Context, id int) (*cloud.ListVNetResponse, error) {
	var resp cloud.ListVNetResponse

	p := path(cloud.VirtualNetworkPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// VNets returns information about all virtual networks.
func (s VNetService) VNets(ctx context.Context) (*cloud.ListVNetsResponse, error) {
	var resp cloud.ListVNetsResponse

	p := path(cloud.VirtualNetworkPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// VNetTemplate returns information about a virtual networks.
func (s VNetService) VNetTemplate(ctx context.Context, id int) (*cloud.ListVNetTemplateResponse, error) {
	var resp cloud.ListVNetTemplateResponse

	p := path(cloud.VirtualNetworkPath, "template", id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// VNetTemplates returns information about all virtual networkss.
func (s VNetService) VNetTemplates(ctx context.Context) (*cloud.ListVNetTemplatesResponse, error) {
	var resp cloud.ListVNetTemplatesResponse

	p := path(cloud.VirtualNetworkPath, "template")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateVNet updates the virtual network with the given ID.
func (s VNetService) UpdateVNet(ctx context.Context, id int, req cloud.UpdateVNetRequest) (*cloud.UpdateVNetResponse, error) {
	var resp cloud.UpdateVNetResponse

	p := path(cloud.VirtualNetworkPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateVNetAddressRange updates an address range from the virtual network with the given ID.
func (s VNetService) UpdateVNetAddressRange(ctx context.Context, id int, req cloud.UpdateVNetAddressRangeRequest) (*cloud.UpdateVNetAddressRangeResponse, error) {
	var resp cloud.UpdateVNetAddressRangeResponse

	p := path(cloud.VirtualNetworkPath, id, "address-range")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HoldVNet holds the virtual network with the given ID.
func (s VNetService) HoldVNet(ctx context.Context, id int, req cloud.HoldVNetRequest) (*cloud.HoldVNetResponse, error) {
	var resp cloud.HoldVNetResponse

	p := path(cloud.VirtualNetworkPath, id, "hold")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockVNet locks the virtual network with the given ID.
func (s VNetService) LockVNet(ctx context.Context, id int, req cloud.LockVNetRequest) (*cloud.LockVNetResponse, error) {
	var resp cloud.LockVNetResponse

	p := path(cloud.VirtualNetworkPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameVNet renames the virtual network with the given ID.
func (s VNetService) RenameVNet(ctx context.Context, id int, req cloud.RenameVNetRequest) (*cloud.RenameVNetResponse, error) {
	var resp cloud.RenameVNetResponse

	p := path(cloud.VirtualNetworkPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeVNetOwnership changes the ownership of the virtual network with the given ID.
func (s VNetService) ChangeVNetOwnership(ctx context.Context, id int, req cloud.ChangeVNetOwnershipRequest) (*cloud.ChangeVNetOwnershipResponse, error) {
	var resp cloud.ChangeVNetOwnershipResponse

	p := path(cloud.VirtualNetworkPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeVNetPermissions changes the permissions of the virtual network with the given ID.
func (s VNetService) ChangeVNetPermissions(ctx context.Context, id int, req cloud.ChangeVNetPermissionsRequest) (*cloud.ChangeVNetPermissionsResponse, error) {
	var resp cloud.ChangeVNetPermissionsResponse

	p := path(cloud.VirtualNetworkPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RecoverVNets recovers the virtual network with the given ID.
func (s VNetService) RecoverVNets(ctx context.Context, id int, req cloud.RecoverVNetRequest) (*cloud.RecoverVNetResponse, error) {
	var resp cloud.RecoverVNetResponse

	p := path(cloud.VirtualNetworkPath, id, "recover")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ReleaseVNet releases the virtual network with the given ID.
func (s VNetService) ReleaseVNet(ctx context.Context, id int, req cloud.ReleaseVNetRequest) (*cloud.ReleaseVNetResponse, error) {
	var resp cloud.ReleaseVNetResponse

	p := path(cloud.VirtualNetworkPath, id, "release")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ReserveVNet reserves the virtual network with the given ID.
func (s VNetService) ReserveVNet(ctx context.Context, id int, req cloud.ReserveVNetRequest) (*cloud.ReserveVNetResponse, error) {
	var resp cloud.ReserveVNetResponse

	p := path(cloud.VirtualNetworkPath, id, "reserve")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateVNetTemplate updates the virtual network template with the given ID.
func (s VNetService) UpdateVNetTemplate(ctx context.Context, id int, req cloud.UpdateVNetTemplateRequest) (*cloud.UpdateVNetTemplateResponse, error) {
	var resp cloud.UpdateVNetTemplateResponse

	p := path(cloud.VirtualNetworkPath, id, "template")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockVNet unlocks the virtual network with the given ID.
func (s VNetService) UnlockVNet(ctx context.Context, id int) (*cloud.UnlockVNetResponse, error) {
	var resp cloud.UnlockVNetResponse

	p := path(cloud.VirtualNetworkPath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// InstantiateVNetTemplate instantiates the virtual network template with the given ID.
func (s VNetService) InstantiateVNetTemplate(ctx context.Context, id int, req cloud.InstantiateVNetTemplateRequest) (*cloud.InstantiateVNetTemplateResponse, error) {
	var resp cloud.InstantiateVNetTemplateResponse

	p := path(cloud.VirtualNetworkPath, "template", id, "instantiate")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// LockVNetTemplate locks the virtual network template with the given ID.
func (s VNetService) LockVNetTemplate(ctx context.Context, id int, req cloud.LockVNetTemplateRequest) (*cloud.LockVNetTemplateResponse, error) {
	var resp cloud.LockVNetTemplateResponse

	p := path(cloud.VirtualNetworkPath, "template", id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// RenameVNetTemplate renames the virtual network template with the given ID.
func (s VNetService) RenameVNetTemplate(ctx context.Context, id int, req cloud.RenameVNetTemplateRequest) (*cloud.RenameVNetTemplateResponse, error) {
	var resp cloud.RenameVNetTemplateResponse

	p := path(cloud.VirtualNetworkPath, "template", id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// ChangeVNetTemplateOwnership changes the ownership of the virtual network template with the given ID.
func (s VNetService) ChangeVNetTemplateOwnership(ctx context.Context, id int, req cloud.ChangeVNetTemplateOwnershipRequest) (*cloud.ChangeVNetTemplateOwnershipResponse, error) {
	var resp cloud.ChangeVNetTemplateOwnershipResponse

	p := path(cloud.VirtualNetworkPath, "template", id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// ChangeVNetTemplatePermissions changes the permissions of the virtual network template with the given ID.
func (s VNetService) ChangeVNetTemplatePermissions(ctx context.Context, id int, req cloud.ChangeVNetTemplatePermissionsRequest) (*cloud.ChangeVNetTemplatePermissionsResponse, error) {
	var resp cloud.ChangeVNetTemplatePermissionsResponse

	p := path(cloud.VirtualNetworkPath, "template", id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// UnlockVNetTemplate unlocks the virtual network template with the given ID.
func (s VNetService) UnlockVNetTemplate(ctx context.Context, id int) (*cloud.UnlockVNetTemplateResponse, error) {
	var resp cloud.UnlockVNetTemplateResponse

	p := path(cloud.VirtualNetworkPath, "template", id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// CreateVNet creates a new virtual network.
func (s VNetService) CreateVNet(ctx context.Context, req cloud.CreateVNetRequest) (*cloud.CreateVNetResponse, error) {
	var resp cloud.CreateVNetResponse

	p := path(cloud.VirtualNetworkPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddVNetAddressRange adds an address range to the virtual network with the given ID.
func (s VNetService) AddVNetAddressRange(ctx context.Context, id int, req cloud.AddVNetAddressRangeRequest) (*cloud.AddVNetAddressRangeResponse, error) {
	var resp cloud.AddVNetAddressRangeResponse

	p := path(cloud.VirtualNetworkPath, id, "address-range")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateVNetTemplate creates a new virtual network template.
func (s VNetService) CreateVNetTemplate(ctx context.Context, req cloud.CreateVNetTemplateRequest) (*cloud.CreateVNetTemplateResponse, error) {
	var resp cloud.CreateVNetTemplateResponse

	p := path(cloud.VirtualNetworkPath, "template")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneVNetTemplate clones the virtual network template with the given ID.
func (s VNetService) CloneVNetTemplate(ctx context.Context, id int, req cloud.CloneVNetTemplateRequest) (*cloud.CloneVNetTemplateResponse, error) {
	var resp cloud.CloneVNetTemplateResponse

	p := path(cloud.VirtualNetworkPath, "template", id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
