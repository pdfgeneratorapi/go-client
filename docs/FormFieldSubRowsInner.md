# FormFieldSubRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | **string** | Comma-separated list of choices. Whitespace around each choice is trimmed and choices must be unique. | 
**Label** | **string** | Column label displayed in the form | 
**Name** | **string** | Data field name of the column | 
**Type** | **string** | Column type. &#x60;select&#x60; columns require &#x60;subRows&#x60; with options and a &#x60;default&#x60; value. | 
**Required** | **bool** | Specifies if the column is required or not | 
**Default** | Pointer to **NullableString** | Value pre-filled in the column. Required for &#x60;select&#x60; columns and must be one of the options. | [optional] 
**SubRows** | Pointer to [**[]FormFieldOptions**](FormFieldOptions.md) | Options for &#x60;select&#x60; columns (a single object). Not used by other column types. | [optional] 

## Methods

### NewFormFieldSubRowsInner

`func NewFormFieldSubRowsInner(options string, label string, name string, type_ string, required bool, ) *FormFieldSubRowsInner`

NewFormFieldSubRowsInner instantiates a new FormFieldSubRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldSubRowsInnerWithDefaults

`func NewFormFieldSubRowsInnerWithDefaults() *FormFieldSubRowsInner`

NewFormFieldSubRowsInnerWithDefaults instantiates a new FormFieldSubRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *FormFieldSubRowsInner) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormFieldSubRowsInner) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormFieldSubRowsInner) SetOptions(v string)`

SetOptions sets Options field to given value.


### GetLabel

`func (o *FormFieldSubRowsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldSubRowsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldSubRowsInner) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *FormFieldSubRowsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormFieldSubRowsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormFieldSubRowsInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FormFieldSubRowsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldSubRowsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldSubRowsInner) SetType(v string)`

SetType sets Type field to given value.


### GetRequired

`func (o *FormFieldSubRowsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormFieldSubRowsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormFieldSubRowsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDefault

`func (o *FormFieldSubRowsInner) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FormFieldSubRowsInner) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FormFieldSubRowsInner) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FormFieldSubRowsInner) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *FormFieldSubRowsInner) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *FormFieldSubRowsInner) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetSubRows

`func (o *FormFieldSubRowsInner) GetSubRows() []FormFieldOptions`

GetSubRows returns the SubRows field if non-nil, zero value otherwise.

### GetSubRowsOk

`func (o *FormFieldSubRowsInner) GetSubRowsOk() (*[]FormFieldOptions, bool)`

GetSubRowsOk returns a tuple with the SubRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRows

`func (o *FormFieldSubRowsInner) SetSubRows(v []FormFieldOptions)`

SetSubRows sets SubRows field to given value.

### HasSubRows

`func (o *FormFieldSubRowsInner) HasSubRows() bool`

HasSubRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


