# EventType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinctId** | **string** |  | 
**Elements** | [**[]ElementType**](ElementType.md) |  | 
**ElementsChain** | Pointer to **NullableString** |  | [optional] 
**Event** | **string** |  | 
**Id** | **string** |  | 
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 
**PersonId** | Pointer to **NullableString** |  | [optional] 
**PersonMode** | Pointer to **NullableString** |  | [optional] 
**Properties** | **map[string]interface{}** |  | 
**Timestamp** | **string** |  | 
**Uuid** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEventType

`func NewEventType(distinctId string, elements []ElementType, event string, id string, properties map[string]interface{}, timestamp string, ) *EventType`

NewEventType instantiates a new EventType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTypeWithDefaults

`func NewEventTypeWithDefaults() *EventType`

NewEventTypeWithDefaults instantiates a new EventType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctId

`func (o *EventType) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *EventType) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *EventType) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.


### GetElements

`func (o *EventType) GetElements() []ElementType`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *EventType) GetElementsOk() (*[]ElementType, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *EventType) SetElements(v []ElementType)`

SetElements sets Elements field to given value.


### GetElementsChain

`func (o *EventType) GetElementsChain() string`

GetElementsChain returns the ElementsChain field if non-nil, zero value otherwise.

### GetElementsChainOk

`func (o *EventType) GetElementsChainOk() (*string, bool)`

GetElementsChainOk returns a tuple with the ElementsChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementsChain

`func (o *EventType) SetElementsChain(v string)`

SetElementsChain sets ElementsChain field to given value.

### HasElementsChain

`func (o *EventType) HasElementsChain() bool`

HasElementsChain returns a boolean if a field has been set.

### SetElementsChainNil

`func (o *EventType) SetElementsChainNil(b bool)`

 SetElementsChainNil sets the value for ElementsChain to be an explicit nil

### UnsetElementsChain
`func (o *EventType) UnsetElementsChain()`

UnsetElementsChain ensures that no value is present for ElementsChain, not even an explicit nil
### GetEvent

`func (o *EventType) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventType) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventType) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetId

`func (o *EventType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventType) SetId(v string)`

SetId sets Id field to given value.


### GetPerson

`func (o *EventType) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *EventType) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *EventType) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *EventType) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetPersonId

`func (o *EventType) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *EventType) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *EventType) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *EventType) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### SetPersonIdNil

`func (o *EventType) SetPersonIdNil(b bool)`

 SetPersonIdNil sets the value for PersonId to be an explicit nil

### UnsetPersonId
`func (o *EventType) UnsetPersonId()`

UnsetPersonId ensures that no value is present for PersonId, not even an explicit nil
### GetPersonMode

`func (o *EventType) GetPersonMode() string`

GetPersonMode returns the PersonMode field if non-nil, zero value otherwise.

### GetPersonModeOk

`func (o *EventType) GetPersonModeOk() (*string, bool)`

GetPersonModeOk returns a tuple with the PersonMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonMode

`func (o *EventType) SetPersonMode(v string)`

SetPersonMode sets PersonMode field to given value.

### HasPersonMode

`func (o *EventType) HasPersonMode() bool`

HasPersonMode returns a boolean if a field has been set.

### SetPersonModeNil

`func (o *EventType) SetPersonModeNil(b bool)`

 SetPersonModeNil sets the value for PersonMode to be an explicit nil

### UnsetPersonMode
`func (o *EventType) UnsetPersonMode()`

UnsetPersonMode ensures that no value is present for PersonMode, not even an explicit nil
### GetProperties

`func (o *EventType) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventType) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventType) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.


### GetTimestamp

`func (o *EventType) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventType) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventType) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetUuid

`func (o *EventType) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EventType) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EventType) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EventType) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EventType) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EventType) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


