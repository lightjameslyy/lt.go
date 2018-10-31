package leet_0382

import (
	"math/rand"
	"time"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	h    *ListNode
	size int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	s := Solution{h: head, size: 0}
	for head != nil {
		s.size++
		head = head.Next
	}
	return s
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(this.size)
	p := this.h
	for count > 0 {
		p = p.Next
		count--
	}
	return p.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
