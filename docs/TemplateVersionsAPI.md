# \TemplateVersionsAPI

All URIs are relative to *https://us1.pdfgeneratorapi.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTemplateVersion**](TemplateVersionsAPI.md#DeleteTemplateVersion) | **Delete** /templates/{templateId}/versions/{templateVersion} | Delete template version
[**GetTemplateVersion**](TemplateVersionsAPI.md#GetTemplateVersion) | **Get** /templates/{templateId}/versions/{templateVersion} | Get template version
[**ListTemplateVersions**](TemplateVersionsAPI.md#ListTemplateVersions) | **Get** /templates/{templateId}/versions | List template versions
[**PromoteTemplateVersion**](TemplateVersionsAPI.md#PromoteTemplateVersion) | **Put** /templates/{templateId}/versions/{templateVersion}/promote | Promote template version



## DeleteTemplateVersion

> DeleteTemplateVersion(ctx, templateId, templateVersion).Execute()

Delete template version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pdfgeneratorapi/go-client/v4"
)

func main() {
	templateId := int32(19375) // int32 | Template unique identifier
	templateVersion := int32(56) // int32 | Unique ID of the template version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TemplateVersionsAPI.DeleteTemplateVersion(context.Background(), templateId, templateVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateVersionsAPI.DeleteTemplateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | Template unique identifier | 
**templateVersion** | **int32** | Unique ID of the template version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateVersion

> InlineObject16 GetTemplateVersion(ctx, templateId, templateVersion).Execute()

Get template version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pdfgeneratorapi/go-client/v4"
)

func main() {
	templateId := int32(19375) // int32 | Template unique identifier
	templateVersion := int32(56) // int32 | Unique ID of the template version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateVersionsAPI.GetTemplateVersion(context.Background(), templateId, templateVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateVersionsAPI.GetTemplateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateVersion`: InlineObject16
	fmt.Fprintf(os.Stdout, "Response from `TemplateVersionsAPI.GetTemplateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | Template unique identifier | 
**templateVersion** | **int32** | Unique ID of the template version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineObject16**](InlineObject16.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplateVersions

> TemplateVersionCollection ListTemplateVersions(ctx, templateId).PerPage(perPage).Page(page).Execute()

List template versions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pdfgeneratorapi/go-client/v4"
)

func main() {
	templateId := int32(19375) // int32 | Template unique identifier
	perPage := int32(56) // int32 | Number of items per page. (optional)
	page := int32(56) // int32 | Page number to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateVersionsAPI.ListTemplateVersions(context.Background(), templateId).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateVersionsAPI.ListTemplateVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTemplateVersions`: TemplateVersionCollection
	fmt.Fprintf(os.Stdout, "Response from `TemplateVersionsAPI.ListTemplateVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | Template unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplateVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | Number of items per page. | 
 **page** | **int32** | Page number to return. | 

### Return type

[**TemplateVersionCollection**](TemplateVersionCollection.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromoteTemplateVersion

> PromoteTemplateVersion200Response PromoteTemplateVersion(ctx, templateId, templateVersion).Execute()

Promote template version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pdfgeneratorapi/go-client/v4"
)

func main() {
	templateId := int32(19375) // int32 | Template unique identifier
	templateVersion := int32(56) // int32 | Unique ID of the template version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateVersionsAPI.PromoteTemplateVersion(context.Background(), templateId, templateVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateVersionsAPI.PromoteTemplateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromoteTemplateVersion`: PromoteTemplateVersion200Response
	fmt.Fprintf(os.Stdout, "Response from `TemplateVersionsAPI.PromoteTemplateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | Template unique identifier | 
**templateVersion** | **int32** | Unique ID of the template version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromoteTemplateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PromoteTemplateVersion200Response**](PromoteTemplateVersion200Response.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

