# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowsCreate**](DefaultApi.md#FlowsCreate) | **Post** /v2/Flows | 
[**FlowsDelete**](DefaultApi.md#FlowsDelete) | **Delete** /v2/Flows/{Sid} | 
[**FlowsRead**](DefaultApi.md#FlowsRead) | **Get** /v2/Flows/{Sid} | 
[**FlowsUpdate**](DefaultApi.md#FlowsUpdate) | **Post** /v2/Flows/{Sid} | 



## FlowsCreate

> StudioV2Flow FlowsCreate(ctx, definition, friendlyName, status, optional)



Create a Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definition** | [**string**](string.md)| JSON representation of flow definition. | 
**friendlyName** | **string**| The string that you assigned to describe the Flow. | 
**status** | **string**| The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;. | 
 **optional** | ***FlowsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FlowsCreateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **commitMessage** | **optional.**| Description on change made in the revision. | 

### Return type

[**StudioV2Flow**](studio.v2.flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsDelete

> FlowsDelete(ctx, sid)



Delete a specific Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Flow resource to delete. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsRead

> StudioV2Flow FlowsRead(ctx, sid)



Retrieve a specific Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Flow resource to fetch. | 

### Return type

[**StudioV2Flow**](studio.v2.flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowsUpdate

> StudioV2Flow FlowsUpdate(ctx, sid, status, optional)



Update a Flow.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Flow resource to fetch. | 
**status** | **string**| The status of the Flow. Can be: &#x60;draft&#x60; or &#x60;published&#x60;. | 
 **optional** | ***FlowsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FlowsUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commitMessage** | **optional.**| Description on change made in the revision. | 
 **definition** | [**optional.Interface of string**](string.md)| JSON representation of flow definition. | 
 **friendlyName** | **optional.**| The string that you assigned to describe the Flow. | 

### Return type

[**StudioV2Flow**](studio.v2.flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

