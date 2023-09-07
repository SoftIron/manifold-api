package hc

// SecurityGroupErrorVMs is the API payload based on the legacy xmlrpc backend.
type SecurityGroupErrorVMs struct {
	ID []int `json:"id" yaml:"id"`
}

// SecurityGroupOutdatedVMs is the API payload based on the legacy xmlrpc backend.
type SecurityGroupOutdatedVMs struct {
	ID []int `json:"id" yaml:"id"`
}

// SecurityGroupPermissions is the API payload based on the legacy xmlrpc backend.
type SecurityGroupPermissions struct {
	OwnerU int `json:"owner_u" yaml:"owner_u"`
	OwnerM int `json:"owner_m" yaml:"owner_m"`
	OwnerA int `json:"owner_a" yaml:"owner_a"`
	GroupU int `json:"group_u" yaml:"group_u"`
	GroupM int `json:"group_m" yaml:"group_m"`
	GroupA int `json:"group_a" yaml:"group_a"`
	OtherU int `json:"other_u" yaml:"other_u"`
	OtherM int `json:"other_m" yaml:"other_m"`
	OtherA int `json:"other_a" yaml:"other_a"`
}

// SecurityGroupRule is the API payload based on the legacy xmlrpc backend.
type SecurityGroupRule struct {
	Protocol string `json:"protocol" yaml:"protocol"`
	RuleType string `json:"rule_type" yaml:"rule_type"`
}

// SecurityGroup is the API payload based on the legacy xmlrpc backend.
type SecurityGroup struct {
	ID          int                      `json:"id" yaml:"id"`
	UID         int                      `json:"uid" yaml:"uid"`
	GID         int                      `json:"gid" yaml:"gid"`
	Uname       string                   `json:"uname" yaml:"uname"`
	Gname       string                   `json:"gname" yaml:"gname"`
	Name        string                   `json:"name" yaml:"name"`
	Permissions SecurityGroupPermissions `json:"permissions" yaml:"permissions"`
	UpdatedVMs  SecurityGroupUpdatedVMs  `json:"updated_vms" yaml:"updated_vms"`
	OutdatedVMs SecurityGroupOutdatedVMs `json:"outdated_vms" yaml:"outdated_vms"`
	UpdatingVMs SecurityGroupUpdatingVMs `json:"updating_vms" yaml:"updating_vms"`
	ErrorVMs    SecurityGroupErrorVMs    `json:"error_vms" yaml:"error_vms"`
	Template    SecurityGroupTemplate    `json:"template" yaml:"template"`
}

// SecurityGroupTemplate is the API payload based on the legacy xmlrpc backend.
type SecurityGroupTemplate struct {
	Values      map[string]string   `json:"values" yaml:"values"`
	Description string              `json:"description" yaml:"description"`
	Rule        []SecurityGroupRule `json:"rule" yaml:"rule"`
}

// SecurityGroupUpdatedVMs is the API payload based on the legacy xmlrpc backend.
type SecurityGroupUpdatedVMs struct {
	ID []int `json:"id" yaml:"id"`
}

// SecurityGroupUpdatingVMs is the API payload based on the legacy xmlrpc backend.
type SecurityGroupUpdatingVMs struct {
	ID []int `json:"id" yaml:"id"`
}
