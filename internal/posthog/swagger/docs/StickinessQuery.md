# StickinessQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompareFilter** | Pointer to [**CompareFilter**](CompareFilter.md) |  | [optional] 
**DataColorTheme** | Pointer to **NullableFloat32** | Colors used in the insight&#39;s visualization | [optional] 
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** | Exclude internal and test users by applying the respective filters | [optional] [default to false]
**Interval** | Pointer to [**IntervalType**](IntervalType.md) |  | [optional] 
**IntervalCount** | Pointer to **NullableInt32** | How many intervals comprise a period. Only used for cohorts, otherwise default 1. | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "StickinessQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Properties** | Pointer to [**NullableProperties1**](Properties1.md) |  | [optional] [default to []]
**Response** | Pointer to [**StickinessQueryResponse**](StickinessQueryResponse.md) |  | [optional] 
**SamplingFactor** | Pointer to **NullableFloat32** | Sampling rate | [optional] 
**Series** | [**[]Series1Inner**](Series1Inner.md) | Events and actions to include | 
**StickinessFilter** | Pointer to [**StickinessFilter**](StickinessFilter.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewStickinessQuery

`func NewStickinessQuery(series []Series1Inner, ) *StickinessQuery`

NewStickinessQuery instantiates a new StickinessQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStickinessQueryWithDefaults

`func NewStickinessQueryWithDefaults() *StickinessQuery`

NewStickinessQueryWithDefaults instantiates a new StickinessQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompareFilter

`func (o *StickinessQuery) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *StickinessQuery) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *StickinessQuery) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *StickinessQuery) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetDataColorTheme

`func (o *StickinessQuery) GetDataColorTheme() float32`

GetDataColorTheme returns the DataColorTheme field if non-nil, zero value otherwise.

### GetDataColorThemeOk

`func (o *StickinessQuery) GetDataColorThemeOk() (*float32, bool)`

GetDataColorThemeOk returns a tuple with the DataColorTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataColorTheme

`func (o *StickinessQuery) SetDataColorTheme(v float32)`

SetDataColorTheme sets DataColorTheme field to given value.

### HasDataColorTheme

`func (o *StickinessQuery) HasDataColorTheme() bool`

HasDataColorTheme returns a boolean if a field has been set.

### SetDataColorThemeNil

`func (o *StickinessQuery) SetDataColorThemeNil(b bool)`

 SetDataColorThemeNil sets the value for DataColorTheme to be an explicit nil

### UnsetDataColorTheme
`func (o *StickinessQuery) UnsetDataColorTheme()`

UnsetDataColorTheme ensures that no value is present for DataColorTheme, not even an explicit nil
### GetDateRange

`func (o *StickinessQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *StickinessQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *StickinessQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *StickinessQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetFilterTestAccounts

`func (o *StickinessQuery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *StickinessQuery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *StickinessQuery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *StickinessQuery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *StickinessQuery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *StickinessQuery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetInterval

`func (o *StickinessQuery) GetInterval() IntervalType`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *StickinessQuery) GetIntervalOk() (*IntervalType, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *StickinessQuery) SetInterval(v IntervalType)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *StickinessQuery) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetIntervalCount

`func (o *StickinessQuery) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *StickinessQuery) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *StickinessQuery) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *StickinessQuery) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### SetIntervalCountNil

`func (o *StickinessQuery) SetIntervalCountNil(b bool)`

 SetIntervalCountNil sets the value for IntervalCount to be an explicit nil

### UnsetIntervalCount
`func (o *StickinessQuery) UnsetIntervalCount()`

UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
### GetKind

`func (o *StickinessQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StickinessQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StickinessQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StickinessQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *StickinessQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *StickinessQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *StickinessQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *StickinessQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetProperties

`func (o *StickinessQuery) GetProperties() Properties1`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *StickinessQuery) GetPropertiesOk() (*Properties1, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *StickinessQuery) SetProperties(v Properties1)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *StickinessQuery) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *StickinessQuery) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *StickinessQuery) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *StickinessQuery) GetResponse() StickinessQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *StickinessQuery) GetResponseOk() (*StickinessQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *StickinessQuery) SetResponse(v StickinessQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *StickinessQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSamplingFactor

`func (o *StickinessQuery) GetSamplingFactor() float32`

GetSamplingFactor returns the SamplingFactor field if non-nil, zero value otherwise.

### GetSamplingFactorOk

`func (o *StickinessQuery) GetSamplingFactorOk() (*float32, bool)`

GetSamplingFactorOk returns a tuple with the SamplingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingFactor

`func (o *StickinessQuery) SetSamplingFactor(v float32)`

SetSamplingFactor sets SamplingFactor field to given value.

### HasSamplingFactor

`func (o *StickinessQuery) HasSamplingFactor() bool`

HasSamplingFactor returns a boolean if a field has been set.

### SetSamplingFactorNil

`func (o *StickinessQuery) SetSamplingFactorNil(b bool)`

 SetSamplingFactorNil sets the value for SamplingFactor to be an explicit nil

### UnsetSamplingFactor
`func (o *StickinessQuery) UnsetSamplingFactor()`

UnsetSamplingFactor ensures that no value is present for SamplingFactor, not even an explicit nil
### GetSeries

`func (o *StickinessQuery) GetSeries() []Series1Inner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *StickinessQuery) GetSeriesOk() (*[]Series1Inner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *StickinessQuery) SetSeries(v []Series1Inner)`

SetSeries sets Series field to given value.


### GetStickinessFilter

`func (o *StickinessQuery) GetStickinessFilter() StickinessFilter`

GetStickinessFilter returns the StickinessFilter field if non-nil, zero value otherwise.

### GetStickinessFilterOk

`func (o *StickinessQuery) GetStickinessFilterOk() (*StickinessFilter, bool)`

GetStickinessFilterOk returns a tuple with the StickinessFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickinessFilter

`func (o *StickinessQuery) SetStickinessFilter(v StickinessFilter)`

SetStickinessFilter sets StickinessFilter field to given value.

### HasStickinessFilter

`func (o *StickinessQuery) HasStickinessFilter() bool`

HasStickinessFilter returns a boolean if a field has been set.

### GetTags

`func (o *StickinessQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StickinessQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StickinessQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StickinessQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *StickinessQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StickinessQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StickinessQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StickinessQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *StickinessQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *StickinessQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


