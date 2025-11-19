# PatchedBatchExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**TeamId** | Pointer to **int32** | The team this belongs to. | [optional] [readonly] 
**Name** | Pointer to **string** | A human-readable name for this BatchExport. | [optional] 
**Model** | Pointer to **NullableString** | Which model this BatchExport is exporting.  * &#x60;events&#x60; - Events * &#x60;persons&#x60; - Persons * &#x60;sessions&#x60; - Sessions | [optional] 
**Destination** | Pointer to [**BatchExportDestination**](BatchExportDestination.md) |  | [optional] 
**Interval** | Pointer to [**IntervalEnum**](IntervalEnum.md) |  | [optional] 
**Paused** | Pointer to **bool** | Whether this BatchExport is paused or not. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp at which this BatchExport was created. | [optional] [readonly] 
**LastUpdatedAt** | Pointer to **time.Time** | The timestamp at which this BatchExport was last updated. | [optional] [readonly] 
**LastPausedAt** | Pointer to **NullableTime** | The timestamp at which this BatchExport was last paused. | [optional] 
**StartAt** | Pointer to **NullableTime** | Time before which any Batch Export runs won&#39;t be triggered. | [optional] 
**EndAt** | Pointer to **NullableTime** | Time after which any Batch Export runs won&#39;t be triggered. | [optional] 
**LatestRuns** | Pointer to [**[]BatchExportRun**](BatchExportRun.md) |  | [optional] [readonly] 
**HogqlQuery** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **interface{}** | A schema of custom fields to select when exporting data. | [optional] [readonly] 
**Filters** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchedBatchExport

`func NewPatchedBatchExport() *PatchedBatchExport`

NewPatchedBatchExport instantiates a new PatchedBatchExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBatchExportWithDefaults

`func NewPatchedBatchExportWithDefaults() *PatchedBatchExport`

NewPatchedBatchExportWithDefaults instantiates a new PatchedBatchExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBatchExport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBatchExport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBatchExport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedBatchExport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTeamId

`func (o *PatchedBatchExport) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *PatchedBatchExport) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *PatchedBatchExport) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *PatchedBatchExport) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetName

`func (o *PatchedBatchExport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBatchExport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBatchExport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBatchExport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModel

`func (o *PatchedBatchExport) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PatchedBatchExport) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PatchedBatchExport) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PatchedBatchExport) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *PatchedBatchExport) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *PatchedBatchExport) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetDestination

`func (o *PatchedBatchExport) GetDestination() BatchExportDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PatchedBatchExport) GetDestinationOk() (*BatchExportDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PatchedBatchExport) SetDestination(v BatchExportDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PatchedBatchExport) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetInterval

`func (o *PatchedBatchExport) GetInterval() IntervalEnum`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PatchedBatchExport) GetIntervalOk() (*IntervalEnum, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PatchedBatchExport) SetInterval(v IntervalEnum)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PatchedBatchExport) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetPaused

`func (o *PatchedBatchExport) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *PatchedBatchExport) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *PatchedBatchExport) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *PatchedBatchExport) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedBatchExport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedBatchExport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedBatchExport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedBatchExport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *PatchedBatchExport) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *PatchedBatchExport) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *PatchedBatchExport) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *PatchedBatchExport) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetLastPausedAt

`func (o *PatchedBatchExport) GetLastPausedAt() time.Time`

GetLastPausedAt returns the LastPausedAt field if non-nil, zero value otherwise.

### GetLastPausedAtOk

`func (o *PatchedBatchExport) GetLastPausedAtOk() (*time.Time, bool)`

GetLastPausedAtOk returns a tuple with the LastPausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPausedAt

`func (o *PatchedBatchExport) SetLastPausedAt(v time.Time)`

SetLastPausedAt sets LastPausedAt field to given value.

### HasLastPausedAt

`func (o *PatchedBatchExport) HasLastPausedAt() bool`

HasLastPausedAt returns a boolean if a field has been set.

### SetLastPausedAtNil

`func (o *PatchedBatchExport) SetLastPausedAtNil(b bool)`

 SetLastPausedAtNil sets the value for LastPausedAt to be an explicit nil

### UnsetLastPausedAt
`func (o *PatchedBatchExport) UnsetLastPausedAt()`

UnsetLastPausedAt ensures that no value is present for LastPausedAt, not even an explicit nil
### GetStartAt

`func (o *PatchedBatchExport) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PatchedBatchExport) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PatchedBatchExport) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PatchedBatchExport) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### SetStartAtNil

`func (o *PatchedBatchExport) SetStartAtNil(b bool)`

 SetStartAtNil sets the value for StartAt to be an explicit nil

### UnsetStartAt
`func (o *PatchedBatchExport) UnsetStartAt()`

UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
### GetEndAt

`func (o *PatchedBatchExport) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *PatchedBatchExport) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *PatchedBatchExport) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *PatchedBatchExport) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *PatchedBatchExport) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *PatchedBatchExport) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetLatestRuns

`func (o *PatchedBatchExport) GetLatestRuns() []BatchExportRun`

GetLatestRuns returns the LatestRuns field if non-nil, zero value otherwise.

### GetLatestRunsOk

`func (o *PatchedBatchExport) GetLatestRunsOk() (*[]BatchExportRun, bool)`

GetLatestRunsOk returns a tuple with the LatestRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRuns

`func (o *PatchedBatchExport) SetLatestRuns(v []BatchExportRun)`

SetLatestRuns sets LatestRuns field to given value.

### HasLatestRuns

`func (o *PatchedBatchExport) HasLatestRuns() bool`

HasLatestRuns returns a boolean if a field has been set.

### GetHogqlQuery

`func (o *PatchedBatchExport) GetHogqlQuery() string`

GetHogqlQuery returns the HogqlQuery field if non-nil, zero value otherwise.

### GetHogqlQueryOk

`func (o *PatchedBatchExport) GetHogqlQueryOk() (*string, bool)`

GetHogqlQueryOk returns a tuple with the HogqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogqlQuery

`func (o *PatchedBatchExport) SetHogqlQuery(v string)`

SetHogqlQuery sets HogqlQuery field to given value.

### HasHogqlQuery

`func (o *PatchedBatchExport) HasHogqlQuery() bool`

HasHogqlQuery returns a boolean if a field has been set.

### GetSchema

`func (o *PatchedBatchExport) GetSchema() interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *PatchedBatchExport) GetSchemaOk() (*interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *PatchedBatchExport) SetSchema(v interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *PatchedBatchExport) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *PatchedBatchExport) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *PatchedBatchExport) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetFilters

`func (o *PatchedBatchExport) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PatchedBatchExport) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PatchedBatchExport) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PatchedBatchExport) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *PatchedBatchExport) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *PatchedBatchExport) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


