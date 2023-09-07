package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/storage"
)

// StorageService owns the /storage methods.
type StorageService struct {
	*client.Client
}

// StorageHosts returns information about all storage hosts.
func (s StorageService) StorageHosts(ctx context.Context) (*storage.ListHostsResponse, error) {
	var resp storage.ListHostsResponse

	p := path(storage.PathPrefix, storage.HostPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageHost returns information about a specific storage host.
func (s StorageService) StorageHost(ctx context.Context, host string) (*storage.ListHostResponse, error) {
	var resp storage.ListHostResponse

	p := path(storage.PathPrefix, storage.HostPath, host)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageHostDisks returns information about all disks on a specific storage host.
func (s StorageService) StorageHostDisks(ctx context.Context, host string) (*storage.ListHostDisksResponse, error) {
	var resp storage.ListHostDisksResponse

	p := path(storage.PathPrefix, storage.HostPath, host, "disk")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageHostDisk returns information about a specific disk on a specific storage host.
func (s StorageService) StorageHostDisk(ctx context.Context, host, disk string) (*storage.ListHostDiskResponse, error) {
	var resp storage.ListHostDiskResponse

	p := path(storage.PathPrefix, storage.HostPath, host, "disk", disk)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageDisks returns information about all disks.
func (s StorageService) StorageDisks(ctx context.Context) (*storage.ListDisksResponse, error) {
	var resp storage.ListDisksResponse

	p := path(storage.PathPrefix, "disk")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageHostVolumeGroups returns information about all volume groups on a specific storage host.
func (s StorageService) StorageHostVolumeGroups(ctx context.Context, host string) (*storage.ListHostVolumeGroupsResponse, error) {
	var resp storage.ListHostVolumeGroupsResponse

	p := path(storage.PathPrefix, storage.HostPath, host, "volgroup")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageHostVolumeGroup returns information about a specific volume group on a specific storage host.
func (s StorageService) StorageHostVolumeGroup(ctx context.Context, host, vg string) (*storage.ListHostVolumeGroupResponse, error) {
	var resp storage.ListHostVolumeGroupResponse

	p := path(storage.PathPrefix, storage.HostPath, host, "volgroup", vg)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageShareRBD returns information about a specific RBD.
func (s StorageService) StorageShareRBD(ctx context.Context, rbd string) (*storage.ListShareRBDResponse, error) {
	var resp storage.ListShareRBDResponse

	p := path(storage.PathPrefix, storage.SharePath, "rbd", rbd)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageShareRBDs returns information about all RBDs.
func (s StorageService) StorageShareRBDs(ctx context.Context) (*storage.ListShareRBDsResponse, error) {
	var resp storage.ListShareRBDsResponse

	p := path(storage.PathPrefix, storage.SharePath, "rbd")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageSummary returns summary information about all storage.
func (s StorageService) StorageSummary(ctx context.Context) (*storage.ListSummaryResponse, error) {
	var resp storage.ListSummaryResponse

	p := path(storage.PathPrefix, "summary")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StoragePools returns information about all storage pools.
func (s StorageService) StoragePools(ctx context.Context) (*storage.ListPoolsResponse, error) {
	var resp storage.ListPoolsResponse

	p := path(storage.PathPrefix, "pool")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StoragePool returns information about a specific storage pool.
func (s StorageService) StoragePool(ctx context.Context, pool uint) (*storage.ListPoolResponse, error) {
	var resp storage.ListPoolResponse

	p := path(storage.PathPrefix, "pool", pool)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageOSDs returns information about all OSDs.
func (s StorageService) StorageOSDs(ctx context.Context) (*storage.ListOSDsResponse, error) {
	var resp storage.ListOSDsResponse

	p := path(storage.PathPrefix, "osd")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageOSD returns information about a specific OSD.
func (s StorageService) StorageOSD(ctx context.Context, osd uint) (*storage.ListOSDResponse, error) {
	var resp storage.ListOSDResponse

	p := path(storage.PathPrefix, "osd", osd)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
