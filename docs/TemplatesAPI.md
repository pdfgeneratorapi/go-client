# \TemplatesAPI

All URIs are relative to *https://us1.pdfgeneratorapi.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyTemplate**](TemplatesAPI.md#CopyTemplate) | **Post** /templates/{templateId}/copy | Copy template
[**CreateTemplate**](TemplatesAPI.md#CreateTemplate) | **Post** /templates | Create template
[**DeleteTemplate**](TemplatesAPI.md#DeleteTemplate) | **Delete** /templates/{templateId} | Delete template
[**GetTemplate**](TemplatesAPI.md#GetTemplate) | **Get** /templates/{templateId} | Get template
[**GetTemplateData**](TemplatesAPI.md#GetTemplateData) | **Get** /templates/{templateId}/data | Get template data fields
[**GetTemplateSchema**](TemplatesAPI.md#GetTemplateSchema) | **Get** /templates/schema | Get schema
[**GetTemplates**](TemplatesAPI.md#GetTemplates) | **Get** /templates | Get templates
[**ImportTemplate**](TemplatesAPI.md#ImportTemplate) | **Post** /templates/import | Import template
[**OpenEditor**](TemplatesAPI.md#OpenEditor) | **Post** /templates/{templateId}/editor | Open editor
[**UpdateTemplate**](TemplatesAPI.md#UpdateTemplate) | **Put** /templates/{templateId} | Update template
[**ValidateTemplate**](TemplatesAPI.md#ValidateTemplate) | **Post** /templates/validate | Validate template



## CopyTemplate

> InlineObject18 CopyTemplate(ctx, templateId).CopyTemplateRequest(copyTemplateRequest).Execute()

Copy template



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
	templateId := int32(19375) // int32 | Template unique identifier
	copyTemplateRequest := *openapiclient.NewCopyTemplateRequest() // CopyTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.CopyTemplate(context.Background(), templateId).CopyTemplateRequest(copyTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CopyTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyTemplate`: InlineObject18
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.CopyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | Template unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **copyTemplateRequest** | [**CopyTemplateRequest**](CopyTemplateRequest.md) |  | 

### Return type

[**InlineObject18**](InlineObject18.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplate

> InlineObject18 CreateTemplate(ctx).TemplateDefinitionNew(templateDefinitionNew).Execute()

Create template



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
	templateDefinitionNew := *openapiclient.NewTemplateDefinitionNew("Invoice template") // TemplateDefinitionNew | Template configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.CreateTemplate(context.Background()).TemplateDefinitionNew(templateDefinitionNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.CreateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplate`: InlineObject18
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.CreateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateDefinitionNew** | [**TemplateDefinitionNew**](TemplateDefinitionNew.md) | Template configuration | 

### Return type

[**InlineObject18**](InlineObject18.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplate

> DeleteTemplate(ctx, templateId).Execute()

Delete template



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
	templateId := int32(19375) // int32 | Template unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TemplatesAPI.DeleteTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.DeleteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | Template unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateRequest struct via the builder pattern


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


## GetTemplate

> InlineObject18 GetTemplate(ctx, templateId).Execute()

Get template



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
	templateId := int32(19375) // int32 | Template unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplate`: InlineObject18
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | Template unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineObject18**](InlineObject18.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateData

> InlineObject2 GetTemplateData(ctx, templateId).Execute()

Get template data fields



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
	templateId := int32(19375) // int32 | Template unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetTemplateData(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetTemplateData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateData`: InlineObject2
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetTemplateData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | Template unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineObject2**](InlineObject2.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplateSchema

> map[string]interface{} GetTemplateSchema(ctx).Execute()

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
	resp, r, err := apiClient.TemplatesAPI.GetTemplateSchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetTemplateSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplateSchema`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetTemplateSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateSchemaRequest struct via the builder pattern


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


## GetTemplates

> InlineObject4 GetTemplates(ctx).Name(name).Tags(tags).Access(access).Page(page).PerPage(perPage).Execute()

Get templates



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
	name := "name_example" // string | Filter template by name (optional)
	tags := "tags_example" // string | Filter template by tags (optional)
	access := "private" // string | Filter template by access type. No values returns all templates. private - returns only private templates, organization - returns only organization templates. (optional) (default to "")
	page := int32(1) // int32 | Pagination: page to return (optional) (default to 1)
	perPage := int32(20) // int32 | Pagination: How many records to return per page (optional) (default to 15)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetTemplates(context.Background()).Name(name).Tags(tags).Access(access).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplates`: InlineObject4
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter template by name | 
 **tags** | **string** | Filter template by tags | 
 **access** | **string** | Filter template by access type. No values returns all templates. private - returns only private templates, organization - returns only organization templates. | [default to &quot;&quot;]
 **page** | **int32** | Pagination: page to return | [default to 1]
 **perPage** | **int32** | Pagination: How many records to return per page | [default to 15]

### Return type

[**InlineObject4**](InlineObject4.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportTemplate

> InlineObject18 ImportTemplate(ctx).ImportTemplateRequest(importTemplateRequest).Execute()

Import template



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
	importTemplateRequest := openapiclient.importTemplate_request{ImportTemplateBase64: openapiclient.NewImportTemplateBase64(*openapiclient.NewImportTemplateUrlTemplate("Invoice template"), "FileBase64_example")} // ImportTemplateRequest | Import a PDF via URL or base64 string as template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.ImportTemplate(context.Background()).ImportTemplateRequest(importTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.ImportTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportTemplate`: InlineObject18
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.ImportTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importTemplateRequest** | [**ImportTemplateRequest**](ImportTemplateRequest.md) | Import a PDF via URL or base64 string as template | 

### Return type

[**InlineObject18**](InlineObject18.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenEditor

> InlineObject3 OpenEditor(ctx, templateId).OpenEditorRequest(openEditorRequest).Execute()

Open editor



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
	templateId := int32(19375) // int32 | Template unique identifier
	openEditorRequest := *openapiclient.NewOpenEditorRequest() // OpenEditorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.OpenEditor(context.Background(), templateId).OpenEditorRequest(openEditorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.OpenEditor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenEditor`: InlineObject3
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.OpenEditor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | Template unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openEditorRequest** | [**OpenEditorRequest**](OpenEditorRequest.md) |  | 

### Return type

[**InlineObject3**](InlineObject3.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplate

> InlineObject18 UpdateTemplate(ctx, templateId).TemplateDefinitionNew(templateDefinitionNew).Execute()

Update template



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
	templateId := int32(19375) // int32 | Template unique identifier
	templateDefinitionNew := *openapiclient.NewTemplateDefinitionNew("Invoice template") // TemplateDefinitionNew | Template configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.UpdateTemplate(context.Background(), templateId).TemplateDefinitionNew(templateDefinitionNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.UpdateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTemplate`: InlineObject18
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.UpdateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | Template unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateDefinitionNew** | [**TemplateDefinitionNew**](TemplateDefinitionNew.md) | Template configuration | 

### Return type

[**InlineObject18**](InlineObject18.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateTemplate

> InlineObject1 ValidateTemplate(ctx).TemplateDefinitionNew(templateDefinitionNew).Execute()

Validate template



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
	templateDefinitionNew := *openapiclient.NewTemplateDefinitionNew("Invoice template") // TemplateDefinitionNew | Template configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.ValidateTemplate(context.Background()).TemplateDefinitionNew(templateDefinitionNew).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.ValidateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateTemplate`: InlineObject1
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.ValidateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateDefinitionNew** | [**TemplateDefinitionNew**](TemplateDefinitionNew.md) | Template configuration | 

### Return type

[**InlineObject1**](InlineObject1.md)

### Authorization

[JSONWebTokenAuth](../README.md#JSONWebTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

