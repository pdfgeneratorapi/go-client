/*
PDF Generator API

# Introduction [PDF Generator API](https://pdfgeneratorapi.com) allows you easily generate transactional PDF documents and reduce the development and support costs by enabling your users to create and manage their document templates using a browser-based drag-and-drop document editor.  The PDF Generator API features a web API architecture, allowing you to code in the language of your choice. This API supports the JSON media type, and uses UTF-8 character encoding.  ## Base URL The base URL for all the API endpoints is `https://us1.pdfgeneratorapi.com/api/v4`  For example * `https://us1.pdfgeneratorapi.com/api/v4/templates` * `https://us1.pdfgeneratorapi.com/api/v4/workspaces` * `https://us1.pdfgeneratorapi.com/api/v4/templates/123123`  ## Editor PDF Generator API comes with a powerful drag & drop editor that allows to create any kind of document templates, from barcode labels to invoices, quotes and reports. You can find tutorials and videos from our [Support Portal](https://support.pdfgeneratorapi.com). * [Component specification](https://support.pdfgeneratorapi.com/en/category/components-1ffseaj/) * [Expression Language documentation](https://support.pdfgeneratorapi.com/en/category/expression-language-q203pa/) * [Frequently asked questions and answers](https://support.pdfgeneratorapi.com/en/category/qanda-1ov519d/)  ## Definitions  ### Organization Organization is a group of workspaces owned by your account.  ### Workspace Workspace contains templates. Each workspace has access to their own templates and organization default templates.  ### Master Workspace Master Workspace is the main/default workspace of your Organization. The Master Workspace identifier is the email you signed up with.  ### Default Template Default template is a template that is available for all workspaces by default. You can set the template access type under Page Setup. If template has \"Organization\" access then your users can use them from the \"New\" menu in the Editor.  ### Data Field Data Field is a placeholder for the specific data in your JSON data set. In this example JSON you can access the buyer name using Data Field `{paymentDetails::buyerName}`. The separator between depth levels is :: (two colons). When designing the template you don’t have to know every Data Field, our editor automatically extracts all the available fields from your data set and provides an easy way to insert them into the template. ``` {     \"documentNumber\": 1,     \"paymentDetails\": {         \"method\": \"Credit Card\",         \"buyerName\": \"John Smith\"     },     \"items\": [         {             \"id\": 1,             \"name\": \"Item one\"         }     ] } ```  ## Rate limiting Our API endpoints use IP-based rate limiting and allow you to make up to 2 requests per second and 60 requests per minute. If you make more requests, you will receive a response with HTTP code 429.  Response headers contain additional values:  | Header   | Description                    | |--------|--------------------------------| | X-RateLimit-Limit    | Maximum requests per minute                   | | X-RateLimit-Remaining    | The requests remaining in the current minute               | | Retry-After     | How many seconds you need to wait until you are allowed to make requests |  *  *  *  *  *  # Libraries and SDKs ## Postman Collection We have created a [Postman Collection](https://www.postman.com/pdfgeneratorapi/workspace/pdf-generator-api-public-workspace/overview) so you can easily test all the API endpoints without developing and code.   ## Client Libraries All our Client Libraries are auto-generated using [OpenAPI Generator](https://openapi-generator.tech/) which uses the OpenAPI v3 specification to automatically generate a client library in specific programming language.  * [PHP Client](https://github.com/pdfgeneratorapi/php-client) * [Java Client](https://github.com/pdfgeneratorapi/java-client) * [Ruby Client](https://github.com/pdfgeneratorapi/ruby-client) * [Python Client](https://github.com/pdfgeneratorapi/python-client) * [Javascript Client](https://github.com/pdfgeneratorapi/javascript-client)  We have validated the generated libraries, but let us know if you find any anomalies in the client code.  ## Model Context Protocol (MCP) Server Integrate document generation directly into your AI agents and LLM applications using our official Model Context Protocol (MCP) Server.  The MCP server provides a standardized interface that allows AI assistants (like Claude Desktop and other MCP-compatible clients) to securely interact with the PDF Generator API. With it, your AI applications can automatically fetch workspaces, retrieve templates, merge data, and generate PDF documents on the fly.  [Get PDF Generator API MCP Server](https://github.com/pdfgeneratorapi/mcp-server) *  *  *  *  *   # Authentication The PDF Generator API uses __JSON Web Tokens (JWT)__ to authenticate all API requests. These tokens offer a method to establish secure server-to-server authentication by transferring a compact JSON object with a signed payload of your account’s API Key and Secret. When authenticating to the PDF Generator API, a JWT should be generated uniquely by a __server-side application__ and included as a __Bearer Token__ in the header of each request.   <SecurityDefinitions />  ## Accessing your API Key and Secret You can find your __API Key__ and __API Secret__ from the __Account Settings__ page after you login to PDF Generator API [here](https://pdfgeneratorapi.com/login).  ## Creating a JWT JSON Web Tokens are composed of three sections: a header, a payload (containing a claim set), and a signature. The header and payload are JSON objects, which are serialized to UTF-8 bytes, then encoded using base64url encoding.  The JWT's header, payload, and signature are concatenated with periods (.). As a result, a JWT typically takes the following form: ``` {Base64url encoded header}.{Base64url encoded payload}.{Base64url encoded signature} ```  We recommend and support libraries provided on [jwt.io](https://jwt.io/). While other libraries can create JWT, these recommended libraries are the most robust.  ### Header Property `alg` defines which signing algorithm is being used. PDF Generator API users HS256. Property `typ` defines the type of token and it is always JWT. ``` {   \"alg\": \"HS256\",   \"typ\": \"JWT\" } ```  ### Payload The second part of the token is the payload, which contains the claims  or the pieces of information being passed about the user and any metadata required. It is mandatory to specify the following claims: * issuer (`iss`): Your API key * subject (`sub`): Workspace identifier * expiration time (`exp`): Timestamp (unix epoch time) until the token is valid. It is highly recommended to set the exp timestamp for a short period, i.e. a matter of seconds. This way, if a token is intercepted or shared, the token will only be valid for a short period of time.  ``` {   \"iss\": \"ad54aaff89ffdfeff178bb8a8f359b29fcb20edb56250b9f584aa2cb0162ed4a\",   \"sub\": \"demo.example@actualreports.com\",   \"exp\": 1586112639 } ```  ### Payload for Partners Our partners can send their unique identifier (provided by us) in JWT's partner_id claim. If the `partner_id` value is specified in the JWT, the organization making the request is automatically connected to the partner account. * Partner ID (`partner_id`): Unique identifier provide by PDF Generator API team  ``` {   \"iss\": \"ad54aaff89ffdfeff178bb8a8f359b29fcb20edb56250b9f584aa2cb0162ed4a\",   \"sub\": \"demo.example@actualreports.com\",   \"partner_id\": \"my-partner-identifier\",   \"exp\": 1586112639 } ```  ### Signature To create the signature part you have to take the encoded header, the encoded payload, a secret, the algorithm specified in the header, and sign that. The signature is used to verify the message wasn't changed along the way, and, in the case of tokens signed with a private key, it can also verify that the sender of the JWT is who it says it is. ``` HMACSHA256(     base64UrlEncode(header) + \".\" +     base64UrlEncode(payload),     API_SECRET) ```  ### Putting all together The output is three Base64-URL strings separated by dots. The following shows a JWT that has the previous header and payload encoded, and it is signed with a secret. ``` eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhZDU0YWFmZjg5ZmZkZmVmZjE3OGJiOGE4ZjM1OWIyOWZjYjIwZWRiNTYyNTBiOWY1ODRhYTJjYjAxNjJlZDRhIiwic3ViIjoiZGVtby5leGFtcGxlQGFjdHVhbHJlcG9ydHMuY29tIn0.SxO-H7UYYYsclS8RGWO1qf0z1cB1m73wF9FLl9RCc1Q  // Base64 encoded header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9 // Base64 encoded payload: eyJpc3MiOiJhZDU0YWFmZjg5ZmZkZmVmZjE3OGJiOGE4ZjM1OWIyOWZjYjIwZWRiNTYyNTBiOWY1ODRhYTJjYjAxNjJlZDRhIiwic3ViIjoiZGVtby5leGFtcGxlQGFjdHVhbHJlcG9ydHMuY29tIn0 // Signature: SxO-H7UYYYsclS8RGWO1qf0z1cB1m73wF9FLl9RCc1Q ```  ## Temporary JWTs You can create a temporary token in [Account Settings](https://pdfgeneratorapi.com/account/organization) page after you login to PDF Generator API. The generated token uses your email address as the subject (`sub`) value and is valid for __15 minutes__. You can also use [jwt.io](https://jwt.io/) to generate test tokens for your API calls. These test tokens should never be used in production applications. *  *  *  *  *  # Error codes  Every error response is a JSON object with a `message` string. Request validation failures (422) also include an `errors` object keyed by the field that failed.  ```json {   \"message\": \"The identifier field is required.\",   \"errors\": {     \"identifier\": [\"The identifier field is required.\"]   } } ```  | Code   | Description                    | |--------|--------------------------------| | 400    | Bad Request                    | | 401    | Unauthorized                   | | 402    | Payment Required               | | 403    | Forbidden                      | | 404    | Not Found                      | | 405    | Method Not Allowed             | | 409    | Conflict                       | | 410    | Gone                           | | 422    | Unprocessable Entity           | | 429    | Too Many Requests              | | 500    | Internal Server Error          | | 503    | Service Unavailable            |  ## 401 Unauthorized | Description                                                                       | |-----------------------------------------------------------------------------------| | Authentication failed: JWT Token missing.                                         | | Authentication failed: Malformed JWT Token.                                       | | Authentication failed: Expired JWT Token.                                         | | Authentication failed: Invalid workspace.                                         | | Authentication failed: Invalid organization.                                      | | Authentication failed: This application no longer has access to the workspace.   | | Authentication failed: The sub field is required. (also `iss` and `exp`)          |  ## 402 Payment Required | Description                                                                       | |-----------------------------------------------------------------------------------| | Your account is suspended, please upgrade your account or contact support@pdfgeneratorapi.com. |  ## 403 Forbidden | Description                                                                       | |-----------------------------------------------------------------------------------| | You don't have permission to access this content. Please get in touch with your organization administrator to update your access permissions. | | This application was not granted permission for this request.                    | | Deleting own user is not allowed.                                                 | | Deleting organization owner is not allowed.                                       |  ## 404 Not Found | Description                                                                       | |-----------------------------------------------------------------------------------| | Entity not found.                                                                 | | Route not found.                                                                  | | Workspace doesn’t have access to read the template.                               |  ## 422 Unprocessable Entity | Description                                                                       | |-----------------------------------------------------------------------------------| | Unable to parse JSON, please check formatting.                                    | | Template id should be a numeric value.                                            | | The per page must be at least 1. (with `errors.per_page`)                         |  ## 429 Too Many Requests | Description                                                                       | |-----------------------------------------------------------------------------------| | Per-second rate limit exceeded. Please slow down your requests.                   | | Per-minute rate limit exceeded. Please try again in a minute.                     | | Global per-second rate limit exceeded. Please slow down your requests.            | | Global per-minute rate limit exceeded. Please try again in a minute.              |  *  *  *  *  * 

API version: 4.0.35
Contact: support@pdfgeneratorapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pdfgeneratorapi

import (
	"encoding/json"
)

// checks if the InlineObject11MetaStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject11MetaStats{}

// InlineObject11MetaStats struct for InlineObject11MetaStats
type InlineObject11MetaStats struct {
	// Document original size in kilobytes
	OriginalSize *float32 `json:"original_size,omitempty"`
	// Document size after optimization in kilobytes
	OptimizedSize *float32 `json:"optimized_size,omitempty"`
	// Unit of the sizes
	Unit *string `json:"unit,omitempty"`
}

// NewInlineObject11MetaStats instantiates a new InlineObject11MetaStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject11MetaStats() *InlineObject11MetaStats {
	this := InlineObject11MetaStats{}
	return &this
}

// NewInlineObject11MetaStatsWithDefaults instantiates a new InlineObject11MetaStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject11MetaStatsWithDefaults() *InlineObject11MetaStats {
	this := InlineObject11MetaStats{}
	return &this
}

// GetOriginalSize returns the OriginalSize field value if set, zero value otherwise.
func (o *InlineObject11MetaStats) GetOriginalSize() float32 {
	if o == nil || IsNil(o.OriginalSize) {
		var ret float32
		return ret
	}
	return *o.OriginalSize
}

// GetOriginalSizeOk returns a tuple with the OriginalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject11MetaStats) GetOriginalSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.OriginalSize) {
		return nil, false
	}
	return o.OriginalSize, true
}

// HasOriginalSize returns a boolean if a field has been set.
func (o *InlineObject11MetaStats) HasOriginalSize() bool {
	if o != nil && !IsNil(o.OriginalSize) {
		return true
	}

	return false
}

// SetOriginalSize gets a reference to the given float32 and assigns it to the OriginalSize field.
func (o *InlineObject11MetaStats) SetOriginalSize(v float32) {
	o.OriginalSize = &v
}

// GetOptimizedSize returns the OptimizedSize field value if set, zero value otherwise.
func (o *InlineObject11MetaStats) GetOptimizedSize() float32 {
	if o == nil || IsNil(o.OptimizedSize) {
		var ret float32
		return ret
	}
	return *o.OptimizedSize
}

// GetOptimizedSizeOk returns a tuple with the OptimizedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject11MetaStats) GetOptimizedSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.OptimizedSize) {
		return nil, false
	}
	return o.OptimizedSize, true
}

// HasOptimizedSize returns a boolean if a field has been set.
func (o *InlineObject11MetaStats) HasOptimizedSize() bool {
	if o != nil && !IsNil(o.OptimizedSize) {
		return true
	}

	return false
}

// SetOptimizedSize gets a reference to the given float32 and assigns it to the OptimizedSize field.
func (o *InlineObject11MetaStats) SetOptimizedSize(v float32) {
	o.OptimizedSize = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *InlineObject11MetaStats) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject11MetaStats) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *InlineObject11MetaStats) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *InlineObject11MetaStats) SetUnit(v string) {
	o.Unit = &v
}

func (o InlineObject11MetaStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject11MetaStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OriginalSize) {
		toSerialize["original_size"] = o.OriginalSize
	}
	if !IsNil(o.OptimizedSize) {
		toSerialize["optimized_size"] = o.OptimizedSize
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	return toSerialize, nil
}

type NullableInlineObject11MetaStats struct {
	value *InlineObject11MetaStats
	isSet bool
}

func (v NullableInlineObject11MetaStats) Get() *InlineObject11MetaStats {
	return v.value
}

func (v *NullableInlineObject11MetaStats) Set(val *InlineObject11MetaStats) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject11MetaStats) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject11MetaStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject11MetaStats(val *InlineObject11MetaStats) *NullableInlineObject11MetaStats {
	return &NullableInlineObject11MetaStats{value: val, isSet: true}
}

func (v NullableInlineObject11MetaStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject11MetaStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


