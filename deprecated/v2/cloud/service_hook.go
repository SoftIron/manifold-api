package cloud

import (
	"context"
	"net/url"
)

// HookService own the /cloud/hook methods.
type HookService struct {
	*service
}

// DeleteHook deletes the hook with the given id.
func (s HookService) DeleteHook(ctx context.Context, id int) error {
	p := s.path(HookPath, id)

	return s.Delete(ctx, p, nil)
}

// Hook returns information about a hook.
func (s HookService) Hook(ctx context.Context, id int) (*HookResponse, error) {
	var resp HookResponse

	p := s.path(HookPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Hooks returns information about all hooks.
func (s HookService) Hooks(ctx context.Context) (*HooksResponse, error) {
	var resp HooksResponse

	p := s.path(HookPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HookLog returns the logs for a hook.
func (s HookService) HookLog(ctx context.Context, id int, start, end string) (*HookLogResponse, error) {
	var resp HookLogResponse

	p := s.path(HookPath, id, "log")

	q := make(url.Values)
	q.Add("min", start)
	q.Add("max", end)

	if err := s.Get(ctx, p+"?"+q.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateHook updates the hook with the given id.
func (s HookService) UpdateHook(ctx context.Context, id int, req UpdateHookRequest) (*UpdateHookResponse, error) {
	var resp UpdateHookResponse

	p := s.path(HookPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockHook locks the hook with the given id.
func (s HookService) LockHook(ctx context.Context, id int, req LockHookRequest) (*LockHookResponse, error) {
	var resp LockHookResponse

	p := s.path(HookPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameHook renames the hook with the given id.
func (s HookService) RenameHook(ctx context.Context, id int, req RenameHookRequest) (*RenameHookResponse, error) {
	var resp RenameHookResponse

	p := s.path(HookPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RetryHook retries the hook with the given id.
func (s HookService) RetryHook(ctx context.Context, id int, req RetryHookRequest) (*RetryHookResponse, error) {
	var resp RetryHookResponse

	p := s.path(HookPath, id, "retry")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockHook unlocks the hook with the given id.
func (s HookService) UnlockHook(ctx context.Context, id int, req UnlockHookRequest) (*UnlockHookResponse, error) {
	var resp UnlockHookResponse

	p := s.path(HookPath, id, "unlock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateHook creates a new hook.
func (s HookService) CreateHook(ctx context.Context, req CreateHookRequest) (*CreateHookResponse, error) {
	var resp CreateHookResponse

	p := s.path(HookPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
