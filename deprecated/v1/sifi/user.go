package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/deprecated/v1/hc"
)

// UserService owns the /user methods.
type UserService struct {
	*client.Client
}

// DeleteUser deletes the user with the given ID.
func (s UserService) DeleteUser(ctx context.Context, id int) error {
	p := path(hc.UserPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteUserGroup deletes the user from the group with the given ID.
func (s UserService) DeleteUserGroup(ctx context.Context, id, group int) error {
	p := path(hc.UserPath, id, "group", group)

	return s.Delete(ctx, p, nil)
}

// User returns information about a user.
func (s UserService) User(ctx context.Context, id int) (*hc.ListUserResponse, error) {
	var resp hc.ListUserResponse

	p := path(hc.UserPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Users returns information about all users.
func (s UserService) Users(ctx context.Context) (*hc.ListUsersResponse, error) {
	var resp hc.ListUsersResponse

	p := path(hc.UserPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UserQuota returns the default user quota.
func (s UserService) UserQuota(ctx context.Context) (*hc.ListUserQuotaResponse, error) {
	var resp hc.ListUserQuotaResponse

	p := path(hc.UserPath, "quota")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateUser updates the user with the given ID.
func (s UserService) UpdateUser(ctx context.Context, id int, req hc.UpdateUserRequest) (*hc.UpdateUserResponse, error) {
	var resp hc.UpdateUserResponse

	p := path(hc.UserPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeUserAuth changes the user's authentication method.
func (s UserService) ChangeUserAuth(ctx context.Context, id int, req hc.ChangeUserAuthRequest) (*hc.ChangeUserAuthResponse, error) {
	var resp hc.ChangeUserAuthResponse

	p := path(hc.UserPath, id, "auth")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableUser enables the user with the given ID.
func (s UserService) EnableUser(ctx context.Context, id int, req hc.EnableUserRequest) (*hc.EnableUserResponse, error) {
	var resp hc.EnableUserResponse

	p := path(hc.UserPath, id, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeUserGroup changes the user's group.
func (s UserService) ChangeUserGroup(ctx context.Context, id, group int) (*hc.ChangeUserGroupResponse, error) {
	var resp hc.ChangeUserGroupResponse

	p := path(hc.UserPath, id, "group", group)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeUserPassword changes the user's password.
func (s UserService) ChangeUserPassword(ctx context.Context, id int, req hc.ChangeUserPasswordRequest) (*hc.ChangeUserPasswordResponse, error) {
	var resp hc.ChangeUserPasswordResponse

	p := path(hc.UserPath, id, "password")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateDefaultUserQuota updates the default user quota.
func (s UserService) UpdateDefaultUserQuota(ctx context.Context, req hc.UpdateDefaultUserQuotaRequest) (*hc.UpdateDefaultUserQuotaResponse, error) {
	var resp hc.UpdateDefaultUserQuotaResponse

	p := path(hc.UserPath, "quota")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateUser creates a new user.
func (s UserService) CreateUser(ctx context.Context, req hc.CreateUserRequest) (*hc.CreateUserResponse, error) {
	var resp hc.CreateUserResponse

	p := path(hc.UserPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddUserGroup adds the user to the group with the given ID.
func (s UserService) AddUserGroup(ctx context.Context, id, group int) (*hc.AddUserGroupResponse, error) {
	var resp hc.AddUserGroupResponse

	p := path(hc.UserPath, id, "group", group)

	if err := s.Post(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetUserQuota sets the user's quota.
func (s UserService) SetUserQuota(ctx context.Context, id int, req hc.SetUserQuotaRequest) (*hc.SetUserQuotaResponse, error) {
	var resp hc.SetUserQuotaResponse

	p := path(hc.UserPath, id, "quota")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
