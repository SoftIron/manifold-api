package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/cloud"
)

// DeleteInstanceGroup deletes the instance group with the given ID.
func (s InstanceService) DeleteInstanceGroup(ctx context.Context, id int) error {
	p := path(cloud.InstanceGroupPath, id)

	return s.Delete(ctx, p, nil)
}

// InstanceGroup returns information about an instance group.
func (s InstanceService) InstanceGroup(ctx context.Context, id int) (*cloud.ListInstanceGroupResponse, error) {
	var resp cloud.ListInstanceGroupResponse

	p := path(cloud.InstanceGroupPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstanceGroups returns information about all instance groups.
func (s InstanceService) InstanceGroups(ctx context.Context) (*cloud.ListInstanceGroupsResponse, error) {
	var resp cloud.ListInstanceGroupsResponse

	p := path(cloud.InstanceGroupPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateInstanceGroup updates the instance group with the given ID.
func (s InstanceService) UpdateInstanceGroup(ctx context.Context, id int, req *cloud.UpdateInstanceGroupRequest) (*cloud.UpdateInstanceGroupResponse, error) {
	var resp cloud.UpdateInstanceGroupResponse

	p := path(cloud.InstanceGroupPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockInstanceGroup locks the instance group with the given ID.
func (s InstanceService) LockInstanceGroup(ctx context.Context, id int, req cloud.LockInstanceGroupRequest) (*cloud.LockInstanceGroupResponse, error) {
	var resp cloud.LockInstanceGroupResponse

	p := path(cloud.InstanceGroupPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameInstanceGroup renames the instance group with the given ID.
func (s InstanceService) RenameInstanceGroup(ctx context.Context, id int, req cloud.RenameInstanceGroupRequest) (*cloud.RenameInstanceGroupResponse, error) {
	var resp cloud.RenameInstanceGroupResponse

	p := path(cloud.InstanceGroupPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeInstanceGroupOwnership changes the ownership of the instance group with the given ID.
func (s InstanceService) ChangeInstanceGroupOwnership(ctx context.Context, id int, req cloud.ChangeInstanceGroupOwnershipRequest) (*cloud.ChangeInstanceGroupOwnershipResponse, error) {
	var resp cloud.ChangeInstanceGroupOwnershipResponse

	p := path(cloud.InstanceGroupPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeInstanceGroupPermissions changes the permissions of the instance group with the given ID.
func (s InstanceService) ChangeInstanceGroupPermissions(ctx context.Context, id int, req cloud.ChangeInstanceGroupPermissionsRequest) (*cloud.ChangeInstanceGroupPermissionsResponse, error) {
	var resp cloud.ChangeInstanceGroupPermissionsResponse

	p := path(cloud.InstanceGroupPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockInstanceGroup unlocks the instance group with the given ID.
func (s InstanceService) UnlockInstanceGroup(ctx context.Context, id int) (*cloud.UnlockInstanceGroupResponse, error) {
	var resp cloud.UnlockInstanceGroupResponse
	p := path(cloud.InstanceGroupPath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
