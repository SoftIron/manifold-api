package hc

// HostCapacity is the API payload based on the legacy xmlrpc backend.
type HostCapacity struct {
	FreeCPU    int `json:"free_cpu" yaml:"free_cpu"`
	FreeMemory int `json:"free_memory" yaml:"free_memory"`
	UsedCPU    int `json:"used_cpu" yaml:"used_cpu"`
	UsedMemory int `json:"used_memory" yaml:"used_memory"`
}

// HostCore is the API payload based on the legacy xmlrpc backend.
type HostCore struct {
	Cpus      string `json:"cpus" yaml:"cpus"`
	Dedicated string `json:"dedicated" yaml:"dedicated"`
	Free      int    `json:"free" yaml:"free"`
	ID        int    `json:"id" yaml:"id"`
}

// HostDatastores is the API payload based on the legacy xmlrpc backend.
type HostDatastores struct {
	DiskUsage int `json:"disk_usage" yaml:"disk_usage"`
	FreeDisk  int `json:"free_disk" yaml:"free_disk"`
	MaxDisk   int `json:"max_disk" yaml:"max_disk"`
	UsedDisk  int `json:"used_disk" yaml:"used_disk"`
}

// Host is the API payload based on the legacy xmlrpc backend.
type Host struct {
	ID         int            `json:"id" yaml:"id"`
	Name       string         `json:"name" yaml:"name"`
	State      int            `json:"state" yaml:"state"`
	PrevState  int            `json:"prev_state" yaml:"prev_state"`
	ImMAD      string         `json:"im_mad" yaml:"im_mad"`
	VMMAD      string         `json:"vmmad" yaml:"vmmad"`
	ClusterID  int            `json:"cluster_id" yaml:"cluster_id"`
	Cluster    string         `json:"cluster" yaml:"cluster"`
	HostShare  HostShare      `json:"host_share" yaml:"host_share"`
	VMs        HostVMs        `json:"vms" yaml:"vms"`
	Template   HostTemplate   `json:"template" yaml:"template"`
	Monitoring HostMonitoring `json:"monitoring" yaml:"monitoring"`
}

// HostShare is the API payload based on the legacy xmlrpc backend.
type HostShare struct {
	MemUsage   int            `json:"mem_usage" yaml:"mem_usage"`
	CPUUsage   int            `json:"cpuusage" yaml:"cpuusage"`
	TotalMem   int            `json:"total_mem" yaml:"total_mem"`
	TotalCPU   int            `json:"total_cpu" yaml:"total_cpu"`
	MaxMem     int            `json:"max_mem" yaml:"max_mem"`
	MaxCPU     int            `json:"max_cpu" yaml:"max_cpu"`
	RunningVMs int            `json:"running_vms" yaml:"running_vms"`
	VMsThread  int            `json:"vms_thread" yaml:"vms_thread"`
	Datastores HostDatastores `json:"datastores" yaml:"datastores"`
	PCIDevices HostPCIDevices `json:"pcidevices" yaml:"pcidevices"`
	NumaNodes  HostNumaNodes  `json:"numa_nodes" yaml:"numa_nodes"`
}

// HostHugepage is the API payload based on the legacy xmlrpc backend.
type HostHugepage struct {
	Free  int `json:"free" yaml:"free"`
	Pages int `json:"pages" yaml:"pages"`
	Size  int `json:"size" yaml:"size"`
	Usage int `json:"usage" yaml:"usage"`
}

// HostMemory is the API payload based on the legacy xmlrpc backend.
type HostMemory struct {
	Distance string `json:"distance" yaml:"distance"`
	Free     int    `json:"free" yaml:"free"`
	Total    int    `json:"total" yaml:"total"`
	Usage    int    `json:"usage" yaml:"usage"`
	Used     int    `json:"used" yaml:"used"`
}

// HostMonitoring is the API payload based on the legacy xmlrpc backend.
type HostMonitoring struct {
	Timestamp int          `json:"timestamp" yaml:"timestamp"`
	ID        int          `json:"id" yaml:"id"`
	Capacity  HostCapacity `json:"capacity" yaml:"capacity"`
	System    HostSystem   `json:"system" yaml:"system"`
}

// HostNode is the API payload based on the legacy xmlrpc backend.
type HostNode struct {
	Core     []HostCore     `json:"core" yaml:"core"`
	Hugepage []HostHugepage `json:"hugepage" yaml:"hugepage"`
	Memory   HostMemory     `json:"memory" yaml:"memory"`
	NodeID   int            `json:"node_id" yaml:"node_id"`
}

// HostNumaNodes is the API payload based on the legacy xmlrpc backend.
type HostNumaNodes struct {
	Node []HostNode `json:"node" yaml:"node"`
}

// HostPCI is the API payload based on the legacy xmlrpc backend.
type HostPCI struct {
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
	VMID         int    `json:"vmid" yaml:"vmid"`
}

// HostPCIDevices is the API payload based on the legacy xmlrpc backend.
type HostPCIDevices struct {
	PCI []HostPCI `json:"pci" yaml:"pci"`
}

// HostSystem is the API payload based on the legacy xmlrpc backend.
type HostSystem struct {
	Netrx int `json:"netrx" yaml:"netrx"`
	Nettx int `json:"nettx" yaml:"nettx"`
}

// HostTemplate is the API payload based on the legacy xmlrpc backend.
type HostTemplate struct {
	Values                  map[string]string `json:"values" yaml:"values"`
	VCenterCCRRef           string            `json:"vcenter_ccrref" yaml:"vcenter_ccrref"`
	VCenterDSRef            []string          `json:"vcenter_dsref" yaml:"vcenter_dsref"`
	VCenterHost             string            `json:"vcenter_host" yaml:"vcenter_host"`
	VCenterInstanceID       string            `json:"vcenter_instance_id" yaml:"vcenter_instance_id"`
	VCenterName             string            `json:"vcenter_name" yaml:"vcenter_name"`
	VCenterPassword         string            `json:"vcenter_password" yaml:"vcenter_password"`
	VCenterResourcePoolInfo []string          `json:"vcenter_resource_pool_info" yaml:"vcenter_resource_pool_info"`
	VCenterUser             string            `json:"vcenter_user" yaml:"vcenter_user"`
	VCenterVersion          string            `json:"vcenter_version" yaml:"vcenter_version"`
}

// HostVMs is the API payload based on the legacy xmlrpc backend.
type HostVMs struct {
	ID []int `json:"id" yaml:"id"`
}
