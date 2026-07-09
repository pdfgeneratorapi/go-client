# \DocumentsAPI

All URIs are relative to *https://us1.pdfgeneratorapi.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDocument**](DocumentsAPI.md#DeleteDocument) | **Delete** /documents/{publicId} | Delete document
[**GenerateDocument**](DocumentsAPI.md#GenerateDocument) | **Post** /documents/generate | Generate document
[**GenerateDocumentAsynchronous**](DocumentsAPI.md#GenerateDocumentAsynchronous) | **Post** /documents/generate/async | Generate document (async)
[**GenerateDocumentBatch**](DocumentsAPI.md#GenerateDocumentBatch) | **Post** /documents/generate/batch | Generate document (batch)
[**GenerateDocumentBatchAsynchronous**](DocumentsAPI.md#GenerateDocumentBatchAsynchronous) | **Post** /documents/generate/batch/async | Generate document (batch + async)
[**GenerateViewerUrl**](DocumentsAPI.md#GenerateViewerUrl) | **Post** /documents/{publicId} | Get document with prefill
[**GetAsyncJobStatus**](DocumentsAPI.md#GetAsyncJobStatus) | **Get** /documents/async/{jobId} | Get job status
[**GetDocument**](DocumentsAPI.md#GetDocument) | **Get** /documents/{publicId} | Get document
[**GetDocumentActions**](DocumentsAPI.md#GetDocumentActions) | **Get** /documents/{publicId}/actions | Get document actions
[**GetDocumentVersions**](DocumentsAPI.md#GetDocumentVersions) | **Get** /documents/{publicId}/versions | Get document versions
[**GetDocuments**](DocumentsAPI.md#GetDocuments) | **Get** /documents | Get documents
[**StoreDocument**](DocumentsAPI.md#StoreDocument) | **Post** /documents | Store document



## DeleteDocument

> DeleteDocument(ctx, publicId).Execute()

Delete document



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
	r, err := apiClient.DocumentsAPI.DeleteDocument(context.Background(), publicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DeleteDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | Resource public id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDocumentRequest struct via the builder pattern


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


## GenerateDocument

> InlineObject9 GenerateDocument(ctx).GenerateDocumentRequest(generateDocumentRequest).Execute()

Generate document



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
	generateDocumentRequest := *openapiclient.NewGenerateDocumentRequest() // GenerateDocumentRequest | Request parameters, including template id, data and formats.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.GenerateDocument(context.Background()).GenerateDocumentRequest(generateDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GenerateDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateDocument`: InlineObject9
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GenerateDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateDocumentRequest** | [**GenerateDocumentRequest**](GenerateDocumentRequest.md) | Request parameters, including template id, data and formats. | 

### Return type

[**InlineObject9**](InlineObject9.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateDocumentAsynchronous

> InlineObject22 GenerateDocumentAsynchronous(ctx).GenerateDocumentAsynchronousRequest(generateDocumentAsynchronousRequest).Execute()

Generate document (async)



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
	generateDocumentAsynchronousRequest := *openapiclient.NewGenerateDocumentAsynchronousRequest() // GenerateDocumentAsynchronousRequest | Request parameters, including template id, data and formats.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.GenerateDocumentAsynchronous(context.Background()).GenerateDocumentAsynchronousRequest(generateDocumentAsynchronousRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GenerateDocumentAsynchronous``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateDocumentAsynchronous`: InlineObject22
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GenerateDocumentAsynchronous`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDocumentAsynchronousRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateDocumentAsynchronousRequest** | [**GenerateDocumentAsynchronousRequest**](GenerateDocumentAsynchronousRequest.md) | Request parameters, including template id, data and formats. | 

### Return type

[**InlineObject22**](InlineObject22.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateDocumentBatch

> InlineObject9 GenerateDocumentBatch(ctx).GenerateDocumentBatchRequest(generateDocumentBatchRequest).Execute()

Generate document (batch)



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
	generateDocumentBatchRequest := *openapiclient.NewGenerateDocumentBatchRequest() // GenerateDocumentBatchRequest | Request parameters, including template id, data and formats.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.GenerateDocumentBatch(context.Background()).GenerateDocumentBatchRequest(generateDocumentBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GenerateDocumentBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateDocumentBatch`: InlineObject9
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GenerateDocumentBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDocumentBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateDocumentBatchRequest** | [**GenerateDocumentBatchRequest**](GenerateDocumentBatchRequest.md) | Request parameters, including template id, data and formats. | 

### Return type

[**InlineObject9**](InlineObject9.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateDocumentBatchAsynchronous

> InlineObject22 GenerateDocumentBatchAsynchronous(ctx).GenerateDocumentBatchAsynchronousRequest(generateDocumentBatchAsynchronousRequest).Execute()

Generate document (batch + async)



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
	generateDocumentBatchAsynchronousRequest := *openapiclient.NewGenerateDocumentBatchAsynchronousRequest() // GenerateDocumentBatchAsynchronousRequest | Request parameters, including template id, data and formats.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.GenerateDocumentBatchAsynchronous(context.Background()).GenerateDocumentBatchAsynchronousRequest(generateDocumentBatchAsynchronousRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GenerateDocumentBatchAsynchronous``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateDocumentBatchAsynchronous`: InlineObject22
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GenerateDocumentBatchAsynchronous`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDocumentBatchAsynchronousRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateDocumentBatchAsynchronousRequest** | [**GenerateDocumentBatchAsynchronousRequest**](GenerateDocumentBatchAsynchronousRequest.md) | Request parameters, including template id, data and formats. | 

### Return type

[**InlineObject22**](InlineObject22.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateViewerUrl

> GenerateViewerUrl200Response GenerateViewerUrl(ctx, publicId).GenerateViewerUrlRequest(generateViewerUrlRequest).Execute()

Get document with prefill



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
	generateViewerUrlRequest := *openapiclient.NewGenerateViewerUrlRequest() // GenerateViewerUrlRequest | Optional response format and viewer prefill data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.GenerateViewerUrl(context.Background(), publicId).GenerateViewerUrlRequest(generateViewerUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GenerateViewerUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateViewerUrl`: GenerateViewerUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GenerateViewerUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | Resource public id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateViewerUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateViewerUrlRequest** | [**GenerateViewerUrlRequest**](GenerateViewerUrlRequest.md) | Optional response format and viewer prefill data. | 

### Return type

[**GenerateViewerUrl200Response**](GenerateViewerUrl200Response.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsyncJobStatus

> InlineObject13 GetAsyncJobStatus(ctx, jobId).Execute()

Get job status



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
	jobId := "968b8a3a-e8ac-49cc-a670-25db545813ed" // string | Job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.GetAsyncJobStatus(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetAsyncJobStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAsyncJobStatus`: InlineObject13
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetAsyncJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineObject13**](InlineObject13.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocument

> InlineObject11 GetDocument(ctx, publicId).Execute()

Get document



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
	resp, r, err := apiClient.DocumentsAPI.GetDocument(context.Background(), publicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocument`: InlineObject11
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | Resource public id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineObject11**](InlineObject11.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentActions

> InlineObject17 GetDocumentActions(ctx, publicId).Execute()

Get document actions



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
	resp, r, err := apiClient.DocumentsAPI.GetDocumentActions(context.Background(), publicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetDocumentActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocumentActions`: InlineObject17
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetDocumentActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | Resource public id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineObject17**](InlineObject17.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentVersions

> InlineObject16 GetDocumentVersions(ctx, publicId).Execute()

Get document versions



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
	resp, r, err := apiClient.DocumentsAPI.GetDocumentVersions(context.Background(), publicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetDocumentVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocumentVersions`: InlineObject16
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetDocumentVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | Resource public id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentVersionsRequest struct via the builder pattern


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


## GetDocuments

> InlineObject15 GetDocuments(ctx).TemplateId(templateId).StartDate(startDate).EndDate(endDate).Page(page).PerPage(perPage).Execute()

Get documents



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
	templateId := int32(19375) // int32 | Template unique identifier (optional)
	startDate := "2022-08-01 12:00:00" // string | Start date. Format: Y-m-d H:i:s (optional)
	endDate := "2022-08-05 12:00:00" // string | End date. Format: Y-m-d H:i:s. Defaults to current timestamp (optional)
	page := int32(1) // int32 | Pagination: page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Pagination: How many records to return per page (optional) (default to 15)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.GetDocuments(context.Background()).TemplateId(templateId).StartDate(startDate).EndDate(endDate).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocuments`: InlineObject15
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **int32** | Template unique identifier | 
 **startDate** | **string** | Start date. Format: Y-m-d H:i:s | 
 **endDate** | **string** | End date. Format: Y-m-d H:i:s. Defaults to current timestamp | 
 **page** | **int32** | Pagination: page to return | [default to 1]
 **perPage** | **int32** | Pagination: How many records to return per page | [default to 15]

### Return type

[**InlineObject15**](InlineObject15.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreDocument

> InlineObject11 StoreDocument(ctx).StoreDocumentRequest(storeDocumentRequest).Execute()

Store document



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
	storeDocumentRequest := *openapiclient.NewStoreDocumentRequest() // StoreDocumentRequest | Document source and optional metadata. Exactly one of `file_base64` or `file_url` is required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.StoreDocument(context.Background()).StoreDocumentRequest(storeDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.StoreDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreDocument`: InlineObject11
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.StoreDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeDocumentRequest** | [**StoreDocumentRequest**](StoreDocumentRequest.md) | Document source and optional metadata. Exactly one of &#x60;file_base64&#x60; or &#x60;file_url&#x60; is required. | 

### Return type

[**InlineObject11**](InlineObject11.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

