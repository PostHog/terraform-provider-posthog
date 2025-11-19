# BatchExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**TeamId** | **int32** | The team this belongs to. | [readonly] 
**Name** | **string** | A human-readable name for this BatchExport. | 
**Model** | Pointer to **NullableString** | Which model this BatchExport is exporting.  * &#x60;events&#x60; - Events * &#x60;persons&#x60; - Persons * &#x60;sessions&#x60; - Sessions | [optional] 
**Destination** | [**BatchExportDestination**](BatchExportDestination.md) |  | 
**Interval** | [**IntervalEnum**](IntervalEnum.md) |  | 
**Paused** | Pointer to **bool** | Whether this BatchExport is paused or not. | [optional] 
**CreatedAt** | **time.Time** | The timestamp at which this BatchExport was created. | [readonly] 
**LastUpdatedAt** | **time.Time** | The timestamp at which this BatchExport was last updated. | [readonly] 
**LastPausedAt** | Pointer to **NullableTime** | The timestamp at which this BatchExport was last paused. | [optional] 
**StartAt** | Pointer to **NullableTime** | Time before which any Batch Export runs won&#39;t be triggered. | [optional] 
**EndAt** | Pointer to **NullableTime** | Time after which any Batch Export runs won&#39;t be triggered. | [optional] 
**LatestRuns** | [**[]BatchExportRun**](BatchExportRun.md) |  | [readonly] 
**HogqlQuery** | Pointer to **string** |  | [optional] 
**Schema** | **interface{}** | A schema of custom fields to select when exporting data. | [readonly] 
**Filters** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewBatchExport

`func NewBatchExport(id string, teamId int32, name string, destination BatchExportDestination, interval IntervalEnum, createdAt time.Time, lastUpdatedAt time.Time, latestRuns []BatchExportRun, schema interface{}, ) *BatchExport`

NewBatchExport instantiates a new BatchExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchExportWithDefaults

`func NewBatchExportWithDefaults() *BatchExport`

NewBatchExportWithDefaults instantiates a new BatchExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchExport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchExport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchExport) SetId(v string)`

SetId sets Id field to given value.


### GetTeamId

`func (o *BatchExport) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *BatchExport) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *BatchExport) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.


### GetName

`func (o *BatchExport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BatchExport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BatchExport) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *BatchExport) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BatchExport) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BatchExport) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BatchExport) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *BatchExport) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *BatchExport) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetDestination

`func (o *BatchExport) GetDestination() BatchExportDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *BatchExport) GetDestinationOk() (*BatchExportDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *BatchExport) SetDestination(v BatchExportDestination)`

SetDestination sets Destination field to given value.


### GetInterval

`func (o *BatchExport) GetInterval() IntervalEnum`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BatchExport) GetIntervalOk() (*IntervalEnum, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BatchExport) SetInterval(v IntervalEnum)`

SetInterval sets Interval field to given value.


### GetPaused

`func (o *BatchExport) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *BatchExport) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *BatchExport) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *BatchExport) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BatchExport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BatchExport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BatchExport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastUpdatedAt

`func (o *BatchExport) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *BatchExport) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *BatchExport) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.


### GetLastPausedAt

`func (o *BatchExport) GetLastPausedAt() time.Time`

GetLastPausedAt returns the LastPausedAt field if non-nil, zero value otherwise.

### GetLastPausedAtOk

`func (o *BatchExport) GetLastPausedAtOk() (*time.Time, bool)`

GetLastPausedAtOk returns a tuple with the LastPausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPausedAt

`func (o *BatchExport) SetLastPausedAt(v time.Time)`

SetLastPausedAt sets LastPausedAt field to given value.

### HasLastPausedAt

`func (o *BatchExport) HasLastPausedAt() bool`

HasLastPausedAt returns a boolean if a field has been set.

### SetLastPausedAtNil

`func (o *BatchExport) SetLastPausedAtNil(b bool)`

 SetLastPausedAtNil sets the value for LastPausedAt to be an explicit nil

### UnsetLastPausedAt
`func (o *BatchExport) UnsetLastPausedAt()`

UnsetLastPausedAt ensures that no value is present for LastPausedAt, not even an explicit nil
### GetStartAt

`func (o *BatchExport) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *BatchExport) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *BatchExport) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *BatchExport) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### SetStartAtNil

`func (o *BatchExport) SetStartAtNil(b bool)`

 SetStartAtNil sets the value for StartAt to be an explicit nil

### UnsetStartAt
`func (o *BatchExport) UnsetStartAt()`

UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
### GetEndAt

`func (o *BatchExport) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *BatchExport) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *BatchExport) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *BatchExport) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *BatchExport) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *BatchExport) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetLatestRuns

`func (o *BatchExport) GetLatestRuns() []BatchExportRun`

GetLatestRuns returns the LatestRuns field if non-nil, zero value otherwise.

### GetLatestRunsOk

`func (o *BatchExport) GetLatestRunsOk() (*[]BatchExportRun, bool)`

GetLatestRunsOk returns a tuple with the LatestRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRuns

`func (o *BatchExport) SetLatestRuns(v []BatchExportRun)`

SetLatestRuns sets LatestRuns field to given value.


### GetHogqlQuery

`func (o *BatchExport) GetHogqlQuery() string`

GetHogqlQuery returns the HogqlQuery field if non-nil, zero value otherwise.

### GetHogqlQueryOk

`func (o *BatchExport) GetHogqlQueryOk() (*string, bool)`

GetHogqlQueryOk returns a tuple with the HogqlQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogqlQuery

`func (o *BatchExport) SetHogqlQuery(v string)`

SetHogqlQuery sets HogqlQuery field to given value.

### HasHogqlQuery

`func (o *BatchExport) HasHogqlQuery() bool`

HasHogqlQuery returns a boolean if a field has been set.

### GetSchema

`func (o *BatchExport) GetSchema() interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *BatchExport) GetSchemaOk() (*interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *BatchExport) SetSchema(v interface{})`

SetSchema sets Schema field to given value.


### SetSchemaNil

`func (o *BatchExport) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *BatchExport) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetFilters

`func (o *BatchExport) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *BatchExport) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *BatchExport) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *BatchExport) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *BatchExport) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *BatchExport) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


