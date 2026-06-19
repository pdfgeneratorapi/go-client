# TemplateDefinitionNewLayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | Defines template page size | [optional] [default to "A4"]
**Width** | Pointer to **float32** | Page width in units | [optional] 
**Height** | Pointer to **float32** | Page height in units | [optional] 
**Unit** | Pointer to **string** | Measure unit | [optional] [default to "cm"]
**Orientation** | Pointer to **string** | Page orientation | [optional] 
**Rotation** | Pointer to **int32** | Page rotation in degrees | [optional] 
**Margins** | Pointer to [**TemplateDefinitionNewLayoutMargins**](TemplateDefinitionNewLayoutMargins.md) |  | [optional] 
**RepeatLayout** | Pointer to [**NullableTemplateDefinitionNewLayoutRepeatLayout**](TemplateDefinitionNewLayoutRepeatLayout.md) |  | [optional] 
**EmptyLabels** | Pointer to **int32** | Specifies how many blank lables to add to sheet label. | [optional] [default to 0]

## Methods

### NewTemplateDefinitionNewLayout

`func NewTemplateDefinitionNewLayout() *TemplateDefinitionNewLayout`

NewTemplateDefinitionNewLayout instantiates a new TemplateDefinitionNewLayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDefinitionNewLayoutWithDefaults

`func NewTemplateDefinitionNewLayoutWithDefaults() *TemplateDefinitionNewLayout`

NewTemplateDefinitionNewLayoutWithDefaults instantiates a new TemplateDefinitionNewLayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *TemplateDefinitionNewLayout) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TemplateDefinitionNewLayout) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TemplateDefinitionNewLayout) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *TemplateDefinitionNewLayout) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetWidth

`func (o *TemplateDefinitionNewLayout) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *TemplateDefinitionNewLayout) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *TemplateDefinitionNewLayout) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *TemplateDefinitionNewLayout) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *TemplateDefinitionNewLayout) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *TemplateDefinitionNewLayout) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *TemplateDefinitionNewLayout) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *TemplateDefinitionNewLayout) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetUnit

`func (o *TemplateDefinitionNewLayout) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TemplateDefinitionNewLayout) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TemplateDefinitionNewLayout) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TemplateDefinitionNewLayout) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetOrientation

`func (o *TemplateDefinitionNewLayout) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *TemplateDefinitionNewLayout) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *TemplateDefinitionNewLayout) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *TemplateDefinitionNewLayout) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetRotation

`func (o *TemplateDefinitionNewLayout) GetRotation() int32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *TemplateDefinitionNewLayout) GetRotationOk() (*int32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *TemplateDefinitionNewLayout) SetRotation(v int32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *TemplateDefinitionNewLayout) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetMargins

`func (o *TemplateDefinitionNewLayout) GetMargins() TemplateDefinitionNewLayoutMargins`

GetMargins returns the Margins field if non-nil, zero value otherwise.

### GetMarginsOk

`func (o *TemplateDefinitionNewLayout) GetMarginsOk() (*TemplateDefinitionNewLayoutMargins, bool)`

GetMarginsOk returns a tuple with the Margins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargins

`func (o *TemplateDefinitionNewLayout) SetMargins(v TemplateDefinitionNewLayoutMargins)`

SetMargins sets Margins field to given value.

### HasMargins

`func (o *TemplateDefinitionNewLayout) HasMargins() bool`

HasMargins returns a boolean if a field has been set.

### GetRepeatLayout

`func (o *TemplateDefinitionNewLayout) GetRepeatLayout() TemplateDefinitionNewLayoutRepeatLayout`

GetRepeatLayout returns the RepeatLayout field if non-nil, zero value otherwise.

### GetRepeatLayoutOk

`func (o *TemplateDefinitionNewLayout) GetRepeatLayoutOk() (*TemplateDefinitionNewLayoutRepeatLayout, bool)`

GetRepeatLayoutOk returns a tuple with the RepeatLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatLayout

`func (o *TemplateDefinitionNewLayout) SetRepeatLayout(v TemplateDefinitionNewLayoutRepeatLayout)`

SetRepeatLayout sets RepeatLayout field to given value.

### HasRepeatLayout

`func (o *TemplateDefinitionNewLayout) HasRepeatLayout() bool`

HasRepeatLayout returns a boolean if a field has been set.

### SetRepeatLayoutNil

`func (o *TemplateDefinitionNewLayout) SetRepeatLayoutNil(b bool)`

 SetRepeatLayoutNil sets the value for RepeatLayout to be an explicit nil

### UnsetRepeatLayout
`func (o *TemplateDefinitionNewLayout) UnsetRepeatLayout()`

UnsetRepeatLayout ensures that no value is present for RepeatLayout, not even an explicit nil
### GetEmptyLabels

`func (o *TemplateDefinitionNewLayout) GetEmptyLabels() int32`

GetEmptyLabels returns the EmptyLabels field if non-nil, zero value otherwise.

### GetEmptyLabelsOk

`func (o *TemplateDefinitionNewLayout) GetEmptyLabelsOk() (*int32, bool)`

GetEmptyLabelsOk returns a tuple with the EmptyLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyLabels

`func (o *TemplateDefinitionNewLayout) SetEmptyLabels(v int32)`

SetEmptyLabels sets EmptyLabels field to given value.

### HasEmptyLabels

`func (o *TemplateDefinitionNewLayout) HasEmptyLabels() bool`

HasEmptyLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


