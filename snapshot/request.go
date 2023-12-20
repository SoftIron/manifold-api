package snapshot

import (
	"time"
)

// Request abstracts the common attributes of ManualRequest and AutoRequest.
type Request interface {
	ID() uint
	Scheduled() time.Time
	Started() time.Time
	Finished() time.Time
	Schedule() Schedule
	Tag() Tag
	VMID() int
	SetID(uint) Request
	SetScheduled(time.Time) Request
	SetStarted(time.Time) Request
	SetFinished(time.Time) Request
	SetSchedule(Schedule) Request
}

// A requestBase holds the state of the snapshot. It includes the Schedule and in
// progress state of the request.
type requestBase struct {
	id        uint
	scheduled time.Time
	started   time.Time
	finished  time.Time
	schedule  Schedule
}

//go:generate go run "github.com/dmarkham/enumer" -type Type -linecomment -text

// Type indicates the type of Snapshot.
type Type int

const (
	// AutoSnapshot is a Snapshot for a HyperCloud VM.
	AutoSnapshot Type = iota // auto
	// ManualSnapshot is a Snapshot for a Ceph RBD image. For example the
	// `hypercloud-dashboard` snapshot is done directly one the rbd image.
	ManualSnapshot // manual
)
