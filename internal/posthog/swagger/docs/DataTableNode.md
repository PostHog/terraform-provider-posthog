# DataTableNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowSorting** | Pointer to **NullableBool** | Can the user click on column headers to sort the table? (default: true) | [optional] 
**Columns** | Pointer to **[]string** | Columns shown in the table, unless the &#x60;source&#x60; provides them. | [optional] 
**Context** | Pointer to [**DataTableNodeViewPropsContext**](DataTableNodeViewPropsContext.md) |  | [optional] 
**DefaultColumns** | Pointer to **[]string** | Default columns to use when resetting column configuration | [optional] 
**Embedded** | Pointer to **NullableBool** | Uses the embedded version of LemonTable | [optional] 
**Expandable** | Pointer to **NullableBool** | Can expand row to show raw event data (default: true) | [optional] 
**Full** | Pointer to **NullableBool** | Show with most visual options enabled. Used in scenes. | [optional] 
**HiddenColumns** | Pointer to **[]string** | Columns that aren&#39;t shown in the table, even if in columns or returned data | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "DataTableNode"]
**PinnedColumns** | Pointer to **[]string** | Columns that are sticky when scrolling horizontally | [optional] 
**PropertiesViaUrl** | Pointer to **NullableBool** | Link properties via the URL (default: false) | [optional] 
**Response** | Pointer to [**NullableResponse1**](Response1.md) |  | [optional] [default to null]
**ShowActions** | Pointer to **NullableBool** | Show the kebab menu at the end of the row | [optional] 
**ShowColumnConfigurator** | Pointer to **NullableBool** | Show a button to configure the table&#39;s columns if possible | [optional] 
**ShowDateRange** | Pointer to **NullableBool** | Show date range selector | [optional] 
**ShowElapsedTime** | Pointer to **NullableBool** | Show the time it takes to run a query | [optional] 
**ShowEventFilter** | Pointer to **NullableBool** | Include an event filter above the table (EventsNode only) | [optional] 
**ShowExport** | Pointer to **NullableBool** | Show the export button | [optional] 
**ShowHogQLEditor** | Pointer to **NullableBool** | Include a HogQL query editor above HogQL tables | [optional] 
**ShowOpenEditorButton** | Pointer to **NullableBool** | Show a button to open the current query as a new insight. (default: true) | [optional] 
**ShowPersistentColumnConfigurator** | Pointer to **NullableBool** | Show a button to configure and persist the table&#39;s default columns if possible | [optional] 
**ShowPropertyFilter** | Pointer to [**NullableShowpropertyfilter**](Showpropertyfilter.md) |  | [optional] [default to null]
**ShowReload** | Pointer to **NullableBool** | Show a reload button | [optional] 
**ShowResultsTable** | Pointer to **NullableBool** | Show a results table | [optional] 
**ShowSavedFilters** | Pointer to **NullableBool** | Show saved filters feature for this table (requires uniqueKey) | [optional] 
**ShowSavedQueries** | Pointer to **NullableBool** | Shows a list of saved queries | [optional] 
**ShowSearch** | Pointer to **NullableBool** | Include a free text search field (PersonsNode only) | [optional] 
**ShowTestAccountFilters** | Pointer to **NullableBool** | Show filter to exclude test accounts | [optional] 
**ShowTimings** | Pointer to **NullableBool** | Show a detailed query timing breakdown | [optional] 
**Source** | [**Source1**](Source1.md) |  | 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewDataTableNode

`func NewDataTableNode(source Source1, ) *DataTableNode`

NewDataTableNode instantiates a new DataTableNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTableNodeWithDefaults

`func NewDataTableNodeWithDefaults() *DataTableNode`

NewDataTableNodeWithDefaults instantiates a new DataTableNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowSorting

`func (o *DataTableNode) GetAllowSorting() bool`

GetAllowSorting returns the AllowSorting field if non-nil, zero value otherwise.

### GetAllowSortingOk

`func (o *DataTableNode) GetAllowSortingOk() (*bool, bool)`

GetAllowSortingOk returns a tuple with the AllowSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSorting

`func (o *DataTableNode) SetAllowSorting(v bool)`

SetAllowSorting sets AllowSorting field to given value.

### HasAllowSorting

`func (o *DataTableNode) HasAllowSorting() bool`

HasAllowSorting returns a boolean if a field has been set.

### SetAllowSortingNil

`func (o *DataTableNode) SetAllowSortingNil(b bool)`

 SetAllowSortingNil sets the value for AllowSorting to be an explicit nil

### UnsetAllowSorting
`func (o *DataTableNode) UnsetAllowSorting()`

UnsetAllowSorting ensures that no value is present for AllowSorting, not even an explicit nil
### GetColumns

`func (o *DataTableNode) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *DataTableNode) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *DataTableNode) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *DataTableNode) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *DataTableNode) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *DataTableNode) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetContext

`func (o *DataTableNode) GetContext() DataTableNodeViewPropsContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DataTableNode) GetContextOk() (*DataTableNodeViewPropsContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DataTableNode) SetContext(v DataTableNodeViewPropsContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *DataTableNode) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetDefaultColumns

`func (o *DataTableNode) GetDefaultColumns() []string`

GetDefaultColumns returns the DefaultColumns field if non-nil, zero value otherwise.

### GetDefaultColumnsOk

`func (o *DataTableNode) GetDefaultColumnsOk() (*[]string, bool)`

GetDefaultColumnsOk returns a tuple with the DefaultColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultColumns

`func (o *DataTableNode) SetDefaultColumns(v []string)`

SetDefaultColumns sets DefaultColumns field to given value.

### HasDefaultColumns

`func (o *DataTableNode) HasDefaultColumns() bool`

HasDefaultColumns returns a boolean if a field has been set.

### SetDefaultColumnsNil

`func (o *DataTableNode) SetDefaultColumnsNil(b bool)`

 SetDefaultColumnsNil sets the value for DefaultColumns to be an explicit nil

### UnsetDefaultColumns
`func (o *DataTableNode) UnsetDefaultColumns()`

UnsetDefaultColumns ensures that no value is present for DefaultColumns, not even an explicit nil
### GetEmbedded

`func (o *DataTableNode) GetEmbedded() bool`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *DataTableNode) GetEmbeddedOk() (*bool, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *DataTableNode) SetEmbedded(v bool)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *DataTableNode) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### SetEmbeddedNil

`func (o *DataTableNode) SetEmbeddedNil(b bool)`

 SetEmbeddedNil sets the value for Embedded to be an explicit nil

### UnsetEmbedded
`func (o *DataTableNode) UnsetEmbedded()`

UnsetEmbedded ensures that no value is present for Embedded, not even an explicit nil
### GetExpandable

`func (o *DataTableNode) GetExpandable() bool`

GetExpandable returns the Expandable field if non-nil, zero value otherwise.

### GetExpandableOk

`func (o *DataTableNode) GetExpandableOk() (*bool, bool)`

GetExpandableOk returns a tuple with the Expandable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandable

`func (o *DataTableNode) SetExpandable(v bool)`

SetExpandable sets Expandable field to given value.

### HasExpandable

`func (o *DataTableNode) HasExpandable() bool`

HasExpandable returns a boolean if a field has been set.

### SetExpandableNil

`func (o *DataTableNode) SetExpandableNil(b bool)`

 SetExpandableNil sets the value for Expandable to be an explicit nil

### UnsetExpandable
`func (o *DataTableNode) UnsetExpandable()`

UnsetExpandable ensures that no value is present for Expandable, not even an explicit nil
### GetFull

`func (o *DataTableNode) GetFull() bool`

GetFull returns the Full field if non-nil, zero value otherwise.

### GetFullOk

`func (o *DataTableNode) GetFullOk() (*bool, bool)`

GetFullOk returns a tuple with the Full field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFull

`func (o *DataTableNode) SetFull(v bool)`

SetFull sets Full field to given value.

### HasFull

`func (o *DataTableNode) HasFull() bool`

HasFull returns a boolean if a field has been set.

### SetFullNil

`func (o *DataTableNode) SetFullNil(b bool)`

 SetFullNil sets the value for Full to be an explicit nil

### UnsetFull
`func (o *DataTableNode) UnsetFull()`

UnsetFull ensures that no value is present for Full, not even an explicit nil
### GetHiddenColumns

`func (o *DataTableNode) GetHiddenColumns() []string`

GetHiddenColumns returns the HiddenColumns field if non-nil, zero value otherwise.

### GetHiddenColumnsOk

`func (o *DataTableNode) GetHiddenColumnsOk() (*[]string, bool)`

GetHiddenColumnsOk returns a tuple with the HiddenColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenColumns

`func (o *DataTableNode) SetHiddenColumns(v []string)`

SetHiddenColumns sets HiddenColumns field to given value.

### HasHiddenColumns

`func (o *DataTableNode) HasHiddenColumns() bool`

HasHiddenColumns returns a boolean if a field has been set.

### SetHiddenColumnsNil

`func (o *DataTableNode) SetHiddenColumnsNil(b bool)`

 SetHiddenColumnsNil sets the value for HiddenColumns to be an explicit nil

### UnsetHiddenColumns
`func (o *DataTableNode) UnsetHiddenColumns()`

UnsetHiddenColumns ensures that no value is present for HiddenColumns, not even an explicit nil
### GetKind

`func (o *DataTableNode) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DataTableNode) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DataTableNode) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DataTableNode) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPinnedColumns

`func (o *DataTableNode) GetPinnedColumns() []string`

GetPinnedColumns returns the PinnedColumns field if non-nil, zero value otherwise.

### GetPinnedColumnsOk

`func (o *DataTableNode) GetPinnedColumnsOk() (*[]string, bool)`

GetPinnedColumnsOk returns a tuple with the PinnedColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedColumns

`func (o *DataTableNode) SetPinnedColumns(v []string)`

SetPinnedColumns sets PinnedColumns field to given value.

### HasPinnedColumns

`func (o *DataTableNode) HasPinnedColumns() bool`

HasPinnedColumns returns a boolean if a field has been set.

### SetPinnedColumnsNil

`func (o *DataTableNode) SetPinnedColumnsNil(b bool)`

 SetPinnedColumnsNil sets the value for PinnedColumns to be an explicit nil

### UnsetPinnedColumns
`func (o *DataTableNode) UnsetPinnedColumns()`

UnsetPinnedColumns ensures that no value is present for PinnedColumns, not even an explicit nil
### GetPropertiesViaUrl

`func (o *DataTableNode) GetPropertiesViaUrl() bool`

GetPropertiesViaUrl returns the PropertiesViaUrl field if non-nil, zero value otherwise.

### GetPropertiesViaUrlOk

`func (o *DataTableNode) GetPropertiesViaUrlOk() (*bool, bool)`

GetPropertiesViaUrlOk returns a tuple with the PropertiesViaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesViaUrl

`func (o *DataTableNode) SetPropertiesViaUrl(v bool)`

SetPropertiesViaUrl sets PropertiesViaUrl field to given value.

### HasPropertiesViaUrl

`func (o *DataTableNode) HasPropertiesViaUrl() bool`

HasPropertiesViaUrl returns a boolean if a field has been set.

### SetPropertiesViaUrlNil

`func (o *DataTableNode) SetPropertiesViaUrlNil(b bool)`

 SetPropertiesViaUrlNil sets the value for PropertiesViaUrl to be an explicit nil

### UnsetPropertiesViaUrl
`func (o *DataTableNode) UnsetPropertiesViaUrl()`

UnsetPropertiesViaUrl ensures that no value is present for PropertiesViaUrl, not even an explicit nil
### GetResponse

`func (o *DataTableNode) GetResponse() Response1`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *DataTableNode) GetResponseOk() (*Response1, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *DataTableNode) SetResponse(v Response1)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *DataTableNode) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *DataTableNode) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *DataTableNode) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetShowActions

`func (o *DataTableNode) GetShowActions() bool`

GetShowActions returns the ShowActions field if non-nil, zero value otherwise.

### GetShowActionsOk

`func (o *DataTableNode) GetShowActionsOk() (*bool, bool)`

GetShowActionsOk returns a tuple with the ShowActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowActions

`func (o *DataTableNode) SetShowActions(v bool)`

SetShowActions sets ShowActions field to given value.

### HasShowActions

`func (o *DataTableNode) HasShowActions() bool`

HasShowActions returns a boolean if a field has been set.

### SetShowActionsNil

`func (o *DataTableNode) SetShowActionsNil(b bool)`

 SetShowActionsNil sets the value for ShowActions to be an explicit nil

### UnsetShowActions
`func (o *DataTableNode) UnsetShowActions()`

UnsetShowActions ensures that no value is present for ShowActions, not even an explicit nil
### GetShowColumnConfigurator

`func (o *DataTableNode) GetShowColumnConfigurator() bool`

GetShowColumnConfigurator returns the ShowColumnConfigurator field if non-nil, zero value otherwise.

### GetShowColumnConfiguratorOk

`func (o *DataTableNode) GetShowColumnConfiguratorOk() (*bool, bool)`

GetShowColumnConfiguratorOk returns a tuple with the ShowColumnConfigurator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowColumnConfigurator

`func (o *DataTableNode) SetShowColumnConfigurator(v bool)`

SetShowColumnConfigurator sets ShowColumnConfigurator field to given value.

### HasShowColumnConfigurator

`func (o *DataTableNode) HasShowColumnConfigurator() bool`

HasShowColumnConfigurator returns a boolean if a field has been set.

### SetShowColumnConfiguratorNil

`func (o *DataTableNode) SetShowColumnConfiguratorNil(b bool)`

 SetShowColumnConfiguratorNil sets the value for ShowColumnConfigurator to be an explicit nil

### UnsetShowColumnConfigurator
`func (o *DataTableNode) UnsetShowColumnConfigurator()`

UnsetShowColumnConfigurator ensures that no value is present for ShowColumnConfigurator, not even an explicit nil
### GetShowDateRange

`func (o *DataTableNode) GetShowDateRange() bool`

GetShowDateRange returns the ShowDateRange field if non-nil, zero value otherwise.

### GetShowDateRangeOk

`func (o *DataTableNode) GetShowDateRangeOk() (*bool, bool)`

GetShowDateRangeOk returns a tuple with the ShowDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDateRange

`func (o *DataTableNode) SetShowDateRange(v bool)`

SetShowDateRange sets ShowDateRange field to given value.

### HasShowDateRange

`func (o *DataTableNode) HasShowDateRange() bool`

HasShowDateRange returns a boolean if a field has been set.

### SetShowDateRangeNil

`func (o *DataTableNode) SetShowDateRangeNil(b bool)`

 SetShowDateRangeNil sets the value for ShowDateRange to be an explicit nil

### UnsetShowDateRange
`func (o *DataTableNode) UnsetShowDateRange()`

UnsetShowDateRange ensures that no value is present for ShowDateRange, not even an explicit nil
### GetShowElapsedTime

`func (o *DataTableNode) GetShowElapsedTime() bool`

GetShowElapsedTime returns the ShowElapsedTime field if non-nil, zero value otherwise.

### GetShowElapsedTimeOk

`func (o *DataTableNode) GetShowElapsedTimeOk() (*bool, bool)`

GetShowElapsedTimeOk returns a tuple with the ShowElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowElapsedTime

`func (o *DataTableNode) SetShowElapsedTime(v bool)`

SetShowElapsedTime sets ShowElapsedTime field to given value.

### HasShowElapsedTime

`func (o *DataTableNode) HasShowElapsedTime() bool`

HasShowElapsedTime returns a boolean if a field has been set.

### SetShowElapsedTimeNil

`func (o *DataTableNode) SetShowElapsedTimeNil(b bool)`

 SetShowElapsedTimeNil sets the value for ShowElapsedTime to be an explicit nil

### UnsetShowElapsedTime
`func (o *DataTableNode) UnsetShowElapsedTime()`

UnsetShowElapsedTime ensures that no value is present for ShowElapsedTime, not even an explicit nil
### GetShowEventFilter

`func (o *DataTableNode) GetShowEventFilter() bool`

GetShowEventFilter returns the ShowEventFilter field if non-nil, zero value otherwise.

### GetShowEventFilterOk

`func (o *DataTableNode) GetShowEventFilterOk() (*bool, bool)`

GetShowEventFilterOk returns a tuple with the ShowEventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowEventFilter

`func (o *DataTableNode) SetShowEventFilter(v bool)`

SetShowEventFilter sets ShowEventFilter field to given value.

### HasShowEventFilter

`func (o *DataTableNode) HasShowEventFilter() bool`

HasShowEventFilter returns a boolean if a field has been set.

### SetShowEventFilterNil

`func (o *DataTableNode) SetShowEventFilterNil(b bool)`

 SetShowEventFilterNil sets the value for ShowEventFilter to be an explicit nil

### UnsetShowEventFilter
`func (o *DataTableNode) UnsetShowEventFilter()`

UnsetShowEventFilter ensures that no value is present for ShowEventFilter, not even an explicit nil
### GetShowExport

`func (o *DataTableNode) GetShowExport() bool`

GetShowExport returns the ShowExport field if non-nil, zero value otherwise.

### GetShowExportOk

`func (o *DataTableNode) GetShowExportOk() (*bool, bool)`

GetShowExportOk returns a tuple with the ShowExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowExport

`func (o *DataTableNode) SetShowExport(v bool)`

SetShowExport sets ShowExport field to given value.

### HasShowExport

`func (o *DataTableNode) HasShowExport() bool`

HasShowExport returns a boolean if a field has been set.

### SetShowExportNil

`func (o *DataTableNode) SetShowExportNil(b bool)`

 SetShowExportNil sets the value for ShowExport to be an explicit nil

### UnsetShowExport
`func (o *DataTableNode) UnsetShowExport()`

UnsetShowExport ensures that no value is present for ShowExport, not even an explicit nil
### GetShowHogQLEditor

`func (o *DataTableNode) GetShowHogQLEditor() bool`

GetShowHogQLEditor returns the ShowHogQLEditor field if non-nil, zero value otherwise.

### GetShowHogQLEditorOk

`func (o *DataTableNode) GetShowHogQLEditorOk() (*bool, bool)`

GetShowHogQLEditorOk returns a tuple with the ShowHogQLEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHogQLEditor

`func (o *DataTableNode) SetShowHogQLEditor(v bool)`

SetShowHogQLEditor sets ShowHogQLEditor field to given value.

### HasShowHogQLEditor

`func (o *DataTableNode) HasShowHogQLEditor() bool`

HasShowHogQLEditor returns a boolean if a field has been set.

### SetShowHogQLEditorNil

`func (o *DataTableNode) SetShowHogQLEditorNil(b bool)`

 SetShowHogQLEditorNil sets the value for ShowHogQLEditor to be an explicit nil

### UnsetShowHogQLEditor
`func (o *DataTableNode) UnsetShowHogQLEditor()`

UnsetShowHogQLEditor ensures that no value is present for ShowHogQLEditor, not even an explicit nil
### GetShowOpenEditorButton

`func (o *DataTableNode) GetShowOpenEditorButton() bool`

GetShowOpenEditorButton returns the ShowOpenEditorButton field if non-nil, zero value otherwise.

### GetShowOpenEditorButtonOk

`func (o *DataTableNode) GetShowOpenEditorButtonOk() (*bool, bool)`

GetShowOpenEditorButtonOk returns a tuple with the ShowOpenEditorButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOpenEditorButton

`func (o *DataTableNode) SetShowOpenEditorButton(v bool)`

SetShowOpenEditorButton sets ShowOpenEditorButton field to given value.

### HasShowOpenEditorButton

`func (o *DataTableNode) HasShowOpenEditorButton() bool`

HasShowOpenEditorButton returns a boolean if a field has been set.

### SetShowOpenEditorButtonNil

`func (o *DataTableNode) SetShowOpenEditorButtonNil(b bool)`

 SetShowOpenEditorButtonNil sets the value for ShowOpenEditorButton to be an explicit nil

### UnsetShowOpenEditorButton
`func (o *DataTableNode) UnsetShowOpenEditorButton()`

UnsetShowOpenEditorButton ensures that no value is present for ShowOpenEditorButton, not even an explicit nil
### GetShowPersistentColumnConfigurator

`func (o *DataTableNode) GetShowPersistentColumnConfigurator() bool`

GetShowPersistentColumnConfigurator returns the ShowPersistentColumnConfigurator field if non-nil, zero value otherwise.

### GetShowPersistentColumnConfiguratorOk

`func (o *DataTableNode) GetShowPersistentColumnConfiguratorOk() (*bool, bool)`

GetShowPersistentColumnConfiguratorOk returns a tuple with the ShowPersistentColumnConfigurator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPersistentColumnConfigurator

`func (o *DataTableNode) SetShowPersistentColumnConfigurator(v bool)`

SetShowPersistentColumnConfigurator sets ShowPersistentColumnConfigurator field to given value.

### HasShowPersistentColumnConfigurator

`func (o *DataTableNode) HasShowPersistentColumnConfigurator() bool`

HasShowPersistentColumnConfigurator returns a boolean if a field has been set.

### SetShowPersistentColumnConfiguratorNil

`func (o *DataTableNode) SetShowPersistentColumnConfiguratorNil(b bool)`

 SetShowPersistentColumnConfiguratorNil sets the value for ShowPersistentColumnConfigurator to be an explicit nil

### UnsetShowPersistentColumnConfigurator
`func (o *DataTableNode) UnsetShowPersistentColumnConfigurator()`

UnsetShowPersistentColumnConfigurator ensures that no value is present for ShowPersistentColumnConfigurator, not even an explicit nil
### GetShowPropertyFilter

`func (o *DataTableNode) GetShowPropertyFilter() Showpropertyfilter`

GetShowPropertyFilter returns the ShowPropertyFilter field if non-nil, zero value otherwise.

### GetShowPropertyFilterOk

`func (o *DataTableNode) GetShowPropertyFilterOk() (*Showpropertyfilter, bool)`

GetShowPropertyFilterOk returns a tuple with the ShowPropertyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPropertyFilter

`func (o *DataTableNode) SetShowPropertyFilter(v Showpropertyfilter)`

SetShowPropertyFilter sets ShowPropertyFilter field to given value.

### HasShowPropertyFilter

`func (o *DataTableNode) HasShowPropertyFilter() bool`

HasShowPropertyFilter returns a boolean if a field has been set.

### SetShowPropertyFilterNil

`func (o *DataTableNode) SetShowPropertyFilterNil(b bool)`

 SetShowPropertyFilterNil sets the value for ShowPropertyFilter to be an explicit nil

### UnsetShowPropertyFilter
`func (o *DataTableNode) UnsetShowPropertyFilter()`

UnsetShowPropertyFilter ensures that no value is present for ShowPropertyFilter, not even an explicit nil
### GetShowReload

`func (o *DataTableNode) GetShowReload() bool`

GetShowReload returns the ShowReload field if non-nil, zero value otherwise.

### GetShowReloadOk

`func (o *DataTableNode) GetShowReloadOk() (*bool, bool)`

GetShowReloadOk returns a tuple with the ShowReload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowReload

`func (o *DataTableNode) SetShowReload(v bool)`

SetShowReload sets ShowReload field to given value.

### HasShowReload

`func (o *DataTableNode) HasShowReload() bool`

HasShowReload returns a boolean if a field has been set.

### SetShowReloadNil

`func (o *DataTableNode) SetShowReloadNil(b bool)`

 SetShowReloadNil sets the value for ShowReload to be an explicit nil

### UnsetShowReload
`func (o *DataTableNode) UnsetShowReload()`

UnsetShowReload ensures that no value is present for ShowReload, not even an explicit nil
### GetShowResultsTable

`func (o *DataTableNode) GetShowResultsTable() bool`

GetShowResultsTable returns the ShowResultsTable field if non-nil, zero value otherwise.

### GetShowResultsTableOk

`func (o *DataTableNode) GetShowResultsTableOk() (*bool, bool)`

GetShowResultsTableOk returns a tuple with the ShowResultsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowResultsTable

`func (o *DataTableNode) SetShowResultsTable(v bool)`

SetShowResultsTable sets ShowResultsTable field to given value.

### HasShowResultsTable

`func (o *DataTableNode) HasShowResultsTable() bool`

HasShowResultsTable returns a boolean if a field has been set.

### SetShowResultsTableNil

`func (o *DataTableNode) SetShowResultsTableNil(b bool)`

 SetShowResultsTableNil sets the value for ShowResultsTable to be an explicit nil

### UnsetShowResultsTable
`func (o *DataTableNode) UnsetShowResultsTable()`

UnsetShowResultsTable ensures that no value is present for ShowResultsTable, not even an explicit nil
### GetShowSavedFilters

`func (o *DataTableNode) GetShowSavedFilters() bool`

GetShowSavedFilters returns the ShowSavedFilters field if non-nil, zero value otherwise.

### GetShowSavedFiltersOk

`func (o *DataTableNode) GetShowSavedFiltersOk() (*bool, bool)`

GetShowSavedFiltersOk returns a tuple with the ShowSavedFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSavedFilters

`func (o *DataTableNode) SetShowSavedFilters(v bool)`

SetShowSavedFilters sets ShowSavedFilters field to given value.

### HasShowSavedFilters

`func (o *DataTableNode) HasShowSavedFilters() bool`

HasShowSavedFilters returns a boolean if a field has been set.

### SetShowSavedFiltersNil

`func (o *DataTableNode) SetShowSavedFiltersNil(b bool)`

 SetShowSavedFiltersNil sets the value for ShowSavedFilters to be an explicit nil

### UnsetShowSavedFilters
`func (o *DataTableNode) UnsetShowSavedFilters()`

UnsetShowSavedFilters ensures that no value is present for ShowSavedFilters, not even an explicit nil
### GetShowSavedQueries

`func (o *DataTableNode) GetShowSavedQueries() bool`

GetShowSavedQueries returns the ShowSavedQueries field if non-nil, zero value otherwise.

### GetShowSavedQueriesOk

`func (o *DataTableNode) GetShowSavedQueriesOk() (*bool, bool)`

GetShowSavedQueriesOk returns a tuple with the ShowSavedQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSavedQueries

`func (o *DataTableNode) SetShowSavedQueries(v bool)`

SetShowSavedQueries sets ShowSavedQueries field to given value.

### HasShowSavedQueries

`func (o *DataTableNode) HasShowSavedQueries() bool`

HasShowSavedQueries returns a boolean if a field has been set.

### SetShowSavedQueriesNil

`func (o *DataTableNode) SetShowSavedQueriesNil(b bool)`

 SetShowSavedQueriesNil sets the value for ShowSavedQueries to be an explicit nil

### UnsetShowSavedQueries
`func (o *DataTableNode) UnsetShowSavedQueries()`

UnsetShowSavedQueries ensures that no value is present for ShowSavedQueries, not even an explicit nil
### GetShowSearch

`func (o *DataTableNode) GetShowSearch() bool`

GetShowSearch returns the ShowSearch field if non-nil, zero value otherwise.

### GetShowSearchOk

`func (o *DataTableNode) GetShowSearchOk() (*bool, bool)`

GetShowSearchOk returns a tuple with the ShowSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSearch

`func (o *DataTableNode) SetShowSearch(v bool)`

SetShowSearch sets ShowSearch field to given value.

### HasShowSearch

`func (o *DataTableNode) HasShowSearch() bool`

HasShowSearch returns a boolean if a field has been set.

### SetShowSearchNil

`func (o *DataTableNode) SetShowSearchNil(b bool)`

 SetShowSearchNil sets the value for ShowSearch to be an explicit nil

### UnsetShowSearch
`func (o *DataTableNode) UnsetShowSearch()`

UnsetShowSearch ensures that no value is present for ShowSearch, not even an explicit nil
### GetShowTestAccountFilters

`func (o *DataTableNode) GetShowTestAccountFilters() bool`

GetShowTestAccountFilters returns the ShowTestAccountFilters field if non-nil, zero value otherwise.

### GetShowTestAccountFiltersOk

`func (o *DataTableNode) GetShowTestAccountFiltersOk() (*bool, bool)`

GetShowTestAccountFiltersOk returns a tuple with the ShowTestAccountFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTestAccountFilters

`func (o *DataTableNode) SetShowTestAccountFilters(v bool)`

SetShowTestAccountFilters sets ShowTestAccountFilters field to given value.

### HasShowTestAccountFilters

`func (o *DataTableNode) HasShowTestAccountFilters() bool`

HasShowTestAccountFilters returns a boolean if a field has been set.

### SetShowTestAccountFiltersNil

`func (o *DataTableNode) SetShowTestAccountFiltersNil(b bool)`

 SetShowTestAccountFiltersNil sets the value for ShowTestAccountFilters to be an explicit nil

### UnsetShowTestAccountFilters
`func (o *DataTableNode) UnsetShowTestAccountFilters()`

UnsetShowTestAccountFilters ensures that no value is present for ShowTestAccountFilters, not even an explicit nil
### GetShowTimings

`func (o *DataTableNode) GetShowTimings() bool`

GetShowTimings returns the ShowTimings field if non-nil, zero value otherwise.

### GetShowTimingsOk

`func (o *DataTableNode) GetShowTimingsOk() (*bool, bool)`

GetShowTimingsOk returns a tuple with the ShowTimings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTimings

`func (o *DataTableNode) SetShowTimings(v bool)`

SetShowTimings sets ShowTimings field to given value.

### HasShowTimings

`func (o *DataTableNode) HasShowTimings() bool`

HasShowTimings returns a boolean if a field has been set.

### SetShowTimingsNil

`func (o *DataTableNode) SetShowTimingsNil(b bool)`

 SetShowTimingsNil sets the value for ShowTimings to be an explicit nil

### UnsetShowTimings
`func (o *DataTableNode) UnsetShowTimings()`

UnsetShowTimings ensures that no value is present for ShowTimings, not even an explicit nil
### GetSource

`func (o *DataTableNode) GetSource() Source1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DataTableNode) GetSourceOk() (*Source1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DataTableNode) SetSource(v Source1)`

SetSource sets Source field to given value.


### GetTags

`func (o *DataTableNode) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DataTableNode) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DataTableNode) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DataTableNode) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *DataTableNode) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DataTableNode) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DataTableNode) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DataTableNode) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DataTableNode) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DataTableNode) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


