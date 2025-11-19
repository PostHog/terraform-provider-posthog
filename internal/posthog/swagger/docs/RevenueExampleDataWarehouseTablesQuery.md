# RevenueExampleDataWarehouseTablesQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] [default to "RevenueExampleDataWarehouseTablesQuery"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**Response** | Pointer to [**RevenueExampleDataWarehouseTablesQueryResponse**](RevenueExampleDataWarehouseTablesQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewRevenueExampleDataWarehouseTablesQuery

`func NewRevenueExampleDataWarehouseTablesQuery() *RevenueExampleDataWarehouseTablesQuery`

NewRevenueExampleDataWarehouseTablesQuery instantiates a new RevenueExampleDataWarehouseTablesQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueExampleDataWarehouseTablesQueryWithDefaults

`func NewRevenueExampleDataWarehouseTablesQueryWithDefaults() *RevenueExampleDataWarehouseTablesQuery`

NewRevenueExampleDataWarehouseTablesQueryWithDefaults instantiates a new RevenueExampleDataWarehouseTablesQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *RevenueExampleDataWarehouseTablesQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RevenueExampleDataWarehouseTablesQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RevenueExampleDataWarehouseTablesQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RevenueExampleDataWarehouseTablesQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *RevenueExampleDataWarehouseTablesQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RevenueExampleDataWarehouseTablesQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RevenueExampleDataWarehouseTablesQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RevenueExampleDataWarehouseTablesQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *RevenueExampleDataWarehouseTablesQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *RevenueExampleDataWarehouseTablesQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *RevenueExampleDataWarehouseTablesQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *RevenueExampleDataWarehouseTablesQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *RevenueExampleDataWarehouseTablesQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *RevenueExampleDataWarehouseTablesQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *RevenueExampleDataWarehouseTablesQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RevenueExampleDataWarehouseTablesQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RevenueExampleDataWarehouseTablesQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RevenueExampleDataWarehouseTablesQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *RevenueExampleDataWarehouseTablesQuery) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *RevenueExampleDataWarehouseTablesQuery) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetResponse

`func (o *RevenueExampleDataWarehouseTablesQuery) GetResponse() RevenueExampleDataWarehouseTablesQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *RevenueExampleDataWarehouseTablesQuery) GetResponseOk() (*RevenueExampleDataWarehouseTablesQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *RevenueExampleDataWarehouseTablesQuery) SetResponse(v RevenueExampleDataWarehouseTablesQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *RevenueExampleDataWarehouseTablesQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *RevenueExampleDataWarehouseTablesQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RevenueExampleDataWarehouseTablesQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RevenueExampleDataWarehouseTablesQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RevenueExampleDataWarehouseTablesQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *RevenueExampleDataWarehouseTablesQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RevenueExampleDataWarehouseTablesQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RevenueExampleDataWarehouseTablesQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RevenueExampleDataWarehouseTablesQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *RevenueExampleDataWarehouseTablesQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *RevenueExampleDataWarehouseTablesQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


