# FunnelCorrelationQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunnelCorrelationEventExcludePropertyNames** | Pointer to **[]string** |  | [optional] 
**FunnelCorrelationEventNames** | Pointer to **[]string** |  | [optional] 
**FunnelCorrelationExcludeEventNames** | Pointer to **[]string** |  | [optional] 
**FunnelCorrelationExcludeNames** | Pointer to **[]string** |  | [optional] 
**FunnelCorrelationNames** | Pointer to **[]string** |  | [optional] 
**FunnelCorrelationType** | [**FunnelCorrelationResultsType**](FunnelCorrelationResultsType.md) |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "FunnelCorrelationQuery"]
**Response** | Pointer to [**FunnelCorrelationResponse**](FunnelCorrelationResponse.md) |  | [optional] 
**Source** | [**FunnelsActorsQuery**](FunnelsActorsQuery.md) |  | 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewFunnelCorrelationQuery

`func NewFunnelCorrelationQuery(funnelCorrelationType FunnelCorrelationResultsType, source FunnelsActorsQuery, ) *FunnelCorrelationQuery`

NewFunnelCorrelationQuery instantiates a new FunnelCorrelationQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunnelCorrelationQueryWithDefaults

`func NewFunnelCorrelationQueryWithDefaults() *FunnelCorrelationQuery`

NewFunnelCorrelationQueryWithDefaults instantiates a new FunnelCorrelationQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunnelCorrelationEventExcludePropertyNames

`func (o *FunnelCorrelationQuery) GetFunnelCorrelationEventExcludePropertyNames() []string`

GetFunnelCorrelationEventExcludePropertyNames returns the FunnelCorrelationEventExcludePropertyNames field if non-nil, zero value otherwise.

### GetFunnelCorrelationEventExcludePropertyNamesOk

`func (o *FunnelCorrelationQuery) GetFunnelCorrelationEventExcludePropertyNamesOk() (*[]string, bool)`

GetFunnelCorrelationEventExcludePropertyNamesOk returns a tuple with the FunnelCorrelationEventExcludePropertyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationEventExcludePropertyNames

`func (o *FunnelCorrelationQuery) SetFunnelCorrelationEventExcludePropertyNames(v []string)`

SetFunnelCorrelationEventExcludePropertyNames sets FunnelCorrelationEventExcludePropertyNames field to given value.

### HasFunnelCorrelationEventExcludePropertyNames

`func (o *FunnelCorrelationQuery) HasFunnelCorrelationEventExcludePropertyNames() bool`

HasFunnelCorrelationEventExcludePropertyNames returns a boolean if a field has been set.

### SetFunnelCorrelationEventExcludePropertyNamesNil

`func (o *FunnelCorrelationQuery) SetFunnelCorrelationEventExcludePropertyNamesNil(b bool)`

 SetFunnelCorrelationEventExcludePropertyNamesNil sets the value for FunnelCorrelationEventExcludePropertyNames to be an explicit nil

### UnsetFunnelCorrelationEventExcludePropertyNames
`func (o *FunnelCorrelationQuery) UnsetFunnelCorrelationEventExcludePropertyNames()`

UnsetFunnelCorrelationEventExcludePropertyNames ensures that no value is present for FunnelCorrelationEventExcludePropertyNames, not even an explicit nil
### GetFunnelCorrelationEventNames

`func (o *FunnelCorrelationQuery) GetFunnelCorrelationEventNames() []string`

GetFunnelCorrelationEventNames returns the FunnelCorrelationEventNames field if non-nil, zero value otherwise.

### GetFunnelCorrelationEventNamesOk

`func (o *FunnelCorrelationQuery) GetFunnelCorrelationEventNamesOk() (*[]string, bool)`

GetFunnelCorrelationEventNamesOk returns a tuple with the FunnelCorrelationEventNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationEventNames

`func (o *FunnelCorrelationQuery) SetFunnelCorrelationEventNames(v []string)`

SetFunnelCorrelationEventNames sets FunnelCorrelationEventNames field to given value.

### HasFunnelCorrelationEventNames

`func (o *FunnelCorrelationQuery) HasFunnelCorrelationEventNames() bool`

HasFunnelCorrelationEventNames returns a boolean if a field has been set.

### SetFunnelCorrelationEventNamesNil

`func (o *FunnelCorrelationQuery) SetFunnelCorrelationEventNamesNil(b bool)`

 SetFunnelCorrelationEventNamesNil sets the value for FunnelCorrelationEventNames to be an explicit nil

### UnsetFunnelCorrelationEventNames
`func (o *FunnelCorrelationQuery) UnsetFunnelCorrelationEventNames()`

UnsetFunnelCorrelationEventNames ensures that no value is present for FunnelCorrelationEventNames, not even an explicit nil
### GetFunnelCorrelationExcludeEventNames

`func (o *FunnelCorrelationQuery) GetFunnelCorrelationExcludeEventNames() []string`

GetFunnelCorrelationExcludeEventNames returns the FunnelCorrelationExcludeEventNames field if non-nil, zero value otherwise.

### GetFunnelCorrelationExcludeEventNamesOk

`func (o *FunnelCorrelationQuery) GetFunnelCorrelationExcludeEventNamesOk() (*[]string, bool)`

GetFunnelCorrelationExcludeEventNamesOk returns a tuple with the FunnelCorrelationExcludeEventNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationExcludeEventNames

`func (o *FunnelCorrelationQuery) SetFunnelCorrelationExcludeEventNames(v []string)`

SetFunnelCorrelationExcludeEventNames sets FunnelCorrelationExcludeEventNames field to given value.

### HasFunnelCorrelationExcludeEventNames

`func (o *FunnelCorrelationQuery) HasFunnelCorrelationExcludeEventNames() bool`

HasFunnelCorrelationExcludeEventNames returns a boolean if a field has been set.

### SetFunnelCorrelationExcludeEventNamesNil

`func (o *FunnelCorrelationQuery) SetFunnelCorrelationExcludeEventNamesNil(b bool)`

 SetFunnelCorrelationExcludeEventNamesNil sets the value for FunnelCorrelationExcludeEventNames to be an explicit nil

### UnsetFunnelCorrelationExcludeEventNames
`func (o *FunnelCorrelationQuery) UnsetFunnelCorrelationExcludeEventNames()`

UnsetFunnelCorrelationExcludeEventNames ensures that no value is present for FunnelCorrelationExcludeEventNames, not even an explicit nil
### GetFunnelCorrelationExcludeNames

`func (o *FunnelCorrelationQuery) GetFunnelCorrelationExcludeNames() []string`

GetFunnelCorrelationExcludeNames returns the FunnelCorrelationExcludeNames field if non-nil, zero value otherwise.

### GetFunnelCorrelationExcludeNamesOk

`func (o *FunnelCorrelationQuery) GetFunnelCorrelationExcludeNamesOk() (*[]string, bool)`

GetFunnelCorrelationExcludeNamesOk returns a tuple with the FunnelCorrelationExcludeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationExcludeNames

`func (o *FunnelCorrelationQuery) SetFunnelCorrelationExcludeNames(v []string)`

SetFunnelCorrelationExcludeNames sets FunnelCorrelationExcludeNames field to given value.

### HasFunnelCorrelationExcludeNames

`func (o *FunnelCorrelationQuery) HasFunnelCorrelationExcludeNames() bool`

HasFunnelCorrelationExcludeNames returns a boolean if a field has been set.

### SetFunnelCorrelationExcludeNamesNil

`func (o *FunnelCorrelationQuery) SetFunnelCorrelationExcludeNamesNil(b bool)`

 SetFunnelCorrelationExcludeNamesNil sets the value for FunnelCorrelationExcludeNames to be an explicit nil

### UnsetFunnelCorrelationExcludeNames
`func (o *FunnelCorrelationQuery) UnsetFunnelCorrelationExcludeNames()`

UnsetFunnelCorrelationExcludeNames ensures that no value is present for FunnelCorrelationExcludeNames, not even an explicit nil
### GetFunnelCorrelationNames

`func (o *FunnelCorrelationQuery) GetFunnelCorrelationNames() []string`

GetFunnelCorrelationNames returns the FunnelCorrelationNames field if non-nil, zero value otherwise.

### GetFunnelCorrelationNamesOk

`func (o *FunnelCorrelationQuery) GetFunnelCorrelationNamesOk() (*[]string, bool)`

GetFunnelCorrelationNamesOk returns a tuple with the FunnelCorrelationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationNames

`func (o *FunnelCorrelationQuery) SetFunnelCorrelationNames(v []string)`

SetFunnelCorrelationNames sets FunnelCorrelationNames field to given value.

### HasFunnelCorrelationNames

`func (o *FunnelCorrelationQuery) HasFunnelCorrelationNames() bool`

HasFunnelCorrelationNames returns a boolean if a field has been set.

### SetFunnelCorrelationNamesNil

`func (o *FunnelCorrelationQuery) SetFunnelCorrelationNamesNil(b bool)`

 SetFunnelCorrelationNamesNil sets the value for FunnelCorrelationNames to be an explicit nil

### UnsetFunnelCorrelationNames
`func (o *FunnelCorrelationQuery) UnsetFunnelCorrelationNames()`

UnsetFunnelCorrelationNames ensures that no value is present for FunnelCorrelationNames, not even an explicit nil
### GetFunnelCorrelationType

`func (o *FunnelCorrelationQuery) GetFunnelCorrelationType() FunnelCorrelationResultsType`

GetFunnelCorrelationType returns the FunnelCorrelationType field if non-nil, zero value otherwise.

### GetFunnelCorrelationTypeOk

`func (o *FunnelCorrelationQuery) GetFunnelCorrelationTypeOk() (*FunnelCorrelationResultsType, bool)`

GetFunnelCorrelationTypeOk returns a tuple with the FunnelCorrelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationType

`func (o *FunnelCorrelationQuery) SetFunnelCorrelationType(v FunnelCorrelationResultsType)`

SetFunnelCorrelationType sets FunnelCorrelationType field to given value.


### GetKind

`func (o *FunnelCorrelationQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FunnelCorrelationQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FunnelCorrelationQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FunnelCorrelationQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetResponse

`func (o *FunnelCorrelationQuery) GetResponse() FunnelCorrelationResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *FunnelCorrelationQuery) GetResponseOk() (*FunnelCorrelationResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *FunnelCorrelationQuery) SetResponse(v FunnelCorrelationResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *FunnelCorrelationQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSource

`func (o *FunnelCorrelationQuery) GetSource() FunnelsActorsQuery`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FunnelCorrelationQuery) GetSourceOk() (*FunnelsActorsQuery, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FunnelCorrelationQuery) SetSource(v FunnelsActorsQuery)`

SetSource sets Source field to given value.


### GetVersion

`func (o *FunnelCorrelationQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FunnelCorrelationQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FunnelCorrelationQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FunnelCorrelationQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *FunnelCorrelationQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *FunnelCorrelationQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


