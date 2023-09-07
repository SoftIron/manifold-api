package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// TemplateService own the /template methods.
type TemplateService struct {
	*client.Client
}

// DeleteTemplate deletes the template with the given ID.
func (s TemplateService) DeleteTemplate(ctx context.Context, id int) error {
	p := path(cloud.TemplatePath, id)

	return s.Delete(ctx, p, nil)
}

// Template returns information about a template.
func (s TemplateService) Template(ctx context.Context, id int) (*cloud.ListTemplateResponse, error) {
	var resp cloud.ListTemplateResponse

	p := path(cloud.TemplatePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Templates returns information about all templates.
func (s TemplateService) Templates(ctx context.Context) (*cloud.ListTemplatesResponse, error) {
	var resp cloud.ListTemplatesResponse

	p := path(cloud.TemplatePath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateTemplate updates the template with the given ID.
func (s TemplateService) UpdateTemplate(ctx context.Context, id int, req cloud.UpdateTemplateRequest) (*cloud.UpdateTemplateResponse, error) {
	var resp cloud.UpdateTemplateResponse

	p := path(cloud.TemplatePath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockTemplate locks the template with the given ID.
func (s TemplateService) LockTemplate(ctx context.Context, id int, req cloud.LockTemplateRequest) (*cloud.LockTemplateResponse, error) {
	var resp cloud.LockTemplateResponse

	p := path(cloud.TemplatePath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstantiateTemplate instantiates the template with the given ID.
func (s TemplateService) InstantiateTemplate(ctx context.Context, id int, req cloud.InstantiateTemplateRequest) (*cloud.InstantiateTemplateResponse, error) {
	var resp cloud.InstantiateTemplateResponse

	p := path(cloud.TemplatePath, id, "instantiate")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameTemplate renames the template with the given ID.
func (s TemplateService) RenameTemplate(ctx context.Context, id int, req cloud.RenameTemplateRequest) (*cloud.RenameTemplateResponse, error) {
	var resp cloud.RenameTemplateResponse

	p := path(cloud.TemplatePath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeTemplateOwnership changes the ownership of the template with the given ID.
func (s TemplateService) ChangeTemplateOwnership(ctx context.Context, id int, req cloud.ChangeTemplateOwnershipRequest) (*cloud.ChangeTemplateOwnershipResponse, error) {
	var resp cloud.ChangeTemplateOwnershipResponse

	p := path(cloud.TemplatePath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeTemplatePermissions changes the permissions of the template with the given ID.
func (s TemplateService) ChangeTemplatePermissions(ctx context.Context, id int, req cloud.ChangeTemplatePermissionsRequest) (*cloud.ChangeTemplatePermissionsResponse, error) {
	var resp cloud.ChangeTemplatePermissionsResponse

	p := path(cloud.TemplatePath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockTemplate unlocks the template with the given ID.
func (s TemplateService) UnlockTemplate(ctx context.Context, id int) (*cloud.UnlockTemplateResponse, error) {
	var resp cloud.UnlockTemplateResponse

	p := path(cloud.TemplatePath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateTemplate creates a new template.
func (s TemplateService) CreateTemplate(ctx context.Context, req cloud.CreateTemplateRequest) (*cloud.CreateTemplateResponse, error) {
	var resp cloud.CreateTemplateResponse

	p := path(cloud.TemplatePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneTemplate clones the template with the given ID.
func (s TemplateService) CloneTemplate(ctx context.Context, id int, req cloud.CloneTemplateRequest) (*cloud.CloneTemplateResponse, error) {
	var resp cloud.CloneTemplateResponse

	p := path(cloud.TemplatePath, id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
