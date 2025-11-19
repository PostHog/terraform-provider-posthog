# FunnelPathsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunnelPathType** | Pointer to [**FunnelPathType**](FunnelPathType.md) |  | [optional] 
**FunnelSource** | [**FunnelsQuery**](FunnelsQuery.md) |  | 
**FunnelStep** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewFunnelPathsFilter

`func NewFunnelPathsFilter(funnelSource FunnelsQuery, ) *FunnelPathsFilter`

NewFunnelPathsFilter instantiates a new FunnelPathsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunnelPathsFilterWithDefaults

`func NewFunnelPathsFilterWithDefaults() *FunnelPathsFilter`

NewFunnelPathsFilterWithDefaults instantiates a new FunnelPathsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunnelPathType

`func (o *FunnelPathsFilter) GetFunnelPathType() FunnelPathType`

GetFunnelPathType returns the FunnelPathType field if non-nil, zero value otherwise.

### GetFunnelPathTypeOk

`func (o *FunnelPathsFilter) GetFunnelPathTypeOk() (*FunnelPathType, bool)`

GetFunnelPathTypeOk returns a tuple with the FunnelPathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelPathType

`func (o *FunnelPathsFilter) SetFunnelPathType(v FunnelPathType)`

SetFunnelPathType sets FunnelPathType field to given value.

### HasFunnelPathType

`func (o *FunnelPathsFilter) HasFunnelPathType() bool`

HasFunnelPathType returns a boolean if a field has been set.

### GetFunnelSource

`func (o *FunnelPathsFilter) GetFunnelSource() FunnelsQuery`

GetFunnelSource returns the FunnelSource field if non-nil, zero value otherwise.

### GetFunnelSourceOk

`func (o *FunnelPathsFilter) GetFunnelSourceOk() (*FunnelsQuery, bool)`

GetFunnelSourceOk returns a tuple with the FunnelSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelSource

`func (o *FunnelPathsFilter) SetFunnelSource(v FunnelsQuery)`

SetFunnelSource sets FunnelSource field to given value.


### GetFunnelStep

`func (o *FunnelPathsFilter) GetFunnelStep() int32`

GetFunnelStep returns the FunnelStep field if non-nil, zero value otherwise.

### GetFunnelStepOk

`func (o *FunnelPathsFilter) GetFunnelStepOk() (*int32, bool)`

GetFunnelStepOk returns a tuple with the FunnelStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelStep

`func (o *FunnelPathsFilter) SetFunnelStep(v int32)`

SetFunnelStep sets FunnelStep field to given value.

### HasFunnelStep

`func (o *FunnelPathsFilter) HasFunnelStep() bool`

HasFunnelStep returns a boolean if a field has been set.

### SetFunnelStepNil

`func (o *FunnelPathsFilter) SetFunnelStepNil(b bool)`

 SetFunnelStepNil sets the value for FunnelStep to be an explicit nil

### UnsetFunnelStep
`func (o *FunnelPathsFilter) UnsetFunnelStep()`

UnsetFunnelStep ensures that no value is present for FunnelStep, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


