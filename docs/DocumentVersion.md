# DocumentVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | Pointer to **string** | Version identifier (hash). | [optional] 
**Name** | Pointer to **string** | Version name. | [optional] 
**Person** | Pointer to [**DocumentUser**](DocumentUser.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | Creation timestamp (Y-m-d H:i:s). | [optional] 

## Methods

### NewDocumentVersion

`func NewDocumentVersion() *DocumentVersion`

NewDocumentVersion instantiates a new DocumentVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentVersionWithDefaults

`func NewDocumentVersionWithDefaults() *DocumentVersion`

NewDocumentVersionWithDefaults instantiates a new DocumentVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *DocumentVersion) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *DocumentVersion) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *DocumentVersion) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *DocumentVersion) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetName

`func (o *DocumentVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPerson

`func (o *DocumentVersion) GetPerson() DocumentUser`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *DocumentVersion) GetPersonOk() (*DocumentUser, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *DocumentVersion) SetPerson(v DocumentUser)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *DocumentVersion) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocumentVersion) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentVersion) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentVersion) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


