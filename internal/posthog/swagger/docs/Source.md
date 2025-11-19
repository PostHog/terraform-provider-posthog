# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Breakdown** | Pointer to [**NullableBreakdown2**](Breakdown2.md) |  | [optional] [default to null]
**Compare** | Pointer to [**Compare**](Compare.md) |  | [optional] 
**Day** | Pointer to [**NullableDay**](Day.md) |  | [optional] [default to null]
**IncludeRecordings** | Pointer to **NullableBool** |  | [optional] 
**Interval** | Pointer to **NullableInt32** | An interval selected out of available intervals in source query. | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "HogQLQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Response** | Pointer to [**HogQLQueryResponse**](HogQLQueryResponse.md) |  | [optional] 
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
**Explain** | Pointer to **NullableBool** |  | [optional] 
**Filters** | Pointer to [**HogQLFilters**](HogQLFilters.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Client provided name of the query | [optional] 
**Query** | **string** |  | 
**Values** | Pointer to **map[string]interface{}** | Constant values that can be referenced with the {placeholder} syntax in the query | [optional] 
**Variables** | Pointer to [**map[string]HogQLVariable**](HogQLVariable.md) | Variables to be substituted into the query | [optional] 

## Methods

### NewSource

`func NewSource(source StickinessQuery, query string, ) *Source`

NewSource instantiates a new Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceWithDefaults

`func NewSourceWithDefaults() *Source`

NewSourceWithDefaults instantiates a new Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdown

`func (o *Source) GetBreakdown() Breakdown2`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *Source) GetBreakdownOk() (*Breakdown2, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *Source) SetBreakdown(v Breakdown2)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *Source) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### SetBreakdownNil

`func (o *Source) SetBreakdownNil(b bool)`

 SetBreakdownNil sets the value for Breakdown to be an explicit nil

### UnsetBreakdown
`func (o *Source) UnsetBreakdown()`

UnsetBreakdown ensures that no value is present for Breakdown, not even an explicit nil
### GetCompare

`func (o *Source) GetCompare() Compare`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *Source) GetCompareOk() (*Compare, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *Source) SetCompare(v Compare)`

SetCompare sets Compare field to given value.

### HasCompare

`func (o *Source) HasCompare() bool`

HasCompare returns a boolean if a field has been set.

### GetDay

`func (o *Source) GetDay() Day`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *Source) GetDayOk() (*Day, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *Source) SetDay(v Day)`

SetDay sets Day field to given value.

### HasDay

`func (o *Source) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *Source) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *Source) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetIncludeRecordings

`func (o *Source) GetIncludeRecordings() bool`

GetIncludeRecordings returns the IncludeRecordings field if non-nil, zero value otherwise.

### GetIncludeRecordingsOk

`func (o *Source) GetIncludeRecordingsOk() (*bool, bool)`

GetIncludeRecordingsOk returns a tuple with the IncludeRecordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecordings

`func (o *Source) SetIncludeRecordings(v bool)`

SetIncludeRecordings sets IncludeRecordings field to given value.

### HasIncludeRecordings

`func (o *Source) HasIncludeRecordings() bool`

HasIncludeRecordings returns a boolean if a field has been set.

### SetIncludeRecordingsNil

`func (o *Source) SetIncludeRecordingsNil(b bool)`

 SetIncludeRecordingsNil sets the value for IncludeRecordings to be an explicit nil

### UnsetIncludeRecordings
`func (o *Source) UnsetIncludeRecordings()`

UnsetIncludeRecordings ensures that no value is present for IncludeRecordings, not even an explicit nil
### GetInterval

`func (o *Source) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Source) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Source) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Source) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *Source) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *Source) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetKind

`func (o *Source) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Source) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Source) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Source) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *Source) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *Source) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *Source) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *Source) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetResponse

`func (o *Source) GetResponse() HogQLQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Source) GetResponseOk() (*HogQLQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Source) SetResponse(v HogQLQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Source) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSeries

`func (o *Source) GetSeries() int32`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Source) GetSeriesOk() (*int32, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Source) SetSeries(v int32)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *Source) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeriesNil

`func (o *Source) SetSeriesNil(b bool)`

 SetSeriesNil sets the value for Series to be an explicit nil

### UnsetSeries
`func (o *Source) UnsetSeries()`

UnsetSeries ensures that no value is present for Series, not even an explicit nil
### GetSource

`func (o *Source) GetSource() StickinessQuery`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Source) GetSourceOk() (*StickinessQuery, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Source) SetSource(v StickinessQuery)`

SetSource sets Source field to given value.


### GetStatus

`func (o *Source) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Source) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Source) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Source) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Source) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Source) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTags

`func (o *Source) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Source) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Source) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Source) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *Source) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Source) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Source) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Source) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Source) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Source) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetFunnelStep

`func (o *Source) GetFunnelStep() int32`

GetFunnelStep returns the FunnelStep field if non-nil, zero value otherwise.

### GetFunnelStepOk

`func (o *Source) GetFunnelStepOk() (*int32, bool)`

GetFunnelStepOk returns a tuple with the FunnelStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelStep

`func (o *Source) SetFunnelStep(v int32)`

SetFunnelStep sets FunnelStep field to given value.

### HasFunnelStep

`func (o *Source) HasFunnelStep() bool`

HasFunnelStep returns a boolean if a field has been set.

### SetFunnelStepNil

`func (o *Source) SetFunnelStepNil(b bool)`

 SetFunnelStepNil sets the value for FunnelStep to be an explicit nil

### UnsetFunnelStep
`func (o *Source) UnsetFunnelStep()`

UnsetFunnelStep ensures that no value is present for FunnelStep, not even an explicit nil
### GetFunnelStepBreakdown

`func (o *Source) GetFunnelStepBreakdown() Funnelstepbreakdown`

GetFunnelStepBreakdown returns the FunnelStepBreakdown field if non-nil, zero value otherwise.

### GetFunnelStepBreakdownOk

`func (o *Source) GetFunnelStepBreakdownOk() (*Funnelstepbreakdown, bool)`

GetFunnelStepBreakdownOk returns a tuple with the FunnelStepBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelStepBreakdown

`func (o *Source) SetFunnelStepBreakdown(v Funnelstepbreakdown)`

SetFunnelStepBreakdown sets FunnelStepBreakdown field to given value.

### HasFunnelStepBreakdown

`func (o *Source) HasFunnelStepBreakdown() bool`

HasFunnelStepBreakdown returns a boolean if a field has been set.

### SetFunnelStepBreakdownNil

`func (o *Source) SetFunnelStepBreakdownNil(b bool)`

 SetFunnelStepBreakdownNil sets the value for FunnelStepBreakdown to be an explicit nil

### UnsetFunnelStepBreakdown
`func (o *Source) UnsetFunnelStepBreakdown()`

UnsetFunnelStepBreakdown ensures that no value is present for FunnelStepBreakdown, not even an explicit nil
### GetFunnelTrendsDropOff

`func (o *Source) GetFunnelTrendsDropOff() bool`

GetFunnelTrendsDropOff returns the FunnelTrendsDropOff field if non-nil, zero value otherwise.

### GetFunnelTrendsDropOffOk

`func (o *Source) GetFunnelTrendsDropOffOk() (*bool, bool)`

GetFunnelTrendsDropOffOk returns a tuple with the FunnelTrendsDropOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelTrendsDropOff

`func (o *Source) SetFunnelTrendsDropOff(v bool)`

SetFunnelTrendsDropOff sets FunnelTrendsDropOff field to given value.

### HasFunnelTrendsDropOff

`func (o *Source) HasFunnelTrendsDropOff() bool`

HasFunnelTrendsDropOff returns a boolean if a field has been set.

### SetFunnelTrendsDropOffNil

`func (o *Source) SetFunnelTrendsDropOffNil(b bool)`

 SetFunnelTrendsDropOffNil sets the value for FunnelTrendsDropOff to be an explicit nil

### UnsetFunnelTrendsDropOff
`func (o *Source) UnsetFunnelTrendsDropOff()`

UnsetFunnelTrendsDropOff ensures that no value is present for FunnelTrendsDropOff, not even an explicit nil
### GetFunnelTrendsEntrancePeriodStart

`func (o *Source) GetFunnelTrendsEntrancePeriodStart() string`

GetFunnelTrendsEntrancePeriodStart returns the FunnelTrendsEntrancePeriodStart field if non-nil, zero value otherwise.

### GetFunnelTrendsEntrancePeriodStartOk

`func (o *Source) GetFunnelTrendsEntrancePeriodStartOk() (*string, bool)`

GetFunnelTrendsEntrancePeriodStartOk returns a tuple with the FunnelTrendsEntrancePeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelTrendsEntrancePeriodStart

`func (o *Source) SetFunnelTrendsEntrancePeriodStart(v string)`

SetFunnelTrendsEntrancePeriodStart sets FunnelTrendsEntrancePeriodStart field to given value.

### HasFunnelTrendsEntrancePeriodStart

`func (o *Source) HasFunnelTrendsEntrancePeriodStart() bool`

HasFunnelTrendsEntrancePeriodStart returns a boolean if a field has been set.

### SetFunnelTrendsEntrancePeriodStartNil

`func (o *Source) SetFunnelTrendsEntrancePeriodStartNil(b bool)`

 SetFunnelTrendsEntrancePeriodStartNil sets the value for FunnelTrendsEntrancePeriodStart to be an explicit nil

### UnsetFunnelTrendsEntrancePeriodStart
`func (o *Source) UnsetFunnelTrendsEntrancePeriodStart()`

UnsetFunnelTrendsEntrancePeriodStart ensures that no value is present for FunnelTrendsEntrancePeriodStart, not even an explicit nil
### GetFunnelCorrelationPersonConverted

`func (o *Source) GetFunnelCorrelationPersonConverted() bool`

GetFunnelCorrelationPersonConverted returns the FunnelCorrelationPersonConverted field if non-nil, zero value otherwise.

### GetFunnelCorrelationPersonConvertedOk

`func (o *Source) GetFunnelCorrelationPersonConvertedOk() (*bool, bool)`

GetFunnelCorrelationPersonConvertedOk returns a tuple with the FunnelCorrelationPersonConverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationPersonConverted

`func (o *Source) SetFunnelCorrelationPersonConverted(v bool)`

SetFunnelCorrelationPersonConverted sets FunnelCorrelationPersonConverted field to given value.

### HasFunnelCorrelationPersonConverted

`func (o *Source) HasFunnelCorrelationPersonConverted() bool`

HasFunnelCorrelationPersonConverted returns a boolean if a field has been set.

### SetFunnelCorrelationPersonConvertedNil

`func (o *Source) SetFunnelCorrelationPersonConvertedNil(b bool)`

 SetFunnelCorrelationPersonConvertedNil sets the value for FunnelCorrelationPersonConverted to be an explicit nil

### UnsetFunnelCorrelationPersonConverted
`func (o *Source) UnsetFunnelCorrelationPersonConverted()`

UnsetFunnelCorrelationPersonConverted ensures that no value is present for FunnelCorrelationPersonConverted, not even an explicit nil
### GetFunnelCorrelationPersonEntity

`func (o *Source) GetFunnelCorrelationPersonEntity() Funnelcorrelationpersonentity`

GetFunnelCorrelationPersonEntity returns the FunnelCorrelationPersonEntity field if non-nil, zero value otherwise.

### GetFunnelCorrelationPersonEntityOk

`func (o *Source) GetFunnelCorrelationPersonEntityOk() (*Funnelcorrelationpersonentity, bool)`

GetFunnelCorrelationPersonEntityOk returns a tuple with the FunnelCorrelationPersonEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationPersonEntity

`func (o *Source) SetFunnelCorrelationPersonEntity(v Funnelcorrelationpersonentity)`

SetFunnelCorrelationPersonEntity sets FunnelCorrelationPersonEntity field to given value.

### HasFunnelCorrelationPersonEntity

`func (o *Source) HasFunnelCorrelationPersonEntity() bool`

HasFunnelCorrelationPersonEntity returns a boolean if a field has been set.

### SetFunnelCorrelationPersonEntityNil

`func (o *Source) SetFunnelCorrelationPersonEntityNil(b bool)`

 SetFunnelCorrelationPersonEntityNil sets the value for FunnelCorrelationPersonEntity to be an explicit nil

### UnsetFunnelCorrelationPersonEntity
`func (o *Source) UnsetFunnelCorrelationPersonEntity()`

UnsetFunnelCorrelationPersonEntity ensures that no value is present for FunnelCorrelationPersonEntity, not even an explicit nil
### GetFunnelCorrelationPropertyValues

`func (o *Source) GetFunnelCorrelationPropertyValues() []FixedpropertiesInner`

GetFunnelCorrelationPropertyValues returns the FunnelCorrelationPropertyValues field if non-nil, zero value otherwise.

### GetFunnelCorrelationPropertyValuesOk

`func (o *Source) GetFunnelCorrelationPropertyValuesOk() (*[]FixedpropertiesInner, bool)`

GetFunnelCorrelationPropertyValuesOk returns a tuple with the FunnelCorrelationPropertyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationPropertyValues

`func (o *Source) SetFunnelCorrelationPropertyValues(v []FixedpropertiesInner)`

SetFunnelCorrelationPropertyValues sets FunnelCorrelationPropertyValues field to given value.

### HasFunnelCorrelationPropertyValues

`func (o *Source) HasFunnelCorrelationPropertyValues() bool`

HasFunnelCorrelationPropertyValues returns a boolean if a field has been set.

### SetFunnelCorrelationPropertyValuesNil

`func (o *Source) SetFunnelCorrelationPropertyValuesNil(b bool)`

 SetFunnelCorrelationPropertyValuesNil sets the value for FunnelCorrelationPropertyValues to be an explicit nil

### UnsetFunnelCorrelationPropertyValues
`func (o *Source) UnsetFunnelCorrelationPropertyValues()`

UnsetFunnelCorrelationPropertyValues ensures that no value is present for FunnelCorrelationPropertyValues, not even an explicit nil
### GetOperator

`func (o *Source) GetOperator() StickinessOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Source) GetOperatorOk() (*StickinessOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Source) SetOperator(v StickinessOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Source) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetExplain

`func (o *Source) GetExplain() bool`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *Source) GetExplainOk() (*bool, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *Source) SetExplain(v bool)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *Source) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### SetExplainNil

`func (o *Source) SetExplainNil(b bool)`

 SetExplainNil sets the value for Explain to be an explicit nil

### UnsetExplain
`func (o *Source) UnsetExplain()`

UnsetExplain ensures that no value is present for Explain, not even an explicit nil
### GetFilters

`func (o *Source) GetFilters() HogQLFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Source) GetFiltersOk() (*HogQLFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Source) SetFilters(v HogQLFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Source) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetName

`func (o *Source) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Source) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Source) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Source) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Source) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Source) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuery

`func (o *Source) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Source) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Source) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetValues

`func (o *Source) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Source) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Source) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *Source) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *Source) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *Source) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetVariables

`func (o *Source) GetVariables() map[string]HogQLVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Source) GetVariablesOk() (*map[string]HogQLVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Source) SetVariables(v map[string]HogQLVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Source) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *Source) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *Source) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


