# {{classname}}

All URIs are relative to *https://gateway.saxobank.com/sim*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSubscriptionAsyncEns**](EnsApi.md#AddSubscriptionAsyncEns) | **Post** /openapi/ens/v1/activities/subscriptions | 
[**DeleteSubscriptionEns3**](EnsApi.md#DeleteSubscriptionEns3) | **Delete** /openapi/ens/v1/activities/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteSubscriptionsEns**](EnsApi.md#DeleteSubscriptionsEns) | **Delete** /openapi/ens/v1/activities/subscriptions/{ContextId}/ | 
[**GetActivitiesAsync**](EnsApi.md#GetActivitiesAsync) | **Get** /openapi/ens/v1/activities/ | 

# **AddSubscriptionAsyncEns**
> interface{} AddSubscriptionAsyncEns(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnsApiAddSubscriptionAsyncEnsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnsApiAddSubscriptionAsyncEnsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body34**](Body34.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscriptionEns3**
> interface{} DeleteSubscriptionEns3(ctx, contextId, referenceId)


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

# **DeleteSubscriptionsEns**
> interface{} DeleteSubscriptionsEns(ctx, contextId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
 **optional** | ***EnsApiDeleteSubscriptionsEnsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnsApiDeleteSubscriptionsEnsOpts struct
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

# **GetActivitiesAsync**
> interface{} GetActivitiesAsync(ctx, activities, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activities** | [**[]string**](string.md)|  | 
 **optional** | ***EnsApiGetActivitiesAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnsApiGetActivitiesAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skiptoken** | **optional.String**|  | 
 **top** | **optional.Int32**|  | 
 **accountKey** | **optional.String**|  | 
 **clientKey** | **optional.String**|  | 
 **duration** | [**optional.Interface of []string**](string.md)|  | 
 **expirationDateTime** | **optional.String**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 
 **fromDateTime** | **optional.String**|  | 
 **includeSubAccounts** | **optional.Bool**|  | 
 **orderStatuses** | [**optional.Interface of []string**](string.md)|  | 
 **orderTypes** | [**optional.Interface of []string**](string.md)|  | 
 **positionEventFilter** | [**optional.Interface of []string**](string.md)|  | 
 **sequenceId** | **optional.String**|  | 
 **toDateTime** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

