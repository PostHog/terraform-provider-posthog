# EventMetadataPropertyFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**Operator** | [**PropertyOperator**](PropertyOperator.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "event_metadata"]
**Value** | Pointer to [**NullableValue2**](Value2.md) |  | [optional] [default to null]

## Methods

### NewEventMetadataPropertyFilter

`func NewEventMetadataPropertyFilter(key string, operator PropertyOperator, ) *EventMetadataPropertyFilter`

NewEventMetadataPropertyFilter instantiates a new EventMetadataPropertyFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventMetadataPropertyFilterWithDefaults

`func NewEventMetadataPropertyFilterWithDefaults() *EventMetadataPropertyFilter`

NewEventMetadataPropertyFilterWithDefaults instantiates a new EventMetadataPropertyFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *EventMetadataPropertyFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EventMetadataPropertyFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EventMetadataPropertyFilter) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *EventMetadataPropertyFilter) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EventMetadataPropertyFilter) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EventMetadataPropertyFilter) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *EventMetadataPropertyFilter) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *EventMetadataPropertyFilter) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *EventMetadataPropertyFilter) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetOperator

`func (o *EventMetadataPropertyFilter) GetOperator() PropertyOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *EventMetadataPropertyFilter) GetOperatorOk() (*PropertyOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *EventMetadataPropertyFilter) SetOperator(v PropertyOperator)`

SetOperator sets Operator field to given value.


### GetType

`func (o *EventMetadataPropertyFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventMetadataPropertyFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventMetadataPropertyFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventMetadataPropertyFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *EventMetadataPropertyFilter) GetValue() Value2`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EventMetadataPropertyFilter) GetValueOk() (*Value2, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EventMetadataPropertyFilter) SetValue(v Value2)`

SetValue sets Value field to given value.

### HasValue

`func (o *EventMetadataPropertyFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *EventMetadataPropertyFilter) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *EventMetadataPropertyFilter) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


