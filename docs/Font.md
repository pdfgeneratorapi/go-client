# Font

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The value of a component&#39;s &#x60;fontFamily&#x60; for this font | 
**Name** | **string** | Font name, as the editor shows it | 
**Parent** | **NullableString** | The &#x60;id&#x60; of the font this one is grouped under, such as another weight of the same family | 
**Styles** | **[]string** | The styles the font has glyphs of its own for. Text in another style is drawn with the regular glyphs. | 
**Blocks** | **map[string]int32** | The Unicode blocks the font has at least half of the characters of, with the share it has in percent. 100 means every character of the block. | 

## Methods

### NewFont

`func NewFont(id string, name string, parent NullableString, styles []string, blocks map[string]int32, ) *Font`

NewFont instantiates a new Font object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFontWithDefaults

`func NewFontWithDefaults() *Font`

NewFontWithDefaults instantiates a new Font object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Font) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Font) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Font) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Font) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Font) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Font) SetName(v string)`

SetName sets Name field to given value.


### GetParent

`func (o *Font) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Font) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Font) SetParent(v string)`

SetParent sets Parent field to given value.


### SetParentNil

`func (o *Font) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *Font) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetStyles

`func (o *Font) GetStyles() []string`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *Font) GetStylesOk() (*[]string, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *Font) SetStyles(v []string)`

SetStyles sets Styles field to given value.


### GetBlocks

`func (o *Font) GetBlocks() map[string]int32`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *Font) GetBlocksOk() (*map[string]int32, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *Font) SetBlocks(v map[string]int32)`

SetBlocks sets Blocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


