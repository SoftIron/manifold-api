package cloud

import (
	"encoding/xml"
	"time"

	"github.com/softiron/hypercloud-api/cloud/instance"
	"github.com/softiron/hypercloud-api/cloud/template"
)

// Image is the API payload based on the legacy xmlrpc backend.
type Image struct {
	ID               int                `json:"id" yaml:"id"`
	UserID           int                `json:"user_id" yaml:"user_id"`
	GroupID          int                `json:"group_id" yaml:"group_id"`
	UserName         string             `json:"user_name" yaml:"user_name"`
	GroupName        string             `json:"group_name" yaml:"group_name"`
	Name             string             `json:"name" yaml:"name"`
	Lock             Lock               `json:"lock" yaml:"lock"`
	Permissions      Permissions        `json:"permissions" yaml:"permissions"`
	Type             int                `json:"type" yaml:"type"`
	DiskType         int                `json:"disk_type" yaml:"disk_type"`
	Persistent       int                `json:"persistent" yaml:"persistent"`
	RegistrationTime time.Time          `json:"registration_time" yaml:"registration_time"`
	Source           string             `json:"source" yaml:"source"`
	Path             string             `json:"path" yaml:"path"`
	Format           string             `json:"format" yaml:"format"`
	Filesystem       string             `json:"filesystem" yaml:"filesystem"`
	Size             int                `json:"size" yaml:"size"`
	State            int                `json:"state" yaml:"state"`
	PrevState        int                `json:"prev_state" yaml:"prev_state"`
	RunningInstances int                `json:"running_instances" yaml:"running_instances"`
	CloningOps       int                `json:"cloning_ops" yaml:"cloning_ops"`
	CloningID        int                `json:"cloning_id" yaml:"cloning_id"`
	TargetSnapshot   int                `json:"target_snapshot" yaml:"target_snapshot"`
	DatastoreID      int                `json:"datastore_id" yaml:"datastore_id"`
	Datastore        string             `json:"datastore" yaml:"datastore"`
	Instances        []int              `json:"instances" yaml:"instances"`
	Clones           []int              `json:"clones" yaml:"clones"`
	AppClones        []int              `json:"app_clones" yaml:"app_clones"`
	Template         ImageTemplate      `json:"template" yaml:"template"`
	Snapshots        instance.Snapshots `json:"snapshots" yaml:"snapshots"`
}

// ImageTemplate is the API payload based on the legacy xmlrpc backend.
type ImageTemplate struct {
	XMLName         xml.Name        `json:"-" yaml:"-" xml:"TEMPLATE"`
	Values          template.Values `json:"values,omitempty" yaml:"values,omitempty"`
	DevPrefix       string          `json:"dev_prefix,omitempty" yaml:"dev_prefix,omitempty" xml:"DEV_PREFIX,omitempty"`
	Driver          string          `json:"driver,omitempty" yaml:"driver,omitempty" xml:"DRIVER,omitempty"`
	FromApp         string          `json:"from_app,omitempty" yaml:"from_app,omitempty" xml:"FROM_APP,omitempty"`
	FromAppMD5      string          `json:"from_app_md5,omitempty" yaml:"from_app_md5,omitempty" xml:"FROM_APP_MD5,omitempty"`
	FromAppName     string          `json:"from_app_name,omitempty" yaml:"from_app_name,omitempty" xml:"FROM_APP_NAME,omitempty"`
	Size            int             `json:"size,omitempty" yaml:"size,omitempty" xml:"SIZE,omitempty"`
	VCenterImported string          `json:"vcenter_imported,omitempty" yaml:"vcenter_imported,omitempty" xml:"VCENTER_IMPORTED,omitempty"`
}
