package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

// singleton pattern in golang
type DbConnection struct{}

var (
	dbConnOnce sync.Once
	conn       *DbConnection
)

func GetConnection() *DbConnection {
	dbConnOnce.Do(func() {
		conn = &DbConnection{}
	})
	return conn
}

//https://blog.chuie.io/posts/synconce/

type CacheEntry struct {
	data []byte
	once *sync.Once
}

type QueryClient struct {
	cache map[string]*CacheEntry
	mutex *sync.RWMutex
	map   *sync.Map
}

// old way of doing this without Map.
func (c *QueryClient) DoQuery(name string) []byte {
	c.mutex.RLock()
	entry, found := c.cache[name]
	if !found {
		// Create a new entry if one does not exist already
		entry = &CacheEntry{
			once: new(sync.Once),
		}
		c.cache[name] = entry
	}
	c.mutex.RUnlock()

	// Now when we invoke `.Do`, if there is an ongoing simultaneous operation,
	// it will block until it has completed (and `entry.data` is populated).
	// Or if the operation has already completed once before, this call is a no-op and does not block.
	entry.once.Do(func() {
		resp, err := http.Get("https://upstream.api/?query=" + url.QueryEscape(name))
		if err != nil {
			fmt.Printf("error %v", err)
		}
		defer resp.Body.Close()
		entry.data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error %v", err)
		}
	})
	return entry.data
}

c
// function using sync.Map 
// tho in the documentation, it is advised that Other than Once and WaitGroup types,
// most are intended for use by low-level library routines.
func (c *QueryClient) DoWithMap(name string) []byte {
	entry, found := c.map.Load(name)
	if !found {
		entry = &CacheEntry{
			once: new(sync.Once),
		})
		c.map.Store(name, entry)
	}
	entry.once.Do(func() {
		resp, err := http.Get("https://upstream.api/?query=" + url.QueryEscape(name))
		if err != nil {
			fmt.Printf("error %v", err)
		}
		defer resp.Body.Close()
		entry.data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error %v", err)
		}
	})
	return entry.data
}