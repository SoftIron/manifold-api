package snapshot_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/softiron/hypercloud-api/deprecated/v2/snapshot"
)

type parseTest struct {
	in, outText, outJSON string
}

// TestOptions tests the parsing of a schedule options, the options apply to
// every period in a given schedule line.
func TestOptions(t *testing.T) {
	_, err := snapshot.Parse("daily:local=1@8 yearly:local=1@015 options:")
	if err != nil { // handle null options
		t.Error(err)
	}

	snaps, err := snapshot.Parse("daily:local=1@8 yearly:local=1@015 options:pauseVM")
	if err != nil {
		t.Error(err)
	}

	for _, s := range snaps {
		if s.Options&snapshot.PauseVM != snapshot.PauseVM {
			t.Errorf("options should have PauseVM")
		}
	}
}

// TestPeriodMatching testing the sched Hourly(), Daily(), ... methods to verify
// they can pick out SnapshotTimes of only their Period using shuffled input.
func TestPeriodMatching(t *testing.T) {
	// this doesn't need to be table driven just one simple test.
	text := `
# comment

daily:local=1@8
yearly:local=1@015
monthly:local=1@11
hourly:local=1@2
monthly:local=1@13
yearly:local=1@016
monthly:local=1@10
weekly:local=1@Mon
weekly:local=1@Tue
weekly:local=1@Wed
daily:local=1@7
daily:local=1@5
hourly:local=1@3
weekly:local=1@Thu
yearly:local=1@014
hourly:local=1@4
daily:local=1@6
monthly:local=1@12
hourly:local=1@1
yearly:local=1@017
`

	s, err := snapshot.Parse(text)
	if err != nil {
		t.Errorf("parsing error %s", err)
	}

	h := s.Hourly()
	d := s.Daily()
	w := s.Weekly()
	m := s.Monthly()
	y := s.Yearly()

	x := "hourly:local=1@1,local=1@2,local=1@3,local=1@4"
	if h.String() != x {
		t.Errorf("%s does not match %s", h.String(), x)
	}

	x = "daily:local=1@5,local=1@6,local=1@7,local=1@8"
	if d.String() != x {
		t.Errorf("%s does not match %s", d.String(), x)
	}

	x = "weekly:local=1@Mon,local=1@Tue,local=1@Wed,local=1@Thu"
	if w.String() != x {
		t.Errorf("%s does not match %s", w.String(), x)
	}

	x = "monthly:local=1@10,local=1@11,local=1@12,local=1@13"
	if m.String() != x {
		t.Errorf("%s does not match %s", m.String(), x)
	}

	x = "yearly:local=1@014,local=1@015,local=1@016,local=1@017"
	if y.String() != x {
		t.Errorf("%s does not match %s", y.String(), x)
	}
}

func TestCommentFiltering(t *testing.T) {
	text := `

 # empty snapshot file

`
	s, err := snapshot.Parse(text)
	if err != nil {
		t.Errorf("parsing error %s", err)
	}

	if len(s) != 0 {
		t.Errorf("should be empty")
	}
}

func TestScheduleParsing(t *testing.T) {
	data := []parseTest{
		//
		// test for missing when
		//
		{
			in:      "hourly:local=1",
			outText: "hourly:local=1", // hourly means all hours, doesn't have a when
			outJSON: `[{"destination":"local","keep":1,"minute":0}]`,
		},
		{
			in:      "hourly:local=1@03",
			outText: "hourly:local=1@3",
			outJSON: `[{"destination":"local","keep":1,"minute":3}]`,
		},
		{
			in:      "hourly:local=1@30",
			outText: "hourly:local=1@30",
			outJSON: `[{"destination":"local","keep":1,"minute":30}]`,
		},
		{
			in:      "daily:local=1",
			outText: "daily:local=1@0", // a missing when time means midnight
			outJSON: `[{"destination":"local","keep":1,"minute":0,"hour":0}]`,
		},
		{
			in:      "daily:local=1@03",
			outText: "daily:local=1@3",
			outJSON: `[{"destination":"local","keep":1,"minute":0,"hour":3}]`,
		},
		{
			in:      "daily:local=1@3",
			outText: "daily:local=1@3",
			outJSON: `[{"destination":"local","keep":1,"minute":0,"hour":3}]`,
		},
		{
			in:      "weekly:local=1@1",
			outText: "weekly:local=1@Mon",
			outJSON: `[{"destination":"local","keep":1,"minute":0,"hour":0,"weekday":1}]`,
		},
		{
			in:      "weekly:local=1",
			outText: "weekly:local=1@Sun",
			outJSON: `[{"destination":"local","keep":1,"minute":0,"hour":0,"weekday":0}]`,
		},
		{
			in:      "monthly:local=1@02",
			outText: "monthly:local=1@2",
			outJSON: `[{"destination":"local","keep":1,"minute":0,"hour":0,"day":2}]`,
		},
		{
			in:      "monthly:local=1",
			outText: "monthly:local=1@1",
			outJSON: `[{"destination":"local","keep":1,"minute":0,"hour":0,"day":1}]`,
		},
		{
			in:      "yearly:local=1",
			outText: "yearly:local=1@Jan",
			outJSON: `[{"destination":"local","keep":1,"minute":0,"hour":0,"day":1,"month":1}]`,
		},
		{
			in:      "yearly:local=1@01",
			outText: "yearly:local=1@Jan",
			outJSON: `[{"destination":"local","keep":1,"minute":0,"hour":0,"day":1,"month":1}]`,
		},
		{
			in:      "yearly:local=1@032",
			outText: "yearly:local=1@032",
			outJSON: `[{"destination":"local","keep":1,"minute":0,"hour":0,"yearday":32}]`,
		},
		//
		// test for sorting the when
		//
		{
			in:      "daily:local=24@5,local=1@3,remote=1",
			outText: "daily:remote=1@0,local=1@3,local=24@5",
			// We test the outText first which means the SnapshotSchedule gets
			// sorted, thats why it's safe to test the sorted JSON as well.
			outJSON: `[{"destination":"remote","keep":1,"minute":0,"hour":0},{"destination":"local","keep":1,"minute":0,"hour":3},{"destination":"local","keep":24,"minute":0,"hour":5}]`, //nolint
		},
		{
			in:      "weekly:local=1@Mon,local=2@Fri,local=3@Sun",
			outText: "weekly:local=3@Sun,local=1@Mon,local=2@Fri",
		},
		{
			in:      "monthly:local=1@5,local=2@3,local=3@1",
			outText: "monthly:local=3@1,local=2@3,local=1@5",
		},
		{
			in:      "yearly:local=1@Sep,local=2@Mar,local=3@Jan",
			outText: "yearly:local=3@Jan,local=2@Mar,local=1@Sep",
		},
		{
			in:      "yearly:local=1@300,local=2@001,local=3@050",
			outText: "yearly:local=2@001,local=3@050,local=1@300",
		},
		//
		// complex lines
		//
		{
			in:      "hourly:local=10,archive=3 daily:local=3,remote=3@10 weekly:local=1@Tuesday",
			outText: "hourly:local=10,archive=3 daily:local=3@0,remote=3@10 weekly:local=1@Tue",
		},
	}

	for _, x := range data {
		s, err := snapshot.Parse(x.in)
		if err != nil {
			t.Errorf("failed to parse %q: %v", x.in, err)
		}

		if s.String() != x.outText {
			t.Errorf("%v does not match %v", s.String(), x.outText)
		}

		if x.outJSON == "" {
			continue // don't test JSON marshaling for all cases
		}

		b, err := json.Marshal(s)
		if err != nil {
			t.Errorf("failed to marshal %s: %v", s, err)
		}

		if string(b) != x.outJSON {
			t.Errorf("%v does not match %v", string(b), x.outJSON)
		}
	}
}

func TestScheduleUnmarshalJSON(t *testing.T) {
	data := []parseTest{
		{
			in:      `[{"destination":"local","keep":1,"minute":0}]`,
			outText: "hourly:local=1",
		},
		{
			in:      `[{"destination":"local","keep":1,"hour":0}]`,
			outText: "daily:local=1@0",
		},
		{
			in:      `[{"destination":"local","keep":1,"weekday":2}]`,
			outText: "weekly:local=1@Tue",
		},
		{
			in:      `[{"destination":"local","keep":1,"day":2}]`,
			outText: "monthly:local=1@2",
		},
		{
			in:      `[{"destination":"local","keep":1,"month":2}]`,
			outText: "yearly:local=1@Feb",
		},
		{
			in:      `[{"destination":"local","keep":1,"yearday":2}]`,
			outText: "yearly:local=1@002",
		},
	}

	for _, x := range data {
		var s snapshot.Schedules

		if err := json.Unmarshal([]byte(x.in), &s); err != nil {
			t.Errorf("cannot unmarshal %s: %v", x.in, err)
		}

		if s.String() != x.outText {
			t.Errorf("%s does not match %v", s, x.outText)
		}

	}
}

func TestIsTime(t *testing.T) {
	data := []struct {
		isTime bool
		time   string
		snap   snapshot.Schedule
	}{
		{
			isTime: false,
			snap:   snapshot.NewSchedule(),
		},
		{
			isTime: false,
			snap:   snapshot.NewSchedule(snapshot.Minute(10)),
		},
		{
			isTime: true,
			snap:   snapshot.NewSchedule(snapshot.Minute(15)),
		},
		{
			isTime: false,
			snap:   snapshot.NewSchedule(snapshot.Hour(3)),
		},
		{
			isTime: true,
			snap:   snapshot.NewSchedule(snapshot.Hour(3), snapshot.Minute(15)),
		},
		{
			isTime: false,
			snap: snapshot.NewSchedule(snapshot.Day(1), snapshot.Hour(3),
				snapshot.Minute(15)),
		},
		{
			isTime: true,
			snap: snapshot.NewSchedule(snapshot.Day(31), snapshot.Hour(3),
				snapshot.Minute(15)),
		},
		{
			isTime: false,
			snap: snapshot.NewSchedule(snapshot.Month(time.August),
				snapshot.Hour(3), snapshot.Minute(15)),
		},
		{
			isTime: true,
			snap: snapshot.NewSchedule(snapshot.Weekday(time.Wednesday),
				snapshot.Hour(3), snapshot.Minute(15)),
		},
		{
			isTime: false, // TODO(katz): this really feels like a valid schedule but is not supported ATM
			snap: snapshot.NewSchedule(snapshot.Month(time.August),
				snapshot.Weekday(time.Wednesday), snapshot.Hour(3), snapshot.Minute(15)),
		},
		{
			isTime: true,
			snap: snapshot.NewSchedule(snapshot.Month(time.August), snapshot.Day(31),
				snapshot.Hour(3), snapshot.Minute(15)),
		},
		{
			isTime: false,
			snap:   snapshot.NewSchedule(snapshot.YearDay(42), snapshot.Hour(3), snapshot.Minute(15)),
		},
		{
			isTime: true,
			snap:   snapshot.NewSchedule(snapshot.YearDay(243), snapshot.Hour(3), snapshot.Minute(15)),
		},
	}

	const layout = "2006-01-02 15:04"

	for _, x := range data {
		if x.time == "" {
			x.time = "2022-08-31 03:15"
		}

		when, err := time.Parse(layout, x.time)
		if err != nil {
			t.Errorf("bad time %s", err)
		}

		if x.snap.IsTime(when) != x.isTime {
			b, err := json.Marshal(x.snap)
			if err != nil {
				t.Errorf("failed to marshal %s: %v", x.snap, err)
			}

			t.Errorf("%s %s isTime should be %v time %s", x.snap.Period(), string(b), x.isTime, x.time)
		}
	}

}
