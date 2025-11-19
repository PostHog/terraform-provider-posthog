# PathsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationGroupTypeIndex** | Pointer to **NullableInt32** | Groups aggregation | [optional] 
**DataColorTheme** | Pointer to **NullableFloat32** | Colors used in the insight&#39;s visualization | [optional] 
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** | Exclude internal and test users by applying the respective filters | [optional] [default to false]
**FunnelPathsFilter** | Pointer to [**FunnelPathsFilter**](FunnelPathsFilter.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "PathsQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**PathsFilter** | [**PathsFilter**](PathsFilter.md) |  | 
**Properties** | Pointer to [**NullableProperties1**](Properties1.md) |  | [optional] [default to []]
**Response** | Pointer to [**PathsQueryResponse**](PathsQueryResponse.md) |  | [optional] 
**SamplingFactor** | Pointer to **NullableFloat32** | Sampling rate | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewPathsQuery

`func NewPathsQuery(pathsFilter PathsFilter, ) *PathsQuery`

NewPathsQuery instantiates a new PathsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathsQueryWithDefaults

`func NewPathsQueryWithDefaults() *PathsQuery`

NewPathsQueryWithDefaults instantiates a new PathsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationGroupTypeIndex

`func (o *PathsQuery) GetAggregationGroupTypeIndex() int32`

GetAggregationGroupTypeIndex returns the AggregationGroupTypeIndex field if non-nil, zero value otherwise.

### GetAggregationGroupTypeIndexOk

`func (o *PathsQuery) GetAggregationGroupTypeIndexOk() (*int32, bool)`

GetAggregationGroupTypeIndexOk returns a tuple with the AggregationGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationGroupTypeIndex

`func (o *PathsQuery) SetAggregationGroupTypeIndex(v int32)`

SetAggregationGroupTypeIndex sets AggregationGroupTypeIndex field to given value.

### HasAggregationGroupTypeIndex

`func (o *PathsQuery) HasAggregationGroupTypeIndex() bool`

HasAggregationGroupTypeIndex returns a boolean if a field has been set.

### SetAggregationGroupTypeIndexNil

`func (o *PathsQuery) SetAggregationGroupTypeIndexNil(b bool)`

 SetAggregationGroupTypeIndexNil sets the value for AggregationGroupTypeIndex to be an explicit nil

### UnsetAggregationGroupTypeIndex
`func (o *PathsQuery) UnsetAggregationGroupTypeIndex()`

UnsetAggregationGroupTypeIndex ensures that no value is present for AggregationGroupTypeIndex, not even an explicit nil
### GetDataColorTheme

`func (o *PathsQuery) GetDataColorTheme() float32`

GetDataColorTheme returns the DataColorTheme field if non-nil, zero value otherwise.

### GetDataColorThemeOk

`func (o *PathsQuery) GetDataColorThemeOk() (*float32, bool)`

GetDataColorThemeOk returns a tuple with the DataColorTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataColorTheme

`func (o *PathsQuery) SetDataColorTheme(v float32)`

SetDataColorTheme sets DataColorTheme field to given value.

### HasDataColorTheme

`func (o *PathsQuery) HasDataColorTheme() bool`

HasDataColorTheme returns a boolean if a field has been set.

### SetDataColorThemeNil

`func (o *PathsQuery) SetDataColorThemeNil(b bool)`

 SetDataColorThemeNil sets the value for DataColorTheme to be an explicit nil

### UnsetDataColorTheme
`func (o *PathsQuery) UnsetDataColorTheme()`

UnsetDataColorTheme ensures that no value is present for DataColorTheme, not even an explicit nil
### GetDateRange

`func (o *PathsQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *PathsQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *PathsQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *PathsQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetFilterTestAccounts

`func (o *PathsQuery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *PathsQuery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *PathsQuery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *PathsQuery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *PathsQuery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *PathsQuery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetFunnelPathsFilter

`func (o *PathsQuery) GetFunnelPathsFilter() FunnelPathsFilter`

GetFunnelPathsFilter returns the FunnelPathsFilter field if non-nil, zero value otherwise.

### GetFunnelPathsFilterOk

`func (o *PathsQuery) GetFunnelPathsFilterOk() (*FunnelPathsFilter, bool)`

GetFunnelPathsFilterOk returns a tuple with the FunnelPathsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelPathsFilter

`func (o *PathsQuery) SetFunnelPathsFilter(v FunnelPathsFilter)`

SetFunnelPathsFilter sets FunnelPathsFilter field to given value.

### HasFunnelPathsFilter

`func (o *PathsQuery) HasFunnelPathsFilter() bool`

HasFunnelPathsFilter returns a boolean if a field has been set.

### GetKind

`func (o *PathsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PathsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PathsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PathsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *PathsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *PathsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *PathsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *PathsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetPathsFilter

`func (o *PathsQuery) GetPathsFilter() PathsFilter`

GetPathsFilter returns the PathsFilter field if non-nil, zero value otherwise.

### GetPathsFilterOk

`func (o *PathsQuery) GetPathsFilterOk() (*PathsFilter, bool)`

GetPathsFilterOk returns a tuple with the PathsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathsFilter

`func (o *PathsQuery) SetPathsFilter(v PathsFilter)`

SetPathsFilter sets PathsFilter field to given value.


### GetProperties

`func (o *PathsQuery) GetProperties() Properties1`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PathsQuery) GetPropertiesOk() (*Properties1, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PathsQuery) SetProperties(v Properties1)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PathsQuery) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *PathsQuery) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *PathsQuery) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *PathsQuery) GetResponse() PathsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *PathsQuery) GetResponseOk() (*PathsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *PathsQuery) SetResponse(v PathsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *PathsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSamplingFactor

`func (o *PathsQuery) GetSamplingFactor() float32`

GetSamplingFactor returns the SamplingFactor field if non-nil, zero value otherwise.

### GetSamplingFactorOk

`func (o *PathsQuery) GetSamplingFactorOk() (*float32, bool)`

GetSamplingFactorOk returns a tuple with the SamplingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingFactor

`func (o *PathsQuery) SetSamplingFactor(v float32)`

SetSamplingFactor sets SamplingFactor field to given value.

### HasSamplingFactor

`func (o *PathsQuery) HasSamplingFactor() bool`

HasSamplingFactor returns a boolean if a field has been set.

### SetSamplingFactorNil

`func (o *PathsQuery) SetSamplingFactorNil(b bool)`

 SetSamplingFactorNil sets the value for SamplingFactor to be an explicit nil

### UnsetSamplingFactor
`func (o *PathsQuery) UnsetSamplingFactor()`

UnsetSamplingFactor ensures that no value is present for SamplingFactor, not even an explicit nil
### GetTags

`func (o *PathsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PathsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PathsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PathsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *PathsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PathsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PathsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PathsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *PathsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *PathsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


