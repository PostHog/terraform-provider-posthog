# StickinessActorsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compare** | Pointer to [**Compare**](Compare.md) |  | [optional] 
**Day** | Pointer to [**NullableDay**](Day.md) |  | [optional] [default to null]
**IncludeRecordings** | Pointer to **NullableBool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "StickinessActorsQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Operator** | Pointer to [**StickinessOperator**](StickinessOperator.md) |  | [optional] 
**Response** | Pointer to [**ActorsQueryResponse**](ActorsQueryResponse.md) |  | [optional] 
**Series** | Pointer to **NullableInt32** |  | [optional] 
**Source** | [**StickinessQuery**](StickinessQuery.md) |  | 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewStickinessActorsQuery

`func NewStickinessActorsQuery(source StickinessQuery, ) *StickinessActorsQuery`

NewStickinessActorsQuery instantiates a new StickinessActorsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStickinessActorsQueryWithDefaults

`func NewStickinessActorsQueryWithDefaults() *StickinessActorsQuery`

NewStickinessActorsQueryWithDefaults instantiates a new StickinessActorsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompare

`func (o *StickinessActorsQuery) GetCompare() Compare`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *StickinessActorsQuery) GetCompareOk() (*Compare, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *StickinessActorsQuery) SetCompare(v Compare)`

SetCompare sets Compare field to given value.

### HasCompare

`func (o *StickinessActorsQuery) HasCompare() bool`

HasCompare returns a boolean if a field has been set.

### GetDay

`func (o *StickinessActorsQuery) GetDay() Day`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *StickinessActorsQuery) GetDayOk() (*Day, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *StickinessActorsQuery) SetDay(v Day)`

SetDay sets Day field to given value.

### HasDay

`func (o *StickinessActorsQuery) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *StickinessActorsQuery) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *StickinessActorsQuery) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetIncludeRecordings

`func (o *StickinessActorsQuery) GetIncludeRecordings() bool`

GetIncludeRecordings returns the IncludeRecordings field if non-nil, zero value otherwise.

### GetIncludeRecordingsOk

`func (o *StickinessActorsQuery) GetIncludeRecordingsOk() (*bool, bool)`

GetIncludeRecordingsOk returns a tuple with the IncludeRecordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecordings

`func (o *StickinessActorsQuery) SetIncludeRecordings(v bool)`

SetIncludeRecordings sets IncludeRecordings field to given value.

### HasIncludeRecordings

`func (o *StickinessActorsQuery) HasIncludeRecordings() bool`

HasIncludeRecordings returns a boolean if a field has been set.

### SetIncludeRecordingsNil

`func (o *StickinessActorsQuery) SetIncludeRecordingsNil(b bool)`

 SetIncludeRecordingsNil sets the value for IncludeRecordings to be an explicit nil

### UnsetIncludeRecordings
`func (o *StickinessActorsQuery) UnsetIncludeRecordings()`

UnsetIncludeRecordings ensures that no value is present for IncludeRecordings, not even an explicit nil
### GetKind

`func (o *StickinessActorsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StickinessActorsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StickinessActorsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StickinessActorsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *StickinessActorsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *StickinessActorsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *StickinessActorsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *StickinessActorsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOperator

`func (o *StickinessActorsQuery) GetOperator() StickinessOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *StickinessActorsQuery) GetOperatorOk() (*StickinessOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *StickinessActorsQuery) SetOperator(v StickinessOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *StickinessActorsQuery) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetResponse

`func (o *StickinessActorsQuery) GetResponse() ActorsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *StickinessActorsQuery) GetResponseOk() (*ActorsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *StickinessActorsQuery) SetResponse(v ActorsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *StickinessActorsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSeries

`func (o *StickinessActorsQuery) GetSeries() int32`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *StickinessActorsQuery) GetSeriesOk() (*int32, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *StickinessActorsQuery) SetSeries(v int32)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *StickinessActorsQuery) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeriesNil

`func (o *StickinessActorsQuery) SetSeriesNil(b bool)`

 SetSeriesNil sets the value for Series to be an explicit nil

### UnsetSeries
`func (o *StickinessActorsQuery) UnsetSeries()`

UnsetSeries ensures that no value is present for Series, not even an explicit nil
### GetSource

`func (o *StickinessActorsQuery) GetSource() StickinessQuery`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StickinessActorsQuery) GetSourceOk() (*StickinessQuery, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StickinessActorsQuery) SetSource(v StickinessQuery)`

SetSource sets Source field to given value.


### GetTags

`func (o *StickinessActorsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StickinessActorsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StickinessActorsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StickinessActorsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *StickinessActorsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StickinessActorsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StickinessActorsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StickinessActorsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *StickinessActorsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *StickinessActorsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


