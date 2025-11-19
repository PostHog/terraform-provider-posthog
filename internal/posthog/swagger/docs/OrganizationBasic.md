# OrganizationBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**LogoMediaId** | **NullableString** |  | [readonly] 
**MembershipLevel** | [**NullableMembershipLevelEnum**](MembershipLevelEnum.md) |  | [readonly] 
**MembersCanUsePersonalApiKeys** | Pointer to **bool** |  | [optional] 

## Methods

### NewOrganizationBasic

`func NewOrganizationBasic(id string, name string, slug string, logoMediaId NullableString, membershipLevel NullableMembershipLevelEnum, ) *OrganizationBasic`

NewOrganizationBasic instantiates a new OrganizationBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationBasicWithDefaults

`func NewOrganizationBasicWithDefaults() *OrganizationBasic`

NewOrganizationBasicWithDefaults instantiates a new OrganizationBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationBasic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationBasic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationBasic) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OrganizationBasic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationBasic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationBasic) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *OrganizationBasic) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OrganizationBasic) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OrganizationBasic) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetLogoMediaId

`func (o *OrganizationBasic) GetLogoMediaId() string`

GetLogoMediaId returns the LogoMediaId field if non-nil, zero value otherwise.

### GetLogoMediaIdOk

`func (o *OrganizationBasic) GetLogoMediaIdOk() (*string, bool)`

GetLogoMediaIdOk returns a tuple with the LogoMediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoMediaId

`func (o *OrganizationBasic) SetLogoMediaId(v string)`

SetLogoMediaId sets LogoMediaId field to given value.


### SetLogoMediaIdNil

`func (o *OrganizationBasic) SetLogoMediaIdNil(b bool)`

 SetLogoMediaIdNil sets the value for LogoMediaId to be an explicit nil

### UnsetLogoMediaId
`func (o *OrganizationBasic) UnsetLogoMediaId()`

UnsetLogoMediaId ensures that no value is present for LogoMediaId, not even an explicit nil
### GetMembershipLevel

`func (o *OrganizationBasic) GetMembershipLevel() MembershipLevelEnum`

GetMembershipLevel returns the MembershipLevel field if non-nil, zero value otherwise.

### GetMembershipLevelOk

`func (o *OrganizationBasic) GetMembershipLevelOk() (*MembershipLevelEnum, bool)`

GetMembershipLevelOk returns a tuple with the MembershipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipLevel

`func (o *OrganizationBasic) SetMembershipLevel(v MembershipLevelEnum)`

SetMembershipLevel sets MembershipLevel field to given value.


### SetMembershipLevelNil

`func (o *OrganizationBasic) SetMembershipLevelNil(b bool)`

 SetMembershipLevelNil sets the value for MembershipLevel to be an explicit nil

### UnsetMembershipLevel
`func (o *OrganizationBasic) UnsetMembershipLevel()`

UnsetMembershipLevel ensures that no value is present for MembershipLevel, not even an explicit nil
### GetMembersCanUsePersonalApiKeys

`func (o *OrganizationBasic) GetMembersCanUsePersonalApiKeys() bool`

GetMembersCanUsePersonalApiKeys returns the MembersCanUsePersonalApiKeys field if non-nil, zero value otherwise.

### GetMembersCanUsePersonalApiKeysOk

`func (o *OrganizationBasic) GetMembersCanUsePersonalApiKeysOk() (*bool, bool)`

GetMembersCanUsePersonalApiKeysOk returns a tuple with the MembersCanUsePersonalApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanUsePersonalApiKeys

`func (o *OrganizationBasic) SetMembersCanUsePersonalApiKeys(v bool)`

SetMembersCanUsePersonalApiKeys sets MembersCanUsePersonalApiKeys field to given value.

### HasMembersCanUsePersonalApiKeys

`func (o *OrganizationBasic) HasMembersCanUsePersonalApiKeys() bool`

HasMembersCanUsePersonalApiKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


