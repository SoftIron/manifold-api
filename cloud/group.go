package cloud

import "strings"

// Group is the API payload based on the legacy xmlrpc backend.
type Group struct {
	ID                 int               `json:"id" yaml:"id"`
	Name               string            `json:"name" yaml:"name"`
	Template           Template          `json:"template" yaml:"template"`
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
	Sunstone *SunstoneTemplate `json:"sunstone,omitempty" yaml:"sunstone,omitempty"`
}

// SunstoneTemplate is the API payload based on the legacy xmlrpc backend.
type SunstoneTemplate struct {
	DefaultView           string   `json:"default_view,omitempty" yaml:"default_view,omitempty"`
	GroupAdminDefaultView string   `json:"group_admin_default_view,omitempty" yaml:"group_admin_default_view,omitempty"`
	GroupAdminViews       []string `json:"group_admin_views,omitempty" yaml:"group_admin_views,omitempty"`
	Views                 []string `json:"views,omitempty" yaml:"views,omitempty"`
}

func (g *Group) ParseTemplate() (*GroupTemplate, error) {
	var t GroupTemplate

	for key, value := range g.Template {
		if m, ok := value.(map[string]any); ok {
			if key == "SUNSTONE" {
				t.Sunstone = newSunstoneTemplate(m)
			}
		}
	}

	return &t, nil
}

func newSunstoneTemplate(m map[string]any) *SunstoneTemplate {
	var t SunstoneTemplate

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "DEFAULT_VIEW":
				t.DefaultView = v
			case "GROUP_ADMIN_DEFAULT_VIEW":
				t.GroupAdminDefaultView = v
			case "GROUP_ADMIN_VIEWS":
				t.GroupAdminViews = strings.Split(v, ",")
			case "VIEWS":
				t.Views = strings.Split(v, ",")
			}
		}
	}

	return &t
}
