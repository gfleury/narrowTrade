# {{classname}}

All URIs are relative to *https://gateway.saxobank.com/sim*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAsync**](ClientReportingApi.md#GetAsync) | **Get** /openapi/cr/v1/reports/AccountStatement/{ClientKey}/ | 
[**GetAsyncClientReporting**](ClientReportingApi.md#GetAsyncClientReporting) | **Get** /openapi/cr/v1/reports/Portfolio/{ClientKey}/{FromDate}/{ToDate}/ | 
[**GetAsyncClientReporting3**](ClientReportingApi.md#GetAsyncClientReporting3) | **Get** /openapi/cr/v1/reports/TradeDetails/{ClientKey}/ | 
[**GetAsyncClientReporting4**](ClientReportingApi.md#GetAsyncClientReporting4) | **Get** /openapi/cr/v1/reports/TradesExecuted/{ClientKey}/ | 
[**GetMeAsync**](ClientReportingApi.md#GetMeAsync) | **Get** /openapi/cr/v1/reports/AccountStatement/me/ | 
[**GetMeAsyncClientReporting**](ClientReportingApi.md#GetMeAsyncClientReporting) | **Get** /openapi/cr/v1/reports/Portfolio/me/{FromDate}/{ToDate}/ | 
[**GetMeAsyncClientReporting3**](ClientReportingApi.md#GetMeAsyncClientReporting3) | **Get** /openapi/cr/v1/reports/TradeDetails/me/ | 
[**GetMeAsyncClientReporting4**](ClientReportingApi.md#GetMeAsyncClientReporting4) | **Get** /openapi/cr/v1/reports/TradesExecuted/me/ | 

# **GetAsync**
> interface{} GetAsync(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***ClientReportingApiGetAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientReportingApiGetAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **fromDate** | **optional.String**|  | 
 **toDate** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsyncClientReporting**
> interface{} GetAsyncClientReporting(ctx, clientKey, fromDate, toDate, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
  **fromDate** | **string**|  | 
  **toDate** | **string**|  | 
 **optional** | ***ClientReportingApiGetAsyncClientReportingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientReportingApiGetAsyncClientReportingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsyncClientReporting3**
> interface{} GetAsyncClientReporting3(ctx, accountKey, clientKey, filterType, filterValue, tradeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountKey** | **string**|  | 
  **clientKey** | **string**|  | 
  **filterType** | [**[]string**](string.md)|  | 
  **filterValue** | **string**|  | 
  **tradeId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsyncClientReporting4**
> interface{} GetAsyncClientReporting4(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***ClientReportingApiGetAsyncClientReporting4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientReportingApiGetAsyncClientReporting4Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **fromDate** | **optional.String**|  | 
 **toDate** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMeAsync**
> interface{} GetMeAsync(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientReportingApiGetMeAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientReportingApiGetMeAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **fromDate** | **optional.String**|  | 
 **toDate** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMeAsyncClientReporting**
> interface{} GetMeAsyncClientReporting(ctx, fromDate, toDate, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fromDate** | **string**|  | 
  **toDate** | **string**|  | 
 **optional** | ***ClientReportingApiGetMeAsyncClientReportingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientReportingApiGetMeAsyncClientReportingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMeAsyncClientReporting3**
> interface{} GetMeAsyncClientReporting3(ctx, accountKey, filterType, filterValue, tradeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountKey** | **string**|  | 
  **filterType** | [**[]string**](string.md)|  | 
  **filterValue** | **string**|  | 
  **tradeId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMeAsyncClientReporting4**
> interface{} GetMeAsyncClientReporting4(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientReportingApiGetMeAsyncClientReporting4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientReportingApiGetMeAsyncClientReporting4Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **fromDate** | **optional.String**|  | 
 **toDate** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

