# FillFormFieldsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | Public URL to a PDF document | 
**Data** | **map[string]interface{}** | Form field data to fill in the PDF | 
**Output** | Pointer to **string** | Returned document output format | [optional] [default to "base64"]
**Name** | Pointer to **string** | Name for the PDF file | [optional] 
**FileBase64** | **string** | PDF document in base64 encoded string format | 

## Methods

### NewFillFormFieldsRequest

`func NewFillFormFieldsRequest(fileUrl string, data map[string]interface{}, fileBase64 string, ) *FillFormFieldsRequest`

NewFillFormFieldsRequest instantiates a new FillFormFieldsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFillFormFieldsRequestWithDefaults

`func NewFillFormFieldsRequestWithDefaults() *FillFormFieldsRequest`

NewFillFormFieldsRequestWithDefaults instantiates a new FillFormFieldsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *FillFormFieldsRequest) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *FillFormFieldsRequest) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *FillFormFieldsRequest) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetData

`func (o *FillFormFieldsRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FillFormFieldsRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FillFormFieldsRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetOutput

`func (o *FillFormFieldsRequest) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *FillFormFieldsRequest) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *FillFormFieldsRequest) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *FillFormFieldsRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetName

`func (o *FillFormFieldsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FillFormFieldsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FillFormFieldsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FillFormFieldsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFileBase64

`func (o *FillFormFieldsRequest) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *FillFormFieldsRequest) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *FillFormFieldsRequest) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


