# FileSystemShortcut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Path** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to **NullableString** |  | [optional] 
**Href** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 

## Methods

### NewFileSystemShortcut

`func NewFileSystemShortcut(id string, path string, createdAt time.Time, ) *FileSystemShortcut`

NewFileSystemShortcut instantiates a new FileSystemShortcut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSystemShortcutWithDefaults

`func NewFileSystemShortcutWithDefaults() *FileSystemShortcut`

NewFileSystemShortcutWithDefaults instantiates a new FileSystemShortcut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileSystemShortcut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileSystemShortcut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileSystemShortcut) SetId(v string)`

SetId sets Id field to given value.


### GetPath

`func (o *FileSystemShortcut) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileSystemShortcut) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileSystemShortcut) SetPath(v string)`

SetPath sets Path field to given value.


### GetType

`func (o *FileSystemShortcut) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileSystemShortcut) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileSystemShortcut) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FileSystemShortcut) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRef

`func (o *FileSystemShortcut) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *FileSystemShortcut) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *FileSystemShortcut) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *FileSystemShortcut) HasRef() bool`

HasRef returns a boolean if a field has been set.

### SetRefNil

`func (o *FileSystemShortcut) SetRefNil(b bool)`

 SetRefNil sets the value for Ref to be an explicit nil

### UnsetRef
`func (o *FileSystemShortcut) UnsetRef()`

UnsetRef ensures that no value is present for Ref, not even an explicit nil
### GetHref

`func (o *FileSystemShortcut) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FileSystemShortcut) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FileSystemShortcut) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FileSystemShortcut) HasHref() bool`

HasHref returns a boolean if a field has been set.

### SetHrefNil

`func (o *FileSystemShortcut) SetHrefNil(b bool)`

 SetHrefNil sets the value for Href to be an explicit nil

### UnsetHref
`func (o *FileSystemShortcut) UnsetHref()`

UnsetHref ensures that no value is present for Href, not even an explicit nil
### GetCreatedAt

`func (o *FileSystemShortcut) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileSystemShortcut) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileSystemShortcut) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


