# PatchedOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] [readonly] 
**LogoMediaId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**MembershipLevel** | Pointer to [**NullableMembershipLevelEnum**](MembershipLevelEnum.md) |  | [optional] [readonly] 
**PluginsAccessLevel** | Pointer to [**PluginsAccessLevelEnum**](PluginsAccessLevelEnum.md) |  | [optional] [readonly] 
**Teams** | Pointer to **[]map[string]interface{}** |  | [optional] [readonly] 
**Projects** | Pointer to **[]map[string]interface{}** |  | [optional] [readonly] 
**AvailableProductFeatures** | Pointer to **[]interface{}** |  | [optional] [readonly] 
**IsMemberJoinEmailEnabled** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] [readonly] 
**CustomerId** | Pointer to **NullableString** |  | [optional] [readonly] 
**Enforce2fa** | Pointer to **NullableBool** |  | [optional] 
**MembersCanInvite** | Pointer to **NullableBool** |  | [optional] 
**MembersCanUsePersonalApiKeys** | Pointer to **bool** |  | [optional] 
**AllowPubliclySharedResources** | Pointer to **bool** |  | [optional] 
**MemberCount** | Pointer to **string** |  | [optional] [readonly] 
**IsAiDataProcessingApproved** | Pointer to **NullableBool** |  | [optional] 
**DefaultExperimentStatsMethod** | Pointer to **NullableString** | Default statistical method for new experiments in this organization.  * &#x60;bayesian&#x60; - Bayesian * &#x60;frequentist&#x60; - Frequentist | [optional] 
**DefaultAnonymizeIps** | Pointer to **bool** | Default setting for &#39;Discard client IP data&#39; for new projects in this organization. | [optional] 
**DefaultRoleId** | Pointer to **NullableString** | ID of the role to automatically assign to new members joining the organization | [optional] 

## Methods

### NewPatchedOrganization

`func NewPatchedOrganization() *PatchedOrganization`

NewPatchedOrganization instantiates a new PatchedOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedOrganizationWithDefaults

`func NewPatchedOrganizationWithDefaults() *PatchedOrganization`

NewPatchedOrganizationWithDefaults instantiates a new PatchedOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedOrganization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedOrganization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedOrganization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedOrganization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedOrganization) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedOrganization) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedOrganization) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedOrganization) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetLogoMediaId

`func (o *PatchedOrganization) GetLogoMediaId() string`

GetLogoMediaId returns the LogoMediaId field if non-nil, zero value otherwise.

### GetLogoMediaIdOk

`func (o *PatchedOrganization) GetLogoMediaIdOk() (*string, bool)`

GetLogoMediaIdOk returns a tuple with the LogoMediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoMediaId

`func (o *PatchedOrganization) SetLogoMediaId(v string)`

SetLogoMediaId sets LogoMediaId field to given value.

### HasLogoMediaId

`func (o *PatchedOrganization) HasLogoMediaId() bool`

HasLogoMediaId returns a boolean if a field has been set.

### SetLogoMediaIdNil

`func (o *PatchedOrganization) SetLogoMediaIdNil(b bool)`

 SetLogoMediaIdNil sets the value for LogoMediaId to be an explicit nil

### UnsetLogoMediaId
`func (o *PatchedOrganization) UnsetLogoMediaId()`

UnsetLogoMediaId ensures that no value is present for LogoMediaId, not even an explicit nil
### GetCreatedAt

`func (o *PatchedOrganization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedOrganization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedOrganization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedOrganization) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedOrganization) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedOrganization) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedOrganization) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedOrganization) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetMembershipLevel

`func (o *PatchedOrganization) GetMembershipLevel() MembershipLevelEnum`

GetMembershipLevel returns the MembershipLevel field if non-nil, zero value otherwise.

### GetMembershipLevelOk

`func (o *PatchedOrganization) GetMembershipLevelOk() (*MembershipLevelEnum, bool)`

GetMembershipLevelOk returns a tuple with the MembershipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipLevel

`func (o *PatchedOrganization) SetMembershipLevel(v MembershipLevelEnum)`

SetMembershipLevel sets MembershipLevel field to given value.

### HasMembershipLevel

`func (o *PatchedOrganization) HasMembershipLevel() bool`

HasMembershipLevel returns a boolean if a field has been set.

### SetMembershipLevelNil

`func (o *PatchedOrganization) SetMembershipLevelNil(b bool)`

 SetMembershipLevelNil sets the value for MembershipLevel to be an explicit nil

### UnsetMembershipLevel
`func (o *PatchedOrganization) UnsetMembershipLevel()`

UnsetMembershipLevel ensures that no value is present for MembershipLevel, not even an explicit nil
### GetPluginsAccessLevel

`func (o *PatchedOrganization) GetPluginsAccessLevel() PluginsAccessLevelEnum`

GetPluginsAccessLevel returns the PluginsAccessLevel field if non-nil, zero value otherwise.

### GetPluginsAccessLevelOk

`func (o *PatchedOrganization) GetPluginsAccessLevelOk() (*PluginsAccessLevelEnum, bool)`

GetPluginsAccessLevelOk returns a tuple with the PluginsAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginsAccessLevel

`func (o *PatchedOrganization) SetPluginsAccessLevel(v PluginsAccessLevelEnum)`

SetPluginsAccessLevel sets PluginsAccessLevel field to given value.

### HasPluginsAccessLevel

`func (o *PatchedOrganization) HasPluginsAccessLevel() bool`

HasPluginsAccessLevel returns a boolean if a field has been set.

### GetTeams

`func (o *PatchedOrganization) GetTeams() []map[string]interface{}`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *PatchedOrganization) GetTeamsOk() (*[]map[string]interface{}, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *PatchedOrganization) SetTeams(v []map[string]interface{})`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *PatchedOrganization) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetProjects

`func (o *PatchedOrganization) GetProjects() []map[string]interface{}`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *PatchedOrganization) GetProjectsOk() (*[]map[string]interface{}, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *PatchedOrganization) SetProjects(v []map[string]interface{})`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *PatchedOrganization) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetAvailableProductFeatures

`func (o *PatchedOrganization) GetAvailableProductFeatures() []interface{}`

GetAvailableProductFeatures returns the AvailableProductFeatures field if non-nil, zero value otherwise.

### GetAvailableProductFeaturesOk

`func (o *PatchedOrganization) GetAvailableProductFeaturesOk() (*[]interface{}, bool)`

GetAvailableProductFeaturesOk returns a tuple with the AvailableProductFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableProductFeatures

`func (o *PatchedOrganization) SetAvailableProductFeatures(v []interface{})`

SetAvailableProductFeatures sets AvailableProductFeatures field to given value.

### HasAvailableProductFeatures

`func (o *PatchedOrganization) HasAvailableProductFeatures() bool`

HasAvailableProductFeatures returns a boolean if a field has been set.

### SetAvailableProductFeaturesNil

`func (o *PatchedOrganization) SetAvailableProductFeaturesNil(b bool)`

 SetAvailableProductFeaturesNil sets the value for AvailableProductFeatures to be an explicit nil

### UnsetAvailableProductFeatures
`func (o *PatchedOrganization) UnsetAvailableProductFeatures()`

UnsetAvailableProductFeatures ensures that no value is present for AvailableProductFeatures, not even an explicit nil
### GetIsMemberJoinEmailEnabled

`func (o *PatchedOrganization) GetIsMemberJoinEmailEnabled() bool`

GetIsMemberJoinEmailEnabled returns the IsMemberJoinEmailEnabled field if non-nil, zero value otherwise.

### GetIsMemberJoinEmailEnabledOk

`func (o *PatchedOrganization) GetIsMemberJoinEmailEnabledOk() (*bool, bool)`

GetIsMemberJoinEmailEnabledOk returns a tuple with the IsMemberJoinEmailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMemberJoinEmailEnabled

`func (o *PatchedOrganization) SetIsMemberJoinEmailEnabled(v bool)`

SetIsMemberJoinEmailEnabled sets IsMemberJoinEmailEnabled field to given value.

### HasIsMemberJoinEmailEnabled

`func (o *PatchedOrganization) HasIsMemberJoinEmailEnabled() bool`

HasIsMemberJoinEmailEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchedOrganization) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchedOrganization) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchedOrganization) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchedOrganization) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCustomerId

`func (o *PatchedOrganization) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PatchedOrganization) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PatchedOrganization) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PatchedOrganization) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *PatchedOrganization) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *PatchedOrganization) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetEnforce2fa

`func (o *PatchedOrganization) GetEnforce2fa() bool`

GetEnforce2fa returns the Enforce2fa field if non-nil, zero value otherwise.

### GetEnforce2faOk

`func (o *PatchedOrganization) GetEnforce2faOk() (*bool, bool)`

GetEnforce2faOk returns a tuple with the Enforce2fa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforce2fa

`func (o *PatchedOrganization) SetEnforce2fa(v bool)`

SetEnforce2fa sets Enforce2fa field to given value.

### HasEnforce2fa

`func (o *PatchedOrganization) HasEnforce2fa() bool`

HasEnforce2fa returns a boolean if a field has been set.

### SetEnforce2faNil

`func (o *PatchedOrganization) SetEnforce2faNil(b bool)`

 SetEnforce2faNil sets the value for Enforce2fa to be an explicit nil

### UnsetEnforce2fa
`func (o *PatchedOrganization) UnsetEnforce2fa()`

UnsetEnforce2fa ensures that no value is present for Enforce2fa, not even an explicit nil
### GetMembersCanInvite

`func (o *PatchedOrganization) GetMembersCanInvite() bool`

GetMembersCanInvite returns the MembersCanInvite field if non-nil, zero value otherwise.

### GetMembersCanInviteOk

`func (o *PatchedOrganization) GetMembersCanInviteOk() (*bool, bool)`

GetMembersCanInviteOk returns a tuple with the MembersCanInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanInvite

`func (o *PatchedOrganization) SetMembersCanInvite(v bool)`

SetMembersCanInvite sets MembersCanInvite field to given value.

### HasMembersCanInvite

`func (o *PatchedOrganization) HasMembersCanInvite() bool`

HasMembersCanInvite returns a boolean if a field has been set.

### SetMembersCanInviteNil

`func (o *PatchedOrganization) SetMembersCanInviteNil(b bool)`

 SetMembersCanInviteNil sets the value for MembersCanInvite to be an explicit nil

### UnsetMembersCanInvite
`func (o *PatchedOrganization) UnsetMembersCanInvite()`

UnsetMembersCanInvite ensures that no value is present for MembersCanInvite, not even an explicit nil
### GetMembersCanUsePersonalApiKeys

`func (o *PatchedOrganization) GetMembersCanUsePersonalApiKeys() bool`

GetMembersCanUsePersonalApiKeys returns the MembersCanUsePersonalApiKeys field if non-nil, zero value otherwise.

### GetMembersCanUsePersonalApiKeysOk

`func (o *PatchedOrganization) GetMembersCanUsePersonalApiKeysOk() (*bool, bool)`

GetMembersCanUsePersonalApiKeysOk returns a tuple with the MembersCanUsePersonalApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanUsePersonalApiKeys

`func (o *PatchedOrganization) SetMembersCanUsePersonalApiKeys(v bool)`

SetMembersCanUsePersonalApiKeys sets MembersCanUsePersonalApiKeys field to given value.

### HasMembersCanUsePersonalApiKeys

`func (o *PatchedOrganization) HasMembersCanUsePersonalApiKeys() bool`

HasMembersCanUsePersonalApiKeys returns a boolean if a field has been set.

### GetAllowPubliclySharedResources

`func (o *PatchedOrganization) GetAllowPubliclySharedResources() bool`

GetAllowPubliclySharedResources returns the AllowPubliclySharedResources field if non-nil, zero value otherwise.

### GetAllowPubliclySharedResourcesOk

`func (o *PatchedOrganization) GetAllowPubliclySharedResourcesOk() (*bool, bool)`

GetAllowPubliclySharedResourcesOk returns a tuple with the AllowPubliclySharedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPubliclySharedResources

`func (o *PatchedOrganization) SetAllowPubliclySharedResources(v bool)`

SetAllowPubliclySharedResources sets AllowPubliclySharedResources field to given value.

### HasAllowPubliclySharedResources

`func (o *PatchedOrganization) HasAllowPubliclySharedResources() bool`

HasAllowPubliclySharedResources returns a boolean if a field has been set.

### GetMemberCount

`func (o *PatchedOrganization) GetMemberCount() string`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *PatchedOrganization) GetMemberCountOk() (*string, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *PatchedOrganization) SetMemberCount(v string)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *PatchedOrganization) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetIsAiDataProcessingApproved

`func (o *PatchedOrganization) GetIsAiDataProcessingApproved() bool`

GetIsAiDataProcessingApproved returns the IsAiDataProcessingApproved field if non-nil, zero value otherwise.

### GetIsAiDataProcessingApprovedOk

`func (o *PatchedOrganization) GetIsAiDataProcessingApprovedOk() (*bool, bool)`

GetIsAiDataProcessingApprovedOk returns a tuple with the IsAiDataProcessingApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAiDataProcessingApproved

`func (o *PatchedOrganization) SetIsAiDataProcessingApproved(v bool)`

SetIsAiDataProcessingApproved sets IsAiDataProcessingApproved field to given value.

### HasIsAiDataProcessingApproved

`func (o *PatchedOrganization) HasIsAiDataProcessingApproved() bool`

HasIsAiDataProcessingApproved returns a boolean if a field has been set.

### SetIsAiDataProcessingApprovedNil

`func (o *PatchedOrganization) SetIsAiDataProcessingApprovedNil(b bool)`

 SetIsAiDataProcessingApprovedNil sets the value for IsAiDataProcessingApproved to be an explicit nil

### UnsetIsAiDataProcessingApproved
`func (o *PatchedOrganization) UnsetIsAiDataProcessingApproved()`

UnsetIsAiDataProcessingApproved ensures that no value is present for IsAiDataProcessingApproved, not even an explicit nil
### GetDefaultExperimentStatsMethod

`func (o *PatchedOrganization) GetDefaultExperimentStatsMethod() string`

GetDefaultExperimentStatsMethod returns the DefaultExperimentStatsMethod field if non-nil, zero value otherwise.

### GetDefaultExperimentStatsMethodOk

`func (o *PatchedOrganization) GetDefaultExperimentStatsMethodOk() (*string, bool)`

GetDefaultExperimentStatsMethodOk returns a tuple with the DefaultExperimentStatsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultExperimentStatsMethod

`func (o *PatchedOrganization) SetDefaultExperimentStatsMethod(v string)`

SetDefaultExperimentStatsMethod sets DefaultExperimentStatsMethod field to given value.

### HasDefaultExperimentStatsMethod

`func (o *PatchedOrganization) HasDefaultExperimentStatsMethod() bool`

HasDefaultExperimentStatsMethod returns a boolean if a field has been set.

### SetDefaultExperimentStatsMethodNil

`func (o *PatchedOrganization) SetDefaultExperimentStatsMethodNil(b bool)`

 SetDefaultExperimentStatsMethodNil sets the value for DefaultExperimentStatsMethod to be an explicit nil

### UnsetDefaultExperimentStatsMethod
`func (o *PatchedOrganization) UnsetDefaultExperimentStatsMethod()`

UnsetDefaultExperimentStatsMethod ensures that no value is present for DefaultExperimentStatsMethod, not even an explicit nil
### GetDefaultAnonymizeIps

`func (o *PatchedOrganization) GetDefaultAnonymizeIps() bool`

GetDefaultAnonymizeIps returns the DefaultAnonymizeIps field if non-nil, zero value otherwise.

### GetDefaultAnonymizeIpsOk

`func (o *PatchedOrganization) GetDefaultAnonymizeIpsOk() (*bool, bool)`

GetDefaultAnonymizeIpsOk returns a tuple with the DefaultAnonymizeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAnonymizeIps

`func (o *PatchedOrganization) SetDefaultAnonymizeIps(v bool)`

SetDefaultAnonymizeIps sets DefaultAnonymizeIps field to given value.

### HasDefaultAnonymizeIps

`func (o *PatchedOrganization) HasDefaultAnonymizeIps() bool`

HasDefaultAnonymizeIps returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *PatchedOrganization) GetDefaultRoleId() string`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *PatchedOrganization) GetDefaultRoleIdOk() (*string, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *PatchedOrganization) SetDefaultRoleId(v string)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *PatchedOrganization) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.

### SetDefaultRoleIdNil

`func (o *PatchedOrganization) SetDefaultRoleIdNil(b bool)`

 SetDefaultRoleIdNil sets the value for DefaultRoleId to be an explicit nil

### UnsetDefaultRoleId
`func (o *PatchedOrganization) UnsetDefaultRoleId()`

UnsetDefaultRoleId ensures that no value is present for DefaultRoleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


