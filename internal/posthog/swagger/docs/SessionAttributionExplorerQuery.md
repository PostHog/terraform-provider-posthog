# SessionAttributionExplorerQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**Filters**](Filters.md) |  | [optional] 
**GroupBy** | [**[]SessionAttributionGroupBy**](SessionAttributionGroupBy.md) |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "SessionAttributionExplorerQuery"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**Response** | Pointer to [**SessionAttributionExplorerQueryResponse**](SessionAttributionExplorerQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewSessionAttributionExplorerQuery

`func NewSessionAttributionExplorerQuery(groupBy []SessionAttributionGroupBy, ) *SessionAttributionExplorerQuery`

NewSessionAttributionExplorerQuery instantiates a new SessionAttributionExplorerQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionAttributionExplorerQueryWithDefaults

`func NewSessionAttributionExplorerQueryWithDefaults() *SessionAttributionExplorerQuery`

NewSessionAttributionExplorerQueryWithDefaults instantiates a new SessionAttributionExplorerQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *SessionAttributionExplorerQuery) GetFilters() Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SessionAttributionExplorerQuery) GetFiltersOk() (*Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SessionAttributionExplorerQuery) SetFilters(v Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SessionAttributionExplorerQuery) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGroupBy

`func (o *SessionAttributionExplorerQuery) GetGroupBy() []SessionAttributionGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *SessionAttributionExplorerQuery) GetGroupByOk() (*[]SessionAttributionGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *SessionAttributionExplorerQuery) SetGroupBy(v []SessionAttributionGroupBy)`

SetGroupBy sets GroupBy field to given value.


### GetKind

`func (o *SessionAttributionExplorerQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SessionAttributionExplorerQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SessionAttributionExplorerQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SessionAttributionExplorerQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *SessionAttributionExplorerQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SessionAttributionExplorerQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SessionAttributionExplorerQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SessionAttributionExplorerQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *SessionAttributionExplorerQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *SessionAttributionExplorerQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *SessionAttributionExplorerQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *SessionAttributionExplorerQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *SessionAttributionExplorerQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *SessionAttributionExplorerQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *SessionAttributionExplorerQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SessionAttributionExplorerQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SessionAttributionExplorerQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SessionAttributionExplorerQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *SessionAttributionExplorerQuery) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *SessionAttributionExplorerQuery) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetResponse

`func (o *SessionAttributionExplorerQuery) GetResponse() SessionAttributionExplorerQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *SessionAttributionExplorerQuery) GetResponseOk() (*SessionAttributionExplorerQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *SessionAttributionExplorerQuery) SetResponse(v SessionAttributionExplorerQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *SessionAttributionExplorerQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *SessionAttributionExplorerQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SessionAttributionExplorerQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SessionAttributionExplorerQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SessionAttributionExplorerQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *SessionAttributionExplorerQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SessionAttributionExplorerQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SessionAttributionExplorerQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SessionAttributionExplorerQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *SessionAttributionExplorerQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *SessionAttributionExplorerQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


