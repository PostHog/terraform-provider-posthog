# ExperimentDataWarehouseNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomName** | Pointer to **NullableString** |  | [optional] 
**DataWarehouseJoinKey** | **string** |  | 
**EventsJoinKey** | **string** |  | 
**FixedProperties** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) | Fixed properties in the query, can&#39;t be edited in the interface (e.g. scoping down by person) | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentDataWarehouseNode"]
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
**TableName** | **string** |  | 
**TimestampField** | **string** |  | 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewExperimentDataWarehouseNode

`func NewExperimentDataWarehouseNode(dataWarehouseJoinKey string, eventsJoinKey string, tableName string, timestampField string, ) *ExperimentDataWarehouseNode`

NewExperimentDataWarehouseNode instantiates a new ExperimentDataWarehouseNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentDataWarehouseNodeWithDefaults

`func NewExperimentDataWarehouseNodeWithDefaults() *ExperimentDataWarehouseNode`

NewExperimentDataWarehouseNodeWithDefaults instantiates a new ExperimentDataWarehouseNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomName

`func (o *ExperimentDataWarehouseNode) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *ExperimentDataWarehouseNode) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *ExperimentDataWarehouseNode) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *ExperimentDataWarehouseNode) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *ExperimentDataWarehouseNode) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *ExperimentDataWarehouseNode) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetDataWarehouseJoinKey

`func (o *ExperimentDataWarehouseNode) GetDataWarehouseJoinKey() string`

GetDataWarehouseJoinKey returns the DataWarehouseJoinKey field if non-nil, zero value otherwise.

### GetDataWarehouseJoinKeyOk

`func (o *ExperimentDataWarehouseNode) GetDataWarehouseJoinKeyOk() (*string, bool)`

GetDataWarehouseJoinKeyOk returns a tuple with the DataWarehouseJoinKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWarehouseJoinKey

`func (o *ExperimentDataWarehouseNode) SetDataWarehouseJoinKey(v string)`

SetDataWarehouseJoinKey sets DataWarehouseJoinKey field to given value.


### GetEventsJoinKey

`func (o *ExperimentDataWarehouseNode) GetEventsJoinKey() string`

GetEventsJoinKey returns the EventsJoinKey field if non-nil, zero value otherwise.

### GetEventsJoinKeyOk

`func (o *ExperimentDataWarehouseNode) GetEventsJoinKeyOk() (*string, bool)`

GetEventsJoinKeyOk returns a tuple with the EventsJoinKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsJoinKey

`func (o *ExperimentDataWarehouseNode) SetEventsJoinKey(v string)`

SetEventsJoinKey sets EventsJoinKey field to given value.


### GetFixedProperties

`func (o *ExperimentDataWarehouseNode) GetFixedProperties() []FixedpropertiesInner`

GetFixedProperties returns the FixedProperties field if non-nil, zero value otherwise.

### GetFixedPropertiesOk

`func (o *ExperimentDataWarehouseNode) GetFixedPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetFixedPropertiesOk returns a tuple with the FixedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedProperties

`func (o *ExperimentDataWarehouseNode) SetFixedProperties(v []FixedpropertiesInner)`

SetFixedProperties sets FixedProperties field to given value.

### HasFixedProperties

`func (o *ExperimentDataWarehouseNode) HasFixedProperties() bool`

HasFixedProperties returns a boolean if a field has been set.

### SetFixedPropertiesNil

`func (o *ExperimentDataWarehouseNode) SetFixedPropertiesNil(b bool)`

 SetFixedPropertiesNil sets the value for FixedProperties to be an explicit nil

### UnsetFixedProperties
`func (o *ExperimentDataWarehouseNode) UnsetFixedProperties()`

UnsetFixedProperties ensures that no value is present for FixedProperties, not even an explicit nil
### GetKind

`func (o *ExperimentDataWarehouseNode) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExperimentDataWarehouseNode) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExperimentDataWarehouseNode) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExperimentDataWarehouseNode) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMath

`func (o *ExperimentDataWarehouseNode) GetMath() string`

GetMath returns the Math field if non-nil, zero value otherwise.

### GetMathOk

`func (o *ExperimentDataWarehouseNode) GetMathOk() (*string, bool)`

GetMathOk returns a tuple with the Math field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMath

`func (o *ExperimentDataWarehouseNode) SetMath(v string)`

SetMath sets Math field to given value.

### HasMath

`func (o *ExperimentDataWarehouseNode) HasMath() bool`

HasMath returns a boolean if a field has been set.

### SetMathNil

`func (o *ExperimentDataWarehouseNode) SetMathNil(b bool)`

 SetMathNil sets the value for Math to be an explicit nil

### UnsetMath
`func (o *ExperimentDataWarehouseNode) UnsetMath()`

UnsetMath ensures that no value is present for Math, not even an explicit nil
### GetMathGroupTypeIndex

`func (o *ExperimentDataWarehouseNode) GetMathGroupTypeIndex() MathGroupTypeIndex`

GetMathGroupTypeIndex returns the MathGroupTypeIndex field if non-nil, zero value otherwise.

### GetMathGroupTypeIndexOk

`func (o *ExperimentDataWarehouseNode) GetMathGroupTypeIndexOk() (*MathGroupTypeIndex, bool)`

GetMathGroupTypeIndexOk returns a tuple with the MathGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathGroupTypeIndex

`func (o *ExperimentDataWarehouseNode) SetMathGroupTypeIndex(v MathGroupTypeIndex)`

SetMathGroupTypeIndex sets MathGroupTypeIndex field to given value.

### HasMathGroupTypeIndex

`func (o *ExperimentDataWarehouseNode) HasMathGroupTypeIndex() bool`

HasMathGroupTypeIndex returns a boolean if a field has been set.

### GetMathHogql

`func (o *ExperimentDataWarehouseNode) GetMathHogql() string`

GetMathHogql returns the MathHogql field if non-nil, zero value otherwise.

### GetMathHogqlOk

`func (o *ExperimentDataWarehouseNode) GetMathHogqlOk() (*string, bool)`

GetMathHogqlOk returns a tuple with the MathHogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathHogql

`func (o *ExperimentDataWarehouseNode) SetMathHogql(v string)`

SetMathHogql sets MathHogql field to given value.

### HasMathHogql

`func (o *ExperimentDataWarehouseNode) HasMathHogql() bool`

HasMathHogql returns a boolean if a field has been set.

### SetMathHogqlNil

`func (o *ExperimentDataWarehouseNode) SetMathHogqlNil(b bool)`

 SetMathHogqlNil sets the value for MathHogql to be an explicit nil

### UnsetMathHogql
`func (o *ExperimentDataWarehouseNode) UnsetMathHogql()`

UnsetMathHogql ensures that no value is present for MathHogql, not even an explicit nil
### GetMathMultiplier

`func (o *ExperimentDataWarehouseNode) GetMathMultiplier() float32`

GetMathMultiplier returns the MathMultiplier field if non-nil, zero value otherwise.

### GetMathMultiplierOk

`func (o *ExperimentDataWarehouseNode) GetMathMultiplierOk() (*float32, bool)`

GetMathMultiplierOk returns a tuple with the MathMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathMultiplier

`func (o *ExperimentDataWarehouseNode) SetMathMultiplier(v float32)`

SetMathMultiplier sets MathMultiplier field to given value.

### HasMathMultiplier

`func (o *ExperimentDataWarehouseNode) HasMathMultiplier() bool`

HasMathMultiplier returns a boolean if a field has been set.

### SetMathMultiplierNil

`func (o *ExperimentDataWarehouseNode) SetMathMultiplierNil(b bool)`

 SetMathMultiplierNil sets the value for MathMultiplier to be an explicit nil

### UnsetMathMultiplier
`func (o *ExperimentDataWarehouseNode) UnsetMathMultiplier()`

UnsetMathMultiplier ensures that no value is present for MathMultiplier, not even an explicit nil
### GetMathProperty

`func (o *ExperimentDataWarehouseNode) GetMathProperty() string`

GetMathProperty returns the MathProperty field if non-nil, zero value otherwise.

### GetMathPropertyOk

`func (o *ExperimentDataWarehouseNode) GetMathPropertyOk() (*string, bool)`

GetMathPropertyOk returns a tuple with the MathProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathProperty

`func (o *ExperimentDataWarehouseNode) SetMathProperty(v string)`

SetMathProperty sets MathProperty field to given value.

### HasMathProperty

`func (o *ExperimentDataWarehouseNode) HasMathProperty() bool`

HasMathProperty returns a boolean if a field has been set.

### SetMathPropertyNil

`func (o *ExperimentDataWarehouseNode) SetMathPropertyNil(b bool)`

 SetMathPropertyNil sets the value for MathProperty to be an explicit nil

### UnsetMathProperty
`func (o *ExperimentDataWarehouseNode) UnsetMathProperty()`

UnsetMathProperty ensures that no value is present for MathProperty, not even an explicit nil
### GetMathPropertyRevenueCurrency

`func (o *ExperimentDataWarehouseNode) GetMathPropertyRevenueCurrency() RevenueCurrencyPropertyConfig`

GetMathPropertyRevenueCurrency returns the MathPropertyRevenueCurrency field if non-nil, zero value otherwise.

### GetMathPropertyRevenueCurrencyOk

`func (o *ExperimentDataWarehouseNode) GetMathPropertyRevenueCurrencyOk() (*RevenueCurrencyPropertyConfig, bool)`

GetMathPropertyRevenueCurrencyOk returns a tuple with the MathPropertyRevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyRevenueCurrency

`func (o *ExperimentDataWarehouseNode) SetMathPropertyRevenueCurrency(v RevenueCurrencyPropertyConfig)`

SetMathPropertyRevenueCurrency sets MathPropertyRevenueCurrency field to given value.

### HasMathPropertyRevenueCurrency

`func (o *ExperimentDataWarehouseNode) HasMathPropertyRevenueCurrency() bool`

HasMathPropertyRevenueCurrency returns a boolean if a field has been set.

### GetMathPropertyType

`func (o *ExperimentDataWarehouseNode) GetMathPropertyType() string`

GetMathPropertyType returns the MathPropertyType field if non-nil, zero value otherwise.

### GetMathPropertyTypeOk

`func (o *ExperimentDataWarehouseNode) GetMathPropertyTypeOk() (*string, bool)`

GetMathPropertyTypeOk returns a tuple with the MathPropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyType

`func (o *ExperimentDataWarehouseNode) SetMathPropertyType(v string)`

SetMathPropertyType sets MathPropertyType field to given value.

### HasMathPropertyType

`func (o *ExperimentDataWarehouseNode) HasMathPropertyType() bool`

HasMathPropertyType returns a boolean if a field has been set.

### SetMathPropertyTypeNil

`func (o *ExperimentDataWarehouseNode) SetMathPropertyTypeNil(b bool)`

 SetMathPropertyTypeNil sets the value for MathPropertyType to be an explicit nil

### UnsetMathPropertyType
`func (o *ExperimentDataWarehouseNode) UnsetMathPropertyType()`

UnsetMathPropertyType ensures that no value is present for MathPropertyType, not even an explicit nil
### GetName

`func (o *ExperimentDataWarehouseNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentDataWarehouseNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentDataWarehouseNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExperimentDataWarehouseNode) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExperimentDataWarehouseNode) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExperimentDataWarehouseNode) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOptionalInFunnel

`func (o *ExperimentDataWarehouseNode) GetOptionalInFunnel() bool`

GetOptionalInFunnel returns the OptionalInFunnel field if non-nil, zero value otherwise.

### GetOptionalInFunnelOk

`func (o *ExperimentDataWarehouseNode) GetOptionalInFunnelOk() (*bool, bool)`

GetOptionalInFunnelOk returns a tuple with the OptionalInFunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalInFunnel

`func (o *ExperimentDataWarehouseNode) SetOptionalInFunnel(v bool)`

SetOptionalInFunnel sets OptionalInFunnel field to given value.

### HasOptionalInFunnel

`func (o *ExperimentDataWarehouseNode) HasOptionalInFunnel() bool`

HasOptionalInFunnel returns a boolean if a field has been set.

### SetOptionalInFunnelNil

`func (o *ExperimentDataWarehouseNode) SetOptionalInFunnelNil(b bool)`

 SetOptionalInFunnelNil sets the value for OptionalInFunnel to be an explicit nil

### UnsetOptionalInFunnel
`func (o *ExperimentDataWarehouseNode) UnsetOptionalInFunnel()`

UnsetOptionalInFunnel ensures that no value is present for OptionalInFunnel, not even an explicit nil
### GetProperties

`func (o *ExperimentDataWarehouseNode) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ExperimentDataWarehouseNode) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ExperimentDataWarehouseNode) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ExperimentDataWarehouseNode) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ExperimentDataWarehouseNode) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ExperimentDataWarehouseNode) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *ExperimentDataWarehouseNode) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ExperimentDataWarehouseNode) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ExperimentDataWarehouseNode) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ExperimentDataWarehouseNode) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *ExperimentDataWarehouseNode) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *ExperimentDataWarehouseNode) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetTableName

`func (o *ExperimentDataWarehouseNode) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *ExperimentDataWarehouseNode) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *ExperimentDataWarehouseNode) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetTimestampField

`func (o *ExperimentDataWarehouseNode) GetTimestampField() string`

GetTimestampField returns the TimestampField field if non-nil, zero value otherwise.

### GetTimestampFieldOk

`func (o *ExperimentDataWarehouseNode) GetTimestampFieldOk() (*string, bool)`

GetTimestampFieldOk returns a tuple with the TimestampField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampField

`func (o *ExperimentDataWarehouseNode) SetTimestampField(v string)`

SetTimestampField sets TimestampField field to given value.


### GetVersion

`func (o *ExperimentDataWarehouseNode) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExperimentDataWarehouseNode) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExperimentDataWarehouseNode) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExperimentDataWarehouseNode) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ExperimentDataWarehouseNode) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ExperimentDataWarehouseNode) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


