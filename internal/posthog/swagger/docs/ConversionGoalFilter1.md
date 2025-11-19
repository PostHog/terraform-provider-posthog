# ConversionGoalFilter1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversionGoalId** | **string** |  | 
**ConversionGoalName** | **string** |  | 
**CustomName** | Pointer to **NullableString** |  | [optional] 
**Event** | Pointer to **NullableString** | The event or &#x60;null&#x60; for all events. | [optional] 
**FixedProperties** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) | Fixed properties in the query, can&#39;t be edited in the interface (e.g. scoping down by person) | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "EventsNode"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Math** | Pointer to **NullableString** |  | [optional] [default to "null"]
**MathGroupTypeIndex** | Pointer to [**MathGroupTypeIndex**](MathGroupTypeIndex.md) |  | [optional] 
**MathHogql** | Pointer to **NullableString** |  | [optional] 
**MathMultiplier** | Pointer to **NullableFloat32** |  | [optional] 
**MathProperty** | Pointer to **NullableString** |  | [optional] 
**MathPropertyRevenueCurrency** | Pointer to [**RevenueCurrencyPropertyConfig**](RevenueCurrencyPropertyConfig.md) |  | [optional] 
**MathPropertyType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OptionalInFunnel** | Pointer to **NullableBool** |  | [optional] 
**OrderBy** | Pointer to **[]string** | Columns to order by | [optional] 
**Properties** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) | Properties configurable in the interface | [optional] 
**Response** | Pointer to **map[string]interface{}** |  | [optional] 
**SchemaMap** | [**map[string]SchemaMapValue**](SchemaMapValue.md) |  | 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewConversionGoalFilter1

`func NewConversionGoalFilter1(conversionGoalId string, conversionGoalName string, schemaMap map[string]SchemaMapValue, ) *ConversionGoalFilter1`

NewConversionGoalFilter1 instantiates a new ConversionGoalFilter1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionGoalFilter1WithDefaults

`func NewConversionGoalFilter1WithDefaults() *ConversionGoalFilter1`

NewConversionGoalFilter1WithDefaults instantiates a new ConversionGoalFilter1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversionGoalId

`func (o *ConversionGoalFilter1) GetConversionGoalId() string`

GetConversionGoalId returns the ConversionGoalId field if non-nil, zero value otherwise.

### GetConversionGoalIdOk

`func (o *ConversionGoalFilter1) GetConversionGoalIdOk() (*string, bool)`

GetConversionGoalIdOk returns a tuple with the ConversionGoalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoalId

`func (o *ConversionGoalFilter1) SetConversionGoalId(v string)`

SetConversionGoalId sets ConversionGoalId field to given value.


### GetConversionGoalName

`func (o *ConversionGoalFilter1) GetConversionGoalName() string`

GetConversionGoalName returns the ConversionGoalName field if non-nil, zero value otherwise.

### GetConversionGoalNameOk

`func (o *ConversionGoalFilter1) GetConversionGoalNameOk() (*string, bool)`

GetConversionGoalNameOk returns a tuple with the ConversionGoalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoalName

`func (o *ConversionGoalFilter1) SetConversionGoalName(v string)`

SetConversionGoalName sets ConversionGoalName field to given value.


### GetCustomName

`func (o *ConversionGoalFilter1) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *ConversionGoalFilter1) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *ConversionGoalFilter1) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *ConversionGoalFilter1) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *ConversionGoalFilter1) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *ConversionGoalFilter1) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetEvent

`func (o *ConversionGoalFilter1) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ConversionGoalFilter1) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ConversionGoalFilter1) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *ConversionGoalFilter1) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *ConversionGoalFilter1) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *ConversionGoalFilter1) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetFixedProperties

`func (o *ConversionGoalFilter1) GetFixedProperties() []FixedpropertiesInner`

GetFixedProperties returns the FixedProperties field if non-nil, zero value otherwise.

### GetFixedPropertiesOk

`func (o *ConversionGoalFilter1) GetFixedPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetFixedPropertiesOk returns a tuple with the FixedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedProperties

`func (o *ConversionGoalFilter1) SetFixedProperties(v []FixedpropertiesInner)`

SetFixedProperties sets FixedProperties field to given value.

### HasFixedProperties

`func (o *ConversionGoalFilter1) HasFixedProperties() bool`

HasFixedProperties returns a boolean if a field has been set.

### SetFixedPropertiesNil

`func (o *ConversionGoalFilter1) SetFixedPropertiesNil(b bool)`

 SetFixedPropertiesNil sets the value for FixedProperties to be an explicit nil

### UnsetFixedProperties
`func (o *ConversionGoalFilter1) UnsetFixedProperties()`

UnsetFixedProperties ensures that no value is present for FixedProperties, not even an explicit nil
### GetKind

`func (o *ConversionGoalFilter1) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConversionGoalFilter1) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConversionGoalFilter1) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConversionGoalFilter1) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *ConversionGoalFilter1) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ConversionGoalFilter1) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ConversionGoalFilter1) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ConversionGoalFilter1) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ConversionGoalFilter1) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ConversionGoalFilter1) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetMath

`func (o *ConversionGoalFilter1) GetMath() string`

GetMath returns the Math field if non-nil, zero value otherwise.

### GetMathOk

`func (o *ConversionGoalFilter1) GetMathOk() (*string, bool)`

GetMathOk returns a tuple with the Math field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMath

`func (o *ConversionGoalFilter1) SetMath(v string)`

SetMath sets Math field to given value.

### HasMath

`func (o *ConversionGoalFilter1) HasMath() bool`

HasMath returns a boolean if a field has been set.

### SetMathNil

`func (o *ConversionGoalFilter1) SetMathNil(b bool)`

 SetMathNil sets the value for Math to be an explicit nil

### UnsetMath
`func (o *ConversionGoalFilter1) UnsetMath()`

UnsetMath ensures that no value is present for Math, not even an explicit nil
### GetMathGroupTypeIndex

`func (o *ConversionGoalFilter1) GetMathGroupTypeIndex() MathGroupTypeIndex`

GetMathGroupTypeIndex returns the MathGroupTypeIndex field if non-nil, zero value otherwise.

### GetMathGroupTypeIndexOk

`func (o *ConversionGoalFilter1) GetMathGroupTypeIndexOk() (*MathGroupTypeIndex, bool)`

GetMathGroupTypeIndexOk returns a tuple with the MathGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathGroupTypeIndex

`func (o *ConversionGoalFilter1) SetMathGroupTypeIndex(v MathGroupTypeIndex)`

SetMathGroupTypeIndex sets MathGroupTypeIndex field to given value.

### HasMathGroupTypeIndex

`func (o *ConversionGoalFilter1) HasMathGroupTypeIndex() bool`

HasMathGroupTypeIndex returns a boolean if a field has been set.

### GetMathHogql

`func (o *ConversionGoalFilter1) GetMathHogql() string`

GetMathHogql returns the MathHogql field if non-nil, zero value otherwise.

### GetMathHogqlOk

`func (o *ConversionGoalFilter1) GetMathHogqlOk() (*string, bool)`

GetMathHogqlOk returns a tuple with the MathHogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathHogql

`func (o *ConversionGoalFilter1) SetMathHogql(v string)`

SetMathHogql sets MathHogql field to given value.

### HasMathHogql

`func (o *ConversionGoalFilter1) HasMathHogql() bool`

HasMathHogql returns a boolean if a field has been set.

### SetMathHogqlNil

`func (o *ConversionGoalFilter1) SetMathHogqlNil(b bool)`

 SetMathHogqlNil sets the value for MathHogql to be an explicit nil

### UnsetMathHogql
`func (o *ConversionGoalFilter1) UnsetMathHogql()`

UnsetMathHogql ensures that no value is present for MathHogql, not even an explicit nil
### GetMathMultiplier

`func (o *ConversionGoalFilter1) GetMathMultiplier() float32`

GetMathMultiplier returns the MathMultiplier field if non-nil, zero value otherwise.

### GetMathMultiplierOk

`func (o *ConversionGoalFilter1) GetMathMultiplierOk() (*float32, bool)`

GetMathMultiplierOk returns a tuple with the MathMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathMultiplier

`func (o *ConversionGoalFilter1) SetMathMultiplier(v float32)`

SetMathMultiplier sets MathMultiplier field to given value.

### HasMathMultiplier

`func (o *ConversionGoalFilter1) HasMathMultiplier() bool`

HasMathMultiplier returns a boolean if a field has been set.

### SetMathMultiplierNil

`func (o *ConversionGoalFilter1) SetMathMultiplierNil(b bool)`

 SetMathMultiplierNil sets the value for MathMultiplier to be an explicit nil

### UnsetMathMultiplier
`func (o *ConversionGoalFilter1) UnsetMathMultiplier()`

UnsetMathMultiplier ensures that no value is present for MathMultiplier, not even an explicit nil
### GetMathProperty

`func (o *ConversionGoalFilter1) GetMathProperty() string`

GetMathProperty returns the MathProperty field if non-nil, zero value otherwise.

### GetMathPropertyOk

`func (o *ConversionGoalFilter1) GetMathPropertyOk() (*string, bool)`

GetMathPropertyOk returns a tuple with the MathProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathProperty

`func (o *ConversionGoalFilter1) SetMathProperty(v string)`

SetMathProperty sets MathProperty field to given value.

### HasMathProperty

`func (o *ConversionGoalFilter1) HasMathProperty() bool`

HasMathProperty returns a boolean if a field has been set.

### SetMathPropertyNil

`func (o *ConversionGoalFilter1) SetMathPropertyNil(b bool)`

 SetMathPropertyNil sets the value for MathProperty to be an explicit nil

### UnsetMathProperty
`func (o *ConversionGoalFilter1) UnsetMathProperty()`

UnsetMathProperty ensures that no value is present for MathProperty, not even an explicit nil
### GetMathPropertyRevenueCurrency

`func (o *ConversionGoalFilter1) GetMathPropertyRevenueCurrency() RevenueCurrencyPropertyConfig`

GetMathPropertyRevenueCurrency returns the MathPropertyRevenueCurrency field if non-nil, zero value otherwise.

### GetMathPropertyRevenueCurrencyOk

`func (o *ConversionGoalFilter1) GetMathPropertyRevenueCurrencyOk() (*RevenueCurrencyPropertyConfig, bool)`

GetMathPropertyRevenueCurrencyOk returns a tuple with the MathPropertyRevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyRevenueCurrency

`func (o *ConversionGoalFilter1) SetMathPropertyRevenueCurrency(v RevenueCurrencyPropertyConfig)`

SetMathPropertyRevenueCurrency sets MathPropertyRevenueCurrency field to given value.

### HasMathPropertyRevenueCurrency

`func (o *ConversionGoalFilter1) HasMathPropertyRevenueCurrency() bool`

HasMathPropertyRevenueCurrency returns a boolean if a field has been set.

### GetMathPropertyType

`func (o *ConversionGoalFilter1) GetMathPropertyType() string`

GetMathPropertyType returns the MathPropertyType field if non-nil, zero value otherwise.

### GetMathPropertyTypeOk

`func (o *ConversionGoalFilter1) GetMathPropertyTypeOk() (*string, bool)`

GetMathPropertyTypeOk returns a tuple with the MathPropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyType

`func (o *ConversionGoalFilter1) SetMathPropertyType(v string)`

SetMathPropertyType sets MathPropertyType field to given value.

### HasMathPropertyType

`func (o *ConversionGoalFilter1) HasMathPropertyType() bool`

HasMathPropertyType returns a boolean if a field has been set.

### SetMathPropertyTypeNil

`func (o *ConversionGoalFilter1) SetMathPropertyTypeNil(b bool)`

 SetMathPropertyTypeNil sets the value for MathPropertyType to be an explicit nil

### UnsetMathPropertyType
`func (o *ConversionGoalFilter1) UnsetMathPropertyType()`

UnsetMathPropertyType ensures that no value is present for MathPropertyType, not even an explicit nil
### GetName

`func (o *ConversionGoalFilter1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConversionGoalFilter1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConversionGoalFilter1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConversionGoalFilter1) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConversionGoalFilter1) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConversionGoalFilter1) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOptionalInFunnel

`func (o *ConversionGoalFilter1) GetOptionalInFunnel() bool`

GetOptionalInFunnel returns the OptionalInFunnel field if non-nil, zero value otherwise.

### GetOptionalInFunnelOk

`func (o *ConversionGoalFilter1) GetOptionalInFunnelOk() (*bool, bool)`

GetOptionalInFunnelOk returns a tuple with the OptionalInFunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalInFunnel

`func (o *ConversionGoalFilter1) SetOptionalInFunnel(v bool)`

SetOptionalInFunnel sets OptionalInFunnel field to given value.

### HasOptionalInFunnel

`func (o *ConversionGoalFilter1) HasOptionalInFunnel() bool`

HasOptionalInFunnel returns a boolean if a field has been set.

### SetOptionalInFunnelNil

`func (o *ConversionGoalFilter1) SetOptionalInFunnelNil(b bool)`

 SetOptionalInFunnelNil sets the value for OptionalInFunnel to be an explicit nil

### UnsetOptionalInFunnel
`func (o *ConversionGoalFilter1) UnsetOptionalInFunnel()`

UnsetOptionalInFunnel ensures that no value is present for OptionalInFunnel, not even an explicit nil
### GetOrderBy

`func (o *ConversionGoalFilter1) GetOrderBy() []string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ConversionGoalFilter1) GetOrderByOk() (*[]string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ConversionGoalFilter1) SetOrderBy(v []string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ConversionGoalFilter1) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *ConversionGoalFilter1) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *ConversionGoalFilter1) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetProperties

`func (o *ConversionGoalFilter1) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ConversionGoalFilter1) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ConversionGoalFilter1) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ConversionGoalFilter1) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ConversionGoalFilter1) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ConversionGoalFilter1) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *ConversionGoalFilter1) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ConversionGoalFilter1) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ConversionGoalFilter1) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ConversionGoalFilter1) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *ConversionGoalFilter1) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *ConversionGoalFilter1) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetSchemaMap

`func (o *ConversionGoalFilter1) GetSchemaMap() map[string]SchemaMapValue`

GetSchemaMap returns the SchemaMap field if non-nil, zero value otherwise.

### GetSchemaMapOk

`func (o *ConversionGoalFilter1) GetSchemaMapOk() (*map[string]SchemaMapValue, bool)`

GetSchemaMapOk returns a tuple with the SchemaMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaMap

`func (o *ConversionGoalFilter1) SetSchemaMap(v map[string]SchemaMapValue)`

SetSchemaMap sets SchemaMap field to given value.


### GetVersion

`func (o *ConversionGoalFilter1) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConversionGoalFilter1) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConversionGoalFilter1) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConversionGoalFilter1) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ConversionGoalFilter1) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ConversionGoalFilter1) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


