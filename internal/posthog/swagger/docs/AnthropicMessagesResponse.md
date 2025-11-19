# AnthropicMessagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**AnthropicMessagesResponseTypeEnum**](AnthropicMessagesResponseTypeEnum.md) |  | 
**Role** | [**AnthropicMessagesResponseRoleEnum**](AnthropicMessagesResponseRoleEnum.md) |  | 
**Content** | **[]map[string]interface{}** |  | 
**Model** | **string** |  | 
**StopReason** | **NullableString** |  | 
**StopSequence** | Pointer to **NullableString** |  | [optional] 
**Usage** | [**AnthropicUsage**](AnthropicUsage.md) |  | 

## Methods

### NewAnthropicMessagesResponse

`func NewAnthropicMessagesResponse(id string, type_ AnthropicMessagesResponseTypeEnum, role AnthropicMessagesResponseRoleEnum, content []map[string]interface{}, model string, stopReason NullableString, usage AnthropicUsage, ) *AnthropicMessagesResponse`

NewAnthropicMessagesResponse instantiates a new AnthropicMessagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnthropicMessagesResponseWithDefaults

`func NewAnthropicMessagesResponseWithDefaults() *AnthropicMessagesResponse`

NewAnthropicMessagesResponseWithDefaults instantiates a new AnthropicMessagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnthropicMessagesResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnthropicMessagesResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnthropicMessagesResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *AnthropicMessagesResponse) GetType() AnthropicMessagesResponseTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnthropicMessagesResponse) GetTypeOk() (*AnthropicMessagesResponseTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnthropicMessagesResponse) SetType(v AnthropicMessagesResponseTypeEnum)`

SetType sets Type field to given value.


### GetRole

`func (o *AnthropicMessagesResponse) GetRole() AnthropicMessagesResponseRoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AnthropicMessagesResponse) GetRoleOk() (*AnthropicMessagesResponseRoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AnthropicMessagesResponse) SetRole(v AnthropicMessagesResponseRoleEnum)`

SetRole sets Role field to given value.


### GetContent

`func (o *AnthropicMessagesResponse) GetContent() []map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AnthropicMessagesResponse) GetContentOk() (*[]map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AnthropicMessagesResponse) SetContent(v []map[string]interface{})`

SetContent sets Content field to given value.


### GetModel

`func (o *AnthropicMessagesResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AnthropicMessagesResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AnthropicMessagesResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetStopReason

`func (o *AnthropicMessagesResponse) GetStopReason() string`

GetStopReason returns the StopReason field if non-nil, zero value otherwise.

### GetStopReasonOk

`func (o *AnthropicMessagesResponse) GetStopReasonOk() (*string, bool)`

GetStopReasonOk returns a tuple with the StopReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopReason

`func (o *AnthropicMessagesResponse) SetStopReason(v string)`

SetStopReason sets StopReason field to given value.


### SetStopReasonNil

`func (o *AnthropicMessagesResponse) SetStopReasonNil(b bool)`

 SetStopReasonNil sets the value for StopReason to be an explicit nil

### UnsetStopReason
`func (o *AnthropicMessagesResponse) UnsetStopReason()`

UnsetStopReason ensures that no value is present for StopReason, not even an explicit nil
### GetStopSequence

`func (o *AnthropicMessagesResponse) GetStopSequence() string`

GetStopSequence returns the StopSequence field if non-nil, zero value otherwise.

### GetStopSequenceOk

`func (o *AnthropicMessagesResponse) GetStopSequenceOk() (*string, bool)`

GetStopSequenceOk returns a tuple with the StopSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopSequence

`func (o *AnthropicMessagesResponse) SetStopSequence(v string)`

SetStopSequence sets StopSequence field to given value.

### HasStopSequence

`func (o *AnthropicMessagesResponse) HasStopSequence() bool`

HasStopSequence returns a boolean if a field has been set.

### SetStopSequenceNil

`func (o *AnthropicMessagesResponse) SetStopSequenceNil(b bool)`

 SetStopSequenceNil sets the value for StopSequence to be an explicit nil

### UnsetStopSequence
`func (o *AnthropicMessagesResponse) UnsetStopSequence()`

UnsetStopSequence ensures that no value is present for StopSequence, not even an explicit nil
### GetUsage

`func (o *AnthropicMessagesResponse) GetUsage() AnthropicUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AnthropicMessagesResponse) GetUsageOk() (*AnthropicUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AnthropicMessagesResponse) SetUsage(v AnthropicUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


