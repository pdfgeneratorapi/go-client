# EncryptDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | Public URL to a PDF document | 
**OwnerPassword** | **string** | An owner password to open the encrypted document | 
**UserPassword** | Pointer to **string** | An user password to open the encrypted document | [optional] 
**Name** | Pointer to **string** | Name for the PDF file | [optional] 
**Output** | Pointer to **string** | Returned document output format | [optional] [default to "base64"]
**FileBase64** | **string** | PDF document in base64 encoded string format | 

## Methods

### NewEncryptDocumentRequest

`func NewEncryptDocumentRequest(fileUrl string, ownerPassword string, fileBase64 string, ) *EncryptDocumentRequest`

NewEncryptDocumentRequest instantiates a new EncryptDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptDocumentRequestWithDefaults

`func NewEncryptDocumentRequestWithDefaults() *EncryptDocumentRequest`

NewEncryptDocumentRequestWithDefaults instantiates a new EncryptDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *EncryptDocumentRequest) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *EncryptDocumentRequest) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *EncryptDocumentRequest) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetOwnerPassword

`func (o *EncryptDocumentRequest) GetOwnerPassword() string`

GetOwnerPassword returns the OwnerPassword field if non-nil, zero value otherwise.

### GetOwnerPasswordOk

`func (o *EncryptDocumentRequest) GetOwnerPasswordOk() (*string, bool)`

GetOwnerPasswordOk returns a tuple with the OwnerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPassword

`func (o *EncryptDocumentRequest) SetOwnerPassword(v string)`

SetOwnerPassword sets OwnerPassword field to given value.


### GetUserPassword

`func (o *EncryptDocumentRequest) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *EncryptDocumentRequest) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *EncryptDocumentRequest) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *EncryptDocumentRequest) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

### GetName

`func (o *EncryptDocumentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EncryptDocumentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EncryptDocumentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EncryptDocumentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *EncryptDocumentRequest) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *EncryptDocumentRequest) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *EncryptDocumentRequest) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *EncryptDocumentRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetFileBase64

`func (o *EncryptDocumentRequest) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *EncryptDocumentRequest) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *EncryptDocumentRequest) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


