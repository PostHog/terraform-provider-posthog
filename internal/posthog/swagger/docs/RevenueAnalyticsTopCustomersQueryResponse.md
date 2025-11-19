# RevenueAnalyticsTopCustomersQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]string** |  | [optional] 
**Error** | Pointer to **NullableString** | Query error. Returned only if &#39;explain&#39; or &#x60;modifiers.debug&#x60; is true. Throws an error otherwise. | [optional] 
**Hogql** | Pointer to **NullableString** | Generated HogQL query. | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**QueryStatus** | Pointer to [**QueryStatus**](QueryStatus.md) |  | [optional] 
**ResolvedDateRange** | Pointer to [**ResolvedDateRangeResponse**](ResolvedDateRangeResponse.md) |  | [optional] 
**Results** | **interface{}** |  | 
**Timings** | Pointer to [**[]QueryTiming**](QueryTiming.md) | Measured timings for different parts of the query generation process | [optional] 

## Methods

### NewRevenueAnalyticsTopCustomersQueryResponse

`func NewRevenueAnalyticsTopCustomersQueryResponse(results interface{}, ) *RevenueAnalyticsTopCustomersQueryResponse`

NewRevenueAnalyticsTopCustomersQueryResponse instantiates a new RevenueAnalyticsTopCustomersQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueAnalyticsTopCustomersQueryResponseWithDefaults

`func NewRevenueAnalyticsTopCustomersQueryResponseWithDefaults() *RevenueAnalyticsTopCustomersQueryResponse`

NewRevenueAnalyticsTopCustomersQueryResponseWithDefaults instantiates a new RevenueAnalyticsTopCustomersQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *RevenueAnalyticsTopCustomersQueryResponse) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *RevenueAnalyticsTopCustomersQueryResponse) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *RevenueAnalyticsTopCustomersQueryResponse) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *RevenueAnalyticsTopCustomersQueryResponse) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetError

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RevenueAnalyticsTopCustomersQueryResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RevenueAnalyticsTopCustomersQueryResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *RevenueAnalyticsTopCustomersQueryResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *RevenueAnalyticsTopCustomersQueryResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetHogql

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *RevenueAnalyticsTopCustomersQueryResponse) SetHogql(v string)`

SetHogql sets Hogql field to given value.

### HasHogql

`func (o *RevenueAnalyticsTopCustomersQueryResponse) HasHogql() bool`

HasHogql returns a boolean if a field has been set.

### SetHogqlNil

`func (o *RevenueAnalyticsTopCustomersQueryResponse) SetHogqlNil(b bool)`

 SetHogqlNil sets the value for Hogql to be an explicit nil

### UnsetHogql
`func (o *RevenueAnalyticsTopCustomersQueryResponse) UnsetHogql()`

UnsetHogql ensures that no value is present for Hogql, not even an explicit nil
### GetModifiers

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *RevenueAnalyticsTopCustomersQueryResponse) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *RevenueAnalyticsTopCustomersQueryResponse) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetQueryStatus

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetQueryStatus() QueryStatus`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetQueryStatusOk() (*QueryStatus, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *RevenueAnalyticsTopCustomersQueryResponse) SetQueryStatus(v QueryStatus)`

SetQueryStatus sets QueryStatus field to given value.

### HasQueryStatus

`func (o *RevenueAnalyticsTopCustomersQueryResponse) HasQueryStatus() bool`

HasQueryStatus returns a boolean if a field has been set.

### GetResolvedDateRange

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetResolvedDateRange() ResolvedDateRangeResponse`

GetResolvedDateRange returns the ResolvedDateRange field if non-nil, zero value otherwise.

### GetResolvedDateRangeOk

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetResolvedDateRangeOk() (*ResolvedDateRangeResponse, bool)`

GetResolvedDateRangeOk returns a tuple with the ResolvedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDateRange

`func (o *RevenueAnalyticsTopCustomersQueryResponse) SetResolvedDateRange(v ResolvedDateRangeResponse)`

SetResolvedDateRange sets ResolvedDateRange field to given value.

### HasResolvedDateRange

`func (o *RevenueAnalyticsTopCustomersQueryResponse) HasResolvedDateRange() bool`

HasResolvedDateRange returns a boolean if a field has been set.

### GetResults

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetResults() interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetResultsOk() (*interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RevenueAnalyticsTopCustomersQueryResponse) SetResults(v interface{})`

SetResults sets Results field to given value.


### SetResultsNil

`func (o *RevenueAnalyticsTopCustomersQueryResponse) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *RevenueAnalyticsTopCustomersQueryResponse) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetTimings

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetTimings() []QueryTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *RevenueAnalyticsTopCustomersQueryResponse) GetTimingsOk() (*[]QueryTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *RevenueAnalyticsTopCustomersQueryResponse) SetTimings(v []QueryTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *RevenueAnalyticsTopCustomersQueryResponse) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *RevenueAnalyticsTopCustomersQueryResponse) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *RevenueAnalyticsTopCustomersQueryResponse) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


