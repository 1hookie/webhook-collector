package dsn

import (
	"strings"
)

func ParseDSN(dsn string) map[string]string {
	res := map[string]string{}
	for _, tok := range strings.Fields(dsn) {
		parts := strings.SplitN(tok, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if key == "" {
			continue
		}
		res[key] = value
	}
	return res
}
