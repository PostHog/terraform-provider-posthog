# Property

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**PropertyTypeEnum**](PropertyTypeEnum.md) |   You can use a simplified version: &#x60;&#x60;&#x60;json {     \&quot;properties\&quot;: [         {             \&quot;key\&quot;: \&quot;email\&quot;,             \&quot;value\&quot;: \&quot;x@y.com\&quot;,             \&quot;operator\&quot;: \&quot;exact\&quot;,             \&quot;type\&quot;: \&quot;event\&quot;         }     ] } &#x60;&#x60;&#x60;  Or you can create more complicated queries with AND and OR: &#x60;&#x60;&#x60;json {     \&quot;properties\&quot;: {         \&quot;type\&quot;: \&quot;AND\&quot;,         \&quot;values\&quot;: [             {                 \&quot;type\&quot;: \&quot;OR\&quot;,                 \&quot;values\&quot;: [                     {\&quot;key\&quot;: \&quot;email\&quot;, ...},                     {\&quot;key\&quot;: \&quot;email\&quot;, ...}                 ]             },             {                 \&quot;type\&quot;: \&quot;AND\&quot;,                 \&quot;values\&quot;: [                     {\&quot;key\&quot;: \&quot;email\&quot;, ...},                     {\&quot;key\&quot;: \&quot;email\&quot;, ...}                 ]             }         ]     ] } &#x60;&#x60;&#x60;   * &#x60;AND&#x60; - AND * &#x60;OR&#x60; - OR | [optional] [default to "AND"]
**Values** | [**[]PropertyItem**](PropertyItem.md) |  | 

## Methods

### NewProperty

`func NewProperty(values []PropertyItem, ) *Property`

NewProperty instantiates a new Property object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyWithDefaults

`func NewPropertyWithDefaults() *Property`

NewPropertyWithDefaults instantiates a new Property object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Property) GetType() PropertyTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Property) GetTypeOk() (*PropertyTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Property) SetType(v PropertyTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *Property) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValues

`func (o *Property) GetValues() []PropertyItem`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Property) GetValuesOk() (*[]PropertyItem, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Property) SetValues(v []PropertyItem)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


