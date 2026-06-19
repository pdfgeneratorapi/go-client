# \ConversionAPI

All URIs are relative to *https://us1.pdfgeneratorapi.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertHTML2PDF**](ConversionAPI.md#ConvertHTML2PDF) | **Post** /conversion/html2pdf | HTML to PDF
[**ConvertPDF2Image**](ConversionAPI.md#ConvertPDF2Image) | **Post** /conversion/pdf2image | PDF to Image
[**ConvertURL2PDF**](ConversionAPI.md#ConvertURL2PDF) | **Post** /conversion/url2pdf | URL to PDF



## ConvertHTML2PDF

> InlineObject9 ConvertHTML2PDF(ctx).ConvertHTML2PDFRequest(convertHTML2PDFRequest).Execute()

HTML to PDF



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
	convertHTML2PDFRequest := *openapiclient.NewConvertHTML2PDFRequest() // ConvertHTML2PDFRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversionAPI.ConvertHTML2PDF(context.Background()).ConvertHTML2PDFRequest(convertHTML2PDFRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversionAPI.ConvertHTML2PDF``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertHTML2PDF`: InlineObject9
	fmt.Fprintf(os.Stdout, "Response from `ConversionAPI.ConvertHTML2PDF`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertHTML2PDFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **convertHTML2PDFRequest** | [**ConvertHTML2PDFRequest**](ConvertHTML2PDFRequest.md) |  | 

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


## ConvertPDF2Image

> InlineObject10 ConvertPDF2Image(ctx).ConvertPDF2ImageRequest(convertPDF2ImageRequest).Execute()

PDF to Image



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
	convertPDF2ImageRequest := *openapiclient.NewConvertPDF2ImageRequest() // ConvertPDF2ImageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversionAPI.ConvertPDF2Image(context.Background()).ConvertPDF2ImageRequest(convertPDF2ImageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversionAPI.ConvertPDF2Image``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertPDF2Image`: InlineObject10
	fmt.Fprintf(os.Stdout, "Response from `ConversionAPI.ConvertPDF2Image`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertPDF2ImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **convertPDF2ImageRequest** | [**ConvertPDF2ImageRequest**](ConvertPDF2ImageRequest.md) |  | 

### Return type

[**InlineObject10**](InlineObject10.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConvertURL2PDF

> InlineObject9 ConvertURL2PDF(ctx).ConvertURL2PDFRequest(convertURL2PDFRequest).Execute()

URL to PDF



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
	convertURL2PDFRequest := *openapiclient.NewConvertURL2PDFRequest() // ConvertURL2PDFRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConversionAPI.ConvertURL2PDF(context.Background()).ConvertURL2PDFRequest(convertURL2PDFRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConversionAPI.ConvertURL2PDF``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertURL2PDF`: InlineObject9
	fmt.Fprintf(os.Stdout, "Response from `ConversionAPI.ConvertURL2PDF`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertURL2PDFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **convertURL2PDFRequest** | [**ConvertURL2PDFRequest**](ConvertURL2PDFRequest.md) |  | 

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

