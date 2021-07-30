package main

import (
	"fmt"
	"sort"
	"sync"
)

// Link: https://blog.golang.org/maps

// Linked list
type Node struct {
	Next  *Node
	Value interface{}
}

type Person struct {
	Name  string
	Likes []string
}

type Key struct {
	Path, Country string
}

func main() {
	// Will result nil
	// It can be used for read, but if you try to write it will runtime panic
	var m map[string]int

	// Best practice to initialize via make() if it is empty
	// make() will allocate and initialize a hash map and return it's pointer
	m = make(map[string]int)

	// set key with value
	// key - route
	// value - 66
	m["route"] = 66

	// if the key does not exist, map will return zero value for it's value type
	// in this case value type is int, so it returns 0
	j := m["root"]
	fmt.Println(j)

	// map key be used in len(), will return number of keys.
	// it will not count underlying maps
	// e.g:
	// {
	//   "a": 1,
	//   "b": 2,
	//   "c": {
	//      "d": 3,
	//    },
	// }
	// map will return 3 instead of 4
	n := len(m)
	fmt.Println(n)

	// delete key from the map
	// if the key does not exist, no-op
	delete(m, "route")

	// retrieve value by key
	// also apply zero value type if key does not exist
	// ok will return true if key exists, false otherwise
	i, ok := m["route"]
	fmt.Println(i, ok)

	// will retrieve only status of key existance in map
	// _ is used as a disposable variable
	_, ok = m["route"]

	// iterating key and value from a map
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	// initialization of a map with some data
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	fmt.Println(commits)

	// initialization of empty map, identical as make().
	m = map[string]int{}

	// initialization of Node as nil struct
	var first *Node

	visited := make(map[*Node]bool)
	for n := first; n != nil; n = n.Next {
		// no need to check if key exists
		// if it does not, return default value type
		// in this case it is false
		if visited[n] {
			fmt.Println("cycle detected")
			break
		}
		visited[n] = true
		fmt.Println(n.Value)
	}

	// empty array of Person's
	var people []*Person

	// map with key string and value array of Person's
	likes := make(map[string][]*Person)
	for _, p := range people {
		for _, l := range p.Likes {
			// it will allocate new slice since it is empty by default
			likes[l] = append(likes[l], p)
		}
	}

	for _, p := range likes["cheese"] {
		fmt.Println(p.Name, "likes cheese")
	}

	fmt.Println(len(likes["bacon"]), "people like bacon.")

	// creation of map with key string, value map with key string, value int
	hits := make(map[string]map[string]int)

	n = hits["/doc/"]["au"]
	// will print 0, because key does not exist
	// default value 0
	fmt.Println(n)

	// map with inner map has to check inner map initialization,
	// when trying to set value to inner map key
	add(hits, "/doc/", "au")

	hhits := make(map[Key]int)
	// When an Vietnamese person visits the home page,
	// incrementing (and possibly creating) the appropriate counter is a one-liner
	hhits[Key{"/", "vn"}]++
	n = hhits[Key{"/ref/spec", "ch"}]
	fmt.Println(n)

	// Concurrency
	// One common way to protect maps is with sync.RWMutex.
	// code below is not concurrent
	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	// To read from the counter, take the read lock
	counter.RLock()
	n = counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("some_key", n)

	// To write to the counter, take the write lock
	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()

	// When iterating over a map with a range loop,
	// the iteration order is not specified and is not guaranteed
	// to be the same from one iteration to the next.
	var mm map[int]string
	// If you require a stable iteration order
	// you must maintain a separate data structure that specifies that order.
	var keys []int
	for k := range mm {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", mm[k])
	}
}

func add(m map[string]map[string]int, path, country string) {
	// check if key exists
	// if not, it means that inner map is not initialized
	mm, ok := m[path]
	if !ok {
		// create inner map as value
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}
