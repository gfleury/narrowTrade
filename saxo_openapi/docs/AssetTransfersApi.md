# {{classname}}

All URIs are relative to *https://gateway.saxobank.com/sim*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CashTransfer**](AssetTransfersApi.md#CashTransfer) | **Post** /openapi/atr/v1/partner/cashtransfers | 
[**CashTransfersById**](AssetTransfersApi.md#CashTransfersById) | **Get** /openapi/atr/v1/partner/cashtransfers/{TransactionId} | 
[**CashTransfersBySearchParameters**](AssetTransfersApi.md#CashTransfersBySearchParameters) | **Get** /openapi/atr/v1/partner/cashtransfers/ | 
[**PrebookFundDeposit**](AssetTransfersApi.md#PrebookFundDeposit) | **Post** /openapi/atr/v1/partner/prebookedfunds | 

# **CashTransfer**
> interface{} CashTransfer(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetTransfersApiCashTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetTransfersApiCashTransferOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body9**](Body9.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CashTransfersById**
> interface{} CashTransfersById(ctx, transactionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transactionId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CashTransfersBySearchParameters**
> interface{} CashTransfersBySearchParameters(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetTransfersApiCashTransfersBySearchParametersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetTransfersApiCashTransfersBySearchParametersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountKey** | **optional.String**|  | 
 **clientKey** | **optional.String**|  | 
 **externalReference** | **optional.String**|  | 
 **fromDateTime** | **optional.String**|  | 
 **skipToken** | **optional.String**|  | 
 **toDateTime** | **optional.String**|  | 
 **top** | **optional.Int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrebookFundDeposit**
> interface{} PrebookFundDeposit(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetTransfersApiPrebookFundDepositOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetTransfersApiPrebookFundDepositOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body10**](Body10.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

