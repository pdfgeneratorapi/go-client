/*
PDF Generator API

# Introduction [PDF Generator API](https://pdfgeneratorapi.com) allows you easily generate transactional PDF documents and reduce the development and support costs by enabling your users to create and manage their document templates using a browser-based drag-and-drop document editor.  The PDF Generator API features a web API architecture, allowing you to code in the language of your choice. This API supports the JSON media type, and uses UTF-8 character encoding.  ## Base URL The base URL for all the API endpoints is `https://us1.pdfgeneratorapi.com/api/v4`  For example * `https://us1.pdfgeneratorapi.com/api/v4/templates` * `https://us1.pdfgeneratorapi.com/api/v4/workspaces` * `https://us1.pdfgeneratorapi.com/api/v4/templates/123123`  ## Editor PDF Generator API comes with a powerful drag & drop editor that allows to create any kind of document templates, from barcode labels to invoices, quotes and reports. You can find tutorials and videos from our [Support Portal](https://support.pdfgeneratorapi.com). * [Component specification](https://support.pdfgeneratorapi.com/en/category/components-1ffseaj/) * [Expression Language documentation](https://support.pdfgeneratorapi.com/en/category/expression-language-q203pa/) * [Frequently asked questions and answers](https://support.pdfgeneratorapi.com/en/category/qanda-1ov519d/)  ## Definitions  ### Organization Organization is a group of workspaces owned by your account.  ### Workspace Workspace contains templates. Each workspace has access to their own templates and organization default templates.  ### Master Workspace Master Workspace is the main/default workspace of your Organization. The Master Workspace identifier is the email you signed up with.  ### Default Template Default template is a template that is available for all workspaces by default. You can set the template access type under Page Setup. If template has \"Organization\" access then your users can use them from the \"New\" menu in the Editor.  ### Data Field Data Field is a placeholder for the specific data in your JSON data set. In this example JSON you can access the buyer name using Data Field `{paymentDetails::buyerName}`. The separator between depth levels is :: (two colons). When designing the template you don’t have to know every Data Field, our editor automatically extracts all the available fields from your data set and provides an easy way to insert them into the template. ``` {     \"documentNumber\": 1,     \"paymentDetails\": {         \"method\": \"Credit Card\",         \"buyerName\": \"John Smith\"     },     \"items\": [         {             \"id\": 1,             \"name\": \"Item one\"         }     ] } ```  ## Rate limiting Our API endpoints use IP-based rate limiting and allow you to make up to 2 requests per second and 60 requests per minute. If you make more requests, you will receive a response with HTTP code 429.  Response headers contain additional values:  | Header   | Description                    | |--------|--------------------------------| | X-RateLimit-Limit    | Maximum requests per minute                   | | X-RateLimit-Remaining    | The requests remaining in the current minute               | | Retry-After     | How many seconds you need to wait until you are allowed to make requests |  *  *  *  *  *  # Libraries and SDKs ## Postman Collection We have created a [Postman Collection](https://www.postman.com/pdfgeneratorapi/workspace/pdf-generator-api-public-workspace/overview) so you can easily test all the API endpoints without developing and code.   ## Client Libraries All our Client Libraries are auto-generated using [OpenAPI Generator](https://openapi-generator.tech/) which uses the OpenAPI v3 specification to automatically generate a client library in specific programming language.  * [PHP Client](https://github.com/pdfgeneratorapi/php-client) * [Java Client](https://github.com/pdfgeneratorapi/java-client) * [Ruby Client](https://github.com/pdfgeneratorapi/ruby-client) * [Python Client](https://github.com/pdfgeneratorapi/python-client) * [Javascript Client](https://github.com/pdfgeneratorapi/javascript-client)  We have validated the generated libraries, but let us know if you find any anomalies in the client code.  ## Model Context Protocol (MCP) Server Integrate document generation directly into your AI agents and LLM applications using our official Model Context Protocol (MCP) Server.  The MCP server provides a standardized interface that allows AI assistants (like Claude Desktop and other MCP-compatible clients) to securely interact with the PDF Generator API. With it, your AI applications can automatically fetch workspaces, retrieve templates, merge data, and generate PDF documents on the fly.  [Get PDF Generator API MCP Server](https://github.com/pdfgeneratorapi/mcp-server) *  *  *  *  *   # Authentication The PDF Generator API uses __JSON Web Tokens (JWT)__ to authenticate all API requests. These tokens offer a method to establish secure server-to-server authentication by transferring a compact JSON object with a signed payload of your account’s API Key and Secret. When authenticating to the PDF Generator API, a JWT should be generated uniquely by a __server-side application__ and included as a __Bearer Token__ in the header of each request.   <SecurityDefinitions />  ## Accessing your API Key and Secret You can find your __API Key__ and __API Secret__ from the __Account Settings__ page after you login to PDF Generator API [here](https://pdfgeneratorapi.com/login).  ## Creating a JWT JSON Web Tokens are composed of three sections: a header, a payload (containing a claim set), and a signature. The header and payload are JSON objects, which are serialized to UTF-8 bytes, then encoded using base64url encoding.  The JWT's header, payload, and signature are concatenated with periods (.). As a result, a JWT typically takes the following form: ``` {Base64url encoded header}.{Base64url encoded payload}.{Base64url encoded signature} ```  We recommend and support libraries provided on [jwt.io](https://jwt.io/). While other libraries can create JWT, these recommended libraries are the most robust.  ### Header Property `alg` defines which signing algorithm is being used. PDF Generator API users HS256. Property `typ` defines the type of token and it is always JWT. ``` {   \"alg\": \"HS256\",   \"typ\": \"JWT\" } ```  ### Payload The second part of the token is the payload, which contains the claims  or the pieces of information being passed about the user and any metadata required. It is mandatory to specify the following claims: * issuer (`iss`): Your API key * subject (`sub`): Workspace identifier * expiration time (`exp`): Timestamp (unix epoch time) until the token is valid. It is highly recommended to set the exp timestamp for a short period, i.e. a matter of seconds. This way, if a token is intercepted or shared, the token will only be valid for a short period of time.  ``` {   \"iss\": \"ad54aaff89ffdfeff178bb8a8f359b29fcb20edb56250b9f584aa2cb0162ed4a\",   \"sub\": \"demo.example@actualreports.com\",   \"exp\": 1586112639 } ```  ### Payload for Partners Our partners can send their unique identifier (provided by us) in JWT's partner_id claim. If the `partner_id` value is specified in the JWT, the organization making the request is automatically connected to the partner account. * Partner ID (`partner_id`): Unique identifier provide by PDF Generator API team  ``` {   \"iss\": \"ad54aaff89ffdfeff178bb8a8f359b29fcb20edb56250b9f584aa2cb0162ed4a\",   \"sub\": \"demo.example@actualreports.com\",   \"partner_id\": \"my-partner-identifier\",   \"exp\": 1586112639 } ```  ### Signature To create the signature part you have to take the encoded header, the encoded payload, a secret, the algorithm specified in the header, and sign that. The signature is used to verify the message wasn't changed along the way, and, in the case of tokens signed with a private key, it can also verify that the sender of the JWT is who it says it is. ``` HMACSHA256(     base64UrlEncode(header) + \".\" +     base64UrlEncode(payload),     API_SECRET) ```  ### Putting all together The output is three Base64-URL strings separated by dots. The following shows a JWT that has the previous header and payload encoded, and it is signed with a secret. ``` eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhZDU0YWFmZjg5ZmZkZmVmZjE3OGJiOGE4ZjM1OWIyOWZjYjIwZWRiNTYyNTBiOWY1ODRhYTJjYjAxNjJlZDRhIiwic3ViIjoiZGVtby5leGFtcGxlQGFjdHVhbHJlcG9ydHMuY29tIn0.SxO-H7UYYYsclS8RGWO1qf0z1cB1m73wF9FLl9RCc1Q  // Base64 encoded header: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9 // Base64 encoded payload: eyJpc3MiOiJhZDU0YWFmZjg5ZmZkZmVmZjE3OGJiOGE4ZjM1OWIyOWZjYjIwZWRiNTYyNTBiOWY1ODRhYTJjYjAxNjJlZDRhIiwic3ViIjoiZGVtby5leGFtcGxlQGFjdHVhbHJlcG9ydHMuY29tIn0 // Signature: SxO-H7UYYYsclS8RGWO1qf0z1cB1m73wF9FLl9RCc1Q ```  ## Temporary JWTs You can create a temporary token in [Account Settings](https://pdfgeneratorapi.com/account/organization) page after you login to PDF Generator API. The generated token uses your email address as the subject (`sub`) value and is valid for __15 minutes__. You can also use [jwt.io](https://jwt.io/) to generate test tokens for your API calls. These test tokens should never be used in production applications. *  *  *  *  *  # Error codes  | Code   | Description                    | |--------|--------------------------------| | 401    | Unauthorized                   | | 402    | Payment Required               | | 403    | Forbidden                      | | 404    | Not Found                      | | 422    | Unprocessable Entity           | | 429    | Too Many Requests              | | 500    | Internal Server Error          |  ## 401 Unauthorized | Description                                                             | |-------------------------------------------------------------------------| | Authentication failed: request expired                                  | | Authentication failed: workspace missing                                | | Authentication failed: key missing                                      | | Authentication failed: property 'iss' (issuer) missing in JWT           | | Authentication failed: property 'sub' (subject) missing in JWT          | | Authentication failed: property 'exp' (expiration time) missing in JWT  | | Authentication failed: incorrect signature                              |  ## 402 Payment Required | Description                                                             | |-------------------------------------------------------------------------| | Your account is suspended, please upgrade your account                  |  ## 403 Forbidden | Description                                                             | |-------------------------------------------------------------------------| | Your account has exceeded the monthly document generation limit.        | | Access not granted: You cannot delete master workspace via API          | | Access not granted: Template is not accessible by this organization     | | Your session has expired, please close and reopen the editor.           |  ## 404 Entity not found | Description                                                             | |-------------------------------------------------------------------------| | Entity not found                                                        | | Resource not found                                                      | | None of the templates is available for the workspace.                   |  ## 422 Unprocessable Entity | Description                                                             | |-------------------------------------------------------------------------| | Unable to parse JSON, please check formatting                           | | Required parameter missing                                              | | Required parameter missing: template definition not defined             | | Required parameter missing: template not defined                        |  ## 429 Too Many Requests | Description                                                             | |-------------------------------------------------------------------------| | You can make up to 2 requests per second and 60 requests per minute.   |  *  *  *  *  * 

API version: 4.0.29
Contact: support@pdfgeneratorapi.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pdfgeneratorapi

import (
	"encoding/json"
)

// checks if the DocumentSignature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentSignature{}

// DocumentSignature struct for DocumentSignature
type DocumentSignature struct {
	// Name of the signature field in the document.
	FieldName *string `json:"field_name,omitempty"`
	// The name the signer signed under.
	SignerName NullableString `json:"signer_name,omitempty"`
	// Subject of the certificate that sealed it — an organization, not a person.
	Signer NullableString `json:"signer,omitempty"`
	// Subject of the certificate authority that issued it.
	Issuer NullableString `json:"issuer,omitempty"`
	// When a timestamp authority attested the signature (ISO-8601). This is the defensible time; the signer's own clock proves nothing. 
	SignedAt NullableString `json:"signed_at,omitempty"`
	// The time the signer's own software recorded (ISO-8601).
	ClaimedSignedAt NullableString `json:"claimed_signed_at,omitempty"`
	// The timestamp authority that attested the signature.
	TimestampAuthority NullableString `json:"timestamp_authority,omitempty"`
	// The bytes this signature covers are unchanged.
	Intact NullableBool `json:"intact,omitempty"`
	// The signature block itself adds up. On its own this does NOT mean the document is unchanged: a tampered file reports `valid` true with `intact` false, so a verdict needs intact AND valid AND trusted. 
	Valid NullableBool `json:"valid,omitempty"`
	// The certificate chains to a trusted root.
	Trusted NullableBool `json:"trusted,omitempty"`
	// How much of the file this signature protects.
	Coverage NullableString `json:"coverage,omitempty"`
	// The ETSI EN 319 102-1 indication, when the AdES engine could run.
	AdesIndication NullableString `json:"ades_indication,omitempty"`
	// The AdES sub-indication, naming why an indication is not PASSED.
	AdesSubIndication NullableString `json:"ades_sub_indication,omitempty"`
}

// NewDocumentSignature instantiates a new DocumentSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentSignature() *DocumentSignature {
	this := DocumentSignature{}
	return &this
}

// NewDocumentSignatureWithDefaults instantiates a new DocumentSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentSignatureWithDefaults() *DocumentSignature {
	this := DocumentSignature{}
	return &this
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *DocumentSignature) GetFieldName() string {
	if o == nil || IsNil(o.FieldName) {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentSignature) GetFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.FieldName) {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *DocumentSignature) HasFieldName() bool {
	if o != nil && !IsNil(o.FieldName) {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *DocumentSignature) SetFieldName(v string) {
	o.FieldName = &v
}

// GetSignerName returns the SignerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentSignature) GetSignerName() string {
	if o == nil || IsNil(o.SignerName.Get()) {
		var ret string
		return ret
	}
	return *o.SignerName.Get()
}

// GetSignerNameOk returns a tuple with the SignerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSignature) GetSignerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SignerName.Get(), o.SignerName.IsSet()
}

// HasSignerName returns a boolean if a field has been set.
func (o *DocumentSignature) HasSignerName() bool {
	if o != nil && o.SignerName.IsSet() {
		return true
	}

	return false
}

// SetSignerName gets a reference to the given NullableString and assigns it to the SignerName field.
func (o *DocumentSignature) SetSignerName(v string) {
	o.SignerName.Set(&v)
}
// SetSignerNameNil sets the value for SignerName to be an explicit nil
func (o *DocumentSignature) SetSignerNameNil() {
	o.SignerName.Set(nil)
}

// UnsetSignerName ensures that no value is present for SignerName, not even an explicit nil
func (o *DocumentSignature) UnsetSignerName() {
	o.SignerName.Unset()
}

// GetSigner returns the Signer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentSignature) GetSigner() string {
	if o == nil || IsNil(o.Signer.Get()) {
		var ret string
		return ret
	}
	return *o.Signer.Get()
}

// GetSignerOk returns a tuple with the Signer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSignature) GetSignerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signer.Get(), o.Signer.IsSet()
}

// HasSigner returns a boolean if a field has been set.
func (o *DocumentSignature) HasSigner() bool {
	if o != nil && o.Signer.IsSet() {
		return true
	}

	return false
}

// SetSigner gets a reference to the given NullableString and assigns it to the Signer field.
func (o *DocumentSignature) SetSigner(v string) {
	o.Signer.Set(&v)
}
// SetSignerNil sets the value for Signer to be an explicit nil
func (o *DocumentSignature) SetSignerNil() {
	o.Signer.Set(nil)
}

// UnsetSigner ensures that no value is present for Signer, not even an explicit nil
func (o *DocumentSignature) UnsetSigner() {
	o.Signer.Unset()
}

// GetIssuer returns the Issuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentSignature) GetIssuer() string {
	if o == nil || IsNil(o.Issuer.Get()) {
		var ret string
		return ret
	}
	return *o.Issuer.Get()
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSignature) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issuer.Get(), o.Issuer.IsSet()
}

// HasIssuer returns a boolean if a field has been set.
func (o *DocumentSignature) HasIssuer() bool {
	if o != nil && o.Issuer.IsSet() {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given NullableString and assigns it to the Issuer field.
func (o *DocumentSignature) SetIssuer(v string) {
	o.Issuer.Set(&v)
}
// SetIssuerNil sets the value for Issuer to be an explicit nil
func (o *DocumentSignature) SetIssuerNil() {
	o.Issuer.Set(nil)
}

// UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
func (o *DocumentSignature) UnsetIssuer() {
	o.Issuer.Unset()
}

// GetSignedAt returns the SignedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentSignature) GetSignedAt() string {
	if o == nil || IsNil(o.SignedAt.Get()) {
		var ret string
		return ret
	}
	return *o.SignedAt.Get()
}

// GetSignedAtOk returns a tuple with the SignedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSignature) GetSignedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SignedAt.Get(), o.SignedAt.IsSet()
}

// HasSignedAt returns a boolean if a field has been set.
func (o *DocumentSignature) HasSignedAt() bool {
	if o != nil && o.SignedAt.IsSet() {
		return true
	}

	return false
}

// SetSignedAt gets a reference to the given NullableString and assigns it to the SignedAt field.
func (o *DocumentSignature) SetSignedAt(v string) {
	o.SignedAt.Set(&v)
}
// SetSignedAtNil sets the value for SignedAt to be an explicit nil
func (o *DocumentSignature) SetSignedAtNil() {
	o.SignedAt.Set(nil)
}

// UnsetSignedAt ensures that no value is present for SignedAt, not even an explicit nil
func (o *DocumentSignature) UnsetSignedAt() {
	o.SignedAt.Unset()
}

// GetClaimedSignedAt returns the ClaimedSignedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentSignature) GetClaimedSignedAt() string {
	if o == nil || IsNil(o.ClaimedSignedAt.Get()) {
		var ret string
		return ret
	}
	return *o.ClaimedSignedAt.Get()
}

// GetClaimedSignedAtOk returns a tuple with the ClaimedSignedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSignature) GetClaimedSignedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClaimedSignedAt.Get(), o.ClaimedSignedAt.IsSet()
}

// HasClaimedSignedAt returns a boolean if a field has been set.
func (o *DocumentSignature) HasClaimedSignedAt() bool {
	if o != nil && o.ClaimedSignedAt.IsSet() {
		return true
	}

	return false
}

// SetClaimedSignedAt gets a reference to the given NullableString and assigns it to the ClaimedSignedAt field.
func (o *DocumentSignature) SetClaimedSignedAt(v string) {
	o.ClaimedSignedAt.Set(&v)
}
// SetClaimedSignedAtNil sets the value for ClaimedSignedAt to be an explicit nil
func (o *DocumentSignature) SetClaimedSignedAtNil() {
	o.ClaimedSignedAt.Set(nil)
}

// UnsetClaimedSignedAt ensures that no value is present for ClaimedSignedAt, not even an explicit nil
func (o *DocumentSignature) UnsetClaimedSignedAt() {
	o.ClaimedSignedAt.Unset()
}

// GetTimestampAuthority returns the TimestampAuthority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentSignature) GetTimestampAuthority() string {
	if o == nil || IsNil(o.TimestampAuthority.Get()) {
		var ret string
		return ret
	}
	return *o.TimestampAuthority.Get()
}

// GetTimestampAuthorityOk returns a tuple with the TimestampAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSignature) GetTimestampAuthorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimestampAuthority.Get(), o.TimestampAuthority.IsSet()
}

// HasTimestampAuthority returns a boolean if a field has been set.
func (o *DocumentSignature) HasTimestampAuthority() bool {
	if o != nil && o.TimestampAuthority.IsSet() {
		return true
	}

	return false
}

// SetTimestampAuthority gets a reference to the given NullableString and assigns it to the TimestampAuthority field.
func (o *DocumentSignature) SetTimestampAuthority(v string) {
	o.TimestampAuthority.Set(&v)
}
// SetTimestampAuthorityNil sets the value for TimestampAuthority to be an explicit nil
func (o *DocumentSignature) SetTimestampAuthorityNil() {
	o.TimestampAuthority.Set(nil)
}

// UnsetTimestampAuthority ensures that no value is present for TimestampAuthority, not even an explicit nil
func (o *DocumentSignature) UnsetTimestampAuthority() {
	o.TimestampAuthority.Unset()
}

// GetIntact returns the Intact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentSignature) GetIntact() bool {
	if o == nil || IsNil(o.Intact.Get()) {
		var ret bool
		return ret
	}
	return *o.Intact.Get()
}

// GetIntactOk returns a tuple with the Intact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSignature) GetIntactOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Intact.Get(), o.Intact.IsSet()
}

// HasIntact returns a boolean if a field has been set.
func (o *DocumentSignature) HasIntact() bool {
	if o != nil && o.Intact.IsSet() {
		return true
	}

	return false
}

// SetIntact gets a reference to the given NullableBool and assigns it to the Intact field.
func (o *DocumentSignature) SetIntact(v bool) {
	o.Intact.Set(&v)
}
// SetIntactNil sets the value for Intact to be an explicit nil
func (o *DocumentSignature) SetIntactNil() {
	o.Intact.Set(nil)
}

// UnsetIntact ensures that no value is present for Intact, not even an explicit nil
func (o *DocumentSignature) UnsetIntact() {
	o.Intact.Unset()
}

// GetValid returns the Valid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentSignature) GetValid() bool {
	if o == nil || IsNil(o.Valid.Get()) {
		var ret bool
		return ret
	}
	return *o.Valid.Get()
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSignature) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Valid.Get(), o.Valid.IsSet()
}

// HasValid returns a boolean if a field has been set.
func (o *DocumentSignature) HasValid() bool {
	if o != nil && o.Valid.IsSet() {
		return true
	}

	return false
}

// SetValid gets a reference to the given NullableBool and assigns it to the Valid field.
func (o *DocumentSignature) SetValid(v bool) {
	o.Valid.Set(&v)
}
// SetValidNil sets the value for Valid to be an explicit nil
func (o *DocumentSignature) SetValidNil() {
	o.Valid.Set(nil)
}

// UnsetValid ensures that no value is present for Valid, not even an explicit nil
func (o *DocumentSignature) UnsetValid() {
	o.Valid.Unset()
}

// GetTrusted returns the Trusted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentSignature) GetTrusted() bool {
	if o == nil || IsNil(o.Trusted.Get()) {
		var ret bool
		return ret
	}
	return *o.Trusted.Get()
}

// GetTrustedOk returns a tuple with the Trusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSignature) GetTrustedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Trusted.Get(), o.Trusted.IsSet()
}

// HasTrusted returns a boolean if a field has been set.
func (o *DocumentSignature) HasTrusted() bool {
	if o != nil && o.Trusted.IsSet() {
		return true
	}

	return false
}

// SetTrusted gets a reference to the given NullableBool and assigns it to the Trusted field.
func (o *DocumentSignature) SetTrusted(v bool) {
	o.Trusted.Set(&v)
}
// SetTrustedNil sets the value for Trusted to be an explicit nil
func (o *DocumentSignature) SetTrustedNil() {
	o.Trusted.Set(nil)
}

// UnsetTrusted ensures that no value is present for Trusted, not even an explicit nil
func (o *DocumentSignature) UnsetTrusted() {
	o.Trusted.Unset()
}

// GetCoverage returns the Coverage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentSignature) GetCoverage() string {
	if o == nil || IsNil(o.Coverage.Get()) {
		var ret string
		return ret
	}
	return *o.Coverage.Get()
}

// GetCoverageOk returns a tuple with the Coverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSignature) GetCoverageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Coverage.Get(), o.Coverage.IsSet()
}

// HasCoverage returns a boolean if a field has been set.
func (o *DocumentSignature) HasCoverage() bool {
	if o != nil && o.Coverage.IsSet() {
		return true
	}

	return false
}

// SetCoverage gets a reference to the given NullableString and assigns it to the Coverage field.
func (o *DocumentSignature) SetCoverage(v string) {
	o.Coverage.Set(&v)
}
// SetCoverageNil sets the value for Coverage to be an explicit nil
func (o *DocumentSignature) SetCoverageNil() {
	o.Coverage.Set(nil)
}

// UnsetCoverage ensures that no value is present for Coverage, not even an explicit nil
func (o *DocumentSignature) UnsetCoverage() {
	o.Coverage.Unset()
}

// GetAdesIndication returns the AdesIndication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentSignature) GetAdesIndication() string {
	if o == nil || IsNil(o.AdesIndication.Get()) {
		var ret string
		return ret
	}
	return *o.AdesIndication.Get()
}

// GetAdesIndicationOk returns a tuple with the AdesIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSignature) GetAdesIndicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdesIndication.Get(), o.AdesIndication.IsSet()
}

// HasAdesIndication returns a boolean if a field has been set.
func (o *DocumentSignature) HasAdesIndication() bool {
	if o != nil && o.AdesIndication.IsSet() {
		return true
	}

	return false
}

// SetAdesIndication gets a reference to the given NullableString and assigns it to the AdesIndication field.
func (o *DocumentSignature) SetAdesIndication(v string) {
	o.AdesIndication.Set(&v)
}
// SetAdesIndicationNil sets the value for AdesIndication to be an explicit nil
func (o *DocumentSignature) SetAdesIndicationNil() {
	o.AdesIndication.Set(nil)
}

// UnsetAdesIndication ensures that no value is present for AdesIndication, not even an explicit nil
func (o *DocumentSignature) UnsetAdesIndication() {
	o.AdesIndication.Unset()
}

// GetAdesSubIndication returns the AdesSubIndication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentSignature) GetAdesSubIndication() string {
	if o == nil || IsNil(o.AdesSubIndication.Get()) {
		var ret string
		return ret
	}
	return *o.AdesSubIndication.Get()
}

// GetAdesSubIndicationOk returns a tuple with the AdesSubIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentSignature) GetAdesSubIndicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdesSubIndication.Get(), o.AdesSubIndication.IsSet()
}

// HasAdesSubIndication returns a boolean if a field has been set.
func (o *DocumentSignature) HasAdesSubIndication() bool {
	if o != nil && o.AdesSubIndication.IsSet() {
		return true
	}

	return false
}

// SetAdesSubIndication gets a reference to the given NullableString and assigns it to the AdesSubIndication field.
func (o *DocumentSignature) SetAdesSubIndication(v string) {
	o.AdesSubIndication.Set(&v)
}
// SetAdesSubIndicationNil sets the value for AdesSubIndication to be an explicit nil
func (o *DocumentSignature) SetAdesSubIndicationNil() {
	o.AdesSubIndication.Set(nil)
}

// UnsetAdesSubIndication ensures that no value is present for AdesSubIndication, not even an explicit nil
func (o *DocumentSignature) UnsetAdesSubIndication() {
	o.AdesSubIndication.Unset()
}

func (o DocumentSignature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentSignature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldName) {
		toSerialize["field_name"] = o.FieldName
	}
	if o.SignerName.IsSet() {
		toSerialize["signer_name"] = o.SignerName.Get()
	}
	if o.Signer.IsSet() {
		toSerialize["signer"] = o.Signer.Get()
	}
	if o.Issuer.IsSet() {
		toSerialize["issuer"] = o.Issuer.Get()
	}
	if o.SignedAt.IsSet() {
		toSerialize["signed_at"] = o.SignedAt.Get()
	}
	if o.ClaimedSignedAt.IsSet() {
		toSerialize["claimed_signed_at"] = o.ClaimedSignedAt.Get()
	}
	if o.TimestampAuthority.IsSet() {
		toSerialize["timestamp_authority"] = o.TimestampAuthority.Get()
	}
	if o.Intact.IsSet() {
		toSerialize["intact"] = o.Intact.Get()
	}
	if o.Valid.IsSet() {
		toSerialize["valid"] = o.Valid.Get()
	}
	if o.Trusted.IsSet() {
		toSerialize["trusted"] = o.Trusted.Get()
	}
	if o.Coverage.IsSet() {
		toSerialize["coverage"] = o.Coverage.Get()
	}
	if o.AdesIndication.IsSet() {
		toSerialize["ades_indication"] = o.AdesIndication.Get()
	}
	if o.AdesSubIndication.IsSet() {
		toSerialize["ades_sub_indication"] = o.AdesSubIndication.Get()
	}
	return toSerialize, nil
}

type NullableDocumentSignature struct {
	value *DocumentSignature
	isSet bool
}

func (v NullableDocumentSignature) Get() *DocumentSignature {
	return v.value
}

func (v *NullableDocumentSignature) Set(val *DocumentSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentSignature(val *DocumentSignature) *NullableDocumentSignature {
	return &NullableDocumentSignature{value: val, isSet: true}
}

func (v NullableDocumentSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


