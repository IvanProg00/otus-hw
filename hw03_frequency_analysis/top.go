package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	words := strings.Fields(str)
	words = cleanNothing(words)

	wordsStat := map[string]int{}
	for _, w := range words {
		wordsStat[w]++
	}

	keys := []string{}
	for key := range wordsStat {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(a, b int) bool {
		if wordsStat[keys[a]] == wordsStat[keys[b]] {
			return keys[a] < keys[b]
		}
		return wordsStat[keys[a]] > wordsStat[keys[b]]
	})

	if len(keys) > 10 {
		keys = keys[:10]
	}

	return keys
}

func cleanNothing(slice []string) []string {
	res := []string{}
	for _, str := range slice {
		if str == "" {
			continue
		}
		res = append(res, str)
	}
	return res
}
