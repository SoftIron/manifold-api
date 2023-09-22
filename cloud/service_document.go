package cloud

import (
	"context"
)

// DocumentService owns the /cloud/document methods.
type DocumentService struct {
	*service
}

// DeleteDocument deletes a document with the given ID.
func (s DocumentService) DeleteDocument(ctx context.Context, id int) error {
	p := s.path(DocumentPath, id)

	return s.Delete(ctx, p, nil)
}

// Document returns information about the document with the given ID.
func (s DocumentService) Document(ctx context.Context, id int) (*DocumentResponse, error) {
	var resp DocumentResponse

	p := s.path(DocumentPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Documents returns information about all documents.
func (s DocumentService) Documents(ctx context.Context) (*DocumentsResponse, error) {
	var resp DocumentsResponse

	p := s.path(DocumentPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateDocument updates the document with the given ID.
func (s DocumentService) UpdateDocument(ctx context.Context, id int, req UpdateDocumentRequest) (*UpdateDocumentResponse, error) {
	var resp UpdateDocumentResponse

	p := s.path(DocumentPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockDocument locks the document with the given ID.
func (s DocumentService) LockDocument(ctx context.Context, id int, req LockDocumentRequest) (*LockDocumentResponse, error) {
	var resp LockDocumentResponse

	p := s.path(DocumentPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameDocument renames the document with the given ID.
func (s DocumentService) RenameDocument(ctx context.Context, id int, req RenameDocumentRequest) (*RenameDocumentResponse, error) {
	var resp RenameDocumentResponse

	p := s.path(DocumentPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeDocumentOwnership changes the ownership of the document with the given ID.
func (s DocumentService) ChangeDocumentOwnership(ctx context.Context, id int, req ChangeDocumentOwnershipRequest) (*ChangeDocumentOwnershipResponse, error) {
	var resp ChangeDocumentOwnershipResponse

	p := s.path(DocumentPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeDocumentPermissions changes the permissions of the document with the given ID.
func (s DocumentService) ChangeDocumentPermissions(ctx context.Context, id int, req ChangeDocumentPermissionsRequest) (*ChangeDocumentPermissionsResponse, error) {
	var resp ChangeDocumentPermissionsResponse

	p := s.path(DocumentPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockDocument unlocks the document with the given ID.
func (s DocumentService) UnlockDocument(ctx context.Context, id int) (*UnlockDocumentResponse, error) {
	var resp UnlockDocumentResponse

	p := s.path(DocumentPath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AllocateDocument allocates a document.
func (s DocumentService) AllocateDocument(ctx context.Context, req AllocateDocumentRequest) (*AllocateDocumentResponse, error) {
	var resp AllocateDocumentResponse

	p := s.path(DocumentPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneDocument clones the document with the given ID.
func (s DocumentService) CloneDocument(ctx context.Context, id int, req CloneDocumentRequest) (*CloneDocumentResponse, error) {
	var resp CloneDocumentResponse

	p := s.path(DocumentPath, id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
