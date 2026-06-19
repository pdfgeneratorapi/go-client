# FormFieldsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Field label displayed in the form | [optional] 
**Name** | Pointer to **string** | Data field name. For example \&quot;name\&quot; can be used as \&quot;{name}\&quot; in the document as placeholder. | [optional] 
**Type** | Pointer to **string** | Field type | [optional] 
**Required** | Pointer to **bool** | Specifies if the field is required or not | [optional] 

## Methods

### NewFormFieldsInner

`func NewFormFieldsInner() *FormFieldsInner`

NewFormFieldsInner instantiates a new FormFieldsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldsInnerWithDefaults

`func NewFormFieldsInnerWithDefaults() *FormFieldsInner`

NewFormFieldsInnerWithDefaults instantiates a new FormFieldsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *FormFieldsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldsInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FormFieldsInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *FormFieldsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormFieldsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormFieldsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormFieldsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *FormFieldsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormFieldsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequired

`func (o *FormFieldsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormFieldsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormFieldsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormFieldsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


