# TemplateVersionCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**[]TemplateVersion**](TemplateVersion.md) |  | [optional] 
**Meta** | Pointer to [**TemplateVersionCollectionMeta**](TemplateVersionCollectionMeta.md) |  | [optional] 

## Methods

### NewTemplateVersionCollection

`func NewTemplateVersionCollection() *TemplateVersionCollection`

NewTemplateVersionCollection instantiates a new TemplateVersionCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateVersionCollectionWithDefaults

`func NewTemplateVersionCollectionWithDefaults() *TemplateVersionCollection`

NewTemplateVersionCollectionWithDefaults instantiates a new TemplateVersionCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *TemplateVersionCollection) GetResponse() []TemplateVersion`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *TemplateVersionCollection) GetResponseOk() (*[]TemplateVersion, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *TemplateVersionCollection) SetResponse(v []TemplateVersion)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *TemplateVersionCollection) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetMeta

`func (o *TemplateVersionCollection) GetMeta() TemplateVersionCollectionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TemplateVersionCollection) GetMetaOk() (*TemplateVersionCollectionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TemplateVersionCollection) SetMeta(v TemplateVersionCollectionMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TemplateVersionCollection) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


