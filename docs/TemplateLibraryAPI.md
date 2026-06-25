# \TemplateLibraryAPI

All URIs are relative to *https://us1.pdfgeneratorapi.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTemplateLibrary**](TemplateLibraryAPI.md#GetTemplateLibrary) | **Get** /templates/library | Get template library
[**GetTemplateLibraryItem**](TemplateLibraryAPI.md#GetTemplateLibraryItem) | **Get** /templates/library/{publicId} | Open template from the library



## GetTemplateLibrary

> GetTemplateLibrary200Response GetTemplateLibrary(ctx).Tags(tags).Execute()

Get template library



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
	tags := "tags_example" // string | Filter template by tags (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateLibraryAPI.GetTemplateLibrary(context.Background()).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateLibraryAPI.GetTemplateLibrary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateLibrary`: GetTemplateLibrary200Response
	fmt.Fprintf(os.Stdout, "Response from `TemplateLibraryAPI.GetTemplateLibrary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateLibraryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | **string** | Filter template by tags | 

### Return type

[**GetTemplateLibrary200Response**](GetTemplateLibrary200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateLibraryItem

> InlineObject18 GetTemplateLibraryItem(ctx, publicId).Execute()

Open template from the library



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
	publicId := "bac8381bce1982e5f6957a0f52371336" // string | Resource public id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplateLibraryAPI.GetTemplateLibraryItem(context.Background(), publicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplateLibraryAPI.GetTemplateLibraryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateLibraryItem`: InlineObject18
	fmt.Fprintf(os.Stdout, "Response from `TemplateLibraryAPI.GetTemplateLibraryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | Resource public id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateLibraryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineObject18**](InlineObject18.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

