package cloud

import "strings"

// Group is the API payload based on the legacy xmlrpc backend.
type Group struct {
	ID                 int               `json:"id" yaml:"id"`
	Name               string            `json:"name" yaml:"name"`
	Template           Template          `json:"template" yaml:"template"`
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
	Sunstone *SunstoneTemplate
}

// SunstoneTemplate is the API payload based on the legacy xmlrpc backend.
type SunstoneTemplate struct {
	DefaultView           string
	GroupAdminDefaultView string
	GroupAdminViews       []string
	Views                 []string
}

// ParseTemplate return a structured Template based on the given map.
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
