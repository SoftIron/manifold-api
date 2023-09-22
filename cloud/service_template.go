package cloud

import (
	"context"
)

// TemplateService own the /cloud/template methods.
type TemplateService struct {
	*service
}

// DeleteTemplate deletes the template with the given ID.
func (s TemplateService) DeleteTemplate(ctx context.Context, id int) error {
	p := s.path(TemplatePath, id)

	return s.Delete(ctx, p, nil)
}

// Template returns information about a template.
func (s TemplateService) Template(ctx context.Context, id int) (*TemplateResponse, error) {
	var resp TemplateResponse

	p := s.path(TemplatePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Templates returns information about all templates.
func (s TemplateService) Templates(ctx context.Context) (*TemplatesResponse, error) {
	var resp TemplatesResponse

	p := s.path(TemplatePath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateTemplate updates the template with the given ID.
func (s TemplateService) UpdateTemplate(ctx context.Context, id int, req UpdateTemplateRequest) (*UpdateTemplateResponse, error) {
	var resp UpdateTemplateResponse

	p := s.path(TemplatePath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockTemplate locks the template with the given ID.
func (s TemplateService) LockTemplate(ctx context.Context, id int, req LockTemplateRequest) (*LockTemplateResponse, error) {
	var resp LockTemplateResponse

	p := s.path(TemplatePath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstantiateTemplate instantiates the template with the given ID.
func (s TemplateService) InstantiateTemplate(ctx context.Context, id int, req InstantiateTemplateRequest) (*InstantiateTemplateResponse, error) {
	var resp InstantiateTemplateResponse

	p := s.path(TemplatePath, id, "instantiate")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameTemplate renames the template with the given ID.
func (s TemplateService) RenameTemplate(ctx context.Context, id int, req RenameTemplateRequest) (*RenameTemplateResponse, error) {
	var resp RenameTemplateResponse

	p := s.path(TemplatePath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeTemplateOwnership changes the ownership of the template with the given ID.
func (s TemplateService) ChangeTemplateOwnership(ctx context.Context, id int, req ChangeTemplateOwnershipRequest) (*ChangeTemplateOwnershipResponse, error) {
	var resp ChangeTemplateOwnershipResponse

	p := s.path(TemplatePath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeTemplatePermissions changes the permissions of the template with the given ID.
func (s TemplateService) ChangeTemplatePermissions(ctx context.Context, id int, req ChangeTemplatePermissionsRequest) (*ChangeTemplatePermissionsResponse, error) {
	var resp ChangeTemplatePermissionsResponse

	p := s.path(TemplatePath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockTemplate unlocks the template with the given ID.
func (s TemplateService) UnlockTemplate(ctx context.Context, id int) (*UnlockTemplateResponse, error) {
	var resp UnlockTemplateResponse

	p := s.path(TemplatePath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateTemplate creates a new template.
func (s TemplateService) CreateTemplate(ctx context.Context, req CreateTemplateRequest) (*CreateTemplateResponse, error) {
	var resp CreateTemplateResponse

	p := s.path(TemplatePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneTemplate clones the template with the given ID.
func (s TemplateService) CloneTemplate(ctx context.Context, id int, req CloneTemplateRequest) (*CloneTemplateResponse, error) {
	var resp CloneTemplateResponse

	p := s.path(TemplatePath, id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
