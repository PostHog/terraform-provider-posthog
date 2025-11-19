# TextReprOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxLength** | Pointer to **int32** | Maximum length of generated text (default: 4000000) | [optional] 
**Truncated** | Pointer to **bool** | Use truncation for long content within events (default: true) | [optional] 
**TruncateBuffer** | Pointer to **int32** | Characters to show at start/end when truncating (default: 1000) | [optional] 
**IncludeMarkers** | Pointer to **bool** | Use interactive markers for frontend vs plain text for backend/LLM (default: true) | [optional] 
**Collapsed** | Pointer to **bool** | Show summary vs full tree hierarchy for traces (default: false) | [optional] 
**IncludeMetadata** | Pointer to **bool** | Include metadata in response | [optional] 
**IncludeHierarchy** | Pointer to **bool** | Include hierarchy information (for traces) | [optional] 
**MaxDepth** | Pointer to **int32** | Maximum depth for hierarchical rendering | [optional] 
**ToolsCollapseThreshold** | Pointer to **int32** | Number of tools before collapsing the list (default: 5) | [optional] 
**IncludeLineNumbers** | Pointer to **bool** | Prefix each line with line number (default: false) | [optional] 

## Methods

### NewTextReprOptions

`func NewTextReprOptions() *TextReprOptions`

NewTextReprOptions instantiates a new TextReprOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextReprOptionsWithDefaults

`func NewTextReprOptionsWithDefaults() *TextReprOptions`

NewTextReprOptionsWithDefaults instantiates a new TextReprOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxLength

`func (o *TextReprOptions) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *TextReprOptions) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *TextReprOptions) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *TextReprOptions) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetTruncated

`func (o *TextReprOptions) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *TextReprOptions) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *TextReprOptions) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *TextReprOptions) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetTruncateBuffer

`func (o *TextReprOptions) GetTruncateBuffer() int32`

GetTruncateBuffer returns the TruncateBuffer field if non-nil, zero value otherwise.

### GetTruncateBufferOk

`func (o *TextReprOptions) GetTruncateBufferOk() (*int32, bool)`

GetTruncateBufferOk returns a tuple with the TruncateBuffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateBuffer

`func (o *TextReprOptions) SetTruncateBuffer(v int32)`

SetTruncateBuffer sets TruncateBuffer field to given value.

### HasTruncateBuffer

`func (o *TextReprOptions) HasTruncateBuffer() bool`

HasTruncateBuffer returns a boolean if a field has been set.

### GetIncludeMarkers

`func (o *TextReprOptions) GetIncludeMarkers() bool`

GetIncludeMarkers returns the IncludeMarkers field if non-nil, zero value otherwise.

### GetIncludeMarkersOk

`func (o *TextReprOptions) GetIncludeMarkersOk() (*bool, bool)`

GetIncludeMarkersOk returns a tuple with the IncludeMarkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMarkers

`func (o *TextReprOptions) SetIncludeMarkers(v bool)`

SetIncludeMarkers sets IncludeMarkers field to given value.

### HasIncludeMarkers

`func (o *TextReprOptions) HasIncludeMarkers() bool`

HasIncludeMarkers returns a boolean if a field has been set.

### GetCollapsed

`func (o *TextReprOptions) GetCollapsed() bool`

GetCollapsed returns the Collapsed field if non-nil, zero value otherwise.

### GetCollapsedOk

`func (o *TextReprOptions) GetCollapsedOk() (*bool, bool)`

GetCollapsedOk returns a tuple with the Collapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapsed

`func (o *TextReprOptions) SetCollapsed(v bool)`

SetCollapsed sets Collapsed field to given value.

### HasCollapsed

`func (o *TextReprOptions) HasCollapsed() bool`

HasCollapsed returns a boolean if a field has been set.

### GetIncludeMetadata

`func (o *TextReprOptions) GetIncludeMetadata() bool`

GetIncludeMetadata returns the IncludeMetadata field if non-nil, zero value otherwise.

### GetIncludeMetadataOk

`func (o *TextReprOptions) GetIncludeMetadataOk() (*bool, bool)`

GetIncludeMetadataOk returns a tuple with the IncludeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMetadata

`func (o *TextReprOptions) SetIncludeMetadata(v bool)`

SetIncludeMetadata sets IncludeMetadata field to given value.

### HasIncludeMetadata

`func (o *TextReprOptions) HasIncludeMetadata() bool`

HasIncludeMetadata returns a boolean if a field has been set.

### GetIncludeHierarchy

`func (o *TextReprOptions) GetIncludeHierarchy() bool`

GetIncludeHierarchy returns the IncludeHierarchy field if non-nil, zero value otherwise.

### GetIncludeHierarchyOk

`func (o *TextReprOptions) GetIncludeHierarchyOk() (*bool, bool)`

GetIncludeHierarchyOk returns a tuple with the IncludeHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHierarchy

`func (o *TextReprOptions) SetIncludeHierarchy(v bool)`

SetIncludeHierarchy sets IncludeHierarchy field to given value.

### HasIncludeHierarchy

`func (o *TextReprOptions) HasIncludeHierarchy() bool`

HasIncludeHierarchy returns a boolean if a field has been set.

### GetMaxDepth

`func (o *TextReprOptions) GetMaxDepth() int32`

GetMaxDepth returns the MaxDepth field if non-nil, zero value otherwise.

### GetMaxDepthOk

`func (o *TextReprOptions) GetMaxDepthOk() (*int32, bool)`

GetMaxDepthOk returns a tuple with the MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDepth

`func (o *TextReprOptions) SetMaxDepth(v int32)`

SetMaxDepth sets MaxDepth field to given value.

### HasMaxDepth

`func (o *TextReprOptions) HasMaxDepth() bool`

HasMaxDepth returns a boolean if a field has been set.

### GetToolsCollapseThreshold

`func (o *TextReprOptions) GetToolsCollapseThreshold() int32`

GetToolsCollapseThreshold returns the ToolsCollapseThreshold field if non-nil, zero value otherwise.

### GetToolsCollapseThresholdOk

`func (o *TextReprOptions) GetToolsCollapseThresholdOk() (*int32, bool)`

GetToolsCollapseThresholdOk returns a tuple with the ToolsCollapseThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsCollapseThreshold

`func (o *TextReprOptions) SetToolsCollapseThreshold(v int32)`

SetToolsCollapseThreshold sets ToolsCollapseThreshold field to given value.

### HasToolsCollapseThreshold

`func (o *TextReprOptions) HasToolsCollapseThreshold() bool`

HasToolsCollapseThreshold returns a boolean if a field has been set.

### GetIncludeLineNumbers

`func (o *TextReprOptions) GetIncludeLineNumbers() bool`

GetIncludeLineNumbers returns the IncludeLineNumbers field if non-nil, zero value otherwise.

### GetIncludeLineNumbersOk

`func (o *TextReprOptions) GetIncludeLineNumbersOk() (*bool, bool)`

GetIncludeLineNumbersOk returns a tuple with the IncludeLineNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLineNumbers

`func (o *TextReprOptions) SetIncludeLineNumbers(v bool)`

SetIncludeLineNumbers sets IncludeLineNumbers field to given value.

### HasIncludeLineNumbers

`func (o *TextReprOptions) HasIncludeLineNumbers() bool`

HasIncludeLineNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


