package portfolio

import (
	"io"
	"io/ioutil"

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
	data, err := yaml.Marshal(t)
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
	return yaml.Unmarshal(data, &t)
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
