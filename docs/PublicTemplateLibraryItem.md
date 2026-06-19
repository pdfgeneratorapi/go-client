# PublicTemplateLibraryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique public identifier of the template | [optional] 
**Name** | Pointer to **string** | Template name | [optional] 
**Tags** | Pointer to **[]string** | A list of tags assigned to a public template | [optional] 
**PreviewImageUrl** | Pointer to **NullableString** | URL to the template preview image | [optional] 
**SampleDataUrl** | Pointer to **NullableString** | URL to the template sample data | [optional] 

## Methods

### NewPublicTemplateLibraryItem

`func NewPublicTemplateLibraryItem() *PublicTemplateLibraryItem`

NewPublicTemplateLibraryItem instantiates a new PublicTemplateLibraryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTemplateLibraryItemWithDefaults

`func NewPublicTemplateLibraryItemWithDefaults() *PublicTemplateLibraryItem`

NewPublicTemplateLibraryItemWithDefaults instantiates a new PublicTemplateLibraryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicTemplateLibraryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicTemplateLibraryItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicTemplateLibraryItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicTemplateLibraryItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PublicTemplateLibraryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicTemplateLibraryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicTemplateLibraryItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicTemplateLibraryItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *PublicTemplateLibraryItem) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PublicTemplateLibraryItem) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PublicTemplateLibraryItem) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PublicTemplateLibraryItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPreviewImageUrl

`func (o *PublicTemplateLibraryItem) GetPreviewImageUrl() string`

GetPreviewImageUrl returns the PreviewImageUrl field if non-nil, zero value otherwise.

### GetPreviewImageUrlOk

`func (o *PublicTemplateLibraryItem) GetPreviewImageUrlOk() (*string, bool)`

GetPreviewImageUrlOk returns a tuple with the PreviewImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImageUrl

`func (o *PublicTemplateLibraryItem) SetPreviewImageUrl(v string)`

SetPreviewImageUrl sets PreviewImageUrl field to given value.

### HasPreviewImageUrl

`func (o *PublicTemplateLibraryItem) HasPreviewImageUrl() bool`

HasPreviewImageUrl returns a boolean if a field has been set.

### SetPreviewImageUrlNil

`func (o *PublicTemplateLibraryItem) SetPreviewImageUrlNil(b bool)`

 SetPreviewImageUrlNil sets the value for PreviewImageUrl to be an explicit nil

### UnsetPreviewImageUrl
`func (o *PublicTemplateLibraryItem) UnsetPreviewImageUrl()`

UnsetPreviewImageUrl ensures that no value is present for PreviewImageUrl, not even an explicit nil
### GetSampleDataUrl

`func (o *PublicTemplateLibraryItem) GetSampleDataUrl() string`

GetSampleDataUrl returns the SampleDataUrl field if non-nil, zero value otherwise.

### GetSampleDataUrlOk

`func (o *PublicTemplateLibraryItem) GetSampleDataUrlOk() (*string, bool)`

GetSampleDataUrlOk returns a tuple with the SampleDataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleDataUrl

`func (o *PublicTemplateLibraryItem) SetSampleDataUrl(v string)`

SetSampleDataUrl sets SampleDataUrl field to given value.

### HasSampleDataUrl

`func (o *PublicTemplateLibraryItem) HasSampleDataUrl() bool`

HasSampleDataUrl returns a boolean if a field has been set.

### SetSampleDataUrlNil

`func (o *PublicTemplateLibraryItem) SetSampleDataUrlNil(b bool)`

 SetSampleDataUrlNil sets the value for SampleDataUrl to be an explicit nil

### UnsetSampleDataUrl
`func (o *PublicTemplateLibraryItem) UnsetSampleDataUrl()`

UnsetSampleDataUrl ensures that no value is present for SampleDataUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


