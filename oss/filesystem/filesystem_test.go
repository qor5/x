package filesystem

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/qor5/x/v3/oss/tests"
)

func TestAll(t *testing.T) {
	fileSystem := New("/tmp")
	tests.TestAll(fileSystem, t)
}

func TestGetFullPath_ValidPaths(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "filesystem_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	fs := New(tempDir)

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "relative path within base",
			path:     "test.txt",
			expected: filepath.Join(tempDir, "test.txt"),
		},
		{
			name:     "relative path with subdirectory",
			path:     "subdir/test.txt",
			expected: filepath.Join(tempDir, "subdir/test.txt"),
		},
		{
			name:     "empty path (should resolve to base)",
			path:     "",
			expected: tempDir,
		},
		{
			name:     "dot path (should resolve to base)",
			path:     ".",
			expected: tempDir,
		},

		{
			name:     "OSS-style path with leading slash",
			path:     "/test.txt",
			expected: filepath.Join(tempDir, "test.txt"),
		},
		{
			name:     "OSS-style path with subdirectory",
			path:     "/subdir/test.txt",
			expected: filepath.Join(tempDir, "subdir/test.txt"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := fs.GetFullPath(tt.path)
			if err != nil {
				t.Errorf("GetFullPath(%q) returned error: %v", tt.path, err)
				return
			}

			// Clean both paths for comparison
			result = filepath.Clean(result)
			expected := filepath.Clean(tt.expected)

			if result != expected {
				t.Errorf("GetFullPath(%q) = %q, want %q", tt.path, result, expected)
			}
		})
	}
}

func TestGetFullPath_InvalidPaths(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "filesystem_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	fs := New(tempDir)

	tests := []struct {
		name string
		path string
	}{
		{
			name: "parent directory traversal with ..",
			path: "../outside.txt",
		},
		{
			name: "multiple parent directory traversal",
			path: "../../outside.txt",
		},
		{
			name: "complex path traversal",
			path: "subdir/../../outside.txt",
		},
		{
			name: "path trying to escape with leading slash and ..",
			path: "/../outside.txt",
		},
		{
			name: "path that would escape if it were truly absolute",
			path: "/../../../etc/passwd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := fs.GetFullPath(tt.path)
			if err == nil {
				t.Errorf("GetFullPath(%q) should have returned error, but got: %q", tt.path, result)
			}

			if err != nil && !strings.Contains(err.Error(), "access denied") {
				t.Errorf("GetFullPath(%q) error should contain 'access denied', but got: %v", tt.path, err)
			}
		})
	}
}

func TestGetFullPath_EdgeCases(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "filesystem_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	fs := New(tempDir)

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "path with ./",
			path:    "./test.txt",
			wantErr: false,
		},
		{
			name:    "path with redundant slashes",
			path:    "subdir//test.txt",
			wantErr: false,
		},
		{
			name:    "path with multiple leading slashes",
			path:    "///test.txt",
			wantErr: false,
		},
		{
			name:    "path attempting to escape with ..",
			path:    "subdir/../../../escape.txt",
			wantErr: true,
		},
		{
			name:    "symlink-like path that resolves within base",
			path:    "link/../inside.txt",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := fs.GetFullPath(tt.path)

			if tt.wantErr {
				if err == nil {
					t.Errorf("GetFullPath(%q) should have returned error, but got: %q", tt.path, result)
				}
			} else {
				if err != nil {
					t.Errorf("GetFullPath(%q) should not have returned error, but got: %v", tt.path, err)
				}

				// Verify the result is within the base directory
				if err == nil && !strings.HasPrefix(result, tempDir) {
					t.Errorf("GetFullPath(%q) result %q is not within base directory %q", tt.path, result, tempDir)
				}
			}
		})
	}
}
