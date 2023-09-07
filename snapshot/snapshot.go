// Package snapshot documents the request and response payloads related to the snapshot daemon.
package snapshot

import (
	pb "github.com/softiron/hypercloud-api/gen/go/proto/snapper/v1"
)

// ListStackResponse is the response body for GET /snapper/stack.
type ListStackResponse pb.ListStackResponse

// ListStatusResponse is the response body for GET /snapper/status.
type ListStatusResponse struct {
	Snapshots []pb.ListStatusResponse `json:"snapshots"`
}

// ListArchiveResponse is the response body for GET /snapper/archive.
type ListArchiveResponse struct {
	Snapshots []pb.ListArchiveSnapshotsResponse `json:"snapshots"`
}

// ListManualResponse is the response body for GET /snapper/manual.
type ListManualResponse struct {
	Snapshots []pb.ListManualSnapshotsResponse `json:"snapshots"`
}

// ListRemoteResponse is the response body for GET /snapper/remote.
type ListRemoteResponse struct {
	Snapshots []pb.ListRemoteSnapshotsResponse `json:"snapshots"`
}
