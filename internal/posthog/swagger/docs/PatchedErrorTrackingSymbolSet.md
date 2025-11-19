# PatchedErrorTrackingSymbolSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Ref** | Pointer to **string** |  | [optional] 
**TeamId** | Pointer to **int32** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**StoragePtr** | Pointer to **NullableString** |  | [optional] 
**FailureReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchedErrorTrackingSymbolSet

`func NewPatchedErrorTrackingSymbolSet() *PatchedErrorTrackingSymbolSet`

NewPatchedErrorTrackingSymbolSet instantiates a new PatchedErrorTrackingSymbolSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedErrorTrackingSymbolSetWithDefaults

`func NewPatchedErrorTrackingSymbolSetWithDefaults() *PatchedErrorTrackingSymbolSet`

NewPatchedErrorTrackingSymbolSetWithDefaults instantiates a new PatchedErrorTrackingSymbolSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedErrorTrackingSymbolSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedErrorTrackingSymbolSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedErrorTrackingSymbolSet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedErrorTrackingSymbolSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRef

`func (o *PatchedErrorTrackingSymbolSet) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PatchedErrorTrackingSymbolSet) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PatchedErrorTrackingSymbolSet) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *PatchedErrorTrackingSymbolSet) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetTeamId

`func (o *PatchedErrorTrackingSymbolSet) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *PatchedErrorTrackingSymbolSet) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *PatchedErrorTrackingSymbolSet) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *PatchedErrorTrackingSymbolSet) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedErrorTrackingSymbolSet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedErrorTrackingSymbolSet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedErrorTrackingSymbolSet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedErrorTrackingSymbolSet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStoragePtr

`func (o *PatchedErrorTrackingSymbolSet) GetStoragePtr() string`

GetStoragePtr returns the StoragePtr field if non-nil, zero value otherwise.

### GetStoragePtrOk

`func (o *PatchedErrorTrackingSymbolSet) GetStoragePtrOk() (*string, bool)`

GetStoragePtrOk returns a tuple with the StoragePtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePtr

`func (o *PatchedErrorTrackingSymbolSet) SetStoragePtr(v string)`

SetStoragePtr sets StoragePtr field to given value.

### HasStoragePtr

`func (o *PatchedErrorTrackingSymbolSet) HasStoragePtr() bool`

HasStoragePtr returns a boolean if a field has been set.

### SetStoragePtrNil

`func (o *PatchedErrorTrackingSymbolSet) SetStoragePtrNil(b bool)`

 SetStoragePtrNil sets the value for StoragePtr to be an explicit nil

### UnsetStoragePtr
`func (o *PatchedErrorTrackingSymbolSet) UnsetStoragePtr()`

UnsetStoragePtr ensures that no value is present for StoragePtr, not even an explicit nil
### GetFailureReason

`func (o *PatchedErrorTrackingSymbolSet) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *PatchedErrorTrackingSymbolSet) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *PatchedErrorTrackingSymbolSet) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *PatchedErrorTrackingSymbolSet) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *PatchedErrorTrackingSymbolSet) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *PatchedErrorTrackingSymbolSet) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


