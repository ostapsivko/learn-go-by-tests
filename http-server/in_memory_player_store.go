package main

import "sync"

type InMemoryPlayerStore struct {
	store map[string]int
	mux   sync.RWMutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mux.RLock()
	defer i.mux.RUnlock()

	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mux.Lock()
	defer i.mux.Unlock()

	i.store[name]++
}
