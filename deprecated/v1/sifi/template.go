package sifi

import (
	"context"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/deprecated/v1/hc"
)

// TemplateService own the /template methods.
type TemplateService struct {
	*client.Client
}

// DeleteTemplate deletes the template with the given ID.
func (s TemplateService) DeleteTemplate(ctx context.Context, id int) error {
	p := path(hc.TemplatePath, id)

	return s.Delete(ctx, p, nil)
}

// Template returns information about a template.
func (s TemplateService) Template(ctx context.Context, id int) (*hc.ListTemplateResponse, error) {
	var resp hc.ListTemplateResponse

	p := path(hc.TemplatePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Templates returns information about all templates.
func (s TemplateService) Templates(ctx context.Context) (*hc.ListTemplatesResponse, error) {
	var resp hc.ListTemplatesResponse

	p := path(hc.TemplatePath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateTemplate updates the template with the given ID.
func (s TemplateService) UpdateTemplate(ctx context.Context, id int, req hc.UpdateTemplateRequest) (*hc.UpdateTemplateResponse, error) {
	var resp hc.UpdateTemplateResponse

	p := path(hc.TemplatePath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockTemplate locks the template with the given ID.
func (s TemplateService) LockTemplate(ctx context.Context, id int, req hc.LockTemplateRequest) (*hc.LockTemplateResponse, error) {
	var resp hc.LockTemplateResponse

	p := path(hc.TemplatePath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstantiateTemplate instantiates the template with the given ID.
func (s TemplateService) InstantiateTemplate(ctx context.Context, id int, req hc.InstantiateTemplateRequest) (*hc.InstantiateTemplateResponse, error) {
	var resp hc.InstantiateTemplateResponse

	p := path(hc.TemplatePath, id, "instantiate")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameTemplate renames the template with the given ID.
func (s TemplateService) RenameTemplate(ctx context.Context, id int, req hc.RenameTemplateRequest) (*hc.RenameTemplateResponse, error) {
	var resp hc.RenameTemplateResponse

	p := path(hc.TemplatePath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeTemplateOwnership changes the ownership of the template with the given ID.
func (s TemplateService) ChangeTemplateOwnership(ctx context.Context, id int, req hc.ChangeTemplateOwnershipRequest) (*hc.ChangeTemplateOwnershipResponse, error) {
	var resp hc.ChangeTemplateOwnershipResponse

	p := path(hc.TemplatePath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeTemplatePermissions changes the permissions of the template with the given ID.
func (s TemplateService) ChangeTemplatePermissions(ctx context.Context, id int, req hc.ChangeTemplatePermissionsRequest) (*hc.ChangeTemplatePermissionsResponse, error) {
	var resp hc.ChangeTemplatePermissionsResponse

	p := path(hc.TemplatePath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockTemplate unlocks the template with the given ID.
func (s TemplateService) UnlockTemplate(ctx context.Context, id int) (*hc.UnlockTemplateResponse, error) {
	var resp hc.UnlockTemplateResponse

	p := path(hc.TemplatePath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateTemplate creates a new template.
func (s TemplateService) CreateTemplate(ctx context.Context, req hc.CreateTemplateRequest) (*hc.CreateTemplateResponse, error) {
	var resp hc.CreateTemplateResponse

	p := path(hc.TemplatePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneTemplate clones the template with the given ID.
func (s TemplateService) CloneTemplate(ctx context.Context, id int, req hc.CloneTemplateRequest) (*hc.CloneTemplateResponse, error) {
	var resp hc.CloneTemplateResponse

	p := path(hc.TemplatePath, id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
