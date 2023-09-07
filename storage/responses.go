package storage

// ListHostsResponse is the response body for GET /storage/host.
type ListHostsResponse struct {
	Nodes []ListHostResponse `json:"nodes"`
}

// ListShareRBDResponse is the response body for GET /storage/share/rbd/{id}.
type ListShareRBDResponse struct {
	RBD RbdInfo `json:"rbd"`
}

// ListShareRBDsResponse is the response body for GET /storage/share/rbd.
type ListShareRBDsResponse struct {
	RBDs []RbdInfo `json:"rbd"`
}

// ListHostDisksResponse is the response body for GET /storage/host/{id}/disk.
type ListHostDisksResponse struct {
	Disks []DiskInfo `json:"disks"`
}

// ListHostDiskResponse is the response body for GET /storage/host/{id}/disk/{id}.
type ListHostDiskResponse struct {
	Disk DiskInfo `json:"disk"`
}

// ListHostVolumeGroupsResponse is the response body for GET /storage/host/{id}/volgroup.
type ListHostVolumeGroupsResponse struct {
	VolumeGroups []VolumeGroupInfo `json:"volume_groups"`
}

// ListHostVolumeGroupResponse is the response body for GET /storage/host/{id}/volgroup/{id}.
type ListHostVolumeGroupResponse struct {
	VolumeGroup VolumeGroupInfo `json:"volume_group"`
}

// ListDisksResponse is the response body for GET /storage/disk.
type ListDisksResponse struct {
	Disks []DiskInfo `json:"disks"`
}

// ListPoolsResponse is the response body for GET /storage/pool.
type ListPoolsResponse struct {
	Pools []PoolInfo `json:"pools"`
}

// ListPoolResponse is the response body for GET /storage/pool/{id}.
type ListPoolResponse struct {
	Pool PoolInfo `json:"pool"`
}

// ListOSDsResponse is the response body for GET /storage/osd.
type ListOSDsResponse struct {
	OSDs []OsdInfo `json:"osds"`
}

// ListOSDResponse is the response body for GET /storage/osd/{id}.
type ListOSDResponse struct {
	OSD OsdInfo `json:"osd"`
}

// ListSummaryResponse is the summary information to display on the dashboard.
type ListSummaryResponse struct {
	Health                    string      `json:"health" enums:"HEALTH_OK,HEALTH_WARN,HEALTH_ERROR"`
	PgCount                   uint        `json:"pg_count"`
	AveragePgsPerOsd          uint        `json:"average_pgs_per_osd"`
	PoolCount                 uint        `json:"pool_count"`
	ObjectCount               uint64      `json:"object_count"`
	StorageTotalMB            uint64      `json:"storage_total_MB"`
	StorageUsedMB             uint64      `json:"storage_used_MB"`
	StorageUtilizationPercent uint        `json:"storage_utilization_percent"`
	OsdCount                  uint        `json:"osd_count"` // The total number of OSDs inthe cluster
	OsdUp                     uint        `json:"osd_up"`
	OsdDown                   uint        `json:"osd_down"`
	OsdIn                     uint        `json:"osd_in"`
	OsdOut                    uint        `json:"osd_out"`
	Flags                     []CephFlag  `json:"flags"`
	Alerts                    []CephAlert `json:"alerts"`
	AlertCount                uint        `json:"alert_count"`
	MonitorsInQuorum          uint        `json:"monitors_in_quorum"`
	VersionName               string      `json:"version_name"`   // The name of the Ceph version
	VersionNumber             string      `json:"version_number"` // The Ceph version
}

// CephFlag is a single Ceph status flag, eg NO_OUT.
type CephFlag struct {
	Name        string       `json:"name"`   // The name of the flag
	IsSet       bool         `json:"is_set"` // True if this flag has been set, false otherwise
	Type        FlagTypeEnum `json:"type" swaggertype:"string" enums:"PRIORITY,EXPECTED,INFORMATIONAL"`
	Description string       `json:"description"`
}

// CephAlert is a single alert message from Ceph.
type CephAlert struct {
	Severity string `json:"severity" enums:"HEALTH_OK,HEALTH_WARN,HEALTH_ERROR"`
	Message  string `json:"message"`
}

// RbdInfo is the full information for a single RBD.
type RbdInfo struct {
	ID                     string                  `json:"id"`
	Name                   string                  `json:"name"`
	Path                   string                  `json:"path"`
	PoolID                 uint                    `json:"pool_id"`
	PoolName               string                  `json:"pool_name"`
	PoolType               string                  `json:"pool_type"`
	PoolCapacityMB         uint64                  `json:"pool_capacity_MB"`
	PoolUsedMB             uint64                  `json:"pool_used_MB"`
	PoolErasureCodeProfile string                  `json:"pool_erasure_code_profile,omitempty"`
	PoolReplicationCount   uint                    `json:"pool_replication_count,omitempty"`
	PoolShareCount         uint                    `json:"pool_share_count"`
	CapacityMB             uint64                  `json:"capacity_MB"`
	ObjectSizeB            uint                    `json:"object_size_B"`
	ObjectCount            uint64                  `json:"object_count"`
	BlockName              string                  `json:"block_name"`
	Features               []string                `json:"features,omitempty"`
	Deleting               bool                    `json:"deleting"` // True if this RBD is being deleted
	IscsiExportable        IscsiExportableTypeEnum `json:"iscsi_exportable" swaggertype:"string" enums:"AVAILABLE,EXPORTED,INVALID_OPTIONS"`
	IscsiShareID           string                  `json:"iscsi_share_id"`
	IscsiShareName         string                  `json:"iscsi_share_name"`
}

// ListHostResponse is the full information we can get for a single node.
type ListHostResponse struct {
	Alerts                    []string           `json:"alerts"`          // A list of alerts firing for this node
	Asset                     string             `json:"asset,omitempty"` // User defined asset tag for the node
	BaseboardID               int                `json:"baseboard_id"`
	BaseboardRev              int                `json:"baseboard_rev"`
	Bcaches                   []BcacheInfo       `json:"bcaches"`
	BootDisk                  *DiskSummary       `json:"boot_disk,omitempty"`
	CaddyIDs                  map[uint]CaddyInfo `json:"caddy_ids"`
	CaddyInfo                 bool               `json:"caddy_info"`
	Capabilities              []string           `json:"capabilities"` // The capabilities this node has
	CephFips                  bool               `json:"ceph_fips"`
	CephMajor                 string             `json:"ceph_major"`
	CephVersion               string             `json:"ceph_version"`
	ChassisType               string             `json:"chassis_type"` // This node's chassis type
	Disks                     []DiskInfo         `json:"disks"`
	Fans                      []FanInfo          `json:"fans"`
	FirmwareBmc               string             `json:"firmware_bmc,omitempty"` // The version of the firmware this node's BMC is running
	FirmwareSoc               string             `json:"firmware_soc,omitempty"` // The version of the firmware this node's SoC is running
	ID                        string             `json:"id"`                     // A unique ID for the node
	IdentifyLed               bool               `json:"identify_led"`           // True if this node's identify LED is on
	InCluster                 bool               `json:"in_cluster"`             // True if this node is in a Ceph cluster
	IPAddress                 string             `json:"ip_address"`             // The public IP address of this node
	LastSeen                  int64              `json:"last_seen,omitempty"`    // The last time we saw this node up (Unix time)
	MainboardID               int                `json:"mainboard_id"`
	MainboardRev              int                `json:"mainboard_rev"`
	MemoryGB                  uint               `json:"memory_GB,omitempty"` // The amount of RAM this node has, in GB
	Model                     string             `json:"model,omitempty"`     // This node's model
	Name                      string             `json:"name"`                // The node's name
	DNSIP                     []string           `json:"dns"`                 // The network's DNS resolvers
	Ntp                       []string           `json:"time_server"`         // The location of the network's NTP time servers
	Networks                  NetInterfacesInfo  `json:"networks"`
	Nic1ID                    int                `json:"nic_id_1"`
	Nic1Rev                   int                `json:"nic_rev_1"`
	Nic2ID                    int                `json:"nic_id_2"`
	Nic2Rev                   int                `json:"nic_rev_2"`
	OperatingSystem           string             `json:"operating_system,omitempty"`
	OsVersion                 string             `json:"os_version"`
	OsdCount                  uint               `json:"osd_count"`
	PendingChange             bool               `json:"pending_change"`
	PowerW                    uint               `json:"power_W,omitempty"`                                                                                                                // The average power draw of this node, in watts
	Roles                     []string           `json:"roles"`                                                                                                                            // The roles this node has been given
	SerialNo                  string             `json:"serial_no"`                                                                                                                        // This node's serial number
	Status                    NodeStatusEnum     `json:"status" swaggertype:"string" enums:"DISCOVERING,MISSING,OFF,CRASHED,DOWN,ON,PENDING_ON,PENDING_OFF,PENDING_REBOOT,PENDING_CHANGE"` // The node status
	StorageTotalMB            uint64             `json:"storage_total_MB"`                                                                                                                 // Amount of storage this node has
	StorageUsedMB             uint64             `json:"storage_used_MB"`                                                                                                                  // Amount of this node's storage that is being used
	StorageUtilizationPercent uint               `json:"storage_utilization_percent"`                                                                                                      // The percentage of this node's storage that is being used
	TemperatureC              uint               `json:"temperature_C"`                                                                                                                    // The average temperature of this node in centigrade
	Timestamp                 string             `json:"timestamp"`
	UpSince                   string             `json:"up_since"`
	VolumeGroups              []VolumeGroupInfo  `json:"volume_groups"`
}

// NetInterfacesInfo is the information we can get for all the networks on a single node.
type NetInterfacesInfo struct {
	Bmc         NetInterfaceInfo `json:"management"`
	Public      NetInterfaceInfo `json:"public"`
	Replication NetInterfaceInfo `json:"replication"`
}

// NetInterfaceInfo is the information we can get for a single network interface.
type NetInterfaceInfo struct {
	Name         string                `json:"name"`              // This network's name
	IPAddress    string                `json:"ip_address"`        // This network's IP address
	Subnet       string                `json:"subnet,omitempty"`  // The network's subnet
	Gateway      string                `json:"gateway,omitempty"` // The network's gateway
	Mac          string                `json:"mac,omitempty"`     // The network's MAC address
	Mtu          uint16                `json:"mtu,omitempty"`     // The network's maximum transmission unit
	LinkSpeedMbs uint                  `json:"link_speed_Mbs,omitempty"`
	Status       NetworkConnStatusEnum `json:"status" swaggertype:"string" enums:"CONNECTED,NOT_CONNECTED"` // The connection status
}

// BcacheInfo is the information we can get for a single bcache device.
type BcacheInfo struct {
	Cache      []DiskOrPartitionRef `json:"caches"`
	CapacityMB uint64               `json:"capacity_MB"`
	Mountpoint string               `json:"mountpoint"`
	Path       string               `json:"path"`
	Storage    []DiskOrPartitionRef `json:"stores"`
	Usage      string               `json:"usage"`
	UsedBy     []OsdRef             `json:"used_by"`
}

// DiskOrPartitionRef is the disks and partitions used by a bcache device.
type DiskOrPartitionRef struct {
	DiskID         string `json:"disk_id"`
	PartitionIndex uint   `json:"partition_index"`
}

// OsdRef is a reference to a single OSD.
type OsdRef struct {
	Name string `json:"name"` // A human-readable nice name for the OSD.
	ID   uint   `json:"id"`   // Sortable ID for the OSD.
}

// DiskSummary is the information for a single disk.
type DiskSummary struct {
	CapacityMB  uint64 `json:"capacity_MB"` // The max capacity of the disk, in MB
	ID          string `json:"id"`          // The disk opaque identifier
	Name        string `json:"name"`        // A human-readable nice name for the disk
	Path        string `json:"path"`        // The path to the disk device
	MediaType   string `json:"type"`
	SmartPassed bool   `json:"smart_passed"`
	Usage       string `json:"usage"`
}

// CaddyInfo is the information for a single caddy.
type CaddyInfo struct {
	ID  int `json:"id"`
	Rev int `json:"revision"`
}

// DiskInfo is the information for a single disk.
type DiskInfo struct {
	Boot            bool                     `json:"boot"` // True if this disk is the boot disk for the node
	CaddyNo         uint                     `json:"caddy_no"`
	CaddyType       string                   `json:"caddy_type"`
	LocationInCaddy uint                     `json:"location_in_caddy"`
	CapacityMB      uint64                   `json:"capacity_MB"` // The max capacity of the disk, in MB.
	ID              string                   `json:"id"`          // The disk opaque identifier
	Name            string                   `json:"name"`        // A human-readable nice name for the disk
	NodeID          string                   `json:"node_id"`     // The ID of the node this disk is inside
	NodeModel       string                   `json:"node_model"`  // The model of the node this disk is inside
	NodeName        string                   `json:"node_name"`   // The name of the node this disk is inside
	Number          uint                     `json:"number"`
	Osds            []OsdRef                 `json:"osds"`       // A list of the OSDs associated with this disk
	Partitions      []PartitionInfo          `json:"partitions"` // A list of partitions
	Path            string                   `json:"path"`       // The path to the disk device
	Sata            string                   `json:"sata"`
	SerialNo        string                   `json:"serial_no"` // This disk's serial number
	SmartPassed     bool                     `json:"smart_passed"`
	SmartMessage    string                   `json:"smart_message"`
	SmartStatus     uint                     `json:"smart_status"`
	MediaType       string                   `json:"type"`
	TemperatureC    uint                     `json:"temperature_C"`
	Usage           string                   `json:"usage"`
	VolumeGroups    []PartialVolumeGroupInfo `json:"volume_groups"`
	Wwn             string                   `json:"wwn"` // A unique ID for this disk
}

// PartitionInfo is the information for a single disk partition.
type PartitionInfo struct {
	CapacityMB  uint64                 `json:"capacity_MB"` // The max capacity of the partition, in MB.
	Index       uint                   `json:"index"`
	Mountpoint  string                 `json:"mountpoint"`
	Path        string                 `json:"path"` // The path to the partition on the disk
	Usage       string                 `json:"usage"`
	UsedBy      []OsdRef               `json:"used_by"` // The OSDs that use this partition
	UUID        string                 `json:"uuid"`    // A unique ID for this partition
	VolumeGroup PartialVolumeGroupInfo `json:"volume_group"`
}

// PartialVolumeGroupInfo is a subset of information for a volume group that's
// relevant to a single physical volume.
type PartialVolumeGroupInfo struct {
	Name           string              `json:"name"`            // The volume group name
	LogicalVolumes []LogicalVolumeInfo `json:"logical_volumes"` // A list of logical volumes
}

// LogicalVolumeInfo is the information for a single logical volume.
type LogicalVolumeInfo struct {
	CapacityMB      uint64               `json:"capacity_MB"` // The max capacity of this logical volume, in MB.
	Mountpoint      string               `json:"mountpoint"`
	Name            string               `json:"name"`      // The logical volume name
	NodeID          string               `json:"node_id"`   // The ID of the node this logical volume is in
	NodeName        string               `json:"node_name"` // The name of the node this logical volume is inside
	Path            string               `json:"path"`      // The path to the logical volume
	Usage           string               `json:"usage"`
	UsedBy          OsdRef               `json:"used_by"`
	UUID            string               `json:"uuid"`             // A unique ID for this partition
	PhysicalVolumes []PhysicalVolumeInfo `json:"physical_volumes"` // The physical volumes used by this logical volume
}

// PhysicalVolumeInfo is the information for a single physical volume.
type PhysicalVolumeInfo struct {
	Bcache         string   `json:"bcache_path"`
	CapacityMB     uint64   `json:"capacity_MB"`     // The max capacity of the physical volume, in MB
	DiskID         string   `json:"disk_id"`         // The disk opaque identifier
	DiskName       string   `json:"disk_name"`       // A human-readable nice name for the disk
	PartitionIndex uint     `json:"partition_index"` // Index of the partition within the disk, 0 for a whole disk
	Path           string   `json:"path"`            // The path to the physical volume device
	Usage          string   `json:"usage"`
	UsedBy         []OsdRef `json:"used_by"` // The OSDs that use this physical volume
}

// FanInfo is the information for a single fan.
type FanInfo struct {
	ID       int  `json:"id"`
	Reserved bool `json:"reserved"` // True if this fan is in reserved mode
	Speed    int  `json:"speed"`    // Current speed fan is running at, in RPM.
}

// VolumeGroupInfo is the information for a single volume group.
type VolumeGroupInfo struct {
	CapacityMB      uint64               `json:"capacity_MB"` // The max capacity of this volume group, in MB.
	Name            string               `json:"name"`        // The volume group name
	NodeID          string               `json:"node_id"`     // The ID of the node this volume groupis in
	NodeName        string               `json:"node_name"`   // The name of the node this volume group is inside
	Usage           string               `json:"usage"`
	UsedBy          []OsdRef             `json:"used_by"`          // The OSDs that use this volume group
	LogicalVolumes  []LogicalVolumeInfo  `json:"logical_volumes"`  // A list of logical volumes
	PhysicalVolumes []PhysicalVolumeInfo `json:"physical_volumes"` // The physical volumes in this volume group
}

// PoolInfo is the full information about a single Ceph pool.
type PoolInfo struct {
	ID                 uint              `json:"id"`
	Name               string            `json:"name"`
	PgCount            uint              `json:"pg_count"`
	PgPlacementNum     uint              `json:"pg_placement_num"`
	PgAutoscale        bool              `json:"pg_autoscale"`
	Type               string            `json:"type"`
	CapacityMB         uint64            `json:"capacity_MB"`
	UsedMB             uint64            `json:"used_MB"`
	UtilizationPercent uint              `json:"utilization_percent"`
	ObjectCount        uint64            `json:"object_count"`
	Compression        string            `json:"compression"`
	ErasureCodeProfile string            `json:"erasure_code_profile,omitempty"`
	ReplicationCount   uint              `json:"replication_count,omitempty"`
	CrushRulesetID     uint              `json:"crush_ruleset_id"`
	CrushRuleset       string            `json:"crush_ruleset"`
	StripeWidth        uint              `json:"stripe_width"`
	QuotaMaxBytes      uint64            `json:"quota_max_bytes"`
	QuotaMaxObjects    uint64            `json:"quota_max_objects"`
	Applications       []string          `json:"applications"`
	Services           []RestServiceInfo `json:"services"`
	Shares             []RestShareInfo   `json:"shares"`
}

// RestServiceInfo is the  information for a single service of any type.
type RestServiceInfo struct {
	ID     string         `json:"id"`   // A unique ID for the service
	Type   ShareTypeEnum  `json:"type"` // The service's share type
	Name   string         `json:"name"`
	Shares []RestShareRef `json:"shares"`
}

// RestShareInfo is the common information for a single share of any type.
type RestShareInfo struct {
	ID          string        `json:"id"`   // A unique ID for the share
	Type        ShareTypeEnum `json:"type"` // The share's type
	Name        string        `json:"name"` // The share's name
	SizeMB      uint64        `json:"size_MB"`
	ServiceID   string        `json:"service_id"`
	ServiceName string        `json:"service_name"`
}

// RestShareRef references to shares used by services.
type RestShareRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OsdInfo is the full information about a single Ceph OSD.
type OsdInfo struct {
	ID                 uint          `json:"id"`          // A unique ID for the OSD
	Name               string        `json:"name"`        // The OSD's name
	CapacityMB         uint64        `json:"capacity_MB"` // Amount of storage this OSD has
	NodeID             string        `json:"node_id"`     // A unique ID for the node
	NodeName           string        `json:"node_name"`   // The containing node's name
	State              OsdStateEnum  `json:"state" swaggertype:"string" enums:"UP,DOWN,OUT"`
	UsedMB             uint64        `json:"used_MB"`             // Amount of this OSD's storage that is being used
	UtilizationPercent uint          `json:"utilization_percent"` // The percentage of this OSD's total storage that is being used
	CrushWeight        float64       `json:"crush_weight"`
	Reweight           float64       `json:"reweight"`
	PgCount            uint          `json:"pg_count"`
	Variance           float64       `json:"variance"`
	NodeIP             string        `json:"node_ip"` // The public IP address of this node
	Class              OsdClassEnum  `json:"class" swaggertype:"string" enums:"SSD,HDD"`
	Alerts             []string      `json:"alerts"` // A list of alerts being fired for this OSD
	Disks              []DiskSummary `json:"disks"`
}
