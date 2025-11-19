# MarketingAnalyticsAggregatedQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **NullableString** | Query error. Returned only if &#39;explain&#39; or &#x60;modifiers.debug&#x60; is true. Throws an error otherwise. | [optional] 
**Hogql** | Pointer to **NullableString** | Generated HogQL query. | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**QueryStatus** | Pointer to [**QueryStatus**](QueryStatus.md) |  | [optional] 
**ResolvedDateRange** | Pointer to [**ResolvedDateRangeResponse**](ResolvedDateRangeResponse.md) |  | [optional] 
**Results** | [**map[string]MarketingAnalyticsItem**](MarketingAnalyticsItem.md) |  | 
**SamplingRate** | Pointer to [**SamplingRate**](SamplingRate.md) |  | [optional] 
**Timings** | Pointer to [**[]QueryTiming**](QueryTiming.md) | Measured timings for different parts of the query generation process | [optional] 

## Methods

### NewMarketingAnalyticsAggregatedQueryResponse

`func NewMarketingAnalyticsAggregatedQueryResponse(results map[string]MarketingAnalyticsItem, ) *MarketingAnalyticsAggregatedQueryResponse`

NewMarketingAnalyticsAggregatedQueryResponse instantiates a new MarketingAnalyticsAggregatedQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingAnalyticsAggregatedQueryResponseWithDefaults

`func NewMarketingAnalyticsAggregatedQueryResponseWithDefaults() *MarketingAnalyticsAggregatedQueryResponse`

NewMarketingAnalyticsAggregatedQueryResponseWithDefaults instantiates a new MarketingAnalyticsAggregatedQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MarketingAnalyticsAggregatedQueryResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MarketingAnalyticsAggregatedQueryResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *MarketingAnalyticsAggregatedQueryResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *MarketingAnalyticsAggregatedQueryResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetHogql

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *MarketingAnalyticsAggregatedQueryResponse) SetHogql(v string)`

SetHogql sets Hogql field to given value.

### HasHogql

`func (o *MarketingAnalyticsAggregatedQueryResponse) HasHogql() bool`

HasHogql returns a boolean if a field has been set.

### SetHogqlNil

`func (o *MarketingAnalyticsAggregatedQueryResponse) SetHogqlNil(b bool)`

 SetHogqlNil sets the value for Hogql to be an explicit nil

### UnsetHogql
`func (o *MarketingAnalyticsAggregatedQueryResponse) UnsetHogql()`

UnsetHogql ensures that no value is present for Hogql, not even an explicit nil
### GetModifiers

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *MarketingAnalyticsAggregatedQueryResponse) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *MarketingAnalyticsAggregatedQueryResponse) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetQueryStatus

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetQueryStatus() QueryStatus`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetQueryStatusOk() (*QueryStatus, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *MarketingAnalyticsAggregatedQueryResponse) SetQueryStatus(v QueryStatus)`

SetQueryStatus sets QueryStatus field to given value.

### HasQueryStatus

`func (o *MarketingAnalyticsAggregatedQueryResponse) HasQueryStatus() bool`

HasQueryStatus returns a boolean if a field has been set.

### GetResolvedDateRange

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetResolvedDateRange() ResolvedDateRangeResponse`

GetResolvedDateRange returns the ResolvedDateRange field if non-nil, zero value otherwise.

### GetResolvedDateRangeOk

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetResolvedDateRangeOk() (*ResolvedDateRangeResponse, bool)`

GetResolvedDateRangeOk returns a tuple with the ResolvedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDateRange

`func (o *MarketingAnalyticsAggregatedQueryResponse) SetResolvedDateRange(v ResolvedDateRangeResponse)`

SetResolvedDateRange sets ResolvedDateRange field to given value.

### HasResolvedDateRange

`func (o *MarketingAnalyticsAggregatedQueryResponse) HasResolvedDateRange() bool`

HasResolvedDateRange returns a boolean if a field has been set.

### GetResults

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetResults() map[string]MarketingAnalyticsItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetResultsOk() (*map[string]MarketingAnalyticsItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MarketingAnalyticsAggregatedQueryResponse) SetResults(v map[string]MarketingAnalyticsItem)`

SetResults sets Results field to given value.


### GetSamplingRate

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetSamplingRate() SamplingRate`

GetSamplingRate returns the SamplingRate field if non-nil, zero value otherwise.

### GetSamplingRateOk

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetSamplingRateOk() (*SamplingRate, bool)`

GetSamplingRateOk returns a tuple with the SamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingRate

`func (o *MarketingAnalyticsAggregatedQueryResponse) SetSamplingRate(v SamplingRate)`

SetSamplingRate sets SamplingRate field to given value.

### HasSamplingRate

`func (o *MarketingAnalyticsAggregatedQueryResponse) HasSamplingRate() bool`

HasSamplingRate returns a boolean if a field has been set.

### GetTimings

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetTimings() []QueryTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *MarketingAnalyticsAggregatedQueryResponse) GetTimingsOk() (*[]QueryTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *MarketingAnalyticsAggregatedQueryResponse) SetTimings(v []QueryTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *MarketingAnalyticsAggregatedQueryResponse) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *MarketingAnalyticsAggregatedQueryResponse) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *MarketingAnalyticsAggregatedQueryResponse) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


