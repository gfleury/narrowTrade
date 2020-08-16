# {{classname}}

All URIs are relative to *https://gateway.saxobank.com/sim*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllCurrencyPairs**](ReferenceDataApi.md#GetAllCurrencyPairs) | **Get** /openapi/ref/v1/currencypairs/ | 
[**GetAllExchanges**](ReferenceDataApi.md#GetAllExchanges) | **Get** /openapi/ref/v1/exchanges/ | 
[**GetAllStrategies**](ReferenceDataApi.md#GetAllStrategies) | **Get** /openapi/ref/v1/algostrategies/ | 
[**GetContractOptionSpace**](ReferenceDataApi.md#GetContractOptionSpace) | **Get** /openapi/ref/v1/instruments/contractoptionspaces/{OptionRootId}/ | 
[**GetCultures**](ReferenceDataApi.md#GetCultures) | **Get** /openapi/ref/v1/cultures | 
[**GetCurrencies**](ReferenceDataApi.md#GetCurrencies) | **Get** /openapi/ref/v1/currencies | 
[**GetDetailsForManyInstruments**](ReferenceDataApi.md#GetDetailsForManyInstruments) | **Get** /openapi/ref/v1/instruments/details/ | 
[**GetDetailsForOneInstrument**](ReferenceDataApi.md#GetDetailsForOneInstrument) | **Get** /openapi/ref/v1/instruments/details/{Uic}/{AssetType}/ | 
[**GetExchange**](ReferenceDataApi.md#GetExchange) | **Get** /openapi/ref/v1/exchanges/{ExchangeId} | 
[**GetForwardTenorDates**](ReferenceDataApi.md#GetForwardTenorDates) | **Get** /openapi/ref/v1/standarddates/forwardtenor/{Uic}/ | 
[**GetFuturesSpace**](ReferenceDataApi.md#GetFuturesSpace) | **Get** /openapi/ref/v1/instruments/futuresspaces/{ContinuousFuturesUic} | 
[**GetFxOptionExpiryDates**](ReferenceDataApi.md#GetFxOptionExpiryDates) | **Get** /openapi/ref/v1/standarddates/fxoptionexpiry/{Uic} | 
[**GetReferenceData6**](ReferenceDataApi.md#GetReferenceData6) | **Get** /openapi/ref/v1/countries | 
[**GetReferenceData7**](ReferenceDataApi.md#GetReferenceData7) | **Get** /openapi/ref/v1/languages | 
[**GetStrategyByName**](ReferenceDataApi.md#GetStrategyByName) | **Get** /openapi/ref/v1/algostrategies/{Name} | 
[**GetSummaries**](ReferenceDataApi.md#GetSummaries) | **Get** /openapi/ref/v1/instruments/ | 
[**GetTimeZones**](ReferenceDataApi.md#GetTimeZones) | **Get** /openapi/ref/v1/timezones | 
[**GetTradingSchedule**](ReferenceDataApi.md#GetTradingSchedule) | **Get** /openapi/ref/v1/instruments/tradingschedule/{Uic}/{AssetType} | 

# **GetAllCurrencyPairs**
> interface{} GetAllCurrencyPairs(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceDataApiGetAllCurrencyPairsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceDataApiGetAllCurrencyPairsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountKey** | **optional.String**|  | 
 **clientKey** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllExchanges**
> interface{} GetAllExchanges(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceDataApiGetAllExchangesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceDataApiGetAllExchangesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllStrategies**
> interface{} GetAllStrategies(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceDataApiGetAllStrategiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceDataApiGetAllStrategiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContractOptionSpace**
> interface{} GetContractOptionSpace(ctx, optionRootId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **optionRootId** | **int32**|  | 
 **optional** | ***ReferenceDataApiGetContractOptionSpaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceDataApiGetContractOptionSpaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **canParticipateInMultiLegOrder** | **optional.Bool**|  | 
 **clientKey** | **optional.String**|  | 
 **expiryDates** | [**optional.Interface of []string**](string.md)|  | 
 **optionSpaceSegment** | **optional.String**|  | 
 **tradingStatus** | [**optional.Interface of []string**](string.md)|  | 
 **underlyingUic** | **optional.Int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCultures**
> interface{} GetCultures(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrencies**
> interface{} GetCurrencies(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailsForManyInstruments**
> interface{} GetDetailsForManyInstruments(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceDataApiGetDetailsForManyInstrumentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceDataApiGetDetailsForManyInstrumentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **accountKey** | **optional.String**|  | 
 **assetTypes** | [**optional.Interface of []string**](string.md)|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 
 **tags** | [**optional.Interface of []string**](string.md)|  | 
 **uics** | [**optional.Interface of []int32**](int32.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailsForOneInstrument**
> interface{} GetDetailsForOneInstrument(ctx, assetType, uic, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assetType** | [**[]string**](string.md)|  | 
  **uic** | **int32**|  | 
 **optional** | ***ReferenceDataApiGetDetailsForOneInstrumentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceDataApiGetDetailsForOneInstrumentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountKey** | **optional.String**|  | 
 **clientKey** | **optional.String**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExchange**
> interface{} GetExchange(ctx, exchangeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **exchangeId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForwardTenorDates**
> interface{} GetForwardTenorDates(ctx, uic, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uic** | **int32**|  | 
 **optional** | ***ReferenceDataApiGetForwardTenorDatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceDataApiGetForwardTenorDatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountKey** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFuturesSpace**
> interface{} GetFuturesSpace(ctx, continuousFuturesUic)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **continuousFuturesUic** | **int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFxOptionExpiryDates**
> interface{} GetFxOptionExpiryDates(ctx, uic)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uic** | **int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReferenceData6**
> interface{} GetReferenceData6(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReferenceData7**
> interface{} GetReferenceData7(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStrategyByName**
> interface{} GetStrategyByName(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSummaries**
> interface{} GetSummaries(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReferenceDataApiGetSummariesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReferenceDataApiGetSummariesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **accountKey** | **optional.String**|  | 
 **assetTypes** | [**optional.Interface of []string**](string.md)|  | 
 **canParticipateInMultiLegOrder** | **optional.Bool**|  | 
 **class** | [**optional.Interface of []string**](string.md)|  | 
 **exchangeId** | **optional.String**|  | 
 **includeNonTradable** | **optional.Bool**|  | 
 **keywords** | **optional.String**|  | 
 **sectorId** | **optional.String**|  | 
 **tags** | [**optional.Interface of []string**](string.md)|  | 
 **uics** | [**optional.Interface of []int32**](int32.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimeZones**
> interface{} GetTimeZones(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTradingSchedule**
> interface{} GetTradingSchedule(ctx, assetType, uic)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assetType** | [**[]string**](string.md)|  | 
  **uic** | **int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

