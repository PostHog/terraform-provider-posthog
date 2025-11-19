# ActorsPropertyTaxonomyQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupTypeIndex** | Pointer to **NullableInt32** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "ActorsPropertyTaxonomyQuery"]
**MaxPropertyValues** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Properties** | **[]string** |  | 
**Response** | Pointer to [**ActorsPropertyTaxonomyQueryResponse**](ActorsPropertyTaxonomyQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewActorsPropertyTaxonomyQuery

`func NewActorsPropertyTaxonomyQuery(properties []string, ) *ActorsPropertyTaxonomyQuery`

NewActorsPropertyTaxonomyQuery instantiates a new ActorsPropertyTaxonomyQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorsPropertyTaxonomyQueryWithDefaults

`func NewActorsPropertyTaxonomyQueryWithDefaults() *ActorsPropertyTaxonomyQuery`

NewActorsPropertyTaxonomyQueryWithDefaults instantiates a new ActorsPropertyTaxonomyQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupTypeIndex

`func (o *ActorsPropertyTaxonomyQuery) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *ActorsPropertyTaxonomyQuery) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *ActorsPropertyTaxonomyQuery) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.

### HasGroupTypeIndex

`func (o *ActorsPropertyTaxonomyQuery) HasGroupTypeIndex() bool`

HasGroupTypeIndex returns a boolean if a field has been set.

### SetGroupTypeIndexNil

`func (o *ActorsPropertyTaxonomyQuery) SetGroupTypeIndexNil(b bool)`

 SetGroupTypeIndexNil sets the value for GroupTypeIndex to be an explicit nil

### UnsetGroupTypeIndex
`func (o *ActorsPropertyTaxonomyQuery) UnsetGroupTypeIndex()`

UnsetGroupTypeIndex ensures that no value is present for GroupTypeIndex, not even an explicit nil
### GetKind

`func (o *ActorsPropertyTaxonomyQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ActorsPropertyTaxonomyQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ActorsPropertyTaxonomyQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ActorsPropertyTaxonomyQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMaxPropertyValues

`func (o *ActorsPropertyTaxonomyQuery) GetMaxPropertyValues() int32`

GetMaxPropertyValues returns the MaxPropertyValues field if non-nil, zero value otherwise.

### GetMaxPropertyValuesOk

`func (o *ActorsPropertyTaxonomyQuery) GetMaxPropertyValuesOk() (*int32, bool)`

GetMaxPropertyValuesOk returns a tuple with the MaxPropertyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPropertyValues

`func (o *ActorsPropertyTaxonomyQuery) SetMaxPropertyValues(v int32)`

SetMaxPropertyValues sets MaxPropertyValues field to given value.

### HasMaxPropertyValues

`func (o *ActorsPropertyTaxonomyQuery) HasMaxPropertyValues() bool`

HasMaxPropertyValues returns a boolean if a field has been set.

### SetMaxPropertyValuesNil

`func (o *ActorsPropertyTaxonomyQuery) SetMaxPropertyValuesNil(b bool)`

 SetMaxPropertyValuesNil sets the value for MaxPropertyValues to be an explicit nil

### UnsetMaxPropertyValues
`func (o *ActorsPropertyTaxonomyQuery) UnsetMaxPropertyValues()`

UnsetMaxPropertyValues ensures that no value is present for MaxPropertyValues, not even an explicit nil
### GetModifiers

`func (o *ActorsPropertyTaxonomyQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *ActorsPropertyTaxonomyQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *ActorsPropertyTaxonomyQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *ActorsPropertyTaxonomyQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetProperties

`func (o *ActorsPropertyTaxonomyQuery) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ActorsPropertyTaxonomyQuery) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ActorsPropertyTaxonomyQuery) SetProperties(v []string)`

SetProperties sets Properties field to given value.


### GetResponse

`func (o *ActorsPropertyTaxonomyQuery) GetResponse() ActorsPropertyTaxonomyQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ActorsPropertyTaxonomyQuery) GetResponseOk() (*ActorsPropertyTaxonomyQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ActorsPropertyTaxonomyQuery) SetResponse(v ActorsPropertyTaxonomyQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ActorsPropertyTaxonomyQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *ActorsPropertyTaxonomyQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ActorsPropertyTaxonomyQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ActorsPropertyTaxonomyQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ActorsPropertyTaxonomyQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *ActorsPropertyTaxonomyQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ActorsPropertyTaxonomyQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ActorsPropertyTaxonomyQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ActorsPropertyTaxonomyQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ActorsPropertyTaxonomyQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ActorsPropertyTaxonomyQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


