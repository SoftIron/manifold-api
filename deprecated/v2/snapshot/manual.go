package snapshot

import "time"

// ManualRequest is the RBD information for the Request. A Request will have
// either a VMInfo or an RBDInfo.
type ManualRequest struct {
	requestBase
	VM      int
	RBDName string
}

// ID return the ID of the request.
func (m ManualRequest) ID() uint {
	return m.requestBase.id
}

// Scheduled returns the time the request was scheduled.
func (m ManualRequest) Scheduled() time.Time {
	return m.requestBase.scheduled
}

// Started returns the time the request was started.
func (m ManualRequest) Started() time.Time {
	return m.requestBase.started
}

// Finished returns the time the request was finished.
func (m ManualRequest) Finished() time.Time {
	return m.requestBase.finished
}

// Schedule returns the request's snapshot schedule.
func (m ManualRequest) Schedule() Schedule {
	return m.requestBase.schedule
}

// SetID sets the ID of the request.
func (m *ManualRequest) SetID(id uint) Request {
	m.requestBase.id = id
	return m
}

// SetScheduled sets the scheduled time of the request.
func (m *ManualRequest) SetScheduled(t time.Time) Request {
	m.requestBase.scheduled = t
	return m
}

// SetStarted sets the started time of the request.
func (m *ManualRequest) SetStarted(t time.Time) Request {
	m.requestBase.started = t
	return m
}

// SetFinished sets the finished time of the request.
func (m *ManualRequest) SetFinished(t time.Time) Request {
	m.requestBase.finished = t
	return m
}

// SetSchedule sets the request's snapshot schedule.
func (m *ManualRequest) SetSchedule(s Schedule) Request {
	m.requestBase.schedule = s
	return m
}

// VMID return the ID of the request's Virtual Machine.
func (m ManualRequest) VMID() int { return m.VM }

// Tag returns the Tag that will be used when processing the request.
func (m ManualRequest) Tag() Tag {
	return Tag{
		Period:  m.requestBase.schedule.Period(),
		Time:    m.requestBase.scheduled,
		RBD:     m.RBDName,
		Version: 3,
	}
}
