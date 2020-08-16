# {{classname}}

All URIs are relative to *https://gateway.saxobank.com/sim*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddActiveSubscription**](PortfolioApi.md#AddActiveSubscription) | **Post** /openapi/port/v1/accounts/subscriptions | 
[**AddActiveSubscriptionPortfolio**](PortfolioApi.md#AddActiveSubscriptionPortfolio) | **Post** /openapi/port/v1/balances/subscriptions | 
[**AddActiveSubscriptionPortfolio3**](PortfolioApi.md#AddActiveSubscriptionPortfolio3) | **Post** /openapi/port/v1/closedpositions/subscriptions/ | 
[**AddActiveSubscriptionPortfolio4**](PortfolioApi.md#AddActiveSubscriptionPortfolio4) | **Post** /openapi/port/v1/exposure/instruments/subscriptions | 
[**AddActiveSubscriptionPortfolio5**](PortfolioApi.md#AddActiveSubscriptionPortfolio5) | **Post** /openapi/port/v1/netpositions/subscriptions | 
[**AddActiveSubscriptionPortfolio6**](PortfolioApi.md#AddActiveSubscriptionPortfolio6) | **Post** /openapi/port/v1/orders/subscriptions/ | 
[**AddActiveSubscriptionPortfolio7**](PortfolioApi.md#AddActiveSubscriptionPortfolio7) | **Post** /openapi/port/v1/positions/subscriptions/ | 
[**DeleteSubscriptionPortfolio10**](PortfolioApi.md#DeleteSubscriptionPortfolio10) | **Delete** /openapi/port/v1/positions/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteSubscriptionPortfolio4**](PortfolioApi.md#DeleteSubscriptionPortfolio4) | **Delete** /openapi/port/v1/accounts/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteSubscriptionPortfolio5**](PortfolioApi.md#DeleteSubscriptionPortfolio5) | **Delete** /openapi/port/v1/balances/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteSubscriptionPortfolio6**](PortfolioApi.md#DeleteSubscriptionPortfolio6) | **Delete** /openapi/port/v1/closedpositions/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteSubscriptionPortfolio7**](PortfolioApi.md#DeleteSubscriptionPortfolio7) | **Delete** /openapi/port/v1/exposure/instruments/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteSubscriptionPortfolio8**](PortfolioApi.md#DeleteSubscriptionPortfolio8) | **Delete** /openapi/port/v1/netpositions/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteSubscriptionPortfolio9**](PortfolioApi.md#DeleteSubscriptionPortfolio9) | **Delete** /openapi/port/v1/orders/subscriptions/{ContextId}/{ReferenceId} | 
[**DeleteSubscriptionsPortfolio3**](PortfolioApi.md#DeleteSubscriptionsPortfolio3) | **Delete** /openapi/port/v1/accounts/subscriptions/{ContextId}/ | 
[**DeleteSubscriptionsPortfolio4**](PortfolioApi.md#DeleteSubscriptionsPortfolio4) | **Delete** /openapi/port/v1/balances/subscriptions/{ContextId}/ | 
[**DeleteSubscriptionsPortfolio5**](PortfolioApi.md#DeleteSubscriptionsPortfolio5) | **Delete** /openapi/port/v1/closedpositions/subscriptions/{ContextId}/ | 
[**DeleteSubscriptionsPortfolio6**](PortfolioApi.md#DeleteSubscriptionsPortfolio6) | **Delete** /openapi/port/v1/exposure/instruments/subscriptions/{ContextId}/ | 
[**DeleteSubscriptionsPortfolio7**](PortfolioApi.md#DeleteSubscriptionsPortfolio7) | **Delete** /openapi/port/v1/netpositions/subscriptions/{ContextId}/ | 
[**DeleteSubscriptionsPortfolio8**](PortfolioApi.md#DeleteSubscriptionsPortfolio8) | **Delete** /openapi/port/v1/orders/subscriptions/{ContextId}/ | 
[**DeleteSubscriptionsPortfolio9**](PortfolioApi.md#DeleteSubscriptionsPortfolio9) | **Delete** /openapi/port/v1/positions/subscriptions/{ContextId}/ | 
[**GetAccount**](PortfolioApi.md#GetAccount) | **Get** /openapi/port/v1/accounts/{AccountKey} | 
[**GetAccountGroup**](PortfolioApi.md#GetAccountGroup) | **Get** /openapi/port/v1/accountgroups/{AccountGroupKey}/ | 
[**GetAccountGroups**](PortfolioApi.md#GetAccountGroups) | **Get** /openapi/port/v1/accountgroups/ | 
[**GetAccountGroupsForLoggedInUser**](PortfolioApi.md#GetAccountGroupsForLoggedInUser) | **Get** /openapi/port/v1/accountgroups/me/ | 
[**GetAccounts**](PortfolioApi.md#GetAccounts) | **Get** /openapi/port/v1/accounts/me/ | 
[**GetAccountsPortfolio**](PortfolioApi.md#GetAccountsPortfolio) | **Get** /openapi/port/v1/accounts/ | 
[**GetAllClientEntitlements**](PortfolioApi.md#GetAllClientEntitlements) | **Get** /openapi/port/v1/users/me/entitlements | 
[**GetBalance**](PortfolioApi.md#GetBalance) | **Get** /openapi/port/v1/balances/me | 
[**GetBalancePortfolio**](PortfolioApi.md#GetBalancePortfolio) | **Get** /openapi/port/v1/balances/ | 
[**GetClient**](PortfolioApi.md#GetClient) | **Get** /openapi/port/v1/clients/me | 
[**GetClientPortfolio**](PortfolioApi.md#GetClientPortfolio) | **Get** /openapi/port/v1/clients/{ClientKey} | 
[**GetClients**](PortfolioApi.md#GetClients) | **Get** /openapi/port/v1/clients/ | 
[**GetClosedPosition**](PortfolioApi.md#GetClosedPosition) | **Get** /openapi/port/v1/closedpositions/{ClosedPositionId}/ | 
[**GetClosedPositions**](PortfolioApi.md#GetClosedPositions) | **Get** /openapi/port/v1/closedpositions/ | 
[**GetClosedPositionsPortfolio**](PortfolioApi.md#GetClosedPositionsPortfolio) | **Get** /openapi/port/v1/closedpositions/me/ | 
[**GetCurrencyExposures**](PortfolioApi.md#GetCurrencyExposures) | **Get** /openapi/port/v1/exposure/currency/me | 
[**GetCurrencyExposuresForLoggedInUser**](PortfolioApi.md#GetCurrencyExposuresForLoggedInUser) | **Get** /openapi/port/v1/exposure/currency/ | 
[**GetFxSpotNetExposures**](PortfolioApi.md#GetFxSpotNetExposures) | **Get** /openapi/port/v1/exposure/fxspot/me | 
[**GetFxSpotNetExposuresForLoggedInUser**](PortfolioApi.md#GetFxSpotNetExposuresForLoggedInUser) | **Get** /openapi/port/v1/exposure/fxspot/ | 
[**GetInstrumentsExposures**](PortfolioApi.md#GetInstrumentsExposures) | **Get** /openapi/port/v1/exposure/instruments/me | 
[**GetInstrumentsExposuresForLoggedInUser**](PortfolioApi.md#GetInstrumentsExposuresForLoggedInUser) | **Get** /openapi/port/v1/exposure/instruments/ | 
[**GetMarginOverview**](PortfolioApi.md#GetMarginOverview) | **Get** /openapi/port/v1/balances/marginoverview/ | 
[**GetNetPosition**](PortfolioApi.md#GetNetPosition) | **Get** /openapi/port/v1/netpositions/{NetPositionId}/ | 
[**GetNetPositionDetails**](PortfolioApi.md#GetNetPositionDetails) | **Get** /openapi/port/v1/netpositions/{NetPositionId}/details/ | 
[**GetNetPositions**](PortfolioApi.md#GetNetPositions) | **Get** /openapi/port/v1/netpositions/me/ | 
[**GetNetPositionsPortfolio**](PortfolioApi.md#GetNetPositionsPortfolio) | **Get** /openapi/port/v1/netpositions/ | 
[**GetOpenOrder**](PortfolioApi.md#GetOpenOrder) | **Get** /openapi/port/v1/orders/{ClientKey}/{OrderId}/ | 
[**GetOpenOrderDetails**](PortfolioApi.md#GetOpenOrderDetails) | **Get** /openapi/port/v1/orders/{OrderId}/details/ | 
[**GetOpenOrders**](PortfolioApi.md#GetOpenOrders) | **Get** /openapi/port/v1/orders/me/ | 
[**GetOpenOrdersPortfolio**](PortfolioApi.md#GetOpenOrdersPortfolio) | **Get** /openapi/port/v1/orders/ | 
[**GetPosition**](PortfolioApi.md#GetPosition) | **Get** /openapi/port/v1/positions/{PositionId}/ | 
[**GetPositionDetails**](PortfolioApi.md#GetPositionDetails) | **Get** /openapi/port/v1/closedpositions/{ClosedPositionId}/details/ | 
[**GetPositionDetailsPortfolio**](PortfolioApi.md#GetPositionDetailsPortfolio) | **Get** /openapi/port/v1/positions/{PositionId}/details/ | 
[**GetPositions**](PortfolioApi.md#GetPositions) | **Get** /openapi/port/v1/positions/me/ | 
[**GetPositionsPortfolio**](PortfolioApi.md#GetPositionsPortfolio) | **Get** /openapi/port/v1/positions/ | 
[**GetUser**](PortfolioApi.md#GetUser) | **Get** /openapi/port/v1/users/me | 
[**GetUserPortfolio**](PortfolioApi.md#GetUserPortfolio) | **Get** /openapi/port/v1/users/{UserKey} | 
[**GetUsers**](PortfolioApi.md#GetUsers) | **Get** /openapi/port/v1/users/ | 
[**ResetAccount**](PortfolioApi.md#ResetAccount) | **Put** /openapi/port/v1/accounts/{AccountKey}/reset | 
[**UpdateAccount**](PortfolioApi.md#UpdateAccount) | **Patch** /openapi/port/v1/accounts/{AccountKey} | 
[**UpdateAccountGroup**](PortfolioApi.md#UpdateAccountGroup) | **Patch** /openapi/port/v1/accountgroups/{AccountGroupKey}/ | 
[**UpdateClientSettings**](PortfolioApi.md#UpdateClientSettings) | **Patch** /openapi/port/v1/clients/me | 
[**UpdateClientSettingsForPartner**](PortfolioApi.md#UpdateClientSettingsForPartner) | **Patch** /openapi/port/v1/clients/ | 
[**UpdateClosedPositionSubscriptionPagesize**](PortfolioApi.md#UpdateClosedPositionSubscriptionPagesize) | **Patch** /openapi/port/v1/closedpositions/subscriptions/{ContextId}/{ReferenceId} | 
[**UpdatePositionSubscriptionPagesize**](PortfolioApi.md#UpdatePositionSubscriptionPagesize) | **Patch** /openapi/port/v1/positions/subscriptions/{ContextId}/{ReferenceId} | 
[**UpdateUserPreferences**](PortfolioApi.md#UpdateUserPreferences) | **Patch** /openapi/port/v1/users/me | 

# **AddActiveSubscription**
> interface{} AddActiveSubscription(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiAddActiveSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiAddActiveSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body40**](Body40.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddActiveSubscriptionPortfolio**
> interface{} AddActiveSubscriptionPortfolio(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiAddActiveSubscriptionPortfolioOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiAddActiveSubscriptionPortfolioOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body41**](Body41.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddActiveSubscriptionPortfolio3**
> interface{} AddActiveSubscriptionPortfolio3(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiAddActiveSubscriptionPortfolio3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiAddActiveSubscriptionPortfolio3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body44**](Body44.md)|  | 
 **top** | **optional.**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddActiveSubscriptionPortfolio4**
> interface{} AddActiveSubscriptionPortfolio4(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiAddActiveSubscriptionPortfolio4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiAddActiveSubscriptionPortfolio4Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body46**](Body46.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddActiveSubscriptionPortfolio5**
> interface{} AddActiveSubscriptionPortfolio5(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiAddActiveSubscriptionPortfolio5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiAddActiveSubscriptionPortfolio5Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body47**](Body47.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddActiveSubscriptionPortfolio6**
> interface{} AddActiveSubscriptionPortfolio6(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiAddActiveSubscriptionPortfolio6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiAddActiveSubscriptionPortfolio6Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body48**](Body48.md)|  | 
 **inlinecount** | [**optional.Interface of []string**](string.md)|  | 
 **skip** | **optional.**|  | 
 **skiptoken** | **optional.**|  | 
 **top** | **optional.**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddActiveSubscriptionPortfolio7**
> interface{} AddActiveSubscriptionPortfolio7(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiAddActiveSubscriptionPortfolio7Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiAddActiveSubscriptionPortfolio7Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body49**](Body49.md)|  | 
 **top** | **optional.**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscriptionPortfolio10**
> interface{} DeleteSubscriptionPortfolio10(ctx, contextId, referenceId)


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

# **DeleteSubscriptionPortfolio4**
> interface{} DeleteSubscriptionPortfolio4(ctx, contextId, referenceId)


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

# **DeleteSubscriptionPortfolio5**
> interface{} DeleteSubscriptionPortfolio5(ctx, contextId, referenceId)


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

# **DeleteSubscriptionPortfolio6**
> interface{} DeleteSubscriptionPortfolio6(ctx, contextId, referenceId)


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

# **DeleteSubscriptionPortfolio7**
> interface{} DeleteSubscriptionPortfolio7(ctx, contextId, referenceId)


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

# **DeleteSubscriptionPortfolio8**
> interface{} DeleteSubscriptionPortfolio8(ctx, contextId, referenceId)


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

# **DeleteSubscriptionPortfolio9**
> interface{} DeleteSubscriptionPortfolio9(ctx, contextId, referenceId)


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

# **DeleteSubscriptionsPortfolio3**
> interface{} DeleteSubscriptionsPortfolio3(ctx, contextId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
 **optional** | ***PortfolioApiDeleteSubscriptionsPortfolio3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiDeleteSubscriptionsPortfolio3Opts struct
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

# **DeleteSubscriptionsPortfolio4**
> interface{} DeleteSubscriptionsPortfolio4(ctx, contextId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
 **optional** | ***PortfolioApiDeleteSubscriptionsPortfolio4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiDeleteSubscriptionsPortfolio4Opts struct
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

# **DeleteSubscriptionsPortfolio5**
> interface{} DeleteSubscriptionsPortfolio5(ctx, contextId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
 **optional** | ***PortfolioApiDeleteSubscriptionsPortfolio5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiDeleteSubscriptionsPortfolio5Opts struct
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

# **DeleteSubscriptionsPortfolio6**
> interface{} DeleteSubscriptionsPortfolio6(ctx, contextId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
 **optional** | ***PortfolioApiDeleteSubscriptionsPortfolio6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiDeleteSubscriptionsPortfolio6Opts struct
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

# **DeleteSubscriptionsPortfolio7**
> interface{} DeleteSubscriptionsPortfolio7(ctx, contextId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
 **optional** | ***PortfolioApiDeleteSubscriptionsPortfolio7Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiDeleteSubscriptionsPortfolio7Opts struct
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

# **DeleteSubscriptionsPortfolio8**
> interface{} DeleteSubscriptionsPortfolio8(ctx, contextId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
 **optional** | ***PortfolioApiDeleteSubscriptionsPortfolio8Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiDeleteSubscriptionsPortfolio8Opts struct
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

# **DeleteSubscriptionsPortfolio9**
> interface{} DeleteSubscriptionsPortfolio9(ctx, contextId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
 **optional** | ***PortfolioApiDeleteSubscriptionsPortfolio9Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiDeleteSubscriptionsPortfolio9Opts struct
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

# **GetAccount**
> interface{} GetAccount(ctx, accountKey)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountKey** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountGroup**
> interface{} GetAccountGroup(ctx, accountGroupKey, clientKey)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountGroupKey** | **string**|  | 
  **clientKey** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountGroups**
> interface{} GetAccountGroups(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiGetAccountGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetAccountGroupsOpts struct
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

# **GetAccountGroupsForLoggedInUser**
> interface{} GetAccountGroupsForLoggedInUser(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiGetAccountGroupsForLoggedInUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetAccountGroupsForLoggedInUserOpts struct
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

# **GetAccounts**
> interface{} GetAccounts(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiGetAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetAccountsOpts struct
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

# **GetAccountsPortfolio**
> interface{} GetAccountsPortfolio(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiGetAccountsPortfolioOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetAccountsPortfolioOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlinecount** | [**optional.Interface of []string**](string.md)|  | 
 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **includeSubAccounts** | **optional.Bool**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllClientEntitlements**
> interface{} GetAllClientEntitlements(ctx, )


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

# **GetBalance**
> interface{} GetBalance(ctx, )


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

# **GetBalancePortfolio**
> interface{} GetBalancePortfolio(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiGetBalancePortfolioOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetBalancePortfolioOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClient**
> interface{} GetClient(ctx, )


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

# **GetClientPortfolio**
> interface{} GetClientPortfolio(ctx, clientKey)


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

# **GetClients**
> interface{} GetClients(ctx, ownerKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***PortfolioApiGetClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetClientsOpts struct
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

# **GetClosedPosition**
> interface{} GetClosedPosition(ctx, clientKey, closedPositionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
  **closedPositionId** | **string**|  | 
 **optional** | ***PortfolioApiGetClosedPositionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetClosedPositionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClosedPositions**
> interface{} GetClosedPositions(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiGetClosedPositionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetClosedPositionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **closedPositionId** | **optional.String**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClosedPositionsPortfolio**
> interface{} GetClosedPositionsPortfolio(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiGetClosedPositionsPortfolioOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetClosedPositionsPortfolioOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrencyExposures**
> interface{} GetCurrencyExposures(ctx, )


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

# **GetCurrencyExposuresForLoggedInUser**
> interface{} GetCurrencyExposuresForLoggedInUser(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiGetCurrencyExposuresForLoggedInUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetCurrencyExposuresForLoggedInUserOpts struct
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

# **GetFxSpotNetExposures**
> interface{} GetFxSpotNetExposures(ctx, )


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

# **GetFxSpotNetExposuresForLoggedInUser**
> interface{} GetFxSpotNetExposuresForLoggedInUser(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiGetFxSpotNetExposuresForLoggedInUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetFxSpotNetExposuresForLoggedInUserOpts struct
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

# **GetInstrumentsExposures**
> interface{} GetInstrumentsExposures(ctx, )


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

# **GetInstrumentsExposuresForLoggedInUser**
> interface{} GetInstrumentsExposuresForLoggedInUser(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiGetInstrumentsExposuresForLoggedInUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetInstrumentsExposuresForLoggedInUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **assetType** | [**optional.Interface of []string**](string.md)|  | 
 **expiryDate** | **optional.String**|  | 
 **lowerBarrier** | **optional.Float64**|  | 
 **putCall** | [**optional.Interface of []string**](string.md)|  | 
 **strike** | **optional.Float64**|  | 
 **uic** | **optional.Int32**|  | 
 **upperBarrier** | **optional.Float64**|  | 
 **valueDate** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMarginOverview**
> interface{} GetMarginOverview(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiGetMarginOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetMarginOverviewOpts struct
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

# **GetNetPosition**
> interface{} GetNetPosition(ctx, clientKey, netPositionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
  **netPositionId** | **string**|  | 
 **optional** | ***PortfolioApiGetNetPositionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetNetPositionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetPositionDetails**
> interface{} GetNetPositionDetails(ctx, clientKey, netPositionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
  **netPositionId** | **string**|  | 
 **optional** | ***PortfolioApiGetNetPositionDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetNetPositionDetailsOpts struct
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

# **GetNetPositions**
> interface{} GetNetPositions(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiGetNetPositionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetNetPositionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetPositionsPortfolio**
> interface{} GetNetPositionsPortfolio(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiGetNetPositionsPortfolioOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetNetPositionsPortfolioOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **assetType** | [**optional.Interface of []string**](string.md)|  | 
 **expiryDate** | **optional.String**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 
 **lowerBarrier** | **optional.Float64**|  | 
 **netPositionId** | **optional.String**|  | 
 **putCall** | [**optional.Interface of []string**](string.md)|  | 
 **strike** | **optional.Float64**|  | 
 **uic** | **optional.Int32**|  | 
 **upperBarrier** | **optional.Float64**|  | 
 **valueDate** | **optional.String**|  | 
 **watchlistId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpenOrder**
> interface{} GetOpenOrder(ctx, clientKey, orderId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
  **orderId** | **string**|  | 
 **optional** | ***PortfolioApiGetOpenOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetOpenOrderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpenOrderDetails**
> interface{} GetOpenOrderDetails(ctx, clientKey, orderId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
  **orderId** | **string**|  | 
 **optional** | ***PortfolioApiGetOpenOrderDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetOpenOrderDetailsOpts struct
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

# **GetOpenOrders**
> interface{} GetOpenOrders(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiGetOpenOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetOpenOrdersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 
 **multiLegOrderId** | **optional.String**|  | 
 **status** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpenOrdersPortfolio**
> interface{} GetOpenOrdersPortfolio(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiGetOpenOrdersPortfolioOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetOpenOrdersPortfolioOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 
 **orderId** | **optional.String**|  | 
 **status** | [**optional.Interface of []string**](string.md)|  | 
 **watchlistId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPosition**
> interface{} GetPosition(ctx, clientKey, positionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
  **positionId** | **string**|  | 
 **optional** | ***PortfolioApiGetPositionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetPositionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPositionDetails**
> interface{} GetPositionDetails(ctx, clientKey, closedPositionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
  **closedPositionId** | **string**|  | 
 **optional** | ***PortfolioApiGetPositionDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetPositionDetailsOpts struct
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

# **GetPositionDetailsPortfolio**
> interface{} GetPositionDetailsPortfolio(ctx, clientKey, positionId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
  **positionId** | **string**|  | 
 **optional** | ***PortfolioApiGetPositionDetailsPortfolioOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetPositionDetailsPortfolioOpts struct
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

# **GetPositions**
> interface{} GetPositions(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiGetPositionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetPositionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPositionsPortfolio**
> interface{} GetPositionsPortfolio(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiGetPositionsPortfolioOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetPositionsPortfolioOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 
 **netPositionId** | **optional.String**|  | 
 **positionId** | **optional.String**|  | 
 **watchlistId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> interface{} GetUser(ctx, )


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

# **GetUserPortfolio**
> interface{} GetUserPortfolio(ctx, userKey)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userKey** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> interface{} GetUsers(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiGetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiGetUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlinecount** | [**optional.Interface of []string**](string.md)|  | 
 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
 **includeSubUsers** | **optional.Bool**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetAccount**
> interface{} ResetAccount(ctx, accountKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountKey** | **string**|  | 
 **optional** | ***PortfolioApiResetAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiResetAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body39**](Body39.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccount**
> interface{} UpdateAccount(ctx, accountKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountKey** | **string**|  | 
 **optional** | ***PortfolioApiUpdateAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiUpdateAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body38**](Body38.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountGroup**
> interface{} UpdateAccountGroup(ctx, accountGroupKey, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountGroupKey** | **string**|  | 
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiUpdateAccountGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiUpdateAccountGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Body37**](Body37.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClientSettings**
> interface{} UpdateClientSettings(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiUpdateClientSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiUpdateClientSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body43**](Body43.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClientSettingsForPartner**
> interface{} UpdateClientSettingsForPartner(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***PortfolioApiUpdateClientSettingsForPartnerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiUpdateClientSettingsForPartnerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body42**](Body42.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClosedPositionSubscriptionPagesize**
> interface{} UpdateClosedPositionSubscriptionPagesize(ctx, contextId, referenceId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
  **referenceId** | **string**|  | 
 **optional** | ***PortfolioApiUpdateClosedPositionSubscriptionPagesizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiUpdateClosedPositionSubscriptionPagesizeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Body45**](Body45.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePositionSubscriptionPagesize**
> interface{} UpdatePositionSubscriptionPagesize(ctx, contextId, referenceId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**|  | 
  **referenceId** | **string**|  | 
 **optional** | ***PortfolioApiUpdatePositionSubscriptionPagesizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiUpdatePositionSubscriptionPagesizeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Body50**](Body50.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserPreferences**
> interface{} UpdateUserPreferences(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortfolioApiUpdateUserPreferencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortfolioApiUpdateUserPreferencesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body51**](Body51.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

