# ImportFormUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | PDF file from remote URL to import | 
**Form** | Pointer to [**FormConfigurationNew**](FormConfigurationNew.md) |  | [optional] 

## Methods

### NewImportFormUrl

`func NewImportFormUrl(fileUrl string, ) *ImportFormUrl`

NewImportFormUrl instantiates a new ImportFormUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportFormUrlWithDefaults

`func NewImportFormUrlWithDefaults() *ImportFormUrl`

NewImportFormUrlWithDefaults instantiates a new ImportFormUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *ImportFormUrl) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ImportFormUrl) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ImportFormUrl) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetForm

`func (o *ImportFormUrl) GetForm() FormConfigurationNew`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *ImportFormUrl) GetFormOk() (*FormConfigurationNew, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *ImportFormUrl) SetForm(v FormConfigurationNew)`

SetForm sets Form field to given value.

### HasForm

`func (o *ImportFormUrl) HasForm() bool`

HasForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


