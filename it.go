package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gfleury/intensiveTrade/saxo_oauth2"
	"github.com/gfleury/intensiveTrade/saxo_openapi"

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

	r, httpResponse, err := client.AccountHistoryApi.GetStandardPeriodAccountValues(auth, "ma6hsMLLkwRlPXxGx|KepQ==", nil)
	if err != nil {
		fmt.Println(httpResponse)
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(r)
}
