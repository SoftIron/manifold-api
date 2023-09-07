package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// DocumentService owns the /document methods.
type DocumentService struct {
	*client.Client
}

// DeleteDocument deletes a document with the given ID.
func (s DocumentService) DeleteDocument(ctx context.Context, id int) error {
	p := path(cloud.DocumentPath, id)

	return s.Delete(ctx, p, nil)
}

// Document returns information about the document with the given ID.
func (s DocumentService) Document(ctx context.Context, id int) (*cloud.ListDocumentResponse, error) {
	var resp cloud.ListDocumentResponse

	p := path(cloud.DocumentPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Documents returns information about all documents.
func (s DocumentService) Documents(ctx context.Context) (*cloud.ListDocumentsResponse, error) {
	var resp cloud.ListDocumentsResponse

	p := path(cloud.DocumentPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateDocument updates the document with the given ID.
func (s DocumentService) UpdateDocument(ctx context.Context, id int, req cloud.UpdateDocumentRequest) (*cloud.UpdateDocumentResponse, error) {
	var resp cloud.UpdateDocumentResponse

	p := path(cloud.DocumentPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockDocument locks the document with the given ID.
func (s DocumentService) LockDocument(ctx context.Context, id int, req cloud.LockDocumentRequest) (*cloud.LockDocumentResponse, error) {
	var resp cloud.LockDocumentResponse

	p := path(cloud.DocumentPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameDocument renames the document with the given ID.
func (s DocumentService) RenameDocument(ctx context.Context, id int, req cloud.RenameDocumentRequest) (*cloud.RenameDocumentResponse, error) {
	var resp cloud.RenameDocumentResponse

	p := path(cloud.DocumentPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeDocumentOwnership changes the ownership of the document with the given ID.
func (s DocumentService) ChangeDocumentOwnership(ctx context.Context, id int, req cloud.ChangeDocumentOwnershipRequest) (*cloud.ChangeDocumentOwnershipResponse, error) {
	var resp cloud.ChangeDocumentOwnershipResponse

	p := path(cloud.DocumentPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeDocumentPermissions changes the permissions of the document with the given ID.
func (s DocumentService) ChangeDocumentPermissions(ctx context.Context, id int, req cloud.ChangeDocumentPermissionsRequest) (*cloud.ChangeDocumentPermissionsResponse, error) {
	var resp cloud.ChangeDocumentPermissionsResponse

	p := path(cloud.DocumentPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockDocument unlocks the document with the given ID.
func (s DocumentService) UnlockDocument(ctx context.Context, id int) (*cloud.UnlockDocumentResponse, error) {
	var resp cloud.UnlockDocumentResponse

	p := path(cloud.DocumentPath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AllocateDocument allocates a document.
func (s DocumentService) AllocateDocument(ctx context.Context, req cloud.AllocateDocumentRequest) (*cloud.AllocateDocumentResponse, error) {
	var resp cloud.AllocateDocumentResponse

	p := path(cloud.DocumentPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneDocument clones the document with the given ID.
func (s DocumentService) CloneDocument(ctx context.Context, id int, req cloud.CloneDocumentRequest) (*cloud.CloneDocumentResponse, error) {
	var resp cloud.CloneDocumentResponse

	p := path(cloud.DocumentPath, id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
