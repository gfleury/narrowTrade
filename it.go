package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gfleury/intensiveTrade/saxo_models"
	"github.com/gfleury/intensiveTrade/saxo_oauth2"
	"github.com/gfleury/intensiveTrade/saxo_openapi"
	"github.com/gfleury/intensiveTrade/trader"

	iex "github.com/goinvest/iexcloud/v2"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	oauth2cfg := &oauth2.Config{
		ClientID:     "4e7b5f42dae64ad889cc4e0ae499c398",
		ClientSecret: "abcdefghijklmn",
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://sim.logonvalidation.net/authorize",
			TokenURL:  "https://sim.logonvalidation.net/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: "http://localhost/redirect",
	}

	token, err := saxo_oauth2.GetToken(ctx, oauth2cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tokenSource := oauth2cfg.TokenSource(ctx, token)

	client := saxo_openapi.NewAPIClient(saxo_openapi.NewConfiguration())
	auth := context.WithValue(oauth2.NoContext, saxo_openapi.ContextOAuth2, tokenSource)
	auth = context.WithValue(auth, saxo_openapi.ContextMockedDataID, "001")

	ma := &saxo_models.ModeledAPI{
		Ctx:    auth,
		Client: client,
	}

	acc, err := ma.GetAccounts()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(acc.GetAccountKey(0))

	bal, err := ma.GetBalanceMe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(bal)

	t := trader.BasicSaxoTrader{
		AccountKey: acc.GetAccountKeyMe(),
		API:        ma,
	}

	iexClient := iex.NewClient("sk_be7d6c55dfb6412e8e8c40bc648b11c1", iex.WithBaseURL("https://cloud.iexapis.com/v1"))
	mostAttractives, err := iexClient.Gainers(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buyingInstruments := make([]saxo_models.Instrument, len(mostAttractives))
	iexQuotes := make([]iex.Quote, len(mostAttractives))
	for idx, quote := range mostAttractives {
		fmt.Println("Getting Saxo Bank symbol for", quote.Symbol)
		i, err := ma.GetInstrument(quote.Symbol)
		if err != nil {
			continue
		}
		buyingInstruments[idx] = i

		fmt.Println("Trying to get price IEXCloud price", i.GetAssetType(), i.GetSymbol())
		iexQuotes[idx], err = iexClient.Quote(ctx, mostAttractives[idx].Symbol)
		if err != nil {
			iexQuotes[idx] = iex.Quote{}
		}
	}

	t.BuyStocksNaive(buyingInstruments, iexQuotes)

	// Always write token back if everything went ok
	token, err = tokenSource.Token()
	if err != nil {
		fmt.Println(err)
	}
	err = saxo_oauth2.PersistToken(token)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
