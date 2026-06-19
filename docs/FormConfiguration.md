# FormConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | Form title displayed above the form | [optional] 
**Intro** | Pointer to **NullableString** | Introductory text displayed above the form fields. Supports rich text. | [optional] 
**Outro** | Pointer to **NullableString** | Closing text displayed below the form fields. Supports rich text. | [optional] 

## Methods

### NewFormConfiguration

`func NewFormConfiguration() *FormConfiguration`

NewFormConfiguration instantiates a new FormConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormConfigurationWithDefaults

`func NewFormConfigurationWithDefaults() *FormConfiguration`

NewFormConfigurationWithDefaults instantiates a new FormConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *FormConfiguration) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FormConfiguration) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FormConfiguration) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FormConfiguration) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FormConfiguration) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FormConfiguration) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetIntro

`func (o *FormConfiguration) GetIntro() string`

GetIntro returns the Intro field if non-nil, zero value otherwise.

### GetIntroOk

`func (o *FormConfiguration) GetIntroOk() (*string, bool)`

GetIntroOk returns a tuple with the Intro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntro

`func (o *FormConfiguration) SetIntro(v string)`

SetIntro sets Intro field to given value.

### HasIntro

`func (o *FormConfiguration) HasIntro() bool`

HasIntro returns a boolean if a field has been set.

### SetIntroNil

`func (o *FormConfiguration) SetIntroNil(b bool)`

 SetIntroNil sets the value for Intro to be an explicit nil

### UnsetIntro
`func (o *FormConfiguration) UnsetIntro()`

UnsetIntro ensures that no value is present for Intro, not even an explicit nil
### GetOutro

`func (o *FormConfiguration) GetOutro() string`

GetOutro returns the Outro field if non-nil, zero value otherwise.

### GetOutroOk

`func (o *FormConfiguration) GetOutroOk() (*string, bool)`

GetOutroOk returns a tuple with the Outro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutro

`func (o *FormConfiguration) SetOutro(v string)`

SetOutro sets Outro field to given value.

### HasOutro

`func (o *FormConfiguration) HasOutro() bool`

HasOutro returns a boolean if a field has been set.

### SetOutroNil

`func (o *FormConfiguration) SetOutroNil(b bool)`

 SetOutroNil sets the value for Outro to be an explicit nil

### UnsetOutro
`func (o *FormConfiguration) UnsetOutro()`

UnsetOutro ensures that no value is present for Outro, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


