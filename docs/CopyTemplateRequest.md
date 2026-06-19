# CopyTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for the copied template. If name is not specified then the original name is used. | [optional] 

## Methods

### NewCopyTemplateRequest

`func NewCopyTemplateRequest() *CopyTemplateRequest`

NewCopyTemplateRequest instantiates a new CopyTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyTemplateRequestWithDefaults

`func NewCopyTemplateRequestWithDefaults() *CopyTemplateRequest`

NewCopyTemplateRequestWithDefaults instantiates a new CopyTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CopyTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CopyTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CopyTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CopyTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


