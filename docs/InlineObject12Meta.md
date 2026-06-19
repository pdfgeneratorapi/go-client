# InlineObject12Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Document name. This value is automatically generated if name attribute is not defined in request. | [optional] 
**DisplayName** | Pointer to **string** | Document name without the file extension. | [optional] 
**Encoding** | Pointer to **string** | Document encoding | [optional] 
**ContentType** | Pointer to **string** | Document content type. | [optional] 
**Stats** | Pointer to [**InlineObject12MetaStats**](InlineObject12MetaStats.md) |  | [optional] 
**PublicId** | Pointer to **string** | Document public id, if output&#x3D;url was used for document storage | [optional] 

## Methods

### NewInlineObject12Meta

`func NewInlineObject12Meta() *InlineObject12Meta`

NewInlineObject12Meta instantiates a new InlineObject12Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject12MetaWithDefaults

`func NewInlineObject12MetaWithDefaults() *InlineObject12Meta`

NewInlineObject12MetaWithDefaults instantiates a new InlineObject12Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject12Meta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject12Meta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject12Meta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject12Meta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *InlineObject12Meta) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InlineObject12Meta) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InlineObject12Meta) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InlineObject12Meta) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEncoding

`func (o *InlineObject12Meta) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *InlineObject12Meta) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *InlineObject12Meta) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *InlineObject12Meta) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetContentType

`func (o *InlineObject12Meta) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *InlineObject12Meta) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *InlineObject12Meta) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *InlineObject12Meta) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetStats

`func (o *InlineObject12Meta) GetStats() InlineObject12MetaStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *InlineObject12Meta) GetStatsOk() (*InlineObject12MetaStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *InlineObject12Meta) SetStats(v InlineObject12MetaStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *InlineObject12Meta) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetPublicId

`func (o *InlineObject12Meta) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *InlineObject12Meta) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *InlineObject12Meta) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *InlineObject12Meta) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


