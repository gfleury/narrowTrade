# {{classname}}

All URIs are relative to *https://gateway.saxobank.com/sim*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachFile**](ClientManagementApi.md#AttachFile) | **Post** /openapi/cm/v2/signups/attachments/{SignUpId} | 
[**AttachFileClientManagement**](ClientManagementApi.md#AttachFileClientManagement) | **Post** /openapi/cm/v1/signups/attachments/{SignUpId}/ | 
[**CompleteApplication**](ClientManagementApi.md#CompleteApplication) | **Put** /openapi/cm/v1/signups/completeapplication/{SignUpId}/ | 
[**GenerateTypedOnboardingPDF**](ClientManagementApi.md#GenerateTypedOnboardingPDF) | **Get** /openapi/cm/v1/signups/onboardingpdf/{ClientKey}/ | 
[**GetClientRenewalData**](ClientManagementApi.md#GetClientRenewalData) | **Get** /openapi/cm/v1/clientrenewals/ | 
[**GetRenewalStatuses**](ClientManagementApi.md#GetRenewalStatuses) | **Get** /openapi/cm/v1/clientrenewals/pending/ | 
[**GetSignUpStatus**](ClientManagementApi.md#GetSignUpStatus) | **Get** /openapi/cm/v1/signups/status/{ClientKey} | 
[**GetSignupOptions**](ClientManagementApi.md#GetSignupOptions) | **Get** /openapi/cm/v2/signups/options | 
[**GetSignupOptionsClientManagement**](ClientManagementApi.md#GetSignupOptionsClientManagement) | **Get** /openapi/cm/v1/signups/options | 
[**InitiateVerification**](ClientManagementApi.md#InitiateVerification) | **Post** /openapi/cm/v1/signups/verification/initiate/{ClientKey} | 
[**ResetPassword**](ClientManagementApi.md#ResetPassword) | **Post** /openapi/cm/v1/users/resetpasswordrequest | 
[**SignUp**](ClientManagementApi.md#SignUp) | **Post** /openapi/cm/v1/signups/ | 
[**UpdateClientRenewalData**](ClientManagementApi.md#UpdateClientRenewalData) | **Patch** /openapi/cm/v1/clientrenewals/{RenewalEntityId} | 

# **AttachFile**
> interface{} AttachFile(ctx, signUpId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **signUpId** | **string**|  | 
 **optional** | ***ClientManagementApiAttachFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientManagementApiAttachFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body16**](Body16.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachFileClientManagement**
> interface{} AttachFileClientManagement(ctx, signUpId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **signUpId** | **string**|  | 
 **optional** | ***ClientManagementApiAttachFileClientManagementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientManagementApiAttachFileClientManagementOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **documentType** | [**optional.Interface of []string**](string.md)|  | 
 **renewalDate** | **optional.String**|  | 
 **title** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompleteApplication**
> interface{} CompleteApplication(ctx, signUpId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **signUpId** | **string**|  | 
 **optional** | ***ClientManagementApiCompleteApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientManagementApiCompleteApplicationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awaitAccountCreation** | **optional.Bool**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateTypedOnboardingPDF**
> interface{} GenerateTypedOnboardingPDF(ctx, clientKey, documentType)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
  **documentType** | [**[]string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientRenewalData**
> interface{} GetClientRenewalData(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientManagementApiGetClientRenewalDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientManagementApiGetClientRenewalDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientKey** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRenewalStatuses**
> interface{} GetRenewalStatuses(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientManagementApiGetRenewalStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientManagementApiGetRenewalStatusesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **mustRenewBy** | **optional.String**|  | 
 **ownerKey** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSignUpStatus**
> interface{} GetSignUpStatus(ctx, clientKey)


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

# **GetSignupOptions**
> interface{} GetSignupOptions(ctx, )


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

# **GetSignupOptionsClientManagement**
> interface{} GetSignupOptionsClientManagement(ctx, )


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

# **InitiateVerification**
> interface{} InitiateVerification(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***ClientManagementApiInitiateVerificationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientManagementApiInitiateVerificationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body14**](Body14.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPassword**
> interface{} ResetPassword(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientManagementApiResetPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientManagementApiResetPasswordOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body15**](Body15.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SignUp**
> interface{} SignUp(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientManagementApiSignUpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientManagementApiSignUpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body13**](Body13.md)|  | 
 **ownerKey** | **optional.**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClientRenewalData**
> interface{} UpdateClientRenewalData(ctx, renewalEntityId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **renewalEntityId** | **string**|  | 
 **optional** | ***ClientManagementApiUpdateClientRenewalDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientManagementApiUpdateClientRenewalDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body12**](Body12.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

