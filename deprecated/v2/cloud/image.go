package cloud

import (
	"time"

	"github.com/softiron/hypercloud-api/deprecated/v2/cloud/instance"
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
	TemplateText     string             `json:"template_text" yaml:"template_text"`
	Snapshots        instance.Snapshots `json:"snapshots" yaml:"snapshots"`
}

// ImageTemplate is the API payload based on the legacy xmlrpc backend.
type ImageTemplate struct {
	Values          map[string]string `json:"values" yaml:"values"`
	DevPrefix       string            `json:"dev_prefix" yaml:"dev_prefix"`
	Driver          string            `json:"driver" yaml:"driver"`
	FromApp         string            `json:"from_app" yaml:"from_app"`
	FromAppMD5      string            `json:"from_app_md5" yaml:"from_app_md5"`
	FromAppName     string            `json:"from_app_name" yaml:"from_app_name"`
	Size            int               `json:"size" yaml:"size"`
	VCenterImported string            `json:"vcenter_imported" yaml:"vcenter_imported"`
}
