# DisplayModeParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zoom** | Pointer to **string** | Initial zoom * &#x60;fullpage&#x60; - Fits the whole page in the window * &#x60;fullwidth&#x60; - Fits the page width to the window * &#x60;real&#x60; - Actual size (100%) * &#x60;default&#x60; - Leaves the zoom to the viewer. The PDF gets no open action, which some e-invoice validators require  | [optional] [default to "fullwidth"]
**Layout** | Pointer to **string** | Page layout * &#x60;SinglePage&#x60; - One page at a time * &#x60;OneColumn&#x60; - Pages in one continuous column * &#x60;TwoColumnLeft&#x60; - Pages in two columns, odd-numbered pages on the left * &#x60;TwoColumnRight&#x60; - Pages in two columns, odd-numbered pages on the right  | [optional] [default to "SinglePage"]
**Mode** | Pointer to **string** | Panel shown when the document opens * &#x60;UseNone&#x60; - No panel * &#x60;UseOutlines&#x60; - Document outline (bookmarks) * &#x60;UseThumbs&#x60; - Page thumbnails * &#x60;FullScreen&#x60; - Full-screen mode, without menu bar or window controls  | [optional] [default to "UseNone"]

## Methods

### NewDisplayModeParam

`func NewDisplayModeParam() *DisplayModeParam`

NewDisplayModeParam instantiates a new DisplayModeParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisplayModeParamWithDefaults

`func NewDisplayModeParamWithDefaults() *DisplayModeParam`

NewDisplayModeParamWithDefaults instantiates a new DisplayModeParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoom

`func (o *DisplayModeParam) GetZoom() string`

GetZoom returns the Zoom field if non-nil, zero value otherwise.

### GetZoomOk

`func (o *DisplayModeParam) GetZoomOk() (*string, bool)`

GetZoomOk returns a tuple with the Zoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoom

`func (o *DisplayModeParam) SetZoom(v string)`

SetZoom sets Zoom field to given value.

### HasZoom

`func (o *DisplayModeParam) HasZoom() bool`

HasZoom returns a boolean if a field has been set.

### GetLayout

`func (o *DisplayModeParam) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *DisplayModeParam) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *DisplayModeParam) SetLayout(v string)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *DisplayModeParam) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetMode

`func (o *DisplayModeParam) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DisplayModeParam) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DisplayModeParam) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DisplayModeParam) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


