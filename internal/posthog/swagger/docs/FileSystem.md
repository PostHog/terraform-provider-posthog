# FileSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Path** | **string** |  | 
**Depth** | **NullableInt32** |  | [readonly] 
**Type** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to **NullableString** |  | [optional] 
**Href** | Pointer to **NullableString** |  | [optional] 
**Meta** | Pointer to **interface{}** |  | [optional] 
**Shortcut** | Pointer to **NullableBool** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**LastViewedAt** | **NullableTime** |  | [readonly] 

## Methods

### NewFileSystem

`func NewFileSystem(id string, path string, depth NullableInt32, createdAt time.Time, lastViewedAt NullableTime, ) *FileSystem`

NewFileSystem instantiates a new FileSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSystemWithDefaults

`func NewFileSystemWithDefaults() *FileSystem`

NewFileSystemWithDefaults instantiates a new FileSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileSystem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileSystem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileSystem) SetId(v string)`

SetId sets Id field to given value.


### GetPath

`func (o *FileSystem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileSystem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileSystem) SetPath(v string)`

SetPath sets Path field to given value.


### GetDepth

`func (o *FileSystem) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *FileSystem) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *FileSystem) SetDepth(v int32)`

SetDepth sets Depth field to given value.


### SetDepthNil

`func (o *FileSystem) SetDepthNil(b bool)`

 SetDepthNil sets the value for Depth to be an explicit nil

### UnsetDepth
`func (o *FileSystem) UnsetDepth()`

UnsetDepth ensures that no value is present for Depth, not even an explicit nil
### GetType

`func (o *FileSystem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileSystem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileSystem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FileSystem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRef

`func (o *FileSystem) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *FileSystem) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *FileSystem) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *FileSystem) HasRef() bool`

HasRef returns a boolean if a field has been set.

### SetRefNil

`func (o *FileSystem) SetRefNil(b bool)`

 SetRefNil sets the value for Ref to be an explicit nil

### UnsetRef
`func (o *FileSystem) UnsetRef()`

UnsetRef ensures that no value is present for Ref, not even an explicit nil
### GetHref

`func (o *FileSystem) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FileSystem) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FileSystem) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FileSystem) HasHref() bool`

HasHref returns a boolean if a field has been set.

### SetHrefNil

`func (o *FileSystem) SetHrefNil(b bool)`

 SetHrefNil sets the value for Href to be an explicit nil

### UnsetHref
`func (o *FileSystem) UnsetHref()`

UnsetHref ensures that no value is present for Href, not even an explicit nil
### GetMeta

`func (o *FileSystem) GetMeta() interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FileSystem) GetMetaOk() (*interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FileSystem) SetMeta(v interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FileSystem) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *FileSystem) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *FileSystem) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetShortcut

`func (o *FileSystem) GetShortcut() bool`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *FileSystem) GetShortcutOk() (*bool, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *FileSystem) SetShortcut(v bool)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *FileSystem) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### SetShortcutNil

`func (o *FileSystem) SetShortcutNil(b bool)`

 SetShortcutNil sets the value for Shortcut to be an explicit nil

### UnsetShortcut
`func (o *FileSystem) UnsetShortcut()`

UnsetShortcut ensures that no value is present for Shortcut, not even an explicit nil
### GetCreatedAt

`func (o *FileSystem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileSystem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileSystem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastViewedAt

`func (o *FileSystem) GetLastViewedAt() time.Time`

GetLastViewedAt returns the LastViewedAt field if non-nil, zero value otherwise.

### GetLastViewedAtOk

`func (o *FileSystem) GetLastViewedAtOk() (*time.Time, bool)`

GetLastViewedAtOk returns a tuple with the LastViewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastViewedAt

`func (o *FileSystem) SetLastViewedAt(v time.Time)`

SetLastViewedAt sets LastViewedAt field to given value.


### SetLastViewedAtNil

`func (o *FileSystem) SetLastViewedAtNil(b bool)`

 SetLastViewedAtNil sets the value for LastViewedAt to be an explicit nil

### UnsetLastViewedAt
`func (o *FileSystem) UnsetLastViewedAt()`

UnsetLastViewedAt ensures that no value is present for LastViewedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


