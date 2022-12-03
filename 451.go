package goLeet

import (
	"sort"
)

// passes the leetcode test beating 29.36% performance and 25.69% memory
func frequencySort(s string) string {
	freqMap := map[rune]int{}
	runeSlice := []rune{}

	// 1. build frequency map and convert string to slice in one pass
	for _, r := range s {
		if frequency, ok := freqMap[r]; ok {
			freqMap[r] = frequency + 1
		} else {
			freqMap[r] = 1
		}
		runeSlice = append(runeSlice, r)
	}
	// 2. sort the slice by frequency
	sort.Slice(runeSlice, func(i int, j int) bool {
		iFreq := freqMap[runeSlice[i]]
		jFreq := freqMap[runeSlice[j]]
		// we need a tie breaker when frequency is the same
		if iFreq == jFreq {
			return runeSlice[i] > runeSlice[j]
		}
		return iFreq>jFreq
	})

	// 4. convert the runes back to a string
	return string (runeSlice)
}
