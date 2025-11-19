# BatchExportBackfill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Progress** | **string** |  | [readonly] 
**StartAt** | Pointer to **NullableTime** | The start of the data interval. | [optional] 
**EndAt** | Pointer to **NullableTime** | The end of the data interval. | [optional] 
**Status** | [**BatchExportBackfillStatusEnum**](BatchExportBackfillStatusEnum.md) | The status of this backfill.  * &#x60;Cancelled&#x60; - Cancelled * &#x60;Completed&#x60; - Completed * &#x60;ContinuedAsNew&#x60; - Continued As New * &#x60;Failed&#x60; - Failed * &#x60;FailedRetryable&#x60; - Failed Retryable * &#x60;Terminated&#x60; - Terminated * &#x60;TimedOut&#x60; - Timedout * &#x60;Running&#x60; - Running * &#x60;Starting&#x60; - Starting | 
**CreatedAt** | **time.Time** | The timestamp at which this BatchExportBackfill was created. | [readonly] 
**FinishedAt** | Pointer to **NullableTime** | The timestamp at which this BatchExportBackfill finished, successfully or not. | [optional] 
**LastUpdatedAt** | **time.Time** | The timestamp at which this BatchExportBackfill was last updated. | [readonly] 
**Team** | **int32** | The team this belongs to. | 
**BatchExport** | **string** | The BatchExport this backfill belongs to. | 

## Methods

### NewBatchExportBackfill

`func NewBatchExportBackfill(id string, progress string, status BatchExportBackfillStatusEnum, createdAt time.Time, lastUpdatedAt time.Time, team int32, batchExport string, ) *BatchExportBackfill`

NewBatchExportBackfill instantiates a new BatchExportBackfill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchExportBackfillWithDefaults

`func NewBatchExportBackfillWithDefaults() *BatchExportBackfill`

NewBatchExportBackfillWithDefaults instantiates a new BatchExportBackfill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchExportBackfill) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchExportBackfill) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchExportBackfill) SetId(v string)`

SetId sets Id field to given value.


### GetProgress

`func (o *BatchExportBackfill) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *BatchExportBackfill) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *BatchExportBackfill) SetProgress(v string)`

SetProgress sets Progress field to given value.


### GetStartAt

`func (o *BatchExportBackfill) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *BatchExportBackfill) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *BatchExportBackfill) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *BatchExportBackfill) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### SetStartAtNil

`func (o *BatchExportBackfill) SetStartAtNil(b bool)`

 SetStartAtNil sets the value for StartAt to be an explicit nil

### UnsetStartAt
`func (o *BatchExportBackfill) UnsetStartAt()`

UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
### GetEndAt

`func (o *BatchExportBackfill) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *BatchExportBackfill) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *BatchExportBackfill) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *BatchExportBackfill) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *BatchExportBackfill) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *BatchExportBackfill) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetStatus

`func (o *BatchExportBackfill) GetStatus() BatchExportBackfillStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchExportBackfill) GetStatusOk() (*BatchExportBackfillStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchExportBackfill) SetStatus(v BatchExportBackfillStatusEnum)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *BatchExportBackfill) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BatchExportBackfill) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BatchExportBackfill) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFinishedAt

`func (o *BatchExportBackfill) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *BatchExportBackfill) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *BatchExportBackfill) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *BatchExportBackfill) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *BatchExportBackfill) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *BatchExportBackfill) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetLastUpdatedAt

`func (o *BatchExportBackfill) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *BatchExportBackfill) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *BatchExportBackfill) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.


### GetTeam

`func (o *BatchExportBackfill) GetTeam() int32`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *BatchExportBackfill) GetTeamOk() (*int32, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *BatchExportBackfill) SetTeam(v int32)`

SetTeam sets Team field to given value.


### GetBatchExport

`func (o *BatchExportBackfill) GetBatchExport() string`

GetBatchExport returns the BatchExport field if non-nil, zero value otherwise.

### GetBatchExportOk

`func (o *BatchExportBackfill) GetBatchExportOk() (*string, bool)`

GetBatchExportOk returns a tuple with the BatchExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchExport

`func (o *BatchExportBackfill) SetBatchExport(v string)`

SetBatchExport sets BatchExport field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


