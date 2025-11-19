# SessionRecordingPlaylist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**ShortId** | **string** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DerivedName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Pinned** | Pointer to **bool** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Filters** | Pointer to **interface{}** |  | [optional] 
**LastModifiedAt** | **time.Time** |  | [readonly] 
**LastModifiedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**RecordingsCounts** | [**map[string]map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue**](map.md) |  | [readonly] 
**Type** | **NullableString** |  | [readonly] 
**IsSynthetic** | **bool** | Return whether this is a synthetic playlist | [readonly] 
**CreateInFolder** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionRecordingPlaylist

`func NewSessionRecordingPlaylist(id int32, shortId string, createdAt time.Time, createdBy UserBasic, lastModifiedAt time.Time, lastModifiedBy UserBasic, recordingsCounts map[string]map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue, type_ NullableString, isSynthetic bool, ) *SessionRecordingPlaylist`

NewSessionRecordingPlaylist instantiates a new SessionRecordingPlaylist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionRecordingPlaylistWithDefaults

`func NewSessionRecordingPlaylistWithDefaults() *SessionRecordingPlaylist`

NewSessionRecordingPlaylistWithDefaults instantiates a new SessionRecordingPlaylist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionRecordingPlaylist) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionRecordingPlaylist) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionRecordingPlaylist) SetId(v int32)`

SetId sets Id field to given value.


### GetShortId

`func (o *SessionRecordingPlaylist) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *SessionRecordingPlaylist) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *SessionRecordingPlaylist) SetShortId(v string)`

SetShortId sets ShortId field to given value.


### GetName

`func (o *SessionRecordingPlaylist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionRecordingPlaylist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionRecordingPlaylist) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SessionRecordingPlaylist) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SessionRecordingPlaylist) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SessionRecordingPlaylist) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDerivedName

`func (o *SessionRecordingPlaylist) GetDerivedName() string`

GetDerivedName returns the DerivedName field if non-nil, zero value otherwise.

### GetDerivedNameOk

`func (o *SessionRecordingPlaylist) GetDerivedNameOk() (*string, bool)`

GetDerivedNameOk returns a tuple with the DerivedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedName

`func (o *SessionRecordingPlaylist) SetDerivedName(v string)`

SetDerivedName sets DerivedName field to given value.

### HasDerivedName

`func (o *SessionRecordingPlaylist) HasDerivedName() bool`

HasDerivedName returns a boolean if a field has been set.

### SetDerivedNameNil

`func (o *SessionRecordingPlaylist) SetDerivedNameNil(b bool)`

 SetDerivedNameNil sets the value for DerivedName to be an explicit nil

### UnsetDerivedName
`func (o *SessionRecordingPlaylist) UnsetDerivedName()`

UnsetDerivedName ensures that no value is present for DerivedName, not even an explicit nil
### GetDescription

`func (o *SessionRecordingPlaylist) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SessionRecordingPlaylist) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SessionRecordingPlaylist) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SessionRecordingPlaylist) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPinned

`func (o *SessionRecordingPlaylist) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *SessionRecordingPlaylist) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *SessionRecordingPlaylist) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *SessionRecordingPlaylist) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SessionRecordingPlaylist) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SessionRecordingPlaylist) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SessionRecordingPlaylist) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *SessionRecordingPlaylist) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SessionRecordingPlaylist) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SessionRecordingPlaylist) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetDeleted

`func (o *SessionRecordingPlaylist) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *SessionRecordingPlaylist) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *SessionRecordingPlaylist) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *SessionRecordingPlaylist) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetFilters

`func (o *SessionRecordingPlaylist) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SessionRecordingPlaylist) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SessionRecordingPlaylist) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SessionRecordingPlaylist) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *SessionRecordingPlaylist) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *SessionRecordingPlaylist) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetLastModifiedAt

`func (o *SessionRecordingPlaylist) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *SessionRecordingPlaylist) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *SessionRecordingPlaylist) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.


### GetLastModifiedBy

`func (o *SessionRecordingPlaylist) GetLastModifiedBy() UserBasic`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *SessionRecordingPlaylist) GetLastModifiedByOk() (*UserBasic, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *SessionRecordingPlaylist) SetLastModifiedBy(v UserBasic)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetRecordingsCounts

`func (o *SessionRecordingPlaylist) GetRecordingsCounts() map[string]map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue`

GetRecordingsCounts returns the RecordingsCounts field if non-nil, zero value otherwise.

### GetRecordingsCountsOk

`func (o *SessionRecordingPlaylist) GetRecordingsCountsOk() (*map[string]map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue, bool)`

GetRecordingsCountsOk returns a tuple with the RecordingsCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingsCounts

`func (o *SessionRecordingPlaylist) SetRecordingsCounts(v map[string]map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue)`

SetRecordingsCounts sets RecordingsCounts field to given value.


### GetType

`func (o *SessionRecordingPlaylist) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SessionRecordingPlaylist) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SessionRecordingPlaylist) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SessionRecordingPlaylist) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SessionRecordingPlaylist) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetIsSynthetic

`func (o *SessionRecordingPlaylist) GetIsSynthetic() bool`

GetIsSynthetic returns the IsSynthetic field if non-nil, zero value otherwise.

### GetIsSyntheticOk

`func (o *SessionRecordingPlaylist) GetIsSyntheticOk() (*bool, bool)`

GetIsSyntheticOk returns a tuple with the IsSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynthetic

`func (o *SessionRecordingPlaylist) SetIsSynthetic(v bool)`

SetIsSynthetic sets IsSynthetic field to given value.


### GetCreateInFolder

`func (o *SessionRecordingPlaylist) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *SessionRecordingPlaylist) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *SessionRecordingPlaylist) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *SessionRecordingPlaylist) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


