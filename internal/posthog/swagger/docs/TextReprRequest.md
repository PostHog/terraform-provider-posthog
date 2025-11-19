# TextReprRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**EventTypeEnum**](EventTypeEnum.md) | Type of LLM event to stringify  * &#x60;$ai_generation&#x60; - $ai_generation * &#x60;$ai_span&#x60; - $ai_span * &#x60;$ai_embedding&#x60; - $ai_embedding * &#x60;$ai_trace&#x60; - $ai_trace | 
**Data** | **interface{}** | Event data to stringify. For traces, should include &#39;trace&#39; and &#39;hierarchy&#39; fields. | 
**Options** | Pointer to [**TextReprOptions**](TextReprOptions.md) | Optional configuration for text generation | [optional] 

## Methods

### NewTextReprRequest

`func NewTextReprRequest(eventType EventTypeEnum, data interface{}, ) *TextReprRequest`

NewTextReprRequest instantiates a new TextReprRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextReprRequestWithDefaults

`func NewTextReprRequestWithDefaults() *TextReprRequest`

NewTextReprRequestWithDefaults instantiates a new TextReprRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *TextReprRequest) GetEventType() EventTypeEnum`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TextReprRequest) GetEventTypeOk() (*EventTypeEnum, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TextReprRequest) SetEventType(v EventTypeEnum)`

SetEventType sets EventType field to given value.


### GetData

`func (o *TextReprRequest) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TextReprRequest) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TextReprRequest) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *TextReprRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *TextReprRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetOptions

`func (o *TextReprRequest) GetOptions() TextReprOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *TextReprRequest) GetOptionsOk() (*TextReprOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *TextReprRequest) SetOptions(v TextReprOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *TextReprRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


