# ErrorTrackingSuppressionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Filters** | **interface{}** |  | 
**OrderKey** | **int32** |  | 

## Methods

### NewErrorTrackingSuppressionRule

`func NewErrorTrackingSuppressionRule(id string, filters interface{}, orderKey int32, ) *ErrorTrackingSuppressionRule`

NewErrorTrackingSuppressionRule instantiates a new ErrorTrackingSuppressionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTrackingSuppressionRuleWithDefaults

`func NewErrorTrackingSuppressionRuleWithDefaults() *ErrorTrackingSuppressionRule`

NewErrorTrackingSuppressionRuleWithDefaults instantiates a new ErrorTrackingSuppressionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ErrorTrackingSuppressionRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorTrackingSuppressionRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorTrackingSuppressionRule) SetId(v string)`

SetId sets Id field to given value.


### GetFilters

`func (o *ErrorTrackingSuppressionRule) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ErrorTrackingSuppressionRule) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ErrorTrackingSuppressionRule) SetFilters(v interface{})`

SetFilters sets Filters field to given value.


### SetFiltersNil

`func (o *ErrorTrackingSuppressionRule) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *ErrorTrackingSuppressionRule) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetOrderKey

`func (o *ErrorTrackingSuppressionRule) GetOrderKey() int32`

GetOrderKey returns the OrderKey field if non-nil, zero value otherwise.

### GetOrderKeyOk

`func (o *ErrorTrackingSuppressionRule) GetOrderKeyOk() (*int32, bool)`

GetOrderKeyOk returns a tuple with the OrderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderKey

`func (o *ErrorTrackingSuppressionRule) SetOrderKey(v int32)`

SetOrderKey sets OrderKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


