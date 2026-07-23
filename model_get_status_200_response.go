/*
PDF Generator API

# Introduction [PDF Generator API](https://pdfgeneratorapi.com) allows you easily generate transactional PDF documents and reduce the development and support costs by enabling your users to create and manage their document templates using a browser-based drag-and-drop document editor.  The PDF Generator API features a web API architecture, allowing you to code in the language of your choice. This API supports the JSON media type, and uses UTF-8 character encoding.  ## Base URL The base URL for all the API endpoints is `https://us1.pdfgeneratorapi.com/api/v4`  For example * `https://us1.pdfgeneratorapi.com/api/v4/templates` * `https://us1.pdfgeneratorapi.com/api/v4/workspaces` * `https://us1.pdfgeneratorapi.com/api/v4/templates/123123`  ## Editor PDF Generator API comes with a powerful drag & drop editor that allows to create any kind of document templates, from barcode labels to invoices, quotes and reports. You can find tutorials and videos from our [Support Portal](https://support.pdfgeneratorapi.com). * [Component specification](https://support.pdfgeneratorapi.com/en/category/components-1ffseaj/) * [Expression Language documentation](https://support.pdfgeneratorapi.com/en/category/expression-language-q203pa/) * [Frequently asked questions and answers](https://support.pdfgeneratorapi.com/en/category/qanda-1ov519d/)  ## Definitions  ### Organization Organization is a group of workspaces owned by your account.  ### Workspace Workspace contains templates. Each workspace has access to their own templates and organization default templates.  ### Master Workspace Master Workspace is the main/default workspace of your Organization. The Master Workspace identifier is the email you signed up with.  ### Default Template Default template is a template that is available for all workspaces by default. You can set the template access type under Page Setup. If template has \"Organization\" access then your users can use them from the \"New\" menu in the Editor.  ### Data Field Data Field is a placeholder for the specific data in your JSON data set. In this example JSON you can access the buyer name using Data Field `{paymentDetails::buyerName}`. The separator between depth levels is :: (two colons). When designing the template you don’t have to know every Data Field, our editor automatically extracts all the available fields from your data set and provides an easy way to insert them into the template. ``` {     \"documentNumber\": 1,     \"paymentDetails\": {         \"method\": \"Credit Card\",         \"buyerName\": \"John Smith\"     },     \"items\": [         {             \"id\": 1,             \"name\": \"Item one\"         }     ] } ```  ## Rate limiting Our API endpoints use IP-based rate limiting and allow you to make up to 2 requests per second and 60 requests per minute. If you make more requests, you will receive a response with HTTP code 429.  Response headers contain additional values:  | Header   | Description                    | |--------|--------------------------------| | X-RateLimit-Limit    | Maximum requests per minute                   | | X-RateLimit-Remaining    | The requests remaining in the current minute               | | Retry-After     | How many seconds you need to wait until you are allowed to make requests |  *  *  *  *  *  # Libraries and SDKs ## Postman Collection We have created a [Postman Collection](https://www.postman.com/pdfgeneratorapi/workspace/pdf-generator-api-public-workspace/overview) so you can easily test all the API endpoints without developing and code.   ## Client Libraries All our Client Libraries are auto-generated using [OpenAPI Generator](https://openapi-generator.tech/) which uses the OpenAPI v3 specification to automatically generate a client library in specific programming language.  * [PHP Client](https://github.com/pdfgeneratorapi/php-client) * [Java Client](https://github.com/pdfgeneratorapi/java-client) * [Ruby Client](https://github.com/pdfgeneratorapi/ruby-client) * [Python Client](https://github.com/pdfgeneratorapi/python-client) * [Javascript Client](https://github.com/pdfgeneratorapi/javascript-client)  We have validated the generated libraries, but let us know if you find any anomalies in the client code.  ## Model Context Protocol (MCP) Server Integrate document generation directly into your AI agents and LLM applications using our official Model Context Protocol (MCP) Server.  The MCP server provides a standardized interface that allows AI assistants (like Claude Desktop and other MCP-compatible clients) to securely interact with the PDF Generator API. With it, your AI applications can automatically fetch workspaces, retrieve templates, merge data, and generate PDF documents on the fly.  [Get PDF Generator API MCP Server](https://github.com/pdfgeneratorapi/mcp-server) *  *  *  *  *   # Authentication The PDF Generator API uses __JSON Web Tokens (JWT)__ to authenticate all API requests. These tokens offer a method to establish secure server-to-server authentication by transferring a compact JSON object with a signed payload of your account’s API Key and Secret. When authenticating to the PDF Generator API, a JWT should be generated uniquely by a __server-side application__ and included as a __Bearer Token__ in the header of each request.   <SecurityDefinitions />  ## Accessing your API Key and Secret You can find your __API Key__ and __API Secret__ from the __Account Settings__ page after you login to PDF Generator API [here](https://pdfgeneratorapi.com/login).  ## Creating a JWT JSON Web Tokens are composed of three sections: a header, a payload (containing a claim set), and a signature. The header and payload are JSON objects, which are serialized to UTF-8 bytes, then encoded using base64url encoding.  The JWT's header, payload, and signature are concatenated with periods (.). As a result, a JWT typically takes the following form: ``` {Base64url encoded header}.{Base64url encoded payload}.{Base64url encoded signature} ```  We recommend and support libraries provided on [jwt.io](https://jwt.io/). While other libraries can create JWT, these recommended libraries are the most robust.  ### Header Property `alg` defines which signing algorithm is being used. PDF Generator API users HS256. Property `typ` defines the type of token and it is always JWT. ``` {   \"alg\": \"HS256\",   \"typ\": \"JWT\" } ```  ### Payload The second part of the token is the payload, which contains the claims  or the pieces of information being passed about the user and any metadata required. It is mandatory to specify the following claims: * issuer (`iss`): Your API key * subject (`sub`): Workspace identifier * expiration time (`exp`): Timestamp (unix epoch time) until the token is valid. It is highly recommended to set the exp timestamp for a short period, i.e. a matter of seconds. This way, if a token is intercepted or shared, the token will only be valid for a short period of time.  ``` {   \"iss\": \"ad54aaff89ffdfeff178bb8a8f359b29fcb20edb56250b9f584aa2cb0162ed4a\",   \"sub\": \"demo.example@actualreports.com\",   \"exp\": 1586112639 } ```  ### Payload for Partners Our partners can send their unique identifier (provided by us) in JWT's partner_id claim. If the `partner_id` value is specified in the JWT, the organization making the request is automatically connected to the partner account. * Partner ID (`partner_id`): Unique identifier provide by PDF Generator API team  ``` {   \"iss\": \"ad54aaff89ffdfeff178bb8a8f359b29fcb20edb56250b9f584aa2cb0162ed4a\",   \"sub\": \"demo.example@actualreports.com\",   \"partner_id\": \"my-partner-identifier\",   \"exp\": 1586112639 } ```  ### Signature To create the signature part you have to take the encoded header, the encoded payload, a secret, the algorithm specified in the header, and sign that. The signature is used to verify the message wasn't changed along the way, and, in the case of tokens signed with a private key, it can also verify that the sender of the JWT is who it says it is. ``` HMACSHA256(     base64UrlEncode(header) + \".\" +     base64UrlEncode(payload),     API_SECRET) ```  ### Putting all together The output is three Base64-URL strings separated by dots. The following shows a JWT that has the previous header and payload encoded, and it is signed with a secret. ``` eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhZDU0YWFmZjg5ZmZkZmVmZjE3OGJiOGE4ZjM1OWIyOWZjYjIwZWRiNTYyNTBiOWY1ODRhYTJjYjAxNjJlZDRhIiwic3ViIjoiZGVtby5leGFtcGxlQGFjdHVhbHJlcG9ydHMuY29tIn0.SxO-H7UYYYsclS8RGWO1qf0z1cB1m73wF9FLl9RCc1Q  // Base64 encoded header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9 // Base64 encoded payload: eyJpc3MiOiJhZDU0YWFmZjg5ZmZkZmVmZjE3OGJiOGE4ZjM1OWIyOWZjYjIwZWRiNTYyNTBiOWY1ODRhYTJjYjAxNjJlZDRhIiwic3ViIjoiZGVtby5leGFtcGxlQGFjdHVhbHJlcG9ydHMuY29tIn0 // Signature: SxO-H7UYYYsclS8RGWO1qf0z1cB1m73wF9FLl9RCc1Q ```  ## Temporary JWTs You can create a temporary token in [Account Settings](https://pdfgeneratorapi.com/account/organization) page after you login to PDF Generator API. The generated token uses your email address as the subject (`sub`) value and is valid for __15 minutes__. You can also use [jwt.io](https://jwt.io/) to generate test tokens for your API calls. These test tokens should never be used in production applications. *  *  *  *  *  # Error codes  | Code   | Description                    | |--------|--------------------------------| | 401    | Unauthorized                   | | 402    | Payment Required               | | 403    | Forbidden                      | | 404    | Not Found                      | | 422    | Unprocessable Entity           | | 429    | Too Many Requests              | | 500    | Internal Server Error          |  ## 401 Unauthorized | Description                                                             | |-------------------------------------------------------------------------| | Authentication failed: request expired                                  | | Authentication failed: workspace missing                                | | Authentication failed: key missing                                      | | Authentication failed: property 'iss' (issuer) missing in JWT           | | Authentication failed: property 'sub' (subject) missing in JWT          | | Authentication failed: property 'exp' (expiration time) missing in JWT  | | Authentication failed: incorrect signature                              |  ## 402 Payment Required | Description                                                             | |-------------------------------------------------------------------------| | Your account is suspended, please upgrade your account                  |  ## 403 Forbidden | Description                                                             | |-------------------------------------------------------------------------| | Your account has exceeded the monthly document generation limit.        | | Access not granted: You cannot delete master workspace via API          | | Access not granted: Template is not accessible by this organization     | | Your session has expired, please close and reopen the editor.           |  ## 404 Entity not found | Description                                                             | |-------------------------------------------------------------------------| | Entity not found                                                        | | Resource not found                                                      | | None of the templates is available for the workspace.                   |  ## 422 Unprocessable Entity | Description                                                             | |-------------------------------------------------------------------------| | Unable to parse JSON, please check formatting                           | | Required parameter missing                                              | | Required parameter missing: template definition not defined             | | Required parameter missing: template not defined                        |  ## 429 Too Many Requests | Description                                                             | |-------------------------------------------------------------------------| | You can make up to 2 requests per second and 60 requests per minute.   |  *  *  *  *  * 

API version: 4.0.28
Contact: support@pdfgeneratorapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pdfgeneratorapi

import (
	"encoding/json"
)

// checks if the GetStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatus200Response{}

// GetStatus200Response struct for GetStatus200Response
type GetStatus200Response struct {
	Api *StatusParam `json:"api,omitempty"`
	ChartApi *StatusParam `json:"chart-api,omitempty"`
	ConversionApi *StatusParam `json:"conversion-api,omitempty"`
	GeneratorApiSync *StatusParam `json:"generator-api-sync,omitempty"`
	GeneratorApiAsync *StatusParam `json:"generator-api-async,omitempty"`
	EInvoice *StatusParam `json:"e-invoice,omitempty"`
}

// NewGetStatus200Response instantiates a new GetStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatus200Response() *GetStatus200Response {
	this := GetStatus200Response{}
	var api StatusParam = STATUSPARAM_OK
	this.Api = &api
	var chartApi StatusParam = STATUSPARAM_OK
	this.ChartApi = &chartApi
	var conversionApi StatusParam = STATUSPARAM_OK
	this.ConversionApi = &conversionApi
	var generatorApiSync StatusParam = STATUSPARAM_OK
	this.GeneratorApiSync = &generatorApiSync
	var generatorApiAsync StatusParam = STATUSPARAM_OK
	this.GeneratorApiAsync = &generatorApiAsync
	var eInvoice StatusParam = STATUSPARAM_OK
	this.EInvoice = &eInvoice
	return &this
}

// NewGetStatus200ResponseWithDefaults instantiates a new GetStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatus200ResponseWithDefaults() *GetStatus200Response {
	this := GetStatus200Response{}
	var api StatusParam = STATUSPARAM_OK
	this.Api = &api
	var chartApi StatusParam = STATUSPARAM_OK
	this.ChartApi = &chartApi
	var conversionApi StatusParam = STATUSPARAM_OK
	this.ConversionApi = &conversionApi
	var generatorApiSync StatusParam = STATUSPARAM_OK
	this.GeneratorApiSync = &generatorApiSync
	var generatorApiAsync StatusParam = STATUSPARAM_OK
	this.GeneratorApiAsync = &generatorApiAsync
	var eInvoice StatusParam = STATUSPARAM_OK
	this.EInvoice = &eInvoice
	return &this
}

// GetApi returns the Api field value if set, zero value otherwise.
func (o *GetStatus200Response) GetApi() StatusParam {
	if o == nil || IsNil(o.Api) {
		var ret StatusParam
		return ret
	}
	return *o.Api
}

// GetApiOk returns a tuple with the Api field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetApiOk() (*StatusParam, bool) {
	if o == nil || IsNil(o.Api) {
		return nil, false
	}
	return o.Api, true
}

// HasApi returns a boolean if a field has been set.
func (o *GetStatus200Response) HasApi() bool {
	if o != nil && !IsNil(o.Api) {
		return true
	}

	return false
}

// SetApi gets a reference to the given StatusParam and assigns it to the Api field.
func (o *GetStatus200Response) SetApi(v StatusParam) {
	o.Api = &v
}

// GetChartApi returns the ChartApi field value if set, zero value otherwise.
func (o *GetStatus200Response) GetChartApi() StatusParam {
	if o == nil || IsNil(o.ChartApi) {
		var ret StatusParam
		return ret
	}
	return *o.ChartApi
}

// GetChartApiOk returns a tuple with the ChartApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetChartApiOk() (*StatusParam, bool) {
	if o == nil || IsNil(o.ChartApi) {
		return nil, false
	}
	return o.ChartApi, true
}

// HasChartApi returns a boolean if a field has been set.
func (o *GetStatus200Response) HasChartApi() bool {
	if o != nil && !IsNil(o.ChartApi) {
		return true
	}

	return false
}

// SetChartApi gets a reference to the given StatusParam and assigns it to the ChartApi field.
func (o *GetStatus200Response) SetChartApi(v StatusParam) {
	o.ChartApi = &v
}

// GetConversionApi returns the ConversionApi field value if set, zero value otherwise.
func (o *GetStatus200Response) GetConversionApi() StatusParam {
	if o == nil || IsNil(o.ConversionApi) {
		var ret StatusParam
		return ret
	}
	return *o.ConversionApi
}

// GetConversionApiOk returns a tuple with the ConversionApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetConversionApiOk() (*StatusParam, bool) {
	if o == nil || IsNil(o.ConversionApi) {
		return nil, false
	}
	return o.ConversionApi, true
}

// HasConversionApi returns a boolean if a field has been set.
func (o *GetStatus200Response) HasConversionApi() bool {
	if o != nil && !IsNil(o.ConversionApi) {
		return true
	}

	return false
}

// SetConversionApi gets a reference to the given StatusParam and assigns it to the ConversionApi field.
func (o *GetStatus200Response) SetConversionApi(v StatusParam) {
	o.ConversionApi = &v
}

// GetGeneratorApiSync returns the GeneratorApiSync field value if set, zero value otherwise.
func (o *GetStatus200Response) GetGeneratorApiSync() StatusParam {
	if o == nil || IsNil(o.GeneratorApiSync) {
		var ret StatusParam
		return ret
	}
	return *o.GeneratorApiSync
}

// GetGeneratorApiSyncOk returns a tuple with the GeneratorApiSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetGeneratorApiSyncOk() (*StatusParam, bool) {
	if o == nil || IsNil(o.GeneratorApiSync) {
		return nil, false
	}
	return o.GeneratorApiSync, true
}

// HasGeneratorApiSync returns a boolean if a field has been set.
func (o *GetStatus200Response) HasGeneratorApiSync() bool {
	if o != nil && !IsNil(o.GeneratorApiSync) {
		return true
	}

	return false
}

// SetGeneratorApiSync gets a reference to the given StatusParam and assigns it to the GeneratorApiSync field.
func (o *GetStatus200Response) SetGeneratorApiSync(v StatusParam) {
	o.GeneratorApiSync = &v
}

// GetGeneratorApiAsync returns the GeneratorApiAsync field value if set, zero value otherwise.
func (o *GetStatus200Response) GetGeneratorApiAsync() StatusParam {
	if o == nil || IsNil(o.GeneratorApiAsync) {
		var ret StatusParam
		return ret
	}
	return *o.GeneratorApiAsync
}

// GetGeneratorApiAsyncOk returns a tuple with the GeneratorApiAsync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetGeneratorApiAsyncOk() (*StatusParam, bool) {
	if o == nil || IsNil(o.GeneratorApiAsync) {
		return nil, false
	}
	return o.GeneratorApiAsync, true
}

// HasGeneratorApiAsync returns a boolean if a field has been set.
func (o *GetStatus200Response) HasGeneratorApiAsync() bool {
	if o != nil && !IsNil(o.GeneratorApiAsync) {
		return true
	}

	return false
}

// SetGeneratorApiAsync gets a reference to the given StatusParam and assigns it to the GeneratorApiAsync field.
func (o *GetStatus200Response) SetGeneratorApiAsync(v StatusParam) {
	o.GeneratorApiAsync = &v
}

// GetEInvoice returns the EInvoice field value if set, zero value otherwise.
func (o *GetStatus200Response) GetEInvoice() StatusParam {
	if o == nil || IsNil(o.EInvoice) {
		var ret StatusParam
		return ret
	}
	return *o.EInvoice
}

// GetEInvoiceOk returns a tuple with the EInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatus200Response) GetEInvoiceOk() (*StatusParam, bool) {
	if o == nil || IsNil(o.EInvoice) {
		return nil, false
	}
	return o.EInvoice, true
}

// HasEInvoice returns a boolean if a field has been set.
func (o *GetStatus200Response) HasEInvoice() bool {
	if o != nil && !IsNil(o.EInvoice) {
		return true
	}

	return false
}

// SetEInvoice gets a reference to the given StatusParam and assigns it to the EInvoice field.
func (o *GetStatus200Response) SetEInvoice(v StatusParam) {
	o.EInvoice = &v
}

func (o GetStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Api) {
		toSerialize["api"] = o.Api
	}
	if !IsNil(o.ChartApi) {
		toSerialize["chart-api"] = o.ChartApi
	}
	if !IsNil(o.ConversionApi) {
		toSerialize["conversion-api"] = o.ConversionApi
	}
	if !IsNil(o.GeneratorApiSync) {
		toSerialize["generator-api-sync"] = o.GeneratorApiSync
	}
	if !IsNil(o.GeneratorApiAsync) {
		toSerialize["generator-api-async"] = o.GeneratorApiAsync
	}
	if !IsNil(o.EInvoice) {
		toSerialize["e-invoice"] = o.EInvoice
	}
	return toSerialize, nil
}

type NullableGetStatus200Response struct {
	value *GetStatus200Response
	isSet bool
}

func (v NullableGetStatus200Response) Get() *GetStatus200Response {
	return v.value
}

func (v *NullableGetStatus200Response) Set(val *GetStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatus200Response(val *GetStatus200Response) *NullableGetStatus200Response {
	return &NullableGetStatus200Response{value: val, isSet: true}
}

func (v NullableGetStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


