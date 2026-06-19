# GenerateQRCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The content which is used to generate QR code | 
**Color** | Pointer to **string** | QR code in hexadecimal format | [optional] [default to "#000000"]
**LogoBase64** | Pointer to **string** | A logo as a base64 image string to add on the QR code | [optional] 
**LogoUrl** | Pointer to **string** | A logo URL to an image to add on the QR code | [optional] 
**Output** | Pointer to **string** | Response format | [optional] [default to "base64"]

## Methods

### NewGenerateQRCodeRequest

`func NewGenerateQRCodeRequest(content string, ) *GenerateQRCodeRequest`

NewGenerateQRCodeRequest instantiates a new GenerateQRCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateQRCodeRequestWithDefaults

`func NewGenerateQRCodeRequestWithDefaults() *GenerateQRCodeRequest`

NewGenerateQRCodeRequestWithDefaults instantiates a new GenerateQRCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *GenerateQRCodeRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GenerateQRCodeRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GenerateQRCodeRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetColor

`func (o *GenerateQRCodeRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *GenerateQRCodeRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *GenerateQRCodeRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *GenerateQRCodeRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLogoBase64

`func (o *GenerateQRCodeRequest) GetLogoBase64() string`

GetLogoBase64 returns the LogoBase64 field if non-nil, zero value otherwise.

### GetLogoBase64Ok

`func (o *GenerateQRCodeRequest) GetLogoBase64Ok() (*string, bool)`

GetLogoBase64Ok returns a tuple with the LogoBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoBase64

`func (o *GenerateQRCodeRequest) SetLogoBase64(v string)`

SetLogoBase64 sets LogoBase64 field to given value.

### HasLogoBase64

`func (o *GenerateQRCodeRequest) HasLogoBase64() bool`

HasLogoBase64 returns a boolean if a field has been set.

### GetLogoUrl

`func (o *GenerateQRCodeRequest) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *GenerateQRCodeRequest) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *GenerateQRCodeRequest) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *GenerateQRCodeRequest) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetOutput

`func (o *GenerateQRCodeRequest) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GenerateQRCodeRequest) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GenerateQRCodeRequest) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GenerateQRCodeRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


