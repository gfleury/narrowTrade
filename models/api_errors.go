package models

import (
	"encoding/json"

	"github.com/gfleury/narrowTrade/saxo_openapi"
)

type OrderError struct {
	ErrorCode  string            `json:"ErrorCode"`
	Message    string            `json:"Message"`
	ModelState map[string]string `json:"ModelState"`
	ErrorInfo  ErrorInfo         `json:"ErrorInfo"`
	Orders     []struct {
		ErrorInfo ErrorInfo `json:"ErrorInfo"`
	} `json:"Orders"`
}

func GetStringError(err error) string {
	if errAPI, ok := err.(saxo_openapi.GenericSwaggerError); ok {
		return string(errAPI.Body())
	}
	return ""
}

func GetOrderError(err error) *OrderError {
	if errAPI, ok := err.(saxo_openapi.GenericSwaggerError); ok {
		orderError := OrderError{}
		errMarsh := json.Unmarshal(errAPI.Body(), &orderError)
		if errMarsh == nil {
			return &orderError
		}
	}
	return nil
}

func BusinessRuleViolation(oe *OrderError) bool {
	for _, violation := range GetBusinessViolations() {
		if oe.ErrorCode == violation {
			return true
		}
		for _, order := range oe.Orders {
			if order.ErrorInfo.ErrorCode == violation {
				return true
			}
		}
	}
	return false
}

func GetBusinessViolations() []string {
	return []string{"AccountIsTradeRestricted",
		"ActiveFollowerCannotCancelOrderManually",
		"AlgoOrderErrorInParameterValue",
		"AlgoOrderIllegalAssetType",
		"AlgoOrderIllegalDurationType",
		"AlgoOrderMissingMandatoryParameter",
		"AllocationKeyNotFound",
		"AllocationKeyPercentDoesNotSumToOneHundred",
		"AlreadyPendingCancelReplace",
		"AmountBelowMinimumLotSize",
		"AmountDiffersFromAmountOnRelatedOrder",
		"AmountLowerThanAllocationKeyUnits",
		"AmountNotInLotSize",
		"BarrierTooCloseToSpot",
		"BrokerOption",
		"ClientCannotCancelMarketExpiryOrder",
		"ClientExposureLimitation",
		"ClmDataFeedClientExceptionCommissionRuleMissing",
		"CloseAllPositionsAlreadyBeingClosed",
		"DirectAccessNotAllowed",
		"DurationNotSupported",
		"ExchangeNotSupported",
		"ExchangeRateNotAvailable",
		"ExpirationDateInPast",
		"ExpirationDateRequired",
		"ForwardDateInPast",
		"ForwardDateRequired",
		"FxSwapLegsValueDateRequired",
		"IllegalAccount",
		"IllegalAmount",
		"IllegalDate",
		"IllegalInstrumentId",
		"IllegalStrike",
		"IllegalWatchlistId",
		"InstrumentDisabledForTrading",
		"InstrumentHasExpired",
		"InstrumentNotAllowed",
		"InstrumentNotFoundOrNotAllowedForMarketTrades",
		"InstrumentNotTradableAsTheSelectedType",
		"InstrumentSuspended",
		"InstrumentTypeNotSupportedException",
		"InsufficentCash",
		"InsufficientTradeLevelException",
		"InternalServerError",
		"InvalidAllocationKeyUsed",
		"InvalidCurrencyPair",
		"InvalidExpiryTimeOnExchange",
		"InvalidFxSwapLegsValueDate",
		"InvalidMessageId",
		"InvalidOptionRootId",
		"InvalidPriceRequest",
		"InvalidRefreshRate",
		"InvalidRequest",
		"InvalidTimeInExpiationDateTime",
		"InvalidTraderId",
		"InvalidUic",
		"MarginBorderline",
		"MarketClosed",
		"MissingTimeInExpirationDateTime",
		"NoChatMessageEntered",
		"NoDataAccess",
		"NotTradableAtPresent",
		"NoValidQuote",
		"OnWrongSideOfMarket",
		"OpeningShortFXOptionPositionsNotAllowed",
		"OptionExerciseAfterCutoff",
		"OptionsChainNotSupportedForAssetType",
		"OrderCannotBeCancelledAtThisTime",
		"OrderNotFound",
		"OrderNotPlaced",
		"OrderRelatedPositionMissMatch",
		"OrderValueTooLarge",
		"OrderValueToSmall",
		"PendingTradeRequests",
		"PositionBuildupNotValidForInstrument",
		"PriceExceedsAggressiveTolerance",
		"PriceNotFound",
		"PriceNotInTickSizeIncrements",
		"PriceRequestRequiresExpiryDate",
		"PriceRequestRequiresPutCall",
		"QuoteHasTimedOut",
		"RelatedOrderWasRejected",
		"RequestForQuoteFailed",
		"RequestForQuoteNotAllowed",
		"RequestMarginImpactOnNextPriceFailed",
		"RequoteRequired",
		"ShortTradeDisabled",
		"TooCloseToEntryOrder",
		"TooCloseToMarket",
		"TooCloseToOcoRelatedOrderPrice",
		"TooFarFromEntryOrder",
		"TooFarFromMarket",
		"TooLateToCancelOrder",
		"TooManyStrikesRequested",
		"TooSmallTrade",
		"TraderIdIsRequired",
		"TradeSessionNotPrimary",
		"UicsInListNotUnique",
		"UnexpectedTimeInExpirationDateTime",
		"Unknown",
		"WarningAmountLargeOrderSize",
		"WouldExceedEquityConcentrationLimit",
		"WouldExceedMargin",
		"WouldExceedMarginCeiling",
		"WouldExceedTradingLine"}
}
