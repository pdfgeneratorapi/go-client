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

// checks if the InlineObject14ResponseValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject14ResponseValue{}

// InlineObject14ResponseValue struct for InlineObject14ResponseValue
type InlineObject14ResponseValue struct {
	// Unique field identifier
	Id *string `json:"id,omitempty"`
	// Field type
	Type *string `json:"type,omitempty"`
	// Field name
	Name *string `json:"name,omitempty"`
	Value *InlineObject14ResponseValueValue `json:"value,omitempty"`
	// Whether the field is locked
	Locked *bool `json:"locked,omitempty"`
	// Pages where the field appears
	Pages []int32 `json:"pages,omitempty"`
	Default *InlineObject14ResponseValueDefault `json:"default,omitempty"`
	// Default values for multi-select fields
	Defaults []string `json:"defaults,omitempty"`
	// Selected values for multi-select fields
	Values []string `json:"values,omitempty"`
	// Available options for select fields
	Options []string `json:"options,omitempty"`
	// Field format (for date fields)
	Format *string `json:"format,omitempty"`
	// Maximum field length
	Maxlen *int32 `json:"maxlen,omitempty"`
	// Whether text field is multiline
	Multiline *bool `json:"multiline,omitempty"`
	// Whether combo box is editable
	Editable *bool `json:"editable,omitempty"`
	// Whether list box allows multiple selections
	Multi *bool `json:"multi,omitempty"`
}

// NewInlineObject14ResponseValue instantiates a new InlineObject14ResponseValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject14ResponseValue() *InlineObject14ResponseValue {
	this := InlineObject14ResponseValue{}
	return &this
}

// NewInlineObject14ResponseValueWithDefaults instantiates a new InlineObject14ResponseValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject14ResponseValueWithDefaults() *InlineObject14ResponseValue {
	this := InlineObject14ResponseValue{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineObject14ResponseValue) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineObject14ResponseValue) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject14ResponseValue) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetValue() InlineObject14ResponseValueValue {
	if o == nil || IsNil(o.Value) {
		var ret InlineObject14ResponseValueValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetValueOk() (*InlineObject14ResponseValueValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given InlineObject14ResponseValueValue and assigns it to the Value field.
func (o *InlineObject14ResponseValue) SetValue(v InlineObject14ResponseValueValue) {
	o.Value = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *InlineObject14ResponseValue) SetLocked(v bool) {
	o.Locked = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetPages() []int32 {
	if o == nil || IsNil(o.Pages) {
		var ret []int32
		return ret
	}
	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetPagesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given []int32 and assigns it to the Pages field.
func (o *InlineObject14ResponseValue) SetPages(v []int32) {
	o.Pages = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetDefault() InlineObject14ResponseValueDefault {
	if o == nil || IsNil(o.Default) {
		var ret InlineObject14ResponseValueDefault
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetDefaultOk() (*InlineObject14ResponseValueDefault, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given InlineObject14ResponseValueDefault and assigns it to the Default field.
func (o *InlineObject14ResponseValue) SetDefault(v InlineObject14ResponseValueDefault) {
	o.Default = &v
}

// GetDefaults returns the Defaults field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetDefaults() []string {
	if o == nil || IsNil(o.Defaults) {
		var ret []string
		return ret
	}
	return o.Defaults
}

// GetDefaultsOk returns a tuple with the Defaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetDefaultsOk() ([]string, bool) {
	if o == nil || IsNil(o.Defaults) {
		return nil, false
	}
	return o.Defaults, true
}

// HasDefaults returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasDefaults() bool {
	if o != nil && !IsNil(o.Defaults) {
		return true
	}

	return false
}

// SetDefaults gets a reference to the given []string and assigns it to the Defaults field.
func (o *InlineObject14ResponseValue) SetDefaults(v []string) {
	o.Defaults = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *InlineObject14ResponseValue) SetValues(v []string) {
	o.Values = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetOptions() []string {
	if o == nil || IsNil(o.Options) {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *InlineObject14ResponseValue) SetOptions(v []string) {
	o.Options = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *InlineObject14ResponseValue) SetFormat(v string) {
	o.Format = &v
}

// GetMaxlen returns the Maxlen field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetMaxlen() int32 {
	if o == nil || IsNil(o.Maxlen) {
		var ret int32
		return ret
	}
	return *o.Maxlen
}

// GetMaxlenOk returns a tuple with the Maxlen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetMaxlenOk() (*int32, bool) {
	if o == nil || IsNil(o.Maxlen) {
		return nil, false
	}
	return o.Maxlen, true
}

// HasMaxlen returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasMaxlen() bool {
	if o != nil && !IsNil(o.Maxlen) {
		return true
	}

	return false
}

// SetMaxlen gets a reference to the given int32 and assigns it to the Maxlen field.
func (o *InlineObject14ResponseValue) SetMaxlen(v int32) {
	o.Maxlen = &v
}

// GetMultiline returns the Multiline field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetMultiline() bool {
	if o == nil || IsNil(o.Multiline) {
		var ret bool
		return ret
	}
	return *o.Multiline
}

// GetMultilineOk returns a tuple with the Multiline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetMultilineOk() (*bool, bool) {
	if o == nil || IsNil(o.Multiline) {
		return nil, false
	}
	return o.Multiline, true
}

// HasMultiline returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasMultiline() bool {
	if o != nil && !IsNil(o.Multiline) {
		return true
	}

	return false
}

// SetMultiline gets a reference to the given bool and assigns it to the Multiline field.
func (o *InlineObject14ResponseValue) SetMultiline(v bool) {
	o.Multiline = &v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetEditable() bool {
	if o == nil || IsNil(o.Editable) {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetEditableOk() (*bool, bool) {
	if o == nil || IsNil(o.Editable) {
		return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasEditable() bool {
	if o != nil && !IsNil(o.Editable) {
		return true
	}

	return false
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *InlineObject14ResponseValue) SetEditable(v bool) {
	o.Editable = &v
}

// GetMulti returns the Multi field value if set, zero value otherwise.
func (o *InlineObject14ResponseValue) GetMulti() bool {
	if o == nil || IsNil(o.Multi) {
		var ret bool
		return ret
	}
	return *o.Multi
}

// GetMultiOk returns a tuple with the Multi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject14ResponseValue) GetMultiOk() (*bool, bool) {
	if o == nil || IsNil(o.Multi) {
		return nil, false
	}
	return o.Multi, true
}

// HasMulti returns a boolean if a field has been set.
func (o *InlineObject14ResponseValue) HasMulti() bool {
	if o != nil && !IsNil(o.Multi) {
		return true
	}

	return false
}

// SetMulti gets a reference to the given bool and assigns it to the Multi field.
func (o *InlineObject14ResponseValue) SetMulti(v bool) {
	o.Multi = &v
}

func (o InlineObject14ResponseValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject14ResponseValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.Pages) {
		toSerialize["pages"] = o.Pages
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Defaults) {
		toSerialize["defaults"] = o.Defaults
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Maxlen) {
		toSerialize["maxlen"] = o.Maxlen
	}
	if !IsNil(o.Multiline) {
		toSerialize["multiline"] = o.Multiline
	}
	if !IsNil(o.Editable) {
		toSerialize["editable"] = o.Editable
	}
	if !IsNil(o.Multi) {
		toSerialize["multi"] = o.Multi
	}
	return toSerialize, nil
}

type NullableInlineObject14ResponseValue struct {
	value *InlineObject14ResponseValue
	isSet bool
}

func (v NullableInlineObject14ResponseValue) Get() *InlineObject14ResponseValue {
	return v.value
}

func (v *NullableInlineObject14ResponseValue) Set(val *InlineObject14ResponseValue) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject14ResponseValue) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject14ResponseValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject14ResponseValue(val *InlineObject14ResponseValue) *NullableInlineObject14ResponseValue {
	return &NullableInlineObject14ResponseValue{value: val, isSet: true}
}

func (v NullableInlineObject14ResponseValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject14ResponseValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


