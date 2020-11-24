// +build !integration
// +build unit

package utils

import (
	"time"

	check "gopkg.in/check.v1"
)

func (s *Suite) TestCacheEntry(c *check.C) {
	cache := &Cache{}

	cache.Put("key", true, time.Millisecond*200)
	v := cache.Get("key")
	c.Assert(v, check.NotNil)
	time.Sleep(time.Millisecond * 200)
	v = cache.Get("key")
	c.Assert(v, check.IsNil)
}
