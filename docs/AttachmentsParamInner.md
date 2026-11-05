# AttachmentsParamInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | File name, must end with &#x60;.pdf&#x60; or &#x60;.xml&#x60; | 
**Content** | **string** | Base64 encoded file content | 

## Methods

### NewAttachmentsParamInner

`func NewAttachmentsParamInner(name string, content string, ) *AttachmentsParamInner`

NewAttachmentsParamInner instantiates a new AttachmentsParamInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentsParamInnerWithDefaults

`func NewAttachmentsParamInnerWithDefaults() *AttachmentsParamInner`

NewAttachmentsParamInnerWithDefaults instantiates a new AttachmentsParamInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AttachmentsParamInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttachmentsParamInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttachmentsParamInner) SetName(v string)`

SetName sets Name field to given value.


### GetContent

`func (o *AttachmentsParamInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AttachmentsParamInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AttachmentsParamInner) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


