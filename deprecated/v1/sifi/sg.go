package sifi

import (
	"context"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/deprecated/v1/hc"
)

// SecurityGroupService owns the /security-group methods.
type SecurityGroupService struct {
	*client.Client
}

// DeleteSecurityGroup deletes the security group with the given ID.
func (s SecurityGroupService) DeleteSecurityGroup(ctx context.Context, id int) error {
	p := path(hc.SecurityGroupPath, id)

	return s.Delete(ctx, p, nil)
}

// SecurityGroup returns information about a image.
func (s SecurityGroupService) SecurityGroup(ctx context.Context, id int) (*hc.ListSecurityGroupResponse, error) {
	var resp hc.ListSecurityGroupResponse

	p := path(hc.SecurityGroupPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SecurityGroups returns information about all images.
func (s SecurityGroupService) SecurityGroups(ctx context.Context) (*hc.ListSecurityGroupsResponse, error) {
	var resp hc.ListSecurityGroupsResponse

	p := path(hc.SecurityGroupPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateSecurityGroup updates the security group with the given ID.
func (s SecurityGroupService) UpdateSecurityGroup(ctx context.Context, id int, req hc.UpdateSecurityGroupRequest) (*hc.UpdateSecurityGroupResponse, error) {
	var resp hc.UpdateSecurityGroupResponse

	p := path(hc.SecurityGroupPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CommitSecurityGroup commits the security group with the given ID.
func (s SecurityGroupService) CommitSecurityGroup(ctx context.Context, id int, req hc.CommitSecurityGroupRequest) (*hc.CommitSecurityGroupResponse, error) {
	var resp hc.CommitSecurityGroupResponse

	p := path(hc.SecurityGroupPath, id, "commit")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeSecurityGroupOwnership changes the ownership of the security group with the given ID.
func (s SecurityGroupService) ChangeSecurityGroupOwnership(ctx context.Context, id int, req hc.ChangeSecurityGroupOwnershipRequest) (*hc.ChangeSecurityGroupOwnershipResponse, error) {
	var resp hc.ChangeSecurityGroupOwnershipResponse

	p := path(hc.SecurityGroupPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeSecurityGroupPermissions changes the permissions of the security group with the given ID.
func (s SecurityGroupService) ChangeSecurityGroupPermissions(ctx context.Context, id int, req hc.ChangeSecurityGroupPermissionsRequest) (*hc.ChangeSecurityGroupPermissionsResponse, error) {
	var resp hc.ChangeSecurityGroupPermissionsResponse

	p := path(hc.SecurityGroupPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameSecurityGroup renames the security group with the given ID.
func (s SecurityGroupService) RenameSecurityGroup(ctx context.Context, id int, req hc.RenameSecurityGroupRequest) (*hc.RenameSecurityGroupResponse, error) {
	var resp hc.RenameSecurityGroupResponse

	p := path(hc.SecurityGroupPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateSecurityGroup creates a new security group.
func (s SecurityGroupService) CreateSecurityGroup(ctx context.Context, req hc.CreateSecurityGroupRequest) (*hc.CreateSecurityGroupResponse, error) {
	var resp hc.CreateSecurityGroupResponse

	p := path(hc.SecurityGroupPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneSecurityGroup clones the security group with the given ID.
func (s SecurityGroupService) CloneSecurityGroup(ctx context.Context, id int, req hc.CloneSecurityGroupRequest) (*hc.CloneSecurityGroupResponse, error) {
	var resp hc.CloneSecurityGroupResponse

	p := path(hc.SecurityGroupPath, id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
