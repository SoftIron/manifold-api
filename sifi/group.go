package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// GroupService owns the /group methods.
type GroupService struct {
	*client.Client
}

// DeleteGroup deletes the group with the given ID.
func (s GroupService) DeleteGroup(ctx context.Context, id int) error {
	p := path(cloud.GroupPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteGroupAdmin deletes the group admin from the group with the given ID.
func (s GroupService) DeleteGroupAdmin(ctx context.Context, id, user int) error {
	p := path(cloud.GroupPath, id, "admin", user)

	return s.Delete(ctx, p, nil)
}

// Group returns information about a group.
func (s GroupService) Group(ctx context.Context, id int) (*cloud.ListGroupResponse, error) {
	var resp cloud.ListGroupResponse

	p := path(cloud.GroupPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Groups returns information about all groups.
func (s GroupService) Groups(ctx context.Context) (*cloud.ListGroupsResponse, error) {
	var resp cloud.ListGroupsResponse

	p := path(cloud.GroupPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GroupQuota returns the default group quota.
func (s GroupService) GroupQuota(ctx context.Context) (*cloud.ListGroupQuotaResponse, error) {
	var resp cloud.ListGroupQuotaResponse

	p := path(cloud.GroupPath, "quota")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateGroup updates the group with the given ID.
func (s GroupService) UpdateGroup(ctx context.Context, id int, req cloud.UpdateGroupRequest) (*cloud.UpdateGroupResponse, error) {
	var resp cloud.UpdateGroupResponse

	p := path(cloud.GroupPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateGroupQuota updates the default group quota.
func (s GroupService) UpdateGroupQuota(ctx context.Context, req cloud.UpdateGroupQuotaRequest) (*cloud.UpdateGroupQuotaResponse, error) {
	var resp cloud.UpdateGroupQuotaResponse

	p := path(cloud.GroupPath, "quota")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateGroup creates a new group.
func (s GroupService) CreateGroup(ctx context.Context, req cloud.CreateGroupRequest) (*cloud.CreateGroupResponse, error) {
	var resp cloud.CreateGroupResponse

	p := path(cloud.GroupPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddGroupAdmin adds a group admin to the group with the given ID.
func (s GroupService) AddGroupAdmin(ctx context.Context, id int, req cloud.AddGroupAdminRequest) (*cloud.AddGroupAdminResponse, error) {
	var resp cloud.AddGroupAdminResponse

	p := path(cloud.GroupPath, id, "admin")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetGroupQuota sets the quota for the group with the given ID.
func (s GroupService) SetGroupQuota(ctx context.Context, id int, req cloud.SetGroupQuotaRequest) (*cloud.SetGroupQuotaResponse, error) {
	var resp cloud.SetGroupQuotaResponse

	p := path(cloud.GroupPath, id, "quota")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
