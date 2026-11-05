# \FontsAPI

All URIs are relative to *https://us1.pdfgeneratorapi.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFonts**](FontsAPI.md#GetFonts) | **Get** /fonts | Get fonts



## GetFonts

> GetFonts200Response GetFonts(ctx).Covers(covers).Execute()

Get fonts



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
	covers := "€₹✓" // string | Text whose every character a font must have a glyph for, in its regular style, to be returned. Control characters are ignored. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FontsAPI.GetFonts(context.Background()).Covers(covers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FontsAPI.GetFonts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFonts`: GetFonts200Response
	fmt.Fprintf(os.Stdout, "Response from `FontsAPI.GetFonts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFontsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **covers** | **string** | Text whose every character a font must have a glyph for, in its regular style, to be returned. Control characters are ignored. | 

### Return type

[**GetFonts200Response**](GetFonts200Response.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

