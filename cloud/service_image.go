package cloud

import (
	"context"
)

// ImageService owns the /cloud/image methods.
type ImageService struct {
	*service
}

// DeleteImage deletes the image with the given ID.
func (s ImageService) DeleteImage(ctx context.Context, id int) error {
	p := s.path(ImagePath, id)

	return s.Delete(ctx, p, nil)
}

// DeleteImageSnapshot deletes a snapshot for the image with the given ID.
func (s ImageService) DeleteImageSnapshot(ctx context.Context, id, snapshot int) error {
	p := s.path(ImagePath, id, "snapshot", snapshot)

	return s.Delete(ctx, p, nil)
}

// Image returns information about a image.
func (s ImageService) Image(ctx context.Context, id int) (*ImageResponse, error) {
	var resp ImageResponse

	p := s.path(ImagePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Images returns information about all images.
func (s ImageService) Images(ctx context.Context) (*ImagesResponse, error) {
	var resp ImagesResponse

	p := s.path(ImagePath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateImage updates the image with the given ID.
func (s ImageService) UpdateImage(ctx context.Context, id int, req UpdateImageRequest) (*UpdateImageResponse, error) {
	var resp UpdateImageResponse

	p := s.path(ImagePath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableImage enables the image with the given ID.
func (s ImageService) EnableImage(ctx context.Context, id int, req EnableImageRequest) (*EnableImageResponse, error) {
	var resp EnableImageResponse

	p := s.path(ImagePath, id, "enable")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockImage locks the image with the given ID.
func (s ImageService) LockImage(ctx context.Context, id int, req LockImageRequest) (*LockImageResponse, error) {
	var resp LockImageResponse

	p := s.path(ImagePath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameImage renames the image with the given ID.
func (s ImageService) RenameImage(ctx context.Context, id int, req RenameImageRequest) (*RenameImageResponse, error) {
	var resp RenameImageResponse

	p := s.path(ImagePath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeImageOwnership changes the ownership of the image with the given ID.
func (s ImageService) ChangeImageOwnership(ctx context.Context, id int, req ChangeImageOwnershipRequest) (*ChangeImageOwnershipResponse, error) {
	var resp ChangeImageOwnershipResponse

	p := s.path(ImagePath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeImagePermissions changes the permissions of the image with the given ID.
func (s ImageService) ChangeImagePermissions(ctx context.Context, id int, req ChangeImagePermissionsRequest) (*ChangeImagePermissionsResponse, error) {
	var resp ChangeImagePermissionsResponse

	p := s.path(ImagePath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetImagePersistent sets the image with the given ID to be persistent.
func (s ImageService) SetImagePersistent(ctx context.Context, id int, req SetImagePersistentRequest) (*SetImagePersistentResponse, error) {
	var resp SetImagePersistentResponse

	p := s.path(ImagePath, id, "persistent")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// FlattenImageSnapshot flattens the snapshot of the image with the given ID.
func (s ImageService) FlattenImageSnapshot(ctx context.Context, id, snapshot int) (*FlattenImageSnapshotResponse, error) {
	var resp FlattenImageSnapshotResponse

	p := s.path(ImagePath, id, "snapshot", snapshot, "flatten")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RevertImageSnapshot reverts the snapshot of the image with the given ID.
func (s ImageService) RevertImageSnapshot(ctx context.Context, id, snapshot int) (*RevertImageSnapshotResponse, error) {
	var resp RevertImageSnapshotResponse

	p := s.path(ImagePath, id, "snapshot", snapshot, "revert")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeImageType changes the type of the image with the given ID.
func (s ImageService) ChangeImageType(ctx context.Context, id int, req ChangeImageTypeRequest) (*ChangeImageTypeResponse, error) {
	var resp ChangeImageTypeResponse

	p := s.path(ImagePath, id, "type")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockImage unlocks the image with the given ID.
func (s ImageService) UnlockImage(ctx context.Context, id int) (*UnlockImageResponse, error) {
	var resp UnlockImageResponse

	p := s.path(ImagePath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateImage creates a new image.
func (s ImageService) CreateImage(ctx context.Context, req CreateImageRequest) (*CreateImageResponse, error) {
	var resp CreateImageResponse

	p := s.path(ImagePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CloneImage clones the image with the given ID.
func (s ImageService) CloneImage(ctx context.Context, id int, req CloneImageRequest) (*CloneImageResponse, error) {
	var resp CloneImageResponse

	p := s.path(ImagePath, id, "clone")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
