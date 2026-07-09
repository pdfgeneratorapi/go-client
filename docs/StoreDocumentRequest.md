# StoreDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileBase64** | Pointer to **string** | Base64 encoded PDF file. Required if file_url is not provided. | [optional] 
**FileUrl** | Pointer to **string** | Public HTTPS URL to a PDF file. Required if file_base64 is not provided. | [optional] 
**Prefill** | Pointer to [**PrefillParam**](PrefillParam.md) |  | [optional] 
**Name** | Pointer to **string** | Generated document name (optional) | [optional] [default to ""]
**Output** | Pointer to **string** | Response format. &#x60;url&#x60; returns a public URL to the stored document; &#x60;viewer&#x60; returns a public URL to the PDF viewer. | [optional] [default to "url"]

## Methods

### NewStoreDocumentRequest

`func NewStoreDocumentRequest() *StoreDocumentRequest`

NewStoreDocumentRequest instantiates a new StoreDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDocumentRequestWithDefaults

`func NewStoreDocumentRequestWithDefaults() *StoreDocumentRequest`

NewStoreDocumentRequestWithDefaults instantiates a new StoreDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileBase64

`func (o *StoreDocumentRequest) GetFileBase64() string`

GetFileBase64 returns the FileBase64 field if non-nil, zero value otherwise.

### GetFileBase64Ok

`func (o *StoreDocumentRequest) GetFileBase64Ok() (*string, bool)`

GetFileBase64Ok returns a tuple with the FileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileBase64

`func (o *StoreDocumentRequest) SetFileBase64(v string)`

SetFileBase64 sets FileBase64 field to given value.

### HasFileBase64

`func (o *StoreDocumentRequest) HasFileBase64() bool`

HasFileBase64 returns a boolean if a field has been set.

### GetFileUrl

`func (o *StoreDocumentRequest) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *StoreDocumentRequest) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *StoreDocumentRequest) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *StoreDocumentRequest) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetPrefill

`func (o *StoreDocumentRequest) GetPrefill() PrefillParam`

GetPrefill returns the Prefill field if non-nil, zero value otherwise.

### GetPrefillOk

`func (o *StoreDocumentRequest) GetPrefillOk() (*PrefillParam, bool)`

GetPrefillOk returns a tuple with the Prefill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefill

`func (o *StoreDocumentRequest) SetPrefill(v PrefillParam)`

SetPrefill sets Prefill field to given value.

### HasPrefill

`func (o *StoreDocumentRequest) HasPrefill() bool`

HasPrefill returns a boolean if a field has been set.

### GetName

`func (o *StoreDocumentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoreDocumentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoreDocumentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoreDocumentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutput

`func (o *StoreDocumentRequest) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *StoreDocumentRequest) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *StoreDocumentRequest) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *StoreDocumentRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


