# \AssetsAPI

All URIs are relative to *https://us1.pdfgeneratorapi.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateQRCode**](AssetsAPI.md#GenerateQRCode) | **Post** /assets/qrcode | Generate QR Code



## GenerateQRCode

> *os.File GenerateQRCode(ctx).GenerateQRCodeRequest(generateQRCodeRequest).Execute()

Generate QR Code



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
	generateQRCodeRequest := *openapiclient.NewGenerateQRCodeRequest("My QR Code") // GenerateQRCodeRequest | QR Code configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.GenerateQRCode(context.Background()).GenerateQRCodeRequest(generateQRCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GenerateQRCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateQRCode`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GenerateQRCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateQRCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateQRCodeRequest** | [**GenerateQRCodeRequest**](GenerateQRCodeRequest.md) | QR Code configuration | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

