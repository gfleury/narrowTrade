package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	s := map[string]interface{}{}
	file, err := ioutil.ReadFile("openapi.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(file, &s)
	if err != nil {
		fmt.Println(err)
		return
	}

	paths := Paths{
		Endpoints: map[string]*Endpoint{},
	}

	operationsIDs := map[string]int{}

	for _, obj := range s["servicegroup"].([]interface{}) {
		key := obj.(map[string]interface{})
		tag := key["Key"].(string)

		for _, obj := range key["Services"].([]interface{}) {
			key := obj.(map[string]interface{})

			for _, obj := range key["Endpoints"].([]interface{}) {
				key := obj.(map[string]interface{})

				urlx := key["Url"].(string)
				u, err := url.Parse(urlx)
				if err != nil {
					log.Fatal(err)
				}
				methodPath := strings.TrimPrefix(u.Path, "/sim")

				method := strings.ToLower(key["HttpMethod"].(string))

				if _, ok := paths.Endpoints[methodPath]; !ok {
					paths.Endpoints[methodPath] = &Endpoint{
						Methods: map[string]*Method{},
					}
				}
				e := paths.Endpoints[methodPath]

				if _, ok := e.Methods[method]; !ok {
					e.Methods[method] = &Method{
						Tags: []string{tag},
					}
				}
				m := e.Methods[method]

				m.OperationID = key["Key"].(string)

				operationsIDs[m.OperationID]++
				if operationsIDs[m.OperationID] == 2 {
					m.OperationID = fmt.Sprintf("%s_%s", m.OperationID, tag)
				} else if operationsIDs[m.OperationID] > 2 {
					m.OperationID = fmt.Sprintf("%s_%s%d", m.OperationID, tag, operationsIDs[m.OperationID])
				}

				for _, oobj := range key["Parameters"].([]interface{}) {
					p := Parameter{}
					param := oobj.(map[string]interface{})
					in := strings.ToLower(param["Origin"].(string))
					dataType := strings.ToLower(param["DataType"].(string))
					if in != "body" {
						switch in {
						case "query-string":
							p.In = "query"
						case "route":
							p.In = "path"
						default:
							p.In = in
						}
						switch dataType {
						case "bool":
							p.Schema.Type = "boolean"
						case "double":
							p.Schema.Type = "number"
						case "decimal":
							p.Schema.Type = "number"
						case "uic":
							p.Schema.Type = "integer"
						case "int":
							p.Schema.Type = "integer"
						case "int []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "integer"
						case "string []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
						case "clientkey":
							p.Schema.Type = "string"
						case "guid":
							p.Schema.Type = "string"
						case "accountkey":
							p.Schema.Type = "string"
						case "userkey":
							p.Schema.Type = "string"
						case "accountgroupkey":
							p.Schema.Type = "string"
						case "datetime":
							p.Schema.Type = "string"
						case "datetime []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
						case "utcdatetime":
							p.Schema.Type = "string"
						case "date":
							p.Schema.Type = "string"
						case "standardperiodtype":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"AllTime", "Month", "Quarter", "Year"}
						case "onboardingdocumenttype":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"Switzerland501", "Switzerland901"}
						case "bookingrequesttype":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"BkAmountId", "CaMasterRecordId", "RelatedTradeId"}
						case "orderentrytype":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"All", "Last"}
						case "scheduledtradingconditionsfieldgroup []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"ScheduledTradingConditions"}
						case "casestatus []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"Canceled", "InformationProvided", "ProblemSolved"}
						case "orderlogstatus []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"Cancelled", "Changed", "DoneForDay",
								"Expired", "Fill", "FinalFill", "Placed", "Working"}
						case "logentrytype []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"Corrections", "DisclaimerActions", "NonTradingActivity",
								"Notifications", "OrderActivity", "Orders", "TradeActivity", "Trades"}
						case "assettype":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"Bond", "Cash", "CfdIndexOption",
								"CfdOnFutures", "CfdOnIndex", "CfdOnStock", "ContractFutures",
								"FuturesOption", "FuturesStrategy", "FxBinaryOption", "FxForwards",
								"FxKnockInOption", "FxKnockOutOption", "FxNoTouchOption", "FxOneTouchOption",
								"FxSpot", "FxVanillaOption", "IpoOnStock", "ManagedFund", "MutualFund",
								"Stock", "StockIndex", "StockIndexOption", "StockOption"}
						case "assettype []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"Bond", "Cash", "CfdIndexOption",
								"CfdOnFutures", "CfdOnIndex", "CfdOnStock", "ContractFutures",
								"FuturesOption", "FuturesStrategy", "FxBinaryOption", "FxForwards",
								"FxKnockInOption", "FxKnockOutOption", "FxNoTouchOption", "FxOneTouchOption",
								"FxSpot", "FxVanillaOption", "IpoOnStock", "ManagedFund", "MutualFund",
								"Stock", "StockIndex", "StockIndexOption", "StockOption"}

						case "chartfieldgroupspec []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"BlockInfo", "ChartInfo", "Data", "DisplayAndFormat"}
						case "chartrequestmode":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"From", "UpTo"}
						case "orderactivityfieldgroup []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"DisplayAndFormat"}
						case "signupflowdocumenttype":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"AccountViewToIb",
								"EsaContract", "FeePaymentAuthorization", "GeneralBusinessTerms",
								"PensionTransferRequest", "PowerOfAttorney", "PowerOfAttorneyToIb",
								"ProofOfIdentity", "ProofOfResidency", "SourceOfFundsDocument",
								"TermsAndConditions", "TermsAndConditionsAldersopsparingPrivate",
								"TermsAndConditionsKapitalPensionPrivate", "TermsAndConditionsRatePensionPrivate"}
						case "tradingconditionfieldgroup []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
						case "inlinecountvalue":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"AllPages", "None"}
						case "activitytype []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"Email", "PhoneCall"}
						case "openorderduration":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{
								"AtTheClose", "AtTheOpening", "DayOrder", "FillOrKill", "GoodForPeriod",
								"GoodTillCancel", "GoodTillDate", "ImmediateOrCancel", "Unknown"}
						case "activityfieldgroup":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"DisplayAndFormat", "ExchangeInfo"}
						case "activityfieldgroup []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"DisplayAndFormat", "ExchangeInfo"}
						case "orderstatus []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"Filled", "LockedPlacementPending",
								"NotWorking", "NotWorkingLockedCancelPending", "NotWorkingLockedChangePending", "Parked",
								"Unknown", "Working", "WorkingLockedCancelPending", "WorkingLockedChangePending"}
						case "placeableordertype []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"Algorithmic",
								"GuaranteedStop", "Limit", "Market", "Stop", "StopIfTraded", "StopLimit",
								"Switch", "TrailingStop", "TrailingStopIfTraded", "Traspaso", "TraspasoIn"}
						case "positioneventfilter":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"All", "TradesOnly"}
						case "accountperformancefieldgroup []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"AccountSummary",
								"All", "Allocation", "AvailableBenchmarks", "BalancePerformance",
								"BalancePerformance_AccountValueTimeSeries", "BenchMark", "BenchmarkPerformance",
								"TimeWeightedPerformance", "TimeWeightedPerformance_AccumulatedTimeWeightedTimeSeries",
								"TotalCashBalancePerCurrency", "TotalPositionsValuePerCurrency",
								"TotalPositionsValuePerProductPerSecurity", "TradeActivity"}
						case "accountperformancestandardperiod":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"AllTime", "Month", "Quarter", "Year"}
						case "optionspacesegment":
							p.Schema.Type = "string"
						case "tradingstatus":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"NonTradable", "NotDefined", "ReduceOnly", "Tradable"}
						case "balancefieldgroup []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"MarginOverview"}
						case "closedpositionfieldgroup []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"ClosedPosition", "DisplayAndFormat", "ExchangeInfo"}
						case "putcall":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"Call", "None", "Put"}
						case "netpositionfieldgroup []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"DisplayAndFormat",
								"ExchangeInfo", "Greeks", "NetPositionBase",
								"NetPositionView", "SinglePosition", "SinglePositionBase", "SinglePositionView"}
						case "orderfieldgroup []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"DisplayAndFormat", "ExchangeInfo", "Greeks"}
						case "orderstatusfilter":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"All", "Working"}
						case "positionfieldgroup []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"DisplayAndFormat", "ExchangeInfo",
								"Greeks", "PositionBase", "PositionIdOnly", "PositionView"}
						case "class []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"Complex", "NonComplex"}
						case "instrumentfieldgroup []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"OrderSetting", "TradingSessions"}
						case "allocationkeystatus":
							p.Schema.Type = "string"
						case "orderamounttype":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"CashAmount", "Quantity"}
						case "infopricegroupspec []":
							p.Schema.Type = "array"
							p.Schema.Items.Type = "string"
							p.Schema.Items.Enum = []string{"Commissions", "DisplayAndFormat", "HistoricalChanges",
								"InstrumentPriceDetails", "MarketDepth", "PriceInfo", "PriceInfoDetails", "Quote"}
						case "toopenclose":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"ToClose", "ToOpen", "Undefined"}
						case "optionsstrategytype":
							p.Schema.Items.Enum = []string{"BackRatio",
								"Butterfly", "CalendarSpread", "Condor", "Custom", "Diagonal",
								"IronButterfly", "IronCondor", "RiskReversal", "Straddle", "Strangle", "Vertical"}
						case "pricealertdefinitionstate":
							p.Schema.Type = "string"
							p.Schema.Items.Enum = []string{"Disabled", "Enabled", "RecentlyTriggered"}
						default:
							p.Schema.Type = dataType
						}
						p.Name = param["Name"].(string)
						p.Required = param["IsRequired"].(bool)

						m.Parameters = append(m.Parameters, p)
					}
				}
			}
		}
	}

	d, err := yaml.Marshal(&paths)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(`
openapi: 3.0.1
info:
	title: Swagger Petstore
	description: xxxx
	version: 1.0.0
servers:
  - url: 'https://gateway.saxobank.com/sim/openapi'
	- url: 'https://gateway.saxobank.com/openapi'`)

	fmt.Println(string(d))
}
