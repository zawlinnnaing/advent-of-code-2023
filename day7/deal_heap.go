package main

import (
	"strings"
)

type dealSorter struct {
	deals    []deal
	cardsMap map[string]int
}

func (h *dealSorter) Len() int {
	return len(h.deals)
}

func (h *dealSorter) Less(i, j int) bool {
	firstDeal := h.deals[i]
	secondDeal := h.deals[j]
	if firstDeal.rank == secondDeal.rank {
		firstDealCards := strings.Split(firstDeal.hand, "")
		secondDealCards := strings.Split(secondDeal.hand, "")
		for idx, firstDealCard := range firstDealCards {
			secondDealCard := secondDealCards[idx]
			if h.cardsMap[firstDealCard] == h.cardsMap[secondDealCard] {
				continue
			}
			return h.cardsMap[firstDealCard] < h.cardsMap[secondDealCard]
		}
	}

	return firstDeal.rank < secondDeal.rank
}

func (h *dealSorter) Swap(i, j int) {
	h.deals[i], h.deals[j] = h.deals[j], h.deals[i]
}

func (h *dealSorter) Push(x deal) {
	h.deals = append(h.deals, x)
}

func (h *dealSorter) Pop() deal {
	old := h.deals
	n := len(old)
	x := old[n-1]
	h.deals = old[0 : n-1]
	return x
}

func (h *dealSorter) Total() int {
	idx := h.Len()
	total := 0
	for h.Len() > 0 {
		output := h.Pop()
		total += idx * output.bid
		idx -= 1
	}
	return total
}
