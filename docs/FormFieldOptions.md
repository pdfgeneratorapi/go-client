# FormFieldOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | **string** | Comma-separated list of choices. Whitespace around each choice is trimmed and choices must be unique. | 

## Methods

### NewFormFieldOptions

`func NewFormFieldOptions(options string, ) *FormFieldOptions`

NewFormFieldOptions instantiates a new FormFieldOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldOptionsWithDefaults

`func NewFormFieldOptionsWithDefaults() *FormFieldOptions`

NewFormFieldOptionsWithDefaults instantiates a new FormFieldOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *FormFieldOptions) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormFieldOptions) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormFieldOptions) SetOptions(v string)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


