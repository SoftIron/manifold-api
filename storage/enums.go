package storage

//go:generate go run "github.com/dmarkham/enumer" -type FlagTypeEnum -linecomment -text

// FlagTypeEnum is the type for Ceph flag types.
type FlagTypeEnum int

// Ceph flag type.
const (
	FlagTypePriority      FlagTypeEnum = iota // PRIORITY
	FlagTypeExpected                          // EXPECTED
	FlagTypeInformational                     // INFORMATIONAL
)

//go:generate go run "github.com/dmarkham/enumer" -type IscsiExportableTypeEnum -linecomment -text

// IscsiExportableTypeEnum is the type for RBD iSCSI exportable types.
type IscsiExportableTypeEnum int

// RBD iSCSI exportable type.
const (
	IscsiExportableTypeAvailable IscsiExportableTypeEnum = iota // AVAILABLE
	IscsiExportableTypeExported                                 // EXPORTED
	IscsiExportableTypeInvalid                                  // INVALID_OPTIONS
)

//go:generate go run "github.com/dmarkham/enumer" -type NetworkConnStatusEnum -linecomment -text

// NetworkConnStatusEnum is the type for network connection status.
type NetworkConnStatusEnum int

// Network connection status.
const (
	NetworkNotConnected NetworkConnStatusEnum = iota // NOT_CONNECTED
	NetworkConnected                                 // CONNECTED
)

//go:generate go run "github.com/dmarkham/enumer" -type NodeStatusEnum -linecomment -text

// NodeStatusEnum is the type for node running status.
type NodeStatusEnum int

// TODO: there is no discovery so several of these are not used.

// Node running status.
const (
	NodeStatusDiscovering   NodeStatusEnum = iota // DISCOVERING
	NodeStatusMissing                             // MISSING
	NodeStatusOff                                 // OFF
	NodeStatusCrashed                             // CRASHED
	NodeStatusBooting                             // DOWN
	NodeStatusOn                                  // ON
	NodeStatusPendingOn                           // PENDING_ON
	NodeStatusPendingOff                          // PENDING_OFF
	NodeStatusPendingReboot                       // PENDING_REBOOT
	NodeStatusPendingChange                       // PENDING_CHANGE
)

//go:generate go run "github.com/dmarkham/enumer" -type ShareTypeEnum -linecomment -text

// ShareTypeEnum is the type of a share.
type ShareTypeEnum int

// Share types.
const (
	ShareTypeSmb    ShareTypeEnum = iota // SMB
	ShareTypeCephFs                      // CephFS
	ShareTypeIscsi                       // iSCSI
	ShareTypeRbd                         // RBD
)

//go:generate go run "github.com/dmarkham/enumer" -type OsdStateEnum -linecomment -text

// OsdStateEnum is the type for OSDs state.
type OsdStateEnum int

// OSDs states.
const (
	OsdStateUp   OsdStateEnum = iota // UP
	OsdStateDown                     // DOWN
	OsdStateOut                      // OUT
)

//go:generate go run "github.com/dmarkham/enumer" -type OsdClassEnum -linecomment -text

// OsdClassEnum is the type for OSDs class.
type OsdClassEnum int

// OSDs classes.
const (
	OsdClassSsd OsdClassEnum = iota // SSD
	OsdClassHdd                     // HDD
)
