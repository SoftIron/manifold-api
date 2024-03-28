package metal

import (
	"context"
	"fmt"
	"strings"

	"github.com/softiron/manifold-api/client"
)

// Service owns the /metal methods.
type Service struct {
	*client.Client
	root string
}

// NewService returns a new Service for bare-metal operations.
func NewService(c *client.Client, path string) *Service {
	return &Service{Client: c, root: path}
}

// path returns a URL path with the correct prefix appended.
func (s Service) path(dirs ...interface{}) string {
	p := make([]string, 2, len(dirs)+1)
	p[0] = s.root
	p[1] = PathPrefix

	for _, d := range dirs {
		p = append(p, fmt.Sprint(d))
	}

	return strings.Join(p, "/")
}

// Hosts returns information about all the hosts in the cluster.
func (s Service) Hosts(ctx context.Context) (*HostsResponse, error) {
	var resp HostsResponse

	p := s.path(HostPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Host returns information about a specific host.
func (s Service) Host(ctx context.Context, host string) (*HostResponse, error) {
	var resp HostResponse

	p := s.path(HostPath, host)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HostDisks returns information about all disks on a specific storage host.
func (s Service) HostDisks(ctx context.Context, host string) (*HostDisksResponse, error) {
	var resp HostDisksResponse

	p := s.path(HostPath, host, "disk")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HostDisk returns information about a specific disk on a specific storage host.
func (s Service) HostDisk(ctx context.Context, host, disk string) (*HostDiskResponse, error) {
	var resp HostDiskResponse

	p := s.path(HostPath, host, "disk", disk)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HostVolumeGroups returns information about all volume groups on a specific storage host.
func (s Service) HostVolumeGroups(ctx context.Context, host string) (*HostVolumeGroupsResponse, error) {
	var resp HostVolumeGroupsResponse

	p := s.path(HostPath, host, "volgroup")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// HostVolumeGroup returns information about a specific volume group on a specific storage host.
func (s Service) HostVolumeGroup(ctx context.Context, host, vg string) (*HostVolumeGroupResponse, error) {
	var resp HostVolumeGroupResponse

	p := s.path(HostPath, host, "volgroup", vg)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ShareRBD returns information about a specific RBD.
func (s Service) ShareRBD(ctx context.Context, rbd string) (*ShareRBDResponse, error) {
	var resp ShareRBDResponse

	p := s.path(SharePath, "rbd", rbd)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ShareRBDs returns information about all RBDs.
func (s Service) ShareRBDs(ctx context.Context) (*ShareRBDsResponse, error) {
	var resp ShareRBDsResponse

	p := s.path(SharePath, "rbd")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageSummary returns summary information about all
func (s Service) StorageSummary(ctx context.Context) (*SummaryResponse, error) {
	var resp SummaryResponse

	p := s.path("summary")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StoragePools returns information about all storage pools.
func (s Service) StoragePools(ctx context.Context) (*StoragePoolsResponse, error) {
	var resp StoragePoolsResponse

	p := s.path(PoolPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StoragePool returns information about a specific storage pool.
func (s Service) StoragePool(ctx context.Context, pool uint) (*StoragePoolResponse, error) {
	var resp StoragePoolResponse

	p := s.path(PoolPath, pool)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageOSDs returns information about all OSDs.
func (s Service) StorageOSDs(ctx context.Context) (*StorageOSDsResponse, error) {
	var resp StorageOSDsResponse

	p := s.path(OSDPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// StorageOSD returns information about a specific OSD.
func (s Service) StorageOSD(ctx context.Context, osd uint) (*StorageOSDResponse, error) {
	var resp StorageOSDResponse

	p := s.path(OSDPath, osd)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SystemCapacity returns total system capacity and allocations.
func (s Service) SystemCapacity(ctx context.Context) (*TotalCapacity, error) {
	var resp TotalCapacity

	p := s.path(CapacityPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
