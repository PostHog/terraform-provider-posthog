# OrganizationDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Domain** | **string** |  | 
**IsVerified** | **bool** | Determines whether a domain is verified or not. | [readonly] 
**VerifiedAt** | **NullableTime** |  | [readonly] 
**VerificationChallenge** | **string** |  | [readonly] 
**JitProvisioningEnabled** | Pointer to **bool** |  | [optional] 
**SsoEnforcement** | Pointer to **string** |  | [optional] 
**HasSaml** | **bool** | Returns whether SAML is configured for the instance. Does not validate the user has the required license (that check is performed in other places). | [readonly] 
**SamlEntityId** | Pointer to **NullableString** |  | [optional] 
**SamlAcsUrl** | Pointer to **NullableString** |  | [optional] 
**SamlX509Cert** | Pointer to **NullableString** |  | [optional] 
**HasScim** | **bool** | Returns whether SCIM is configured and enabled for this domain. | [readonly] 
**ScimEnabled** | Pointer to **bool** |  | [optional] 
**ScimBaseUrl** | **NullableString** |  | [readonly] 
**ScimBearerToken** | **NullableString** |  | [readonly] 

## Methods

### NewOrganizationDomain

`func NewOrganizationDomain(id string, domain string, isVerified bool, verifiedAt NullableTime, verificationChallenge string, hasSaml bool, hasScim bool, scimBaseUrl NullableString, scimBearerToken NullableString, ) *OrganizationDomain`

NewOrganizationDomain instantiates a new OrganizationDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDomainWithDefaults

`func NewOrganizationDomainWithDefaults() *OrganizationDomain`

NewOrganizationDomainWithDefaults instantiates a new OrganizationDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationDomain) SetId(v string)`

SetId sets Id field to given value.


### GetDomain

`func (o *OrganizationDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OrganizationDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OrganizationDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetIsVerified

`func (o *OrganizationDomain) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *OrganizationDomain) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *OrganizationDomain) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.


### GetVerifiedAt

`func (o *OrganizationDomain) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *OrganizationDomain) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *OrganizationDomain) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.


### SetVerifiedAtNil

`func (o *OrganizationDomain) SetVerifiedAtNil(b bool)`

 SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil

### UnsetVerifiedAt
`func (o *OrganizationDomain) UnsetVerifiedAt()`

UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil
### GetVerificationChallenge

`func (o *OrganizationDomain) GetVerificationChallenge() string`

GetVerificationChallenge returns the VerificationChallenge field if non-nil, zero value otherwise.

### GetVerificationChallengeOk

`func (o *OrganizationDomain) GetVerificationChallengeOk() (*string, bool)`

GetVerificationChallengeOk returns a tuple with the VerificationChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationChallenge

`func (o *OrganizationDomain) SetVerificationChallenge(v string)`

SetVerificationChallenge sets VerificationChallenge field to given value.


### GetJitProvisioningEnabled

`func (o *OrganizationDomain) GetJitProvisioningEnabled() bool`

GetJitProvisioningEnabled returns the JitProvisioningEnabled field if non-nil, zero value otherwise.

### GetJitProvisioningEnabledOk

`func (o *OrganizationDomain) GetJitProvisioningEnabledOk() (*bool, bool)`

GetJitProvisioningEnabledOk returns a tuple with the JitProvisioningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitProvisioningEnabled

`func (o *OrganizationDomain) SetJitProvisioningEnabled(v bool)`

SetJitProvisioningEnabled sets JitProvisioningEnabled field to given value.

### HasJitProvisioningEnabled

`func (o *OrganizationDomain) HasJitProvisioningEnabled() bool`

HasJitProvisioningEnabled returns a boolean if a field has been set.

### GetSsoEnforcement

`func (o *OrganizationDomain) GetSsoEnforcement() string`

GetSsoEnforcement returns the SsoEnforcement field if non-nil, zero value otherwise.

### GetSsoEnforcementOk

`func (o *OrganizationDomain) GetSsoEnforcementOk() (*string, bool)`

GetSsoEnforcementOk returns a tuple with the SsoEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEnforcement

`func (o *OrganizationDomain) SetSsoEnforcement(v string)`

SetSsoEnforcement sets SsoEnforcement field to given value.

### HasSsoEnforcement

`func (o *OrganizationDomain) HasSsoEnforcement() bool`

HasSsoEnforcement returns a boolean if a field has been set.

### GetHasSaml

`func (o *OrganizationDomain) GetHasSaml() bool`

GetHasSaml returns the HasSaml field if non-nil, zero value otherwise.

### GetHasSamlOk

`func (o *OrganizationDomain) GetHasSamlOk() (*bool, bool)`

GetHasSamlOk returns a tuple with the HasSaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSaml

`func (o *OrganizationDomain) SetHasSaml(v bool)`

SetHasSaml sets HasSaml field to given value.


### GetSamlEntityId

`func (o *OrganizationDomain) GetSamlEntityId() string`

GetSamlEntityId returns the SamlEntityId field if non-nil, zero value otherwise.

### GetSamlEntityIdOk

`func (o *OrganizationDomain) GetSamlEntityIdOk() (*string, bool)`

GetSamlEntityIdOk returns a tuple with the SamlEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlEntityId

`func (o *OrganizationDomain) SetSamlEntityId(v string)`

SetSamlEntityId sets SamlEntityId field to given value.

### HasSamlEntityId

`func (o *OrganizationDomain) HasSamlEntityId() bool`

HasSamlEntityId returns a boolean if a field has been set.

### SetSamlEntityIdNil

`func (o *OrganizationDomain) SetSamlEntityIdNil(b bool)`

 SetSamlEntityIdNil sets the value for SamlEntityId to be an explicit nil

### UnsetSamlEntityId
`func (o *OrganizationDomain) UnsetSamlEntityId()`

UnsetSamlEntityId ensures that no value is present for SamlEntityId, not even an explicit nil
### GetSamlAcsUrl

`func (o *OrganizationDomain) GetSamlAcsUrl() string`

GetSamlAcsUrl returns the SamlAcsUrl field if non-nil, zero value otherwise.

### GetSamlAcsUrlOk

`func (o *OrganizationDomain) GetSamlAcsUrlOk() (*string, bool)`

GetSamlAcsUrlOk returns a tuple with the SamlAcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAcsUrl

`func (o *OrganizationDomain) SetSamlAcsUrl(v string)`

SetSamlAcsUrl sets SamlAcsUrl field to given value.

### HasSamlAcsUrl

`func (o *OrganizationDomain) HasSamlAcsUrl() bool`

HasSamlAcsUrl returns a boolean if a field has been set.

### SetSamlAcsUrlNil

`func (o *OrganizationDomain) SetSamlAcsUrlNil(b bool)`

 SetSamlAcsUrlNil sets the value for SamlAcsUrl to be an explicit nil

### UnsetSamlAcsUrl
`func (o *OrganizationDomain) UnsetSamlAcsUrl()`

UnsetSamlAcsUrl ensures that no value is present for SamlAcsUrl, not even an explicit nil
### GetSamlX509Cert

`func (o *OrganizationDomain) GetSamlX509Cert() string`

GetSamlX509Cert returns the SamlX509Cert field if non-nil, zero value otherwise.

### GetSamlX509CertOk

`func (o *OrganizationDomain) GetSamlX509CertOk() (*string, bool)`

GetSamlX509CertOk returns a tuple with the SamlX509Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlX509Cert

`func (o *OrganizationDomain) SetSamlX509Cert(v string)`

SetSamlX509Cert sets SamlX509Cert field to given value.

### HasSamlX509Cert

`func (o *OrganizationDomain) HasSamlX509Cert() bool`

HasSamlX509Cert returns a boolean if a field has been set.

### SetSamlX509CertNil

`func (o *OrganizationDomain) SetSamlX509CertNil(b bool)`

 SetSamlX509CertNil sets the value for SamlX509Cert to be an explicit nil

### UnsetSamlX509Cert
`func (o *OrganizationDomain) UnsetSamlX509Cert()`

UnsetSamlX509Cert ensures that no value is present for SamlX509Cert, not even an explicit nil
### GetHasScim

`func (o *OrganizationDomain) GetHasScim() bool`

GetHasScim returns the HasScim field if non-nil, zero value otherwise.

### GetHasScimOk

`func (o *OrganizationDomain) GetHasScimOk() (*bool, bool)`

GetHasScimOk returns a tuple with the HasScim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasScim

`func (o *OrganizationDomain) SetHasScim(v bool)`

SetHasScim sets HasScim field to given value.


### GetScimEnabled

`func (o *OrganizationDomain) GetScimEnabled() bool`

GetScimEnabled returns the ScimEnabled field if non-nil, zero value otherwise.

### GetScimEnabledOk

`func (o *OrganizationDomain) GetScimEnabledOk() (*bool, bool)`

GetScimEnabledOk returns a tuple with the ScimEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimEnabled

`func (o *OrganizationDomain) SetScimEnabled(v bool)`

SetScimEnabled sets ScimEnabled field to given value.

### HasScimEnabled

`func (o *OrganizationDomain) HasScimEnabled() bool`

HasScimEnabled returns a boolean if a field has been set.

### GetScimBaseUrl

`func (o *OrganizationDomain) GetScimBaseUrl() string`

GetScimBaseUrl returns the ScimBaseUrl field if non-nil, zero value otherwise.

### GetScimBaseUrlOk

`func (o *OrganizationDomain) GetScimBaseUrlOk() (*string, bool)`

GetScimBaseUrlOk returns a tuple with the ScimBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimBaseUrl

`func (o *OrganizationDomain) SetScimBaseUrl(v string)`

SetScimBaseUrl sets ScimBaseUrl field to given value.


### SetScimBaseUrlNil

`func (o *OrganizationDomain) SetScimBaseUrlNil(b bool)`

 SetScimBaseUrlNil sets the value for ScimBaseUrl to be an explicit nil

### UnsetScimBaseUrl
`func (o *OrganizationDomain) UnsetScimBaseUrl()`

UnsetScimBaseUrl ensures that no value is present for ScimBaseUrl, not even an explicit nil
### GetScimBearerToken

`func (o *OrganizationDomain) GetScimBearerToken() string`

GetScimBearerToken returns the ScimBearerToken field if non-nil, zero value otherwise.

### GetScimBearerTokenOk

`func (o *OrganizationDomain) GetScimBearerTokenOk() (*string, bool)`

GetScimBearerTokenOk returns a tuple with the ScimBearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimBearerToken

`func (o *OrganizationDomain) SetScimBearerToken(v string)`

SetScimBearerToken sets ScimBearerToken field to given value.


### SetScimBearerTokenNil

`func (o *OrganizationDomain) SetScimBearerTokenNil(b bool)`

 SetScimBearerTokenNil sets the value for ScimBearerToken to be an explicit nil

### UnsetScimBearerToken
`func (o *OrganizationDomain) UnsetScimBearerToken()`

UnsetScimBearerToken ensures that no value is present for ScimBearerToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


