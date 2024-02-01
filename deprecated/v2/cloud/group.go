package cloud

// Group is the API payload based on the legacy xmlrpc backend.
type Group struct {
	ID                 int               `json:"id" yaml:"id"`
	Name               string            `json:"name" yaml:"name"`
	Template           GroupTemplate     `json:"template" yaml:"template"`
	TemplateText       string            `json:"template_text" yaml:"template_text"`
	Users              []int             `json:"users" yaml:"users"`
	Admins             []int             `json:"admins" yaml:"admins"`
	DatastoreQuota     []UserDatastore   `json:"datastore_quota" yaml:"datastore_quota"`
	NetworkQuota       []UserNetwork     `json:"network_quota" yaml:"network_quota"`
	InstanceQuota      UserInstance      `json:"instance_quota" yaml:"instance_quota"`
	ImageQuota         []UserImage       `json:"image_quota" yaml:"image_quota"`
	DefaultGroupQuotas UserDefaultQuotas `json:"default_group_quotas" yaml:"default_group_quotas"`
}

// GroupTemplate is the API payload based on the legacy xmlrpc backend.
type GroupTemplate struct {
	Values   map[string]string `json:"values" yaml:"values"`
	Sunstone *SunstoneTemplate `json:"sunstone" yaml:"sunstone"`
}

// SunstoneTemplate is the API payload based on the legacy xmlrpc backend.
type SunstoneTemplate struct {
	Values                map[string]string `json:"values" yaml:"values"`
	DefaultView           string            `json:"default_view" yaml:"default_view"`
	GroupAdminDefaultView string            `json:"group_admin_default_view" yaml:"group_admin_default_view"`
	GroupAdminViews       []string          `json:"group_admin_views" yaml:"group_admin_views"`
	Views                 []string          `json:"views" yaml:"views"`
}
