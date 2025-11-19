# ErrorTrackingGroupingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Filters** | **interface{}** |  | 
**Assignee** | **string** |  | [readonly] 
**OrderKey** | **int32** |  | 
**DisabledData** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewErrorTrackingGroupingRule

`func NewErrorTrackingGroupingRule(id string, filters interface{}, assignee string, orderKey int32, ) *ErrorTrackingGroupingRule`

NewErrorTrackingGroupingRule instantiates a new ErrorTrackingGroupingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTrackingGroupingRuleWithDefaults

`func NewErrorTrackingGroupingRuleWithDefaults() *ErrorTrackingGroupingRule`

NewErrorTrackingGroupingRuleWithDefaults instantiates a new ErrorTrackingGroupingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ErrorTrackingGroupingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorTrackingGroupingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorTrackingGroupingRule) SetId(v string)`

SetId sets Id field to given value.


### GetFilters

`func (o *ErrorTrackingGroupingRule) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ErrorTrackingGroupingRule) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ErrorTrackingGroupingRule) SetFilters(v interface{})`

SetFilters sets Filters field to given value.


### SetFiltersNil

`func (o *ErrorTrackingGroupingRule) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *ErrorTrackingGroupingRule) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetAssignee

`func (o *ErrorTrackingGroupingRule) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ErrorTrackingGroupingRule) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ErrorTrackingGroupingRule) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.


### GetOrderKey

`func (o *ErrorTrackingGroupingRule) GetOrderKey() int32`

GetOrderKey returns the OrderKey field if non-nil, zero value otherwise.

### GetOrderKeyOk

`func (o *ErrorTrackingGroupingRule) GetOrderKeyOk() (*int32, bool)`

GetOrderKeyOk returns a tuple with the OrderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderKey

`func (o *ErrorTrackingGroupingRule) SetOrderKey(v int32)`

SetOrderKey sets OrderKey field to given value.


### GetDisabledData

`func (o *ErrorTrackingGroupingRule) GetDisabledData() interface{}`

GetDisabledData returns the DisabledData field if non-nil, zero value otherwise.

### GetDisabledDataOk

`func (o *ErrorTrackingGroupingRule) GetDisabledDataOk() (*interface{}, bool)`

GetDisabledDataOk returns a tuple with the DisabledData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledData

`func (o *ErrorTrackingGroupingRule) SetDisabledData(v interface{})`

SetDisabledData sets DisabledData field to given value.

### HasDisabledData

`func (o *ErrorTrackingGroupingRule) HasDisabledData() bool`

HasDisabledData returns a boolean if a field has been set.

### SetDisabledDataNil

`func (o *ErrorTrackingGroupingRule) SetDisabledDataNil(b bool)`

 SetDisabledDataNil sets the value for DisabledData to be an explicit nil

### UnsetDisabledData
`func (o *ErrorTrackingGroupingRule) UnsetDisabledData()`

UnsetDisabledData ensures that no value is present for DisabledData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


