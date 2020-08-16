# {{classname}}

All URIs are relative to *https://gateway.saxobank.com/sim*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseCase**](ClientServicesApi.md#CloseCase) | **Put** /openapi/cs/v1/partner/support/cases/{CaseId}/caseclose | 
[**CreateCase**](ClientServicesApi.md#CreateCase) | **Post** /openapi/cs/v1/partner/support/cases | 
[**CreateInternalComment**](ClientServicesApi.md#CreateInternalComment) | **Post** /openapi/cs/v1/partner/support/cases/{CaseId}/internalcomment | 
[**CreateNote**](ClientServicesApi.md#CreateNote) | **Post** /openapi/cs/v1/partner/support/cases/{CaseId}/note | 
[**Get**](ClientServicesApi.md#Get) | **Get** /openapi/cs/v1/audit/activities/ | 
[**GetCase**](ClientServicesApi.md#GetCase) | **Get** /openapi/cs/v1/partner/support/cases/{CaseId} | 
[**GetCases**](ClientServicesApi.md#GetCases) | **Get** /openapi/cs/v1/partner/support/cases/ | 
[**GetClientServices**](ClientServicesApi.md#GetClientServices) | **Get** /openapi/cs/v2/cashmanagement/wiretransfers/instructions/ | 
[**GetClientServices3**](ClientServicesApi.md#GetClientServices3) | **Get** /openapi/cs/v1/reports/aggregatedAmounts/{ClientKey}/{FromDate}/{ToDate}/ | 
[**GetClientServices4**](ClientServicesApi.md#GetClientServices4) | **Get** /openapi/cs/v1/reports/bookings/{ClientKey}/ | 
[**GetClientServices5**](ClientServicesApi.md#GetClientServices5) | **Get** /openapi/cs/v1/reports/closedPositions/{ClientKey}/{FromDate}/{ToDate}/ | 
[**GetDetailsAsync**](ClientServicesApi.md#GetDetailsAsync) | **Get** /openapi/cs/v1/reports/trades/{ClientKey}/ | 
[**GetForContractOption**](ClientServicesApi.md#GetForContractOption) | **Get** /openapi/cs/v1/tradingconditions/ContractOptionSpaces/{AccountKey}/{OptionRootId}/ | 
[**GetForInstrument**](ClientServicesApi.md#GetForInstrument) | **Get** /openapi/cs/v1/tradingconditions/instrument/{AccountKey}/{Uic}/{AssetType}/ | 
[**GetOrderStatesAsync**](ClientServicesApi.md#GetOrderStatesAsync) | **Get** /openapi/cs/v1/audit/orderactivities/ | 
[**GetTradingConditionCost**](ClientServicesApi.md#GetTradingConditionCost) | **Get** /openapi/cs/v1/tradingconditions/cost/{AccountKey}/{Uic}/{AssetType}/ | 
[**Search**](ClientServicesApi.md#Search) | **Post** /openapi/cs/v2/clientinfo/clients/search/ | 
[**Transfer**](ClientServicesApi.md#Transfer) | **Post** /openapi/cs/v2/cashmanagement/interaccounttransfers | 
[**UpdateNote**](ClientServicesApi.md#UpdateNote) | **Patch** /openapi/cs/v1/partner/support/cases/{CaseId} | 
[**WithdrawIBanSwift**](ClientServicesApi.md#WithdrawIBanSwift) | **Post** /openapi/cs/v2/cashmanagement/withdrawals/WithdrawIBanSwift | 
[**WithdrawalAccountNumberBeneficiaryBank**](ClientServicesApi.md#WithdrawalAccountNumberBeneficiaryBank) | **Post** /openapi/cs/v2/cashmanagement/withdrawals/WithdrawalAccountNumberBeneficiaryBank | 
[**WithdrawalAccountNumberSwift**](ClientServicesApi.md#WithdrawalAccountNumberSwift) | **Post** /openapi/cs/v2/cashmanagement/withdrawals/WithdrawalAccountNumberSwift | 
[**WithdrawalIBanBeneficiaryBank**](ClientServicesApi.md#WithdrawalIBanBeneficiaryBank) | **Post** /openapi/cs/v2/cashmanagement/withdrawals/WithdrawalIbanBeneficiaryBank | 

# **CloseCase**
> interface{} CloseCase(ctx, caseId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **caseId** | **string**|  | 
 **optional** | ***ClientServicesApiCloseCaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiCloseCaseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body19**](Body19.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCase**
> interface{} CreateCase(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientServicesApiCreateCaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiCreateCaseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body17**](Body17.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInternalComment**
> interface{} CreateInternalComment(ctx, caseId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **caseId** | **string**|  | 
 **optional** | ***ClientServicesApiCreateInternalCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiCreateInternalCommentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body20**](Body20.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNote**
> interface{} CreateNote(ctx, caseId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **caseId** | **string**|  | 
 **optional** | ***ClientServicesApiCreateNoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiCreateNoteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body21**](Body21.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> interface{} Get(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***ClientServicesApiGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skiptoken** | **optional.String**|  | 
 **top** | **optional.Int32**|  | 
 **accountKey** | **optional.String**|  | 
 **correlationKey** | **optional.String**|  | 
 **from** | **optional.String**|  | 
 **logEntryTypes** | [**optional.Interface of []string**](string.md)|  | 
 **postingId** | **optional.String**|  | 
 **to** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCase**
> interface{} GetCase(ctx, caseId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **caseId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCases**
> interface{} GetCases(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientServicesApiGetCasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiGetCasesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **optional.Int32**|  | 
 **fromDateTime** | **optional.String**|  | 
 **status** | [**optional.Interface of []string**](string.md)|  | 
 **toDateTime** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientServices**
> interface{} GetClientServices(ctx, accountKey, clientKey, currencyCode)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountKey** | **string**|  | 
  **clientKey** | **string**|  | 
  **currencyCode** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientServices3**
> interface{} GetClientServices3(ctx, clientKey, fromDate, toDate, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
  **fromDate** | **string**|  | 
  **toDate** | **string**|  | 
 **optional** | ***ClientServicesApiGetClientServices3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiGetClientServices3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **skip** | **optional.Int32**|  | 
 **skiptoken** | **optional.String**|  | 
 **top** | **optional.Int32**|  | 
 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **mockDataId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientServices4**
> interface{} GetClientServices4(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***ClientServicesApiGetClientServices4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiGetClientServices4Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **optional.Int32**|  | 
 **skiptoken** | **optional.String**|  | 
 **top** | **optional.Int32**|  | 
 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **filterType** | [**optional.Interface of []string**](string.md)|  | 
 **filterValue** | **optional.String**|  | 
 **fromDate** | **optional.String**|  | 
 **mockDataId** | **optional.String**|  | 
 **toDate** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientServices5**
> interface{} GetClientServices5(ctx, clientKey, fromDate, toDate, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
  **fromDate** | **string**|  | 
  **toDate** | **string**|  | 
 **optional** | ***ClientServicesApiGetClientServices5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiGetClientServices5Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **skip** | **optional.Int32**|  | 
 **top** | **optional.Int32**|  | 
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

# **GetDetailsAsync**
> interface{} GetDetailsAsync(ctx, clientKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientKey** | **string**|  | 
 **optional** | ***ClientServicesApiGetDetailsAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiGetDetailsAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **optional.Int32**|  | 
 **skiptoken** | **optional.String**|  | 
 **top** | **optional.Int32**|  | 
 **accountGroupKey** | **optional.String**|  | 
 **accountKey** | **optional.String**|  | 
 **fromDate** | **optional.String**|  | 
 **mockDataId** | **optional.String**|  | 
 **toDate** | **optional.String**|  | 
 **tradeId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForContractOption**
> interface{} GetForContractOption(ctx, accountKey, optionRootId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountKey** | **string**|  | 
  **optionRootId** | **int32**|  | 
 **optional** | ***ClientServicesApiGetForContractOptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiGetForContractOptionOpts struct
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

# **GetForInstrument**
> interface{} GetForInstrument(ctx, accountKey, assetType, uic, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountKey** | **string**|  | 
  **assetType** | [**[]string**](string.md)|  | 
  **uic** | **int32**|  | 
 **optional** | ***ClientServicesApiGetForInstrumentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiGetForInstrumentOpts struct
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

# **GetOrderStatesAsync**
> interface{} GetOrderStatesAsync(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientServicesApiGetOrderStatesAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiGetOrderStatesAsyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skiptoken** | **optional.String**|  | 
 **top** | **optional.Int32**|  | 
 **accountKey** | **optional.String**|  | 
 **clientKey** | **optional.String**|  | 
 **correlationKey** | **optional.String**|  | 
 **entryType** | [**optional.Interface of []string**](string.md)|  | 
 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 
 **fromDateTime** | **optional.String**|  | 
 **includeSubAccounts** | **optional.Bool**|  | 
 **orderId** | **optional.String**|  | 
 **status** | [**optional.Interface of []string**](string.md)|  | 
 **toDateTime** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTradingConditionCost**
> interface{} GetTradingConditionCost(ctx, accountKey, amount, assetType, uic, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountKey** | **string**|  | 
  **amount** | **float64**|  | 
  **assetType** | [**[]string**](string.md)|  | 
  **uic** | **int32**|  | 
 **optional** | ***ClientServicesApiGetTradingConditionCostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiGetTradingConditionCostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **fieldGroups** | [**optional.Interface of []string**](string.md)|  | 
 **holdingPeriodInDays** | **optional.Int32**|  | 
 **price** | **optional.Float64**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Search**
> interface{} Search(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientServicesApiSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body27**](Body27.md)|  | 
 **inlinecount** | [**optional.Interface of []string**](string.md)|  | 
 **skip** | **optional.**|  | 
 **top** | **optional.**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Transfer**
> interface{} Transfer(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientServicesApiTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiTransferOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body22**](Body22.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNote**
> interface{} UpdateNote(ctx, caseId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **caseId** | **string**|  | 
 **optional** | ***ClientServicesApiUpdateNoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiUpdateNoteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body18**](Body18.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WithdrawIBanSwift**
> interface{} WithdrawIBanSwift(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientServicesApiWithdrawIBanSwiftOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiWithdrawIBanSwiftOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body23**](Body23.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WithdrawalAccountNumberBeneficiaryBank**
> interface{} WithdrawalAccountNumberBeneficiaryBank(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientServicesApiWithdrawalAccountNumberBeneficiaryBankOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiWithdrawalAccountNumberBeneficiaryBankOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body24**](Body24.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WithdrawalAccountNumberSwift**
> interface{} WithdrawalAccountNumberSwift(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientServicesApiWithdrawalAccountNumberSwiftOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiWithdrawalAccountNumberSwiftOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body25**](Body25.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WithdrawalIBanBeneficiaryBank**
> interface{} WithdrawalIBanBeneficiaryBank(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientServicesApiWithdrawalIBanBeneficiaryBankOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientServicesApiWithdrawalIBanBeneficiaryBankOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body26**](Body26.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[openapi_auth](../README.md#openapi_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

