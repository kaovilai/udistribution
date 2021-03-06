// Package tmpdir is a TESTING-ONLY utility.
//
// Some tests directly or indirectly exercising the directory/explicitfilepath
// subpackage expect the path returned by os.MkdirTemp to be canonical in the
// directory/explicitfilepath sense (absolute, no symlinks, cleaned up).
//
// os.MkdirTemp uses $TMPDIR by default, and on macOS, $TMPDIR is by
// default set to /var/folders/…, with /var a symlink to /private/var ,
// which does not match our expectations.  So, tests which want to use
// os.MkdirTemp that way, can
// import _ "github.com/containers/image/internal/testing/explicitfilepath-tmpdir"
// to ensure that $TMPDIR is canonical and usable as a base for testing
// path canonicalization in its subdirectories.
//
// NEVER use this in non-testing subpackages!
package tmpdir

import (
	"os"
	"path/filepath"
)

func init() {
	tmpDir := os.TempDir()
	explicitTmpDir, err := filepath.EvalSymlinks(tmpDir)
	if err == nil {
		os.Setenv("TMPDIR", explicitTmpDir)
	}
}
