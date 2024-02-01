package cloud

// Zone is the API payload based on the legacy xmlrpc backend.
type Zone struct {
	ID       int      `json:"id" yaml:"id"`
	Name     string   `json:"name" yaml:"name"`
	State    int      `json:"state" yaml:"state"`
	Template Template `json:"template" yaml:"template"`
	Servers  []Server `json:"servers" yaml:"servers"`
}

// ZoneTemplate is the API payload based on the legacy xmlrpc backend.
type ZoneTemplate struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
}

// Server is the API payload based on the legacy xmlrpc backend.
type Server struct {
	Endpoint    string `json:"endpoint" yaml:"endpoint"`
	ID          int    `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	State       int    `json:"state" yaml:"state"`
	Term        int    `json:"term" yaml:"term"`
	VotedFor    int    `json:"voted_for" yaml:"voted_for"`
	Commit      int    `json:"commit" yaml:"commit"`
	LogIndex    int    `json:"log_index" yaml:"log_index"`
	FedLogIndex int    `json:"fed_log_index" yaml:"fed_log_index"`
}

// ParseTemplate returns a structured subset of the nested key x value pair map.
func (z *Zone) ParseTemplate() (*ZoneTemplate, error) {
	var t ZoneTemplate

	for key, value := range z.Template {
		if v, ok := value.(string); ok {
			if key == "ENDPOINT" {
				t.Endpoint = v
			}
		}
	}

	return &t, nil
}
