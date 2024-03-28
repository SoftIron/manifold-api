package sifi

import (
	"context"
	"net/url"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/deprecated/v1/hc"
)

// HookService own the /hook methods.
type HookService struct {
	*client.Client
}

// DeleteHook deletes the hook with the given id.
func (s HookService) DeleteHook(ctx context.Context, id int) error {
	p := path(hc.HookPath, id)

	return s.Delete(ctx, p, nil)
}

// Hook returns information about a hook.
func (s HookService) Hook(ctx context.Context, id int) (*hc.ListHookResponse, error) {
	var resp hc.ListHookResponse

	p := path(hc.HookPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Hooks returns information about all hooks.
func (s HookService) Hooks(ctx context.Context) (*hc.ListHooksResponse, error) {
	var resp hc.ListHooksResponse

	p := path(hc.HookPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HookLog returns the logs for a hook.
func (s HookService) HookLog(ctx context.Context, id int, start, end string) (*hc.ListHookLogResponse, error) {
	var resp hc.ListHookLogResponse

	p := path(hc.HookPath, id, "log")

	q := make(url.Values)
	q.Add("min", start)
	q.Add("max", end)

	if err := s.Get(ctx, p+"?"+q.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateHook updates the hook with the given id.
func (s HookService) UpdateHook(ctx context.Context, id int, req hc.UpdateHookRequest) (*hc.UpdateHookResponse, error) {
	var resp hc.UpdateHookResponse

	p := path(hc.HookPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockHook locks the hook with the given id.
func (s HookService) LockHook(ctx context.Context, id int, req hc.LockHookRequest) (*hc.LockHookResponse, error) {
	var resp hc.LockHookResponse

	p := path(hc.HookPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameHook renames the hook with the given id.
func (s HookService) RenameHook(ctx context.Context, id int, req hc.RenameHookRequest) (*hc.RenameHookResponse, error) {
	var resp hc.RenameHookResponse

	p := path(hc.HookPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RetryHook retries the hook with the given id.
func (s HookService) RetryHook(ctx context.Context, id int, req hc.RetryHookRequest) (*hc.RetryHookResponse, error) {
	var resp hc.RetryHookResponse

	p := path(hc.HookPath, id, "retry")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockHook unlocks the hook with the given id.
func (s HookService) UnlockHook(ctx context.Context, id int, req hc.UnlockHookRequest) (*hc.UnlockHookResponse, error) {
	var resp hc.UnlockHookResponse

	p := path(hc.HookPath, id, "unlock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateHook creates a new hook.
func (s HookService) CreateHook(ctx context.Context, req hc.CreateHookRequest) (*hc.CreateHookResponse, error) {
	var resp hc.CreateHookResponse

	p := path(hc.HookPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
