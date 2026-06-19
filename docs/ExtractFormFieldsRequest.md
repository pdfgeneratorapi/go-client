# ExtractFormFieldsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | Public URL to a PDF document | 
**KeyField** | Pointer to [**KeyFieldParam**](KeyFieldParam.md) |  | [optional] [default to KEYFIELDPARAM_NAME]
**FileBase64** | **string** | PDF document in base64 encoded string format | 

## Methods

### NewExtractFormFieldsRequest

`func NewExtractFormFieldsRequest(fileUrl string, fileBase64 string, ) *ExtractFormFieldsRequest`

NewExtractFormFieldsRequest instantiates a new ExtractFormFieldsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractFormFieldsRequestWithDefaults

`func NewExtractFormFieldsRequestWithDefaults() *ExtractFormFieldsRequest`

NewExtractFormFieldsRequestWithDefaults instantiates a new ExtractFormFieldsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *ExtractFormFieldsRequest) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ExtractFormFieldsRequest) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ExtractFormFieldsRequest) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetKeyField

`func (o *ExtractFormFieldsRequest) GetKeyField() KeyFieldParam`

GetKeyField returns the KeyField field if non-nil, zero value otherwise.

### GetKeyFieldOk

`func (o *ExtractFormFieldsRequest) GetKeyFieldOk() (*KeyFieldParam, bool)`

GetKeyFieldOk returns a tuple with the KeyField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyField

`func (o *ExtractFormFieldsRequest) SetKeyField(v KeyFieldParam)`

SetKeyField sets KeyField field to given value.

### HasKeyField

`func (o *ExtractFormFieldsRequest) HasKeyField() bool`

HasKeyField returns a boolean if a field has been set.

### GetFileBase64

`func (o *ExtractFormFieldsRequest) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *ExtractFormFieldsRequest) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *ExtractFormFieldsRequest) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


