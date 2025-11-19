# QueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Async** | Pointer to **NullableBool** |  | [optional] 
**ClientQueryId** | Pointer to **NullableString** | Client provided query ID. Can be used to retrieve the status or cancel the query. | [optional] 
**FiltersOverride** | Pointer to [**DashboardFilter**](DashboardFilter.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Name given to a query. It&#39;s used to identify the query in the UI. Up to 128 characters for a name. | [optional] 
**Query** | [**Query1**](Query1.md) |  | 
**Refresh** | Pointer to [**RefreshType**](RefreshType.md) |  | [optional] 
**VariablesOverride** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewQueryRequest

`func NewQueryRequest(query Query1, ) *QueryRequest`

NewQueryRequest instantiates a new QueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryRequestWithDefaults

`func NewQueryRequestWithDefaults() *QueryRequest`

NewQueryRequestWithDefaults instantiates a new QueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsync

`func (o *QueryRequest) GetAsync() bool`

GetAsync returns the Async field if non-nil, zero value otherwise.

### GetAsyncOk

`func (o *QueryRequest) GetAsyncOk() (*bool, bool)`

GetAsyncOk returns a tuple with the Async field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsync

`func (o *QueryRequest) SetAsync(v bool)`

SetAsync sets Async field to given value.

### HasAsync

`func (o *QueryRequest) HasAsync() bool`

HasAsync returns a boolean if a field has been set.

### SetAsyncNil

`func (o *QueryRequest) SetAsyncNil(b bool)`

 SetAsyncNil sets the value for Async to be an explicit nil

### UnsetAsync
`func (o *QueryRequest) UnsetAsync()`

UnsetAsync ensures that no value is present for Async, not even an explicit nil
### GetClientQueryId

`func (o *QueryRequest) GetClientQueryId() string`

GetClientQueryId returns the ClientQueryId field if non-nil, zero value otherwise.

### GetClientQueryIdOk

`func (o *QueryRequest) GetClientQueryIdOk() (*string, bool)`

GetClientQueryIdOk returns a tuple with the ClientQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientQueryId

`func (o *QueryRequest) SetClientQueryId(v string)`

SetClientQueryId sets ClientQueryId field to given value.

### HasClientQueryId

`func (o *QueryRequest) HasClientQueryId() bool`

HasClientQueryId returns a boolean if a field has been set.

### SetClientQueryIdNil

`func (o *QueryRequest) SetClientQueryIdNil(b bool)`

 SetClientQueryIdNil sets the value for ClientQueryId to be an explicit nil

### UnsetClientQueryId
`func (o *QueryRequest) UnsetClientQueryId()`

UnsetClientQueryId ensures that no value is present for ClientQueryId, not even an explicit nil
### GetFiltersOverride

`func (o *QueryRequest) GetFiltersOverride() DashboardFilter`

GetFiltersOverride returns the FiltersOverride field if non-nil, zero value otherwise.

### GetFiltersOverrideOk

`func (o *QueryRequest) GetFiltersOverrideOk() (*DashboardFilter, bool)`

GetFiltersOverrideOk returns a tuple with the FiltersOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersOverride

`func (o *QueryRequest) SetFiltersOverride(v DashboardFilter)`

SetFiltersOverride sets FiltersOverride field to given value.

### HasFiltersOverride

`func (o *QueryRequest) HasFiltersOverride() bool`

HasFiltersOverride returns a boolean if a field has been set.

### GetName

`func (o *QueryRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QueryRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *QueryRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *QueryRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuery

`func (o *QueryRequest) GetQuery() Query1`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryRequest) GetQueryOk() (*Query1, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryRequest) SetQuery(v Query1)`

SetQuery sets Query field to given value.


### GetRefresh

`func (o *QueryRequest) GetRefresh() RefreshType`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *QueryRequest) GetRefreshOk() (*RefreshType, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *QueryRequest) SetRefresh(v RefreshType)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *QueryRequest) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetVariablesOverride

`func (o *QueryRequest) GetVariablesOverride() map[string]map[string]interface{}`

GetVariablesOverride returns the VariablesOverride field if non-nil, zero value otherwise.

### GetVariablesOverrideOk

`func (o *QueryRequest) GetVariablesOverrideOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOverrideOk returns a tuple with the VariablesOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesOverride

`func (o *QueryRequest) SetVariablesOverride(v map[string]map[string]interface{})`

SetVariablesOverride sets VariablesOverride field to given value.

### HasVariablesOverride

`func (o *QueryRequest) HasVariablesOverride() bool`

HasVariablesOverride returns a boolean if a field has been set.

### SetVariablesOverrideNil

`func (o *QueryRequest) SetVariablesOverrideNil(b bool)`

 SetVariablesOverrideNil sets the value for VariablesOverride to be an explicit nil

### UnsetVariablesOverride
`func (o *QueryRequest) UnsetVariablesOverride()`

UnsetVariablesOverride ensures that no value is present for VariablesOverride, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


