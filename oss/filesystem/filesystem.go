package filesystem

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/qor5/x/v3/filepathx"
	"github.com/qor5/x/v3/oss"
)

// FileSystem file system storage
type FileSystem struct {
	Base string
}

// New initialize FileSystem storage
func New(base string) *FileSystem {
	absbase, err := filepath.Abs(base)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize FileSystem storage: cannot get absolute path for base directory %q: %v", base, err))
	}
	return &FileSystem{Base: absbase}
}

// Get receive file with given path
func (fileSystem FileSystem) Get(ctx context.Context, path string) (*os.File, error) {
	fullPath, err := filepathx.Join(fileSystem.Base, path)
	if err != nil {
		return nil, err
	}
	return os.Open(fullPath)
}

// GetStream get file as stream
func (fileSystem FileSystem) GetStream(ctx context.Context, path string) (io.ReadCloser, error) {
	fullPath, err := filepathx.Join(fileSystem.Base, path)
	if err != nil {
		return nil, err
	}
	return os.Open(fullPath)
}

// Put store a reader into given path
func (fileSystem FileSystem) Put(ctx context.Context, path string, reader io.Reader) (*oss.Object, error) {
	fullpath, err := filepathx.Join(fileSystem.Base, path)
	if err != nil {
		return nil, err
	}
	if err = os.MkdirAll(filepath.Dir(fullpath), os.ModePerm); err != nil {
		return nil, err
	}

	if seeker, ok := reader.(io.ReadSeeker); ok {
		seeker.Seek(0, 0)
	}
	buf := bytes.NewBuffer([]byte{})
	if _, err = io.Copy(buf, reader); err != nil {
		return nil, err
	}

	dst, err := os.Create(fullpath)

	if err == nil {
		_, err = io.Copy(dst, buf)
	}

	return &oss.Object{Path: path, Name: filepath.Base(path), StorageInterface: fileSystem}, err
}

// Delete delete file
func (fileSystem FileSystem) Delete(ctx context.Context, path string) error {
	fullPath, err := filepathx.Join(fileSystem.Base, path)
	if err != nil {
		return err
	}
	return os.Remove(fullPath)
}

// List list all objects under current path
func (fileSystem FileSystem) List(ctx context.Context, path string) ([]*oss.Object, error) {
	var objects []*oss.Object
	fullpath, err := filepathx.Join(fileSystem.Base, path)
	if err != nil {
		return nil, err
	}

	filepath.Walk(fullpath, func(path string, info os.FileInfo, err error) error {
		if path == fullpath {
			return nil
		}

		if err == nil && !info.IsDir() {
			modTime := info.ModTime()
			objects = append(objects, &oss.Object{
				Path:             strings.TrimPrefix(path, fileSystem.Base),
				Name:             info.Name(),
				LastModified:     &modTime,
				StorageInterface: fileSystem,
			})
		}
		return nil
	})

	return objects, nil
}

// GetEndpoint get endpoint, FileSystem's endpoint is /
func (fileSystem FileSystem) GetEndpoint(ctx context.Context) string {
	return "/"
}

// GetURL get public accessible URL
func (fileSystem FileSystem) GetURL(ctx context.Context, path string) (url string, err error) {
	return path, nil
}
