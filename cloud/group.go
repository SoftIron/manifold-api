package cloud

// Group is the API payload based on the legacy xmlrpc backend.
type Group struct {
	ID                 int               `json:"id" yaml:"id"`
	Name               string            `json:"name" yaml:"name"`
	Template           GroupTemplate     `json:"template" yaml:"template"`
	TemplateText       string            `json:"template_text,omitempty" yaml:"template_text,omitempty"`
	Users              []int             `json:"users,omitempty" yaml:"users,omitempty"`
	Admins             []int             `json:"admins,omitempty" yaml:"admins,omitempty"`
	DatastoreQuota     []UserDatastore   `json:"datastore_quota,omitempty" yaml:"datastore_quota,omitempty"`
	NetworkQuota       []UserNetwork     `json:"network_quota,omitempty" yaml:"network_quota,omitempty"`
	InstanceQuota      UserInstance      `json:"instance_quota" yaml:"instance_quota"`
	ImageQuota         []UserImage       `json:"image_quota,omitempty" yaml:"image_quota,omitempty"`
	DefaultGroupQuotas UserDefaultQuotas `json:"default_group_quotas" yaml:"default_group_quotas"`
}

// GroupTemplate is the API payload based on the legacy xmlrpc backend.
type GroupTemplate struct {
	Values   map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	Sunstone *SunstoneTemplate `json:"sunstone,omitempty" yaml:"sunstone,omitempty"`
}

// SunstoneTemplate is the API payload based on the legacy xmlrpc backend.
type SunstoneTemplate struct {
	Values                map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
	DefaultView           string            `json:"default_view,omitempty" yaml:"default_view,omitempty"`
	GroupAdminDefaultView string            `json:"group_admin_default_view,omitempty" yaml:"group_admin_default_view,omitempty"`
	GroupAdminViews       []string          `json:"group_admin_views,omitempty" yaml:"group_admin_views,omitempty"`
	Views                 []string          `json:"views,omitempty" yaml:"views,omitempty"`
}
