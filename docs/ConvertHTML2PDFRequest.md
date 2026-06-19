# ConvertHTML2PDFRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | HTML content | [optional] 
**PaperSize** | Pointer to **string** | PDF page size | [optional] [default to "a4"]
**Orientation** | Pointer to **string** |  | [optional] [default to "portrait"]
**Output** | Pointer to **string** | Output | [optional] [default to "base64"]
**Filename** | Pointer to **string** | Document name | [optional] 

## Methods

### NewConvertHTML2PDFRequest

`func NewConvertHTML2PDFRequest() *ConvertHTML2PDFRequest`

NewConvertHTML2PDFRequest instantiates a new ConvertHTML2PDFRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertHTML2PDFRequestWithDefaults

`func NewConvertHTML2PDFRequestWithDefaults() *ConvertHTML2PDFRequest`

NewConvertHTML2PDFRequestWithDefaults instantiates a new ConvertHTML2PDFRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ConvertHTML2PDFRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ConvertHTML2PDFRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ConvertHTML2PDFRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ConvertHTML2PDFRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPaperSize

`func (o *ConvertHTML2PDFRequest) GetPaperSize() string`

GetPaperSize returns the PaperSize field if non-nil, zero value otherwise.

### GetPaperSizeOk

`func (o *ConvertHTML2PDFRequest) GetPaperSizeOk() (*string, bool)`

GetPaperSizeOk returns a tuple with the PaperSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaperSize

`func (o *ConvertHTML2PDFRequest) SetPaperSize(v string)`

SetPaperSize sets PaperSize field to given value.

### HasPaperSize

`func (o *ConvertHTML2PDFRequest) HasPaperSize() bool`

HasPaperSize returns a boolean if a field has been set.

### GetOrientation

`func (o *ConvertHTML2PDFRequest) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *ConvertHTML2PDFRequest) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *ConvertHTML2PDFRequest) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *ConvertHTML2PDFRequest) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetOutput

`func (o *ConvertHTML2PDFRequest) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ConvertHTML2PDFRequest) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ConvertHTML2PDFRequest) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ConvertHTML2PDFRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetFilename

`func (o *ConvertHTML2PDFRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ConvertHTML2PDFRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ConvertHTML2PDFRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ConvertHTML2PDFRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


