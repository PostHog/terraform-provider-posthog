# QueryResponseAlternative29

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string]interface{}** |  | 
**Error** | Pointer to [**ExternalQueryError**](ExternalQueryError.md) |  | [optional] 
**Status** | [**ExternalQueryStatus**](ExternalQueryStatus.md) |  | 

## Methods

### NewQueryResponseAlternative29

`func NewQueryResponseAlternative29(data map[string]interface{}, status ExternalQueryStatus, ) *QueryResponseAlternative29`

NewQueryResponseAlternative29 instantiates a new QueryResponseAlternative29 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseAlternative29WithDefaults

`func NewQueryResponseAlternative29WithDefaults() *QueryResponseAlternative29`

NewQueryResponseAlternative29WithDefaults instantiates a new QueryResponseAlternative29 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *QueryResponseAlternative29) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QueryResponseAlternative29) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QueryResponseAlternative29) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetError

`func (o *QueryResponseAlternative29) GetError() ExternalQueryError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *QueryResponseAlternative29) GetErrorOk() (*ExternalQueryError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *QueryResponseAlternative29) SetError(v ExternalQueryError)`

SetError sets Error field to given value.

### HasError

`func (o *QueryResponseAlternative29) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStatus

`func (o *QueryResponseAlternative29) GetStatus() ExternalQueryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryResponseAlternative29) GetStatusOk() (*ExternalQueryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryResponseAlternative29) SetStatus(v ExternalQueryStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


