package cloud

// SecurityGroup is the API payload based on the legacy xmlrpc backend.
type SecurityGroup struct {
	ID                int                   `json:"id" yaml:"id"`
	UserID            int                   `json:"user_id" yaml:"user_id"`
	GroupID           int                   `json:"group_id" yaml:"group_id"`
	UserName          string                `json:"user_name" yaml:"user_name"`
	GroupName         string                `json:"group_name" yaml:"group_name"`
	Name              string                `json:"name" yaml:"name"`
	Permissions       Permissions           `json:"permissions" yaml:"permissions"`
	UpdatedInstances  []int                 `json:"updated_instances" yaml:"updated_instances"`
	OutdatedInstances []int                 `json:"outdated_instances" yaml:"outdated_instances"`
	UpdatingInstances []int                 `json:"updating_instances" yaml:"updating_instances"`
	ErrorInstances    []int                 `json:"error_instances" yaml:"error_instances"`
	Template          SecurityGroupTemplate `json:"template" yaml:"template"`
}

// SecurityGroupTemplate is the API payload based on the legacy xmlrpc backend.
type SecurityGroupTemplate struct {
	Values      map[string]string   `json:"values" yaml:"values"`
	Description string              `json:"description" yaml:"description"`
	Rule        []SecurityGroupRule `json:"rule" yaml:"rule"`
}

// SecurityGroupRule is the API payload based on the legacy xmlrpc backend.
type SecurityGroupRule struct {
	Protocol string `json:"protocol" yaml:"protocol"`
	RuleType string `json:"rule_type" yaml:"rule_type"`
}
