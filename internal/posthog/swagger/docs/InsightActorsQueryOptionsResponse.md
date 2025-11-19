# InsightActorsQueryOptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Breakdown** | Pointer to [**[]BreakdownItem**](BreakdownItem.md) |  | [optional] 
**Breakdowns** | Pointer to [**[]MultipleBreakdownOptions**](MultipleBreakdownOptions.md) |  | [optional] 
**Compare** | Pointer to [**[]CompareItem**](CompareItem.md) |  | [optional] 
**Day** | Pointer to [**[]DayItem**](DayItem.md) |  | [optional] 
**Interval** | Pointer to [**[]IntervalItem**](IntervalItem.md) |  | [optional] 
**Series** | Pointer to [**[]Series**](Series.md) |  | [optional] 
**Status** | Pointer to [**[]StatusItem**](StatusItem.md) |  | [optional] 

## Methods

### NewInsightActorsQueryOptionsResponse

`func NewInsightActorsQueryOptionsResponse() *InsightActorsQueryOptionsResponse`

NewInsightActorsQueryOptionsResponse instantiates a new InsightActorsQueryOptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightActorsQueryOptionsResponseWithDefaults

`func NewInsightActorsQueryOptionsResponseWithDefaults() *InsightActorsQueryOptionsResponse`

NewInsightActorsQueryOptionsResponseWithDefaults instantiates a new InsightActorsQueryOptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdown

`func (o *InsightActorsQueryOptionsResponse) GetBreakdown() []BreakdownItem`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *InsightActorsQueryOptionsResponse) GetBreakdownOk() (*[]BreakdownItem, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *InsightActorsQueryOptionsResponse) SetBreakdown(v []BreakdownItem)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *InsightActorsQueryOptionsResponse) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### SetBreakdownNil

`func (o *InsightActorsQueryOptionsResponse) SetBreakdownNil(b bool)`

 SetBreakdownNil sets the value for Breakdown to be an explicit nil

### UnsetBreakdown
`func (o *InsightActorsQueryOptionsResponse) UnsetBreakdown()`

UnsetBreakdown ensures that no value is present for Breakdown, not even an explicit nil
### GetBreakdowns

`func (o *InsightActorsQueryOptionsResponse) GetBreakdowns() []MultipleBreakdownOptions`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *InsightActorsQueryOptionsResponse) GetBreakdownsOk() (*[]MultipleBreakdownOptions, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *InsightActorsQueryOptionsResponse) SetBreakdowns(v []MultipleBreakdownOptions)`

SetBreakdowns sets Breakdowns field to given value.

### HasBreakdowns

`func (o *InsightActorsQueryOptionsResponse) HasBreakdowns() bool`

HasBreakdowns returns a boolean if a field has been set.

### SetBreakdownsNil

`func (o *InsightActorsQueryOptionsResponse) SetBreakdownsNil(b bool)`

 SetBreakdownsNil sets the value for Breakdowns to be an explicit nil

### UnsetBreakdowns
`func (o *InsightActorsQueryOptionsResponse) UnsetBreakdowns()`

UnsetBreakdowns ensures that no value is present for Breakdowns, not even an explicit nil
### GetCompare

`func (o *InsightActorsQueryOptionsResponse) GetCompare() []CompareItem`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *InsightActorsQueryOptionsResponse) GetCompareOk() (*[]CompareItem, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *InsightActorsQueryOptionsResponse) SetCompare(v []CompareItem)`

SetCompare sets Compare field to given value.

### HasCompare

`func (o *InsightActorsQueryOptionsResponse) HasCompare() bool`

HasCompare returns a boolean if a field has been set.

### SetCompareNil

`func (o *InsightActorsQueryOptionsResponse) SetCompareNil(b bool)`

 SetCompareNil sets the value for Compare to be an explicit nil

### UnsetCompare
`func (o *InsightActorsQueryOptionsResponse) UnsetCompare()`

UnsetCompare ensures that no value is present for Compare, not even an explicit nil
### GetDay

`func (o *InsightActorsQueryOptionsResponse) GetDay() []DayItem`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *InsightActorsQueryOptionsResponse) GetDayOk() (*[]DayItem, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *InsightActorsQueryOptionsResponse) SetDay(v []DayItem)`

SetDay sets Day field to given value.

### HasDay

`func (o *InsightActorsQueryOptionsResponse) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *InsightActorsQueryOptionsResponse) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *InsightActorsQueryOptionsResponse) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetInterval

`func (o *InsightActorsQueryOptionsResponse) GetInterval() []IntervalItem`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InsightActorsQueryOptionsResponse) GetIntervalOk() (*[]IntervalItem, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InsightActorsQueryOptionsResponse) SetInterval(v []IntervalItem)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *InsightActorsQueryOptionsResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *InsightActorsQueryOptionsResponse) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *InsightActorsQueryOptionsResponse) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetSeries

`func (o *InsightActorsQueryOptionsResponse) GetSeries() []Series`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *InsightActorsQueryOptionsResponse) GetSeriesOk() (*[]Series, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *InsightActorsQueryOptionsResponse) SetSeries(v []Series)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *InsightActorsQueryOptionsResponse) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeriesNil

`func (o *InsightActorsQueryOptionsResponse) SetSeriesNil(b bool)`

 SetSeriesNil sets the value for Series to be an explicit nil

### UnsetSeries
`func (o *InsightActorsQueryOptionsResponse) UnsetSeries()`

UnsetSeries ensures that no value is present for Series, not even an explicit nil
### GetStatus

`func (o *InsightActorsQueryOptionsResponse) GetStatus() []StatusItem`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InsightActorsQueryOptionsResponse) GetStatusOk() (*[]StatusItem, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InsightActorsQueryOptionsResponse) SetStatus(v []StatusItem)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InsightActorsQueryOptionsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *InsightActorsQueryOptionsResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *InsightActorsQueryOptionsResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


