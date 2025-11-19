# Source4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Breakdown** | Pointer to [**NullableBreakdown2**](Breakdown2.md) |  | [optional] [default to null]
**Compare** | Pointer to [**Compare**](Compare.md) |  | [optional] 
**Day** | Pointer to [**NullableDay**](Day.md) |  | [optional] [default to null]
**IncludeRecordings** | Pointer to **NullableBool** |  | [optional] 
**Interval** | Pointer to **NullableInt32** | An interval selected out of available intervals in source query. | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "StickinessActorsQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Response** | Pointer to [**ActorsQueryResponse**](ActorsQueryResponse.md) |  | [optional] 
**Series** | Pointer to **NullableInt32** |  | [optional] 
**Source** | [**StickinessQuery**](StickinessQuery.md) |  | 
**Status** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 
**FunnelStep** | Pointer to **NullableInt32** | Index of the step for which we want to get the timestamp for, per person. Positive for converted persons, negative for dropped of persons. | [optional] 
**FunnelStepBreakdown** | Pointer to [**NullableFunnelstepbreakdown**](Funnelstepbreakdown.md) |  | [optional] [default to null]
**FunnelTrendsDropOff** | Pointer to **NullableBool** |  | [optional] 
**FunnelTrendsEntrancePeriodStart** | Pointer to **NullableString** | Used together with &#x60;funnelTrendsDropOff&#x60; for funnels time conversion date for the persons modal. | [optional] 
**FunnelCorrelationPersonConverted** | Pointer to **NullableBool** |  | [optional] 
**FunnelCorrelationPersonEntity** | Pointer to [**NullableFunnelcorrelationpersonentity**](Funnelcorrelationpersonentity.md) |  | [optional] [default to null]
**FunnelCorrelationPropertyValues** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) |  | [optional] 
**Operator** | Pointer to [**StickinessOperator**](StickinessOperator.md) |  | [optional] 

## Methods

### NewSource4

`func NewSource4(source StickinessQuery, ) *Source4`

NewSource4 instantiates a new Source4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSource4WithDefaults

`func NewSource4WithDefaults() *Source4`

NewSource4WithDefaults instantiates a new Source4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdown

`func (o *Source4) GetBreakdown() Breakdown2`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *Source4) GetBreakdownOk() (*Breakdown2, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *Source4) SetBreakdown(v Breakdown2)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *Source4) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### SetBreakdownNil

`func (o *Source4) SetBreakdownNil(b bool)`

 SetBreakdownNil sets the value for Breakdown to be an explicit nil

### UnsetBreakdown
`func (o *Source4) UnsetBreakdown()`

UnsetBreakdown ensures that no value is present for Breakdown, not even an explicit nil
### GetCompare

`func (o *Source4) GetCompare() Compare`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *Source4) GetCompareOk() (*Compare, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *Source4) SetCompare(v Compare)`

SetCompare sets Compare field to given value.

### HasCompare

`func (o *Source4) HasCompare() bool`

HasCompare returns a boolean if a field has been set.

### GetDay

`func (o *Source4) GetDay() Day`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *Source4) GetDayOk() (*Day, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *Source4) SetDay(v Day)`

SetDay sets Day field to given value.

### HasDay

`func (o *Source4) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *Source4) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *Source4) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetIncludeRecordings

`func (o *Source4) GetIncludeRecordings() bool`

GetIncludeRecordings returns the IncludeRecordings field if non-nil, zero value otherwise.

### GetIncludeRecordingsOk

`func (o *Source4) GetIncludeRecordingsOk() (*bool, bool)`

GetIncludeRecordingsOk returns a tuple with the IncludeRecordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecordings

`func (o *Source4) SetIncludeRecordings(v bool)`

SetIncludeRecordings sets IncludeRecordings field to given value.

### HasIncludeRecordings

`func (o *Source4) HasIncludeRecordings() bool`

HasIncludeRecordings returns a boolean if a field has been set.

### SetIncludeRecordingsNil

`func (o *Source4) SetIncludeRecordingsNil(b bool)`

 SetIncludeRecordingsNil sets the value for IncludeRecordings to be an explicit nil

### UnsetIncludeRecordings
`func (o *Source4) UnsetIncludeRecordings()`

UnsetIncludeRecordings ensures that no value is present for IncludeRecordings, not even an explicit nil
### GetInterval

`func (o *Source4) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Source4) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Source4) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Source4) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *Source4) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *Source4) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetKind

`func (o *Source4) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Source4) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Source4) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Source4) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *Source4) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *Source4) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *Source4) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *Source4) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetResponse

`func (o *Source4) GetResponse() ActorsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Source4) GetResponseOk() (*ActorsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Source4) SetResponse(v ActorsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Source4) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSeries

`func (o *Source4) GetSeries() int32`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Source4) GetSeriesOk() (*int32, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Source4) SetSeries(v int32)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *Source4) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeriesNil

`func (o *Source4) SetSeriesNil(b bool)`

 SetSeriesNil sets the value for Series to be an explicit nil

### UnsetSeries
`func (o *Source4) UnsetSeries()`

UnsetSeries ensures that no value is present for Series, not even an explicit nil
### GetSource

`func (o *Source4) GetSource() StickinessQuery`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Source4) GetSourceOk() (*StickinessQuery, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Source4) SetSource(v StickinessQuery)`

SetSource sets Source field to given value.


### GetStatus

`func (o *Source4) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Source4) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Source4) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Source4) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Source4) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Source4) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTags

`func (o *Source4) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Source4) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Source4) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Source4) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *Source4) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Source4) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Source4) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Source4) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Source4) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Source4) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetFunnelStep

`func (o *Source4) GetFunnelStep() int32`

GetFunnelStep returns the FunnelStep field if non-nil, zero value otherwise.

### GetFunnelStepOk

`func (o *Source4) GetFunnelStepOk() (*int32, bool)`

GetFunnelStepOk returns a tuple with the FunnelStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelStep

`func (o *Source4) SetFunnelStep(v int32)`

SetFunnelStep sets FunnelStep field to given value.

### HasFunnelStep

`func (o *Source4) HasFunnelStep() bool`

HasFunnelStep returns a boolean if a field has been set.

### SetFunnelStepNil

`func (o *Source4) SetFunnelStepNil(b bool)`

 SetFunnelStepNil sets the value for FunnelStep to be an explicit nil

### UnsetFunnelStep
`func (o *Source4) UnsetFunnelStep()`

UnsetFunnelStep ensures that no value is present for FunnelStep, not even an explicit nil
### GetFunnelStepBreakdown

`func (o *Source4) GetFunnelStepBreakdown() Funnelstepbreakdown`

GetFunnelStepBreakdown returns the FunnelStepBreakdown field if non-nil, zero value otherwise.

### GetFunnelStepBreakdownOk

`func (o *Source4) GetFunnelStepBreakdownOk() (*Funnelstepbreakdown, bool)`

GetFunnelStepBreakdownOk returns a tuple with the FunnelStepBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelStepBreakdown

`func (o *Source4) SetFunnelStepBreakdown(v Funnelstepbreakdown)`

SetFunnelStepBreakdown sets FunnelStepBreakdown field to given value.

### HasFunnelStepBreakdown

`func (o *Source4) HasFunnelStepBreakdown() bool`

HasFunnelStepBreakdown returns a boolean if a field has been set.

### SetFunnelStepBreakdownNil

`func (o *Source4) SetFunnelStepBreakdownNil(b bool)`

 SetFunnelStepBreakdownNil sets the value for FunnelStepBreakdown to be an explicit nil

### UnsetFunnelStepBreakdown
`func (o *Source4) UnsetFunnelStepBreakdown()`

UnsetFunnelStepBreakdown ensures that no value is present for FunnelStepBreakdown, not even an explicit nil
### GetFunnelTrendsDropOff

`func (o *Source4) GetFunnelTrendsDropOff() bool`

GetFunnelTrendsDropOff returns the FunnelTrendsDropOff field if non-nil, zero value otherwise.

### GetFunnelTrendsDropOffOk

`func (o *Source4) GetFunnelTrendsDropOffOk() (*bool, bool)`

GetFunnelTrendsDropOffOk returns a tuple with the FunnelTrendsDropOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelTrendsDropOff

`func (o *Source4) SetFunnelTrendsDropOff(v bool)`

SetFunnelTrendsDropOff sets FunnelTrendsDropOff field to given value.

### HasFunnelTrendsDropOff

`func (o *Source4) HasFunnelTrendsDropOff() bool`

HasFunnelTrendsDropOff returns a boolean if a field has been set.

### SetFunnelTrendsDropOffNil

`func (o *Source4) SetFunnelTrendsDropOffNil(b bool)`

 SetFunnelTrendsDropOffNil sets the value for FunnelTrendsDropOff to be an explicit nil

### UnsetFunnelTrendsDropOff
`func (o *Source4) UnsetFunnelTrendsDropOff()`

UnsetFunnelTrendsDropOff ensures that no value is present for FunnelTrendsDropOff, not even an explicit nil
### GetFunnelTrendsEntrancePeriodStart

`func (o *Source4) GetFunnelTrendsEntrancePeriodStart() string`

GetFunnelTrendsEntrancePeriodStart returns the FunnelTrendsEntrancePeriodStart field if non-nil, zero value otherwise.

### GetFunnelTrendsEntrancePeriodStartOk

`func (o *Source4) GetFunnelTrendsEntrancePeriodStartOk() (*string, bool)`

GetFunnelTrendsEntrancePeriodStartOk returns a tuple with the FunnelTrendsEntrancePeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelTrendsEntrancePeriodStart

`func (o *Source4) SetFunnelTrendsEntrancePeriodStart(v string)`

SetFunnelTrendsEntrancePeriodStart sets FunnelTrendsEntrancePeriodStart field to given value.

### HasFunnelTrendsEntrancePeriodStart

`func (o *Source4) HasFunnelTrendsEntrancePeriodStart() bool`

HasFunnelTrendsEntrancePeriodStart returns a boolean if a field has been set.

### SetFunnelTrendsEntrancePeriodStartNil

`func (o *Source4) SetFunnelTrendsEntrancePeriodStartNil(b bool)`

 SetFunnelTrendsEntrancePeriodStartNil sets the value for FunnelTrendsEntrancePeriodStart to be an explicit nil

### UnsetFunnelTrendsEntrancePeriodStart
`func (o *Source4) UnsetFunnelTrendsEntrancePeriodStart()`

UnsetFunnelTrendsEntrancePeriodStart ensures that no value is present for FunnelTrendsEntrancePeriodStart, not even an explicit nil
### GetFunnelCorrelationPersonConverted

`func (o *Source4) GetFunnelCorrelationPersonConverted() bool`

GetFunnelCorrelationPersonConverted returns the FunnelCorrelationPersonConverted field if non-nil, zero value otherwise.

### GetFunnelCorrelationPersonConvertedOk

`func (o *Source4) GetFunnelCorrelationPersonConvertedOk() (*bool, bool)`

GetFunnelCorrelationPersonConvertedOk returns a tuple with the FunnelCorrelationPersonConverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationPersonConverted

`func (o *Source4) SetFunnelCorrelationPersonConverted(v bool)`

SetFunnelCorrelationPersonConverted sets FunnelCorrelationPersonConverted field to given value.

### HasFunnelCorrelationPersonConverted

`func (o *Source4) HasFunnelCorrelationPersonConverted() bool`

HasFunnelCorrelationPersonConverted returns a boolean if a field has been set.

### SetFunnelCorrelationPersonConvertedNil

`func (o *Source4) SetFunnelCorrelationPersonConvertedNil(b bool)`

 SetFunnelCorrelationPersonConvertedNil sets the value for FunnelCorrelationPersonConverted to be an explicit nil

### UnsetFunnelCorrelationPersonConverted
`func (o *Source4) UnsetFunnelCorrelationPersonConverted()`

UnsetFunnelCorrelationPersonConverted ensures that no value is present for FunnelCorrelationPersonConverted, not even an explicit nil
### GetFunnelCorrelationPersonEntity

`func (o *Source4) GetFunnelCorrelationPersonEntity() Funnelcorrelationpersonentity`

GetFunnelCorrelationPersonEntity returns the FunnelCorrelationPersonEntity field if non-nil, zero value otherwise.

### GetFunnelCorrelationPersonEntityOk

`func (o *Source4) GetFunnelCorrelationPersonEntityOk() (*Funnelcorrelationpersonentity, bool)`

GetFunnelCorrelationPersonEntityOk returns a tuple with the FunnelCorrelationPersonEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationPersonEntity

`func (o *Source4) SetFunnelCorrelationPersonEntity(v Funnelcorrelationpersonentity)`

SetFunnelCorrelationPersonEntity sets FunnelCorrelationPersonEntity field to given value.

### HasFunnelCorrelationPersonEntity

`func (o *Source4) HasFunnelCorrelationPersonEntity() bool`

HasFunnelCorrelationPersonEntity returns a boolean if a field has been set.

### SetFunnelCorrelationPersonEntityNil

`func (o *Source4) SetFunnelCorrelationPersonEntityNil(b bool)`

 SetFunnelCorrelationPersonEntityNil sets the value for FunnelCorrelationPersonEntity to be an explicit nil

### UnsetFunnelCorrelationPersonEntity
`func (o *Source4) UnsetFunnelCorrelationPersonEntity()`

UnsetFunnelCorrelationPersonEntity ensures that no value is present for FunnelCorrelationPersonEntity, not even an explicit nil
### GetFunnelCorrelationPropertyValues

`func (o *Source4) GetFunnelCorrelationPropertyValues() []FixedpropertiesInner`

GetFunnelCorrelationPropertyValues returns the FunnelCorrelationPropertyValues field if non-nil, zero value otherwise.

### GetFunnelCorrelationPropertyValuesOk

`func (o *Source4) GetFunnelCorrelationPropertyValuesOk() (*[]FixedpropertiesInner, bool)`

GetFunnelCorrelationPropertyValuesOk returns a tuple with the FunnelCorrelationPropertyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationPropertyValues

`func (o *Source4) SetFunnelCorrelationPropertyValues(v []FixedpropertiesInner)`

SetFunnelCorrelationPropertyValues sets FunnelCorrelationPropertyValues field to given value.

### HasFunnelCorrelationPropertyValues

`func (o *Source4) HasFunnelCorrelationPropertyValues() bool`

HasFunnelCorrelationPropertyValues returns a boolean if a field has been set.

### SetFunnelCorrelationPropertyValuesNil

`func (o *Source4) SetFunnelCorrelationPropertyValuesNil(b bool)`

 SetFunnelCorrelationPropertyValuesNil sets the value for FunnelCorrelationPropertyValues to be an explicit nil

### UnsetFunnelCorrelationPropertyValues
`func (o *Source4) UnsetFunnelCorrelationPropertyValues()`

UnsetFunnelCorrelationPropertyValues ensures that no value is present for FunnelCorrelationPropertyValues, not even an explicit nil
### GetOperator

`func (o *Source4) GetOperator() StickinessOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Source4) GetOperatorOk() (*StickinessOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Source4) SetOperator(v StickinessOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Source4) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


