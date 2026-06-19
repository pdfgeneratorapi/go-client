# ConvertURL2PDFRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Public URL | [optional] 
**PaperSize** | Pointer to **string** | PDF page size | [optional] [default to "a4"]
**Orientation** | Pointer to **string** |  | [optional] [default to "portrait"]
**Output** | Pointer to **string** | Output | [optional] [default to "base64"]
**Filename** | Pointer to **string** | Document name | [optional] 

## Methods

### NewConvertURL2PDFRequest

`func NewConvertURL2PDFRequest() *ConvertURL2PDFRequest`

NewConvertURL2PDFRequest instantiates a new ConvertURL2PDFRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertURL2PDFRequestWithDefaults

`func NewConvertURL2PDFRequestWithDefaults() *ConvertURL2PDFRequest`

NewConvertURL2PDFRequestWithDefaults instantiates a new ConvertURL2PDFRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ConvertURL2PDFRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConvertURL2PDFRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConvertURL2PDFRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConvertURL2PDFRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPaperSize

`func (o *ConvertURL2PDFRequest) GetPaperSize() string`

GetPaperSize returns the PaperSize field if non-nil, zero value otherwise.

### GetPaperSizeOk

`func (o *ConvertURL2PDFRequest) GetPaperSizeOk() (*string, bool)`

GetPaperSizeOk returns a tuple with the PaperSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaperSize

`func (o *ConvertURL2PDFRequest) SetPaperSize(v string)`

SetPaperSize sets PaperSize field to given value.

### HasPaperSize

`func (o *ConvertURL2PDFRequest) HasPaperSize() bool`

HasPaperSize returns a boolean if a field has been set.

### GetOrientation

`func (o *ConvertURL2PDFRequest) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *ConvertURL2PDFRequest) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *ConvertURL2PDFRequest) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *ConvertURL2PDFRequest) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetOutput

`func (o *ConvertURL2PDFRequest) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ConvertURL2PDFRequest) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ConvertURL2PDFRequest) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ConvertURL2PDFRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetFilename

`func (o *ConvertURL2PDFRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ConvertURL2PDFRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ConvertURL2PDFRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ConvertURL2PDFRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


