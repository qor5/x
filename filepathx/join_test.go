package filepathx

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestJoin_ValidPaths(t *testing.T) {
	tests := []struct {
		name     string
		base     string
		path     string
		expected string
	}{
		{
			name:     "simple relative path",
			base:     "/home/user",
			path:     "file.txt",
			expected: "/home/user/file.txt",
		},
		{
			name:     "nested relative path",
			base:     "/home/user",
			path:     "dir/subdir/file.txt",
			expected: "/home/user/dir/subdir/file.txt",
		},
		{
			name:     "empty path",
			base:     "/home/user",
			path:     "",
			expected: "/home/user",
		},
		{
			name:     "dot path",
			base:     "/home/user",
			path:     ".",
			expected: "/home/user",
		},
		{
			name:     "OSS-style path with leading slash",
			base:     "/home/user",
			path:     "/file.txt",
			expected: "/home/user/file.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Join(tt.base, tt.path)
			if err != nil {
				t.Errorf("Join(%q, %q) returned unexpected error: %v", tt.base, tt.path, err)
				return
			}

			// Clean paths for comparison to handle OS differences
			result = filepath.Clean(result)
			expected := filepath.Clean(tt.expected)

			if result != expected {
				t.Errorf("Join(%q, %q) = %q, want %q", tt.base, tt.path, result, expected)
			}
		})
	}
}

func TestJoin_PathTraversalAttacks(t *testing.T) {
	tests := []struct {
		name string
		base string
		path string
	}{
		{
			name: "simple parent directory traversal",
			base: "/home/user",
			path: "../secret.txt",
		},
		{
			name: "multiple parent directory traversal",
			base: "/home/user",
			path: "../../etc/passwd",
		},
		{
			name: "complex traversal with subdirectories",
			base: "/home/user",
			path: "subdir/../../../etc/passwd",
		},
		{
			name: "traversal with current directory",
			base: "/home/user",
			path: "./../../secret.txt",
		},
		{
			name: "mixed traversal with absolute-like path",
			base: "/home/user",
			path: "subdir/../../etc/passwd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Join(tt.base, tt.path)
			if err == nil {
				t.Errorf("Join(%q, %q) should have returned error for path traversal attack, but got: %q", tt.base, tt.path, result)
				return
			}
			if !strings.Contains(err.Error(), "illegal file path") {
				t.Errorf("Join(%q, %q) should return 'illegal file path' error, got: %v", tt.base, tt.path, err)
			}
		})
	}
}

func TestJoin_CrossDriveAttacks(t *testing.T) {
	// These tests are primarily for Windows systems but should be safe on all platforms
	if runtime.GOOS != "windows" {
		t.Skip("Cross-drive attacks are Windows-specific")
	}

	tests := []struct {
		name string
		base string
		path string
	}{
		{
			name: "cross drive attack C to D",
			base: "C:\\Users\\user",
			path: "D:\\secrets\\file.txt",
		},
		{
			name: "cross drive attack with UNC path",
			base: "C:\\Users\\user",
			path: "\\\\server\\share\\file.txt",
		},
		{
			name: "cross drive via path traversal",
			base: "C:\\Users\\user",
			path: "..\\..\\..\\D:\\secrets\\file.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Join(tt.base, tt.path)
			if err == nil {
				t.Errorf("Join(%q, %q) should have returned error for cross-drive attack, but got: %q", tt.base, tt.path, result)
			}
			if !strings.Contains(err.Error(), "illegal file path") {
				t.Errorf("Join(%q, %q) should return 'illegal file path' error, got: %v", tt.base, tt.path, err)
			}
		})
	}
}

func TestJoin_EdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		base      string
		path      string
		shouldErr bool
		reason    string
	}{
		{
			name:      "symlink-like traversal",
			base:      "/home/user",
			path:      "link/../../../etc/passwd",
			shouldErr: true,
			reason:    "should block symlink-like traversal",
		},
		{
			name:      "mixed separators (Unix treats backslash as filename)",
			base:      "/home/user",
			path:      "dir\\..\\..\\secret.txt",
			shouldErr: false,
			reason:    "backslashes are valid filename chars on Unix",
		},
		{
			name:      "double dot in filename",
			base:      "/home/user",
			path:      "file..txt",
			shouldErr: false,
			reason:    "double dot in filename should be allowed",
		},
		{
			name:      "trailing dot dot (resolves to base)",
			base:      "/home/user",
			path:      "validdir/..",
			shouldErr: false,
			reason:    "validdir/.. resolves back to base directory, which is allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Join(tt.base, tt.path)

			if tt.shouldErr && err == nil {
				t.Errorf("Join(%q, %q) should have returned error: %s, but got: %q", tt.base, tt.path, tt.reason, result)
			}

			if !tt.shouldErr && err != nil {
				t.Errorf("Join(%q, %q) should not have returned error: %s, but got: %v", tt.base, tt.path, tt.reason, err)
			}

			if tt.shouldErr && err != nil && !strings.Contains(err.Error(), "illegal file path") {
				t.Errorf("Join(%q, %q) should return 'illegal file path' error, got: %v", tt.base, tt.path, err)
			}
		})
	}
}

func TestJoin_RelativePathErrors(t *testing.T) {
	// Test cases that specifically trigger filepath.Rel errors

	tests := []struct {
		name   string
		base   string
		path   string
		reason string
	}{
		{
			name:   "base and path on different drives (Windows)",
			base:   "C:\\Users\\user",
			path:   "..\\..\\..\\D:\\secrets\\file.txt",
			reason: "should trigger filepath.Rel error on cross-drive scenarios",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Skip non-Windows tests that are Windows-specific
			if runtime.GOOS != "windows" && strings.Contains(tt.name, "Windows") {
				t.Skip("Skipping Windows-specific test on non-Windows platform")
			}

			result, err := Join(tt.base, tt.path)
			if err == nil {
				t.Errorf("Join(%q, %q) should have returned error: %s, but got: %q", tt.base, tt.path, tt.reason, result)
			}

			// The error could be either a Rel error or illegal file path error
			if !strings.Contains(err.Error(), "illegal file path") && !strings.Contains(err.Error(), "failed to get relative path") {
				t.Errorf("Join(%q, %q) should return appropriate error message, got: %v", tt.base, tt.path, err)
			}
		})
	}
}

// TestJoin_SecurityValidation tests the specific security checks
// TestJoin_IsAbsEdgeCases tests scenarios where filepath.IsAbs(relPath) check is important
func TestJoin_IsAbsEdgeCases(t *testing.T) {
	// On Windows, test cross-drive scenarios where Rel might return absolute paths
	if runtime.GOOS == "windows" {
		tests := []struct {
			name string
			base string
			path string
		}{
			{
				name: "cross drive scenario",
				base: "C:\\temp",
				path: "..\\..\\D:\\secrets\\file.txt",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				fullpath := filepath.Join(tt.base, tt.path)
				relPath, err := filepath.Rel(tt.base, fullpath)

				// Debug information to understand what happens
				t.Logf("Base: %s", tt.base)
				t.Logf("Input path: %s", tt.path)
				t.Logf("Full path after Join: %s", fullpath)
				t.Logf("Relative path: %s", relPath)
				t.Logf("Rel error: %v", err)
				if err == nil {
					t.Logf("IsAbs(relPath): %v", filepath.IsAbs(relPath))
					t.Logf("HasPrefix(relPath, '..'): %v", strings.HasPrefix(relPath, ".."))
				}

				// Test actual function
				result, joinErr := Join(tt.base, tt.path)
				if joinErr == nil {
					t.Errorf("Join(%q, %q) should have returned error but got: %q", tt.base, tt.path, result)
				}
			})
		}
	} else {
		t.Skip("Cross-drive tests are Windows-specific")
	}
}

func TestJoin_SecurityValidation(t *testing.T) {
	base := "/home/user/storage"

	// Test the three security conditions in Join
	testCases := []struct {
		name      string
		path      string
		checkType string // "rel_error", "is_abs", "has_prefix_dotdot"
	}{
		{
			name:      "parent directory traversal triggers HasPrefix check",
			path:      "../../../etc/passwd",
			checkType: "has_prefix_dotdot",
		},
		{
			name:      "simple parent directory",
			path:      "../file.txt",
			checkType: "has_prefix_dotdot",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fullpath := filepath.Join(base, tc.path)
			relPath, err := filepath.Rel(base, fullpath)

			// Debug information
			t.Logf("Input path: %s", tc.path)
			t.Logf("Full path: %s", fullpath)
			t.Logf("Relative path: %s", relPath)
			t.Logf("Rel error: %v", err)
			if err == nil {
				t.Logf("IsAbs(relPath): %v", filepath.IsAbs(relPath))
				t.Logf("HasPrefix(relPath, '..'): %v", strings.HasPrefix(relPath, ".."))
			}

			// Test actual function
			result, joinErr := Join(base, tc.path)
			if joinErr == nil {
				t.Errorf("Join(%q, %q) should have returned error but got: %q", base, tc.path, result)
			}
		})
	}
}

// TestJoin_EmptyBase tests behavior when base directory is empty
func TestJoin_EmptyBase(t *testing.T) {
	tests := []struct {
		name      string
		base      string
		path      string
		shouldErr bool
		reason    string
	}{
		{
			name:      "empty base with relative path",
			base:      "",
			path:      "file.txt",
			shouldErr: false,
			reason:    "empty base with relative path should work",
		},
		{
			name:      "empty base with absolute path",
			base:      "",
			path:      "/etc/passwd",
			shouldErr: true,
			reason:    "empty base with absolute path should be blocked",
		},
		{
			name:      "empty base with parent traversal",
			base:      "",
			path:      "../secret.txt",
			shouldErr: true,
			reason:    "empty base with parent traversal should be blocked",
		},
		{
			name:      "empty base with empty path",
			base:      "",
			path:      "",
			shouldErr: false,
			reason:    "both empty should result in current directory",
		},
		{
			name:      "empty base with dot path",
			base:      "",
			path:      ".",
			shouldErr: false,
			reason:    "empty base with dot should work",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Join(tt.base, tt.path)

			if tt.shouldErr && err == nil {
				t.Errorf("Join(%q, %q) should have returned error: %s, but got: %q", tt.base, tt.path, tt.reason, result)
				return
			}

			if !tt.shouldErr && err != nil {
				t.Errorf("Join(%q, %q) should not have returned error: %s, but got: %v", tt.base, tt.path, tt.reason, err)
				return
			}

			// Debug output for successful cases
			if !tt.shouldErr && err == nil {
				t.Logf("Join(%q, %q) = %q", tt.base, tt.path, result)
			}
		})
	}
}
