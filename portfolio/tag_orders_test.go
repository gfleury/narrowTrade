// +build !integration
// +build unit

package portfolio

import (
	"strings"

	check "gopkg.in/check.v1"
)

func (s *Suite) TestTagsSave(c *check.C) {
	t := &Tags{
		Tag{},
		Tag{},
	}

	c.Assert(t, check.NotNil)

	err := t.Save()
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestTagsLoad(c *check.C) {
	data := strings.NewReader(`
- ordersids: ["1223321"]
  id: 0
- ordersids: ["1223321"]
  id: 1
`)
	t := &Tags{}

	err := t.Load(data)
	c.Assert(err, check.IsNil)

}
