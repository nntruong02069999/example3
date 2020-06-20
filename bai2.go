package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type DataStore struct {
	sync.Mutex
	data      map[string]string
	indentity int64
}

func usingMutex() {
	dataStore := newData()
	for i := 1; i <= 3; i++ {
		go func() {
			for i := 1; i <= 1000; i++ {
				dataStore.set("Key", "Value")
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(dataStore.data)
}

// Khoi tao DataStore

func newData() *DataStore {
	return &DataStore{
		data: make(map[string]string),
		indentity: 0,
	}
}

func (ds *DataStore) set(key string, value string) {
	ds.Lock()
	changeKey := key + strconv.FormatInt(ds.indentity, 10)
	changeValue := value + strconv.FormatInt(ds.indentity, 10)
	ds.indentity++
	defer ds.Unlock()
	ds.data[changeKey] = changeValue
}
