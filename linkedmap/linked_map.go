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

package linkedmap

import (
	"container/list"

	"github.com/thinkgos/container/comparator"
)

// Interface is a type of linked map, and LinkedMap implements this interface.
type Interface interface {
	// Cap returns the capacity of elements of list l.
	Cap() int
	// Len returns the number of elements in the collection.
	Len() int
	// IsEmpty returns true if this container contains no elements.
	IsEmpty() bool
	// Clear initializes or clears all of the elements from this container.
	Clear()
	// Push associates the specified value with the specified key in this map.
	// If the map previously contained a mapping for the key,
	// the old value is replaced by the specified value. and then move the item to the back of the list.
	// If over the cap, it will remove the back item then push new item to back
	// It returns the previous value associated with the specified key, or nil if there was no mapping for the key.
	// A nil return can also indicate that the map previously associated nil with the specified key.
	Push(k, v interface{}) interface{}
	// PushFront associates the specified value with the specified key in this map.
	// If the map previously contained a mapping for the key,
	// the old value is replaced by the specified value. and then move the item to the front of the list.
	// If over the cap, it will remove the back item then push new item to front
	// It returns the previous value associated with the specified key, or nil if there was no mapping for the key.
	// A nil return can also indicate that the map previously associated nil with the specified key.
	PushFront(k, v interface{}) interface{}
	// PushBack associates the specified value with the specified key in this map.
	// If the map previously contained a mapping for the key,
	// the old value is replaced by the specified value. and then move the item to the back of the list.
	// If over the cap, it will remove the back item then push new item to back
	PushBack(k, v interface{}) interface{}

	// Poll removes the first element from this map, which is the head of the list.
	// It returns the (key, value, true) if the map isn't empty, or (nil, nil, false) if the map is empty.
	Poll() (interface{}, interface{}, bool)
	// PollFront return the front element value and then remove from list
	PollFront() (k interface{}, v interface{}, exist bool)
	// PollBack removes the last element from this map, which is the tail of the list.
	// It returns the (key, value, true) if the map isn't empty, or (nil, nil, false) if the map is empty.
	PollBack() (interface{}, interface{}, bool)
	// Remove removes the mapping for a key from this map if it is present.
	// It returns the value to which this map previously associated the key, and true,
	// or nil and false if the map contained no mapping for the key.
	Remove(k interface{}) (interface{}, bool)
	// Get returns the value to which the specified key is mapped, or nil if this map contains no mapping for the key.
	Get(k interface{}, defaultValue ...interface{}) interface{}

	// ContainsKey returns true if this map contains a mapping for the specified key.
	ContainsKey(k interface{}) bool
	// ContainsValue returns true if this map maps one or more keys to the specified value.
	ContainsValue(v interface{}) bool

	// Iterator returns an iterator over the elements in this map in proper sequence.
	Iterator(cb func(k interface{}, v interface{}) bool)
	// ReverseIterator returns an iterator over the elements in this map in reverse sequence as Iterator.
	ReverseIterator(cb func(k interface{}, v interface{}) bool)
}

type store struct {
	key, value interface{}
}

// LinkedMap implements the Interface.
type LinkedMap struct {
	data map[interface{}]*list.Element
	l    *list.List
	cmp  comparator.Comparator
	cap  int
}

var _ Interface = (*LinkedMap)(nil)

// Option option for New
type Option func(lm *LinkedMap)

// WithCap with limit capacity
func WithCap(cap int) Option {
	return func(lm *LinkedMap) {
		lm.cap = cap
	}
}

// WithComparator with user's Comparator
func WithComparator(cmp comparator.Comparator) Option {
	return func(lm *LinkedMap) {
		lm.cmp = cmp
	}
}

// New creates a LinkedMap.
func New(opts ...Option) *LinkedMap {
	lm := &LinkedMap{
		data: make(map[interface{}]*list.Element),
		l:    list.New(),
	}
	for _, opt := range opts {
		opt(lm)
	}
	return lm
}

// Cap returns the capacity of elements of list l.
// The complexity is O(1).
func (sf *LinkedMap) Cap() int { return sf.cap }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (sf *LinkedMap) Len() int { return sf.l.Len() }

// IsEmpty returns the list l is empty or not
func (sf *LinkedMap) IsEmpty() bool { return sf.Len() == 0 }

// Clear initializes or clears list l.
func (sf *LinkedMap) Clear() {
	sf.data = make(map[interface{}]*list.Element)
	sf.l.Init()
}

// Push associates the specified value with the specified key in this map.
// If the map previously contained a mapping for the key,
// the old value is replaced by the specified value. and then move the item to the back of the list.
// If over the cap, it will remove the back item then push new item to back
// It returns the previous value associated with the specified key, or nil if there was no mapping for the key.
// A nil return can also indicate that the map previously associated nil with the specified key.
func (sf *LinkedMap) Push(k, v interface{}) interface{} {
	return sf.PushBack(k, v)
}

// PushFront associates the specified value with the specified key in this map.
// If the map previously contained a mapping for the key,
// the old value is replaced by the specified value. and then move the item to the front of the list.
// If over the cap, it will remove the back item then push new item to front
// It returns the previous value associated with the specified key, or nil if there was no mapping for the key.
// A nil return can also indicate that the map previously associated nil with the specified key.
func (sf *LinkedMap) PushFront(k, v interface{}) interface{} {
	var retVal interface{}

	if old, ok := sf.data[k]; ok {
		retVal = old.Value.(store).value
		old.Value = store{k, v}
		sf.l.MoveToFront(old)
	} else {
		if sf.cap != 0 && sf.l.Len() >= sf.cap {
			e := sf.l.Back()
			delete(sf.data, e.Value.(store).key)
			sf.l.Remove(e)
		}
		sf.data[k] = sf.l.PushFront(store{k, v})
	}
	return retVal
}

// PushBack associates the specified value with the specified key in this map.
// If the map previously contained a mapping for the key,
// the old value is replaced by the specified value. and then move the item to the back of the list.
// If over the cap, it will remove the back item then push new item to back
func (sf *LinkedMap) PushBack(k, v interface{}) interface{} {
	var retVal interface{}

	if old, ok := sf.data[k]; ok {
		retVal = old.Value.(store).value
		old.Value = store{k, v}
		sf.l.MoveToBack(old)
	} else {
		if sf.cap != 0 && sf.l.Len() >= sf.cap {
			e := sf.l.Front()
			delete(sf.data, e.Value.(store).key)
			sf.l.Remove(e)
		}
		sf.data[k] = sf.l.PushBack(store{k, v})
	}
	return retVal
}

// Poll return the front element value and then remove from list
func (sf *LinkedMap) Poll() (k interface{}, v interface{}, exist bool) {
	return sf.PollFront()
}

// PollFront return the front element value and then remove from list
func (sf *LinkedMap) PollFront() (k interface{}, v interface{}, exist bool) {
	if e := sf.l.Front(); e != nil {
		st := e.Value.(store)
		delete(sf.data, st.key)
		sf.l.Remove(e)
		return st.key, st.value, true
	}
	return nil, nil, false
}

// PollBack return the back element value and then remove from list
func (sf *LinkedMap) PollBack() (k interface{}, v interface{}, exist bool) {
	if e := sf.l.Back(); e != nil {
		st := e.Value.(store)
		delete(sf.data, st.key)
		sf.l.Remove(e)
		return st.key, st.value, true
	}
	return nil, nil, false
}

// Remove removes the mapping for a key from this map if it is present.
// It returns the value to which this map previously associated the key, and true,
// or nil and false if the map contained no mapping for the key.
func (sf *LinkedMap) Remove(k interface{}) (interface{}, bool) {
	if oldElement, ok := sf.data[k]; ok {
		retVal := oldElement.Value.(store).value
		delete(sf.data, k)
		sf.l.Remove(oldElement)
		return retVal, true
	}
	return nil, false
}

// ContainsKey returns true if this map contains a mapping for the specified key.
func (sf *LinkedMap) ContainsKey(k interface{}) bool {
	_, ok := sf.data[k]
	return ok
}

// ContainsValue returns true if this map maps one or more keys to the specified value.
func (sf *LinkedMap) ContainsValue(v interface{}) bool {
	for e := sf.l.Front(); e != nil; e = e.Next() {
		if sf.compare(e.Value.(store).value, v) {
			return true
		}
	}
	return false
}

// Get returns the value to which the specified key is mapped, or nil if this map contains no mapping for the key.
func (sf *LinkedMap) Get(k interface{}, defaultValue ...interface{}) interface{} {
	if old, ok := sf.data[k]; ok {
		sf.l.MoveToBack(old)
		return old.Value.(store).value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return nil
}

// Peek return the front element value .
func (sf *LinkedMap) Peek() (k interface{}, v interface{}, exist bool) {
	return sf.PeekFront()
}

// PeekFront return the front element value.
func (sf *LinkedMap) PeekFront() (k interface{}, v interface{}, exist bool) {
	if e := sf.l.Front(); e != nil {
		st := e.Value.(store)
		return st.key, st.value, true
	}
	return nil, nil, false
}

// PeekBack return the back element value .
func (sf *LinkedMap) PeekBack() (k interface{}, v interface{}, exist bool) {
	if e := sf.l.Back(); e != nil {
		st := e.Value.(store)
		return st.key, st.value, true
	}
	return nil, nil, false
}

// Iterator iterator the list
func (sf *LinkedMap) Iterator(cb func(k interface{}, v interface{}) bool) {
	for e := sf.l.Front(); e != nil; e = e.Next() {
		st := e.Value.(store)
		if cb == nil || !cb(st.key, st.value) {
			return
		}
	}
}

// ReverseIterator reverse iterator the list
func (sf *LinkedMap) ReverseIterator(cb func(k interface{}, v interface{}) bool) {
	for e := sf.l.Back(); e != nil; e = e.Prev() {
		st := e.Value.(store)
		if cb == nil || !cb(st.key, st.value) {
			return
		}
	}
}

func (sf *LinkedMap) compare(v1, v2 interface{}) bool {
	if sf.cmp != nil {
		return sf.cmp.Compare(v1, v2) == 0
	}
	return v1 == v2
}
