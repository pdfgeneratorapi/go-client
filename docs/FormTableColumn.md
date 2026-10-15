# FormTableColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Column label displayed in the form | 
**Name** | **string** | Data field name of the column | 
**Type** | **string** | Column type. &#x60;select&#x60; columns require &#x60;subRows&#x60; with options and a &#x60;default&#x60; value. | 
**Required** | **bool** | Specifies if the column is required or not | 
**Default** | Pointer to **NullableString** | Value pre-filled in the column. Required for &#x60;select&#x60; columns and must be one of the options. | [optional] 
**SubRows** | Pointer to [**[]FormFieldOptions**](FormFieldOptions.md) | Options for &#x60;select&#x60; columns (a single object). Not used by other column types. | [optional] 

## Methods

### NewFormTableColumn

`func NewFormTableColumn(label string, name string, type_ string, required bool, ) *FormTableColumn`

NewFormTableColumn instantiates a new FormTableColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormTableColumnWithDefaults

`func NewFormTableColumnWithDefaults() *FormTableColumn`

NewFormTableColumnWithDefaults instantiates a new FormTableColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *FormTableColumn) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormTableColumn) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormTableColumn) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *FormTableColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormTableColumn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormTableColumn) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FormTableColumn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormTableColumn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormTableColumn) SetType(v string)`

SetType sets Type field to given value.


### GetRequired

`func (o *FormTableColumn) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormTableColumn) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormTableColumn) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDefault

`func (o *FormTableColumn) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FormTableColumn) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FormTableColumn) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FormTableColumn) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *FormTableColumn) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *FormTableColumn) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetSubRows

`func (o *FormTableColumn) GetSubRows() []FormFieldOptions`

GetSubRows returns the SubRows field if non-nil, zero value otherwise.

### GetSubRowsOk

`func (o *FormTableColumn) GetSubRowsOk() (*[]FormFieldOptions, bool)`

GetSubRowsOk returns a tuple with the SubRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRows

`func (o *FormTableColumn) SetSubRows(v []FormFieldOptions)`

SetSubRows sets SubRows field to given value.

### HasSubRows

`func (o *FormTableColumn) HasSubRows() bool`

HasSubRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


