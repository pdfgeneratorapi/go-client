# InlineObject12

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to **string** | Base64 encoded document if the output&#x3D;base64 is used or URL to the document when the output&#x3D;url is used. | [optional] 
**Meta** | Pointer to [**InlineObject12Meta**](InlineObject12Meta.md) |  | [optional] 

## Methods

### NewInlineObject12

`func NewInlineObject12() *InlineObject12`

NewInlineObject12 instantiates a new InlineObject12 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject12WithDefaults

`func NewInlineObject12WithDefaults() *InlineObject12`

NewInlineObject12WithDefaults instantiates a new InlineObject12 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *InlineObject12) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *InlineObject12) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *InlineObject12) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *InlineObject12) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetMeta

`func (o *InlineObject12) GetMeta() InlineObject12Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineObject12) GetMetaOk() (*InlineObject12Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineObject12) SetMeta(v InlineObject12Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InlineObject12) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


