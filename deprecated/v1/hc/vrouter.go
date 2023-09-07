package hc

// VRouterLock is the API payload based on the legacy xmlrpc backend.
type VRouterLock struct {
	Locked int `json:"locked" yaml:"locked"`
	Owner  int `json:"owner" yaml:"owner"`
	Time   int `json:"time" yaml:"time"`
	ReqID  int `json:"req_id" yaml:"req_id"`
}

// VRouterPermissions is the API payload based on the legacy xmlrpc backend.
type VRouterPermissions struct {
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

// VRouterVMs is the API payload based on the legacy xmlrpc backend.
type VRouterVMs struct {
	ID []int `json:"id" yaml:"id"`
}

// VRouter is the API payload based on the legacy xmlrpc backend.
type VRouter struct {
	ID          int                `json:"id" yaml:"id"`
	UID         int                `json:"uid" yaml:"uid"`
	GID         int                `json:"gid" yaml:"gid"`
	Uname       string             `json:"uname" yaml:"uname"`
	Gname       string             `json:"gname" yaml:"gname"`
	Name        string             `json:"name" yaml:"name"`
	Permissions VRouterPermissions `json:"permissions" yaml:"permissions"`
	Lock        VRouterLock        `json:"lock" yaml:"lock"`
	VMs         VRouterVMs         `json:"vms" yaml:"vms"`
	Template    string             `json:"template" yaml:"template"`
}
