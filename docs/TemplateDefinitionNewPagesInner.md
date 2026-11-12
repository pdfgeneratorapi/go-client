# TemplateDefinitionNewPagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Width** | Pointer to **float32** | Page width in units | [optional] 
**Height** | Pointer to **float32** | Page height in units | [optional] 
**Margins** | Pointer to [**TemplateDefinitionNewPagesInnerMargins**](TemplateDefinitionNewPagesInnerMargins.md) |  | [optional] 
**Border** | Pointer to **bool** | Stored for older templates. It draws nothing. | [optional] [default to false]
**Components** | Pointer to **[]map[string]interface{}** | The page&#39;s components. The fields of each depend on its cls; the template schema (Get schema endpoint) lists them. | [optional] 
**Layout** | Pointer to [**TemplateDefinitionNewPagesInnerLayout**](TemplateDefinitionNewPagesInnerLayout.md) |  | [optional] 
**ConditionalFormats** | Pointer to **[]map[string]interface{}** |  | [optional] 
**BackgroundImage** | Pointer to **NullableString** | The page&#39;s background image: an image URL, a data URI or a {data field} placeholder, PNG, JPEG, GIF or SVG. It is drawn over the background color and under everything else on the page, as backgroundImageFit says. A template with a background PDF does not draw it. Empty or null for none. | [optional] 
**BackgroundColor** | Pointer to **NullableString** | The page&#39;s background color in hex color code, #RRGGBB or #RGB, or with an alpha, #RRGGBBAA, #RGBA or rgba(r, g, b, a); a translucent one is lighter, the white of the paper showing through. It fills every sheet the page runs onto, edge to edge, under everything on it, the background image included; on a sheet of labels it fills each label. A template with a background PDF is not filled. Empty or null for none. | [optional] 
**BackgroundImageFit** | Pointer to **NullableString** | How the background image covers the page, centred: &#x60;stretch&#x60;, stretched to the page, as when none is set; &#x60;contain&#x60;, fitted inside the page, keeping its proportions; &#x60;cover&#x60;, filling the page, keeping its proportions, with what overflows cut off. Null for stretch. | [optional] 

## Methods

### NewTemplateDefinitionNewPagesInner

`func NewTemplateDefinitionNewPagesInner() *TemplateDefinitionNewPagesInner`

NewTemplateDefinitionNewPagesInner instantiates a new TemplateDefinitionNewPagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDefinitionNewPagesInnerWithDefaults

`func NewTemplateDefinitionNewPagesInnerWithDefaults() *TemplateDefinitionNewPagesInner`

NewTemplateDefinitionNewPagesInnerWithDefaults instantiates a new TemplateDefinitionNewPagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWidth

`func (o *TemplateDefinitionNewPagesInner) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *TemplateDefinitionNewPagesInner) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *TemplateDefinitionNewPagesInner) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *TemplateDefinitionNewPagesInner) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *TemplateDefinitionNewPagesInner) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *TemplateDefinitionNewPagesInner) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *TemplateDefinitionNewPagesInner) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *TemplateDefinitionNewPagesInner) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetMargins

`func (o *TemplateDefinitionNewPagesInner) GetMargins() TemplateDefinitionNewPagesInnerMargins`

GetMargins returns the Margins field if non-nil, zero value otherwise.

### GetMarginsOk

`func (o *TemplateDefinitionNewPagesInner) GetMarginsOk() (*TemplateDefinitionNewPagesInnerMargins, bool)`

GetMarginsOk returns a tuple with the Margins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargins

`func (o *TemplateDefinitionNewPagesInner) SetMargins(v TemplateDefinitionNewPagesInnerMargins)`

SetMargins sets Margins field to given value.

### HasMargins

`func (o *TemplateDefinitionNewPagesInner) HasMargins() bool`

HasMargins returns a boolean if a field has been set.

### GetBorder

`func (o *TemplateDefinitionNewPagesInner) GetBorder() bool`

GetBorder returns the Border field if non-nil, zero value otherwise.

### GetBorderOk

`func (o *TemplateDefinitionNewPagesInner) GetBorderOk() (*bool, bool)`

GetBorderOk returns a tuple with the Border field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorder

`func (o *TemplateDefinitionNewPagesInner) SetBorder(v bool)`

SetBorder sets Border field to given value.

### HasBorder

`func (o *TemplateDefinitionNewPagesInner) HasBorder() bool`

HasBorder returns a boolean if a field has been set.

### GetComponents

`func (o *TemplateDefinitionNewPagesInner) GetComponents() []map[string]interface{}`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *TemplateDefinitionNewPagesInner) GetComponentsOk() (*[]map[string]interface{}, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *TemplateDefinitionNewPagesInner) SetComponents(v []map[string]interface{})`

SetComponents sets Components field to given value.

### HasComponents

`func (o *TemplateDefinitionNewPagesInner) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetLayout

`func (o *TemplateDefinitionNewPagesInner) GetLayout() TemplateDefinitionNewPagesInnerLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *TemplateDefinitionNewPagesInner) GetLayoutOk() (*TemplateDefinitionNewPagesInnerLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *TemplateDefinitionNewPagesInner) SetLayout(v TemplateDefinitionNewPagesInnerLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *TemplateDefinitionNewPagesInner) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetConditionalFormats

`func (o *TemplateDefinitionNewPagesInner) GetConditionalFormats() []map[string]interface{}`

GetConditionalFormats returns the ConditionalFormats field if non-nil, zero value otherwise.

### GetConditionalFormatsOk

`func (o *TemplateDefinitionNewPagesInner) GetConditionalFormatsOk() (*[]map[string]interface{}, bool)`

GetConditionalFormatsOk returns a tuple with the ConditionalFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalFormats

`func (o *TemplateDefinitionNewPagesInner) SetConditionalFormats(v []map[string]interface{})`

SetConditionalFormats sets ConditionalFormats field to given value.

### HasConditionalFormats

`func (o *TemplateDefinitionNewPagesInner) HasConditionalFormats() bool`

HasConditionalFormats returns a boolean if a field has been set.

### GetBackgroundImage

`func (o *TemplateDefinitionNewPagesInner) GetBackgroundImage() string`

GetBackgroundImage returns the BackgroundImage field if non-nil, zero value otherwise.

### GetBackgroundImageOk

`func (o *TemplateDefinitionNewPagesInner) GetBackgroundImageOk() (*string, bool)`

GetBackgroundImageOk returns a tuple with the BackgroundImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImage

`func (o *TemplateDefinitionNewPagesInner) SetBackgroundImage(v string)`

SetBackgroundImage sets BackgroundImage field to given value.

### HasBackgroundImage

`func (o *TemplateDefinitionNewPagesInner) HasBackgroundImage() bool`

HasBackgroundImage returns a boolean if a field has been set.

### SetBackgroundImageNil

`func (o *TemplateDefinitionNewPagesInner) SetBackgroundImageNil(b bool)`

 SetBackgroundImageNil sets the value for BackgroundImage to be an explicit nil

### UnsetBackgroundImage
`func (o *TemplateDefinitionNewPagesInner) UnsetBackgroundImage()`

UnsetBackgroundImage ensures that no value is present for BackgroundImage, not even an explicit nil
### GetBackgroundColor

`func (o *TemplateDefinitionNewPagesInner) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *TemplateDefinitionNewPagesInner) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *TemplateDefinitionNewPagesInner) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *TemplateDefinitionNewPagesInner) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *TemplateDefinitionNewPagesInner) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *TemplateDefinitionNewPagesInner) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetBackgroundImageFit

`func (o *TemplateDefinitionNewPagesInner) GetBackgroundImageFit() string`

GetBackgroundImageFit returns the BackgroundImageFit field if non-nil, zero value otherwise.

### GetBackgroundImageFitOk

`func (o *TemplateDefinitionNewPagesInner) GetBackgroundImageFitOk() (*string, bool)`

GetBackgroundImageFitOk returns a tuple with the BackgroundImageFit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImageFit

`func (o *TemplateDefinitionNewPagesInner) SetBackgroundImageFit(v string)`

SetBackgroundImageFit sets BackgroundImageFit field to given value.

### HasBackgroundImageFit

`func (o *TemplateDefinitionNewPagesInner) HasBackgroundImageFit() bool`

HasBackgroundImageFit returns a boolean if a field has been set.

### SetBackgroundImageFitNil

`func (o *TemplateDefinitionNewPagesInner) SetBackgroundImageFitNil(b bool)`

 SetBackgroundImageFitNil sets the value for BackgroundImageFit to be an explicit nil

### UnsetBackgroundImageFit
`func (o *TemplateDefinitionNewPagesInner) UnsetBackgroundImageFit()`

UnsetBackgroundImageFit ensures that no value is present for BackgroundImageFit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


