# EncryptAndDecryptUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | Public URL to a PDF document | 
**OwnerPassword** | **string** | An owner password to open the encrypted document | 
**UserPassword** | Pointer to **string** | An user password to open the encrypted document | [optional] 
**Name** | Pointer to **string** | Name for the PDF file | [optional] 
**Output** | Pointer to **string** | Returned document output format | [optional] [default to "base64"]

## Methods

### NewEncryptAndDecryptUrl

`func NewEncryptAndDecryptUrl(fileUrl string, ownerPassword string, ) *EncryptAndDecryptUrl`

NewEncryptAndDecryptUrl instantiates a new EncryptAndDecryptUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptAndDecryptUrlWithDefaults

`func NewEncryptAndDecryptUrlWithDefaults() *EncryptAndDecryptUrl`

NewEncryptAndDecryptUrlWithDefaults instantiates a new EncryptAndDecryptUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *EncryptAndDecryptUrl) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *EncryptAndDecryptUrl) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *EncryptAndDecryptUrl) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetOwnerPassword

`func (o *EncryptAndDecryptUrl) GetOwnerPassword() string`

GetOwnerPassword returns the OwnerPassword field if non-nil, zero value otherwise.

### GetOwnerPasswordOk

`func (o *EncryptAndDecryptUrl) GetOwnerPasswordOk() (*string, bool)`

GetOwnerPasswordOk returns a tuple with the OwnerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPassword

`func (o *EncryptAndDecryptUrl) SetOwnerPassword(v string)`

SetOwnerPassword sets OwnerPassword field to given value.


### GetUserPassword

`func (o *EncryptAndDecryptUrl) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *EncryptAndDecryptUrl) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *EncryptAndDecryptUrl) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *EncryptAndDecryptUrl) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

### GetName

`func (o *EncryptAndDecryptUrl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EncryptAndDecryptUrl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EncryptAndDecryptUrl) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EncryptAndDecryptUrl) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *EncryptAndDecryptUrl) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *EncryptAndDecryptUrl) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *EncryptAndDecryptUrl) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *EncryptAndDecryptUrl) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


