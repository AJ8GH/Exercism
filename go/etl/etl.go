package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}
	for k, v := range in {
		for _, w := range v {
			out[strings.ToLower(w)] = k
		}
	}
	return out
}
