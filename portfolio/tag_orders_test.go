// +build !integration
// +build unit

package portfolio

import (
	"strings"

	check "gopkg.in/check.v1"
)

func (s *Suite) TestTagsSave(c *check.C) {
	t := &Tags{
		Tag{
			ID: 0,
			OrdersIDs: map[string]float64{
				"123": 1.7,
			},
		},
		Tag{
			ID: 1,
		},
	}
	c.Assert(t, check.NotNil)

	err := t.Save()
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestTagsLoad(c *check.C) {
	data := strings.NewReader(`
- ordersids:
    "123": 1.7
  id: 0
- ordersids:
    "123": 1.7
  id: 1
`)
	t := &Tags{}

	err := t.Load(data)
	c.Assert(err, check.IsNil)

}

func (s *Suite) TestAddTag(c *check.C) {
	t := make(Tags, 2)
	t.AddTag(0, map[string]float64{
		"123": 1.0,
		"234": 1.0,
	})

	c.Assert(t, check.DeepEquals, Tags{
		{ID: 0, OrdersIDs: map[string]float64{
			"123": 1.0,
			"234": 1.0,
		}},
		{},
	})
}

func (s *Suite) TestTagsLoadAndSave(c *check.C) {
	data := strings.NewReader(`
- ordersids:
    "123": 1.7
    "234": 1.7
  id: 0
- ordersids:
    "123": 1.7
    "234": 1.7
  id: 1
`)
	t := make(Tags, 2)

	err := t.Load(data)
	c.Assert(err, check.IsNil)

	c.Assert(t, check.DeepEquals, Tags{
		{ID: 0, OrdersIDs: map[string]float64{
			"123": 1.7,
			"234": 1.7,
		}},
		{ID: 1, OrdersIDs: map[string]float64{
			"123": 1.7,
			"234": 1.7,
		}},
	})

	err = t.Save()
	c.Assert(err, check.IsNil)
}
