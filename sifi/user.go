package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// UserService owns the /user methods.
type UserService struct {
	*client.Client
}

// DeleteUser deletes the user with the given ID.
func (s UserService) DeleteUser(ctx context.Context, id int) error {
	p := path(cloud.UserPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteUserGroup deletes the user from the group with the given ID.
func (s UserService) DeleteUserGroup(ctx context.Context, id, group int) error {
	p := path(cloud.UserPath, id, "group", group)

	return s.Delete(ctx, p, nil)
}

// User returns information about a user.
func (s UserService) User(ctx context.Context, id int) (*cloud.ListUserResponse, error) {
	var resp cloud.ListUserResponse

	p := path(cloud.UserPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Users returns information about all users.
func (s UserService) Users(ctx context.Context) (*cloud.ListUsersResponse, error) {
	var resp cloud.ListUsersResponse

	p := path(cloud.UserPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UserQuota returns the default user quota.
func (s UserService) UserQuota(ctx context.Context) (*cloud.ListUserQuotaResponse, error) {
	var resp cloud.ListUserQuotaResponse

	p := path(cloud.UserPath, "quota")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateUser updates the user with the given ID.
func (s UserService) UpdateUser(ctx context.Context, id int, req cloud.UpdateUserRequest) (*cloud.UpdateUserResponse, error) {
	var resp cloud.UpdateUserResponse

	p := path(cloud.UserPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeUserAuth changes the user's authentication method.
func (s UserService) ChangeUserAuth(ctx context.Context, id int, req cloud.ChangeUserAuthRequest) (*cloud.ChangeUserAuthResponse, error) {
	var resp cloud.ChangeUserAuthResponse

	p := path(cloud.UserPath, id, "auth")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableUser enables the user with the given ID.
func (s UserService) EnableUser(ctx context.Context, id int, req cloud.EnableUserRequest) (*cloud.EnableUserResponse, error) {
	var resp cloud.EnableUserResponse

	p := path(cloud.UserPath, id, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeUserGroup changes the user's group.
func (s UserService) ChangeUserGroup(ctx context.Context, id, group int) (*cloud.ChangeUserGroupResponse, error) {
	var resp cloud.ChangeUserGroupResponse

	p := path(cloud.UserPath, id, "group", group)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeUserPassword changes the user's password.
func (s UserService) ChangeUserPassword(ctx context.Context, id int, req cloud.ChangeUserPasswordRequest) (*cloud.ChangeUserPasswordResponse, error) {
	var resp cloud.ChangeUserPasswordResponse

	p := path(cloud.UserPath, id, "password")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UserLogin sets the user's login token.
// func (s UserService) UserLogin(ctx context.Context, req cloud.UserLoginRequest) (*cloud.UserLoginResponse, error) {
// 	var resp cloud.UserLoginResponse

// 	p := path(cloud.UserPath, "login")

// 	if err := s.Post(ctx, p, req, &resp); err != nil {
// 		return nil, err
// 	}

// 	return &resp, nil
// }

// UpdateDefaultUserQuota updates the default user quota.
func (s UserService) UpdateDefaultUserQuota(ctx context.Context, req cloud.UpdateDefaultUserQuotaRequest) (*cloud.UpdateDefaultUserQuotaResponse, error) {
	var resp cloud.UpdateDefaultUserQuotaResponse

	p := path(cloud.UserPath, "quota")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateUser creates a new user.
func (s UserService) CreateUser(ctx context.Context, req cloud.CreateUserRequest) (*cloud.CreateUserResponse, error) {
	var resp cloud.CreateUserResponse

	p := path(cloud.UserPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddUserGroup adds the user to the group with the given ID.
func (s UserService) AddUserGroup(ctx context.Context, id, group int) (*cloud.AddUserGroupResponse, error) {
	var resp cloud.AddUserGroupResponse

	p := path(cloud.UserPath, id, "group", group)

	if err := s.Post(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetUserQuota sets the user's quota.
func (s UserService) SetUserQuota(ctx context.Context, id int, req cloud.SetUserQuotaRequest) (*cloud.SetUserQuotaResponse, error) {
	var resp cloud.SetUserQuotaResponse

	p := path(cloud.UserPath, id, "quota")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
