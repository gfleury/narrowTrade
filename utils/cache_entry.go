package utils

import "time"

type CacheEntry struct {
	Value  interface{}
	Expire time.Time
}

type Cache map[string]*CacheEntry

func (c Cache) Put(k string, v interface{}, ex time.Duration) {
	c[k] = &CacheEntry{v, time.Now().Add(ex)}
}

func (c Cache) Get(k string) interface{} {
	if ce, ok := c[k]; ok {
		if !time.Now().After(ce.Expire) {
			return ce.Value
		}
		delete(c, k)
	}
	return nil
}
