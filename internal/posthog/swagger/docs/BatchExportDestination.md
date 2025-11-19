# BatchExportDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**BatchExportDestinationTypeEnum**](BatchExportDestinationTypeEnum.md) | A choice of supported BatchExportDestination types.  * &#x60;S3&#x60; - S3 * &#x60;Snowflake&#x60; - Snowflake * &#x60;Postgres&#x60; - Postgres * &#x60;Redshift&#x60; - Redshift * &#x60;BigQuery&#x60; - Bigquery * &#x60;Databricks&#x60; - Databricks * &#x60;HTTP&#x60; - Http * &#x60;NoOp&#x60; - Noop | 
**Config** | Pointer to **interface{}** | A JSON field to store all configuration parameters required to access a BatchExportDestination. | [optional] 
**Integration** | Pointer to **NullableInt32** | The integration for this destination. | [optional] 
**IntegrationId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBatchExportDestination

`func NewBatchExportDestination(type_ BatchExportDestinationTypeEnum, ) *BatchExportDestination`

NewBatchExportDestination instantiates a new BatchExportDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchExportDestinationWithDefaults

`func NewBatchExportDestinationWithDefaults() *BatchExportDestination`

NewBatchExportDestinationWithDefaults instantiates a new BatchExportDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BatchExportDestination) GetType() BatchExportDestinationTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BatchExportDestination) GetTypeOk() (*BatchExportDestinationTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BatchExportDestination) SetType(v BatchExportDestinationTypeEnum)`

SetType sets Type field to given value.


### GetConfig

`func (o *BatchExportDestination) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BatchExportDestination) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BatchExportDestination) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *BatchExportDestination) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *BatchExportDestination) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *BatchExportDestination) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetIntegration

`func (o *BatchExportDestination) GetIntegration() int32`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *BatchExportDestination) GetIntegrationOk() (*int32, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *BatchExportDestination) SetIntegration(v int32)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *BatchExportDestination) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegrationNil

`func (o *BatchExportDestination) SetIntegrationNil(b bool)`

 SetIntegrationNil sets the value for Integration to be an explicit nil

### UnsetIntegration
`func (o *BatchExportDestination) UnsetIntegration()`

UnsetIntegration ensures that no value is present for Integration, not even an explicit nil
### GetIntegrationId

`func (o *BatchExportDestination) GetIntegrationId() int32`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *BatchExportDestination) GetIntegrationIdOk() (*int32, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *BatchExportDestination) SetIntegrationId(v int32)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *BatchExportDestination) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationIdNil

`func (o *BatchExportDestination) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *BatchExportDestination) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


