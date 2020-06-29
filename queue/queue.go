// Copyright [2020] [thinkgos]
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package queue

// Interface is a type of Queue, which is FIFO(first-in-first-out).
type Interface interface {
	// Len returns the number of elements in the collection.
	Len() int
	// IsEmpty returns true if this container contains no elements.
	IsEmpty() bool
	// Clear initializes or clears all of the elements from this container.
	Clear()
	// Add inserts an element into the tail of this Queue.
	Add(interface{})
	// Peek retrieves, but does not remove, the head of this Queue, or return nil if this Queue is empty.
	Peek() interface{}
	// Poll retrieves and removes the head of the this Queue, or return nil if this Queue is empty.
	Poll() interface{}
}

// element is an element of the Queue
type element struct {
	next  *element
	value interface{}
}

// Queue represents a singly linked list.
type Queue struct {
	head   *element
	tail   *element
	length int
}

var _ Interface = (*Queue)(nil)

// New creates a Queue. which implement queue.Interface
func New() *Queue { return new(Queue) }

// Len returns the length of this priority queue.
func (sf *Queue) Len() int { return sf.length }

// IsEmpty returns true if this Queue contains no elements.
func (sf *Queue) IsEmpty() bool { return sf.Len() == 0 }

// Clear initializes or clears queue.
func (sf *Queue) Clear() { sf.head, sf.tail, sf.length = nil, nil, 0 }

// Add items to the queue
func (sf *Queue) Add(v interface{}) {
	e := &element{value: v}
	if sf.tail == nil {
		sf.head, sf.tail = e, e
	} else {
		sf.tail.next = e
		sf.tail = e
	}
	sf.length++
}

// Peek retrieves, but does not remove, the head of this Queue, or return nil if this Queue is empty.
func (sf *Queue) Peek() interface{} {
	if sf.head != nil {
		return sf.head.value
	}
	return nil
}

// Poll retrieves and removes the head of the this Queue, or return nil if this Queue is empty.
func (sf *Queue) Poll() interface{} {
	var val interface{}

	if sf.head != nil {
		val = sf.head.value
		sf.head = sf.head.next
		if sf.head == nil {
			sf.tail = nil
		}
		sf.length--
	}
	return val
}
