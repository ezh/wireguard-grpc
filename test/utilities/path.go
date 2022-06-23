package utilities

import (
	"os"
	"runtime"
	"strings"
)

// Get test root directory
func TestGetRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	if !strings.HasSuffix(filename,
		strings.Join([]string{"test", "utilities", "path.go"}, string(os.PathSeparator))) {
		panic("Unable to locate test root directory")
	}
	return filename[:len(filename)-len("/utilities/path.go")]
}
