// Package config contains struct for the Configuration payload.
package config

// AuthMAD is the API payload based on the legacy xmlrpc backend.
type AuthMAD struct {
	Authn      string `json:"authn" yaml:"authn"`
	Executable string `json:"executable" yaml:"executable"`
}

// AuthMADConf is the API payload based on the legacy xmlrpc backend.
type AuthMADConf struct {
	DriverManagedGroups     string `json:"driver_managed_groups" yaml:"driver_managed_groups"`
	DriverManagedGroupAdmin string `json:"driver_managed_group_admin" yaml:"driver_managed_group_admin"`
	MaxTokenTime            int    `json:"max_token_time" yaml:"max_token_time"`
	Name                    string `json:"name" yaml:"name"`
	PasswordChange          string `json:"password_change" yaml:"password_change"`
}

// DatastoreMAD is the API payload based on the legacy xmlrpc backend.
type DatastoreMAD struct {
	Arguments  string `json:"arguments" yaml:"arguments"`
	Executable string `json:"executable" yaml:"executable"`
}

// DB is the API payload based on the legacy xmlrpc backend.
type DB struct {
	Backend       string `json:"backend" yaml:"backend"`
	CompareBinary string `json:"compare_binary" yaml:"compare_binary"`
	Connections   int    `json:"connections" yaml:"connections"`
	DBName        string `json:"dbname" yaml:"dbname"`
	Passwd        string `json:"passwd" yaml:"passwd"`
	Port          int    `json:"port" yaml:"port"`
	Server        string `json:"server" yaml:"server"`
	User          string `json:"user" yaml:"user"`
	Timeout       int    `json:"timeout" yaml:"timeout"`
}

// DefaultCost is the API payload based on the legacy xmlrpc backend.
type DefaultCost struct {
	CPUCost    int `json:"cpucost" yaml:"cpucost"`
	DiskCost   int `json:"disk_cost" yaml:"disk_cost"`
	MemoryCost int `json:"memory_cost" yaml:"memory_cost"`
}

// DSMADConf is the API payload based on the legacy xmlrpc backend.
type DSMADConf struct {
	MarketplaceActions string `json:"marketplace_actions" yaml:"marketplace_actions"`
	Name               string `json:"name" yaml:"name"`
	PersistentOnly     string `json:"persistent_only" yaml:"persistent_only"`
	RequiredAttrs      string `json:"required_attrs" yaml:"required_attrs"`
}

// Federation is the API payload based on the legacy xmlrpc backend.
type Federation struct {
	MasterOned string `json:"master_oned" yaml:"master_oned"`
	Mode       string `json:"mode" yaml:"mode"`
	ServerID   int    `json:"server_id" yaml:"server_id"`
	ZoneID     int    `json:"zone_id" yaml:"zone_id"`
}

// HookManagerMAD is the API payload based on the legacy xmlrpc backend.
type HookManagerMAD struct {
	Arguments  string `json:"arguments" yaml:"arguments"`
	Executable string `json:"executable" yaml:"executable"`
}

// HookLogConf is the API payload based on the legacy xmlrpc backend.
type HookLogConf struct {
	LogRetention int `json:"log_retention" yaml:"log_retention"`
}

// ImMAD is the API payload based on the legacy xmlrpc backend.
type ImMAD struct {
	Arguments  string `json:"arguments" yaml:"arguments"`
	Executable string `json:"executable" yaml:"executable"`
	Name       string `json:"name" yaml:"name"`
	Threads    int    `json:"threads" yaml:"threads"`
}

// IpamMAD is the API payload based on the legacy xmlrpc backend.
type IpamMAD struct {
	Arguments  string `json:"arguments" yaml:"arguments"`
	Executable string `json:"executable" yaml:"executable"`
}

// Log is the API payload based on the legacy xmlrpc backend.
type Log struct {
	DebugLevel           int    `json:"debug_level" yaml:"debug_level"`
	System               string `json:"system" yaml:"system"`
	UseInstancesLocation string `json:"use_instances_location" yaml:"use_instances_location"`
}

// MarketMAD is the API payload based on the legacy xmlrpc backend.
type MarketMAD struct {
	Arguments  string `json:"arguments" yaml:"arguments"`
	Executable string `json:"executable" yaml:"executable"`
}

// MarketMADConf is the API payload based on the legacy xmlrpc backend.
type MarketMADConf struct {
	AppActions    string `json:"app_actions" yaml:"app_actions"`
	Name          string `json:"name" yaml:"name"`
	Public        string `json:"public" yaml:"public"`
	RequiredAttrs string `json:"required_attrs" yaml:"required_attrs"`
	SunstoneName  string `json:"sunstone_name" yaml:"sunstone_name"`
}

// Raft is the API payload based on the legacy xmlrpc backend.
type Raft struct {
	BroadcastTimeoutMS int `json:"broadcast_timeout_ms" yaml:"broadcast_timeout_ms"`
	ElectionTimeoutMS  int `json:"election_timeout_ms" yaml:"election_timeout_ms"`
	LimitPurge         int `json:"limit_purge" yaml:"limit_purge"`
	LogPurgeTimeout    int `json:"log_purge_timeout" yaml:"log_purge_timeout"`
	LogRetention       int `json:"log_retention" yaml:"log_retention"`
	XmlrpcTimeoutMS    int `json:"xmlrpc_timeout_ms" yaml:"xmlrpc_timeout_ms"`
}

// TmMAD is the API payload based on the legacy xmlrpc backend.
type TmMAD struct {
	Arguments  string `json:"arguments" yaml:"arguments"`
	Executable string `json:"executable" yaml:"executable"`
}

// TmMADConf is the API payload based on the legacy xmlrpc backend.
type TmMADConf struct {
	AllowOrphans      string `json:"allow_orphans" yaml:"allow_orphans"`
	CloneTarget       string `json:"clone_target" yaml:"clone_target"`
	CloneTargetShared string `json:"clone_target_shared" yaml:"clone_target_shared"`
	CloneTargetSSH    string `json:"clone_target_ssh" yaml:"clone_target_ssh"`
	DiskTypeShared    string `json:"disk_type_shared" yaml:"disk_type_shared"`
	DiskTypeSSH       string `json:"disk_type_ssh" yaml:"disk_type_ssh"`
	Driver            string `json:"driver" yaml:"driver"`
	DSMigrate         string `json:"dsmigrate" yaml:"dsmigrate"`
	LnTarget          string `json:"ln_target" yaml:"ln_target"`
	LnTargetShared    string `json:"ln_target_shared" yaml:"ln_target_shared"`
	LnTargetSSH       string `json:"ln_target_ssh" yaml:"ln_target_ssh"`
	Name              string `json:"name" yaml:"name"`
	Shared            string `json:"shared" yaml:"shared"`
	TmMADSystem       string `json:"tm_madsystem" yaml:"tm_madsystem"`
}

// VLANIDs is the API payload based on the legacy xmlrpc backend.
type VLANIDs struct {
	Reserved string `json:"reserved" yaml:"reserved"`
	Start    int    `json:"start" yaml:"start"`
}

// InstanceMAD is the API payload based on the legacy xmlrpc backend.
type InstanceMAD struct {
	Arguments                string `json:"arguments" yaml:"arguments"`
	Default                  string `json:"default" yaml:"default"`
	Executable               string `json:"executable" yaml:"executable"`
	ImportedInstancesActions string `json:"imported_instances_actions" yaml:"imported_instances_actions"`
	Name                     string `json:"name" yaml:"name"`
	SunstoneName             string `json:"sunstone_name" yaml:"sunstone_name"`
	Type                     string `json:"type" yaml:"type"`
	KeepSnapshots            string `json:"keep_snapshots" yaml:"keep_snapshots"`
	ColdNICAttach            string `json:"cold_nicattach" yaml:"cold_nicattach"`
	DSLiveMigration          string `json:"dslive_migration" yaml:"dslive_migration"`
	LiveResize               string `json:"live_resize" yaml:"live_resize"`
}

// VNCPorts is the API payload based on the legacy xmlrpc backend.
type VNCPorts struct {
	Reserved string `json:"reserved" yaml:"reserved"`
	Start    int    `json:"start" yaml:"start"`
}

// VNetMADConf is the API payload based on the legacy xmlrpc backend.
type VNetMADConf struct {
	BridgeType string `json:"bridge_type" yaml:"bridge_type"`
	Name       string `json:"name" yaml:"name"`
}

// VxlanIDs is the API payload based on the legacy xmlrpc backend.
type VxlanIDs struct {
	Start int `json:"start" yaml:"start"`
}
