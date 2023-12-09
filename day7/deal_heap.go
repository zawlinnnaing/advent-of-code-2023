package main

import "strings"

type dealHeap []deal

func (h *dealHeap) Len() int {
	return len(*h)
}

func (h *dealHeap) Less(i, j int) bool {
	firstDeal := (*h)[i]
	secondDeal := (*h)[j]
	if firstDeal.rank == secondDeal.rank {
		firstDealCards := strings.Split(firstDeal.hand, "")
		secondDealCards := strings.Split(secondDeal.hand, "")
		for idx, firstDealCard := range firstDealCards {
			secondDealCard := secondDealCards[idx]
			if cardsMap[firstDealCard] == cardsMap[secondDealCard] {
				continue
			}
			return cardsMap[firstDealCard] < cardsMap[secondDealCard]
		}
	}

	return firstDeal.rank < secondDeal.rank
}

func (h *dealHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *dealHeap) Push(x deal) {
	*h = append(*h, x)
}

func (h *dealHeap) Pop() deal {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
