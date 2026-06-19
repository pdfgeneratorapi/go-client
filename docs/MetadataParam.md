# MetadataParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | Document author | [optional] [default to "Organization name"]
**Language** | Pointer to **string** | Document language | [optional] [default to "en"]

## Methods

### NewMetadataParam

`func NewMetadataParam() *MetadataParam`

NewMetadataParam instantiates a new MetadataParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataParamWithDefaults

`func NewMetadataParamWithDefaults() *MetadataParam`

NewMetadataParamWithDefaults instantiates a new MetadataParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *MetadataParam) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *MetadataParam) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *MetadataParam) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *MetadataParam) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetLanguage

`func (o *MetadataParam) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MetadataParam) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MetadataParam) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MetadataParam) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


