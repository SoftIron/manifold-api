package upload

// CreateResponse is the response from the Create method.
type CreateResponse struct {
	Version string
	ID      string
}

// UploadResponse is the response from the Upload method.
type UploadResponse struct { // nolint:revive
	Version   string
	ID        string
	Offset    uint64
	Completed bool
}

// ResumeResponse is the response from the Resume method.
type ResumeResponse struct {
	Version   string
	Offset    uint64
	Completed bool
}

// StatusResponse is the response from the Status method.
type StatusResponse struct {
	Version string
	Offset  uint64
}

// InfoResponse is the response from the Info method.
type InfoResponse struct {
	Version     string
	Supported   []string
	MaxFileSize int
	Extensions  []string
}
