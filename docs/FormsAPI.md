# \FormsAPI

All URIs are relative to *https://us1.pdfgeneratorapi.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateForm**](FormsAPI.md#CreateForm) | **Post** /forms | Create form
[**DeleteForm**](FormsAPI.md#DeleteForm) | **Delete** /forms/{formId} | Delete form
[**GetForm**](FormsAPI.md#GetForm) | **Get** /forms/{formId} | Get form
[**GetForms**](FormsAPI.md#GetForms) | **Get** /forms | Get forms
[**ImportForm**](FormsAPI.md#ImportForm) | **Post** /forms/import | Import Form
[**OpenFormBuilder**](FormsAPI.md#OpenFormBuilder) | **Post** /forms/open | Open new form builder
[**OpenFormBuilderForExistingForm**](FormsAPI.md#OpenFormBuilderForExistingForm) | **Post** /forms/{formId}/open | Open existing form builder
[**ShareForm**](FormsAPI.md#ShareForm) | **Post** /forms/{formId}/share | Share form
[**UpdateForm**](FormsAPI.md#UpdateForm) | **Put** /forms/{formId} | Update form



## CreateForm

> InlineObject19 CreateForm(ctx).FormConfigurationNew(formConfigurationNew).Execute()

Create form



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
	formConfigurationNew := *openapiclient.NewFormConfigurationNew() // FormConfigurationNew | Form configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormsAPI.CreateForm(context.Background()).FormConfigurationNew(formConfigurationNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.CreateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateForm`: InlineObject19
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.CreateForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **formConfigurationNew** | [**FormConfigurationNew**](FormConfigurationNew.md) | Form configuration | 

### Return type

[**InlineObject19**](InlineObject19.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteForm

> DeleteForm(ctx, formId).Execute()

Delete form



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
	formId := int32(1) // int32 | Form unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FormsAPI.DeleteForm(context.Background(), formId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.DeleteForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **int32** | Form unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFormRequest struct via the builder pattern


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


## GetForm

> InlineObject19 GetForm(ctx, formId).Execute()

Get form



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
	formId := int32(1) // int32 | Form unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormsAPI.GetForm(context.Background(), formId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.GetForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForm`: InlineObject19
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.GetForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **int32** | Form unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineObject19**](InlineObject19.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForms

> InlineObject6 GetForms(ctx).Page(page).PerPage(perPage).Execute()

Get forms



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormsAPI.GetForms(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.GetForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForms`: InlineObject6
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.GetForms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Pagination: page to return | [default to 1]
 **perPage** | **int32** | Pagination: How many records to return per page | [default to 15]

### Return type

[**InlineObject6**](InlineObject6.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportForm

> InlineObject19 ImportForm(ctx).ImportFormRequest(importFormRequest).Execute()

Import Form



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
	importFormRequest := openapiclient.importForm_request{ImportFormBase64: openapiclient.NewImportFormBase64("FileBase64_example")} // ImportFormRequest | Import editable PDF via URL or base64 string as form

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormsAPI.ImportForm(context.Background()).ImportFormRequest(importFormRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.ImportForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportForm`: InlineObject19
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.ImportForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importFormRequest** | [**ImportFormRequest**](ImportFormRequest.md) | Import editable PDF via URL or base64 string as form | 

### Return type

[**InlineObject19**](InlineObject19.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenFormBuilder

> InlineObject21 OpenFormBuilder(ctx).Execute()

Open new form builder



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
	resp, r, err := apiClient.FormsAPI.OpenFormBuilder(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.OpenFormBuilder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenFormBuilder`: InlineObject21
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.OpenFormBuilder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenFormBuilderRequest struct via the builder pattern


### Return type

[**InlineObject21**](InlineObject21.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenFormBuilderForExistingForm

> InlineObject21 OpenFormBuilderForExistingForm(ctx, formId).Execute()

Open existing form builder



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
	formId := int32(1) // int32 | Form unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormsAPI.OpenFormBuilderForExistingForm(context.Background(), formId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.OpenFormBuilderForExistingForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenFormBuilderForExistingForm`: InlineObject21
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.OpenFormBuilderForExistingForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **int32** | Form unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenFormBuilderForExistingFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineObject21**](InlineObject21.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareForm

> InlineObject20 ShareForm(ctx, formId).Execute()

Share form



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
	formId := int32(1) // int32 | Form unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormsAPI.ShareForm(context.Background(), formId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.ShareForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShareForm`: InlineObject20
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.ShareForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **int32** | Form unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineObject20**](InlineObject20.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateForm

> InlineObject19 UpdateForm(ctx, formId).FormConfigurationNew(formConfigurationNew).Execute()

Update form



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
	formId := int32(1) // int32 | Form unique identifier
	formConfigurationNew := *openapiclient.NewFormConfigurationNew() // FormConfigurationNew | Form configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormsAPI.UpdateForm(context.Background(), formId).FormConfigurationNew(formConfigurationNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.UpdateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateForm`: InlineObject19
	fmt.Fprintf(os.Stdout, "Response from `FormsAPI.UpdateForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formId** | **int32** | Form unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **formConfigurationNew** | [**FormConfigurationNew**](FormConfigurationNew.md) | Form configuration | 

### Return type

[**InlineObject19**](InlineObject19.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

