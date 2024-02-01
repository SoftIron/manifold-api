// Package snapshot documents the request and response payloads related to the snapshot daemon.
package snapshot

// PathPrefix is the root path for API endpoint.
const PathPrefix = "snapper"

// Destination of the Snapshot.
type Destination int

const (
	// Local means take snapshot of a HyperCloud VM.
	Local Destination = iota
	// Remote means copy a Local to a remote Ceph cluster.
	Remote
	// Archive means copy a Local to a remote S3 bucket.
	Archive
)

//go:generate go run "github.com/dmarkham/enumer" -type Period -transform lower -text

// Period indicates the time interval of a snapshot.
type Period int

// Snapshots are taken at one of more of these intervals.
const (
	_ Period = iota
	Minutely
	Hourly
	Daily
	Weekly
	Monthly
	Yearly
)

// CatCompress is to support the legacy 'no compression' option
const CatCompress = "cat"
