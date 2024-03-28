package sifi

import (
	"context"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/deprecated/v1/hc"
)

// ImageService owns the /image methods.
type ImageService struct {
	*client.Client
}

// DeleteImage deletes the image with the given ID.
func (s ImageService) DeleteImage(ctx context.Context, id int) error {
	p := path(hc.ImagePath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteImageSnapshot deletes a snapshot for the image with the given ID.
func (s ImageService) DeleteImageSnapshot(ctx context.Context, id, snapshot int) error {
	p := path(hc.ImagePath, id, "snapshot", snapshot)

	return s.Delete(ctx, p, nil)
}

// Image returns information about a image.
func (s ImageService) Image(ctx context.Context, id int) (*hc.ListImageResponse, error) {
	var resp hc.ListImageResponse

	p := path(hc.ImagePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Images returns information about all images.
func (s ImageService) Images(ctx context.Context) (*hc.ListImagesResponse, error) {
	var resp hc.ListImagesResponse

	p := path(hc.ImagePath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateImage updates the image with the given ID.
func (s ImageService) UpdateImage(ctx context.Context, id int, req hc.UpdateImageRequest) (*hc.UpdateImageResponse, error) {
	var resp hc.UpdateImageResponse

	p := path(hc.ImagePath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableImage enables the image with the given ID.
func (s ImageService) EnableImage(ctx context.Context, id int, req hc.EnableImageRequest) (*hc.EnableImageResponse, error) {
	var resp hc.EnableImageResponse

	p := path(hc.ImagePath, id, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockImage locks the image with the given ID.
func (s ImageService) LockImage(ctx context.Context, id int, req hc.LockImageRequest) (*hc.LockImageResponse, error) {
	var resp hc.LockImageResponse

	p := path(hc.ImagePath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameImage renames the image with the given ID.
func (s ImageService) RenameImage(ctx context.Context, id int, req hc.RenameImageRequest) (*hc.RenameImageResponse, error) {
	var resp hc.RenameImageResponse

	p := path(hc.ImagePath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeImageOwnership changes the ownership of the image with the given ID.
func (s ImageService) ChangeImageOwnership(ctx context.Context, id int, req hc.ChangeImageOwnershipRequest) (*hc.ChangeImageOwnershipResponse, error) {
	var resp hc.ChangeImageOwnershipResponse

	p := path(hc.ImagePath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeImagePermissions changes the permissions of the image with the given ID.
func (s ImageService) ChangeImagePermissions(ctx context.Context, id int, req hc.ChangeImagePermissionsRequest) (*hc.ChangeImagePermissionsResponse, error) {
	var resp hc.ChangeImagePermissionsResponse

	p := path(hc.ImagePath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetImagePersistent sets the image with the given ID to be persistent.
func (s ImageService) SetImagePersistent(ctx context.Context, id int, req hc.SetImagePersistentRequest) (*hc.SetImagePersistentResponse, error) {
	var resp hc.SetImagePersistentResponse

	p := path(hc.ImagePath, id, "persistent")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// FlattenImageSnapshot flattens the snapshot of the image with the given ID.
func (s ImageService) FlattenImageSnapshot(ctx context.Context, id, snapshot int) (*hc.FlattenImageSnapshotResponse, error) {
	var resp hc.FlattenImageSnapshotResponse

	p := path(hc.ImagePath, id, "snapshot", snapshot, "flatten")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RevertImageSnapshot reverts the snapshot of the image with the given ID.
func (s ImageService) RevertImageSnapshot(ctx context.Context, id, snapshot int) (*hc.RevertImageSnapshotResponse, error) {
	var resp hc.RevertImageSnapshotResponse

	p := path(hc.ImagePath, id, "snapshot", snapshot, "revert")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeImageType changes the type of the image with the given ID.
func (s ImageService) ChangeImageType(ctx context.Context, id int, req hc.ChangeImageTypeRequest) (*hc.ChangeImageTypeResponse, error) {
	var resp hc.ChangeImageTypeResponse

	p := path(hc.ImagePath, id, "type")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockImage unlocks the image with the given ID.
func (s ImageService) UnlockImage(ctx context.Context, id int) (*hc.UnlockImageResponse, error) {
	var resp hc.UnlockImageResponse

	p := path(hc.ImagePath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateImage creates a new image.
func (s ImageService) CreateImage(ctx context.Context, req hc.CreateImageRequest) (*hc.CreateImageResponse, error) {
	var resp hc.CreateImageResponse

	p := path(hc.ImagePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneImage clones the image with the given ID.
func (s ImageService) CloneImage(ctx context.Context, id int, req hc.CloneImageRequest) (*hc.CloneImageResponse, error) {
	var resp hc.CloneImageResponse

	p := path(hc.ImagePath, id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
