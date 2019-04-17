package main

import (
	"container/heap"
	"fmt"
)

// NodeId: id of the node in the graph
// Level: BFS level of this node
type BFSPair struct {
	NodeId   int64
	ParentId int64
	Level    int64
}

type BFSQueue []*BFSPair

func (pq BFSQueue) Len() int { return len(pq) }

// minimun heap
func (pq BFSQueue) Less(i, j int) bool {
	return pq[i].Level < pq[j].Level
}

func (pq BFSQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *BFSQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*BFSPair))
}

func (pq *BFSQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	p := old[n-1]
	*pq = old[0 : n-1]
	return p
}

func main() {
	pq := make(BFSQueue, 0)
	levels := []int64{3, 4, 5, 1, 2, 2}
	for _, lev := range levels {
		heap.Push(&pq, &BFSPair{
			NodeId:   lev,
			ParentId: lev,
			Level:    lev,
		})
	}
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*BFSPair)
		fmt.Println(node)
	}
}
