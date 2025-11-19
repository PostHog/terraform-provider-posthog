# RevenueAnalyticsMetricsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Breakdown** | [**[]RevenueAnalyticsBreakdown**](RevenueAnalyticsBreakdown.md) |  | 
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**Interval** | [**SimpleIntervalType**](SimpleIntervalType.md) |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "RevenueAnalyticsMetricsQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Properties** | [**[]RevenueAnalyticsPropertyFilter**](RevenueAnalyticsPropertyFilter.md) |  | 
**Response** | Pointer to [**RevenueAnalyticsMetricsQueryResponse**](RevenueAnalyticsMetricsQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewRevenueAnalyticsMetricsQuery

`func NewRevenueAnalyticsMetricsQuery(breakdown []RevenueAnalyticsBreakdown, interval SimpleIntervalType, properties []RevenueAnalyticsPropertyFilter, ) *RevenueAnalyticsMetricsQuery`

NewRevenueAnalyticsMetricsQuery instantiates a new RevenueAnalyticsMetricsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueAnalyticsMetricsQueryWithDefaults

`func NewRevenueAnalyticsMetricsQueryWithDefaults() *RevenueAnalyticsMetricsQuery`

NewRevenueAnalyticsMetricsQueryWithDefaults instantiates a new RevenueAnalyticsMetricsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdown

`func (o *RevenueAnalyticsMetricsQuery) GetBreakdown() []RevenueAnalyticsBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *RevenueAnalyticsMetricsQuery) GetBreakdownOk() (*[]RevenueAnalyticsBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *RevenueAnalyticsMetricsQuery) SetBreakdown(v []RevenueAnalyticsBreakdown)`

SetBreakdown sets Breakdown field to given value.


### GetDateRange

`func (o *RevenueAnalyticsMetricsQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *RevenueAnalyticsMetricsQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *RevenueAnalyticsMetricsQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *RevenueAnalyticsMetricsQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetInterval

`func (o *RevenueAnalyticsMetricsQuery) GetInterval() SimpleIntervalType`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *RevenueAnalyticsMetricsQuery) GetIntervalOk() (*SimpleIntervalType, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *RevenueAnalyticsMetricsQuery) SetInterval(v SimpleIntervalType)`

SetInterval sets Interval field to given value.


### GetKind

`func (o *RevenueAnalyticsMetricsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RevenueAnalyticsMetricsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RevenueAnalyticsMetricsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RevenueAnalyticsMetricsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *RevenueAnalyticsMetricsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *RevenueAnalyticsMetricsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *RevenueAnalyticsMetricsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *RevenueAnalyticsMetricsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetProperties

`func (o *RevenueAnalyticsMetricsQuery) GetProperties() []RevenueAnalyticsPropertyFilter`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RevenueAnalyticsMetricsQuery) GetPropertiesOk() (*[]RevenueAnalyticsPropertyFilter, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RevenueAnalyticsMetricsQuery) SetProperties(v []RevenueAnalyticsPropertyFilter)`

SetProperties sets Properties field to given value.


### GetResponse

`func (o *RevenueAnalyticsMetricsQuery) GetResponse() RevenueAnalyticsMetricsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *RevenueAnalyticsMetricsQuery) GetResponseOk() (*RevenueAnalyticsMetricsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *RevenueAnalyticsMetricsQuery) SetResponse(v RevenueAnalyticsMetricsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *RevenueAnalyticsMetricsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *RevenueAnalyticsMetricsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RevenueAnalyticsMetricsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RevenueAnalyticsMetricsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RevenueAnalyticsMetricsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *RevenueAnalyticsMetricsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RevenueAnalyticsMetricsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RevenueAnalyticsMetricsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RevenueAnalyticsMetricsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *RevenueAnalyticsMetricsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *RevenueAnalyticsMetricsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


