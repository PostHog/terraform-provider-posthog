# EventTaxonomyQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionId** | Pointer to **NullableInt32** |  | [optional] 
**Event** | Pointer to **NullableString** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "EventTaxonomyQuery"]
**MaxPropertyValues** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Properties** | Pointer to **[]string** |  | [optional] 
**Response** | Pointer to [**EventTaxonomyQueryResponse**](EventTaxonomyQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewEventTaxonomyQuery

`func NewEventTaxonomyQuery() *EventTaxonomyQuery`

NewEventTaxonomyQuery instantiates a new EventTaxonomyQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTaxonomyQueryWithDefaults

`func NewEventTaxonomyQueryWithDefaults() *EventTaxonomyQuery`

NewEventTaxonomyQueryWithDefaults instantiates a new EventTaxonomyQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionId

`func (o *EventTaxonomyQuery) GetActionId() int32`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *EventTaxonomyQuery) GetActionIdOk() (*int32, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *EventTaxonomyQuery) SetActionId(v int32)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *EventTaxonomyQuery) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### SetActionIdNil

`func (o *EventTaxonomyQuery) SetActionIdNil(b bool)`

 SetActionIdNil sets the value for ActionId to be an explicit nil

### UnsetActionId
`func (o *EventTaxonomyQuery) UnsetActionId()`

UnsetActionId ensures that no value is present for ActionId, not even an explicit nil
### GetEvent

`func (o *EventTaxonomyQuery) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventTaxonomyQuery) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventTaxonomyQuery) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *EventTaxonomyQuery) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *EventTaxonomyQuery) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *EventTaxonomyQuery) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetKind

`func (o *EventTaxonomyQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EventTaxonomyQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EventTaxonomyQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EventTaxonomyQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMaxPropertyValues

`func (o *EventTaxonomyQuery) GetMaxPropertyValues() int32`

GetMaxPropertyValues returns the MaxPropertyValues field if non-nil, zero value otherwise.

### GetMaxPropertyValuesOk

`func (o *EventTaxonomyQuery) GetMaxPropertyValuesOk() (*int32, bool)`

GetMaxPropertyValuesOk returns a tuple with the MaxPropertyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPropertyValues

`func (o *EventTaxonomyQuery) SetMaxPropertyValues(v int32)`

SetMaxPropertyValues sets MaxPropertyValues field to given value.

### HasMaxPropertyValues

`func (o *EventTaxonomyQuery) HasMaxPropertyValues() bool`

HasMaxPropertyValues returns a boolean if a field has been set.

### SetMaxPropertyValuesNil

`func (o *EventTaxonomyQuery) SetMaxPropertyValuesNil(b bool)`

 SetMaxPropertyValuesNil sets the value for MaxPropertyValues to be an explicit nil

### UnsetMaxPropertyValues
`func (o *EventTaxonomyQuery) UnsetMaxPropertyValues()`

UnsetMaxPropertyValues ensures that no value is present for MaxPropertyValues, not even an explicit nil
### GetModifiers

`func (o *EventTaxonomyQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *EventTaxonomyQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *EventTaxonomyQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *EventTaxonomyQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetProperties

`func (o *EventTaxonomyQuery) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventTaxonomyQuery) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventTaxonomyQuery) SetProperties(v []string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EventTaxonomyQuery) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *EventTaxonomyQuery) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *EventTaxonomyQuery) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *EventTaxonomyQuery) GetResponse() EventTaxonomyQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *EventTaxonomyQuery) GetResponseOk() (*EventTaxonomyQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *EventTaxonomyQuery) SetResponse(v EventTaxonomyQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *EventTaxonomyQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *EventTaxonomyQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EventTaxonomyQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EventTaxonomyQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EventTaxonomyQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *EventTaxonomyQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EventTaxonomyQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EventTaxonomyQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EventTaxonomyQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *EventTaxonomyQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *EventTaxonomyQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


