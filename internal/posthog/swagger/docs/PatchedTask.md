# PatchedTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**TaskNumber** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**Slug** | Pointer to **string** |  | [optional] [readonly] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OriginProduct** | Pointer to [**OriginProductEnum**](OriginProductEnum.md) |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**GithubIntegration** | Pointer to **NullableInt32** | GitHub integration for this task | [optional] 
**RepositoryConfig** | Pointer to **interface{}** | Repository configuration with organization and repository fields | [optional] 
**RepositoryList** | Pointer to **string** |  | [optional] [readonly] 
**PrimaryRepository** | Pointer to **string** |  | [optional] [readonly] 
**LatestRun** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 

## Methods

### NewPatchedTask

`func NewPatchedTask() *PatchedTask`

NewPatchedTask instantiates a new PatchedTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTaskWithDefaults

`func NewPatchedTaskWithDefaults() *PatchedTask`

NewPatchedTaskWithDefaults instantiates a new PatchedTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTaskNumber

`func (o *PatchedTask) GetTaskNumber() int32`

GetTaskNumber returns the TaskNumber field if non-nil, zero value otherwise.

### GetTaskNumberOk

`func (o *PatchedTask) GetTaskNumberOk() (*int32, bool)`

GetTaskNumberOk returns a tuple with the TaskNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskNumber

`func (o *PatchedTask) SetTaskNumber(v int32)`

SetTaskNumber sets TaskNumber field to given value.

### HasTaskNumber

`func (o *PatchedTask) HasTaskNumber() bool`

HasTaskNumber returns a boolean if a field has been set.

### SetTaskNumberNil

`func (o *PatchedTask) SetTaskNumberNil(b bool)`

 SetTaskNumberNil sets the value for TaskNumber to be an explicit nil

### UnsetTaskNumber
`func (o *PatchedTask) UnsetTaskNumber()`

UnsetTaskNumber ensures that no value is present for TaskNumber, not even an explicit nil
### GetSlug

`func (o *PatchedTask) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedTask) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedTask) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedTask) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTitle

`func (o *PatchedTask) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PatchedTask) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PatchedTask) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PatchedTask) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOriginProduct

`func (o *PatchedTask) GetOriginProduct() OriginProductEnum`

GetOriginProduct returns the OriginProduct field if non-nil, zero value otherwise.

### GetOriginProductOk

`func (o *PatchedTask) GetOriginProductOk() (*OriginProductEnum, bool)`

GetOriginProductOk returns a tuple with the OriginProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProduct

`func (o *PatchedTask) SetOriginProduct(v OriginProductEnum)`

SetOriginProduct sets OriginProduct field to given value.

### HasOriginProduct

`func (o *PatchedTask) HasOriginProduct() bool`

HasOriginProduct returns a boolean if a field has been set.

### GetPosition

`func (o *PatchedTask) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PatchedTask) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PatchedTask) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PatchedTask) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetGithubIntegration

`func (o *PatchedTask) GetGithubIntegration() int32`

GetGithubIntegration returns the GithubIntegration field if non-nil, zero value otherwise.

### GetGithubIntegrationOk

`func (o *PatchedTask) GetGithubIntegrationOk() (*int32, bool)`

GetGithubIntegrationOk returns a tuple with the GithubIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubIntegration

`func (o *PatchedTask) SetGithubIntegration(v int32)`

SetGithubIntegration sets GithubIntegration field to given value.

### HasGithubIntegration

`func (o *PatchedTask) HasGithubIntegration() bool`

HasGithubIntegration returns a boolean if a field has been set.

### SetGithubIntegrationNil

`func (o *PatchedTask) SetGithubIntegrationNil(b bool)`

 SetGithubIntegrationNil sets the value for GithubIntegration to be an explicit nil

### UnsetGithubIntegration
`func (o *PatchedTask) UnsetGithubIntegration()`

UnsetGithubIntegration ensures that no value is present for GithubIntegration, not even an explicit nil
### GetRepositoryConfig

`func (o *PatchedTask) GetRepositoryConfig() interface{}`

GetRepositoryConfig returns the RepositoryConfig field if non-nil, zero value otherwise.

### GetRepositoryConfigOk

`func (o *PatchedTask) GetRepositoryConfigOk() (*interface{}, bool)`

GetRepositoryConfigOk returns a tuple with the RepositoryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryConfig

`func (o *PatchedTask) SetRepositoryConfig(v interface{})`

SetRepositoryConfig sets RepositoryConfig field to given value.

### HasRepositoryConfig

`func (o *PatchedTask) HasRepositoryConfig() bool`

HasRepositoryConfig returns a boolean if a field has been set.

### SetRepositoryConfigNil

`func (o *PatchedTask) SetRepositoryConfigNil(b bool)`

 SetRepositoryConfigNil sets the value for RepositoryConfig to be an explicit nil

### UnsetRepositoryConfig
`func (o *PatchedTask) UnsetRepositoryConfig()`

UnsetRepositoryConfig ensures that no value is present for RepositoryConfig, not even an explicit nil
### GetRepositoryList

`func (o *PatchedTask) GetRepositoryList() string`

GetRepositoryList returns the RepositoryList field if non-nil, zero value otherwise.

### GetRepositoryListOk

`func (o *PatchedTask) GetRepositoryListOk() (*string, bool)`

GetRepositoryListOk returns a tuple with the RepositoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryList

`func (o *PatchedTask) SetRepositoryList(v string)`

SetRepositoryList sets RepositoryList field to given value.

### HasRepositoryList

`func (o *PatchedTask) HasRepositoryList() bool`

HasRepositoryList returns a boolean if a field has been set.

### GetPrimaryRepository

`func (o *PatchedTask) GetPrimaryRepository() string`

GetPrimaryRepository returns the PrimaryRepository field if non-nil, zero value otherwise.

### GetPrimaryRepositoryOk

`func (o *PatchedTask) GetPrimaryRepositoryOk() (*string, bool)`

GetPrimaryRepositoryOk returns a tuple with the PrimaryRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRepository

`func (o *PatchedTask) SetPrimaryRepository(v string)`

SetPrimaryRepository sets PrimaryRepository field to given value.

### HasPrimaryRepository

`func (o *PatchedTask) HasPrimaryRepository() bool`

HasPrimaryRepository returns a boolean if a field has been set.

### GetLatestRun

`func (o *PatchedTask) GetLatestRun() string`

GetLatestRun returns the LatestRun field if non-nil, zero value otherwise.

### GetLatestRunOk

`func (o *PatchedTask) GetLatestRunOk() (*string, bool)`

GetLatestRunOk returns a tuple with the LatestRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRun

`func (o *PatchedTask) SetLatestRun(v string)`

SetLatestRun sets LatestRun field to given value.

### HasLatestRun

`func (o *PatchedTask) HasLatestRun() bool`

HasLatestRun returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedTask) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedTask) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedTask) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedTask) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedTask) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedTask) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedTask) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedTask) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedTask) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedTask) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedTask) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedTask) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


