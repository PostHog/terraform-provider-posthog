# PersistedFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Type** | [**PersistedFolderTypeEnum**](PersistedFolderTypeEnum.md) |  | 
**Protocol** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 

## Methods

### NewPersistedFolder

`func NewPersistedFolder(id string, type_ PersistedFolderTypeEnum, createdAt time.Time, updatedAt time.Time, ) *PersistedFolder`

NewPersistedFolder instantiates a new PersistedFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersistedFolderWithDefaults

`func NewPersistedFolderWithDefaults() *PersistedFolder`

NewPersistedFolderWithDefaults instantiates a new PersistedFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersistedFolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersistedFolder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersistedFolder) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PersistedFolder) GetType() PersistedFolderTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PersistedFolder) GetTypeOk() (*PersistedFolderTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PersistedFolder) SetType(v PersistedFolderTypeEnum)`

SetType sets Type field to given value.


### GetProtocol

`func (o *PersistedFolder) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PersistedFolder) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PersistedFolder) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PersistedFolder) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPath

`func (o *PersistedFolder) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PersistedFolder) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PersistedFolder) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PersistedFolder) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PersistedFolder) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PersistedFolder) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PersistedFolder) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PersistedFolder) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PersistedFolder) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PersistedFolder) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


