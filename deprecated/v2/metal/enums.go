package metal

//go:generate go run "github.com/dmarkham/enumer" -type FlagTypeEnum -linecomment -text

// FlagTypeEnum is the type for Ceph flag types.
type FlagTypeEnum int

// Ceph flag type.
const (
	FlagTypePriority      FlagTypeEnum = iota // PRIORITY
	FlagTypeExpected                          // EXPECTED
	FlagTypeInformational                     // INFORMATIONAL
)

//go:generate go run "github.com/dmarkham/enumer" -type ISCSIExportableTypeEnum -linecomment -text

// ISCSIExportableTypeEnum is the type for RBD iSCSI exportable types.
type ISCSIExportableTypeEnum int

// RBD iSCSI exportable type.
const (
	ISCSIExportableTypeAvailable ISCSIExportableTypeEnum = iota // AVAILABLE
	ISCSIExportableTypeExported                                 // EXPORTED
	ISCSIExportableTypeInvalid                                  // INVALID_OPTIONS
)

//go:generate go run "github.com/dmarkham/enumer" -type NetworkConnStatusEnum -linecomment -text

// NetworkConnStatusEnum is the type for network connection status.
type NetworkConnStatusEnum int

// Network connection status.
const (
	NetworkNotConnected NetworkConnStatusEnum = iota // NOT_CONNECTED
	NetworkConnected                                 // CONNECTED
)

//go:generate go run "github.com/dmarkham/enumer" -type ShareTypeEnum -linecomment -text

// ShareTypeEnum is the type of a share.
type ShareTypeEnum int

// Share types.
const (
	ShareTypeSSM    ShareTypeEnum = iota // SMB
	ShareTypeCephFS                      // CephFS
	ShareTypeISCSI                       // iSCSI
	ShareTypeRBD                         // RBD
)

//go:generate go run "github.com/dmarkham/enumer" -type OSDStateEnum -linecomment -text

// OSDStateEnum is the type for OSDs state.
type OSDStateEnum int

// OSDs states.
const (
	OSDStateUp   OSDStateEnum = iota // UP
	OSDStateDown                     // DOWN
	OSDStateOut                      // OUT
)

//go:generate go run "github.com/dmarkham/enumer" -type OSDClassEnum -linecomment -text

// OSDClassEnum is the type for OSDs class.
type OSDClassEnum int

// OSDs classes.
const (
	OSDClassSSD OSDClassEnum = iota // SSD
	OSDClassHDD                     // HDD
)

//go:generate go run "github.com/dmarkham/enumer" -type DatastoreSchemeEnum -linecomment -text

// DatastoreSchemeEnum is the protection scheme for a Ceph pool.
type DatastoreSchemeEnum int

// OSDs classes.
const (
	TripleReplication DatastoreSchemeEnum = iota // triple_replication
	EC4Plus2                                     // ec4+2
	EC8Plus3                                     // ec8+3
	EC8Plus4                                     // ec8+4
)
