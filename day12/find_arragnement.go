package main

import (
	"slices"
)

func count(cfg []string, records []int) int {
	if len(cfg) == 0 {
		if len(records) == 0 {
			return 1
		}
		return 0
	}
	if len(records) == 0 {
		if slices.Contains(cfg, DAMAGED) {
			return 0
		}
		return 1
	}
	result := 0
	if slices.Contains([]string{OPERATION, UNKNOWN}, cfg[0]) {
		result += count(cfg[1:], records)
	}
	if slices.Contains([]string{DAMAGED, UNKNOWN}, cfg[0]) {
		if records[0] <= len(cfg) && !slices.Contains(cfg[:records[0]], OPERATION) && (records[0] == len(cfg) || cfg[records[0]] != DAMAGED) {
			if records[0]+1 >= len(cfg) {
				result += count([]string{}, records[1:])
			} else {
				result += count(cfg[records[0]+1:], records[1:])
			}
		}
	}
	return result
}
