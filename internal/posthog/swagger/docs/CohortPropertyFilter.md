# CohortPropertyFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CohortName** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] [default to "id"]
**Label** | Pointer to **NullableString** |  | [optional] 
**Operator** | Pointer to [**PropertyOperator**](PropertyOperator.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "cohort"]
**Value** | **int32** |  | 

## Methods

### NewCohortPropertyFilter

`func NewCohortPropertyFilter(value int32, ) *CohortPropertyFilter`

NewCohortPropertyFilter instantiates a new CohortPropertyFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCohortPropertyFilterWithDefaults

`func NewCohortPropertyFilterWithDefaults() *CohortPropertyFilter`

NewCohortPropertyFilterWithDefaults instantiates a new CohortPropertyFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCohortName

`func (o *CohortPropertyFilter) GetCohortName() string`

GetCohortName returns the CohortName field if non-nil, zero value otherwise.

### GetCohortNameOk

`func (o *CohortPropertyFilter) GetCohortNameOk() (*string, bool)`

GetCohortNameOk returns a tuple with the CohortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortName

`func (o *CohortPropertyFilter) SetCohortName(v string)`

SetCohortName sets CohortName field to given value.

### HasCohortName

`func (o *CohortPropertyFilter) HasCohortName() bool`

HasCohortName returns a boolean if a field has been set.

### SetCohortNameNil

`func (o *CohortPropertyFilter) SetCohortNameNil(b bool)`

 SetCohortNameNil sets the value for CohortName to be an explicit nil

### UnsetCohortName
`func (o *CohortPropertyFilter) UnsetCohortName()`

UnsetCohortName ensures that no value is present for CohortName, not even an explicit nil
### GetKey

`func (o *CohortPropertyFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CohortPropertyFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CohortPropertyFilter) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CohortPropertyFilter) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *CohortPropertyFilter) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CohortPropertyFilter) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CohortPropertyFilter) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CohortPropertyFilter) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *CohortPropertyFilter) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *CohortPropertyFilter) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetOperator

`func (o *CohortPropertyFilter) GetOperator() PropertyOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CohortPropertyFilter) GetOperatorOk() (*PropertyOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CohortPropertyFilter) SetOperator(v PropertyOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CohortPropertyFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetType

`func (o *CohortPropertyFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CohortPropertyFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CohortPropertyFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CohortPropertyFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *CohortPropertyFilter) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CohortPropertyFilter) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CohortPropertyFilter) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


