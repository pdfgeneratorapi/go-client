# ConvertURL2PDFRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Public URL | [optional] 
**PaperSize** | Pointer to **string** | PDF page size | [optional] [default to "a4"]
**Orientation** | Pointer to **string** |  | [optional] [default to "portrait"]
**Output** | Pointer to **string** | Output | [optional] [default to "base64"]
**Filename** | Pointer to **string** | Document name | [optional] 
**WaitForCss** | Pointer to **string** | CSS selector of an element to wait for before the page is printed; see &#x60;wait_for_css_state&#x60; | [optional] 
**WaitForCssState** | Pointer to **string** | How &#x60;wait_for_css&#x60; matches. &#x60;present&#x60; waits until the element exists, even if it is empty or hidden (e.g. a &#x60;&lt;div id&#x3D;\&quot;report-ready\&quot;&gt;&lt;/div&gt;&#x60; marker); &#x60;visible&#x60; waits until it is also visible with a non-empty size (e.g. a chart that appears when data has loaded). | [optional] [default to "present"]
**WaitForJs** | Pointer to **string** | JavaScript expression evaluated in the page every 100ms; the page is printed once it is truthy. Promises are awaited, and an expression that throws counts as not ready yet. | [optional] 
**WaitTimeout** | Pointer to **int32** | Seconds to wait for &#x60;wait_for_css&#x60; / &#x60;wait_for_js&#x60; before applying &#x60;on_wait_timeout&#x60; | [optional] [default to 30]
**OnWaitTimeout** | Pointer to **string** | What to do when the wait signal is not met within &#x60;wait_timeout&#x60;. &#x60;error&#x60; returns a 422 response; &#x60;print&#x60; prints the page as it is. | [optional] [default to "error"]

## Methods

### NewConvertURL2PDFRequest

`func NewConvertURL2PDFRequest() *ConvertURL2PDFRequest`

NewConvertURL2PDFRequest instantiates a new ConvertURL2PDFRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertURL2PDFRequestWithDefaults

`func NewConvertURL2PDFRequestWithDefaults() *ConvertURL2PDFRequest`

NewConvertURL2PDFRequestWithDefaults instantiates a new ConvertURL2PDFRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ConvertURL2PDFRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConvertURL2PDFRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConvertURL2PDFRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConvertURL2PDFRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPaperSize

`func (o *ConvertURL2PDFRequest) GetPaperSize() string`

GetPaperSize returns the PaperSize field if non-nil, zero value otherwise.

### GetPaperSizeOk

`func (o *ConvertURL2PDFRequest) GetPaperSizeOk() (*string, bool)`

GetPaperSizeOk returns a tuple with the PaperSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaperSize

`func (o *ConvertURL2PDFRequest) SetPaperSize(v string)`

SetPaperSize sets PaperSize field to given value.

### HasPaperSize

`func (o *ConvertURL2PDFRequest) HasPaperSize() bool`

HasPaperSize returns a boolean if a field has been set.

### GetOrientation

`func (o *ConvertURL2PDFRequest) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *ConvertURL2PDFRequest) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *ConvertURL2PDFRequest) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *ConvertURL2PDFRequest) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetOutput

`func (o *ConvertURL2PDFRequest) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ConvertURL2PDFRequest) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ConvertURL2PDFRequest) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ConvertURL2PDFRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetFilename

`func (o *ConvertURL2PDFRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ConvertURL2PDFRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ConvertURL2PDFRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ConvertURL2PDFRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetWaitForCss

`func (o *ConvertURL2PDFRequest) GetWaitForCss() string`

GetWaitForCss returns the WaitForCss field if non-nil, zero value otherwise.

### GetWaitForCssOk

`func (o *ConvertURL2PDFRequest) GetWaitForCssOk() (*string, bool)`

GetWaitForCssOk returns a tuple with the WaitForCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForCss

`func (o *ConvertURL2PDFRequest) SetWaitForCss(v string)`

SetWaitForCss sets WaitForCss field to given value.

### HasWaitForCss

`func (o *ConvertURL2PDFRequest) HasWaitForCss() bool`

HasWaitForCss returns a boolean if a field has been set.

### GetWaitForCssState

`func (o *ConvertURL2PDFRequest) GetWaitForCssState() string`

GetWaitForCssState returns the WaitForCssState field if non-nil, zero value otherwise.

### GetWaitForCssStateOk

`func (o *ConvertURL2PDFRequest) GetWaitForCssStateOk() (*string, bool)`

GetWaitForCssStateOk returns a tuple with the WaitForCssState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForCssState

`func (o *ConvertURL2PDFRequest) SetWaitForCssState(v string)`

SetWaitForCssState sets WaitForCssState field to given value.

### HasWaitForCssState

`func (o *ConvertURL2PDFRequest) HasWaitForCssState() bool`

HasWaitForCssState returns a boolean if a field has been set.

### GetWaitForJs

`func (o *ConvertURL2PDFRequest) GetWaitForJs() string`

GetWaitForJs returns the WaitForJs field if non-nil, zero value otherwise.

### GetWaitForJsOk

`func (o *ConvertURL2PDFRequest) GetWaitForJsOk() (*string, bool)`

GetWaitForJsOk returns a tuple with the WaitForJs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForJs

`func (o *ConvertURL2PDFRequest) SetWaitForJs(v string)`

SetWaitForJs sets WaitForJs field to given value.

### HasWaitForJs

`func (o *ConvertURL2PDFRequest) HasWaitForJs() bool`

HasWaitForJs returns a boolean if a field has been set.

### GetWaitTimeout

`func (o *ConvertURL2PDFRequest) GetWaitTimeout() int32`

GetWaitTimeout returns the WaitTimeout field if non-nil, zero value otherwise.

### GetWaitTimeoutOk

`func (o *ConvertURL2PDFRequest) GetWaitTimeoutOk() (*int32, bool)`

GetWaitTimeoutOk returns a tuple with the WaitTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimeout

`func (o *ConvertURL2PDFRequest) SetWaitTimeout(v int32)`

SetWaitTimeout sets WaitTimeout field to given value.

### HasWaitTimeout

`func (o *ConvertURL2PDFRequest) HasWaitTimeout() bool`

HasWaitTimeout returns a boolean if a field has been set.

### GetOnWaitTimeout

`func (o *ConvertURL2PDFRequest) GetOnWaitTimeout() string`

GetOnWaitTimeout returns the OnWaitTimeout field if non-nil, zero value otherwise.

### GetOnWaitTimeoutOk

`func (o *ConvertURL2PDFRequest) GetOnWaitTimeoutOk() (*string, bool)`

GetOnWaitTimeoutOk returns a tuple with the OnWaitTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnWaitTimeout

`func (o *ConvertURL2PDFRequest) SetOnWaitTimeout(v string)`

SetOnWaitTimeout sets OnWaitTimeout field to given value.

### HasOnWaitTimeout

`func (o *ConvertURL2PDFRequest) HasOnWaitTimeout() bool`

HasOnWaitTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


