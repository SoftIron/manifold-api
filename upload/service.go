package upload

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/internal/api"
)

// Service owns the /upload methods.
type Service struct {
	ChunkSize int // ChunkSize is the size in bytes (and memory) for each file segment to upload.

	*client.Client
	root string
}

// NewService returns a new Service for upload operations.
func NewService(c *client.Client, path string) *Service {
	return &Service{
		ChunkSize: 1024 * 1024 * 32, // in bytes
		Client:    c,
		root:      path,
	}
}

// path returns a URL path with the correct prefix appended.
func (s Service) path(dirs ...interface{}) string {
	return api.Path(s.root, PathPrefix, dirs...)
}

// Create creates a new upload target. It will return server generated file ID
// that must be used to `Resume` the upload. The advantage of using `Create`
// over `Upload` is the client can obtain the file ID before sending any data.
// This means all failed uploads can be restarted continuing from the last
// successful write.
func (s Service) Create(ctx context.Context, path string) (*CreateResponse, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	header := newClientHeader()
	header[uploadLength] = strconv.FormatInt(stat.Size(), 10)
	header[uploadMetadata] = "filename " + base64.StdEncoding.EncodeToString([]byte(stat.Name()))

	resp, err := s.Client.Request(ctx, client.AccessTokenAuth, http.MethodPost, s.path(), header, http.NoBody)
	if err != nil {
		return nil, err
	}

	var create CreateResponse

	for key, values := range resp.Header {
		for _, value := range values {
			switch key {
			case tusResumable:
				create.Version = value
			case location:
				create.ID = filepath.Base(value)
			}
		}
	}

	return &create, nil
}

// Upload uploads the given file to the server. The upload is sent in chunks of
// `s.ChunkSize` bytes. If the upload fails the client can `Resume` the upload
// using the returned file `ID`.
func (s Service) Upload(ctx context.Context, path string) (*UploadResponse, error) {
	create, err := s.Create(ctx, path)
	if err != nil {
		return nil, err
	}

	upload := UploadResponse{
		Version: create.Version,
		ID:      create.ID,
	}

	fin, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fin.Close()

	status, err := s.upload(ctx, create.ID, 0, fin)
	if err != nil {
		return &upload, err // need to return the upload response so the caller can get the ID
	}

	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	upload.Offset = status.Offset
	upload.Completed = int64(status.Offset) == stat.Size()

	return &upload, nil
}

// Resume resumes partial `Upload` of a file. The `id` is the server assigned
// file id returned by `Upload`, or `Create`.
func (s Service) Resume(ctx context.Context, id, path string) (*ResumeResponse, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	status, err := s.Status(ctx, id)
	if err != nil {
		return nil, err
	}

	fin, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fin.Close()

	status, err = s.upload(ctx, id, int64(status.Offset), fin)
	if err != nil {
		return nil, err
	}

	resume := ResumeResponse{
		Version:   status.Version,
		Offset:    status.Offset,
		Completed: int64(status.Offset) == stat.Size(),
	}

	return &resume, nil
}

// Status returns the status of the upload. Specifically the offset that should
// be used to resume the upload.
func (s Service) Status(ctx context.Context, path string) (*StatusResponse, error) {
	header := newClientHeader()
	resp, err := s.Client.Request(ctx, client.AccessTokenAuth, http.MethodHead, s.path(path), header, http.NoBody)
	if err != nil {
		return nil, err
	}

	var status StatusResponse

	for key, values := range resp.Header {
		for _, value := range values {
			switch key {
			case tusResumable:
				status.Version = value
			case uploadOffset:
				offset, err := strconv.ParseUint(value, 10, 64)
				if err != nil {
					return nil, err
				}
				status.Offset = offset
			}
		}
	}

	return &status, nil
}

// Info returns information about the server's capabilities.
func (s Service) Info(ctx context.Context) (*InfoResponse, error) {
	resp, err := s.Client.Request(ctx, client.AccessTokenAuth, http.MethodOptions, s.path(), nil, http.NoBody)
	if err != nil {
		return nil, err
	}

	var info InfoResponse

	for key, values := range resp.Header {
		for _, value := range values {
			switch key {
			case tusResumable:
				info.Version = value
			case tusVersion:
				info.Supported = strings.Split(value, ",")
			case tusMaxSize:
				n, err := strconv.Atoi(value)
				if err != nil {
					return nil, err
				}
				info.MaxFileSize = n
			case tusExtension:
				info.Extensions = strings.Split(value, ",")
			}
		}
	}

	return &info, nil
}

func (s Service) upload(ctx context.Context, id string, offset int64, r io.ReadSeeker) (*StatusResponse, error) {
	header := newClientHeader()
	header["Content-Type"] = "application/offset+octet-stream"

	for {
		if _, err := r.Seek(offset, 0); err != nil {
			return nil, fmt.Errorf("failed to seek to offset: %w", err)
		}

		b := make([]byte, s.ChunkSize)
		n, err := r.Read(b)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("failed to read from file: %w", err)
		}
		if n == 0 {
			break
		}

		header[uploadOffset] = strconv.FormatInt(offset, 10)
		header["Content-Length"] = strconv.Itoa(n)

		resp, err := s.Client.Request(ctx, client.AccessTokenAuth, http.MethodPatch, s.path(id), header, bytes.NewReader(b[:n]))
		if err != nil {
			return nil, err
		}

		for key, values := range resp.Header {
			for _, value := range values {
				if key == uploadOffset {
					offset, err = strconv.ParseInt(value, 10, 64)
					if err != nil {
						return nil, fmt.Errorf("failed to parse offset: %w", err)
					}
					break
				}
			}
		}
	}

	return s.Status(ctx, id)
}

func newClientHeader() map[string]string {
	h := make(map[string]string)
	h[tusResumable] = "1.0.0" // TODO: should really be negotiated, but we know it's 1.0.0

	return h
}
