# BatchExportRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Status** | [**BatchExportRunStatusEnum**](BatchExportRunStatusEnum.md) | The status of this run.  * &#x60;Cancelled&#x60; - Cancelled * &#x60;Completed&#x60; - Completed * &#x60;ContinuedAsNew&#x60; - Continued As New * &#x60;Failed&#x60; - Failed * &#x60;FailedRetryable&#x60; - Failed Retryable * &#x60;FailedBilling&#x60; - Failed Billing * &#x60;Terminated&#x60; - Terminated * &#x60;TimedOut&#x60; - Timedout * &#x60;Running&#x60; - Running * &#x60;Starting&#x60; - Starting | 
**RecordsCompleted** | Pointer to **NullableInt32** | The number of records that have been exported. | [optional] 
**LatestError** | Pointer to **NullableString** | The latest error that occurred during this run. | [optional] 
**DataIntervalStart** | Pointer to **NullableTime** | The start of the data interval. | [optional] 
**DataIntervalEnd** | **time.Time** | The end of the data interval. | 
**Cursor** | Pointer to **NullableString** | An opaque cursor that may be used to resume. | [optional] 
**CreatedAt** | **time.Time** | The timestamp at which this BatchExportRun was created. | [readonly] 
**FinishedAt** | Pointer to **NullableTime** | The timestamp at which this BatchExportRun finished, successfully or not. | [optional] 
**LastUpdatedAt** | **time.Time** | The timestamp at which this BatchExportRun was last updated. | [readonly] 
**RecordsTotalCount** | Pointer to **NullableInt32** | The total count of records that should be exported in this BatchExportRun. | [optional] 
**BytesExported** | Pointer to **NullableInt64** | The number of bytes that have been exported in this BatchExportRun. | [optional] 
**BatchExport** | **string** | The BatchExport this run belongs to. | [readonly] 
**Backfill** | Pointer to **NullableString** | The backfill this run belongs to. | [optional] 

## Methods

### NewBatchExportRun

`func NewBatchExportRun(id string, status BatchExportRunStatusEnum, dataIntervalEnd time.Time, createdAt time.Time, lastUpdatedAt time.Time, batchExport string, ) *BatchExportRun`

NewBatchExportRun instantiates a new BatchExportRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchExportRunWithDefaults

`func NewBatchExportRunWithDefaults() *BatchExportRun`

NewBatchExportRunWithDefaults instantiates a new BatchExportRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchExportRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchExportRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchExportRun) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *BatchExportRun) GetStatus() BatchExportRunStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchExportRun) GetStatusOk() (*BatchExportRunStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchExportRun) SetStatus(v BatchExportRunStatusEnum)`

SetStatus sets Status field to given value.


### GetRecordsCompleted

`func (o *BatchExportRun) GetRecordsCompleted() int32`

GetRecordsCompleted returns the RecordsCompleted field if non-nil, zero value otherwise.

### GetRecordsCompletedOk

`func (o *BatchExportRun) GetRecordsCompletedOk() (*int32, bool)`

GetRecordsCompletedOk returns a tuple with the RecordsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsCompleted

`func (o *BatchExportRun) SetRecordsCompleted(v int32)`

SetRecordsCompleted sets RecordsCompleted field to given value.

### HasRecordsCompleted

`func (o *BatchExportRun) HasRecordsCompleted() bool`

HasRecordsCompleted returns a boolean if a field has been set.

### SetRecordsCompletedNil

`func (o *BatchExportRun) SetRecordsCompletedNil(b bool)`

 SetRecordsCompletedNil sets the value for RecordsCompleted to be an explicit nil

### UnsetRecordsCompleted
`func (o *BatchExportRun) UnsetRecordsCompleted()`

UnsetRecordsCompleted ensures that no value is present for RecordsCompleted, not even an explicit nil
### GetLatestError

`func (o *BatchExportRun) GetLatestError() string`

GetLatestError returns the LatestError field if non-nil, zero value otherwise.

### GetLatestErrorOk

`func (o *BatchExportRun) GetLatestErrorOk() (*string, bool)`

GetLatestErrorOk returns a tuple with the LatestError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestError

`func (o *BatchExportRun) SetLatestError(v string)`

SetLatestError sets LatestError field to given value.

### HasLatestError

`func (o *BatchExportRun) HasLatestError() bool`

HasLatestError returns a boolean if a field has been set.

### SetLatestErrorNil

`func (o *BatchExportRun) SetLatestErrorNil(b bool)`

 SetLatestErrorNil sets the value for LatestError to be an explicit nil

### UnsetLatestError
`func (o *BatchExportRun) UnsetLatestError()`

UnsetLatestError ensures that no value is present for LatestError, not even an explicit nil
### GetDataIntervalStart

`func (o *BatchExportRun) GetDataIntervalStart() time.Time`

GetDataIntervalStart returns the DataIntervalStart field if non-nil, zero value otherwise.

### GetDataIntervalStartOk

`func (o *BatchExportRun) GetDataIntervalStartOk() (*time.Time, bool)`

GetDataIntervalStartOk returns a tuple with the DataIntervalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIntervalStart

`func (o *BatchExportRun) SetDataIntervalStart(v time.Time)`

SetDataIntervalStart sets DataIntervalStart field to given value.

### HasDataIntervalStart

`func (o *BatchExportRun) HasDataIntervalStart() bool`

HasDataIntervalStart returns a boolean if a field has been set.

### SetDataIntervalStartNil

`func (o *BatchExportRun) SetDataIntervalStartNil(b bool)`

 SetDataIntervalStartNil sets the value for DataIntervalStart to be an explicit nil

### UnsetDataIntervalStart
`func (o *BatchExportRun) UnsetDataIntervalStart()`

UnsetDataIntervalStart ensures that no value is present for DataIntervalStart, not even an explicit nil
### GetDataIntervalEnd

`func (o *BatchExportRun) GetDataIntervalEnd() time.Time`

GetDataIntervalEnd returns the DataIntervalEnd field if non-nil, zero value otherwise.

### GetDataIntervalEndOk

`func (o *BatchExportRun) GetDataIntervalEndOk() (*time.Time, bool)`

GetDataIntervalEndOk returns a tuple with the DataIntervalEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIntervalEnd

`func (o *BatchExportRun) SetDataIntervalEnd(v time.Time)`

SetDataIntervalEnd sets DataIntervalEnd field to given value.


### GetCursor

`func (o *BatchExportRun) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *BatchExportRun) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *BatchExportRun) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *BatchExportRun) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### SetCursorNil

`func (o *BatchExportRun) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *BatchExportRun) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
### GetCreatedAt

`func (o *BatchExportRun) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BatchExportRun) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BatchExportRun) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFinishedAt

`func (o *BatchExportRun) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *BatchExportRun) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *BatchExportRun) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *BatchExportRun) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *BatchExportRun) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *BatchExportRun) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetLastUpdatedAt

`func (o *BatchExportRun) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *BatchExportRun) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *BatchExportRun) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.


### GetRecordsTotalCount

`func (o *BatchExportRun) GetRecordsTotalCount() int32`

GetRecordsTotalCount returns the RecordsTotalCount field if non-nil, zero value otherwise.

### GetRecordsTotalCountOk

`func (o *BatchExportRun) GetRecordsTotalCountOk() (*int32, bool)`

GetRecordsTotalCountOk returns a tuple with the RecordsTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsTotalCount

`func (o *BatchExportRun) SetRecordsTotalCount(v int32)`

SetRecordsTotalCount sets RecordsTotalCount field to given value.

### HasRecordsTotalCount

`func (o *BatchExportRun) HasRecordsTotalCount() bool`

HasRecordsTotalCount returns a boolean if a field has been set.

### SetRecordsTotalCountNil

`func (o *BatchExportRun) SetRecordsTotalCountNil(b bool)`

 SetRecordsTotalCountNil sets the value for RecordsTotalCount to be an explicit nil

### UnsetRecordsTotalCount
`func (o *BatchExportRun) UnsetRecordsTotalCount()`

UnsetRecordsTotalCount ensures that no value is present for RecordsTotalCount, not even an explicit nil
### GetBytesExported

`func (o *BatchExportRun) GetBytesExported() int64`

GetBytesExported returns the BytesExported field if non-nil, zero value otherwise.

### GetBytesExportedOk

`func (o *BatchExportRun) GetBytesExportedOk() (*int64, bool)`

GetBytesExportedOk returns a tuple with the BytesExported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesExported

`func (o *BatchExportRun) SetBytesExported(v int64)`

SetBytesExported sets BytesExported field to given value.

### HasBytesExported

`func (o *BatchExportRun) HasBytesExported() bool`

HasBytesExported returns a boolean if a field has been set.

### SetBytesExportedNil

`func (o *BatchExportRun) SetBytesExportedNil(b bool)`

 SetBytesExportedNil sets the value for BytesExported to be an explicit nil

### UnsetBytesExported
`func (o *BatchExportRun) UnsetBytesExported()`

UnsetBytesExported ensures that no value is present for BytesExported, not even an explicit nil
### GetBatchExport

`func (o *BatchExportRun) GetBatchExport() string`

GetBatchExport returns the BatchExport field if non-nil, zero value otherwise.

### GetBatchExportOk

`func (o *BatchExportRun) GetBatchExportOk() (*string, bool)`

GetBatchExportOk returns a tuple with the BatchExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchExport

`func (o *BatchExportRun) SetBatchExport(v string)`

SetBatchExport sets BatchExport field to given value.


### GetBackfill

`func (o *BatchExportRun) GetBackfill() string`

GetBackfill returns the Backfill field if non-nil, zero value otherwise.

### GetBackfillOk

`func (o *BatchExportRun) GetBackfillOk() (*string, bool)`

GetBackfillOk returns a tuple with the Backfill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfill

`func (o *BatchExportRun) SetBackfill(v string)`

SetBackfill sets Backfill field to given value.

### HasBackfill

`func (o *BatchExportRun) HasBackfill() bool`

HasBackfill returns a boolean if a field has been set.

### SetBackfillNil

`func (o *BatchExportRun) SetBackfillNil(b bool)`

 SetBackfillNil sets the value for Backfill to be an explicit nil

### UnsetBackfill
`func (o *BatchExportRun) UnsetBackfill()`

UnsetBackfill ensures that no value is present for Backfill, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


