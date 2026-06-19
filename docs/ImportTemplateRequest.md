# ImportTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | [**ImportTemplateUrlTemplate**](ImportTemplateUrlTemplate.md) |  | 
**FileUrl** | **string** | PDF file from remote URL to import | 
**FileBase64** | **string** | PDF file from base64 string to import | 

## Methods

### NewImportTemplateRequest

`func NewImportTemplateRequest(template ImportTemplateUrlTemplate, fileUrl string, fileBase64 string, ) *ImportTemplateRequest`

NewImportTemplateRequest instantiates a new ImportTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportTemplateRequestWithDefaults

`func NewImportTemplateRequestWithDefaults() *ImportTemplateRequest`

NewImportTemplateRequestWithDefaults instantiates a new ImportTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *ImportTemplateRequest) GetTemplate() ImportTemplateUrlTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ImportTemplateRequest) GetTemplateOk() (*ImportTemplateUrlTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ImportTemplateRequest) SetTemplate(v ImportTemplateUrlTemplate)`

SetTemplate sets Template field to given value.


### GetFileUrl

`func (o *ImportTemplateRequest) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ImportTemplateRequest) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ImportTemplateRequest) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetFileBase64

`func (o *ImportTemplateRequest) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *ImportTemplateRequest) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *ImportTemplateRequest) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


