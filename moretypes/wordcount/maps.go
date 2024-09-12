package wordcount

import (
	"strings"
)

// https://go.dev/tour/moretypes/23

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	ret := map[string]int{}
	for _, word := range fields {
		ret[word]++ //
		// alternative equivalents:
		// ret[word] = ret[word] + 1
		// OR
		// el := ret[word]
		// ret[word] = el+1
	}
	return ret
}
