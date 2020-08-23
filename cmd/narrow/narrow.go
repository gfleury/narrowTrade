package main

import (
	"context"
	"log"
	"os"

	"github.com/gfleury/narrowTrade/models"
	"github.com/gfleury/narrowTrade/saxo_oauth2"
	"github.com/gfleury/narrowTrade/saxo_openapi"
	"github.com/gfleury/narrowTrade/trader"

	"github.com/gfleury/narrowTrade/tests"

	iex "github.com/goinvest/iexcloud/v2"
	"golang.org/x/oauth2"
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
	auth := context.WithValue(oauth2.NoContext, saxo_openapi.ContextOAuth2, tokenSource)
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

	t := trader.BasicSaxoTrader{
		AccountKey: acc.GetAccountKeyMe(),
		API:        ma,
	}

	iexClient := iex.NewClient("sk_be7d6c55dfb6412e8e8c40bc648b11c1", iex.WithBaseURL("https://cloud.iexapis.com/v1"))
	mostAttractives, err := iexClient.Gainers(ctx)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	buyingInstruments := make([]models.Instrument, len(mostAttractives))
	iexQuotes := make([]iex.Quote, len(mostAttractives))
	for idx, quote := range mostAttractives {
		log.Println("Getting Saxo Bank symbol for", quote.Symbol)
		i, err := ma.GetInstrument(quote.Symbol)
		if err != nil {
			continue
		}
		buyingInstruments[idx] = i

		log.Println("Trying to get price IEXCloud price", i.GetAssetType(), i.GetSymbol())
		iexQuotes[idx], err = iexClient.Quote(ctx, mostAttractives[idx].Symbol)
		if err != nil {
			iexQuotes[idx] = iex.Quote{}
		}
	}

	t.BuyStocksNaive(buyingInstruments, iexQuotes)

	// Always write token back if everything went ok
	token, err = tokenSource.Token()
	if err != nil {
		log.Println(err)
	}
	err = saxo_oauth2.PersistToken(token)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
