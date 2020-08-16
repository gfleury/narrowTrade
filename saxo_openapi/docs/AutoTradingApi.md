# {{classname}}

All URIs are relative to *https://gateway.saxobank.com/sim*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSubscription**](AutoTradingApi.md#AddSubscription) | **Post** /openapi/at/v3/investments/subscriptions | 
[**AddSuggestionSubscriptionAsync**](AutoTradingApi.md#AddSuggestionSubscriptionAsync) | **Post** /openapi/at/v3/investments/subscriptions/suggestions | 
[**DeleteSubscription**](AutoTradingApi.md#DeleteSubscription) | **Delete** /openapi/at/v3/investments/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteSuggestionSubscription**](AutoTradingApi.md#DeleteSuggestionSubscription) | **Delete** /openapi/at/v3/investments/subscriptions/suggestions/{ContextId}/{ReferenceId} | 
[**GetAllTradeLeaders**](AutoTradingApi.md#GetAllTradeLeaders) | **Get** /openapi/at/v3/tradeLeaders/ | 
[**GetAllTradeLeadersPublic**](AutoTradingApi.md#GetAllTradeLeadersPublic) | **Post** /openapi/at/v3/tradeLeaders/public | 
[**GetInvestmentByLeader**](AutoTradingApi.md#GetInvestmentByLeader) | **Get** /openapi/at/v3/investments/ | 
[**GetInvestments**](AutoTradingApi.md#GetInvestments) | **Get** /openapi/at/v3/investments/me/ | 
[**GetInvestmentsummariesAsync**](AutoTradingApi.md#GetInvestmentsummariesAsync) | **Get** /openapi/at/v3/investments/summaries | 
[**GetMaxWithdrawalAmountAsync**](AutoTradingApi.md#GetMaxWithdrawalAmountAsync) | **Get** /openapi/at/v3/investments/{InvestmentId}/maxWithdrawalAmount/ | 
[**GetSuitabilityStatus**](AutoTradingApi.md#GetSuitabilityStatus) | **Get** /openapi/at/v3/tradeFollowers/{ClientKey}/suitabilitystatus | 
[**GetTermsAndConditions**](AutoTradingApi.md#GetTermsAndConditions) | **Get** /openapi/at/v3/tradeFollowers/{ClientKey}/termsandconditions | 
[**GetTradeFollowers**](AutoTradingApi.md#GetTradeFollowers) | **Get** /openapi/at/v3/tradeFollowers/me | 
[**GetTradeLeader**](AutoTradingApi.md#GetTradeLeader) | **Get** /openapi/at/v3/tradeLeaders/{TradeLeaderId}/ | 
[**GetTradeLeaderLogo**](AutoTradingApi.md#GetTradeLeaderLogo) | **Get** /openapi/at/v3/tradeLeaders/{TradeLeaderId}/logo | 
[**GetTradeLeaderLogoByLogoKey**](AutoTradingApi.md#GetTradeLeaderLogoByLogoKey) | **Get** /openapi/at/v3/tradeLeaders/logo/{LogoKey} | 
[**GetTradeLeaderPublic**](AutoTradingApi.md#GetTradeLeaderPublic) | **Post** /openapi/at/v3/tradeLeaders/public/{TradeLeaderId} | 
[**GetTradeLeadersCustomSelectionPublic**](AutoTradingApi.md#GetTradeLeadersCustomSelectionPublic) | **Get** /openapi/at/v3/tradeLeaders/public/customselection/{SelectionId}/ | 
[**PatchInvestmentAsync**](AutoTradingApi.md#PatchInvestmentAsync) | **Patch** /openapi/at/v3/investments/{InvestmentId} | 
[**PostInvestment**](AutoTradingApi.md#PostInvestment) | **Post** /openapi/at/v3/investments | 
[**PutClientTermsAndConditions**](AutoTradingApi.md#PutClientTermsAndConditions) | **Put** /openapi/at/v3/tradeFollowers/{ClientKey}/termsandconditions | 
[**PutInvestmentAsync**](AutoTradingApi.md#PutInvestmentAsync) | **Put** /openapi/at/v3/investments/{InvestmentId}/withdraw | 
[**PutSuitabilityStatus**](AutoTradingApi.md#PutSuitabilityStatus) | **Put** /openapi/at/v3/tradeFollowers/{ClientKey}/suitabilitystatus | 

# **AddSubscription**
> interface{} AddSubscription(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutoTradingApiAddSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoTradingApiAddSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body3**](Body3.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSuggestionSubscriptionAsync**
> interface{} AddSuggestionSubscriptionAsync(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutoTradingApiAddSuggestionSubscriptionAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoTradingApiAddSuggestionSubscriptionAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body4**](Body4.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscription**
> interface{} DeleteSubscription(ctx, contextId, referenceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
  **referenceId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSuggestionSubscription**
> interface{} DeleteSuggestionSubscription(ctx, contextId, referenceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
  **referenceId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTradeLeaders**
> interface{} GetAllTradeLeaders(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutoTradingApiGetAllTradeLeadersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoTradingApiGetAllTradeLeadersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **standardPeriod** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTradeLeadersPublic**
> interface{} GetAllTradeLeadersPublic(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutoTradingApiGetAllTradeLeadersPublicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoTradingApiGetAllTradeLeadersPublicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body7**](Body7.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvestmentByLeader**
> interface{} GetInvestmentByLeader(ctx, investmentId, tradeLeaderId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **investmentId** | **string**|  | 
  **tradeLeaderId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvestments**
> interface{} GetInvestments(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutoTradingApiGetInvestmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoTradingApiGetInvestmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tradeLeaderId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvestmentsummariesAsync**
> interface{} GetInvestmentsummariesAsync(ctx, )


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

# **GetMaxWithdrawalAmountAsync**
> interface{} GetMaxWithdrawalAmountAsync(ctx, investmentId, targetAccountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **investmentId** | **string**|  | 
  **targetAccountId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSuitabilityStatus**
> interface{} GetSuitabilityStatus(ctx, clientKey)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTermsAndConditions**
> interface{} GetTermsAndConditions(ctx, clientKey)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTradeFollowers**
> interface{} GetTradeFollowers(ctx, )


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

# **GetTradeLeader**
> interface{} GetTradeLeader(ctx, tradeLeaderId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tradeLeaderId** | **string**|  | 
 **optional** | ***AutoTradingApiGetTradeLeaderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoTradingApiGetTradeLeaderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **standardPeriod** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTradeLeaderLogo**
> interface{} GetTradeLeaderLogo(ctx, tradeLeaderId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tradeLeaderId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTradeLeaderLogoByLogoKey**
> interface{} GetTradeLeaderLogoByLogoKey(ctx, logoKey)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logoKey** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTradeLeaderPublic**
> interface{} GetTradeLeaderPublic(ctx, tradeLeaderId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tradeLeaderId** | **int32**|  | 
 **optional** | ***AutoTradingApiGetTradeLeaderPublicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoTradingApiGetTradeLeaderPublicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body8**](Body8.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTradeLeadersCustomSelectionPublic**
> interface{} GetTradeLeadersCustomSelectionPublic(ctx, selectionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **selectionId** | **string**|  | 
 **optional** | ***AutoTradingApiGetTradeLeadersCustomSelectionPublicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoTradingApiGetTradeLeadersCustomSelectionPublicOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **standardPeriod** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchInvestmentAsync**
> interface{} PatchInvestmentAsync(ctx, investmentId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **investmentId** | **string**|  | 
 **optional** | ***AutoTradingApiPatchInvestmentAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoTradingApiPatchInvestmentAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body1**](Body1.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostInvestment**
> interface{} PostInvestment(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AutoTradingApiPostInvestmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoTradingApiPostInvestmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body**](Body.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutClientTermsAndConditions**
> interface{} PutClientTermsAndConditions(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***AutoTradingApiPutClientTermsAndConditionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoTradingApiPutClientTermsAndConditionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body6**](Body6.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutInvestmentAsync**
> interface{} PutInvestmentAsync(ctx, investmentId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **investmentId** | **int32**|  | 
 **optional** | ***AutoTradingApiPutInvestmentAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoTradingApiPutInvestmentAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body2**](Body2.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutSuitabilityStatus**
> interface{} PutSuitabilityStatus(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***AutoTradingApiPutSuitabilityStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AutoTradingApiPutSuitabilityStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body5**](Body5.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

