package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// ImageService owns the /image methods.
type ImageService struct {
	*client.Client
}

// DeleteImage deletes the image with the given ID.
func (s ImageService) DeleteImage(ctx context.Context, id int) error {
	p := path(cloud.ImagePath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteImageSnapshot deletes a snapshot for the image with the given ID.
func (s ImageService) DeleteImageSnapshot(ctx context.Context, id, snapshot int) error {
	p := path(cloud.ImagePath, id, "snapshot", snapshot)

	return s.Delete(ctx, p, nil)
}

// Image returns information about a image.
func (s ImageService) Image(ctx context.Context, id int) (*cloud.ListImageResponse, error) {
	var resp cloud.ListImageResponse

	p := path(cloud.ImagePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Images returns information about all images.
func (s ImageService) Images(ctx context.Context) (*cloud.ListImagesResponse, error) {
	var resp cloud.ListImagesResponse

	p := path(cloud.ImagePath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateImage updates the image with the given ID.
func (s ImageService) UpdateImage(ctx context.Context, id int, req cloud.UpdateImageRequest) (*cloud.UpdateImageResponse, error) {
	var resp cloud.UpdateImageResponse

	p := path(cloud.ImagePath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableImage enables the image with the given ID.
func (s ImageService) EnableImage(ctx context.Context, id int, req cloud.EnableImageRequest) (*cloud.EnableImageResponse, error) {
	var resp cloud.EnableImageResponse

	p := path(cloud.ImagePath, id, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockImage locks the image with the given ID.
func (s ImageService) LockImage(ctx context.Context, id int, req cloud.LockImageRequest) (*cloud.LockImageResponse, error) {
	var resp cloud.LockImageResponse

	p := path(cloud.ImagePath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameImage renames the image with the given ID.
func (s ImageService) RenameImage(ctx context.Context, id int, req cloud.RenameImageRequest) (*cloud.RenameImageResponse, error) {
	var resp cloud.RenameImageResponse

	p := path(cloud.ImagePath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeImageOwnership changes the ownership of the image with the given ID.
func (s ImageService) ChangeImageOwnership(ctx context.Context, id int, req cloud.ChangeImageOwnershipRequest) (*cloud.ChangeImageOwnershipResponse, error) {
	var resp cloud.ChangeImageOwnershipResponse

	p := path(cloud.ImagePath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeImagePermissions changes the permissions of the image with the given ID.
func (s ImageService) ChangeImagePermissions(ctx context.Context, id int, req cloud.ChangeImagePermissionsRequest) (*cloud.ChangeImagePermissionsResponse, error) {
	var resp cloud.ChangeImagePermissionsResponse

	p := path(cloud.ImagePath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetImagePersistent sets the image with the given ID to be persistent.
func (s ImageService) SetImagePersistent(ctx context.Context, id int, req cloud.SetImagePersistentRequest) (*cloud.SetImagePersistentResponse, error) {
	var resp cloud.SetImagePersistentResponse

	p := path(cloud.ImagePath, id, "persistent")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// FlattenImageSnapshot flattens the snapshot of the image with the given ID.
func (s ImageService) FlattenImageSnapshot(ctx context.Context, id, snapshot int) (*cloud.FlattenImageSnapshotResponse, error) {
	var resp cloud.FlattenImageSnapshotResponse

	p := path(cloud.ImagePath, id, "snapshot", snapshot, "flatten")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RevertImageSnapshot reverts the snapshot of the image with the given ID.
func (s ImageService) RevertImageSnapshot(ctx context.Context, id, snapshot int) (*cloud.RevertImageSnapshotResponse, error) {
	var resp cloud.RevertImageSnapshotResponse

	p := path(cloud.ImagePath, id, "snapshot", snapshot, "revert")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeImageType changes the type of the image with the given ID.
func (s ImageService) ChangeImageType(ctx context.Context, id int, req cloud.ChangeImageTypeRequest) (*cloud.ChangeImageTypeResponse, error) {
	var resp cloud.ChangeImageTypeResponse

	p := path(cloud.ImagePath, id, "type")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockImage unlocks the image with the given ID.
func (s ImageService) UnlockImage(ctx context.Context, id int) (*cloud.UnlockImageResponse, error) {
	var resp cloud.UnlockImageResponse

	p := path(cloud.ImagePath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateImage creates a new image.
func (s ImageService) CreateImage(ctx context.Context, req cloud.CreateImageRequest) (*cloud.CreateImageResponse, error) {
	var resp cloud.CreateImageResponse

	p := path(cloud.ImagePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneImage clones the image with the given ID.
func (s ImageService) CloneImage(ctx context.Context, id int, req cloud.CloneImageRequest) (*cloud.CloneImageResponse, error) {
	var resp cloud.CloneImageResponse

	p := path(cloud.ImagePath, id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
