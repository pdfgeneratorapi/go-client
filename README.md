# Go API client for pdfgeneratorapi

# Introduction
[PDF Generator API](https://pdfgeneratorapi.com) allows you easily generate transactional PDF documents and reduce the development and support costs by enabling your users to create and manage their document templates using a browser-based drag-and-drop document editor.

The PDF Generator API features a web API architecture, allowing you to code in the language of your choice. This API supports the JSON media type, and uses UTF-8 character encoding.

## Base URL
The base URL for all the API endpoints is `https://us1.pdfgeneratorapi.com/api/v4`

For example
* `https://us1.pdfgeneratorapi.com/api/v4/templates`
* `https://us1.pdfgeneratorapi.com/api/v4/workspaces`
* `https://us1.pdfgeneratorapi.com/api/v4/templates/123123`

## Editor
PDF Generator API comes with a powerful drag & drop editor that allows to create any kind of document templates, from barcode labels to invoices, quotes and reports. You can find tutorials and videos from our [Support Portal](https://support.pdfgeneratorapi.com).
* [Component specification](https://support.pdfgeneratorapi.com/en/category/components-1ffseaj/)
* [Expression Language documentation](https://support.pdfgeneratorapi.com/en/category/expression-language-q203pa/)
* [Frequently asked questions and answers](https://support.pdfgeneratorapi.com/en/category/qanda-1ov519d/)

## Definitions

### Organization
Organization is a group of workspaces owned by your account.

### Workspace
Workspace contains templates. Each workspace has access to their own templates and organization default templates.

### Master Workspace
Master Workspace is the main/default workspace of your Organization. The Master Workspace identifier is the email you signed up with.

### Default Template
Default template is a template that is available for all workspaces by default. You can set the template access type under Page Setup. If template has "Organization" access then your users can use them from the "New" menu in the Editor.

### Data Field
Data Field is a placeholder for the specific data in your JSON data set. In this example JSON you can access the buyer name using Data Field `{paymentDetails::buyerName}`. The separator between depth levels is :: (two colons). When designing the template you don’t have to know every Data Field, our editor automatically extracts all the available fields from your data set and provides an easy way to insert them into the template.
```
{
    "documentNumber": 1,
    "paymentDetails": {
        "method": "Credit Card",
        "buyerName": "John Smith"
    },
    "items": [
        {
            "id": 1,
            "name": "Item one"
        }
    ]
}
```

## Rate limiting
Our API endpoints use IP-based rate limiting and allow you to make up to 2 requests per second and 60 requests per minute. If you make more requests, you will receive a response with HTTP code 429.

Response headers contain additional values:

| Header   | Description                    |
|--------|--------------------------------|
| X-RateLimit-Limit    | Maximum requests per minute                   |
| X-RateLimit-Remaining    | The requests remaining in the current minute               |
| Retry-After     | How many seconds you need to wait until you are allowed to make requests |

*  *  *  *  *

# Libraries and SDKs
## Postman Collection
We have created a [Postman Collection](https://www.postman.com/pdfgeneratorapi/workspace/pdf-generator-api-public-workspace/overview) so you can easily test all the API endpoints without developing and code.


## Client Libraries
All our Client Libraries are auto-generated using [OpenAPI Generator](https://openapi-generator.tech/) which uses the OpenAPI v3 specification to automatically generate a client library in specific programming language.

* [PHP Client](https://github.com/pdfgeneratorapi/php-client)
* [Java Client](https://github.com/pdfgeneratorapi/java-client)
* [Ruby Client](https://github.com/pdfgeneratorapi/ruby-client)
* [Python Client](https://github.com/pdfgeneratorapi/python-client)
* [Javascript Client](https://github.com/pdfgeneratorapi/javascript-client)

We have validated the generated libraries, but let us know if you find any anomalies in the client code.

## Model Context Protocol (MCP) Server
Integrate document generation directly into your AI agents and LLM applications using our official Model Context Protocol (MCP) Server.

The MCP server provides a standardized interface that allows AI assistants (like Claude Desktop and other MCP-compatible clients) to securely interact with the PDF Generator API. With it, your AI applications can automatically fetch workspaces, retrieve templates, merge data, and generate PDF documents on the fly.

[Get PDF Generator API MCP Server](https://github.com/pdfgeneratorapi/mcp-server)
*  *  *  *  *


# Authentication
The PDF Generator API uses __JSON Web Tokens (JWT)__ to authenticate all API requests. These tokens offer a method to establish secure server-to-server authentication by transferring a compact JSON object with a signed payload of your account’s API Key and Secret.
When authenticating to the PDF Generator API, a JWT should be generated uniquely by a __server-side application__ and included as a __Bearer Token__ in the header of each request.


<SecurityDefinitions />

## Accessing your API Key and Secret
You can find your __API Key__ and __API Secret__ from the __Account Settings__ page after you login to PDF Generator API [here](https://pdfgeneratorapi.com/login).

## Creating a JWT
JSON Web Tokens are composed of three sections: a header, a payload (containing a claim set), and a signature. The header and payload are JSON objects, which are serialized to UTF-8 bytes, then encoded using base64url encoding.

The JWT's header, payload, and signature are concatenated with periods (.). As a result, a JWT typically takes the following form:
```
{Base64url encoded header}.{Base64url encoded payload}.{Base64url encoded signature}
```

We recommend and support libraries provided on [jwt.io](https://jwt.io/). While other libraries can create JWT, these recommended libraries are the most robust.

### Header
Property `alg` defines which signing algorithm is being used. PDF Generator API users HS256.
Property `typ` defines the type of token and it is always JWT.
```
{
  "alg": "HS256",
  "typ": "JWT"
}
```

### Payload
The second part of the token is the payload, which contains the claims  or the pieces of information being passed about the user and any metadata required.
It is mandatory to specify the following claims:
* issuer (`iss`): Your API key
* subject (`sub`): Workspace identifier
* expiration time (`exp`): Timestamp (unix epoch time) until the token is valid. It is highly recommended to set the exp timestamp for a short period, i.e. a matter of seconds. This way, if a token is intercepted or shared, the token will only be valid for a short period of time.

```
{
  "iss": "ad54aaff89ffdfeff178bb8a8f359b29fcb20edb56250b9f584aa2cb0162ed4a",
  "sub": "demo.example@actualreports.com",
  "exp": 1586112639
}
```

### Payload for Partners
Our partners can send their unique identifier (provided by us) in JWT's partner_id claim. If the `partner_id` value is specified in the JWT, the organization making the request is automatically connected to the partner account.
* Partner ID (`partner_id`): Unique identifier provide by PDF Generator API team

```
{
  "iss": "ad54aaff89ffdfeff178bb8a8f359b29fcb20edb56250b9f584aa2cb0162ed4a",
  "sub": "demo.example@actualreports.com",
  "partner_id": "my-partner-identifier",
  "exp": 1586112639
}
```

### Signature
To create the signature part you have to take the encoded header, the encoded payload, a secret, the algorithm specified in the header, and sign that. The signature is used to verify the message wasn't changed along the way, and, in the case of tokens signed with a private key, it can also verify that the sender of the JWT is who it says it is.
```
HMACSHA256(
    base64UrlEncode(header) + "." +
    base64UrlEncode(payload),
    API_SECRET)
```

### Putting all together
The output is three Base64-URL strings separated by dots. The following shows a JWT that has the previous header and payload encoded, and it is signed with a secret.
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhZDU0YWFmZjg5ZmZkZmVmZjE3OGJiOGE4ZjM1OWIyOWZjYjIwZWRiNTYyNTBiOWY1ODRhYTJjYjAxNjJlZDRhIiwic3ViIjoiZGVtby5leGFtcGxlQGFjdHVhbHJlcG9ydHMuY29tIn0.SxO-H7UYYYsclS8RGWO1qf0z1cB1m73wF9FLl9RCc1Q

// Base64 encoded header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
// Base64 encoded payload: eyJpc3MiOiJhZDU0YWFmZjg5ZmZkZmVmZjE3OGJiOGE4ZjM1OWIyOWZjYjIwZWRiNTYyNTBiOWY1ODRhYTJjYjAxNjJlZDRhIiwic3ViIjoiZGVtby5leGFtcGxlQGFjdHVhbHJlcG9ydHMuY29tIn0
// Signature: SxO-H7UYYYsclS8RGWO1qf0z1cB1m73wF9FLl9RCc1Q
```

## Temporary JWTs
You can create a temporary token in [Account Settings](https://pdfgeneratorapi.com/account/organization) page after you login to PDF Generator API. The generated token uses your email address as the subject (`sub`) value and is valid for __15 minutes__.
You can also use [jwt.io](https://jwt.io/) to generate test tokens for your API calls. These test tokens should never be used in production applications.
*  *  *  *  *

# Error codes

| Code   | Description                    |
|--------|--------------------------------|
| 401    | Unauthorized                   |
| 402    | Payment Required               |
| 403    | Forbidden                      |
| 404    | Not Found                      |
| 422    | Unprocessable Entity           |
| 429    | Too Many Requests              |
| 500    | Internal Server Error          |

## 401 Unauthorized
| Description                                                             |
|-------------------------------------------------------------------------|
| Authentication failed: request expired                                  |
| Authentication failed: workspace missing                                |
| Authentication failed: key missing                                      |
| Authentication failed: property 'iss' (issuer) missing in JWT           |
| Authentication failed: property 'sub' (subject) missing in JWT          |
| Authentication failed: property 'exp' (expiration time) missing in JWT  |
| Authentication failed: incorrect signature                              |

## 402 Payment Required
| Description                                                             |
|-------------------------------------------------------------------------|
| Your account is suspended, please upgrade your account                  |

## 403 Forbidden
| Description                                                             |
|-------------------------------------------------------------------------|
| Your account has exceeded the monthly document generation limit.        |
| Access not granted: You cannot delete master workspace via API          |
| Access not granted: Template is not accessible by this organization     |
| Your session has expired, please close and reopen the editor.           |

## 404 Entity not found
| Description                                                             |
|-------------------------------------------------------------------------|
| Entity not found                                                        |
| Resource not found                                                      |
| None of the templates is available for the workspace.                   |

## 422 Unprocessable Entity
| Description                                                             |
|-------------------------------------------------------------------------|
| Unable to parse JSON, please check formatting                           |
| Required parameter missing                                              |
| Required parameter missing: template definition not defined             |
| Required parameter missing: template not defined                        |

## 429 Too Many Requests
| Description                                                             |
|-------------------------------------------------------------------------|
| You can make up to 2 requests per second and 60 requests per minute.   |

*  *  *  *  *


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 4.0.25
- Package version: 1.0.0
- Generator version: 7.14.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://support.pdfgeneratorapi.com](https://support.pdfgeneratorapi.com)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import pdfgeneratorapi "github.com/pdfgeneratorapi/go-client/v4"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `pdfgeneratorapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), pdfgeneratorapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `pdfgeneratorapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), pdfgeneratorapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `pdfgeneratorapi.ContextOperationServerIndices` and `pdfgeneratorapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), pdfgeneratorapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), pdfgeneratorapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://us1.pdfgeneratorapi.com/api/v4*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AssetsAPI* | [**GenerateQRCode**](docs/AssetsAPI.md#generateqrcode) | **Post** /assets/qrcode | Generate QR Code
*ConversionAPI* | [**ConvertHTML2PDF**](docs/ConversionAPI.md#converthtml2pdf) | **Post** /conversion/html2pdf | HTML to PDF
*ConversionAPI* | [**ConvertPDF2Image**](docs/ConversionAPI.md#convertpdf2image) | **Post** /conversion/pdf2image | PDF to Image
*ConversionAPI* | [**ConvertURL2PDF**](docs/ConversionAPI.md#converturl2pdf) | **Post** /conversion/url2pdf | URL to PDF
*DocumentsAPI* | [**DeleteDocument**](docs/DocumentsAPI.md#deletedocument) | **Delete** /documents/{publicId} | Delete document
*DocumentsAPI* | [**GenerateDocument**](docs/DocumentsAPI.md#generatedocument) | **Post** /documents/generate | Generate document
*DocumentsAPI* | [**GenerateDocumentAsynchronous**](docs/DocumentsAPI.md#generatedocumentasynchronous) | **Post** /documents/generate/async | Generate document (async)
*DocumentsAPI* | [**GenerateDocumentBatch**](docs/DocumentsAPI.md#generatedocumentbatch) | **Post** /documents/generate/batch | Generate document (batch)
*DocumentsAPI* | [**GenerateDocumentBatchAsynchronous**](docs/DocumentsAPI.md#generatedocumentbatchasynchronous) | **Post** /documents/generate/batch/async | Generate document (batch + async)
*DocumentsAPI* | [**GetAsyncJobStatus**](docs/DocumentsAPI.md#getasyncjobstatus) | **Get** /documents/async/{jobId} | Get job status
*DocumentsAPI* | [**GetDocument**](docs/DocumentsAPI.md#getdocument) | **Get** /documents/{publicId} | Get document
*DocumentsAPI* | [**GetDocuments**](docs/DocumentsAPI.md#getdocuments) | **Get** /documents | Get documents
*DocumentsAPI* | [**StoreDocument**](docs/DocumentsAPI.md#storedocument) | **Post** /documents | Store document
*EInvoicesAPI* | [**CreateEInvoice**](docs/EInvoicesAPI.md#createeinvoice) | **Post** /einvoice | Create eInvoice
*EInvoicesAPI* | [**CreateFacturXEInvoice**](docs/EInvoicesAPI.md#createfacturxeinvoice) | **Post** /einvoice/facturx | Create Factur-X eInvoice
*EInvoicesAPI* | [**CreateXRechnungEInvoice**](docs/EInvoicesAPI.md#createxrechnungeinvoice) | **Post** /einvoice/xrechnung | Create XRechnung eInvoice
*EInvoicesAPI* | [**GetEInvoiceSchema**](docs/EInvoicesAPI.md#geteinvoiceschema) | **Get** /einvoice/schema | Get schema
*FormsAPI* | [**CreateForm**](docs/FormsAPI.md#createform) | **Post** /forms | Create form
*FormsAPI* | [**DeleteForm**](docs/FormsAPI.md#deleteform) | **Delete** /forms/{formId} | Delete form
*FormsAPI* | [**GetForm**](docs/FormsAPI.md#getform) | **Get** /forms/{formId} | Get form
*FormsAPI* | [**GetForms**](docs/FormsAPI.md#getforms) | **Get** /forms | Get forms
*FormsAPI* | [**ImportForm**](docs/FormsAPI.md#importform) | **Post** /forms/import | Import Form
*FormsAPI* | [**OpenFormBuilder**](docs/FormsAPI.md#openformbuilder) | **Post** /forms/open | Open new form builder
*FormsAPI* | [**OpenFormBuilderForExistingForm**](docs/FormsAPI.md#openformbuilderforexistingform) | **Post** /forms/{formId}/open | Open existing form builder
*FormsAPI* | [**ShareForm**](docs/FormsAPI.md#shareform) | **Post** /forms/{formId}/share | Share form
*FormsAPI* | [**UpdateForm**](docs/FormsAPI.md#updateform) | **Put** /forms/{formId} | Update form
*MiscAPI* | [**GetStatus**](docs/MiscAPI.md#getstatus) | **Get** /status | Get Service Status
*ServicesAPI* | [**AddWatermark**](docs/ServicesAPI.md#addwatermark) | **Post** /pdfservices/watermark | Add watermark
*ServicesAPI* | [**DecryptDocument**](docs/ServicesAPI.md#decryptdocument) | **Post** /pdfservices/decrypt | Decrypt document
*ServicesAPI* | [**EncryptDocument**](docs/ServicesAPI.md#encryptdocument) | **Post** /pdfservices/encrypt | Encrypt document
*ServicesAPI* | [**ExtractFormFields**](docs/ServicesAPI.md#extractformfields) | **Post** /pdfservices/form/fields | Extract form fields
*ServicesAPI* | [**FillFormFields**](docs/ServicesAPI.md#fillformfields) | **Post** /pdfservices/form/fill | Fill form fields
*ServicesAPI* | [**MakeAccessible**](docs/ServicesAPI.md#makeaccessible) | **Post** /pdfservices/make-accessible | Make accessible
*ServicesAPI* | [**OptimizeDocument**](docs/ServicesAPI.md#optimizedocument) | **Post** /pdfservices/optimize | Optimize document
*TemplateLibraryAPI* | [**GetTemplateLibrary**](docs/TemplateLibraryAPI.md#gettemplatelibrary) | **Get** /templates/library | Get template library
*TemplateLibraryAPI* | [**GetTemplateLibraryItem**](docs/TemplateLibraryAPI.md#gettemplatelibraryitem) | **Get** /templates/library/{publicId} | Open template from the library
*TemplateVersionsAPI* | [**DeleteTemplateVersion**](docs/TemplateVersionsAPI.md#deletetemplateversion) | **Delete** /templates/{templateId}/versions/{templateVersion} | Delete template version
*TemplateVersionsAPI* | [**GetTemplateVersion**](docs/TemplateVersionsAPI.md#gettemplateversion) | **Get** /templates/{templateId}/versions/{templateVersion} | Get template version
*TemplateVersionsAPI* | [**ListTemplateVersions**](docs/TemplateVersionsAPI.md#listtemplateversions) | **Get** /templates/{templateId}/versions | List template versions
*TemplateVersionsAPI* | [**PromoteTemplateVersion**](docs/TemplateVersionsAPI.md#promotetemplateversion) | **Put** /templates/{templateId}/versions/{templateVersion}/promote | Promote template version
*TemplatesAPI* | [**CopyTemplate**](docs/TemplatesAPI.md#copytemplate) | **Post** /templates/{templateId}/copy | Copy template
*TemplatesAPI* | [**CreateTemplate**](docs/TemplatesAPI.md#createtemplate) | **Post** /templates | Create template
*TemplatesAPI* | [**DeleteTemplate**](docs/TemplatesAPI.md#deletetemplate) | **Delete** /templates/{templateId} | Delete template
*TemplatesAPI* | [**GetTemplate**](docs/TemplatesAPI.md#gettemplate) | **Get** /templates/{templateId} | Get template
*TemplatesAPI* | [**GetTemplateData**](docs/TemplatesAPI.md#gettemplatedata) | **Get** /templates/{templateId}/data | Get template data fields
*TemplatesAPI* | [**GetTemplateSchema**](docs/TemplatesAPI.md#gettemplateschema) | **Get** /templates/schema | Get schema
*TemplatesAPI* | [**GetTemplates**](docs/TemplatesAPI.md#gettemplates) | **Get** /templates | Get templates
*TemplatesAPI* | [**ImportTemplate**](docs/TemplatesAPI.md#importtemplate) | **Post** /templates/import | Import template
*TemplatesAPI* | [**OpenEditor**](docs/TemplatesAPI.md#openeditor) | **Post** /templates/{templateId}/editor | Open editor
*TemplatesAPI* | [**UpdateTemplate**](docs/TemplatesAPI.md#updatetemplate) | **Put** /templates/{templateId} | Update template
*TemplatesAPI* | [**ValidateTemplate**](docs/TemplatesAPI.md#validatetemplate) | **Post** /templates/validate | Validate template
*WorkspacesAPI* | [**CreateWorkspace**](docs/WorkspacesAPI.md#createworkspace) | **Post** /workspaces | Create workspace
*WorkspacesAPI* | [**DeleteWorkspace**](docs/WorkspacesAPI.md#deleteworkspace) | **Delete** /workspaces/{workspaceIdentifier} | Delete workspace
*WorkspacesAPI* | [**GetWorkspace**](docs/WorkspacesAPI.md#getworkspace) | **Get** /workspaces/{workspaceIdentifier} | Get workspace
*WorkspacesAPI* | [**GetWorkspaces**](docs/WorkspacesAPI.md#getworkspaces) | **Get** /workspaces | Get workspaces


## Documentation For Models

 - [AccessibilityOption](docs/AccessibilityOption.md)
 - [AddWatermarkRequest](docs/AddWatermarkRequest.md)
 - [AsyncOutputParam](docs/AsyncOutputParam.md)
 - [CallbackParam](docs/CallbackParam.md)
 - [Component](docs/Component.md)
 - [ConvertHTML2PDFRequest](docs/ConvertHTML2PDFRequest.md)
 - [ConvertPDF2ImageRequest](docs/ConvertPDF2ImageRequest.md)
 - [ConvertURL2PDFRequest](docs/ConvertURL2PDFRequest.md)
 - [CopyTemplateRequest](docs/CopyTemplateRequest.md)
 - [CreateEInvoiceRequest](docs/CreateEInvoiceRequest.md)
 - [CreateFacturXEInvoiceRequest](docs/CreateFacturXEInvoiceRequest.md)
 - [CreateFacturXEInvoiceRequestTemplate](docs/CreateFacturXEInvoiceRequestTemplate.md)
 - [CreateWorkspaceRequest](docs/CreateWorkspaceRequest.md)
 - [DataBatchInner](docs/DataBatchInner.md)
 - [Document](docs/Document.md)
 - [EncryptAndDecryptBase64](docs/EncryptAndDecryptBase64.md)
 - [EncryptAndDecryptUrl](docs/EncryptAndDecryptUrl.md)
 - [EncryptDocumentRequest](docs/EncryptDocumentRequest.md)
 - [ExtractFormFieldsRequest](docs/ExtractFormFieldsRequest.md)
 - [FillFormFieldsRequest](docs/FillFormFieldsRequest.md)
 - [FormActionDownload](docs/FormActionDownload.md)
 - [FormActionSend](docs/FormActionSend.md)
 - [FormActionSendSendDocument](docs/FormActionSendSendDocument.md)
 - [FormActionSendSendDocumentHeadersInner](docs/FormActionSendSendDocumentHeadersInner.md)
 - [FormActionStore](docs/FormActionStore.md)
 - [FormConfiguration](docs/FormConfiguration.md)
 - [FormConfigurationNew](docs/FormConfigurationNew.md)
 - [FormConfigurationNewActionsInner](docs/FormConfigurationNewActionsInner.md)
 - [FormFieldsBase64](docs/FormFieldsBase64.md)
 - [FormFieldsInner](docs/FormFieldsInner.md)
 - [FormFieldsUrl](docs/FormFieldsUrl.md)
 - [FormFillBase64](docs/FormFillBase64.md)
 - [FormFillUrl](docs/FormFillUrl.md)
 - [FormatParam](docs/FormatParam.md)
 - [GenerateDocumentAsynchronousRequest](docs/GenerateDocumentAsynchronousRequest.md)
 - [GenerateDocumentBatchAsynchronousRequest](docs/GenerateDocumentBatchAsynchronousRequest.md)
 - [GenerateDocumentBatchRequest](docs/GenerateDocumentBatchRequest.md)
 - [GenerateDocumentRequest](docs/GenerateDocumentRequest.md)
 - [GenerateQRCode201Response](docs/GenerateQRCode201Response.md)
 - [GenerateQRCode201ResponseMeta](docs/GenerateQRCode201ResponseMeta.md)
 - [GenerateQRCodeRequest](docs/GenerateQRCodeRequest.md)
 - [GetStatus200Response](docs/GetStatus200Response.md)
 - [GetTemplateLibrary200Response](docs/GetTemplateLibrary200Response.md)
 - [GetTemplateVersion422Response](docs/GetTemplateVersion422Response.md)
 - [ImportFormBase64](docs/ImportFormBase64.md)
 - [ImportFormRequest](docs/ImportFormRequest.md)
 - [ImportFormUrl](docs/ImportFormUrl.md)
 - [ImportTemplateBase64](docs/ImportTemplateBase64.md)
 - [ImportTemplateRequest](docs/ImportTemplateRequest.md)
 - [ImportTemplateUrl](docs/ImportTemplateUrl.md)
 - [ImportTemplateUrlTemplate](docs/ImportTemplateUrlTemplate.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject10](docs/InlineObject10.md)
 - [InlineObject10Meta](docs/InlineObject10Meta.md)
 - [InlineObject11](docs/InlineObject11.md)
 - [InlineObject11Meta](docs/InlineObject11Meta.md)
 - [InlineObject12](docs/InlineObject12.md)
 - [InlineObject12Meta](docs/InlineObject12Meta.md)
 - [InlineObject12MetaStats](docs/InlineObject12MetaStats.md)
 - [InlineObject13](docs/InlineObject13.md)
 - [InlineObject14](docs/InlineObject14.md)
 - [InlineObject14ResponseValue](docs/InlineObject14ResponseValue.md)
 - [InlineObject14ResponseValueDefault](docs/InlineObject14ResponseValueDefault.md)
 - [InlineObject14ResponseValueValue](docs/InlineObject14ResponseValueValue.md)
 - [InlineObject15](docs/InlineObject15.md)
 - [InlineObject16](docs/InlineObject16.md)
 - [InlineObject17](docs/InlineObject17.md)
 - [InlineObject18](docs/InlineObject18.md)
 - [InlineObject18Meta](docs/InlineObject18Meta.md)
 - [InlineObject19](docs/InlineObject19.md)
 - [InlineObject1Response](docs/InlineObject1Response.md)
 - [InlineObject2](docs/InlineObject2.md)
 - [InlineObject20](docs/InlineObject20.md)
 - [InlineObject20Response](docs/InlineObject20Response.md)
 - [InlineObject21](docs/InlineObject21.md)
 - [InlineObject22](docs/InlineObject22.md)
 - [InlineObject23](docs/InlineObject23.md)
 - [InlineObject24](docs/InlineObject24.md)
 - [InlineObject25](docs/InlineObject25.md)
 - [InlineObject26](docs/InlineObject26.md)
 - [InlineObject27](docs/InlineObject27.md)
 - [InlineObject3](docs/InlineObject3.md)
 - [InlineObject4](docs/InlineObject4.md)
 - [InlineObject5](docs/InlineObject5.md)
 - [InlineObject6](docs/InlineObject6.md)
 - [InlineObject7](docs/InlineObject7.md)
 - [InlineObject7Response](docs/InlineObject7Response.md)
 - [InlineObject8](docs/InlineObject8.md)
 - [InlineObject9](docs/InlineObject9.md)
 - [InlineObject9Meta](docs/InlineObject9Meta.md)
 - [InlineObjectMeta](docs/InlineObjectMeta.md)
 - [KeyFieldParam](docs/KeyFieldParam.md)
 - [MakeAccessibleBase64](docs/MakeAccessibleBase64.md)
 - [MakeAccessibleRequest](docs/MakeAccessibleRequest.md)
 - [MakeAccessibleUrl](docs/MakeAccessibleUrl.md)
 - [MetadataParam](docs/MetadataParam.md)
 - [OpenEditorRequest](docs/OpenEditorRequest.md)
 - [OpenEditorRequestData](docs/OpenEditorRequestData.md)
 - [OptimizeBase64](docs/OptimizeBase64.md)
 - [OptimizeDocumentRequest](docs/OptimizeDocumentRequest.md)
 - [OptimizeUrl](docs/OptimizeUrl.md)
 - [OutputParam](docs/OutputParam.md)
 - [PaginationMeta](docs/PaginationMeta.md)
 - [PromoteTemplateVersion200Response](docs/PromoteTemplateVersion200Response.md)
 - [PublicTemplateLibraryItem](docs/PublicTemplateLibraryItem.md)
 - [StatusParam](docs/StatusParam.md)
 - [StoreDocumentRequest](docs/StoreDocumentRequest.md)
 - [Template](docs/Template.md)
 - [TemplateDefinition](docs/TemplateDefinition.md)
 - [TemplateDefinitionDataSettings](docs/TemplateDefinitionDataSettings.md)
 - [TemplateDefinitionEditor](docs/TemplateDefinitionEditor.md)
 - [TemplateDefinitionNew](docs/TemplateDefinitionNew.md)
 - [TemplateDefinitionNewDataSettings](docs/TemplateDefinitionNewDataSettings.md)
 - [TemplateDefinitionNewEditor](docs/TemplateDefinitionNewEditor.md)
 - [TemplateDefinitionNewLayout](docs/TemplateDefinitionNewLayout.md)
 - [TemplateDefinitionNewLayoutMargins](docs/TemplateDefinitionNewLayoutMargins.md)
 - [TemplateDefinitionNewLayoutRepeatLayout](docs/TemplateDefinitionNewLayoutRepeatLayout.md)
 - [TemplateDefinitionNewPagesInner](docs/TemplateDefinitionNewPagesInner.md)
 - [TemplateDefinitionNewPagesInnerMargins](docs/TemplateDefinitionNewPagesInnerMargins.md)
 - [TemplateDefinitionPagesInner](docs/TemplateDefinitionPagesInner.md)
 - [TemplateParam](docs/TemplateParam.md)
 - [TemplateParamData](docs/TemplateParamData.md)
 - [TemplateParamId](docs/TemplateParamId.md)
 - [TemplateVersion](docs/TemplateVersion.md)
 - [TemplateVersionCollection](docs/TemplateVersionCollection.md)
 - [TemplateVersionCollectionMeta](docs/TemplateVersionCollectionMeta.md)
 - [WatermarkBase64](docs/WatermarkBase64.md)
 - [WatermarkFileUrl](docs/WatermarkFileUrl.md)
 - [WatermarkFileUrlWatermark](docs/WatermarkFileUrlWatermark.md)
 - [WatermarkImage](docs/WatermarkImage.md)
 - [WatermarkImageContentBase64](docs/WatermarkImageContentBase64.md)
 - [WatermarkImageContentUrl](docs/WatermarkImageContentUrl.md)
 - [WatermarkPosition](docs/WatermarkPosition.md)
 - [WatermarkText](docs/WatermarkText.md)
 - [Workspace](docs/Workspace.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### JSONWebTokenAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), pdfgeneratorapi.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@pdfgeneratorapi.com

