# Workspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal workspace id | [optional] 
**Identifier** | Pointer to **string** | The unique workspace idenfitifer specified by your application | [optional] 
**IsMasterWorkspace** | Pointer to **bool** | True if a master workspace | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkspace

`func NewWorkspace() *Workspace`

NewWorkspace instantiates a new Workspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceWithDefaults

`func NewWorkspaceWithDefaults() *Workspace`

NewWorkspaceWithDefaults instantiates a new Workspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Workspace) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Workspace) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Workspace) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Workspace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *Workspace) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Workspace) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Workspace) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Workspace) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIsMasterWorkspace

`func (o *Workspace) GetIsMasterWorkspace() bool`

GetIsMasterWorkspace returns the IsMasterWorkspace field if non-nil, zero value otherwise.

### GetIsMasterWorkspaceOk

`func (o *Workspace) GetIsMasterWorkspaceOk() (*bool, bool)`

GetIsMasterWorkspaceOk returns a tuple with the IsMasterWorkspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterWorkspace

`func (o *Workspace) SetIsMasterWorkspace(v bool)`

SetIsMasterWorkspace sets IsMasterWorkspace field to given value.

### HasIsMasterWorkspace

`func (o *Workspace) HasIsMasterWorkspace() bool`

HasIsMasterWorkspace returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Workspace) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Workspace) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Workspace) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Workspace) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


