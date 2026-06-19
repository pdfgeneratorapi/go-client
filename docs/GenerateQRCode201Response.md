# GenerateQRCode201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to **string** | Base64 encoded image if the output&#x3D;base64. Raw image string if output&#x3D;file is used | [optional] 
**Meta** | Pointer to [**GenerateQRCode201ResponseMeta**](GenerateQRCode201ResponseMeta.md) |  | [optional] 

## Methods

### NewGenerateQRCode201Response

`func NewGenerateQRCode201Response() *GenerateQRCode201Response`

NewGenerateQRCode201Response instantiates a new GenerateQRCode201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateQRCode201ResponseWithDefaults

`func NewGenerateQRCode201ResponseWithDefaults() *GenerateQRCode201Response`

NewGenerateQRCode201ResponseWithDefaults instantiates a new GenerateQRCode201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *GenerateQRCode201Response) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *GenerateQRCode201Response) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *GenerateQRCode201Response) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *GenerateQRCode201Response) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetMeta

`func (o *GenerateQRCode201Response) GetMeta() GenerateQRCode201ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GenerateQRCode201Response) GetMetaOk() (*GenerateQRCode201ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GenerateQRCode201Response) SetMeta(v GenerateQRCode201ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GenerateQRCode201Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


