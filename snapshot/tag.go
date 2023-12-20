package snapshot

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const tagTimeLayout = "200601021504" // YYYYMMDDHHMM

// Tag is the named handle to uniquely name all snapshot images. The Time is the
// scheduled time, not the actual time the snapshot took place.
type Tag struct {
	VM      int    // VM is used for HyperCloud snapshots (AKA local)
	RBD     string // RBD is used for Ceph snapshots (AKA manual)
	Period  Period
	Time    time.Time
	Version int
}

// Type returns the type of snapshot represented by t.
func (t Tag) Type() Type {
	if t.RBD != "" {
		return ManualSnapshot
	}

	return AutoSnapshot
}

// ID returns the either the VM number or RBD name for the Tag.
func (t Tag) ID() string {
	if t.Type() == AutoSnapshot {
		return strconv.Itoa(t.VM)
	}

	return t.RBD
}

// String implements the Stringer interface for s. See ParseSnapshotTag for the
// format of this string.
func (t Tag) String() string {
	switch t.Version {
	case 0, 2:
		if t.Type() == AutoSnapshot {
			return fmt.Sprintf("%v-v2-%s-%s-%d", t.Type(), t.Period, t.Timestamp(), t.VM)
		}

		return fmt.Sprintf("%v-v2-%s-%s-m%s", AutoSnapshot, t.Period, t.Timestamp(), t.RBD)
	case 3:
		return fmt.Sprintf("%v-v3-%s-%s-%s", t.Type(), t.Period, t.Timestamp(), t.ID())
	}

	return ""
}

// MarshalText implements the encoding.TextMarshaler interface for t.
func (t Tag) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for t.
func (t *Tag) UnmarshalText(text []byte) error {
	var err error

	*t, err = ParseTag(string(text))

	return err
}

// Timestamp returns the Time in the form YYYYMMDDHHMM.
func (t Tag) Timestamp() string {
	return t.Time.Format(tagTimeLayout)
}

// Schedule returns a Schedule for the Tag based on it's time. The Schedules
// Destination and Keep values are the defaults since the Tag has no such
// information.
func (t Tag) Schedule() Schedule {
	_, month, day := t.Time.Date()

	switch t.Period {
	case Yearly:
		if day == 1 { // first of the month
			return NewSchedule(Month(month), Day(day))
		}
		return NewSchedule(YearDay(t.Time.YearDay()))
	case Monthly:
		return NewSchedule(Day(day))
	case Weekly:
		return NewSchedule(Weekday(t.Time.Weekday()))
	case Daily:
		return NewSchedule(Hour(t.Time.Hour()))
	case Hourly:
		return NewSchedule(Minute(t.Time.Minute()))
	case Minutely:
		return NewSchedule(EveryMinute())
	}

	return NewSchedule()
}

// IsZero returns true if a tag hasn't been initialized and returns false if it has
func (t Tag) IsZero() bool {
	if t.Period == 0 && t.VM == 0 && t.Time.IsZero() {
		return true
	}
	return false
}

// ParseTagTime returns the time.Time parsed from a timestamp string. The
// timestamp must be of the form YYYMMDDHHMM, otherwise an error is returned.
func ParseTagTime(t string) (time.Time, error) {
	return time.Parse(tagTimeLayout, t)
}

// ParseTag parses s to return a SnapshotTag.
func ParseTag(s string) (Tag, error) {
	fields := strings.SplitN(s, "-", 5)

	if len(fields) != 5 {
		return Tag{}, fmt.Errorf("invalid tag %q", s)
	}

	switch fields[1] {
	case "v2":
		tag, err := parseTag2(fields)
		if err != nil {
			err = fmt.Errorf("cannot parse %q: %w", s, err)
		}
		return tag, err
	case "v3":
		tag, err := parseTag3(fields)
		if err != nil {
			err = fmt.Errorf("cannot parse %q: %w", s, err)
		}
		return tag, err
	}

	return Tag{}, fmt.Errorf("invalid tag %q", s)
}

// parseTag2 parses s as a V2 tag to return a SnapshotTag. The format of s is
// one of the following formats:
//
//	auto-v2-PERIOD-TIME-ID
//	auto-v2-PERIOD-TIME-mID
//
// Where
//
//	PERIOD is one of yearly, monthly, weekly, daily, hourly
//	TIME is the timestamp in the form YYYYMMDDHHMM
//	ID is the VM ID
//	mID is the ID of the manual snapshot prefixed with "m"
func parseTag2(fields []string) (Tag, error) {
	if fields[0] != AutoSnapshot.String() {
		return Tag{}, fmt.Errorf("invalid prefix")
	}

	period, err := PeriodString(fields[2])
	if err != nil {
		return Tag{}, fmt.Errorf("invalid period %q", fields[2])
	}

	tm, err := time.Parse(tagTimeLayout, fields[3])
	if err != nil {
		return Tag{}, fmt.Errorf("invalid time %q", fields[3])
	}

	var (
		vmid int
		rbd  string
	)

	if strings.HasPrefix(fields[4], "m") {
		fields[4] = strings.Replace(fields[4], "m", "", 1)
		rbd = fields[4]
	}

	vmid, err = strconv.Atoi(fields[4])
	if err != nil {
		return Tag{}, fmt.Errorf("invalid VM ID %q", fields[4])
	}

	tag := Tag{
		Period:  period,
		Time:    tm,
		VM:      vmid,
		RBD:     rbd,
		Version: 2,
	}

	return tag, nil
}

// parseTag3 parses s as a V2 tag to return a SnapshotTag. The format of s is
// one of the following formats:
//
//	auto-v3-PERIOD-TIME-ID
//	manual-v3-PERIOD-TIME-VM
//
// where
//
//	PERIOD is one of yearly, monthly, weekly, daily, hourly
//	TIME is the timestamp in the form YYYYMMDDHHMM
//	ID is the VM ID
//	VM is the name of the VM
func parseTag3(fields []string) (Tag, error) {
	var (
		vmid int
		rbd  string
		err  error
	)

	switch fields[0] {
	case AutoSnapshot.String():
		vmid, err = strconv.Atoi(fields[4])
		if err != nil {
			return Tag{}, fmt.Errorf("invalid VM ID %q", fields[4])
		}
	case ManualSnapshot.String():
		rbd = fields[4]
	default:
		return Tag{}, fmt.Errorf("invalid tag")
	}

	period, err := PeriodString(fields[2])
	if err != nil {
		return Tag{}, fmt.Errorf("invalid period %q", fields[2])
	}

	tm, err := time.Parse(tagTimeLayout, fields[3])
	if err != nil {
		return Tag{}, fmt.Errorf("invalid time %q", fields[3])
	}

	tag := Tag{
		Period:  period,
		Time:    tm,
		VM:      vmid,
		RBD:     rbd,
		Version: 3,
	}

	return tag, nil
}
