package portfolio

import (
	"io"
	"io/ioutil"
	"sort"

	"github.com/gfleury/narrowTrade/models"
	yaml "gopkg.in/yaml.v2"
)

type Tag struct {
	OrdersIDs map[string]float64
	ID        int
}

type Tags []Tag

func (t Tags) AddTag(id int, ordersIDs map[string]float64) {
	t[id].ID = id
	if t[id].OrdersIDs == nil {
		t[id].OrdersIDs = map[string]float64{}
	}
	for orderID, price := range ordersIDs {
		t[id].OrdersIDs[orderID] = price
	}
}

func (t *Tags) Save() error {
	sortT := *t
	sort.Slice(sortT, func(i, j int) bool {
		return sortT[i].ID < sortT[j].ID
	})

	data, err := yaml.Marshal(sortT)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(".narrow_tags", data, 0644)
}

func (t *Tags) Load(r io.Reader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, t)
	if err != nil {
		return err
	}
	idxT := *t
	for idx := range idxT {
		idxT[idx].ID = idx
	}
	return nil
}

func (t *Tag) GetOpenOrdersTotal(ol *models.OrderList) float64 {
	total := 0.0
	for _, order := range ol.Data {
		if price, ok := t.OrdersIDs[order.OrderID]; ok {
			total += price
		}
	}
	return total
}

func generateTags(s []string, p []float64) map[string]float64 {
	t := map[string]float64{}

	for idx, orderID := range s {
		t[orderID] = p[idx]
	}

	return t
}
