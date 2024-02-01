package cloud

import (
	"context"
)

// SecurityGroupService owns the /cloud/security-group methods.
type SecurityGroupService struct {
	*service
}

// DeleteSecurityGroup deletes the security group with the given ID.
func (s SecurityGroupService) DeleteSecurityGroup(ctx context.Context, id int) error {
	p := s.path(SecurityGroupPath, id)

	return s.Delete(ctx, p, nil)
}

// SecurityGroup returns information about a image.
func (s SecurityGroupService) SecurityGroup(ctx context.Context, id int) (*SecurityGroupResponse, error) {
	var resp SecurityGroupResponse

	p := s.path(SecurityGroupPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SecurityGroups returns information about all images.
func (s SecurityGroupService) SecurityGroups(ctx context.Context) (*SecurityGroupsResponse, error) {
	var resp SecurityGroupsResponse

	p := s.path(SecurityGroupPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateSecurityGroup updates the security group with the given ID.
func (s SecurityGroupService) UpdateSecurityGroup(ctx context.Context, id int, req UpdateSecurityGroupRequest) (*UpdateSecurityGroupResponse, error) {
	var resp UpdateSecurityGroupResponse

	p := s.path(SecurityGroupPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CommitSecurityGroup commits the security group with the given ID.
func (s SecurityGroupService) CommitSecurityGroup(ctx context.Context, id int, req CommitSecurityGroupRequest) (*CommitSecurityGroupResponse, error) {
	var resp CommitSecurityGroupResponse

	p := s.path(SecurityGroupPath, id, "commit")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeSecurityGroupOwnership changes the ownership of the security group with the given ID.
func (s SecurityGroupService) ChangeSecurityGroupOwnership(ctx context.Context, id int, req ChangeSecurityGroupOwnershipRequest) (*ChangeSecurityGroupOwnershipResponse, error) {
	var resp ChangeSecurityGroupOwnershipResponse

	p := s.path(SecurityGroupPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeSecurityGroupPermissions changes the permissions of the security group with the given ID.
func (s SecurityGroupService) ChangeSecurityGroupPermissions(ctx context.Context, id int, req ChangeSecurityGroupPermissionsRequest) (*ChangeSecurityGroupPermissionsResponse, error) {
	var resp ChangeSecurityGroupPermissionsResponse

	p := s.path(SecurityGroupPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameSecurityGroup renames the security group with the given ID.
func (s SecurityGroupService) RenameSecurityGroup(ctx context.Context, id int, req RenameSecurityGroupRequest) (*RenameSecurityGroupResponse, error) {
	var resp RenameSecurityGroupResponse

	p := s.path(SecurityGroupPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateSecurityGroup creates a new security group.
func (s SecurityGroupService) CreateSecurityGroup(ctx context.Context, req CreateSecurityGroupRequest) (*CreateSecurityGroupResponse, error) {
	var resp CreateSecurityGroupResponse

	p := s.path(SecurityGroupPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneSecurityGroup clones the security group with the given ID.
func (s SecurityGroupService) CloneSecurityGroup(ctx context.Context, id int, req CloneSecurityGroupRequest) (*CloneSecurityGroupResponse, error) {
	var resp CloneSecurityGroupResponse

	p := s.path(SecurityGroupPath, id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
