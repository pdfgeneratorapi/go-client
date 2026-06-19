# DataBatchInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to **int32** | Template id | [optional] 
**Data** | Pointer to **map[string]interface{}** | JSON data used to replace data fields in the template | [optional] 

## Methods

### NewDataBatchInner

`func NewDataBatchInner() *DataBatchInner`

NewDataBatchInner instantiates a new DataBatchInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataBatchInnerWithDefaults

`func NewDataBatchInnerWithDefaults() *DataBatchInner`

NewDataBatchInnerWithDefaults instantiates a new DataBatchInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *DataBatchInner) GetTemplate() int32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *DataBatchInner) GetTemplateOk() (*int32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *DataBatchInner) SetTemplate(v int32)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *DataBatchInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetData

`func (o *DataBatchInner) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DataBatchInner) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DataBatchInner) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *DataBatchInner) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


