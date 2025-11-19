# PatchedPersistedFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to [**PersistedFolderTypeEnum**](PersistedFolderTypeEnum.md) |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedPersistedFolder

`func NewPatchedPersistedFolder() *PatchedPersistedFolder`

NewPatchedPersistedFolder instantiates a new PatchedPersistedFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPersistedFolderWithDefaults

`func NewPatchedPersistedFolderWithDefaults() *PatchedPersistedFolder`

NewPatchedPersistedFolderWithDefaults instantiates a new PatchedPersistedFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedPersistedFolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedPersistedFolder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedPersistedFolder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedPersistedFolder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PatchedPersistedFolder) GetType() PersistedFolderTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedPersistedFolder) GetTypeOk() (*PersistedFolderTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedPersistedFolder) SetType(v PersistedFolderTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedPersistedFolder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProtocol

`func (o *PatchedPersistedFolder) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedPersistedFolder) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedPersistedFolder) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedPersistedFolder) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPath

`func (o *PatchedPersistedFolder) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchedPersistedFolder) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchedPersistedFolder) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PatchedPersistedFolder) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedPersistedFolder) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedPersistedFolder) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedPersistedFolder) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedPersistedFolder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedPersistedFolder) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedPersistedFolder) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedPersistedFolder) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedPersistedFolder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


