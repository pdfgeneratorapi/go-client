/*
PDF Generator API

# Introduction [PDF Generator API](https://pdfgeneratorapi.com) allows you easily generate transactional PDF documents and reduce the development and support costs by enabling your users to create and manage their document templates using a browser-based drag-and-drop document editor.  The PDF Generator API features a web API architecture, allowing you to code in the language of your choice. This API supports the JSON media type, and uses UTF-8 character encoding.  ## Base URL The base URL for all the API endpoints is `https://us1.pdfgeneratorapi.com/api/v4`  For example * `https://us1.pdfgeneratorapi.com/api/v4/templates` * `https://us1.pdfgeneratorapi.com/api/v4/workspaces` * `https://us1.pdfgeneratorapi.com/api/v4/templates/123123`  ## Editor PDF Generator API comes with a powerful drag & drop editor that allows to create any kind of document templates, from barcode labels to invoices, quotes and reports. You can find tutorials and videos from our [Support Portal](https://support.pdfgeneratorapi.com). * [Component specification](https://support.pdfgeneratorapi.com/en/category/components-1ffseaj/) * [Expression Language documentation](https://support.pdfgeneratorapi.com/en/category/expression-language-q203pa/) * [Frequently asked questions and answers](https://support.pdfgeneratorapi.com/en/category/qanda-1ov519d/)  ## Definitions  ### Organization Organization is a group of workspaces owned by your account.  ### Workspace Workspace contains templates. Each workspace has access to their own templates and organization default templates.  ### Master Workspace Master Workspace is the main/default workspace of your Organization. The Master Workspace identifier is the email you signed up with.  ### Default Template Default template is a template that is available for all workspaces by default. You can set the template access type under Page Setup. If template has \"Organization\" access then your users can use them from the \"New\" menu in the Editor.  ### Data Field Data Field is a placeholder for the specific data in your JSON data set. In this example JSON you can access the buyer name using Data Field `{paymentDetails::buyerName}`. The separator between depth levels is :: (two colons). When designing the template you don’t have to know every Data Field, our editor automatically extracts all the available fields from your data set and provides an easy way to insert them into the template. ``` {     \"documentNumber\": 1,     \"paymentDetails\": {         \"method\": \"Credit Card\",         \"buyerName\": \"John Smith\"     },     \"items\": [         {             \"id\": 1,             \"name\": \"Item one\"         }     ] } ```  ## Rate limiting Our API endpoints use IP-based rate limiting and allow you to make up to 2 requests per second and 60 requests per minute. If you make more requests, you will receive a response with HTTP code 429.  Response headers contain additional values:  | Header   | Description                    | |--------|--------------------------------| | X-RateLimit-Limit    | Maximum requests per minute                   | | X-RateLimit-Remaining    | The requests remaining in the current minute               | | Retry-After     | How many seconds you need to wait until you are allowed to make requests |  *  *  *  *  *  # Libraries and SDKs ## Postman Collection We have created a [Postman Collection](https://www.postman.com/pdfgeneratorapi/workspace/pdf-generator-api-public-workspace/overview) so you can easily test all the API endpoints without developing and code.   ## Client Libraries All our Client Libraries are auto-generated using [OpenAPI Generator](https://openapi-generator.tech/) which uses the OpenAPI v3 specification to automatically generate a client library in specific programming language.  * [PHP Client](https://github.com/pdfgeneratorapi/php-client) * [Java Client](https://github.com/pdfgeneratorapi/java-client) * [Ruby Client](https://github.com/pdfgeneratorapi/ruby-client) * [Python Client](https://github.com/pdfgeneratorapi/python-client) * [Javascript Client](https://github.com/pdfgeneratorapi/javascript-client)  We have validated the generated libraries, but let us know if you find any anomalies in the client code.  ## Model Context Protocol (MCP) Server Integrate document generation directly into your AI agents and LLM applications using our official Model Context Protocol (MCP) Server.  The MCP server provides a standardized interface that allows AI assistants (like Claude Desktop and other MCP-compatible clients) to securely interact with the PDF Generator API. With it, your AI applications can automatically fetch workspaces, retrieve templates, merge data, and generate PDF documents on the fly.  [Get PDF Generator API MCP Server](https://github.com/pdfgeneratorapi/mcp-server) *  *  *  *  *   # Authentication The PDF Generator API uses __JSON Web Tokens (JWT)__ to authenticate all API requests. These tokens offer a method to establish secure server-to-server authentication by transferring a compact JSON object with a signed payload of your account’s API Key and Secret. When authenticating to the PDF Generator API, a JWT should be generated uniquely by a __server-side application__ and included as a __Bearer Token__ in the header of each request.   <SecurityDefinitions />  ## Accessing your API Key and Secret You can find your __API Key__ and __API Secret__ from the __Account Settings__ page after you login to PDF Generator API [here](https://pdfgeneratorapi.com/login).  ## Creating a JWT JSON Web Tokens are composed of three sections: a header, a payload (containing a claim set), and a signature. The header and payload are JSON objects, which are serialized to UTF-8 bytes, then encoded using base64url encoding.  The JWT's header, payload, and signature are concatenated with periods (.). As a result, a JWT typically takes the following form: ``` {Base64url encoded header}.{Base64url encoded payload}.{Base64url encoded signature} ```  We recommend and support libraries provided on [jwt.io](https://jwt.io/). While other libraries can create JWT, these recommended libraries are the most robust.  ### Header Property `alg` defines which signing algorithm is being used. PDF Generator API users HS256. Property `typ` defines the type of token and it is always JWT. ``` {   \"alg\": \"HS256\",   \"typ\": \"JWT\" } ```  ### Payload The second part of the token is the payload, which contains the claims  or the pieces of information being passed about the user and any metadata required. It is mandatory to specify the following claims: * issuer (`iss`): Your API key * subject (`sub`): Workspace identifier * expiration time (`exp`): Timestamp (unix epoch time) until the token is valid. It is highly recommended to set the exp timestamp for a short period, i.e. a matter of seconds. This way, if a token is intercepted or shared, the token will only be valid for a short period of time.  ``` {   \"iss\": \"ad54aaff89ffdfeff178bb8a8f359b29fcb20edb56250b9f584aa2cb0162ed4a\",   \"sub\": \"demo.example@actualreports.com\",   \"exp\": 1586112639 } ```  ### Payload for Partners Our partners can send their unique identifier (provided by us) in JWT's partner_id claim. If the `partner_id` value is specified in the JWT, the organization making the request is automatically connected to the partner account. * Partner ID (`partner_id`): Unique identifier provide by PDF Generator API team  ``` {   \"iss\": \"ad54aaff89ffdfeff178bb8a8f359b29fcb20edb56250b9f584aa2cb0162ed4a\",   \"sub\": \"demo.example@actualreports.com\",   \"partner_id\": \"my-partner-identifier\",   \"exp\": 1586112639 } ```  ### Signature To create the signature part you have to take the encoded header, the encoded payload, a secret, the algorithm specified in the header, and sign that. The signature is used to verify the message wasn't changed along the way, and, in the case of tokens signed with a private key, it can also verify that the sender of the JWT is who it says it is. ``` HMACSHA256(     base64UrlEncode(header) + \".\" +     base64UrlEncode(payload),     API_SECRET) ```  ### Putting all together The output is three Base64-URL strings separated by dots. The following shows a JWT that has the previous header and payload encoded, and it is signed with a secret. ``` eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhZDU0YWFmZjg5ZmZkZmVmZjE3OGJiOGE4ZjM1OWIyOWZjYjIwZWRiNTYyNTBiOWY1ODRhYTJjYjAxNjJlZDRhIiwic3ViIjoiZGVtby5leGFtcGxlQGFjdHVhbHJlcG9ydHMuY29tIn0.SxO-H7UYYYsclS8RGWO1qf0z1cB1m73wF9FLl9RCc1Q  // Base64 encoded header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9 // Base64 encoded payload: eyJpc3MiOiJhZDU0YWFmZjg5ZmZkZmVmZjE3OGJiOGE4ZjM1OWIyOWZjYjIwZWRiNTYyNTBiOWY1ODRhYTJjYjAxNjJlZDRhIiwic3ViIjoiZGVtby5leGFtcGxlQGFjdHVhbHJlcG9ydHMuY29tIn0 // Signature: SxO-H7UYYYsclS8RGWO1qf0z1cB1m73wF9FLl9RCc1Q ```  ## Temporary JWTs You can create a temporary token in [Account Settings](https://pdfgeneratorapi.com/account/organization) page after you login to PDF Generator API. The generated token uses your email address as the subject (`sub`) value and is valid for __15 minutes__. You can also use [jwt.io](https://jwt.io/) to generate test tokens for your API calls. These test tokens should never be used in production applications. *  *  *  *  *  # Error codes  | Code   | Description                    | |--------|--------------------------------| | 401    | Unauthorized                   | | 402    | Payment Required               | | 403    | Forbidden                      | | 404    | Not Found                      | | 422    | Unprocessable Entity           | | 429    | Too Many Requests              | | 500    | Internal Server Error          |  ## 401 Unauthorized | Description                                                             | |-------------------------------------------------------------------------| | Authentication failed: request expired                                  | | Authentication failed: workspace missing                                | | Authentication failed: key missing                                      | | Authentication failed: property 'iss' (issuer) missing in JWT           | | Authentication failed: property 'sub' (subject) missing in JWT          | | Authentication failed: property 'exp' (expiration time) missing in JWT  | | Authentication failed: incorrect signature                              |  ## 402 Payment Required | Description                                                             | |-------------------------------------------------------------------------| | Your account is suspended, please upgrade your account                  |  ## 403 Forbidden | Description                                                             | |-------------------------------------------------------------------------| | Your account has exceeded the monthly document generation limit.        | | Access not granted: You cannot delete master workspace via API          | | Access not granted: Template is not accessible by this organization     | | Your session has expired, please close and reopen the editor.           |  ## 404 Entity not found | Description                                                             | |-------------------------------------------------------------------------| | Entity not found                                                        | | Resource not found                                                      | | None of the templates is available for the workspace.                   |  ## 422 Unprocessable Entity | Description                                                             | |-------------------------------------------------------------------------| | Unable to parse JSON, please check formatting                           | | Required parameter missing                                              | | Required parameter missing: template definition not defined             | | Required parameter missing: template not defined                        |  ## 429 Too Many Requests | Description                                                             | |-------------------------------------------------------------------------| | You can make up to 2 requests per second and 60 requests per minute.   |  *  *  *  *  * 

API version: 4.0.25
Contact: support@pdfgeneratorapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pdfgeneratorapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TemplateDefinitionNew type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateDefinitionNew{}

// TemplateDefinitionNew Template configuration. Use the Get schema endpoint to see the detailed payload structure.
type TemplateDefinitionNew struct {
	// Template name
	Name string `json:"name"`
	// A list of tags assigned to a template
	Tags []string `json:"tags,omitempty"`
	// Indicates if the template is a draft or published.
	IsDraft *bool `json:"isDraft,omitempty"`
	Layout *TemplateDefinitionNewLayout `json:"layout,omitempty"`
	// Defines page or label size, margins and components on page or label
	Pages []TemplateDefinitionNewPagesInner `json:"pages,omitempty"`
	DataSettings *TemplateDefinitionNewDataSettings `json:"dataSettings,omitempty"`
	Editor *TemplateDefinitionNewEditor `json:"editor,omitempty"`
	// If font-subsetting is applied to document when generated
	FontSubsetting *bool `json:"fontSubsetting,omitempty"`
	// Defines if barcodes are rendered as raster images instead of vector graphics.
	BarcodeAsImage *bool `json:"barcodeAsImage,omitempty"`
}

type _TemplateDefinitionNew TemplateDefinitionNew

// NewTemplateDefinitionNew instantiates a new TemplateDefinitionNew object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateDefinitionNew(name string) *TemplateDefinitionNew {
	this := TemplateDefinitionNew{}
	this.Name = name
	var fontSubsetting bool = false
	this.FontSubsetting = &fontSubsetting
	var barcodeAsImage bool = false
	this.BarcodeAsImage = &barcodeAsImage
	return &this
}

// NewTemplateDefinitionNewWithDefaults instantiates a new TemplateDefinitionNew object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateDefinitionNewWithDefaults() *TemplateDefinitionNew {
	this := TemplateDefinitionNew{}
	var fontSubsetting bool = false
	this.FontSubsetting = &fontSubsetting
	var barcodeAsImage bool = false
	this.BarcodeAsImage = &barcodeAsImage
	return &this
}

// GetName returns the Name field value
func (o *TemplateDefinitionNew) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNew) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateDefinitionNew) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TemplateDefinitionNew) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNew) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TemplateDefinitionNew) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TemplateDefinitionNew) SetTags(v []string) {
	o.Tags = v
}

// GetIsDraft returns the IsDraft field value if set, zero value otherwise.
func (o *TemplateDefinitionNew) GetIsDraft() bool {
	if o == nil || IsNil(o.IsDraft) {
		var ret bool
		return ret
	}
	return *o.IsDraft
}

// GetIsDraftOk returns a tuple with the IsDraft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNew) GetIsDraftOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDraft) {
		return nil, false
	}
	return o.IsDraft, true
}

// HasIsDraft returns a boolean if a field has been set.
func (o *TemplateDefinitionNew) HasIsDraft() bool {
	if o != nil && !IsNil(o.IsDraft) {
		return true
	}

	return false
}

// SetIsDraft gets a reference to the given bool and assigns it to the IsDraft field.
func (o *TemplateDefinitionNew) SetIsDraft(v bool) {
	o.IsDraft = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *TemplateDefinitionNew) GetLayout() TemplateDefinitionNewLayout {
	if o == nil || IsNil(o.Layout) {
		var ret TemplateDefinitionNewLayout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNew) GetLayoutOk() (*TemplateDefinitionNewLayout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *TemplateDefinitionNew) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given TemplateDefinitionNewLayout and assigns it to the Layout field.
func (o *TemplateDefinitionNew) SetLayout(v TemplateDefinitionNewLayout) {
	o.Layout = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *TemplateDefinitionNew) GetPages() []TemplateDefinitionNewPagesInner {
	if o == nil || IsNil(o.Pages) {
		var ret []TemplateDefinitionNewPagesInner
		return ret
	}
	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNew) GetPagesOk() ([]TemplateDefinitionNewPagesInner, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *TemplateDefinitionNew) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given []TemplateDefinitionNewPagesInner and assigns it to the Pages field.
func (o *TemplateDefinitionNew) SetPages(v []TemplateDefinitionNewPagesInner) {
	o.Pages = v
}

// GetDataSettings returns the DataSettings field value if set, zero value otherwise.
func (o *TemplateDefinitionNew) GetDataSettings() TemplateDefinitionNewDataSettings {
	if o == nil || IsNil(o.DataSettings) {
		var ret TemplateDefinitionNewDataSettings
		return ret
	}
	return *o.DataSettings
}

// GetDataSettingsOk returns a tuple with the DataSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNew) GetDataSettingsOk() (*TemplateDefinitionNewDataSettings, bool) {
	if o == nil || IsNil(o.DataSettings) {
		return nil, false
	}
	return o.DataSettings, true
}

// HasDataSettings returns a boolean if a field has been set.
func (o *TemplateDefinitionNew) HasDataSettings() bool {
	if o != nil && !IsNil(o.DataSettings) {
		return true
	}

	return false
}

// SetDataSettings gets a reference to the given TemplateDefinitionNewDataSettings and assigns it to the DataSettings field.
func (o *TemplateDefinitionNew) SetDataSettings(v TemplateDefinitionNewDataSettings) {
	o.DataSettings = &v
}

// GetEditor returns the Editor field value if set, zero value otherwise.
func (o *TemplateDefinitionNew) GetEditor() TemplateDefinitionNewEditor {
	if o == nil || IsNil(o.Editor) {
		var ret TemplateDefinitionNewEditor
		return ret
	}
	return *o.Editor
}

// GetEditorOk returns a tuple with the Editor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNew) GetEditorOk() (*TemplateDefinitionNewEditor, bool) {
	if o == nil || IsNil(o.Editor) {
		return nil, false
	}
	return o.Editor, true
}

// HasEditor returns a boolean if a field has been set.
func (o *TemplateDefinitionNew) HasEditor() bool {
	if o != nil && !IsNil(o.Editor) {
		return true
	}

	return false
}

// SetEditor gets a reference to the given TemplateDefinitionNewEditor and assigns it to the Editor field.
func (o *TemplateDefinitionNew) SetEditor(v TemplateDefinitionNewEditor) {
	o.Editor = &v
}

// GetFontSubsetting returns the FontSubsetting field value if set, zero value otherwise.
func (o *TemplateDefinitionNew) GetFontSubsetting() bool {
	if o == nil || IsNil(o.FontSubsetting) {
		var ret bool
		return ret
	}
	return *o.FontSubsetting
}

// GetFontSubsettingOk returns a tuple with the FontSubsetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNew) GetFontSubsettingOk() (*bool, bool) {
	if o == nil || IsNil(o.FontSubsetting) {
		return nil, false
	}
	return o.FontSubsetting, true
}

// HasFontSubsetting returns a boolean if a field has been set.
func (o *TemplateDefinitionNew) HasFontSubsetting() bool {
	if o != nil && !IsNil(o.FontSubsetting) {
		return true
	}

	return false
}

// SetFontSubsetting gets a reference to the given bool and assigns it to the FontSubsetting field.
func (o *TemplateDefinitionNew) SetFontSubsetting(v bool) {
	o.FontSubsetting = &v
}

// GetBarcodeAsImage returns the BarcodeAsImage field value if set, zero value otherwise.
func (o *TemplateDefinitionNew) GetBarcodeAsImage() bool {
	if o == nil || IsNil(o.BarcodeAsImage) {
		var ret bool
		return ret
	}
	return *o.BarcodeAsImage
}

// GetBarcodeAsImageOk returns a tuple with the BarcodeAsImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateDefinitionNew) GetBarcodeAsImageOk() (*bool, bool) {
	if o == nil || IsNil(o.BarcodeAsImage) {
		return nil, false
	}
	return o.BarcodeAsImage, true
}

// HasBarcodeAsImage returns a boolean if a field has been set.
func (o *TemplateDefinitionNew) HasBarcodeAsImage() bool {
	if o != nil && !IsNil(o.BarcodeAsImage) {
		return true
	}

	return false
}

// SetBarcodeAsImage gets a reference to the given bool and assigns it to the BarcodeAsImage field.
func (o *TemplateDefinitionNew) SetBarcodeAsImage(v bool) {
	o.BarcodeAsImage = &v
}

func (o TemplateDefinitionNew) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateDefinitionNew) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.IsDraft) {
		toSerialize["isDraft"] = o.IsDraft
	}
	if !IsNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}
	if !IsNil(o.Pages) {
		toSerialize["pages"] = o.Pages
	}
	if !IsNil(o.DataSettings) {
		toSerialize["dataSettings"] = o.DataSettings
	}
	if !IsNil(o.Editor) {
		toSerialize["editor"] = o.Editor
	}
	if !IsNil(o.FontSubsetting) {
		toSerialize["fontSubsetting"] = o.FontSubsetting
	}
	if !IsNil(o.BarcodeAsImage) {
		toSerialize["barcodeAsImage"] = o.BarcodeAsImage
	}
	return toSerialize, nil
}

func (o *TemplateDefinitionNew) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTemplateDefinitionNew := _TemplateDefinitionNew{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTemplateDefinitionNew)

	if err != nil {
		return err
	}

	*o = TemplateDefinitionNew(varTemplateDefinitionNew)

	return err
}

type NullableTemplateDefinitionNew struct {
	value *TemplateDefinitionNew
	isSet bool
}

func (v NullableTemplateDefinitionNew) Get() *TemplateDefinitionNew {
	return v.value
}

func (v *NullableTemplateDefinitionNew) Set(val *TemplateDefinitionNew) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateDefinitionNew) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateDefinitionNew) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateDefinitionNew(val *TemplateDefinitionNew) *NullableTemplateDefinitionNew {
	return &NullableTemplateDefinitionNew{value: val, isSet: true}
}

func (v NullableTemplateDefinitionNew) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateDefinitionNew) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


