package main

import (
	"fmt"
	"math"
	"strings"
)

var cacheMap = CacheMap{cache: make(map[string]map[string]int)}

func count(cfg string, records *[]int) int {
	cfgStr := cfg
	recordStr := fmt.Sprint(records)
	if cfg == "" {
		if len(*records) == 0 {
			return 1
		}
		return 0
	}
	if len(*records) == 0 {
		if strings.Contains(cfg, DAMAGED) {
			return 0
		}
		return 1
	}
	if cacheMap.Get(cfgStr, recordStr) != 0 {
		return cacheMap.Get(cfgStr, recordStr)
	}
	result := 0

	if strings.Contains(fmt.Sprintf("%s%s", OPERATION, UNKNOWN), cfg[0:1]) {
		result += count(cfg[1:], records)
	}
	if strings.Contains(fmt.Sprintf("%s%s", DAMAGED, UNKNOWN), cfg[0:1]) {
		if (*records)[0] <= len(cfg) && !strings.Contains(cfg[:(*records)[0]], OPERATION) && ((*records)[0] == len(cfg) || cfg[(*records)[0]:(*records)[0]+1] != DAMAGED) {
			nextRecords := (*records)[1:]
			result += count(cfg[int(math.Min(float64((*records)[0]+1), float64(len(cfg)))):], &nextRecords)
		}
	}
	cacheMap.Set(cfgStr, recordStr, result)
	return result
}
