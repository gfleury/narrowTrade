# {{classname}}

All URIs are relative to *https://gateway.saxobank.com/sim*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdNowInteractiveVerificationComplete**](PartnerIntegrationApi.md#IdNowInteractiveVerificationComplete) | **Post** /openapi/partnerintegration/v1/InteractiveIdVerification/idnow/complete | 
[**VerifyInteractiveVerification**](PartnerIntegrationApi.md#VerifyInteractiveVerification) | **Post** /openapi/partnerintegration/v1/InteractiveIdVerification/verify/{Correlationid} | 

# **IdNowInteractiveVerificationComplete**
> interface{} IdNowInteractiveVerificationComplete(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PartnerIntegrationApiIdNowInteractiveVerificationCompleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerIntegrationApiIdNowInteractiveVerificationCompleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body35**](Body35.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyInteractiveVerification**
> interface{} VerifyInteractiveVerification(ctx, correlationid, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **correlationid** | **string**|  | 
 **optional** | ***PartnerIntegrationApiVerifyInteractiveVerificationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerIntegrationApiVerifyInteractiveVerificationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body36**](Body36.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

