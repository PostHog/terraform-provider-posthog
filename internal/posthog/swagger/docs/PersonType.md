# PersonType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**DistinctIds** | **[]string** |  | 
**Id** | Pointer to **NullableString** |  | [optional] 
**IsIdentified** | Pointer to **NullableBool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Properties** | **map[string]interface{}** |  | 
**Uuid** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPersonType

`func NewPersonType(distinctIds []string, properties map[string]interface{}, ) *PersonType`

NewPersonType instantiates a new PersonType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonTypeWithDefaults

`func NewPersonTypeWithDefaults() *PersonType`

NewPersonTypeWithDefaults instantiates a new PersonType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PersonType) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PersonType) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PersonType) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PersonType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PersonType) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PersonType) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDistinctIds

`func (o *PersonType) GetDistinctIds() []string`

GetDistinctIds returns the DistinctIds field if non-nil, zero value otherwise.

### GetDistinctIdsOk

`func (o *PersonType) GetDistinctIdsOk() (*[]string, bool)`

GetDistinctIdsOk returns a tuple with the DistinctIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctIds

`func (o *PersonType) SetDistinctIds(v []string)`

SetDistinctIds sets DistinctIds field to given value.


### GetId

`func (o *PersonType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PersonType) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PersonType) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PersonType) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsIdentified

`func (o *PersonType) GetIsIdentified() bool`

GetIsIdentified returns the IsIdentified field if non-nil, zero value otherwise.

### GetIsIdentifiedOk

`func (o *PersonType) GetIsIdentifiedOk() (*bool, bool)`

GetIsIdentifiedOk returns a tuple with the IsIdentified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIdentified

`func (o *PersonType) SetIsIdentified(v bool)`

SetIsIdentified sets IsIdentified field to given value.

### HasIsIdentified

`func (o *PersonType) HasIsIdentified() bool`

HasIsIdentified returns a boolean if a field has been set.

### SetIsIdentifiedNil

`func (o *PersonType) SetIsIdentifiedNil(b bool)`

 SetIsIdentifiedNil sets the value for IsIdentified to be an explicit nil

### UnsetIsIdentified
`func (o *PersonType) UnsetIsIdentified()`

UnsetIsIdentified ensures that no value is present for IsIdentified, not even an explicit nil
### GetName

`func (o *PersonType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PersonType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PersonType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PersonType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PersonType) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PersonType) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProperties

`func (o *PersonType) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PersonType) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PersonType) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.


### GetUuid

`func (o *PersonType) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PersonType) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PersonType) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PersonType) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *PersonType) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *PersonType) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


