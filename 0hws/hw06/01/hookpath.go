package hookpath

import (
	"strings"
)

func ValidHookPath(path string) bool {
	if !strings.HasPrefix(path, "/wh/") {
		return false
	}
	id := strings.TrimPrefix(path, "/wh/")
	if id != "" && strings.ContainsAny(id, "0123456789") {
		return true
	}
	return false

}
