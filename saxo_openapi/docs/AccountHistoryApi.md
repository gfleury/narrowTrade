# {{classname}}

All URIs are relative to *https://gateway.saxobank.com/sim*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoricalPositions**](AccountHistoryApi.md#GetHistoricalPositions) | **Get** /openapi/hist/v3/positions/{ClientKey}/ | 
[**GetPerformance**](AccountHistoryApi.md#GetPerformance) | **Get** /openapi/hist/v3/perf/{ClientKey}/ | 
[**GetStandardPeriodAccountValues**](AccountHistoryApi.md#GetStandardPeriodAccountValues) | **Get** /openapi/hist/v3/accountValues/{ClientKey}/ | 

# **GetHistoricalPositions**
> interface{} GetHistoricalPositions(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***AccountHistoryApiGetHistoricalPositionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountHistoryApiGetHistoricalPositionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlinecount** | [**optional.Interface of []string**](string.md)|  | 
 **skip** | **optional.Int32**|  | 
 **skiptoken** | **optional.String**|  | 
 **top** | **optional.Int32**|  | 
 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **assetType** | [**optional.Interface of []string**](string.md)|  | 
 **fromDate** | **optional.String**|  | 
 **mockDataId** | **optional.String**|  | 
 **standardPeriod** | [**optional.Interface of []string**](string.md)|  | 
 **symbol** | [**optional.Interface of []string**](string.md)|  | 
 **toDate** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPerformance**
> interface{} GetPerformance(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***AccountHistoryApiGetPerformanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountHistoryApiGetPerformanceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 
 **fromDate** | **optional.String**|  | 
 **mockDataId** | **optional.String**|  | 
 **standardPeriod** | [**optional.Interface of []string**](string.md)|  | 
 **toDate** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStandardPeriodAccountValues**
> interface{} GetStandardPeriodAccountValues(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***AccountHistoryApiGetStandardPeriodAccountValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountHistoryApiGetStandardPeriodAccountValuesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mockDataId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

