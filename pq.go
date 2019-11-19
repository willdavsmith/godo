package main

import (
	"container/heap"
)

type Item struct {
  Priority    int    `json:"priority"`
  Description string `json:"description"`
  Index       int    `json:"index"`
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	i := x.(*Item)
	i.Index = n
	*pq = append(*pq, i)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	i := old[n-1]
	old[n-1] = nil
	i.Index = -1
	*pq = old[0 : n-1]
	return i
}

func (pq *PriorityQueue) Update(i *Item, description string, priority int) {
	i.Description = description
	i.Priority = priority
	heap.Fix(pq, i.Index)
}

