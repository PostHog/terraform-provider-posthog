# EndpointRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientQueryId** | Pointer to **NullableString** | Client provided query ID. Can be used to retrieve the status or cancel the query. | [optional] 
**FiltersOverride** | Pointer to [**DashboardFilter**](DashboardFilter.md) |  | [optional] 
**QueryOverride** | Pointer to **map[string]interface{}** |  | [optional] 
**Refresh** | Pointer to [**RefreshType**](RefreshType.md) |  | [optional] 
**VariablesOverride** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**VariablesValues** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEndpointRunRequest

`func NewEndpointRunRequest() *EndpointRunRequest`

NewEndpointRunRequest instantiates a new EndpointRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointRunRequestWithDefaults

`func NewEndpointRunRequestWithDefaults() *EndpointRunRequest`

NewEndpointRunRequestWithDefaults instantiates a new EndpointRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientQueryId

`func (o *EndpointRunRequest) GetClientQueryId() string`

GetClientQueryId returns the ClientQueryId field if non-nil, zero value otherwise.

### GetClientQueryIdOk

`func (o *EndpointRunRequest) GetClientQueryIdOk() (*string, bool)`

GetClientQueryIdOk returns a tuple with the ClientQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientQueryId

`func (o *EndpointRunRequest) SetClientQueryId(v string)`

SetClientQueryId sets ClientQueryId field to given value.

### HasClientQueryId

`func (o *EndpointRunRequest) HasClientQueryId() bool`

HasClientQueryId returns a boolean if a field has been set.

### SetClientQueryIdNil

`func (o *EndpointRunRequest) SetClientQueryIdNil(b bool)`

 SetClientQueryIdNil sets the value for ClientQueryId to be an explicit nil

### UnsetClientQueryId
`func (o *EndpointRunRequest) UnsetClientQueryId()`

UnsetClientQueryId ensures that no value is present for ClientQueryId, not even an explicit nil
### GetFiltersOverride

`func (o *EndpointRunRequest) GetFiltersOverride() DashboardFilter`

GetFiltersOverride returns the FiltersOverride field if non-nil, zero value otherwise.

### GetFiltersOverrideOk

`func (o *EndpointRunRequest) GetFiltersOverrideOk() (*DashboardFilter, bool)`

GetFiltersOverrideOk returns a tuple with the FiltersOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersOverride

`func (o *EndpointRunRequest) SetFiltersOverride(v DashboardFilter)`

SetFiltersOverride sets FiltersOverride field to given value.

### HasFiltersOverride

`func (o *EndpointRunRequest) HasFiltersOverride() bool`

HasFiltersOverride returns a boolean if a field has been set.

### GetQueryOverride

`func (o *EndpointRunRequest) GetQueryOverride() map[string]interface{}`

GetQueryOverride returns the QueryOverride field if non-nil, zero value otherwise.

### GetQueryOverrideOk

`func (o *EndpointRunRequest) GetQueryOverrideOk() (*map[string]interface{}, bool)`

GetQueryOverrideOk returns a tuple with the QueryOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOverride

`func (o *EndpointRunRequest) SetQueryOverride(v map[string]interface{})`

SetQueryOverride sets QueryOverride field to given value.

### HasQueryOverride

`func (o *EndpointRunRequest) HasQueryOverride() bool`

HasQueryOverride returns a boolean if a field has been set.

### SetQueryOverrideNil

`func (o *EndpointRunRequest) SetQueryOverrideNil(b bool)`

 SetQueryOverrideNil sets the value for QueryOverride to be an explicit nil

### UnsetQueryOverride
`func (o *EndpointRunRequest) UnsetQueryOverride()`

UnsetQueryOverride ensures that no value is present for QueryOverride, not even an explicit nil
### GetRefresh

`func (o *EndpointRunRequest) GetRefresh() RefreshType`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *EndpointRunRequest) GetRefreshOk() (*RefreshType, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *EndpointRunRequest) SetRefresh(v RefreshType)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *EndpointRunRequest) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetVariablesOverride

`func (o *EndpointRunRequest) GetVariablesOverride() map[string]map[string]interface{}`

GetVariablesOverride returns the VariablesOverride field if non-nil, zero value otherwise.

### GetVariablesOverrideOk

`func (o *EndpointRunRequest) GetVariablesOverrideOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOverrideOk returns a tuple with the VariablesOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesOverride

`func (o *EndpointRunRequest) SetVariablesOverride(v map[string]map[string]interface{})`

SetVariablesOverride sets VariablesOverride field to given value.

### HasVariablesOverride

`func (o *EndpointRunRequest) HasVariablesOverride() bool`

HasVariablesOverride returns a boolean if a field has been set.

### SetVariablesOverrideNil

`func (o *EndpointRunRequest) SetVariablesOverrideNil(b bool)`

 SetVariablesOverrideNil sets the value for VariablesOverride to be an explicit nil

### UnsetVariablesOverride
`func (o *EndpointRunRequest) UnsetVariablesOverride()`

UnsetVariablesOverride ensures that no value is present for VariablesOverride, not even an explicit nil
### GetVariablesValues

`func (o *EndpointRunRequest) GetVariablesValues() map[string]interface{}`

GetVariablesValues returns the VariablesValues field if non-nil, zero value otherwise.

### GetVariablesValuesOk

`func (o *EndpointRunRequest) GetVariablesValuesOk() (*map[string]interface{}, bool)`

GetVariablesValuesOk returns a tuple with the VariablesValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesValues

`func (o *EndpointRunRequest) SetVariablesValues(v map[string]interface{})`

SetVariablesValues sets VariablesValues field to given value.

### HasVariablesValues

`func (o *EndpointRunRequest) HasVariablesValues() bool`

HasVariablesValues returns a boolean if a field has been set.

### SetVariablesValuesNil

`func (o *EndpointRunRequest) SetVariablesValuesNil(b bool)`

 SetVariablesValuesNil sets the value for VariablesValues to be an explicit nil

### UnsetVariablesValues
`func (o *EndpointRunRequest) UnsetVariablesValues()`

UnsetVariablesValues ensures that no value is present for VariablesValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


