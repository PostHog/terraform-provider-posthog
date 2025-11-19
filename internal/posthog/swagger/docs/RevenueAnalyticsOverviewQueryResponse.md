# RevenueAnalyticsOverviewQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **NullableString** | Query error. Returned only if &#39;explain&#39; or &#x60;modifiers.debug&#x60; is true. Throws an error otherwise. | [optional] 
**Hogql** | Pointer to **NullableString** | Generated HogQL query. | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**QueryStatus** | Pointer to [**QueryStatus**](QueryStatus.md) |  | [optional] 
**ResolvedDateRange** | Pointer to [**ResolvedDateRangeResponse**](ResolvedDateRangeResponse.md) |  | [optional] 
**Results** | [**[]RevenueAnalyticsOverviewItem**](RevenueAnalyticsOverviewItem.md) |  | 
**Timings** | Pointer to [**[]QueryTiming**](QueryTiming.md) | Measured timings for different parts of the query generation process | [optional] 

## Methods

### NewRevenueAnalyticsOverviewQueryResponse

`func NewRevenueAnalyticsOverviewQueryResponse(results []RevenueAnalyticsOverviewItem, ) *RevenueAnalyticsOverviewQueryResponse`

NewRevenueAnalyticsOverviewQueryResponse instantiates a new RevenueAnalyticsOverviewQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueAnalyticsOverviewQueryResponseWithDefaults

`func NewRevenueAnalyticsOverviewQueryResponseWithDefaults() *RevenueAnalyticsOverviewQueryResponse`

NewRevenueAnalyticsOverviewQueryResponseWithDefaults instantiates a new RevenueAnalyticsOverviewQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *RevenueAnalyticsOverviewQueryResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RevenueAnalyticsOverviewQueryResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RevenueAnalyticsOverviewQueryResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RevenueAnalyticsOverviewQueryResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *RevenueAnalyticsOverviewQueryResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *RevenueAnalyticsOverviewQueryResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetHogql

`func (o *RevenueAnalyticsOverviewQueryResponse) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *RevenueAnalyticsOverviewQueryResponse) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *RevenueAnalyticsOverviewQueryResponse) SetHogql(v string)`

SetHogql sets Hogql field to given value.

### HasHogql

`func (o *RevenueAnalyticsOverviewQueryResponse) HasHogql() bool`

HasHogql returns a boolean if a field has been set.

### SetHogqlNil

`func (o *RevenueAnalyticsOverviewQueryResponse) SetHogqlNil(b bool)`

 SetHogqlNil sets the value for Hogql to be an explicit nil

### UnsetHogql
`func (o *RevenueAnalyticsOverviewQueryResponse) UnsetHogql()`

UnsetHogql ensures that no value is present for Hogql, not even an explicit nil
### GetModifiers

`func (o *RevenueAnalyticsOverviewQueryResponse) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *RevenueAnalyticsOverviewQueryResponse) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *RevenueAnalyticsOverviewQueryResponse) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *RevenueAnalyticsOverviewQueryResponse) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetQueryStatus

`func (o *RevenueAnalyticsOverviewQueryResponse) GetQueryStatus() QueryStatus`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *RevenueAnalyticsOverviewQueryResponse) GetQueryStatusOk() (*QueryStatus, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *RevenueAnalyticsOverviewQueryResponse) SetQueryStatus(v QueryStatus)`

SetQueryStatus sets QueryStatus field to given value.

### HasQueryStatus

`func (o *RevenueAnalyticsOverviewQueryResponse) HasQueryStatus() bool`

HasQueryStatus returns a boolean if a field has been set.

### GetResolvedDateRange

`func (o *RevenueAnalyticsOverviewQueryResponse) GetResolvedDateRange() ResolvedDateRangeResponse`

GetResolvedDateRange returns the ResolvedDateRange field if non-nil, zero value otherwise.

### GetResolvedDateRangeOk

`func (o *RevenueAnalyticsOverviewQueryResponse) GetResolvedDateRangeOk() (*ResolvedDateRangeResponse, bool)`

GetResolvedDateRangeOk returns a tuple with the ResolvedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDateRange

`func (o *RevenueAnalyticsOverviewQueryResponse) SetResolvedDateRange(v ResolvedDateRangeResponse)`

SetResolvedDateRange sets ResolvedDateRange field to given value.

### HasResolvedDateRange

`func (o *RevenueAnalyticsOverviewQueryResponse) HasResolvedDateRange() bool`

HasResolvedDateRange returns a boolean if a field has been set.

### GetResults

`func (o *RevenueAnalyticsOverviewQueryResponse) GetResults() []RevenueAnalyticsOverviewItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RevenueAnalyticsOverviewQueryResponse) GetResultsOk() (*[]RevenueAnalyticsOverviewItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RevenueAnalyticsOverviewQueryResponse) SetResults(v []RevenueAnalyticsOverviewItem)`

SetResults sets Results field to given value.


### GetTimings

`func (o *RevenueAnalyticsOverviewQueryResponse) GetTimings() []QueryTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *RevenueAnalyticsOverviewQueryResponse) GetTimingsOk() (*[]QueryTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *RevenueAnalyticsOverviewQueryResponse) SetTimings(v []QueryTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *RevenueAnalyticsOverviewQueryResponse) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *RevenueAnalyticsOverviewQueryResponse) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *RevenueAnalyticsOverviewQueryResponse) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


