package snapshot_test

import (
	"testing"
	"time"

	"git.softiron.com/sw/hc/snapper.git/snapshot"
)

func TestTagParsing(t *testing.T) {
	vm := "auto-v2-hourly-202001020300-4"
	rbd := "manual-v3-hourly-202001020300-image-name"

	tag, err := snapshot.ParseTag(vm)
	if err != nil {
		t.Error(err)
	}

	if tag.Period != snapshot.Hourly {
		t.Errorf("incorrect Period")
	}

	if tag.VM != 4 {
		t.Errorf("incorrect vm ID")
	}

	// YYYYMMDDHHMM
	tm, err := time.Parse("200601021504", "202001020300")
	if err != nil {
		t.Error(err)
	}

	if !tag.Time.Equal(tm) {
		t.Errorf("incorrect time")
	}

	tag, err = snapshot.ParseTag(rbd)
	if err != nil {
		t.Error(err)
	}

	if tag.Type() != snapshot.ManualSnapshot {
		t.Errorf("not a manual snapshot")
	}

	if tag.RBD != "image-name" {
		t.Errorf("incorrect rbd")
	}
}

// TestRBDTag checks that the Tag for a RBD snapshot uses the RBDInfo.ID as the
// virtual machine number, and include a "m" prefix.
func TestRBDTag(t *testing.T) {
	tm, err := time.Parse("01/02/2006 15:04", "01/02/2020 03:00")
	if err != nil {
		t.Error(err)
	}

	r := snapshot.ManualRequest{
		VM:      6,
		RBDName: "imagename",
	}
	r.SetScheduled(tm)
	r.SetSchedule(snapshot.NewSchedule(snapshot.Minute(0)))

	if r.Tag().String() != "manual-v3-hourly-202001020300-imagename" {
		t.Errorf("incorrect tag %q", r.Tag().String())
	}
}
