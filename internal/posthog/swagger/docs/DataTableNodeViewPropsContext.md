# DataTableNodeViewPropsContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventDefinitionId** | Pointer to **NullableString** |  | [optional] 
**Type** | [**DataTableNodeViewPropsContextType**](DataTableNodeViewPropsContextType.md) |  | 

## Methods

### NewDataTableNodeViewPropsContext

`func NewDataTableNodeViewPropsContext(type_ DataTableNodeViewPropsContextType, ) *DataTableNodeViewPropsContext`

NewDataTableNodeViewPropsContext instantiates a new DataTableNodeViewPropsContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTableNodeViewPropsContextWithDefaults

`func NewDataTableNodeViewPropsContextWithDefaults() *DataTableNodeViewPropsContext`

NewDataTableNodeViewPropsContextWithDefaults instantiates a new DataTableNodeViewPropsContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventDefinitionId

`func (o *DataTableNodeViewPropsContext) GetEventDefinitionId() string`

GetEventDefinitionId returns the EventDefinitionId field if non-nil, zero value otherwise.

### GetEventDefinitionIdOk

`func (o *DataTableNodeViewPropsContext) GetEventDefinitionIdOk() (*string, bool)`

GetEventDefinitionIdOk returns a tuple with the EventDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDefinitionId

`func (o *DataTableNodeViewPropsContext) SetEventDefinitionId(v string)`

SetEventDefinitionId sets EventDefinitionId field to given value.

### HasEventDefinitionId

`func (o *DataTableNodeViewPropsContext) HasEventDefinitionId() bool`

HasEventDefinitionId returns a boolean if a field has been set.

### SetEventDefinitionIdNil

`func (o *DataTableNodeViewPropsContext) SetEventDefinitionIdNil(b bool)`

 SetEventDefinitionIdNil sets the value for EventDefinitionId to be an explicit nil

### UnsetEventDefinitionId
`func (o *DataTableNodeViewPropsContext) UnsetEventDefinitionId()`

UnsetEventDefinitionId ensures that no value is present for EventDefinitionId, not even an explicit nil
### GetType

`func (o *DataTableNodeViewPropsContext) GetType() DataTableNodeViewPropsContextType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataTableNodeViewPropsContext) GetTypeOk() (*DataTableNodeViewPropsContextType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataTableNodeViewPropsContext) SetType(v DataTableNodeViewPropsContextType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


