package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// SecurityGroupService owns the /security-group methods.
type SecurityGroupService struct {
	*client.Client
}

// DeleteSecurityGroup deletes the security group with the given ID.
func (s SecurityGroupService) DeleteSecurityGroup(ctx context.Context, id int) error {
	p := path(cloud.SecurityGroupPath, id)

	return s.Delete(ctx, p, nil)
}

// SecurityGroup returns information about a image.
func (s SecurityGroupService) SecurityGroup(ctx context.Context, id int) (*cloud.ListSecurityGroupResponse, error) {
	var resp cloud.ListSecurityGroupResponse

	p := path(cloud.SecurityGroupPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SecurityGroups returns information about all images.
func (s SecurityGroupService) SecurityGroups(ctx context.Context) (*cloud.ListSecurityGroupsResponse, error) {
	var resp cloud.ListSecurityGroupsResponse

	p := path(cloud.SecurityGroupPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateSecurityGroup updates the security group with the given ID.
func (s SecurityGroupService) UpdateSecurityGroup(ctx context.Context, id int, req cloud.UpdateSecurityGroupRequest) (*cloud.UpdateSecurityGroupResponse, error) {
	var resp cloud.UpdateSecurityGroupResponse

	p := path(cloud.SecurityGroupPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CommitSecurityGroup commits the security group with the given ID.
func (s SecurityGroupService) CommitSecurityGroup(ctx context.Context, id int, req cloud.CommitSecurityGroupRequest) (*cloud.CommitSecurityGroupResponse, error) {
	var resp cloud.CommitSecurityGroupResponse

	p := path(cloud.SecurityGroupPath, id, "commit")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeSecurityGroupOwnership changes the ownership of the security group with the given ID.
func (s SecurityGroupService) ChangeSecurityGroupOwnership(ctx context.Context, id int, req cloud.ChangeSecurityGroupOwnershipRequest) (*cloud.ChangeSecurityGroupOwnershipResponse, error) {
	var resp cloud.ChangeSecurityGroupOwnershipResponse

	p := path(cloud.SecurityGroupPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeSecurityGroupPermissions changes the permissions of the security group with the given ID.
func (s SecurityGroupService) ChangeSecurityGroupPermissions(ctx context.Context, id int, req cloud.ChangeSecurityGroupPermissionsRequest) (*cloud.ChangeSecurityGroupPermissionsResponse, error) {
	var resp cloud.ChangeSecurityGroupPermissionsResponse

	p := path(cloud.SecurityGroupPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameSecurityGroup renames the security group with the given ID.
func (s SecurityGroupService) RenameSecurityGroup(ctx context.Context, id int, req cloud.RenameSecurityGroupRequest) (*cloud.RenameSecurityGroupResponse, error) {
	var resp cloud.RenameSecurityGroupResponse

	p := path(cloud.SecurityGroupPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateSecurityGroup creates a new security group.
func (s SecurityGroupService) CreateSecurityGroup(ctx context.Context, req cloud.CreateSecurityGroupRequest) (*cloud.CreateSecurityGroupResponse, error) {
	var resp cloud.CreateSecurityGroupResponse

	p := path(cloud.SecurityGroupPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneSecurityGroup clones the security group with the given ID.
func (s SecurityGroupService) CloneSecurityGroup(ctx context.Context, id int, req cloud.CloneSecurityGroupRequest) (*cloud.CloneSecurityGroupResponse, error) {
	var resp cloud.CloneSecurityGroupResponse

	p := path(cloud.SecurityGroupPath, id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
