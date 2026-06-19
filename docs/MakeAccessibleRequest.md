# MakeAccessibleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | Public URL to a PDF document | 
**Options** | Pointer to [**[]AccessibilityOption**](AccessibilityOption.md) | Accessibility processing options | [optional] 
**Name** | Pointer to **string** | Name for the PDF file | [optional] 
**Output** | Pointer to **string** | Returned document output format | [optional] [default to "base64"]
**Metadata** | Pointer to [**MetadataParam**](MetadataParam.md) |  | [optional] 
**FileBase64** | **string** | PDF document in base64 encoded string format | 

## Methods

### NewMakeAccessibleRequest

`func NewMakeAccessibleRequest(fileUrl string, fileBase64 string, ) *MakeAccessibleRequest`

NewMakeAccessibleRequest instantiates a new MakeAccessibleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMakeAccessibleRequestWithDefaults

`func NewMakeAccessibleRequestWithDefaults() *MakeAccessibleRequest`

NewMakeAccessibleRequestWithDefaults instantiates a new MakeAccessibleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *MakeAccessibleRequest) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *MakeAccessibleRequest) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *MakeAccessibleRequest) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetOptions

`func (o *MakeAccessibleRequest) GetOptions() []AccessibilityOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MakeAccessibleRequest) GetOptionsOk() (*[]AccessibilityOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MakeAccessibleRequest) SetOptions(v []AccessibilityOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MakeAccessibleRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetName

`func (o *MakeAccessibleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MakeAccessibleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MakeAccessibleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MakeAccessibleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *MakeAccessibleRequest) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *MakeAccessibleRequest) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *MakeAccessibleRequest) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *MakeAccessibleRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetMetadata

`func (o *MakeAccessibleRequest) GetMetadata() MetadataParam`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MakeAccessibleRequest) GetMetadataOk() (*MetadataParam, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MakeAccessibleRequest) SetMetadata(v MetadataParam)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MakeAccessibleRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFileBase64

`func (o *MakeAccessibleRequest) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *MakeAccessibleRequest) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *MakeAccessibleRequest) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


