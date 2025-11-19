# TeamBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Uuid** | **string** |  | [readonly] 
**Organization** | **string** |  | [readonly] 
**ProjectId** | **int64** |  | [readonly] 
**ApiToken** | **string** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**CompletedSnippetOnboarding** | **bool** |  | [readonly] 
**HasCompletedOnboardingFor** | **interface{}** |  | [readonly] 
**IngestedEvent** | **bool** |  | [readonly] 
**IsDemo** | **bool** |  | [readonly] 
**Timezone** | [**TimezoneEnum**](TimezoneEnum.md) |  | [readonly] 
**AccessControl** | **bool** |  | [readonly] 

## Methods

### NewTeamBasic

`func NewTeamBasic(id int32, uuid string, organization string, projectId int64, apiToken string, name string, completedSnippetOnboarding bool, hasCompletedOnboardingFor interface{}, ingestedEvent bool, isDemo bool, timezone TimezoneEnum, accessControl bool, ) *TeamBasic`

NewTeamBasic instantiates a new TeamBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamBasicWithDefaults

`func NewTeamBasicWithDefaults() *TeamBasic`

NewTeamBasicWithDefaults instantiates a new TeamBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamBasic) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamBasic) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamBasic) SetId(v int32)`

SetId sets Id field to given value.


### GetUuid

`func (o *TeamBasic) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TeamBasic) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TeamBasic) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetOrganization

`func (o *TeamBasic) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *TeamBasic) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *TeamBasic) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetProjectId

`func (o *TeamBasic) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TeamBasic) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TeamBasic) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetApiToken

`func (o *TeamBasic) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *TeamBasic) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *TeamBasic) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetName

`func (o *TeamBasic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamBasic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamBasic) SetName(v string)`

SetName sets Name field to given value.


### GetCompletedSnippetOnboarding

`func (o *TeamBasic) GetCompletedSnippetOnboarding() bool`

GetCompletedSnippetOnboarding returns the CompletedSnippetOnboarding field if non-nil, zero value otherwise.

### GetCompletedSnippetOnboardingOk

`func (o *TeamBasic) GetCompletedSnippetOnboardingOk() (*bool, bool)`

GetCompletedSnippetOnboardingOk returns a tuple with the CompletedSnippetOnboarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedSnippetOnboarding

`func (o *TeamBasic) SetCompletedSnippetOnboarding(v bool)`

SetCompletedSnippetOnboarding sets CompletedSnippetOnboarding field to given value.


### GetHasCompletedOnboardingFor

`func (o *TeamBasic) GetHasCompletedOnboardingFor() interface{}`

GetHasCompletedOnboardingFor returns the HasCompletedOnboardingFor field if non-nil, zero value otherwise.

### GetHasCompletedOnboardingForOk

`func (o *TeamBasic) GetHasCompletedOnboardingForOk() (*interface{}, bool)`

GetHasCompletedOnboardingForOk returns a tuple with the HasCompletedOnboardingFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCompletedOnboardingFor

`func (o *TeamBasic) SetHasCompletedOnboardingFor(v interface{})`

SetHasCompletedOnboardingFor sets HasCompletedOnboardingFor field to given value.


### SetHasCompletedOnboardingForNil

`func (o *TeamBasic) SetHasCompletedOnboardingForNil(b bool)`

 SetHasCompletedOnboardingForNil sets the value for HasCompletedOnboardingFor to be an explicit nil

### UnsetHasCompletedOnboardingFor
`func (o *TeamBasic) UnsetHasCompletedOnboardingFor()`

UnsetHasCompletedOnboardingFor ensures that no value is present for HasCompletedOnboardingFor, not even an explicit nil
### GetIngestedEvent

`func (o *TeamBasic) GetIngestedEvent() bool`

GetIngestedEvent returns the IngestedEvent field if non-nil, zero value otherwise.

### GetIngestedEventOk

`func (o *TeamBasic) GetIngestedEventOk() (*bool, bool)`

GetIngestedEventOk returns a tuple with the IngestedEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedEvent

`func (o *TeamBasic) SetIngestedEvent(v bool)`

SetIngestedEvent sets IngestedEvent field to given value.


### GetIsDemo

`func (o *TeamBasic) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *TeamBasic) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *TeamBasic) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.


### GetTimezone

`func (o *TeamBasic) GetTimezone() TimezoneEnum`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *TeamBasic) GetTimezoneOk() (*TimezoneEnum, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *TeamBasic) SetTimezone(v TimezoneEnum)`

SetTimezone sets Timezone field to given value.


### GetAccessControl

`func (o *TeamBasic) GetAccessControl() bool`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *TeamBasic) GetAccessControlOk() (*bool, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *TeamBasic) SetAccessControl(v bool)`

SetAccessControl sets AccessControl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


