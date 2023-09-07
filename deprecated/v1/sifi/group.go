package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/deprecated/v1/hc"
)

// GroupService owns the /group methods.
type GroupService struct {
	*client.Client
}

// DeleteGroup deletes the group with the given ID.
func (s GroupService) DeleteGroup(ctx context.Context, id int) error {
	p := path(hc.GroupPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteGroupAdmin deletes the group admin from the group with the given ID.
func (s GroupService) DeleteGroupAdmin(ctx context.Context, id, user int) error {
	p := path(hc.GroupPath, id, "admin", user)

	return s.Delete(ctx, p, nil)
}

// Group returns information about a group.
func (s GroupService) Group(ctx context.Context, id int) (*hc.ListGroupResponse, error) {
	var resp hc.ListGroupResponse

	p := path(hc.GroupPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Groups returns information about all groups.
func (s GroupService) Groups(ctx context.Context) (*hc.ListGroupsResponse, error) {
	var resp hc.ListGroupsResponse

	p := path(hc.GroupPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GroupQuota returns the default group quota.
func (s GroupService) GroupQuota(ctx context.Context) (*hc.ListGroupQuotaResponse, error) {
	var resp hc.ListGroupQuotaResponse

	p := path(hc.GroupPath, "quota")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateGroup updates the group with the given ID.
func (s GroupService) UpdateGroup(ctx context.Context, id int, req hc.UpdateGroupRequest) (*hc.UpdateGroupResponse, error) {
	var resp hc.UpdateGroupResponse

	p := path(hc.GroupPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateGroupQuota updates the default group quota.
func (s GroupService) UpdateGroupQuota(ctx context.Context, req hc.UpdateGroupQuotaRequest) (*hc.UpdateGroupQuotaResponse, error) {
	var resp hc.UpdateGroupQuotaResponse

	p := path(hc.GroupPath, "quota")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateGroup creates a new group.
func (s GroupService) CreateGroup(ctx context.Context, req hc.CreateGroupRequest) (*hc.CreateGroupResponse, error) {
	var resp hc.CreateGroupResponse

	p := path(hc.GroupPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddGroupAdmin adds a group admin to the group with the given ID.
func (s GroupService) AddGroupAdmin(ctx context.Context, id int, req hc.AddGroupAdminRequest) (*hc.AddGroupAdminResponse, error) {
	var resp hc.AddGroupAdminResponse

	p := path(hc.GroupPath, id, "admin")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetGroupQuota sets the quota for the group with the given ID.
func (s GroupService) SetGroupQuota(ctx context.Context, id int, req hc.SetGroupQuotaRequest) (*hc.SetGroupQuotaResponse, error) {
	var resp hc.SetGroupQuotaResponse

	p := path(hc.GroupPath, id, "quota")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
