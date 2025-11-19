# SessionsTimelineQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **NullableString** | Only fetch sessions that started after this timestamp (default: &#39;-24h&#39;) | [optional] 
**Before** | Pointer to **NullableString** | Only fetch sessions that started before this timestamp (default: &#39;+5s&#39;) | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "SessionsTimelineQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**PersonId** | Pointer to **NullableString** | Fetch sessions only for a given person | [optional] 
**Response** | Pointer to [**SessionsTimelineQueryResponse**](SessionsTimelineQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewSessionsTimelineQuery

`func NewSessionsTimelineQuery() *SessionsTimelineQuery`

NewSessionsTimelineQuery instantiates a new SessionsTimelineQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionsTimelineQueryWithDefaults

`func NewSessionsTimelineQueryWithDefaults() *SessionsTimelineQuery`

NewSessionsTimelineQueryWithDefaults instantiates a new SessionsTimelineQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *SessionsTimelineQuery) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *SessionsTimelineQuery) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *SessionsTimelineQuery) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *SessionsTimelineQuery) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *SessionsTimelineQuery) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *SessionsTimelineQuery) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetBefore

`func (o *SessionsTimelineQuery) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *SessionsTimelineQuery) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *SessionsTimelineQuery) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *SessionsTimelineQuery) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### SetBeforeNil

`func (o *SessionsTimelineQuery) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *SessionsTimelineQuery) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil
### GetKind

`func (o *SessionsTimelineQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SessionsTimelineQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SessionsTimelineQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SessionsTimelineQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *SessionsTimelineQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *SessionsTimelineQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *SessionsTimelineQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *SessionsTimelineQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetPersonId

`func (o *SessionsTimelineQuery) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *SessionsTimelineQuery) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *SessionsTimelineQuery) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *SessionsTimelineQuery) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### SetPersonIdNil

`func (o *SessionsTimelineQuery) SetPersonIdNil(b bool)`

 SetPersonIdNil sets the value for PersonId to be an explicit nil

### UnsetPersonId
`func (o *SessionsTimelineQuery) UnsetPersonId()`

UnsetPersonId ensures that no value is present for PersonId, not even an explicit nil
### GetResponse

`func (o *SessionsTimelineQuery) GetResponse() SessionsTimelineQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *SessionsTimelineQuery) GetResponseOk() (*SessionsTimelineQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *SessionsTimelineQuery) SetResponse(v SessionsTimelineQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *SessionsTimelineQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *SessionsTimelineQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SessionsTimelineQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SessionsTimelineQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SessionsTimelineQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *SessionsTimelineQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SessionsTimelineQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SessionsTimelineQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SessionsTimelineQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *SessionsTimelineQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *SessionsTimelineQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


