# PatchedSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Dashboard** | Pointer to **NullableInt32** |  | [optional] 
**Insight** | Pointer to **NullableInt32** |  | [optional] 
**TargetType** | Pointer to [**TargetTypeEnum**](TargetTypeEnum.md) |  | [optional] 
**TargetValue** | Pointer to **string** |  | [optional] 
**Frequency** | Pointer to [**FrequencyEnum**](FrequencyEnum.md) |  | [optional] 
**Interval** | Pointer to **int32** |  | [optional] 
**Byweekday** | Pointer to [**[]ByweekdayEnum**](ByweekdayEnum.md) |  | [optional] 
**Bysetpos** | Pointer to **NullableInt32** |  | [optional] 
**Count** | Pointer to **NullableInt32** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**UntilDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] [readonly] 
**NextDeliveryDate** | Pointer to **NullableTime** |  | [optional] [readonly] 
**InviteMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchedSubscription

`func NewPatchedSubscription() *PatchedSubscription`

NewPatchedSubscription instantiates a new PatchedSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSubscriptionWithDefaults

`func NewPatchedSubscriptionWithDefaults() *PatchedSubscription`

NewPatchedSubscriptionWithDefaults instantiates a new PatchedSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedSubscription) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedSubscription) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedSubscription) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDashboard

`func (o *PatchedSubscription) GetDashboard() int32`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *PatchedSubscription) GetDashboardOk() (*int32, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *PatchedSubscription) SetDashboard(v int32)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *PatchedSubscription) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### SetDashboardNil

`func (o *PatchedSubscription) SetDashboardNil(b bool)`

 SetDashboardNil sets the value for Dashboard to be an explicit nil

### UnsetDashboard
`func (o *PatchedSubscription) UnsetDashboard()`

UnsetDashboard ensures that no value is present for Dashboard, not even an explicit nil
### GetInsight

`func (o *PatchedSubscription) GetInsight() int32`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *PatchedSubscription) GetInsightOk() (*int32, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *PatchedSubscription) SetInsight(v int32)`

SetInsight sets Insight field to given value.

### HasInsight

`func (o *PatchedSubscription) HasInsight() bool`

HasInsight returns a boolean if a field has been set.

### SetInsightNil

`func (o *PatchedSubscription) SetInsightNil(b bool)`

 SetInsightNil sets the value for Insight to be an explicit nil

### UnsetInsight
`func (o *PatchedSubscription) UnsetInsight()`

UnsetInsight ensures that no value is present for Insight, not even an explicit nil
### GetTargetType

`func (o *PatchedSubscription) GetTargetType() TargetTypeEnum`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *PatchedSubscription) GetTargetTypeOk() (*TargetTypeEnum, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *PatchedSubscription) SetTargetType(v TargetTypeEnum)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *PatchedSubscription) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTargetValue

`func (o *PatchedSubscription) GetTargetValue() string`

GetTargetValue returns the TargetValue field if non-nil, zero value otherwise.

### GetTargetValueOk

`func (o *PatchedSubscription) GetTargetValueOk() (*string, bool)`

GetTargetValueOk returns a tuple with the TargetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetValue

`func (o *PatchedSubscription) SetTargetValue(v string)`

SetTargetValue sets TargetValue field to given value.

### HasTargetValue

`func (o *PatchedSubscription) HasTargetValue() bool`

HasTargetValue returns a boolean if a field has been set.

### GetFrequency

`func (o *PatchedSubscription) GetFrequency() FrequencyEnum`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *PatchedSubscription) GetFrequencyOk() (*FrequencyEnum, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *PatchedSubscription) SetFrequency(v FrequencyEnum)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *PatchedSubscription) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetInterval

`func (o *PatchedSubscription) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PatchedSubscription) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PatchedSubscription) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PatchedSubscription) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetByweekday

`func (o *PatchedSubscription) GetByweekday() []ByweekdayEnum`

GetByweekday returns the Byweekday field if non-nil, zero value otherwise.

### GetByweekdayOk

`func (o *PatchedSubscription) GetByweekdayOk() (*[]ByweekdayEnum, bool)`

GetByweekdayOk returns a tuple with the Byweekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByweekday

`func (o *PatchedSubscription) SetByweekday(v []ByweekdayEnum)`

SetByweekday sets Byweekday field to given value.

### HasByweekday

`func (o *PatchedSubscription) HasByweekday() bool`

HasByweekday returns a boolean if a field has been set.

### SetByweekdayNil

`func (o *PatchedSubscription) SetByweekdayNil(b bool)`

 SetByweekdayNil sets the value for Byweekday to be an explicit nil

### UnsetByweekday
`func (o *PatchedSubscription) UnsetByweekday()`

UnsetByweekday ensures that no value is present for Byweekday, not even an explicit nil
### GetBysetpos

`func (o *PatchedSubscription) GetBysetpos() int32`

GetBysetpos returns the Bysetpos field if non-nil, zero value otherwise.

### GetBysetposOk

`func (o *PatchedSubscription) GetBysetposOk() (*int32, bool)`

GetBysetposOk returns a tuple with the Bysetpos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBysetpos

`func (o *PatchedSubscription) SetBysetpos(v int32)`

SetBysetpos sets Bysetpos field to given value.

### HasBysetpos

`func (o *PatchedSubscription) HasBysetpos() bool`

HasBysetpos returns a boolean if a field has been set.

### SetBysetposNil

`func (o *PatchedSubscription) SetBysetposNil(b bool)`

 SetBysetposNil sets the value for Bysetpos to be an explicit nil

### UnsetBysetpos
`func (o *PatchedSubscription) UnsetBysetpos()`

UnsetBysetpos ensures that no value is present for Bysetpos, not even an explicit nil
### GetCount

`func (o *PatchedSubscription) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PatchedSubscription) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PatchedSubscription) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PatchedSubscription) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *PatchedSubscription) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PatchedSubscription) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetStartDate

`func (o *PatchedSubscription) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PatchedSubscription) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PatchedSubscription) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PatchedSubscription) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetUntilDate

`func (o *PatchedSubscription) GetUntilDate() time.Time`

GetUntilDate returns the UntilDate field if non-nil, zero value otherwise.

### GetUntilDateOk

`func (o *PatchedSubscription) GetUntilDateOk() (*time.Time, bool)`

GetUntilDateOk returns a tuple with the UntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntilDate

`func (o *PatchedSubscription) SetUntilDate(v time.Time)`

SetUntilDate sets UntilDate field to given value.

### HasUntilDate

`func (o *PatchedSubscription) HasUntilDate() bool`

HasUntilDate returns a boolean if a field has been set.

### SetUntilDateNil

`func (o *PatchedSubscription) SetUntilDateNil(b bool)`

 SetUntilDateNil sets the value for UntilDate to be an explicit nil

### UnsetUntilDate
`func (o *PatchedSubscription) UnsetUntilDate()`

UnsetUntilDate ensures that no value is present for UntilDate, not even an explicit nil
### GetCreatedAt

`func (o *PatchedSubscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedSubscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedSubscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedSubscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedSubscription) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedSubscription) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedSubscription) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedSubscription) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeleted

`func (o *PatchedSubscription) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedSubscription) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedSubscription) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedSubscription) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetTitle

`func (o *PatchedSubscription) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PatchedSubscription) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PatchedSubscription) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PatchedSubscription) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PatchedSubscription) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PatchedSubscription) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSummary

`func (o *PatchedSubscription) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PatchedSubscription) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PatchedSubscription) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PatchedSubscription) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetNextDeliveryDate

`func (o *PatchedSubscription) GetNextDeliveryDate() time.Time`

GetNextDeliveryDate returns the NextDeliveryDate field if non-nil, zero value otherwise.

### GetNextDeliveryDateOk

`func (o *PatchedSubscription) GetNextDeliveryDateOk() (*time.Time, bool)`

GetNextDeliveryDateOk returns a tuple with the NextDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDeliveryDate

`func (o *PatchedSubscription) SetNextDeliveryDate(v time.Time)`

SetNextDeliveryDate sets NextDeliveryDate field to given value.

### HasNextDeliveryDate

`func (o *PatchedSubscription) HasNextDeliveryDate() bool`

HasNextDeliveryDate returns a boolean if a field has been set.

### SetNextDeliveryDateNil

`func (o *PatchedSubscription) SetNextDeliveryDateNil(b bool)`

 SetNextDeliveryDateNil sets the value for NextDeliveryDate to be an explicit nil

### UnsetNextDeliveryDate
`func (o *PatchedSubscription) UnsetNextDeliveryDate()`

UnsetNextDeliveryDate ensures that no value is present for NextDeliveryDate, not even an explicit nil
### GetInviteMessage

`func (o *PatchedSubscription) GetInviteMessage() string`

GetInviteMessage returns the InviteMessage field if non-nil, zero value otherwise.

### GetInviteMessageOk

`func (o *PatchedSubscription) GetInviteMessageOk() (*string, bool)`

GetInviteMessageOk returns a tuple with the InviteMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteMessage

`func (o *PatchedSubscription) SetInviteMessage(v string)`

SetInviteMessage sets InviteMessage field to given value.

### HasInviteMessage

`func (o *PatchedSubscription) HasInviteMessage() bool`

HasInviteMessage returns a boolean if a field has been set.

### SetInviteMessageNil

`func (o *PatchedSubscription) SetInviteMessageNil(b bool)`

 SetInviteMessageNil sets the value for InviteMessage to be an explicit nil

### UnsetInviteMessage
`func (o *PatchedSubscription) UnsetInviteMessage()`

UnsetInviteMessage ensures that no value is present for InviteMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


