# SavedInsightNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowSorting** | Pointer to **NullableBool** | Can the user click on column headers to sort the table? (default: true) | [optional] 
**Context** | Pointer to [**DataTableNodeViewPropsContext**](DataTableNodeViewPropsContext.md) |  | [optional] 
**DefaultColumns** | Pointer to **[]string** | Default columns to use when resetting column configuration | [optional] 
**Embedded** | Pointer to **NullableBool** | Query is embedded inside another bordered component | [optional] 
**Expandable** | Pointer to **NullableBool** | Can expand row to show raw event data (default: true) | [optional] 
**Full** | Pointer to **NullableBool** | Show with most visual options enabled. Used in insight scene. | [optional] 
**HidePersonsModal** | Pointer to **NullableBool** |  | [optional] 
**HideTooltipOnScroll** | Pointer to **NullableBool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "SavedInsightNode"]
**PropertiesViaUrl** | Pointer to **NullableBool** | Link properties via the URL (default: false) | [optional] 
**ShortId** | **string** |  | 
**ShowActions** | Pointer to **NullableBool** | Show the kebab menu at the end of the row | [optional] 
**ShowColumnConfigurator** | Pointer to **NullableBool** | Show a button to configure the table&#39;s columns if possible | [optional] 
**ShowCorrelationTable** | Pointer to **NullableBool** |  | [optional] 
**ShowDateRange** | Pointer to **NullableBool** | Show date range selector | [optional] 
**ShowElapsedTime** | Pointer to **NullableBool** | Show the time it takes to run a query | [optional] 
**ShowEventFilter** | Pointer to **NullableBool** | Include an event filter above the table (EventsNode only) | [optional] 
**ShowExport** | Pointer to **NullableBool** | Show the export button | [optional] 
**ShowFilters** | Pointer to **NullableBool** |  | [optional] 
**ShowHeader** | Pointer to **NullableBool** |  | [optional] 
**ShowHogQLEditor** | Pointer to **NullableBool** | Include a HogQL query editor above HogQL tables | [optional] 
**ShowLastComputation** | Pointer to **NullableBool** |  | [optional] 
**ShowLastComputationRefresh** | Pointer to **NullableBool** |  | [optional] 
**ShowOpenEditorButton** | Pointer to **NullableBool** | Show a button to open the current query as a new insight. (default: true) | [optional] 
**ShowPersistentColumnConfigurator** | Pointer to **NullableBool** | Show a button to configure and persist the table&#39;s default columns if possible | [optional] 
**ShowPropertyFilter** | Pointer to [**NullableShowpropertyfilter**](Showpropertyfilter.md) |  | [optional] [default to null]
**ShowReload** | Pointer to **NullableBool** | Show a reload button | [optional] 
**ShowResults** | Pointer to **NullableBool** |  | [optional] 
**ShowResultsTable** | Pointer to **NullableBool** | Show a results table | [optional] 
**ShowSavedFilters** | Pointer to **NullableBool** | Show saved filters feature for this table (requires uniqueKey) | [optional] 
**ShowSavedQueries** | Pointer to **NullableBool** | Shows a list of saved queries | [optional] 
**ShowSearch** | Pointer to **NullableBool** | Include a free text search field (PersonsNode only) | [optional] 
**ShowTable** | Pointer to **NullableBool** |  | [optional] 
**ShowTestAccountFilters** | Pointer to **NullableBool** | Show filter to exclude test accounts | [optional] 
**ShowTimings** | Pointer to **NullableBool** | Show a detailed query timing breakdown | [optional] 
**SuppressSessionAnalysisWarning** | Pointer to **NullableBool** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 
**VizSpecificOptions** | Pointer to [**VizSpecificOptions**](VizSpecificOptions.md) |  | [optional] 

## Methods

### NewSavedInsightNode

`func NewSavedInsightNode(shortId string, ) *SavedInsightNode`

NewSavedInsightNode instantiates a new SavedInsightNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedInsightNodeWithDefaults

`func NewSavedInsightNodeWithDefaults() *SavedInsightNode`

NewSavedInsightNodeWithDefaults instantiates a new SavedInsightNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowSorting

`func (o *SavedInsightNode) GetAllowSorting() bool`

GetAllowSorting returns the AllowSorting field if non-nil, zero value otherwise.

### GetAllowSortingOk

`func (o *SavedInsightNode) GetAllowSortingOk() (*bool, bool)`

GetAllowSortingOk returns a tuple with the AllowSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSorting

`func (o *SavedInsightNode) SetAllowSorting(v bool)`

SetAllowSorting sets AllowSorting field to given value.

### HasAllowSorting

`func (o *SavedInsightNode) HasAllowSorting() bool`

HasAllowSorting returns a boolean if a field has been set.

### SetAllowSortingNil

`func (o *SavedInsightNode) SetAllowSortingNil(b bool)`

 SetAllowSortingNil sets the value for AllowSorting to be an explicit nil

### UnsetAllowSorting
`func (o *SavedInsightNode) UnsetAllowSorting()`

UnsetAllowSorting ensures that no value is present for AllowSorting, not even an explicit nil
### GetContext

`func (o *SavedInsightNode) GetContext() DataTableNodeViewPropsContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SavedInsightNode) GetContextOk() (*DataTableNodeViewPropsContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SavedInsightNode) SetContext(v DataTableNodeViewPropsContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *SavedInsightNode) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetDefaultColumns

`func (o *SavedInsightNode) GetDefaultColumns() []string`

GetDefaultColumns returns the DefaultColumns field if non-nil, zero value otherwise.

### GetDefaultColumnsOk

`func (o *SavedInsightNode) GetDefaultColumnsOk() (*[]string, bool)`

GetDefaultColumnsOk returns a tuple with the DefaultColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultColumns

`func (o *SavedInsightNode) SetDefaultColumns(v []string)`

SetDefaultColumns sets DefaultColumns field to given value.

### HasDefaultColumns

`func (o *SavedInsightNode) HasDefaultColumns() bool`

HasDefaultColumns returns a boolean if a field has been set.

### SetDefaultColumnsNil

`func (o *SavedInsightNode) SetDefaultColumnsNil(b bool)`

 SetDefaultColumnsNil sets the value for DefaultColumns to be an explicit nil

### UnsetDefaultColumns
`func (o *SavedInsightNode) UnsetDefaultColumns()`

UnsetDefaultColumns ensures that no value is present for DefaultColumns, not even an explicit nil
### GetEmbedded

`func (o *SavedInsightNode) GetEmbedded() bool`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *SavedInsightNode) GetEmbeddedOk() (*bool, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *SavedInsightNode) SetEmbedded(v bool)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *SavedInsightNode) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### SetEmbeddedNil

`func (o *SavedInsightNode) SetEmbeddedNil(b bool)`

 SetEmbeddedNil sets the value for Embedded to be an explicit nil

### UnsetEmbedded
`func (o *SavedInsightNode) UnsetEmbedded()`

UnsetEmbedded ensures that no value is present for Embedded, not even an explicit nil
### GetExpandable

`func (o *SavedInsightNode) GetExpandable() bool`

GetExpandable returns the Expandable field if non-nil, zero value otherwise.

### GetExpandableOk

`func (o *SavedInsightNode) GetExpandableOk() (*bool, bool)`

GetExpandableOk returns a tuple with the Expandable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandable

`func (o *SavedInsightNode) SetExpandable(v bool)`

SetExpandable sets Expandable field to given value.

### HasExpandable

`func (o *SavedInsightNode) HasExpandable() bool`

HasExpandable returns a boolean if a field has been set.

### SetExpandableNil

`func (o *SavedInsightNode) SetExpandableNil(b bool)`

 SetExpandableNil sets the value for Expandable to be an explicit nil

### UnsetExpandable
`func (o *SavedInsightNode) UnsetExpandable()`

UnsetExpandable ensures that no value is present for Expandable, not even an explicit nil
### GetFull

`func (o *SavedInsightNode) GetFull() bool`

GetFull returns the Full field if non-nil, zero value otherwise.

### GetFullOk

`func (o *SavedInsightNode) GetFullOk() (*bool, bool)`

GetFullOk returns a tuple with the Full field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFull

`func (o *SavedInsightNode) SetFull(v bool)`

SetFull sets Full field to given value.

### HasFull

`func (o *SavedInsightNode) HasFull() bool`

HasFull returns a boolean if a field has been set.

### SetFullNil

`func (o *SavedInsightNode) SetFullNil(b bool)`

 SetFullNil sets the value for Full to be an explicit nil

### UnsetFull
`func (o *SavedInsightNode) UnsetFull()`

UnsetFull ensures that no value is present for Full, not even an explicit nil
### GetHidePersonsModal

`func (o *SavedInsightNode) GetHidePersonsModal() bool`

GetHidePersonsModal returns the HidePersonsModal field if non-nil, zero value otherwise.

### GetHidePersonsModalOk

`func (o *SavedInsightNode) GetHidePersonsModalOk() (*bool, bool)`

GetHidePersonsModalOk returns a tuple with the HidePersonsModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePersonsModal

`func (o *SavedInsightNode) SetHidePersonsModal(v bool)`

SetHidePersonsModal sets HidePersonsModal field to given value.

### HasHidePersonsModal

`func (o *SavedInsightNode) HasHidePersonsModal() bool`

HasHidePersonsModal returns a boolean if a field has been set.

### SetHidePersonsModalNil

`func (o *SavedInsightNode) SetHidePersonsModalNil(b bool)`

 SetHidePersonsModalNil sets the value for HidePersonsModal to be an explicit nil

### UnsetHidePersonsModal
`func (o *SavedInsightNode) UnsetHidePersonsModal()`

UnsetHidePersonsModal ensures that no value is present for HidePersonsModal, not even an explicit nil
### GetHideTooltipOnScroll

`func (o *SavedInsightNode) GetHideTooltipOnScroll() bool`

GetHideTooltipOnScroll returns the HideTooltipOnScroll field if non-nil, zero value otherwise.

### GetHideTooltipOnScrollOk

`func (o *SavedInsightNode) GetHideTooltipOnScrollOk() (*bool, bool)`

GetHideTooltipOnScrollOk returns a tuple with the HideTooltipOnScroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTooltipOnScroll

`func (o *SavedInsightNode) SetHideTooltipOnScroll(v bool)`

SetHideTooltipOnScroll sets HideTooltipOnScroll field to given value.

### HasHideTooltipOnScroll

`func (o *SavedInsightNode) HasHideTooltipOnScroll() bool`

HasHideTooltipOnScroll returns a boolean if a field has been set.

### SetHideTooltipOnScrollNil

`func (o *SavedInsightNode) SetHideTooltipOnScrollNil(b bool)`

 SetHideTooltipOnScrollNil sets the value for HideTooltipOnScroll to be an explicit nil

### UnsetHideTooltipOnScroll
`func (o *SavedInsightNode) UnsetHideTooltipOnScroll()`

UnsetHideTooltipOnScroll ensures that no value is present for HideTooltipOnScroll, not even an explicit nil
### GetKind

`func (o *SavedInsightNode) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SavedInsightNode) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SavedInsightNode) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SavedInsightNode) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPropertiesViaUrl

`func (o *SavedInsightNode) GetPropertiesViaUrl() bool`

GetPropertiesViaUrl returns the PropertiesViaUrl field if non-nil, zero value otherwise.

### GetPropertiesViaUrlOk

`func (o *SavedInsightNode) GetPropertiesViaUrlOk() (*bool, bool)`

GetPropertiesViaUrlOk returns a tuple with the PropertiesViaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesViaUrl

`func (o *SavedInsightNode) SetPropertiesViaUrl(v bool)`

SetPropertiesViaUrl sets PropertiesViaUrl field to given value.

### HasPropertiesViaUrl

`func (o *SavedInsightNode) HasPropertiesViaUrl() bool`

HasPropertiesViaUrl returns a boolean if a field has been set.

### SetPropertiesViaUrlNil

`func (o *SavedInsightNode) SetPropertiesViaUrlNil(b bool)`

 SetPropertiesViaUrlNil sets the value for PropertiesViaUrl to be an explicit nil

### UnsetPropertiesViaUrl
`func (o *SavedInsightNode) UnsetPropertiesViaUrl()`

UnsetPropertiesViaUrl ensures that no value is present for PropertiesViaUrl, not even an explicit nil
### GetShortId

`func (o *SavedInsightNode) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *SavedInsightNode) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *SavedInsightNode) SetShortId(v string)`

SetShortId sets ShortId field to given value.


### GetShowActions

`func (o *SavedInsightNode) GetShowActions() bool`

GetShowActions returns the ShowActions field if non-nil, zero value otherwise.

### GetShowActionsOk

`func (o *SavedInsightNode) GetShowActionsOk() (*bool, bool)`

GetShowActionsOk returns a tuple with the ShowActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowActions

`func (o *SavedInsightNode) SetShowActions(v bool)`

SetShowActions sets ShowActions field to given value.

### HasShowActions

`func (o *SavedInsightNode) HasShowActions() bool`

HasShowActions returns a boolean if a field has been set.

### SetShowActionsNil

`func (o *SavedInsightNode) SetShowActionsNil(b bool)`

 SetShowActionsNil sets the value for ShowActions to be an explicit nil

### UnsetShowActions
`func (o *SavedInsightNode) UnsetShowActions()`

UnsetShowActions ensures that no value is present for ShowActions, not even an explicit nil
### GetShowColumnConfigurator

`func (o *SavedInsightNode) GetShowColumnConfigurator() bool`

GetShowColumnConfigurator returns the ShowColumnConfigurator field if non-nil, zero value otherwise.

### GetShowColumnConfiguratorOk

`func (o *SavedInsightNode) GetShowColumnConfiguratorOk() (*bool, bool)`

GetShowColumnConfiguratorOk returns a tuple with the ShowColumnConfigurator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowColumnConfigurator

`func (o *SavedInsightNode) SetShowColumnConfigurator(v bool)`

SetShowColumnConfigurator sets ShowColumnConfigurator field to given value.

### HasShowColumnConfigurator

`func (o *SavedInsightNode) HasShowColumnConfigurator() bool`

HasShowColumnConfigurator returns a boolean if a field has been set.

### SetShowColumnConfiguratorNil

`func (o *SavedInsightNode) SetShowColumnConfiguratorNil(b bool)`

 SetShowColumnConfiguratorNil sets the value for ShowColumnConfigurator to be an explicit nil

### UnsetShowColumnConfigurator
`func (o *SavedInsightNode) UnsetShowColumnConfigurator()`

UnsetShowColumnConfigurator ensures that no value is present for ShowColumnConfigurator, not even an explicit nil
### GetShowCorrelationTable

`func (o *SavedInsightNode) GetShowCorrelationTable() bool`

GetShowCorrelationTable returns the ShowCorrelationTable field if non-nil, zero value otherwise.

### GetShowCorrelationTableOk

`func (o *SavedInsightNode) GetShowCorrelationTableOk() (*bool, bool)`

GetShowCorrelationTableOk returns a tuple with the ShowCorrelationTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCorrelationTable

`func (o *SavedInsightNode) SetShowCorrelationTable(v bool)`

SetShowCorrelationTable sets ShowCorrelationTable field to given value.

### HasShowCorrelationTable

`func (o *SavedInsightNode) HasShowCorrelationTable() bool`

HasShowCorrelationTable returns a boolean if a field has been set.

### SetShowCorrelationTableNil

`func (o *SavedInsightNode) SetShowCorrelationTableNil(b bool)`

 SetShowCorrelationTableNil sets the value for ShowCorrelationTable to be an explicit nil

### UnsetShowCorrelationTable
`func (o *SavedInsightNode) UnsetShowCorrelationTable()`

UnsetShowCorrelationTable ensures that no value is present for ShowCorrelationTable, not even an explicit nil
### GetShowDateRange

`func (o *SavedInsightNode) GetShowDateRange() bool`

GetShowDateRange returns the ShowDateRange field if non-nil, zero value otherwise.

### GetShowDateRangeOk

`func (o *SavedInsightNode) GetShowDateRangeOk() (*bool, bool)`

GetShowDateRangeOk returns a tuple with the ShowDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDateRange

`func (o *SavedInsightNode) SetShowDateRange(v bool)`

SetShowDateRange sets ShowDateRange field to given value.

### HasShowDateRange

`func (o *SavedInsightNode) HasShowDateRange() bool`

HasShowDateRange returns a boolean if a field has been set.

### SetShowDateRangeNil

`func (o *SavedInsightNode) SetShowDateRangeNil(b bool)`

 SetShowDateRangeNil sets the value for ShowDateRange to be an explicit nil

### UnsetShowDateRange
`func (o *SavedInsightNode) UnsetShowDateRange()`

UnsetShowDateRange ensures that no value is present for ShowDateRange, not even an explicit nil
### GetShowElapsedTime

`func (o *SavedInsightNode) GetShowElapsedTime() bool`

GetShowElapsedTime returns the ShowElapsedTime field if non-nil, zero value otherwise.

### GetShowElapsedTimeOk

`func (o *SavedInsightNode) GetShowElapsedTimeOk() (*bool, bool)`

GetShowElapsedTimeOk returns a tuple with the ShowElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowElapsedTime

`func (o *SavedInsightNode) SetShowElapsedTime(v bool)`

SetShowElapsedTime sets ShowElapsedTime field to given value.

### HasShowElapsedTime

`func (o *SavedInsightNode) HasShowElapsedTime() bool`

HasShowElapsedTime returns a boolean if a field has been set.

### SetShowElapsedTimeNil

`func (o *SavedInsightNode) SetShowElapsedTimeNil(b bool)`

 SetShowElapsedTimeNil sets the value for ShowElapsedTime to be an explicit nil

### UnsetShowElapsedTime
`func (o *SavedInsightNode) UnsetShowElapsedTime()`

UnsetShowElapsedTime ensures that no value is present for ShowElapsedTime, not even an explicit nil
### GetShowEventFilter

`func (o *SavedInsightNode) GetShowEventFilter() bool`

GetShowEventFilter returns the ShowEventFilter field if non-nil, zero value otherwise.

### GetShowEventFilterOk

`func (o *SavedInsightNode) GetShowEventFilterOk() (*bool, bool)`

GetShowEventFilterOk returns a tuple with the ShowEventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowEventFilter

`func (o *SavedInsightNode) SetShowEventFilter(v bool)`

SetShowEventFilter sets ShowEventFilter field to given value.

### HasShowEventFilter

`func (o *SavedInsightNode) HasShowEventFilter() bool`

HasShowEventFilter returns a boolean if a field has been set.

### SetShowEventFilterNil

`func (o *SavedInsightNode) SetShowEventFilterNil(b bool)`

 SetShowEventFilterNil sets the value for ShowEventFilter to be an explicit nil

### UnsetShowEventFilter
`func (o *SavedInsightNode) UnsetShowEventFilter()`

UnsetShowEventFilter ensures that no value is present for ShowEventFilter, not even an explicit nil
### GetShowExport

`func (o *SavedInsightNode) GetShowExport() bool`

GetShowExport returns the ShowExport field if non-nil, zero value otherwise.

### GetShowExportOk

`func (o *SavedInsightNode) GetShowExportOk() (*bool, bool)`

GetShowExportOk returns a tuple with the ShowExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowExport

`func (o *SavedInsightNode) SetShowExport(v bool)`

SetShowExport sets ShowExport field to given value.

### HasShowExport

`func (o *SavedInsightNode) HasShowExport() bool`

HasShowExport returns a boolean if a field has been set.

### SetShowExportNil

`func (o *SavedInsightNode) SetShowExportNil(b bool)`

 SetShowExportNil sets the value for ShowExport to be an explicit nil

### UnsetShowExport
`func (o *SavedInsightNode) UnsetShowExport()`

UnsetShowExport ensures that no value is present for ShowExport, not even an explicit nil
### GetShowFilters

`func (o *SavedInsightNode) GetShowFilters() bool`

GetShowFilters returns the ShowFilters field if non-nil, zero value otherwise.

### GetShowFiltersOk

`func (o *SavedInsightNode) GetShowFiltersOk() (*bool, bool)`

GetShowFiltersOk returns a tuple with the ShowFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFilters

`func (o *SavedInsightNode) SetShowFilters(v bool)`

SetShowFilters sets ShowFilters field to given value.

### HasShowFilters

`func (o *SavedInsightNode) HasShowFilters() bool`

HasShowFilters returns a boolean if a field has been set.

### SetShowFiltersNil

`func (o *SavedInsightNode) SetShowFiltersNil(b bool)`

 SetShowFiltersNil sets the value for ShowFilters to be an explicit nil

### UnsetShowFilters
`func (o *SavedInsightNode) UnsetShowFilters()`

UnsetShowFilters ensures that no value is present for ShowFilters, not even an explicit nil
### GetShowHeader

`func (o *SavedInsightNode) GetShowHeader() bool`

GetShowHeader returns the ShowHeader field if non-nil, zero value otherwise.

### GetShowHeaderOk

`func (o *SavedInsightNode) GetShowHeaderOk() (*bool, bool)`

GetShowHeaderOk returns a tuple with the ShowHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHeader

`func (o *SavedInsightNode) SetShowHeader(v bool)`

SetShowHeader sets ShowHeader field to given value.

### HasShowHeader

`func (o *SavedInsightNode) HasShowHeader() bool`

HasShowHeader returns a boolean if a field has been set.

### SetShowHeaderNil

`func (o *SavedInsightNode) SetShowHeaderNil(b bool)`

 SetShowHeaderNil sets the value for ShowHeader to be an explicit nil

### UnsetShowHeader
`func (o *SavedInsightNode) UnsetShowHeader()`

UnsetShowHeader ensures that no value is present for ShowHeader, not even an explicit nil
### GetShowHogQLEditor

`func (o *SavedInsightNode) GetShowHogQLEditor() bool`

GetShowHogQLEditor returns the ShowHogQLEditor field if non-nil, zero value otherwise.

### GetShowHogQLEditorOk

`func (o *SavedInsightNode) GetShowHogQLEditorOk() (*bool, bool)`

GetShowHogQLEditorOk returns a tuple with the ShowHogQLEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHogQLEditor

`func (o *SavedInsightNode) SetShowHogQLEditor(v bool)`

SetShowHogQLEditor sets ShowHogQLEditor field to given value.

### HasShowHogQLEditor

`func (o *SavedInsightNode) HasShowHogQLEditor() bool`

HasShowHogQLEditor returns a boolean if a field has been set.

### SetShowHogQLEditorNil

`func (o *SavedInsightNode) SetShowHogQLEditorNil(b bool)`

 SetShowHogQLEditorNil sets the value for ShowHogQLEditor to be an explicit nil

### UnsetShowHogQLEditor
`func (o *SavedInsightNode) UnsetShowHogQLEditor()`

UnsetShowHogQLEditor ensures that no value is present for ShowHogQLEditor, not even an explicit nil
### GetShowLastComputation

`func (o *SavedInsightNode) GetShowLastComputation() bool`

GetShowLastComputation returns the ShowLastComputation field if non-nil, zero value otherwise.

### GetShowLastComputationOk

`func (o *SavedInsightNode) GetShowLastComputationOk() (*bool, bool)`

GetShowLastComputationOk returns a tuple with the ShowLastComputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLastComputation

`func (o *SavedInsightNode) SetShowLastComputation(v bool)`

SetShowLastComputation sets ShowLastComputation field to given value.

### HasShowLastComputation

`func (o *SavedInsightNode) HasShowLastComputation() bool`

HasShowLastComputation returns a boolean if a field has been set.

### SetShowLastComputationNil

`func (o *SavedInsightNode) SetShowLastComputationNil(b bool)`

 SetShowLastComputationNil sets the value for ShowLastComputation to be an explicit nil

### UnsetShowLastComputation
`func (o *SavedInsightNode) UnsetShowLastComputation()`

UnsetShowLastComputation ensures that no value is present for ShowLastComputation, not even an explicit nil
### GetShowLastComputationRefresh

`func (o *SavedInsightNode) GetShowLastComputationRefresh() bool`

GetShowLastComputationRefresh returns the ShowLastComputationRefresh field if non-nil, zero value otherwise.

### GetShowLastComputationRefreshOk

`func (o *SavedInsightNode) GetShowLastComputationRefreshOk() (*bool, bool)`

GetShowLastComputationRefreshOk returns a tuple with the ShowLastComputationRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLastComputationRefresh

`func (o *SavedInsightNode) SetShowLastComputationRefresh(v bool)`

SetShowLastComputationRefresh sets ShowLastComputationRefresh field to given value.

### HasShowLastComputationRefresh

`func (o *SavedInsightNode) HasShowLastComputationRefresh() bool`

HasShowLastComputationRefresh returns a boolean if a field has been set.

### SetShowLastComputationRefreshNil

`func (o *SavedInsightNode) SetShowLastComputationRefreshNil(b bool)`

 SetShowLastComputationRefreshNil sets the value for ShowLastComputationRefresh to be an explicit nil

### UnsetShowLastComputationRefresh
`func (o *SavedInsightNode) UnsetShowLastComputationRefresh()`

UnsetShowLastComputationRefresh ensures that no value is present for ShowLastComputationRefresh, not even an explicit nil
### GetShowOpenEditorButton

`func (o *SavedInsightNode) GetShowOpenEditorButton() bool`

GetShowOpenEditorButton returns the ShowOpenEditorButton field if non-nil, zero value otherwise.

### GetShowOpenEditorButtonOk

`func (o *SavedInsightNode) GetShowOpenEditorButtonOk() (*bool, bool)`

GetShowOpenEditorButtonOk returns a tuple with the ShowOpenEditorButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOpenEditorButton

`func (o *SavedInsightNode) SetShowOpenEditorButton(v bool)`

SetShowOpenEditorButton sets ShowOpenEditorButton field to given value.

### HasShowOpenEditorButton

`func (o *SavedInsightNode) HasShowOpenEditorButton() bool`

HasShowOpenEditorButton returns a boolean if a field has been set.

### SetShowOpenEditorButtonNil

`func (o *SavedInsightNode) SetShowOpenEditorButtonNil(b bool)`

 SetShowOpenEditorButtonNil sets the value for ShowOpenEditorButton to be an explicit nil

### UnsetShowOpenEditorButton
`func (o *SavedInsightNode) UnsetShowOpenEditorButton()`

UnsetShowOpenEditorButton ensures that no value is present for ShowOpenEditorButton, not even an explicit nil
### GetShowPersistentColumnConfigurator

`func (o *SavedInsightNode) GetShowPersistentColumnConfigurator() bool`

GetShowPersistentColumnConfigurator returns the ShowPersistentColumnConfigurator field if non-nil, zero value otherwise.

### GetShowPersistentColumnConfiguratorOk

`func (o *SavedInsightNode) GetShowPersistentColumnConfiguratorOk() (*bool, bool)`

GetShowPersistentColumnConfiguratorOk returns a tuple with the ShowPersistentColumnConfigurator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPersistentColumnConfigurator

`func (o *SavedInsightNode) SetShowPersistentColumnConfigurator(v bool)`

SetShowPersistentColumnConfigurator sets ShowPersistentColumnConfigurator field to given value.

### HasShowPersistentColumnConfigurator

`func (o *SavedInsightNode) HasShowPersistentColumnConfigurator() bool`

HasShowPersistentColumnConfigurator returns a boolean if a field has been set.

### SetShowPersistentColumnConfiguratorNil

`func (o *SavedInsightNode) SetShowPersistentColumnConfiguratorNil(b bool)`

 SetShowPersistentColumnConfiguratorNil sets the value for ShowPersistentColumnConfigurator to be an explicit nil

### UnsetShowPersistentColumnConfigurator
`func (o *SavedInsightNode) UnsetShowPersistentColumnConfigurator()`

UnsetShowPersistentColumnConfigurator ensures that no value is present for ShowPersistentColumnConfigurator, not even an explicit nil
### GetShowPropertyFilter

`func (o *SavedInsightNode) GetShowPropertyFilter() Showpropertyfilter`

GetShowPropertyFilter returns the ShowPropertyFilter field if non-nil, zero value otherwise.

### GetShowPropertyFilterOk

`func (o *SavedInsightNode) GetShowPropertyFilterOk() (*Showpropertyfilter, bool)`

GetShowPropertyFilterOk returns a tuple with the ShowPropertyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPropertyFilter

`func (o *SavedInsightNode) SetShowPropertyFilter(v Showpropertyfilter)`

SetShowPropertyFilter sets ShowPropertyFilter field to given value.

### HasShowPropertyFilter

`func (o *SavedInsightNode) HasShowPropertyFilter() bool`

HasShowPropertyFilter returns a boolean if a field has been set.

### SetShowPropertyFilterNil

`func (o *SavedInsightNode) SetShowPropertyFilterNil(b bool)`

 SetShowPropertyFilterNil sets the value for ShowPropertyFilter to be an explicit nil

### UnsetShowPropertyFilter
`func (o *SavedInsightNode) UnsetShowPropertyFilter()`

UnsetShowPropertyFilter ensures that no value is present for ShowPropertyFilter, not even an explicit nil
### GetShowReload

`func (o *SavedInsightNode) GetShowReload() bool`

GetShowReload returns the ShowReload field if non-nil, zero value otherwise.

### GetShowReloadOk

`func (o *SavedInsightNode) GetShowReloadOk() (*bool, bool)`

GetShowReloadOk returns a tuple with the ShowReload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowReload

`func (o *SavedInsightNode) SetShowReload(v bool)`

SetShowReload sets ShowReload field to given value.

### HasShowReload

`func (o *SavedInsightNode) HasShowReload() bool`

HasShowReload returns a boolean if a field has been set.

### SetShowReloadNil

`func (o *SavedInsightNode) SetShowReloadNil(b bool)`

 SetShowReloadNil sets the value for ShowReload to be an explicit nil

### UnsetShowReload
`func (o *SavedInsightNode) UnsetShowReload()`

UnsetShowReload ensures that no value is present for ShowReload, not even an explicit nil
### GetShowResults

`func (o *SavedInsightNode) GetShowResults() bool`

GetShowResults returns the ShowResults field if non-nil, zero value otherwise.

### GetShowResultsOk

`func (o *SavedInsightNode) GetShowResultsOk() (*bool, bool)`

GetShowResultsOk returns a tuple with the ShowResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowResults

`func (o *SavedInsightNode) SetShowResults(v bool)`

SetShowResults sets ShowResults field to given value.

### HasShowResults

`func (o *SavedInsightNode) HasShowResults() bool`

HasShowResults returns a boolean if a field has been set.

### SetShowResultsNil

`func (o *SavedInsightNode) SetShowResultsNil(b bool)`

 SetShowResultsNil sets the value for ShowResults to be an explicit nil

### UnsetShowResults
`func (o *SavedInsightNode) UnsetShowResults()`

UnsetShowResults ensures that no value is present for ShowResults, not even an explicit nil
### GetShowResultsTable

`func (o *SavedInsightNode) GetShowResultsTable() bool`

GetShowResultsTable returns the ShowResultsTable field if non-nil, zero value otherwise.

### GetShowResultsTableOk

`func (o *SavedInsightNode) GetShowResultsTableOk() (*bool, bool)`

GetShowResultsTableOk returns a tuple with the ShowResultsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowResultsTable

`func (o *SavedInsightNode) SetShowResultsTable(v bool)`

SetShowResultsTable sets ShowResultsTable field to given value.

### HasShowResultsTable

`func (o *SavedInsightNode) HasShowResultsTable() bool`

HasShowResultsTable returns a boolean if a field has been set.

### SetShowResultsTableNil

`func (o *SavedInsightNode) SetShowResultsTableNil(b bool)`

 SetShowResultsTableNil sets the value for ShowResultsTable to be an explicit nil

### UnsetShowResultsTable
`func (o *SavedInsightNode) UnsetShowResultsTable()`

UnsetShowResultsTable ensures that no value is present for ShowResultsTable, not even an explicit nil
### GetShowSavedFilters

`func (o *SavedInsightNode) GetShowSavedFilters() bool`

GetShowSavedFilters returns the ShowSavedFilters field if non-nil, zero value otherwise.

### GetShowSavedFiltersOk

`func (o *SavedInsightNode) GetShowSavedFiltersOk() (*bool, bool)`

GetShowSavedFiltersOk returns a tuple with the ShowSavedFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSavedFilters

`func (o *SavedInsightNode) SetShowSavedFilters(v bool)`

SetShowSavedFilters sets ShowSavedFilters field to given value.

### HasShowSavedFilters

`func (o *SavedInsightNode) HasShowSavedFilters() bool`

HasShowSavedFilters returns a boolean if a field has been set.

### SetShowSavedFiltersNil

`func (o *SavedInsightNode) SetShowSavedFiltersNil(b bool)`

 SetShowSavedFiltersNil sets the value for ShowSavedFilters to be an explicit nil

### UnsetShowSavedFilters
`func (o *SavedInsightNode) UnsetShowSavedFilters()`

UnsetShowSavedFilters ensures that no value is present for ShowSavedFilters, not even an explicit nil
### GetShowSavedQueries

`func (o *SavedInsightNode) GetShowSavedQueries() bool`

GetShowSavedQueries returns the ShowSavedQueries field if non-nil, zero value otherwise.

### GetShowSavedQueriesOk

`func (o *SavedInsightNode) GetShowSavedQueriesOk() (*bool, bool)`

GetShowSavedQueriesOk returns a tuple with the ShowSavedQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSavedQueries

`func (o *SavedInsightNode) SetShowSavedQueries(v bool)`

SetShowSavedQueries sets ShowSavedQueries field to given value.

### HasShowSavedQueries

`func (o *SavedInsightNode) HasShowSavedQueries() bool`

HasShowSavedQueries returns a boolean if a field has been set.

### SetShowSavedQueriesNil

`func (o *SavedInsightNode) SetShowSavedQueriesNil(b bool)`

 SetShowSavedQueriesNil sets the value for ShowSavedQueries to be an explicit nil

### UnsetShowSavedQueries
`func (o *SavedInsightNode) UnsetShowSavedQueries()`

UnsetShowSavedQueries ensures that no value is present for ShowSavedQueries, not even an explicit nil
### GetShowSearch

`func (o *SavedInsightNode) GetShowSearch() bool`

GetShowSearch returns the ShowSearch field if non-nil, zero value otherwise.

### GetShowSearchOk

`func (o *SavedInsightNode) GetShowSearchOk() (*bool, bool)`

GetShowSearchOk returns a tuple with the ShowSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSearch

`func (o *SavedInsightNode) SetShowSearch(v bool)`

SetShowSearch sets ShowSearch field to given value.

### HasShowSearch

`func (o *SavedInsightNode) HasShowSearch() bool`

HasShowSearch returns a boolean if a field has been set.

### SetShowSearchNil

`func (o *SavedInsightNode) SetShowSearchNil(b bool)`

 SetShowSearchNil sets the value for ShowSearch to be an explicit nil

### UnsetShowSearch
`func (o *SavedInsightNode) UnsetShowSearch()`

UnsetShowSearch ensures that no value is present for ShowSearch, not even an explicit nil
### GetShowTable

`func (o *SavedInsightNode) GetShowTable() bool`

GetShowTable returns the ShowTable field if non-nil, zero value otherwise.

### GetShowTableOk

`func (o *SavedInsightNode) GetShowTableOk() (*bool, bool)`

GetShowTableOk returns a tuple with the ShowTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTable

`func (o *SavedInsightNode) SetShowTable(v bool)`

SetShowTable sets ShowTable field to given value.

### HasShowTable

`func (o *SavedInsightNode) HasShowTable() bool`

HasShowTable returns a boolean if a field has been set.

### SetShowTableNil

`func (o *SavedInsightNode) SetShowTableNil(b bool)`

 SetShowTableNil sets the value for ShowTable to be an explicit nil

### UnsetShowTable
`func (o *SavedInsightNode) UnsetShowTable()`

UnsetShowTable ensures that no value is present for ShowTable, not even an explicit nil
### GetShowTestAccountFilters

`func (o *SavedInsightNode) GetShowTestAccountFilters() bool`

GetShowTestAccountFilters returns the ShowTestAccountFilters field if non-nil, zero value otherwise.

### GetShowTestAccountFiltersOk

`func (o *SavedInsightNode) GetShowTestAccountFiltersOk() (*bool, bool)`

GetShowTestAccountFiltersOk returns a tuple with the ShowTestAccountFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTestAccountFilters

`func (o *SavedInsightNode) SetShowTestAccountFilters(v bool)`

SetShowTestAccountFilters sets ShowTestAccountFilters field to given value.

### HasShowTestAccountFilters

`func (o *SavedInsightNode) HasShowTestAccountFilters() bool`

HasShowTestAccountFilters returns a boolean if a field has been set.

### SetShowTestAccountFiltersNil

`func (o *SavedInsightNode) SetShowTestAccountFiltersNil(b bool)`

 SetShowTestAccountFiltersNil sets the value for ShowTestAccountFilters to be an explicit nil

### UnsetShowTestAccountFilters
`func (o *SavedInsightNode) UnsetShowTestAccountFilters()`

UnsetShowTestAccountFilters ensures that no value is present for ShowTestAccountFilters, not even an explicit nil
### GetShowTimings

`func (o *SavedInsightNode) GetShowTimings() bool`

GetShowTimings returns the ShowTimings field if non-nil, zero value otherwise.

### GetShowTimingsOk

`func (o *SavedInsightNode) GetShowTimingsOk() (*bool, bool)`

GetShowTimingsOk returns a tuple with the ShowTimings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTimings

`func (o *SavedInsightNode) SetShowTimings(v bool)`

SetShowTimings sets ShowTimings field to given value.

### HasShowTimings

`func (o *SavedInsightNode) HasShowTimings() bool`

HasShowTimings returns a boolean if a field has been set.

### SetShowTimingsNil

`func (o *SavedInsightNode) SetShowTimingsNil(b bool)`

 SetShowTimingsNil sets the value for ShowTimings to be an explicit nil

### UnsetShowTimings
`func (o *SavedInsightNode) UnsetShowTimings()`

UnsetShowTimings ensures that no value is present for ShowTimings, not even an explicit nil
### GetSuppressSessionAnalysisWarning

`func (o *SavedInsightNode) GetSuppressSessionAnalysisWarning() bool`

GetSuppressSessionAnalysisWarning returns the SuppressSessionAnalysisWarning field if non-nil, zero value otherwise.

### GetSuppressSessionAnalysisWarningOk

`func (o *SavedInsightNode) GetSuppressSessionAnalysisWarningOk() (*bool, bool)`

GetSuppressSessionAnalysisWarningOk returns a tuple with the SuppressSessionAnalysisWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressSessionAnalysisWarning

`func (o *SavedInsightNode) SetSuppressSessionAnalysisWarning(v bool)`

SetSuppressSessionAnalysisWarning sets SuppressSessionAnalysisWarning field to given value.

### HasSuppressSessionAnalysisWarning

`func (o *SavedInsightNode) HasSuppressSessionAnalysisWarning() bool`

HasSuppressSessionAnalysisWarning returns a boolean if a field has been set.

### SetSuppressSessionAnalysisWarningNil

`func (o *SavedInsightNode) SetSuppressSessionAnalysisWarningNil(b bool)`

 SetSuppressSessionAnalysisWarningNil sets the value for SuppressSessionAnalysisWarning to be an explicit nil

### UnsetSuppressSessionAnalysisWarning
`func (o *SavedInsightNode) UnsetSuppressSessionAnalysisWarning()`

UnsetSuppressSessionAnalysisWarning ensures that no value is present for SuppressSessionAnalysisWarning, not even an explicit nil
### GetVersion

`func (o *SavedInsightNode) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SavedInsightNode) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SavedInsightNode) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SavedInsightNode) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *SavedInsightNode) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *SavedInsightNode) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetVizSpecificOptions

`func (o *SavedInsightNode) GetVizSpecificOptions() VizSpecificOptions`

GetVizSpecificOptions returns the VizSpecificOptions field if non-nil, zero value otherwise.

### GetVizSpecificOptionsOk

`func (o *SavedInsightNode) GetVizSpecificOptionsOk() (*VizSpecificOptions, bool)`

GetVizSpecificOptionsOk returns a tuple with the VizSpecificOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizSpecificOptions

`func (o *SavedInsightNode) SetVizSpecificOptions(v VizSpecificOptions)`

SetVizSpecificOptions sets VizSpecificOptions field to given value.

### HasVizSpecificOptions

`func (o *SavedInsightNode) HasVizSpecificOptions() bool`

HasVizSpecificOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


