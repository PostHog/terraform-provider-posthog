# PatchedOrganizationDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Domain** | Pointer to **string** |  | [optional] 
**IsVerified** | Pointer to **bool** | Determines whether a domain is verified or not. | [optional] [readonly] 
**VerifiedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**VerificationChallenge** | Pointer to **string** |  | [optional] [readonly] 
**JitProvisioningEnabled** | Pointer to **bool** |  | [optional] 
**SsoEnforcement** | Pointer to **string** |  | [optional] 
**HasSaml** | Pointer to **bool** | Returns whether SAML is configured for the instance. Does not validate the user has the required license (that check is performed in other places). | [optional] [readonly] 
**SamlEntityId** | Pointer to **NullableString** |  | [optional] 
**SamlAcsUrl** | Pointer to **NullableString** |  | [optional] 
**SamlX509Cert** | Pointer to **NullableString** |  | [optional] 
**HasScim** | Pointer to **bool** | Returns whether SCIM is configured and enabled for this domain. | [optional] [readonly] 
**ScimEnabled** | Pointer to **bool** |  | [optional] 
**ScimBaseUrl** | Pointer to **NullableString** |  | [optional] [readonly] 
**ScimBearerToken** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewPatchedOrganizationDomain

`func NewPatchedOrganizationDomain() *PatchedOrganizationDomain`

NewPatchedOrganizationDomain instantiates a new PatchedOrganizationDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedOrganizationDomainWithDefaults

`func NewPatchedOrganizationDomainWithDefaults() *PatchedOrganizationDomain`

NewPatchedOrganizationDomainWithDefaults instantiates a new PatchedOrganizationDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedOrganizationDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedOrganizationDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedOrganizationDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedOrganizationDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *PatchedOrganizationDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PatchedOrganizationDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PatchedOrganizationDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PatchedOrganizationDomain) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetIsVerified

`func (o *PatchedOrganizationDomain) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *PatchedOrganizationDomain) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *PatchedOrganizationDomain) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *PatchedOrganizationDomain) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *PatchedOrganizationDomain) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *PatchedOrganizationDomain) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *PatchedOrganizationDomain) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *PatchedOrganizationDomain) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### SetVerifiedAtNil

`func (o *PatchedOrganizationDomain) SetVerifiedAtNil(b bool)`

 SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil

### UnsetVerifiedAt
`func (o *PatchedOrganizationDomain) UnsetVerifiedAt()`

UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil
### GetVerificationChallenge

`func (o *PatchedOrganizationDomain) GetVerificationChallenge() string`

GetVerificationChallenge returns the VerificationChallenge field if non-nil, zero value otherwise.

### GetVerificationChallengeOk

`func (o *PatchedOrganizationDomain) GetVerificationChallengeOk() (*string, bool)`

GetVerificationChallengeOk returns a tuple with the VerificationChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationChallenge

`func (o *PatchedOrganizationDomain) SetVerificationChallenge(v string)`

SetVerificationChallenge sets VerificationChallenge field to given value.

### HasVerificationChallenge

`func (o *PatchedOrganizationDomain) HasVerificationChallenge() bool`

HasVerificationChallenge returns a boolean if a field has been set.

### GetJitProvisioningEnabled

`func (o *PatchedOrganizationDomain) GetJitProvisioningEnabled() bool`

GetJitProvisioningEnabled returns the JitProvisioningEnabled field if non-nil, zero value otherwise.

### GetJitProvisioningEnabledOk

`func (o *PatchedOrganizationDomain) GetJitProvisioningEnabledOk() (*bool, bool)`

GetJitProvisioningEnabledOk returns a tuple with the JitProvisioningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitProvisioningEnabled

`func (o *PatchedOrganizationDomain) SetJitProvisioningEnabled(v bool)`

SetJitProvisioningEnabled sets JitProvisioningEnabled field to given value.

### HasJitProvisioningEnabled

`func (o *PatchedOrganizationDomain) HasJitProvisioningEnabled() bool`

HasJitProvisioningEnabled returns a boolean if a field has been set.

### GetSsoEnforcement

`func (o *PatchedOrganizationDomain) GetSsoEnforcement() string`

GetSsoEnforcement returns the SsoEnforcement field if non-nil, zero value otherwise.

### GetSsoEnforcementOk

`func (o *PatchedOrganizationDomain) GetSsoEnforcementOk() (*string, bool)`

GetSsoEnforcementOk returns a tuple with the SsoEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEnforcement

`func (o *PatchedOrganizationDomain) SetSsoEnforcement(v string)`

SetSsoEnforcement sets SsoEnforcement field to given value.

### HasSsoEnforcement

`func (o *PatchedOrganizationDomain) HasSsoEnforcement() bool`

HasSsoEnforcement returns a boolean if a field has been set.

### GetHasSaml

`func (o *PatchedOrganizationDomain) GetHasSaml() bool`

GetHasSaml returns the HasSaml field if non-nil, zero value otherwise.

### GetHasSamlOk

`func (o *PatchedOrganizationDomain) GetHasSamlOk() (*bool, bool)`

GetHasSamlOk returns a tuple with the HasSaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSaml

`func (o *PatchedOrganizationDomain) SetHasSaml(v bool)`

SetHasSaml sets HasSaml field to given value.

### HasHasSaml

`func (o *PatchedOrganizationDomain) HasHasSaml() bool`

HasHasSaml returns a boolean if a field has been set.

### GetSamlEntityId

`func (o *PatchedOrganizationDomain) GetSamlEntityId() string`

GetSamlEntityId returns the SamlEntityId field if non-nil, zero value otherwise.

### GetSamlEntityIdOk

`func (o *PatchedOrganizationDomain) GetSamlEntityIdOk() (*string, bool)`

GetSamlEntityIdOk returns a tuple with the SamlEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlEntityId

`func (o *PatchedOrganizationDomain) SetSamlEntityId(v string)`

SetSamlEntityId sets SamlEntityId field to given value.

### HasSamlEntityId

`func (o *PatchedOrganizationDomain) HasSamlEntityId() bool`

HasSamlEntityId returns a boolean if a field has been set.

### SetSamlEntityIdNil

`func (o *PatchedOrganizationDomain) SetSamlEntityIdNil(b bool)`

 SetSamlEntityIdNil sets the value for SamlEntityId to be an explicit nil

### UnsetSamlEntityId
`func (o *PatchedOrganizationDomain) UnsetSamlEntityId()`

UnsetSamlEntityId ensures that no value is present for SamlEntityId, not even an explicit nil
### GetSamlAcsUrl

`func (o *PatchedOrganizationDomain) GetSamlAcsUrl() string`

GetSamlAcsUrl returns the SamlAcsUrl field if non-nil, zero value otherwise.

### GetSamlAcsUrlOk

`func (o *PatchedOrganizationDomain) GetSamlAcsUrlOk() (*string, bool)`

GetSamlAcsUrlOk returns a tuple with the SamlAcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAcsUrl

`func (o *PatchedOrganizationDomain) SetSamlAcsUrl(v string)`

SetSamlAcsUrl sets SamlAcsUrl field to given value.

### HasSamlAcsUrl

`func (o *PatchedOrganizationDomain) HasSamlAcsUrl() bool`

HasSamlAcsUrl returns a boolean if a field has been set.

### SetSamlAcsUrlNil

`func (o *PatchedOrganizationDomain) SetSamlAcsUrlNil(b bool)`

 SetSamlAcsUrlNil sets the value for SamlAcsUrl to be an explicit nil

### UnsetSamlAcsUrl
`func (o *PatchedOrganizationDomain) UnsetSamlAcsUrl()`

UnsetSamlAcsUrl ensures that no value is present for SamlAcsUrl, not even an explicit nil
### GetSamlX509Cert

`func (o *PatchedOrganizationDomain) GetSamlX509Cert() string`

GetSamlX509Cert returns the SamlX509Cert field if non-nil, zero value otherwise.

### GetSamlX509CertOk

`func (o *PatchedOrganizationDomain) GetSamlX509CertOk() (*string, bool)`

GetSamlX509CertOk returns a tuple with the SamlX509Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlX509Cert

`func (o *PatchedOrganizationDomain) SetSamlX509Cert(v string)`

SetSamlX509Cert sets SamlX509Cert field to given value.

### HasSamlX509Cert

`func (o *PatchedOrganizationDomain) HasSamlX509Cert() bool`

HasSamlX509Cert returns a boolean if a field has been set.

### SetSamlX509CertNil

`func (o *PatchedOrganizationDomain) SetSamlX509CertNil(b bool)`

 SetSamlX509CertNil sets the value for SamlX509Cert to be an explicit nil

### UnsetSamlX509Cert
`func (o *PatchedOrganizationDomain) UnsetSamlX509Cert()`

UnsetSamlX509Cert ensures that no value is present for SamlX509Cert, not even an explicit nil
### GetHasScim

`func (o *PatchedOrganizationDomain) GetHasScim() bool`

GetHasScim returns the HasScim field if non-nil, zero value otherwise.

### GetHasScimOk

`func (o *PatchedOrganizationDomain) GetHasScimOk() (*bool, bool)`

GetHasScimOk returns a tuple with the HasScim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasScim

`func (o *PatchedOrganizationDomain) SetHasScim(v bool)`

SetHasScim sets HasScim field to given value.

### HasHasScim

`func (o *PatchedOrganizationDomain) HasHasScim() bool`

HasHasScim returns a boolean if a field has been set.

### GetScimEnabled

`func (o *PatchedOrganizationDomain) GetScimEnabled() bool`

GetScimEnabled returns the ScimEnabled field if non-nil, zero value otherwise.

### GetScimEnabledOk

`func (o *PatchedOrganizationDomain) GetScimEnabledOk() (*bool, bool)`

GetScimEnabledOk returns a tuple with the ScimEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimEnabled

`func (o *PatchedOrganizationDomain) SetScimEnabled(v bool)`

SetScimEnabled sets ScimEnabled field to given value.

### HasScimEnabled

`func (o *PatchedOrganizationDomain) HasScimEnabled() bool`

HasScimEnabled returns a boolean if a field has been set.

### GetScimBaseUrl

`func (o *PatchedOrganizationDomain) GetScimBaseUrl() string`

GetScimBaseUrl returns the ScimBaseUrl field if non-nil, zero value otherwise.

### GetScimBaseUrlOk

`func (o *PatchedOrganizationDomain) GetScimBaseUrlOk() (*string, bool)`

GetScimBaseUrlOk returns a tuple with the ScimBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimBaseUrl

`func (o *PatchedOrganizationDomain) SetScimBaseUrl(v string)`

SetScimBaseUrl sets ScimBaseUrl field to given value.

### HasScimBaseUrl

`func (o *PatchedOrganizationDomain) HasScimBaseUrl() bool`

HasScimBaseUrl returns a boolean if a field has been set.

### SetScimBaseUrlNil

`func (o *PatchedOrganizationDomain) SetScimBaseUrlNil(b bool)`

 SetScimBaseUrlNil sets the value for ScimBaseUrl to be an explicit nil

### UnsetScimBaseUrl
`func (o *PatchedOrganizationDomain) UnsetScimBaseUrl()`

UnsetScimBaseUrl ensures that no value is present for ScimBaseUrl, not even an explicit nil
### GetScimBearerToken

`func (o *PatchedOrganizationDomain) GetScimBearerToken() string`

GetScimBearerToken returns the ScimBearerToken field if non-nil, zero value otherwise.

### GetScimBearerTokenOk

`func (o *PatchedOrganizationDomain) GetScimBearerTokenOk() (*string, bool)`

GetScimBearerTokenOk returns a tuple with the ScimBearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimBearerToken

`func (o *PatchedOrganizationDomain) SetScimBearerToken(v string)`

SetScimBearerToken sets ScimBearerToken field to given value.

### HasScimBearerToken

`func (o *PatchedOrganizationDomain) HasScimBearerToken() bool`

HasScimBearerToken returns a boolean if a field has been set.

### SetScimBearerTokenNil

`func (o *PatchedOrganizationDomain) SetScimBearerTokenNil(b bool)`

 SetScimBearerTokenNil sets the value for ScimBearerToken to be an explicit nil

### UnsetScimBearerToken
`func (o *PatchedOrganizationDomain) UnsetScimBearerToken()`

UnsetScimBearerToken ensures that no value is present for ScimBearerToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


