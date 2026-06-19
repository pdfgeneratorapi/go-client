# ImportFormRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | PDF file from remote URL to import | 
**Form** | Pointer to [**FormConfigurationNew**](FormConfigurationNew.md) |  | [optional] 
**FileBase64** | **string** | PDF file from base64 string to import | 

## Methods

### NewImportFormRequest

`func NewImportFormRequest(fileUrl string, fileBase64 string, ) *ImportFormRequest`

NewImportFormRequest instantiates a new ImportFormRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportFormRequestWithDefaults

`func NewImportFormRequestWithDefaults() *ImportFormRequest`

NewImportFormRequestWithDefaults instantiates a new ImportFormRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *ImportFormRequest) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ImportFormRequest) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ImportFormRequest) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetForm

`func (o *ImportFormRequest) GetForm() FormConfigurationNew`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *ImportFormRequest) GetFormOk() (*FormConfigurationNew, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *ImportFormRequest) SetForm(v FormConfigurationNew)`

SetForm sets Form field to given value.

### HasForm

`func (o *ImportFormRequest) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetFileBase64

`func (o *ImportFormRequest) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *ImportFormRequest) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *ImportFormRequest) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


