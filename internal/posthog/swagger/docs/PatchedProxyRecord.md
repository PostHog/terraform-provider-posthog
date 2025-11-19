# PatchedProxyRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Domain** | Pointer to **string** |  | [optional] 
**TargetCname** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to [**ProxyRecordStatusEnum**](ProxyRecordStatusEnum.md) |  | [optional] [readonly] 
**Message** | Pointer to **NullableString** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to **NullableInt32** |  | [optional] [readonly] 

## Methods

### NewPatchedProxyRecord

`func NewPatchedProxyRecord() *PatchedProxyRecord`

NewPatchedProxyRecord instantiates a new PatchedProxyRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedProxyRecordWithDefaults

`func NewPatchedProxyRecordWithDefaults() *PatchedProxyRecord`

NewPatchedProxyRecordWithDefaults instantiates a new PatchedProxyRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedProxyRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedProxyRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedProxyRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedProxyRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *PatchedProxyRecord) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PatchedProxyRecord) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PatchedProxyRecord) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PatchedProxyRecord) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTargetCname

`func (o *PatchedProxyRecord) GetTargetCname() string`

GetTargetCname returns the TargetCname field if non-nil, zero value otherwise.

### GetTargetCnameOk

`func (o *PatchedProxyRecord) GetTargetCnameOk() (*string, bool)`

GetTargetCnameOk returns a tuple with the TargetCname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCname

`func (o *PatchedProxyRecord) SetTargetCname(v string)`

SetTargetCname sets TargetCname field to given value.

### HasTargetCname

`func (o *PatchedProxyRecord) HasTargetCname() bool`

HasTargetCname returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedProxyRecord) GetStatus() ProxyRecordStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedProxyRecord) GetStatusOk() (*ProxyRecordStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedProxyRecord) SetStatus(v ProxyRecordStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedProxyRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *PatchedProxyRecord) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PatchedProxyRecord) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PatchedProxyRecord) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PatchedProxyRecord) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *PatchedProxyRecord) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *PatchedProxyRecord) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCreatedAt

`func (o *PatchedProxyRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedProxyRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedProxyRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedProxyRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedProxyRecord) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedProxyRecord) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedProxyRecord) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedProxyRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedProxyRecord) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedProxyRecord) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedProxyRecord) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedProxyRecord) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *PatchedProxyRecord) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *PatchedProxyRecord) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


