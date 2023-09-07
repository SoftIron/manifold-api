package sifi

import (
	"context"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/deprecated/v1/hc"
)

// InstanceService owns the /instance and /instance-group methods.
type InstanceService struct {
	*client.Client
}

// DeleteInstanceDisk deletes a disk from the instance with the given id.
func (s InstanceService) DeleteInstanceDisk(ctx context.Context, id, disk int) error {
	p := path(hc.InstancePath, id, "disk", disk)

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceDiskSnapshot deletes a snapshot from the instance with the given id.
func (s InstanceService) DeleteInstanceDiskSnapshot(ctx context.Context, id, disk, snapshot int) error {
	p := path(hc.InstancePath, id, "disk", disk, "snapshot", snapshot)

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceNIC deletes a NIC from the instance with the given id.
func (s InstanceService) DeleteInstanceNIC(ctx context.Context, id, nic int) error {
	p := path(hc.InstancePath, id, "nic", nic)

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceSecurityGroup deletes a security group from the instance with the given id.
func (s InstanceService) DeleteInstanceSecurityGroup(ctx context.Context, id, nic int) error {
	p := path(hc.InstancePath, id, "nic", nic, "security-group")

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceSchedule deletes a schedule from the instance with the given id.
func (s InstanceService) DeleteInstanceSchedule(ctx context.Context, id, schedule int) error {
	p := path(hc.InstancePath, id, "schedule", schedule)

	return s.Delete(ctx, p, nil)
}

// DeleteInstanceSnapshot deletes a snapshot from the instance with the given id.
func (s InstanceService) DeleteInstanceSnapshot(ctx context.Context, id, snapshot int) error {
	p := path(hc.InstancePath, id, "snapshot", snapshot)

	return s.Delete(ctx, p, nil)
}

// Instance returns the Instance for the given id.
func (s InstanceService) Instance(ctx context.Context, id int) (*hc.ListInstanceResponse, error) {
	var resp hc.ListInstanceResponse

	p := path(hc.InstancePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstancesSet returns a slice of Instances.
func (s InstanceService) InstancesSet(ctx context.Context, ids ...int) (*hc.ListInstancesResponse, error) {
	var resp hc.ListInstancesResponse

	p, err := url.ParseQuery(path(hc.InstancePath))
	if err != nil {
		return nil, err
	}

	if len(ids) > 0 {
		s := make([]string, len(ids))
		for i := range ids {
			s[i] = strconv.Itoa(ids[i])
		}

		p.Add(hc.InstancePath, strings.Join(s, ","))
	}

	if err := s.Get(ctx, p.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Instances returns a slice of all the Instances in the given state.
func (s InstanceService) Instances(ctx context.Context, state hc.InstanceState, extended bool) (*hc.ListInstancesResponse, error) {
	var resp hc.ListInstancesResponse

	p := path(hc.InstancePath)

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
func (s InstanceService) InstanceMonitoringData(ctx context.Context, id int) (*hc.ListInstanceMonitoringResponse, error) {
	var resp hc.ListInstanceMonitoringResponse

	p := path(hc.InstancePath, "monitoring", id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstancesMonitoringData returns monitoring data for a set of instances.
func (s InstanceService) InstancesMonitoringData(ctx context.Context, filter hc.Filter, seconds int) (*hc.ListInstancesMonitoringResponse, error) {

	var resp hc.ListInstancesMonitoringResponse

	p := path(hc.InstancePath, "monitoring")

	q := make(url.Values)
	q.Add("filter", filter.String())
	q.Add("seconds", strconv.Itoa(seconds))

	if err := s.Get(ctx, p+"?"+q.Encode(), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// InstancesAccountingData returns accounting data for a set of instances.
func (s InstanceService) InstancesAccountingData(ctx context.Context, filter hc.Filter, start, end time.Time) (*hc.ListInstancesAccountingResponse, error) {

	var resp hc.ListInstancesAccountingResponse

	p := path(hc.InstancePath, "accounting")

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
func (s InstanceService) InstanceShowback(ctx context.Context, filter hc.Filter, monthStart, monthEnd time.Month, yearStart, yearEnd int) (*hc.ListInstancesShowbackResponse, error) {

	var resp hc.ListInstancesShowbackResponse

	p := path(hc.InstancePath, "showback")

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
func (s InstanceService) SetInstanceAction(ctx context.Context, id int, req hc.SetInstanceActionRequest) (*hc.SetInstanceActionResponse, error) {
	var resp hc.SetInstanceActionResponse

	p := path(hc.InstancePath, id, "action")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateInstanceConfig updates the config for the instance with the given id.
func (s InstanceService) UpdateInstanceConfig(ctx context.Context, id int, req hc.UpdateInstanceConfigRequest) (*hc.UpdateInstanceConfigResponse, error) {
	var resp hc.UpdateInstanceConfigResponse

	p := path(hc.InstancePath, id, "attributes")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// DeployInstance deploys the instance with the given id.
func (s InstanceService) DeployInstance(ctx context.Context, id int, req hc.DeployInstanceRequest) (*hc.DeployInstanceResponse, error) {
	var resp hc.DeployInstanceResponse

	p := path(hc.InstancePath, id, "deploy")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameInstanceDiskSnapshot renames the disk snapshot for the instance with the given id.
func (s InstanceService) RenameInstanceDiskSnapshot(ctx context.Context, id, disk, snapshot int, req hc.RenameInstanceDiskSnapshotRequest) (*hc.RenameInstanceDiskSnapshotResponse, error) {
	var resp hc.RenameInstanceDiskSnapshotResponse

	p := path(hc.InstancePath, id, "disk", disk, "snapshot", snapshot, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RevertInstanceDiskSnapshot reverts the disk snapshot for the instance with the given id.
func (s InstanceService) RevertInstanceDiskSnapshot(ctx context.Context, id, disk, snapshot int) (*hc.RevertInstanceDiskSnapshotResponse, error) {
	var resp hc.RevertInstanceDiskSnapshotResponse

	p := path(hc.InstancePath, id, "disk", disk, "snapshot", snapshot, "revert")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ResizeInstanceDisk resizes the disk for the instance with the given id.
func (s InstanceService) ResizeInstanceDisk(ctx context.Context, id, disk int, req hc.ResizeInstanceDiskRequest) (*hc.ResizeInstanceDiskResponse, error) {
	var resp hc.ResizeInstanceDiskResponse

	p := path(hc.InstancePath, id, "disk", disk, "size")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// LockInstance locks the instance with the given id.
func (s InstanceService) LockInstance(ctx context.Context, id int, req hc.LockInstanceRequest) (*hc.LockInstanceResponse, error) {
	var resp hc.LockInstanceResponse

	p := path(hc.InstancePath, id, "lock")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// MoveInstance moves the instance with the given id.
func (s InstanceService) MoveInstance(ctx context.Context, id int, req hc.MoveInstanceRequest) (*hc.MoveInstanceResponse, error) {
	var resp hc.MoveInstanceResponse

	p := path(hc.InstancePath, id, "move")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameInstance renames the instance with the given id.
func (s InstanceService) RenameInstance(ctx context.Context, id int, req hc.RenameInstanceRequest) (*hc.RenameInstanceResponse, error) {
	var resp hc.RenameInstanceResponse

	p := path(hc.InstancePath, id, "name")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeInstanceOwnership changes the ownership of the instance with the given id.
func (s InstanceService) ChangeInstanceOwnership(ctx context.Context, id int, req hc.ChangeInstanceOwnershipRequest) (*hc.ChangeInstanceOwnershipResponse, error) {
	var resp hc.ChangeInstanceOwnershipResponse

	p := path(hc.InstancePath, id, "ownership")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ChangeInstancePermissions changes the permissions of the instance with the given id.
func (s InstanceService) ChangeInstancePermissions(ctx context.Context, id int, req hc.ChangeInstancePermissionsRequest) (*hc.ChangeInstancePermissionsResponse, error) {
	var resp hc.ChangeInstancePermissionsResponse

	p := path(hc.InstancePath, id, "permissions")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RecoverInstance recovers the instance with the given id.
func (s InstanceService) RecoverInstance(ctx context.Context, id int, req hc.RecoverInstanceRequest) (*hc.RecoverInstanceResponse, error) {
	var resp hc.RecoverInstanceResponse

	p := path(hc.InstancePath, id, "recover")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateInstanceSchedule updates the schedule for the instance with the given id.
func (s InstanceService) UpdateInstanceSchedule(ctx context.Context, id int, req hc.UpdateInstanceScheduleRequest) (*hc.UpdateInstanceScheduleResponse, error) {
	var resp hc.UpdateInstanceScheduleResponse

	p := path(hc.InstancePath, id, "schedule")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ResizeInstance resizes the instance with the given id.
func (s InstanceService) ResizeInstance(ctx context.Context, id int, req hc.ResizeInstanceRequest) (*hc.ResizeInstanceResponse, error) {
	var resp hc.ResizeInstanceResponse

	p := path(hc.InstancePath, id, "size")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RevertInstanceSnapshot reverts the snapshot for the instance with the given id.
func (s InstanceService) RevertInstanceSnapshot(ctx context.Context, id, snapshot int) (*hc.RevertInstanceSnapshotResponse, error) {
	var resp hc.RevertInstanceSnapshotResponse

	p := path(hc.InstancePath, id, "snapshot", snapshot, "revert")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateInstanceTemplate updates the template for the instance with the given id.
func (s InstanceService) UpdateInstanceTemplate(ctx context.Context, id int, req hc.UpdateInstanceTemplateRequest) (*hc.UpdateInstanceTemplateResponse, error) {
	var resp hc.UpdateInstanceTemplateResponse

	p := path(hc.InstancePath, id, "template")

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UnlockInstance unlocks the instance with the given id.
func (s InstanceService) UnlockInstance(ctx context.Context, id int) (*hc.UnlockInstanceResponse, error) {
	var resp hc.UnlockInstanceResponse

	p := path(hc.InstancePath, id, "unlock")

	if err := s.Patch(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstance creates a new instance.
func (s InstanceService) CreateInstance(ctx context.Context, req hc.CreateInstanceRequest) (*hc.CreateInstanceResponse, error) {
	var resp hc.CreateInstanceResponse

	if err := s.Post(ctx, hc.InstancePath, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CalculateInstancesShowback calculates the showback for all the instances.
func (s InstanceService) CalculateInstancesShowback(ctx context.Context, req hc.CalculateInstancesShowbackRequest) error {
	p := path(hc.InstancePath, "showback")

	return s.Post(ctx, p, req, nil)
}

// CreateInstanceDisk creates a new disk for the instance with the given id.
func (s InstanceService) CreateInstanceDisk(ctx context.Context, id int, req hc.CreateInstanceDiskRequest) (*hc.CreateInstanceDiskResponse, error) {
	var resp hc.CreateInstanceDiskResponse

	p := path(hc.InstancePath, id, "disk")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstanceDiskImage creates a new disk image for the instance with the given id.
func (s InstanceService) CreateInstanceDiskImage(ctx context.Context, id, disk int, req hc.CreateInstanceDiskImageRequest) (*hc.CreateInstanceDiskImageResponse, error) {
	var resp hc.CreateInstanceDiskImageResponse

	p := path(hc.InstancePath, id, "disk", disk, "image")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstanceDiskSnapshot creates a new disk snapshot for the instance with the given id.
func (s InstanceService) CreateInstanceDiskSnapshot(ctx context.Context, id, disk int, req hc.CreateInstanceDiskSnapshotRequest) (*hc.CreateInstanceDiskSnapshotResponse, error) {
	var resp hc.CreateInstanceDiskSnapshotResponse

	p := path(hc.InstancePath, id, "disk", disk, "snapshot")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstanceNIC creates a new NIC for the instance with the given id.
func (s InstanceService) CreateInstanceNIC(ctx context.Context, id int, req hc.CreateInstanceNICRequest) (*hc.CreateInstanceNICResponse, error) {
	var resp hc.CreateInstanceNICResponse

	p := path(hc.InstancePath, id, "nic")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddInstanceSecurityGroup adds a security group to the instance with the given id.
func (s InstanceService) AddInstanceSecurityGroup(ctx context.Context, id, nic, sg int) (*hc.AddInstanceSecurityGroupResponse, error) {
	var resp hc.AddInstanceSecurityGroupResponse

	p := path(hc.InstancePath, id, "nic", nic, "security-group", sg)

	if err := s.Post(ctx, p, nil, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// AddInstanceSchedule adds a schedule to the instance with the given id.
func (s InstanceService) AddInstanceSchedule(ctx context.Context, id int, req hc.AddInstanceScheduleRequest) (*hc.AddInstanceScheduleResponse, error) {
	var resp hc.AddInstanceScheduleResponse

	p := path(hc.InstancePath, id, "schedule")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateInstanceSnapshot creates a new snapshot for the instance with the given id.
func (s InstanceService) CreateInstanceSnapshot(ctx context.Context, id int, req hc.CreateInstanceSnapshotRequest) (*hc.CreateInstanceSnapshotResponse, error) {
	var resp hc.CreateInstanceSnapshotResponse

	p := path(hc.InstancePath, id, "snapshot")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
