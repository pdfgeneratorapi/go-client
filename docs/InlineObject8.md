# InlineObject8

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to **string** | Base64 encoded document when output&#x3D;base64, URL to the document when output&#x3D;url, or URL to the PDF viewer when output&#x3D;viewer. | [optional] 
**Meta** | Pointer to [**InlineObject8Meta**](InlineObject8Meta.md) |  | [optional] 

## Methods

### NewInlineObject8

`func NewInlineObject8() *InlineObject8`

NewInlineObject8 instantiates a new InlineObject8 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject8WithDefaults

`func NewInlineObject8WithDefaults() *InlineObject8`

NewInlineObject8WithDefaults instantiates a new InlineObject8 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *InlineObject8) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *InlineObject8) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *InlineObject8) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *InlineObject8) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetMeta

`func (o *InlineObject8) GetMeta() InlineObject8Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineObject8) GetMetaOk() (*InlineObject8Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineObject8) SetMeta(v InlineObject8Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineObject8) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


