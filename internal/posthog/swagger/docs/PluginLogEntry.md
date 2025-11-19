# PluginLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TeamId** | **int32** |  | 
**PluginId** | **int32** |  | 
**PluginConfigId** | **int32** |  | 
**Timestamp** | **time.Time** |  | 
**Source** | [**PluginLogEntrySourceEnum**](PluginLogEntrySourceEnum.md) |  | 
**Type** | [**PluginLogEntryTypeEnum**](PluginLogEntryTypeEnum.md) |  | 
**Message** | **string** |  | 
**InstanceId** | **string** |  | 

## Methods

### NewPluginLogEntry

`func NewPluginLogEntry(id string, teamId int32, pluginId int32, pluginConfigId int32, timestamp time.Time, source PluginLogEntrySourceEnum, type_ PluginLogEntryTypeEnum, message string, instanceId string, ) *PluginLogEntry`

NewPluginLogEntry instantiates a new PluginLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginLogEntryWithDefaults

`func NewPluginLogEntryWithDefaults() *PluginLogEntry`

NewPluginLogEntryWithDefaults instantiates a new PluginLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PluginLogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PluginLogEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PluginLogEntry) SetId(v string)`

SetId sets Id field to given value.


### GetTeamId

`func (o *PluginLogEntry) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *PluginLogEntry) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *PluginLogEntry) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.


### GetPluginId

`func (o *PluginLogEntry) GetPluginId() int32`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *PluginLogEntry) GetPluginIdOk() (*int32, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *PluginLogEntry) SetPluginId(v int32)`

SetPluginId sets PluginId field to given value.


### GetPluginConfigId

`func (o *PluginLogEntry) GetPluginConfigId() int32`

GetPluginConfigId returns the PluginConfigId field if non-nil, zero value otherwise.

### GetPluginConfigIdOk

`func (o *PluginLogEntry) GetPluginConfigIdOk() (*int32, bool)`

GetPluginConfigIdOk returns a tuple with the PluginConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginConfigId

`func (o *PluginLogEntry) SetPluginConfigId(v int32)`

SetPluginConfigId sets PluginConfigId field to given value.


### GetTimestamp

`func (o *PluginLogEntry) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PluginLogEntry) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PluginLogEntry) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSource

`func (o *PluginLogEntry) GetSource() PluginLogEntrySourceEnum`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PluginLogEntry) GetSourceOk() (*PluginLogEntrySourceEnum, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PluginLogEntry) SetSource(v PluginLogEntrySourceEnum)`

SetSource sets Source field to given value.


### GetType

`func (o *PluginLogEntry) GetType() PluginLogEntryTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PluginLogEntry) GetTypeOk() (*PluginLogEntryTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PluginLogEntry) SetType(v PluginLogEntryTypeEnum)`

SetType sets Type field to given value.


### GetMessage

`func (o *PluginLogEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PluginLogEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PluginLogEntry) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetInstanceId

`func (o *PluginLogEntry) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *PluginLogEntry) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *PluginLogEntry) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


