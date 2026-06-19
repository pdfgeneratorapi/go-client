# FormFillUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | Public URL to a PDF document | 
**Data** | **map[string]interface{}** | Form field data to fill in the PDF | 
**Output** | Pointer to **string** | Returned document output format | [optional] [default to "base64"]
**Name** | Pointer to **string** | Name for the PDF file | [optional] 

## Methods

### NewFormFillUrl

`func NewFormFillUrl(fileUrl string, data map[string]interface{}, ) *FormFillUrl`

NewFormFillUrl instantiates a new FormFillUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFillUrlWithDefaults

`func NewFormFillUrlWithDefaults() *FormFillUrl`

NewFormFillUrlWithDefaults instantiates a new FormFillUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *FormFillUrl) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *FormFillUrl) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *FormFillUrl) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetData

`func (o *FormFillUrl) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FormFillUrl) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FormFillUrl) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetOutput

`func (o *FormFillUrl) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *FormFillUrl) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *FormFillUrl) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *FormFillUrl) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetName

`func (o *FormFillUrl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormFillUrl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormFillUrl) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormFillUrl) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


