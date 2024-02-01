package cloud

// Zone is the API payload based on the legacy xmlrpc backend.
type Zone struct {
	ID           int          `json:"id" yaml:"id"`
	Name         string       `json:"name" yaml:"name"`
	State        int          `json:"state" yaml:"state"`
	Template     ZoneTemplate `json:"template" yaml:"template"`
	TemplateText string       `json:"template_text" yaml:"template_text"`
	Servers      []Server     `json:"servers" yaml:"servers"`
}

// ZoneTemplate is the API payload based on the legacy xmlrpc backend.
type ZoneTemplate struct {
	Values   map[string]string `json:"values" yaml:"values"`
	Endpoint string            `json:"endpoint" yaml:"endpoint"`
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
