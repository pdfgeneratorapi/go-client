# InlineObject13ResponseValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique field identifier | [optional] 
**Type** | Pointer to **string** | Field type | [optional] 
**Name** | Pointer to **string** | Field name | [optional] 
**Value** | Pointer to [**InlineObject13ResponseValueValue**](InlineObject13ResponseValueValue.md) |  | [optional] 
**Locked** | Pointer to **bool** | Whether the field is locked | [optional] 
**Pages** | Pointer to **[]int32** | Pages where the field appears | [optional] 
**Default** | Pointer to [**InlineObject13ResponseValueDefault**](InlineObject13ResponseValueDefault.md) |  | [optional] 
**Defaults** | Pointer to **[]string** | Default values for multi-select fields | [optional] 
**Values** | Pointer to **[]string** | Selected values for multi-select fields | [optional] 
**Options** | Pointer to **[]string** | Available options for select fields | [optional] 
**Format** | Pointer to **string** | Field format (for date fields) | [optional] 
**Maxlen** | Pointer to **int32** | Maximum field length | [optional] 
**Multiline** | Pointer to **bool** | Whether text field is multiline | [optional] 
**Editable** | Pointer to **bool** | Whether combo box is editable | [optional] 
**Multi** | Pointer to **bool** | Whether list box allows multiple selections | [optional] 

## Methods

### NewInlineObject13ResponseValue

`func NewInlineObject13ResponseValue() *InlineObject13ResponseValue`

NewInlineObject13ResponseValue instantiates a new InlineObject13ResponseValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject13ResponseValueWithDefaults

`func NewInlineObject13ResponseValueWithDefaults() *InlineObject13ResponseValue`

NewInlineObject13ResponseValueWithDefaults instantiates a new InlineObject13ResponseValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineObject13ResponseValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineObject13ResponseValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineObject13ResponseValue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineObject13ResponseValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *InlineObject13ResponseValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject13ResponseValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject13ResponseValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineObject13ResponseValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InlineObject13ResponseValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject13ResponseValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject13ResponseValue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject13ResponseValue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *InlineObject13ResponseValue) GetValue() InlineObject13ResponseValueValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InlineObject13ResponseValue) GetValueOk() (*InlineObject13ResponseValueValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InlineObject13ResponseValue) SetValue(v InlineObject13ResponseValueValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *InlineObject13ResponseValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLocked

`func (o *InlineObject13ResponseValue) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *InlineObject13ResponseValue) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *InlineObject13ResponseValue) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *InlineObject13ResponseValue) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPages

`func (o *InlineObject13ResponseValue) GetPages() []int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *InlineObject13ResponseValue) GetPagesOk() (*[]int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *InlineObject13ResponseValue) SetPages(v []int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *InlineObject13ResponseValue) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetDefault

`func (o *InlineObject13ResponseValue) GetDefault() InlineObject13ResponseValueDefault`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *InlineObject13ResponseValue) GetDefaultOk() (*InlineObject13ResponseValueDefault, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *InlineObject13ResponseValue) SetDefault(v InlineObject13ResponseValueDefault)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *InlineObject13ResponseValue) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDefaults

`func (o *InlineObject13ResponseValue) GetDefaults() []string`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *InlineObject13ResponseValue) GetDefaultsOk() (*[]string, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *InlineObject13ResponseValue) SetDefaults(v []string)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *InlineObject13ResponseValue) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### GetValues

`func (o *InlineObject13ResponseValue) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InlineObject13ResponseValue) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InlineObject13ResponseValue) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *InlineObject13ResponseValue) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetOptions

`func (o *InlineObject13ResponseValue) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *InlineObject13ResponseValue) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *InlineObject13ResponseValue) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *InlineObject13ResponseValue) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFormat

`func (o *InlineObject13ResponseValue) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *InlineObject13ResponseValue) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *InlineObject13ResponseValue) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *InlineObject13ResponseValue) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMaxlen

`func (o *InlineObject13ResponseValue) GetMaxlen() int32`

GetMaxlen returns the Maxlen field if non-nil, zero value otherwise.

### GetMaxlenOk

`func (o *InlineObject13ResponseValue) GetMaxlenOk() (*int32, bool)`

GetMaxlenOk returns a tuple with the Maxlen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxlen

`func (o *InlineObject13ResponseValue) SetMaxlen(v int32)`

SetMaxlen sets Maxlen field to given value.

### HasMaxlen

`func (o *InlineObject13ResponseValue) HasMaxlen() bool`

HasMaxlen returns a boolean if a field has been set.

### GetMultiline

`func (o *InlineObject13ResponseValue) GetMultiline() bool`

GetMultiline returns the Multiline field if non-nil, zero value otherwise.

### GetMultilineOk

`func (o *InlineObject13ResponseValue) GetMultilineOk() (*bool, bool)`

GetMultilineOk returns a tuple with the Multiline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiline

`func (o *InlineObject13ResponseValue) SetMultiline(v bool)`

SetMultiline sets Multiline field to given value.

### HasMultiline

`func (o *InlineObject13ResponseValue) HasMultiline() bool`

HasMultiline returns a boolean if a field has been set.

### GetEditable

`func (o *InlineObject13ResponseValue) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *InlineObject13ResponseValue) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *InlineObject13ResponseValue) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *InlineObject13ResponseValue) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetMulti

`func (o *InlineObject13ResponseValue) GetMulti() bool`

GetMulti returns the Multi field if non-nil, zero value otherwise.

### GetMultiOk

`func (o *InlineObject13ResponseValue) GetMultiOk() (*bool, bool)`

GetMultiOk returns a tuple with the Multi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulti

`func (o *InlineObject13ResponseValue) SetMulti(v bool)`

SetMulti sets Multi field to given value.

### HasMulti

`func (o *InlineObject13ResponseValue) HasMulti() bool`

HasMulti returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


