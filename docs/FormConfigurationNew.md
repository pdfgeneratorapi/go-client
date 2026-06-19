# FormConfigurationNew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | Pointer to **int32** | Template ID which is connected to the form | [optional] 
**Name** | Pointer to **string** | Form name | [optional] 
**Actions** | Pointer to [**[]FormConfigurationNewActionsInner**](FormConfigurationNewActionsInner.md) | Array of action configurations | [optional] 
**Fields** | Pointer to [**[]FormFieldsInner**](FormFieldsInner.md) | A list of form field objects | [optional] 
**Configuration** | Pointer to [**FormConfiguration**](FormConfiguration.md) |  | [optional] 

## Methods

### NewFormConfigurationNew

`func NewFormConfigurationNew() *FormConfigurationNew`

NewFormConfigurationNew instantiates a new FormConfigurationNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormConfigurationNewWithDefaults

`func NewFormConfigurationNewWithDefaults() *FormConfigurationNew`

NewFormConfigurationNewWithDefaults instantiates a new FormConfigurationNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *FormConfigurationNew) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *FormConfigurationNew) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *FormConfigurationNew) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *FormConfigurationNew) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetName

`func (o *FormConfigurationNew) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormConfigurationNew) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormConfigurationNew) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormConfigurationNew) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActions

`func (o *FormConfigurationNew) GetActions() []FormConfigurationNewActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *FormConfigurationNew) GetActionsOk() (*[]FormConfigurationNewActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *FormConfigurationNew) SetActions(v []FormConfigurationNewActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *FormConfigurationNew) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetFields

`func (o *FormConfigurationNew) GetFields() []FormFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FormConfigurationNew) GetFieldsOk() (*[]FormFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FormConfigurationNew) SetFields(v []FormFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *FormConfigurationNew) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *FormConfigurationNew) GetConfiguration() FormConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *FormConfigurationNew) GetConfigurationOk() (*FormConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *FormConfigurationNew) SetConfiguration(v FormConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *FormConfigurationNew) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


