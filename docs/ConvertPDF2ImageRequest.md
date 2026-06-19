# ConvertPDF2ImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileBase64** | Pointer to **string** | Base64 encoded PDF file. Required if file_url is not provided. | [optional] 
**FileUrl** | Pointer to **string** | Public HTTPS URL to a PDF file. Required if file_base64 is not provided. | [optional] 
**Format** | Pointer to **string** | Output image format | [optional] [default to "png"]
**Quality** | Pointer to **int32** | Image quality (1-100) | [optional] [default to 85]
**Resolution** | Pointer to **int32** | Image resolution in DPI (72-600) | [optional] [default to 150]
**Pages** | Pointer to **string** | Page number or range to convert. Use a single number (e.g. \&quot;1\&quot;) or a range (e.g. \&quot;1-4\&quot;). Defaults to all pages. | [optional] [default to "all"]
**Output** | Pointer to **string** | Output format | [optional] [default to "base64"]
**Name** | Pointer to **string** | Document name (max 120 characters). Auto-generated if not provided. | [optional] 

## Methods

### NewConvertPDF2ImageRequest

`func NewConvertPDF2ImageRequest() *ConvertPDF2ImageRequest`

NewConvertPDF2ImageRequest instantiates a new ConvertPDF2ImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertPDF2ImageRequestWithDefaults

`func NewConvertPDF2ImageRequestWithDefaults() *ConvertPDF2ImageRequest`

NewConvertPDF2ImageRequestWithDefaults instantiates a new ConvertPDF2ImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileBase64

`func (o *ConvertPDF2ImageRequest) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *ConvertPDF2ImageRequest) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *ConvertPDF2ImageRequest) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.

### HasFileBase64

`func (o *ConvertPDF2ImageRequest) HasFileBase64() bool`

HasFileBase64 returns a boolean if a field has been set.

### GetFileUrl

`func (o *ConvertPDF2ImageRequest) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ConvertPDF2ImageRequest) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ConvertPDF2ImageRequest) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *ConvertPDF2ImageRequest) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetFormat

`func (o *ConvertPDF2ImageRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ConvertPDF2ImageRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ConvertPDF2ImageRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ConvertPDF2ImageRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetQuality

`func (o *ConvertPDF2ImageRequest) GetQuality() int32`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ConvertPDF2ImageRequest) GetQualityOk() (*int32, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ConvertPDF2ImageRequest) SetQuality(v int32)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ConvertPDF2ImageRequest) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetResolution

`func (o *ConvertPDF2ImageRequest) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ConvertPDF2ImageRequest) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ConvertPDF2ImageRequest) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *ConvertPDF2ImageRequest) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetPages

`func (o *ConvertPDF2ImageRequest) GetPages() string`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ConvertPDF2ImageRequest) GetPagesOk() (*string, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ConvertPDF2ImageRequest) SetPages(v string)`

SetPages sets Pages field to given value.

### HasPages

`func (o *ConvertPDF2ImageRequest) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetOutput

`func (o *ConvertPDF2ImageRequest) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ConvertPDF2ImageRequest) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ConvertPDF2ImageRequest) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ConvertPDF2ImageRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetName

`func (o *ConvertPDF2ImageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConvertPDF2ImageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConvertPDF2ImageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConvertPDF2ImageRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


