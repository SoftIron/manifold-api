package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/deprecated/v1/hc"
)

// DeleteInstanceGroup deletes the instance group with the given ID.
func (s InstanceService) DeleteInstanceGroup(ctx context.Context, id int) error {
	p := path(hc.InstanceGroupPath, id)

	return s.Delete(ctx, p, nil)
}

// InstanceGroup returns information about an instance group.
func (s InstanceService) InstanceGroup(ctx context.Context, id int) (*hc.ListInstanceGroupResponse, error) {
	var resp hc.ListInstanceGroupResponse

	p := path(hc.InstanceGroupPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstanceGroups returns information about all instance groups.
func (s InstanceService) InstanceGroups(ctx context.Context) (*hc.ListInstanceGroupsResponse, error) {
	var resp hc.ListInstanceGroupsResponse

	p := path(hc.InstanceGroupPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateInstanceGroup updates the instance group with the given ID.
func (s InstanceService) UpdateInstanceGroup(ctx context.Context, id int, req *hc.UpdateInstanceGroupRequest) (*hc.UpdateInstanceGroupResponse, error) {
	var resp hc.UpdateInstanceGroupResponse

	p := path(hc.InstanceGroupPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockInstanceGroup locks the instance group with the given ID.
func (s InstanceService) LockInstanceGroup(ctx context.Context, id int, req hc.LockInstanceGroupRequest) (*hc.LockInstanceGroupResponse, error) {
	var resp hc.LockInstanceGroupResponse

	p := path(hc.InstanceGroupPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameInstanceGroup renames the instance group with the given ID.
func (s InstanceService) RenameInstanceGroup(ctx context.Context, id int, req hc.RenameInstanceGroupRequest) (*hc.RenameInstanceGroupResponse, error) {
	var resp hc.RenameInstanceGroupResponse

	p := path(hc.InstanceGroupPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeInstanceGroupOwnership changes the ownership of the instance group with the given ID.
func (s InstanceService) ChangeInstanceGroupOwnership(ctx context.Context, id int, req hc.ChangeInstanceGroupOwnershipRequest) (*hc.ChangeInstanceGroupOwnershipResponse, error) {
	var resp hc.ChangeInstanceGroupOwnershipResponse

	p := path(hc.InstanceGroupPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeInstanceGroupPermissions changes the permissions of the instance group with the given ID.
func (s InstanceService) ChangeInstanceGroupPermissions(ctx context.Context, id int, req hc.ChangeInstanceGroupPermissionsRequest) (*hc.ChangeInstanceGroupPermissionsResponse, error) {
	var resp hc.ChangeInstanceGroupPermissionsResponse

	p := path(hc.InstanceGroupPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockInstanceGroup unlocks the instance group with the given ID.
func (s InstanceService) UnlockInstanceGroup(ctx context.Context, id int) (*hc.UnlockInstanceGroupResponse, error) {
	var resp hc.UnlockInstanceGroupResponse
	p := path(hc.InstanceGroupPath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
