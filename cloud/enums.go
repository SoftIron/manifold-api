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

//go:generate go run "github.com/dmarkham/enumer" -type NetworkRecovery -linecomment -text

// NetworkRecovery is the recovery action to take when a vnet is in error.
type NetworkRecovery int

// VNetRecovery values.
const (
	FailureNetworkRecovery NetworkRecovery = iota // failure
	SuccessNetworkRecovery                        // success
	RetryNetworkRecovery                          // retry_vnet
	DeleteNetworkRecovery                         // delete_vnet
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

//go:generate go run "github.com/dmarkham/enumer" -type LCMState -linecomment -text

// LCMState is the Life Cycle Manager state of an instance.
type LCMState int

// LCMState values.
const (
	InitLCMState          LCMState = iota // init
	PrologLCMState                        // prolog
	BootLCMState                          // boot
	RunningLCMState                       // running
	MigrateLCMState                       // migrate
	SaveStopLCMState                      // save_stop
	SaveSuspendLCMState                   // save_suspend
	SaveMigrateLCMState                   // save_migrate
	PrologMigrateLCMState                 // prolog_migrate
	PrologResumeLCMState                  // prolog_resume
	EpilogStopLCMState                    // epilog_stop
	EpilogLCMState                        // epilog
	ShutdownLCMState                      // shutdown
	_
	_
	CleanupResubmitLCMState              // cleanup_resubmit
	UnknownLCMState                      // unknown
	HotplugLCMState                      // hotplug
	ShutdownPowerOffLCMState             // shutdown_poweroff
	BootUnknownLCMState                  // boot_unknown
	BootPowerOffLCMState                 // boot_poweroff
	BootSuspendedLCMState                // boot_suspended
	BootStoppedLCMState                  // boot_stopped
	CleanupDeleteLCMState                // cleanup_delete
	HotplugSnapshotLCMState              // hotplug_snapshot
	HotplugNICLCMState                   // hotplug_nic
	HotplugSaveAsLCMState                // hotplug_saveas
	HotplugSaveAsPowerOffLCMState        // hotplug_saveas_poweroff
	HotplutSaveAsSuspendedLCMState       // hotplug_saveas_suspended
	ShutdownUndeployLCMState             // shutdown_undeploy
	EpilogUndeployLCMState               // epilog_undeploy
	PrologUndeployLCMState               // prolog_undeploy
	BootUndeployLCMState                 // boot_undeploy
	HotplugPrologPowerOffLCMState        // hotplug_prolog_poweroff
	HotplugEpilogPowerOffLCMState        // hotplug_epilog_poweroff
	BootMigrateLCMState                  // boot_migrate
	BootFailureLCMState                  // boot_failure
	BootMigrateFailureLCMState           // boot_migrate_failure
	PrologMigrateFailureLCMState         // prolog_migrate_failure
	PrologFailureLCMState                // prolog_failure
	EpilogFailureLCMState                // epilog_failure
	EpilogStopFailureLCMState            // epilog_stop_failure
	EpilogUndeployFailureLCMState        // epilog_undeploy_failure
	PrologMigratePowerOffLCMState        // prolog_migrate_poweroff
	PrologMigratePowerOffFailureLCMState // prolog_migrate_poweroff_failure
	PrologMigrageSuspendLCMState         // prolog_migrate_suspend
	PrologMigrageSuspendFailureLCMState  // prolog_migrate_suspend_failure
	BootUndeployFailureLCMState          // boot_undeploy_failure
	BootStoppedFailureLCMState           // boot_stopped_failure
	PrologResumeFailureLCMState          // prolog_resume_failure
	PrologUndeployFailureLCMState        // prolog_undeploy_failure
	DiskSnapshotPowerOffLCMState         // disk_snapshot_poweroff
	DiskSnapshotRevertPowerOffLCMState   // disk_snapshot_revert_poweroff
	DiskSnapshotDeletePowerOffLCMState   // disk_snapshot_delete_poweroff
	DiskSnapshotSuspendLCMState          // disk_snapshot_suspend
	DiskSnapshotRevertSuspendedLCMState  // disk_snapshot_revert_suspended
	DiskSnapshotDeleteSuspendedLCMState  // disk_snapshot_delete_suspended
	DiskSnapshotLCMState                 // disk_snapshot
	_
	DiskSnapshotDeleteLCMState          // disk_snapshot_delete
	PrologMigrateUnknownLCMState        // prolog_migrate_unknown
	PrologMigrateUnknownFailureLCMState // prolog_migrate_unknown_failure
)
