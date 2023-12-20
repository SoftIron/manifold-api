package snapshot

import (
	"time"

	"github.com/softiron/hypercloud-api/cloud"
)

// AutoRequest is a HyperCloud VM snapshot request.
type AutoRequest struct {
	requestBase
	VM         *cloud.LockedInstance
	Hypervisor string
	Images     []ImageSpec
}

// ID return the ID of the request.
func (a AutoRequest) ID() uint {
	return a.requestBase.id
}

// Scheduled returns the time the request was scheduled.
func (a AutoRequest) Scheduled() time.Time {
	return a.requestBase.scheduled
}

// Started returns the time the request was started.
func (a AutoRequest) Started() time.Time {
	return a.requestBase.started
}

// Finished returns the time the request was finished.
func (a AutoRequest) Finished() time.Time {
	return a.requestBase.finished
}

// Schedule returns the request's snapshot schedule.
func (a AutoRequest) Schedule() Schedule {
	return a.requestBase.schedule
}

// SetID sets the ID of the request.
func (a *AutoRequest) SetID(id uint) Request {
	a.requestBase.id = id
	return a
}

// SetScheduled sets the scheduled time of the request.
func (a *AutoRequest) SetScheduled(t time.Time) Request {
	a.requestBase.scheduled = t
	return a
}

// SetStarted sets the started time of the request.
func (a *AutoRequest) SetStarted(t time.Time) Request {
	a.requestBase.started = t
	return a
}

// SetFinished sets the finished time of the request.
func (a *AutoRequest) SetFinished(t time.Time) Request {
	a.requestBase.finished = t
	return a
}

// SetSchedule sets the request's snapshot schedule.
func (a *AutoRequest) SetSchedule(s Schedule) Request {
	a.requestBase.schedule = s
	return a
}

// VMID return the ID of the request's Virtual Machine.
func (a AutoRequest) VMID() int {
	if a.VM != nil {
		return a.VM.ID
	}
	return -1
}

// Tag returns the Tag that will be used when processing the request.
func (a AutoRequest) Tag() Tag {
	tag := Tag{
		Period:  a.requestBase.schedule.Period(),
		Time:    a.requestBase.scheduled,
		VM:      -1,
		Version: 2,
	}

	if a.VM != nil {
		tag.VM = a.VM.ID
	}

	return tag
}
