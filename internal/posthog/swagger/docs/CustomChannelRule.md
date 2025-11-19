# CustomChannelRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelType** | **string** |  | 
**Combiner** | [**FilterLogicalOperator**](FilterLogicalOperator.md) |  | 
**Id** | **string** |  | 
**Items** | [**[]CustomChannelCondition**](CustomChannelCondition.md) |  | 

## Methods

### NewCustomChannelRule

`func NewCustomChannelRule(channelType string, combiner FilterLogicalOperator, id string, items []CustomChannelCondition, ) *CustomChannelRule`

NewCustomChannelRule instantiates a new CustomChannelRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomChannelRuleWithDefaults

`func NewCustomChannelRuleWithDefaults() *CustomChannelRule`

NewCustomChannelRuleWithDefaults instantiates a new CustomChannelRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelType

`func (o *CustomChannelRule) GetChannelType() string`

GetChannelType returns the ChannelType field if non-nil, zero value otherwise.

### GetChannelTypeOk

`func (o *CustomChannelRule) GetChannelTypeOk() (*string, bool)`

GetChannelTypeOk returns a tuple with the ChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelType

`func (o *CustomChannelRule) SetChannelType(v string)`

SetChannelType sets ChannelType field to given value.


### GetCombiner

`func (o *CustomChannelRule) GetCombiner() FilterLogicalOperator`

GetCombiner returns the Combiner field if non-nil, zero value otherwise.

### GetCombinerOk

`func (o *CustomChannelRule) GetCombinerOk() (*FilterLogicalOperator, bool)`

GetCombinerOk returns a tuple with the Combiner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombiner

`func (o *CustomChannelRule) SetCombiner(v FilterLogicalOperator)`

SetCombiner sets Combiner field to given value.


### GetId

`func (o *CustomChannelRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomChannelRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomChannelRule) SetId(v string)`

SetId sets Id field to given value.


### GetItems

`func (o *CustomChannelRule) GetItems() []CustomChannelCondition`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CustomChannelRule) GetItemsOk() (*[]CustomChannelCondition, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CustomChannelRule) SetItems(v []CustomChannelCondition)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


