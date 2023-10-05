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
	ID                 int               `json:"id" yaml:"id"`
	UserID             int               `json:"user_id" yaml:"user_id"`
	GroupID            int               `json:"group_id" yaml:"group_id"`
	UserName           string            `json:"user_name" yaml:"user_name"`
	GroupName          string            `json:"group_name" yaml:"group_name"`
	Name               string            `json:"name" yaml:"name"`
	Permissions        Permissions       `json:"permissions" yaml:"permissions"`
	DatastoreMAD       string            `json:"datastore_mad" yaml:"datastore_mad"`
	TransferManagerMAD string            `json:"transfer_manager_mad" yaml:"transfer_manager_mad"`
	BasePath           string            `json:"base_path" yaml:"base_path"`
	Type               DatastoreType     `json:"type" yaml:"type"`
	DiskType           int               `json:"disk_type" yaml:"disk_type"`
	State              DatastoreState    `json:"state" yaml:"state"`
	Clusters           []int             `json:"clusters,omitempty" yaml:"clusters,omitempty"`
	TotalMB            int               `json:"total_mb" yaml:"total_mb"`
	FreeMB             int               `json:"free_mb" yaml:"free_mb"`
	UsedMB             int               `json:"used_mb" yaml:"used_mb"`
	Images             []int             `json:"images,omitempty" yaml:"images,omitempty"`
	Template           DatastoreTemplate `json:"template" yaml:"template"`
	TemplateText       string            `json:"template_text,omitempty" yaml:"template_text,omitempty"`
}

// DatastoreTemplate is the API payload based on the legacy xmlrpc backend.
type DatastoreTemplate struct {
	Values                   map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	AllowOrphans             string            `json:"allow_orphans,omitempty" yaml:"allow_orphans,omitempty"`
	BridgeList               string            `json:"bridge_list,omitempty" yaml:"bridge_list,omitempty"`
	CloneTarget              string            `json:"clone_target,omitempty" yaml:"clone_target,omitempty"`
	CloneTargetShared        string            `json:"clone_target_shared,omitempty" yaml:"clone_target_shared,omitempty"`
	CloneTargetSSH           string            `json:"clone_target_ssh,omitempty" yaml:"clone_target_ssh,omitempty"`
	DiskType                 string            `json:"disk_type,omitempty" yaml:"disk_type,omitempty"`
	DiskTypeShared           string            `json:"disk_type_shared,omitempty" yaml:"disk_type_shared,omitempty"`
	DiskTypeSSH              string            `json:"disk_type_ssh,omitempty" yaml:"disk_type_ssh,omitempty"`
	Driver                   string            `json:"driver,omitempty" yaml:"driver,omitempty"`
	DatastoreMAD             string            `json:"ds_mad,omitempty" yaml:"ds_mad,omitempty"`
	DatastoreMigrate         bool              `json:"ds_migrate,omitempty" yaml:"ds_migrate,omitempty"`
	LNTarget                 string            `json:"ln_target,omitempty" yaml:"ln_target,omitempty"`
	LNTargetShared           string            `json:"ln_target_shared,omitempty" yaml:"ln_target_shared,omitempty"`
	LNTargetSSH              string            `json:"ln_target_ssh,omitempty" yaml:"ln_target_ssh,omitempty"`
	PoolName                 string            `json:"pool_name,omitempty" yaml:"pool_name,omitempty"`
	RestrictedDirs           string            `json:"restricted_dirs,omitempty" yaml:"restricted_dirs,omitempty"`
	SafeDirs                 string            `json:"safe_dirs,omitempty" yaml:"safe_dirs,omitempty"`
	Shared                   bool              `json:"shared,omitempty" yaml:"shared,omitempty"`
	StagingDir               string            `json:"staging_dir,omitempty" yaml:"staging_dir,omitempty"`
	TransferManagerMAD       string            `json:"tm_mad,omitempty" yaml:"tm_mad,omitempty"`
	TransferManagerMADSystem []string          `json:"tm_mad_system,omitempty" yaml:"tm_mad_system,omitempty"`
	Type                     string            `json:"type,omitempty" yaml:"type,omitempty"`
	VCenterDCName            string            `json:"vcenter_dc_name,omitempty" yaml:"vcenter_dc_name,omitempty"`
	VCenterDCRef             string            `json:"vcenter_dc_ref,omitempty" yaml:"vcenter_dc_ref,omitempty"`
	VCenterDSName            string            `json:"vcenter_ds_name,omitempty" yaml:"vcenter_ds_name,omitempty"`
	VCenterDSRef             string            `json:"vcenter_ds_ref,omitempty" yaml:"vcenter_ds_ref,omitempty"`
	VCenterHost              string            `json:"vcenter_host,omitempty" yaml:"vcenter_host,omitempty"`
	VCenterInstanceID        string            `json:"vcenter_instance_id,omitempty" yaml:"vcenter_instance_id,omitempty"`
}
