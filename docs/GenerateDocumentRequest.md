# GenerateDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to [**TemplateParam**](TemplateParam.md) |  | [optional] 
**Format** | Pointer to [**FormatParam**](FormatParam.md) |  | [optional] [default to FORMATPARAM_PDF]
**Output** | Pointer to [**OutputParam**](OutputParam.md) |  | [optional] [default to OUTPUTPARAM_BASE64]
**Name** | Pointer to **string** | Generated document name (optional) | [optional] [default to ""]
**Testing** | Pointer to **bool** | When set to true the generation is not counted as merge (monthly usage), but a large PREVIEW stamp is added. | [optional] [default to false]
**MakeAccessible** | Pointer to **bool** | Enables semantic document tagging. When enabled, a separate Make Accessible action is executed, which consumes additional credits. | [optional] [default to false]
**Metadata** | Pointer to [**MetadataParam**](MetadataParam.md) |  | [optional] 

## Methods

### NewGenerateDocumentRequest

`func NewGenerateDocumentRequest() *GenerateDocumentRequest`

NewGenerateDocumentRequest instantiates a new GenerateDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateDocumentRequestWithDefaults

`func NewGenerateDocumentRequestWithDefaults() *GenerateDocumentRequest`

NewGenerateDocumentRequestWithDefaults instantiates a new GenerateDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *GenerateDocumentRequest) GetTemplate() TemplateParam`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GenerateDocumentRequest) GetTemplateOk() (*TemplateParam, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GenerateDocumentRequest) SetTemplate(v TemplateParam)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GenerateDocumentRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetFormat

`func (o *GenerateDocumentRequest) GetFormat() FormatParam`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GenerateDocumentRequest) GetFormatOk() (*FormatParam, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GenerateDocumentRequest) SetFormat(v FormatParam)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GenerateDocumentRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetOutput

`func (o *GenerateDocumentRequest) GetOutput() OutputParam`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GenerateDocumentRequest) GetOutputOk() (*OutputParam, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GenerateDocumentRequest) SetOutput(v OutputParam)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GenerateDocumentRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetName

`func (o *GenerateDocumentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenerateDocumentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenerateDocumentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenerateDocumentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTesting

`func (o *GenerateDocumentRequest) GetTesting() bool`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *GenerateDocumentRequest) GetTestingOk() (*bool, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *GenerateDocumentRequest) SetTesting(v bool)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *GenerateDocumentRequest) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetMakeAccessible

`func (o *GenerateDocumentRequest) GetMakeAccessible() bool`

GetMakeAccessible returns the MakeAccessible field if non-nil, zero value otherwise.

### GetMakeAccessibleOk

`func (o *GenerateDocumentRequest) GetMakeAccessibleOk() (*bool, bool)`

GetMakeAccessibleOk returns a tuple with the MakeAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeAccessible

`func (o *GenerateDocumentRequest) SetMakeAccessible(v bool)`

SetMakeAccessible sets MakeAccessible field to given value.

### HasMakeAccessible

`func (o *GenerateDocumentRequest) HasMakeAccessible() bool`

HasMakeAccessible returns a boolean if a field has been set.

### GetMetadata

`func (o *GenerateDocumentRequest) GetMetadata() MetadataParam`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GenerateDocumentRequest) GetMetadataOk() (*MetadataParam, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GenerateDocumentRequest) SetMetadata(v MetadataParam)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GenerateDocumentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


