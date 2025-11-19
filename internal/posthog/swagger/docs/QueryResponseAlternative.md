# QueryResponseAlternative

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
**Results** | [**[]UsageMetric**](UsageMetric.md) |  | 
**Timings** | Pointer to [**[]QueryTiming**](QueryTiming.md) | Measured timings for different parts of the query generation process | [optional] 
**Types** | **[]map[string]interface{}** |  | 
**MissingActorsCount** | Pointer to **NullableInt32** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentTrendsQuery"]
**Breakdown** | Pointer to [**[]BreakdownItem**](BreakdownItem.md) |  | [optional] 
**Breakdowns** | Pointer to [**[]MultipleBreakdownOptions**](MultipleBreakdownOptions.md) |  | [optional] 
**Compare** | Pointer to [**[]CompareItem**](CompareItem.md) |  | [optional] 
**Day** | Pointer to [**[]DayItem**](DayItem.md) |  | [optional] 
**Interval** | Pointer to [**[]IntervalItem**](IntervalItem.md) |  | [optional] 
**Series** | Pointer to [**[]Series**](Series.md) |  | [optional] 
**Status** | [**ExternalQueryStatus**](ExternalQueryStatus.md) |  | 
**Bytecode** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ColoredBytecode** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Stdout** | Pointer to **NullableString** |  | [optional] 
**Clickhouse** | Pointer to **NullableString** | Executed ClickHouse query | [optional] 
**Explain** | Pointer to **[]string** | Query explanation output | [optional] 
**Metadata** | Pointer to [**HogQLMetadataResponse**](HogQLMetadataResponse.md) |  | [optional] 
**Query** | Pointer to **NullableString** | Input query string | [optional] 
**ChTableNames** | Pointer to **[]string** |  | [optional] 
**Errors** | [**[]HogQLNotice**](HogQLNotice.md) |  | 
**IsUsingIndices** | Pointer to [**QueryIndexUsage**](QueryIndexUsage.md) |  | [optional] 
**IsValid** | Pointer to **NullableBool** |  | [optional] 
**Notices** | [**[]HogQLNotice**](HogQLNotice.md) |  | 
**TableNames** | Pointer to **[]string** |  | [optional] 
**Warnings** | [**[]HogQLNotice**](HogQLNotice.md) |  | 
**IncompleteList** | **bool** | Whether or not the suggestions returned are complete | 
**Suggestions** | [**[]AutocompleteCompletionItem**](AutocompleteCompletionItem.md) |  | 
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
**Baseline** | Pointer to [**ExperimentStatsBaseValidated**](ExperimentStatsBaseValidated.md) |  | [optional] 
**BreakdownResults** | Pointer to [**[]ExperimentBreakdownResult**](ExperimentBreakdownResult.md) | Results grouped by breakdown value. When present, baseline and variant_results contain aggregated data. | [optional] 
**Metric** | Pointer to [**NullableMetric1**](Metric1.md) |  | [optional] [default to null]
**VariantResults** | Pointer to [**NullableVariantResults**](VariantResults.md) |  | [optional] [default to null]
**DateRange** | [**DateRange**](DateRange.md) |  | 
**Timeseries** | [**[]ExperimentExposureTimeSeries**](ExperimentExposureTimeSeries.md) |  | 
**TotalExposures** | **map[string]float32** |  | 
**DateFrom** | Pointer to **NullableString** |  | [optional] 
**DateTo** | Pointer to **NullableString** |  | [optional] 
**SamplingRate** | Pointer to [**SamplingRate**](SamplingRate.md) |  | [optional] 
**UsedPreAggregatedTables** | Pointer to **NullableBool** |  | [optional] 
**Data** | **map[string]interface{}** |  | 
**IsUdf** | Pointer to **NullableBool** |  | [optional] 
**Joins** | [**[]DataWarehouseViewLink**](DataWarehouseViewLink.md) |  | 
**Tables** | [**map[string]TablesValue**](TablesValue.md) |  | 
**Questions** | **[]string** |  | 

## Methods

### NewQueryResponseAlternative

`func NewQueryResponseAlternative(columns []string, hogql NullableString, limit NullableInt32, offset NullableInt32, results []UsageMetric, types []map[string]interface{}, status ExternalQueryStatus, errors []HogQLNotice, notices []HogQLNotice, warnings []HogQLNotice, incompleteList bool, suggestions []AutocompleteCompletionItem, credibleIntervals map[string][]float32, expectedLoss float32, insight []map[string]interface{}, probability map[string]float32, significanceCode ExperimentSignificanceCode, significant bool, variants []ExperimentVariantTrendsBaseStats, pValue float32, dateRange DateRange, timeseries []ExperimentExposureTimeSeries, totalExposures map[string]float32, data map[string]interface{}, joins []DataWarehouseViewLink, tables map[string]TablesValue, questions []string, ) *QueryResponseAlternative`

NewQueryResponseAlternative instantiates a new QueryResponseAlternative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseAlternativeWithDefaults

`func NewQueryResponseAlternativeWithDefaults() *QueryResponseAlternative`

NewQueryResponseAlternativeWithDefaults instantiates a new QueryResponseAlternative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *QueryResponseAlternative) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *QueryResponseAlternative) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *QueryResponseAlternative) SetColumns(v []string)`

SetColumns sets Columns field to given value.


### SetColumnsNil

`func (o *QueryResponseAlternative) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *QueryResponseAlternative) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetError

`func (o *QueryResponseAlternative) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *QueryResponseAlternative) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *QueryResponseAlternative) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *QueryResponseAlternative) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *QueryResponseAlternative) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *QueryResponseAlternative) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetHasMore

`func (o *QueryResponseAlternative) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *QueryResponseAlternative) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *QueryResponseAlternative) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *QueryResponseAlternative) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### SetHasMoreNil

`func (o *QueryResponseAlternative) SetHasMoreNil(b bool)`

 SetHasMoreNil sets the value for HasMore to be an explicit nil

### UnsetHasMore
`func (o *QueryResponseAlternative) UnsetHasMore()`

UnsetHasMore ensures that no value is present for HasMore, not even an explicit nil
### GetHogql

`func (o *QueryResponseAlternative) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *QueryResponseAlternative) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *QueryResponseAlternative) SetHogql(v string)`

SetHogql sets Hogql field to given value.


### SetHogqlNil

`func (o *QueryResponseAlternative) SetHogqlNil(b bool)`

 SetHogqlNil sets the value for Hogql to be an explicit nil

### UnsetHogql
`func (o *QueryResponseAlternative) UnsetHogql()`

UnsetHogql ensures that no value is present for Hogql, not even an explicit nil
### GetLimit

`func (o *QueryResponseAlternative) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryResponseAlternative) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryResponseAlternative) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### SetLimitNil

`func (o *QueryResponseAlternative) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *QueryResponseAlternative) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *QueryResponseAlternative) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *QueryResponseAlternative) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *QueryResponseAlternative) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *QueryResponseAlternative) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *QueryResponseAlternative) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QueryResponseAlternative) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QueryResponseAlternative) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### SetOffsetNil

`func (o *QueryResponseAlternative) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *QueryResponseAlternative) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetQueryStatus

`func (o *QueryResponseAlternative) GetQueryStatus() QueryStatus`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *QueryResponseAlternative) GetQueryStatusOk() (*QueryStatus, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *QueryResponseAlternative) SetQueryStatus(v QueryStatus)`

SetQueryStatus sets QueryStatus field to given value.

### HasQueryStatus

`func (o *QueryResponseAlternative) HasQueryStatus() bool`

HasQueryStatus returns a boolean if a field has been set.

### GetResolvedDateRange

`func (o *QueryResponseAlternative) GetResolvedDateRange() ResolvedDateRangeResponse`

GetResolvedDateRange returns the ResolvedDateRange field if non-nil, zero value otherwise.

### GetResolvedDateRangeOk

`func (o *QueryResponseAlternative) GetResolvedDateRangeOk() (*ResolvedDateRangeResponse, bool)`

GetResolvedDateRangeOk returns a tuple with the ResolvedDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDateRange

`func (o *QueryResponseAlternative) SetResolvedDateRange(v ResolvedDateRangeResponse)`

SetResolvedDateRange sets ResolvedDateRange field to given value.

### HasResolvedDateRange

`func (o *QueryResponseAlternative) HasResolvedDateRange() bool`

HasResolvedDateRange returns a boolean if a field has been set.

### GetResults

`func (o *QueryResponseAlternative) GetResults() []UsageMetric`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *QueryResponseAlternative) GetResultsOk() (*[]UsageMetric, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *QueryResponseAlternative) SetResults(v []UsageMetric)`

SetResults sets Results field to given value.


### GetTimings

`func (o *QueryResponseAlternative) GetTimings() []QueryTiming`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *QueryResponseAlternative) GetTimingsOk() (*[]QueryTiming, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *QueryResponseAlternative) SetTimings(v []QueryTiming)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *QueryResponseAlternative) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *QueryResponseAlternative) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *QueryResponseAlternative) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil
### GetTypes

`func (o *QueryResponseAlternative) GetTypes() []map[string]interface{}`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *QueryResponseAlternative) GetTypesOk() (*[]map[string]interface{}, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *QueryResponseAlternative) SetTypes(v []map[string]interface{})`

SetTypes sets Types field to given value.


### SetTypesNil

`func (o *QueryResponseAlternative) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *QueryResponseAlternative) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetMissingActorsCount

`func (o *QueryResponseAlternative) GetMissingActorsCount() int32`

GetMissingActorsCount returns the MissingActorsCount field if non-nil, zero value otherwise.

### GetMissingActorsCountOk

`func (o *QueryResponseAlternative) GetMissingActorsCountOk() (*int32, bool)`

GetMissingActorsCountOk returns a tuple with the MissingActorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingActorsCount

`func (o *QueryResponseAlternative) SetMissingActorsCount(v int32)`

SetMissingActorsCount sets MissingActorsCount field to given value.

### HasMissingActorsCount

`func (o *QueryResponseAlternative) HasMissingActorsCount() bool`

HasMissingActorsCount returns a boolean if a field has been set.

### SetMissingActorsCountNil

`func (o *QueryResponseAlternative) SetMissingActorsCountNil(b bool)`

 SetMissingActorsCountNil sets the value for MissingActorsCount to be an explicit nil

### UnsetMissingActorsCount
`func (o *QueryResponseAlternative) UnsetMissingActorsCount()`

UnsetMissingActorsCount ensures that no value is present for MissingActorsCount, not even an explicit nil
### GetKind

`func (o *QueryResponseAlternative) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QueryResponseAlternative) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QueryResponseAlternative) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *QueryResponseAlternative) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetBreakdown

`func (o *QueryResponseAlternative) GetBreakdown() []BreakdownItem`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *QueryResponseAlternative) GetBreakdownOk() (*[]BreakdownItem, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *QueryResponseAlternative) SetBreakdown(v []BreakdownItem)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *QueryResponseAlternative) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### SetBreakdownNil

`func (o *QueryResponseAlternative) SetBreakdownNil(b bool)`

 SetBreakdownNil sets the value for Breakdown to be an explicit nil

### UnsetBreakdown
`func (o *QueryResponseAlternative) UnsetBreakdown()`

UnsetBreakdown ensures that no value is present for Breakdown, not even an explicit nil
### GetBreakdowns

`func (o *QueryResponseAlternative) GetBreakdowns() []MultipleBreakdownOptions`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *QueryResponseAlternative) GetBreakdownsOk() (*[]MultipleBreakdownOptions, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *QueryResponseAlternative) SetBreakdowns(v []MultipleBreakdownOptions)`

SetBreakdowns sets Breakdowns field to given value.

### HasBreakdowns

`func (o *QueryResponseAlternative) HasBreakdowns() bool`

HasBreakdowns returns a boolean if a field has been set.

### SetBreakdownsNil

`func (o *QueryResponseAlternative) SetBreakdownsNil(b bool)`

 SetBreakdownsNil sets the value for Breakdowns to be an explicit nil

### UnsetBreakdowns
`func (o *QueryResponseAlternative) UnsetBreakdowns()`

UnsetBreakdowns ensures that no value is present for Breakdowns, not even an explicit nil
### GetCompare

`func (o *QueryResponseAlternative) GetCompare() []CompareItem`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *QueryResponseAlternative) GetCompareOk() (*[]CompareItem, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *QueryResponseAlternative) SetCompare(v []CompareItem)`

SetCompare sets Compare field to given value.

### HasCompare

`func (o *QueryResponseAlternative) HasCompare() bool`

HasCompare returns a boolean if a field has been set.

### SetCompareNil

`func (o *QueryResponseAlternative) SetCompareNil(b bool)`

 SetCompareNil sets the value for Compare to be an explicit nil

### UnsetCompare
`func (o *QueryResponseAlternative) UnsetCompare()`

UnsetCompare ensures that no value is present for Compare, not even an explicit nil
### GetDay

`func (o *QueryResponseAlternative) GetDay() []DayItem`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *QueryResponseAlternative) GetDayOk() (*[]DayItem, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *QueryResponseAlternative) SetDay(v []DayItem)`

SetDay sets Day field to given value.

### HasDay

`func (o *QueryResponseAlternative) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *QueryResponseAlternative) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *QueryResponseAlternative) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetInterval

`func (o *QueryResponseAlternative) GetInterval() []IntervalItem`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *QueryResponseAlternative) GetIntervalOk() (*[]IntervalItem, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *QueryResponseAlternative) SetInterval(v []IntervalItem)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *QueryResponseAlternative) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *QueryResponseAlternative) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *QueryResponseAlternative) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetSeries

`func (o *QueryResponseAlternative) GetSeries() []Series`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *QueryResponseAlternative) GetSeriesOk() (*[]Series, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *QueryResponseAlternative) SetSeries(v []Series)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *QueryResponseAlternative) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeriesNil

`func (o *QueryResponseAlternative) SetSeriesNil(b bool)`

 SetSeriesNil sets the value for Series to be an explicit nil

### UnsetSeries
`func (o *QueryResponseAlternative) UnsetSeries()`

UnsetSeries ensures that no value is present for Series, not even an explicit nil
### GetStatus

`func (o *QueryResponseAlternative) GetStatus() ExternalQueryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryResponseAlternative) GetStatusOk() (*ExternalQueryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryResponseAlternative) SetStatus(v ExternalQueryStatus)`

SetStatus sets Status field to given value.


### GetBytecode

`func (o *QueryResponseAlternative) GetBytecode() []map[string]interface{}`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *QueryResponseAlternative) GetBytecodeOk() (*[]map[string]interface{}, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *QueryResponseAlternative) SetBytecode(v []map[string]interface{})`

SetBytecode sets Bytecode field to given value.

### HasBytecode

`func (o *QueryResponseAlternative) HasBytecode() bool`

HasBytecode returns a boolean if a field has been set.

### SetBytecodeNil

`func (o *QueryResponseAlternative) SetBytecodeNil(b bool)`

 SetBytecodeNil sets the value for Bytecode to be an explicit nil

### UnsetBytecode
`func (o *QueryResponseAlternative) UnsetBytecode()`

UnsetBytecode ensures that no value is present for Bytecode, not even an explicit nil
### GetColoredBytecode

`func (o *QueryResponseAlternative) GetColoredBytecode() []map[string]interface{}`

GetColoredBytecode returns the ColoredBytecode field if non-nil, zero value otherwise.

### GetColoredBytecodeOk

`func (o *QueryResponseAlternative) GetColoredBytecodeOk() (*[]map[string]interface{}, bool)`

GetColoredBytecodeOk returns a tuple with the ColoredBytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColoredBytecode

`func (o *QueryResponseAlternative) SetColoredBytecode(v []map[string]interface{})`

SetColoredBytecode sets ColoredBytecode field to given value.

### HasColoredBytecode

`func (o *QueryResponseAlternative) HasColoredBytecode() bool`

HasColoredBytecode returns a boolean if a field has been set.

### SetColoredBytecodeNil

`func (o *QueryResponseAlternative) SetColoredBytecodeNil(b bool)`

 SetColoredBytecodeNil sets the value for ColoredBytecode to be an explicit nil

### UnsetColoredBytecode
`func (o *QueryResponseAlternative) UnsetColoredBytecode()`

UnsetColoredBytecode ensures that no value is present for ColoredBytecode, not even an explicit nil
### GetStdout

`func (o *QueryResponseAlternative) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *QueryResponseAlternative) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *QueryResponseAlternative) SetStdout(v string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *QueryResponseAlternative) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### SetStdoutNil

`func (o *QueryResponseAlternative) SetStdoutNil(b bool)`

 SetStdoutNil sets the value for Stdout to be an explicit nil

### UnsetStdout
`func (o *QueryResponseAlternative) UnsetStdout()`

UnsetStdout ensures that no value is present for Stdout, not even an explicit nil
### GetClickhouse

`func (o *QueryResponseAlternative) GetClickhouse() string`

GetClickhouse returns the Clickhouse field if non-nil, zero value otherwise.

### GetClickhouseOk

`func (o *QueryResponseAlternative) GetClickhouseOk() (*string, bool)`

GetClickhouseOk returns a tuple with the Clickhouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickhouse

`func (o *QueryResponseAlternative) SetClickhouse(v string)`

SetClickhouse sets Clickhouse field to given value.

### HasClickhouse

`func (o *QueryResponseAlternative) HasClickhouse() bool`

HasClickhouse returns a boolean if a field has been set.

### SetClickhouseNil

`func (o *QueryResponseAlternative) SetClickhouseNil(b bool)`

 SetClickhouseNil sets the value for Clickhouse to be an explicit nil

### UnsetClickhouse
`func (o *QueryResponseAlternative) UnsetClickhouse()`

UnsetClickhouse ensures that no value is present for Clickhouse, not even an explicit nil
### GetExplain

`func (o *QueryResponseAlternative) GetExplain() []string`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *QueryResponseAlternative) GetExplainOk() (*[]string, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *QueryResponseAlternative) SetExplain(v []string)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *QueryResponseAlternative) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### SetExplainNil

`func (o *QueryResponseAlternative) SetExplainNil(b bool)`

 SetExplainNil sets the value for Explain to be an explicit nil

### UnsetExplain
`func (o *QueryResponseAlternative) UnsetExplain()`

UnsetExplain ensures that no value is present for Explain, not even an explicit nil
### GetMetadata

`func (o *QueryResponseAlternative) GetMetadata() HogQLMetadataResponse`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QueryResponseAlternative) GetMetadataOk() (*HogQLMetadataResponse, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QueryResponseAlternative) SetMetadata(v HogQLMetadataResponse)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QueryResponseAlternative) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQuery

`func (o *QueryResponseAlternative) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryResponseAlternative) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryResponseAlternative) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *QueryResponseAlternative) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *QueryResponseAlternative) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *QueryResponseAlternative) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetChTableNames

`func (o *QueryResponseAlternative) GetChTableNames() []string`

GetChTableNames returns the ChTableNames field if non-nil, zero value otherwise.

### GetChTableNamesOk

`func (o *QueryResponseAlternative) GetChTableNamesOk() (*[]string, bool)`

GetChTableNamesOk returns a tuple with the ChTableNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChTableNames

`func (o *QueryResponseAlternative) SetChTableNames(v []string)`

SetChTableNames sets ChTableNames field to given value.

### HasChTableNames

`func (o *QueryResponseAlternative) HasChTableNames() bool`

HasChTableNames returns a boolean if a field has been set.

### SetChTableNamesNil

`func (o *QueryResponseAlternative) SetChTableNamesNil(b bool)`

 SetChTableNamesNil sets the value for ChTableNames to be an explicit nil

### UnsetChTableNames
`func (o *QueryResponseAlternative) UnsetChTableNames()`

UnsetChTableNames ensures that no value is present for ChTableNames, not even an explicit nil
### GetErrors

`func (o *QueryResponseAlternative) GetErrors() []HogQLNotice`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *QueryResponseAlternative) GetErrorsOk() (*[]HogQLNotice, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *QueryResponseAlternative) SetErrors(v []HogQLNotice)`

SetErrors sets Errors field to given value.


### GetIsUsingIndices

`func (o *QueryResponseAlternative) GetIsUsingIndices() QueryIndexUsage`

GetIsUsingIndices returns the IsUsingIndices field if non-nil, zero value otherwise.

### GetIsUsingIndicesOk

`func (o *QueryResponseAlternative) GetIsUsingIndicesOk() (*QueryIndexUsage, bool)`

GetIsUsingIndicesOk returns a tuple with the IsUsingIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsingIndices

`func (o *QueryResponseAlternative) SetIsUsingIndices(v QueryIndexUsage)`

SetIsUsingIndices sets IsUsingIndices field to given value.

### HasIsUsingIndices

`func (o *QueryResponseAlternative) HasIsUsingIndices() bool`

HasIsUsingIndices returns a boolean if a field has been set.

### GetIsValid

`func (o *QueryResponseAlternative) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *QueryResponseAlternative) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *QueryResponseAlternative) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *QueryResponseAlternative) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### SetIsValidNil

`func (o *QueryResponseAlternative) SetIsValidNil(b bool)`

 SetIsValidNil sets the value for IsValid to be an explicit nil

### UnsetIsValid
`func (o *QueryResponseAlternative) UnsetIsValid()`

UnsetIsValid ensures that no value is present for IsValid, not even an explicit nil
### GetNotices

`func (o *QueryResponseAlternative) GetNotices() []HogQLNotice`

GetNotices returns the Notices field if non-nil, zero value otherwise.

### GetNoticesOk

`func (o *QueryResponseAlternative) GetNoticesOk() (*[]HogQLNotice, bool)`

GetNoticesOk returns a tuple with the Notices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotices

`func (o *QueryResponseAlternative) SetNotices(v []HogQLNotice)`

SetNotices sets Notices field to given value.


### GetTableNames

`func (o *QueryResponseAlternative) GetTableNames() []string`

GetTableNames returns the TableNames field if non-nil, zero value otherwise.

### GetTableNamesOk

`func (o *QueryResponseAlternative) GetTableNamesOk() (*[]string, bool)`

GetTableNamesOk returns a tuple with the TableNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableNames

`func (o *QueryResponseAlternative) SetTableNames(v []string)`

SetTableNames sets TableNames field to given value.

### HasTableNames

`func (o *QueryResponseAlternative) HasTableNames() bool`

HasTableNames returns a boolean if a field has been set.

### SetTableNamesNil

`func (o *QueryResponseAlternative) SetTableNamesNil(b bool)`

 SetTableNamesNil sets the value for TableNames to be an explicit nil

### UnsetTableNames
`func (o *QueryResponseAlternative) UnsetTableNames()`

UnsetTableNames ensures that no value is present for TableNames, not even an explicit nil
### GetWarnings

`func (o *QueryResponseAlternative) GetWarnings() []HogQLNotice`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *QueryResponseAlternative) GetWarningsOk() (*[]HogQLNotice, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *QueryResponseAlternative) SetWarnings(v []HogQLNotice)`

SetWarnings sets Warnings field to given value.


### GetIncompleteList

`func (o *QueryResponseAlternative) GetIncompleteList() bool`

GetIncompleteList returns the IncompleteList field if non-nil, zero value otherwise.

### GetIncompleteListOk

`func (o *QueryResponseAlternative) GetIncompleteListOk() (*bool, bool)`

GetIncompleteListOk returns a tuple with the IncompleteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteList

`func (o *QueryResponseAlternative) SetIncompleteList(v bool)`

SetIncompleteList sets IncompleteList field to given value.


### GetSuggestions

`func (o *QueryResponseAlternative) GetSuggestions() []AutocompleteCompletionItem`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *QueryResponseAlternative) GetSuggestionsOk() (*[]AutocompleteCompletionItem, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *QueryResponseAlternative) SetSuggestions(v []AutocompleteCompletionItem)`

SetSuggestions sets Suggestions field to given value.


### GetCredibleIntervals

`func (o *QueryResponseAlternative) GetCredibleIntervals() map[string][]float32`

GetCredibleIntervals returns the CredibleIntervals field if non-nil, zero value otherwise.

### GetCredibleIntervalsOk

`func (o *QueryResponseAlternative) GetCredibleIntervalsOk() (*map[string][]float32, bool)`

GetCredibleIntervalsOk returns a tuple with the CredibleIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredibleIntervals

`func (o *QueryResponseAlternative) SetCredibleIntervals(v map[string][]float32)`

SetCredibleIntervals sets CredibleIntervals field to given value.


### GetExpectedLoss

`func (o *QueryResponseAlternative) GetExpectedLoss() float32`

GetExpectedLoss returns the ExpectedLoss field if non-nil, zero value otherwise.

### GetExpectedLossOk

`func (o *QueryResponseAlternative) GetExpectedLossOk() (*float32, bool)`

GetExpectedLossOk returns a tuple with the ExpectedLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedLoss

`func (o *QueryResponseAlternative) SetExpectedLoss(v float32)`

SetExpectedLoss sets ExpectedLoss field to given value.


### GetFunnelsQuery

`func (o *QueryResponseAlternative) GetFunnelsQuery() FunnelsQuery`

GetFunnelsQuery returns the FunnelsQuery field if non-nil, zero value otherwise.

### GetFunnelsQueryOk

`func (o *QueryResponseAlternative) GetFunnelsQueryOk() (*FunnelsQuery, bool)`

GetFunnelsQueryOk returns a tuple with the FunnelsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelsQuery

`func (o *QueryResponseAlternative) SetFunnelsQuery(v FunnelsQuery)`

SetFunnelsQuery sets FunnelsQuery field to given value.

### HasFunnelsQuery

`func (o *QueryResponseAlternative) HasFunnelsQuery() bool`

HasFunnelsQuery returns a boolean if a field has been set.

### GetInsight

`func (o *QueryResponseAlternative) GetInsight() []map[string]interface{}`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *QueryResponseAlternative) GetInsightOk() (*[]map[string]interface{}, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *QueryResponseAlternative) SetInsight(v []map[string]interface{})`

SetInsight sets Insight field to given value.


### GetProbability

`func (o *QueryResponseAlternative) GetProbability() map[string]float32`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *QueryResponseAlternative) GetProbabilityOk() (*map[string]float32, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *QueryResponseAlternative) SetProbability(v map[string]float32)`

SetProbability sets Probability field to given value.


### GetSignificanceCode

`func (o *QueryResponseAlternative) GetSignificanceCode() ExperimentSignificanceCode`

GetSignificanceCode returns the SignificanceCode field if non-nil, zero value otherwise.

### GetSignificanceCodeOk

`func (o *QueryResponseAlternative) GetSignificanceCodeOk() (*ExperimentSignificanceCode, bool)`

GetSignificanceCodeOk returns a tuple with the SignificanceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificanceCode

`func (o *QueryResponseAlternative) SetSignificanceCode(v ExperimentSignificanceCode)`

SetSignificanceCode sets SignificanceCode field to given value.


### GetSignificant

`func (o *QueryResponseAlternative) GetSignificant() bool`

GetSignificant returns the Significant field if non-nil, zero value otherwise.

### GetSignificantOk

`func (o *QueryResponseAlternative) GetSignificantOk() (*bool, bool)`

GetSignificantOk returns a tuple with the Significant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificant

`func (o *QueryResponseAlternative) SetSignificant(v bool)`

SetSignificant sets Significant field to given value.


### GetStatsVersion

`func (o *QueryResponseAlternative) GetStatsVersion() int32`

GetStatsVersion returns the StatsVersion field if non-nil, zero value otherwise.

### GetStatsVersionOk

`func (o *QueryResponseAlternative) GetStatsVersionOk() (*int32, bool)`

GetStatsVersionOk returns a tuple with the StatsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsVersion

`func (o *QueryResponseAlternative) SetStatsVersion(v int32)`

SetStatsVersion sets StatsVersion field to given value.

### HasStatsVersion

`func (o *QueryResponseAlternative) HasStatsVersion() bool`

HasStatsVersion returns a boolean if a field has been set.

### SetStatsVersionNil

`func (o *QueryResponseAlternative) SetStatsVersionNil(b bool)`

 SetStatsVersionNil sets the value for StatsVersion to be an explicit nil

### UnsetStatsVersion
`func (o *QueryResponseAlternative) UnsetStatsVersion()`

UnsetStatsVersion ensures that no value is present for StatsVersion, not even an explicit nil
### GetVariants

`func (o *QueryResponseAlternative) GetVariants() []ExperimentVariantTrendsBaseStats`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *QueryResponseAlternative) GetVariantsOk() (*[]ExperimentVariantTrendsBaseStats, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *QueryResponseAlternative) SetVariants(v []ExperimentVariantTrendsBaseStats)`

SetVariants sets Variants field to given value.


### GetCountQuery

`func (o *QueryResponseAlternative) GetCountQuery() TrendsQuery`

GetCountQuery returns the CountQuery field if non-nil, zero value otherwise.

### GetCountQueryOk

`func (o *QueryResponseAlternative) GetCountQueryOk() (*TrendsQuery, bool)`

GetCountQueryOk returns a tuple with the CountQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountQuery

`func (o *QueryResponseAlternative) SetCountQuery(v TrendsQuery)`

SetCountQuery sets CountQuery field to given value.

### HasCountQuery

`func (o *QueryResponseAlternative) HasCountQuery() bool`

HasCountQuery returns a boolean if a field has been set.

### GetExposureQuery

`func (o *QueryResponseAlternative) GetExposureQuery() TrendsQuery`

GetExposureQuery returns the ExposureQuery field if non-nil, zero value otherwise.

### GetExposureQueryOk

`func (o *QueryResponseAlternative) GetExposureQueryOk() (*TrendsQuery, bool)`

GetExposureQueryOk returns a tuple with the ExposureQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureQuery

`func (o *QueryResponseAlternative) SetExposureQuery(v TrendsQuery)`

SetExposureQuery sets ExposureQuery field to given value.

### HasExposureQuery

`func (o *QueryResponseAlternative) HasExposureQuery() bool`

HasExposureQuery returns a boolean if a field has been set.

### GetPValue

`func (o *QueryResponseAlternative) GetPValue() float32`

GetPValue returns the PValue field if non-nil, zero value otherwise.

### GetPValueOk

`func (o *QueryResponseAlternative) GetPValueOk() (*float32, bool)`

GetPValueOk returns a tuple with the PValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPValue

`func (o *QueryResponseAlternative) SetPValue(v float32)`

SetPValue sets PValue field to given value.


### GetBaseline

`func (o *QueryResponseAlternative) GetBaseline() ExperimentStatsBaseValidated`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *QueryResponseAlternative) GetBaselineOk() (*ExperimentStatsBaseValidated, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *QueryResponseAlternative) SetBaseline(v ExperimentStatsBaseValidated)`

SetBaseline sets Baseline field to given value.

### HasBaseline

`func (o *QueryResponseAlternative) HasBaseline() bool`

HasBaseline returns a boolean if a field has been set.

### GetBreakdownResults

`func (o *QueryResponseAlternative) GetBreakdownResults() []ExperimentBreakdownResult`

GetBreakdownResults returns the BreakdownResults field if non-nil, zero value otherwise.

### GetBreakdownResultsOk

`func (o *QueryResponseAlternative) GetBreakdownResultsOk() (*[]ExperimentBreakdownResult, bool)`

GetBreakdownResultsOk returns a tuple with the BreakdownResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownResults

`func (o *QueryResponseAlternative) SetBreakdownResults(v []ExperimentBreakdownResult)`

SetBreakdownResults sets BreakdownResults field to given value.

### HasBreakdownResults

`func (o *QueryResponseAlternative) HasBreakdownResults() bool`

HasBreakdownResults returns a boolean if a field has been set.

### SetBreakdownResultsNil

`func (o *QueryResponseAlternative) SetBreakdownResultsNil(b bool)`

 SetBreakdownResultsNil sets the value for BreakdownResults to be an explicit nil

### UnsetBreakdownResults
`func (o *QueryResponseAlternative) UnsetBreakdownResults()`

UnsetBreakdownResults ensures that no value is present for BreakdownResults, not even an explicit nil
### GetMetric

`func (o *QueryResponseAlternative) GetMetric() Metric1`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *QueryResponseAlternative) GetMetricOk() (*Metric1, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *QueryResponseAlternative) SetMetric(v Metric1)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *QueryResponseAlternative) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### SetMetricNil

`func (o *QueryResponseAlternative) SetMetricNil(b bool)`

 SetMetricNil sets the value for Metric to be an explicit nil

### UnsetMetric
`func (o *QueryResponseAlternative) UnsetMetric()`

UnsetMetric ensures that no value is present for Metric, not even an explicit nil
### GetVariantResults

`func (o *QueryResponseAlternative) GetVariantResults() VariantResults`

GetVariantResults returns the VariantResults field if non-nil, zero value otherwise.

### GetVariantResultsOk

`func (o *QueryResponseAlternative) GetVariantResultsOk() (*VariantResults, bool)`

GetVariantResultsOk returns a tuple with the VariantResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantResults

`func (o *QueryResponseAlternative) SetVariantResults(v VariantResults)`

SetVariantResults sets VariantResults field to given value.

### HasVariantResults

`func (o *QueryResponseAlternative) HasVariantResults() bool`

HasVariantResults returns a boolean if a field has been set.

### SetVariantResultsNil

`func (o *QueryResponseAlternative) SetVariantResultsNil(b bool)`

 SetVariantResultsNil sets the value for VariantResults to be an explicit nil

### UnsetVariantResults
`func (o *QueryResponseAlternative) UnsetVariantResults()`

UnsetVariantResults ensures that no value is present for VariantResults, not even an explicit nil
### GetDateRange

`func (o *QueryResponseAlternative) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *QueryResponseAlternative) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *QueryResponseAlternative) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.


### GetTimeseries

`func (o *QueryResponseAlternative) GetTimeseries() []ExperimentExposureTimeSeries`

GetTimeseries returns the Timeseries field if non-nil, zero value otherwise.

### GetTimeseriesOk

`func (o *QueryResponseAlternative) GetTimeseriesOk() (*[]ExperimentExposureTimeSeries, bool)`

GetTimeseriesOk returns a tuple with the Timeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseries

`func (o *QueryResponseAlternative) SetTimeseries(v []ExperimentExposureTimeSeries)`

SetTimeseries sets Timeseries field to given value.


### GetTotalExposures

`func (o *QueryResponseAlternative) GetTotalExposures() map[string]float32`

GetTotalExposures returns the TotalExposures field if non-nil, zero value otherwise.

### GetTotalExposuresOk

`func (o *QueryResponseAlternative) GetTotalExposuresOk() (*map[string]float32, bool)`

GetTotalExposuresOk returns a tuple with the TotalExposures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExposures

`func (o *QueryResponseAlternative) SetTotalExposures(v map[string]float32)`

SetTotalExposures sets TotalExposures field to given value.


### GetDateFrom

`func (o *QueryResponseAlternative) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *QueryResponseAlternative) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *QueryResponseAlternative) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *QueryResponseAlternative) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### SetDateFromNil

`func (o *QueryResponseAlternative) SetDateFromNil(b bool)`

 SetDateFromNil sets the value for DateFrom to be an explicit nil

### UnsetDateFrom
`func (o *QueryResponseAlternative) UnsetDateFrom()`

UnsetDateFrom ensures that no value is present for DateFrom, not even an explicit nil
### GetDateTo

`func (o *QueryResponseAlternative) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *QueryResponseAlternative) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *QueryResponseAlternative) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *QueryResponseAlternative) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### SetDateToNil

`func (o *QueryResponseAlternative) SetDateToNil(b bool)`

 SetDateToNil sets the value for DateTo to be an explicit nil

### UnsetDateTo
`func (o *QueryResponseAlternative) UnsetDateTo()`

UnsetDateTo ensures that no value is present for DateTo, not even an explicit nil
### GetSamplingRate

`func (o *QueryResponseAlternative) GetSamplingRate() SamplingRate`

GetSamplingRate returns the SamplingRate field if non-nil, zero value otherwise.

### GetSamplingRateOk

`func (o *QueryResponseAlternative) GetSamplingRateOk() (*SamplingRate, bool)`

GetSamplingRateOk returns a tuple with the SamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingRate

`func (o *QueryResponseAlternative) SetSamplingRate(v SamplingRate)`

SetSamplingRate sets SamplingRate field to given value.

### HasSamplingRate

`func (o *QueryResponseAlternative) HasSamplingRate() bool`

HasSamplingRate returns a boolean if a field has been set.

### GetUsedPreAggregatedTables

`func (o *QueryResponseAlternative) GetUsedPreAggregatedTables() bool`

GetUsedPreAggregatedTables returns the UsedPreAggregatedTables field if non-nil, zero value otherwise.

### GetUsedPreAggregatedTablesOk

`func (o *QueryResponseAlternative) GetUsedPreAggregatedTablesOk() (*bool, bool)`

GetUsedPreAggregatedTablesOk returns a tuple with the UsedPreAggregatedTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPreAggregatedTables

`func (o *QueryResponseAlternative) SetUsedPreAggregatedTables(v bool)`

SetUsedPreAggregatedTables sets UsedPreAggregatedTables field to given value.

### HasUsedPreAggregatedTables

`func (o *QueryResponseAlternative) HasUsedPreAggregatedTables() bool`

HasUsedPreAggregatedTables returns a boolean if a field has been set.

### SetUsedPreAggregatedTablesNil

`func (o *QueryResponseAlternative) SetUsedPreAggregatedTablesNil(b bool)`

 SetUsedPreAggregatedTablesNil sets the value for UsedPreAggregatedTables to be an explicit nil

### UnsetUsedPreAggregatedTables
`func (o *QueryResponseAlternative) UnsetUsedPreAggregatedTables()`

UnsetUsedPreAggregatedTables ensures that no value is present for UsedPreAggregatedTables, not even an explicit nil
### GetData

`func (o *QueryResponseAlternative) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QueryResponseAlternative) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QueryResponseAlternative) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetIsUdf

`func (o *QueryResponseAlternative) GetIsUdf() bool`

GetIsUdf returns the IsUdf field if non-nil, zero value otherwise.

### GetIsUdfOk

`func (o *QueryResponseAlternative) GetIsUdfOk() (*bool, bool)`

GetIsUdfOk returns a tuple with the IsUdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUdf

`func (o *QueryResponseAlternative) SetIsUdf(v bool)`

SetIsUdf sets IsUdf field to given value.

### HasIsUdf

`func (o *QueryResponseAlternative) HasIsUdf() bool`

HasIsUdf returns a boolean if a field has been set.

### SetIsUdfNil

`func (o *QueryResponseAlternative) SetIsUdfNil(b bool)`

 SetIsUdfNil sets the value for IsUdf to be an explicit nil

### UnsetIsUdf
`func (o *QueryResponseAlternative) UnsetIsUdf()`

UnsetIsUdf ensures that no value is present for IsUdf, not even an explicit nil
### GetJoins

`func (o *QueryResponseAlternative) GetJoins() []DataWarehouseViewLink`

GetJoins returns the Joins field if non-nil, zero value otherwise.

### GetJoinsOk

`func (o *QueryResponseAlternative) GetJoinsOk() (*[]DataWarehouseViewLink, bool)`

GetJoinsOk returns a tuple with the Joins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoins

`func (o *QueryResponseAlternative) SetJoins(v []DataWarehouseViewLink)`

SetJoins sets Joins field to given value.


### GetTables

`func (o *QueryResponseAlternative) GetTables() map[string]TablesValue`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *QueryResponseAlternative) GetTablesOk() (*map[string]TablesValue, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *QueryResponseAlternative) SetTables(v map[string]TablesValue)`

SetTables sets Tables field to given value.


### GetQuestions

`func (o *QueryResponseAlternative) GetQuestions() []string`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *QueryResponseAlternative) GetQuestionsOk() (*[]string, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *QueryResponseAlternative) SetQuestions(v []string)`

SetQuestions sets Questions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


