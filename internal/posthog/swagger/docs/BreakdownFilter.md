# BreakdownFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Breakdown** | Pointer to [**NullableBreakdown1**](Breakdown1.md) |  | [optional] [default to null]
**BreakdownGroupTypeIndex** | Pointer to **NullableInt32** |  | [optional] 
**BreakdownHideOtherAggregation** | Pointer to **NullableBool** |  | [optional] 
**BreakdownHistogramBinCount** | Pointer to **NullableInt32** |  | [optional] 
**BreakdownLimit** | Pointer to **NullableInt32** |  | [optional] 
**BreakdownNormalizeUrl** | Pointer to **NullableBool** |  | [optional] 
**BreakdownType** | Pointer to [**BreakdownType**](BreakdownType.md) |  | [optional] 
**Breakdowns** | Pointer to [**[]Breakdown**](Breakdown.md) |  | [optional] 

## Methods

### NewBreakdownFilter

`func NewBreakdownFilter() *BreakdownFilter`

NewBreakdownFilter instantiates a new BreakdownFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBreakdownFilterWithDefaults

`func NewBreakdownFilterWithDefaults() *BreakdownFilter`

NewBreakdownFilterWithDefaults instantiates a new BreakdownFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdown

`func (o *BreakdownFilter) GetBreakdown() Breakdown1`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *BreakdownFilter) GetBreakdownOk() (*Breakdown1, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *BreakdownFilter) SetBreakdown(v Breakdown1)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *BreakdownFilter) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### SetBreakdownNil

`func (o *BreakdownFilter) SetBreakdownNil(b bool)`

 SetBreakdownNil sets the value for Breakdown to be an explicit nil

### UnsetBreakdown
`func (o *BreakdownFilter) UnsetBreakdown()`

UnsetBreakdown ensures that no value is present for Breakdown, not even an explicit nil
### GetBreakdownGroupTypeIndex

`func (o *BreakdownFilter) GetBreakdownGroupTypeIndex() int32`

GetBreakdownGroupTypeIndex returns the BreakdownGroupTypeIndex field if non-nil, zero value otherwise.

### GetBreakdownGroupTypeIndexOk

`func (o *BreakdownFilter) GetBreakdownGroupTypeIndexOk() (*int32, bool)`

GetBreakdownGroupTypeIndexOk returns a tuple with the BreakdownGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownGroupTypeIndex

`func (o *BreakdownFilter) SetBreakdownGroupTypeIndex(v int32)`

SetBreakdownGroupTypeIndex sets BreakdownGroupTypeIndex field to given value.

### HasBreakdownGroupTypeIndex

`func (o *BreakdownFilter) HasBreakdownGroupTypeIndex() bool`

HasBreakdownGroupTypeIndex returns a boolean if a field has been set.

### SetBreakdownGroupTypeIndexNil

`func (o *BreakdownFilter) SetBreakdownGroupTypeIndexNil(b bool)`

 SetBreakdownGroupTypeIndexNil sets the value for BreakdownGroupTypeIndex to be an explicit nil

### UnsetBreakdownGroupTypeIndex
`func (o *BreakdownFilter) UnsetBreakdownGroupTypeIndex()`

UnsetBreakdownGroupTypeIndex ensures that no value is present for BreakdownGroupTypeIndex, not even an explicit nil
### GetBreakdownHideOtherAggregation

`func (o *BreakdownFilter) GetBreakdownHideOtherAggregation() bool`

GetBreakdownHideOtherAggregation returns the BreakdownHideOtherAggregation field if non-nil, zero value otherwise.

### GetBreakdownHideOtherAggregationOk

`func (o *BreakdownFilter) GetBreakdownHideOtherAggregationOk() (*bool, bool)`

GetBreakdownHideOtherAggregationOk returns a tuple with the BreakdownHideOtherAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownHideOtherAggregation

`func (o *BreakdownFilter) SetBreakdownHideOtherAggregation(v bool)`

SetBreakdownHideOtherAggregation sets BreakdownHideOtherAggregation field to given value.

### HasBreakdownHideOtherAggregation

`func (o *BreakdownFilter) HasBreakdownHideOtherAggregation() bool`

HasBreakdownHideOtherAggregation returns a boolean if a field has been set.

### SetBreakdownHideOtherAggregationNil

`func (o *BreakdownFilter) SetBreakdownHideOtherAggregationNil(b bool)`

 SetBreakdownHideOtherAggregationNil sets the value for BreakdownHideOtherAggregation to be an explicit nil

### UnsetBreakdownHideOtherAggregation
`func (o *BreakdownFilter) UnsetBreakdownHideOtherAggregation()`

UnsetBreakdownHideOtherAggregation ensures that no value is present for BreakdownHideOtherAggregation, not even an explicit nil
### GetBreakdownHistogramBinCount

`func (o *BreakdownFilter) GetBreakdownHistogramBinCount() int32`

GetBreakdownHistogramBinCount returns the BreakdownHistogramBinCount field if non-nil, zero value otherwise.

### GetBreakdownHistogramBinCountOk

`func (o *BreakdownFilter) GetBreakdownHistogramBinCountOk() (*int32, bool)`

GetBreakdownHistogramBinCountOk returns a tuple with the BreakdownHistogramBinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownHistogramBinCount

`func (o *BreakdownFilter) SetBreakdownHistogramBinCount(v int32)`

SetBreakdownHistogramBinCount sets BreakdownHistogramBinCount field to given value.

### HasBreakdownHistogramBinCount

`func (o *BreakdownFilter) HasBreakdownHistogramBinCount() bool`

HasBreakdownHistogramBinCount returns a boolean if a field has been set.

### SetBreakdownHistogramBinCountNil

`func (o *BreakdownFilter) SetBreakdownHistogramBinCountNil(b bool)`

 SetBreakdownHistogramBinCountNil sets the value for BreakdownHistogramBinCount to be an explicit nil

### UnsetBreakdownHistogramBinCount
`func (o *BreakdownFilter) UnsetBreakdownHistogramBinCount()`

UnsetBreakdownHistogramBinCount ensures that no value is present for BreakdownHistogramBinCount, not even an explicit nil
### GetBreakdownLimit

`func (o *BreakdownFilter) GetBreakdownLimit() int32`

GetBreakdownLimit returns the BreakdownLimit field if non-nil, zero value otherwise.

### GetBreakdownLimitOk

`func (o *BreakdownFilter) GetBreakdownLimitOk() (*int32, bool)`

GetBreakdownLimitOk returns a tuple with the BreakdownLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownLimit

`func (o *BreakdownFilter) SetBreakdownLimit(v int32)`

SetBreakdownLimit sets BreakdownLimit field to given value.

### HasBreakdownLimit

`func (o *BreakdownFilter) HasBreakdownLimit() bool`

HasBreakdownLimit returns a boolean if a field has been set.

### SetBreakdownLimitNil

`func (o *BreakdownFilter) SetBreakdownLimitNil(b bool)`

 SetBreakdownLimitNil sets the value for BreakdownLimit to be an explicit nil

### UnsetBreakdownLimit
`func (o *BreakdownFilter) UnsetBreakdownLimit()`

UnsetBreakdownLimit ensures that no value is present for BreakdownLimit, not even an explicit nil
### GetBreakdownNormalizeUrl

`func (o *BreakdownFilter) GetBreakdownNormalizeUrl() bool`

GetBreakdownNormalizeUrl returns the BreakdownNormalizeUrl field if non-nil, zero value otherwise.

### GetBreakdownNormalizeUrlOk

`func (o *BreakdownFilter) GetBreakdownNormalizeUrlOk() (*bool, bool)`

GetBreakdownNormalizeUrlOk returns a tuple with the BreakdownNormalizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownNormalizeUrl

`func (o *BreakdownFilter) SetBreakdownNormalizeUrl(v bool)`

SetBreakdownNormalizeUrl sets BreakdownNormalizeUrl field to given value.

### HasBreakdownNormalizeUrl

`func (o *BreakdownFilter) HasBreakdownNormalizeUrl() bool`

HasBreakdownNormalizeUrl returns a boolean if a field has been set.

### SetBreakdownNormalizeUrlNil

`func (o *BreakdownFilter) SetBreakdownNormalizeUrlNil(b bool)`

 SetBreakdownNormalizeUrlNil sets the value for BreakdownNormalizeUrl to be an explicit nil

### UnsetBreakdownNormalizeUrl
`func (o *BreakdownFilter) UnsetBreakdownNormalizeUrl()`

UnsetBreakdownNormalizeUrl ensures that no value is present for BreakdownNormalizeUrl, not even an explicit nil
### GetBreakdownType

`func (o *BreakdownFilter) GetBreakdownType() BreakdownType`

GetBreakdownType returns the BreakdownType field if non-nil, zero value otherwise.

### GetBreakdownTypeOk

`func (o *BreakdownFilter) GetBreakdownTypeOk() (*BreakdownType, bool)`

GetBreakdownTypeOk returns a tuple with the BreakdownType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownType

`func (o *BreakdownFilter) SetBreakdownType(v BreakdownType)`

SetBreakdownType sets BreakdownType field to given value.

### HasBreakdownType

`func (o *BreakdownFilter) HasBreakdownType() bool`

HasBreakdownType returns a boolean if a field has been set.

### GetBreakdowns

`func (o *BreakdownFilter) GetBreakdowns() []Breakdown`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *BreakdownFilter) GetBreakdownsOk() (*[]Breakdown, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *BreakdownFilter) SetBreakdowns(v []Breakdown)`

SetBreakdowns sets Breakdowns field to given value.

### HasBreakdowns

`func (o *BreakdownFilter) HasBreakdowns() bool`

HasBreakdowns returns a boolean if a field has been set.

### SetBreakdownsNil

`func (o *BreakdownFilter) SetBreakdownsNil(b bool)`

 SetBreakdownsNil sets the value for Breakdowns to be an explicit nil

### UnsetBreakdowns
`func (o *BreakdownFilter) UnsetBreakdowns()`

UnsetBreakdowns ensures that no value is present for Breakdowns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


