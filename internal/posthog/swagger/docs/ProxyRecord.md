# ProxyRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Domain** | **string** |  | 
**TargetCname** | **string** |  | [readonly] 
**Status** | [**ProxyRecordStatusEnum**](ProxyRecordStatusEnum.md) |  | [readonly] 
**Message** | **NullableString** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**CreatedBy** | **NullableInt32** |  | [readonly] 

## Methods

### NewProxyRecord

`func NewProxyRecord(id string, domain string, targetCname string, status ProxyRecordStatusEnum, message NullableString, createdAt time.Time, updatedAt time.Time, createdBy NullableInt32, ) *ProxyRecord`

NewProxyRecord instantiates a new ProxyRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyRecordWithDefaults

`func NewProxyRecordWithDefaults() *ProxyRecord`

NewProxyRecordWithDefaults instantiates a new ProxyRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProxyRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProxyRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProxyRecord) SetId(v string)`

SetId sets Id field to given value.


### GetDomain

`func (o *ProxyRecord) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ProxyRecord) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ProxyRecord) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetTargetCname

`func (o *ProxyRecord) GetTargetCname() string`

GetTargetCname returns the TargetCname field if non-nil, zero value otherwise.

### GetTargetCnameOk

`func (o *ProxyRecord) GetTargetCnameOk() (*string, bool)`

GetTargetCnameOk returns a tuple with the TargetCname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCname

`func (o *ProxyRecord) SetTargetCname(v string)`

SetTargetCname sets TargetCname field to given value.


### GetStatus

`func (o *ProxyRecord) GetStatus() ProxyRecordStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProxyRecord) GetStatusOk() (*ProxyRecordStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProxyRecord) SetStatus(v ProxyRecordStatusEnum)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ProxyRecord) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProxyRecord) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProxyRecord) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *ProxyRecord) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ProxyRecord) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCreatedAt

`func (o *ProxyRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProxyRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProxyRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ProxyRecord) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProxyRecord) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProxyRecord) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedBy

`func (o *ProxyRecord) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProxyRecord) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProxyRecord) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *ProxyRecord) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ProxyRecord) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


