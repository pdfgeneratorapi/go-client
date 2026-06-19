# FormFieldsUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | Public URL to a PDF document | 
**KeyField** | Pointer to [**KeyFieldParam**](KeyFieldParam.md) |  | [optional] [default to KEYFIELDPARAM_NAME]

## Methods

### NewFormFieldsUrl

`func NewFormFieldsUrl(fileUrl string, ) *FormFieldsUrl`

NewFormFieldsUrl instantiates a new FormFieldsUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldsUrlWithDefaults

`func NewFormFieldsUrlWithDefaults() *FormFieldsUrl`

NewFormFieldsUrlWithDefaults instantiates a new FormFieldsUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *FormFieldsUrl) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *FormFieldsUrl) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *FormFieldsUrl) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetKeyField

`func (o *FormFieldsUrl) GetKeyField() KeyFieldParam`

GetKeyField returns the KeyField field if non-nil, zero value otherwise.

### GetKeyFieldOk

`func (o *FormFieldsUrl) GetKeyFieldOk() (*KeyFieldParam, bool)`

GetKeyFieldOk returns a tuple with the KeyField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyField

`func (o *FormFieldsUrl) SetKeyField(v KeyFieldParam)`

SetKeyField sets KeyField field to given value.

### HasKeyField

`func (o *FormFieldsUrl) HasKeyField() bool`

HasKeyField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


