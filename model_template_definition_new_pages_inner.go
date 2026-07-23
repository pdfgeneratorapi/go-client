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

// checks if the TemplateDefinitionNewPagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateDefinitionNewPagesInner{}

// TemplateDefinitionNewPagesInner struct for TemplateDefinitionNewPagesInner
type TemplateDefinitionNewPagesInner struct {
	// Page width in units
	Width *float32 `json:"width,omitempty"`
	// Page height in units
	Height *float32 `json:"height,omitempty"`
	Margins *TemplateDefinitionNewPagesInnerMargins `json:"margins,omitempty"`
	Border *bool `json:"border,omitempty"`
	Components []map[string]interface{} `json:"components,omitempty"`
	// Defines page specific layout which can differ from the main template layout (e.g page format, margins).
	Layout map[string]interface{} `json:"layout,omitempty"`
	ConditionalFormats []map[string]interface{} `json:"conditionalFormats,omitempty"`
	// Defines background image for the page.
	BackgroundImage NullableString `json:"backgroundImage,omitempty"`
}

// NewTemplateDefinitionNewPagesInner instantiates a new TemplateDefinitionNewPagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateDefinitionNewPagesInner() *TemplateDefinitionNewPagesInner {
	this := TemplateDefinitionNewPagesInner{}
	var border bool = false
	this.Border = &border
	return &this
}

// NewTemplateDefinitionNewPagesInnerWithDefaults instantiates a new TemplateDefinitionNewPagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateDefinitionNewPagesInnerWithDefaults() *TemplateDefinitionNewPagesInner {
	this := TemplateDefinitionNewPagesInner{}
	var border bool = false
	this.Border = &border
	return &this
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *TemplateDefinitionNewPagesInner) GetWidth() float32 {
	if o == nil || IsNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNewPagesInner) GetWidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *TemplateDefinitionNewPagesInner) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *TemplateDefinitionNewPagesInner) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *TemplateDefinitionNewPagesInner) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNewPagesInner) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *TemplateDefinitionNewPagesInner) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *TemplateDefinitionNewPagesInner) SetHeight(v float32) {
	o.Height = &v
}

// GetMargins returns the Margins field value if set, zero value otherwise.
func (o *TemplateDefinitionNewPagesInner) GetMargins() TemplateDefinitionNewPagesInnerMargins {
	if o == nil || IsNil(o.Margins) {
		var ret TemplateDefinitionNewPagesInnerMargins
		return ret
	}
	return *o.Margins
}

// GetMarginsOk returns a tuple with the Margins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNewPagesInner) GetMarginsOk() (*TemplateDefinitionNewPagesInnerMargins, bool) {
	if o == nil || IsNil(o.Margins) {
		return nil, false
	}
	return o.Margins, true
}

// HasMargins returns a boolean if a field has been set.
func (o *TemplateDefinitionNewPagesInner) HasMargins() bool {
	if o != nil && !IsNil(o.Margins) {
		return true
	}

	return false
}

// SetMargins gets a reference to the given TemplateDefinitionNewPagesInnerMargins and assigns it to the Margins field.
func (o *TemplateDefinitionNewPagesInner) SetMargins(v TemplateDefinitionNewPagesInnerMargins) {
	o.Margins = &v
}

// GetBorder returns the Border field value if set, zero value otherwise.
func (o *TemplateDefinitionNewPagesInner) GetBorder() bool {
	if o == nil || IsNil(o.Border) {
		var ret bool
		return ret
	}
	return *o.Border
}

// GetBorderOk returns a tuple with the Border field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNewPagesInner) GetBorderOk() (*bool, bool) {
	if o == nil || IsNil(o.Border) {
		return nil, false
	}
	return o.Border, true
}

// HasBorder returns a boolean if a field has been set.
func (o *TemplateDefinitionNewPagesInner) HasBorder() bool {
	if o != nil && !IsNil(o.Border) {
		return true
	}

	return false
}

// SetBorder gets a reference to the given bool and assigns it to the Border field.
func (o *TemplateDefinitionNewPagesInner) SetBorder(v bool) {
	o.Border = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *TemplateDefinitionNewPagesInner) GetComponents() []map[string]interface{} {
	if o == nil || IsNil(o.Components) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNewPagesInner) GetComponentsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *TemplateDefinitionNewPagesInner) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []map[string]interface{} and assigns it to the Components field.
func (o *TemplateDefinitionNewPagesInner) SetComponents(v []map[string]interface{}) {
	o.Components = v
}

// GetLayout returns the Layout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateDefinitionNewPagesInner) GetLayout() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateDefinitionNewPagesInner) GetLayoutOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Layout) {
		return map[string]interface{}{}, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *TemplateDefinitionNewPagesInner) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given map[string]interface{} and assigns it to the Layout field.
func (o *TemplateDefinitionNewPagesInner) SetLayout(v map[string]interface{}) {
	o.Layout = v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *TemplateDefinitionNewPagesInner) GetConditionalFormats() []map[string]interface{} {
	if o == nil || IsNil(o.ConditionalFormats) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNewPagesInner) GetConditionalFormatsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ConditionalFormats) {
		return nil, false
	}
	return o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *TemplateDefinitionNewPagesInner) HasConditionalFormats() bool {
	if o != nil && !IsNil(o.ConditionalFormats) {
		return true
	}

	return false
}

// SetConditionalFormats gets a reference to the given []map[string]interface{} and assigns it to the ConditionalFormats field.
func (o *TemplateDefinitionNewPagesInner) SetConditionalFormats(v []map[string]interface{}) {
	o.ConditionalFormats = v
}

// GetBackgroundImage returns the BackgroundImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateDefinitionNewPagesInner) GetBackgroundImage() string {
	if o == nil || IsNil(o.BackgroundImage.Get()) {
		var ret string
		return ret
	}
	return *o.BackgroundImage.Get()
}

// GetBackgroundImageOk returns a tuple with the BackgroundImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateDefinitionNewPagesInner) GetBackgroundImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackgroundImage.Get(), o.BackgroundImage.IsSet()
}

// HasBackgroundImage returns a boolean if a field has been set.
func (o *TemplateDefinitionNewPagesInner) HasBackgroundImage() bool {
	if o != nil && o.BackgroundImage.IsSet() {
		return true
	}

	return false
}

// SetBackgroundImage gets a reference to the given NullableString and assigns it to the BackgroundImage field.
func (o *TemplateDefinitionNewPagesInner) SetBackgroundImage(v string) {
	o.BackgroundImage.Set(&v)
}
// SetBackgroundImageNil sets the value for BackgroundImage to be an explicit nil
func (o *TemplateDefinitionNewPagesInner) SetBackgroundImageNil() {
	o.BackgroundImage.Set(nil)
}

// UnsetBackgroundImage ensures that no value is present for BackgroundImage, not even an explicit nil
func (o *TemplateDefinitionNewPagesInner) UnsetBackgroundImage() {
	o.BackgroundImage.Unset()
}

func (o TemplateDefinitionNewPagesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateDefinitionNewPagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Margins) {
		toSerialize["margins"] = o.Margins
	}
	if !IsNil(o.Border) {
		toSerialize["border"] = o.Border
	}
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	if !IsNil(o.ConditionalFormats) {
		toSerialize["conditionalFormats"] = o.ConditionalFormats
	}
	if o.BackgroundImage.IsSet() {
		toSerialize["backgroundImage"] = o.BackgroundImage.Get()
	}
	return toSerialize, nil
}

type NullableTemplateDefinitionNewPagesInner struct {
	value *TemplateDefinitionNewPagesInner
	isSet bool
}

func (v NullableTemplateDefinitionNewPagesInner) Get() *TemplateDefinitionNewPagesInner {
	return v.value
}

func (v *NullableTemplateDefinitionNewPagesInner) Set(val *TemplateDefinitionNewPagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateDefinitionNewPagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateDefinitionNewPagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateDefinitionNewPagesInner(val *TemplateDefinitionNewPagesInner) *NullableTemplateDefinitionNewPagesInner {
	return &NullableTemplateDefinitionNewPagesInner{value: val, isSet: true}
}

func (v NullableTemplateDefinitionNewPagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateDefinitionNewPagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


