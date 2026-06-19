# EncryptAndDecryptBase64

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileBase64** | **string** | PDF document in base64 encoded string format | 
**OwnerPassword** | **string** | An owner password to open the encrypted document | 
**UserPassword** | Pointer to **string** | An user password to open the encrypted document | [optional] 
**Name** | Pointer to **string** | Name for the PDF file | [optional] 
**Output** | Pointer to **string** | Returned document output format | [optional] [default to "base64"]

## Methods

### NewEncryptAndDecryptBase64

`func NewEncryptAndDecryptBase64(fileBase64 string, ownerPassword string, ) *EncryptAndDecryptBase64`

NewEncryptAndDecryptBase64 instantiates a new EncryptAndDecryptBase64 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptAndDecryptBase64WithDefaults

`func NewEncryptAndDecryptBase64WithDefaults() *EncryptAndDecryptBase64`

NewEncryptAndDecryptBase64WithDefaults instantiates a new EncryptAndDecryptBase64 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileBase64

`func (o *EncryptAndDecryptBase64) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *EncryptAndDecryptBase64) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *EncryptAndDecryptBase64) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.


### GetOwnerPassword

`func (o *EncryptAndDecryptBase64) GetOwnerPassword() string`

GetOwnerPassword returns the OwnerPassword field if non-nil, zero value otherwise.

### GetOwnerPasswordOk

`func (o *EncryptAndDecryptBase64) GetOwnerPasswordOk() (*string, bool)`

GetOwnerPasswordOk returns a tuple with the OwnerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPassword

`func (o *EncryptAndDecryptBase64) SetOwnerPassword(v string)`

SetOwnerPassword sets OwnerPassword field to given value.


### GetUserPassword

`func (o *EncryptAndDecryptBase64) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *EncryptAndDecryptBase64) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *EncryptAndDecryptBase64) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *EncryptAndDecryptBase64) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

### GetName

`func (o *EncryptAndDecryptBase64) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EncryptAndDecryptBase64) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EncryptAndDecryptBase64) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EncryptAndDecryptBase64) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *EncryptAndDecryptBase64) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *EncryptAndDecryptBase64) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *EncryptAndDecryptBase64) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *EncryptAndDecryptBase64) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


