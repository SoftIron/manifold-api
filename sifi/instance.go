package sifi

import (
	"context"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
	"github.com/softiron/hypercloud-api/cloud/instance"
)

// InstanceService owns the /instance and /instance-group methods.
type InstanceService struct {
	*client.Client
}

// DeleteInstanceDisk deletes a disk from the instance with the given id.
func (s InstanceService) DeleteInstanceDisk(ctx context.Context, id, disk int) error {
	p := path(cloud.InstancePath, id, "disk", disk)

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceDiskSnapshot deletes a snapshot from the instance with the given id.
func (s InstanceService) DeleteInstanceDiskSnapshot(ctx context.Context, id, disk, snapshot int) error {
	p := path(cloud.InstancePath, id, "disk", disk, "snapshot", snapshot)

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceNIC deletes a NIC from the instance with the given id.
func (s InstanceService) DeleteInstanceNIC(ctx context.Context, id, nic int) error {
	p := path(cloud.InstancePath, id, "nic", nic)

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceSecurityGroup deletes a security group from the instance with the given id.
func (s InstanceService) DeleteInstanceSecurityGroup(ctx context.Context, id, nic int) error {
	p := path(cloud.InstancePath, id, "nic", nic, "security-group")

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceSchedule deletes a schedule from the instance with the given id.
func (s InstanceService) DeleteInstanceSchedule(ctx context.Context, id, schedule int) error {
	p := path(cloud.InstancePath, id, "schedule", schedule)

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceSnapshot deletes a snapshot from the instance with the given id.
func (s InstanceService) DeleteInstanceSnapshot(ctx context.Context, id, snapshot int) error {
	p := path(cloud.InstancePath, id, "snapshot", snapshot)

	return s.Delete(ctx, p, nil)
}

// Instance returns the Instance for the given id.
func (s InstanceService) Instance(ctx context.Context, id int) (*cloud.ListInstanceResponse, error) {
	var resp cloud.ListInstanceResponse

	p := path(cloud.InstancePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstancesSet returns a slice of Instances.
func (s InstanceService) InstancesSet(ctx context.Context, ids ...int) (*cloud.ListInstancesResponse, error) {
	var resp cloud.ListInstancesResponse

	p, err := url.ParseQuery(path(cloud.InstancePath))
	if err != nil {
		return nil, err
	}

	if len(ids) > 0 {
		s := make([]string, len(ids))
		for i := range ids {
			s[i] = strconv.Itoa(ids[i])
		}

		p.Add(cloud.InstancePath, strings.Join(s, ","))
	}

	if err := s.Get(ctx, p.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Instances returns a slice of all the Instances in the given state.
func (s InstanceService) Instances(ctx context.Context, state instance.State, extended bool) (*cloud.ListInstancesResponse, error) {
	var resp cloud.ListInstancesResponse

	p := path(cloud.InstancePath)

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
func (s InstanceService) InstanceMonitoringData(ctx context.Context, id int) (*cloud.ListInstanceMonitoringResponse, error) {
	var resp cloud.ListInstanceMonitoringResponse

	p := path(cloud.InstancePath, "monitoring", id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstancesMonitoringData returns monitoring data for a set of instances.
func (s InstanceService) InstancesMonitoringData(ctx context.Context, filter cloud.Filter, seconds int) (*cloud.ListInstancesMonitoringResponse, error) {

	var resp cloud.ListInstancesMonitoringResponse

	p := path(cloud.InstancePath, "monitoring")

	q := make(url.Values)
	q.Add("filter", filter.String())
	q.Add("seconds", strconv.Itoa(seconds))

	if err := s.Get(ctx, p+"?"+q.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstancesAccountingData returns accounting data for a set of instances.
func (s InstanceService) InstancesAccountingData(ctx context.Context, filter cloud.Filter, start, end time.Time) (*cloud.ListInstancesAccountingResponse, error) {

	var resp cloud.ListInstancesAccountingResponse

	p := path(cloud.InstancePath, "accounting")

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
func (s InstanceService) InstanceShowback(ctx context.Context, filter cloud.Filter, monthStart, monthEnd time.Month, yearStart, yearEnd int) (*cloud.ListInstancesShowbackResponse, error) {

	var resp cloud.ListInstancesShowbackResponse

	p := path(cloud.InstancePath, "showback")

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
func (s InstanceService) SetInstanceAction(ctx context.Context, id int, req cloud.SetInstanceActionRequest) (*cloud.SetInstanceActionResponse, error) {
	var resp cloud.SetInstanceActionResponse

	p := path(cloud.InstancePath, id, "action")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateInstanceConfig updates the config for the instance with the given id.
func (s InstanceService) UpdateInstanceConfig(ctx context.Context, id int, req cloud.UpdateInstanceConfigRequest) (*cloud.UpdateInstanceConfigResponse, error) {
	var resp cloud.UpdateInstanceConfigResponse

	p := path(cloud.InstancePath, id, "attributes")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// DeployInstance deploys the instance with the given id.
func (s InstanceService) DeployInstance(ctx context.Context, id int, req cloud.DeployInstanceRequest) (*cloud.DeployInstanceResponse, error) {
	var resp cloud.DeployInstanceResponse

	p := path(cloud.InstancePath, id, "deploy")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameInstanceDiskSnapshot renames the disk snapshot for the instance with the given id.
func (s InstanceService) RenameInstanceDiskSnapshot(ctx context.Context, id, disk, snapshot int, req cloud.RenameInstanceDiskSnapshotRequest) (*cloud.RenameInstanceDiskSnapshotResponse, error) {
	var resp cloud.RenameInstanceDiskSnapshotResponse

	p := path(cloud.InstancePath, id, "disk", disk, "snapshot", snapshot, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RevertInstanceDiskSnapshot reverts the disk snapshot for the instance with the given id.
func (s InstanceService) RevertInstanceDiskSnapshot(ctx context.Context, id, disk, snapshot int) (*cloud.RevertInstanceDiskSnapshotResponse, error) {
	var resp cloud.RevertInstanceDiskSnapshotResponse

	p := path(cloud.InstancePath, id, "disk", disk, "snapshot", snapshot, "revert")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ResizeInstanceDisk resizes the disk for the instance with the given id.
func (s InstanceService) ResizeInstanceDisk(ctx context.Context, id, disk int, req cloud.ResizeInstanceDiskRequest) (*cloud.ResizeInstanceDiskResponse, error) {
	var resp cloud.ResizeInstanceDiskResponse

	p := path(cloud.InstancePath, id, "disk", disk, "size")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockInstance locks the instance with the given id.
func (s InstanceService) LockInstance(ctx context.Context, id int, req cloud.LockInstanceRequest) (*cloud.LockInstanceResponse, error) {
	var resp cloud.LockInstanceResponse

	p := path(cloud.InstancePath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MoveInstance moves the instance with the given id.
func (s InstanceService) MoveInstance(ctx context.Context, id int, req cloud.MoveInstanceRequest) (*cloud.MoveInstanceResponse, error) {
	var resp cloud.MoveInstanceResponse

	p := path(cloud.InstancePath, id, "move")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameInstance renames the instance with the given id.
func (s InstanceService) RenameInstance(ctx context.Context, id int, req cloud.RenameInstanceRequest) (*cloud.RenameInstanceResponse, error) {
	var resp cloud.RenameInstanceResponse

	p := path(cloud.InstancePath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeInstanceOwnership changes the ownership of the instance with the given id.
func (s InstanceService) ChangeInstanceOwnership(ctx context.Context, id int, req cloud.ChangeInstanceOwnershipRequest) (*cloud.ChangeInstanceOwnershipResponse, error) {
	var resp cloud.ChangeInstanceOwnershipResponse

	p := path(cloud.InstancePath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeInstancePermissions changes the permissions of the instance with the given id.
func (s InstanceService) ChangeInstancePermissions(ctx context.Context, id int, req cloud.ChangeInstancePermissionsRequest) (*cloud.ChangeInstancePermissionsResponse, error) {
	var resp cloud.ChangeInstancePermissionsResponse

	p := path(cloud.InstancePath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RecoverInstance recovers the instance with the given id.
func (s InstanceService) RecoverInstance(ctx context.Context, id int, req cloud.RecoverInstanceRequest) (*cloud.RecoverInstanceResponse, error) {
	var resp cloud.RecoverInstanceResponse

	p := path(cloud.InstancePath, id, "recover")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateInstanceSchedule updates the schedule for the instance with the given id.
func (s InstanceService) UpdateInstanceSchedule(ctx context.Context, id int, req cloud.UpdateInstanceScheduleRequest) (*cloud.UpdateInstanceScheduleResponse, error) {
	var resp cloud.UpdateInstanceScheduleResponse

	p := path(cloud.InstancePath, id, "schedule")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ResizeInstance resizes the instance with the given id.
func (s InstanceService) ResizeInstance(ctx context.Context, id int, req cloud.ResizeInstanceRequest) (*cloud.ResizeInstanceResponse, error) {
	var resp cloud.ResizeInstanceResponse

	p := path(cloud.InstancePath, id, "size")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RevertInstanceSnapshot reverts the snapshot for the instance with the given id.
func (s InstanceService) RevertInstanceSnapshot(ctx context.Context, id, snapshot int) (*cloud.RevertInstanceSnapshotResponse, error) {
	var resp cloud.RevertInstanceSnapshotResponse

	p := path(cloud.InstancePath, id, "snapshot", snapshot, "revert")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateInstanceTemplate updates the template for the instance with the given id.
func (s InstanceService) UpdateInstanceTemplate(ctx context.Context, id int, req cloud.UpdateInstanceTemplateRequest) (*cloud.UpdateInstanceTemplateResponse, error) {
	var resp cloud.UpdateInstanceTemplateResponse

	p := path(cloud.InstancePath, id, "template")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockInstance unlocks the instance with the given id.
func (s InstanceService) UnlockInstance(ctx context.Context, id int) (*cloud.UnlockInstanceResponse, error) {
	var resp cloud.UnlockInstanceResponse

	p := path(cloud.InstancePath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstance creates a new instance.
func (s InstanceService) CreateInstance(ctx context.Context, req cloud.CreateInstanceRequest) (*cloud.CreateInstanceResponse, error) {
	var resp cloud.CreateInstanceResponse

	if err := s.Post(ctx, cloud.InstancePath, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CalculateInstancesShowback calculates the showback for all the instances.
func (s InstanceService) CalculateInstancesShowback(ctx context.Context, req cloud.CalculateInstancesShowbackRequest) error {
	p := path(cloud.InstancePath, "showback")

	return s.Post(ctx, p, req, nil)
}

// CreateInstanceDisk creates a new disk for the instance with the given id.
func (s InstanceService) CreateInstanceDisk(ctx context.Context, id int, req cloud.CreateInstanceDiskRequest) (*cloud.CreateInstanceDiskResponse, error) {
	var resp cloud.CreateInstanceDiskResponse

	p := path(cloud.InstancePath, id, "disk")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstanceDiskImage creates a new disk image for the instance with the given id.
func (s InstanceService) CreateInstanceDiskImage(ctx context.Context, id, disk int, req cloud.CreateInstanceDiskImageRequest) (*cloud.CreateInstanceDiskImageResponse, error) {
	var resp cloud.CreateInstanceDiskImageResponse

	p := path(cloud.InstancePath, id, "disk", disk, "image")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstanceDiskSnapshot creates a new disk snapshot for the instance with the given id.
func (s InstanceService) CreateInstanceDiskSnapshot(ctx context.Context, id, disk int, req cloud.CreateInstanceDiskSnapshotRequest) (*cloud.CreateInstanceDiskSnapshotResponse, error) {
	var resp cloud.CreateInstanceDiskSnapshotResponse

	p := path(cloud.InstancePath, id, "disk", disk, "snapshot")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstanceNIC creates a new NIC for the instance with the given id.
func (s InstanceService) CreateInstanceNIC(ctx context.Context, id int, req cloud.CreateInstanceNICRequest) (*cloud.CreateInstanceNICResponse, error) {
	var resp cloud.CreateInstanceNICResponse

	p := path(cloud.InstancePath, id, "nic")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddInstanceSecurityGroup adds a security group to the instance with the given id.
func (s InstanceService) AddInstanceSecurityGroup(ctx context.Context, id, nic, sg int) (*cloud.AddInstanceSecurityGroupResponse, error) {
	var resp cloud.AddInstanceSecurityGroupResponse

	p := path(cloud.InstancePath, id, "nic", nic, "security-group", sg)

	if err := s.Post(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddInstanceSchedule adds a schedule to the instance with the given id.
func (s InstanceService) AddInstanceSchedule(ctx context.Context, id int, req cloud.AddInstanceScheduleRequest) (*cloud.AddInstanceScheduleResponse, error) {
	var resp cloud.AddInstanceScheduleResponse

	p := path(cloud.InstancePath, id, "schedule")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstanceSnapshot creates a new snapshot for the instance with the given id.
func (s InstanceService) CreateInstanceSnapshot(ctx context.Context, id int, req cloud.CreateInstanceSnapshotRequest) (*cloud.CreateInstanceSnapshotResponse, error) {
	var resp cloud.CreateInstanceSnapshotResponse

	p := path(cloud.InstancePath, id, "snapshot")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateVNCProxy creates a new VNC proxy on the server for a specific VM.
func (s InstanceService) CreateVNCProxy(ctx context.Context, id int, req cloud.CreateVNCProxyRequest) (*cloud.CreateVNCProxyResponse, error) {
	var resp cloud.CreateVNCProxyResponse

	p := path(cloud.InstancePath, id, "vnc")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
