# ImportTemplateUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | [**ImportTemplateUrlTemplate**](ImportTemplateUrlTemplate.md) |  | 
**FileUrl** | **string** | PDF file from remote URL to import | 

## Methods

### NewImportTemplateUrl

`func NewImportTemplateUrl(template ImportTemplateUrlTemplate, fileUrl string, ) *ImportTemplateUrl`

NewImportTemplateUrl instantiates a new ImportTemplateUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportTemplateUrlWithDefaults

`func NewImportTemplateUrlWithDefaults() *ImportTemplateUrl`

NewImportTemplateUrlWithDefaults instantiates a new ImportTemplateUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *ImportTemplateUrl) GetTemplate() ImportTemplateUrlTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ImportTemplateUrl) GetTemplateOk() (*ImportTemplateUrlTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ImportTemplateUrl) SetTemplate(v ImportTemplateUrlTemplate)`

SetTemplate sets Template field to given value.


### GetFileUrl

`func (o *ImportTemplateUrl) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ImportTemplateUrl) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ImportTemplateUrl) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


