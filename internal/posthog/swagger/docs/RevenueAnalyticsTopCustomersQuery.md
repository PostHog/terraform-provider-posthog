# RevenueAnalyticsTopCustomersQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**GroupBy** | [**RevenueAnalyticsTopCustomersGroupBy**](RevenueAnalyticsTopCustomersGroupBy.md) |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "RevenueAnalyticsTopCustomersQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Properties** | [**[]RevenueAnalyticsPropertyFilter**](RevenueAnalyticsPropertyFilter.md) |  | 
**Response** | Pointer to [**RevenueAnalyticsTopCustomersQueryResponse**](RevenueAnalyticsTopCustomersQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewRevenueAnalyticsTopCustomersQuery

`func NewRevenueAnalyticsTopCustomersQuery(groupBy RevenueAnalyticsTopCustomersGroupBy, properties []RevenueAnalyticsPropertyFilter, ) *RevenueAnalyticsTopCustomersQuery`

NewRevenueAnalyticsTopCustomersQuery instantiates a new RevenueAnalyticsTopCustomersQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueAnalyticsTopCustomersQueryWithDefaults

`func NewRevenueAnalyticsTopCustomersQueryWithDefaults() *RevenueAnalyticsTopCustomersQuery`

NewRevenueAnalyticsTopCustomersQueryWithDefaults instantiates a new RevenueAnalyticsTopCustomersQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRange

`func (o *RevenueAnalyticsTopCustomersQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *RevenueAnalyticsTopCustomersQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *RevenueAnalyticsTopCustomersQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *RevenueAnalyticsTopCustomersQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetGroupBy

`func (o *RevenueAnalyticsTopCustomersQuery) GetGroupBy() RevenueAnalyticsTopCustomersGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *RevenueAnalyticsTopCustomersQuery) GetGroupByOk() (*RevenueAnalyticsTopCustomersGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *RevenueAnalyticsTopCustomersQuery) SetGroupBy(v RevenueAnalyticsTopCustomersGroupBy)`

SetGroupBy sets GroupBy field to given value.


### GetKind

`func (o *RevenueAnalyticsTopCustomersQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RevenueAnalyticsTopCustomersQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RevenueAnalyticsTopCustomersQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RevenueAnalyticsTopCustomersQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *RevenueAnalyticsTopCustomersQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *RevenueAnalyticsTopCustomersQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *RevenueAnalyticsTopCustomersQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *RevenueAnalyticsTopCustomersQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetProperties

`func (o *RevenueAnalyticsTopCustomersQuery) GetProperties() []RevenueAnalyticsPropertyFilter`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RevenueAnalyticsTopCustomersQuery) GetPropertiesOk() (*[]RevenueAnalyticsPropertyFilter, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RevenueAnalyticsTopCustomersQuery) SetProperties(v []RevenueAnalyticsPropertyFilter)`

SetProperties sets Properties field to given value.


### GetResponse

`func (o *RevenueAnalyticsTopCustomersQuery) GetResponse() RevenueAnalyticsTopCustomersQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *RevenueAnalyticsTopCustomersQuery) GetResponseOk() (*RevenueAnalyticsTopCustomersQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *RevenueAnalyticsTopCustomersQuery) SetResponse(v RevenueAnalyticsTopCustomersQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *RevenueAnalyticsTopCustomersQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *RevenueAnalyticsTopCustomersQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RevenueAnalyticsTopCustomersQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RevenueAnalyticsTopCustomersQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RevenueAnalyticsTopCustomersQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *RevenueAnalyticsTopCustomersQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RevenueAnalyticsTopCustomersQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RevenueAnalyticsTopCustomersQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RevenueAnalyticsTopCustomersQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *RevenueAnalyticsTopCustomersQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *RevenueAnalyticsTopCustomersQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


