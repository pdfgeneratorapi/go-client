# FormFieldsBase64

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileBase64** | **string** | PDF document in base64 encoded string format | 
**KeyField** | Pointer to [**KeyFieldParam**](KeyFieldParam.md) |  | [optional] [default to KEYFIELDPARAM_NAME]

## Methods

### NewFormFieldsBase64

`func NewFormFieldsBase64(fileBase64 string, ) *FormFieldsBase64`

NewFormFieldsBase64 instantiates a new FormFieldsBase64 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldsBase64WithDefaults

`func NewFormFieldsBase64WithDefaults() *FormFieldsBase64`

NewFormFieldsBase64WithDefaults instantiates a new FormFieldsBase64 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileBase64

`func (o *FormFieldsBase64) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *FormFieldsBase64) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *FormFieldsBase64) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.


### GetKeyField

`func (o *FormFieldsBase64) GetKeyField() KeyFieldParam`

GetKeyField returns the KeyField field if non-nil, zero value otherwise.

### GetKeyFieldOk

`func (o *FormFieldsBase64) GetKeyFieldOk() (*KeyFieldParam, bool)`

GetKeyFieldOk returns a tuple with the KeyField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyField

`func (o *FormFieldsBase64) SetKeyField(v KeyFieldParam)`

SetKeyField sets KeyField field to given value.

### HasKeyField

`func (o *FormFieldsBase64) HasKeyField() bool`

HasKeyField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


