# GenerateDocumentBatchAsynchronousRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | [**[]TemplateParam**](TemplateParam.md) |  | 
**Callback** | [**CallbackParam**](CallbackParam.md) |  | 
**Format** | Pointer to [**FormatParam**](FormatParam.md) |  | [optional] [default to FORMATPARAM_PDF]
**Output** | Pointer to [**AsyncOutputParam**](AsyncOutputParam.md) |  | [optional] [default to ASYNCOUTPUTPARAM_BASE64]
**Name** | Pointer to **string** | Generated document name (optional) | [optional] [default to ""]
**Testing** | Pointer to **bool** | When set to true the generation is not counted as merge (monthly usage), but a large PREVIEW stamp is added. | [optional] [default to false]
**Print** | Pointer to **bool** | When set to true, the PDF opens the print dialog when it is opened in a viewer. | [optional] [default to false]
**EmptyLabels** | Pointer to **int32** | Number of label positions to leave empty at the start of the first sheet. Only applies to label templates. | [optional] [default to 0]
**Attachments** | Pointer to [**[]AttachmentsParamInner**](AttachmentsParamInner.md) | Files embedded into the generated PDF as attachments. Only PDF and XML files are accepted. | [optional] 
**Metadata** | Pointer to [**MetadataParam**](MetadataParam.md) |  | [optional] 

## Methods

### NewGenerateDocumentBatchAsynchronousRequest

`func NewGenerateDocumentBatchAsynchronousRequest(template []TemplateParam, callback CallbackParam, ) *GenerateDocumentBatchAsynchronousRequest`

NewGenerateDocumentBatchAsynchronousRequest instantiates a new GenerateDocumentBatchAsynchronousRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateDocumentBatchAsynchronousRequestWithDefaults

`func NewGenerateDocumentBatchAsynchronousRequestWithDefaults() *GenerateDocumentBatchAsynchronousRequest`

NewGenerateDocumentBatchAsynchronousRequestWithDefaults instantiates a new GenerateDocumentBatchAsynchronousRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *GenerateDocumentBatchAsynchronousRequest) GetTemplate() []TemplateParam`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GenerateDocumentBatchAsynchronousRequest) GetTemplateOk() (*[]TemplateParam, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GenerateDocumentBatchAsynchronousRequest) SetTemplate(v []TemplateParam)`

SetTemplate sets Template field to given value.


### GetCallback

`func (o *GenerateDocumentBatchAsynchronousRequest) GetCallback() CallbackParam`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *GenerateDocumentBatchAsynchronousRequest) GetCallbackOk() (*CallbackParam, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *GenerateDocumentBatchAsynchronousRequest) SetCallback(v CallbackParam)`

SetCallback sets Callback field to given value.


### GetFormat

`func (o *GenerateDocumentBatchAsynchronousRequest) GetFormat() FormatParam`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GenerateDocumentBatchAsynchronousRequest) GetFormatOk() (*FormatParam, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GenerateDocumentBatchAsynchronousRequest) SetFormat(v FormatParam)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GenerateDocumentBatchAsynchronousRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetOutput

`func (o *GenerateDocumentBatchAsynchronousRequest) GetOutput() AsyncOutputParam`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GenerateDocumentBatchAsynchronousRequest) GetOutputOk() (*AsyncOutputParam, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GenerateDocumentBatchAsynchronousRequest) SetOutput(v AsyncOutputParam)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GenerateDocumentBatchAsynchronousRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetName

`func (o *GenerateDocumentBatchAsynchronousRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenerateDocumentBatchAsynchronousRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenerateDocumentBatchAsynchronousRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenerateDocumentBatchAsynchronousRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTesting

`func (o *GenerateDocumentBatchAsynchronousRequest) GetTesting() bool`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *GenerateDocumentBatchAsynchronousRequest) GetTestingOk() (*bool, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *GenerateDocumentBatchAsynchronousRequest) SetTesting(v bool)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *GenerateDocumentBatchAsynchronousRequest) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetPrint

`func (o *GenerateDocumentBatchAsynchronousRequest) GetPrint() bool`

GetPrint returns the Print field if non-nil, zero value otherwise.

### GetPrintOk

`func (o *GenerateDocumentBatchAsynchronousRequest) GetPrintOk() (*bool, bool)`

GetPrintOk returns a tuple with the Print field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrint

`func (o *GenerateDocumentBatchAsynchronousRequest) SetPrint(v bool)`

SetPrint sets Print field to given value.

### HasPrint

`func (o *GenerateDocumentBatchAsynchronousRequest) HasPrint() bool`

HasPrint returns a boolean if a field has been set.

### GetEmptyLabels

`func (o *GenerateDocumentBatchAsynchronousRequest) GetEmptyLabels() int32`

GetEmptyLabels returns the EmptyLabels field if non-nil, zero value otherwise.

### GetEmptyLabelsOk

`func (o *GenerateDocumentBatchAsynchronousRequest) GetEmptyLabelsOk() (*int32, bool)`

GetEmptyLabelsOk returns a tuple with the EmptyLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyLabels

`func (o *GenerateDocumentBatchAsynchronousRequest) SetEmptyLabels(v int32)`

SetEmptyLabels sets EmptyLabels field to given value.

### HasEmptyLabels

`func (o *GenerateDocumentBatchAsynchronousRequest) HasEmptyLabels() bool`

HasEmptyLabels returns a boolean if a field has been set.

### GetAttachments

`func (o *GenerateDocumentBatchAsynchronousRequest) GetAttachments() []AttachmentsParamInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GenerateDocumentBatchAsynchronousRequest) GetAttachmentsOk() (*[]AttachmentsParamInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GenerateDocumentBatchAsynchronousRequest) SetAttachments(v []AttachmentsParamInner)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GenerateDocumentBatchAsynchronousRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetMetadata

`func (o *GenerateDocumentBatchAsynchronousRequest) GetMetadata() MetadataParam`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GenerateDocumentBatchAsynchronousRequest) GetMetadataOk() (*MetadataParam, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GenerateDocumentBatchAsynchronousRequest) SetMetadata(v MetadataParam)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GenerateDocumentBatchAsynchronousRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


