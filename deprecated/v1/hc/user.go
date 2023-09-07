package hc

// UserDatastore is the API payload based on the legacy xmlrpc backend.
type UserDatastore struct {
	ID         string `json:"id" yaml:"id"`
	Images     string `json:"images" yaml:"images"`
	ImagesUsed string `json:"images_used" yaml:"images_used"`
	Size       string `json:"size" yaml:"size"`
	SizeUsed   string `json:"size_used" yaml:"size_used"`
}

// UserDatastoreQuota is the API payload based on the legacy xmlrpc backend.
type UserDatastoreQuota struct {
	Datastore []UserDatastore `json:"datastore" yaml:"datastore"`
}

// DefaultUserQuotas is the API payload based on the legacy xmlrpc backend.
type DefaultUserQuotas struct {
	DatastoreQuota UserAnon1 `json:"datastore_quota" yaml:"datastore_quota"`
	NetworkQuota   UserAnon3 `json:"network_quota" yaml:"network_quota"`
	VMQuota        UserAnon5 `json:"vmquota" yaml:"vmquota"`
	ImageQuota     UserAnon7 `json:"image_quota" yaml:"image_quota"`
}

// UserGroups is the API payload based on the legacy xmlrpc backend.
type UserGroups struct {
	ID []int `json:"id" yaml:"id"`
}

// UserImage is the API payload based on the legacy xmlrpc backend.
type UserImage struct {
	ID       string `json:"id" yaml:"id"`
	Rvms     string `json:"rvms" yaml:"rvms"`
	RvmsUsed string `json:"rvms_used" yaml:"rvms_used"`
}

// UserImageQuota is the API payload based on the legacy xmlrpc backend.
type UserImageQuota struct {
	Image []UserImage `json:"image" yaml:"image"`
}

// UserLoginToken is the API payload based on the legacy xmlrpc backend.
type UserLoginToken struct {
	Token          string `json:"token" yaml:"token"`
	ExpirationTime int    `json:"expiration_time" yaml:"expiration_time"`
	Egid           int    `json:"egid" yaml:"egid"`
}

// UserNetwork is the API payload based on the legacy xmlrpc backend.
type UserNetwork struct {
	ID         string `json:"id" yaml:"id"`
	Leases     string `json:"leases" yaml:"leases"`
	LeasesUsed string `json:"leases_used" yaml:"leases_used"`
}

// UserNetworkQuota is the API payload based on the legacy xmlrpc backend.
type UserNetworkQuota struct {
	Network []UserNetwork `json:"network" yaml:"network"`
}

// User is the API payload based on the legacy xmlrpc backend.
type User struct {
	ID                int                `json:"id" yaml:"id"`
	GID               int                `json:"gid" yaml:"gid"`
	Groups            UserGroups         `json:"groups" yaml:"groups"`
	Gname             string             `json:"gname" yaml:"gname"`
	Name              string             `json:"name" yaml:"name"`
	Password          string             `json:"password" yaml:"password"`
	AuthDriver        string             `json:"auth_driver" yaml:"auth_driver"`
	Enabled           int                `json:"enabled" yaml:"enabled"`
	LoginToken        []UserLoginToken   `json:"login_token" yaml:"login_token"`
	Template          string             `json:"template" yaml:"template"`
	DatastoreQuota    UserDatastoreQuota `json:"datastore_quota" yaml:"datastore_quota"`
	NetworkQuota      UserNetworkQuota   `json:"network_quota" yaml:"network_quota"`
	VMQuota           UserVMQuota        `json:"vmquota" yaml:"vmquota"`
	ImageQuota        UserImageQuota     `json:"image_quota" yaml:"image_quota"`
	DefaultUserQuotas DefaultUserQuotas  `json:"default_user_quotas" yaml:"default_user_quotas"`
}

// UserVM is the API payload based on the legacy xmlrpc backend.
type UserVM struct {
	CPU                string `json:"cpu" yaml:"cpu"`
	CPUUsed            string `json:"cpuused" yaml:"cpuused"`
	Memory             string `json:"memory" yaml:"memory"`
	MemoryUsed         string `json:"memory_used" yaml:"memory_used"`
	RunningCPU         string `json:"running_cpu" yaml:"running_cpu"`
	RunningCPUUsed     string `json:"running_cpuused" yaml:"running_cpuused"`
	RunningMemory      string `json:"running_memory" yaml:"running_memory"`
	RunningMemoryUsed  string `json:"running_memory_used" yaml:"running_memory_used"`
	RunningVMs         string `json:"running_vms" yaml:"running_vms"`
	RunningVMsUsed     string `json:"running_vms_used" yaml:"running_vms_used"`
	SystemDiskSize     string `json:"system_disk_size" yaml:"system_disk_size"`
	SystemDiskSizeUsed string `json:"system_disk_size_used" yaml:"system_disk_size_used"`
	VMs                string `json:"vms" yaml:"vms"`
	VMsUsed            string `json:"vms_used" yaml:"vms_used"`
}

// UserVMQuota is the API payload based on the legacy xmlrpc backend.
type UserVMQuota struct {
	VM UserVM `json:"vm" yaml:"vm"`
}

// UserAnon1 is the API payload based on the legacy xmlrpc backend.
type UserAnon1 struct {
	Datastore []UserAnon2 `json:"datastore" yaml:"datastore"`
}

// UserAnon2 is the API payload based on the legacy xmlrpc backend.
type UserAnon2 struct {
	ID         string `json:"id" yaml:"id"`
	Images     string `json:"images" yaml:"images"`
	ImagesUsed string `json:"images_used" yaml:"images_used"`
	Size       string `json:"size" yaml:"size"`
	SizeUsed   string `json:"size_used" yaml:"size_used"`
}

// UserAnon3 is the API payload based on the legacy xmlrpc backend.
type UserAnon3 struct {
	Network []UserAnon4 `json:"network" yaml:"network"`
}

// UserAnon4 is the API payload based on the legacy xmlrpc backend.
type UserAnon4 struct {
	ID         string `json:"id" yaml:"id"`
	Leases     string `json:"leases" yaml:"leases"`
	LeasesUsed string `json:"leases_used" yaml:"leases_used"`
}

// UserAnon5 is the API payload based on the legacy xmlrpc backend.
type UserAnon5 struct {
	VM UserAnon6 `json:"vm" yaml:"vm"`
}

// UserAnon6 is the API payload based on the legacy xmlrpc backend.
type UserAnon6 struct {
	CPU                string `json:"cpu" yaml:"cpu"`
	CPUUsed            string `json:"cpuused" yaml:"cpuused"`
	Memory             string `json:"memory" yaml:"memory"`
	MemoryUsed         string `json:"memory_used" yaml:"memory_used"`
	RunningCPU         string `json:"running_cpu" yaml:"running_cpu"`
	RunningCPUUsed     string `json:"running_cpuused" yaml:"running_cpuused"`
	RunningMemory      string `json:"running_memory" yaml:"running_memory"`
	RunningMemoryUsed  string `json:"running_memory_used" yaml:"running_memory_used"`
	RunningVMs         string `json:"running_vms" yaml:"running_vms"`
	RunningVMsUsed     string `json:"running_vms_used" yaml:"running_vms_used"`
	SystemDiskSize     string `json:"system_disk_size" yaml:"system_disk_size"`
	SystemDiskSizeUsed string `json:"system_disk_size_used" yaml:"system_disk_size_used"`
	VMs                string `json:"vms" yaml:"vms"`
	VMsUsed            string `json:"vms_used" yaml:"vms_used"`
}

// UserAnon7 is the API payload based on the legacy xmlrpc backend.
type UserAnon7 struct {
	Image []UserAnon8 `json:"image" yaml:"image"`
}

// UserAnon8 is the API payload based on the legacy xmlrpc backend.
type UserAnon8 struct {
	ID       string `json:"id" yaml:"id"`
	Rvms     string `json:"rvms" yaml:"rvms"`
	RvmsUsed string `json:"rvms_used" yaml:"rvms_used"`
}
