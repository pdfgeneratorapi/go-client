# TemplateDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier | [optional] 
**Name** | Pointer to **string** | Template name | [optional] 
**Tags** | Pointer to **[]string** | A list of tags assigned to a template | [optional] 
**IsDraft** | Pointer to **bool** | Indicates if the template is a draft or published. | [optional] 
**Layout** | Pointer to [**TemplateDefinitionNewLayout**](TemplateDefinitionNewLayout.md) |  | [optional] 
**Pages** | Pointer to [**[]TemplateDefinitionPagesInner**](TemplateDefinitionPagesInner.md) | Defines page or label size, margins and components on page or label | [optional] 
**DataSettings** | Pointer to [**TemplateDefinitionDataSettings**](TemplateDefinitionDataSettings.md) |  | [optional] 
**Editor** | Pointer to [**TemplateDefinitionEditor**](TemplateDefinitionEditor.md) |  | [optional] 

## Methods

### NewTemplateDefinition

`func NewTemplateDefinition() *TemplateDefinition`

NewTemplateDefinition instantiates a new TemplateDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDefinitionWithDefaults

`func NewTemplateDefinitionWithDefaults() *TemplateDefinition`

NewTemplateDefinitionWithDefaults instantiates a new TemplateDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateDefinition) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateDefinition) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateDefinition) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TemplateDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *TemplateDefinition) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TemplateDefinition) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TemplateDefinition) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TemplateDefinition) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsDraft

`func (o *TemplateDefinition) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *TemplateDefinition) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *TemplateDefinition) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *TemplateDefinition) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### GetLayout

`func (o *TemplateDefinition) GetLayout() TemplateDefinitionNewLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *TemplateDefinition) GetLayoutOk() (*TemplateDefinitionNewLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *TemplateDefinition) SetLayout(v TemplateDefinitionNewLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *TemplateDefinition) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetPages

`func (o *TemplateDefinition) GetPages() []TemplateDefinitionPagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *TemplateDefinition) GetPagesOk() (*[]TemplateDefinitionPagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *TemplateDefinition) SetPages(v []TemplateDefinitionPagesInner)`

SetPages sets Pages field to given value.

### HasPages

`func (o *TemplateDefinition) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetDataSettings

`func (o *TemplateDefinition) GetDataSettings() TemplateDefinitionDataSettings`

GetDataSettings returns the DataSettings field if non-nil, zero value otherwise.

### GetDataSettingsOk

`func (o *TemplateDefinition) GetDataSettingsOk() (*TemplateDefinitionDataSettings, bool)`

GetDataSettingsOk returns a tuple with the DataSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSettings

`func (o *TemplateDefinition) SetDataSettings(v TemplateDefinitionDataSettings)`

SetDataSettings sets DataSettings field to given value.

### HasDataSettings

`func (o *TemplateDefinition) HasDataSettings() bool`

HasDataSettings returns a boolean if a field has been set.

### GetEditor

`func (o *TemplateDefinition) GetEditor() TemplateDefinitionEditor`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *TemplateDefinition) GetEditorOk() (*TemplateDefinitionEditor, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *TemplateDefinition) SetEditor(v TemplateDefinitionEditor)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *TemplateDefinition) HasEditor() bool`

HasEditor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


