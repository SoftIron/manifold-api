package hc

// Marketplace is the API payload based on the legacy xmlrpc backend.
type Marketplace struct {
	ID              int                    `json:"id" yaml:"id"`
	UID             int                    `json:"uid" yaml:"uid"`
	GID             int                    `json:"gid" yaml:"gid"`
	Uname           string                 `json:"uname" yaml:"uname"`
	Gname           string                 `json:"gname" yaml:"gname"`
	Name            string                 `json:"name" yaml:"name"`
	State           int                    `json:"state" yaml:"state"`
	MarketMAD       string                 `json:"market_mad" yaml:"market_mad"`
	ZoneID          string                 `json:"zone_id" yaml:"zone_id"`
	TotalMB         int                    `json:"total_mb" yaml:"total_mb"`
	FreeMB          int                    `json:"free_mb" yaml:"free_mb"`
	UsedMB          int                    `json:"used_mb" yaml:"used_mb"`
	Marketplaceapps MarketplaceApps        `json:"marketplaceapps" yaml:"marketplaceapps"`
	Permissions     MarketplacePermissions `json:"permissions" yaml:"permissions"`
	Template        string                 `json:"template" yaml:"template"`
}

// MarketplaceApps is the API payload based on the legacy xmlrpc backend.
type MarketplaceApps struct {
	ID []int `json:"id" yaml:"id"`
}

// MarketplacePermissions is the API payload based on the legacy xmlrpc backend.
type MarketplacePermissions struct {
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
