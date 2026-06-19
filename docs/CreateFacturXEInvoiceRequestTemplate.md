# CreateFacturXEInvoiceRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Template ID | [optional] 
**VersionId** | Pointer to **int32** | Template version ID | [optional] 
**Data** | Pointer to **map[string]interface{}** | JSON payload that represents the Peppol BIS Billing 3.0 UBL Invoice (https://docs.peppol.eu/poacc/billing/3.0/) Use the Get schema endpoint to see the detailed payload structure. | [optional] 

## Methods

### NewCreateFacturXEInvoiceRequestTemplate

`func NewCreateFacturXEInvoiceRequestTemplate() *CreateFacturXEInvoiceRequestTemplate`

NewCreateFacturXEInvoiceRequestTemplate instantiates a new CreateFacturXEInvoiceRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFacturXEInvoiceRequestTemplateWithDefaults

`func NewCreateFacturXEInvoiceRequestTemplateWithDefaults() *CreateFacturXEInvoiceRequestTemplate`

NewCreateFacturXEInvoiceRequestTemplateWithDefaults instantiates a new CreateFacturXEInvoiceRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateFacturXEInvoiceRequestTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateFacturXEInvoiceRequestTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateFacturXEInvoiceRequestTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateFacturXEInvoiceRequestTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionId

`func (o *CreateFacturXEInvoiceRequestTemplate) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *CreateFacturXEInvoiceRequestTemplate) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *CreateFacturXEInvoiceRequestTemplate) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *CreateFacturXEInvoiceRequestTemplate) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetData

`func (o *CreateFacturXEInvoiceRequestTemplate) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateFacturXEInvoiceRequestTemplate) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateFacturXEInvoiceRequestTemplate) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CreateFacturXEInvoiceRequestTemplate) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


