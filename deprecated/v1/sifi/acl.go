package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/deprecated/v1/hc"
)

// ACLService owns the /acl methods.
type ACLService struct {
	*client.Client
}

// DeleteACL deletes the Access Control List specified by the id.
func (s ACLService) DeleteACL(ctx context.Context, id int) error {
	p := path(hc.AccessControlListPath, id)

	return s.Delete(ctx, p, nil)
}

// ACLs returns a slice containing all of the Access Control Lists.
func (s ACLService) ACLs(ctx context.Context) (*hc.ListACLsResponse, error) {
	var resp hc.ListACLsResponse

	p := path(hc.AccessControlListPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateACL creates a new Access Control List and returns its id.
func (s ACLService) CreateACL(ctx context.Context, req hc.CreateACLRequest) (*hc.CreateACLResponse, error) {
	var resp hc.CreateACLResponse

	p := path(hc.AccessControlListPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
