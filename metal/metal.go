// Package metal documents the request and response payloads for the bare metal
// layer. This was previously part of the SoftIron Storage Manager.
package metal

// Root path for API endpoint.
const (
	PathPrefix = "metal"

	CapacityPath    = "capacity"
	DatastorePath   = "datastore"
	HostPath        = "host"
	LicensePath     = "license"
	OSDPath         = "osd"
	PoolPath        = "pool"
	SharePath       = "share"
	DeviceClassPath = "deviceclasses"
)
