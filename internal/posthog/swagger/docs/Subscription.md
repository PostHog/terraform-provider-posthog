# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Dashboard** | Pointer to **NullableInt32** |  | [optional] 
**Insight** | Pointer to **NullableInt32** |  | [optional] 
**TargetType** | [**TargetTypeEnum**](TargetTypeEnum.md) |  | 
**TargetValue** | **string** |  | 
**Frequency** | [**FrequencyEnum**](FrequencyEnum.md) |  | 
**Interval** | Pointer to **int32** |  | [optional] 
**Byweekday** | Pointer to [**[]ByweekdayEnum**](ByweekdayEnum.md) |  | [optional] 
**Bysetpos** | Pointer to **NullableInt32** |  | [optional] 
**Count** | Pointer to **NullableInt32** |  | [optional] 
**StartDate** | **time.Time** |  | 
**UntilDate** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Summary** | **string** |  | [readonly] 
**NextDeliveryDate** | **NullableTime** |  | [readonly] 
**InviteMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSubscription

`func NewSubscription(id int32, targetType TargetTypeEnum, targetValue string, frequency FrequencyEnum, startDate time.Time, createdAt time.Time, createdBy UserBasic, summary string, nextDeliveryDate NullableTime, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscription) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v int32)`

SetId sets Id field to given value.


### GetDashboard

`func (o *Subscription) GetDashboard() int32`

GetDashboard returns the Dashboard field if non-nil, zero value otherwise.

### GetDashboardOk

`func (o *Subscription) GetDashboardOk() (*int32, bool)`

GetDashboardOk returns a tuple with the Dashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboard

`func (o *Subscription) SetDashboard(v int32)`

SetDashboard sets Dashboard field to given value.

### HasDashboard

`func (o *Subscription) HasDashboard() bool`

HasDashboard returns a boolean if a field has been set.

### SetDashboardNil

`func (o *Subscription) SetDashboardNil(b bool)`

 SetDashboardNil sets the value for Dashboard to be an explicit nil

### UnsetDashboard
`func (o *Subscription) UnsetDashboard()`

UnsetDashboard ensures that no value is present for Dashboard, not even an explicit nil
### GetInsight

`func (o *Subscription) GetInsight() int32`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *Subscription) GetInsightOk() (*int32, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *Subscription) SetInsight(v int32)`

SetInsight sets Insight field to given value.

### HasInsight

`func (o *Subscription) HasInsight() bool`

HasInsight returns a boolean if a field has been set.

### SetInsightNil

`func (o *Subscription) SetInsightNil(b bool)`

 SetInsightNil sets the value for Insight to be an explicit nil

### UnsetInsight
`func (o *Subscription) UnsetInsight()`

UnsetInsight ensures that no value is present for Insight, not even an explicit nil
### GetTargetType

`func (o *Subscription) GetTargetType() TargetTypeEnum`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *Subscription) GetTargetTypeOk() (*TargetTypeEnum, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *Subscription) SetTargetType(v TargetTypeEnum)`

SetTargetType sets TargetType field to given value.


### GetTargetValue

`func (o *Subscription) GetTargetValue() string`

GetTargetValue returns the TargetValue field if non-nil, zero value otherwise.

### GetTargetValueOk

`func (o *Subscription) GetTargetValueOk() (*string, bool)`

GetTargetValueOk returns a tuple with the TargetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetValue

`func (o *Subscription) SetTargetValue(v string)`

SetTargetValue sets TargetValue field to given value.


### GetFrequency

`func (o *Subscription) GetFrequency() FrequencyEnum`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *Subscription) GetFrequencyOk() (*FrequencyEnum, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *Subscription) SetFrequency(v FrequencyEnum)`

SetFrequency sets Frequency field to given value.


### GetInterval

`func (o *Subscription) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Subscription) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Subscription) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Subscription) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetByweekday

`func (o *Subscription) GetByweekday() []ByweekdayEnum`

GetByweekday returns the Byweekday field if non-nil, zero value otherwise.

### GetByweekdayOk

`func (o *Subscription) GetByweekdayOk() (*[]ByweekdayEnum, bool)`

GetByweekdayOk returns a tuple with the Byweekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByweekday

`func (o *Subscription) SetByweekday(v []ByweekdayEnum)`

SetByweekday sets Byweekday field to given value.

### HasByweekday

`func (o *Subscription) HasByweekday() bool`

HasByweekday returns a boolean if a field has been set.

### SetByweekdayNil

`func (o *Subscription) SetByweekdayNil(b bool)`

 SetByweekdayNil sets the value for Byweekday to be an explicit nil

### UnsetByweekday
`func (o *Subscription) UnsetByweekday()`

UnsetByweekday ensures that no value is present for Byweekday, not even an explicit nil
### GetBysetpos

`func (o *Subscription) GetBysetpos() int32`

GetBysetpos returns the Bysetpos field if non-nil, zero value otherwise.

### GetBysetposOk

`func (o *Subscription) GetBysetposOk() (*int32, bool)`

GetBysetposOk returns a tuple with the Bysetpos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBysetpos

`func (o *Subscription) SetBysetpos(v int32)`

SetBysetpos sets Bysetpos field to given value.

### HasBysetpos

`func (o *Subscription) HasBysetpos() bool`

HasBysetpos returns a boolean if a field has been set.

### SetBysetposNil

`func (o *Subscription) SetBysetposNil(b bool)`

 SetBysetposNil sets the value for Bysetpos to be an explicit nil

### UnsetBysetpos
`func (o *Subscription) UnsetBysetpos()`

UnsetBysetpos ensures that no value is present for Bysetpos, not even an explicit nil
### GetCount

`func (o *Subscription) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Subscription) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Subscription) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Subscription) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *Subscription) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *Subscription) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetStartDate

`func (o *Subscription) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Subscription) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Subscription) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetUntilDate

`func (o *Subscription) GetUntilDate() time.Time`

GetUntilDate returns the UntilDate field if non-nil, zero value otherwise.

### GetUntilDateOk

`func (o *Subscription) GetUntilDateOk() (*time.Time, bool)`

GetUntilDateOk returns a tuple with the UntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntilDate

`func (o *Subscription) SetUntilDate(v time.Time)`

SetUntilDate sets UntilDate field to given value.

### HasUntilDate

`func (o *Subscription) HasUntilDate() bool`

HasUntilDate returns a boolean if a field has been set.

### SetUntilDateNil

`func (o *Subscription) SetUntilDateNil(b bool)`

 SetUntilDateNil sets the value for UntilDate to be an explicit nil

### UnsetUntilDate
`func (o *Subscription) UnsetUntilDate()`

UnsetUntilDate ensures that no value is present for UntilDate, not even an explicit nil
### GetCreatedAt

`func (o *Subscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Subscription) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Subscription) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Subscription) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetDeleted

`func (o *Subscription) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Subscription) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Subscription) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Subscription) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetTitle

`func (o *Subscription) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Subscription) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Subscription) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Subscription) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Subscription) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Subscription) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSummary

`func (o *Subscription) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Subscription) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Subscription) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetNextDeliveryDate

`func (o *Subscription) GetNextDeliveryDate() time.Time`

GetNextDeliveryDate returns the NextDeliveryDate field if non-nil, zero value otherwise.

### GetNextDeliveryDateOk

`func (o *Subscription) GetNextDeliveryDateOk() (*time.Time, bool)`

GetNextDeliveryDateOk returns a tuple with the NextDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDeliveryDate

`func (o *Subscription) SetNextDeliveryDate(v time.Time)`

SetNextDeliveryDate sets NextDeliveryDate field to given value.


### SetNextDeliveryDateNil

`func (o *Subscription) SetNextDeliveryDateNil(b bool)`

 SetNextDeliveryDateNil sets the value for NextDeliveryDate to be an explicit nil

### UnsetNextDeliveryDate
`func (o *Subscription) UnsetNextDeliveryDate()`

UnsetNextDeliveryDate ensures that no value is present for NextDeliveryDate, not even an explicit nil
### GetInviteMessage

`func (o *Subscription) GetInviteMessage() string`

GetInviteMessage returns the InviteMessage field if non-nil, zero value otherwise.

### GetInviteMessageOk

`func (o *Subscription) GetInviteMessageOk() (*string, bool)`

GetInviteMessageOk returns a tuple with the InviteMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteMessage

`func (o *Subscription) SetInviteMessage(v string)`

SetInviteMessage sets InviteMessage field to given value.

### HasInviteMessage

`func (o *Subscription) HasInviteMessage() bool`

HasInviteMessage returns a boolean if a field has been set.

### SetInviteMessageNil

`func (o *Subscription) SetInviteMessageNil(b bool)`

 SetInviteMessageNil sets the value for InviteMessage to be an explicit nil

### UnsetInviteMessage
`func (o *Subscription) UnsetInviteMessage()`

UnsetInviteMessage ensures that no value is present for InviteMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


