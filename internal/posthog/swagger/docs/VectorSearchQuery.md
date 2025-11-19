# VectorSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedding** | **[]float32** |  | 
**EmbeddingVersion** | Pointer to **NullableFloat32** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "VectorSearchQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Response** | Pointer to [**VectorSearchQueryResponse**](VectorSearchQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewVectorSearchQuery

`func NewVectorSearchQuery(embedding []float32, ) *VectorSearchQuery`

NewVectorSearchQuery instantiates a new VectorSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorSearchQueryWithDefaults

`func NewVectorSearchQueryWithDefaults() *VectorSearchQuery`

NewVectorSearchQueryWithDefaults instantiates a new VectorSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedding

`func (o *VectorSearchQuery) GetEmbedding() []float32`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *VectorSearchQuery) GetEmbeddingOk() (*[]float32, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *VectorSearchQuery) SetEmbedding(v []float32)`

SetEmbedding sets Embedding field to given value.


### GetEmbeddingVersion

`func (o *VectorSearchQuery) GetEmbeddingVersion() float32`

GetEmbeddingVersion returns the EmbeddingVersion field if non-nil, zero value otherwise.

### GetEmbeddingVersionOk

`func (o *VectorSearchQuery) GetEmbeddingVersionOk() (*float32, bool)`

GetEmbeddingVersionOk returns a tuple with the EmbeddingVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingVersion

`func (o *VectorSearchQuery) SetEmbeddingVersion(v float32)`

SetEmbeddingVersion sets EmbeddingVersion field to given value.

### HasEmbeddingVersion

`func (o *VectorSearchQuery) HasEmbeddingVersion() bool`

HasEmbeddingVersion returns a boolean if a field has been set.

### SetEmbeddingVersionNil

`func (o *VectorSearchQuery) SetEmbeddingVersionNil(b bool)`

 SetEmbeddingVersionNil sets the value for EmbeddingVersion to be an explicit nil

### UnsetEmbeddingVersion
`func (o *VectorSearchQuery) UnsetEmbeddingVersion()`

UnsetEmbeddingVersion ensures that no value is present for EmbeddingVersion, not even an explicit nil
### GetKind

`func (o *VectorSearchQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VectorSearchQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VectorSearchQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *VectorSearchQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *VectorSearchQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *VectorSearchQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *VectorSearchQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *VectorSearchQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetResponse

`func (o *VectorSearchQuery) GetResponse() VectorSearchQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *VectorSearchQuery) GetResponseOk() (*VectorSearchQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *VectorSearchQuery) SetResponse(v VectorSearchQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *VectorSearchQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *VectorSearchQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VectorSearchQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VectorSearchQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VectorSearchQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *VectorSearchQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VectorSearchQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VectorSearchQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VectorSearchQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *VectorSearchQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *VectorSearchQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


