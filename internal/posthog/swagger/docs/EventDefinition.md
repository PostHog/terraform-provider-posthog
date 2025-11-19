# EventDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | **[]interface{}** |  | 
**Event** | **string** |  | 
**Properties** | **map[string]interface{}** |  | 

## Methods

### NewEventDefinition

`func NewEventDefinition(elements []interface{}, event string, properties map[string]interface{}, ) *EventDefinition`

NewEventDefinition instantiates a new EventDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDefinitionWithDefaults

`func NewEventDefinitionWithDefaults() *EventDefinition`

NewEventDefinitionWithDefaults instantiates a new EventDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *EventDefinition) GetElements() []interface{}`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *EventDefinition) GetElementsOk() (*[]interface{}, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *EventDefinition) SetElements(v []interface{})`

SetElements sets Elements field to given value.


### GetEvent

`func (o *EventDefinition) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventDefinition) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventDefinition) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetProperties

`func (o *EventDefinition) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventDefinition) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventDefinition) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


