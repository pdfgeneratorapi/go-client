# \EInvoicesAPI

All URIs are relative to *https://us1.pdfgeneratorapi.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEInvoice**](EInvoicesAPI.md#CreateEInvoice) | **Post** /einvoice | Create eInvoice
[**CreateFacturXEInvoice**](EInvoicesAPI.md#CreateFacturXEInvoice) | **Post** /einvoice/facturx | Create Factur-X eInvoice
[**CreateXRechnungEInvoice**](EInvoicesAPI.md#CreateXRechnungEInvoice) | **Post** /einvoice/xrechnung | Create XRechnung eInvoice
[**GetEInvoiceSchema**](EInvoicesAPI.md#GetEInvoiceSchema) | **Get** /einvoice/schema | Get schema



## CreateEInvoice

> InlineObject CreateEInvoice(ctx).CreateEInvoiceRequest(createEInvoiceRequest).Execute()

Create eInvoice



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
	createEInvoiceRequest := *openapiclient.NewCreateEInvoiceRequest(map[string]interface{}(123)) // CreateEInvoiceRequest | eInvoice conversion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EInvoicesAPI.CreateEInvoice(context.Background()).CreateEInvoiceRequest(createEInvoiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EInvoicesAPI.CreateEInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEInvoice`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `EInvoicesAPI.CreateEInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEInvoiceRequest** | [**CreateEInvoiceRequest**](CreateEInvoiceRequest.md) | eInvoice conversion | 

### Return type

[**InlineObject**](InlineObject.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFacturXEInvoice

> InlineObject9 CreateFacturXEInvoice(ctx).CreateFacturXEInvoiceRequest(createFacturXEInvoiceRequest).Execute()

Create Factur-X eInvoice



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
	createFacturXEInvoiceRequest := *openapiclient.NewCreateFacturXEInvoiceRequest(*openapiclient.NewCreateFacturXEInvoiceRequestTemplate()) // CreateFacturXEInvoiceRequest | eInvoice conversion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EInvoicesAPI.CreateFacturXEInvoice(context.Background()).CreateFacturXEInvoiceRequest(createFacturXEInvoiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EInvoicesAPI.CreateFacturXEInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFacturXEInvoice`: InlineObject9
	fmt.Fprintf(os.Stdout, "Response from `EInvoicesAPI.CreateFacturXEInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFacturXEInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFacturXEInvoiceRequest** | [**CreateFacturXEInvoiceRequest**](CreateFacturXEInvoiceRequest.md) | eInvoice conversion | 

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


## CreateXRechnungEInvoice

> InlineObject CreateXRechnungEInvoice(ctx).CreateEInvoiceRequest(createEInvoiceRequest).Execute()

Create XRechnung eInvoice



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
	createEInvoiceRequest := *openapiclient.NewCreateEInvoiceRequest(map[string]interface{}(123)) // CreateEInvoiceRequest | eInvoice conversion

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EInvoicesAPI.CreateXRechnungEInvoice(context.Background()).CreateEInvoiceRequest(createEInvoiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EInvoicesAPI.CreateXRechnungEInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateXRechnungEInvoice`: InlineObject
	fmt.Fprintf(os.Stdout, "Response from `EInvoicesAPI.CreateXRechnungEInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateXRechnungEInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEInvoiceRequest** | [**CreateEInvoiceRequest**](CreateEInvoiceRequest.md) | eInvoice conversion | 

### Return type

[**InlineObject**](InlineObject.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEInvoiceSchema

> map[string]interface{} GetEInvoiceSchema(ctx).Execute()

Get schema



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EInvoicesAPI.GetEInvoiceSchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EInvoicesAPI.GetEInvoiceSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEInvoiceSchema`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EInvoicesAPI.GetEInvoiceSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEInvoiceSchemaRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

