# TrendsFormulaNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomName** | Pointer to **NullableString** | Optional user-defined name for the formula | [optional] 
**Formula** | **string** |  | 

## Methods

### NewTrendsFormulaNode

`func NewTrendsFormulaNode(formula string, ) *TrendsFormulaNode`

NewTrendsFormulaNode instantiates a new TrendsFormulaNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendsFormulaNodeWithDefaults

`func NewTrendsFormulaNodeWithDefaults() *TrendsFormulaNode`

NewTrendsFormulaNodeWithDefaults instantiates a new TrendsFormulaNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomName

`func (o *TrendsFormulaNode) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *TrendsFormulaNode) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *TrendsFormulaNode) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *TrendsFormulaNode) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *TrendsFormulaNode) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *TrendsFormulaNode) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetFormula

`func (o *TrendsFormulaNode) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *TrendsFormulaNode) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *TrendsFormulaNode) SetFormula(v string)`

SetFormula sets Formula field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


