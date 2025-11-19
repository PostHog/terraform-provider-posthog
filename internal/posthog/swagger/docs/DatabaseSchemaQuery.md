# DatabaseSchemaQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] [default to "DatabaseSchemaQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Response** | Pointer to [**DatabaseSchemaQueryResponse**](DatabaseSchemaQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewDatabaseSchemaQuery

`func NewDatabaseSchemaQuery() *DatabaseSchemaQuery`

NewDatabaseSchemaQuery instantiates a new DatabaseSchemaQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSchemaQueryWithDefaults

`func NewDatabaseSchemaQueryWithDefaults() *DatabaseSchemaQuery`

NewDatabaseSchemaQueryWithDefaults instantiates a new DatabaseSchemaQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *DatabaseSchemaQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DatabaseSchemaQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DatabaseSchemaQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DatabaseSchemaQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *DatabaseSchemaQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *DatabaseSchemaQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *DatabaseSchemaQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *DatabaseSchemaQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetResponse

`func (o *DatabaseSchemaQuery) GetResponse() DatabaseSchemaQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *DatabaseSchemaQuery) GetResponseOk() (*DatabaseSchemaQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *DatabaseSchemaQuery) SetResponse(v DatabaseSchemaQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *DatabaseSchemaQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *DatabaseSchemaQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DatabaseSchemaQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DatabaseSchemaQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DatabaseSchemaQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *DatabaseSchemaQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DatabaseSchemaQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DatabaseSchemaQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DatabaseSchemaQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DatabaseSchemaQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DatabaseSchemaQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


