# EventsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionId** | Pointer to **NullableInt32** | Show events matching a given action | [optional] 
**After** | Pointer to **NullableString** | Only fetch events that happened after this timestamp | [optional] 
**Before** | Pointer to **NullableString** | Only fetch events that happened before this timestamp | [optional] 
**Event** | Pointer to **NullableString** | Limit to events matching this string | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** | Filter test accounts | [optional] 
**FixedProperties** | Pointer to [**[]FixedpropertiesInner2**](FixedpropertiesInner2.md) | Fixed properties in the query, can&#39;t be edited in the interface (e.g. scoping down by person) | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "EventsQuery"]
**Limit** | Pointer to **NullableInt32** | Number of rows to return | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** | Number of rows to skip before returning rows | [optional] 
**OrderBy** | Pointer to **[]string** | Columns to order by | [optional] 
**PersonId** | Pointer to **NullableString** | Show events for a given person | [optional] 
**Properties** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) | Properties configurable in the interface | [optional] 
**Response** | Pointer to [**EventsQueryResponse**](EventsQueryResponse.md) |  | [optional] 
**Select** | **[]string** | Return a limited set of data. Required. | 
**Source** | Pointer to [**InsightActorsQuery**](InsightActorsQuery.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 
**Where** | Pointer to **[]string** | HogQL filters to apply on returned data | [optional] 

## Methods

### NewEventsQuery

`func NewEventsQuery(select_ []string, ) *EventsQuery`

NewEventsQuery instantiates a new EventsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsQueryWithDefaults

`func NewEventsQueryWithDefaults() *EventsQuery`

NewEventsQueryWithDefaults instantiates a new EventsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionId

`func (o *EventsQuery) GetActionId() int32`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *EventsQuery) GetActionIdOk() (*int32, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *EventsQuery) SetActionId(v int32)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *EventsQuery) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### SetActionIdNil

`func (o *EventsQuery) SetActionIdNil(b bool)`

 SetActionIdNil sets the value for ActionId to be an explicit nil

### UnsetActionId
`func (o *EventsQuery) UnsetActionId()`

UnsetActionId ensures that no value is present for ActionId, not even an explicit nil
### GetAfter

`func (o *EventsQuery) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *EventsQuery) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *EventsQuery) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *EventsQuery) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *EventsQuery) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *EventsQuery) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetBefore

`func (o *EventsQuery) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *EventsQuery) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *EventsQuery) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *EventsQuery) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### SetBeforeNil

`func (o *EventsQuery) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *EventsQuery) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil
### GetEvent

`func (o *EventsQuery) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventsQuery) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventsQuery) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *EventsQuery) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *EventsQuery) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *EventsQuery) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetFilterTestAccounts

`func (o *EventsQuery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *EventsQuery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *EventsQuery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *EventsQuery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *EventsQuery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *EventsQuery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetFixedProperties

`func (o *EventsQuery) GetFixedProperties() []FixedpropertiesInner2`

GetFixedProperties returns the FixedProperties field if non-nil, zero value otherwise.

### GetFixedPropertiesOk

`func (o *EventsQuery) GetFixedPropertiesOk() (*[]FixedpropertiesInner2, bool)`

GetFixedPropertiesOk returns a tuple with the FixedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedProperties

`func (o *EventsQuery) SetFixedProperties(v []FixedpropertiesInner2)`

SetFixedProperties sets FixedProperties field to given value.

### HasFixedProperties

`func (o *EventsQuery) HasFixedProperties() bool`

HasFixedProperties returns a boolean if a field has been set.

### SetFixedPropertiesNil

`func (o *EventsQuery) SetFixedPropertiesNil(b bool)`

 SetFixedPropertiesNil sets the value for FixedProperties to be an explicit nil

### UnsetFixedProperties
`func (o *EventsQuery) UnsetFixedProperties()`

UnsetFixedProperties ensures that no value is present for FixedProperties, not even an explicit nil
### GetKind

`func (o *EventsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EventsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EventsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EventsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *EventsQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *EventsQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *EventsQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *EventsQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *EventsQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *EventsQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *EventsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *EventsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *EventsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *EventsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *EventsQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *EventsQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *EventsQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *EventsQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *EventsQuery) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *EventsQuery) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetOrderBy

`func (o *EventsQuery) GetOrderBy() []string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *EventsQuery) GetOrderByOk() (*[]string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *EventsQuery) SetOrderBy(v []string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *EventsQuery) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *EventsQuery) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *EventsQuery) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetPersonId

`func (o *EventsQuery) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *EventsQuery) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *EventsQuery) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *EventsQuery) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### SetPersonIdNil

`func (o *EventsQuery) SetPersonIdNil(b bool)`

 SetPersonIdNil sets the value for PersonId to be an explicit nil

### UnsetPersonId
`func (o *EventsQuery) UnsetPersonId()`

UnsetPersonId ensures that no value is present for PersonId, not even an explicit nil
### GetProperties

`func (o *EventsQuery) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventsQuery) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventsQuery) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EventsQuery) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *EventsQuery) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *EventsQuery) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *EventsQuery) GetResponse() EventsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *EventsQuery) GetResponseOk() (*EventsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *EventsQuery) SetResponse(v EventsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *EventsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSelect

`func (o *EventsQuery) GetSelect() []string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *EventsQuery) GetSelectOk() (*[]string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *EventsQuery) SetSelect(v []string)`

SetSelect sets Select field to given value.


### GetSource

`func (o *EventsQuery) GetSource() InsightActorsQuery`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventsQuery) GetSourceOk() (*InsightActorsQuery, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventsQuery) SetSource(v InsightActorsQuery)`

SetSource sets Source field to given value.

### HasSource

`func (o *EventsQuery) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTags

`func (o *EventsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EventsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EventsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EventsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *EventsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EventsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EventsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EventsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *EventsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *EventsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetWhere

`func (o *EventsQuery) GetWhere() []string`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *EventsQuery) GetWhereOk() (*[]string, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *EventsQuery) SetWhere(v []string)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *EventsQuery) HasWhere() bool`

HasWhere returns a boolean if a field has been set.

### SetWhereNil

`func (o *EventsQuery) SetWhereNil(b bool)`

 SetWhereNil sets the value for Where to be an explicit nil

### UnsetWhere
`func (o *EventsQuery) UnsetWhere()`

UnsetWhere ensures that no value is present for Where, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


