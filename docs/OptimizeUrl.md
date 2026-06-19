# OptimizeUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | Public URL to a PDF document | 
**Name** | Pointer to **string** | Name for the PDF file | [optional] 
**Output** | Pointer to **string** | Returned document output format | [optional] [default to "base64"]

## Methods

### NewOptimizeUrl

`func NewOptimizeUrl(fileUrl string, ) *OptimizeUrl`

NewOptimizeUrl instantiates a new OptimizeUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizeUrlWithDefaults

`func NewOptimizeUrlWithDefaults() *OptimizeUrl`

NewOptimizeUrlWithDefaults instantiates a new OptimizeUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *OptimizeUrl) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *OptimizeUrl) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *OptimizeUrl) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetName

`func (o *OptimizeUrl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptimizeUrl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptimizeUrl) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptimizeUrl) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *OptimizeUrl) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *OptimizeUrl) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *OptimizeUrl) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *OptimizeUrl) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


