# ActorsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FixedProperties** | Pointer to [**[]FixedpropertiesInner1**](FixedpropertiesInner1.md) | Currently only person filters supported. No filters for querying groups. See &#x60;filter_conditions()&#x60; in actor_strategies.py. | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "ActorsQuery"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**OrderBy** | Pointer to **[]string** |  | [optional] 
**Properties** | Pointer to [**NullableProperties**](Properties.md) |  | [optional] [default to null]
**Response** | Pointer to [**ActorsQueryResponse**](ActorsQueryResponse.md) |  | [optional] 
**Search** | Pointer to **NullableString** |  | [optional] 
**Select** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to [**NullableSource**](Source.md) |  | [optional] [default to null]
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewActorsQuery

`func NewActorsQuery() *ActorsQuery`

NewActorsQuery instantiates a new ActorsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorsQueryWithDefaults

`func NewActorsQueryWithDefaults() *ActorsQuery`

NewActorsQueryWithDefaults instantiates a new ActorsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFixedProperties

`func (o *ActorsQuery) GetFixedProperties() []FixedpropertiesInner1`

GetFixedProperties returns the FixedProperties field if non-nil, zero value otherwise.

### GetFixedPropertiesOk

`func (o *ActorsQuery) GetFixedPropertiesOk() (*[]FixedpropertiesInner1, bool)`

GetFixedPropertiesOk returns a tuple with the FixedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedProperties

`func (o *ActorsQuery) SetFixedProperties(v []FixedpropertiesInner1)`

SetFixedProperties sets FixedProperties field to given value.

### HasFixedProperties

`func (o *ActorsQuery) HasFixedProperties() bool`

HasFixedProperties returns a boolean if a field has been set.

### SetFixedPropertiesNil

`func (o *ActorsQuery) SetFixedPropertiesNil(b bool)`

 SetFixedPropertiesNil sets the value for FixedProperties to be an explicit nil

### UnsetFixedProperties
`func (o *ActorsQuery) UnsetFixedProperties()`

UnsetFixedProperties ensures that no value is present for FixedProperties, not even an explicit nil
### GetKind

`func (o *ActorsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ActorsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ActorsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ActorsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *ActorsQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ActorsQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ActorsQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ActorsQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ActorsQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ActorsQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *ActorsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *ActorsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *ActorsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *ActorsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *ActorsQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ActorsQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ActorsQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ActorsQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *ActorsQuery) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *ActorsQuery) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetOrderBy

`func (o *ActorsQuery) GetOrderBy() []string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ActorsQuery) GetOrderByOk() (*[]string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ActorsQuery) SetOrderBy(v []string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ActorsQuery) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *ActorsQuery) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *ActorsQuery) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetProperties

`func (o *ActorsQuery) GetProperties() Properties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ActorsQuery) GetPropertiesOk() (*Properties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ActorsQuery) SetProperties(v Properties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ActorsQuery) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ActorsQuery) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ActorsQuery) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *ActorsQuery) GetResponse() ActorsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ActorsQuery) GetResponseOk() (*ActorsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ActorsQuery) SetResponse(v ActorsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ActorsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSearch

`func (o *ActorsQuery) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *ActorsQuery) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *ActorsQuery) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *ActorsQuery) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *ActorsQuery) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *ActorsQuery) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetSelect

`func (o *ActorsQuery) GetSelect() []string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *ActorsQuery) GetSelectOk() (*[]string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *ActorsQuery) SetSelect(v []string)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *ActorsQuery) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### SetSelectNil

`func (o *ActorsQuery) SetSelectNil(b bool)`

 SetSelectNil sets the value for Select to be an explicit nil

### UnsetSelect
`func (o *ActorsQuery) UnsetSelect()`

UnsetSelect ensures that no value is present for Select, not even an explicit nil
### GetSource

`func (o *ActorsQuery) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ActorsQuery) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ActorsQuery) SetSource(v Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *ActorsQuery) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *ActorsQuery) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *ActorsQuery) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetTags

`func (o *ActorsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ActorsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ActorsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ActorsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *ActorsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ActorsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ActorsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ActorsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ActorsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ActorsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


