package substr

import (
	"strings"
)

type Stats struct {
	Count    int
	FirstIdx int
	LastIdx  int
}

func SubstringStats(s, sub string) Stats {
	count := strings.Count(s, sub)
	firstIdx := strings.Index(s, sub)
	lastIdx := strings.LastIndex(s, sub)
	return Stats{
		Count:    count,
		FirstIdx: firstIdx,
		LastIdx:  lastIdx}
}
