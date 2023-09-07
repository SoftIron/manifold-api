package hc

import (
	"time"
)

// InstanceDisk is the API payload based on the legacy xmlrpc backend.
type InstanceDisk struct {
	Values                map[string]string `json:"values" yaml:"values"`
	AllowOrphans          string            `json:"allow_orphans" yaml:"allow_orphans"`
	Clone                 bool              `json:"clone" yaml:"clone"`
	CloneTarget           string            `json:"clone_target" yaml:"clone_target"`
	ClusterID             string            `json:"cluster_id" yaml:"cluster_id"`
	Datastore             string            `json:"datastore" yaml:"datastore"`
	DatastoreID           string            `json:"datastore_id" yaml:"datastore_id"`
	DevPrefix             string            `json:"dev_prefix" yaml:"dev_prefix"`
	DiskID                string            `json:"disk_id" yaml:"disk_id"`
	DiskSnapshotTotalSize string            `json:"disk_snapshot_total_size" yaml:"disk_snapshot_total_size"`
	DiskType              string            `json:"disk_type" yaml:"disk_type"`
	Driver                string            `json:"driver" yaml:"driver"`
	Image                 string            `json:"image" yaml:"image"`
	ImageID               string            `json:"image_id" yaml:"image_id"`
	ImageState            string            `json:"image_state" yaml:"image_state"`
	LnTarget              string            `json:"ln_target" yaml:"ln_target"`
	OriginalSize          string            `json:"original_size" yaml:"original_size"`
	PoolName              string            `json:"pool_name" yaml:"pool_name"`
	Readonly              bool              `json:"readonly" yaml:"readonly"`
	Save                  bool              `json:"save" yaml:"save"`
	Size                  string            `json:"size" yaml:"size"`
	Source                string            `json:"source" yaml:"source"`
	Target                string            `json:"target" yaml:"target"`
	TmMAD                 string            `json:"tm_mad" yaml:"tm_mad"`
	TmMADSystem           string            `json:"tm_madsystem" yaml:"tm_madsystem"`
	Type                  string            `json:"type" yaml:"type"`
	VCenterDSRef          string            `json:"vcenter_dsref" yaml:"vcenter_dsref"`
	VCenterInstanceID     string            `json:"vcenter_instance_id" yaml:"vcenter_instance_id"`
}

// InstanceDiskSize is the API payload based on the legacy xmlrpc backend.
type InstanceDiskSize struct {
	ID   int `json:"id" yaml:"id"`
	Size int `json:"size" yaml:"size"`
}

// InstanceHistory is the API payload based on the legacy xmlrpc backend.
type InstanceHistory struct {
	OID              int       `json:"oid" yaml:"oid"`
	SequenceNumber   int       `json:"sequence_number" yaml:"sequence_number"`
	Hostname         string    `json:"hostname" yaml:"hostname"`
	HID              int       `json:"hid" yaml:"hid"`
	CID              int       `json:"cid" yaml:"cid"`
	StartTime        time.Time `json:"start_time" yaml:"start_time"`
	EndTime          time.Time `json:"end_time" yaml:"end_time"`
	InstanceMAD      string    `json:"instance_mad" yaml:"instance_mad"`
	TmMAD            string    `json:"tm_mad" yaml:"tm_mad"`
	DSID             int       `json:"dsid" yaml:"dsid"`
	PrologStartTime  time.Time `json:"prolog_start_time,omitempty" yaml:"prolog_start_time"`
	PrologEndtime    time.Time `json:"prolog_end_time" yaml:"prolog_end_time"`
	RunningStartTime time.Time `json:"running_start_time" yaml:"running_start_time"`
	RunningEndTime   time.Time `json:"running_end_time" yaml:"running_end_time"`
	EpilogStartTime  time.Time `json:"epilog_start_time" yaml:"epilog_start_time"`
	EpilogEndTime    time.Time `json:"epilog_end_time" yaml:"epilog_end_time"`
	Action           int       `json:"action" yaml:"action"`
	UID              int       `json:"uid" yaml:"uid"`
	GID              int       `json:"gid" yaml:"gid"`
	RequestID        string    `json:"request_id" yaml:"request_id"`
}

// InstanceHistoryRecords is the API payload based on the legacy xmlrpc backend.
type InstanceHistoryRecords struct {
	History []InstanceHistory `json:"history" yaml:"history"`
}

// InstanceLock is the API payload based on the legacy xmlrpc backend.
type InstanceLock struct {
	Locked int `json:"locked" yaml:"locked"`
	Owner  int `json:"owner" yaml:"owner"`
	Time   int `json:"time" yaml:"time"`
	ReqID  int `json:"req_id" yaml:"req_id"`
}

// InstanceMonitoring is the API payload based on the legacy xmlrpc backend.
type InstanceMonitoring struct {
	CPU                             float64            `json:"cpu" yaml:"cpu"`
	DiskReadBytes                   int                `json:"disk_read_bytes" yaml:"disk_read_bytes"`
	DiskReadIOps                    int                `json:"disk_read_iops" yaml:"disk_read_iops"`
	DiskWriteBytes                  int                `json:"disk_write_bytes" yaml:"disk_write_bytes"`
	DiskWriteIOps                   int                `json:"disk_write_iops" yaml:"disk_write_iops"`
	DiskSize                        []InstanceDiskSize `json:"disk_size" yaml:"disk_size"`
	ID                              int                `json:"id" yaml:"id"`
	Memory                          int                `json:"memory" yaml:"memory"`
	NetRX                           int                `json:"netrx" yaml:"netrx"`
	NetTX                           int                `json:"nettx" yaml:"nettx"`
	Timestamp                       int                `json:"timestamp" yaml:"timestamp"`
	VCenterESXHost                  string             `json:"vcenter_esxhost" yaml:"vcenter_esxhost"`
	VCenterGuestState               string             `json:"vcenter_guest_state" yaml:"vcenter_guest_state"`
	VCenterRPName                   string             `json:"vcenter_rpname" yaml:"vcenter_rpname"`
	VCenterVMWareToolsRunningStatus string             `json:"vcenter_vmware_tools_running_status" yaml:"vcenter_vmware_tools_running_status"`
	VCenterVMWareToolsVersion       string             `json:"vcenter_vmware_tools_version" yaml:"vcenter_vmware_tools_version"`
	VCenterVMWareToolsVersionStatus string             `json:"vcenter_vmware_tools_version_status" yaml:"vcenter_vmware_tools_version_status"`
	VCenterVMName                   string             `json:"vcenter_vmname" yaml:"vcenter_vmname"`
}

// InstanceNIC is the API payload based on the legacy xmlrpc backend.
type InstanceNIC struct {
	Values               map[string]string `json:"values" yaml:"values"`
	VCenterInstanceID    string            `json:"vcenter_instance_id" yaml:"vcenter_instance_id"`
	VCenterNetRef        string            `json:"vcenter_net_ref" yaml:"vcenter_net_ref"`
	VCenterPortgroupType string            `json:"vcenter_portgroup_type" yaml:"vcenter_portgroup_type"`
}

// InstanceNICAlias is the API payload based on the legacy xmlrpc backend.
type InstanceNICAlias struct {
	Values               map[string]string `json:"values" yaml:"values"`
	AliasID              string            `json:"alias_id" yaml:"alias_id"`
	Parent               string            `json:"parent" yaml:"parent"`
	ParentID             string            `json:"parent_id" yaml:"parent_id"`
	VCenterInstanceID    string            `json:"vcenter_instance_id" yaml:"vcenter_instance_id"`
	VCenterNetRef        string            `json:"vcenter_net_ref" yaml:"vcenter_net_ref"`
	VCenterPortgroupType string            `json:"vcenter_portgroup_type" yaml:"vcenter_portgroup_type"`
}

// InstancePermissions is the API payload based on the legacy xmlrpc backend.
type InstancePermissions struct {
	OwnerUse    bool `json:"owner_use,omitempty" yaml:"owner_use,omitempty"`
	OwnerManage bool `json:"owner_manage,omitempty" yaml:"owner_manage,omitempty"`
	OwnerAdmin  bool `json:"owner_admin,omitempty" yaml:"owner_admin,omitempty"`
	GroupUse    bool `json:"group_use,omitempty" yaml:"group_use,omitempty"`
	GroupManage bool `json:"group_manage,omitempty" yaml:"group_manage,omitempty"`
	GroupAdmin  bool `json:"group_admin,omitempty" yaml:"group_admin,omitempty"`
	OtherUse    bool `json:"other_use,omitempty" yaml:"other_use,omitempty"`
	OtherManage bool `json:"other_manage,omitempty" yaml:"other_manage,omitempty"`
	OtherAdmin  bool `json:"other_admin,omitempty" yaml:"other_admin,omitempty"`
}

// InstanceSchedAction is the API payload based on the legacy xmlrpc backend.
type InstanceSchedAction struct {
	Action   string `json:"action" yaml:"action"`
	Args     string `json:"args" yaml:"args"`
	Days     string `json:"days" yaml:"days"`
	EndType  string `json:"end_type" yaml:"end_type"`
	EndValue string `json:"end_value" yaml:"end_value"`
	ID       string `json:"id" yaml:"id"`
	Repeat   string `json:"repeat" yaml:"repeat"`
	Time     string `json:"time" yaml:"time"`
	Warning  string `json:"warning" yaml:"warning"`
}

// InstanceSnapshot is the API payload based on the legacy xmlrpc backend.
type InstanceSnapshot struct {
	Action         string    `json:"action" yaml:"action"`
	Active         bool      `json:"active" yaml:"active"`
	HypervisorID   string    `json:"hypervisor_id" yaml:"hypervisor_id"`
	Name           string    `json:"name" yaml:"name"`
	SnapshotID     int       `json:"snapshot_id" yaml:"snapshot_id"`
	SystemDiskSize string    `json:"system_disk_size" yaml:"system_disk_size"`
	Time           time.Time `json:"time" yaml:"time"`
}

// InstanceSnapshots is the API payload based on the legacy xmlrpc backend.
type InstanceSnapshots struct {
	AllowOrphans string         `json:"allow_orphans" yaml:"allow_orphans"`
	CurrentBase  int            `json:"current_base" yaml:"current_base"`
	DiskID       int            `json:"disk_id" yaml:"disk_id"`
	NextSnapshot int            `json:"next_snapshot" yaml:"next_snapshot"`
	Snapshot     []InstanceAnon `json:"snapshot" yaml:"snapshot"`
}

// InstanceAnon is the API payload based on the legacy xmlrpc backend.
type InstanceAnon struct {
	Active   string `json:"active" yaml:"active"`
	Children string `json:"children" yaml:"children"`
	Date     int    `json:"date" yaml:"date"`
	ID       int    `json:"id" yaml:"id"`
	Name     string `json:"name" yaml:"name"`
	Parent   int    `json:"parent" yaml:"parent"`
	Size     int    `json:"size" yaml:"size"`
}

// InstanceTemplateData is the API payload based on the legacy xmlrpc backend.
type InstanceTemplateData struct {
	AutomaticDSRequirements   string                `json:"automatic_dsrequirements" yaml:"automatic_dsrequirements"`
	AutomaticNICRequirements  string                `json:"automatic_nicrequirements" yaml:"automatic_nicrequirements"`
	AutomaticRequirements     string                `json:"automatic_requirements" yaml:"automatic_requirements"`
	CloningTemplateID         string                `json:"cloning_template_id" yaml:"cloning_template_id"`
	Context                   string                `json:"context" yaml:"context"`
	CPU                       string                `json:"cpu" yaml:"cpu"`
	CPUCost                   string                `json:"cpucost" yaml:"cpucost"`
	Disk                      []InstanceDisk        `json:"disk" yaml:"disk"`
	DiskCost                  string                `json:"disk_cost" yaml:"disk_cost"`
	Emulator                  string                `json:"emulator" yaml:"emulator"`
	Features                  string                `json:"features" yaml:"features"`
	HypervOptions             string                `json:"hyperv_options" yaml:"hyperv_options"`
	Graphics                  string                `json:"graphics" yaml:"graphics"`
	Imported                  string                `json:"imported" yaml:"imported"`
	Input                     string                `json:"input" yaml:"input"`
	Memory                    string                `json:"memory" yaml:"memory"`
	MemoryCost                string                `json:"memory_cost" yaml:"memory_cost"`
	MemoryMax                 string                `json:"memory_max" yaml:"memory_max"`
	MemorySlots               string                `json:"memory_slots" yaml:"memory_slots"`
	MemoryResizeMode          string                `json:"memory_resize_mode" yaml:"memory_resize_mode"`
	NIC                       []InstanceNIC         `json:"nic" yaml:"nic"`
	NICAlias                  []InstanceNICAlias    `json:"nic_alias" yaml:"nic_alias"`
	NICDefault                string                `json:"nic_default" yaml:"nic_default"`
	NumaNode                  string                `json:"numa_node" yaml:"numa_node"`
	OS                        string                `json:"os" yaml:"os"`
	PCI                       string                `json:"pci" yaml:"pci"`
	Raw                       string                `json:"raw" yaml:"raw"`
	SchedAction               []InstanceSchedAction `json:"sched_action" yaml:"sched_action"`
	SecurityGroupRule         []string              `json:"security_group_rule" yaml:"security_group_rule"`
	Snapshot                  []InstanceSnapshot    `json:"snapshot" yaml:"snapshot"`
	SpiceOptions              string                `json:"spice_options" yaml:"spice_options"`
	SubmitOnHold              string                `json:"submit_on_hold" yaml:"submit_on_hold"`
	TemplateID                string                `json:"template_id" yaml:"template_id"`
	TmMADSystem               string                `json:"tm_madsystem" yaml:"tm_madsystem"`
	Topology                  string                `json:"topology" yaml:"topology"`
	VCPU                      string                `json:"vcpu" yaml:"vcpu"`
	VCPUMax                   string                `json:"vcpu_max" yaml:"vcpu_max"`
	InstanceGroup             []string              `json:"instance_group" yaml:"instance_group"`
	InstanceID                string                `json:"instance_id" yaml:"instance_id"`
	VRouterID                 string                `json:"vrouter_id" yaml:"vrouter_id"`
	VRouterKeepalivedID       string                `json:"vrouter_keepalived_id" yaml:"vrouter_keepalived_id"`
	VRouterKeepalivedPassword string                `json:"vrouter_keepalived_password" yaml:"vrouter_keepalived_password"`
}

// InstanceUserTemplate is the API payload based on the legacy xmlrpc backend.
type InstanceUserTemplate struct {
	Values                map[string]string `json:"values" yaml:"values"`
	VCenterCCRRef         string            `json:"vcenter_ccrref" yaml:"vcenter_ccrref"`
	VCenterDSRef          string            `json:"vcenter_dsref" yaml:"vcenter_dsref"`
	VCenterInstanceID     string            `json:"vcenter_instance_id" yaml:"vcenter_instance_id"`
	Error                 string            `json:"error" yaml:"error"`
	Hypervisor            string            `json:"hypervisor" yaml:"hypervisor"`
	Info                  string            `json:"info" yaml:"info"`
	Logo                  string            `json:"logo" yaml:"logo"`
	LXDSecurityPrivileged bool              `json:"lxdsecurity_privileged" yaml:"lxdsecurity_privileged"`
	SchedRequirements     string            `json:"sched_requirements" yaml:"sched_requirements"`
}

// Instance is the API payload based on the legacy xmlrpc backend.
type Instance struct {
	ID             int                    `json:"id" yaml:"id"`
	UID            int                    `json:"uid" yaml:"uid"`
	GID            int                    `json:"gid" yaml:"gid"`
	UserName       string                 `json:"user_name" yaml:"user_name"`
	GroupName      string                 `json:"group_name" yaml:"group_name"`
	Name           string                 `json:"name" yaml:"name"`
	Permissions    InstancePermissions    `json:"permissions" yaml:"permissions"`
	LastPoll       int                    `json:"last_poll" yaml:"last_poll"`
	State          InstanceState          `json:"state" yaml:"state"`
	LCMState       LCMState               `json:"lcm_state" yaml:"lcm_state"`
	PrevState      InstanceState          `json:"prev_state" yaml:"prev_state"`
	PrevLcmState   LCMState               `json:"prev_lcm_state" yaml:"prev_lcm_state"`
	Resched        int                    `json:"resched" yaml:"resched"`
	StartTime      time.Time              `json:"start_time" yaml:"start_time"`
	EndTime        time.Time              `json:"end_time" yaml:"end_time"`
	DeployID       string                 `json:"deploy_id" yaml:"deploy_id"`
	Lock           InstanceLock           `json:"lock" yaml:"lock"`
	Monitoring     InstanceMonitoring     `json:"monitoring" yaml:"monitoring"`
	Template       InstanceTemplateData   `json:"template" yaml:"template"`
	UserTemplate   InstanceUserTemplate   `json:"user_template" yaml:"user_template"`
	HistoryRecords InstanceHistoryRecords `json:"history_records" yaml:"history_records"`
	Snapshots      []InstanceSnapshots    `json:"snapshots" yaml:"snapshots"`
}
