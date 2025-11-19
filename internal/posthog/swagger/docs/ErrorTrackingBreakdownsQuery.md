# ErrorTrackingBreakdownsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakdownProperties** | **[]string** |  | 
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** |  | [optional] 
**IssueId** | **string** |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "ErrorTrackingBreakdownsQuery"]
**MaxValuesPerProperty** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Response** | Pointer to [**ErrorTrackingBreakdownsQueryResponse**](ErrorTrackingBreakdownsQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewErrorTrackingBreakdownsQuery

`func NewErrorTrackingBreakdownsQuery(breakdownProperties []string, issueId string, ) *ErrorTrackingBreakdownsQuery`

NewErrorTrackingBreakdownsQuery instantiates a new ErrorTrackingBreakdownsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTrackingBreakdownsQueryWithDefaults

`func NewErrorTrackingBreakdownsQueryWithDefaults() *ErrorTrackingBreakdownsQuery`

NewErrorTrackingBreakdownsQueryWithDefaults instantiates a new ErrorTrackingBreakdownsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdownProperties

`func (o *ErrorTrackingBreakdownsQuery) GetBreakdownProperties() []string`

GetBreakdownProperties returns the BreakdownProperties field if non-nil, zero value otherwise.

### GetBreakdownPropertiesOk

`func (o *ErrorTrackingBreakdownsQuery) GetBreakdownPropertiesOk() (*[]string, bool)`

GetBreakdownPropertiesOk returns a tuple with the BreakdownProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownProperties

`func (o *ErrorTrackingBreakdownsQuery) SetBreakdownProperties(v []string)`

SetBreakdownProperties sets BreakdownProperties field to given value.


### GetDateRange

`func (o *ErrorTrackingBreakdownsQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *ErrorTrackingBreakdownsQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *ErrorTrackingBreakdownsQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *ErrorTrackingBreakdownsQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetFilterTestAccounts

`func (o *ErrorTrackingBreakdownsQuery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *ErrorTrackingBreakdownsQuery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *ErrorTrackingBreakdownsQuery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *ErrorTrackingBreakdownsQuery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *ErrorTrackingBreakdownsQuery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *ErrorTrackingBreakdownsQuery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetIssueId

`func (o *ErrorTrackingBreakdownsQuery) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *ErrorTrackingBreakdownsQuery) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *ErrorTrackingBreakdownsQuery) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.


### GetKind

`func (o *ErrorTrackingBreakdownsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ErrorTrackingBreakdownsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ErrorTrackingBreakdownsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ErrorTrackingBreakdownsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMaxValuesPerProperty

`func (o *ErrorTrackingBreakdownsQuery) GetMaxValuesPerProperty() int32`

GetMaxValuesPerProperty returns the MaxValuesPerProperty field if non-nil, zero value otherwise.

### GetMaxValuesPerPropertyOk

`func (o *ErrorTrackingBreakdownsQuery) GetMaxValuesPerPropertyOk() (*int32, bool)`

GetMaxValuesPerPropertyOk returns a tuple with the MaxValuesPerProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValuesPerProperty

`func (o *ErrorTrackingBreakdownsQuery) SetMaxValuesPerProperty(v int32)`

SetMaxValuesPerProperty sets MaxValuesPerProperty field to given value.

### HasMaxValuesPerProperty

`func (o *ErrorTrackingBreakdownsQuery) HasMaxValuesPerProperty() bool`

HasMaxValuesPerProperty returns a boolean if a field has been set.

### SetMaxValuesPerPropertyNil

`func (o *ErrorTrackingBreakdownsQuery) SetMaxValuesPerPropertyNil(b bool)`

 SetMaxValuesPerPropertyNil sets the value for MaxValuesPerProperty to be an explicit nil

### UnsetMaxValuesPerProperty
`func (o *ErrorTrackingBreakdownsQuery) UnsetMaxValuesPerProperty()`

UnsetMaxValuesPerProperty ensures that no value is present for MaxValuesPerProperty, not even an explicit nil
### GetModifiers

`func (o *ErrorTrackingBreakdownsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *ErrorTrackingBreakdownsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *ErrorTrackingBreakdownsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *ErrorTrackingBreakdownsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetResponse

`func (o *ErrorTrackingBreakdownsQuery) GetResponse() ErrorTrackingBreakdownsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ErrorTrackingBreakdownsQuery) GetResponseOk() (*ErrorTrackingBreakdownsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ErrorTrackingBreakdownsQuery) SetResponse(v ErrorTrackingBreakdownsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ErrorTrackingBreakdownsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *ErrorTrackingBreakdownsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ErrorTrackingBreakdownsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ErrorTrackingBreakdownsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ErrorTrackingBreakdownsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *ErrorTrackingBreakdownsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ErrorTrackingBreakdownsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ErrorTrackingBreakdownsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ErrorTrackingBreakdownsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ErrorTrackingBreakdownsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ErrorTrackingBreakdownsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


