# RetentionValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRetentionValue

`func NewRetentionValue(count int32, ) *RetentionValue`

NewRetentionValue instantiates a new RetentionValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionValueWithDefaults

`func NewRetentionValueWithDefaults() *RetentionValue`

NewRetentionValueWithDefaults instantiates a new RetentionValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RetentionValue) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RetentionValue) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RetentionValue) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLabel

`func (o *RetentionValue) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RetentionValue) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RetentionValue) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RetentionValue) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *RetentionValue) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *RetentionValue) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


