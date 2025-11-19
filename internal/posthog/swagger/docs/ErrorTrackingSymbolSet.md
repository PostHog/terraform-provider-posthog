# ErrorTrackingSymbolSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Ref** | **string** |  | 
**TeamId** | **int32** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**StoragePtr** | Pointer to **NullableString** |  | [optional] 
**FailureReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewErrorTrackingSymbolSet

`func NewErrorTrackingSymbolSet(id string, ref string, teamId int32, createdAt time.Time, ) *ErrorTrackingSymbolSet`

NewErrorTrackingSymbolSet instantiates a new ErrorTrackingSymbolSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTrackingSymbolSetWithDefaults

`func NewErrorTrackingSymbolSetWithDefaults() *ErrorTrackingSymbolSet`

NewErrorTrackingSymbolSetWithDefaults instantiates a new ErrorTrackingSymbolSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ErrorTrackingSymbolSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorTrackingSymbolSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorTrackingSymbolSet) SetId(v string)`

SetId sets Id field to given value.


### GetRef

`func (o *ErrorTrackingSymbolSet) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ErrorTrackingSymbolSet) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ErrorTrackingSymbolSet) SetRef(v string)`

SetRef sets Ref field to given value.


### GetTeamId

`func (o *ErrorTrackingSymbolSet) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ErrorTrackingSymbolSet) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ErrorTrackingSymbolSet) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.


### GetCreatedAt

`func (o *ErrorTrackingSymbolSet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ErrorTrackingSymbolSet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ErrorTrackingSymbolSet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStoragePtr

`func (o *ErrorTrackingSymbolSet) GetStoragePtr() string`

GetStoragePtr returns the StoragePtr field if non-nil, zero value otherwise.

### GetStoragePtrOk

`func (o *ErrorTrackingSymbolSet) GetStoragePtrOk() (*string, bool)`

GetStoragePtrOk returns a tuple with the StoragePtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePtr

`func (o *ErrorTrackingSymbolSet) SetStoragePtr(v string)`

SetStoragePtr sets StoragePtr field to given value.

### HasStoragePtr

`func (o *ErrorTrackingSymbolSet) HasStoragePtr() bool`

HasStoragePtr returns a boolean if a field has been set.

### SetStoragePtrNil

`func (o *ErrorTrackingSymbolSet) SetStoragePtrNil(b bool)`

 SetStoragePtrNil sets the value for StoragePtr to be an explicit nil

### UnsetStoragePtr
`func (o *ErrorTrackingSymbolSet) UnsetStoragePtr()`

UnsetStoragePtr ensures that no value is present for StoragePtr, not even an explicit nil
### GetFailureReason

`func (o *ErrorTrackingSymbolSet) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ErrorTrackingSymbolSet) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ErrorTrackingSymbolSet) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ErrorTrackingSymbolSet) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *ErrorTrackingSymbolSet) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *ErrorTrackingSymbolSet) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


