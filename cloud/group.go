package cloud

import (
	"encoding/xml"

	"github.com/softiron/hypercloud-api/cloud/template"
)

// Group is the API payload based on the legacy xmlrpc backend.
type Group struct {
	ID                 int               `json:"id" yaml:"id"`
	Name               string            `json:"name" yaml:"name"`
	Template           GroupTemplate     `json:"template" yaml:"template"`
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
	XMLName  xml.Name          `json:"-" yaml:"-" xml:"TEMPLATE"`
	Values   template.Values   `json:"values,omitempty" yaml:"values,omitempty"`
	Sunstone *SunstoneTemplate `json:"sunstone,omitempty" yaml:"sunstone,omitempty" xml:"SUNSTONE,omitempty"`
}

// SunstoneTemplate is the API payload based on the legacy xmlrpc backend.
type SunstoneTemplate struct {
	Values                template.Values         `json:"values,omitempty" yaml:"values,omitempty"`
	DefaultView           string                  `json:"default_view,omitempty" yaml:"default_view,omitempty" xml:"DEFAULT_VIEW,omitempty"`
	GroupAdminDefaultView string                  `json:"group_admin_default_view,omitempty" yaml:"group_admin_default_view,omitempty" xml:"GROUP_ADMIN_DEFAULT_VIEW,omitempty"`
	GroupAdminViews       []string                `json:"group_admin_views,omitempty" yaml:"group_admin_views,omitempty" xml:"GROUP_ADMIN_VIEWS,omitempty"`
	Views                 template.CommaDelimited `json:"views,omitempty" yaml:"views,omitempty" xml:"VIEWS,omitempty"`
}
