# SuggestedQuestionsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] [default to "SuggestedQuestionsQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Response** | Pointer to [**SuggestedQuestionsQueryResponse**](SuggestedQuestionsQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewSuggestedQuestionsQuery

`func NewSuggestedQuestionsQuery() *SuggestedQuestionsQuery`

NewSuggestedQuestionsQuery instantiates a new SuggestedQuestionsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedQuestionsQueryWithDefaults

`func NewSuggestedQuestionsQueryWithDefaults() *SuggestedQuestionsQuery`

NewSuggestedQuestionsQueryWithDefaults instantiates a new SuggestedQuestionsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SuggestedQuestionsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SuggestedQuestionsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SuggestedQuestionsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SuggestedQuestionsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *SuggestedQuestionsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *SuggestedQuestionsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *SuggestedQuestionsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *SuggestedQuestionsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetResponse

`func (o *SuggestedQuestionsQuery) GetResponse() SuggestedQuestionsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *SuggestedQuestionsQuery) GetResponseOk() (*SuggestedQuestionsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *SuggestedQuestionsQuery) SetResponse(v SuggestedQuestionsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *SuggestedQuestionsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *SuggestedQuestionsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SuggestedQuestionsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SuggestedQuestionsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SuggestedQuestionsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *SuggestedQuestionsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SuggestedQuestionsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SuggestedQuestionsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SuggestedQuestionsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *SuggestedQuestionsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *SuggestedQuestionsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


