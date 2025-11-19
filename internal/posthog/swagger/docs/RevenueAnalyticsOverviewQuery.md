# RevenueAnalyticsOverviewQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "RevenueAnalyticsOverviewQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Properties** | [**[]RevenueAnalyticsPropertyFilter**](RevenueAnalyticsPropertyFilter.md) |  | 
**Response** | Pointer to [**RevenueAnalyticsOverviewQueryResponse**](RevenueAnalyticsOverviewQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewRevenueAnalyticsOverviewQuery

`func NewRevenueAnalyticsOverviewQuery(properties []RevenueAnalyticsPropertyFilter, ) *RevenueAnalyticsOverviewQuery`

NewRevenueAnalyticsOverviewQuery instantiates a new RevenueAnalyticsOverviewQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueAnalyticsOverviewQueryWithDefaults

`func NewRevenueAnalyticsOverviewQueryWithDefaults() *RevenueAnalyticsOverviewQuery`

NewRevenueAnalyticsOverviewQueryWithDefaults instantiates a new RevenueAnalyticsOverviewQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateRange

`func (o *RevenueAnalyticsOverviewQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *RevenueAnalyticsOverviewQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *RevenueAnalyticsOverviewQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *RevenueAnalyticsOverviewQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetKind

`func (o *RevenueAnalyticsOverviewQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RevenueAnalyticsOverviewQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RevenueAnalyticsOverviewQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RevenueAnalyticsOverviewQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *RevenueAnalyticsOverviewQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *RevenueAnalyticsOverviewQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *RevenueAnalyticsOverviewQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *RevenueAnalyticsOverviewQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetProperties

`func (o *RevenueAnalyticsOverviewQuery) GetProperties() []RevenueAnalyticsPropertyFilter`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RevenueAnalyticsOverviewQuery) GetPropertiesOk() (*[]RevenueAnalyticsPropertyFilter, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RevenueAnalyticsOverviewQuery) SetProperties(v []RevenueAnalyticsPropertyFilter)`

SetProperties sets Properties field to given value.


### GetResponse

`func (o *RevenueAnalyticsOverviewQuery) GetResponse() RevenueAnalyticsOverviewQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *RevenueAnalyticsOverviewQuery) GetResponseOk() (*RevenueAnalyticsOverviewQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *RevenueAnalyticsOverviewQuery) SetResponse(v RevenueAnalyticsOverviewQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *RevenueAnalyticsOverviewQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *RevenueAnalyticsOverviewQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RevenueAnalyticsOverviewQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RevenueAnalyticsOverviewQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RevenueAnalyticsOverviewQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *RevenueAnalyticsOverviewQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RevenueAnalyticsOverviewQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RevenueAnalyticsOverviewQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RevenueAnalyticsOverviewQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *RevenueAnalyticsOverviewQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *RevenueAnalyticsOverviewQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


