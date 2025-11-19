# ResultCustomizationByValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentBy** | Pointer to **string** |  | [optional] [default to "value"]
**Color** | Pointer to [**DataColorToken**](DataColorToken.md) |  | [optional] 
**Hidden** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewResultCustomizationByValue

`func NewResultCustomizationByValue() *ResultCustomizationByValue`

NewResultCustomizationByValue instantiates a new ResultCustomizationByValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultCustomizationByValueWithDefaults

`func NewResultCustomizationByValueWithDefaults() *ResultCustomizationByValue`

NewResultCustomizationByValueWithDefaults instantiates a new ResultCustomizationByValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentBy

`func (o *ResultCustomizationByValue) GetAssignmentBy() string`

GetAssignmentBy returns the AssignmentBy field if non-nil, zero value otherwise.

### GetAssignmentByOk

`func (o *ResultCustomizationByValue) GetAssignmentByOk() (*string, bool)`

GetAssignmentByOk returns a tuple with the AssignmentBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentBy

`func (o *ResultCustomizationByValue) SetAssignmentBy(v string)`

SetAssignmentBy sets AssignmentBy field to given value.

### HasAssignmentBy

`func (o *ResultCustomizationByValue) HasAssignmentBy() bool`

HasAssignmentBy returns a boolean if a field has been set.

### GetColor

`func (o *ResultCustomizationByValue) GetColor() DataColorToken`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ResultCustomizationByValue) GetColorOk() (*DataColorToken, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ResultCustomizationByValue) SetColor(v DataColorToken)`

SetColor sets Color field to given value.

### HasColor

`func (o *ResultCustomizationByValue) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetHidden

`func (o *ResultCustomizationByValue) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ResultCustomizationByValue) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ResultCustomizationByValue) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ResultCustomizationByValue) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHiddenNil

`func (o *ResultCustomizationByValue) SetHiddenNil(b bool)`

 SetHiddenNil sets the value for Hidden to be an explicit nil

### UnsetHidden
`func (o *ResultCustomizationByValue) UnsetHidden()`

UnsetHidden ensures that no value is present for Hidden, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


