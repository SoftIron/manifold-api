package hc

// GroupAdmins is the API payload based on the legacy xmlrpc backend.
type GroupAdmins struct {
	ID []int `json:"id" yaml:"id"`
}

// GroupDatastore is the API payload based on the legacy xmlrpc backend.
type GroupDatastore struct {
	ID         string `json:"id" yaml:"id"`
	Images     string `json:"images" yaml:"images"`
	ImagesUsed string `json:"images_used" yaml:"images_used"`
	Size       string `json:"size" yaml:"size"`
	SizeUsed   string `json:"size_used" yaml:"size_used"`
}

// GroupDatastoreQuota is the API payload based on the legacy xmlrpc backend.
type GroupDatastoreQuota struct {
	Datastore []GroupDatastore `json:"datastore" yaml:"datastore"`
}

// DefaultGroupQuotas is the API payload based on the legacy xmlrpc backend.
type DefaultGroupQuotas struct {
	DatastoreQuota GroupAnon1 `json:"datastore_quota" yaml:"datastore_quota"`
	NetworkQuota   GroupAnon3 `json:"network_quota" yaml:"network_quota"`
	VMQuota        GroupAnon5 `json:"vmquota" yaml:"vmquota"`
	ImageQuota     GroupAnon7 `json:"image_quota" yaml:"image_quota"`
}

// Group is the API payload based on the legacy xmlrpc backend.
type Group struct {
	ID                 int                 `json:"id" yaml:"id"`
	Name               string              `json:"name" yaml:"name"`
	Template           string              `json:"template" yaml:"template"`
	Users              GroupUsers          `json:"users" yaml:"users"`
	Admins             GroupAdmins         `json:"admins" yaml:"admins"`
	DatastoreQuota     GroupDatastoreQuota `json:"datastore_quota" yaml:"datastore_quota"`
	NetworkQuota       GroupNetworkQuota   `json:"network_quota" yaml:"network_quota"`
	VMQuota            GroupVMQuota        `json:"vmquota" yaml:"vmquota"`
	ImageQuota         GroupImageQuota     `json:"image_quota" yaml:"image_quota"`
	DefaultGroupQuotas DefaultGroupQuotas  `json:"default_group_quotas" yaml:"default_group_quotas"`
}

// GroupImage is the API payload based on the legacy xmlrpc backend.
type GroupImage struct {
	ID       string `json:"id" yaml:"id"`
	Rvms     string `json:"rvms" yaml:"rvms"`
	RvmsUsed string `json:"rvms_used" yaml:"rvms_used"`
}

// GroupImageQuota is the API payload based on the legacy xmlrpc backend.
type GroupImageQuota struct {
	Image []GroupImage `json:"image" yaml:"image"`
}

// GroupNetwork is the API payload based on the legacy xmlrpc backend.
type GroupNetwork struct {
	ID         string `json:"id" yaml:"id"`
	Leases     string `json:"leases" yaml:"leases"`
	LeasesUsed string `json:"leases_used" yaml:"leases_used"`
}

// GroupNetworkQuota is the API payload based on the legacy xmlrpc backend.
type GroupNetworkQuota struct {
	Network []GroupNetwork `json:"network" yaml:"network"`
}

// GroupUsers is the API payload based on the legacy xmlrpc backend.
type GroupUsers struct {
	ID []int `json:"id" yaml:"id"`
}

// GroupVM is the API payload based on the legacy xmlrpc backend.
type GroupVM struct {
	CPU                float32 `json:"cpu" yaml:"cpu"`
	CPUUsed            float32 `json:"cpuused" yaml:"cpuused"`
	Memory             int     `json:"memory" yaml:"memory"`
	MemoryUsed         int     `json:"memory_used" yaml:"memory_used"`
	RunningCPU         float32 `json:"running_cpu" yaml:"running_cpu"`
	RunningCPUUsed     float32 `json:"running_cpuused" yaml:"running_cpuused"`
	RunningMemory      int     `json:"running_memory" yaml:"running_memory"`
	RunningMemoryUsed  int     `json:"running_memory_used" yaml:"running_memory_used"`
	RunningVMs         int     `json:"running_vms" yaml:"running_vms"`
	RunningVMsUsed     int     `json:"running_vms_used" yaml:"running_vms_used"`
	SystemDiskSize     int64   `json:"system_disk_size" yaml:"system_disk_size"`
	SystemDiskSizeUsed int64   `json:"system_disk_size_used" yaml:"system_disk_size_used"`
	VMs                int     `json:"vms" yaml:"vms"`
	VMsUsed            int     `json:"vms_used" yaml:"vms_used"`
}

// GroupVMQuota is the API payload based on the legacy xmlrpc backend.
type GroupVMQuota struct {
	VM GroupVM `json:"vm" yaml:"vm"`
}

// GroupAnon1 is the API payload based on the legacy xmlrpc backend.
type GroupAnon1 struct {
	Datastore []GroupAnon2 `json:"datastore" yaml:"datastore"`
}

// GroupAnon2 is the API payload based on the legacy xmlrpc backend.
type GroupAnon2 struct {
	ID         string `json:"id" yaml:"id"`
	Images     string `json:"images" yaml:"images"`
	ImagesUsed string `json:"images_used" yaml:"images_used"`
	Size       string `json:"size" yaml:"size"`
	SizeUsed   string `json:"size_used" yaml:"size_used"`
}

// GroupAnon3 is the API payload based on the legacy xmlrpc backend.
type GroupAnon3 struct {
	Network []GroupAnon4 `json:"network" yaml:"network"`
}

// GroupAnon4 is the API payload based on the legacy xmlrpc backend.
type GroupAnon4 struct {
	ID         string `json:"id" yaml:"id"`
	Leases     string `json:"leases" yaml:"leases"`
	LeasesUsed string `json:"leases_used" yaml:"leases_used"`
}

// GroupAnon5 is the API payload based on the legacy xmlrpc backend.
type GroupAnon5 struct {
	VM GroupAnon6 `json:"vm" yaml:"vm"`
}

// GroupAnon6 is the API payload based on the legacy xmlrpc backend.
type GroupAnon6 struct {
	CPU                float32 `json:"cpu" yaml:"cpu"`
	CPUUsed            float32 `json:"cpuused" yaml:"cpuused"`
	Memory             int     `json:"memory" yaml:"memory"`
	MemoryUsed         int     `json:"memory_used" yaml:"memory_used"`
	RunningCPU         float32 `json:"running_cpu" yaml:"running_cpu"`
	RunningCPUUsed     float32 `json:"running_cpuused" yaml:"running_cpuused"`
	RunningMemory      int     `json:"running_memory" yaml:"running_memory"`
	RunningMemoryUsed  int     `json:"running_memory_used" yaml:"running_memory_used"`
	RunningVMs         int     `json:"running_vms" yaml:"running_vms"`
	RunningVMsUsed     int     `json:"running_vms_used" yaml:"running_vms_used"`
	SystemDiskSize     int64   `json:"system_disk_size" yaml:"system_disk_size"`
	SystemDiskSizeUsed int64   `json:"system_disk_size_used" yaml:"system_disk_size_used"`
	VMs                int     `json:"vms" yaml:"vms"`
	VMsUsed            int     `json:"vms_used" yaml:"vms_used"`
}

// GroupAnon7 is the API payload based on the legacy xmlrpc backend.
type GroupAnon7 struct {
	Image []GroupAnon8 `json:"image" yaml:"image"`
}

// GroupAnon8 is the API payload based on the legacy xmlrpc backend.
type GroupAnon8 struct {
	ID       string `json:"id" yaml:"id"`
	Rvms     string `json:"rvms" yaml:"rvms"`
	RvmsUsed string `json:"rvms_used" yaml:"rvms_used"`
}
