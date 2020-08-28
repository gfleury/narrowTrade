package main

import (
	"context"
	"log"
	"os"
	"sort"

	"github.com/gfleury/narrowTrade/models"
	"github.com/gfleury/narrowTrade/saxo_oauth2"
	"github.com/gfleury/narrowTrade/saxo_openapi"
	"github.com/gfleury/narrowTrade/trader"

	"github.com/gfleury/narrowTrade/tests"

	iex "github.com/goinvest/iexcloud/v2"
)

func main() {
	ctx := context.Background()

	oauth2cfg := tests.GetTestOauth()

	token, err := saxo_oauth2.GetToken(ctx, oauth2cfg)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	tokenSource := oauth2cfg.TokenSource(ctx, token)

	client := saxo_openapi.NewAPIClient(saxo_openapi.NewConfiguration())
	auth := context.WithValue(ctx, saxo_openapi.ContextOAuth2, tokenSource)
	auth = context.WithValue(auth, saxo_openapi.ContextMockedDataID, "001")

	ma := &models.ModeledAPI{
		Ctx:    auth,
		Client: client,
	}

	acc, err := ma.GetAccounts()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println(acc.GetAccountKey(0))

	t := trader.StockNaive{
		BasicSaxoTrader: &trader.BasicSaxoTrader{
			AccountKey: acc.GetAccountKeyMe(),
			ModeledAPI: ma,
			IEXAPI:     iex.NewClient("sk_be7d6c55dfb6412e8e8c40bc648b11c1", iex.WithBaseURL("https://cloud.iexapis.com/v1")),
		},
	}

	watchlistRequest := &models.WatchlistRequest{
		Arguments: models.Arguments{
			WatchlistID: "2491746",
			AssetTypes:  "Bond,Stock,StockIndex,CfdOnStock,CfdOnIndex,FxSpot,FxForwards,FxVanillaOption,FxKnockInOption,FxKnockOutOption,FxOneTouchOption,FxNoTouchOption",
			Index:       0,
			RowCount:    100,
		},
		RefreshRate: 500,
		Format:      "application/json",
		ContextID:   "7865091208",
		ReferenceID: "10",
	}

	watchlist, err := t.ModeledAPI.GetWatchlist(watchlistRequest)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	uics := make([]int, len(watchlist.Snapshot.Rows))

	sort.Slice(watchlist.Snapshot.Rows, func(i, j int) bool {
		return watchlist.Snapshot.Rows[i].ThreeMonthsReturnPct*watchlist.Snapshot.Rows[i].Price >
			watchlist.Snapshot.Rows[j].ThreeMonthsReturnPct*watchlist.Snapshot.Rows[j].Price
	})

	for idx, instrument := range watchlist.Snapshot.Rows {
		uics[idx] = instrument.Uic
	}

	err = t.Trade(
		trader.StockNaiveParameter{
			Symbols:       uics[0:10],
			PercentLoss:   2,
			PercentProfit: 5,
			TotalInvest:   0,
		})

	if err != nil {
		log.Println(err)
		log.Println(models.GetStringError(err))
		os.Exit(1)
	}

	// Always write token back if everything went ok
	token, err = tokenSource.Token()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = saxo_oauth2.PersistToken(token)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
