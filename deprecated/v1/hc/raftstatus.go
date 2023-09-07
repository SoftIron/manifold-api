package hc

// RaftStatus is the API payload based on the legacy xmlrpc backend.
type RaftStatus struct {
	ServerID    int `json:"server_id" yaml:"server_id"`
	State       int `json:"state" yaml:"state"`
	Term        int `json:"term" yaml:"term"`
	Votedfor    int `json:"votedfor" yaml:"votedfor"`
	Commit      int `json:"commit" yaml:"commit"`
	LogIndex    int `json:"log_index" yaml:"log_index"`
	LogTerm     int `json:"log_term" yaml:"log_term"`
	FedlogIndex int `json:"fedlog_index" yaml:"fedlog_index"`
}
