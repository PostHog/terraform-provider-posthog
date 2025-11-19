# ResultCustomizationByPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentBy** | Pointer to **string** |  | [optional] [default to "position"]
**Color** | Pointer to [**DataColorToken**](DataColorToken.md) |  | [optional] 
**Hidden** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewResultCustomizationByPosition

`func NewResultCustomizationByPosition() *ResultCustomizationByPosition`

NewResultCustomizationByPosition instantiates a new ResultCustomizationByPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultCustomizationByPositionWithDefaults

`func NewResultCustomizationByPositionWithDefaults() *ResultCustomizationByPosition`

NewResultCustomizationByPositionWithDefaults instantiates a new ResultCustomizationByPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentBy

`func (o *ResultCustomizationByPosition) GetAssignmentBy() string`

GetAssignmentBy returns the AssignmentBy field if non-nil, zero value otherwise.

### GetAssignmentByOk

`func (o *ResultCustomizationByPosition) GetAssignmentByOk() (*string, bool)`

GetAssignmentByOk returns a tuple with the AssignmentBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentBy

`func (o *ResultCustomizationByPosition) SetAssignmentBy(v string)`

SetAssignmentBy sets AssignmentBy field to given value.

### HasAssignmentBy

`func (o *ResultCustomizationByPosition) HasAssignmentBy() bool`

HasAssignmentBy returns a boolean if a field has been set.

### GetColor

`func (o *ResultCustomizationByPosition) GetColor() DataColorToken`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ResultCustomizationByPosition) GetColorOk() (*DataColorToken, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ResultCustomizationByPosition) SetColor(v DataColorToken)`

SetColor sets Color field to given value.

### HasColor

`func (o *ResultCustomizationByPosition) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetHidden

`func (o *ResultCustomizationByPosition) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ResultCustomizationByPosition) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ResultCustomizationByPosition) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ResultCustomizationByPosition) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHiddenNil

`func (o *ResultCustomizationByPosition) SetHiddenNil(b bool)`

 SetHiddenNil sets the value for Hidden to be an explicit nil

### UnsetHidden
`func (o *ResultCustomizationByPosition) UnsetHidden()`

UnsetHidden ensures that no value is present for Hidden, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


