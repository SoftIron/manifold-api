package hc

// VNLock is the API payload based on the legacy xmlrpc backend.
type VNLock struct {
	Locked int `json:"locked" yaml:"locked"`
	Owner  int `json:"owner" yaml:"owner"`
	Time   int `json:"time" yaml:"time"`
	ReqID  int `json:"req_id" yaml:"req_id"`
}

// VNPermissions is the API payload based on the legacy xmlrpc backend.
type VNPermissions struct {
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

// VNTemplateData is the API payload based on the legacy xmlrpc backend.
type VNTemplateData struct {
	Values map[string]string `json:"values" yaml:"values"`
	VNMAD  string            `json:"vnmad" yaml:"vnmad"`
}

// VNTemplate is the API payload based on the legacy xmlrpc backend.
type VNTemplate struct {
	ID          int            `json:"id" yaml:"id"`
	UID         int            `json:"uid" yaml:"uid"`
	GID         int            `json:"gid" yaml:"gid"`
	Uname       string         `json:"uname" yaml:"uname"`
	Gname       string         `json:"gname" yaml:"gname"`
	Name        string         `json:"name" yaml:"name"`
	Lock        VNLock         `json:"lock" yaml:"lock"`
	Permissions VNPermissions  `json:"permissions" yaml:"permissions"`
	Regtime     int            `json:"regtime" yaml:"regtime"`
	Template    VNTemplateData `json:"template" yaml:"template"`
}
