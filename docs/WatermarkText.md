# WatermarkText

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Watermark text | 
**Color** | Pointer to **string** | Watermark text color in hexadecimal format | [optional] [default to "#000000"]
**Size** | Pointer to **int32** | Watermark text font size in px | [optional] [default to 48]
**Opacity** | Pointer to **float32** | Watermark text opaxity | [optional] [default to 0.5]
**Position** | Pointer to [**WatermarkPosition**](WatermarkPosition.md) |  | [optional] [default to WATERMARKPOSITION_CENTER]
**Rotation** | Pointer to **int32** | Watermark rotation | [optional] [default to 0]

## Methods

### NewWatermarkText

`func NewWatermarkText(content string, ) *WatermarkText`

NewWatermarkText instantiates a new WatermarkText object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatermarkTextWithDefaults

`func NewWatermarkTextWithDefaults() *WatermarkText`

NewWatermarkTextWithDefaults instantiates a new WatermarkText object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *WatermarkText) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WatermarkText) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WatermarkText) SetContent(v string)`

SetContent sets Content field to given value.


### GetColor

`func (o *WatermarkText) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WatermarkText) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WatermarkText) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WatermarkText) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetSize

`func (o *WatermarkText) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *WatermarkText) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *WatermarkText) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *WatermarkText) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetOpacity

`func (o *WatermarkText) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *WatermarkText) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *WatermarkText) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *WatermarkText) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.

### GetPosition

`func (o *WatermarkText) GetPosition() WatermarkPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WatermarkText) GetPositionOk() (*WatermarkPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WatermarkText) SetPosition(v WatermarkPosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *WatermarkText) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetRotation

`func (o *WatermarkText) GetRotation() int32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *WatermarkText) GetRotationOk() (*int32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *WatermarkText) SetRotation(v int32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *WatermarkText) HasRotation() bool`

HasRotation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


