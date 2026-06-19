# CreateWorkspaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | A unique identifier of the workspace. Make sure that you can generate the same identifier for your user. | [optional] 

## Methods

### NewCreateWorkspaceRequest

`func NewCreateWorkspaceRequest() *CreateWorkspaceRequest`

NewCreateWorkspaceRequest instantiates a new CreateWorkspaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkspaceRequestWithDefaults

`func NewCreateWorkspaceRequestWithDefaults() *CreateWorkspaceRequest`

NewCreateWorkspaceRequestWithDefaults instantiates a new CreateWorkspaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CreateWorkspaceRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateWorkspaceRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateWorkspaceRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CreateWorkspaceRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


