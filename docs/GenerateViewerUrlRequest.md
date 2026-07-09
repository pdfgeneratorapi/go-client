# GenerateViewerUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Output** | Pointer to **string** | Response format. &#x60;url&#x60; returns a public URL to the stored document; &#x60;viewer&#x60; returns a public URL to the PDF viewer (with encrypted prefill when provided). | [optional] [default to "url"]
**Prefill** | Pointer to [**PrefillParam**](PrefillParam.md) |  | [optional] 

## Methods

### NewGenerateViewerUrlRequest

`func NewGenerateViewerUrlRequest() *GenerateViewerUrlRequest`

NewGenerateViewerUrlRequest instantiates a new GenerateViewerUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateViewerUrlRequestWithDefaults

`func NewGenerateViewerUrlRequestWithDefaults() *GenerateViewerUrlRequest`

NewGenerateViewerUrlRequestWithDefaults instantiates a new GenerateViewerUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutput

`func (o *GenerateViewerUrlRequest) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GenerateViewerUrlRequest) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GenerateViewerUrlRequest) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GenerateViewerUrlRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetPrefill

`func (o *GenerateViewerUrlRequest) GetPrefill() PrefillParam`

GetPrefill returns the Prefill field if non-nil, zero value otherwise.

### GetPrefillOk

`func (o *GenerateViewerUrlRequest) GetPrefillOk() (*PrefillParam, bool)`

GetPrefillOk returns a tuple with the Prefill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefill

`func (o *GenerateViewerUrlRequest) SetPrefill(v PrefillParam)`

SetPrefill sets Prefill field to given value.

### HasPrefill

`func (o *GenerateViewerUrlRequest) HasPrefill() bool`

HasPrefill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


