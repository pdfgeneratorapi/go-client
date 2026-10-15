# FormField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Field label displayed in the form | 
**Name** | **string** | Data field name. For example \&quot;name\&quot; can be used as \&quot;{name}\&quot; in the document as placeholder. | 
**Type** | **string** | Field type * &#x60;text&#x60; - Free text input * &#x60;integer&#x60; - Whole number input * &#x60;number&#x60; - Decimal number input * &#x60;date&#x60; - Date picker * &#x60;signature&#x60; - Hand-drawn signature, stored as an image * &#x60;select&#x60; - Dropdown with a single choice. Requires &#x60;subRows&#x60; with options and a &#x60;default&#x60; value * &#x60;multi-select&#x60; - Dropdown with multiple choices. Requires &#x60;subRows&#x60; with options * &#x60;table&#x60; - Repeating rows of values. Its columns are defined in &#x60;subRows&#x60; * &#x60;image&#x60; - Image upload * &#x60;pdf&#x60; - PDF file upload  | 
**Required** | **bool** | Specifies if the field is required or not | 
**Default** | Pointer to **NullableString** | Value pre-filled in the field. Required for &#x60;select&#x60; fields and must be one of the options. Not used by &#x60;table&#x60; fields. | [optional] 
**Description** | Pointer to **NullableString** | Help text shown with the field | [optional] 
**SubRows** | Pointer to [**[]FormFieldSubRowsInner**](FormFieldSubRowsInner.md) | Options for &#x60;select&#x60; and &#x60;multi-select&#x60; fields (a single object), columns for &#x60;table&#x60; fields. Not used by other field types. | [optional] 

## Methods

### NewFormField

`func NewFormField(label string, name string, type_ string, required bool, ) *FormField`

NewFormField instantiates a new FormField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldWithDefaults

`func NewFormFieldWithDefaults() *FormField`

NewFormFieldWithDefaults instantiates a new FormField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *FormField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *FormField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormField) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FormField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormField) SetType(v string)`

SetType sets Type field to given value.


### GetRequired

`func (o *FormField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormField) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDefault

`func (o *FormField) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FormField) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FormField) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FormField) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *FormField) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *FormField) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetDescription

`func (o *FormField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FormField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FormField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FormField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FormField) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FormField) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSubRows

`func (o *FormField) GetSubRows() []FormFieldSubRowsInner`

GetSubRows returns the SubRows field if non-nil, zero value otherwise.

### GetSubRowsOk

`func (o *FormField) GetSubRowsOk() (*[]FormFieldSubRowsInner, bool)`

GetSubRowsOk returns a tuple with the SubRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRows

`func (o *FormField) SetSubRows(v []FormFieldSubRowsInner)`

SetSubRows sets SubRows field to given value.

### HasSubRows

`func (o *FormField) HasSubRows() bool`

HasSubRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


