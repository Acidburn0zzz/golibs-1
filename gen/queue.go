// Copyright (c) 2018 Sung Pae <self@sungpae.com>
// Distributed under the MIT license.
// http://www.opensource.org/licenses/mit-license.php

package gen

import (
	"math/bits"

	"github.com/cheekybits/genny/generic"
)

type Type generic.Type

// A TypeQueue is an auto-growing queue backed by a ring buffer.
type TypeQueue struct {
	a          []Type
	head, tail int
}

// DefaultQueueLen is the default size of a TypeQueue that is created with a
// non-positive size.
const DefaultQueueLen = 8

// NewTypeQueue returns a new queue with the given size rounded to the next
// power of two, or DefaultQueueLen if size <= 0.
func NewTypeQueue(size int) *TypeQueue {
	if size <= 0 {
		size = DefaultQueueLen
	}
	return &TypeQueue{
		a:    make([]Type, 1<<uint(bits.Len(uint(size-1)))),
		head: -1,
		tail: -1,
	}
}

// Len returns the current number of queued elements.
func (q *TypeQueue) Len() int {
	switch {
	case q.head == -1:
		return 0
	case q.head < q.tail:
		return q.tail - q.head
	default:
		return len(q.a) - q.head + q.tail
	}
}

// Enqueue adds a new element into the queue. If adding this element would
// overflow the queue, the current queue is moved to a new TypeQueue twice the
// size of the original before adding the element.
func (q *TypeQueue) Enqueue(x Type) {
	if q.tail == -1 {
		q.head = 0
		q.tail = 0
	} else if q.head == q.tail {
		q.Grow(1)
	}

	q.a[q.tail] = x

	q.tail++
	if q.tail >= len(q.a) {
		q.tail -= len(q.a)
	}
}

// Dequeue removes and returns the next element from the queue. Calling
// Dequeue on an empty queue results in a panic.
func (q *TypeQueue) Dequeue() Type {
	if q.head == -1 {
		panic("TypeQueue underflow")
	}

	x := q.a[q.head]

	q.head++
	if q.head >= len(q.a) {
		q.head -= len(q.a)
	}

	if q.head == q.tail {
		q.Reset()
	}

	return x
}

// Peek returns the next element from the queue without removing it. Peeking
// an empty queue results in a panic.
func (q *TypeQueue) Peek() Type {
	if q.head == -1 {
		panic("cannot peek empty TypeQueue")
	}

	return q.a[q.head]
}

// Reset the queue so that its length is zero. Note that the internal slice is
// NOT cleared.
func (q *TypeQueue) Reset() {
	q.head = -1
	q.tail = -1
}

// Grow internal slice to accommodate at least n more items.
func (q *TypeQueue) Grow(n int) {
	n -= cap(q.a) - len(q.a)
	if n <= 0 {
		return
	}

	r := TypeQueue{
		a:    make([]Type, 1<<uint(bits.Len(uint(cap(q.a)+n-1)))),
		head: -1,
		tail: -1,
	}

	for q.Len() > 0 {
		r.Enqueue(q.Dequeue())
	}

	*q = r
}
