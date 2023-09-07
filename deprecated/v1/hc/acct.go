package hc

// AcctHistory is the API payload based on the legacy xmlrpc backend.
type AcctHistory struct {
	OID       int    `json:"oid" yaml:"oid"`
	Seq       int    `json:"seq" yaml:"seq"`
	Hostname  string `json:"hostname" yaml:"hostname"`
	HID       int    `json:"hid" yaml:"hid"`
	CID       int    `json:"cid" yaml:"cid"`
	Stime     int    `json:"stime" yaml:"stime"`
	Etime     int    `json:"etime" yaml:"etime"`
	VMMAD     string `json:"vmmad" yaml:"vmmad"`
	TmMAD     string `json:"tm_mad" yaml:"tm_mad"`
	DSID      int    `json:"dsid" yaml:"dsid"`
	Pstime    int    `json:"pstime" yaml:"pstime"`
	Petime    int    `json:"petime" yaml:"petime"`
	Rstime    int    `json:"rstime" yaml:"rstime"`
	Retime    int    `json:"retime" yaml:"retime"`
	Estime    int    `json:"estime" yaml:"estime"`
	Eetime    int    `json:"eetime" yaml:"eetime"`
	Action    int    `json:"action" yaml:"action"`
	UID       int    `json:"uid" yaml:"uid"`
	GID       int    `json:"gid" yaml:"gid"`
	RequestID string `json:"request_id" yaml:"request_id"`
	VM        AcctVM `json:"vm" yaml:"vm"`
}

// AcctHistoryRecords is the API payload based on the legacy xmlrpc backend.
type AcctHistoryRecords struct {
	History []string `json:"history" yaml:"history"`
}

// AcctPermissions is the API payload based on the legacy xmlrpc backend.
type AcctPermissions struct {
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

// AcctSnapshot is the API payload based on the legacy xmlrpc backend.
type AcctSnapshot struct {
	Active   string `json:"active" yaml:"active"`
	Children string `json:"children" yaml:"children"`
	Date     int    `json:"date" yaml:"date"`
	ID       int    `json:"id" yaml:"id"`
	Name     string `json:"name" yaml:"name"`
	Parent   int    `json:"parent" yaml:"parent"`
	Size     int    `json:"size" yaml:"size"`
}

// AcctSnapshots is the API payload based on the legacy xmlrpc backend.
type AcctSnapshots struct {
	AllowOrphans string         `json:"allow_orphans" yaml:"allow_orphans"`
	CurrentBase  int            `json:"current_base" yaml:"current_base"`
	DiskID       int            `json:"disk_id" yaml:"disk_id"`
	NextSnapshot int            `json:"next_snapshot" yaml:"next_snapshot"`
	Snapshot     []AcctSnapshot `json:"snapshot" yaml:"snapshot"`
}

// AcctVM is the API payload based on the legacy xmlrpc backend.
type AcctVM struct {
	ID             int             `json:"id" yaml:"id"`
	UID            int             `json:"uid" yaml:"uid"`
	GID            int             `json:"gid" yaml:"gid"`
	Uname          string          `json:"uname" yaml:"uname"`
	Gname          string          `json:"gname" yaml:"gname"`
	Name           string          `json:"name" yaml:"name"`
	Permissions    AcctPermissions `json:"permissions" yaml:"permissions"`
	LastPoll       int             `json:"last_poll" yaml:"last_poll"`
	State          int             `json:"state" yaml:"state"`
	LcmState       int             `json:"lcm_state" yaml:"lcm_state"`
	PrevState      int             `json:"prev_state" yaml:"prev_state"`
	PrevLcmState   int             `json:"prev_lcm_state" yaml:"prev_lcm_state"`
	Resched        int             `json:"resched" yaml:"resched"`
	Stime          int             `json:"stime" yaml:"stime"`
	Etime          int             `json:"etime" yaml:"etime"`
	DeployID       string          `json:"deploy_id" yaml:"deploy_id"`
	Monitoring     string          `json:"monitoring" yaml:"monitoring"`
	Template       string          `json:"template" yaml:"template"`
	UserTemplate   string          `json:"user_template" yaml:"user_template"`
	HistoryRecords string          `json:"history_records" yaml:"history_records"`
	Snapshots      []AcctSnapshots `json:"snapshots" yaml:"snapshots"`
}
