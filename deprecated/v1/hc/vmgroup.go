package hc

// InstanceGroupLock is the API payload based on the legacy xmlrpc backend.
type InstanceGroupLock struct {
	Locked int `json:"locked" yaml:"locked"`
	Owner  int `json:"owner" yaml:"owner"`
	Time   int `json:"time" yaml:"time"`
	ReqID  int `json:"req_id" yaml:"req_id"`
}

// InstanceGroupPermissions is the API payload based on the legacy xmlrpc backend.
type InstanceGroupPermissions struct {
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

// InstanceGroupRole is the API payload based on the legacy xmlrpc backend.
type InstanceGroupRole struct {
	HostAffined     string `json:"host_affined" yaml:"host_affined"`
	HostAntiAffined string `json:"host_anti_affined" yaml:"host_anti_affined"`
	ID              int    `json:"id" yaml:"id"`
	Name            string `json:"name" yaml:"name"`
	Policy          string `json:"policy" yaml:"policy"`
}

// InstanceGroupRoles is the API payload based on the legacy xmlrpc backend.
type InstanceGroupRoles struct {
	Role []InstanceGroupRole `json:"role" yaml:"role"`
}

// InstanceGroup is the API payload based on the legacy xmlrpc backend.
type InstanceGroup struct {
	ID          int                      `json:"id" yaml:"id"`
	UID         int                      `json:"uid" yaml:"uid"`
	GID         int                      `json:"gid" yaml:"gid"`
	Uname       string                   `json:"uname" yaml:"uname"`
	Gname       string                   `json:"gname" yaml:"gname"`
	Name        string                   `json:"name" yaml:"name"`
	Permissions InstanceGroupPermissions `json:"permissions" yaml:"permissions"`
	Lock        InstanceGroupLock        `json:"lock" yaml:"lock"`
	Roles       InstanceGroupRoles       `json:"roles" yaml:"roles"`
	Template    string                   `json:"template" yaml:"template"`
}
