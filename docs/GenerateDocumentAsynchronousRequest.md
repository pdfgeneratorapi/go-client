# GenerateDocumentAsynchronousRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to [**TemplateParam**](TemplateParam.md) |  | [optional] 
**Callback** | Pointer to [**CallbackParam**](CallbackParam.md) |  | [optional] 
**Format** | Pointer to [**FormatParam**](FormatParam.md) |  | [optional] [default to FORMATPARAM_PDF]
**Output** | Pointer to [**AsyncOutputParam**](AsyncOutputParam.md) |  | [optional] [default to ASYNCOUTPUTPARAM_BASE64]
**Name** | Pointer to **string** | Generated document name (optional) | [optional] [default to ""]
**Testing** | Pointer to **bool** | When set to true the generation is not counted as merge (monthly usage), but a large PREVIEW stamp is added. | [optional] [default to false]
**MakeAccessible** | Pointer to **bool** | Enables semantic document tagging. When enabled, a separate Make Accessible action is executed, which consumes additional credits. | [optional] [default to false]
**Metadata** | Pointer to [**MetadataParam**](MetadataParam.md) |  | [optional] 

## Methods

### NewGenerateDocumentAsynchronousRequest

`func NewGenerateDocumentAsynchronousRequest() *GenerateDocumentAsynchronousRequest`

NewGenerateDocumentAsynchronousRequest instantiates a new GenerateDocumentAsynchronousRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateDocumentAsynchronousRequestWithDefaults

`func NewGenerateDocumentAsynchronousRequestWithDefaults() *GenerateDocumentAsynchronousRequest`

NewGenerateDocumentAsynchronousRequestWithDefaults instantiates a new GenerateDocumentAsynchronousRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *GenerateDocumentAsynchronousRequest) GetTemplate() TemplateParam`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GenerateDocumentAsynchronousRequest) GetTemplateOk() (*TemplateParam, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GenerateDocumentAsynchronousRequest) SetTemplate(v TemplateParam)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GenerateDocumentAsynchronousRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetCallback

`func (o *GenerateDocumentAsynchronousRequest) GetCallback() CallbackParam`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *GenerateDocumentAsynchronousRequest) GetCallbackOk() (*CallbackParam, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *GenerateDocumentAsynchronousRequest) SetCallback(v CallbackParam)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *GenerateDocumentAsynchronousRequest) HasCallback() bool`

HasCallback returns a boolean if a field has been set.

### GetFormat

`func (o *GenerateDocumentAsynchronousRequest) GetFormat() FormatParam`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GenerateDocumentAsynchronousRequest) GetFormatOk() (*FormatParam, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GenerateDocumentAsynchronousRequest) SetFormat(v FormatParam)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GenerateDocumentAsynchronousRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetOutput

`func (o *GenerateDocumentAsynchronousRequest) GetOutput() AsyncOutputParam`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GenerateDocumentAsynchronousRequest) GetOutputOk() (*AsyncOutputParam, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GenerateDocumentAsynchronousRequest) SetOutput(v AsyncOutputParam)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GenerateDocumentAsynchronousRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetName

`func (o *GenerateDocumentAsynchronousRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenerateDocumentAsynchronousRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenerateDocumentAsynchronousRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenerateDocumentAsynchronousRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTesting

`func (o *GenerateDocumentAsynchronousRequest) GetTesting() bool`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *GenerateDocumentAsynchronousRequest) GetTestingOk() (*bool, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *GenerateDocumentAsynchronousRequest) SetTesting(v bool)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *GenerateDocumentAsynchronousRequest) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetMakeAccessible

`func (o *GenerateDocumentAsynchronousRequest) GetMakeAccessible() bool`

GetMakeAccessible returns the MakeAccessible field if non-nil, zero value otherwise.

### GetMakeAccessibleOk

`func (o *GenerateDocumentAsynchronousRequest) GetMakeAccessibleOk() (*bool, bool)`

GetMakeAccessibleOk returns a tuple with the MakeAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakeAccessible

`func (o *GenerateDocumentAsynchronousRequest) SetMakeAccessible(v bool)`

SetMakeAccessible sets MakeAccessible field to given value.

### HasMakeAccessible

`func (o *GenerateDocumentAsynchronousRequest) HasMakeAccessible() bool`

HasMakeAccessible returns a boolean if a field has been set.

### GetMetadata

`func (o *GenerateDocumentAsynchronousRequest) GetMetadata() MetadataParam`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GenerateDocumentAsynchronousRequest) GetMetadataOk() (*MetadataParam, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GenerateDocumentAsynchronousRequest) SetMetadata(v MetadataParam)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GenerateDocumentAsynchronousRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


