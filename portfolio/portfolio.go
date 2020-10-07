package portfolio

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/gfleury/narrowTrade/config"
	"github.com/gfleury/narrowTrade/external_lists"
	"github.com/gfleury/narrowTrade/models"
	"github.com/gfleury/narrowTrade/trader"

	yaml "gopkg.in/yaml.v2"
)

type TraderName string

const (
	StockNaive TraderName = "StockNaive"
	ForexNaive            = "ForexNaive"
)

type Investment struct {
	ExternalList    string                `yaml:",omitempty"`
	WatchlistID     string                `yaml:",omitempty"`
	WatchlistItems  int                   `yaml:",omitempty"`
	Symbols         []string              `yaml:",omitempty"`
	ValuePercentage float64               `yaml:",omitempty"`
	ValueAbsolute   float64               `yaml:",omitempty"`
	Trader          TraderName            `yaml:",omitempty"`
	Parameters      trader.TradeParameter `yaml:",omitempty"`
}

type Portfolio struct {
	Distribution []Investment
	Traders      map[TraderName]trader.ComplexTrader `yaml:"-"`
	Tags         Tags                                `yaml:"-"`
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
		if investment.Parameters.PercentLoss == 0 ||
			investment.Parameters.PercentProfit == 0 {
			return fmt.Errorf("Investment parameters %v must contain PercentProfit and PercentLoss", investment)
		}
		if investment.ValuePercentage == 0 &&
			investment.ValueAbsolute == 0 {
			return fmt.Errorf("Investment %v must contain ValuePercentage or ValueAbsolute", investment)
		}
		if investment.ExternalList == "" && investment.WatchlistID == "" &&
			(investment.Symbols == nil || len(investment.Symbols) == 0) {
			return fmt.Errorf("Investment %v must contain ExternalList, WatchlistID or Symbols", investment)
		}
	}
	if totalPercentage > 100 {
		return fmt.Errorf("Portifolio investments can't be more than 100%%, right now is %f", totalPercentage)
	}
	return nil
}

func (p *Portfolio) Rebalance() error {
	err := p.ReadTags()
	if err != nil {
		return err
	}

	for id, investment := range p.Distribution {
		t, ok := p.Traders[investment.Trader]
		if !ok {
			log.Printf("Trader not found: %s, skipping it", investment.Trader)
			continue
		}

		openOrders, err := t.Api().GetOrdersMe()
		if err != nil {
			return fmt.Errorf("Unable to get open orders, bouncing back: %s", err)
		}

		totalInvested := p.Tags[id].GetOpenOrdersTotal(openOrders)

		balance, err := t.Api().GetBalanceMe()
		if err != nil {
			return fmt.Errorf("Unable to get balance, bouncing back: %s", err)
		}

		if balance.MarginAvailableForTrading < config.MINIMUM_TRADE_VALUE {
			log.Printf("Available margin is lower than Minimum amount for trading: %f < %f",
				balance.MarginAvailableForTrading, config.MINIMUM_TRADE_VALUE)
			continue
		}

		if totalInvested >= balance.MarginAvailableForTrading*(investment.ValuePercentage/100) {
			log.Printf("Invested amount is already higher than requested: %f >= %f * %f ",
				totalInvested, balance.MarginAvailableForTrading, investment.ValuePercentage/100)
			continue
		}
		log.Printf("Balancing invested amount as it is lower than requested: %f >= %f * %f ",
			totalInvested, balance.MarginAvailableForTrading, investment.ValuePercentage/100)
		// Get available cash from account (cash that is not hold by positions/orders)
		availableCash := balance.InitialMargin.MarginAvailable

		if investment.WatchlistID != "" {
			watchlist, err := t.Api().GetWatchlistByID(investment.WatchlistID)
			if err != nil {
				return err
			}

			uics := make([]int, len(watchlist.Snapshot.Rows))

			sort.Slice(watchlist.Snapshot.Rows, func(i, j int) bool {
				return watchlist.Snapshot.Rows[i].ThreeMonthsReturnPct*watchlist.Snapshot.Rows[i].Price >
					watchlist.Snapshot.Rows[j].ThreeMonthsReturnPct*watchlist.Snapshot.Rows[j].Price
			})

			for idx, instrument := range watchlist.Snapshot.Rows {
				uics[idx] = instrument.Uic
			}

			// Default list to first 10 items of list
			if investment.WatchlistItems == 0 {
				investment.WatchlistItems = 10
			}

			if len(uics) < investment.WatchlistItems {
				investment.WatchlistItems = len(uics)
			}

			investment.Parameters.Symbols = uics[:investment.WatchlistItems]

			if investment.ValuePercentage != 0 {
				investment.Parameters.TotalInvest = availableCash * (investment.ValuePercentage / 100)
			} else {
				investment.Parameters.TotalInvest = investment.ValueAbsolute
			}
		} else if investment.ExternalList != "" {
			var externalList external_lists.ExternalList

			switch investment.ExternalList {
			case "SP100":
				externalList = external_lists.NewWikipediaSP100()
			case "SP500":
				externalList = external_lists.NewWikipediaSP500()
			default:
				return fmt.Errorf("Unknown ExternalList provider %s", investment.ExternalList)
			}

			elSymbols, err := externalList.GetSymbols()
			if err != nil {
				return err
			}

			log.Printf("Getting instruments for symbols, it may take a while: %s",
				strings.Join(elSymbols, ", "))

			// Default list to first 10 items of list
			if investment.WatchlistItems == 0 {
				investment.WatchlistItems = 10
			}

			uics := []int{}

			for idx, symbol := range elSymbols {
				if idx > investment.WatchlistItems {
					break
				}
				i, err := t.Api().GetInstrument(symbol)
				if err != nil {
					log.Printf("Failed to get Symbol %s from %s", symbol, investment.ExternalList)
					continue
				}
				uics = append(uics, int(i.GetID()))
			}

			if len(uics) < investment.WatchlistItems {
				investment.WatchlistItems = len(uics)
			}

			investment.Parameters.Symbols = uics[:investment.WatchlistItems]

			if investment.ValuePercentage != 0 {
				investment.Parameters.TotalInvest = availableCash * (investment.ValuePercentage / 100)
			} else {
				investment.Parameters.TotalInvest = investment.ValueAbsolute
			}
		}

		// Read Symbols From Investment
		for _, symbol := range investment.Symbols {
			i, err := t.Api().GetInstrument(symbol)
			if err != nil {
				return fmt.Errorf("Unable to fetch Symbol: %s , bouncing back: %s", symbol, err)
			}
			investment.Parameters.Symbols = append(investment.Parameters.Symbols, int(i.GetID()))
		}

		tradeErr := t.Trade(investment.Parameters)

		orders, prices := t.GetOrders()

		p.Tags.AddTag(id, generateTags(orders, prices))
		err = p.Tags.Save()

		if tradeErr != nil {
			log.Printf("Trading failed with: %s", models.GetStringError(tradeErr))
			return tradeErr
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Portfolio) ReadTags() error {
	f, err := os.OpenFile(".narrow_tags", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	tags := make(Tags, len(p.Distribution))
	p.Tags = tags

	err = p.Tags.Load(f)

	return err
}
