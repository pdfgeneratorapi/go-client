# CreateFacturXEInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | [**CreateFacturXEInvoiceRequestTemplate**](CreateFacturXEInvoiceRequestTemplate.md) |  | 
**Profile** | Pointer to **string** | Factur-X conformance level. | [optional] [default to "basic"]
**Output** | Pointer to [**OutputParam**](OutputParam.md) |  | [optional] [default to OUTPUTPARAM_BASE64]
**Name** | Pointer to **string** | Generated document name (optional) | [optional] [default to ""]
**Metadata** | Pointer to [**MetadataParam**](MetadataParam.md) |  | [optional] 

## Methods

### NewCreateFacturXEInvoiceRequest

`func NewCreateFacturXEInvoiceRequest(template CreateFacturXEInvoiceRequestTemplate, ) *CreateFacturXEInvoiceRequest`

NewCreateFacturXEInvoiceRequest instantiates a new CreateFacturXEInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFacturXEInvoiceRequestWithDefaults

`func NewCreateFacturXEInvoiceRequestWithDefaults() *CreateFacturXEInvoiceRequest`

NewCreateFacturXEInvoiceRequestWithDefaults instantiates a new CreateFacturXEInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *CreateFacturXEInvoiceRequest) GetTemplate() CreateFacturXEInvoiceRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CreateFacturXEInvoiceRequest) GetTemplateOk() (*CreateFacturXEInvoiceRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CreateFacturXEInvoiceRequest) SetTemplate(v CreateFacturXEInvoiceRequestTemplate)`

SetTemplate sets Template field to given value.


### GetProfile

`func (o *CreateFacturXEInvoiceRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CreateFacturXEInvoiceRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CreateFacturXEInvoiceRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CreateFacturXEInvoiceRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetOutput

`func (o *CreateFacturXEInvoiceRequest) GetOutput() OutputParam`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *CreateFacturXEInvoiceRequest) GetOutputOk() (*OutputParam, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *CreateFacturXEInvoiceRequest) SetOutput(v OutputParam)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *CreateFacturXEInvoiceRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetName

`func (o *CreateFacturXEInvoiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFacturXEInvoiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFacturXEInvoiceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateFacturXEInvoiceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateFacturXEInvoiceRequest) GetMetadata() MetadataParam`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateFacturXEInvoiceRequest) GetMetadataOk() (*MetadataParam, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateFacturXEInvoiceRequest) SetMetadata(v MetadataParam)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateFacturXEInvoiceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


