# CallbackParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Public callback URL that is used to make a POST request when the document is generated. | [optional] 
**Headers** | Pointer to **map[string]interface{}** | A key-value pairs of header values sent with the POST request. | [optional] 

## Methods

### NewCallbackParam

`func NewCallbackParam() *CallbackParam`

NewCallbackParam instantiates a new CallbackParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallbackParamWithDefaults

`func NewCallbackParamWithDefaults() *CallbackParam`

NewCallbackParamWithDefaults instantiates a new CallbackParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CallbackParam) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CallbackParam) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CallbackParam) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CallbackParam) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHeaders

`func (o *CallbackParam) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CallbackParam) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CallbackParam) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CallbackParam) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


