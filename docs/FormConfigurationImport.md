# FormConfigurationImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | Pointer to **NullableInt32** | Template ID which is connected to the form | [optional] 
**Name** | Pointer to **string** | Form name | [optional] 
**Actions** | Pointer to [**[]FormConfigurationNewActionsInner**](FormConfigurationNewActionsInner.md) | Array of action configurations. Each action is a separate object with a single key; leave an action out to disable it. | [optional] 
**Configuration** | Pointer to [**FormConfiguration**](FormConfiguration.md) |  | [optional] 

## Methods

### NewFormConfigurationImport

`func NewFormConfigurationImport() *FormConfigurationImport`

NewFormConfigurationImport instantiates a new FormConfigurationImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormConfigurationImportWithDefaults

`func NewFormConfigurationImportWithDefaults() *FormConfigurationImport`

NewFormConfigurationImportWithDefaults instantiates a new FormConfigurationImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *FormConfigurationImport) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *FormConfigurationImport) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *FormConfigurationImport) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *FormConfigurationImport) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *FormConfigurationImport) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *FormConfigurationImport) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetName

`func (o *FormConfigurationImport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormConfigurationImport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormConfigurationImport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormConfigurationImport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActions

`func (o *FormConfigurationImport) GetActions() []FormConfigurationNewActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *FormConfigurationImport) GetActionsOk() (*[]FormConfigurationNewActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *FormConfigurationImport) SetActions(v []FormConfigurationNewActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *FormConfigurationImport) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConfiguration

`func (o *FormConfigurationImport) GetConfiguration() FormConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *FormConfigurationImport) GetConfigurationOk() (*FormConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *FormConfigurationImport) SetConfiguration(v FormConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *FormConfigurationImport) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


