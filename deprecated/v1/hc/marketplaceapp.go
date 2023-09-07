package hc

// MarketplaceAppLock is the API payload based on the legacy xmlrpc backend.
type MarketplaceAppLock struct {
	Locked int `json:"locked" yaml:"locked"`
	Owner  int `json:"owner" yaml:"owner"`
	Time   int `json:"time" yaml:"time"`
	ReqID  int `json:"req_id" yaml:"req_id"`
}

// MarketplaceApp is the API payload based on the legacy xmlrpc backend.
type MarketplaceApp struct {
	ID            int                       `json:"id" yaml:"id"`
	UID           int                       `json:"uid" yaml:"uid"`
	GID           int                       `json:"gid" yaml:"gid"`
	Uname         string                    `json:"uname" yaml:"uname"`
	Gname         string                    `json:"gname" yaml:"gname"`
	Lock          MarketplaceAppLock        `json:"lock" yaml:"lock"`
	Regtime       int                       `json:"regtime" yaml:"regtime"`
	Name          string                    `json:"name" yaml:"name"`
	ZoneID        string                    `json:"zone_id" yaml:"zone_id"`
	OriginID      string                    `json:"origin_id" yaml:"origin_id"`
	Source        string                    `json:"source" yaml:"source"`
	MD5           string                    `json:"md5" yaml:"md5"`
	Size          int                       `json:"size" yaml:"size"`
	Description   string                    `json:"description" yaml:"description"`
	Version       string                    `json:"version" yaml:"version"`
	Format        string                    `json:"format" yaml:"format"`
	Apptemplate64 string                    `json:"apptemplate64" yaml:"apptemplate64"`
	MarketplaceID int                       `json:"marketplace_id" yaml:"marketplace_id"`
	Marketplace   string                    `json:"marketplace" yaml:"marketplace"`
	State         int                       `json:"state" yaml:"state"`
	Type          int                       `json:"type" yaml:"type"`
	Permissions   MarketPlaceAppPermissions `json:"permissions" yaml:"permissions"`
	Template      string                    `json:"template" yaml:"template"`
}

// MarketPlaceAppPermissions is the API payload based on the legacy xmlrpc backend.
type MarketPlaceAppPermissions struct {
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
