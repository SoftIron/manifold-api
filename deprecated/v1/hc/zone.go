package hc

// ZoneServer is the API payload based on the legacy xmlrpc backend.
type ZoneServer struct {
	Endpoint    string `json:"endpoint" yaml:"endpoint"`
	ID          int    `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	State       int    `json:"state" yaml:"state"`
	Term        int    `json:"term" yaml:"term"`
	Votedfor    int    `json:"votedfor" yaml:"votedfor"`
	Commit      int    `json:"commit" yaml:"commit"`
	LogIndex    int    `json:"log_index" yaml:"log_index"`
	FedlogIndex int    `json:"fedlog_index" yaml:"fedlog_index"`
}

// ZoneServerPool is the API payload based on the legacy xmlrpc backend.
type ZoneServerPool struct {
	Server []ZoneServer `json:"server" yaml:"server"`
}

// ZoneTemplate is the API payload based on the legacy xmlrpc backend.
type ZoneTemplate struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
}

// Zone is the API payload based on the legacy xmlrpc backend.
type Zone struct {
	ID         int            `json:"id" yaml:"id"`
	Name       string         `json:"name" yaml:"name"`
	State      int            `json:"state" yaml:"state"`
	Template   ZoneTemplate   `json:"template" yaml:"template"`
	ServerPool ZoneServerPool `json:"server_pool" yaml:"server_pool"`
}
