package registry

import (
	"fmt"
	"sync"
)

type Registry[K comparable, V any] struct {
	lock  sync.RWMutex
	store map[K]V
}

func New[K comparable, V any]() *Registry[K, V] {
	return &Registry[K, V]{
		store: make(map[K]V),
	}
}

func (r *Registry[K, V]) Register(key K, value V) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if _, exists := r.store[key]; exists {
		panic(fmt.Sprintf("registry: key %v already registered", key))
	}
	r.store[key] = value
}

func (r *Registry[K, V]) Get(key K) V {
	r.lock.RLock()
	defer r.lock.RUnlock()

	v, ok := r.store[key]
	if !ok {
		panic(fmt.Sprintf("registry: key %v not found", key))
	}
	return v
}

func (r *Registry[K, V]) TryGet(key K) (V, bool) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	v, ok := r.store[key]
	return v, ok
}

func (r *Registry[K, V]) All() []V {
	r.lock.RLock()
	defer r.lock.RUnlock()

	list := make([]V, 0, len(r.store))
	for _, v := range r.store {
		list = append(list, v)
	}
	return list
}

func (r *Registry[K, V]) Keys() []K {
	r.lock.RLock()
	defer r.lock.RUnlock()

	keys := make([]K, 0, len(r.store))
	for k := range r.store {
		keys = append(keys, k)
	}
	return keys
}
