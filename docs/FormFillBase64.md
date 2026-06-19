# FormFillBase64

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileBase64** | **string** | PDF document in base64 encoded string format | 
**Data** | **map[string]interface{}** | Form field data to fill in the PDF | 
**Output** | Pointer to **string** | Returned document output format | [optional] [default to "base64"]
**Name** | Pointer to **string** | Name for the PDF file | [optional] 

## Methods

### NewFormFillBase64

`func NewFormFillBase64(fileBase64 string, data map[string]interface{}, ) *FormFillBase64`

NewFormFillBase64 instantiates a new FormFillBase64 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFillBase64WithDefaults

`func NewFormFillBase64WithDefaults() *FormFillBase64`

NewFormFillBase64WithDefaults instantiates a new FormFillBase64 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileBase64

`func (o *FormFillBase64) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *FormFillBase64) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *FormFillBase64) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.


### GetData

`func (o *FormFillBase64) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FormFillBase64) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FormFillBase64) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetOutput

`func (o *FormFillBase64) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *FormFillBase64) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *FormFillBase64) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *FormFillBase64) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetName

`func (o *FormFillBase64) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormFillBase64) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormFillBase64) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormFillBase64) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


