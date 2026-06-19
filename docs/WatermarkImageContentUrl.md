# WatermarkImageContentUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentUrl** | **string** | URL to an image | 
**Position** | Pointer to [**WatermarkPosition**](WatermarkPosition.md) |  | [optional] [default to WATERMARKPOSITION_CENTER]
**Rotation** | Pointer to **int32** | Watermark rotation | [optional] [default to 0]
**Scale** | Pointer to **float32** | Watermark image scale | [optional] [default to 1]

## Methods

### NewWatermarkImageContentUrl

`func NewWatermarkImageContentUrl(contentUrl string, ) *WatermarkImageContentUrl`

NewWatermarkImageContentUrl instantiates a new WatermarkImageContentUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatermarkImageContentUrlWithDefaults

`func NewWatermarkImageContentUrlWithDefaults() *WatermarkImageContentUrl`

NewWatermarkImageContentUrlWithDefaults instantiates a new WatermarkImageContentUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentUrl

`func (o *WatermarkImageContentUrl) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *WatermarkImageContentUrl) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *WatermarkImageContentUrl) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.


### GetPosition

`func (o *WatermarkImageContentUrl) GetPosition() WatermarkPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WatermarkImageContentUrl) GetPositionOk() (*WatermarkPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *WatermarkImageContentUrl) SetPosition(v WatermarkPosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *WatermarkImageContentUrl) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetRotation

`func (o *WatermarkImageContentUrl) GetRotation() int32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *WatermarkImageContentUrl) GetRotationOk() (*int32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *WatermarkImageContentUrl) SetRotation(v int32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *WatermarkImageContentUrl) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetScale

`func (o *WatermarkImageContentUrl) GetScale() float32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *WatermarkImageContentUrl) GetScaleOk() (*float32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *WatermarkImageContentUrl) SetScale(v float32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *WatermarkImageContentUrl) HasScale() bool`

HasScale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


