# ConversionGoalFilter3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversionGoalId** | **string** |  | 
**ConversionGoalName** | **string** |  | 
**CustomName** | Pointer to **NullableString** |  | [optional] 
**DistinctIdField** | **string** |  | 
**DwSourceType** | Pointer to **NullableString** |  | [optional] 
**FixedProperties** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) | Fixed properties in the query, can&#39;t be edited in the interface (e.g. scoping down by person) | [optional] 
**Id** | **string** |  | 
**IdField** | **string** |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "DataWarehouseNode"]
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
**TableName** | **string** |  | 
**TimestampField** | **string** |  | 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewConversionGoalFilter3

`func NewConversionGoalFilter3(conversionGoalId string, conversionGoalName string, distinctIdField string, id string, idField string, schemaMap map[string]SchemaMapValue, tableName string, timestampField string, ) *ConversionGoalFilter3`

NewConversionGoalFilter3 instantiates a new ConversionGoalFilter3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionGoalFilter3WithDefaults

`func NewConversionGoalFilter3WithDefaults() *ConversionGoalFilter3`

NewConversionGoalFilter3WithDefaults instantiates a new ConversionGoalFilter3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversionGoalId

`func (o *ConversionGoalFilter3) GetConversionGoalId() string`

GetConversionGoalId returns the ConversionGoalId field if non-nil, zero value otherwise.

### GetConversionGoalIdOk

`func (o *ConversionGoalFilter3) GetConversionGoalIdOk() (*string, bool)`

GetConversionGoalIdOk returns a tuple with the ConversionGoalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoalId

`func (o *ConversionGoalFilter3) SetConversionGoalId(v string)`

SetConversionGoalId sets ConversionGoalId field to given value.


### GetConversionGoalName

`func (o *ConversionGoalFilter3) GetConversionGoalName() string`

GetConversionGoalName returns the ConversionGoalName field if non-nil, zero value otherwise.

### GetConversionGoalNameOk

`func (o *ConversionGoalFilter3) GetConversionGoalNameOk() (*string, bool)`

GetConversionGoalNameOk returns a tuple with the ConversionGoalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoalName

`func (o *ConversionGoalFilter3) SetConversionGoalName(v string)`

SetConversionGoalName sets ConversionGoalName field to given value.


### GetCustomName

`func (o *ConversionGoalFilter3) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *ConversionGoalFilter3) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *ConversionGoalFilter3) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *ConversionGoalFilter3) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *ConversionGoalFilter3) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *ConversionGoalFilter3) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetDistinctIdField

`func (o *ConversionGoalFilter3) GetDistinctIdField() string`

GetDistinctIdField returns the DistinctIdField field if non-nil, zero value otherwise.

### GetDistinctIdFieldOk

`func (o *ConversionGoalFilter3) GetDistinctIdFieldOk() (*string, bool)`

GetDistinctIdFieldOk returns a tuple with the DistinctIdField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctIdField

`func (o *ConversionGoalFilter3) SetDistinctIdField(v string)`

SetDistinctIdField sets DistinctIdField field to given value.


### GetDwSourceType

`func (o *ConversionGoalFilter3) GetDwSourceType() string`

GetDwSourceType returns the DwSourceType field if non-nil, zero value otherwise.

### GetDwSourceTypeOk

`func (o *ConversionGoalFilter3) GetDwSourceTypeOk() (*string, bool)`

GetDwSourceTypeOk returns a tuple with the DwSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDwSourceType

`func (o *ConversionGoalFilter3) SetDwSourceType(v string)`

SetDwSourceType sets DwSourceType field to given value.

### HasDwSourceType

`func (o *ConversionGoalFilter3) HasDwSourceType() bool`

HasDwSourceType returns a boolean if a field has been set.

### SetDwSourceTypeNil

`func (o *ConversionGoalFilter3) SetDwSourceTypeNil(b bool)`

 SetDwSourceTypeNil sets the value for DwSourceType to be an explicit nil

### UnsetDwSourceType
`func (o *ConversionGoalFilter3) UnsetDwSourceType()`

UnsetDwSourceType ensures that no value is present for DwSourceType, not even an explicit nil
### GetFixedProperties

`func (o *ConversionGoalFilter3) GetFixedProperties() []FixedpropertiesInner`

GetFixedProperties returns the FixedProperties field if non-nil, zero value otherwise.

### GetFixedPropertiesOk

`func (o *ConversionGoalFilter3) GetFixedPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetFixedPropertiesOk returns a tuple with the FixedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedProperties

`func (o *ConversionGoalFilter3) SetFixedProperties(v []FixedpropertiesInner)`

SetFixedProperties sets FixedProperties field to given value.

### HasFixedProperties

`func (o *ConversionGoalFilter3) HasFixedProperties() bool`

HasFixedProperties returns a boolean if a field has been set.

### SetFixedPropertiesNil

`func (o *ConversionGoalFilter3) SetFixedPropertiesNil(b bool)`

 SetFixedPropertiesNil sets the value for FixedProperties to be an explicit nil

### UnsetFixedProperties
`func (o *ConversionGoalFilter3) UnsetFixedProperties()`

UnsetFixedProperties ensures that no value is present for FixedProperties, not even an explicit nil
### GetId

`func (o *ConversionGoalFilter3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConversionGoalFilter3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConversionGoalFilter3) SetId(v string)`

SetId sets Id field to given value.


### GetIdField

`func (o *ConversionGoalFilter3) GetIdField() string`

GetIdField returns the IdField field if non-nil, zero value otherwise.

### GetIdFieldOk

`func (o *ConversionGoalFilter3) GetIdFieldOk() (*string, bool)`

GetIdFieldOk returns a tuple with the IdField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdField

`func (o *ConversionGoalFilter3) SetIdField(v string)`

SetIdField sets IdField field to given value.


### GetKind

`func (o *ConversionGoalFilter3) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConversionGoalFilter3) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConversionGoalFilter3) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConversionGoalFilter3) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMath

`func (o *ConversionGoalFilter3) GetMath() string`

GetMath returns the Math field if non-nil, zero value otherwise.

### GetMathOk

`func (o *ConversionGoalFilter3) GetMathOk() (*string, bool)`

GetMathOk returns a tuple with the Math field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMath

`func (o *ConversionGoalFilter3) SetMath(v string)`

SetMath sets Math field to given value.

### HasMath

`func (o *ConversionGoalFilter3) HasMath() bool`

HasMath returns a boolean if a field has been set.

### SetMathNil

`func (o *ConversionGoalFilter3) SetMathNil(b bool)`

 SetMathNil sets the value for Math to be an explicit nil

### UnsetMath
`func (o *ConversionGoalFilter3) UnsetMath()`

UnsetMath ensures that no value is present for Math, not even an explicit nil
### GetMathGroupTypeIndex

`func (o *ConversionGoalFilter3) GetMathGroupTypeIndex() MathGroupTypeIndex`

GetMathGroupTypeIndex returns the MathGroupTypeIndex field if non-nil, zero value otherwise.

### GetMathGroupTypeIndexOk

`func (o *ConversionGoalFilter3) GetMathGroupTypeIndexOk() (*MathGroupTypeIndex, bool)`

GetMathGroupTypeIndexOk returns a tuple with the MathGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathGroupTypeIndex

`func (o *ConversionGoalFilter3) SetMathGroupTypeIndex(v MathGroupTypeIndex)`

SetMathGroupTypeIndex sets MathGroupTypeIndex field to given value.

### HasMathGroupTypeIndex

`func (o *ConversionGoalFilter3) HasMathGroupTypeIndex() bool`

HasMathGroupTypeIndex returns a boolean if a field has been set.

### GetMathHogql

`func (o *ConversionGoalFilter3) GetMathHogql() string`

GetMathHogql returns the MathHogql field if non-nil, zero value otherwise.

### GetMathHogqlOk

`func (o *ConversionGoalFilter3) GetMathHogqlOk() (*string, bool)`

GetMathHogqlOk returns a tuple with the MathHogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathHogql

`func (o *ConversionGoalFilter3) SetMathHogql(v string)`

SetMathHogql sets MathHogql field to given value.

### HasMathHogql

`func (o *ConversionGoalFilter3) HasMathHogql() bool`

HasMathHogql returns a boolean if a field has been set.

### SetMathHogqlNil

`func (o *ConversionGoalFilter3) SetMathHogqlNil(b bool)`

 SetMathHogqlNil sets the value for MathHogql to be an explicit nil

### UnsetMathHogql
`func (o *ConversionGoalFilter3) UnsetMathHogql()`

UnsetMathHogql ensures that no value is present for MathHogql, not even an explicit nil
### GetMathMultiplier

`func (o *ConversionGoalFilter3) GetMathMultiplier() float32`

GetMathMultiplier returns the MathMultiplier field if non-nil, zero value otherwise.

### GetMathMultiplierOk

`func (o *ConversionGoalFilter3) GetMathMultiplierOk() (*float32, bool)`

GetMathMultiplierOk returns a tuple with the MathMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathMultiplier

`func (o *ConversionGoalFilter3) SetMathMultiplier(v float32)`

SetMathMultiplier sets MathMultiplier field to given value.

### HasMathMultiplier

`func (o *ConversionGoalFilter3) HasMathMultiplier() bool`

HasMathMultiplier returns a boolean if a field has been set.

### SetMathMultiplierNil

`func (o *ConversionGoalFilter3) SetMathMultiplierNil(b bool)`

 SetMathMultiplierNil sets the value for MathMultiplier to be an explicit nil

### UnsetMathMultiplier
`func (o *ConversionGoalFilter3) UnsetMathMultiplier()`

UnsetMathMultiplier ensures that no value is present for MathMultiplier, not even an explicit nil
### GetMathProperty

`func (o *ConversionGoalFilter3) GetMathProperty() string`

GetMathProperty returns the MathProperty field if non-nil, zero value otherwise.

### GetMathPropertyOk

`func (o *ConversionGoalFilter3) GetMathPropertyOk() (*string, bool)`

GetMathPropertyOk returns a tuple with the MathProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathProperty

`func (o *ConversionGoalFilter3) SetMathProperty(v string)`

SetMathProperty sets MathProperty field to given value.

### HasMathProperty

`func (o *ConversionGoalFilter3) HasMathProperty() bool`

HasMathProperty returns a boolean if a field has been set.

### SetMathPropertyNil

`func (o *ConversionGoalFilter3) SetMathPropertyNil(b bool)`

 SetMathPropertyNil sets the value for MathProperty to be an explicit nil

### UnsetMathProperty
`func (o *ConversionGoalFilter3) UnsetMathProperty()`

UnsetMathProperty ensures that no value is present for MathProperty, not even an explicit nil
### GetMathPropertyRevenueCurrency

`func (o *ConversionGoalFilter3) GetMathPropertyRevenueCurrency() RevenueCurrencyPropertyConfig`

GetMathPropertyRevenueCurrency returns the MathPropertyRevenueCurrency field if non-nil, zero value otherwise.

### GetMathPropertyRevenueCurrencyOk

`func (o *ConversionGoalFilter3) GetMathPropertyRevenueCurrencyOk() (*RevenueCurrencyPropertyConfig, bool)`

GetMathPropertyRevenueCurrencyOk returns a tuple with the MathPropertyRevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyRevenueCurrency

`func (o *ConversionGoalFilter3) SetMathPropertyRevenueCurrency(v RevenueCurrencyPropertyConfig)`

SetMathPropertyRevenueCurrency sets MathPropertyRevenueCurrency field to given value.

### HasMathPropertyRevenueCurrency

`func (o *ConversionGoalFilter3) HasMathPropertyRevenueCurrency() bool`

HasMathPropertyRevenueCurrency returns a boolean if a field has been set.

### GetMathPropertyType

`func (o *ConversionGoalFilter3) GetMathPropertyType() string`

GetMathPropertyType returns the MathPropertyType field if non-nil, zero value otherwise.

### GetMathPropertyTypeOk

`func (o *ConversionGoalFilter3) GetMathPropertyTypeOk() (*string, bool)`

GetMathPropertyTypeOk returns a tuple with the MathPropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyType

`func (o *ConversionGoalFilter3) SetMathPropertyType(v string)`

SetMathPropertyType sets MathPropertyType field to given value.

### HasMathPropertyType

`func (o *ConversionGoalFilter3) HasMathPropertyType() bool`

HasMathPropertyType returns a boolean if a field has been set.

### SetMathPropertyTypeNil

`func (o *ConversionGoalFilter3) SetMathPropertyTypeNil(b bool)`

 SetMathPropertyTypeNil sets the value for MathPropertyType to be an explicit nil

### UnsetMathPropertyType
`func (o *ConversionGoalFilter3) UnsetMathPropertyType()`

UnsetMathPropertyType ensures that no value is present for MathPropertyType, not even an explicit nil
### GetName

`func (o *ConversionGoalFilter3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConversionGoalFilter3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConversionGoalFilter3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConversionGoalFilter3) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConversionGoalFilter3) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConversionGoalFilter3) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOptionalInFunnel

`func (o *ConversionGoalFilter3) GetOptionalInFunnel() bool`

GetOptionalInFunnel returns the OptionalInFunnel field if non-nil, zero value otherwise.

### GetOptionalInFunnelOk

`func (o *ConversionGoalFilter3) GetOptionalInFunnelOk() (*bool, bool)`

GetOptionalInFunnelOk returns a tuple with the OptionalInFunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalInFunnel

`func (o *ConversionGoalFilter3) SetOptionalInFunnel(v bool)`

SetOptionalInFunnel sets OptionalInFunnel field to given value.

### HasOptionalInFunnel

`func (o *ConversionGoalFilter3) HasOptionalInFunnel() bool`

HasOptionalInFunnel returns a boolean if a field has been set.

### SetOptionalInFunnelNil

`func (o *ConversionGoalFilter3) SetOptionalInFunnelNil(b bool)`

 SetOptionalInFunnelNil sets the value for OptionalInFunnel to be an explicit nil

### UnsetOptionalInFunnel
`func (o *ConversionGoalFilter3) UnsetOptionalInFunnel()`

UnsetOptionalInFunnel ensures that no value is present for OptionalInFunnel, not even an explicit nil
### GetProperties

`func (o *ConversionGoalFilter3) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ConversionGoalFilter3) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ConversionGoalFilter3) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ConversionGoalFilter3) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ConversionGoalFilter3) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ConversionGoalFilter3) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *ConversionGoalFilter3) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ConversionGoalFilter3) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ConversionGoalFilter3) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ConversionGoalFilter3) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *ConversionGoalFilter3) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *ConversionGoalFilter3) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetSchemaMap

`func (o *ConversionGoalFilter3) GetSchemaMap() map[string]SchemaMapValue`

GetSchemaMap returns the SchemaMap field if non-nil, zero value otherwise.

### GetSchemaMapOk

`func (o *ConversionGoalFilter3) GetSchemaMapOk() (*map[string]SchemaMapValue, bool)`

GetSchemaMapOk returns a tuple with the SchemaMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaMap

`func (o *ConversionGoalFilter3) SetSchemaMap(v map[string]SchemaMapValue)`

SetSchemaMap sets SchemaMap field to given value.


### GetTableName

`func (o *ConversionGoalFilter3) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *ConversionGoalFilter3) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *ConversionGoalFilter3) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetTimestampField

`func (o *ConversionGoalFilter3) GetTimestampField() string`

GetTimestampField returns the TimestampField field if non-nil, zero value otherwise.

### GetTimestampFieldOk

`func (o *ConversionGoalFilter3) GetTimestampFieldOk() (*string, bool)`

GetTimestampFieldOk returns a tuple with the TimestampField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampField

`func (o *ConversionGoalFilter3) SetTimestampField(v string)`

SetTimestampField sets TimestampField field to given value.


### GetVersion

`func (o *ConversionGoalFilter3) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConversionGoalFilter3) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConversionGoalFilter3) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConversionGoalFilter3) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ConversionGoalFilter3) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ConversionGoalFilter3) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


