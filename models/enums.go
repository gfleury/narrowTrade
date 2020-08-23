package models

type DurationType string

const (
	AtTheClose        DurationType = "AtTheClose"
	AtTheOpening                   = "AtTheOpening"
	DayOrder                       = "DayOrder"
	FillOrKill                     = "FillOrKill"
	GoodForPeriod                  = "GoodForPeriod"
	GoodTillCancel                 = "GoodTillCancel"
	GoodTillDate                   = "GoodTillDate"
	ImmediateOrCancel              = "ImmediateOrCancel"
	Unknown                        = "Unknown"
)

type BuySell string

const (
	Buy  BuySell = "Buy"
	Sell         = "Sell"
)

type OrderType string

const (
	Algorithmic           OrderType = "Algorithmic"
	GuaranteedStop                  = "GuaranteedStop"
	Limit                           = "Limit"
	Market                          = "Market"
	Stop                            = "Stop"
	StopIfBid                       = "StopIfBid"
	StopIfOffered                   = "StopIfOffered"
	StopIfTraded                    = "StopIfTraded"
	StopLimit                       = "StopLimit"
	Switch                          = "Switch"
	TrailingStop                    = "TrailingStop"
	TrailingStopIfBid               = "TrailingStopIfBid"
	TrailingStopIfOffered           = "TrailingStopIfOffered"
	TrailingStopIfTraded            = "TrailingStopIfTraded"
	Traspaso                        = "Traspaso"
	TraspasoIn                      = "TraspasoIn"
)

type AssetType string

const (
	Bond             AssetType = "Bond"
	Cash                       = "Cash"
	CfdIndexOption             = "CfdIndexOption"
	CfdOnFutures               = "CfdOnFutures"
	CfdOnIndex                 = "CfdOnIndex"
	CfdOnStock                 = "CfdOnStock"
	ContractFutures            = "ContractFutures"
	FuturesOption              = "FuturesOption"
	FuturesStrategy            = "FuturesStrategy"
	FxBinaryOption             = "FxBinaryOption"
	FxForwards                 = "FxForwards"
	FxKnockInOption            = "FxKnockInOption"
	FxKnockOutOption           = "FxKnockOutOption"
	FxNoTouchOption            = "FxNoTouchOption"
	FxOneTouchOption           = "FxOneTouchOption"
	FxSpot                     = "FxSpot"
	FxVanillaOption            = "FxVanillaOption"
	IpoOnStock                 = "IpoOnStock"
	ManagedFund                = "ManagedFund"
	MutualFund                 = "MutualFund"
	Stock                      = "Stock"
	StockIndex                 = "StockIndex"
	StockIndexOption           = "StockIndexOption"
	StockOption                = "StockOption"
)
