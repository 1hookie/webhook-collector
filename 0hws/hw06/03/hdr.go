package main

import (
	"fmt"
	"strings"
)

func HeaderHas(accept, proto string) bool {
	for _, part := range strings.Split(accept, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		if strings.EqualFold(part, proto) {
			return true
		}
	}
	return false
}

func main() {
	str := "application/json, text/html"
	fmt.Printf("%q\n", strings.Split(str, ","))
	fmt.Printf("%q\n", strings.TrimSpace(str))
}
