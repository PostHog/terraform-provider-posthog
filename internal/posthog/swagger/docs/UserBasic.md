# UserBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Uuid** | **string** |  | [readonly] 
**DistinctId** | Pointer to **NullableString** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**IsEmailVerified** | Pointer to **NullableBool** |  | [optional] 
**HedgehogConfig** | **map[string]interface{}** |  | [readonly] 
**RoleAtOrganization** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserBasic

`func NewUserBasic(id int32, uuid string, email string, hedgehogConfig map[string]interface{}, ) *UserBasic`

NewUserBasic instantiates a new UserBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBasicWithDefaults

`func NewUserBasicWithDefaults() *UserBasic`

NewUserBasicWithDefaults instantiates a new UserBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserBasic) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserBasic) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserBasic) SetId(v int32)`

SetId sets Id field to given value.


### GetUuid

`func (o *UserBasic) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UserBasic) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UserBasic) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetDistinctId

`func (o *UserBasic) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *UserBasic) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *UserBasic) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *UserBasic) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### SetDistinctIdNil

`func (o *UserBasic) SetDistinctIdNil(b bool)`

 SetDistinctIdNil sets the value for DistinctId to be an explicit nil

### UnsetDistinctId
`func (o *UserBasic) UnsetDistinctId()`

UnsetDistinctId ensures that no value is present for DistinctId, not even an explicit nil
### GetFirstName

`func (o *UserBasic) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserBasic) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserBasic) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserBasic) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserBasic) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserBasic) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserBasic) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserBasic) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *UserBasic) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserBasic) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserBasic) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIsEmailVerified

`func (o *UserBasic) GetIsEmailVerified() bool`

GetIsEmailVerified returns the IsEmailVerified field if non-nil, zero value otherwise.

### GetIsEmailVerifiedOk

`func (o *UserBasic) GetIsEmailVerifiedOk() (*bool, bool)`

GetIsEmailVerifiedOk returns a tuple with the IsEmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmailVerified

`func (o *UserBasic) SetIsEmailVerified(v bool)`

SetIsEmailVerified sets IsEmailVerified field to given value.

### HasIsEmailVerified

`func (o *UserBasic) HasIsEmailVerified() bool`

HasIsEmailVerified returns a boolean if a field has been set.

### SetIsEmailVerifiedNil

`func (o *UserBasic) SetIsEmailVerifiedNil(b bool)`

 SetIsEmailVerifiedNil sets the value for IsEmailVerified to be an explicit nil

### UnsetIsEmailVerified
`func (o *UserBasic) UnsetIsEmailVerified()`

UnsetIsEmailVerified ensures that no value is present for IsEmailVerified, not even an explicit nil
### GetHedgehogConfig

`func (o *UserBasic) GetHedgehogConfig() map[string]interface{}`

GetHedgehogConfig returns the HedgehogConfig field if non-nil, zero value otherwise.

### GetHedgehogConfigOk

`func (o *UserBasic) GetHedgehogConfigOk() (*map[string]interface{}, bool)`

GetHedgehogConfigOk returns a tuple with the HedgehogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHedgehogConfig

`func (o *UserBasic) SetHedgehogConfig(v map[string]interface{})`

SetHedgehogConfig sets HedgehogConfig field to given value.


### SetHedgehogConfigNil

`func (o *UserBasic) SetHedgehogConfigNil(b bool)`

 SetHedgehogConfigNil sets the value for HedgehogConfig to be an explicit nil

### UnsetHedgehogConfig
`func (o *UserBasic) UnsetHedgehogConfig()`

UnsetHedgehogConfig ensures that no value is present for HedgehogConfig, not even an explicit nil
### GetRoleAtOrganization

`func (o *UserBasic) GetRoleAtOrganization() string`

GetRoleAtOrganization returns the RoleAtOrganization field if non-nil, zero value otherwise.

### GetRoleAtOrganizationOk

`func (o *UserBasic) GetRoleAtOrganizationOk() (*string, bool)`

GetRoleAtOrganizationOk returns a tuple with the RoleAtOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAtOrganization

`func (o *UserBasic) SetRoleAtOrganization(v string)`

SetRoleAtOrganization sets RoleAtOrganization field to given value.

### HasRoleAtOrganization

`func (o *UserBasic) HasRoleAtOrganization() bool`

HasRoleAtOrganization returns a boolean if a field has been set.

### SetRoleAtOrganizationNil

`func (o *UserBasic) SetRoleAtOrganizationNil(b bool)`

 SetRoleAtOrganizationNil sets the value for RoleAtOrganization to be an explicit nil

### UnsetRoleAtOrganization
`func (o *UserBasic) UnsetRoleAtOrganization()`

UnsetRoleAtOrganization ensures that no value is present for RoleAtOrganization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


