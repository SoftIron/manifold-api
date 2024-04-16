// Package upload provides bindings for the file upload service.
package upload

const (
	// PathPrefix is the path prefix for the upload service.
	PathPrefix = "upload"

	uploadOffset         = "Upload-Offset"
	uploadLength         = "Upload-Length"
	uploadDeferLength    = "Upload-Defer-Length"
	uploadMetadata       = "Upload-Metadata"
	uploadChecksum       = "Upload-Checksum"
	location             = "Location"
	tusVersion           = "Tus-Version"
	tusResumable         = "Tus-Resumable"
	tusExtension         = "Tus-Extension"
	tusMaxSize           = "Tus-Max-Size"
	tusChecksumAlgorithm = "Tus-Checksum-Algorithm"
)
