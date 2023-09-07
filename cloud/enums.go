package cloud

//go:generate go run "github.com/dmarkham/enumer" -type ImageType -transform upper -text

// ImageType is the type of image.
type ImageType int

// ImageType values.
const (
	_ ImageType = iota
	OS
	CDROM
	DataBlock
	Kernel
	RAMDisk
	Context
)

//go:generate go run "github.com/dmarkham/enumer" -type LockLevel -linecomment -text

// LockLevel is the level of lock.
type LockLevel int

// Lock levels.
const (
	_               LockLevel = iota
	UseLockLevel              // use
	ManageLockLevel           // manage
	AdminLockLevel            // admin
	AllLockLevel              // all
)

//go:generate go run "github.com/dmarkham/enumer" -type MigrationType -linecomment -text

// MigrationType is the type of migration.
type MigrationType int

// MigrationTypes values.
const (
	_                         MigrationType = iota
	SaveMigrationType                       // save
	PowerOffMigrationType                   // poweroff
	PowerOffHardMigrationType               // poweroff_hard
)

//go:generate go run "github.com/dmarkham/enumer" -type Filter -linecomment -text

// Filter is to filter list of object by user and group ownership.
type Filter int

// Filter type values.
const (
	GroupFilter        Filter = -4 // group
	UserFilter         Filter = -3 // user
	AllFilter          Filter = -2 // all
	UserAndGroupFilter Filter = -1 // user_group
)

//go:generate go run "github.com/dmarkham/enumer" -type InstanceRecovery -linecomment -text

// InstanceRecovery is the recovery action to take when an instance is in error.
type InstanceRecovery int

// InstanceRecovery values.
const (
	FailureRecovery        InstanceRecovery = iota // failure
	SuccessRecovery                                // success
	RetryRecovery                                  // retry
	DeleteRecovery                                 // delete
	DeleteRecreateRecovery                         // delete_recreate
	DeleteDBVNRecovery                             // delete_dbvn
)

//go:generate go run "github.com/dmarkham/enumer" -type VNetRecovery -linecomment -text

// VNetRecovery is the recovery action to take when a vnet is in error.
type VNetRecovery int

// VNetRecovery values.
const (
	FailureVNetRecovery VNetRecovery = iota // failure
	SuccessVNetRecovery                     // success
	RetryVNetRecovery                       // retry_vnet
	DeleteVNetRecovery                      // delete_vnet
)

//go:generate go run "github.com/dmarkham/enumer" -type Status -linecomment -text

// Status is the status of an object (usually an Instance).
type Status int

// Status values.
const (
	EnabledStatus  Status = iota // enabled
	DisabledStatus               // disabled
	OfflineStatus                // offline
)
