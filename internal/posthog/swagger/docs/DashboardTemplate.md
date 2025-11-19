# DashboardTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**TemplateName** | Pointer to **NullableString** |  | [optional] 
**DashboardDescription** | Pointer to **NullableString** |  | [optional] 
**DashboardFilters** | Pointer to **interface{}** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Tiles** | Pointer to **interface{}** |  | [optional] 
**Variables** | Pointer to **interface{}** |  | [optional] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 
**CreatedAt** | **NullableTime** |  | [readonly] 
**CreatedBy** | Pointer to **NullableInt32** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**TeamId** | **NullableInt32** |  | [readonly] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**AvailabilityContexts** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDashboardTemplate

`func NewDashboardTemplate(id string, createdAt NullableTime, teamId NullableInt32, ) *DashboardTemplate`

NewDashboardTemplate instantiates a new DashboardTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardTemplateWithDefaults

`func NewDashboardTemplateWithDefaults() *DashboardTemplate`

NewDashboardTemplateWithDefaults instantiates a new DashboardTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetTemplateName

`func (o *DashboardTemplate) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *DashboardTemplate) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *DashboardTemplate) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *DashboardTemplate) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### SetTemplateNameNil

`func (o *DashboardTemplate) SetTemplateNameNil(b bool)`

 SetTemplateNameNil sets the value for TemplateName to be an explicit nil

### UnsetTemplateName
`func (o *DashboardTemplate) UnsetTemplateName()`

UnsetTemplateName ensures that no value is present for TemplateName, not even an explicit nil
### GetDashboardDescription

`func (o *DashboardTemplate) GetDashboardDescription() string`

GetDashboardDescription returns the DashboardDescription field if non-nil, zero value otherwise.

### GetDashboardDescriptionOk

`func (o *DashboardTemplate) GetDashboardDescriptionOk() (*string, bool)`

GetDashboardDescriptionOk returns a tuple with the DashboardDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardDescription

`func (o *DashboardTemplate) SetDashboardDescription(v string)`

SetDashboardDescription sets DashboardDescription field to given value.

### HasDashboardDescription

`func (o *DashboardTemplate) HasDashboardDescription() bool`

HasDashboardDescription returns a boolean if a field has been set.

### SetDashboardDescriptionNil

`func (o *DashboardTemplate) SetDashboardDescriptionNil(b bool)`

 SetDashboardDescriptionNil sets the value for DashboardDescription to be an explicit nil

### UnsetDashboardDescription
`func (o *DashboardTemplate) UnsetDashboardDescription()`

UnsetDashboardDescription ensures that no value is present for DashboardDescription, not even an explicit nil
### GetDashboardFilters

`func (o *DashboardTemplate) GetDashboardFilters() interface{}`

GetDashboardFilters returns the DashboardFilters field if non-nil, zero value otherwise.

### GetDashboardFiltersOk

`func (o *DashboardTemplate) GetDashboardFiltersOk() (*interface{}, bool)`

GetDashboardFiltersOk returns a tuple with the DashboardFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardFilters

`func (o *DashboardTemplate) SetDashboardFilters(v interface{})`

SetDashboardFilters sets DashboardFilters field to given value.

### HasDashboardFilters

`func (o *DashboardTemplate) HasDashboardFilters() bool`

HasDashboardFilters returns a boolean if a field has been set.

### SetDashboardFiltersNil

`func (o *DashboardTemplate) SetDashboardFiltersNil(b bool)`

 SetDashboardFiltersNil sets the value for DashboardFilters to be an explicit nil

### UnsetDashboardFilters
`func (o *DashboardTemplate) UnsetDashboardFilters()`

UnsetDashboardFilters ensures that no value is present for DashboardFilters, not even an explicit nil
### GetTags

`func (o *DashboardTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DashboardTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DashboardTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DashboardTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DashboardTemplate) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DashboardTemplate) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTiles

`func (o *DashboardTemplate) GetTiles() interface{}`

GetTiles returns the Tiles field if non-nil, zero value otherwise.

### GetTilesOk

`func (o *DashboardTemplate) GetTilesOk() (*interface{}, bool)`

GetTilesOk returns a tuple with the Tiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiles

`func (o *DashboardTemplate) SetTiles(v interface{})`

SetTiles sets Tiles field to given value.

### HasTiles

`func (o *DashboardTemplate) HasTiles() bool`

HasTiles returns a boolean if a field has been set.

### SetTilesNil

`func (o *DashboardTemplate) SetTilesNil(b bool)`

 SetTilesNil sets the value for Tiles to be an explicit nil

### UnsetTiles
`func (o *DashboardTemplate) UnsetTiles()`

UnsetTiles ensures that no value is present for Tiles, not even an explicit nil
### GetVariables

`func (o *DashboardTemplate) GetVariables() interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *DashboardTemplate) GetVariablesOk() (*interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *DashboardTemplate) SetVariables(v interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *DashboardTemplate) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *DashboardTemplate) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *DashboardTemplate) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetDeleted

`func (o *DashboardTemplate) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DashboardTemplate) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DashboardTemplate) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DashboardTemplate) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *DashboardTemplate) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *DashboardTemplate) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetCreatedAt

`func (o *DashboardTemplate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DashboardTemplate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DashboardTemplate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *DashboardTemplate) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DashboardTemplate) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *DashboardTemplate) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DashboardTemplate) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DashboardTemplate) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DashboardTemplate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *DashboardTemplate) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *DashboardTemplate) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetImageUrl

`func (o *DashboardTemplate) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *DashboardTemplate) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *DashboardTemplate) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *DashboardTemplate) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *DashboardTemplate) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *DashboardTemplate) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetTeamId

`func (o *DashboardTemplate) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *DashboardTemplate) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *DashboardTemplate) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.


### SetTeamIdNil

`func (o *DashboardTemplate) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *DashboardTemplate) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
### GetScope

`func (o *DashboardTemplate) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DashboardTemplate) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DashboardTemplate) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DashboardTemplate) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *DashboardTemplate) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *DashboardTemplate) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetAvailabilityContexts

`func (o *DashboardTemplate) GetAvailabilityContexts() []string`

GetAvailabilityContexts returns the AvailabilityContexts field if non-nil, zero value otherwise.

### GetAvailabilityContextsOk

`func (o *DashboardTemplate) GetAvailabilityContextsOk() (*[]string, bool)`

GetAvailabilityContextsOk returns a tuple with the AvailabilityContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityContexts

`func (o *DashboardTemplate) SetAvailabilityContexts(v []string)`

SetAvailabilityContexts sets AvailabilityContexts field to given value.

### HasAvailabilityContexts

`func (o *DashboardTemplate) HasAvailabilityContexts() bool`

HasAvailabilityContexts returns a boolean if a field has been set.

### SetAvailabilityContextsNil

`func (o *DashboardTemplate) SetAvailabilityContextsNil(b bool)`

 SetAvailabilityContextsNil sets the value for AvailabilityContexts to be an explicit nil

### UnsetAvailabilityContexts
`func (o *DashboardTemplate) UnsetAvailabilityContexts()`

UnsetAvailabilityContexts ensures that no value is present for AvailabilityContexts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


