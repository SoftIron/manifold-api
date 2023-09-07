package hc

// ACL is the API payload based on the legacy xmlrpc backend.
type ACL struct {
	ID       int    `json:"id" yaml:"id"`
	User     string `json:"user" yaml:"user"`
	Resource string `json:"resource" yaml:"resource"`
	Rights   string `json:"rights" yaml:"rights"`
	Zone     string `json:"zone" yaml:"zone"`
	String   string `json:"string" yaml:"string"`
}

// ACLPool is the API payload based on the legacy xmlrpc backend.
type ACLPool struct {
	ACL []ACL `json:"acl" yaml:"acl"`
}
