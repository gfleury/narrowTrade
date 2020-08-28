package portfolio

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/gfleury/narrowTrade/trader"
	yaml "gopkg.in/yaml.v2"
)

type TraderName string

const (
	StockNaive TraderName = "StockNaive"
	ForexNaive            = "ForexNaive"
)

type Investment struct {
	WatchlistID     string      `yaml:",omitempty"`
	ValuePercentage float64     `yaml:",omitempty"`
	ValueAbsolute   float64     `yaml:",omitempty"`
	Trader          TraderName  `yaml:",omitempty"`
	Parameters      interface{} `yaml:",omitempty"`
}

type Portfolio struct {
	Distribution []Investment
	Traders      map[TraderName]trader.ComplexTrader
}

func (p *Portfolio) Save() error {
	data, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("testdata/portfolio", data, 0644)
}

func (p *Portfolio) Load(r io.Reader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, p)
}

func (p *Portfolio) Validate() error {
	totalPercentage := 0.0
	for _, investment := range p.Distribution {
		totalPercentage += investment.ValuePercentage
		if investment.ValuePercentage == 0 &&
			investment.ValueAbsolute == 0 {
			return fmt.Errorf("Investment %v must contain ValuePercentage or ValueAbsolute", investment)
		}
	}
	if totalPercentage > 100 {
		return fmt.Errorf("Portifolio investments can't be more than 100%%, right now is %f", totalPercentage)
	}
	return nil
}

func (p *Portfolio) Rebalance() error {
	for _, investment := range p.Distribution {
		t := p.Traders[investment.Trader]
		err := t.Trade(investment.Parameters)
		if err != nil {
			return err
		}
	}
	return nil
}
