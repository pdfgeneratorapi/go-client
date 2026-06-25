# DocumentUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**IdCode** | Pointer to **NullableString** | National identification code, when provided. | [optional] 
**Ip** | Pointer to **NullableString** | IP address captured when the action was performed. | [optional] 

## Methods

### NewDocumentUser

`func NewDocumentUser() *DocumentUser`

NewDocumentUser instantiates a new DocumentUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentUserWithDefaults

`func NewDocumentUserWithDefaults() *DocumentUser`

NewDocumentUserWithDefaults instantiates a new DocumentUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *DocumentUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DocumentUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DocumentUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *DocumentUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *DocumentUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DocumentUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DocumentUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *DocumentUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *DocumentUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DocumentUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DocumentUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DocumentUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIdCode

`func (o *DocumentUser) GetIdCode() string`

GetIdCode returns the IdCode field if non-nil, zero value otherwise.

### GetIdCodeOk

`func (o *DocumentUser) GetIdCodeOk() (*string, bool)`

GetIdCodeOk returns a tuple with the IdCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdCode

`func (o *DocumentUser) SetIdCode(v string)`

SetIdCode sets IdCode field to given value.

### HasIdCode

`func (o *DocumentUser) HasIdCode() bool`

HasIdCode returns a boolean if a field has been set.

### SetIdCodeNil

`func (o *DocumentUser) SetIdCodeNil(b bool)`

 SetIdCodeNil sets the value for IdCode to be an explicit nil

### UnsetIdCode
`func (o *DocumentUser) UnsetIdCode()`

UnsetIdCode ensures that no value is present for IdCode, not even an explicit nil
### GetIp

`func (o *DocumentUser) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DocumentUser) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DocumentUser) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DocumentUser) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *DocumentUser) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *DocumentUser) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


