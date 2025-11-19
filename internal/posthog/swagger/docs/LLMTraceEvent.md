# LLMTraceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** |  | 
**Event** | [**Event**](Event.md) |  | 
**Id** | **string** |  | 
**Properties** | **map[string]interface{}** |  | 

## Methods

### NewLLMTraceEvent

`func NewLLMTraceEvent(createdAt string, event Event, id string, properties map[string]interface{}, ) *LLMTraceEvent`

NewLLMTraceEvent instantiates a new LLMTraceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLLMTraceEventWithDefaults

`func NewLLMTraceEventWithDefaults() *LLMTraceEvent`

NewLLMTraceEventWithDefaults instantiates a new LLMTraceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LLMTraceEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LLMTraceEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LLMTraceEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEvent

`func (o *LLMTraceEvent) GetEvent() Event`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *LLMTraceEvent) GetEventOk() (*Event, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *LLMTraceEvent) SetEvent(v Event)`

SetEvent sets Event field to given value.


### GetId

`func (o *LLMTraceEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LLMTraceEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LLMTraceEvent) SetId(v string)`

SetId sets Id field to given value.


### GetProperties

`func (o *LLMTraceEvent) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *LLMTraceEvent) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *LLMTraceEvent) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


