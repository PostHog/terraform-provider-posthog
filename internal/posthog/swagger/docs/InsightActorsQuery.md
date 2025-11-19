# InsightActorsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Breakdown** | Pointer to [**NullableBreakdown2**](Breakdown2.md) |  | [optional] [default to null]
**Compare** | Pointer to [**Compare**](Compare.md) |  | [optional] 
**Day** | Pointer to [**NullableDay**](Day.md) |  | [optional] [default to null]
**IncludeRecordings** | Pointer to **NullableBool** |  | [optional] 
**Interval** | Pointer to **NullableInt32** | An interval selected out of available intervals in source query. | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "InsightActorsQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Response** | Pointer to [**ActorsQueryResponse**](ActorsQueryResponse.md) |  | [optional] 
**Series** | Pointer to **NullableInt32** |  | [optional] 
**Source** | [**Source3**](Source3.md) |  | 
**Status** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewInsightActorsQuery

`func NewInsightActorsQuery(source Source3, ) *InsightActorsQuery`

NewInsightActorsQuery instantiates a new InsightActorsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightActorsQueryWithDefaults

`func NewInsightActorsQueryWithDefaults() *InsightActorsQuery`

NewInsightActorsQueryWithDefaults instantiates a new InsightActorsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdown

`func (o *InsightActorsQuery) GetBreakdown() Breakdown2`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *InsightActorsQuery) GetBreakdownOk() (*Breakdown2, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *InsightActorsQuery) SetBreakdown(v Breakdown2)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *InsightActorsQuery) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### SetBreakdownNil

`func (o *InsightActorsQuery) SetBreakdownNil(b bool)`

 SetBreakdownNil sets the value for Breakdown to be an explicit nil

### UnsetBreakdown
`func (o *InsightActorsQuery) UnsetBreakdown()`

UnsetBreakdown ensures that no value is present for Breakdown, not even an explicit nil
### GetCompare

`func (o *InsightActorsQuery) GetCompare() Compare`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *InsightActorsQuery) GetCompareOk() (*Compare, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *InsightActorsQuery) SetCompare(v Compare)`

SetCompare sets Compare field to given value.

### HasCompare

`func (o *InsightActorsQuery) HasCompare() bool`

HasCompare returns a boolean if a field has been set.

### GetDay

`func (o *InsightActorsQuery) GetDay() Day`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *InsightActorsQuery) GetDayOk() (*Day, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *InsightActorsQuery) SetDay(v Day)`

SetDay sets Day field to given value.

### HasDay

`func (o *InsightActorsQuery) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *InsightActorsQuery) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *InsightActorsQuery) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetIncludeRecordings

`func (o *InsightActorsQuery) GetIncludeRecordings() bool`

GetIncludeRecordings returns the IncludeRecordings field if non-nil, zero value otherwise.

### GetIncludeRecordingsOk

`func (o *InsightActorsQuery) GetIncludeRecordingsOk() (*bool, bool)`

GetIncludeRecordingsOk returns a tuple with the IncludeRecordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecordings

`func (o *InsightActorsQuery) SetIncludeRecordings(v bool)`

SetIncludeRecordings sets IncludeRecordings field to given value.

### HasIncludeRecordings

`func (o *InsightActorsQuery) HasIncludeRecordings() bool`

HasIncludeRecordings returns a boolean if a field has been set.

### SetIncludeRecordingsNil

`func (o *InsightActorsQuery) SetIncludeRecordingsNil(b bool)`

 SetIncludeRecordingsNil sets the value for IncludeRecordings to be an explicit nil

### UnsetIncludeRecordings
`func (o *InsightActorsQuery) UnsetIncludeRecordings()`

UnsetIncludeRecordings ensures that no value is present for IncludeRecordings, not even an explicit nil
### GetInterval

`func (o *InsightActorsQuery) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InsightActorsQuery) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InsightActorsQuery) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *InsightActorsQuery) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *InsightActorsQuery) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *InsightActorsQuery) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetKind

`func (o *InsightActorsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InsightActorsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InsightActorsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *InsightActorsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *InsightActorsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *InsightActorsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *InsightActorsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *InsightActorsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetResponse

`func (o *InsightActorsQuery) GetResponse() ActorsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *InsightActorsQuery) GetResponseOk() (*ActorsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *InsightActorsQuery) SetResponse(v ActorsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *InsightActorsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSeries

`func (o *InsightActorsQuery) GetSeries() int32`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *InsightActorsQuery) GetSeriesOk() (*int32, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *InsightActorsQuery) SetSeries(v int32)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *InsightActorsQuery) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeriesNil

`func (o *InsightActorsQuery) SetSeriesNil(b bool)`

 SetSeriesNil sets the value for Series to be an explicit nil

### UnsetSeries
`func (o *InsightActorsQuery) UnsetSeries()`

UnsetSeries ensures that no value is present for Series, not even an explicit nil
### GetSource

`func (o *InsightActorsQuery) GetSource() Source3`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InsightActorsQuery) GetSourceOk() (*Source3, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InsightActorsQuery) SetSource(v Source3)`

SetSource sets Source field to given value.


### GetStatus

`func (o *InsightActorsQuery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InsightActorsQuery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InsightActorsQuery) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InsightActorsQuery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *InsightActorsQuery) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *InsightActorsQuery) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTags

`func (o *InsightActorsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InsightActorsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InsightActorsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InsightActorsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *InsightActorsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InsightActorsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InsightActorsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InsightActorsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *InsightActorsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *InsightActorsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


