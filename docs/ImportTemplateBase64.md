# ImportTemplateBase64

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | [**ImportTemplateUrlTemplate**](ImportTemplateUrlTemplate.md) |  | 
**FileBase64** | **string** | PDF file from base64 string to import | 

## Methods

### NewImportTemplateBase64

`func NewImportTemplateBase64(template ImportTemplateUrlTemplate, fileBase64 string, ) *ImportTemplateBase64`

NewImportTemplateBase64 instantiates a new ImportTemplateBase64 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportTemplateBase64WithDefaults

`func NewImportTemplateBase64WithDefaults() *ImportTemplateBase64`

NewImportTemplateBase64WithDefaults instantiates a new ImportTemplateBase64 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *ImportTemplateBase64) GetTemplate() ImportTemplateUrlTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ImportTemplateBase64) GetTemplateOk() (*ImportTemplateUrlTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ImportTemplateBase64) SetTemplate(v ImportTemplateUrlTemplate)`

SetTemplate sets Template field to given value.


### GetFileBase64

`func (o *ImportTemplateBase64) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *ImportTemplateBase64) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *ImportTemplateBase64) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


