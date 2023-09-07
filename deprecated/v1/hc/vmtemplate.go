package hc

// InstanceTemplateLock is the API payload based on the legacy xmlrpc backend.
type InstanceTemplateLock struct {
	Locked int `json:"locked" yaml:"locked"`
	Owner  int `json:"owner" yaml:"owner"`
	Time   int `json:"time" yaml:"time"`
	ReqID  int `json:"req_id" yaml:"req_id"`
}

// InstanceTemplatePermissions is the API payload based on the legacy xmlrpc backend.
type InstanceTemplatePermissions struct {
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

// VMTemplateData is the API payload based on the legacy xmlrpc backend.
type VMTemplateData struct {
	Values             map[string]string `json:"values" yaml:"values"`
	VCenterCCRRef      string            `json:"vcenter_ccrref" yaml:"vcenter_ccrref"`
	VCenterInstanceID  string            `json:"vcenter_instance_id" yaml:"vcenter_instance_id"`
	VCenterTemplateRef string            `json:"vcenter_template_ref" yaml:"vcenter_template_ref"`
}

// InstanceTemplate is the API payload based on the legacy xmlrpc backend.
type InstanceTemplate struct {
	ID          int                         `json:"id" yaml:"id"`
	UID         int                         `json:"uid" yaml:"uid"`
	GID         int                         `json:"gid" yaml:"gid"`
	Uname       string                      `json:"uname" yaml:"uname"`
	Gname       string                      `json:"gname" yaml:"gname"`
	Name        string                      `json:"name" yaml:"name"`
	Lock        InstanceTemplateLock        `json:"lock" yaml:"lock"`
	Permissions InstanceTemplatePermissions `json:"permissions" yaml:"permissions"`
	Regtime     int                         `json:"regtime" yaml:"regtime"`
	Template    VMTemplateData              `json:"template" yaml:"template"`
}
