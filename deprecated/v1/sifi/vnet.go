package sifi

import (
	"context"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/deprecated/v1/hc"
)

// VNetService owns the /vnet methods.
type VNetService struct {
	*client.Client
}

// FreeVNetAddressRange free an address range from the virtual network with the given ID.
func (s VNetService) FreeVNetAddressRange(ctx context.Context, id, ar int) error {
	p := path(hc.VirtualNetworkPath, id, "reservation", ar)

	return s.Delete(ctx, p, nil)
}

// DeleteVNet deletes the virtual network with the given ID.
func (s VNetService) DeleteVNet(ctx context.Context, id int) error {
	p := path(hc.VirtualNetworkPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteVNetAddressRange deletes an address range from the virtual network with the given ID.
func (s VNetService) DeleteVNetAddressRange(ctx context.Context, id, ar int) error {
	p := path(hc.VirtualNetworkPath, id, "address-range", ar)

	return s.Delete(ctx, p, nil)
}

// DeleteVNetTemplate deletes the virtual network template with the given ID.
func (s VNetService) DeleteVNetTemplate(ctx context.Context, id int) error {
	p := path(hc.VirtualNetworkPath, "template", id)

	return s.Delete(ctx, p, nil)
}

// VNet returns information about a virtual network.
func (s VNetService) VNet(ctx context.Context, id int) (*hc.ListVNetResponse, error) {
	var resp hc.ListVNetResponse

	p := path(hc.VirtualNetworkPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// VNets returns information about all virtual networks.
func (s VNetService) VNets(ctx context.Context) (*hc.ListVNetsResponse, error) {
	var resp hc.ListVNetsResponse

	p := path(hc.VirtualNetworkPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// VNetTemplate returns information about a virtual networks.
func (s VNetService) VNetTemplate(ctx context.Context, id int) (*hc.ListVNetTemplateResponse, error) {
	var resp hc.ListVNetTemplateResponse

	p := path(hc.VirtualNetworkPath, "template", id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// VNetTemplates returns information about all virtual networkss.
func (s VNetService) VNetTemplates(ctx context.Context) (*hc.ListVNetTemplatesResponse, error) {
	var resp hc.ListVNetTemplatesResponse

	p := path(hc.VirtualNetworkPath, "template")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateVNet updates the virtual network with the given ID.
func (s VNetService) UpdateVNet(ctx context.Context, id int, req hc.UpdateVNetRequest) (*hc.UpdateVNetResponse, error) {
	var resp hc.UpdateVNetResponse

	p := path(hc.VirtualNetworkPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateVNetAddressRange updates an address range from the virtual network with the given ID.
func (s VNetService) UpdateVNetAddressRange(ctx context.Context, id int, req hc.UpdateVNetAddressRangeRequest) (*hc.UpdateVNetAddressRangeResponse, error) {
	var resp hc.UpdateVNetAddressRangeResponse

	p := path(hc.VirtualNetworkPath, id, "address-range")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HoldVNet holds the virtual network with the given ID.
func (s VNetService) HoldVNet(ctx context.Context, id int, req hc.HoldVNetRequest) (*hc.HoldVNetResponse, error) {
	var resp hc.HoldVNetResponse

	p := path(hc.VirtualNetworkPath, id, "hold")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockVNet locks the virtual network with the given ID.
func (s VNetService) LockVNet(ctx context.Context, id int, req hc.LockVNetRequest) (*hc.LockVNetResponse, error) {
	var resp hc.LockVNetResponse

	p := path(hc.VirtualNetworkPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameVNet renames the virtual network with the given ID.
func (s VNetService) RenameVNet(ctx context.Context, id int, req hc.RenameVNetRequest) (*hc.RenameVNetResponse, error) {
	var resp hc.RenameVNetResponse

	p := path(hc.VirtualNetworkPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeVNetOwnership changes the ownership of the virtual network with the given ID.
func (s VNetService) ChangeVNetOwnership(ctx context.Context, id int, req hc.ChangeVNetOwnershipRequest) (*hc.ChangeVNetOwnershipResponse, error) {
	var resp hc.ChangeVNetOwnershipResponse

	p := path(hc.VirtualNetworkPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeVNetPermissions changes the permissions of the virtual network with the given ID.
func (s VNetService) ChangeVNetPermissions(ctx context.Context, id int, req hc.ChangeVNetPermissionsRequest) (*hc.ChangeVNetPermissionsResponse, error) {
	var resp hc.ChangeVNetPermissionsResponse

	p := path(hc.VirtualNetworkPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RecoverVNets recovers the virtual network with the given ID.
func (s VNetService) RecoverVNets(ctx context.Context, id int, req hc.RecoverVNetRequest) (*hc.RecoverVNetResponse, error) {
	var resp hc.RecoverVNetResponse

	p := path(hc.VirtualNetworkPath, id, "recover")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ReleaseVNet releases the virtual network with the given ID.
func (s VNetService) ReleaseVNet(ctx context.Context, id int, req hc.ReleaseVNetRequest) (*hc.ReleaseVNetResponse, error) {
	var resp hc.ReleaseVNetResponse

	p := path(hc.VirtualNetworkPath, id, "release")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ReserveVNet reserves the virtual network with the given ID.
func (s VNetService) ReserveVNet(ctx context.Context, id int, req hc.ReserveVNetRequest) (*hc.ReserveVNetResponse, error) {
	var resp hc.ReserveVNetResponse

	p := path(hc.VirtualNetworkPath, id, "reserve")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateVNetTemplate updates the virtual network template with the given ID.
func (s VNetService) UpdateVNetTemplate(ctx context.Context, id int, req hc.UpdateVNetTemplateRequest) (*hc.UpdateVNetTemplateResponse, error) {
	var resp hc.UpdateVNetTemplateResponse

	p := path(hc.VirtualNetworkPath, id, "template")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockVNet unlocks the virtual network with the given ID.
func (s VNetService) UnlockVNet(ctx context.Context, id int) (*hc.UnlockVNetResponse, error) {
	var resp hc.UnlockVNetResponse

	p := path(hc.VirtualNetworkPath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// InstantiateVNetTemplate instantiates the virtual network template with the given ID.
func (s VNetService) InstantiateVNetTemplate(ctx context.Context, id int, req hc.InstantiateVNetTemplateRequest) (*hc.InstantiateVNetTemplateResponse, error) {
	var resp hc.InstantiateVNetTemplateResponse

	p := path(hc.VirtualNetworkPath, "template", id, "instantiate")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// LockVNetTemplate locks the virtual network template with the given ID.
func (s VNetService) LockVNetTemplate(ctx context.Context, id int, req hc.LockVNetTemplateRequest) (*hc.LockVNetTemplateResponse, error) {
	var resp hc.LockVNetTemplateResponse

	p := path(hc.VirtualNetworkPath, "template", id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// RenameVNetTemplate renames the virtual network template with the given ID.
func (s VNetService) RenameVNetTemplate(ctx context.Context, id int, req hc.RenameVNetTemplateRequest) (*hc.RenameVNetTemplateResponse, error) {
	var resp hc.RenameVNetTemplateResponse

	p := path(hc.VirtualNetworkPath, "template", id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// ChangeVNetTemplateOwnership changes the ownership of the virtual network template with the given ID.
func (s VNetService) ChangeVNetTemplateOwnership(ctx context.Context, id int, req hc.ChangeVNetTemplateOwnershipRequest) (*hc.ChangeVNetTemplateOwnershipResponse, error) {
	var resp hc.ChangeVNetTemplateOwnershipResponse

	p := path(hc.VirtualNetworkPath, "template", id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// ChangeVNetTemplatePermissions changes the permissions of the virtual network template with the given ID.
func (s VNetService) ChangeVNetTemplatePermissions(ctx context.Context, id int, req hc.ChangeVNetTemplatePermissionsRequest) (*hc.ChangeVNetTemplatePermissionsResponse, error) {
	var resp hc.ChangeVNetTemplatePermissionsResponse

	p := path(hc.VirtualNetworkPath, "template", id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// UnlockVNetTemplate unlocks the virtual network template with the given ID.
func (s VNetService) UnlockVNetTemplate(ctx context.Context, id int) (*hc.UnlockVNetTemplateResponse, error) {
	var resp hc.UnlockVNetTemplateResponse

	p := path(hc.VirtualNetworkPath, "template", id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil

}

// CreateVNet creates a new virtual network.
func (s VNetService) CreateVNet(ctx context.Context, req hc.CreateVNetRequest) (*hc.CreateVNetResponse, error) {
	var resp hc.CreateVNetResponse

	p := path(hc.VirtualNetworkPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddVNetAddressRange adds an address range to the virtual network with the given ID.
func (s VNetService) AddVNetAddressRange(ctx context.Context, id int, req hc.AddVNetAddressRangeRequest) (*hc.AddVNetAddressRangeResponse, error) {
	var resp hc.AddVNetAddressRangeResponse

	p := path(hc.VirtualNetworkPath, id, "address-range")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateVNetTemplate creates a new virtual network template.
func (s VNetService) CreateVNetTemplate(ctx context.Context, req hc.CreateVNetTemplateRequest) (*hc.CreateVNetTemplateResponse, error) {
	var resp hc.CreateVNetTemplateResponse

	p := path(hc.VirtualNetworkPath, "template")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneVNetTemplate clones the virtual network template with the given ID.
func (s VNetService) CloneVNetTemplate(ctx context.Context, id int, req hc.CloneVNetTemplateRequest) (*hc.CloneVNetTemplateResponse, error) {
	var resp hc.CloneVNetTemplateResponse

	p := path(hc.VirtualNetworkPath, "template", id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
