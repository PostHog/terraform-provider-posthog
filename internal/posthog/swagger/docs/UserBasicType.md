# UserBasicType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinctId** | **string** |  | 
**Email** | **string** |  | 
**FirstName** | **string** |  | 
**HedgehogConfig** | Pointer to [**MinimalHedgehogConfig**](MinimalHedgehogConfig.md) |  | [optional] 
**Id** | **float32** |  | 
**IsEmailVerified** | Pointer to **interface{}** |  | [optional] [default to null]
**LastName** | Pointer to **NullableString** |  | [optional] 
**RoleAtOrganization** | Pointer to **NullableString** |  | [optional] 
**Uuid** | **string** |  | 

## Methods

### NewUserBasicType

`func NewUserBasicType(distinctId string, email string, firstName string, id float32, uuid string, ) *UserBasicType`

NewUserBasicType instantiates a new UserBasicType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBasicTypeWithDefaults

`func NewUserBasicTypeWithDefaults() *UserBasicType`

NewUserBasicTypeWithDefaults instantiates a new UserBasicType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctId

`func (o *UserBasicType) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *UserBasicType) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *UserBasicType) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.


### GetEmail

`func (o *UserBasicType) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserBasicType) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserBasicType) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *UserBasicType) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserBasicType) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserBasicType) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetHedgehogConfig

`func (o *UserBasicType) GetHedgehogConfig() MinimalHedgehogConfig`

GetHedgehogConfig returns the HedgehogConfig field if non-nil, zero value otherwise.

### GetHedgehogConfigOk

`func (o *UserBasicType) GetHedgehogConfigOk() (*MinimalHedgehogConfig, bool)`

GetHedgehogConfigOk returns a tuple with the HedgehogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHedgehogConfig

`func (o *UserBasicType) SetHedgehogConfig(v MinimalHedgehogConfig)`

SetHedgehogConfig sets HedgehogConfig field to given value.

### HasHedgehogConfig

`func (o *UserBasicType) HasHedgehogConfig() bool`

HasHedgehogConfig returns a boolean if a field has been set.

### GetId

`func (o *UserBasicType) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserBasicType) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserBasicType) SetId(v float32)`

SetId sets Id field to given value.


### GetIsEmailVerified

`func (o *UserBasicType) GetIsEmailVerified() interface{}`

GetIsEmailVerified returns the IsEmailVerified field if non-nil, zero value otherwise.

### GetIsEmailVerifiedOk

`func (o *UserBasicType) GetIsEmailVerifiedOk() (*interface{}, bool)`

GetIsEmailVerifiedOk returns a tuple with the IsEmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmailVerified

`func (o *UserBasicType) SetIsEmailVerified(v interface{})`

SetIsEmailVerified sets IsEmailVerified field to given value.

### HasIsEmailVerified

`func (o *UserBasicType) HasIsEmailVerified() bool`

HasIsEmailVerified returns a boolean if a field has been set.

### SetIsEmailVerifiedNil

`func (o *UserBasicType) SetIsEmailVerifiedNil(b bool)`

 SetIsEmailVerifiedNil sets the value for IsEmailVerified to be an explicit nil

### UnsetIsEmailVerified
`func (o *UserBasicType) UnsetIsEmailVerified()`

UnsetIsEmailVerified ensures that no value is present for IsEmailVerified, not even an explicit nil
### GetLastName

`func (o *UserBasicType) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserBasicType) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserBasicType) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserBasicType) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UserBasicType) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UserBasicType) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetRoleAtOrganization

`func (o *UserBasicType) GetRoleAtOrganization() string`

GetRoleAtOrganization returns the RoleAtOrganization field if non-nil, zero value otherwise.

### GetRoleAtOrganizationOk

`func (o *UserBasicType) GetRoleAtOrganizationOk() (*string, bool)`

GetRoleAtOrganizationOk returns a tuple with the RoleAtOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAtOrganization

`func (o *UserBasicType) SetRoleAtOrganization(v string)`

SetRoleAtOrganization sets RoleAtOrganization field to given value.

### HasRoleAtOrganization

`func (o *UserBasicType) HasRoleAtOrganization() bool`

HasRoleAtOrganization returns a boolean if a field has been set.

### SetRoleAtOrganizationNil

`func (o *UserBasicType) SetRoleAtOrganizationNil(b bool)`

 SetRoleAtOrganizationNil sets the value for RoleAtOrganization to be an explicit nil

### UnsetRoleAtOrganization
`func (o *UserBasicType) UnsetRoleAtOrganization()`

UnsetRoleAtOrganization ensures that no value is present for RoleAtOrganization, not even an explicit nil
### GetUuid

`func (o *UserBasicType) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UserBasicType) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UserBasicType) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


