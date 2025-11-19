# PatchedFileSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Path** | Pointer to **string** |  | [optional] 
**Depth** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to **NullableString** |  | [optional] 
**Href** | Pointer to **NullableString** |  | [optional] 
**Meta** | Pointer to **interface{}** |  | [optional] 
**Shortcut** | Pointer to **NullableBool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastViewedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewPatchedFileSystem

`func NewPatchedFileSystem() *PatchedFileSystem`

NewPatchedFileSystem instantiates a new PatchedFileSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedFileSystemWithDefaults

`func NewPatchedFileSystemWithDefaults() *PatchedFileSystem`

NewPatchedFileSystemWithDefaults instantiates a new PatchedFileSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedFileSystem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedFileSystem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedFileSystem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedFileSystem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *PatchedFileSystem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchedFileSystem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchedFileSystem) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PatchedFileSystem) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetDepth

`func (o *PatchedFileSystem) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *PatchedFileSystem) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *PatchedFileSystem) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *PatchedFileSystem) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### SetDepthNil

`func (o *PatchedFileSystem) SetDepthNil(b bool)`

 SetDepthNil sets the value for Depth to be an explicit nil

### UnsetDepth
`func (o *PatchedFileSystem) UnsetDepth()`

UnsetDepth ensures that no value is present for Depth, not even an explicit nil
### GetType

`func (o *PatchedFileSystem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedFileSystem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedFileSystem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedFileSystem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRef

`func (o *PatchedFileSystem) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PatchedFileSystem) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PatchedFileSystem) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *PatchedFileSystem) HasRef() bool`

HasRef returns a boolean if a field has been set.

### SetRefNil

`func (o *PatchedFileSystem) SetRefNil(b bool)`

 SetRefNil sets the value for Ref to be an explicit nil

### UnsetRef
`func (o *PatchedFileSystem) UnsetRef()`

UnsetRef ensures that no value is present for Ref, not even an explicit nil
### GetHref

`func (o *PatchedFileSystem) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PatchedFileSystem) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PatchedFileSystem) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PatchedFileSystem) HasHref() bool`

HasHref returns a boolean if a field has been set.

### SetHrefNil

`func (o *PatchedFileSystem) SetHrefNil(b bool)`

 SetHrefNil sets the value for Href to be an explicit nil

### UnsetHref
`func (o *PatchedFileSystem) UnsetHref()`

UnsetHref ensures that no value is present for Href, not even an explicit nil
### GetMeta

`func (o *PatchedFileSystem) GetMeta() interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PatchedFileSystem) GetMetaOk() (*interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PatchedFileSystem) SetMeta(v interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PatchedFileSystem) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *PatchedFileSystem) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *PatchedFileSystem) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetShortcut

`func (o *PatchedFileSystem) GetShortcut() bool`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *PatchedFileSystem) GetShortcutOk() (*bool, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *PatchedFileSystem) SetShortcut(v bool)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *PatchedFileSystem) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### SetShortcutNil

`func (o *PatchedFileSystem) SetShortcutNil(b bool)`

 SetShortcutNil sets the value for Shortcut to be an explicit nil

### UnsetShortcut
`func (o *PatchedFileSystem) UnsetShortcut()`

UnsetShortcut ensures that no value is present for Shortcut, not even an explicit nil
### GetCreatedAt

`func (o *PatchedFileSystem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedFileSystem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedFileSystem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedFileSystem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastViewedAt

`func (o *PatchedFileSystem) GetLastViewedAt() time.Time`

GetLastViewedAt returns the LastViewedAt field if non-nil, zero value otherwise.

### GetLastViewedAtOk

`func (o *PatchedFileSystem) GetLastViewedAtOk() (*time.Time, bool)`

GetLastViewedAtOk returns a tuple with the LastViewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastViewedAt

`func (o *PatchedFileSystem) SetLastViewedAt(v time.Time)`

SetLastViewedAt sets LastViewedAt field to given value.

### HasLastViewedAt

`func (o *PatchedFileSystem) HasLastViewedAt() bool`

HasLastViewedAt returns a boolean if a field has been set.

### SetLastViewedAtNil

`func (o *PatchedFileSystem) SetLastViewedAtNil(b bool)`

 SetLastViewedAtNil sets the value for LastViewedAt to be an explicit nil

### UnsetLastViewedAt
`func (o *PatchedFileSystem) UnsetLastViewedAt()`

UnsetLastViewedAt ensures that no value is present for LastViewedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


