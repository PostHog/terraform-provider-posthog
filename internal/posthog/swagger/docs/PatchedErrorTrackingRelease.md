# PatchedErrorTrackingRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**HashId** | Pointer to **string** |  | [optional] 
**TeamId** | Pointer to **int32** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedErrorTrackingRelease

`func NewPatchedErrorTrackingRelease() *PatchedErrorTrackingRelease`

NewPatchedErrorTrackingRelease instantiates a new PatchedErrorTrackingRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedErrorTrackingReleaseWithDefaults

`func NewPatchedErrorTrackingReleaseWithDefaults() *PatchedErrorTrackingRelease`

NewPatchedErrorTrackingReleaseWithDefaults instantiates a new PatchedErrorTrackingRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedErrorTrackingRelease) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedErrorTrackingRelease) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedErrorTrackingRelease) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedErrorTrackingRelease) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHashId

`func (o *PatchedErrorTrackingRelease) GetHashId() string`

GetHashId returns the HashId field if non-nil, zero value otherwise.

### GetHashIdOk

`func (o *PatchedErrorTrackingRelease) GetHashIdOk() (*string, bool)`

GetHashIdOk returns a tuple with the HashId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashId

`func (o *PatchedErrorTrackingRelease) SetHashId(v string)`

SetHashId sets HashId field to given value.

### HasHashId

`func (o *PatchedErrorTrackingRelease) HasHashId() bool`

HasHashId returns a boolean if a field has been set.

### GetTeamId

`func (o *PatchedErrorTrackingRelease) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *PatchedErrorTrackingRelease) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *PatchedErrorTrackingRelease) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *PatchedErrorTrackingRelease) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedErrorTrackingRelease) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedErrorTrackingRelease) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedErrorTrackingRelease) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedErrorTrackingRelease) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchedErrorTrackingRelease) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchedErrorTrackingRelease) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchedErrorTrackingRelease) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchedErrorTrackingRelease) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PatchedErrorTrackingRelease) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PatchedErrorTrackingRelease) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetVersion

`func (o *PatchedErrorTrackingRelease) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchedErrorTrackingRelease) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchedErrorTrackingRelease) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchedErrorTrackingRelease) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetProject

`func (o *PatchedErrorTrackingRelease) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatchedErrorTrackingRelease) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatchedErrorTrackingRelease) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatchedErrorTrackingRelease) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


