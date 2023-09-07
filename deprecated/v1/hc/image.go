package hc

// ImageAppClones is the API payload based on the legacy xmlrpc backend.
type ImageAppClones struct {
	ID []int `json:"id" yaml:"id"`
}

// ImageClones is the API payload based on the legacy xmlrpc backend.
type ImageClones struct {
	ID []int `json:"id" yaml:"id"`
}

// Image is the API payload based on the legacy xmlrpc backend.
type Image struct {
	ID             int              `json:"id" yaml:"id"`
	UID            int              `json:"uid" yaml:"uid"`
	GID            int              `json:"gid" yaml:"gid"`
	Uname          string           `json:"uname" yaml:"uname"`
	Gname          string           `json:"gname" yaml:"gname"`
	Name           string           `json:"name" yaml:"name"`
	Lock           ImageLock        `json:"lock" yaml:"lock"`
	Permissions    ImagePermissions `json:"permissions" yaml:"permissions"`
	Type           int              `json:"type" yaml:"type"`
	DiskType       int              `json:"disk_type" yaml:"disk_type"`
	Persistent     int              `json:"persistent" yaml:"persistent"`
	Regtime        int              `json:"regtime" yaml:"regtime"`
	Source         string           `json:"source" yaml:"source"`
	Path           string           `json:"path" yaml:"path"`
	Format         string           `json:"format" yaml:"format"`
	FS             string           `json:"fs" yaml:"fs"`
	Size           int              `json:"size" yaml:"size"`
	State          int              `json:"state" yaml:"state"`
	PrevState      int              `json:"prev_state" yaml:"prev_state"`
	RunningVMs     int              `json:"running_vms" yaml:"running_vms"`
	CloningOps     int              `json:"cloning_ops" yaml:"cloning_ops"`
	CloningID      int              `json:"cloning_id" yaml:"cloning_id"`
	TargetSnapshot int              `json:"target_snapshot" yaml:"target_snapshot"`
	DatastoreID    int              `json:"datastore_id" yaml:"datastore_id"`
	Datastore      string           `json:"datastore" yaml:"datastore"`
	VMs            ImageVMs         `json:"vms" yaml:"vms"`
	Clones         ImageClones      `json:"clones" yaml:"clones"`
	AppClones      ImageAppClones   `json:"app_clones" yaml:"app_clones"`
	Template       ImageTemplate    `json:"template" yaml:"template"`
	Snapshots      ImageSnapshots   `json:"snapshots" yaml:"snapshots"`
}

// ImageLock is the API payload based on the legacy xmlrpc backend.
type ImageLock struct {
	Locked int `json:"locked" yaml:"locked"`
	Owner  int `json:"owner" yaml:"owner"`
	Time   int `json:"time" yaml:"time"`
	ReqID  int `json:"req_id" yaml:"req_id"`
}

// ImagePermissions is the API payload based on the legacy xmlrpc backend.
type ImagePermissions struct {
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

// ImageSnapshot is the API payload based on the legacy xmlrpc backend.
type ImageSnapshot struct {
	Children string `json:"children" yaml:"children"`
	Active   string `json:"active" yaml:"active"`
	Date     int    `json:"date" yaml:"date"`
	ID       int    `json:"id" yaml:"id"`
	Name     string `json:"name" yaml:"name"`
	Parent   int    `json:"parent" yaml:"parent"`
	Size     int    `json:"size" yaml:"size"`
}

// ImageSnapshots is the API payload based on the legacy xmlrpc backend.
type ImageSnapshots struct {
	AllowOrphans string          `json:"allow_orphans" yaml:"allow_orphans"`
	CurrentBase  int             `json:"current_base" yaml:"current_base"`
	NextSnapshot string          `json:"next_snapshot" yaml:"next_snapshot"`
	Snapshot     []ImageSnapshot `json:"snapshot" yaml:"snapshot"`
}

// ImageTemplate is the API payload based on the legacy xmlrpc backend.
type ImageTemplate struct {
	Values          map[string]string `json:"values" yaml:"values"`
	VCenterImported string            `json:"vcenter_imported" yaml:"vcenter_imported"`
}

// ImageVMs is the API payload based on the legacy xmlrpc backend.
type ImageVMs struct {
	ID []int `json:"id" yaml:"id"`
}
