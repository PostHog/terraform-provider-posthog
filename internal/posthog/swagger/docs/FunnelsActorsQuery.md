# FunnelsActorsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunnelStep** | Pointer to **NullableInt32** | Index of the step for which we want to get the timestamp for, per person. Positive for converted persons, negative for dropped of persons. | [optional] 
**FunnelStepBreakdown** | Pointer to [**NullableFunnelstepbreakdown**](Funnelstepbreakdown.md) |  | [optional] [default to null]
**FunnelTrendsDropOff** | Pointer to **NullableBool** |  | [optional] 
**FunnelTrendsEntrancePeriodStart** | Pointer to **NullableString** | Used together with &#x60;funnelTrendsDropOff&#x60; for funnels time conversion date for the persons modal. | [optional] 
**IncludeRecordings** | Pointer to **NullableBool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "FunnelsActorsQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Response** | Pointer to [**ActorsQueryResponse**](ActorsQueryResponse.md) |  | [optional] 
**Source** | [**FunnelsQuery**](FunnelsQuery.md) |  | 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewFunnelsActorsQuery

`func NewFunnelsActorsQuery(source FunnelsQuery, ) *FunnelsActorsQuery`

NewFunnelsActorsQuery instantiates a new FunnelsActorsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunnelsActorsQueryWithDefaults

`func NewFunnelsActorsQueryWithDefaults() *FunnelsActorsQuery`

NewFunnelsActorsQueryWithDefaults instantiates a new FunnelsActorsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunnelStep

`func (o *FunnelsActorsQuery) GetFunnelStep() int32`

GetFunnelStep returns the FunnelStep field if non-nil, zero value otherwise.

### GetFunnelStepOk

`func (o *FunnelsActorsQuery) GetFunnelStepOk() (*int32, bool)`

GetFunnelStepOk returns a tuple with the FunnelStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelStep

`func (o *FunnelsActorsQuery) SetFunnelStep(v int32)`

SetFunnelStep sets FunnelStep field to given value.

### HasFunnelStep

`func (o *FunnelsActorsQuery) HasFunnelStep() bool`

HasFunnelStep returns a boolean if a field has been set.

### SetFunnelStepNil

`func (o *FunnelsActorsQuery) SetFunnelStepNil(b bool)`

 SetFunnelStepNil sets the value for FunnelStep to be an explicit nil

### UnsetFunnelStep
`func (o *FunnelsActorsQuery) UnsetFunnelStep()`

UnsetFunnelStep ensures that no value is present for FunnelStep, not even an explicit nil
### GetFunnelStepBreakdown

`func (o *FunnelsActorsQuery) GetFunnelStepBreakdown() Funnelstepbreakdown`

GetFunnelStepBreakdown returns the FunnelStepBreakdown field if non-nil, zero value otherwise.

### GetFunnelStepBreakdownOk

`func (o *FunnelsActorsQuery) GetFunnelStepBreakdownOk() (*Funnelstepbreakdown, bool)`

GetFunnelStepBreakdownOk returns a tuple with the FunnelStepBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelStepBreakdown

`func (o *FunnelsActorsQuery) SetFunnelStepBreakdown(v Funnelstepbreakdown)`

SetFunnelStepBreakdown sets FunnelStepBreakdown field to given value.

### HasFunnelStepBreakdown

`func (o *FunnelsActorsQuery) HasFunnelStepBreakdown() bool`

HasFunnelStepBreakdown returns a boolean if a field has been set.

### SetFunnelStepBreakdownNil

`func (o *FunnelsActorsQuery) SetFunnelStepBreakdownNil(b bool)`

 SetFunnelStepBreakdownNil sets the value for FunnelStepBreakdown to be an explicit nil

### UnsetFunnelStepBreakdown
`func (o *FunnelsActorsQuery) UnsetFunnelStepBreakdown()`

UnsetFunnelStepBreakdown ensures that no value is present for FunnelStepBreakdown, not even an explicit nil
### GetFunnelTrendsDropOff

`func (o *FunnelsActorsQuery) GetFunnelTrendsDropOff() bool`

GetFunnelTrendsDropOff returns the FunnelTrendsDropOff field if non-nil, zero value otherwise.

### GetFunnelTrendsDropOffOk

`func (o *FunnelsActorsQuery) GetFunnelTrendsDropOffOk() (*bool, bool)`

GetFunnelTrendsDropOffOk returns a tuple with the FunnelTrendsDropOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelTrendsDropOff

`func (o *FunnelsActorsQuery) SetFunnelTrendsDropOff(v bool)`

SetFunnelTrendsDropOff sets FunnelTrendsDropOff field to given value.

### HasFunnelTrendsDropOff

`func (o *FunnelsActorsQuery) HasFunnelTrendsDropOff() bool`

HasFunnelTrendsDropOff returns a boolean if a field has been set.

### SetFunnelTrendsDropOffNil

`func (o *FunnelsActorsQuery) SetFunnelTrendsDropOffNil(b bool)`

 SetFunnelTrendsDropOffNil sets the value for FunnelTrendsDropOff to be an explicit nil

### UnsetFunnelTrendsDropOff
`func (o *FunnelsActorsQuery) UnsetFunnelTrendsDropOff()`

UnsetFunnelTrendsDropOff ensures that no value is present for FunnelTrendsDropOff, not even an explicit nil
### GetFunnelTrendsEntrancePeriodStart

`func (o *FunnelsActorsQuery) GetFunnelTrendsEntrancePeriodStart() string`

GetFunnelTrendsEntrancePeriodStart returns the FunnelTrendsEntrancePeriodStart field if non-nil, zero value otherwise.

### GetFunnelTrendsEntrancePeriodStartOk

`func (o *FunnelsActorsQuery) GetFunnelTrendsEntrancePeriodStartOk() (*string, bool)`

GetFunnelTrendsEntrancePeriodStartOk returns a tuple with the FunnelTrendsEntrancePeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelTrendsEntrancePeriodStart

`func (o *FunnelsActorsQuery) SetFunnelTrendsEntrancePeriodStart(v string)`

SetFunnelTrendsEntrancePeriodStart sets FunnelTrendsEntrancePeriodStart field to given value.

### HasFunnelTrendsEntrancePeriodStart

`func (o *FunnelsActorsQuery) HasFunnelTrendsEntrancePeriodStart() bool`

HasFunnelTrendsEntrancePeriodStart returns a boolean if a field has been set.

### SetFunnelTrendsEntrancePeriodStartNil

`func (o *FunnelsActorsQuery) SetFunnelTrendsEntrancePeriodStartNil(b bool)`

 SetFunnelTrendsEntrancePeriodStartNil sets the value for FunnelTrendsEntrancePeriodStart to be an explicit nil

### UnsetFunnelTrendsEntrancePeriodStart
`func (o *FunnelsActorsQuery) UnsetFunnelTrendsEntrancePeriodStart()`

UnsetFunnelTrendsEntrancePeriodStart ensures that no value is present for FunnelTrendsEntrancePeriodStart, not even an explicit nil
### GetIncludeRecordings

`func (o *FunnelsActorsQuery) GetIncludeRecordings() bool`

GetIncludeRecordings returns the IncludeRecordings field if non-nil, zero value otherwise.

### GetIncludeRecordingsOk

`func (o *FunnelsActorsQuery) GetIncludeRecordingsOk() (*bool, bool)`

GetIncludeRecordingsOk returns a tuple with the IncludeRecordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecordings

`func (o *FunnelsActorsQuery) SetIncludeRecordings(v bool)`

SetIncludeRecordings sets IncludeRecordings field to given value.

### HasIncludeRecordings

`func (o *FunnelsActorsQuery) HasIncludeRecordings() bool`

HasIncludeRecordings returns a boolean if a field has been set.

### SetIncludeRecordingsNil

`func (o *FunnelsActorsQuery) SetIncludeRecordingsNil(b bool)`

 SetIncludeRecordingsNil sets the value for IncludeRecordings to be an explicit nil

### UnsetIncludeRecordings
`func (o *FunnelsActorsQuery) UnsetIncludeRecordings()`

UnsetIncludeRecordings ensures that no value is present for IncludeRecordings, not even an explicit nil
### GetKind

`func (o *FunnelsActorsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FunnelsActorsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FunnelsActorsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FunnelsActorsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *FunnelsActorsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *FunnelsActorsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *FunnelsActorsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *FunnelsActorsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetResponse

`func (o *FunnelsActorsQuery) GetResponse() ActorsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *FunnelsActorsQuery) GetResponseOk() (*ActorsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *FunnelsActorsQuery) SetResponse(v ActorsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *FunnelsActorsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSource

`func (o *FunnelsActorsQuery) GetSource() FunnelsQuery`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FunnelsActorsQuery) GetSourceOk() (*FunnelsQuery, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FunnelsActorsQuery) SetSource(v FunnelsQuery)`

SetSource sets Source field to given value.


### GetTags

`func (o *FunnelsActorsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FunnelsActorsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FunnelsActorsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FunnelsActorsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *FunnelsActorsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FunnelsActorsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FunnelsActorsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FunnelsActorsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *FunnelsActorsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *FunnelsActorsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


