package cloud

//go:generate go run "github.com/dmarkham/enumer" -type DatastoreType -linecomment -text

// DatastoreType is the type of datastore.
type DatastoreType int

// Type values.
const (
	DatastoreImage  DatastoreType = iota // image
	DatastoreSystem                      // system
	DatastoreFile                        // file
)

//go:generate go run "github.com/dmarkham/enumer" -type DatastoreState -linecomment -text

// DatastoreState is the state of datastore.
type DatastoreState int

// State values.
const (
	DatastoreReady    DatastoreState = iota // ready
	DatastoreDisabled                       // disabled
)

// Datastore is the API payload based on the legacy xmlrpc backend.
type Datastore struct {
	ID                 int            `json:"id" yaml:"id"`
	UserID             int            `json:"user_id" yaml:"user_id"`
	GroupID            int            `json:"group_id" yaml:"group_id"`
	UserName           string         `json:"user_name" yaml:"user_name"`
	GroupName          string         `json:"group_name" yaml:"group_name"`
	Name               string         `json:"name" yaml:"name"`
	Permissions        Permissions    `json:"permissions" yaml:"permissions"`
	DatastoreMAD       string         `json:"datastore_mad" yaml:"datastore_mad"`
	TransferManagerMAD string         `json:"transfer_manager_mad" yaml:"transfer_manager_mad"`
	BasePath           string         `json:"base_path" yaml:"base_path"`
	Type               DatastoreType  `json:"type" yaml:"type" enums:"image,system,file"`
	DiskType           int            `json:"disk_type" yaml:"disk_type"`
	State              DatastoreState `json:"state" yaml:"state" enums:"ready,disabled"`
	Clusters           []int          `json:"clusters" yaml:"clusters"`
	TotalMB            int            `json:"total_mb" yaml:"total_mb"`
	FreeMB             int            `json:"free_mb" yaml:"free_mb"`
	UsedMB             int            `json:"used_mb" yaml:"used_mb"`
	Images             []int          `json:"images" yaml:"images"`
	Template           Template       `json:"template" yaml:"template"`
}

// DatastoreTemplate is the API payload based on the legacy xmlrpc backend.
type DatastoreTemplate struct {
	AllowOrphans             string   `json:"allow_orphans" yaml:"allow_orphans"`
	BridgeList               string   `json:"bridge_list" yaml:"bridge_list"`
	CloneTarget              string   `json:"clone_target" yaml:"clone_target"`
	CloneTargetShared        string   `json:"clone_target_shared" yaml:"clone_target_shared"`
	CloneTargetSSH           string   `json:"clone_target_ssh" yaml:"clone_target_ssh"`
	DiskType                 string   `json:"disk_type" yaml:"disk_type"`
	DiskTypeShared           string   `json:"disk_type_shared" yaml:"disk_type_shared"`
	DiskTypeSSH              string   `json:"disk_type_ssh" yaml:"disk_type_ssh"`
	Driver                   string   `json:"driver" yaml:"driver"`
	DatastoreMAD             string   `json:"ds_mad" yaml:"ds_mad"`
	DatastoreMigrate         bool     `json:"ds_migrate" yaml:"ds_migrate"`
	LNTarget                 string   `json:"ln_target" yaml:"ln_target"`
	LNTargetShared           string   `json:"ln_target_shared" yaml:"ln_target_shared"`
	LNTargetSSH              string   `json:"ln_target_ssh" yaml:"ln_target_ssh"`
	PoolName                 string   `json:"pool_name" yaml:"pool_name"`
	RestrictedDirs           string   `json:"restricted_dirs" yaml:"restricted_dirs"`
	SafeDirs                 string   `json:"safe_dirs" yaml:"safe_dirs"`
	Shared                   bool     `json:"shared" yaml:"shared"`
	StagingDir               string   `json:"staging_dir" yaml:"staging_dir"`
	TransferManagerMAD       string   `json:"tm_mad" yaml:"tm_mad"`
	TransferManagerMADSystem []string `json:"tm_mad_system" yaml:"tm_mad_system"`
	Type                     string   `json:"type" yaml:"type"`
	VCenterDCName            string   `json:"vcenter_dc_name" yaml:"vcenter_dc_name"`
	VCenterDCRef             string   `json:"vcenter_dc_ref" yaml:"vcenter_dc_ref"`
	VCenterDSName            string   `json:"vcenter_ds_name" yaml:"vcenter_ds_name"`
	VCenterDSRef             string   `json:"vcenter_ds_ref" yaml:"vcenter_ds_ref"`
	VCenterHost              string   `json:"vcenter_host" yaml:"vcenter_host"`
	VCenterInstanceID        string   `json:"vcenter_instance_id" yaml:"vcenter_instance_id"`
}
