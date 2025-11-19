# LogsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRange** | [**DateRange**](DateRange.md) |  | 
**FilterGroup** | [**PropertyGroupFilter**](PropertyGroupFilter.md) |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "LogsQuery"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**OrderBy** | Pointer to [**OrderBy3**](OrderBy3.md) |  | [optional] 
**Response** | Pointer to [**LogsQueryResponse**](LogsQueryResponse.md) |  | [optional] 
**SearchTerm** | Pointer to **NullableString** |  | [optional] 
**ServiceNames** | **[]string** |  | 
**SeverityLevels** | [**[]LogSeverityLevel**](LogSeverityLevel.md) |  | 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewLogsQuery

`func NewLogsQuery(dateRange DateRange, filterGroup PropertyGroupFilter, serviceNames []string, severityLevels []LogSeverityLevel, ) *LogsQuery`

NewLogsQuery instantiates a new LogsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsQueryWithDefaults

`func NewLogsQueryWithDefaults() *LogsQuery`

NewLogsQueryWithDefaults instantiates a new LogsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRange

`func (o *LogsQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *LogsQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *LogsQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.


### GetFilterGroup

`func (o *LogsQuery) GetFilterGroup() PropertyGroupFilter`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *LogsQuery) GetFilterGroupOk() (*PropertyGroupFilter, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *LogsQuery) SetFilterGroup(v PropertyGroupFilter)`

SetFilterGroup sets FilterGroup field to given value.


### GetKind

`func (o *LogsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LogsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LogsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LogsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *LogsQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LogsQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LogsQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *LogsQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *LogsQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *LogsQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *LogsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *LogsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *LogsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *LogsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *LogsQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *LogsQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *LogsQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *LogsQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *LogsQuery) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *LogsQuery) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetOrderBy

`func (o *LogsQuery) GetOrderBy() OrderBy3`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *LogsQuery) GetOrderByOk() (*OrderBy3, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *LogsQuery) SetOrderBy(v OrderBy3)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *LogsQuery) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetResponse

`func (o *LogsQuery) GetResponse() LogsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *LogsQuery) GetResponseOk() (*LogsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *LogsQuery) SetResponse(v LogsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *LogsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSearchTerm

`func (o *LogsQuery) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *LogsQuery) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *LogsQuery) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.

### HasSearchTerm

`func (o *LogsQuery) HasSearchTerm() bool`

HasSearchTerm returns a boolean if a field has been set.

### SetSearchTermNil

`func (o *LogsQuery) SetSearchTermNil(b bool)`

 SetSearchTermNil sets the value for SearchTerm to be an explicit nil

### UnsetSearchTerm
`func (o *LogsQuery) UnsetSearchTerm()`

UnsetSearchTerm ensures that no value is present for SearchTerm, not even an explicit nil
### GetServiceNames

`func (o *LogsQuery) GetServiceNames() []string`

GetServiceNames returns the ServiceNames field if non-nil, zero value otherwise.

### GetServiceNamesOk

`func (o *LogsQuery) GetServiceNamesOk() (*[]string, bool)`

GetServiceNamesOk returns a tuple with the ServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNames

`func (o *LogsQuery) SetServiceNames(v []string)`

SetServiceNames sets ServiceNames field to given value.


### GetSeverityLevels

`func (o *LogsQuery) GetSeverityLevels() []LogSeverityLevel`

GetSeverityLevels returns the SeverityLevels field if non-nil, zero value otherwise.

### GetSeverityLevelsOk

`func (o *LogsQuery) GetSeverityLevelsOk() (*[]LogSeverityLevel, bool)`

GetSeverityLevelsOk returns a tuple with the SeverityLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevels

`func (o *LogsQuery) SetSeverityLevels(v []LogSeverityLevel)`

SetSeverityLevels sets SeverityLevels field to given value.


### GetTags

`func (o *LogsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LogsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LogsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LogsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *LogsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LogsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LogsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LogsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *LogsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *LogsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


