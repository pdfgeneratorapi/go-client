# DocumentSignatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The one-line answer, reduced from every signature worst case first. &#x60;unavailable&#x60; means signature validation is not enabled on this deployment — it says nothing about the document.  | [optional] 
**StatusLabel** | Pointer to **string** | The status in words, ready to display. | [optional] 
**HasLongTermValidation** | Pointer to **bool** | The document carries the certificate and revocation data needed to verify it after the signing certificates expire.  | [optional] 
**CoversWholeDocument** | Pointer to **bool** | The last signature covers every byte, so nothing was appended after it. | [optional] 
**DocumentTimestamps** | Pointer to **int32** | Timestamp-only signatures, counted rather than listed: they are machinery, not people, and would double every signer.  | [optional] 
**Signatures** | Pointer to [**[]DocumentSignature**](DocumentSignature.md) | One entry per signature, excluding document timestamps. | [optional] 

## Methods

### NewDocumentSignatures

`func NewDocumentSignatures() *DocumentSignatures`

NewDocumentSignatures instantiates a new DocumentSignatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentSignaturesWithDefaults

`func NewDocumentSignaturesWithDefaults() *DocumentSignatures`

NewDocumentSignaturesWithDefaults instantiates a new DocumentSignatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DocumentSignatures) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentSignatures) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentSignatures) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocumentSignatures) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusLabel

`func (o *DocumentSignatures) GetStatusLabel() string`

GetStatusLabel returns the StatusLabel field if non-nil, zero value otherwise.

### GetStatusLabelOk

`func (o *DocumentSignatures) GetStatusLabelOk() (*string, bool)`

GetStatusLabelOk returns a tuple with the StatusLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLabel

`func (o *DocumentSignatures) SetStatusLabel(v string)`

SetStatusLabel sets StatusLabel field to given value.

### HasStatusLabel

`func (o *DocumentSignatures) HasStatusLabel() bool`

HasStatusLabel returns a boolean if a field has been set.

### GetHasLongTermValidation

`func (o *DocumentSignatures) GetHasLongTermValidation() bool`

GetHasLongTermValidation returns the HasLongTermValidation field if non-nil, zero value otherwise.

### GetHasLongTermValidationOk

`func (o *DocumentSignatures) GetHasLongTermValidationOk() (*bool, bool)`

GetHasLongTermValidationOk returns a tuple with the HasLongTermValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLongTermValidation

`func (o *DocumentSignatures) SetHasLongTermValidation(v bool)`

SetHasLongTermValidation sets HasLongTermValidation field to given value.

### HasHasLongTermValidation

`func (o *DocumentSignatures) HasHasLongTermValidation() bool`

HasHasLongTermValidation returns a boolean if a field has been set.

### GetCoversWholeDocument

`func (o *DocumentSignatures) GetCoversWholeDocument() bool`

GetCoversWholeDocument returns the CoversWholeDocument field if non-nil, zero value otherwise.

### GetCoversWholeDocumentOk

`func (o *DocumentSignatures) GetCoversWholeDocumentOk() (*bool, bool)`

GetCoversWholeDocumentOk returns a tuple with the CoversWholeDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoversWholeDocument

`func (o *DocumentSignatures) SetCoversWholeDocument(v bool)`

SetCoversWholeDocument sets CoversWholeDocument field to given value.

### HasCoversWholeDocument

`func (o *DocumentSignatures) HasCoversWholeDocument() bool`

HasCoversWholeDocument returns a boolean if a field has been set.

### GetDocumentTimestamps

`func (o *DocumentSignatures) GetDocumentTimestamps() int32`

GetDocumentTimestamps returns the DocumentTimestamps field if non-nil, zero value otherwise.

### GetDocumentTimestampsOk

`func (o *DocumentSignatures) GetDocumentTimestampsOk() (*int32, bool)`

GetDocumentTimestampsOk returns a tuple with the DocumentTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTimestamps

`func (o *DocumentSignatures) SetDocumentTimestamps(v int32)`

SetDocumentTimestamps sets DocumentTimestamps field to given value.

### HasDocumentTimestamps

`func (o *DocumentSignatures) HasDocumentTimestamps() bool`

HasDocumentTimestamps returns a boolean if a field has been set.

### GetSignatures

`func (o *DocumentSignatures) GetSignatures() []DocumentSignature`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *DocumentSignatures) GetSignaturesOk() (*[]DocumentSignature, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *DocumentSignatures) SetSignatures(v []DocumentSignature)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *DocumentSignatures) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


