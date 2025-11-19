# ErrorTrackingIssueCorrelationQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | **[]string** |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "ErrorTrackingIssueCorrelationQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Response** | Pointer to [**ErrorTrackingIssueCorrelationQueryResponse**](ErrorTrackingIssueCorrelationQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewErrorTrackingIssueCorrelationQuery

`func NewErrorTrackingIssueCorrelationQuery(events []string, ) *ErrorTrackingIssueCorrelationQuery`

NewErrorTrackingIssueCorrelationQuery instantiates a new ErrorTrackingIssueCorrelationQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTrackingIssueCorrelationQueryWithDefaults

`func NewErrorTrackingIssueCorrelationQueryWithDefaults() *ErrorTrackingIssueCorrelationQuery`

NewErrorTrackingIssueCorrelationQueryWithDefaults instantiates a new ErrorTrackingIssueCorrelationQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ErrorTrackingIssueCorrelationQuery) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ErrorTrackingIssueCorrelationQuery) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ErrorTrackingIssueCorrelationQuery) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetKind

`func (o *ErrorTrackingIssueCorrelationQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ErrorTrackingIssueCorrelationQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ErrorTrackingIssueCorrelationQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ErrorTrackingIssueCorrelationQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *ErrorTrackingIssueCorrelationQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *ErrorTrackingIssueCorrelationQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *ErrorTrackingIssueCorrelationQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *ErrorTrackingIssueCorrelationQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetResponse

`func (o *ErrorTrackingIssueCorrelationQuery) GetResponse() ErrorTrackingIssueCorrelationQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ErrorTrackingIssueCorrelationQuery) GetResponseOk() (*ErrorTrackingIssueCorrelationQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ErrorTrackingIssueCorrelationQuery) SetResponse(v ErrorTrackingIssueCorrelationQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ErrorTrackingIssueCorrelationQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *ErrorTrackingIssueCorrelationQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ErrorTrackingIssueCorrelationQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ErrorTrackingIssueCorrelationQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ErrorTrackingIssueCorrelationQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *ErrorTrackingIssueCorrelationQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ErrorTrackingIssueCorrelationQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ErrorTrackingIssueCorrelationQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ErrorTrackingIssueCorrelationQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ErrorTrackingIssueCorrelationQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ErrorTrackingIssueCorrelationQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


