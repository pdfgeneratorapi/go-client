# OptimizeDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | Public URL to a PDF document | 
**Name** | Pointer to **string** | Name for the PDF file | [optional] 
**Output** | Pointer to **string** | Returned document output format | [optional] [default to "base64"]
**FileBase64** | **string** | PDF document in base64 encoded string format | 

## Methods

### NewOptimizeDocumentRequest

`func NewOptimizeDocumentRequest(fileUrl string, fileBase64 string, ) *OptimizeDocumentRequest`

NewOptimizeDocumentRequest instantiates a new OptimizeDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizeDocumentRequestWithDefaults

`func NewOptimizeDocumentRequestWithDefaults() *OptimizeDocumentRequest`

NewOptimizeDocumentRequestWithDefaults instantiates a new OptimizeDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *OptimizeDocumentRequest) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *OptimizeDocumentRequest) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *OptimizeDocumentRequest) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetName

`func (o *OptimizeDocumentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptimizeDocumentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptimizeDocumentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptimizeDocumentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *OptimizeDocumentRequest) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *OptimizeDocumentRequest) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *OptimizeDocumentRequest) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *OptimizeDocumentRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetFileBase64

`func (o *OptimizeDocumentRequest) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *OptimizeDocumentRequest) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *OptimizeDocumentRequest) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


