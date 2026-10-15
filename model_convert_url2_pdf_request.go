/*
PDF Generator API

# Introduction [PDF Generator API](https://pdfgeneratorapi.com) allows you easily generate transactional PDF documents and reduce the development and support costs by enabling your users to create and manage their document templates using a browser-based drag-and-drop document editor.  The PDF Generator API features a web API architecture, allowing you to code in the language of your choice. This API supports the JSON media type, and uses UTF-8 character encoding.  ## Base URL The base URL for all the API endpoints is `https://us1.pdfgeneratorapi.com/api/v4`  For example * `https://us1.pdfgeneratorapi.com/api/v4/templates` * `https://us1.pdfgeneratorapi.com/api/v4/workspaces` * `https://us1.pdfgeneratorapi.com/api/v4/templates/123123`  ## Editor PDF Generator API comes with a powerful drag & drop editor that allows to create any kind of document templates, from barcode labels to invoices, quotes and reports. You can find tutorials and videos from our [Support Portal](https://support.pdfgeneratorapi.com). * [Component specification](https://support.pdfgeneratorapi.com/en/category/components-1ffseaj/) * [Expression Language documentation](https://support.pdfgeneratorapi.com/en/category/expression-language-q203pa/) * [Frequently asked questions and answers](https://support.pdfgeneratorapi.com/en/category/qanda-1ov519d/)  ## Definitions  ### Organization Organization is a group of workspaces owned by your account.  ### Workspace Workspace contains templates. Each workspace has access to their own templates and organization default templates.  ### Master Workspace Master Workspace is the main/default workspace of your Organization. The Master Workspace identifier is the email you signed up with.  ### Default Template Default template is a template that is available for all workspaces by default. You can set the template access type under Page Setup. If template has \"Organization\" access then your users can use them from the \"New\" menu in the Editor.  ### Data Field Data Field is a placeholder for the specific data in your JSON data set. In this example JSON you can access the buyer name using Data Field `{paymentDetails::buyerName}`. The separator between depth levels is :: (two colons). When designing the template you don’t have to know every Data Field, our editor automatically extracts all the available fields from your data set and provides an easy way to insert them into the template. ``` {     \"documentNumber\": 1,     \"paymentDetails\": {         \"method\": \"Credit Card\",         \"buyerName\": \"John Smith\"     },     \"items\": [         {             \"id\": 1,             \"name\": \"Item one\"         }     ] } ```  ## Rate limiting Our API endpoints use IP-based rate limiting and allow you to make up to 2 requests per second and 60 requests per minute. If you make more requests, you will receive a response with HTTP code 429.  Response headers contain additional values:  | Header   | Description                    | |--------|--------------------------------| | X-RateLimit-Limit    | Maximum requests per minute                   | | X-RateLimit-Remaining    | The requests remaining in the current minute               | | Retry-After     | How many seconds you need to wait until you are allowed to make requests |  *  *  *  *  *  # Libraries and SDKs ## Postman Collection We have created a [Postman Collection](https://www.postman.com/pdfgeneratorapi/workspace/pdf-generator-api-public-workspace/overview) so you can easily test all the API endpoints without developing and code.   ## Client Libraries All our Client Libraries are auto-generated using [OpenAPI Generator](https://openapi-generator.tech/) which uses the OpenAPI v3 specification to automatically generate a client library in specific programming language.  * [PHP Client](https://github.com/pdfgeneratorapi/php-client) * [Java Client](https://github.com/pdfgeneratorapi/java-client) * [Ruby Client](https://github.com/pdfgeneratorapi/ruby-client) * [Python Client](https://github.com/pdfgeneratorapi/python-client) * [Javascript Client](https://github.com/pdfgeneratorapi/javascript-client)  We have validated the generated libraries, but let us know if you find any anomalies in the client code.  ## Model Context Protocol (MCP) Server Integrate document generation directly into your AI agents and LLM applications using our official Model Context Protocol (MCP) Server.  The MCP server provides a standardized interface that allows AI assistants (like Claude Desktop and other MCP-compatible clients) to securely interact with the PDF Generator API. With it, your AI applications can automatically fetch workspaces, retrieve templates, merge data, and generate PDF documents on the fly.  [Get PDF Generator API MCP Server](https://github.com/pdfgeneratorapi/mcp-server) *  *  *  *  *   # Authentication The PDF Generator API uses __JSON Web Tokens (JWT)__ to authenticate all API requests. These tokens offer a method to establish secure server-to-server authentication by transferring a compact JSON object with a signed payload of your account’s API Key and Secret. When authenticating to the PDF Generator API, a JWT should be generated uniquely by a __server-side application__ and included as a __Bearer Token__ in the header of each request.   <SecurityDefinitions />  ## Accessing your API Key and Secret You can find your __API Key__ and __API Secret__ from the __Account Settings__ page after you login to PDF Generator API [here](https://pdfgeneratorapi.com/login).  ## Creating a JWT JSON Web Tokens are composed of three sections: a header, a payload (containing a claim set), and a signature. The header and payload are JSON objects, which are serialized to UTF-8 bytes, then encoded using base64url encoding.  The JWT's header, payload, and signature are concatenated with periods (.). As a result, a JWT typically takes the following form: ``` {Base64url encoded header}.{Base64url encoded payload}.{Base64url encoded signature} ```  We recommend and support libraries provided on [jwt.io](https://jwt.io/). While other libraries can create JWT, these recommended libraries are the most robust.  ### Header Property `alg` defines which signing algorithm is being used. PDF Generator API users HS256. Property `typ` defines the type of token and it is always JWT. ``` {   \"alg\": \"HS256\",   \"typ\": \"JWT\" } ```  ### Payload The second part of the token is the payload, which contains the claims  or the pieces of information being passed about the user and any metadata required. It is mandatory to specify the following claims: * issuer (`iss`): Your API key * subject (`sub`): Workspace identifier * expiration time (`exp`): Timestamp (unix epoch time) until the token is valid. It is highly recommended to set the exp timestamp for a short period, i.e. a matter of seconds. This way, if a token is intercepted or shared, the token will only be valid for a short period of time.  ``` {   \"iss\": \"ad54aaff89ffdfeff178bb8a8f359b29fcb20edb56250b9f584aa2cb0162ed4a\",   \"sub\": \"demo.example@actualreports.com\",   \"exp\": 1586112639 } ```  ### Payload for Partners Our partners can send their unique identifier (provided by us) in JWT's partner_id claim. If the `partner_id` value is specified in the JWT, the organization making the request is automatically connected to the partner account. * Partner ID (`partner_id`): Unique identifier provide by PDF Generator API team  ``` {   \"iss\": \"ad54aaff89ffdfeff178bb8a8f359b29fcb20edb56250b9f584aa2cb0162ed4a\",   \"sub\": \"demo.example@actualreports.com\",   \"partner_id\": \"my-partner-identifier\",   \"exp\": 1586112639 } ```  ### Signature To create the signature part you have to take the encoded header, the encoded payload, a secret, the algorithm specified in the header, and sign that. The signature is used to verify the message wasn't changed along the way, and, in the case of tokens signed with a private key, it can also verify that the sender of the JWT is who it says it is. ``` HMACSHA256(     base64UrlEncode(header) + \".\" +     base64UrlEncode(payload),     API_SECRET) ```  ### Putting all together The output is three Base64-URL strings separated by dots. The following shows a JWT that has the previous header and payload encoded, and it is signed with a secret. ``` eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhZDU0YWFmZjg5ZmZkZmVmZjE3OGJiOGE4ZjM1OWIyOWZjYjIwZWRiNTYyNTBiOWY1ODRhYTJjYjAxNjJlZDRhIiwic3ViIjoiZGVtby5leGFtcGxlQGFjdHVhbHJlcG9ydHMuY29tIn0.SxO-H7UYYYsclS8RGWO1qf0z1cB1m73wF9FLl9RCc1Q  // Base64 encoded header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9 // Base64 encoded payload: eyJpc3MiOiJhZDU0YWFmZjg5ZmZkZmVmZjE3OGJiOGE4ZjM1OWIyOWZjYjIwZWRiNTYyNTBiOWY1ODRhYTJjYjAxNjJlZDRhIiwic3ViIjoiZGVtby5leGFtcGxlQGFjdHVhbHJlcG9ydHMuY29tIn0 // Signature: SxO-H7UYYYsclS8RGWO1qf0z1cB1m73wF9FLl9RCc1Q ```  ## Temporary JWTs You can create a temporary token in [Account Settings](https://pdfgeneratorapi.com/account/organization) page after you login to PDF Generator API. The generated token uses your email address as the subject (`sub`) value and is valid for __15 minutes__. You can also use [jwt.io](https://jwt.io/) to generate test tokens for your API calls. These test tokens should never be used in production applications. *  *  *  *  *  # Error codes  | Code   | Description                    | |--------|--------------------------------| | 401    | Unauthorized                   | | 402    | Payment Required               | | 403    | Forbidden                      | | 404    | Not Found                      | | 422    | Unprocessable Entity           | | 429    | Too Many Requests              | | 500    | Internal Server Error          |  ## 401 Unauthorized | Description                                                             | |-------------------------------------------------------------------------| | Authentication failed: request expired                                  | | Authentication failed: workspace missing                                | | Authentication failed: key missing                                      | | Authentication failed: property 'iss' (issuer) missing in JWT           | | Authentication failed: property 'sub' (subject) missing in JWT          | | Authentication failed: property 'exp' (expiration time) missing in JWT  | | Authentication failed: incorrect signature                              |  ## 402 Payment Required | Description                                                             | |-------------------------------------------------------------------------| | Your account is suspended, please upgrade your account                  |  ## 403 Forbidden | Description                                                             | |-------------------------------------------------------------------------| | Your account has exceeded the monthly document generation limit.        | | Access not granted: You cannot delete master workspace via API          | | Access not granted: Template is not accessible by this organization     | | Your session has expired, please close and reopen the editor.           |  ## 404 Entity not found | Description                                                             | |-------------------------------------------------------------------------| | Entity not found                                                        | | Resource not found                                                      | | None of the templates is available for the workspace.                   |  ## 422 Unprocessable Entity | Description                                                             | |-------------------------------------------------------------------------| | Unable to parse JSON, please check formatting                           | | Required parameter missing                                              | | Required parameter missing: template definition not defined             | | Required parameter missing: template not defined                        |  ## 429 Too Many Requests | Description                                                             | |-------------------------------------------------------------------------| | You can make up to 2 requests per second and 60 requests per minute.   |  *  *  *  *  * 

API version: 4.0.31
Contact: support@pdfgeneratorapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pdfgeneratorapi

import (
	"encoding/json"
)

// checks if the ConvertURL2PDFRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvertURL2PDFRequest{}

// ConvertURL2PDFRequest struct for ConvertURL2PDFRequest
type ConvertURL2PDFRequest struct {
	// Public URL
	Url *string `json:"url,omitempty"`
	// PDF page size
	PaperSize *string `json:"paper_size,omitempty"`
	Orientation *string `json:"orientation,omitempty"`
	// Output
	Output *string `json:"output,omitempty"`
	// Document name
	Filename *string `json:"filename,omitempty"`
	// CSS selector of an element to wait for before the page is printed; see `wait_for_css_state`
	WaitForCss *string `json:"wait_for_css,omitempty"`
	// How `wait_for_css` matches. `present` waits until the element exists, even if it is empty or hidden (e.g. a `<div id=\"report-ready\"></div>` marker); `visible` waits until it is also visible with a non-empty size (e.g. a chart that appears when data has loaded).
	WaitForCssState *string `json:"wait_for_css_state,omitempty"`
	// JavaScript expression evaluated in the page every 100ms; the page is printed once it is truthy. Promises are awaited, and an expression that throws counts as not ready yet.
	WaitForJs *string `json:"wait_for_js,omitempty"`
	// Seconds to wait for `wait_for_css` / `wait_for_js` before applying `on_wait_timeout`
	WaitTimeout *int32 `json:"wait_timeout,omitempty"`
	// What to do when the wait signal is not met within `wait_timeout`. `error` returns a 422 response; `print` prints the page as it is.
	OnWaitTimeout *string `json:"on_wait_timeout,omitempty"`
}

// NewConvertURL2PDFRequest instantiates a new ConvertURL2PDFRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertURL2PDFRequest() *ConvertURL2PDFRequest {
	this := ConvertURL2PDFRequest{}
	var paperSize string = "a4"
	this.PaperSize = &paperSize
	var orientation string = "portrait"
	this.Orientation = &orientation
	var output string = "base64"
	this.Output = &output
	var waitForCssState string = "present"
	this.WaitForCssState = &waitForCssState
	var waitTimeout int32 = 30
	this.WaitTimeout = &waitTimeout
	var onWaitTimeout string = "error"
	this.OnWaitTimeout = &onWaitTimeout
	return &this
}

// NewConvertURL2PDFRequestWithDefaults instantiates a new ConvertURL2PDFRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertURL2PDFRequestWithDefaults() *ConvertURL2PDFRequest {
	this := ConvertURL2PDFRequest{}
	var paperSize string = "a4"
	this.PaperSize = &paperSize
	var orientation string = "portrait"
	this.Orientation = &orientation
	var output string = "base64"
	this.Output = &output
	var waitForCssState string = "present"
	this.WaitForCssState = &waitForCssState
	var waitTimeout int32 = 30
	this.WaitTimeout = &waitTimeout
	var onWaitTimeout string = "error"
	this.OnWaitTimeout = &onWaitTimeout
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ConvertURL2PDFRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertURL2PDFRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ConvertURL2PDFRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ConvertURL2PDFRequest) SetUrl(v string) {
	o.Url = &v
}

// GetPaperSize returns the PaperSize field value if set, zero value otherwise.
func (o *ConvertURL2PDFRequest) GetPaperSize() string {
	if o == nil || IsNil(o.PaperSize) {
		var ret string
		return ret
	}
	return *o.PaperSize
}

// GetPaperSizeOk returns a tuple with the PaperSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertURL2PDFRequest) GetPaperSizeOk() (*string, bool) {
	if o == nil || IsNil(o.PaperSize) {
		return nil, false
	}
	return o.PaperSize, true
}

// HasPaperSize returns a boolean if a field has been set.
func (o *ConvertURL2PDFRequest) HasPaperSize() bool {
	if o != nil && !IsNil(o.PaperSize) {
		return true
	}

	return false
}

// SetPaperSize gets a reference to the given string and assigns it to the PaperSize field.
func (o *ConvertURL2PDFRequest) SetPaperSize(v string) {
	o.PaperSize = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *ConvertURL2PDFRequest) GetOrientation() string {
	if o == nil || IsNil(o.Orientation) {
		var ret string
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertURL2PDFRequest) GetOrientationOk() (*string, bool) {
	if o == nil || IsNil(o.Orientation) {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *ConvertURL2PDFRequest) HasOrientation() bool {
	if o != nil && !IsNil(o.Orientation) {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given string and assigns it to the Orientation field.
func (o *ConvertURL2PDFRequest) SetOrientation(v string) {
	o.Orientation = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *ConvertURL2PDFRequest) GetOutput() string {
	if o == nil || IsNil(o.Output) {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertURL2PDFRequest) GetOutputOk() (*string, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *ConvertURL2PDFRequest) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *ConvertURL2PDFRequest) SetOutput(v string) {
	o.Output = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *ConvertURL2PDFRequest) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertURL2PDFRequest) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *ConvertURL2PDFRequest) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *ConvertURL2PDFRequest) SetFilename(v string) {
	o.Filename = &v
}

// GetWaitForCss returns the WaitForCss field value if set, zero value otherwise.
func (o *ConvertURL2PDFRequest) GetWaitForCss() string {
	if o == nil || IsNil(o.WaitForCss) {
		var ret string
		return ret
	}
	return *o.WaitForCss
}

// GetWaitForCssOk returns a tuple with the WaitForCss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertURL2PDFRequest) GetWaitForCssOk() (*string, bool) {
	if o == nil || IsNil(o.WaitForCss) {
		return nil, false
	}
	return o.WaitForCss, true
}

// HasWaitForCss returns a boolean if a field has been set.
func (o *ConvertURL2PDFRequest) HasWaitForCss() bool {
	if o != nil && !IsNil(o.WaitForCss) {
		return true
	}

	return false
}

// SetWaitForCss gets a reference to the given string and assigns it to the WaitForCss field.
func (o *ConvertURL2PDFRequest) SetWaitForCss(v string) {
	o.WaitForCss = &v
}

// GetWaitForCssState returns the WaitForCssState field value if set, zero value otherwise.
func (o *ConvertURL2PDFRequest) GetWaitForCssState() string {
	if o == nil || IsNil(o.WaitForCssState) {
		var ret string
		return ret
	}
	return *o.WaitForCssState
}

// GetWaitForCssStateOk returns a tuple with the WaitForCssState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertURL2PDFRequest) GetWaitForCssStateOk() (*string, bool) {
	if o == nil || IsNil(o.WaitForCssState) {
		return nil, false
	}
	return o.WaitForCssState, true
}

// HasWaitForCssState returns a boolean if a field has been set.
func (o *ConvertURL2PDFRequest) HasWaitForCssState() bool {
	if o != nil && !IsNil(o.WaitForCssState) {
		return true
	}

	return false
}

// SetWaitForCssState gets a reference to the given string and assigns it to the WaitForCssState field.
func (o *ConvertURL2PDFRequest) SetWaitForCssState(v string) {
	o.WaitForCssState = &v
}

// GetWaitForJs returns the WaitForJs field value if set, zero value otherwise.
func (o *ConvertURL2PDFRequest) GetWaitForJs() string {
	if o == nil || IsNil(o.WaitForJs) {
		var ret string
		return ret
	}
	return *o.WaitForJs
}

// GetWaitForJsOk returns a tuple with the WaitForJs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertURL2PDFRequest) GetWaitForJsOk() (*string, bool) {
	if o == nil || IsNil(o.WaitForJs) {
		return nil, false
	}
	return o.WaitForJs, true
}

// HasWaitForJs returns a boolean if a field has been set.
func (o *ConvertURL2PDFRequest) HasWaitForJs() bool {
	if o != nil && !IsNil(o.WaitForJs) {
		return true
	}

	return false
}

// SetWaitForJs gets a reference to the given string and assigns it to the WaitForJs field.
func (o *ConvertURL2PDFRequest) SetWaitForJs(v string) {
	o.WaitForJs = &v
}

// GetWaitTimeout returns the WaitTimeout field value if set, zero value otherwise.
func (o *ConvertURL2PDFRequest) GetWaitTimeout() int32 {
	if o == nil || IsNil(o.WaitTimeout) {
		var ret int32
		return ret
	}
	return *o.WaitTimeout
}

// GetWaitTimeoutOk returns a tuple with the WaitTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertURL2PDFRequest) GetWaitTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitTimeout) {
		return nil, false
	}
	return o.WaitTimeout, true
}

// HasWaitTimeout returns a boolean if a field has been set.
func (o *ConvertURL2PDFRequest) HasWaitTimeout() bool {
	if o != nil && !IsNil(o.WaitTimeout) {
		return true
	}

	return false
}

// SetWaitTimeout gets a reference to the given int32 and assigns it to the WaitTimeout field.
func (o *ConvertURL2PDFRequest) SetWaitTimeout(v int32) {
	o.WaitTimeout = &v
}

// GetOnWaitTimeout returns the OnWaitTimeout field value if set, zero value otherwise.
func (o *ConvertURL2PDFRequest) GetOnWaitTimeout() string {
	if o == nil || IsNil(o.OnWaitTimeout) {
		var ret string
		return ret
	}
	return *o.OnWaitTimeout
}

// GetOnWaitTimeoutOk returns a tuple with the OnWaitTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertURL2PDFRequest) GetOnWaitTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.OnWaitTimeout) {
		return nil, false
	}
	return o.OnWaitTimeout, true
}

// HasOnWaitTimeout returns a boolean if a field has been set.
func (o *ConvertURL2PDFRequest) HasOnWaitTimeout() bool {
	if o != nil && !IsNil(o.OnWaitTimeout) {
		return true
	}

	return false
}

// SetOnWaitTimeout gets a reference to the given string and assigns it to the OnWaitTimeout field.
func (o *ConvertURL2PDFRequest) SetOnWaitTimeout(v string) {
	o.OnWaitTimeout = &v
}

func (o ConvertURL2PDFRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvertURL2PDFRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.PaperSize) {
		toSerialize["paper_size"] = o.PaperSize
	}
	if !IsNil(o.Orientation) {
		toSerialize["orientation"] = o.Orientation
	}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.WaitForCss) {
		toSerialize["wait_for_css"] = o.WaitForCss
	}
	if !IsNil(o.WaitForCssState) {
		toSerialize["wait_for_css_state"] = o.WaitForCssState
	}
	if !IsNil(o.WaitForJs) {
		toSerialize["wait_for_js"] = o.WaitForJs
	}
	if !IsNil(o.WaitTimeout) {
		toSerialize["wait_timeout"] = o.WaitTimeout
	}
	if !IsNil(o.OnWaitTimeout) {
		toSerialize["on_wait_timeout"] = o.OnWaitTimeout
	}
	return toSerialize, nil
}

type NullableConvertURL2PDFRequest struct {
	value *ConvertURL2PDFRequest
	isSet bool
}

func (v NullableConvertURL2PDFRequest) Get() *ConvertURL2PDFRequest {
	return v.value
}

func (v *NullableConvertURL2PDFRequest) Set(val *ConvertURL2PDFRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertURL2PDFRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertURL2PDFRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertURL2PDFRequest(val *ConvertURL2PDFRequest) *NullableConvertURL2PDFRequest {
	return &NullableConvertURL2PDFRequest{value: val, isSet: true}
}

func (v NullableConvertURL2PDFRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertURL2PDFRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


