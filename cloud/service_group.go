package cloud

import (
	"context"
)

// GroupService owns the /cloud/group methods.
type GroupService struct {
	*service
}

// DeleteGroup deletes the group with the given ID.
func (s GroupService) DeleteGroup(ctx context.Context, id int) error {
	p := s.path(GroupPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteGroupAdmin deletes the group admin from the group with the given ID.
func (s GroupService) DeleteGroupAdmin(ctx context.Context, id, user int) error {
	p := s.path(GroupPath, id, "admin", user)

	return s.Delete(ctx, p, nil)
}

// Group returns information about a group.
func (s GroupService) Group(ctx context.Context, id int) (*GroupResponse, error) {
	var resp GroupResponse

	p := s.path(GroupPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Groups returns information about all groups.
func (s GroupService) Groups(ctx context.Context) (*GroupsResponse, error) {
	var resp GroupsResponse

	p := s.path(GroupPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GroupQuota returns the default group quota.
func (s GroupService) GroupQuota(ctx context.Context) (*GroupQuotaResponse, error) {
	var resp GroupQuotaResponse

	p := s.path(GroupPath, "quota")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateGroup updates the group with the given ID.
func (s GroupService) UpdateGroup(ctx context.Context, id int, req UpdateGroupRequest) (*UpdateGroupResponse, error) {
	var resp UpdateGroupResponse

	p := s.path(GroupPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateGroupQuota updates the default group quota.
func (s GroupService) UpdateGroupQuota(ctx context.Context, req UpdateGroupQuotaRequest) (*UpdateGroupQuotaResponse, error) {
	var resp UpdateGroupQuotaResponse

	p := s.path(GroupPath, "quota")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateGroup creates a new group.
func (s GroupService) CreateGroup(ctx context.Context, req CreateGroupRequest) (*CreateGroupResponse, error) {
	var resp CreateGroupResponse

	p := s.path(GroupPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddGroupAdmin adds a group admin to the group with the given ID.
func (s GroupService) AddGroupAdmin(ctx context.Context, id int, req AddGroupAdminRequest) (*AddGroupAdminResponse, error) {
	var resp AddGroupAdminResponse

	p := s.path(GroupPath, id, "admin")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetGroupQuota sets the quota for the group with the given ID.
func (s GroupService) SetGroupQuota(ctx context.Context, id int, req SetGroupQuotaRequest) (*SetGroupQuotaResponse, error) {
	var resp SetGroupQuotaResponse

	p := s.path(GroupPath, id, "quota")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
