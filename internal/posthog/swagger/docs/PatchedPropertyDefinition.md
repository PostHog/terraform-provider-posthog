# PatchedPropertyDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**IsNumerical** | Pointer to **bool** |  | [optional] 
**PropertyType** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]interface{}** |  | [optional] 
**IsSeenOnFilteredEvents** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedPropertyDefinition

`func NewPatchedPropertyDefinition() *PatchedPropertyDefinition`

NewPatchedPropertyDefinition instantiates a new PatchedPropertyDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPropertyDefinitionWithDefaults

`func NewPatchedPropertyDefinitionWithDefaults() *PatchedPropertyDefinition`

NewPatchedPropertyDefinitionWithDefaults instantiates a new PatchedPropertyDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedPropertyDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedPropertyDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedPropertyDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedPropertyDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedPropertyDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedPropertyDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedPropertyDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedPropertyDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsNumerical

`func (o *PatchedPropertyDefinition) GetIsNumerical() bool`

GetIsNumerical returns the IsNumerical field if non-nil, zero value otherwise.

### GetIsNumericalOk

`func (o *PatchedPropertyDefinition) GetIsNumericalOk() (*bool, bool)`

GetIsNumericalOk returns a tuple with the IsNumerical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNumerical

`func (o *PatchedPropertyDefinition) SetIsNumerical(v bool)`

SetIsNumerical sets IsNumerical field to given value.

### HasIsNumerical

`func (o *PatchedPropertyDefinition) HasIsNumerical() bool`

HasIsNumerical returns a boolean if a field has been set.

### GetPropertyType

`func (o *PatchedPropertyDefinition) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *PatchedPropertyDefinition) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *PatchedPropertyDefinition) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.

### HasPropertyType

`func (o *PatchedPropertyDefinition) HasPropertyType() bool`

HasPropertyType returns a boolean if a field has been set.

### SetPropertyTypeNil

`func (o *PatchedPropertyDefinition) SetPropertyTypeNil(b bool)`

 SetPropertyTypeNil sets the value for PropertyType to be an explicit nil

### UnsetPropertyType
`func (o *PatchedPropertyDefinition) UnsetPropertyType()`

UnsetPropertyType ensures that no value is present for PropertyType, not even an explicit nil
### GetTags

`func (o *PatchedPropertyDefinition) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedPropertyDefinition) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedPropertyDefinition) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedPropertyDefinition) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsSeenOnFilteredEvents

`func (o *PatchedPropertyDefinition) GetIsSeenOnFilteredEvents() string`

GetIsSeenOnFilteredEvents returns the IsSeenOnFilteredEvents field if non-nil, zero value otherwise.

### GetIsSeenOnFilteredEventsOk

`func (o *PatchedPropertyDefinition) GetIsSeenOnFilteredEventsOk() (*string, bool)`

GetIsSeenOnFilteredEventsOk returns a tuple with the IsSeenOnFilteredEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeenOnFilteredEvents

`func (o *PatchedPropertyDefinition) SetIsSeenOnFilteredEvents(v string)`

SetIsSeenOnFilteredEvents sets IsSeenOnFilteredEvents field to given value.

### HasIsSeenOnFilteredEvents

`func (o *PatchedPropertyDefinition) HasIsSeenOnFilteredEvents() bool`

HasIsSeenOnFilteredEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


