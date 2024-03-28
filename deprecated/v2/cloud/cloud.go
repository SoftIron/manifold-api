// Package cloud documents the request and response payloads for version 2 of the sifi API.
package cloud

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/softiron/manifold-api/internal/api"
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
	Locked bool     `json:"locked" yaml:"locked"`
	Owner  int      `json:"owner" yaml:"owner"`
	Time   api.Time `json:"time" yaml:"time"`
	ReqID  int      `json:"req_id" yaml:"req_id"`
}

// Permissions is the API payload based on the legacy xmlrpc backend.
type Permissions struct {
	Owner perms `json:"owner" yaml:"owner"`
	Group perms `json:"group" yaml:"group"`
	Other perms `json:"other" yaml:"other"`
}

type perms struct {
	Use    bool `json:"use" yaml:"use"`
	Manage bool `json:"manage" yaml:"manage"`
	Admin  bool `json:"admin" yaml:"admin"`
}

// Period is a time interval with optional start and end times.
type Period struct { // TODO: use a time type for v2?
	Start *int `json:"start"`
	End   *int `json:"end"`
}

// Perms is a set of owner (user), group, and other permissions. Think UNIX.
type Perms struct {
	OwnerUse    *bool `json:"owner_use"`
	OwnerManage *bool `json:"owner_manage"`
	OwnerAdmin  *bool `json:"owner_admin"`
	GroupUse    *bool `json:"group_use"`
	GroupManage *bool `json:"group_manage"`
	GroupAdmin  *bool `json:"group_admin"`
	OtherUse    *bool `json:"other_use"`
	OtherManage *bool `json:"other_manage"`
	OtherAdmin  *bool `json:"other_admin"`
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
