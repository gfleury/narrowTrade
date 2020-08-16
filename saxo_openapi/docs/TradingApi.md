# {{classname}}

All URIs are relative to *https://gateway.saxobank.com/sim*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAllocationKey**](TradingApi.md#AddAllocationKey) | **Post** /openapi/trade/v1/allocationkeys | 
[**AddMultiLegPricesSubscriptionAsync**](TradingApi.md#AddMultiLegPricesSubscriptionAsync) | **Post** /openapi/trade/v1/prices/multileg/subscriptions | 
[**AddSubscriptionAsyncTrading3**](TradingApi.md#AddSubscriptionAsyncTrading3) | **Post** /openapi/trade/v1/infoprices/subscriptions | 
[**AddSubscriptionAsyncTrading4**](TradingApi.md#AddSubscriptionAsyncTrading4) | **Post** /openapi/trade/v1/messages/subscriptions | 
[**AddSubscriptionAsyncTrading5**](TradingApi.md#AddSubscriptionAsyncTrading5) | **Post** /openapi/trade/v1/optionschain/subscriptions | 
[**AddSubscriptionAsyncTrading6**](TradingApi.md#AddSubscriptionAsyncTrading6) | **Post** /openapi/trade/v1/prices/subscriptions | 
[**CancelAllOrder**](TradingApi.md#CancelAllOrder) | **Delete** /openapi/trade/v2/orders/ | 
[**CancelMultiLegStrategyOrder**](TradingApi.md#CancelMultiLegStrategyOrder) | **Delete** /openapi/trade/v2/orders/multileg/{MultiLegOrderId}/ | 
[**CancelOrder**](TradingApi.md#CancelOrder) | **Delete** /openapi/trade/v2/orders/{OrderIds}/ | 
[**ChangeMultiLegStrategyOrder**](TradingApi.md#ChangeMultiLegStrategyOrder) | **Patch** /openapi/trade/v2/orders/multileg | 
[**ChangeOrder**](TradingApi.md#ChangeOrder) | **Patch** /openapi/trade/v2/orders | 
[**CreatePosition**](TradingApi.md#CreatePosition) | **Post** /openapi/trade/v1/positions | 
[**DeleteAllocationKey**](TradingApi.md#DeleteAllocationKey) | **Delete** /openapi/trade/v1/allocationkeys/{AllocationKeyId} | 
[**DeleteMultiLegPricesSubscription**](TradingApi.md#DeleteMultiLegPricesSubscription) | **Delete** /openapi/trade/v1/prices/multileg/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteMultiLegPricesSubscriptions**](TradingApi.md#DeleteMultiLegPricesSubscriptions) | **Delete** /openapi/trade/v1/prices/multileg/subscriptions/{ContextId}/ | 
[**DeleteSubscriptionTrading12**](TradingApi.md#DeleteSubscriptionTrading12) | **Delete** /openapi/trade/v1/infoprices/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteSubscriptionTrading13**](TradingApi.md#DeleteSubscriptionTrading13) | **Delete** /openapi/trade/v1/messages/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteSubscriptionTrading14**](TradingApi.md#DeleteSubscriptionTrading14) | **Delete** /openapi/trade/v1/optionschain/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteSubscriptionTrading15**](TradingApi.md#DeleteSubscriptionTrading15) | **Delete** /openapi/trade/v1/prices/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteSubscriptionsTrading11**](TradingApi.md#DeleteSubscriptionsTrading11) | **Delete** /openapi/trade/v1/infoprices/subscriptions/{ContextId}/ | 
[**DeleteSubscriptionsTrading12**](TradingApi.md#DeleteSubscriptionsTrading12) | **Delete** /openapi/trade/v1/messages/subscriptions/{ContextId}/ | 
[**DeleteSubscriptionsTrading13**](TradingApi.md#DeleteSubscriptionsTrading13) | **Delete** /openapi/trade/v1/prices/subscriptions/{ContextId}/ | 
[**DistributeAmountAccordingToAllocationKey**](TradingApi.md#DistributeAmountAccordingToAllocationKey) | **Get** /openapi/trade/v1/allocationkeys/distributions/{AllocationKeyId}/ | 
[**ExerciseAmountAsync**](TradingApi.md#ExerciseAmountAsync) | **Put** /openapi/trade/v1/positions/exercise | 
[**ExercisePositionAsync**](TradingApi.md#ExercisePositionAsync) | **Put** /openapi/trade/v1/positions/{PositionId}/exercise | 
[**GetAllocationKey**](TradingApi.md#GetAllocationKey) | **Get** /openapi/trade/v1/allocationkeys/{AllocationKeyId} | 
[**GetAllocationKeys**](TradingApi.md#GetAllocationKeys) | **Get** /openapi/trade/v1/allocationkeys/ | 
[**GetInfoPriceAsync**](TradingApi.md#GetInfoPriceAsync) | **Get** /openapi/trade/v1/infoprices/ | 
[**GetInfoPriceListAsync**](TradingApi.md#GetInfoPriceListAsync) | **Get** /openapi/trade/v1/infoprices/list/ | 
[**GetMessagesAsync**](TradingApi.md#GetMessagesAsync) | **Get** /openapi/trade/v1/messages | 
[**GetMultiLegOrderStrategyDefaults**](TradingApi.md#GetMultiLegOrderStrategyDefaults) | **Get** /openapi/trade/v2/orders/multileg/defaults/ | 
[**GetMultiLegPriceAsync**](TradingApi.md#GetMultiLegPriceAsync) | **Post** /openapi/trade/v1/prices/multileg | 
[**MarkMessageAsSeen**](TradingApi.md#MarkMessageAsSeen) | **Put** /openapi/trade/v1/messages/seen/{MessageId} | 
[**ModifySubscription**](TradingApi.md#ModifySubscription) | **Patch** /openapi/trade/v1/optionschain/subscriptions/{ContextId}/{ReferenceId} | 
[**PlaceMultiLegStrategyOrder**](TradingApi.md#PlaceMultiLegStrategyOrder) | **Post** /openapi/trade/v2/orders/multileg | 
[**PlaceOrder**](TradingApi.md#PlaceOrder) | **Post** /openapi/trade/v2/orders | 
[**PreCheckMultilegOrder**](TradingApi.md#PreCheckMultilegOrder) | **Post** /openapi/trade/v2/orders/multileg/precheck | 
[**PreCheckOrder**](TradingApi.md#PreCheckOrder) | **Post** /openapi/trade/v2/orders/precheck | 
[**RequestMarginImpactOnSubscription**](TradingApi.md#RequestMarginImpactOnSubscription) | **Put** /openapi/trade/v1/prices/subscriptions/{ContextId}/{ReferenceId}/MarginImpact | 
[**ResetSubscriptionAtTheMoney**](TradingApi.md#ResetSubscriptionAtTheMoney) | **Put** /openapi/trade/v1/optionschain/subscriptions/{ContextId}/{ReferenceId}/ResetATM | 
[**UpdatePosition**](TradingApi.md#UpdatePosition) | **Patch** /openapi/trade/v1/positions/{PositionId} | 

# **AddAllocationKey**
> interface{} AddAllocationKey(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiAddAllocationKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiAddAllocationKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body56**](Body56.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddMultiLegPricesSubscriptionAsync**
> interface{} AddMultiLegPricesSubscriptionAsync(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiAddMultiLegPricesSubscriptionAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiAddMultiLegPricesSubscriptionAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body66**](Body66.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSubscriptionAsyncTrading3**
> interface{} AddSubscriptionAsyncTrading3(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiAddSubscriptionAsyncTrading3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiAddSubscriptionAsyncTrading3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body57**](Body57.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSubscriptionAsyncTrading4**
> interface{} AddSubscriptionAsyncTrading4(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiAddSubscriptionAsyncTrading4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiAddSubscriptionAsyncTrading4Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body58**](Body58.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSubscriptionAsyncTrading5**
> interface{} AddSubscriptionAsyncTrading5(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiAddSubscriptionAsyncTrading5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiAddSubscriptionAsyncTrading5Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body59**](Body59.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSubscriptionAsyncTrading6**
> interface{} AddSubscriptionAsyncTrading6(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiAddSubscriptionAsyncTrading6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiAddSubscriptionAsyncTrading6Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body67**](Body67.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelAllOrder**
> interface{} CancelAllOrder(ctx, accountKey, assetType, uic)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountKey** | **string**|  | 
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

# **CancelMultiLegStrategyOrder**
> interface{} CancelMultiLegStrategyOrder(ctx, accountKey, multiLegOrderId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountKey** | **string**|  | 
  **multiLegOrderId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelOrder**
> interface{} CancelOrder(ctx, accountKey, orderIds)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountKey** | **string**|  | 
  **orderIds** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeMultiLegStrategyOrder**
> interface{} ChangeMultiLegStrategyOrder(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiChangeMultiLegStrategyOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiChangeMultiLegStrategyOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body71**](Body71.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeOrder**
> interface{} ChangeOrder(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiChangeOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiChangeOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body69**](Body69.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePosition**
> interface{} CreatePosition(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiCreatePositionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiCreatePositionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body61**](Body61.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllocationKey**
> interface{} DeleteAllocationKey(ctx, allocationKeyId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **allocationKeyId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMultiLegPricesSubscription**
> interface{} DeleteMultiLegPricesSubscription(ctx, contextId, referenceId)


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

# **DeleteMultiLegPricesSubscriptions**
> interface{} DeleteMultiLegPricesSubscriptions(ctx, contextId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
 **optional** | ***TradingApiDeleteMultiLegPricesSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiDeleteMultiLegPricesSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscriptionTrading12**
> interface{} DeleteSubscriptionTrading12(ctx, contextId, referenceId)


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

# **DeleteSubscriptionTrading13**
> interface{} DeleteSubscriptionTrading13(ctx, contextId, referenceId)


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

# **DeleteSubscriptionTrading14**
> interface{} DeleteSubscriptionTrading14(ctx, contextId, referenceId)


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

# **DeleteSubscriptionTrading15**
> interface{} DeleteSubscriptionTrading15(ctx, contextId, referenceId)


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

# **DeleteSubscriptionsTrading11**
> interface{} DeleteSubscriptionsTrading11(ctx, contextId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
 **optional** | ***TradingApiDeleteSubscriptionsTrading11Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiDeleteSubscriptionsTrading11Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscriptionsTrading12**
> interface{} DeleteSubscriptionsTrading12(ctx, contextId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
 **optional** | ***TradingApiDeleteSubscriptionsTrading12Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiDeleteSubscriptionsTrading12Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscriptionsTrading13**
> interface{} DeleteSubscriptionsTrading13(ctx, contextId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
 **optional** | ***TradingApiDeleteSubscriptionsTrading13Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiDeleteSubscriptionsTrading13Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DistributeAmountAccordingToAllocationKey**
> interface{} DistributeAmountAccordingToAllocationKey(ctx, allocationKeyId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **allocationKeyId** | **string**|  | 
 **optional** | ***TradingApiDistributeAmountAccordingToAllocationKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiDistributeAmountAccordingToAllocationKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **totalamount** | **optional.Float64**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExerciseAmountAsync**
> interface{} ExerciseAmountAsync(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiExerciseAmountAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiExerciseAmountAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body64**](Body64.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExercisePositionAsync**
> interface{} ExercisePositionAsync(ctx, positionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **positionId** | **string**|  | 
 **optional** | ***TradingApiExercisePositionAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiExercisePositionAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body63**](Body63.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllocationKey**
> interface{} GetAllocationKey(ctx, allocationKeyId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **allocationKeyId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllocationKeys**
> interface{} GetAllocationKeys(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiGetAllocationKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiGetAllocationKeysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **accountKey** | **optional.String**|  | 
 **clientKey** | **optional.String**|  | 
 **statuses** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfoPriceAsync**
> interface{} GetInfoPriceAsync(ctx, assetType, uic, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assetType** | [**[]string**](string.md)|  | 
  **uic** | **int32**|  | 
 **optional** | ***TradingApiGetInfoPriceAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiGetInfoPriceAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountKey** | **optional.String**|  | 
 **amount** | **optional.Float64**|  | 
 **amountType** | [**optional.Interface of []string**](string.md)|  | 
 **expiryDate** | **optional.String**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 
 **forwardDate** | **optional.String**|  | 
 **lowerBarrier** | **optional.Float64**|  | 
 **orderAskPrice** | **optional.Float64**|  | 
 **orderBidPrice** | **optional.Float64**|  | 
 **putCall** | [**optional.Interface of []string**](string.md)|  | 
 **quoteCurrency** | **optional.Bool**|  | 
 **strikePrice** | **optional.Float64**|  | 
 **toOpenClose** | [**optional.Interface of []string**](string.md)|  | 
 **upperBarrier** | **optional.Float64**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfoPriceListAsync**
> interface{} GetInfoPriceListAsync(ctx, assetType, uics, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assetType** | [**[]string**](string.md)|  | 
  **uics** | **string**|  | 
 **optional** | ***TradingApiGetInfoPriceListAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiGetInfoPriceListAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountKey** | **optional.String**|  | 
 **amount** | **optional.Float64**|  | 
 **amountType** | [**optional.Interface of []string**](string.md)|  | 
 **expiryDate** | **optional.String**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 
 **forwardDate** | **optional.String**|  | 
 **forwardDateFarLeg** | **optional.String**|  | 
 **forwardDateNearLeg** | **optional.String**|  | 
 **lowerBarrier** | **optional.Float64**|  | 
 **orderAskPrice** | **optional.Float64**|  | 
 **orderBidPrice** | **optional.Float64**|  | 
 **putCall** | [**optional.Interface of []string**](string.md)|  | 
 **quoteCurrency** | **optional.Bool**|  | 
 **strikePrice** | **optional.Float64**|  | 
 **toOpenClose** | [**optional.Interface of []string**](string.md)|  | 
 **upperBarrier** | **optional.Float64**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagesAsync**
> interface{} GetMessagesAsync(ctx, )


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

# **GetMultiLegOrderStrategyDefaults**
> interface{} GetMultiLegOrderStrategyDefaults(ctx, accountKey, optionRootId, optionsStrategyType)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountKey** | **string**|  | 
  **optionRootId** | **int32**|  | 
  **optionsStrategyType** | [**[]string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMultiLegPriceAsync**
> interface{} GetMultiLegPriceAsync(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiGetMultiLegPriceAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiGetMultiLegPriceAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body65**](Body65.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarkMessageAsSeen**
> interface{} MarkMessageAsSeen(ctx, messageId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **messageId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifySubscription**
> interface{} ModifySubscription(ctx, contextId, referenceId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
  **referenceId** | **string**|  | 
 **optional** | ***TradingApiModifySubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiModifySubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Body60**](Body60.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlaceMultiLegStrategyOrder**
> interface{} PlaceMultiLegStrategyOrder(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiPlaceMultiLegStrategyOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiPlaceMultiLegStrategyOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body70**](Body70.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlaceOrder**
> interface{} PlaceOrder(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiPlaceOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiPlaceOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body68**](Body68.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PreCheckMultilegOrder**
> interface{} PreCheckMultilegOrder(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiPreCheckMultilegOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiPreCheckMultilegOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body72**](Body72.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PreCheckOrder**
> interface{} PreCheckOrder(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TradingApiPreCheckOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiPreCheckOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body73**](Body73.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestMarginImpactOnSubscription**
> interface{} RequestMarginImpactOnSubscription(ctx, contextId, referenceId)


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

# **ResetSubscriptionAtTheMoney**
> interface{} ResetSubscriptionAtTheMoney(ctx, contextId, referenceId)


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

# **UpdatePosition**
> interface{} UpdatePosition(ctx, positionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **positionId** | **string**|  | 
 **optional** | ***TradingApiUpdatePositionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TradingApiUpdatePositionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body62**](Body62.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

