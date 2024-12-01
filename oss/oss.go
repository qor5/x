package oss

import (
	"context"
	"io"
	"os"
	"time"
)

// StorageInterface define common API to operate storage
type StorageInterface interface {
	Get(ctx context.Context, path string) (*os.File, error)
	GetStream(ctx context.Context, path string) (io.ReadCloser, error)
	Put(ctx context.Context, path string, reader io.Reader) (*Object, error)
	Delete(ctx context.Context, path string) error
	List(ctx context.Context, path string) ([]*Object, error)
	GetURL(ctx context.Context, path string) (string, error)
	GetEndpoint(ctx context.Context) string
}

// Object content object
type Object struct {
	Path             string
	Name             string
	LastModified     *time.Time
	StorageInterface StorageInterface
}

// Get retrieve object's content
func (object Object) Get(ctx context.Context) (*os.File, error) {
	return object.StorageInterface.Get(ctx, object.Path)
}
