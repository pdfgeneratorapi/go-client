# WatermarkImageContentBase64

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentBase64** | **string** | Base64 image string | 
**Position** | Pointer to [**WatermarkPosition**](WatermarkPosition.md) |  | [optional] [default to WATERMARKPOSITION_CENTER]
**Rotation** | Pointer to **int32** | Watermark rotation | [optional] [default to 0]
**Scale** | Pointer to **float32** | Watermark image scale | [optional] [default to 1]

## Methods

### NewWatermarkImageContentBase64

`func NewWatermarkImageContentBase64(contentBase64 string, ) *WatermarkImageContentBase64`

NewWatermarkImageContentBase64 instantiates a new WatermarkImageContentBase64 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatermarkImageContentBase64WithDefaults

`func NewWatermarkImageContentBase64WithDefaults() *WatermarkImageContentBase64`

NewWatermarkImageContentBase64WithDefaults instantiates a new WatermarkImageContentBase64 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentBase64

`func (o *WatermarkImageContentBase64) GetContentBase64() string`

GetContentBase64 returns the ContentBase64 field if non-nil, zero value otherwise.

### GetContentBase64Ok

`func (o *WatermarkImageContentBase64) GetContentBase64Ok() (*string, bool)`

GetContentBase64Ok returns a tuple with the ContentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBase64

`func (o *WatermarkImageContentBase64) SetContentBase64(v string)`

SetContentBase64 sets ContentBase64 field to given value.


### GetPosition

`func (o *WatermarkImageContentBase64) GetPosition() WatermarkPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WatermarkImageContentBase64) GetPositionOk() (*WatermarkPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WatermarkImageContentBase64) SetPosition(v WatermarkPosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *WatermarkImageContentBase64) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetRotation

`func (o *WatermarkImageContentBase64) GetRotation() int32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *WatermarkImageContentBase64) GetRotationOk() (*int32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *WatermarkImageContentBase64) SetRotation(v int32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *WatermarkImageContentBase64) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetScale

`func (o *WatermarkImageContentBase64) GetScale() float32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *WatermarkImageContentBase64) GetScaleOk() (*float32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *WatermarkImageContentBase64) SetScale(v float32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *WatermarkImageContentBase64) HasScale() bool`

HasScale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


