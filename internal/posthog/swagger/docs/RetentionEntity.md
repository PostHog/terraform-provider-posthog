# RetentionEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomName** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to [**NullableId1**](Id1.md) |  | [optional] [default to null]
**Kind** | Pointer to [**RetentionEntityKind**](RetentionEntityKind.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **NullableInt32** |  | [optional] 
**Properties** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) | filters on the event | [optional] 
**Type** | Pointer to [**EntityType**](EntityType.md) |  | [optional] 
**Uuid** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRetentionEntity

`func NewRetentionEntity() *RetentionEntity`

NewRetentionEntity instantiates a new RetentionEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionEntityWithDefaults

`func NewRetentionEntityWithDefaults() *RetentionEntity`

NewRetentionEntityWithDefaults instantiates a new RetentionEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomName

`func (o *RetentionEntity) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *RetentionEntity) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *RetentionEntity) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *RetentionEntity) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *RetentionEntity) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *RetentionEntity) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetId

`func (o *RetentionEntity) GetId() Id1`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RetentionEntity) GetIdOk() (*Id1, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RetentionEntity) SetId(v Id1)`

SetId sets Id field to given value.

### HasId

`func (o *RetentionEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RetentionEntity) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RetentionEntity) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKind

`func (o *RetentionEntity) GetKind() RetentionEntityKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RetentionEntity) GetKindOk() (*RetentionEntityKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RetentionEntity) SetKind(v RetentionEntityKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RetentionEntity) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *RetentionEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RetentionEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RetentionEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RetentionEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RetentionEntity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RetentionEntity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrder

`func (o *RetentionEntity) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RetentionEntity) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RetentionEntity) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RetentionEntity) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *RetentionEntity) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *RetentionEntity) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetProperties

`func (o *RetentionEntity) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RetentionEntity) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RetentionEntity) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *RetentionEntity) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *RetentionEntity) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *RetentionEntity) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetType

`func (o *RetentionEntity) GetType() EntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RetentionEntity) GetTypeOk() (*EntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RetentionEntity) SetType(v EntityType)`

SetType sets Type field to given value.

### HasType

`func (o *RetentionEntity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *RetentionEntity) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RetentionEntity) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RetentionEntity) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RetentionEntity) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *RetentionEntity) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *RetentionEntity) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


