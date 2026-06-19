# FormActionSendSendDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | HTTPS URL where the generated document will be delivered. | 
**Headers** | Pointer to [**[]FormActionSendSendDocumentHeadersInner**](FormActionSendSendDocumentHeadersInner.md) | Optional HTTP headers to include with the delivery request. | [optional] 

## Methods

### NewFormActionSendSendDocument

`func NewFormActionSendSendDocument(url string, ) *FormActionSendSendDocument`

NewFormActionSendSendDocument instantiates a new FormActionSendSendDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormActionSendSendDocumentWithDefaults

`func NewFormActionSendSendDocumentWithDefaults() *FormActionSendSendDocument`

NewFormActionSendSendDocumentWithDefaults instantiates a new FormActionSendSendDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *FormActionSendSendDocument) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FormActionSendSendDocument) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FormActionSendSendDocument) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *FormActionSendSendDocument) GetHeaders() []FormActionSendSendDocumentHeadersInner`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *FormActionSendSendDocument) GetHeadersOk() (*[]FormActionSendSendDocumentHeadersInner, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *FormActionSendSendDocument) SetHeaders(v []FormActionSendSendDocumentHeadersInner)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *FormActionSendSendDocument) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


