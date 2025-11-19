# MinimalPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**DistinctIds** | **string** |  | [readonly] 
**Properties** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**Uuid** | **string** |  | [readonly] 

## Methods

### NewMinimalPerson

`func NewMinimalPerson(id int32, name string, distinctIds string, createdAt time.Time, uuid string, ) *MinimalPerson`

NewMinimalPerson instantiates a new MinimalPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalPersonWithDefaults

`func NewMinimalPersonWithDefaults() *MinimalPerson`

NewMinimalPersonWithDefaults instantiates a new MinimalPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MinimalPerson) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalPerson) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalPerson) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *MinimalPerson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MinimalPerson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MinimalPerson) SetName(v string)`

SetName sets Name field to given value.


### GetDistinctIds

`func (o *MinimalPerson) GetDistinctIds() string`

GetDistinctIds returns the DistinctIds field if non-nil, zero value otherwise.

### GetDistinctIdsOk

`func (o *MinimalPerson) GetDistinctIdsOk() (*string, bool)`

GetDistinctIdsOk returns a tuple with the DistinctIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctIds

`func (o *MinimalPerson) SetDistinctIds(v string)`

SetDistinctIds sets DistinctIds field to given value.


### GetProperties

`func (o *MinimalPerson) GetProperties() interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MinimalPerson) GetPropertiesOk() (*interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MinimalPerson) SetProperties(v interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MinimalPerson) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *MinimalPerson) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *MinimalPerson) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetCreatedAt

`func (o *MinimalPerson) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MinimalPerson) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MinimalPerson) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUuid

`func (o *MinimalPerson) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MinimalPerson) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MinimalPerson) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


