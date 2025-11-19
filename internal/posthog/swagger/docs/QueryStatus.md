# QueryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complete** | Pointer to **NullableBool** | Whether the query is still running. Will be true if the query is complete, even if it errored. Either result or error will be set. | [optional] [default to false]
**DashboardId** | Pointer to **NullableInt32** |  | [optional] 
**EndTime** | Pointer to **NullableTime** | When did the query execution task finish (whether successfully or not). | [optional] 
**Error** | Pointer to **NullableBool** | If the query failed, this will be set to true. More information can be found in the error_message field. | [optional] [default to false]
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**ExpirationTime** | Pointer to **NullableTime** |  | [optional] 
**Id** | **string** |  | 
**InsightId** | Pointer to **NullableInt32** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**PickupTime** | Pointer to **NullableTime** | When was the query execution task picked up by a worker. | [optional] 
**QueryAsync** | Pointer to **bool** | ONLY async queries use QueryStatus. | [optional] [default to true]
**QueryProgress** | Pointer to [**ClickhouseQueryProgress**](ClickhouseQueryProgress.md) |  | [optional] 
**Results** | Pointer to **interface{}** |  | [optional] [default to null]
**StartTime** | Pointer to **NullableTime** | When was query execution task enqueued. | [optional] 
**TaskId** | Pointer to **NullableString** |  | [optional] 
**TeamId** | **int32** |  | 

## Methods

### NewQueryStatus

`func NewQueryStatus(id string, teamId int32, ) *QueryStatus`

NewQueryStatus instantiates a new QueryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryStatusWithDefaults

`func NewQueryStatusWithDefaults() *QueryStatus`

NewQueryStatusWithDefaults instantiates a new QueryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplete

`func (o *QueryStatus) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *QueryStatus) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *QueryStatus) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *QueryStatus) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### SetCompleteNil

`func (o *QueryStatus) SetCompleteNil(b bool)`

 SetCompleteNil sets the value for Complete to be an explicit nil

### UnsetComplete
`func (o *QueryStatus) UnsetComplete()`

UnsetComplete ensures that no value is present for Complete, not even an explicit nil
### GetDashboardId

`func (o *QueryStatus) GetDashboardId() int32`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *QueryStatus) GetDashboardIdOk() (*int32, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *QueryStatus) SetDashboardId(v int32)`

SetDashboardId sets DashboardId field to given value.

### HasDashboardId

`func (o *QueryStatus) HasDashboardId() bool`

HasDashboardId returns a boolean if a field has been set.

### SetDashboardIdNil

`func (o *QueryStatus) SetDashboardIdNil(b bool)`

 SetDashboardIdNil sets the value for DashboardId to be an explicit nil

### UnsetDashboardId
`func (o *QueryStatus) UnsetDashboardId()`

UnsetDashboardId ensures that no value is present for DashboardId, not even an explicit nil
### GetEndTime

`func (o *QueryStatus) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *QueryStatus) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *QueryStatus) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *QueryStatus) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *QueryStatus) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *QueryStatus) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetError

`func (o *QueryStatus) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *QueryStatus) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *QueryStatus) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *QueryStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *QueryStatus) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *QueryStatus) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetErrorMessage

`func (o *QueryStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *QueryStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *QueryStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *QueryStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *QueryStatus) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *QueryStatus) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetExpirationTime

`func (o *QueryStatus) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *QueryStatus) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *QueryStatus) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *QueryStatus) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### SetExpirationTimeNil

`func (o *QueryStatus) SetExpirationTimeNil(b bool)`

 SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil

### UnsetExpirationTime
`func (o *QueryStatus) UnsetExpirationTime()`

UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
### GetId

`func (o *QueryStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueryStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueryStatus) SetId(v string)`

SetId sets Id field to given value.


### GetInsightId

`func (o *QueryStatus) GetInsightId() int32`

GetInsightId returns the InsightId field if non-nil, zero value otherwise.

### GetInsightIdOk

`func (o *QueryStatus) GetInsightIdOk() (*int32, bool)`

GetInsightIdOk returns a tuple with the InsightId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightId

`func (o *QueryStatus) SetInsightId(v int32)`

SetInsightId sets InsightId field to given value.

### HasInsightId

`func (o *QueryStatus) HasInsightId() bool`

HasInsightId returns a boolean if a field has been set.

### SetInsightIdNil

`func (o *QueryStatus) SetInsightIdNil(b bool)`

 SetInsightIdNil sets the value for InsightId to be an explicit nil

### UnsetInsightId
`func (o *QueryStatus) UnsetInsightId()`

UnsetInsightId ensures that no value is present for InsightId, not even an explicit nil
### GetLabels

`func (o *QueryStatus) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *QueryStatus) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *QueryStatus) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *QueryStatus) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *QueryStatus) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *QueryStatus) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetPickupTime

`func (o *QueryStatus) GetPickupTime() time.Time`

GetPickupTime returns the PickupTime field if non-nil, zero value otherwise.

### GetPickupTimeOk

`func (o *QueryStatus) GetPickupTimeOk() (*time.Time, bool)`

GetPickupTimeOk returns a tuple with the PickupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupTime

`func (o *QueryStatus) SetPickupTime(v time.Time)`

SetPickupTime sets PickupTime field to given value.

### HasPickupTime

`func (o *QueryStatus) HasPickupTime() bool`

HasPickupTime returns a boolean if a field has been set.

### SetPickupTimeNil

`func (o *QueryStatus) SetPickupTimeNil(b bool)`

 SetPickupTimeNil sets the value for PickupTime to be an explicit nil

### UnsetPickupTime
`func (o *QueryStatus) UnsetPickupTime()`

UnsetPickupTime ensures that no value is present for PickupTime, not even an explicit nil
### GetQueryAsync

`func (o *QueryStatus) GetQueryAsync() bool`

GetQueryAsync returns the QueryAsync field if non-nil, zero value otherwise.

### GetQueryAsyncOk

`func (o *QueryStatus) GetQueryAsyncOk() (*bool, bool)`

GetQueryAsyncOk returns a tuple with the QueryAsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAsync

`func (o *QueryStatus) SetQueryAsync(v bool)`

SetQueryAsync sets QueryAsync field to given value.

### HasQueryAsync

`func (o *QueryStatus) HasQueryAsync() bool`

HasQueryAsync returns a boolean if a field has been set.

### GetQueryProgress

`func (o *QueryStatus) GetQueryProgress() ClickhouseQueryProgress`

GetQueryProgress returns the QueryProgress field if non-nil, zero value otherwise.

### GetQueryProgressOk

`func (o *QueryStatus) GetQueryProgressOk() (*ClickhouseQueryProgress, bool)`

GetQueryProgressOk returns a tuple with the QueryProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryProgress

`func (o *QueryStatus) SetQueryProgress(v ClickhouseQueryProgress)`

SetQueryProgress sets QueryProgress field to given value.

### HasQueryProgress

`func (o *QueryStatus) HasQueryProgress() bool`

HasQueryProgress returns a boolean if a field has been set.

### GetResults

`func (o *QueryStatus) GetResults() interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *QueryStatus) GetResultsOk() (*interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *QueryStatus) SetResults(v interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *QueryStatus) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *QueryStatus) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *QueryStatus) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetStartTime

`func (o *QueryStatus) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *QueryStatus) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *QueryStatus) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *QueryStatus) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *QueryStatus) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *QueryStatus) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetTaskId

`func (o *QueryStatus) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *QueryStatus) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *QueryStatus) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *QueryStatus) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *QueryStatus) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *QueryStatus) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetTeamId

`func (o *QueryStatus) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *QueryStatus) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *QueryStatus) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


