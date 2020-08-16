# {{classname}}

All URIs are relative to *https://gateway.saxobank.com/sim*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApp**](DeveloperApi.md#CreateApp) | **Post** /openapi/developer/apps | 
[**CreateRedirectUri**](DeveloperApi.md#CreateRedirectUri) | **Post** /openapi/developer/apps/{AppKey}/redirecturis | 
[**CreateSecret**](DeveloperApi.md#CreateSecret) | **Post** /openapi/developer/apps/{AppKey}/secrets | 
[**DeactivateApp**](DeveloperApi.md#DeactivateApp) | **Delete** /openapi/developer/apps/{AppKey} | 
[**DeleteRedirectUri**](DeveloperApi.md#DeleteRedirectUri) | **Delete** /openapi/developer/apps/{AppKey}/redirecturis/{RedirectUriId} | 
[**DeleteSecret**](DeveloperApi.md#DeleteSecret) | **Delete** /openapi/developer/apps/{AppKey}/secrets/{SecretId} | 
[**GetAppByAppKey**](DeveloperApi.md#GetAppByAppKey) | **Get** /openapi/developer/apps/{AppKey} | 
[**GetApps**](DeveloperApi.md#GetApps) | **Get** /openapi/developer/apps/ | 
[**GetFlowTypes**](DeveloperApi.md#GetFlowTypes) | **Get** /openapi/developer/apps/flowtypes/ | 
[**GetRedirectUriFromId**](DeveloperApi.md#GetRedirectUriFromId) | **Get** /openapi/developer/apps/{AppKey}/redirecturis/{RedirectUriId}/ | 
[**GetRedirectUris**](DeveloperApi.md#GetRedirectUris) | **Get** /openapi/developer/apps/{AppKey}/redirecturis/ | 
[**GetSecretById**](DeveloperApi.md#GetSecretById) | **Get** /openapi/developer/apps/{AppKey}/secrets/{SecretId} | 
[**GetSecrets**](DeveloperApi.md#GetSecrets) | **Get** /openapi/developer/apps/{AppKey}/secrets/ | 
[**UpdateApp**](DeveloperApi.md#UpdateApp) | **Patch** /openapi/developer/apps/{AppKey} | 
[**UpdateRedirectUri**](DeveloperApi.md#UpdateRedirectUri) | **Patch** /openapi/developer/apps/{AppKey}/redirecturis/{RedirectUriId} | 
[**UpdateSecret**](DeveloperApi.md#UpdateSecret) | **Patch** /openapi/developer/apps/{AppKey}/secrets/{SecretId} | 

# **CreateApp**
> interface{} CreateApp(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeveloperApiCreateAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeveloperApiCreateAppOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body28**](Body28.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRedirectUri**
> interface{} CreateRedirectUri(ctx, appKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appKey** | **string**|  | 
 **optional** | ***DeveloperApiCreateRedirectUriOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeveloperApiCreateRedirectUriOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body30**](Body30.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSecret**
> interface{} CreateSecret(ctx, appKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appKey** | **string**|  | 
 **optional** | ***DeveloperApiCreateSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeveloperApiCreateSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body32**](Body32.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateApp**
> interface{} DeactivateApp(ctx, appKey)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appKey** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRedirectUri**
> interface{} DeleteRedirectUri(ctx, appKey, redirectUriId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appKey** | **string**|  | 
  **redirectUriId** | **int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecret**
> interface{} DeleteSecret(ctx, appKey, secretId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appKey** | **string**|  | 
  **secretId** | **int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppByAppKey**
> interface{} GetAppByAppKey(ctx, appKey)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appKey** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApps**
> interface{} GetApps(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeveloperApiGetAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeveloperApiGetAppsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlinecount** | [**optional.Interface of []string**](string.md)|  | 
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

# **GetFlowTypes**
> interface{} GetFlowTypes(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeveloperApiGetFlowTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeveloperApiGetFlowTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlinecount** | [**optional.Interface of []string**](string.md)|  | 
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

# **GetRedirectUriFromId**
> interface{} GetRedirectUriFromId(ctx, appKey, redirectUriId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appKey** | **string**|  | 
  **redirectUriId** | **int32**|  | 
 **optional** | ***DeveloperApiGetRedirectUriFromIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeveloperApiGetRedirectUriFromIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlinecount** | [**optional.Interface of []string**](string.md)|  | 
 **skip** | **optional.Int32**|  | 
 **skiptoken** | **optional.String**|  | 
 **top** | **optional.Int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRedirectUris**
> interface{} GetRedirectUris(ctx, appKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appKey** | **string**|  | 
 **optional** | ***DeveloperApiGetRedirectUrisOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeveloperApiGetRedirectUrisOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlinecount** | [**optional.Interface of []string**](string.md)|  | 
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

# **GetSecretById**
> interface{} GetSecretById(ctx, appKey, secretId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appKey** | **string**|  | 
  **secretId** | **int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecrets**
> interface{} GetSecrets(ctx, appKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appKey** | **string**|  | 
 **optional** | ***DeveloperApiGetSecretsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeveloperApiGetSecretsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlinecount** | [**optional.Interface of []string**](string.md)|  | 
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

# **UpdateApp**
> interface{} UpdateApp(ctx, appKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appKey** | **string**|  | 
 **optional** | ***DeveloperApiUpdateAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeveloperApiUpdateAppOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body29**](Body29.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRedirectUri**
> interface{} UpdateRedirectUri(ctx, appKey, redirectUriId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appKey** | **string**|  | 
  **redirectUriId** | **int32**|  | 
 **optional** | ***DeveloperApiUpdateRedirectUriOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeveloperApiUpdateRedirectUriOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Body31**](Body31.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecret**
> interface{} UpdateSecret(ctx, appKey, secretId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appKey** | **string**|  | 
  **secretId** | **int32**|  | 
 **optional** | ***DeveloperApiUpdateSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeveloperApiUpdateSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Body33**](Body33.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

