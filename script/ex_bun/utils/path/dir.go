package path

import (
	"path/filepath"
	"runtime"
)

func DirScript() string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	wd, _ := filepath.Abs(filepath.Dir(currentFilePath))
	return wd
}
