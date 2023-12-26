package main

import (
	"slices"
)

func RunPart2(plane Plane) int {
	caches := []string{}
	hitIdx := 0
	for i := 0; i < 1e9; i++ {
		planeStr := plane.String()
		if slices.Contains(caches, planeStr) {
			// Hit index is the start of cycle inside cycle
			hitIdx = slices.IndexFunc(caches, func(s string) bool {
				return s == planeStr
			})
			break
		}
		caches = append(caches, planeStr)
		for _, direction := range directions {
			plane.Spin(direction)
		}
	}

	/**
	 * Illustrations
	 * Let [] to 1 billion cycle, let () to be small cycles inside [].
	 * [. . (. . . .)( . . . .). . . * 1e9]
	 * First we need to find small cycle length: `len(caches) - hitIdx`
	 * Then we disregard first `hitIndex` cycles, as we want to start beginning of small cycle: `1e9-hitIdx`
	 * [(. . . .)(. . . . .). . . * 1e9-hitIdx]
	 * Then, we use module to find last transformation in 1e9-hitIdx: 1e9-hitIdx % (len(caches) - hitIdx)
	 */
	cacheIdx := (1e9 - hitIdx) % (len(caches) - hitIdx)
	plane.parse(caches[hitIdx:][cacheIdx])
	return plane.TotalLoad()
}
