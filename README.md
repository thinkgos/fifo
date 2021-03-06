# container

container implements some containers, currently the containers are not thread-safe. `safe` path support thread-safe.

[![GoDoc](https://godoc.org/github.com/thinkgos/container?status.svg)](https://godoc.org/github.com/thinkgos/container)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/thinkgos/container?tab=doc)
[![Build Status](https://www.travis-ci.org/thinkgos/container.svg?branch=master)](https://www.travis-ci.org/thinkgos/container)
[![codecov](https://codecov.io/gh/thinkgos/container/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkgos/container)
![Action Status](https://github.com/thinkgos/container/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/container)](https://goreportcard.com/report/github.com/thinkgos/container)
[![License](https://img.shields.io/github/license/thinkgos/container)](https://github.com/thinkgos/container/raw/master/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/thinkgos/container)](https://github.com/thinkgos/container/tags)

- **[How to use this repo](#how-to-use-this-package)**
- **[Containers](#Containers-Interface)**
  - [Sets](#sets) move to [sets](https://github.com/things-go/sets)
  - [Stack](#stack) 
    - stack use container/list.
    - quick stack use builtin slice.
  - [Queue](#queue) 
    - queue use container/list
    - quick queue use builtin slice.
  - [PriorityQueue](#priorityqueue) use builtin slice with container/heap
  - [LinkedList](#linkedlist) use container/list
  - [ArrayList](#arraylist) use builtin slice.
  - [LinkedMap](#linkedMap) use container/list and builtin map.
  - [topic](#topic) topic tree like MQTT topic
  - [trie](#trie) trie tree
- **[safe container](#safe-container)**
  - [fifo](#fifo) 
    > FIFO solves this use case:
    > * You want to process every object (exactly) once.
    > * You want to process the most recent version of the object when you process it.
    > * You do not want to process deleted objects, they should be removed from the queue.
    > * You do not want to periodically reprocess objects.

  - [heap](#heap) Heap is a thread-safe producer/consumer queue that implements a heap data structure.It can be used to implement priority queues and similar data structures.
- **[others](#others)**
  - [Comparator](#Comparator) 
    - [Sort](#sort) sort with Comparator interface
    - [Heap](#heap) heap with Comparator interface
    
## Donation

if package help you a lot,you can support us by:

**Alipay**

![alipay](https://github.com/thinkgos/thinkgos/blob/master/asserts/alipay.jpg)

**WeChat Pay**

![wxpay](https://github.com/thinkgos/thinkgos/blob/master/asserts/wxpay.jpg)