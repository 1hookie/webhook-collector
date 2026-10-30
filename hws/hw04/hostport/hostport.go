// Задача 1. «host:port» — надёжный разбор строки.
package hostport

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseHostPort(s string) (host string, port int, err error) {
	parts := strings.SplitN(s, ":", 2)
	//проверка на наличие ровно двух частией - хост и порт, иначе err
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid %q: expected host:port", s)
	}
	host, portStr := parts[0], parts[1]

	switch {
	case strings.Contains(portStr, ":"):
		return "", 0, fmt.Errorf("invalid %q: more than one ':'", s)
	case host == "":
		return "", 0, fmt.Errorf("invalid %q: empty host", s)
	case portStr == "":
		return "", 0, fmt.Errorf("invalid %q: empty port", s)
	}

	port, err = strconv.Atoi(portStr)
	if err != nil {
		return "", 0, fmt.Errorf("invalid port %q: %w", portStr, err)
	}
	if port < 1 || port > 65535 {
		return "", 0, fmt.Errorf("port %d out of range 1..65535", port)
	}
	return host, port, nil
}
