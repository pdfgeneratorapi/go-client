# TemplateDefinition

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
**Id** | Pointer to **NullableInt32** | Unique identifier. &#x60;null&#x60; for a template library preview, which is not a template of the workspace. | [optional] 
**ImprovedHtmlSpacing** | Pointer to **bool** | Measures HTML text the way it is drawn. New templates have it on; one saved before it existed keeps it off, and the layout it was designed with. | [optional] 
**BackgroundPdf** | Pointer to **NullableString** | A base64 PDF the pages are drawn on. Set when the template is imported from a PDF. | [optional] 

## Methods

### NewTemplateDefinition

`func NewTemplateDefinition() *TemplateDefinition`

NewTemplateDefinition instantiates a new TemplateDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDefinitionWithDefaults

`func NewTemplateDefinitionWithDefaults() *TemplateDefinition`

NewTemplateDefinitionWithDefaults instantiates a new TemplateDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *TemplateDefinition) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TemplateDefinition) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TemplateDefinition) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TemplateDefinition) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsDraft

`func (o *TemplateDefinition) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *TemplateDefinition) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *TemplateDefinition) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *TemplateDefinition) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### GetLayout

`func (o *TemplateDefinition) GetLayout() TemplateDefinitionNewLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *TemplateDefinition) GetLayoutOk() (*TemplateDefinitionNewLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *TemplateDefinition) SetLayout(v TemplateDefinitionNewLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *TemplateDefinition) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetPages

`func (o *TemplateDefinition) GetPages() []TemplateDefinitionNewPagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *TemplateDefinition) GetPagesOk() (*[]TemplateDefinitionNewPagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *TemplateDefinition) SetPages(v []TemplateDefinitionNewPagesInner)`

SetPages sets Pages field to given value.

### HasPages

`func (o *TemplateDefinition) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetDataSettings

`func (o *TemplateDefinition) GetDataSettings() TemplateDefinitionNewDataSettings`

GetDataSettings returns the DataSettings field if non-nil, zero value otherwise.

### GetDataSettingsOk

`func (o *TemplateDefinition) GetDataSettingsOk() (*TemplateDefinitionNewDataSettings, bool)`

GetDataSettingsOk returns a tuple with the DataSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSettings

`func (o *TemplateDefinition) SetDataSettings(v TemplateDefinitionNewDataSettings)`

SetDataSettings sets DataSettings field to given value.

### HasDataSettings

`func (o *TemplateDefinition) HasDataSettings() bool`

HasDataSettings returns a boolean if a field has been set.

### GetEditor

`func (o *TemplateDefinition) GetEditor() TemplateDefinitionNewEditor`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *TemplateDefinition) GetEditorOk() (*TemplateDefinitionNewEditor, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *TemplateDefinition) SetEditor(v TemplateDefinitionNewEditor)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *TemplateDefinition) HasEditor() bool`

HasEditor returns a boolean if a field has been set.

### GetFontSubsetting

`func (o *TemplateDefinition) GetFontSubsetting() bool`

GetFontSubsetting returns the FontSubsetting field if non-nil, zero value otherwise.

### GetFontSubsettingOk

`func (o *TemplateDefinition) GetFontSubsettingOk() (*bool, bool)`

GetFontSubsettingOk returns a tuple with the FontSubsetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSubsetting

`func (o *TemplateDefinition) SetFontSubsetting(v bool)`

SetFontSubsetting sets FontSubsetting field to given value.

### HasFontSubsetting

`func (o *TemplateDefinition) HasFontSubsetting() bool`

HasFontSubsetting returns a boolean if a field has been set.

### GetBarcodeAsImage

`func (o *TemplateDefinition) GetBarcodeAsImage() bool`

GetBarcodeAsImage returns the BarcodeAsImage field if non-nil, zero value otherwise.

### GetBarcodeAsImageOk

`func (o *TemplateDefinition) GetBarcodeAsImageOk() (*bool, bool)`

GetBarcodeAsImageOk returns a tuple with the BarcodeAsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeAsImage

`func (o *TemplateDefinition) SetBarcodeAsImage(v bool)`

SetBarcodeAsImage sets BarcodeAsImage field to given value.

### HasBarcodeAsImage

`func (o *TemplateDefinition) HasBarcodeAsImage() bool`

HasBarcodeAsImage returns a boolean if a field has been set.

### GetId

`func (o *TemplateDefinition) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateDefinition) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateDefinition) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TemplateDefinition) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TemplateDefinition) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetImprovedHtmlSpacing

`func (o *TemplateDefinition) GetImprovedHtmlSpacing() bool`

GetImprovedHtmlSpacing returns the ImprovedHtmlSpacing field if non-nil, zero value otherwise.

### GetImprovedHtmlSpacingOk

`func (o *TemplateDefinition) GetImprovedHtmlSpacingOk() (*bool, bool)`

GetImprovedHtmlSpacingOk returns a tuple with the ImprovedHtmlSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImprovedHtmlSpacing

`func (o *TemplateDefinition) SetImprovedHtmlSpacing(v bool)`

SetImprovedHtmlSpacing sets ImprovedHtmlSpacing field to given value.

### HasImprovedHtmlSpacing

`func (o *TemplateDefinition) HasImprovedHtmlSpacing() bool`

HasImprovedHtmlSpacing returns a boolean if a field has been set.

### GetBackgroundPdf

`func (o *TemplateDefinition) GetBackgroundPdf() string`

GetBackgroundPdf returns the BackgroundPdf field if non-nil, zero value otherwise.

### GetBackgroundPdfOk

`func (o *TemplateDefinition) GetBackgroundPdfOk() (*string, bool)`

GetBackgroundPdfOk returns a tuple with the BackgroundPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundPdf

`func (o *TemplateDefinition) SetBackgroundPdf(v string)`

SetBackgroundPdf sets BackgroundPdf field to given value.

### HasBackgroundPdf

`func (o *TemplateDefinition) HasBackgroundPdf() bool`

HasBackgroundPdf returns a boolean if a field has been set.

### SetBackgroundPdfNil

`func (o *TemplateDefinition) SetBackgroundPdfNil(b bool)`

 SetBackgroundPdfNil sets the value for BackgroundPdf to be an explicit nil

### UnsetBackgroundPdf
`func (o *TemplateDefinition) UnsetBackgroundPdf()`

UnsetBackgroundPdf ensures that no value is present for BackgroundPdf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


