# RecordingsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**After** | Pointer to **NullableString** | Cursor for pagination. Contains the ordering value and session_id from the last record of the previous page. | [optional] 
**CommentText** | Pointer to [**RecordingPropertyFilter**](RecordingPropertyFilter.md) |  | [optional] 
**ConsoleLogFilters** | Pointer to [**[]LogEntryPropertyFilter**](LogEntryPropertyFilter.md) |  | [optional] 
**DateFrom** | Pointer to **NullableString** |  | [optional] [default to "-3d"]
**DateTo** | Pointer to **NullableString** |  | [optional] 
**DistinctIds** | Pointer to **[]string** |  | [optional] 
**Events** | Pointer to **[]map[string]interface{}** |  | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** |  | [optional] 
**HavingPredicates** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "RecordingsQuery"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**Operand** | Pointer to [**FilterLogicalOperator**](FilterLogicalOperator.md) |  | [optional] 
**Order** | Pointer to [**RecordingOrder**](RecordingOrder.md) |  | [optional] 
**OrderDirection** | Pointer to [**RecordingOrderDirection**](RecordingOrderDirection.md) |  | [optional] 
**PersonUuid** | Pointer to **NullableString** |  | [optional] 
**Properties** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) |  | [optional] 
**Response** | Pointer to [**RecordingsQueryResponse**](RecordingsQueryResponse.md) |  | [optional] 
**SessionIds** | Pointer to **[]string** |  | [optional] 
**SessionRecordingId** | Pointer to **NullableString** | If provided, this recording will be fetched and prepended to the results, even if it doesn&#39;t match the filters | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**UserModifiedFilters** | Pointer to **map[string]interface{}** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewRecordingsQuery

`func NewRecordingsQuery() *RecordingsQuery`

NewRecordingsQuery instantiates a new RecordingsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingsQueryWithDefaults

`func NewRecordingsQueryWithDefaults() *RecordingsQuery`

NewRecordingsQueryWithDefaults instantiates a new RecordingsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *RecordingsQuery) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *RecordingsQuery) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *RecordingsQuery) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *RecordingsQuery) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *RecordingsQuery) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *RecordingsQuery) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetAfter

`func (o *RecordingsQuery) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *RecordingsQuery) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *RecordingsQuery) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *RecordingsQuery) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *RecordingsQuery) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *RecordingsQuery) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetCommentText

`func (o *RecordingsQuery) GetCommentText() RecordingPropertyFilter`

GetCommentText returns the CommentText field if non-nil, zero value otherwise.

### GetCommentTextOk

`func (o *RecordingsQuery) GetCommentTextOk() (*RecordingPropertyFilter, bool)`

GetCommentTextOk returns a tuple with the CommentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentText

`func (o *RecordingsQuery) SetCommentText(v RecordingPropertyFilter)`

SetCommentText sets CommentText field to given value.

### HasCommentText

`func (o *RecordingsQuery) HasCommentText() bool`

HasCommentText returns a boolean if a field has been set.

### GetConsoleLogFilters

`func (o *RecordingsQuery) GetConsoleLogFilters() []LogEntryPropertyFilter`

GetConsoleLogFilters returns the ConsoleLogFilters field if non-nil, zero value otherwise.

### GetConsoleLogFiltersOk

`func (o *RecordingsQuery) GetConsoleLogFiltersOk() (*[]LogEntryPropertyFilter, bool)`

GetConsoleLogFiltersOk returns a tuple with the ConsoleLogFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleLogFilters

`func (o *RecordingsQuery) SetConsoleLogFilters(v []LogEntryPropertyFilter)`

SetConsoleLogFilters sets ConsoleLogFilters field to given value.

### HasConsoleLogFilters

`func (o *RecordingsQuery) HasConsoleLogFilters() bool`

HasConsoleLogFilters returns a boolean if a field has been set.

### SetConsoleLogFiltersNil

`func (o *RecordingsQuery) SetConsoleLogFiltersNil(b bool)`

 SetConsoleLogFiltersNil sets the value for ConsoleLogFilters to be an explicit nil

### UnsetConsoleLogFilters
`func (o *RecordingsQuery) UnsetConsoleLogFilters()`

UnsetConsoleLogFilters ensures that no value is present for ConsoleLogFilters, not even an explicit nil
### GetDateFrom

`func (o *RecordingsQuery) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *RecordingsQuery) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *RecordingsQuery) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *RecordingsQuery) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### SetDateFromNil

`func (o *RecordingsQuery) SetDateFromNil(b bool)`

 SetDateFromNil sets the value for DateFrom to be an explicit nil

### UnsetDateFrom
`func (o *RecordingsQuery) UnsetDateFrom()`

UnsetDateFrom ensures that no value is present for DateFrom, not even an explicit nil
### GetDateTo

`func (o *RecordingsQuery) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *RecordingsQuery) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *RecordingsQuery) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *RecordingsQuery) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### SetDateToNil

`func (o *RecordingsQuery) SetDateToNil(b bool)`

 SetDateToNil sets the value for DateTo to be an explicit nil

### UnsetDateTo
`func (o *RecordingsQuery) UnsetDateTo()`

UnsetDateTo ensures that no value is present for DateTo, not even an explicit nil
### GetDistinctIds

`func (o *RecordingsQuery) GetDistinctIds() []string`

GetDistinctIds returns the DistinctIds field if non-nil, zero value otherwise.

### GetDistinctIdsOk

`func (o *RecordingsQuery) GetDistinctIdsOk() (*[]string, bool)`

GetDistinctIdsOk returns a tuple with the DistinctIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctIds

`func (o *RecordingsQuery) SetDistinctIds(v []string)`

SetDistinctIds sets DistinctIds field to given value.

### HasDistinctIds

`func (o *RecordingsQuery) HasDistinctIds() bool`

HasDistinctIds returns a boolean if a field has been set.

### SetDistinctIdsNil

`func (o *RecordingsQuery) SetDistinctIdsNil(b bool)`

 SetDistinctIdsNil sets the value for DistinctIds to be an explicit nil

### UnsetDistinctIds
`func (o *RecordingsQuery) UnsetDistinctIds()`

UnsetDistinctIds ensures that no value is present for DistinctIds, not even an explicit nil
### GetEvents

`func (o *RecordingsQuery) GetEvents() []map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RecordingsQuery) GetEventsOk() (*[]map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RecordingsQuery) SetEvents(v []map[string]interface{})`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RecordingsQuery) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *RecordingsQuery) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *RecordingsQuery) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetFilterTestAccounts

`func (o *RecordingsQuery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *RecordingsQuery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *RecordingsQuery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *RecordingsQuery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *RecordingsQuery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *RecordingsQuery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetHavingPredicates

`func (o *RecordingsQuery) GetHavingPredicates() []FixedpropertiesInner`

GetHavingPredicates returns the HavingPredicates field if non-nil, zero value otherwise.

### GetHavingPredicatesOk

`func (o *RecordingsQuery) GetHavingPredicatesOk() (*[]FixedpropertiesInner, bool)`

GetHavingPredicatesOk returns a tuple with the HavingPredicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHavingPredicates

`func (o *RecordingsQuery) SetHavingPredicates(v []FixedpropertiesInner)`

SetHavingPredicates sets HavingPredicates field to given value.

### HasHavingPredicates

`func (o *RecordingsQuery) HasHavingPredicates() bool`

HasHavingPredicates returns a boolean if a field has been set.

### SetHavingPredicatesNil

`func (o *RecordingsQuery) SetHavingPredicatesNil(b bool)`

 SetHavingPredicatesNil sets the value for HavingPredicates to be an explicit nil

### UnsetHavingPredicates
`func (o *RecordingsQuery) UnsetHavingPredicates()`

UnsetHavingPredicates ensures that no value is present for HavingPredicates, not even an explicit nil
### GetKind

`func (o *RecordingsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RecordingsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RecordingsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RecordingsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *RecordingsQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RecordingsQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RecordingsQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RecordingsQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *RecordingsQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *RecordingsQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *RecordingsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *RecordingsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *RecordingsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *RecordingsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *RecordingsQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RecordingsQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RecordingsQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RecordingsQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *RecordingsQuery) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *RecordingsQuery) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetOperand

`func (o *RecordingsQuery) GetOperand() FilterLogicalOperator`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *RecordingsQuery) GetOperandOk() (*FilterLogicalOperator, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *RecordingsQuery) SetOperand(v FilterLogicalOperator)`

SetOperand sets Operand field to given value.

### HasOperand

`func (o *RecordingsQuery) HasOperand() bool`

HasOperand returns a boolean if a field has been set.

### GetOrder

`func (o *RecordingsQuery) GetOrder() RecordingOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RecordingsQuery) GetOrderOk() (*RecordingOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RecordingsQuery) SetOrder(v RecordingOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RecordingsQuery) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrderDirection

`func (o *RecordingsQuery) GetOrderDirection() RecordingOrderDirection`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *RecordingsQuery) GetOrderDirectionOk() (*RecordingOrderDirection, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *RecordingsQuery) SetOrderDirection(v RecordingOrderDirection)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *RecordingsQuery) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetPersonUuid

`func (o *RecordingsQuery) GetPersonUuid() string`

GetPersonUuid returns the PersonUuid field if non-nil, zero value otherwise.

### GetPersonUuidOk

`func (o *RecordingsQuery) GetPersonUuidOk() (*string, bool)`

GetPersonUuidOk returns a tuple with the PersonUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonUuid

`func (o *RecordingsQuery) SetPersonUuid(v string)`

SetPersonUuid sets PersonUuid field to given value.

### HasPersonUuid

`func (o *RecordingsQuery) HasPersonUuid() bool`

HasPersonUuid returns a boolean if a field has been set.

### SetPersonUuidNil

`func (o *RecordingsQuery) SetPersonUuidNil(b bool)`

 SetPersonUuidNil sets the value for PersonUuid to be an explicit nil

### UnsetPersonUuid
`func (o *RecordingsQuery) UnsetPersonUuid()`

UnsetPersonUuid ensures that no value is present for PersonUuid, not even an explicit nil
### GetProperties

`func (o *RecordingsQuery) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RecordingsQuery) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RecordingsQuery) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *RecordingsQuery) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *RecordingsQuery) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *RecordingsQuery) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *RecordingsQuery) GetResponse() RecordingsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *RecordingsQuery) GetResponseOk() (*RecordingsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *RecordingsQuery) SetResponse(v RecordingsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *RecordingsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSessionIds

`func (o *RecordingsQuery) GetSessionIds() []string`

GetSessionIds returns the SessionIds field if non-nil, zero value otherwise.

### GetSessionIdsOk

`func (o *RecordingsQuery) GetSessionIdsOk() (*[]string, bool)`

GetSessionIdsOk returns a tuple with the SessionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIds

`func (o *RecordingsQuery) SetSessionIds(v []string)`

SetSessionIds sets SessionIds field to given value.

### HasSessionIds

`func (o *RecordingsQuery) HasSessionIds() bool`

HasSessionIds returns a boolean if a field has been set.

### SetSessionIdsNil

`func (o *RecordingsQuery) SetSessionIdsNil(b bool)`

 SetSessionIdsNil sets the value for SessionIds to be an explicit nil

### UnsetSessionIds
`func (o *RecordingsQuery) UnsetSessionIds()`

UnsetSessionIds ensures that no value is present for SessionIds, not even an explicit nil
### GetSessionRecordingId

`func (o *RecordingsQuery) GetSessionRecordingId() string`

GetSessionRecordingId returns the SessionRecordingId field if non-nil, zero value otherwise.

### GetSessionRecordingIdOk

`func (o *RecordingsQuery) GetSessionRecordingIdOk() (*string, bool)`

GetSessionRecordingIdOk returns a tuple with the SessionRecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingId

`func (o *RecordingsQuery) SetSessionRecordingId(v string)`

SetSessionRecordingId sets SessionRecordingId field to given value.

### HasSessionRecordingId

`func (o *RecordingsQuery) HasSessionRecordingId() bool`

HasSessionRecordingId returns a boolean if a field has been set.

### SetSessionRecordingIdNil

`func (o *RecordingsQuery) SetSessionRecordingIdNil(b bool)`

 SetSessionRecordingIdNil sets the value for SessionRecordingId to be an explicit nil

### UnsetSessionRecordingId
`func (o *RecordingsQuery) UnsetSessionRecordingId()`

UnsetSessionRecordingId ensures that no value is present for SessionRecordingId, not even an explicit nil
### GetTags

`func (o *RecordingsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RecordingsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RecordingsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RecordingsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUserModifiedFilters

`func (o *RecordingsQuery) GetUserModifiedFilters() map[string]interface{}`

GetUserModifiedFilters returns the UserModifiedFilters field if non-nil, zero value otherwise.

### GetUserModifiedFiltersOk

`func (o *RecordingsQuery) GetUserModifiedFiltersOk() (*map[string]interface{}, bool)`

GetUserModifiedFiltersOk returns a tuple with the UserModifiedFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserModifiedFilters

`func (o *RecordingsQuery) SetUserModifiedFilters(v map[string]interface{})`

SetUserModifiedFilters sets UserModifiedFilters field to given value.

### HasUserModifiedFilters

`func (o *RecordingsQuery) HasUserModifiedFilters() bool`

HasUserModifiedFilters returns a boolean if a field has been set.

### SetUserModifiedFiltersNil

`func (o *RecordingsQuery) SetUserModifiedFiltersNil(b bool)`

 SetUserModifiedFiltersNil sets the value for UserModifiedFilters to be an explicit nil

### UnsetUserModifiedFilters
`func (o *RecordingsQuery) UnsetUserModifiedFilters()`

UnsetUserModifiedFilters ensures that no value is present for UserModifiedFilters, not even an explicit nil
### GetVersion

`func (o *RecordingsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RecordingsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RecordingsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RecordingsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *RecordingsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *RecordingsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


