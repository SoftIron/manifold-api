package cloud

import (
	"github.com/softiron/hypercloud-api/cloud/network"
)

// Network is the API payload based on the legacy xmlrpc backend.
type Network struct {
	ID                   int                    `json:"id" yaml:"id"`
	UID                  int                    `json:"uid" yaml:"uid"`
	GID                  int                    `json:"gid" yaml:"gid"`
	UserName             string                 `json:"user_name" yaml:"user_name"`
	GroupName            string                 `json:"group_name" yaml:"group_name"`
	Name                 string                 `json:"name" yaml:"name"`
	Lock                 Lock                   `json:"lock" yaml:"lock"`
	Permissions          Permissions            `json:"permissions" yaml:"permissions"`
	Clusters             []int                  `json:"clusters" yaml:"clusters"`
	Bridge               string                 `json:"bridge" yaml:"bridge"`
	BridgeType           string                 `json:"bridge_type" yaml:"bridge_type"`
	State                int                    `json:"state" yaml:"state"`
	PrevState            int                    `json:"prev_state" yaml:"prev_state"`
	ParentNetworkID      string                 `json:"parent_network_id" yaml:"parent_network_id"`
	VNMAD                string                 `json:"vnmad" yaml:"vnmad"`
	Phydev               string                 `json:"phydev" yaml:"phydev"`
	VLANID               string                 `json:"vlanid" yaml:"vlanid"`
	OuterVLANID          string                 `json:"outer_vlanid" yaml:"outer_vlanid"`
	VLANIDAutomatic      string                 `json:"vlanidautomatic" yaml:"vlanidautomatic"`
	OuterVLANIDAutomatic string                 `json:"outer_vlanidautomatic" yaml:"outer_vlanidautomatic"`
	UsedLeases           int                    `json:"used_leases" yaml:"used_leases"`
	VRouters             []int                  `json:"vrouters" yaml:"vrouters"`
	Template             network.Template       `json:"template" yaml:"template"`
	TemplateText         string                 `json:"template_text" yaml:"template_text"`
	AddressRanges        []network.AddressRange `json:"arpool" yaml:"arpool"`
}
