# OpenEditorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**OpenEditorRequestData**](OpenEditorRequestData.md) |  | [optional] 
**Language** | Pointer to **string** | Specify the editor UI language. Defaults to organization editor language. | [optional] 

## Methods

### NewOpenEditorRequest

`func NewOpenEditorRequest() *OpenEditorRequest`

NewOpenEditorRequest instantiates a new OpenEditorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenEditorRequestWithDefaults

`func NewOpenEditorRequestWithDefaults() *OpenEditorRequest`

NewOpenEditorRequestWithDefaults instantiates a new OpenEditorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OpenEditorRequest) GetData() OpenEditorRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OpenEditorRequest) GetDataOk() (*OpenEditorRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OpenEditorRequest) SetData(v OpenEditorRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *OpenEditorRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLanguage

`func (o *OpenEditorRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *OpenEditorRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *OpenEditorRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *OpenEditorRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


