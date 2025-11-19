# PatchedPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**DistinctIds** | Pointer to **[]string** |  | [optional] [readonly] 
**Properties** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedPerson

`func NewPatchedPerson() *PatchedPerson`

NewPatchedPerson instantiates a new PatchedPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPersonWithDefaults

`func NewPatchedPersonWithDefaults() *PatchedPerson`

NewPatchedPersonWithDefaults instantiates a new PatchedPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedPerson) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedPerson) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedPerson) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedPerson) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedPerson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedPerson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedPerson) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedPerson) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDistinctIds

`func (o *PatchedPerson) GetDistinctIds() []string`

GetDistinctIds returns the DistinctIds field if non-nil, zero value otherwise.

### GetDistinctIdsOk

`func (o *PatchedPerson) GetDistinctIdsOk() (*[]string, bool)`

GetDistinctIdsOk returns a tuple with the DistinctIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctIds

`func (o *PatchedPerson) SetDistinctIds(v []string)`

SetDistinctIds sets DistinctIds field to given value.

### HasDistinctIds

`func (o *PatchedPerson) HasDistinctIds() bool`

HasDistinctIds returns a boolean if a field has been set.

### GetProperties

`func (o *PatchedPerson) GetProperties() interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PatchedPerson) GetPropertiesOk() (*interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PatchedPerson) SetProperties(v interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PatchedPerson) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *PatchedPerson) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *PatchedPerson) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetCreatedAt

`func (o *PatchedPerson) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedPerson) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedPerson) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedPerson) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUuid

`func (o *PatchedPerson) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PatchedPerson) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PatchedPerson) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PatchedPerson) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


