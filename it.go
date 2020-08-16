package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gfleury/intensiveTrade/saxo_models"
	"github.com/gfleury/intensiveTrade/saxo_oauth2"
	"github.com/gfleury/intensiveTrade/saxo_openapi"

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

	i, err := ma.GetInstrument("VWAGY")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(i)

	// trader := trader.BasicSaxoTrader{
	// 	AccountKey: acc.GetAccountKeyMe(),
	// 	API:        ma,
	// }

	// or, err := trader.Buy(i, saxo_models.Market, 1000)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(or)

	iexClient := iex.NewClient("Tsk_c1ee184113dc42bdae6907741c07d6ac", iex.WithBaseURL("https://sandbox.iexapis.com/v1"))
	bs, err := iexClient.TechIndicators(ctx, "VWAGY", "ema", "3m", iex.NewTechIndicatorOpt("period", "10"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(bs)

	// Always write token back if everything went ok
	err = saxo_oauth2.PersistToken(token)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
