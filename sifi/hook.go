package sifi

import (
	"context"
	"net/url"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// HookService own the /hook methods.
type HookService struct {
	*client.Client
}

// DeleteHook deletes the hook with the given id.
func (s HookService) DeleteHook(ctx context.Context, id int) error {
	p := path(cloud.HookPath, id)

	return s.Delete(ctx, p, nil)
}

// Hook returns information about a hook.
func (s HookService) Hook(ctx context.Context, id int) (*cloud.ListHookResponse, error) {
	var resp cloud.ListHookResponse

	p := path(cloud.HookPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Hooks returns information about all hooks.
func (s HookService) Hooks(ctx context.Context) (*cloud.ListHooksResponse, error) {
	var resp cloud.ListHooksResponse

	p := path(cloud.HookPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HookLog returns the logs for a hook.
func (s HookService) HookLog(ctx context.Context, id int, start, end string) (*cloud.ListHookLogResponse, error) {
	var resp cloud.ListHookLogResponse

	p := path(cloud.HookPath, id, "log")

	q := make(url.Values)
	q.Add("min", start)
	q.Add("max", end)

	if err := s.Get(ctx, p+"?"+q.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateHook updates the hook with the given id.
func (s HookService) UpdateHook(ctx context.Context, id int, req cloud.UpdateHookRequest) (*cloud.UpdateHookResponse, error) {
	var resp cloud.UpdateHookResponse

	p := path(cloud.HookPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockHook locks the hook with the given id.
func (s HookService) LockHook(ctx context.Context, id int, req cloud.LockHookRequest) (*cloud.LockHookResponse, error) {
	var resp cloud.LockHookResponse

	p := path(cloud.HookPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameHook renames the hook with the given id.
func (s HookService) RenameHook(ctx context.Context, id int, req cloud.RenameHookRequest) (*cloud.RenameHookResponse, error) {
	var resp cloud.RenameHookResponse

	p := path(cloud.HookPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RetryHook retries the hook with the given id.
func (s HookService) RetryHook(ctx context.Context, id int, req cloud.RetryHookRequest) (*cloud.RetryHookResponse, error) {
	var resp cloud.RetryHookResponse

	p := path(cloud.HookPath, id, "retry")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockHook unlocks the hook with the given id.
func (s HookService) UnlockHook(ctx context.Context, id int, req cloud.UnlockHookRequest) (*cloud.UnlockHookResponse, error) {
	var resp cloud.UnlockHookResponse

	p := path(cloud.HookPath, id, "unlock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateHook creates a new hook.
func (s HookService) CreateHook(ctx context.Context, req cloud.CreateHookRequest) (*cloud.CreateHookResponse, error) {
	var resp cloud.CreateHookResponse

	p := path(cloud.HookPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
