# DocumentAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action performed on the document. | [optional] 
**Person** | Pointer to [**DocumentUser**](DocumentUser.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | Action timestamp (Y-m-d H:i:s). | [optional] 

## Methods

### NewDocumentAction

`func NewDocumentAction() *DocumentAction`

NewDocumentAction instantiates a new DocumentAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentActionWithDefaults

`func NewDocumentActionWithDefaults() *DocumentAction`

NewDocumentActionWithDefaults instantiates a new DocumentAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DocumentAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DocumentAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DocumentAction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DocumentAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPerson

`func (o *DocumentAction) GetPerson() DocumentUser`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *DocumentAction) GetPersonOk() (*DocumentUser, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *DocumentAction) SetPerson(v DocumentUser)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *DocumentAction) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocumentAction) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentAction) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentAction) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentAction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


