# InlineObject9

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to **string** | Base64 encoded document if the output&#x3D;base64 is used or URL to the document when the output&#x3D;url is used. | [optional] 
**Meta** | Pointer to [**InlineObject9Meta**](InlineObject9Meta.md) |  | [optional] 

## Methods

### NewInlineObject9

`func NewInlineObject9() *InlineObject9`

NewInlineObject9 instantiates a new InlineObject9 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject9WithDefaults

`func NewInlineObject9WithDefaults() *InlineObject9`

NewInlineObject9WithDefaults instantiates a new InlineObject9 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *InlineObject9) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *InlineObject9) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *InlineObject9) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *InlineObject9) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetMeta

`func (o *InlineObject9) GetMeta() InlineObject9Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineObject9) GetMetaOk() (*InlineObject9Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineObject9) SetMeta(v InlineObject9Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineObject9) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


