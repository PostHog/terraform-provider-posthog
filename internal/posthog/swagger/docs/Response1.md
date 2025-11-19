# Response1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | **[]string** |  | 
**Error** | Pointer to **NullableString** | Query error. Returned only if &#39;explain&#39; or &#x60;modifiers.debug&#x60; is true. Throws an error otherwise. | [optional] 
**HasMore** | Pointer to **NullableBool** |  | [optional] 
**Hogql** | **NullableString** | Generated HogQL query. | 
**Limit** | **NullableInt32** |  | 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | **NullableInt32** |  | 
**QueryStatus** | Pointer to [**QueryStatus**](QueryStatus.md) |  | [optional] 
**ResolvedDateRange** | Pointer to [**ResolvedDateRangeResponse**](ResolvedDateRangeResponse.md) |  | [optional] 
**Results** | [**[]LLMTrace**](LLMTrace.md) |  | 
**Timings** | Pointer to [**[]QueryTiming**](QueryTiming.md) | Measured timings for different parts of the query generation process | [optional] 
**Types** | **[]map[string]interface{}** |  | 
**MissingActorsCount** | Pointer to **NullableInt32** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentTrendsQuery"]
**Clickhouse** | Pointer to **NullableString** | Executed ClickHouse query | [optional] 
**Explain** | Pointer to **[]string** | Query explanation output | [optional] 
**Metadata** | Pointer to [**HogQLMetadataResponse**](HogQLMetadataResponse.md) |  | [optional] 
**Query** | Pointer to **NullableString** | Input query string | [optional] 
**DateFrom** | Pointer to **NullableString** |  | [optional] 
**DateTo** | Pointer to **NullableString** |  | [optional] 
**SamplingRate** | Pointer to [**SamplingRate**](SamplingRate.md) |  | [optional] 
**UsedPreAggregatedTables** | Pointer to **NullableBool** |  | [optional] 
**CredibleIntervals** | **map[string][]float32** |  | 
**ExpectedLoss** | **float32** |  | 
**FunnelsQuery** | Pointer to [**FunnelsQuery**](FunnelsQuery.md) |  | [optional] 
**Insight** | **[]map[string]interface{}** |  | 
**Probability** | **map[string]float32** |  | 
**SignificanceCode** | [**ExperimentSignificanceCode**](ExperimentSignificanceCode.md) |  | 
**Significant** | **bool** |  | 
**StatsVersion** | Pointer to **NullableInt32** |  | [optional] 
**Variants** | [**[]ExperimentVariantTrendsBaseStats**](ExperimentVariantTrendsBaseStats.md) |  | 
**CountQuery** | Pointer to [**TrendsQuery**](TrendsQuery.md) |  | [optional] 
**ExposureQuery** | Pointer to [**TrendsQuery**](TrendsQuery.md) |  | [optional] 
**PValue** | **float32** |  | 

## Methods

### NewResponse1

`func NewResponse1(columns []string, hogql NullableString, limit NullableInt32, offset NullableInt32, results []LLMTrace, types []map[string]interface{}, credibleIntervals map[string][]float32, expectedLoss float32, insight []map[string]interface{}, probability map[string]float32, significanceCode ExperimentSignificanceCode, significant bool, variants []ExperimentVariantTrendsBaseStats, pValue float32, ) *Response1`

NewResponse1 instantiates a new Response1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponse1WithDefaults

`func NewResponse1WithDefaults() *Response1`

NewResponse1WithDefaults instantiates a new Response1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *Response1) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Response1) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Response1) SetColumns(v []string)`

SetColumns sets Columns field to given value.


### SetColumnsNil

`func (o *Response1) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *Response1) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetError

`func (o *Response1) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Response1) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Response1) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Response1) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *Response1) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *Response1) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetHasMore

`func (o *Response1) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *Response1) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *Response1) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *Response1) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### SetHasMoreNil

`func (o *Response1) SetHasMoreNil(b bool)`

 SetHasMoreNil sets the value for HasMore to be an explicit nil

### UnsetHasMore
`func (o *Response1) UnsetHasMore()`

UnsetHasMore ensures that no value is present for HasMore, not even an explicit nil
### GetHogql

`func (o *Response1) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *Response1) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *Response1) SetHogql(v string)`

SetHogql sets Hogql field to given value.


### SetHogqlNil

`func (o *Response1) SetHogqlNil(b bool)`

 SetHogqlNil sets the value for Hogql to be an explicit nil

### UnsetHogql
`func (o *Response1) UnsetHogql()`

UnsetHogql ensures that no value is present for Hogql, not even an explicit nil
### GetLimit

`func (o *Response1) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Response1) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Response1) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### SetLimitNil

`func (o *Response1) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *Response1) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *Response1) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *Response1) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *Response1) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *Response1) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *Response1) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Response1) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Response1) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### SetOffsetNil

`func (o *Response1) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *Response1) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetQueryStatus

`func (o *Response1) GetQueryStatus() QueryStatus`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *Response1) GetQueryStatusOk() (*QueryStatus, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *Response1) SetQueryStatus(v QueryStatus)`

SetQueryStatus sets QueryStatus field to given value.

### HasQueryStatus

`func (o *Response1) HasQueryStatus() bool`

HasQueryStatus returns a boolean if a field has been set.

### GetResolvedDateRange

`func (o *Response1) GetResolvedDateRange() ResolvedDateRangeResponse`

GetResolvedDateRange returns the ResolvedDateRange field if non-nil, zero value otherwise.

### GetResolvedDateRangeOk

`func (o *Response1) GetResolvedDateRangeOk() (*ResolvedDateRangeResponse, bool)`

GetResolvedDateRangeOk returns a tuple with the ResolvedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDateRange

`func (o *Response1) SetResolvedDateRange(v ResolvedDateRangeResponse)`

SetResolvedDateRange sets ResolvedDateRange field to given value.

### HasResolvedDateRange

`func (o *Response1) HasResolvedDateRange() bool`

HasResolvedDateRange returns a boolean if a field has been set.

### GetResults

`func (o *Response1) GetResults() []LLMTrace`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Response1) GetResultsOk() (*[]LLMTrace, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Response1) SetResults(v []LLMTrace)`

SetResults sets Results field to given value.


### GetTimings

`func (o *Response1) GetTimings() []QueryTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *Response1) GetTimingsOk() (*[]QueryTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *Response1) SetTimings(v []QueryTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *Response1) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *Response1) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *Response1) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil
### GetTypes

`func (o *Response1) GetTypes() []map[string]interface{}`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *Response1) GetTypesOk() (*[]map[string]interface{}, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *Response1) SetTypes(v []map[string]interface{})`

SetTypes sets Types field to given value.


### SetTypesNil

`func (o *Response1) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *Response1) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetMissingActorsCount

`func (o *Response1) GetMissingActorsCount() int32`

GetMissingActorsCount returns the MissingActorsCount field if non-nil, zero value otherwise.

### GetMissingActorsCountOk

`func (o *Response1) GetMissingActorsCountOk() (*int32, bool)`

GetMissingActorsCountOk returns a tuple with the MissingActorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingActorsCount

`func (o *Response1) SetMissingActorsCount(v int32)`

SetMissingActorsCount sets MissingActorsCount field to given value.

### HasMissingActorsCount

`func (o *Response1) HasMissingActorsCount() bool`

HasMissingActorsCount returns a boolean if a field has been set.

### SetMissingActorsCountNil

`func (o *Response1) SetMissingActorsCountNil(b bool)`

 SetMissingActorsCountNil sets the value for MissingActorsCount to be an explicit nil

### UnsetMissingActorsCount
`func (o *Response1) UnsetMissingActorsCount()`

UnsetMissingActorsCount ensures that no value is present for MissingActorsCount, not even an explicit nil
### GetKind

`func (o *Response1) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Response1) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Response1) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Response1) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetClickhouse

`func (o *Response1) GetClickhouse() string`

GetClickhouse returns the Clickhouse field if non-nil, zero value otherwise.

### GetClickhouseOk

`func (o *Response1) GetClickhouseOk() (*string, bool)`

GetClickhouseOk returns a tuple with the Clickhouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickhouse

`func (o *Response1) SetClickhouse(v string)`

SetClickhouse sets Clickhouse field to given value.

### HasClickhouse

`func (o *Response1) HasClickhouse() bool`

HasClickhouse returns a boolean if a field has been set.

### SetClickhouseNil

`func (o *Response1) SetClickhouseNil(b bool)`

 SetClickhouseNil sets the value for Clickhouse to be an explicit nil

### UnsetClickhouse
`func (o *Response1) UnsetClickhouse()`

UnsetClickhouse ensures that no value is present for Clickhouse, not even an explicit nil
### GetExplain

`func (o *Response1) GetExplain() []string`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *Response1) GetExplainOk() (*[]string, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *Response1) SetExplain(v []string)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *Response1) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### SetExplainNil

`func (o *Response1) SetExplainNil(b bool)`

 SetExplainNil sets the value for Explain to be an explicit nil

### UnsetExplain
`func (o *Response1) UnsetExplain()`

UnsetExplain ensures that no value is present for Explain, not even an explicit nil
### GetMetadata

`func (o *Response1) GetMetadata() HogQLMetadataResponse`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Response1) GetMetadataOk() (*HogQLMetadataResponse, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Response1) SetMetadata(v HogQLMetadataResponse)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Response1) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQuery

`func (o *Response1) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Response1) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Response1) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Response1) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *Response1) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *Response1) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetDateFrom

`func (o *Response1) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *Response1) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *Response1) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *Response1) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### SetDateFromNil

`func (o *Response1) SetDateFromNil(b bool)`

 SetDateFromNil sets the value for DateFrom to be an explicit nil

### UnsetDateFrom
`func (o *Response1) UnsetDateFrom()`

UnsetDateFrom ensures that no value is present for DateFrom, not even an explicit nil
### GetDateTo

`func (o *Response1) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *Response1) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *Response1) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *Response1) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### SetDateToNil

`func (o *Response1) SetDateToNil(b bool)`

 SetDateToNil sets the value for DateTo to be an explicit nil

### UnsetDateTo
`func (o *Response1) UnsetDateTo()`

UnsetDateTo ensures that no value is present for DateTo, not even an explicit nil
### GetSamplingRate

`func (o *Response1) GetSamplingRate() SamplingRate`

GetSamplingRate returns the SamplingRate field if non-nil, zero value otherwise.

### GetSamplingRateOk

`func (o *Response1) GetSamplingRateOk() (*SamplingRate, bool)`

GetSamplingRateOk returns a tuple with the SamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingRate

`func (o *Response1) SetSamplingRate(v SamplingRate)`

SetSamplingRate sets SamplingRate field to given value.

### HasSamplingRate

`func (o *Response1) HasSamplingRate() bool`

HasSamplingRate returns a boolean if a field has been set.

### GetUsedPreAggregatedTables

`func (o *Response1) GetUsedPreAggregatedTables() bool`

GetUsedPreAggregatedTables returns the UsedPreAggregatedTables field if non-nil, zero value otherwise.

### GetUsedPreAggregatedTablesOk

`func (o *Response1) GetUsedPreAggregatedTablesOk() (*bool, bool)`

GetUsedPreAggregatedTablesOk returns a tuple with the UsedPreAggregatedTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPreAggregatedTables

`func (o *Response1) SetUsedPreAggregatedTables(v bool)`

SetUsedPreAggregatedTables sets UsedPreAggregatedTables field to given value.

### HasUsedPreAggregatedTables

`func (o *Response1) HasUsedPreAggregatedTables() bool`

HasUsedPreAggregatedTables returns a boolean if a field has been set.

### SetUsedPreAggregatedTablesNil

`func (o *Response1) SetUsedPreAggregatedTablesNil(b bool)`

 SetUsedPreAggregatedTablesNil sets the value for UsedPreAggregatedTables to be an explicit nil

### UnsetUsedPreAggregatedTables
`func (o *Response1) UnsetUsedPreAggregatedTables()`

UnsetUsedPreAggregatedTables ensures that no value is present for UsedPreAggregatedTables, not even an explicit nil
### GetCredibleIntervals

`func (o *Response1) GetCredibleIntervals() map[string][]float32`

GetCredibleIntervals returns the CredibleIntervals field if non-nil, zero value otherwise.

### GetCredibleIntervalsOk

`func (o *Response1) GetCredibleIntervalsOk() (*map[string][]float32, bool)`

GetCredibleIntervalsOk returns a tuple with the CredibleIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredibleIntervals

`func (o *Response1) SetCredibleIntervals(v map[string][]float32)`

SetCredibleIntervals sets CredibleIntervals field to given value.


### GetExpectedLoss

`func (o *Response1) GetExpectedLoss() float32`

GetExpectedLoss returns the ExpectedLoss field if non-nil, zero value otherwise.

### GetExpectedLossOk

`func (o *Response1) GetExpectedLossOk() (*float32, bool)`

GetExpectedLossOk returns a tuple with the ExpectedLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedLoss

`func (o *Response1) SetExpectedLoss(v float32)`

SetExpectedLoss sets ExpectedLoss field to given value.


### GetFunnelsQuery

`func (o *Response1) GetFunnelsQuery() FunnelsQuery`

GetFunnelsQuery returns the FunnelsQuery field if non-nil, zero value otherwise.

### GetFunnelsQueryOk

`func (o *Response1) GetFunnelsQueryOk() (*FunnelsQuery, bool)`

GetFunnelsQueryOk returns a tuple with the FunnelsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelsQuery

`func (o *Response1) SetFunnelsQuery(v FunnelsQuery)`

SetFunnelsQuery sets FunnelsQuery field to given value.

### HasFunnelsQuery

`func (o *Response1) HasFunnelsQuery() bool`

HasFunnelsQuery returns a boolean if a field has been set.

### GetInsight

`func (o *Response1) GetInsight() []map[string]interface{}`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *Response1) GetInsightOk() (*[]map[string]interface{}, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *Response1) SetInsight(v []map[string]interface{})`

SetInsight sets Insight field to given value.


### GetProbability

`func (o *Response1) GetProbability() map[string]float32`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *Response1) GetProbabilityOk() (*map[string]float32, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *Response1) SetProbability(v map[string]float32)`

SetProbability sets Probability field to given value.


### GetSignificanceCode

`func (o *Response1) GetSignificanceCode() ExperimentSignificanceCode`

GetSignificanceCode returns the SignificanceCode field if non-nil, zero value otherwise.

### GetSignificanceCodeOk

`func (o *Response1) GetSignificanceCodeOk() (*ExperimentSignificanceCode, bool)`

GetSignificanceCodeOk returns a tuple with the SignificanceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificanceCode

`func (o *Response1) SetSignificanceCode(v ExperimentSignificanceCode)`

SetSignificanceCode sets SignificanceCode field to given value.


### GetSignificant

`func (o *Response1) GetSignificant() bool`

GetSignificant returns the Significant field if non-nil, zero value otherwise.

### GetSignificantOk

`func (o *Response1) GetSignificantOk() (*bool, bool)`

GetSignificantOk returns a tuple with the Significant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificant

`func (o *Response1) SetSignificant(v bool)`

SetSignificant sets Significant field to given value.


### GetStatsVersion

`func (o *Response1) GetStatsVersion() int32`

GetStatsVersion returns the StatsVersion field if non-nil, zero value otherwise.

### GetStatsVersionOk

`func (o *Response1) GetStatsVersionOk() (*int32, bool)`

GetStatsVersionOk returns a tuple with the StatsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsVersion

`func (o *Response1) SetStatsVersion(v int32)`

SetStatsVersion sets StatsVersion field to given value.

### HasStatsVersion

`func (o *Response1) HasStatsVersion() bool`

HasStatsVersion returns a boolean if a field has been set.

### SetStatsVersionNil

`func (o *Response1) SetStatsVersionNil(b bool)`

 SetStatsVersionNil sets the value for StatsVersion to be an explicit nil

### UnsetStatsVersion
`func (o *Response1) UnsetStatsVersion()`

UnsetStatsVersion ensures that no value is present for StatsVersion, not even an explicit nil
### GetVariants

`func (o *Response1) GetVariants() []ExperimentVariantTrendsBaseStats`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *Response1) GetVariantsOk() (*[]ExperimentVariantTrendsBaseStats, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *Response1) SetVariants(v []ExperimentVariantTrendsBaseStats)`

SetVariants sets Variants field to given value.


### GetCountQuery

`func (o *Response1) GetCountQuery() TrendsQuery`

GetCountQuery returns the CountQuery field if non-nil, zero value otherwise.

### GetCountQueryOk

`func (o *Response1) GetCountQueryOk() (*TrendsQuery, bool)`

GetCountQueryOk returns a tuple with the CountQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountQuery

`func (o *Response1) SetCountQuery(v TrendsQuery)`

SetCountQuery sets CountQuery field to given value.

### HasCountQuery

`func (o *Response1) HasCountQuery() bool`

HasCountQuery returns a boolean if a field has been set.

### GetExposureQuery

`func (o *Response1) GetExposureQuery() TrendsQuery`

GetExposureQuery returns the ExposureQuery field if non-nil, zero value otherwise.

### GetExposureQueryOk

`func (o *Response1) GetExposureQueryOk() (*TrendsQuery, bool)`

GetExposureQueryOk returns a tuple with the ExposureQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureQuery

`func (o *Response1) SetExposureQuery(v TrendsQuery)`

SetExposureQuery sets ExposureQuery field to given value.

### HasExposureQuery

`func (o *Response1) HasExposureQuery() bool`

HasExposureQuery returns a boolean if a field has been set.

### GetPValue

`func (o *Response1) GetPValue() float32`

GetPValue returns the PValue field if non-nil, zero value otherwise.

### GetPValueOk

`func (o *Response1) GetPValueOk() (*float32, bool)`

GetPValueOk returns a tuple with the PValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPValue

`func (o *Response1) SetPValue(v float32)`

SetPValue sets PValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


