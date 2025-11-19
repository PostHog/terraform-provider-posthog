# InsightActorsQueryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] [default to "InsightActorsQueryOptions"]
**Response** | Pointer to [**InsightActorsQueryOptionsResponse**](InsightActorsQueryOptionsResponse.md) |  | [optional] 
**Source** | [**Source4**](Source4.md) |  | 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewInsightActorsQueryOptions

`func NewInsightActorsQueryOptions(source Source4, ) *InsightActorsQueryOptions`

NewInsightActorsQueryOptions instantiates a new InsightActorsQueryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightActorsQueryOptionsWithDefaults

`func NewInsightActorsQueryOptionsWithDefaults() *InsightActorsQueryOptions`

NewInsightActorsQueryOptionsWithDefaults instantiates a new InsightActorsQueryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *InsightActorsQueryOptions) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InsightActorsQueryOptions) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InsightActorsQueryOptions) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *InsightActorsQueryOptions) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetResponse

`func (o *InsightActorsQueryOptions) GetResponse() InsightActorsQueryOptionsResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *InsightActorsQueryOptions) GetResponseOk() (*InsightActorsQueryOptionsResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *InsightActorsQueryOptions) SetResponse(v InsightActorsQueryOptionsResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *InsightActorsQueryOptions) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSource

`func (o *InsightActorsQueryOptions) GetSource() Source4`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InsightActorsQueryOptions) GetSourceOk() (*Source4, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InsightActorsQueryOptions) SetSource(v Source4)`

SetSource sets Source field to given value.


### GetVersion

`func (o *InsightActorsQueryOptions) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InsightActorsQueryOptions) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InsightActorsQueryOptions) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InsightActorsQueryOptions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *InsightActorsQueryOptions) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *InsightActorsQueryOptions) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


