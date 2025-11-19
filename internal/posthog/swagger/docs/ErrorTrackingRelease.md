# ErrorTrackingRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**HashId** | **string** |  | 
**TeamId** | **int32** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**Metadata** | Pointer to **interface{}** |  | [optional] 
**Version** | **string** |  | 
**Project** | **string** |  | 

## Methods

### NewErrorTrackingRelease

`func NewErrorTrackingRelease(id string, hashId string, teamId int32, createdAt time.Time, version string, project string, ) *ErrorTrackingRelease`

NewErrorTrackingRelease instantiates a new ErrorTrackingRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTrackingReleaseWithDefaults

`func NewErrorTrackingReleaseWithDefaults() *ErrorTrackingRelease`

NewErrorTrackingReleaseWithDefaults instantiates a new ErrorTrackingRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ErrorTrackingRelease) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorTrackingRelease) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorTrackingRelease) SetId(v string)`

SetId sets Id field to given value.


### GetHashId

`func (o *ErrorTrackingRelease) GetHashId() string`

GetHashId returns the HashId field if non-nil, zero value otherwise.

### GetHashIdOk

`func (o *ErrorTrackingRelease) GetHashIdOk() (*string, bool)`

GetHashIdOk returns a tuple with the HashId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashId

`func (o *ErrorTrackingRelease) SetHashId(v string)`

SetHashId sets HashId field to given value.


### GetTeamId

`func (o *ErrorTrackingRelease) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ErrorTrackingRelease) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ErrorTrackingRelease) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.


### GetCreatedAt

`func (o *ErrorTrackingRelease) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ErrorTrackingRelease) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ErrorTrackingRelease) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMetadata

`func (o *ErrorTrackingRelease) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ErrorTrackingRelease) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ErrorTrackingRelease) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ErrorTrackingRelease) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ErrorTrackingRelease) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ErrorTrackingRelease) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetVersion

`func (o *ErrorTrackingRelease) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ErrorTrackingRelease) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ErrorTrackingRelease) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetProject

`func (o *ErrorTrackingRelease) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ErrorTrackingRelease) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ErrorTrackingRelease) SetProject(v string)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


