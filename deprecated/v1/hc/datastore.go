package hc

// DatastoreClusters is the API payload based on the legacy xmlrpc backend.
type DatastoreClusters struct {
	ID []int `json:"id" yaml:"id"`
}

// Datastore is the API payload based on the legacy xmlrpc backend.
type Datastore struct {
	ID          int                  `json:"id" yaml:"id"`
	UID         int                  `json:"uid" yaml:"uid"`
	GID         int                  `json:"gid" yaml:"gid"`
	Uname       string               `json:"uname" yaml:"uname"`
	Gname       string               `json:"gname" yaml:"gname"`
	Name        string               `json:"name" yaml:"name"`
	Permissions DatastorePermissions `json:"permissions" yaml:"permissions"`
	DSMAD       string               `json:"dsmad" yaml:"dsmad"`
	TmMAD       string               `json:"tm_mad" yaml:"tm_mad"`
	BasePath    string               `json:"base_path" yaml:"base_path"`
	Type        int                  `json:"type" yaml:"type"`
	DiskType    int                  `json:"disk_type" yaml:"disk_type"`
	State       int                  `json:"state" yaml:"state"`
	Clusters    DatastoreClusters    `json:"clusters" yaml:"clusters"`
	TotalMB     int                  `json:"total_mb" yaml:"total_mb"`
	FreeMB      int                  `json:"free_mb" yaml:"free_mb"`
	UsedMB      int                  `json:"used_mb" yaml:"used_mb"`
	Images      DatastoreImages      `json:"images" yaml:"images"`
	Template    DatastoreTemplate    `json:"template" yaml:"template"`
}

// DatastoreImages is the API payload based on the legacy xmlrpc backend.
type DatastoreImages struct {
	ID []int `json:"id" yaml:"id"`
}

// DatastorePermissions is the API payload based on the legacy xmlrpc backend.
type DatastorePermissions struct {
	OwnerU int `json:"owner_u" yaml:"owner_u"`
	OwnerM int `json:"owner_m" yaml:"owner_m"`
	OwnerA int `json:"owner_a" yaml:"owner_a"`
	GroupU int `json:"group_u" yaml:"group_u"`
	GroupM int `json:"group_m" yaml:"group_m"`
	GroupA int `json:"group_a" yaml:"group_a"`
	OtherU int `json:"other_u" yaml:"other_u"`
	OtherM int `json:"other_m" yaml:"other_m"`
	OtherA int `json:"other_a" yaml:"other_a"`
}

// DatastoreTemplate is the API payload based on the legacy xmlrpc backend.
type DatastoreTemplate struct {
	Values            map[string]string `json:"values" yaml:"values"`
	VCenterDCName     string            `json:"vcenter_dcname" yaml:"vcenter_dcname"`
	VCenterDCRef      string            `json:"vcenter_dcref" yaml:"vcenter_dcref"`
	VCenterDSName     string            `json:"vcenter_dsname" yaml:"vcenter_dsname"`
	VCenterDSRef      string            `json:"vcenter_dsref" yaml:"vcenter_dsref"`
	VCenterHost       string            `json:"vcenter_host" yaml:"vcenter_host"`
	VCenterInstanceID string            `json:"vcenter_instance_id" yaml:"vcenter_instance_id"`
}
