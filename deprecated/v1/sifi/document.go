package sifi

import (
	"context"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/deprecated/v1/hc"
)

// DocumentService owns the /document methods.
type DocumentService struct {
	*client.Client
}

// DeleteDocument deletes a document with the given ID.
func (s DocumentService) DeleteDocument(ctx context.Context, id int) error {
	p := path(hc.DocumentPath, id)

	return s.Delete(ctx, p, nil)
}

// Document returns information about the document with the given ID.
func (s DocumentService) Document(ctx context.Context, id int) (*hc.ListDocumentResponse, error) {
	var resp hc.ListDocumentResponse

	p := path(hc.DocumentPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Documents returns information about all documents.
func (s DocumentService) Documents(ctx context.Context) (*hc.ListDocumentsResponse, error) {
	var resp hc.ListDocumentsResponse

	p := path(hc.DocumentPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateDocument updates the document with the given ID.
func (s DocumentService) UpdateDocument(ctx context.Context, id int, req hc.UpdateDocumentRequest) (*hc.UpdateDocumentResponse, error) {
	var resp hc.UpdateDocumentResponse

	p := path(hc.DocumentPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockDocument locks the document with the given ID.
func (s DocumentService) LockDocument(ctx context.Context, id int, req hc.LockDocumentRequest) (*hc.LockDocumentResponse, error) {
	var resp hc.LockDocumentResponse

	p := path(hc.DocumentPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameDocument renames the document with the given ID.
func (s DocumentService) RenameDocument(ctx context.Context, id int, req hc.RenameDocumentRequest) (*hc.RenameDocumentResponse, error) {
	var resp hc.RenameDocumentResponse

	p := path(hc.DocumentPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeDocumentOwnership changes the ownership of the document with the given ID.
func (s DocumentService) ChangeDocumentOwnership(ctx context.Context, id int, req hc.ChangeDocumentOwnershipRequest) (*hc.ChangeDocumentOwnershipResponse, error) {
	var resp hc.ChangeDocumentOwnershipResponse

	p := path(hc.DocumentPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeDocumentPermissions changes the permissions of the document with the given ID.
func (s DocumentService) ChangeDocumentPermissions(ctx context.Context, id int, req hc.ChangeDocumentPermissionsRequest) (*hc.ChangeDocumentPermissionsResponse, error) {
	var resp hc.ChangeDocumentPermissionsResponse

	p := path(hc.DocumentPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockDocument unlocks the document with the given ID.
func (s DocumentService) UnlockDocument(ctx context.Context, id int) (*hc.UnlockDocumentResponse, error) {
	var resp hc.UnlockDocumentResponse

	p := path(hc.DocumentPath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AllocateDocument allocates a document.
func (s DocumentService) AllocateDocument(ctx context.Context, req hc.AllocateDocumentRequest) (*hc.AllocateDocumentResponse, error) {
	var resp hc.AllocateDocumentResponse

	p := path(hc.DocumentPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneDocument clones the document with the given ID.
func (s DocumentService) CloneDocument(ctx context.Context, id int, req hc.CloneDocumentRequest) (*hc.CloneDocumentResponse, error) {
	var resp hc.CloneDocumentResponse

	p := path(hc.DocumentPath, id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
