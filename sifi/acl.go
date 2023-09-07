package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// ACLService owns the /acl methods.
type ACLService struct {
	*client.Client
}

// DeleteACL deletes the Access Control List specified by the id.
func (s ACLService) DeleteACL(ctx context.Context, id int) error {
	p := path(cloud.AccessControlListPath, id)

	return s.Delete(ctx, p, nil)
}

// ACLs returns a slice containing all of the Access Control Lists.
func (s ACLService) ACLs(ctx context.Context) (*cloud.ListACLsResponse, error) {
	var resp cloud.ListACLsResponse

	p := path(cloud.AccessControlListPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateACL creates a new Access Control List and returns its id.
func (s ACLService) CreateACL(ctx context.Context, req cloud.CreateACLRequest) (*cloud.CreateACLResponse, error) {
	var resp cloud.CreateACLResponse

	p := path(cloud.AccessControlListPath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
