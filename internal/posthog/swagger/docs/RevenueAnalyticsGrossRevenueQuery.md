# RevenueAnalyticsGrossRevenueQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Breakdown** | [**[]RevenueAnalyticsBreakdown**](RevenueAnalyticsBreakdown.md) |  | 
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**Interval** | [**SimpleIntervalType**](SimpleIntervalType.md) |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "RevenueAnalyticsGrossRevenueQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Properties** | [**[]RevenueAnalyticsPropertyFilter**](RevenueAnalyticsPropertyFilter.md) |  | 
**Response** | Pointer to [**RevenueAnalyticsGrossRevenueQueryResponse**](RevenueAnalyticsGrossRevenueQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewRevenueAnalyticsGrossRevenueQuery

`func NewRevenueAnalyticsGrossRevenueQuery(breakdown []RevenueAnalyticsBreakdown, interval SimpleIntervalType, properties []RevenueAnalyticsPropertyFilter, ) *RevenueAnalyticsGrossRevenueQuery`

NewRevenueAnalyticsGrossRevenueQuery instantiates a new RevenueAnalyticsGrossRevenueQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueAnalyticsGrossRevenueQueryWithDefaults

`func NewRevenueAnalyticsGrossRevenueQueryWithDefaults() *RevenueAnalyticsGrossRevenueQuery`

NewRevenueAnalyticsGrossRevenueQueryWithDefaults instantiates a new RevenueAnalyticsGrossRevenueQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdown

`func (o *RevenueAnalyticsGrossRevenueQuery) GetBreakdown() []RevenueAnalyticsBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *RevenueAnalyticsGrossRevenueQuery) GetBreakdownOk() (*[]RevenueAnalyticsBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *RevenueAnalyticsGrossRevenueQuery) SetBreakdown(v []RevenueAnalyticsBreakdown)`

SetBreakdown sets Breakdown field to given value.


### GetDateRange

`func (o *RevenueAnalyticsGrossRevenueQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *RevenueAnalyticsGrossRevenueQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *RevenueAnalyticsGrossRevenueQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *RevenueAnalyticsGrossRevenueQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetInterval

`func (o *RevenueAnalyticsGrossRevenueQuery) GetInterval() SimpleIntervalType`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RevenueAnalyticsGrossRevenueQuery) GetIntervalOk() (*SimpleIntervalType, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RevenueAnalyticsGrossRevenueQuery) SetInterval(v SimpleIntervalType)`

SetInterval sets Interval field to given value.


### GetKind

`func (o *RevenueAnalyticsGrossRevenueQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RevenueAnalyticsGrossRevenueQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RevenueAnalyticsGrossRevenueQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RevenueAnalyticsGrossRevenueQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *RevenueAnalyticsGrossRevenueQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *RevenueAnalyticsGrossRevenueQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *RevenueAnalyticsGrossRevenueQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *RevenueAnalyticsGrossRevenueQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetProperties

`func (o *RevenueAnalyticsGrossRevenueQuery) GetProperties() []RevenueAnalyticsPropertyFilter`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RevenueAnalyticsGrossRevenueQuery) GetPropertiesOk() (*[]RevenueAnalyticsPropertyFilter, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RevenueAnalyticsGrossRevenueQuery) SetProperties(v []RevenueAnalyticsPropertyFilter)`

SetProperties sets Properties field to given value.


### GetResponse

`func (o *RevenueAnalyticsGrossRevenueQuery) GetResponse() RevenueAnalyticsGrossRevenueQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *RevenueAnalyticsGrossRevenueQuery) GetResponseOk() (*RevenueAnalyticsGrossRevenueQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *RevenueAnalyticsGrossRevenueQuery) SetResponse(v RevenueAnalyticsGrossRevenueQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *RevenueAnalyticsGrossRevenueQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *RevenueAnalyticsGrossRevenueQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RevenueAnalyticsGrossRevenueQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RevenueAnalyticsGrossRevenueQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RevenueAnalyticsGrossRevenueQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *RevenueAnalyticsGrossRevenueQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RevenueAnalyticsGrossRevenueQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RevenueAnalyticsGrossRevenueQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RevenueAnalyticsGrossRevenueQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *RevenueAnalyticsGrossRevenueQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *RevenueAnalyticsGrossRevenueQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


