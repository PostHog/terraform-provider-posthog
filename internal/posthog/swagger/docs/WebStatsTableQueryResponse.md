# WebStatsTableQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]interface{}** |  | [optional] 
**Error** | Pointer to **NullableString** | Query error. Returned only if &#39;explain&#39; or &#x60;modifiers.debug&#x60; is true. Throws an error otherwise. | [optional] 
**HasMore** | Pointer to **NullableBool** |  | [optional] 
**Hogql** | Pointer to **NullableString** | Generated HogQL query. | [optional] 
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**QueryStatus** | Pointer to [**QueryStatus**](QueryStatus.md) |  | [optional] 
**ResolvedDateRange** | Pointer to [**ResolvedDateRangeResponse**](ResolvedDateRangeResponse.md) |  | [optional] 
**Results** | **[]interface{}** |  | 
**SamplingRate** | Pointer to [**SamplingRate**](SamplingRate.md) |  | [optional] 
**Timings** | Pointer to [**[]QueryTiming**](QueryTiming.md) | Measured timings for different parts of the query generation process | [optional] 
**Types** | Pointer to **[]interface{}** |  | [optional] 
**UsedPreAggregatedTables** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewWebStatsTableQueryResponse

`func NewWebStatsTableQueryResponse(results []interface{}, ) *WebStatsTableQueryResponse`

NewWebStatsTableQueryResponse instantiates a new WebStatsTableQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebStatsTableQueryResponseWithDefaults

`func NewWebStatsTableQueryResponseWithDefaults() *WebStatsTableQueryResponse`

NewWebStatsTableQueryResponseWithDefaults instantiates a new WebStatsTableQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *WebStatsTableQueryResponse) GetColumns() []interface{}`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *WebStatsTableQueryResponse) GetColumnsOk() (*[]interface{}, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *WebStatsTableQueryResponse) SetColumns(v []interface{})`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *WebStatsTableQueryResponse) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *WebStatsTableQueryResponse) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *WebStatsTableQueryResponse) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetError

`func (o *WebStatsTableQueryResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebStatsTableQueryResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebStatsTableQueryResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WebStatsTableQueryResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *WebStatsTableQueryResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *WebStatsTableQueryResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetHasMore

`func (o *WebStatsTableQueryResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *WebStatsTableQueryResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *WebStatsTableQueryResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *WebStatsTableQueryResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### SetHasMoreNil

`func (o *WebStatsTableQueryResponse) SetHasMoreNil(b bool)`

 SetHasMoreNil sets the value for HasMore to be an explicit nil

### UnsetHasMore
`func (o *WebStatsTableQueryResponse) UnsetHasMore()`

UnsetHasMore ensures that no value is present for HasMore, not even an explicit nil
### GetHogql

`func (o *WebStatsTableQueryResponse) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *WebStatsTableQueryResponse) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *WebStatsTableQueryResponse) SetHogql(v string)`

SetHogql sets Hogql field to given value.

### HasHogql

`func (o *WebStatsTableQueryResponse) HasHogql() bool`

HasHogql returns a boolean if a field has been set.

### SetHogqlNil

`func (o *WebStatsTableQueryResponse) SetHogqlNil(b bool)`

 SetHogqlNil sets the value for Hogql to be an explicit nil

### UnsetHogql
`func (o *WebStatsTableQueryResponse) UnsetHogql()`

UnsetHogql ensures that no value is present for Hogql, not even an explicit nil
### GetLimit

`func (o *WebStatsTableQueryResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WebStatsTableQueryResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WebStatsTableQueryResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *WebStatsTableQueryResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *WebStatsTableQueryResponse) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *WebStatsTableQueryResponse) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *WebStatsTableQueryResponse) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *WebStatsTableQueryResponse) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *WebStatsTableQueryResponse) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *WebStatsTableQueryResponse) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *WebStatsTableQueryResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WebStatsTableQueryResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WebStatsTableQueryResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *WebStatsTableQueryResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *WebStatsTableQueryResponse) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *WebStatsTableQueryResponse) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetQueryStatus

`func (o *WebStatsTableQueryResponse) GetQueryStatus() QueryStatus`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *WebStatsTableQueryResponse) GetQueryStatusOk() (*QueryStatus, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *WebStatsTableQueryResponse) SetQueryStatus(v QueryStatus)`

SetQueryStatus sets QueryStatus field to given value.

### HasQueryStatus

`func (o *WebStatsTableQueryResponse) HasQueryStatus() bool`

HasQueryStatus returns a boolean if a field has been set.

### GetResolvedDateRange

`func (o *WebStatsTableQueryResponse) GetResolvedDateRange() ResolvedDateRangeResponse`

GetResolvedDateRange returns the ResolvedDateRange field if non-nil, zero value otherwise.

### GetResolvedDateRangeOk

`func (o *WebStatsTableQueryResponse) GetResolvedDateRangeOk() (*ResolvedDateRangeResponse, bool)`

GetResolvedDateRangeOk returns a tuple with the ResolvedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDateRange

`func (o *WebStatsTableQueryResponse) SetResolvedDateRange(v ResolvedDateRangeResponse)`

SetResolvedDateRange sets ResolvedDateRange field to given value.

### HasResolvedDateRange

`func (o *WebStatsTableQueryResponse) HasResolvedDateRange() bool`

HasResolvedDateRange returns a boolean if a field has been set.

### GetResults

`func (o *WebStatsTableQueryResponse) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *WebStatsTableQueryResponse) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *WebStatsTableQueryResponse) SetResults(v []interface{})`

SetResults sets Results field to given value.


### GetSamplingRate

`func (o *WebStatsTableQueryResponse) GetSamplingRate() SamplingRate`

GetSamplingRate returns the SamplingRate field if non-nil, zero value otherwise.

### GetSamplingRateOk

`func (o *WebStatsTableQueryResponse) GetSamplingRateOk() (*SamplingRate, bool)`

GetSamplingRateOk returns a tuple with the SamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingRate

`func (o *WebStatsTableQueryResponse) SetSamplingRate(v SamplingRate)`

SetSamplingRate sets SamplingRate field to given value.

### HasSamplingRate

`func (o *WebStatsTableQueryResponse) HasSamplingRate() bool`

HasSamplingRate returns a boolean if a field has been set.

### GetTimings

`func (o *WebStatsTableQueryResponse) GetTimings() []QueryTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *WebStatsTableQueryResponse) GetTimingsOk() (*[]QueryTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *WebStatsTableQueryResponse) SetTimings(v []QueryTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *WebStatsTableQueryResponse) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *WebStatsTableQueryResponse) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *WebStatsTableQueryResponse) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil
### GetTypes

`func (o *WebStatsTableQueryResponse) GetTypes() []interface{}`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *WebStatsTableQueryResponse) GetTypesOk() (*[]interface{}, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *WebStatsTableQueryResponse) SetTypes(v []interface{})`

SetTypes sets Types field to given value.

### HasTypes

`func (o *WebStatsTableQueryResponse) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### SetTypesNil

`func (o *WebStatsTableQueryResponse) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *WebStatsTableQueryResponse) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetUsedPreAggregatedTables

`func (o *WebStatsTableQueryResponse) GetUsedPreAggregatedTables() bool`

GetUsedPreAggregatedTables returns the UsedPreAggregatedTables field if non-nil, zero value otherwise.

### GetUsedPreAggregatedTablesOk

`func (o *WebStatsTableQueryResponse) GetUsedPreAggregatedTablesOk() (*bool, bool)`

GetUsedPreAggregatedTablesOk returns a tuple with the UsedPreAggregatedTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPreAggregatedTables

`func (o *WebStatsTableQueryResponse) SetUsedPreAggregatedTables(v bool)`

SetUsedPreAggregatedTables sets UsedPreAggregatedTables field to given value.

### HasUsedPreAggregatedTables

`func (o *WebStatsTableQueryResponse) HasUsedPreAggregatedTables() bool`

HasUsedPreAggregatedTables returns a boolean if a field has been set.

### SetUsedPreAggregatedTablesNil

`func (o *WebStatsTableQueryResponse) SetUsedPreAggregatedTablesNil(b bool)`

 SetUsedPreAggregatedTablesNil sets the value for UsedPreAggregatedTables to be an explicit nil

### UnsetUsedPreAggregatedTables
`func (o *WebStatsTableQueryResponse) UnsetUsedPreAggregatedTables()`

UnsetUsedPreAggregatedTables ensures that no value is present for UsedPreAggregatedTables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


