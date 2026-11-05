# CreateXRechnungEInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string]interface{}** | JSON payload that represents the Peppol BIS Billing 3.0 UBL Invoice (https://docs.peppol.eu/poacc/billing/3.0/) Use the Get schema endpoint to see the detailed payload structure. | 
**Type** | Pointer to **string** | Syntax of the XRechnung invoice. &#x60;CII&#x60; produces XRechnung CII; any other value produces XRechnung UBL. | [optional] [default to "UBL"]
**Output** | Pointer to **string** | Response format. When the \&quot;file\&quot; option is used the API returns the file inline. | [optional] [default to "base64"]

## Methods

### NewCreateXRechnungEInvoiceRequest

`func NewCreateXRechnungEInvoiceRequest(data map[string]interface{}, ) *CreateXRechnungEInvoiceRequest`

NewCreateXRechnungEInvoiceRequest instantiates a new CreateXRechnungEInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateXRechnungEInvoiceRequestWithDefaults

`func NewCreateXRechnungEInvoiceRequestWithDefaults() *CreateXRechnungEInvoiceRequest`

NewCreateXRechnungEInvoiceRequestWithDefaults instantiates a new CreateXRechnungEInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateXRechnungEInvoiceRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateXRechnungEInvoiceRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateXRechnungEInvoiceRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetType

`func (o *CreateXRechnungEInvoiceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateXRechnungEInvoiceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateXRechnungEInvoiceRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateXRechnungEInvoiceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOutput

`func (o *CreateXRechnungEInvoiceRequest) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *CreateXRechnungEInvoiceRequest) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *CreateXRechnungEInvoiceRequest) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *CreateXRechnungEInvoiceRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


