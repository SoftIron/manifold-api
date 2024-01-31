// Package cloud documents the request and response payloads for version 2 of the sifi API.
package cloud

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/softiron/hypercloud-api/internal/time"
)

// Root path for API endpoint.
const (
	PathPrefix = "cloud"

	AccessControlListPath = "acl"
	ClusterPath           = "cluster"
	DataCenterPath        = "datacenter"
	DatastorePath         = "datastore"
	DocumentPath          = "document"
	GroupPath             = "group"
	HookPath              = "hook"
	HostPath              = "host"
	ImagePath             = "image"
	InstanceGroupPath     = "instance-group"
	InstancePath          = "instance"
	MarketPath            = "market"
	NetworkPath           = "network"
	RouterPath            = "router"
	SecurityGroupPath     = "security-group"
	SystemPath            = "system"
	TemplatePath          = "template"
	UserPath              = "user"
	ZonePath              = "zone"
)

// Lock is the API payload based on the legacy xmlrpc backend.
type Lock struct {
	Locked bool      `json:"locked" yaml:"locked"`
	Owner  int       `json:"owner" yaml:"owner"`
	Time   time.Time `json:"time,omitempty" yaml:"time,omitempty"`
	ReqID  int       `json:"req_id" yaml:"req_id"`
}

// Permissions is the API payload based on the legacy xmlrpc backend.
type Permissions struct {
	Owner perms `json:"owner,omitempty" yaml:"owner,omitempty"`
	Group perms `json:"group,omitempty" yaml:"group,omitempty"`
	Other perms `json:"other,omitempty" yaml:"other,omitempty"`
}

type perms struct {
	Use    bool `json:"use,omitempty" yaml:"use,omitempty"`
	Manage bool `json:"manage,omitempty" yaml:"manage,omitempty"`
	Admin  bool `json:"admin,omitempty" yaml:"admin,omitempty"`
}

// Period is a time interval with optional start and end times.
type Period struct { // TODO: use a time type for v2?
	Start *int `json:"start,omitempty"`
	End   *int `json:"end,omitempty"`
}

// Perms is a set of owner (user), group, and other permissions. Think UNIX.
type Perms struct {
	OwnerUse    *bool `json:"owner_use,omitempty"`
	OwnerManage *bool `json:"owner_manage,omitempty"`
	OwnerAdmin  *bool `json:"owner_admin,omitempty"`
	GroupUse    *bool `json:"group_use,omitempty"`
	GroupManage *bool `json:"group_manage,omitempty"`
	GroupAdmin  *bool `json:"group_admin,omitempty"`
	OtherUse    *bool `json:"other_use,omitempty"`
	OtherManage *bool `json:"other_manage,omitempty"`
	OtherAdmin  *bool `json:"other_admin,omitempty"`
}

// str2Bool converts a string to a bool extending strconv.ParseBool to
// include "yes"/"no". It returns an error if the string is not a
// valid bool.
func str2Bool(s string) (bool, error) {
	if s == "" || strings.EqualFold(s, "no") {
		return false, nil
	}

	if strings.EqualFold(s, "yes") {
		return true, nil
	}

	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, fmt.Errorf("str2bool: %w", err)
	}

	return b, nil
}
