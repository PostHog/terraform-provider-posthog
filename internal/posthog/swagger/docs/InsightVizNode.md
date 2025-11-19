# InsightVizNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to **NullableBool** | Query is embedded inside another bordered component | [optional] 
**Full** | Pointer to **NullableBool** | Show with most visual options enabled. Used in insight scene. | [optional] 
**HidePersonsModal** | Pointer to **NullableBool** |  | [optional] 
**HideTooltipOnScroll** | Pointer to **NullableBool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "InsightVizNode"]
**ShowCorrelationTable** | Pointer to **NullableBool** |  | [optional] 
**ShowFilters** | Pointer to **NullableBool** |  | [optional] 
**ShowHeader** | Pointer to **NullableBool** |  | [optional] 
**ShowLastComputation** | Pointer to **NullableBool** |  | [optional] 
**ShowLastComputationRefresh** | Pointer to **NullableBool** |  | [optional] 
**ShowResults** | Pointer to **NullableBool** |  | [optional] 
**ShowTable** | Pointer to **NullableBool** |  | [optional] 
**Source** | [**Source3**](Source3.md) |  | 
**SuppressSessionAnalysisWarning** | Pointer to **NullableBool** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 
**VizSpecificOptions** | Pointer to [**VizSpecificOptions**](VizSpecificOptions.md) |  | [optional] 

## Methods

### NewInsightVizNode

`func NewInsightVizNode(source Source3, ) *InsightVizNode`

NewInsightVizNode instantiates a new InsightVizNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightVizNodeWithDefaults

`func NewInsightVizNodeWithDefaults() *InsightVizNode`

NewInsightVizNodeWithDefaults instantiates a new InsightVizNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *InsightVizNode) GetEmbedded() bool`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *InsightVizNode) GetEmbeddedOk() (*bool, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *InsightVizNode) SetEmbedded(v bool)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *InsightVizNode) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### SetEmbeddedNil

`func (o *InsightVizNode) SetEmbeddedNil(b bool)`

 SetEmbeddedNil sets the value for Embedded to be an explicit nil

### UnsetEmbedded
`func (o *InsightVizNode) UnsetEmbedded()`

UnsetEmbedded ensures that no value is present for Embedded, not even an explicit nil
### GetFull

`func (o *InsightVizNode) GetFull() bool`

GetFull returns the Full field if non-nil, zero value otherwise.

### GetFullOk

`func (o *InsightVizNode) GetFullOk() (*bool, bool)`

GetFullOk returns a tuple with the Full field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFull

`func (o *InsightVizNode) SetFull(v bool)`

SetFull sets Full field to given value.

### HasFull

`func (o *InsightVizNode) HasFull() bool`

HasFull returns a boolean if a field has been set.

### SetFullNil

`func (o *InsightVizNode) SetFullNil(b bool)`

 SetFullNil sets the value for Full to be an explicit nil

### UnsetFull
`func (o *InsightVizNode) UnsetFull()`

UnsetFull ensures that no value is present for Full, not even an explicit nil
### GetHidePersonsModal

`func (o *InsightVizNode) GetHidePersonsModal() bool`

GetHidePersonsModal returns the HidePersonsModal field if non-nil, zero value otherwise.

### GetHidePersonsModalOk

`func (o *InsightVizNode) GetHidePersonsModalOk() (*bool, bool)`

GetHidePersonsModalOk returns a tuple with the HidePersonsModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePersonsModal

`func (o *InsightVizNode) SetHidePersonsModal(v bool)`

SetHidePersonsModal sets HidePersonsModal field to given value.

### HasHidePersonsModal

`func (o *InsightVizNode) HasHidePersonsModal() bool`

HasHidePersonsModal returns a boolean if a field has been set.

### SetHidePersonsModalNil

`func (o *InsightVizNode) SetHidePersonsModalNil(b bool)`

 SetHidePersonsModalNil sets the value for HidePersonsModal to be an explicit nil

### UnsetHidePersonsModal
`func (o *InsightVizNode) UnsetHidePersonsModal()`

UnsetHidePersonsModal ensures that no value is present for HidePersonsModal, not even an explicit nil
### GetHideTooltipOnScroll

`func (o *InsightVizNode) GetHideTooltipOnScroll() bool`

GetHideTooltipOnScroll returns the HideTooltipOnScroll field if non-nil, zero value otherwise.

### GetHideTooltipOnScrollOk

`func (o *InsightVizNode) GetHideTooltipOnScrollOk() (*bool, bool)`

GetHideTooltipOnScrollOk returns a tuple with the HideTooltipOnScroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTooltipOnScroll

`func (o *InsightVizNode) SetHideTooltipOnScroll(v bool)`

SetHideTooltipOnScroll sets HideTooltipOnScroll field to given value.

### HasHideTooltipOnScroll

`func (o *InsightVizNode) HasHideTooltipOnScroll() bool`

HasHideTooltipOnScroll returns a boolean if a field has been set.

### SetHideTooltipOnScrollNil

`func (o *InsightVizNode) SetHideTooltipOnScrollNil(b bool)`

 SetHideTooltipOnScrollNil sets the value for HideTooltipOnScroll to be an explicit nil

### UnsetHideTooltipOnScroll
`func (o *InsightVizNode) UnsetHideTooltipOnScroll()`

UnsetHideTooltipOnScroll ensures that no value is present for HideTooltipOnScroll, not even an explicit nil
### GetKind

`func (o *InsightVizNode) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InsightVizNode) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InsightVizNode) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *InsightVizNode) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetShowCorrelationTable

`func (o *InsightVizNode) GetShowCorrelationTable() bool`

GetShowCorrelationTable returns the ShowCorrelationTable field if non-nil, zero value otherwise.

### GetShowCorrelationTableOk

`func (o *InsightVizNode) GetShowCorrelationTableOk() (*bool, bool)`

GetShowCorrelationTableOk returns a tuple with the ShowCorrelationTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCorrelationTable

`func (o *InsightVizNode) SetShowCorrelationTable(v bool)`

SetShowCorrelationTable sets ShowCorrelationTable field to given value.

### HasShowCorrelationTable

`func (o *InsightVizNode) HasShowCorrelationTable() bool`

HasShowCorrelationTable returns a boolean if a field has been set.

### SetShowCorrelationTableNil

`func (o *InsightVizNode) SetShowCorrelationTableNil(b bool)`

 SetShowCorrelationTableNil sets the value for ShowCorrelationTable to be an explicit nil

### UnsetShowCorrelationTable
`func (o *InsightVizNode) UnsetShowCorrelationTable()`

UnsetShowCorrelationTable ensures that no value is present for ShowCorrelationTable, not even an explicit nil
### GetShowFilters

`func (o *InsightVizNode) GetShowFilters() bool`

GetShowFilters returns the ShowFilters field if non-nil, zero value otherwise.

### GetShowFiltersOk

`func (o *InsightVizNode) GetShowFiltersOk() (*bool, bool)`

GetShowFiltersOk returns a tuple with the ShowFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFilters

`func (o *InsightVizNode) SetShowFilters(v bool)`

SetShowFilters sets ShowFilters field to given value.

### HasShowFilters

`func (o *InsightVizNode) HasShowFilters() bool`

HasShowFilters returns a boolean if a field has been set.

### SetShowFiltersNil

`func (o *InsightVizNode) SetShowFiltersNil(b bool)`

 SetShowFiltersNil sets the value for ShowFilters to be an explicit nil

### UnsetShowFilters
`func (o *InsightVizNode) UnsetShowFilters()`

UnsetShowFilters ensures that no value is present for ShowFilters, not even an explicit nil
### GetShowHeader

`func (o *InsightVizNode) GetShowHeader() bool`

GetShowHeader returns the ShowHeader field if non-nil, zero value otherwise.

### GetShowHeaderOk

`func (o *InsightVizNode) GetShowHeaderOk() (*bool, bool)`

GetShowHeaderOk returns a tuple with the ShowHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHeader

`func (o *InsightVizNode) SetShowHeader(v bool)`

SetShowHeader sets ShowHeader field to given value.

### HasShowHeader

`func (o *InsightVizNode) HasShowHeader() bool`

HasShowHeader returns a boolean if a field has been set.

### SetShowHeaderNil

`func (o *InsightVizNode) SetShowHeaderNil(b bool)`

 SetShowHeaderNil sets the value for ShowHeader to be an explicit nil

### UnsetShowHeader
`func (o *InsightVizNode) UnsetShowHeader()`

UnsetShowHeader ensures that no value is present for ShowHeader, not even an explicit nil
### GetShowLastComputation

`func (o *InsightVizNode) GetShowLastComputation() bool`

GetShowLastComputation returns the ShowLastComputation field if non-nil, zero value otherwise.

### GetShowLastComputationOk

`func (o *InsightVizNode) GetShowLastComputationOk() (*bool, bool)`

GetShowLastComputationOk returns a tuple with the ShowLastComputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLastComputation

`func (o *InsightVizNode) SetShowLastComputation(v bool)`

SetShowLastComputation sets ShowLastComputation field to given value.

### HasShowLastComputation

`func (o *InsightVizNode) HasShowLastComputation() bool`

HasShowLastComputation returns a boolean if a field has been set.

### SetShowLastComputationNil

`func (o *InsightVizNode) SetShowLastComputationNil(b bool)`

 SetShowLastComputationNil sets the value for ShowLastComputation to be an explicit nil

### UnsetShowLastComputation
`func (o *InsightVizNode) UnsetShowLastComputation()`

UnsetShowLastComputation ensures that no value is present for ShowLastComputation, not even an explicit nil
### GetShowLastComputationRefresh

`func (o *InsightVizNode) GetShowLastComputationRefresh() bool`

GetShowLastComputationRefresh returns the ShowLastComputationRefresh field if non-nil, zero value otherwise.

### GetShowLastComputationRefreshOk

`func (o *InsightVizNode) GetShowLastComputationRefreshOk() (*bool, bool)`

GetShowLastComputationRefreshOk returns a tuple with the ShowLastComputationRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLastComputationRefresh

`func (o *InsightVizNode) SetShowLastComputationRefresh(v bool)`

SetShowLastComputationRefresh sets ShowLastComputationRefresh field to given value.

### HasShowLastComputationRefresh

`func (o *InsightVizNode) HasShowLastComputationRefresh() bool`

HasShowLastComputationRefresh returns a boolean if a field has been set.

### SetShowLastComputationRefreshNil

`func (o *InsightVizNode) SetShowLastComputationRefreshNil(b bool)`

 SetShowLastComputationRefreshNil sets the value for ShowLastComputationRefresh to be an explicit nil

### UnsetShowLastComputationRefresh
`func (o *InsightVizNode) UnsetShowLastComputationRefresh()`

UnsetShowLastComputationRefresh ensures that no value is present for ShowLastComputationRefresh, not even an explicit nil
### GetShowResults

`func (o *InsightVizNode) GetShowResults() bool`

GetShowResults returns the ShowResults field if non-nil, zero value otherwise.

### GetShowResultsOk

`func (o *InsightVizNode) GetShowResultsOk() (*bool, bool)`

GetShowResultsOk returns a tuple with the ShowResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowResults

`func (o *InsightVizNode) SetShowResults(v bool)`

SetShowResults sets ShowResults field to given value.

### HasShowResults

`func (o *InsightVizNode) HasShowResults() bool`

HasShowResults returns a boolean if a field has been set.

### SetShowResultsNil

`func (o *InsightVizNode) SetShowResultsNil(b bool)`

 SetShowResultsNil sets the value for ShowResults to be an explicit nil

### UnsetShowResults
`func (o *InsightVizNode) UnsetShowResults()`

UnsetShowResults ensures that no value is present for ShowResults, not even an explicit nil
### GetShowTable

`func (o *InsightVizNode) GetShowTable() bool`

GetShowTable returns the ShowTable field if non-nil, zero value otherwise.

### GetShowTableOk

`func (o *InsightVizNode) GetShowTableOk() (*bool, bool)`

GetShowTableOk returns a tuple with the ShowTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTable

`func (o *InsightVizNode) SetShowTable(v bool)`

SetShowTable sets ShowTable field to given value.

### HasShowTable

`func (o *InsightVizNode) HasShowTable() bool`

HasShowTable returns a boolean if a field has been set.

### SetShowTableNil

`func (o *InsightVizNode) SetShowTableNil(b bool)`

 SetShowTableNil sets the value for ShowTable to be an explicit nil

### UnsetShowTable
`func (o *InsightVizNode) UnsetShowTable()`

UnsetShowTable ensures that no value is present for ShowTable, not even an explicit nil
### GetSource

`func (o *InsightVizNode) GetSource() Source3`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InsightVizNode) GetSourceOk() (*Source3, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InsightVizNode) SetSource(v Source3)`

SetSource sets Source field to given value.


### GetSuppressSessionAnalysisWarning

`func (o *InsightVizNode) GetSuppressSessionAnalysisWarning() bool`

GetSuppressSessionAnalysisWarning returns the SuppressSessionAnalysisWarning field if non-nil, zero value otherwise.

### GetSuppressSessionAnalysisWarningOk

`func (o *InsightVizNode) GetSuppressSessionAnalysisWarningOk() (*bool, bool)`

GetSuppressSessionAnalysisWarningOk returns a tuple with the SuppressSessionAnalysisWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressSessionAnalysisWarning

`func (o *InsightVizNode) SetSuppressSessionAnalysisWarning(v bool)`

SetSuppressSessionAnalysisWarning sets SuppressSessionAnalysisWarning field to given value.

### HasSuppressSessionAnalysisWarning

`func (o *InsightVizNode) HasSuppressSessionAnalysisWarning() bool`

HasSuppressSessionAnalysisWarning returns a boolean if a field has been set.

### SetSuppressSessionAnalysisWarningNil

`func (o *InsightVizNode) SetSuppressSessionAnalysisWarningNil(b bool)`

 SetSuppressSessionAnalysisWarningNil sets the value for SuppressSessionAnalysisWarning to be an explicit nil

### UnsetSuppressSessionAnalysisWarning
`func (o *InsightVizNode) UnsetSuppressSessionAnalysisWarning()`

UnsetSuppressSessionAnalysisWarning ensures that no value is present for SuppressSessionAnalysisWarning, not even an explicit nil
### GetVersion

`func (o *InsightVizNode) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InsightVizNode) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InsightVizNode) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InsightVizNode) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *InsightVizNode) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *InsightVizNode) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetVizSpecificOptions

`func (o *InsightVizNode) GetVizSpecificOptions() VizSpecificOptions`

GetVizSpecificOptions returns the VizSpecificOptions field if non-nil, zero value otherwise.

### GetVizSpecificOptionsOk

`func (o *InsightVizNode) GetVizSpecificOptionsOk() (*VizSpecificOptions, bool)`

GetVizSpecificOptionsOk returns a tuple with the VizSpecificOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizSpecificOptions

`func (o *InsightVizNode) SetVizSpecificOptions(v VizSpecificOptions)`

SetVizSpecificOptions sets VizSpecificOptions field to given value.

### HasVizSpecificOptions

`func (o *InsightVizNode) HasVizSpecificOptions() bool`

HasVizSpecificOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


