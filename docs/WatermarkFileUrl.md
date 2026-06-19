# WatermarkFileUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | PDF file from remote URL to add the watermark to | 
**Watermark** | [**WatermarkFileUrlWatermark**](WatermarkFileUrlWatermark.md) |  | 
**Output** | Pointer to **string** | Returned document output | [optional] [default to "base64"]
**Name** | Pointer to **string** | File name of the returned document | [optional] 

## Methods

### NewWatermarkFileUrl

`func NewWatermarkFileUrl(fileUrl string, watermark WatermarkFileUrlWatermark, ) *WatermarkFileUrl`

NewWatermarkFileUrl instantiates a new WatermarkFileUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatermarkFileUrlWithDefaults

`func NewWatermarkFileUrlWithDefaults() *WatermarkFileUrl`

NewWatermarkFileUrlWithDefaults instantiates a new WatermarkFileUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *WatermarkFileUrl) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *WatermarkFileUrl) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *WatermarkFileUrl) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetWatermark

`func (o *WatermarkFileUrl) GetWatermark() WatermarkFileUrlWatermark`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *WatermarkFileUrl) GetWatermarkOk() (*WatermarkFileUrlWatermark, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *WatermarkFileUrl) SetWatermark(v WatermarkFileUrlWatermark)`

SetWatermark sets Watermark field to given value.


### GetOutput

`func (o *WatermarkFileUrl) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WatermarkFileUrl) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WatermarkFileUrl) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WatermarkFileUrl) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetName

`func (o *WatermarkFileUrl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WatermarkFileUrl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WatermarkFileUrl) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WatermarkFileUrl) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


