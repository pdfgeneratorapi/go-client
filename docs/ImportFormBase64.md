# ImportFormBase64

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileBase64** | **string** | PDF file from base64 string to import | 
**Form** | Pointer to [**FormConfigurationNew**](FormConfigurationNew.md) |  | [optional] 

## Methods

### NewImportFormBase64

`func NewImportFormBase64(fileBase64 string, ) *ImportFormBase64`

NewImportFormBase64 instantiates a new ImportFormBase64 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportFormBase64WithDefaults

`func NewImportFormBase64WithDefaults() *ImportFormBase64`

NewImportFormBase64WithDefaults instantiates a new ImportFormBase64 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileBase64

`func (o *ImportFormBase64) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *ImportFormBase64) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *ImportFormBase64) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.


### GetForm

`func (o *ImportFormBase64) GetForm() FormConfigurationNew`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *ImportFormBase64) GetFormOk() (*FormConfigurationNew, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *ImportFormBase64) SetForm(v FormConfigurationNew)`

SetForm sets Form field to given value.

### HasForm

`func (o *ImportFormBase64) HasForm() bool`

HasForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


