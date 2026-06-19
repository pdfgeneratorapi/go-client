# \ServicesAPI

All URIs are relative to *https://us1.pdfgeneratorapi.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWatermark**](ServicesAPI.md#AddWatermark) | **Post** /pdfservices/watermark | Add watermark
[**DecryptDocument**](ServicesAPI.md#DecryptDocument) | **Post** /pdfservices/decrypt | Decrypt document
[**EncryptDocument**](ServicesAPI.md#EncryptDocument) | **Post** /pdfservices/encrypt | Encrypt document
[**ExtractFormFields**](ServicesAPI.md#ExtractFormFields) | **Post** /pdfservices/form/fields | Extract form fields
[**FillFormFields**](ServicesAPI.md#FillFormFields) | **Post** /pdfservices/form/fill | Fill form fields
[**MakeAccessible**](ServicesAPI.md#MakeAccessible) | **Post** /pdfservices/make-accessible | Make accessible
[**OptimizeDocument**](ServicesAPI.md#OptimizeDocument) | **Post** /pdfservices/optimize | Optimize document



## AddWatermark

> InlineObject9 AddWatermark(ctx).AddWatermarkRequest(addWatermarkRequest).Execute()

Add watermark



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
	addWatermarkRequest := openapiclient.addWatermark_request{WatermarkBase64: openapiclient.NewWatermarkBase64("FileBase64_example", *openapiclient.NewWatermarkFileUrlWatermark())} // AddWatermarkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.AddWatermark(context.Background()).AddWatermarkRequest(addWatermarkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.AddWatermark``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddWatermark`: InlineObject9
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.AddWatermark`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddWatermarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addWatermarkRequest** | [**AddWatermarkRequest**](AddWatermarkRequest.md) |  | 

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


## DecryptDocument

> InlineObject9 DecryptDocument(ctx).EncryptDocumentRequest(encryptDocumentRequest).Execute()

Decrypt document



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
	encryptDocumentRequest := openapiclient.encryptDocument_request{EncryptAndDecryptBase64: openapiclient.NewEncryptAndDecryptBase64("FileBase64_example", "OwnerPassword_example")} // EncryptDocumentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.DecryptDocument(context.Background()).EncryptDocumentRequest(encryptDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.DecryptDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DecryptDocument`: InlineObject9
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.DecryptDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDecryptDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **encryptDocumentRequest** | [**EncryptDocumentRequest**](EncryptDocumentRequest.md) |  | 

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


## EncryptDocument

> InlineObject9 EncryptDocument(ctx).EncryptDocumentRequest(encryptDocumentRequest).Execute()

Encrypt document



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
	encryptDocumentRequest := openapiclient.encryptDocument_request{EncryptAndDecryptBase64: openapiclient.NewEncryptAndDecryptBase64("FileBase64_example", "OwnerPassword_example")} // EncryptDocumentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.EncryptDocument(context.Background()).EncryptDocumentRequest(encryptDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.EncryptDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EncryptDocument`: InlineObject9
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.EncryptDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEncryptDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **encryptDocumentRequest** | [**EncryptDocumentRequest**](EncryptDocumentRequest.md) |  | 

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


## ExtractFormFields

> InlineObject14 ExtractFormFields(ctx).ExtractFormFieldsRequest(extractFormFieldsRequest).Execute()

Extract form fields



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
	extractFormFieldsRequest := openapiclient.extractFormFields_request{FormFieldsBase64: openapiclient.NewFormFieldsBase64("FileBase64_example")} // ExtractFormFieldsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.ExtractFormFields(context.Background()).ExtractFormFieldsRequest(extractFormFieldsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.ExtractFormFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExtractFormFields`: InlineObject14
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.ExtractFormFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExtractFormFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extractFormFieldsRequest** | [**ExtractFormFieldsRequest**](ExtractFormFieldsRequest.md) |  | 

### Return type

[**InlineObject14**](InlineObject14.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FillFormFields

> InlineObject9 FillFormFields(ctx).FillFormFieldsRequest(fillFormFieldsRequest).Execute()

Fill form fields



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
	fillFormFieldsRequest := openapiclient.fillFormFields_request{FormFillBase64: openapiclient.NewFormFillBase64("FileBase64_example", map[string]interface{}({"firstName":"John","lastName":"Doe","email":"john.doe@example.com"}))} // FillFormFieldsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.FillFormFields(context.Background()).FillFormFieldsRequest(fillFormFieldsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.FillFormFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FillFormFields`: InlineObject9
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.FillFormFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFillFormFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fillFormFieldsRequest** | [**FillFormFieldsRequest**](FillFormFieldsRequest.md) |  | 

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


## MakeAccessible

> InlineObject9 MakeAccessible(ctx).MakeAccessibleRequest(makeAccessibleRequest).Execute()

Make accessible



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
	makeAccessibleRequest := openapiclient.makeAccessible_request{MakeAccessibleBase64: openapiclient.NewMakeAccessibleBase64("FileBase64_example")} // MakeAccessibleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.MakeAccessible(context.Background()).MakeAccessibleRequest(makeAccessibleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.MakeAccessible``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MakeAccessible`: InlineObject9
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.MakeAccessible`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMakeAccessibleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **makeAccessibleRequest** | [**MakeAccessibleRequest**](MakeAccessibleRequest.md) |  | 

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


## OptimizeDocument

> InlineObject12 OptimizeDocument(ctx).OptimizeDocumentRequest(optimizeDocumentRequest).Execute()

Optimize document



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
	optimizeDocumentRequest := openapiclient.optimizeDocument_request{OptimizeBase64: openapiclient.NewOptimizeBase64("FileBase64_example")} // OptimizeDocumentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServicesAPI.OptimizeDocument(context.Background()).OptimizeDocumentRequest(optimizeDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServicesAPI.OptimizeDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptimizeDocument`: InlineObject12
	fmt.Fprintf(os.Stdout, "Response from `ServicesAPI.OptimizeDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptimizeDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optimizeDocumentRequest** | [**OptimizeDocumentRequest**](OptimizeDocumentRequest.md) |  | 

### Return type

[**InlineObject12**](InlineObject12.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

