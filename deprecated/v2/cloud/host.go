package cloud

// Host is the API payload based on the legacy xmlrpc backend.
type Host struct {
	ID                    int            `json:"id" yaml:"id"`
	Name                  string         `json:"name" yaml:"name"`
	State                 int            `json:"state" yaml:"state"`
	PrevState             int            `json:"prev_state" yaml:"prev_state"`
	InformationManagerMAD string         `json:"information_manager_mad" yaml:"information_manager_mad"`
	InstanceMAD           string         `json:"instance_mad" yaml:"instance_mad"`
	ClusterID             int            `json:"cluster_id" yaml:"cluster_id"`
	Cluster               string         `json:"cluster" yaml:"cluster"`
	HostShare             HostShare      `json:"host_share" yaml:"host_share"`
	Instances             []int          `json:"instances" yaml:"instances"`
	Template              HostTemplate   `json:"template" yaml:"template"`
	TemplateText          string         `json:"template_text" yaml:"template_text"`
	Monitoring            HostMonitoring `json:"monitoring" yaml:"monitoring"`
}

// HostShare is the API payload based on the legacy xmlrpc backend.
type HostShare struct {
	MemUsage         int        `json:"mem_usage" yaml:"mem_usage"`
	CPUUsage         int        `json:"cpuusage" yaml:"cpuusage"`
	TotalMem         int        `json:"total_mem" yaml:"total_mem"`
	TotalCPU         int        `json:"total_cpu" yaml:"total_cpu"`
	MaxMem           int        `json:"max_mem" yaml:"max_mem"`
	MaxCPU           int        `json:"max_cpu" yaml:"max_cpu"`
	RunningInstances int        `json:"running_instances" yaml:"running_instances"`
	InstancessThread int        `json:"instances_thread" yaml:"instances_thread"`
	Datastores       Datastores `json:"datastores" yaml:"datastores"`
	PCIDevices       []PCI      `json:"pcidevices" yaml:"pcidevices"`
	NumaNodes        []Node     `json:"numa_nodes" yaml:"numa_nodes"`
}

// PCI is the API payload based on the legacy xmlrpc backend.
type PCI struct {
	Address      string `json:"address" yaml:"address"`
	Bus          string `json:"bus" yaml:"bus"`
	Class        string `json:"class" yaml:"class"`
	ClassName    string `json:"class_name" yaml:"class_name"`
	Device       string `json:"device" yaml:"device"`
	DeviceName   string `json:"device_name" yaml:"device_name"`
	Domain       string `json:"domain" yaml:"domain"`
	Function     string `json:"function" yaml:"function"`
	NumaNode     string `json:"numa_node" yaml:"numa_node"`
	ShortAddress string `json:"short_address" yaml:"short_address"`
	Slot         string `json:"slot" yaml:"slot"`
	Type         string `json:"type" yaml:"type"`
	Vendor       string `json:"vendor" yaml:"vendor"`
	VendorName   string `json:"vendor_name" yaml:"vendor_name"`
	InstanceID   int    `json:"instance_id" yaml:"instance_id"`
}

// Node is the API payload based on the legacy xmlrpc backend.
type Node struct {
	Core     []Core     `json:"core" yaml:"core"`
	Hugepage []HugePage `json:"hugepage" yaml:"hugepage"`
	Memory   Memory     `json:"memory" yaml:"memory"`
	NodeID   int        `json:"node_id" yaml:"node_id"`
}

// Capacity is the API payload based on the legacy xmlrpc backend.
type Capacity struct {
	FreeCPU    int `json:"free_cpu" yaml:"free_cpu"`
	FreeMemory int `json:"free_memory" yaml:"free_memory"`
	UsedCPU    int `json:"used_cpu" yaml:"used_cpu"`
	UsedMemory int `json:"used_memory" yaml:"used_memory"`
}

// Core is the API payload based on the legacy xmlrpc backend.
type Core struct {
	CPUs      string `json:"cpus" yaml:"cpus"`
	Dedicated bool   `json:"dedicated" yaml:"dedicated"`
	Free      int    `json:"free" yaml:"free"`
	ID        int    `json:"id" yaml:"id"`
}

// Datastores is the API payload based on the legacy xmlrpc backend.
type Datastores struct {
	DiskUsage int `json:"disk_usage" yaml:"disk_usage"`
	FreeDisk  int `json:"free_disk" yaml:"free_disk"`
	MaxDisk   int `json:"max_disk" yaml:"max_disk"`
	UsedDisk  int `json:"used_disk" yaml:"used_disk"`
}

// HugePage is the API payload based on the legacy xmlrpc backend.
type HugePage struct {
	Free  int `json:"free" yaml:"free"`
	Pages int `json:"pages" yaml:"pages"`
	Size  int `json:"size" yaml:"size"`
	Usage int `json:"usage" yaml:"usage"`
}

// Memory is the API payload based on the legacy xmlrpc backend.
type Memory struct {
	Distance string `json:"distance" yaml:"distance"`
	Free     int    `json:"free" yaml:"free"`
	Total    int    `json:"total" yaml:"total"`
	Usage    int    `json:"usage" yaml:"usage"`
	Used     int    `json:"used" yaml:"used"`
}

// HostMonitoring is the API payload based on the legacy xmlrpc backend.
type HostMonitoring struct {
	Timestamp int      `json:"timestamp" yaml:"timestamp"`
	ID        int      `json:"id" yaml:"id"`
	Capacity  Capacity `json:"capacity" yaml:"capacity"`
	System    System   `json:"system" yaml:"system"`
}

// System is the API payload based on the legacy xmlrpc backend.
type System struct {
	Netrx int `json:"netrx" yaml:"netrx"`
	Nettx int `json:"nettx" yaml:"nettx"`
}

// HostTemplate is the API payload based on the legacy xmlrpc backend.
type HostTemplate struct {
	Values                  map[string]string      `json:"values" yaml:"values"`
	Arch                    string                 `json:"arch" yaml:"arch"`
	CGroupsVersion          string                 `json:"cgroups_version" yaml:"cgroups_version"`
	CPUSpeed                int                    `json:"cpuspeed" yaml:"cpuspeed"`
	Hostname                string                 `json:"hostname" yaml:"hostname"`
	Hypervisor              string                 `json:"hypervisor" yaml:"hypervisor"`
	IMMAD                   string                 `json:"im_mad" yaml:"im_mad"`
	KVMCPUModel             string                 `json:"kvm_cpumodel" yaml:"kvm_cpumodel"`
	KVMCPUModels            []string               `json:"kvm_cpumodels" yaml:"kvm_cpumodels"`
	KVMMachines             []string               `json:"kvm_machines" yaml:"kvm_machines"`
	ModelName               string                 `json:"modelname" yaml:"modelname"`
	ReservedCPU             string                 `json:"reserved_cpu" yaml:"reserved_cpu"`
	ReservedMem             string                 `json:"reserved_mem" yaml:"reserved_mem"`
	TotalWilds              int                    `json:"total_wilds" yaml:"total_wilds"`
	Version                 string                 `json:"version" yaml:"version"`
	Instance                []HostTemplateInstance `json:"instance" yaml:"instance"`
	InstanceMAD             string                 `json:"instance_mad" yaml:"instance_mad"`
	Wilds                   string                 `json:"wilds" yaml:"wilds"`
	VCenterCCRRef           string                 `json:"vcenter_ccrref" yaml:"vcenter_ccrref"`
	VCenterDSRef            []string               `json:"vcenter_dsref" yaml:"vcenter_dsref"`
	VCenterHost             string                 `json:"vcenter_host" yaml:"vcenter_host"`
	VCenterInstanceID       string                 `json:"vcenter_instance_id" yaml:"vcenter_instance_id"`
	VCenterName             string                 `json:"vcenter_name" yaml:"vcenter_name"`
	VCenterPassword         string                 `json:"vcenter_password" yaml:"vcenter_password"`
	VCenterResourcePoolInfo []string               `json:"vcenter_resource_pool_info" yaml:"vcenter_resource_pool_info"`
	VCenterUser             string                 `json:"vcenter_user" yaml:"vcenter_user"`
	VCenterVersion          string                 `json:"vcenter_version" yaml:"vcenter_version"`
}

// HostTemplateInstance is the API payload based on the legacy xmlrpc backend.
type HostTemplateInstance struct {
	Values   map[string]string `json:"values" yaml:"values"`
	DeployID string            `json:"deploy_id" yaml:"deploy_id"`
	ID       int               `json:"id" yaml:"id"`
	Monitor  string            `json:"monitor" yaml:"monitor"`
}
