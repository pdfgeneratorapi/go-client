# TemplateVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TemplateId** | Pointer to **int32** |  | [optional] 
**TemplateName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Checksum** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**IsProduction** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewTemplateVersion

`func NewTemplateVersion() *TemplateVersion`

NewTemplateVersion instantiates a new TemplateVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateVersionWithDefaults

`func NewTemplateVersionWithDefaults() *TemplateVersion`

NewTemplateVersionWithDefaults instantiates a new TemplateVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateVersion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateVersion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateVersion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTemplateId

`func (o *TemplateVersion) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplateVersion) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplateVersion) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *TemplateVersion) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateName

`func (o *TemplateVersion) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *TemplateVersion) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *TemplateVersion) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *TemplateVersion) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetName

`func (o *TemplateVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetChecksum

`func (o *TemplateVersion) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *TemplateVersion) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *TemplateVersion) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *TemplateVersion) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetUserId

`func (o *TemplateVersion) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TemplateVersion) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TemplateVersion) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TemplateVersion) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *TemplateVersion) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *TemplateVersion) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *TemplateVersion) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *TemplateVersion) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetIsProduction

`func (o *TemplateVersion) GetIsProduction() bool`

GetIsProduction returns the IsProduction field if non-nil, zero value otherwise.

### GetIsProductionOk

`func (o *TemplateVersion) GetIsProductionOk() (*bool, bool)`

GetIsProductionOk returns a tuple with the IsProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProduction

`func (o *TemplateVersion) SetIsProduction(v bool)`

SetIsProduction sets IsProduction field to given value.

### HasIsProduction

`func (o *TemplateVersion) HasIsProduction() bool`

HasIsProduction returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TemplateVersion) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateVersion) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateVersion) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TemplateVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TemplateVersion) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemplateVersion) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemplateVersion) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TemplateVersion) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


