package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gfleury/narrowTrade/analysis"

	"github.com/gfleury/narrowTrade/models"
	"github.com/gfleury/narrowTrade/portfolio"
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

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			ExitSaveToken(tokenSource, fmt.Errorf("CTRL-Ced"))
		case syscall.SIGTERM:
			ExitSaveToken(tokenSource, fmt.Errorf("CTRL-Ced"))
		}
	}()

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

	IEXAnalyser := &analysis.IEXAnalyser{
		Ctx:    ctx,
		Client: iex.NewClient("sk_be7d6c55dfb6412e8e8c40bc648b11c1", iex.WithBaseURL("https://cloud.iexapis.com/v1")),
	}
	IEXAnalyser.Init()

	basicTrader := &trader.BasicSaxoTrader{
		AccountKey:         acc.GetAccountKeyMe(),
		ModeledAPI:         ma,
		InstrumentAnalyser: IEXAnalyser,
	}

	stockNaive := trader.StockNaive{
		BasicSaxoTrader: basicTrader,
	}

	forexNaive := trader.ForexNaive{
		StockNaive: &stockNaive,
	}

	p := &portfolio.Portfolio{
		Traders: map[portfolio.TraderName]trader.ComplexTrader{
			portfolio.StockNaive: &stockNaive,
			portfolio.ForexNaive: &forexNaive,
		},
	}

	f, err := os.Open("portfolio.yaml")
	if err != nil {
		ExitSaveToken(tokenSource, err)
	}
	defer f.Close()

	err = p.Load(f)
	if err != nil {
		ExitSaveToken(tokenSource, err)
	}

	err = p.Validate()
	if err != nil {
		ExitSaveToken(tokenSource, err)
	}

	for {
		err = p.Rebalance()
		if err != nil {
			break
		}

		token, err := tokenSource.Token()
		if err != nil {
			break
		}

		interval := time.Until(token.Expiry) - 30*time.Second
		time.Sleep(interval)

		log.Println("Slept", interval)
	}

	ExitSaveToken(tokenSource, err)
}

func ExitSaveToken(tokenSource oauth2.TokenSource, previousErr error) {

	// Always write token back if everything went ok
	token, err := tokenSource.Token()
	if err != nil {
		log.Println(previousErr)
		log.Println(err)
		os.Exit(1)
	}
	err = saxo_oauth2.PersistToken(token)
	if err != nil {
		log.Println(previousErr)
		log.Println(err)
		os.Exit(1)
	}

	if previousErr != nil {
		log.Println(previousErr)
		os.Exit(1)
	}
}
