# TemplateDefinitionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Template name | [optional] 
**Tags** | Pointer to **[]string** | A list of tags assigned to a template | [optional] 
**IsDraft** | Pointer to **bool** | Indicates if the template is a draft or published. | [optional] 
**Layout** | Pointer to [**TemplateDefinitionNewLayout**](TemplateDefinitionNewLayout.md) |  | [optional] 
**Pages** | Pointer to [**[]TemplateDefinitionNewPagesInner**](TemplateDefinitionNewPagesInner.md) | Defines page or label size, margins and components on page or label | [optional] 
**DataSettings** | Pointer to [**TemplateDefinitionNewDataSettings**](TemplateDefinitionNewDataSettings.md) |  | [optional] 
**Editor** | Pointer to [**TemplateDefinitionNewEditor**](TemplateDefinitionNewEditor.md) |  | [optional] 
**FontSubsetting** | Pointer to **bool** | If font-subsetting is applied to document when generated | [optional] [default to false]
**BarcodeAsImage** | Pointer to **bool** | Defines if barcodes are rendered as raster images instead of vector graphics. | [optional] [default to false]

## Methods

### NewTemplateDefinitionUpdate

`func NewTemplateDefinitionUpdate() *TemplateDefinitionUpdate`

NewTemplateDefinitionUpdate instantiates a new TemplateDefinitionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDefinitionUpdateWithDefaults

`func NewTemplateDefinitionUpdateWithDefaults() *TemplateDefinitionUpdate`

NewTemplateDefinitionUpdateWithDefaults instantiates a new TemplateDefinitionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateDefinitionUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateDefinitionUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateDefinitionUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateDefinitionUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *TemplateDefinitionUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TemplateDefinitionUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TemplateDefinitionUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TemplateDefinitionUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsDraft

`func (o *TemplateDefinitionUpdate) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *TemplateDefinitionUpdate) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *TemplateDefinitionUpdate) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *TemplateDefinitionUpdate) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### GetLayout

`func (o *TemplateDefinitionUpdate) GetLayout() TemplateDefinitionNewLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *TemplateDefinitionUpdate) GetLayoutOk() (*TemplateDefinitionNewLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *TemplateDefinitionUpdate) SetLayout(v TemplateDefinitionNewLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *TemplateDefinitionUpdate) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetPages

`func (o *TemplateDefinitionUpdate) GetPages() []TemplateDefinitionNewPagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *TemplateDefinitionUpdate) GetPagesOk() (*[]TemplateDefinitionNewPagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *TemplateDefinitionUpdate) SetPages(v []TemplateDefinitionNewPagesInner)`

SetPages sets Pages field to given value.

### HasPages

`func (o *TemplateDefinitionUpdate) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetDataSettings

`func (o *TemplateDefinitionUpdate) GetDataSettings() TemplateDefinitionNewDataSettings`

GetDataSettings returns the DataSettings field if non-nil, zero value otherwise.

### GetDataSettingsOk

`func (o *TemplateDefinitionUpdate) GetDataSettingsOk() (*TemplateDefinitionNewDataSettings, bool)`

GetDataSettingsOk returns a tuple with the DataSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSettings

`func (o *TemplateDefinitionUpdate) SetDataSettings(v TemplateDefinitionNewDataSettings)`

SetDataSettings sets DataSettings field to given value.

### HasDataSettings

`func (o *TemplateDefinitionUpdate) HasDataSettings() bool`

HasDataSettings returns a boolean if a field has been set.

### GetEditor

`func (o *TemplateDefinitionUpdate) GetEditor() TemplateDefinitionNewEditor`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *TemplateDefinitionUpdate) GetEditorOk() (*TemplateDefinitionNewEditor, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *TemplateDefinitionUpdate) SetEditor(v TemplateDefinitionNewEditor)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *TemplateDefinitionUpdate) HasEditor() bool`

HasEditor returns a boolean if a field has been set.

### GetFontSubsetting

`func (o *TemplateDefinitionUpdate) GetFontSubsetting() bool`

GetFontSubsetting returns the FontSubsetting field if non-nil, zero value otherwise.

### GetFontSubsettingOk

`func (o *TemplateDefinitionUpdate) GetFontSubsettingOk() (*bool, bool)`

GetFontSubsettingOk returns a tuple with the FontSubsetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSubsetting

`func (o *TemplateDefinitionUpdate) SetFontSubsetting(v bool)`

SetFontSubsetting sets FontSubsetting field to given value.

### HasFontSubsetting

`func (o *TemplateDefinitionUpdate) HasFontSubsetting() bool`

HasFontSubsetting returns a boolean if a field has been set.

### GetBarcodeAsImage

`func (o *TemplateDefinitionUpdate) GetBarcodeAsImage() bool`

GetBarcodeAsImage returns the BarcodeAsImage field if non-nil, zero value otherwise.

### GetBarcodeAsImageOk

`func (o *TemplateDefinitionUpdate) GetBarcodeAsImageOk() (*bool, bool)`

GetBarcodeAsImageOk returns a tuple with the BarcodeAsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeAsImage

`func (o *TemplateDefinitionUpdate) SetBarcodeAsImage(v bool)`

SetBarcodeAsImage sets BarcodeAsImage field to given value.

### HasBarcodeAsImage

`func (o *TemplateDefinitionUpdate) HasBarcodeAsImage() bool`

HasBarcodeAsImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


