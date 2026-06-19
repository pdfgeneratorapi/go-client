# AddWatermarkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | PDF file from remote URL to add the watermark to | 
**Watermark** | [**WatermarkFileUrlWatermark**](WatermarkFileUrlWatermark.md) |  | 
**Output** | Pointer to **string** | Returned document output | [optional] [default to "base64"]
**Name** | Pointer to **string** | File name of the returned document | [optional] 
**FileBase64** | **string** | PDF file from base64 string to add the watermark to | 

## Methods

### NewAddWatermarkRequest

`func NewAddWatermarkRequest(fileUrl string, watermark WatermarkFileUrlWatermark, fileBase64 string, ) *AddWatermarkRequest`

NewAddWatermarkRequest instantiates a new AddWatermarkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddWatermarkRequestWithDefaults

`func NewAddWatermarkRequestWithDefaults() *AddWatermarkRequest`

NewAddWatermarkRequestWithDefaults instantiates a new AddWatermarkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *AddWatermarkRequest) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *AddWatermarkRequest) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *AddWatermarkRequest) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetWatermark

`func (o *AddWatermarkRequest) GetWatermark() WatermarkFileUrlWatermark`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *AddWatermarkRequest) GetWatermarkOk() (*WatermarkFileUrlWatermark, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *AddWatermarkRequest) SetWatermark(v WatermarkFileUrlWatermark)`

SetWatermark sets Watermark field to given value.


### GetOutput

`func (o *AddWatermarkRequest) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *AddWatermarkRequest) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *AddWatermarkRequest) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *AddWatermarkRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetName

`func (o *AddWatermarkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddWatermarkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddWatermarkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddWatermarkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFileBase64

`func (o *AddWatermarkRequest) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *AddWatermarkRequest) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *AddWatermarkRequest) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


