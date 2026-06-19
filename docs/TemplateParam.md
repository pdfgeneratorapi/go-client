# TemplateParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**TemplateParamId**](TemplateParamId.md) |  | [optional] 
**VersionId** | Pointer to **int32** | Template version ID | [optional] 
**Data** | Pointer to [**TemplateParamData**](TemplateParamData.md) |  | [optional] 

## Methods

### NewTemplateParam

`func NewTemplateParam() *TemplateParam`

NewTemplateParam instantiates a new TemplateParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateParamWithDefaults

`func NewTemplateParamWithDefaults() *TemplateParam`

NewTemplateParamWithDefaults instantiates a new TemplateParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateParam) GetId() TemplateParamId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateParam) GetIdOk() (*TemplateParamId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateParam) SetId(v TemplateParamId)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateParam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionId

`func (o *TemplateParam) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *TemplateParam) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *TemplateParam) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *TemplateParam) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetData

`func (o *TemplateParam) GetData() TemplateParamData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TemplateParam) GetDataOk() (*TemplateParamData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TemplateParam) SetData(v TemplateParamData)`

SetData sets Data field to given value.

### HasData

`func (o *TemplateParam) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


