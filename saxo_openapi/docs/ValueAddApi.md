# {{classname}}

All URIs are relative to *https://gateway.saxobank.com/sim*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAlertDefinitions**](ValueAddApi.md#DeleteAlertDefinitions) | **Delete** /openapi/vas/v1/pricealerts/definitions/{AlertDefinitionIds} | 
[**GetAlertDefinition**](ValueAddApi.md#GetAlertDefinition) | **Get** /openapi/vas/v1/pricealerts/definitions/{AlertDefinitionId} | 
[**GetAllAlertDefinitionsFilteredByState**](ValueAddApi.md#GetAllAlertDefinitionsFilteredByState) | **Get** /openapi/vas/v1/pricealerts/definitions/ | 
[**GetUserSettings**](ValueAddApi.md#GetUserSettings) | **Get** /openapi/vas/v1/pricealerts/usersettings | 
[**PostAlertDefinition**](ValueAddApi.md#PostAlertDefinition) | **Post** /openapi/vas/v1/pricealerts/definitions | 
[**PutAlertDefinition**](ValueAddApi.md#PutAlertDefinition) | **Put** /openapi/vas/v1/pricealerts/definitions/{AlertDefinitionId} | 
[**PutUserSettings**](ValueAddApi.md#PutUserSettings) | **Put** /openapi/vas/v1/pricealerts/usersettings | 

# **DeleteAlertDefinitions**
> interface{} DeleteAlertDefinitions(ctx, alertDefinitionIds)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionIds** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertDefinition**
> interface{} GetAlertDefinition(ctx, alertDefinitionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionId** | **int32**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAlertDefinitionsFilteredByState**
> interface{} GetAllAlertDefinitionsFilteredByState(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValueAddApiGetAllAlertDefinitionsFilteredByStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValueAddApiGetAllAlertDefinitionsFilteredByStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlinecount** | [**optional.Interface of []string**](string.md)|  | 
 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **state** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserSettings**
> interface{} GetUserSettings(ctx, )


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

# **PostAlertDefinition**
> interface{} PostAlertDefinition(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValueAddApiPostAlertDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValueAddApiPostAlertDefinitionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body74**](Body74.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAlertDefinition**
> interface{} PutAlertDefinition(ctx, alertDefinitionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertDefinitionId** | **int32**|  | 
 **optional** | ***ValueAddApiPutAlertDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValueAddApiPutAlertDefinitionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body75**](Body75.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUserSettings**
> interface{} PutUserSettings(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValueAddApiPutUserSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValueAddApiPutUserSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body76**](Body76.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

