# OptimizeBase64

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileBase64** | **string** | PDF document in base64 encoded string format | 
**Name** | Pointer to **string** | Name for the PDF file | [optional] 
**Output** | Pointer to **string** | Returned document output format | [optional] [default to "base64"]

## Methods

### NewOptimizeBase64

`func NewOptimizeBase64(fileBase64 string, ) *OptimizeBase64`

NewOptimizeBase64 instantiates a new OptimizeBase64 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizeBase64WithDefaults

`func NewOptimizeBase64WithDefaults() *OptimizeBase64`

NewOptimizeBase64WithDefaults instantiates a new OptimizeBase64 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileBase64

`func (o *OptimizeBase64) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *OptimizeBase64) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *OptimizeBase64) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.


### GetName

`func (o *OptimizeBase64) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptimizeBase64) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptimizeBase64) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptimizeBase64) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *OptimizeBase64) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *OptimizeBase64) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *OptimizeBase64) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *OptimizeBase64) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


