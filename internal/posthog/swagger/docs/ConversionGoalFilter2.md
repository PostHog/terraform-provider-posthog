# ConversionGoalFilter2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversionGoalId** | **string** |  | 
**ConversionGoalName** | **string** |  | 
**CustomName** | Pointer to **NullableString** |  | [optional] 
**FixedProperties** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) | Fixed properties in the query, can&#39;t be edited in the interface (e.g. scoping down by person) | [optional] 
**Id** | **int32** |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "ActionsNode"]
**Math** | Pointer to **NullableString** |  | [optional] [default to "null"]
**MathGroupTypeIndex** | Pointer to [**MathGroupTypeIndex**](MathGroupTypeIndex.md) |  | [optional] 
**MathHogql** | Pointer to **NullableString** |  | [optional] 
**MathMultiplier** | Pointer to **NullableFloat32** |  | [optional] 
**MathProperty** | Pointer to **NullableString** |  | [optional] 
**MathPropertyRevenueCurrency** | Pointer to [**RevenueCurrencyPropertyConfig**](RevenueCurrencyPropertyConfig.md) |  | [optional] 
**MathPropertyType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OptionalInFunnel** | Pointer to **NullableBool** |  | [optional] 
**Properties** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) | Properties configurable in the interface | [optional] 
**Response** | Pointer to **map[string]interface{}** |  | [optional] 
**SchemaMap** | [**map[string]SchemaMapValue**](SchemaMapValue.md) |  | 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewConversionGoalFilter2

`func NewConversionGoalFilter2(conversionGoalId string, conversionGoalName string, id int32, schemaMap map[string]SchemaMapValue, ) *ConversionGoalFilter2`

NewConversionGoalFilter2 instantiates a new ConversionGoalFilter2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionGoalFilter2WithDefaults

`func NewConversionGoalFilter2WithDefaults() *ConversionGoalFilter2`

NewConversionGoalFilter2WithDefaults instantiates a new ConversionGoalFilter2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversionGoalId

`func (o *ConversionGoalFilter2) GetConversionGoalId() string`

GetConversionGoalId returns the ConversionGoalId field if non-nil, zero value otherwise.

### GetConversionGoalIdOk

`func (o *ConversionGoalFilter2) GetConversionGoalIdOk() (*string, bool)`

GetConversionGoalIdOk returns a tuple with the ConversionGoalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoalId

`func (o *ConversionGoalFilter2) SetConversionGoalId(v string)`

SetConversionGoalId sets ConversionGoalId field to given value.


### GetConversionGoalName

`func (o *ConversionGoalFilter2) GetConversionGoalName() string`

GetConversionGoalName returns the ConversionGoalName field if non-nil, zero value otherwise.

### GetConversionGoalNameOk

`func (o *ConversionGoalFilter2) GetConversionGoalNameOk() (*string, bool)`

GetConversionGoalNameOk returns a tuple with the ConversionGoalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoalName

`func (o *ConversionGoalFilter2) SetConversionGoalName(v string)`

SetConversionGoalName sets ConversionGoalName field to given value.


### GetCustomName

`func (o *ConversionGoalFilter2) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *ConversionGoalFilter2) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *ConversionGoalFilter2) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *ConversionGoalFilter2) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *ConversionGoalFilter2) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *ConversionGoalFilter2) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetFixedProperties

`func (o *ConversionGoalFilter2) GetFixedProperties() []FixedpropertiesInner`

GetFixedProperties returns the FixedProperties field if non-nil, zero value otherwise.

### GetFixedPropertiesOk

`func (o *ConversionGoalFilter2) GetFixedPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetFixedPropertiesOk returns a tuple with the FixedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedProperties

`func (o *ConversionGoalFilter2) SetFixedProperties(v []FixedpropertiesInner)`

SetFixedProperties sets FixedProperties field to given value.

### HasFixedProperties

`func (o *ConversionGoalFilter2) HasFixedProperties() bool`

HasFixedProperties returns a boolean if a field has been set.

### SetFixedPropertiesNil

`func (o *ConversionGoalFilter2) SetFixedPropertiesNil(b bool)`

 SetFixedPropertiesNil sets the value for FixedProperties to be an explicit nil

### UnsetFixedProperties
`func (o *ConversionGoalFilter2) UnsetFixedProperties()`

UnsetFixedProperties ensures that no value is present for FixedProperties, not even an explicit nil
### GetId

`func (o *ConversionGoalFilter2) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConversionGoalFilter2) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConversionGoalFilter2) SetId(v int32)`

SetId sets Id field to given value.


### GetKind

`func (o *ConversionGoalFilter2) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConversionGoalFilter2) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConversionGoalFilter2) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConversionGoalFilter2) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMath

`func (o *ConversionGoalFilter2) GetMath() string`

GetMath returns the Math field if non-nil, zero value otherwise.

### GetMathOk

`func (o *ConversionGoalFilter2) GetMathOk() (*string, bool)`

GetMathOk returns a tuple with the Math field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMath

`func (o *ConversionGoalFilter2) SetMath(v string)`

SetMath sets Math field to given value.

### HasMath

`func (o *ConversionGoalFilter2) HasMath() bool`

HasMath returns a boolean if a field has been set.

### SetMathNil

`func (o *ConversionGoalFilter2) SetMathNil(b bool)`

 SetMathNil sets the value for Math to be an explicit nil

### UnsetMath
`func (o *ConversionGoalFilter2) UnsetMath()`

UnsetMath ensures that no value is present for Math, not even an explicit nil
### GetMathGroupTypeIndex

`func (o *ConversionGoalFilter2) GetMathGroupTypeIndex() MathGroupTypeIndex`

GetMathGroupTypeIndex returns the MathGroupTypeIndex field if non-nil, zero value otherwise.

### GetMathGroupTypeIndexOk

`func (o *ConversionGoalFilter2) GetMathGroupTypeIndexOk() (*MathGroupTypeIndex, bool)`

GetMathGroupTypeIndexOk returns a tuple with the MathGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathGroupTypeIndex

`func (o *ConversionGoalFilter2) SetMathGroupTypeIndex(v MathGroupTypeIndex)`

SetMathGroupTypeIndex sets MathGroupTypeIndex field to given value.

### HasMathGroupTypeIndex

`func (o *ConversionGoalFilter2) HasMathGroupTypeIndex() bool`

HasMathGroupTypeIndex returns a boolean if a field has been set.

### GetMathHogql

`func (o *ConversionGoalFilter2) GetMathHogql() string`

GetMathHogql returns the MathHogql field if non-nil, zero value otherwise.

### GetMathHogqlOk

`func (o *ConversionGoalFilter2) GetMathHogqlOk() (*string, bool)`

GetMathHogqlOk returns a tuple with the MathHogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathHogql

`func (o *ConversionGoalFilter2) SetMathHogql(v string)`

SetMathHogql sets MathHogql field to given value.

### HasMathHogql

`func (o *ConversionGoalFilter2) HasMathHogql() bool`

HasMathHogql returns a boolean if a field has been set.

### SetMathHogqlNil

`func (o *ConversionGoalFilter2) SetMathHogqlNil(b bool)`

 SetMathHogqlNil sets the value for MathHogql to be an explicit nil

### UnsetMathHogql
`func (o *ConversionGoalFilter2) UnsetMathHogql()`

UnsetMathHogql ensures that no value is present for MathHogql, not even an explicit nil
### GetMathMultiplier

`func (o *ConversionGoalFilter2) GetMathMultiplier() float32`

GetMathMultiplier returns the MathMultiplier field if non-nil, zero value otherwise.

### GetMathMultiplierOk

`func (o *ConversionGoalFilter2) GetMathMultiplierOk() (*float32, bool)`

GetMathMultiplierOk returns a tuple with the MathMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathMultiplier

`func (o *ConversionGoalFilter2) SetMathMultiplier(v float32)`

SetMathMultiplier sets MathMultiplier field to given value.

### HasMathMultiplier

`func (o *ConversionGoalFilter2) HasMathMultiplier() bool`

HasMathMultiplier returns a boolean if a field has been set.

### SetMathMultiplierNil

`func (o *ConversionGoalFilter2) SetMathMultiplierNil(b bool)`

 SetMathMultiplierNil sets the value for MathMultiplier to be an explicit nil

### UnsetMathMultiplier
`func (o *ConversionGoalFilter2) UnsetMathMultiplier()`

UnsetMathMultiplier ensures that no value is present for MathMultiplier, not even an explicit nil
### GetMathProperty

`func (o *ConversionGoalFilter2) GetMathProperty() string`

GetMathProperty returns the MathProperty field if non-nil, zero value otherwise.

### GetMathPropertyOk

`func (o *ConversionGoalFilter2) GetMathPropertyOk() (*string, bool)`

GetMathPropertyOk returns a tuple with the MathProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathProperty

`func (o *ConversionGoalFilter2) SetMathProperty(v string)`

SetMathProperty sets MathProperty field to given value.

### HasMathProperty

`func (o *ConversionGoalFilter2) HasMathProperty() bool`

HasMathProperty returns a boolean if a field has been set.

### SetMathPropertyNil

`func (o *ConversionGoalFilter2) SetMathPropertyNil(b bool)`

 SetMathPropertyNil sets the value for MathProperty to be an explicit nil

### UnsetMathProperty
`func (o *ConversionGoalFilter2) UnsetMathProperty()`

UnsetMathProperty ensures that no value is present for MathProperty, not even an explicit nil
### GetMathPropertyRevenueCurrency

`func (o *ConversionGoalFilter2) GetMathPropertyRevenueCurrency() RevenueCurrencyPropertyConfig`

GetMathPropertyRevenueCurrency returns the MathPropertyRevenueCurrency field if non-nil, zero value otherwise.

### GetMathPropertyRevenueCurrencyOk

`func (o *ConversionGoalFilter2) GetMathPropertyRevenueCurrencyOk() (*RevenueCurrencyPropertyConfig, bool)`

GetMathPropertyRevenueCurrencyOk returns a tuple with the MathPropertyRevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyRevenueCurrency

`func (o *ConversionGoalFilter2) SetMathPropertyRevenueCurrency(v RevenueCurrencyPropertyConfig)`

SetMathPropertyRevenueCurrency sets MathPropertyRevenueCurrency field to given value.

### HasMathPropertyRevenueCurrency

`func (o *ConversionGoalFilter2) HasMathPropertyRevenueCurrency() bool`

HasMathPropertyRevenueCurrency returns a boolean if a field has been set.

### GetMathPropertyType

`func (o *ConversionGoalFilter2) GetMathPropertyType() string`

GetMathPropertyType returns the MathPropertyType field if non-nil, zero value otherwise.

### GetMathPropertyTypeOk

`func (o *ConversionGoalFilter2) GetMathPropertyTypeOk() (*string, bool)`

GetMathPropertyTypeOk returns a tuple with the MathPropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyType

`func (o *ConversionGoalFilter2) SetMathPropertyType(v string)`

SetMathPropertyType sets MathPropertyType field to given value.

### HasMathPropertyType

`func (o *ConversionGoalFilter2) HasMathPropertyType() bool`

HasMathPropertyType returns a boolean if a field has been set.

### SetMathPropertyTypeNil

`func (o *ConversionGoalFilter2) SetMathPropertyTypeNil(b bool)`

 SetMathPropertyTypeNil sets the value for MathPropertyType to be an explicit nil

### UnsetMathPropertyType
`func (o *ConversionGoalFilter2) UnsetMathPropertyType()`

UnsetMathPropertyType ensures that no value is present for MathPropertyType, not even an explicit nil
### GetName

`func (o *ConversionGoalFilter2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConversionGoalFilter2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConversionGoalFilter2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConversionGoalFilter2) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConversionGoalFilter2) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConversionGoalFilter2) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOptionalInFunnel

`func (o *ConversionGoalFilter2) GetOptionalInFunnel() bool`

GetOptionalInFunnel returns the OptionalInFunnel field if non-nil, zero value otherwise.

### GetOptionalInFunnelOk

`func (o *ConversionGoalFilter2) GetOptionalInFunnelOk() (*bool, bool)`

GetOptionalInFunnelOk returns a tuple with the OptionalInFunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalInFunnel

`func (o *ConversionGoalFilter2) SetOptionalInFunnel(v bool)`

SetOptionalInFunnel sets OptionalInFunnel field to given value.

### HasOptionalInFunnel

`func (o *ConversionGoalFilter2) HasOptionalInFunnel() bool`

HasOptionalInFunnel returns a boolean if a field has been set.

### SetOptionalInFunnelNil

`func (o *ConversionGoalFilter2) SetOptionalInFunnelNil(b bool)`

 SetOptionalInFunnelNil sets the value for OptionalInFunnel to be an explicit nil

### UnsetOptionalInFunnel
`func (o *ConversionGoalFilter2) UnsetOptionalInFunnel()`

UnsetOptionalInFunnel ensures that no value is present for OptionalInFunnel, not even an explicit nil
### GetProperties

`func (o *ConversionGoalFilter2) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ConversionGoalFilter2) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ConversionGoalFilter2) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ConversionGoalFilter2) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ConversionGoalFilter2) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ConversionGoalFilter2) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *ConversionGoalFilter2) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ConversionGoalFilter2) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ConversionGoalFilter2) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ConversionGoalFilter2) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *ConversionGoalFilter2) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *ConversionGoalFilter2) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetSchemaMap

`func (o *ConversionGoalFilter2) GetSchemaMap() map[string]SchemaMapValue`

GetSchemaMap returns the SchemaMap field if non-nil, zero value otherwise.

### GetSchemaMapOk

`func (o *ConversionGoalFilter2) GetSchemaMapOk() (*map[string]SchemaMapValue, bool)`

GetSchemaMapOk returns a tuple with the SchemaMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaMap

`func (o *ConversionGoalFilter2) SetSchemaMap(v map[string]SchemaMapValue)`

SetSchemaMap sets SchemaMap field to given value.


### GetVersion

`func (o *ConversionGoalFilter2) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConversionGoalFilter2) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConversionGoalFilter2) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConversionGoalFilter2) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ConversionGoalFilter2) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ConversionGoalFilter2) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


