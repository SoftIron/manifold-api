package cloud

import (
	"context"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/softiron/manifold-api/cloud/instance"
)

// InstanceService owns the /cloud/instance and /cloud/instance-group methods.
type InstanceService struct {
	*service
}

// DeleteInstanceDisk deletes a disk from the instance with the given id.
func (s InstanceService) DeleteInstanceDisk(ctx context.Context, id, disk int) error {
	p := s.path(InstancePath, id, "disk", disk)

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceDiskSnapshot deletes a snapshot from the instance with the given id.
func (s InstanceService) DeleteInstanceDiskSnapshot(ctx context.Context, id, disk, snapshot int) error {
	p := s.path(InstancePath, id, "disk", disk, "snapshot", snapshot)

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceNIC deletes a NIC from the instance with the given id.
func (s InstanceService) DeleteInstanceNIC(ctx context.Context, id, nic int) error {
	p := s.path(InstancePath, id, "nic", nic)

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceSecurityGroup deletes a security group from the instance with the given id.
func (s InstanceService) DeleteInstanceSecurityGroup(ctx context.Context, id, nic int) error {
	p := s.path(InstancePath, id, "nic", nic, "security-group")

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceSchedule deletes a schedule from the instance with the given id.
func (s InstanceService) DeleteInstanceSchedule(ctx context.Context, id, schedule int) error {
	p := s.path(InstancePath, id, "schedule", schedule)

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceSnapshot deletes a snapshot from the instance with the given id.
func (s InstanceService) DeleteInstanceSnapshot(ctx context.Context, id, snapshot int) error {
	p := s.path(InstancePath, id, "snapshot", snapshot)

	return s.Delete(ctx, p, nil)
}

// Instance returns the Instance for the given id.
func (s InstanceService) Instance(ctx context.Context, id int) (*InstanceResponse, error) {
	var resp InstanceResponse

	p := s.path(InstancePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstancesSet returns a slice of Instances.
func (s InstanceService) InstancesSet(ctx context.Context, ids ...int) (*InstancesResponse, error) {
	var resp InstancesResponse

	p, err := url.ParseQuery(s.path(InstancePath))
	if err != nil {
		return nil, err
	}

	if len(ids) > 0 {
		s := make([]string, len(ids))
		for i := range ids {
			s[i] = strconv.Itoa(ids[i])
		}

		p.Add(InstancePath, strings.Join(s, ","))
	}

	if err := s.Get(ctx, p.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Instances returns a slice of all the Instances in the given state.
func (s InstanceService) Instances(ctx context.Context, state instance.State, extended bool) (*InstancesResponse, error) {
	var resp InstancesResponse

	p := s.path(InstancePath)

	q := make(url.Values)
	q.Add("state", state.String())

	if extended {
		q.Add("extended", strconv.FormatBool(extended))
	}

	if err := s.Get(ctx, p+"?"+q.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstanceMonitoringData returns monitoring data for id.
func (s InstanceService) InstanceMonitoringData(ctx context.Context, id int) (*InstanceMonitoringResponse, error) {
	var resp InstanceMonitoringResponse

	p := s.path(InstancePath, "monitoring", id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstancesMonitoringData returns monitoring data for a set of instances.
func (s InstanceService) InstancesMonitoringData(ctx context.Context, filter Filter, seconds int) (*InstancesMonitoringResponse, error) {

	var resp InstancesMonitoringResponse

	p := s.path(InstancePath, "monitoring")

	q := make(url.Values)
	q.Add("filter", filter.String())
	q.Add("seconds", strconv.Itoa(seconds))

	if err := s.Get(ctx, p+"?"+q.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstancesAccountingData returns accounting data for a set of instances.
func (s InstanceService) InstancesAccountingData(ctx context.Context, filter Filter, start, end time.Time) (*InstancesAccountingResponse, error) {

	var resp InstancesAccountingResponse

	p := s.path(InstancePath, "accounting")

	q := make(url.Values)
	q.Add("filter", filter.String())
	q.Add("start", strconv.FormatInt(start.Unix(), 10))
	q.Add("end", strconv.FormatInt(end.Unix(), 10))

	if err := s.Get(ctx, p+"?"+q.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstanceShowback returns showback data for a set of instances.
func (s InstanceService) InstanceShowback(ctx context.Context, filter Filter, monthStart, monthEnd time.Month, yearStart, yearEnd int) (*InstancesShowbackResponse, error) {

	var resp InstancesShowbackResponse

	p := s.path(InstancePath, "showback")

	q := make(url.Values)
	q.Add("filter", filter.String())
	q.Add("month-start", strconv.Itoa(int(monthStart)))
	q.Add("month-end", strconv.Itoa(int(monthEnd)))
	q.Add("year-start", strconv.Itoa(yearStart))
	q.Add("year-start", strconv.Itoa(yearEnd))

	if err := s.Get(ctx, p+"?"+q.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// SetInstanceAction sets the action for the instance with the given id.
func (s InstanceService) SetInstanceAction(ctx context.Context, id int, req SetInstanceActionRequest) (*SetInstanceActionResponse, error) {
	var resp SetInstanceActionResponse

	p := s.path(InstancePath, id, "action")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateInstanceConfig updates the config for the instance with the given id.
func (s InstanceService) UpdateInstanceConfig(ctx context.Context, id int, req UpdateInstanceConfigRequest) (*UpdateInstanceConfigResponse, error) {
	var resp UpdateInstanceConfigResponse

	p := s.path(InstancePath, id, "attributes")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// DeployInstance deploys the instance with the given id.
func (s InstanceService) DeployInstance(ctx context.Context, id int, req DeployInstanceRequest) (*DeployInstanceResponse, error) {
	var resp DeployInstanceResponse

	p := s.path(InstancePath, id, "deploy")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameInstanceDiskSnapshot renames the disk snapshot for the instance with the given id.
func (s InstanceService) RenameInstanceDiskSnapshot(ctx context.Context, id, disk, snapshot int, req RenameInstanceDiskSnapshotRequest) (*RenameInstanceDiskSnapshotResponse, error) {
	var resp RenameInstanceDiskSnapshotResponse

	p := s.path(InstancePath, id, "disk", disk, "snapshot", snapshot, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RevertInstanceDiskSnapshot reverts the disk snapshot for the instance with the given id.
func (s InstanceService) RevertInstanceDiskSnapshot(ctx context.Context, id, disk, snapshot int) (*RevertInstanceDiskSnapshotResponse, error) {
	var resp RevertInstanceDiskSnapshotResponse

	p := s.path(InstancePath, id, "disk", disk, "snapshot", snapshot, "revert")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ResizeInstanceDisk resizes the disk for the instance with the given id.
func (s InstanceService) ResizeInstanceDisk(ctx context.Context, id, disk int, req ResizeInstanceDiskRequest) (*ResizeInstanceDiskResponse, error) {
	var resp ResizeInstanceDiskResponse

	p := s.path(InstancePath, id, "disk", disk, "size")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockInstance locks the instance with the given id.
func (s InstanceService) LockInstance(ctx context.Context, id int, req LockInstanceRequest) (*LockInstanceResponse, error) {
	var resp LockInstanceResponse

	p := s.path(InstancePath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MoveInstance moves the instance with the given id.
func (s InstanceService) MoveInstance(ctx context.Context, id int, req MoveInstanceRequest) (*MoveInstanceResponse, error) {
	var resp MoveInstanceResponse

	p := s.path(InstancePath, id, "move")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameInstance renames the instance with the given id.
func (s InstanceService) RenameInstance(ctx context.Context, id int, req RenameInstanceRequest) (*RenameInstanceResponse, error) {
	var resp RenameInstanceResponse

	p := s.path(InstancePath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeInstanceOwnership changes the ownership of the instance with the given id.
func (s InstanceService) ChangeInstanceOwnership(ctx context.Context, id int, req ChangeInstanceOwnershipRequest) (*ChangeInstanceOwnershipResponse, error) {
	var resp ChangeInstanceOwnershipResponse

	p := s.path(InstancePath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeInstancePermissions changes the permissions of the instance with the given id.
func (s InstanceService) ChangeInstancePermissions(ctx context.Context, id int, req ChangeInstancePermissionsRequest) (*ChangeInstancePermissionsResponse, error) {
	var resp ChangeInstancePermissionsResponse

	p := s.path(InstancePath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RecoverInstance recovers the instance with the given id.
func (s InstanceService) RecoverInstance(ctx context.Context, id int, req RecoverInstanceRequest) (*RecoverInstanceResponse, error) {
	var resp RecoverInstanceResponse

	p := s.path(InstancePath, id, "recover")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateInstanceSchedule updates the schedule for the instance with the given id.
func (s InstanceService) UpdateInstanceSchedule(ctx context.Context, id int, req UpdateInstanceScheduleRequest) (*UpdateInstanceScheduleResponse, error) {
	var resp UpdateInstanceScheduleResponse

	p := s.path(InstancePath, id, "schedule")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ResizeInstance resizes the instance with the given id.
func (s InstanceService) ResizeInstance(ctx context.Context, id int, req ResizeInstanceRequest) (*ResizeInstanceResponse, error) {
	var resp ResizeInstanceResponse

	p := s.path(InstancePath, id, "size")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RevertInstanceSnapshot reverts the snapshot for the instance with the given id.
func (s InstanceService) RevertInstanceSnapshot(ctx context.Context, id, snapshot int) (*RevertInstanceSnapshotResponse, error) {
	var resp RevertInstanceSnapshotResponse

	p := s.path(InstancePath, id, "snapshot", snapshot, "revert")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateInstanceTemplate updates the template for the instance with the given id.
func (s InstanceService) UpdateInstanceTemplate(ctx context.Context, id int, req UpdateInstanceTemplateRequest) (*UpdateInstanceTemplateResponse, error) {
	var resp UpdateInstanceTemplateResponse

	p := s.path(InstancePath, id, "template")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockInstance unlocks the instance with the given id.
func (s InstanceService) UnlockInstance(ctx context.Context, id int) (*UnlockInstanceResponse, error) {
	var resp UnlockInstanceResponse

	p := s.path(InstancePath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstance creates a new instance.
func (s InstanceService) CreateInstance(ctx context.Context, req CreateInstanceRequest) (*CreateInstanceResponse, error) {
	var resp CreateInstanceResponse

	p := s.path(InstancePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CalculateInstancesShowback calculates the showback for all the instances.
func (s InstanceService) CalculateInstancesShowback(ctx context.Context, req CalculateInstancesShowbackRequest) error {
	p := s.path(InstancePath, "showback")

	return s.Post(ctx, p, req, nil)
}

// CreateInstanceDisk creates a new disk for the instance with the given id.
func (s InstanceService) CreateInstanceDisk(ctx context.Context, id int, req CreateInstanceDiskRequest) (*CreateInstanceDiskResponse, error) {
	var resp CreateInstanceDiskResponse

	p := s.path(InstancePath, id, "disk")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstanceDiskImage creates a new disk image for the instance with the given id.
func (s InstanceService) CreateInstanceDiskImage(ctx context.Context, id, disk int, req CreateInstanceDiskImageRequest) (*CreateInstanceDiskImageResponse, error) {
	var resp CreateInstanceDiskImageResponse

	p := s.path(InstancePath, id, "disk", disk, "image")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstanceDiskSnapshot creates a new disk snapshot for the instance with the given id.
func (s InstanceService) CreateInstanceDiskSnapshot(ctx context.Context, id, disk int, req CreateInstanceDiskSnapshotRequest) (*CreateInstanceDiskSnapshotResponse, error) {
	var resp CreateInstanceDiskSnapshotResponse

	p := s.path(InstancePath, id, "disk", disk, "snapshot")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstanceNIC creates a new NIC for the instance with the given id.
func (s InstanceService) CreateInstanceNIC(ctx context.Context, id int, req CreateInstanceNICRequest) (*CreateInstanceNICResponse, error) {
	var resp CreateInstanceNICResponse

	p := s.path(InstancePath, id, "nic")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddInstanceSecurityGroup adds a security group to the instance with the given id.
func (s InstanceService) AddInstanceSecurityGroup(ctx context.Context, id, nic, sg int) (*AddInstanceSecurityGroupResponse, error) {
	var resp AddInstanceSecurityGroupResponse

	p := s.path(InstancePath, id, "nic", nic, "security-group", sg)

	if err := s.Post(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddInstanceSchedule adds a schedule to the instance with the given id.
func (s InstanceService) AddInstanceSchedule(ctx context.Context, id int, req AddInstanceScheduleRequest) (*AddInstanceScheduleResponse, error) {
	var resp AddInstanceScheduleResponse

	p := s.path(InstancePath, id, "schedule")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstanceSnapshot creates a new snapshot for the instance with the given id.
func (s InstanceService) CreateInstanceSnapshot(ctx context.Context, id int, req CreateInstanceSnapshotRequest) (*CreateInstanceSnapshotResponse, error) {
	var resp CreateInstanceSnapshotResponse

	p := s.path(InstancePath, id, "snapshot")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateVNCProxy creates a new VNC proxy on the server for a specific VM.
func (s InstanceService) CreateVNCProxy(ctx context.Context, id int, req CreateVNCProxyRequest) (*CreateVNCProxyResponse, error) {
	var resp CreateVNCProxyResponse

	p := s.path(InstancePath, id, "vnc")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// DeleteInstanceGroup deletes the instance group with the given ID.
func (s InstanceService) DeleteInstanceGroup(ctx context.Context, id int) error {
	p := s.path(InstanceGroupPath, id)

	return s.Delete(ctx, p, nil)
}

// InstanceGroup returns information about an instance group.
func (s InstanceService) InstanceGroup(ctx context.Context, id int) (*InstanceGroupResponse, error) {
	var resp InstanceGroupResponse

	p := s.path(InstanceGroupPath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstanceGroups returns information about all instance groups.
func (s InstanceService) InstanceGroups(ctx context.Context) (*InstanceGroupsResponse, error) {
	var resp InstanceGroupsResponse

	p := s.path(InstanceGroupPath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateInstanceGroup updates the instance group with the given ID.
func (s InstanceService) UpdateInstanceGroup(ctx context.Context, id int, req *UpdateInstanceGroupRequest) (*UpdateInstanceGroupResponse, error) {
	var resp UpdateInstanceGroupResponse

	p := s.path(InstanceGroupPath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockInstanceGroup locks the instance group with the given ID.
func (s InstanceService) LockInstanceGroup(ctx context.Context, id int, req LockInstanceGroupRequest) (*LockInstanceGroupResponse, error) {
	var resp LockInstanceGroupResponse

	p := s.path(InstanceGroupPath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameInstanceGroup renames the instance group with the given ID.
func (s InstanceService) RenameInstanceGroup(ctx context.Context, id int, req RenameInstanceGroupRequest) (*RenameInstanceGroupResponse, error) {
	var resp RenameInstanceGroupResponse

	p := s.path(InstanceGroupPath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeInstanceGroupOwnership changes the ownership of the instance group with the given ID.
func (s InstanceService) ChangeInstanceGroupOwnership(ctx context.Context, id int, req ChangeInstanceGroupOwnershipRequest) (*ChangeInstanceGroupOwnershipResponse, error) {
	var resp ChangeInstanceGroupOwnershipResponse

	p := s.path(InstanceGroupPath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeInstanceGroupPermissions changes the permissions of the instance group with the given ID.
func (s InstanceService) ChangeInstanceGroupPermissions(ctx context.Context, id int, req ChangeInstanceGroupPermissionsRequest) (*ChangeInstanceGroupPermissionsResponse, error) {
	var resp ChangeInstanceGroupPermissionsResponse

	p := s.path(InstanceGroupPath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockInstanceGroup unlocks the instance group with the given ID.
func (s InstanceService) UnlockInstanceGroup(ctx context.Context, id int) (*UnlockInstanceGroupResponse, error) {
	var resp UnlockInstanceGroupResponse
	p := s.path(InstanceGroupPath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
