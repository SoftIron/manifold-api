package cloud

import (
	"context"
)

// UserService owns the /cloud/user methods.
type UserService struct {
	*service
}

// DeleteUser deletes the user with the given ID.
func (s UserService) DeleteUser(ctx context.Context, id int) error {
	p := s.path(UserPath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteUserGroup deletes the user from the group with the given ID.
func (s UserService) DeleteUserGroup(ctx context.Context, id, group int) error {
	p := s.path(UserPath, id, "group", group)

	return s.Delete(ctx, p, nil)
}

// User returns information about a user.
func (s UserService) User(ctx context.Context, id int) (*UserResponse, error) {
	var resp UserResponse

	p := s.path(UserPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Users returns information about all users.
func (s UserService) Users(ctx context.Context) (*UsersResponse, error) {
	var resp UsersResponse

	p := s.path(UserPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UserQuota returns the default user quota.
func (s UserService) UserQuota(ctx context.Context) (*UserQuotaResponse, error) {
	var resp UserQuotaResponse

	p := s.path(UserPath, "quota")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateUser updates the user with the given ID.
func (s UserService) UpdateUser(ctx context.Context, id int, req UpdateUserRequest) (*UpdateUserResponse, error) {
	var resp UpdateUserResponse

	p := s.path(UserPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeUserAuth changes the user's authentication method.
func (s UserService) ChangeUserAuth(ctx context.Context, id int, req ChangeUserAuthRequest) (*ChangeUserAuthResponse, error) {
	var resp ChangeUserAuthResponse

	p := s.path(UserPath, id, "auth")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableUser enables the user with the given ID.
func (s UserService) EnableUser(ctx context.Context, id int, req EnableUserRequest) (*EnableUserResponse, error) {
	var resp EnableUserResponse

	p := s.path(UserPath, id, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeUserGroup changes the user's group.
func (s UserService) ChangeUserGroup(ctx context.Context, id, group int) (*ChangeUserGroupResponse, error) {
	var resp ChangeUserGroupResponse

	p := s.path(UserPath, id, "group", group)

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeUserPassword changes the user's password.
func (s UserService) ChangeUserPassword(ctx context.Context, id int, req ChangeUserPasswordRequest) (*ChangeUserPasswordResponse, error) {
	var resp ChangeUserPasswordResponse

	p := s.path(UserPath, id, "password")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UserLogin sets the user's login token.
// func (s UserService) UserLogin(ctx context.Context, req UserLoginRequest) (*UserLoginResponse, error) {
// 	var resp UserLoginResponse

// 	p := s.path( UserPath, "login")

// 	if err := s.Post(ctx, p, req, &resp); err != nil {
// 		return nil, err
// 	}

// 	return &resp, nil
// }

// UpdateDefaultUserQuota updates the default user quota.
func (s UserService) UpdateDefaultUserQuota(ctx context.Context, req UpdateDefaultUserQuotaRequest) (*UpdateDefaultUserQuotaResponse, error) {
	var resp UpdateDefaultUserQuotaResponse

	p := s.path(UserPath, "quota")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateUser creates a new user.
func (s UserService) CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error) {
	var resp CreateUserResponse

	p := s.path(UserPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddUserGroup adds the user to the group with the given ID.
func (s UserService) AddUserGroup(ctx context.Context, id, group int) (*AddUserGroupResponse, error) {
	var resp AddUserGroupResponse

	p := s.path(UserPath, id, "group", group)

	if err := s.Post(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetUserQuota sets the user's quota.
func (s UserService) SetUserQuota(ctx context.Context, id int, req SetUserQuotaRequest) (*SetUserQuotaResponse, error) {
	var resp SetUserQuotaResponse

	p := s.path(UserPath, id, "quota")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
