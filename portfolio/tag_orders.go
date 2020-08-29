package portfolio

import (
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Tag struct {
	OrdersIDs []string
	ID        int
}

type Tags []Tag

func (t *Tags) AddTag(id int, ordersIDs []string) {
	*t = append(*t, Tag{ID: id, OrdersIDs: ordersIDs})
}

func (t Tags) Save() error {
	data, err := yaml.Marshal(&t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(".narrow_tags", data, 0644)
}

func (t Tags) Load(r io.Reader) error {
	if t == nil {
		t = Tags{}
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &t)
}
