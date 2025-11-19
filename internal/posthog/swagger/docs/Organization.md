# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | [readonly] 
**LogoMediaId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**MembershipLevel** | [**NullableMembershipLevelEnum**](MembershipLevelEnum.md) |  | [readonly] 
**PluginsAccessLevel** | [**PluginsAccessLevelEnum**](PluginsAccessLevelEnum.md) |  | [readonly] 
**Teams** | **[]map[string]interface{}** |  | [readonly] 
**Projects** | **[]map[string]interface{}** |  | [readonly] 
**AvailableProductFeatures** | **[]interface{}** |  | [readonly] 
**IsMemberJoinEmailEnabled** | Pointer to **bool** |  | [optional] 
**Metadata** | **string** |  | [readonly] 
**CustomerId** | **NullableString** |  | [readonly] 
**Enforce2fa** | Pointer to **NullableBool** |  | [optional] 
**MembersCanInvite** | Pointer to **NullableBool** |  | [optional] 
**MembersCanUsePersonalApiKeys** | Pointer to **bool** |  | [optional] 
**AllowPubliclySharedResources** | Pointer to **bool** |  | [optional] 
**MemberCount** | **string** |  | [readonly] 
**IsAiDataProcessingApproved** | Pointer to **NullableBool** |  | [optional] 
**DefaultExperimentStatsMethod** | Pointer to **NullableString** | Default statistical method for new experiments in this organization.  * &#x60;bayesian&#x60; - Bayesian * &#x60;frequentist&#x60; - Frequentist | [optional] 
**DefaultAnonymizeIps** | Pointer to **bool** | Default setting for &#39;Discard client IP data&#39; for new projects in this organization. | [optional] 
**DefaultRoleId** | Pointer to **NullableString** | ID of the role to automatically assign to new members joining the organization | [optional] 

## Methods

### NewOrganization

`func NewOrganization(id string, name string, slug string, createdAt time.Time, updatedAt time.Time, membershipLevel NullableMembershipLevelEnum, pluginsAccessLevel PluginsAccessLevelEnum, teams []map[string]interface{}, projects []map[string]interface{}, availableProductFeatures []interface{}, metadata string, customerId NullableString, memberCount string, ) *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Organization) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Organization) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Organization) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetLogoMediaId

`func (o *Organization) GetLogoMediaId() string`

GetLogoMediaId returns the LogoMediaId field if non-nil, zero value otherwise.

### GetLogoMediaIdOk

`func (o *Organization) GetLogoMediaIdOk() (*string, bool)`

GetLogoMediaIdOk returns a tuple with the LogoMediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoMediaId

`func (o *Organization) SetLogoMediaId(v string)`

SetLogoMediaId sets LogoMediaId field to given value.

### HasLogoMediaId

`func (o *Organization) HasLogoMediaId() bool`

HasLogoMediaId returns a boolean if a field has been set.

### SetLogoMediaIdNil

`func (o *Organization) SetLogoMediaIdNil(b bool)`

 SetLogoMediaIdNil sets the value for LogoMediaId to be an explicit nil

### UnsetLogoMediaId
`func (o *Organization) UnsetLogoMediaId()`

UnsetLogoMediaId ensures that no value is present for LogoMediaId, not even an explicit nil
### GetCreatedAt

`func (o *Organization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Organization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Organization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Organization) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Organization) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Organization) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetMembershipLevel

`func (o *Organization) GetMembershipLevel() MembershipLevelEnum`

GetMembershipLevel returns the MembershipLevel field if non-nil, zero value otherwise.

### GetMembershipLevelOk

`func (o *Organization) GetMembershipLevelOk() (*MembershipLevelEnum, bool)`

GetMembershipLevelOk returns a tuple with the MembershipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipLevel

`func (o *Organization) SetMembershipLevel(v MembershipLevelEnum)`

SetMembershipLevel sets MembershipLevel field to given value.


### SetMembershipLevelNil

`func (o *Organization) SetMembershipLevelNil(b bool)`

 SetMembershipLevelNil sets the value for MembershipLevel to be an explicit nil

### UnsetMembershipLevel
`func (o *Organization) UnsetMembershipLevel()`

UnsetMembershipLevel ensures that no value is present for MembershipLevel, not even an explicit nil
### GetPluginsAccessLevel

`func (o *Organization) GetPluginsAccessLevel() PluginsAccessLevelEnum`

GetPluginsAccessLevel returns the PluginsAccessLevel field if non-nil, zero value otherwise.

### GetPluginsAccessLevelOk

`func (o *Organization) GetPluginsAccessLevelOk() (*PluginsAccessLevelEnum, bool)`

GetPluginsAccessLevelOk returns a tuple with the PluginsAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginsAccessLevel

`func (o *Organization) SetPluginsAccessLevel(v PluginsAccessLevelEnum)`

SetPluginsAccessLevel sets PluginsAccessLevel field to given value.


### GetTeams

`func (o *Organization) GetTeams() []map[string]interface{}`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *Organization) GetTeamsOk() (*[]map[string]interface{}, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *Organization) SetTeams(v []map[string]interface{})`

SetTeams sets Teams field to given value.


### GetProjects

`func (o *Organization) GetProjects() []map[string]interface{}`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *Organization) GetProjectsOk() (*[]map[string]interface{}, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *Organization) SetProjects(v []map[string]interface{})`

SetProjects sets Projects field to given value.


### GetAvailableProductFeatures

`func (o *Organization) GetAvailableProductFeatures() []interface{}`

GetAvailableProductFeatures returns the AvailableProductFeatures field if non-nil, zero value otherwise.

### GetAvailableProductFeaturesOk

`func (o *Organization) GetAvailableProductFeaturesOk() (*[]interface{}, bool)`

GetAvailableProductFeaturesOk returns a tuple with the AvailableProductFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableProductFeatures

`func (o *Organization) SetAvailableProductFeatures(v []interface{})`

SetAvailableProductFeatures sets AvailableProductFeatures field to given value.


### SetAvailableProductFeaturesNil

`func (o *Organization) SetAvailableProductFeaturesNil(b bool)`

 SetAvailableProductFeaturesNil sets the value for AvailableProductFeatures to be an explicit nil

### UnsetAvailableProductFeatures
`func (o *Organization) UnsetAvailableProductFeatures()`

UnsetAvailableProductFeatures ensures that no value is present for AvailableProductFeatures, not even an explicit nil
### GetIsMemberJoinEmailEnabled

`func (o *Organization) GetIsMemberJoinEmailEnabled() bool`

GetIsMemberJoinEmailEnabled returns the IsMemberJoinEmailEnabled field if non-nil, zero value otherwise.

### GetIsMemberJoinEmailEnabledOk

`func (o *Organization) GetIsMemberJoinEmailEnabledOk() (*bool, bool)`

GetIsMemberJoinEmailEnabledOk returns a tuple with the IsMemberJoinEmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMemberJoinEmailEnabled

`func (o *Organization) SetIsMemberJoinEmailEnabled(v bool)`

SetIsMemberJoinEmailEnabled sets IsMemberJoinEmailEnabled field to given value.

### HasIsMemberJoinEmailEnabled

`func (o *Organization) HasIsMemberJoinEmailEnabled() bool`

HasIsMemberJoinEmailEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *Organization) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Organization) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Organization) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.


### GetCustomerId

`func (o *Organization) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Organization) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Organization) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *Organization) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *Organization) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetEnforce2fa

`func (o *Organization) GetEnforce2fa() bool`

GetEnforce2fa returns the Enforce2fa field if non-nil, zero value otherwise.

### GetEnforce2faOk

`func (o *Organization) GetEnforce2faOk() (*bool, bool)`

GetEnforce2faOk returns a tuple with the Enforce2fa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforce2fa

`func (o *Organization) SetEnforce2fa(v bool)`

SetEnforce2fa sets Enforce2fa field to given value.

### HasEnforce2fa

`func (o *Organization) HasEnforce2fa() bool`

HasEnforce2fa returns a boolean if a field has been set.

### SetEnforce2faNil

`func (o *Organization) SetEnforce2faNil(b bool)`

 SetEnforce2faNil sets the value for Enforce2fa to be an explicit nil

### UnsetEnforce2fa
`func (o *Organization) UnsetEnforce2fa()`

UnsetEnforce2fa ensures that no value is present for Enforce2fa, not even an explicit nil
### GetMembersCanInvite

`func (o *Organization) GetMembersCanInvite() bool`

GetMembersCanInvite returns the MembersCanInvite field if non-nil, zero value otherwise.

### GetMembersCanInviteOk

`func (o *Organization) GetMembersCanInviteOk() (*bool, bool)`

GetMembersCanInviteOk returns a tuple with the MembersCanInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanInvite

`func (o *Organization) SetMembersCanInvite(v bool)`

SetMembersCanInvite sets MembersCanInvite field to given value.

### HasMembersCanInvite

`func (o *Organization) HasMembersCanInvite() bool`

HasMembersCanInvite returns a boolean if a field has been set.

### SetMembersCanInviteNil

`func (o *Organization) SetMembersCanInviteNil(b bool)`

 SetMembersCanInviteNil sets the value for MembersCanInvite to be an explicit nil

### UnsetMembersCanInvite
`func (o *Organization) UnsetMembersCanInvite()`

UnsetMembersCanInvite ensures that no value is present for MembersCanInvite, not even an explicit nil
### GetMembersCanUsePersonalApiKeys

`func (o *Organization) GetMembersCanUsePersonalApiKeys() bool`

GetMembersCanUsePersonalApiKeys returns the MembersCanUsePersonalApiKeys field if non-nil, zero value otherwise.

### GetMembersCanUsePersonalApiKeysOk

`func (o *Organization) GetMembersCanUsePersonalApiKeysOk() (*bool, bool)`

GetMembersCanUsePersonalApiKeysOk returns a tuple with the MembersCanUsePersonalApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanUsePersonalApiKeys

`func (o *Organization) SetMembersCanUsePersonalApiKeys(v bool)`

SetMembersCanUsePersonalApiKeys sets MembersCanUsePersonalApiKeys field to given value.

### HasMembersCanUsePersonalApiKeys

`func (o *Organization) HasMembersCanUsePersonalApiKeys() bool`

HasMembersCanUsePersonalApiKeys returns a boolean if a field has been set.

### GetAllowPubliclySharedResources

`func (o *Organization) GetAllowPubliclySharedResources() bool`

GetAllowPubliclySharedResources returns the AllowPubliclySharedResources field if non-nil, zero value otherwise.

### GetAllowPubliclySharedResourcesOk

`func (o *Organization) GetAllowPubliclySharedResourcesOk() (*bool, bool)`

GetAllowPubliclySharedResourcesOk returns a tuple with the AllowPubliclySharedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPubliclySharedResources

`func (o *Organization) SetAllowPubliclySharedResources(v bool)`

SetAllowPubliclySharedResources sets AllowPubliclySharedResources field to given value.

### HasAllowPubliclySharedResources

`func (o *Organization) HasAllowPubliclySharedResources() bool`

HasAllowPubliclySharedResources returns a boolean if a field has been set.

### GetMemberCount

`func (o *Organization) GetMemberCount() string`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *Organization) GetMemberCountOk() (*string, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *Organization) SetMemberCount(v string)`

SetMemberCount sets MemberCount field to given value.


### GetIsAiDataProcessingApproved

`func (o *Organization) GetIsAiDataProcessingApproved() bool`

GetIsAiDataProcessingApproved returns the IsAiDataProcessingApproved field if non-nil, zero value otherwise.

### GetIsAiDataProcessingApprovedOk

`func (o *Organization) GetIsAiDataProcessingApprovedOk() (*bool, bool)`

GetIsAiDataProcessingApprovedOk returns a tuple with the IsAiDataProcessingApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAiDataProcessingApproved

`func (o *Organization) SetIsAiDataProcessingApproved(v bool)`

SetIsAiDataProcessingApproved sets IsAiDataProcessingApproved field to given value.

### HasIsAiDataProcessingApproved

`func (o *Organization) HasIsAiDataProcessingApproved() bool`

HasIsAiDataProcessingApproved returns a boolean if a field has been set.

### SetIsAiDataProcessingApprovedNil

`func (o *Organization) SetIsAiDataProcessingApprovedNil(b bool)`

 SetIsAiDataProcessingApprovedNil sets the value for IsAiDataProcessingApproved to be an explicit nil

### UnsetIsAiDataProcessingApproved
`func (o *Organization) UnsetIsAiDataProcessingApproved()`

UnsetIsAiDataProcessingApproved ensures that no value is present for IsAiDataProcessingApproved, not even an explicit nil
### GetDefaultExperimentStatsMethod

`func (o *Organization) GetDefaultExperimentStatsMethod() string`

GetDefaultExperimentStatsMethod returns the DefaultExperimentStatsMethod field if non-nil, zero value otherwise.

### GetDefaultExperimentStatsMethodOk

`func (o *Organization) GetDefaultExperimentStatsMethodOk() (*string, bool)`

GetDefaultExperimentStatsMethodOk returns a tuple with the DefaultExperimentStatsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExperimentStatsMethod

`func (o *Organization) SetDefaultExperimentStatsMethod(v string)`

SetDefaultExperimentStatsMethod sets DefaultExperimentStatsMethod field to given value.

### HasDefaultExperimentStatsMethod

`func (o *Organization) HasDefaultExperimentStatsMethod() bool`

HasDefaultExperimentStatsMethod returns a boolean if a field has been set.

### SetDefaultExperimentStatsMethodNil

`func (o *Organization) SetDefaultExperimentStatsMethodNil(b bool)`

 SetDefaultExperimentStatsMethodNil sets the value for DefaultExperimentStatsMethod to be an explicit nil

### UnsetDefaultExperimentStatsMethod
`func (o *Organization) UnsetDefaultExperimentStatsMethod()`

UnsetDefaultExperimentStatsMethod ensures that no value is present for DefaultExperimentStatsMethod, not even an explicit nil
### GetDefaultAnonymizeIps

`func (o *Organization) GetDefaultAnonymizeIps() bool`

GetDefaultAnonymizeIps returns the DefaultAnonymizeIps field if non-nil, zero value otherwise.

### GetDefaultAnonymizeIpsOk

`func (o *Organization) GetDefaultAnonymizeIpsOk() (*bool, bool)`

GetDefaultAnonymizeIpsOk returns a tuple with the DefaultAnonymizeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAnonymizeIps

`func (o *Organization) SetDefaultAnonymizeIps(v bool)`

SetDefaultAnonymizeIps sets DefaultAnonymizeIps field to given value.

### HasDefaultAnonymizeIps

`func (o *Organization) HasDefaultAnonymizeIps() bool`

HasDefaultAnonymizeIps returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *Organization) GetDefaultRoleId() string`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *Organization) GetDefaultRoleIdOk() (*string, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *Organization) SetDefaultRoleId(v string)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *Organization) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.

### SetDefaultRoleIdNil

`func (o *Organization) SetDefaultRoleIdNil(b bool)`

 SetDefaultRoleIdNil sets the value for DefaultRoleId to be an explicit nil

### UnsetDefaultRoleId
`func (o *Organization) UnsetDefaultRoleId()`

UnsetDefaultRoleId ensures that no value is present for DefaultRoleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


