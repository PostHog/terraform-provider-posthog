# ElementType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttrClass** | Pointer to **[]string** |  | [optional] 
**AttrId** | Pointer to **NullableString** |  | [optional] 
**Attributes** | **map[string]string** |  | 
**Href** | Pointer to **NullableString** |  | [optional] 
**NthChild** | Pointer to **NullableFloat32** |  | [optional] 
**NthOfType** | Pointer to **NullableFloat32** |  | [optional] 
**Order** | Pointer to **NullableFloat32** |  | [optional] 
**TagName** | **string** |  | 
**Text** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewElementType

`func NewElementType(attributes map[string]string, tagName string, ) *ElementType`

NewElementType instantiates a new ElementType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElementTypeWithDefaults

`func NewElementTypeWithDefaults() *ElementType`

NewElementTypeWithDefaults instantiates a new ElementType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttrClass

`func (o *ElementType) GetAttrClass() []string`

GetAttrClass returns the AttrClass field if non-nil, zero value otherwise.

### GetAttrClassOk

`func (o *ElementType) GetAttrClassOk() (*[]string, bool)`

GetAttrClassOk returns a tuple with the AttrClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrClass

`func (o *ElementType) SetAttrClass(v []string)`

SetAttrClass sets AttrClass field to given value.

### HasAttrClass

`func (o *ElementType) HasAttrClass() bool`

HasAttrClass returns a boolean if a field has been set.

### SetAttrClassNil

`func (o *ElementType) SetAttrClassNil(b bool)`

 SetAttrClassNil sets the value for AttrClass to be an explicit nil

### UnsetAttrClass
`func (o *ElementType) UnsetAttrClass()`

UnsetAttrClass ensures that no value is present for AttrClass, not even an explicit nil
### GetAttrId

`func (o *ElementType) GetAttrId() string`

GetAttrId returns the AttrId field if non-nil, zero value otherwise.

### GetAttrIdOk

`func (o *ElementType) GetAttrIdOk() (*string, bool)`

GetAttrIdOk returns a tuple with the AttrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrId

`func (o *ElementType) SetAttrId(v string)`

SetAttrId sets AttrId field to given value.

### HasAttrId

`func (o *ElementType) HasAttrId() bool`

HasAttrId returns a boolean if a field has been set.

### SetAttrIdNil

`func (o *ElementType) SetAttrIdNil(b bool)`

 SetAttrIdNil sets the value for AttrId to be an explicit nil

### UnsetAttrId
`func (o *ElementType) UnsetAttrId()`

UnsetAttrId ensures that no value is present for AttrId, not even an explicit nil
### GetAttributes

`func (o *ElementType) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ElementType) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ElementType) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.


### GetHref

`func (o *ElementType) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ElementType) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ElementType) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ElementType) HasHref() bool`

HasHref returns a boolean if a field has been set.

### SetHrefNil

`func (o *ElementType) SetHrefNil(b bool)`

 SetHrefNil sets the value for Href to be an explicit nil

### UnsetHref
`func (o *ElementType) UnsetHref()`

UnsetHref ensures that no value is present for Href, not even an explicit nil
### GetNthChild

`func (o *ElementType) GetNthChild() float32`

GetNthChild returns the NthChild field if non-nil, zero value otherwise.

### GetNthChildOk

`func (o *ElementType) GetNthChildOk() (*float32, bool)`

GetNthChildOk returns a tuple with the NthChild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNthChild

`func (o *ElementType) SetNthChild(v float32)`

SetNthChild sets NthChild field to given value.

### HasNthChild

`func (o *ElementType) HasNthChild() bool`

HasNthChild returns a boolean if a field has been set.

### SetNthChildNil

`func (o *ElementType) SetNthChildNil(b bool)`

 SetNthChildNil sets the value for NthChild to be an explicit nil

### UnsetNthChild
`func (o *ElementType) UnsetNthChild()`

UnsetNthChild ensures that no value is present for NthChild, not even an explicit nil
### GetNthOfType

`func (o *ElementType) GetNthOfType() float32`

GetNthOfType returns the NthOfType field if non-nil, zero value otherwise.

### GetNthOfTypeOk

`func (o *ElementType) GetNthOfTypeOk() (*float32, bool)`

GetNthOfTypeOk returns a tuple with the NthOfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNthOfType

`func (o *ElementType) SetNthOfType(v float32)`

SetNthOfType sets NthOfType field to given value.

### HasNthOfType

`func (o *ElementType) HasNthOfType() bool`

HasNthOfType returns a boolean if a field has been set.

### SetNthOfTypeNil

`func (o *ElementType) SetNthOfTypeNil(b bool)`

 SetNthOfTypeNil sets the value for NthOfType to be an explicit nil

### UnsetNthOfType
`func (o *ElementType) UnsetNthOfType()`

UnsetNthOfType ensures that no value is present for NthOfType, not even an explicit nil
### GetOrder

`func (o *ElementType) GetOrder() float32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ElementType) GetOrderOk() (*float32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ElementType) SetOrder(v float32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ElementType) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *ElementType) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *ElementType) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetTagName

`func (o *ElementType) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *ElementType) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *ElementType) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetText

`func (o *ElementType) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ElementType) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ElementType) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ElementType) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *ElementType) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *ElementType) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


