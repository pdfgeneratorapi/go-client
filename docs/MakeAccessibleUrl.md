# MakeAccessibleUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | Public URL to a PDF document | 
**Options** | Pointer to [**[]AccessibilityOption**](AccessibilityOption.md) | Accessibility processing options | [optional] 
**Name** | Pointer to **string** | Name for the PDF file | [optional] 
**Output** | Pointer to **string** | Returned document output format | [optional] [default to "base64"]
**Metadata** | Pointer to [**MetadataParam**](MetadataParam.md) |  | [optional] 

## Methods

### NewMakeAccessibleUrl

`func NewMakeAccessibleUrl(fileUrl string, ) *MakeAccessibleUrl`

NewMakeAccessibleUrl instantiates a new MakeAccessibleUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMakeAccessibleUrlWithDefaults

`func NewMakeAccessibleUrlWithDefaults() *MakeAccessibleUrl`

NewMakeAccessibleUrlWithDefaults instantiates a new MakeAccessibleUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *MakeAccessibleUrl) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *MakeAccessibleUrl) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *MakeAccessibleUrl) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetOptions

`func (o *MakeAccessibleUrl) GetOptions() []AccessibilityOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MakeAccessibleUrl) GetOptionsOk() (*[]AccessibilityOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MakeAccessibleUrl) SetOptions(v []AccessibilityOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MakeAccessibleUrl) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetName

`func (o *MakeAccessibleUrl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MakeAccessibleUrl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MakeAccessibleUrl) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MakeAccessibleUrl) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *MakeAccessibleUrl) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *MakeAccessibleUrl) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *MakeAccessibleUrl) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *MakeAccessibleUrl) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetMetadata

`func (o *MakeAccessibleUrl) GetMetadata() MetadataParam`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MakeAccessibleUrl) GetMetadataOk() (*MetadataParam, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MakeAccessibleUrl) SetMetadata(v MetadataParam)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MakeAccessibleUrl) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


