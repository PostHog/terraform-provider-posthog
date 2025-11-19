# PatchedDashboardTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**TemplateName** | Pointer to **NullableString** |  | [optional] 
**DashboardDescription** | Pointer to **NullableString** |  | [optional] 
**DashboardFilters** | Pointer to **interface{}** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Tiles** | Pointer to **interface{}** |  | [optional] 
**Variables** | Pointer to **interface{}** |  | [optional] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**CreatedBy** | Pointer to **NullableInt32** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**TeamId** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**AvailabilityContexts** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPatchedDashboardTemplate

`func NewPatchedDashboardTemplate() *PatchedDashboardTemplate`

NewPatchedDashboardTemplate instantiates a new PatchedDashboardTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDashboardTemplateWithDefaults

`func NewPatchedDashboardTemplateWithDefaults() *PatchedDashboardTemplate`

NewPatchedDashboardTemplateWithDefaults instantiates a new PatchedDashboardTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedDashboardTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedDashboardTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedDashboardTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedDashboardTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTemplateName

`func (o *PatchedDashboardTemplate) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *PatchedDashboardTemplate) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *PatchedDashboardTemplate) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *PatchedDashboardTemplate) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### SetTemplateNameNil

`func (o *PatchedDashboardTemplate) SetTemplateNameNil(b bool)`

 SetTemplateNameNil sets the value for TemplateName to be an explicit nil

### UnsetTemplateName
`func (o *PatchedDashboardTemplate) UnsetTemplateName()`

UnsetTemplateName ensures that no value is present for TemplateName, not even an explicit nil
### GetDashboardDescription

`func (o *PatchedDashboardTemplate) GetDashboardDescription() string`

GetDashboardDescription returns the DashboardDescription field if non-nil, zero value otherwise.

### GetDashboardDescriptionOk

`func (o *PatchedDashboardTemplate) GetDashboardDescriptionOk() (*string, bool)`

GetDashboardDescriptionOk returns a tuple with the DashboardDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardDescription

`func (o *PatchedDashboardTemplate) SetDashboardDescription(v string)`

SetDashboardDescription sets DashboardDescription field to given value.

### HasDashboardDescription

`func (o *PatchedDashboardTemplate) HasDashboardDescription() bool`

HasDashboardDescription returns a boolean if a field has been set.

### SetDashboardDescriptionNil

`func (o *PatchedDashboardTemplate) SetDashboardDescriptionNil(b bool)`

 SetDashboardDescriptionNil sets the value for DashboardDescription to be an explicit nil

### UnsetDashboardDescription
`func (o *PatchedDashboardTemplate) UnsetDashboardDescription()`

UnsetDashboardDescription ensures that no value is present for DashboardDescription, not even an explicit nil
### GetDashboardFilters

`func (o *PatchedDashboardTemplate) GetDashboardFilters() interface{}`

GetDashboardFilters returns the DashboardFilters field if non-nil, zero value otherwise.

### GetDashboardFiltersOk

`func (o *PatchedDashboardTemplate) GetDashboardFiltersOk() (*interface{}, bool)`

GetDashboardFiltersOk returns a tuple with the DashboardFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardFilters

`func (o *PatchedDashboardTemplate) SetDashboardFilters(v interface{})`

SetDashboardFilters sets DashboardFilters field to given value.

### HasDashboardFilters

`func (o *PatchedDashboardTemplate) HasDashboardFilters() bool`

HasDashboardFilters returns a boolean if a field has been set.

### SetDashboardFiltersNil

`func (o *PatchedDashboardTemplate) SetDashboardFiltersNil(b bool)`

 SetDashboardFiltersNil sets the value for DashboardFilters to be an explicit nil

### UnsetDashboardFilters
`func (o *PatchedDashboardTemplate) UnsetDashboardFilters()`

UnsetDashboardFilters ensures that no value is present for DashboardFilters, not even an explicit nil
### GetTags

`func (o *PatchedDashboardTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedDashboardTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedDashboardTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedDashboardTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *PatchedDashboardTemplate) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *PatchedDashboardTemplate) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTiles

`func (o *PatchedDashboardTemplate) GetTiles() interface{}`

GetTiles returns the Tiles field if non-nil, zero value otherwise.

### GetTilesOk

`func (o *PatchedDashboardTemplate) GetTilesOk() (*interface{}, bool)`

GetTilesOk returns a tuple with the Tiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiles

`func (o *PatchedDashboardTemplate) SetTiles(v interface{})`

SetTiles sets Tiles field to given value.

### HasTiles

`func (o *PatchedDashboardTemplate) HasTiles() bool`

HasTiles returns a boolean if a field has been set.

### SetTilesNil

`func (o *PatchedDashboardTemplate) SetTilesNil(b bool)`

 SetTilesNil sets the value for Tiles to be an explicit nil

### UnsetTiles
`func (o *PatchedDashboardTemplate) UnsetTiles()`

UnsetTiles ensures that no value is present for Tiles, not even an explicit nil
### GetVariables

`func (o *PatchedDashboardTemplate) GetVariables() interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PatchedDashboardTemplate) GetVariablesOk() (*interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PatchedDashboardTemplate) SetVariables(v interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *PatchedDashboardTemplate) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *PatchedDashboardTemplate) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *PatchedDashboardTemplate) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetDeleted

`func (o *PatchedDashboardTemplate) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedDashboardTemplate) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedDashboardTemplate) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedDashboardTemplate) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *PatchedDashboardTemplate) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *PatchedDashboardTemplate) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetCreatedAt

`func (o *PatchedDashboardTemplate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedDashboardTemplate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedDashboardTemplate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedDashboardTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PatchedDashboardTemplate) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PatchedDashboardTemplate) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *PatchedDashboardTemplate) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedDashboardTemplate) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedDashboardTemplate) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedDashboardTemplate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *PatchedDashboardTemplate) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *PatchedDashboardTemplate) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetImageUrl

`func (o *PatchedDashboardTemplate) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *PatchedDashboardTemplate) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *PatchedDashboardTemplate) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *PatchedDashboardTemplate) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *PatchedDashboardTemplate) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *PatchedDashboardTemplate) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetTeamId

`func (o *PatchedDashboardTemplate) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *PatchedDashboardTemplate) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *PatchedDashboardTemplate) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *PatchedDashboardTemplate) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *PatchedDashboardTemplate) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *PatchedDashboardTemplate) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
### GetScope

`func (o *PatchedDashboardTemplate) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PatchedDashboardTemplate) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PatchedDashboardTemplate) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PatchedDashboardTemplate) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *PatchedDashboardTemplate) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *PatchedDashboardTemplate) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetAvailabilityContexts

`func (o *PatchedDashboardTemplate) GetAvailabilityContexts() []string`

GetAvailabilityContexts returns the AvailabilityContexts field if non-nil, zero value otherwise.

### GetAvailabilityContextsOk

`func (o *PatchedDashboardTemplate) GetAvailabilityContextsOk() (*[]string, bool)`

GetAvailabilityContextsOk returns a tuple with the AvailabilityContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityContexts

`func (o *PatchedDashboardTemplate) SetAvailabilityContexts(v []string)`

SetAvailabilityContexts sets AvailabilityContexts field to given value.

### HasAvailabilityContexts

`func (o *PatchedDashboardTemplate) HasAvailabilityContexts() bool`

HasAvailabilityContexts returns a boolean if a field has been set.

### SetAvailabilityContextsNil

`func (o *PatchedDashboardTemplate) SetAvailabilityContextsNil(b bool)`

 SetAvailabilityContextsNil sets the value for AvailabilityContexts to be an explicit nil

### UnsetAvailabilityContexts
`func (o *PatchedDashboardTemplate) UnsetAvailabilityContexts()`

UnsetAvailabilityContexts ensures that no value is present for AvailabilityContexts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


