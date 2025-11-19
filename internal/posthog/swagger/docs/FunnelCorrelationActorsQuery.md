# FunnelCorrelationActorsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunnelCorrelationPersonConverted** | Pointer to **NullableBool** |  | [optional] 
**FunnelCorrelationPersonEntity** | Pointer to [**NullableFunnelcorrelationpersonentity**](Funnelcorrelationpersonentity.md) |  | [optional] [default to null]
**FunnelCorrelationPropertyValues** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) |  | [optional] 
**IncludeRecordings** | Pointer to **NullableBool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "FunnelCorrelationActorsQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Response** | Pointer to [**ActorsQueryResponse**](ActorsQueryResponse.md) |  | [optional] 
**Source** | [**FunnelCorrelationQuery**](FunnelCorrelationQuery.md) |  | 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewFunnelCorrelationActorsQuery

`func NewFunnelCorrelationActorsQuery(source FunnelCorrelationQuery, ) *FunnelCorrelationActorsQuery`

NewFunnelCorrelationActorsQuery instantiates a new FunnelCorrelationActorsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunnelCorrelationActorsQueryWithDefaults

`func NewFunnelCorrelationActorsQueryWithDefaults() *FunnelCorrelationActorsQuery`

NewFunnelCorrelationActorsQueryWithDefaults instantiates a new FunnelCorrelationActorsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunnelCorrelationPersonConverted

`func (o *FunnelCorrelationActorsQuery) GetFunnelCorrelationPersonConverted() bool`

GetFunnelCorrelationPersonConverted returns the FunnelCorrelationPersonConverted field if non-nil, zero value otherwise.

### GetFunnelCorrelationPersonConvertedOk

`func (o *FunnelCorrelationActorsQuery) GetFunnelCorrelationPersonConvertedOk() (*bool, bool)`

GetFunnelCorrelationPersonConvertedOk returns a tuple with the FunnelCorrelationPersonConverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationPersonConverted

`func (o *FunnelCorrelationActorsQuery) SetFunnelCorrelationPersonConverted(v bool)`

SetFunnelCorrelationPersonConverted sets FunnelCorrelationPersonConverted field to given value.

### HasFunnelCorrelationPersonConverted

`func (o *FunnelCorrelationActorsQuery) HasFunnelCorrelationPersonConverted() bool`

HasFunnelCorrelationPersonConverted returns a boolean if a field has been set.

### SetFunnelCorrelationPersonConvertedNil

`func (o *FunnelCorrelationActorsQuery) SetFunnelCorrelationPersonConvertedNil(b bool)`

 SetFunnelCorrelationPersonConvertedNil sets the value for FunnelCorrelationPersonConverted to be an explicit nil

### UnsetFunnelCorrelationPersonConverted
`func (o *FunnelCorrelationActorsQuery) UnsetFunnelCorrelationPersonConverted()`

UnsetFunnelCorrelationPersonConverted ensures that no value is present for FunnelCorrelationPersonConverted, not even an explicit nil
### GetFunnelCorrelationPersonEntity

`func (o *FunnelCorrelationActorsQuery) GetFunnelCorrelationPersonEntity() Funnelcorrelationpersonentity`

GetFunnelCorrelationPersonEntity returns the FunnelCorrelationPersonEntity field if non-nil, zero value otherwise.

### GetFunnelCorrelationPersonEntityOk

`func (o *FunnelCorrelationActorsQuery) GetFunnelCorrelationPersonEntityOk() (*Funnelcorrelationpersonentity, bool)`

GetFunnelCorrelationPersonEntityOk returns a tuple with the FunnelCorrelationPersonEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationPersonEntity

`func (o *FunnelCorrelationActorsQuery) SetFunnelCorrelationPersonEntity(v Funnelcorrelationpersonentity)`

SetFunnelCorrelationPersonEntity sets FunnelCorrelationPersonEntity field to given value.

### HasFunnelCorrelationPersonEntity

`func (o *FunnelCorrelationActorsQuery) HasFunnelCorrelationPersonEntity() bool`

HasFunnelCorrelationPersonEntity returns a boolean if a field has been set.

### SetFunnelCorrelationPersonEntityNil

`func (o *FunnelCorrelationActorsQuery) SetFunnelCorrelationPersonEntityNil(b bool)`

 SetFunnelCorrelationPersonEntityNil sets the value for FunnelCorrelationPersonEntity to be an explicit nil

### UnsetFunnelCorrelationPersonEntity
`func (o *FunnelCorrelationActorsQuery) UnsetFunnelCorrelationPersonEntity()`

UnsetFunnelCorrelationPersonEntity ensures that no value is present for FunnelCorrelationPersonEntity, not even an explicit nil
### GetFunnelCorrelationPropertyValues

`func (o *FunnelCorrelationActorsQuery) GetFunnelCorrelationPropertyValues() []FixedpropertiesInner`

GetFunnelCorrelationPropertyValues returns the FunnelCorrelationPropertyValues field if non-nil, zero value otherwise.

### GetFunnelCorrelationPropertyValuesOk

`func (o *FunnelCorrelationActorsQuery) GetFunnelCorrelationPropertyValuesOk() (*[]FixedpropertiesInner, bool)`

GetFunnelCorrelationPropertyValuesOk returns a tuple with the FunnelCorrelationPropertyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationPropertyValues

`func (o *FunnelCorrelationActorsQuery) SetFunnelCorrelationPropertyValues(v []FixedpropertiesInner)`

SetFunnelCorrelationPropertyValues sets FunnelCorrelationPropertyValues field to given value.

### HasFunnelCorrelationPropertyValues

`func (o *FunnelCorrelationActorsQuery) HasFunnelCorrelationPropertyValues() bool`

HasFunnelCorrelationPropertyValues returns a boolean if a field has been set.

### SetFunnelCorrelationPropertyValuesNil

`func (o *FunnelCorrelationActorsQuery) SetFunnelCorrelationPropertyValuesNil(b bool)`

 SetFunnelCorrelationPropertyValuesNil sets the value for FunnelCorrelationPropertyValues to be an explicit nil

### UnsetFunnelCorrelationPropertyValues
`func (o *FunnelCorrelationActorsQuery) UnsetFunnelCorrelationPropertyValues()`

UnsetFunnelCorrelationPropertyValues ensures that no value is present for FunnelCorrelationPropertyValues, not even an explicit nil
### GetIncludeRecordings

`func (o *FunnelCorrelationActorsQuery) GetIncludeRecordings() bool`

GetIncludeRecordings returns the IncludeRecordings field if non-nil, zero value otherwise.

### GetIncludeRecordingsOk

`func (o *FunnelCorrelationActorsQuery) GetIncludeRecordingsOk() (*bool, bool)`

GetIncludeRecordingsOk returns a tuple with the IncludeRecordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecordings

`func (o *FunnelCorrelationActorsQuery) SetIncludeRecordings(v bool)`

SetIncludeRecordings sets IncludeRecordings field to given value.

### HasIncludeRecordings

`func (o *FunnelCorrelationActorsQuery) HasIncludeRecordings() bool`

HasIncludeRecordings returns a boolean if a field has been set.

### SetIncludeRecordingsNil

`func (o *FunnelCorrelationActorsQuery) SetIncludeRecordingsNil(b bool)`

 SetIncludeRecordingsNil sets the value for IncludeRecordings to be an explicit nil

### UnsetIncludeRecordings
`func (o *FunnelCorrelationActorsQuery) UnsetIncludeRecordings()`

UnsetIncludeRecordings ensures that no value is present for IncludeRecordings, not even an explicit nil
### GetKind

`func (o *FunnelCorrelationActorsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FunnelCorrelationActorsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FunnelCorrelationActorsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FunnelCorrelationActorsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *FunnelCorrelationActorsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *FunnelCorrelationActorsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *FunnelCorrelationActorsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *FunnelCorrelationActorsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetResponse

`func (o *FunnelCorrelationActorsQuery) GetResponse() ActorsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *FunnelCorrelationActorsQuery) GetResponseOk() (*ActorsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *FunnelCorrelationActorsQuery) SetResponse(v ActorsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *FunnelCorrelationActorsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSource

`func (o *FunnelCorrelationActorsQuery) GetSource() FunnelCorrelationQuery`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FunnelCorrelationActorsQuery) GetSourceOk() (*FunnelCorrelationQuery, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FunnelCorrelationActorsQuery) SetSource(v FunnelCorrelationQuery)`

SetSource sets Source field to given value.


### GetTags

`func (o *FunnelCorrelationActorsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FunnelCorrelationActorsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FunnelCorrelationActorsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FunnelCorrelationActorsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *FunnelCorrelationActorsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FunnelCorrelationActorsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FunnelCorrelationActorsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FunnelCorrelationActorsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *FunnelCorrelationActorsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *FunnelCorrelationActorsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


