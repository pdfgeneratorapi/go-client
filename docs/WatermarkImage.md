# WatermarkImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentUrl** | **string** | URL to an image | 
**Position** | Pointer to [**WatermarkPosition**](WatermarkPosition.md) |  | [optional] [default to WATERMARKPOSITION_CENTER]
**Rotation** | Pointer to **int32** | Watermark rotation | [optional] [default to 0]
**Scale** | Pointer to **float32** | Watermark image scale | [optional] [default to 1]
**ContentBase64** | **string** | Base64 image string | 

## Methods

### NewWatermarkImage

`func NewWatermarkImage(contentUrl string, contentBase64 string, ) *WatermarkImage`

NewWatermarkImage instantiates a new WatermarkImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatermarkImageWithDefaults

`func NewWatermarkImageWithDefaults() *WatermarkImage`

NewWatermarkImageWithDefaults instantiates a new WatermarkImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentUrl

`func (o *WatermarkImage) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *WatermarkImage) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *WatermarkImage) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.


### GetPosition

`func (o *WatermarkImage) GetPosition() WatermarkPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WatermarkImage) GetPositionOk() (*WatermarkPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WatermarkImage) SetPosition(v WatermarkPosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *WatermarkImage) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetRotation

`func (o *WatermarkImage) GetRotation() int32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *WatermarkImage) GetRotationOk() (*int32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *WatermarkImage) SetRotation(v int32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *WatermarkImage) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetScale

`func (o *WatermarkImage) GetScale() float32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *WatermarkImage) GetScaleOk() (*float32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *WatermarkImage) SetScale(v float32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *WatermarkImage) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetContentBase64

`func (o *WatermarkImage) GetContentBase64() string`

GetContentBase64 returns the ContentBase64 field if non-nil, zero value otherwise.

### GetContentBase64Ok

`func (o *WatermarkImage) GetContentBase64Ok() (*string, bool)`

GetContentBase64Ok returns a tuple with the ContentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBase64

`func (o *WatermarkImage) SetContentBase64(v string)`

SetContentBase64 sets ContentBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


