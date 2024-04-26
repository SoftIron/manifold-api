package cloud

// SecurityGroup is the API payload based on the legacy xmlrpc backend.
type SecurityGroup struct {
	ID                int         `json:"id" yaml:"id"`
	UserID            int         `json:"user_id" yaml:"user_id"`
	GroupID           int         `json:"group_id" yaml:"group_id"`
	UserName          string      `json:"user_name" yaml:"user_name"`
	GroupName         string      `json:"group_name" yaml:"group_name"`
	Name              string      `json:"name" yaml:"name"`
	Permissions       Permissions `json:"permissions" yaml:"permissions"`
	UpdatedInstances  []int       `json:"updated_instances" yaml:"updated_instances"`
	OutdatedInstances []int       `json:"outdated_instances" yaml:"outdated_instances"`
	UpdatingInstances []int       `json:"updating_instances" yaml:"updating_instances"`
	ErrorInstances    []int       `json:"error_instances" yaml:"error_instances"`
	Template          Template    `json:"template" yaml:"template"`
}

// SecurityGroupTemplate is the API payload based on the legacy xmlrpc backend.
type SecurityGroupTemplate struct {
	Name        string
	Description string
	Rule        []SecurityGroupRule
}

// SecurityGroupRule is the API payload based on the legacy xmlrpc backend.
type SecurityGroupRule struct {
	ICMPType  string `json:"icmp_type" yaml:"icmp_type"`
	IP        string `json:"ip" yaml:"ip"`
	NetworkID string `json:"network_id" yaml:"network_id"`
	Protocol  string `json:"protocol" yaml:"protocol"`
	Range     string `json:"range" yaml:"range"`
	RuleType  string `json:"rule_type" yaml:"rule_type"`
	Size      string `json:"size" yaml:"size"`
}

// ParseTemplate returns a structured subset of the nested key x value pair map.
func (g *SecurityGroup) ParseTemplate() (*SecurityGroupTemplate, error) {
	var t SecurityGroupTemplate

	for key, value := range g.Template {
		switch v := value.(type) {
		case string:
			if key == "DESCRIPTION" {
				t.Description = v
			}

			if key == "NAME" {
				t.Name = v
			}
		case map[string]any:
			if key == "RULE" {
				t.Rule = []SecurityGroupRule{*newSecurityGroupRule(v)}
			}
		case []map[string]any:
			if key == "RULE" {
				for _, m := range v {
					t.Rule = append(t.Rule, *newSecurityGroupRule(m))
				}
			}
		}
	}

	return &t, nil
}

func newSecurityGroupRule(m map[string]any) *SecurityGroupRule {
	var r SecurityGroupRule

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "PROTOCOL":
				r.Protocol = v
			case "RULE_TYPE":
				r.RuleType = v
			case "IP":
				r.IP = v
			case "RANGE":
				r.Range = v
			case "ICMP_TYPE":
				r.ICMPType = v
			case "NETWORK_ID":
				r.NetworkID = v
			}
		}
	}

	return &r
}
