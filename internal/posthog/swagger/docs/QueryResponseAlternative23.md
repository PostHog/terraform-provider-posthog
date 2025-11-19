# QueryResponseAlternative23

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateFrom** | Pointer to **NullableString** |  | [optional] 
**DateTo** | Pointer to **NullableString** |  | [optional] 
**Error** | Pointer to **NullableString** | Query error. Returned only if &#39;explain&#39; or &#x60;modifiers.debug&#x60; is true. Throws an error otherwise. | [optional] 
**Hogql** | Pointer to **NullableString** | Generated HogQL query. | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**QueryStatus** | Pointer to [**QueryStatus**](QueryStatus.md) |  | [optional] 
**ResolvedDateRange** | Pointer to [**ResolvedDateRangeResponse**](ResolvedDateRangeResponse.md) |  | [optional] 
**Results** | [**[]WebOverviewItem**](WebOverviewItem.md) |  | 
**SamplingRate** | Pointer to [**SamplingRate**](SamplingRate.md) |  | [optional] 
**Timings** | Pointer to [**[]QueryTiming**](QueryTiming.md) | Measured timings for different parts of the query generation process | [optional] 
**UsedPreAggregatedTables** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewQueryResponseAlternative23

`func NewQueryResponseAlternative23(results []WebOverviewItem, ) *QueryResponseAlternative23`

NewQueryResponseAlternative23 instantiates a new QueryResponseAlternative23 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseAlternative23WithDefaults

`func NewQueryResponseAlternative23WithDefaults() *QueryResponseAlternative23`

NewQueryResponseAlternative23WithDefaults instantiates a new QueryResponseAlternative23 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateFrom

`func (o *QueryResponseAlternative23) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *QueryResponseAlternative23) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *QueryResponseAlternative23) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *QueryResponseAlternative23) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### SetDateFromNil

`func (o *QueryResponseAlternative23) SetDateFromNil(b bool)`

 SetDateFromNil sets the value for DateFrom to be an explicit nil

### UnsetDateFrom
`func (o *QueryResponseAlternative23) UnsetDateFrom()`

UnsetDateFrom ensures that no value is present for DateFrom, not even an explicit nil
### GetDateTo

`func (o *QueryResponseAlternative23) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *QueryResponseAlternative23) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *QueryResponseAlternative23) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *QueryResponseAlternative23) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### SetDateToNil

`func (o *QueryResponseAlternative23) SetDateToNil(b bool)`

 SetDateToNil sets the value for DateTo to be an explicit nil

### UnsetDateTo
`func (o *QueryResponseAlternative23) UnsetDateTo()`

UnsetDateTo ensures that no value is present for DateTo, not even an explicit nil
### GetError

`func (o *QueryResponseAlternative23) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *QueryResponseAlternative23) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *QueryResponseAlternative23) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *QueryResponseAlternative23) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *QueryResponseAlternative23) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *QueryResponseAlternative23) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetHogql

`func (o *QueryResponseAlternative23) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *QueryResponseAlternative23) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *QueryResponseAlternative23) SetHogql(v string)`

SetHogql sets Hogql field to given value.

### HasHogql

`func (o *QueryResponseAlternative23) HasHogql() bool`

HasHogql returns a boolean if a field has been set.

### SetHogqlNil

`func (o *QueryResponseAlternative23) SetHogqlNil(b bool)`

 SetHogqlNil sets the value for Hogql to be an explicit nil

### UnsetHogql
`func (o *QueryResponseAlternative23) UnsetHogql()`

UnsetHogql ensures that no value is present for Hogql, not even an explicit nil
### GetModifiers

`func (o *QueryResponseAlternative23) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *QueryResponseAlternative23) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *QueryResponseAlternative23) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *QueryResponseAlternative23) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetQueryStatus

`func (o *QueryResponseAlternative23) GetQueryStatus() QueryStatus`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *QueryResponseAlternative23) GetQueryStatusOk() (*QueryStatus, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *QueryResponseAlternative23) SetQueryStatus(v QueryStatus)`

SetQueryStatus sets QueryStatus field to given value.

### HasQueryStatus

`func (o *QueryResponseAlternative23) HasQueryStatus() bool`

HasQueryStatus returns a boolean if a field has been set.

### GetResolvedDateRange

`func (o *QueryResponseAlternative23) GetResolvedDateRange() ResolvedDateRangeResponse`

GetResolvedDateRange returns the ResolvedDateRange field if non-nil, zero value otherwise.

### GetResolvedDateRangeOk

`func (o *QueryResponseAlternative23) GetResolvedDateRangeOk() (*ResolvedDateRangeResponse, bool)`

GetResolvedDateRangeOk returns a tuple with the ResolvedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDateRange

`func (o *QueryResponseAlternative23) SetResolvedDateRange(v ResolvedDateRangeResponse)`

SetResolvedDateRange sets ResolvedDateRange field to given value.

### HasResolvedDateRange

`func (o *QueryResponseAlternative23) HasResolvedDateRange() bool`

HasResolvedDateRange returns a boolean if a field has been set.

### GetResults

`func (o *QueryResponseAlternative23) GetResults() []WebOverviewItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *QueryResponseAlternative23) GetResultsOk() (*[]WebOverviewItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *QueryResponseAlternative23) SetResults(v []WebOverviewItem)`

SetResults sets Results field to given value.


### GetSamplingRate

`func (o *QueryResponseAlternative23) GetSamplingRate() SamplingRate`

GetSamplingRate returns the SamplingRate field if non-nil, zero value otherwise.

### GetSamplingRateOk

`func (o *QueryResponseAlternative23) GetSamplingRateOk() (*SamplingRate, bool)`

GetSamplingRateOk returns a tuple with the SamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingRate

`func (o *QueryResponseAlternative23) SetSamplingRate(v SamplingRate)`

SetSamplingRate sets SamplingRate field to given value.

### HasSamplingRate

`func (o *QueryResponseAlternative23) HasSamplingRate() bool`

HasSamplingRate returns a boolean if a field has been set.

### GetTimings

`func (o *QueryResponseAlternative23) GetTimings() []QueryTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *QueryResponseAlternative23) GetTimingsOk() (*[]QueryTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *QueryResponseAlternative23) SetTimings(v []QueryTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *QueryResponseAlternative23) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *QueryResponseAlternative23) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *QueryResponseAlternative23) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil
### GetUsedPreAggregatedTables

`func (o *QueryResponseAlternative23) GetUsedPreAggregatedTables() bool`

GetUsedPreAggregatedTables returns the UsedPreAggregatedTables field if non-nil, zero value otherwise.

### GetUsedPreAggregatedTablesOk

`func (o *QueryResponseAlternative23) GetUsedPreAggregatedTablesOk() (*bool, bool)`

GetUsedPreAggregatedTablesOk returns a tuple with the UsedPreAggregatedTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPreAggregatedTables

`func (o *QueryResponseAlternative23) SetUsedPreAggregatedTables(v bool)`

SetUsedPreAggregatedTables sets UsedPreAggregatedTables field to given value.

### HasUsedPreAggregatedTables

`func (o *QueryResponseAlternative23) HasUsedPreAggregatedTables() bool`

HasUsedPreAggregatedTables returns a boolean if a field has been set.

### SetUsedPreAggregatedTablesNil

`func (o *QueryResponseAlternative23) SetUsedPreAggregatedTablesNil(b bool)`

 SetUsedPreAggregatedTablesNil sets the value for UsedPreAggregatedTables to be an explicit nil

### UnsetUsedPreAggregatedTables
`func (o *QueryResponseAlternative23) UnsetUsedPreAggregatedTables()`

UnsetUsedPreAggregatedTables ensures that no value is present for UsedPreAggregatedTables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


