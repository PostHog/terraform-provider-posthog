# TeamTaxonomyQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] [default to "TeamTaxonomyQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Response** | Pointer to [**TeamTaxonomyQueryResponse**](TeamTaxonomyQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewTeamTaxonomyQuery

`func NewTeamTaxonomyQuery() *TeamTaxonomyQuery`

NewTeamTaxonomyQuery instantiates a new TeamTaxonomyQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTaxonomyQueryWithDefaults

`func NewTeamTaxonomyQueryWithDefaults() *TeamTaxonomyQuery`

NewTeamTaxonomyQueryWithDefaults instantiates a new TeamTaxonomyQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TeamTaxonomyQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TeamTaxonomyQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TeamTaxonomyQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TeamTaxonomyQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *TeamTaxonomyQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *TeamTaxonomyQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *TeamTaxonomyQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *TeamTaxonomyQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetResponse

`func (o *TeamTaxonomyQuery) GetResponse() TeamTaxonomyQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *TeamTaxonomyQuery) GetResponseOk() (*TeamTaxonomyQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *TeamTaxonomyQuery) SetResponse(v TeamTaxonomyQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *TeamTaxonomyQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *TeamTaxonomyQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TeamTaxonomyQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TeamTaxonomyQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TeamTaxonomyQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *TeamTaxonomyQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TeamTaxonomyQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TeamTaxonomyQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TeamTaxonomyQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *TeamTaxonomyQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *TeamTaxonomyQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


