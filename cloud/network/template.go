package network

import "strconv"

// ParseTemplate return a structured Template based on the given map.
func ParseTemplate(m map[string]any) (*Template, error) {
	var t Template

	for key, value := range m {
		if v, ok := value.(string); ok {
			switch key {
			case "BRIDGE":
				t.Bridge = v
			case "BRIDGE_TYPE":
				t.BridgeType = v
			case "DESCRIPTION":
				t.Description = v
			case "OUTER_VLAN_ID":
				t.OuterVLANID = v
			case "SECURITY_GROUPS":
				t.SecurityGroups = v
			case "VLAN":
				t.VLANID = v
			case "VLAN_ID":
				t.VLANID = v
			case "VN_MAD":
				t.VNMAD = v
			case "DNS":
				t.DNS = v
			case "GATEWAY":
				t.Gateway = v
			case "GATEWAY6":
				t.Gateway6 = v
			case "GUEST_MTU":
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				t.GuestMTU = n
			case "IP6_METHOD":
				t.IP6Method = v
			case "IP6_METRIC":
				t.IP6Metric = v
			case "METHOD":
				t.Method = v
			case "METRIC":
				t.Metric = v
			case "NETWORK_ADDRESS":
				t.NetworkAddress = v
			case "NETWORK_MASK":
				t.NetworkMask = v
			case "SEARCH_DOMAIN":
				t.SearchDomain = v
			case "VCENTER_FROM_WILD":
				t.VCenterFromWild = v
			case "VCENTER_INSTANCE_ID":
				t.VCenterInstanceID = v
			case "VCENTER_NET_REF":
				t.VCenterNetRef = v
			case "VCENTER_PORTGROUP_TYPE":
				t.VCenterPortgroupType = v
			case "VCENTER_TEMPLATE_REF":
				t.VCenterTemplateRef = v
			}
		}
	}

	return &t, nil
}
