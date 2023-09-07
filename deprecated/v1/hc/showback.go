package hc

// Showback is the API payload based on the legacy xmlrpc backend.
type Showback struct {
	VMID       int     `json:"vmid" yaml:"vmid"`
	Vmname     string  `json:"vmname" yaml:"vmname"`
	UID        int     `json:"uid" yaml:"uid"`
	GID        int     `json:"gid" yaml:"gid"`
	Uname      string  `json:"uname" yaml:"uname"`
	Gname      string  `json:"gname" yaml:"gname"`
	Year       int     `json:"year" yaml:"year"`
	Month      int     `json:"month" yaml:"month"`
	CPUCost    float32 `json:"cpucost" yaml:"cpucost"`
	MemoryCost float32 `json:"memory_cost" yaml:"memory_cost"`
	DiskCost   float32 `json:"disk_cost" yaml:"disk_cost"`
	TotalCost  float32 `json:"total_cost" yaml:"total_cost"`
	Hours      float32 `json:"hours" yaml:"hours"`
	Rhours     float32 `json:"rhours" yaml:"rhours"`
}

// ShowbackRecords is the API payload based on the legacy xmlrpc backend.
type ShowbackRecords struct {
	Showback []Showback `json:"showback" yaml:"showback"`
}
