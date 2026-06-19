# WatermarkBase64

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileBase64** | **string** | PDF file from base64 string to add the watermark to | 
**Watermark** | [**WatermarkFileUrlWatermark**](WatermarkFileUrlWatermark.md) |  | 
**Output** | Pointer to **string** | Returned document output | [optional] [default to "base64"]
**Name** | Pointer to **string** | File name of the returned document | [optional] 

## Methods

### NewWatermarkBase64

`func NewWatermarkBase64(fileBase64 string, watermark WatermarkFileUrlWatermark, ) *WatermarkBase64`

NewWatermarkBase64 instantiates a new WatermarkBase64 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatermarkBase64WithDefaults

`func NewWatermarkBase64WithDefaults() *WatermarkBase64`

NewWatermarkBase64WithDefaults instantiates a new WatermarkBase64 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileBase64

`func (o *WatermarkBase64) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *WatermarkBase64) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *WatermarkBase64) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.


### GetWatermark

`func (o *WatermarkBase64) GetWatermark() WatermarkFileUrlWatermark`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *WatermarkBase64) GetWatermarkOk() (*WatermarkFileUrlWatermark, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *WatermarkBase64) SetWatermark(v WatermarkFileUrlWatermark)`

SetWatermark sets Watermark field to given value.


### GetOutput

`func (o *WatermarkBase64) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WatermarkBase64) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WatermarkBase64) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WatermarkBase64) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetName

`func (o *WatermarkBase64) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WatermarkBase64) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WatermarkBase64) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WatermarkBase64) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


