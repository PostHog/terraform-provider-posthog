# ClickhouseEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**DistinctId** | **string** |  | [readonly] 
**Properties** | **string** |  | [readonly] 
**Event** | **string** |  | [readonly] 
**Timestamp** | **string** |  | [readonly] 
**Person** | **string** |  | [readonly] 
**Elements** | **string** |  | [readonly] 
**ElementsChain** | **string** |  | [readonly] 

## Methods

### NewClickhouseEvent

`func NewClickhouseEvent(id string, distinctId string, properties string, event string, timestamp string, person string, elements string, elementsChain string, ) *ClickhouseEvent`

NewClickhouseEvent instantiates a new ClickhouseEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickhouseEventWithDefaults

`func NewClickhouseEventWithDefaults() *ClickhouseEvent`

NewClickhouseEventWithDefaults instantiates a new ClickhouseEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClickhouseEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClickhouseEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClickhouseEvent) SetId(v string)`

SetId sets Id field to given value.


### GetDistinctId

`func (o *ClickhouseEvent) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *ClickhouseEvent) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *ClickhouseEvent) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.


### GetProperties

`func (o *ClickhouseEvent) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ClickhouseEvent) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ClickhouseEvent) SetProperties(v string)`

SetProperties sets Properties field to given value.


### GetEvent

`func (o *ClickhouseEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ClickhouseEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ClickhouseEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetTimestamp

`func (o *ClickhouseEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ClickhouseEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ClickhouseEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetPerson

`func (o *ClickhouseEvent) GetPerson() string`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *ClickhouseEvent) GetPersonOk() (*string, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *ClickhouseEvent) SetPerson(v string)`

SetPerson sets Person field to given value.


### GetElements

`func (o *ClickhouseEvent) GetElements() string`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *ClickhouseEvent) GetElementsOk() (*string, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *ClickhouseEvent) SetElements(v string)`

SetElements sets Elements field to given value.


### GetElementsChain

`func (o *ClickhouseEvent) GetElementsChain() string`

GetElementsChain returns the ElementsChain field if non-nil, zero value otherwise.

### GetElementsChainOk

`func (o *ClickhouseEvent) GetElementsChainOk() (*string, bool)`

GetElementsChainOk returns a tuple with the ElementsChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementsChain

`func (o *ClickhouseEvent) SetElementsChain(v string)`

SetElementsChain sets ElementsChain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


