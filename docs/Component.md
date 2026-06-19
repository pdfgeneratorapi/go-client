# Component

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cls** | Pointer to **string** | Defines component class/type | [optional] 
**Id** | Pointer to **string** | Component id | [optional] 
**Width** | Pointer to **float32** | Width in units | [optional] 
**Height** | Pointer to **float32** | Height in units | [optional] 
**Top** | Pointer to **float32** | Position from the page top in units | [optional] 
**Left** | Pointer to **float32** | Position from the page left in units | [optional] 
**Zindex** | Pointer to **int32** | Defines the rendering order on page. Components with smaller zindex are rendered before | [optional] 
**Value** | Pointer to **string** | Component value | [optional] 
**DataIndex** | Pointer to **string** | Defines data field for Table and Container components which are used to iterate over list of items | [optional] 

## Methods

### NewComponent

`func NewComponent() *Component`

NewComponent instantiates a new Component object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentWithDefaults

`func NewComponentWithDefaults() *Component`

NewComponentWithDefaults instantiates a new Component object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCls

`func (o *Component) GetCls() string`

GetCls returns the Cls field if non-nil, zero value otherwise.

### GetClsOk

`func (o *Component) GetClsOk() (*string, bool)`

GetClsOk returns a tuple with the Cls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCls

`func (o *Component) SetCls(v string)`

SetCls sets Cls field to given value.

### HasCls

`func (o *Component) HasCls() bool`

HasCls returns a boolean if a field has been set.

### GetId

`func (o *Component) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Component) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Component) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Component) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWidth

`func (o *Component) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Component) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Component) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Component) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *Component) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Component) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Component) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Component) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetTop

`func (o *Component) GetTop() float32`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *Component) GetTopOk() (*float32, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *Component) SetTop(v float32)`

SetTop sets Top field to given value.

### HasTop

`func (o *Component) HasTop() bool`

HasTop returns a boolean if a field has been set.

### GetLeft

`func (o *Component) GetLeft() float32`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *Component) GetLeftOk() (*float32, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *Component) SetLeft(v float32)`

SetLeft sets Left field to given value.

### HasLeft

`func (o *Component) HasLeft() bool`

HasLeft returns a boolean if a field has been set.

### GetZindex

`func (o *Component) GetZindex() int32`

GetZindex returns the Zindex field if non-nil, zero value otherwise.

### GetZindexOk

`func (o *Component) GetZindexOk() (*int32, bool)`

GetZindexOk returns a tuple with the Zindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZindex

`func (o *Component) SetZindex(v int32)`

SetZindex sets Zindex field to given value.

### HasZindex

`func (o *Component) HasZindex() bool`

HasZindex returns a boolean if a field has been set.

### GetValue

`func (o *Component) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Component) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Component) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Component) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDataIndex

`func (o *Component) GetDataIndex() string`

GetDataIndex returns the DataIndex field if non-nil, zero value otherwise.

### GetDataIndexOk

`func (o *Component) GetDataIndexOk() (*string, bool)`

GetDataIndexOk returns a tuple with the DataIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIndex

`func (o *Component) SetDataIndex(v string)`

SetDataIndex sets DataIndex field to given value.

### HasDataIndex

`func (o *Component) HasDataIndex() bool`

HasDataIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


