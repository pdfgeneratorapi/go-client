# GenerateViewerUrl200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to **string** | Public URL to the PDF viewer. | [optional] 
**Meta** | Pointer to [**GenerateViewerUrl200ResponseMeta**](GenerateViewerUrl200ResponseMeta.md) |  | [optional] 

## Methods

### NewGenerateViewerUrl200Response

`func NewGenerateViewerUrl200Response() *GenerateViewerUrl200Response`

NewGenerateViewerUrl200Response instantiates a new GenerateViewerUrl200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateViewerUrl200ResponseWithDefaults

`func NewGenerateViewerUrl200ResponseWithDefaults() *GenerateViewerUrl200Response`

NewGenerateViewerUrl200ResponseWithDefaults instantiates a new GenerateViewerUrl200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *GenerateViewerUrl200Response) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *GenerateViewerUrl200Response) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *GenerateViewerUrl200Response) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *GenerateViewerUrl200Response) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetMeta

`func (o *GenerateViewerUrl200Response) GetMeta() GenerateViewerUrl200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GenerateViewerUrl200Response) GetMetaOk() (*GenerateViewerUrl200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GenerateViewerUrl200Response) SetMeta(v GenerateViewerUrl200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GenerateViewerUrl200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


