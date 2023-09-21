package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/metal"
)

// MetalService owns the /metal methods.
type MetalService struct {
	*client.Client
}

// MetalHosts returns information about all storage hosts.
func (s MetalService) MetalHosts(ctx context.Context) (*metal.ListHostsResponse, error) {
	var resp metal.ListHostsResponse

	p := path(metal.PathPrefix, metal.HostPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MetalHost returns information about a specific storage host.
func (s MetalService) MetalHost(ctx context.Context, host string) (*metal.ListHostResponse, error) {
	var resp metal.ListHostResponse

	p := path(metal.PathPrefix, metal.HostPath, host)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MetalHostDisks returns information about all disks on a specific storage host.
func (s MetalService) MetalHostDisks(ctx context.Context, host string) (*metal.ListHostDisksResponse, error) {
	var resp metal.ListHostDisksResponse

	p := path(metal.PathPrefix, metal.HostPath, host, "disk")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MetalHostDisk returns information about a specific disk on a specific storage host.
func (s MetalService) MetalHostDisk(ctx context.Context, host, disk string) (*metal.ListHostDiskResponse, error) {
	var resp metal.ListHostDiskResponse

	p := path(metal.PathPrefix, metal.HostPath, host, "disk", disk)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MetalDisks returns information about all disks.
func (s MetalService) MetalDisks(ctx context.Context) (*metal.ListDisksResponse, error) {
	var resp metal.ListDisksResponse

	p := path(metal.PathPrefix, "disk")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MetalHostVolumeGroups returns information about all volume groups on a specific storage host.
func (s MetalService) MetalHostVolumeGroups(ctx context.Context, host string) (*metal.ListHostVolumeGroupsResponse, error) {
	var resp metal.ListHostVolumeGroupsResponse

	p := path(metal.PathPrefix, metal.HostPath, host, "volgroup")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MetalHostVolumeGroup returns information about a specific volume group on a specific storage host.
func (s MetalService) MetalHostVolumeGroup(ctx context.Context, host, vg string) (*metal.ListHostVolumeGroupResponse, error) {
	var resp metal.ListHostVolumeGroupResponse

	p := path(metal.PathPrefix, metal.HostPath, host, "volgroup", vg)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MetalShareRBD returns information about a specific RBD.
func (s MetalService) MetalShareRBD(ctx context.Context, rbd string) (*metal.ListShareRBDResponse, error) {
	var resp metal.ListShareRBDResponse

	p := path(metal.PathPrefix, metal.SharePath, "rbd", rbd)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MetalShareRBDs returns information about all RBDs.
func (s MetalService) MetalShareRBDs(ctx context.Context) (*metal.ListShareRBDsResponse, error) {
	var resp metal.ListShareRBDsResponse

	p := path(metal.PathPrefix, metal.SharePath, "rbd")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MetalStorageSummary returns summary information about all metal.
func (s MetalService) MetalStorageSummary(ctx context.Context) (*metal.ListSummaryResponse, error) {
	var resp metal.ListSummaryResponse

	p := path(metal.PathPrefix, "summary")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MetalStoragePools returns information about all storage pools.
func (s MetalService) MetalStoragePools(ctx context.Context) (*metal.ListPoolsResponse, error) {
	var resp metal.ListPoolsResponse

	p := path(metal.PathPrefix, "pool")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MetalStoragePool returns information about a specific storage pool.
func (s MetalService) MetalStoragePool(ctx context.Context, pool uint) (*metal.ListPoolResponse, error) {
	var resp metal.ListPoolResponse

	p := path(metal.PathPrefix, "pool", pool)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MetalStorageOSDs returns information about all OSDs.
func (s MetalService) MetalStorageOSDs(ctx context.Context) (*metal.ListOSDsResponse, error) {
	var resp metal.ListOSDsResponse

	p := path(metal.PathPrefix, "osd")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MetalStorageOSD returns information about a specific OSD.
func (s MetalService) MetalStorageOSD(ctx context.Context, osd uint) (*metal.ListOSDResponse, error) {
	var resp metal.ListOSDResponse

	p := path(metal.PathPrefix, "osd", osd)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
