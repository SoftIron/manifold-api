package cloud

// Host is a HyperCloud physical host.
type Host struct {
	Name    string   `json:"name"`
	IPs     []string `json:"ips"`
	Static  bool     `json:"static"`
	Compute bool     `json:"compute"`
	Storage bool     `json:"storage"`
}
