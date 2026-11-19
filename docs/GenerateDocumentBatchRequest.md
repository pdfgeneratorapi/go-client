# GenerateDocumentBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | [**[]TemplateParam**](TemplateParam.md) |  | 
**Format** | Pointer to [**FormatParam**](FormatParam.md) |  | [optional] [default to FORMATPARAM_PDF]
**Output** | Pointer to [**OutputParam**](OutputParam.md) |  | [optional] [default to OUTPUTPARAM_BASE64]
**Name** | Pointer to **string** | Generated document name (optional) | [optional] [default to ""]
**Testing** | Pointer to **bool** | When set to true the generation is not counted as merge (monthly usage), but a large PREVIEW stamp is added. | [optional] [default to false]
**Print** | Pointer to **bool** | When set to true, the PDF opens the print dialog when it is opened in a viewer. | [optional] [default to false]
**EmptyLabels** | Pointer to **int32** | Number of label positions to leave empty at the start of the first sheet. Only applies to label templates. | [optional] [default to 0]
**Attachments** | Pointer to [**[]AttachmentsParamInner**](AttachmentsParamInner.md) | Files embedded into the generated PDF as attachments. Only PDF and XML files are accepted. | [optional] 
**Metadata** | Pointer to [**MetadataParam**](MetadataParam.md) |  | [optional] 
**DisplayMode** | Pointer to [**DisplayModeParam**](DisplayModeParam.md) |  | [optional] 

## Methods

### NewGenerateDocumentBatchRequest

`func NewGenerateDocumentBatchRequest(template []TemplateParam, ) *GenerateDocumentBatchRequest`

NewGenerateDocumentBatchRequest instantiates a new GenerateDocumentBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateDocumentBatchRequestWithDefaults

`func NewGenerateDocumentBatchRequestWithDefaults() *GenerateDocumentBatchRequest`

NewGenerateDocumentBatchRequestWithDefaults instantiates a new GenerateDocumentBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *GenerateDocumentBatchRequest) GetTemplate() []TemplateParam`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GenerateDocumentBatchRequest) GetTemplateOk() (*[]TemplateParam, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GenerateDocumentBatchRequest) SetTemplate(v []TemplateParam)`

SetTemplate sets Template field to given value.


### GetFormat

`func (o *GenerateDocumentBatchRequest) GetFormat() FormatParam`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GenerateDocumentBatchRequest) GetFormatOk() (*FormatParam, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GenerateDocumentBatchRequest) SetFormat(v FormatParam)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GenerateDocumentBatchRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetOutput

`func (o *GenerateDocumentBatchRequest) GetOutput() OutputParam`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GenerateDocumentBatchRequest) GetOutputOk() (*OutputParam, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GenerateDocumentBatchRequest) SetOutput(v OutputParam)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GenerateDocumentBatchRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetName

`func (o *GenerateDocumentBatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenerateDocumentBatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenerateDocumentBatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenerateDocumentBatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTesting

`func (o *GenerateDocumentBatchRequest) GetTesting() bool`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *GenerateDocumentBatchRequest) GetTestingOk() (*bool, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *GenerateDocumentBatchRequest) SetTesting(v bool)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *GenerateDocumentBatchRequest) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetPrint

`func (o *GenerateDocumentBatchRequest) GetPrint() bool`

GetPrint returns the Print field if non-nil, zero value otherwise.

### GetPrintOk

`func (o *GenerateDocumentBatchRequest) GetPrintOk() (*bool, bool)`

GetPrintOk returns a tuple with the Print field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrint

`func (o *GenerateDocumentBatchRequest) SetPrint(v bool)`

SetPrint sets Print field to given value.

### HasPrint

`func (o *GenerateDocumentBatchRequest) HasPrint() bool`

HasPrint returns a boolean if a field has been set.

### GetEmptyLabels

`func (o *GenerateDocumentBatchRequest) GetEmptyLabels() int32`

GetEmptyLabels returns the EmptyLabels field if non-nil, zero value otherwise.

### GetEmptyLabelsOk

`func (o *GenerateDocumentBatchRequest) GetEmptyLabelsOk() (*int32, bool)`

GetEmptyLabelsOk returns a tuple with the EmptyLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyLabels

`func (o *GenerateDocumentBatchRequest) SetEmptyLabels(v int32)`

SetEmptyLabels sets EmptyLabels field to given value.

### HasEmptyLabels

`func (o *GenerateDocumentBatchRequest) HasEmptyLabels() bool`

HasEmptyLabels returns a boolean if a field has been set.

### GetAttachments

`func (o *GenerateDocumentBatchRequest) GetAttachments() []AttachmentsParamInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GenerateDocumentBatchRequest) GetAttachmentsOk() (*[]AttachmentsParamInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GenerateDocumentBatchRequest) SetAttachments(v []AttachmentsParamInner)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GenerateDocumentBatchRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetMetadata

`func (o *GenerateDocumentBatchRequest) GetMetadata() MetadataParam`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GenerateDocumentBatchRequest) GetMetadataOk() (*MetadataParam, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GenerateDocumentBatchRequest) SetMetadata(v MetadataParam)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GenerateDocumentBatchRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDisplayMode

`func (o *GenerateDocumentBatchRequest) GetDisplayMode() DisplayModeParam`

GetDisplayMode returns the DisplayMode field if non-nil, zero value otherwise.

### GetDisplayModeOk

`func (o *GenerateDocumentBatchRequest) GetDisplayModeOk() (*DisplayModeParam, bool)`

GetDisplayModeOk returns a tuple with the DisplayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMode

`func (o *GenerateDocumentBatchRequest) SetDisplayMode(v DisplayModeParam)`

SetDisplayMode sets DisplayMode field to given value.

### HasDisplayMode

`func (o *GenerateDocumentBatchRequest) HasDisplayMode() bool`

HasDisplayMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


