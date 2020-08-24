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

	t := trader.BasicSaxoTrader{
		AccountKey: acc.GetAccountKeyMe(),
		ModeledAPI: ma,
		IEXClient:  iex.NewClient("sk_be7d6c55dfb6412e8e8c40bc648b11c1", iex.WithBaseURL("https://cloud.iexapis.com/v1")),
	}

	gainers, err := t.IEXClient.Gainers(ctx)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	symbols := []string{}

	for _, instrument := range gainers {
		symbols = append(symbols, instrument.Symbol)
	}

	err = t.BuyStocksNaive(symbols[0:5], 2, 5)

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
