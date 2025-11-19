# PatchedSessionRecordingPlaylist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**ShortId** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DerivedName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Pinned** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Filters** | Pointer to **interface{}** |  | [optional] 
**LastModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastModifiedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**RecordingsCounts** | Pointer to [**map[string]map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue**](map.md) |  | [optional] [readonly] 
**Type** | Pointer to **NullableString** |  | [optional] [readonly] 
**IsSynthetic** | Pointer to **bool** | Return whether this is a synthetic playlist | [optional] [readonly] 
**CreateInFolder** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedSessionRecordingPlaylist

`func NewPatchedSessionRecordingPlaylist() *PatchedSessionRecordingPlaylist`

NewPatchedSessionRecordingPlaylist instantiates a new PatchedSessionRecordingPlaylist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSessionRecordingPlaylistWithDefaults

`func NewPatchedSessionRecordingPlaylistWithDefaults() *PatchedSessionRecordingPlaylist`

NewPatchedSessionRecordingPlaylistWithDefaults instantiates a new PatchedSessionRecordingPlaylist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedSessionRecordingPlaylist) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedSessionRecordingPlaylist) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedSessionRecordingPlaylist) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedSessionRecordingPlaylist) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShortId

`func (o *PatchedSessionRecordingPlaylist) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *PatchedSessionRecordingPlaylist) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *PatchedSessionRecordingPlaylist) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *PatchedSessionRecordingPlaylist) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetName

`func (o *PatchedSessionRecordingPlaylist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedSessionRecordingPlaylist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedSessionRecordingPlaylist) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedSessionRecordingPlaylist) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedSessionRecordingPlaylist) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedSessionRecordingPlaylist) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDerivedName

`func (o *PatchedSessionRecordingPlaylist) GetDerivedName() string`

GetDerivedName returns the DerivedName field if non-nil, zero value otherwise.

### GetDerivedNameOk

`func (o *PatchedSessionRecordingPlaylist) GetDerivedNameOk() (*string, bool)`

GetDerivedNameOk returns a tuple with the DerivedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedName

`func (o *PatchedSessionRecordingPlaylist) SetDerivedName(v string)`

SetDerivedName sets DerivedName field to given value.

### HasDerivedName

`func (o *PatchedSessionRecordingPlaylist) HasDerivedName() bool`

HasDerivedName returns a boolean if a field has been set.

### SetDerivedNameNil

`func (o *PatchedSessionRecordingPlaylist) SetDerivedNameNil(b bool)`

 SetDerivedNameNil sets the value for DerivedName to be an explicit nil

### UnsetDerivedName
`func (o *PatchedSessionRecordingPlaylist) UnsetDerivedName()`

UnsetDerivedName ensures that no value is present for DerivedName, not even an explicit nil
### GetDescription

`func (o *PatchedSessionRecordingPlaylist) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedSessionRecordingPlaylist) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedSessionRecordingPlaylist) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedSessionRecordingPlaylist) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPinned

`func (o *PatchedSessionRecordingPlaylist) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *PatchedSessionRecordingPlaylist) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *PatchedSessionRecordingPlaylist) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *PatchedSessionRecordingPlaylist) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedSessionRecordingPlaylist) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedSessionRecordingPlaylist) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedSessionRecordingPlaylist) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedSessionRecordingPlaylist) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedSessionRecordingPlaylist) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedSessionRecordingPlaylist) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedSessionRecordingPlaylist) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedSessionRecordingPlaylist) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeleted

`func (o *PatchedSessionRecordingPlaylist) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedSessionRecordingPlaylist) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedSessionRecordingPlaylist) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedSessionRecordingPlaylist) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetFilters

`func (o *PatchedSessionRecordingPlaylist) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PatchedSessionRecordingPlaylist) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PatchedSessionRecordingPlaylist) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PatchedSessionRecordingPlaylist) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *PatchedSessionRecordingPlaylist) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *PatchedSessionRecordingPlaylist) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetLastModifiedAt

`func (o *PatchedSessionRecordingPlaylist) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *PatchedSessionRecordingPlaylist) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *PatchedSessionRecordingPlaylist) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *PatchedSessionRecordingPlaylist) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *PatchedSessionRecordingPlaylist) GetLastModifiedBy() UserBasic`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *PatchedSessionRecordingPlaylist) GetLastModifiedByOk() (*UserBasic, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *PatchedSessionRecordingPlaylist) SetLastModifiedBy(v UserBasic)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *PatchedSessionRecordingPlaylist) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetRecordingsCounts

`func (o *PatchedSessionRecordingPlaylist) GetRecordingsCounts() map[string]map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue`

GetRecordingsCounts returns the RecordingsCounts field if non-nil, zero value otherwise.

### GetRecordingsCountsOk

`func (o *PatchedSessionRecordingPlaylist) GetRecordingsCountsOk() (*map[string]map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue, bool)`

GetRecordingsCountsOk returns a tuple with the RecordingsCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingsCounts

`func (o *PatchedSessionRecordingPlaylist) SetRecordingsCounts(v map[string]map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue)`

SetRecordingsCounts sets RecordingsCounts field to given value.

### HasRecordingsCounts

`func (o *PatchedSessionRecordingPlaylist) HasRecordingsCounts() bool`

HasRecordingsCounts returns a boolean if a field has been set.

### GetType

`func (o *PatchedSessionRecordingPlaylist) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedSessionRecordingPlaylist) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedSessionRecordingPlaylist) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedSessionRecordingPlaylist) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PatchedSessionRecordingPlaylist) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PatchedSessionRecordingPlaylist) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetIsSynthetic

`func (o *PatchedSessionRecordingPlaylist) GetIsSynthetic() bool`

GetIsSynthetic returns the IsSynthetic field if non-nil, zero value otherwise.

### GetIsSyntheticOk

`func (o *PatchedSessionRecordingPlaylist) GetIsSyntheticOk() (*bool, bool)`

GetIsSyntheticOk returns a tuple with the IsSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynthetic

`func (o *PatchedSessionRecordingPlaylist) SetIsSynthetic(v bool)`

SetIsSynthetic sets IsSynthetic field to given value.

### HasIsSynthetic

`func (o *PatchedSessionRecordingPlaylist) HasIsSynthetic() bool`

HasIsSynthetic returns a boolean if a field has been set.

### GetCreateInFolder

`func (o *PatchedSessionRecordingPlaylist) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *PatchedSessionRecordingPlaylist) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *PatchedSessionRecordingPlaylist) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *PatchedSessionRecordingPlaylist) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


