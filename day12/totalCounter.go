package main

import (
	"sync"
)

type TotalCounter struct {
	sync.Mutex
	total int
}

func (tc *TotalCounter) Add(num int) {
	tc.Lock()
	tc.total += num
	tc.Unlock()
}

func (tc *TotalCounter) Get() int {
	return tc.total
}

func NewTotalCounter() *TotalCounter {
	return &TotalCounter{total: 0}
}
