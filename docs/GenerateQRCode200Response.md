# GenerateQRCode200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to **string** | Base64 encoded PNG image | [optional] 
**Meta** | Pointer to [**GenerateQRCode200ResponseMeta**](GenerateQRCode200ResponseMeta.md) |  | [optional] 

## Methods

### NewGenerateQRCode200Response

`func NewGenerateQRCode200Response() *GenerateQRCode200Response`

NewGenerateQRCode200Response instantiates a new GenerateQRCode200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateQRCode200ResponseWithDefaults

`func NewGenerateQRCode200ResponseWithDefaults() *GenerateQRCode200Response`

NewGenerateQRCode200ResponseWithDefaults instantiates a new GenerateQRCode200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *GenerateQRCode200Response) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *GenerateQRCode200Response) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *GenerateQRCode200Response) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *GenerateQRCode200Response) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetMeta

`func (o *GenerateQRCode200Response) GetMeta() GenerateQRCode200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GenerateQRCode200Response) GetMetaOk() (*GenerateQRCode200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GenerateQRCode200Response) SetMeta(v GenerateQRCode200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GenerateQRCode200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


