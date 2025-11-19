# PropertyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Key of the property you&#39;re filtering on. For example &#x60;email&#x60; or &#x60;$current_url&#x60; | 
**Value** | **string** | Value of your filter. For example &#x60;test@example.com&#x60; or &#x60;https://example.com/test/&#x60;. Can be an array for an OR query, like &#x60;[\&quot;test@example.com\&quot;,\&quot;ok@example.com\&quot;]&#x60; | 
**Operator** | Pointer to **NullableString** |  | [optional] [default to "exact"]
**Type** | Pointer to **string** |  | [optional] [default to "event"]

## Methods

### NewPropertyItem

`func NewPropertyItem(key string, value string, ) *PropertyItem`

NewPropertyItem instantiates a new PropertyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyItemWithDefaults

`func NewPropertyItemWithDefaults() *PropertyItem`

NewPropertyItemWithDefaults instantiates a new PropertyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PropertyItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PropertyItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PropertyItem) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *PropertyItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PropertyItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PropertyItem) SetValue(v string)`

SetValue sets Value field to given value.


### GetOperator

`func (o *PropertyItem) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PropertyItem) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PropertyItem) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *PropertyItem) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *PropertyItem) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *PropertyItem) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetType

`func (o *PropertyItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropertyItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropertyItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PropertyItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


