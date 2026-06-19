# MakeAccessibleBase64

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileBase64** | **string** | PDF document in base64 encoded string format | 
**Options** | Pointer to [**[]AccessibilityOption**](AccessibilityOption.md) | Accessibility processing options | [optional] 
**Name** | Pointer to **string** | Name for the PDF file | [optional] 
**Output** | Pointer to **string** | Returned document output format | [optional] [default to "base64"]
**Metadata** | Pointer to [**MetadataParam**](MetadataParam.md) |  | [optional] 

## Methods

### NewMakeAccessibleBase64

`func NewMakeAccessibleBase64(fileBase64 string, ) *MakeAccessibleBase64`

NewMakeAccessibleBase64 instantiates a new MakeAccessibleBase64 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMakeAccessibleBase64WithDefaults

`func NewMakeAccessibleBase64WithDefaults() *MakeAccessibleBase64`

NewMakeAccessibleBase64WithDefaults instantiates a new MakeAccessibleBase64 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileBase64

`func (o *MakeAccessibleBase64) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *MakeAccessibleBase64) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *MakeAccessibleBase64) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.


### GetOptions

`func (o *MakeAccessibleBase64) GetOptions() []AccessibilityOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MakeAccessibleBase64) GetOptionsOk() (*[]AccessibilityOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MakeAccessibleBase64) SetOptions(v []AccessibilityOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MakeAccessibleBase64) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetName

`func (o *MakeAccessibleBase64) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MakeAccessibleBase64) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MakeAccessibleBase64) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MakeAccessibleBase64) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *MakeAccessibleBase64) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *MakeAccessibleBase64) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *MakeAccessibleBase64) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *MakeAccessibleBase64) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetMetadata

`func (o *MakeAccessibleBase64) GetMetadata() MetadataParam`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MakeAccessibleBase64) GetMetadataOk() (*MetadataParam, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MakeAccessibleBase64) SetMetadata(v MetadataParam)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MakeAccessibleBase64) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


