# PropertyDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** |  | 
**IsNumerical** | Pointer to **bool** |  | [optional] 
**PropertyType** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]interface{}** |  | [optional] 
**IsSeenOnFilteredEvents** | **string** |  | [readonly] 

## Methods

### NewPropertyDefinition

`func NewPropertyDefinition(id string, name string, isSeenOnFilteredEvents string, ) *PropertyDefinition`

NewPropertyDefinition instantiates a new PropertyDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyDefinitionWithDefaults

`func NewPropertyDefinitionWithDefaults() *PropertyDefinition`

NewPropertyDefinitionWithDefaults instantiates a new PropertyDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PropertyDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PropertyDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PropertyDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PropertyDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropertyDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropertyDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetIsNumerical

`func (o *PropertyDefinition) GetIsNumerical() bool`

GetIsNumerical returns the IsNumerical field if non-nil, zero value otherwise.

### GetIsNumericalOk

`func (o *PropertyDefinition) GetIsNumericalOk() (*bool, bool)`

GetIsNumericalOk returns a tuple with the IsNumerical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumerical

`func (o *PropertyDefinition) SetIsNumerical(v bool)`

SetIsNumerical sets IsNumerical field to given value.

### HasIsNumerical

`func (o *PropertyDefinition) HasIsNumerical() bool`

HasIsNumerical returns a boolean if a field has been set.

### GetPropertyType

`func (o *PropertyDefinition) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *PropertyDefinition) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *PropertyDefinition) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.

### HasPropertyType

`func (o *PropertyDefinition) HasPropertyType() bool`

HasPropertyType returns a boolean if a field has been set.

### SetPropertyTypeNil

`func (o *PropertyDefinition) SetPropertyTypeNil(b bool)`

 SetPropertyTypeNil sets the value for PropertyType to be an explicit nil

### UnsetPropertyType
`func (o *PropertyDefinition) UnsetPropertyType()`

UnsetPropertyType ensures that no value is present for PropertyType, not even an explicit nil
### GetTags

`func (o *PropertyDefinition) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PropertyDefinition) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PropertyDefinition) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *PropertyDefinition) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsSeenOnFilteredEvents

`func (o *PropertyDefinition) GetIsSeenOnFilteredEvents() string`

GetIsSeenOnFilteredEvents returns the IsSeenOnFilteredEvents field if non-nil, zero value otherwise.

### GetIsSeenOnFilteredEventsOk

`func (o *PropertyDefinition) GetIsSeenOnFilteredEventsOk() (*string, bool)`

GetIsSeenOnFilteredEventsOk returns a tuple with the IsSeenOnFilteredEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeenOnFilteredEvents

`func (o *PropertyDefinition) SetIsSeenOnFilteredEvents(v string)`

SetIsSeenOnFilteredEvents sets IsSeenOnFilteredEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


