package cloud

import (
	"context"
)

// ACLService owns the /cloud/acl methods.
type ACLService struct {
	*service
}

// DeleteACL deletes the Access Control List specified by the id.
func (s ACLService) DeleteACL(ctx context.Context, id int) error {
	p := s.path(AccessControlListPath, id)

	return s.Delete(ctx, p, nil)
}

// ACLs returns a slice containing all of the Access Control Lists.
func (s ACLService) ACLs(ctx context.Context) (*ACLsResponse, error) {
	var resp ACLsResponse

	p := s.path(AccessControlListPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateACL creates a new Access Control List and returns its id.
func (s ACLService) CreateACL(ctx context.Context, req CreateACLRequest) (*CreateACLResponse, error) {
	var resp CreateACLResponse

	p := s.path(AccessControlListPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
