# narrowTrader

## Roadmap:
- Full-fledged robo-trading (non low latency) - OK
- Portfolio Automation - OK
- rebalancing portfolio - OK
- auto investing - OK
- remove all IEX code and create Instrument analysis interface. - OK
- Stock Trading based on dynamic data from external systems or from the API - OK
- Forex Trading based on dynamic data from external systems or from the API - OK
- Rebalance CashPerSymbol if any order fail during the Trade() - OK
- StockNaive with pricing per percentage and moving averages. - OK
- Portfolio rebalancing and monitoring across clients for smaller wealth managers or hedge funds
- Check how to implement analysis for resistance/support
- Trading on specific times
- Only Trade when exchange is about 20 minutes to open
- StockNaive GetCashPerSymbol testing for different edge cases
- API's context with timeout and exponential backoff 
  `sleep = min(2000, random.uniform(5, sleep * 3))`
- Removed closed orders from .narrow_tags

Standalone tools:
- FX hedging - sheets can be configured to automatically hedge other positions or transactions once these are committed

## Nice to check technical projects:
- https://github.com/sdcoffey/techan
- https://github.com/ivopetiz/technical-indicators

## Strategies to implement
### Mean Reversion
The mean reversion strategy is based on a straightforward assumption – if the price of a coin shift from its average, then it’s eventually going to revert back to it. 

### Momentum Trading
The core philosophy behind this is the belief that the prices of an asset will spike above its average and then run out of momentum and fall down. 

### Naive Bayes
