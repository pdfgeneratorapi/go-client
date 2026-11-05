# \WorkspacesAPI

All URIs are relative to *https://us1.pdfgeneratorapi.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkspace**](WorkspacesAPI.md#CreateWorkspace) | **Post** /workspaces | Create workspace
[**DeleteWorkspace**](WorkspacesAPI.md#DeleteWorkspace) | **Delete** /workspaces/{workspaceIdentifier} | Delete workspace
[**GetWorkspace**](WorkspacesAPI.md#GetWorkspace) | **Get** /workspaces/{workspaceIdentifier} | Get workspace
[**GetWorkspaces**](WorkspacesAPI.md#GetWorkspaces) | **Get** /workspaces | Get workspaces



## CreateWorkspace

> InlineObject7 CreateWorkspace(ctx).CreateWorkspaceRequest(createWorkspaceRequest).Execute()

Create workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pdfgeneratorapi/go-client/v8"
)

func main() {
	createWorkspaceRequest := *openapiclient.NewCreateWorkspaceRequest("demo.example@actualreports.com") // CreateWorkspaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.CreateWorkspace(context.Background()).CreateWorkspaceRequest(createWorkspaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.CreateWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspace`: InlineObject7
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.CreateWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkspaceRequest** | [**CreateWorkspaceRequest**](CreateWorkspaceRequest.md) |  | 

### Return type

[**InlineObject7**](InlineObject7.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkspace

> DeleteWorkspace(ctx, workspaceIdentifier).Execute()

Delete workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pdfgeneratorapi/go-client/v8"
)

func main() {
	workspaceIdentifier := "demo.example@actualreports.com" // string | Workspace identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspacesAPI.DeleteWorkspace(context.Background(), workspaceIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.DeleteWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceIdentifier** | **string** | Workspace identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceRequest struct via the builder pattern


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


## GetWorkspace

> InlineObject7 GetWorkspace(ctx, workspaceIdentifier).Execute()

Get workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pdfgeneratorapi/go-client/v8"
)

func main() {
	workspaceIdentifier := "demo.example@actualreports.com" // string | Workspace identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.GetWorkspace(context.Background(), workspaceIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.GetWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspace`: InlineObject7
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.GetWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceIdentifier** | **string** | Workspace identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineObject7**](InlineObject7.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspaces

> InlineObject5 GetWorkspaces(ctx).Page(page).PerPage(perPage).OrderBy(orderBy).OrderDirection(orderDirection).Activated(activated).Search(search).Execute()

Get workspaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pdfgeneratorapi/go-client/v8"
)

func main() {
	page := int32(1) // int32 | Pagination: page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Pagination: How many records to return per page (optional) (default to 15)
	orderBy := "identifier" // string | Field to sort the workspaces by (optional) (default to "id")
	orderDirection := "asc" // string | Sort direction (optional) (default to "desc")
	activated := int32(1) // int32 | Returns only active (`1`) or deactivated (`0`) workspaces (optional)
	search := "demo.example" // string | Returns workspaces whose ID, name or identifier contains the search value (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.GetWorkspaces(context.Background()).Page(page).PerPage(perPage).OrderBy(orderBy).OrderDirection(orderDirection).Activated(activated).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.GetWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaces`: InlineObject5
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.GetWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Pagination: page to return | [default to 1]
 **perPage** | **int32** | Pagination: How many records to return per page | [default to 15]
 **orderBy** | **string** | Field to sort the workspaces by | [default to &quot;id&quot;]
 **orderDirection** | **string** | Sort direction | [default to &quot;desc&quot;]
 **activated** | **int32** | Returns only active (&#x60;1&#x60;) or deactivated (&#x60;0&#x60;) workspaces | 
 **search** | **string** | Returns workspaces whose ID, name or identifier contains the search value | 

### Return type

[**InlineObject5**](InlineObject5.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

