package main

func DataTypeMap(dataType string, s *Schema) {
	switch dataType {
	case "bool":
		s.Type = "boolean"
	case "double":
		s.Type = "number"
	case "decimal":
		s.Type = "number"
	case "uic":
		s.Type = "integer"
	case "int":
		s.Type = "integer"
	case "int []":
		s.Type = "array"
		s.Items.Type = "integer"
	case "string []":
		s.Type = "array"
		s.Items.Type = "string"
	case "clientkey":
		s.Type = "string"
	case "guid":
		s.Type = "string"
	case "accountkey":
		s.Type = "string"
	case "userkey":
		s.Type = "string"
	case "accountgroupkey":
		s.Type = "string"
	case "datetime":
		s.Type = "string"
	case "datetime []":
		s.Type = "array"
		s.Items.Type = "string"
	case "utcdatetime":
		s.Type = "string"
	case "date":
		s.Type = "string"
	case "standardperiodtype":
		s.Type = "string"
		s.Items.Enum = []string{"AllTime", "Month", "Quarter", "Year"}
	case "onboardingdocumenttype":
		s.Type = "string"
		s.Items.Enum = []string{"Switzerland501", "Switzerland901"}
	case "bookingrequesttype":
		s.Type = "string"
		s.Items.Enum = []string{"BkAmountId", "CaMasterRecordId", "RelatedTradeId"}
	case "orderentrytype":
		s.Type = "string"
		s.Items.Enum = []string{"All", "Last"}
	case "scheduledtradingconditionsfieldgroup []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"ScheduledTradingConditions"}
	case "casestatus []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"Canceled", "InformationProvided", "ProblemSolved"}
	case "orderlogstatus []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"Cancelled", "Changed", "DoneForDay",
			"Expired", "Fill", "FinalFill", "Placed", "Working"}
	case "logentrytype []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"Corrections", "DisclaimerActions", "NonTradingActivity",
			"Notifications", "OrderActivity", "Orders", "TradeActivity", "Trades"}
	case "assettype":
		s.Type = "string"
		s.Items.Enum = []string{"Bond", "Cash", "CfdIndexOption",
			"CfdOnFutures", "CfdOnIndex", "CfdOnStock", "ContractFutures",
			"FuturesOption", "FuturesStrategy", "FxBinaryOption", "FxForwards",
			"FxKnockInOption", "FxKnockOutOption", "FxNoTouchOption", "FxOneTouchOption",
			"FxSpot", "FxVanillaOption", "IpoOnStock", "ManagedFund", "MutualFund",
			"Stock", "StockIndex", "StockIndexOption", "StockOption"}
	case "assettype []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"Bond", "Cash", "CfdIndexOption",
			"CfdOnFutures", "CfdOnIndex", "CfdOnStock", "ContractFutures",
			"FuturesOption", "FuturesStrategy", "FxBinaryOption", "FxForwards",
			"FxKnockInOption", "FxKnockOutOption", "FxNoTouchOption", "FxOneTouchOption",
			"FxSpot", "FxVanillaOption", "IpoOnStock", "ManagedFund", "MutualFund",
			"Stock", "StockIndex", "StockIndexOption", "StockOption"}

	case "chartfieldgroupspec []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"BlockInfo", "ChartInfo", "Data", "DisplayAndFormat"}
	case "chartrequestmode":
		s.Type = "string"
		s.Items.Enum = []string{"From", "UpTo"}
	case "orderactivityfieldgroup []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"DisplayAndFormat"}
	case "signupflowdocumenttype":
		s.Type = "string"
		s.Items.Enum = []string{"AccountViewToIb",
			"EsaContract", "FeePaymentAuthorization", "GeneralBusinessTerms",
			"PensionTransferRequest", "PowerOfAttorney", "PowerOfAttorneyToIb",
			"ProofOfIdentity", "ProofOfResidency", "SourceOfFundsDocument",
			"TermsAndConditions", "TermsAndConditionsAldersopsparingPrivate",
			"TermsAndConditionsKapitalPensionPrivate", "TermsAndConditionsRatePensionPrivate"}
	case "tradingconditionfieldgroup []":
		s.Type = "array"
		s.Items.Type = "string"
	case "inlinecountvalue":
		s.Type = "string"
		s.Items.Enum = []string{"AllPages", "None"}
	case "activitytype []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"Email", "PhoneCall"}
	case "openorderduration":
		s.Type = "string"
		s.Items.Enum = []string{
			"AtTheClose", "AtTheOpening", "DayOrder", "FillOrKill", "GoodForPeriod",
			"GoodTillCancel", "GoodTillDate", "ImmediateOrCancel", "Unknown"}
	case "activityfieldgroup":
		s.Type = "string"
		s.Items.Enum = []string{"DisplayAndFormat", "ExchangeInfo"}
	case "activityfieldgroup []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"DisplayAndFormat", "ExchangeInfo"}
	case "orderstatus []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"Filled", "LockedPlacementPending",
			"NotWorking", "NotWorkingLockedCancelPending", "NotWorkingLockedChangePending", "Parked",
			"Unknown", "Working", "WorkingLockedCancelPending", "WorkingLockedChangePending"}
	case "placeableordertype []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"Algorithmic",
			"GuaranteedStop", "Limit", "Market", "Stop", "StopIfTraded", "StopLimit",
			"Switch", "TrailingStop", "TrailingStopIfTraded", "Traspaso", "TraspasoIn"}
	case "positioneventfilter":
		s.Type = "string"
		s.Items.Enum = []string{"All", "TradesOnly"}
	case "accountperformancefieldgroup []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"AccountSummary",
			"All", "Allocation", "AvailableBenchmarks", "BalancePerformance",
			"BalancePerformance_AccountValueTimeSeries", "BenchMark", "BenchmarkPerformance",
			"TimeWeightedPerformance", "TimeWeightedPerformance_AccumulatedTimeWeightedTimeSeries",
			"TotalCashBalancePerCurrency", "TotalPositionsValuePerCurrency",
			"TotalPositionsValuePerProductPerSecurity", "TradeActivity"}
	case "accountperformancestandardperiod":
		s.Type = "string"
		s.Items.Enum = []string{"AllTime", "Month", "Quarter", "Year"}
	case "optionspacesegment":
		s.Type = "string"
	case "tradingstatus":
		s.Type = "string"
		s.Items.Enum = []string{"NonTradable", "NotDefined", "ReduceOnly", "Tradable"}
	case "balancefieldgroup []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"MarginOverview"}
	case "closedpositionfieldgroup []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"ClosedPosition", "DisplayAndFormat", "ExchangeInfo"}
	case "putcall":
		s.Type = "string"
		s.Items.Enum = []string{"Call", "None", "Put"}
	case "netpositionfieldgroup []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"DisplayAndFormat",
			"ExchangeInfo", "Greeks", "NetPositionBase",
			"NetPositionView", "SinglePosition", "SinglePositionBase", "SinglePositionView"}
	case "orderfieldgroup []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"DisplayAndFormat", "ExchangeInfo", "Greeks"}
	case "orderstatusfilter":
		s.Type = "string"
		s.Items.Enum = []string{"All", "Working"}
	case "positionfieldgroup []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"DisplayAndFormat", "ExchangeInfo",
			"Greeks", "PositionBase", "PositionIdOnly", "PositionView"}
	case "class []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"Complex", "NonComplex"}
	case "instrumentfieldgroup []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"OrderSetting", "TradingSessions"}
	case "allocationkeystatus":
		s.Type = "string"
	case "orderamounttype":
		s.Type = "string"
		s.Items.Enum = []string{"CashAmount", "Quantity"}
	case "infopricegroupspec []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"Commissions", "DisplayAndFormat", "HistoricalChanges",
			"InstrumentPriceDetails", "MarketDepth", "PriceInfo", "PriceInfoDetails", "Quote"}
	case "toopenclose":
		s.Type = "string"
		s.Items.Enum = []string{"ToClose", "ToOpen", "Undefined"}
	case "optionsstrategytype":
		s.Type = "string"
		s.Items.Enum = []string{"BackRatio",
			"Butterfly", "CalendarSpread", "Condor", "Custom", "Diagonal",
			"IronButterfly", "IronCondor", "RiskReversal", "Straddle", "Strangle", "Vertical"}
	case "pricealertdefinitionstate":
		s.Type = "string"
		s.Items.Enum = []string{"Disabled", "Enabled", "RecentlyTriggered"}

	// New
	case "investmentssubscriptionarguments":
		s.Type = "object"
	case "suitabilitylevel":
		s.Type = "string"
		s.Items.Enum = []string{"High", "Low", "Medium", "Unknown", "VeryHigh"}
	case "productarea []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"Bonds", "CfdCommodities", "CFDs",
			"ContractOptions", "Futures", "FXForwards", "FXOptions", "FXSpot", "MutualFunds", "Stocks"}
	case "fundingcheck":
		s.Type = "string"
		s.Items.Enum = []string{"Enforce", "Ignore"}
	case "chartsubscriptionrequest":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"AssetType": Schema{
				Type: "string",
				Items: struct {
					Type    string   `json:"type" yaml:",omitempty"`
					Default string   `json:"default" yaml:",omitempty"`
					Enum    []string `json:"enum" yaml:",omitempty"`
				}{"", "", []string{"Bond", "Cash", "CfdIndexOption",
					"CfdOnFutures", "CfdOnIndex", "CfdOnStock", "ContractFutures", "FuturesOption",
					"FuturesStrategy", "FxBinaryOption", "FxForwards", "FxKnockInOption", "FxKnockOutOption",
					"FxNoTouchOption", "FxOneTouchOption", "FxSpot", "FxVanillaOption", "IpoOnStock",
					"ManagedFund", "MutualFund", "Stock", "StockIndex", "StockIndexOption", "StockOption"},
				},
			},
			"Count": Schema{Type: "integer"},
			"FieldGroups": Schema{
				Type: "array",
				Items: struct {
					Type    string   `json:"type" yaml:",omitempty"`
					Default string   `json:"default" yaml:",omitempty"`
					Enum    []string `json:"enum" yaml:",omitempty"`
				}{"object", "", nil},
			},
			"Horizon": Schema{Type: "integer"},
			"Uic":     Schema{Type: "integer"},
		}
	case "renewaldocument []":
		s.Type = "array"
		s.Items.Type = "string"
	case "renewaldata":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"AustraliaData": Schema{
				Type:       "object",
				Properties: map[string]Schema{},
			},
			"PersonalInformation": Schema{
				Type:       "object",
				Properties: map[string]Schema{},
			},
			"ProfileInformation": Schema{
				Type:       "object",
				Properties: map[string]Schema{},
			},
			"RegulatoryInformation": Schema{
				Type:       "object",
				Properties: map[string]Schema{},
			},
			"SingaporeData": Schema{
				Type:       "object",
				Properties: map[string]Schema{},
			},
		}
	case "accountinformation":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"AdditionalChoiceOfAccounts": Schema{
				Type: "array",
				Items: struct {
					Type    string   `json:"type" yaml:",omitempty"`
					Default string   `json:"default" yaml:",omitempty"`
					Enum    []string `json:"enum" yaml:",omitempty"`
				}{"string", "", nil},
			},
			"ClientCategoryId":          Schema{Type: "string"},
			"CurrencyCode":              Schema{Type: "string"},
			"IntendedCommissionGroupId": Schema{Type: "string"},
			"IntendedTemplateId":        Schema{Type: "string"},
			"OtherInstructions":         Schema{Type: "string"},
		}
	case "bankinformation":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"BeneficiaryBankName":     Schema{Type: "string"},
			"ReceivingCountryIsoCode": Schema{Type: "string"},
		}
	case "finlanddata":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"EuroclearJudicialForm": Schema{Type: "string"},
			"EuroclearSectorCode":   Schema{Type: "string"},
		}
	case "italydata":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"EmploymentInformation": Schema{Type: "object"},
			"MailContactPreference": Schema{Type: "boolean"},
			"ProfileInformation":    Schema{Type: "object"},
		}
	case "onboardinginformation":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"ElectronicVerification": Schema{Type: "object"},
		}
	case "pensiondata":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"AdditionalPensionData": Schema{Type: "string"},
			"EmployerReference":     Schema{Type: "string"},
			"PensionProductTypes": Schema{
				Type: "array",
				Items: struct {
					Type    string   `json:"type" yaml:",omitempty"`
					Default string   `json:"default" yaml:",omitempty"`
					Enum    []string `json:"enum" yaml:",omitempty"`
				}{"string", "", nil},
			},
		}
	case "personalinformation":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"AdditionalTaxCountryCode": Schema{Type: "string"},
		}
	case "profileinformation":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"AnnualIncomeInformation": Schema{Type: "object"},
			"InvestableAssets":        Schema{Type: "object"},
			"InvestmentPurpose":       Schema{Type: "object"},
		}
	case "regulatoryinformation":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"FatcaDeclaration": Schema{Type: "object"},
		}
	case "singaporedata":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"AnnualIncomeSgd": Schema{Type: "string"},
		}
	case "switzerlanddata":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"AnnualIncomeChf": Schema{Type: "string"},
			"NetWorthChf":     Schema{Type: "string"},
		}
	case "userid":
		s.Type = "string"
	case "signupdocument []":
		s.Type = "array"
		s.Items.Type = "object"
	case "updatecasestatus":
		s.Type = "string"
		s.Items.Enum = []string{"ExternallyPending", "InProgress", "InternallyPending"}
	case "casetype":
		s.Type = "string"
		s.Items.Enum = []string{"Faq", "Problem", "Question", "Request"}
	case "closecasestatus":
		s.Type = "string"
		s.Items.Enum = []string{"Canceled", "InformationProvided", "ProblemSolved"}
	case "file":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"FileName": Schema{Type: "string"},
			"MimeType": Schema{Type: "string"},
			"Data":     Schema{Type: "string"},
		}
	case "advancedoptions":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"IntermediaryBank":        Schema{Type: "string"},
			"RegulatedBroker":         Schema{Type: "object"},
			"RequestClosureOfAccount": Schema{Type: "string"},
		}
	case "searchcriteriafieldgroups []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"Accounts", "Default", "Users"}
	case "subscriptionactivityrequest":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "formdatacollection":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "accountbenchmarkinstrument":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"BenchmarkInstrumentAssetType": Schema{Type: "string"},
			"BenchmarkInstrumentUic":       Schema{Type: "integer"},
		}
	case "accountsrequest":
		s.Type = "object"
		s.Properties = map[string]Schema{
			"AccountGroupKey": Schema{Type: "string"},
			"AccountKey":      Schema{Type: "string"},
			"ClientKey":       Schema{Type: "string"},
		}
	case "balancerequest":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "clientpositionnettingmode":
		s.Type = "string"
		s.Items.Enum = []string{"EndOfDay", "Intraday"}
	case "closedpositionrequest":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "instrumentexposurerequest":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "netpositionrequest":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "openordersrequest":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "positionrequest":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "tradelevel":
		s.Type = "string"
		s.Items.Enum = []string{"FullTradingAndChat", "OrdersOnly"}
	case "allocationunittype":
		s.Type = "string"
		s.Items.Enum = []string{"Percentage", "Unit"}
	case "marginhandlingtype":
		s.Type = "string"
		s.Items.Enum = []string{"Disabled", "Enabled", "RecentlyTriggered"}
	case "allocationkeyaccountinfo []":
		s.Type = "array"
		s.Items.Type = "object"
	case "infopricelistrequest":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "optionschainrequest":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "requestexpiryspecification []":
		s.Type = "array"
		s.Items.Type = "object"
	case "buysell":
		s.Type = "string"
		s.Items.Enum = []string{"buy", "sell"}
	case "exercisemethod":
		s.Type = "string"
		s.Items.Enum = []string{"Cash", "Spot"}
	case "multilegpricefieldgroups []":
		s.Type = "array"
		s.Items.Type = "string"
		s.Items.Enum = []string{"Commissions", "DisplayAndFormat", "Greeks", "InstrumentPriceDetails",
			"MarginImpactBuySell", "Quote"}
	case "multilegstrategylegrequest []":
		s.Type = "array"
		s.Items.Type = "object"
	case "multilegpricerequest":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "pricerequest":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "algorithmicorderdata":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "orderduration":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "placeableordertype":
		s.Type = "string"
		s.Items.Enum = []string{"Algorithmic", "GuaranteedStop", "Limit", "Market", "Stop", "StopIfTraded",
			"StopLimit", "Switch", "TrailingStop", "TrailingStopIfTraded", "Traspaso", "TraspasoIn"}
	case "relatedorocoorder []":
		s.Type = "array"
		s.Items.Type = "object"
	case "placerelatedorocoorder []":
		s.Type = "array"
		s.Items.Type = "object"
	case "traspasoindetails":
		s.Type = "object"
		s.Properties = map[string]Schema{}
	case "orderleg []":
		s.Type = "array"
		s.Items.Type = "object"
	case "precheckorderspec []":
		s.Type = "array"
		s.Items.Type = "object"
	case "alertdefinitionassettype":
		s.Type = "string"
		s.Items.Enum = []string{"CfdOnFutures", "CfdOnIndex", "CfdOnStock",
			"ContractFutures", "FxSpot", "SrdOnEtf", "SrdOnStock", "Stock", "StockIndex"}
	case "pricealertcomparisonoperator":
		s.Type = "string"
		s.Items.Enum = []string{"GreaterOrEqual", "LessOrEqual"}
	case "pricevariable":
		s.Type = "string"
		s.Items.Enum = []string{"AskTick", "BidTick", "PercentChange", "Traded"}
	case "alertsound":
		s.Type = "string"
		s.Items.Enum = []string{"Asterisk", "Beep", "Exclamation", "Hand", "None", "Question"}
	default:
		s.Type = dataType
	}
}
