# DocumentSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** | Name of the signature field in the document. | [optional] 
**SignerName** | Pointer to **NullableString** | The name the signer signed under. | [optional] 
**Signer** | Pointer to **NullableString** | Subject of the certificate that sealed it — an organization, not a person. | [optional] 
**Issuer** | Pointer to **NullableString** | Subject of the certificate authority that issued it. | [optional] 
**SignedAt** | Pointer to **NullableString** | When a timestamp authority attested the signature (ISO-8601). This is the defensible time; the signer&#39;s own clock proves nothing.  | [optional] 
**ClaimedSignedAt** | Pointer to **NullableString** | The time the signer&#39;s own software recorded (ISO-8601). | [optional] 
**TimestampAuthority** | Pointer to **NullableString** | The timestamp authority that attested the signature. | [optional] 
**Intact** | Pointer to **NullableBool** | The bytes this signature covers are unchanged. | [optional] 
**Valid** | Pointer to **NullableBool** | The signature block itself adds up. On its own this does NOT mean the document is unchanged: a tampered file reports &#x60;valid&#x60; true with &#x60;intact&#x60; false, so a verdict needs intact AND valid AND trusted.  | [optional] 
**Trusted** | Pointer to **NullableBool** | The certificate chains to a trusted root. | [optional] 
**Coverage** | Pointer to **NullableString** | How much of the file this signature protects. | [optional] 
**AdesIndication** | Pointer to **NullableString** | The ETSI EN 319 102-1 indication, when the AdES engine could run. | [optional] 
**AdesSubIndication** | Pointer to **NullableString** | The AdES sub-indication, naming why an indication is not PASSED. | [optional] 

## Methods

### NewDocumentSignature

`func NewDocumentSignature() *DocumentSignature`

NewDocumentSignature instantiates a new DocumentSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentSignatureWithDefaults

`func NewDocumentSignatureWithDefaults() *DocumentSignature`

NewDocumentSignatureWithDefaults instantiates a new DocumentSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *DocumentSignature) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *DocumentSignature) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *DocumentSignature) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *DocumentSignature) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetSignerName

`func (o *DocumentSignature) GetSignerName() string`

GetSignerName returns the SignerName field if non-nil, zero value otherwise.

### GetSignerNameOk

`func (o *DocumentSignature) GetSignerNameOk() (*string, bool)`

GetSignerNameOk returns a tuple with the SignerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerName

`func (o *DocumentSignature) SetSignerName(v string)`

SetSignerName sets SignerName field to given value.

### HasSignerName

`func (o *DocumentSignature) HasSignerName() bool`

HasSignerName returns a boolean if a field has been set.

### SetSignerNameNil

`func (o *DocumentSignature) SetSignerNameNil(b bool)`

 SetSignerNameNil sets the value for SignerName to be an explicit nil

### UnsetSignerName
`func (o *DocumentSignature) UnsetSignerName()`

UnsetSignerName ensures that no value is present for SignerName, not even an explicit nil
### GetSigner

`func (o *DocumentSignature) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *DocumentSignature) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *DocumentSignature) SetSigner(v string)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *DocumentSignature) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### SetSignerNil

`func (o *DocumentSignature) SetSignerNil(b bool)`

 SetSignerNil sets the value for Signer to be an explicit nil

### UnsetSigner
`func (o *DocumentSignature) UnsetSigner()`

UnsetSigner ensures that no value is present for Signer, not even an explicit nil
### GetIssuer

`func (o *DocumentSignature) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *DocumentSignature) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *DocumentSignature) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *DocumentSignature) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *DocumentSignature) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *DocumentSignature) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetSignedAt

`func (o *DocumentSignature) GetSignedAt() string`

GetSignedAt returns the SignedAt field if non-nil, zero value otherwise.

### GetSignedAtOk

`func (o *DocumentSignature) GetSignedAtOk() (*string, bool)`

GetSignedAtOk returns a tuple with the SignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedAt

`func (o *DocumentSignature) SetSignedAt(v string)`

SetSignedAt sets SignedAt field to given value.

### HasSignedAt

`func (o *DocumentSignature) HasSignedAt() bool`

HasSignedAt returns a boolean if a field has been set.

### SetSignedAtNil

`func (o *DocumentSignature) SetSignedAtNil(b bool)`

 SetSignedAtNil sets the value for SignedAt to be an explicit nil

### UnsetSignedAt
`func (o *DocumentSignature) UnsetSignedAt()`

UnsetSignedAt ensures that no value is present for SignedAt, not even an explicit nil
### GetClaimedSignedAt

`func (o *DocumentSignature) GetClaimedSignedAt() string`

GetClaimedSignedAt returns the ClaimedSignedAt field if non-nil, zero value otherwise.

### GetClaimedSignedAtOk

`func (o *DocumentSignature) GetClaimedSignedAtOk() (*string, bool)`

GetClaimedSignedAtOk returns a tuple with the ClaimedSignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedSignedAt

`func (o *DocumentSignature) SetClaimedSignedAt(v string)`

SetClaimedSignedAt sets ClaimedSignedAt field to given value.

### HasClaimedSignedAt

`func (o *DocumentSignature) HasClaimedSignedAt() bool`

HasClaimedSignedAt returns a boolean if a field has been set.

### SetClaimedSignedAtNil

`func (o *DocumentSignature) SetClaimedSignedAtNil(b bool)`

 SetClaimedSignedAtNil sets the value for ClaimedSignedAt to be an explicit nil

### UnsetClaimedSignedAt
`func (o *DocumentSignature) UnsetClaimedSignedAt()`

UnsetClaimedSignedAt ensures that no value is present for ClaimedSignedAt, not even an explicit nil
### GetTimestampAuthority

`func (o *DocumentSignature) GetTimestampAuthority() string`

GetTimestampAuthority returns the TimestampAuthority field if non-nil, zero value otherwise.

### GetTimestampAuthorityOk

`func (o *DocumentSignature) GetTimestampAuthorityOk() (*string, bool)`

GetTimestampAuthorityOk returns a tuple with the TimestampAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampAuthority

`func (o *DocumentSignature) SetTimestampAuthority(v string)`

SetTimestampAuthority sets TimestampAuthority field to given value.

### HasTimestampAuthority

`func (o *DocumentSignature) HasTimestampAuthority() bool`

HasTimestampAuthority returns a boolean if a field has been set.

### SetTimestampAuthorityNil

`func (o *DocumentSignature) SetTimestampAuthorityNil(b bool)`

 SetTimestampAuthorityNil sets the value for TimestampAuthority to be an explicit nil

### UnsetTimestampAuthority
`func (o *DocumentSignature) UnsetTimestampAuthority()`

UnsetTimestampAuthority ensures that no value is present for TimestampAuthority, not even an explicit nil
### GetIntact

`func (o *DocumentSignature) GetIntact() bool`

GetIntact returns the Intact field if non-nil, zero value otherwise.

### GetIntactOk

`func (o *DocumentSignature) GetIntactOk() (*bool, bool)`

GetIntactOk returns a tuple with the Intact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntact

`func (o *DocumentSignature) SetIntact(v bool)`

SetIntact sets Intact field to given value.

### HasIntact

`func (o *DocumentSignature) HasIntact() bool`

HasIntact returns a boolean if a field has been set.

### SetIntactNil

`func (o *DocumentSignature) SetIntactNil(b bool)`

 SetIntactNil sets the value for Intact to be an explicit nil

### UnsetIntact
`func (o *DocumentSignature) UnsetIntact()`

UnsetIntact ensures that no value is present for Intact, not even an explicit nil
### GetValid

`func (o *DocumentSignature) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *DocumentSignature) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *DocumentSignature) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *DocumentSignature) HasValid() bool`

HasValid returns a boolean if a field has been set.

### SetValidNil

`func (o *DocumentSignature) SetValidNil(b bool)`

 SetValidNil sets the value for Valid to be an explicit nil

### UnsetValid
`func (o *DocumentSignature) UnsetValid()`

UnsetValid ensures that no value is present for Valid, not even an explicit nil
### GetTrusted

`func (o *DocumentSignature) GetTrusted() bool`

GetTrusted returns the Trusted field if non-nil, zero value otherwise.

### GetTrustedOk

`func (o *DocumentSignature) GetTrustedOk() (*bool, bool)`

GetTrustedOk returns a tuple with the Trusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrusted

`func (o *DocumentSignature) SetTrusted(v bool)`

SetTrusted sets Trusted field to given value.

### HasTrusted

`func (o *DocumentSignature) HasTrusted() bool`

HasTrusted returns a boolean if a field has been set.

### SetTrustedNil

`func (o *DocumentSignature) SetTrustedNil(b bool)`

 SetTrustedNil sets the value for Trusted to be an explicit nil

### UnsetTrusted
`func (o *DocumentSignature) UnsetTrusted()`

UnsetTrusted ensures that no value is present for Trusted, not even an explicit nil
### GetCoverage

`func (o *DocumentSignature) GetCoverage() string`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *DocumentSignature) GetCoverageOk() (*string, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *DocumentSignature) SetCoverage(v string)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *DocumentSignature) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### SetCoverageNil

`func (o *DocumentSignature) SetCoverageNil(b bool)`

 SetCoverageNil sets the value for Coverage to be an explicit nil

### UnsetCoverage
`func (o *DocumentSignature) UnsetCoverage()`

UnsetCoverage ensures that no value is present for Coverage, not even an explicit nil
### GetAdesIndication

`func (o *DocumentSignature) GetAdesIndication() string`

GetAdesIndication returns the AdesIndication field if non-nil, zero value otherwise.

### GetAdesIndicationOk

`func (o *DocumentSignature) GetAdesIndicationOk() (*string, bool)`

GetAdesIndicationOk returns a tuple with the AdesIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdesIndication

`func (o *DocumentSignature) SetAdesIndication(v string)`

SetAdesIndication sets AdesIndication field to given value.

### HasAdesIndication

`func (o *DocumentSignature) HasAdesIndication() bool`

HasAdesIndication returns a boolean if a field has been set.

### SetAdesIndicationNil

`func (o *DocumentSignature) SetAdesIndicationNil(b bool)`

 SetAdesIndicationNil sets the value for AdesIndication to be an explicit nil

### UnsetAdesIndication
`func (o *DocumentSignature) UnsetAdesIndication()`

UnsetAdesIndication ensures that no value is present for AdesIndication, not even an explicit nil
### GetAdesSubIndication

`func (o *DocumentSignature) GetAdesSubIndication() string`

GetAdesSubIndication returns the AdesSubIndication field if non-nil, zero value otherwise.

### GetAdesSubIndicationOk

`func (o *DocumentSignature) GetAdesSubIndicationOk() (*string, bool)`

GetAdesSubIndicationOk returns a tuple with the AdesSubIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdesSubIndication

`func (o *DocumentSignature) SetAdesSubIndication(v string)`

SetAdesSubIndication sets AdesSubIndication field to given value.

### HasAdesSubIndication

`func (o *DocumentSignature) HasAdesSubIndication() bool`

HasAdesSubIndication returns a boolean if a field has been set.

### SetAdesSubIndicationNil

`func (o *DocumentSignature) SetAdesSubIndicationNil(b bool)`

 SetAdesSubIndicationNil sets the value for AdesSubIndication to be an explicit nil

### UnsetAdesSubIndication
`func (o *DocumentSignature) UnsetAdesSubIndication()`

UnsetAdesSubIndication ensures that no value is present for AdesSubIndication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


