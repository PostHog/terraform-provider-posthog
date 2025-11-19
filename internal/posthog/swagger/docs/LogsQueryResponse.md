# LogsQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]string** |  | [optional] 
**Error** | Pointer to **NullableString** | Query error. Returned only if &#39;explain&#39; or &#x60;modifiers.debug&#x60; is true. Throws an error otherwise. | [optional] 
**HasMore** | Pointer to **NullableBool** |  | [optional] 
**Hogql** | Pointer to **NullableString** | Generated HogQL query. | [optional] 
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**QueryStatus** | Pointer to [**QueryStatus**](QueryStatus.md) |  | [optional] 
**ResolvedDateRange** | Pointer to [**ResolvedDateRangeResponse**](ResolvedDateRangeResponse.md) |  | [optional] 
**Results** | **interface{}** |  | 
**Timings** | Pointer to [**[]QueryTiming**](QueryTiming.md) | Measured timings for different parts of the query generation process | [optional] 

## Methods

### NewLogsQueryResponse

`func NewLogsQueryResponse(results interface{}, ) *LogsQueryResponse`

NewLogsQueryResponse instantiates a new LogsQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsQueryResponseWithDefaults

`func NewLogsQueryResponseWithDefaults() *LogsQueryResponse`

NewLogsQueryResponseWithDefaults instantiates a new LogsQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *LogsQueryResponse) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *LogsQueryResponse) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *LogsQueryResponse) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *LogsQueryResponse) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *LogsQueryResponse) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *LogsQueryResponse) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetError

`func (o *LogsQueryResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LogsQueryResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LogsQueryResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *LogsQueryResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *LogsQueryResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *LogsQueryResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetHasMore

`func (o *LogsQueryResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *LogsQueryResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *LogsQueryResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *LogsQueryResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### SetHasMoreNil

`func (o *LogsQueryResponse) SetHasMoreNil(b bool)`

 SetHasMoreNil sets the value for HasMore to be an explicit nil

### UnsetHasMore
`func (o *LogsQueryResponse) UnsetHasMore()`

UnsetHasMore ensures that no value is present for HasMore, not even an explicit nil
### GetHogql

`func (o *LogsQueryResponse) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *LogsQueryResponse) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *LogsQueryResponse) SetHogql(v string)`

SetHogql sets Hogql field to given value.

### HasHogql

`func (o *LogsQueryResponse) HasHogql() bool`

HasHogql returns a boolean if a field has been set.

### SetHogqlNil

`func (o *LogsQueryResponse) SetHogqlNil(b bool)`

 SetHogqlNil sets the value for Hogql to be an explicit nil

### UnsetHogql
`func (o *LogsQueryResponse) UnsetHogql()`

UnsetHogql ensures that no value is present for Hogql, not even an explicit nil
### GetLimit

`func (o *LogsQueryResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LogsQueryResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LogsQueryResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *LogsQueryResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *LogsQueryResponse) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *LogsQueryResponse) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *LogsQueryResponse) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *LogsQueryResponse) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *LogsQueryResponse) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *LogsQueryResponse) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *LogsQueryResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *LogsQueryResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *LogsQueryResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *LogsQueryResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *LogsQueryResponse) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *LogsQueryResponse) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetQueryStatus

`func (o *LogsQueryResponse) GetQueryStatus() QueryStatus`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *LogsQueryResponse) GetQueryStatusOk() (*QueryStatus, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *LogsQueryResponse) SetQueryStatus(v QueryStatus)`

SetQueryStatus sets QueryStatus field to given value.

### HasQueryStatus

`func (o *LogsQueryResponse) HasQueryStatus() bool`

HasQueryStatus returns a boolean if a field has been set.

### GetResolvedDateRange

`func (o *LogsQueryResponse) GetResolvedDateRange() ResolvedDateRangeResponse`

GetResolvedDateRange returns the ResolvedDateRange field if non-nil, zero value otherwise.

### GetResolvedDateRangeOk

`func (o *LogsQueryResponse) GetResolvedDateRangeOk() (*ResolvedDateRangeResponse, bool)`

GetResolvedDateRangeOk returns a tuple with the ResolvedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDateRange

`func (o *LogsQueryResponse) SetResolvedDateRange(v ResolvedDateRangeResponse)`

SetResolvedDateRange sets ResolvedDateRange field to given value.

### HasResolvedDateRange

`func (o *LogsQueryResponse) HasResolvedDateRange() bool`

HasResolvedDateRange returns a boolean if a field has been set.

### GetResults

`func (o *LogsQueryResponse) GetResults() interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *LogsQueryResponse) GetResultsOk() (*interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *LogsQueryResponse) SetResults(v interface{})`

SetResults sets Results field to given value.


### SetResultsNil

`func (o *LogsQueryResponse) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *LogsQueryResponse) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetTimings

`func (o *LogsQueryResponse) GetTimings() []QueryTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *LogsQueryResponse) GetTimingsOk() (*[]QueryTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *LogsQueryResponse) SetTimings(v []QueryTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *LogsQueryResponse) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *LogsQueryResponse) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *LogsQueryResponse) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


