package cloud

// Hook is the API payload based on the legacy xmlrpc backend.
type Hook struct {
	ID           int          `json:"id" yaml:"id"`
	Name         string       `json:"name" yaml:"name"`
	Type         string       `json:"type" yaml:"type"`
	Template     HookTemplate `json:"template" yaml:"template"`
	TemplateText string       `json:"template_text" yaml:"template_text"`
	HookLog      HookLog      `json:"hook_log" yaml:"hook_log"`
}

// HookLog is the API payload based on the legacy xmlrpc backend.
type HookLog struct {
	HookExecutionRecord []HookExecutionRecord `json:"execution_record" yaml:"execution_record"`
}

// HookExecutionResult is the API payload based on the legacy xmlrpc backend.
type HookExecutionResult struct {
	Command string `json:"command" yaml:"command"`
	Stdout  string `json:"stdout" yaml:"stdout"`
	Stderr  string `json:"stderr" yaml:"stderr"`
	Code    string `json:"code" yaml:"code"`
}

// HookExecutionRecord is the API payload based on the legacy xmlrpc backend.
type HookExecutionRecord struct {
	Values          map[string]string   `json:"values" yaml:"values"`
	HookID          int                 `json:"hook_id" yaml:"hook_id"`
	ExecutionID     int                 `json:"execution_id" yaml:"execution_id"`
	Timestamp       int                 `json:"timestamp" yaml:"timestamp"`
	Arguments       string              `json:"arguments" yaml:"arguments"`
	ExecutionResult HookExecutionResult `json:"execution_result" yaml:"execution_result"`
	RemoteHost      string              `json:"remote_host" yaml:"remote_host"`
	Retry           string              `json:"retry" yaml:"retry"`
}

// HookTemplate is the API payload based on the legacy xmlrpc backend.
type HookTemplate struct {
	Values         map[string]string `json:"values" yaml:"values"`
	Arguments      string            `json:"arguments" yaml:"arguments"`
	ArgumentsStdin string            `json:"arguments_stdin" yaml:"arguments_stdin"`
	Call           string            `json:"call" yaml:"call"`
	Command        string            `json:"command" yaml:"command"`
	Remote         string            `json:"remote" yaml:"remote"`
	Resource       string            `json:"resource" yaml:"resource"`
	State          string            `json:"state" yaml:"state"`
	LcmState       string            `json:"lcm_state" yaml:"lcm_state"`
}
