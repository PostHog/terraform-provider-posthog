# Draftconversiongoal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversionGoalId** | **string** |  | 
**ConversionGoalName** | **string** |  | 
**CustomName** | Pointer to **NullableString** |  | [optional] 
**Event** | Pointer to **NullableString** | The event or &#x60;null&#x60; for all events. | [optional] 
**FixedProperties** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) | Fixed properties in the query, can&#39;t be edited in the interface (e.g. scoping down by person) | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "DataWarehouseNode"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Math** | Pointer to **NullableString** |  | [optional] 
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
**Id** | **string** |  | 
**DistinctIdField** | **string** |  | 
**DwSourceType** | Pointer to **NullableString** |  | [optional] 
**IdField** | **string** |  | 
**TableName** | **string** |  | 
**TimestampField** | **string** |  | 

## Methods

### NewDraftconversiongoal

`func NewDraftconversiongoal(conversionGoalId string, conversionGoalName string, schemaMap map[string]SchemaMapValue, id string, distinctIdField string, idField string, tableName string, timestampField string, ) *Draftconversiongoal`

NewDraftconversiongoal instantiates a new Draftconversiongoal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDraftconversiongoalWithDefaults

`func NewDraftconversiongoalWithDefaults() *Draftconversiongoal`

NewDraftconversiongoalWithDefaults instantiates a new Draftconversiongoal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversionGoalId

`func (o *Draftconversiongoal) GetConversionGoalId() string`

GetConversionGoalId returns the ConversionGoalId field if non-nil, zero value otherwise.

### GetConversionGoalIdOk

`func (o *Draftconversiongoal) GetConversionGoalIdOk() (*string, bool)`

GetConversionGoalIdOk returns a tuple with the ConversionGoalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoalId

`func (o *Draftconversiongoal) SetConversionGoalId(v string)`

SetConversionGoalId sets ConversionGoalId field to given value.


### GetConversionGoalName

`func (o *Draftconversiongoal) GetConversionGoalName() string`

GetConversionGoalName returns the ConversionGoalName field if non-nil, zero value otherwise.

### GetConversionGoalNameOk

`func (o *Draftconversiongoal) GetConversionGoalNameOk() (*string, bool)`

GetConversionGoalNameOk returns a tuple with the ConversionGoalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoalName

`func (o *Draftconversiongoal) SetConversionGoalName(v string)`

SetConversionGoalName sets ConversionGoalName field to given value.


### GetCustomName

`func (o *Draftconversiongoal) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *Draftconversiongoal) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *Draftconversiongoal) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *Draftconversiongoal) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *Draftconversiongoal) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *Draftconversiongoal) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetEvent

`func (o *Draftconversiongoal) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Draftconversiongoal) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Draftconversiongoal) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *Draftconversiongoal) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *Draftconversiongoal) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *Draftconversiongoal) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetFixedProperties

`func (o *Draftconversiongoal) GetFixedProperties() []FixedpropertiesInner`

GetFixedProperties returns the FixedProperties field if non-nil, zero value otherwise.

### GetFixedPropertiesOk

`func (o *Draftconversiongoal) GetFixedPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetFixedPropertiesOk returns a tuple with the FixedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedProperties

`func (o *Draftconversiongoal) SetFixedProperties(v []FixedpropertiesInner)`

SetFixedProperties sets FixedProperties field to given value.

### HasFixedProperties

`func (o *Draftconversiongoal) HasFixedProperties() bool`

HasFixedProperties returns a boolean if a field has been set.

### SetFixedPropertiesNil

`func (o *Draftconversiongoal) SetFixedPropertiesNil(b bool)`

 SetFixedPropertiesNil sets the value for FixedProperties to be an explicit nil

### UnsetFixedProperties
`func (o *Draftconversiongoal) UnsetFixedProperties()`

UnsetFixedProperties ensures that no value is present for FixedProperties, not even an explicit nil
### GetKind

`func (o *Draftconversiongoal) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Draftconversiongoal) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Draftconversiongoal) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Draftconversiongoal) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *Draftconversiongoal) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Draftconversiongoal) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Draftconversiongoal) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Draftconversiongoal) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *Draftconversiongoal) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *Draftconversiongoal) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetMath

`func (o *Draftconversiongoal) GetMath() string`

GetMath returns the Math field if non-nil, zero value otherwise.

### GetMathOk

`func (o *Draftconversiongoal) GetMathOk() (*string, bool)`

GetMathOk returns a tuple with the Math field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMath

`func (o *Draftconversiongoal) SetMath(v string)`

SetMath sets Math field to given value.

### HasMath

`func (o *Draftconversiongoal) HasMath() bool`

HasMath returns a boolean if a field has been set.

### SetMathNil

`func (o *Draftconversiongoal) SetMathNil(b bool)`

 SetMathNil sets the value for Math to be an explicit nil

### UnsetMath
`func (o *Draftconversiongoal) UnsetMath()`

UnsetMath ensures that no value is present for Math, not even an explicit nil
### GetMathGroupTypeIndex

`func (o *Draftconversiongoal) GetMathGroupTypeIndex() MathGroupTypeIndex`

GetMathGroupTypeIndex returns the MathGroupTypeIndex field if non-nil, zero value otherwise.

### GetMathGroupTypeIndexOk

`func (o *Draftconversiongoal) GetMathGroupTypeIndexOk() (*MathGroupTypeIndex, bool)`

GetMathGroupTypeIndexOk returns a tuple with the MathGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathGroupTypeIndex

`func (o *Draftconversiongoal) SetMathGroupTypeIndex(v MathGroupTypeIndex)`

SetMathGroupTypeIndex sets MathGroupTypeIndex field to given value.

### HasMathGroupTypeIndex

`func (o *Draftconversiongoal) HasMathGroupTypeIndex() bool`

HasMathGroupTypeIndex returns a boolean if a field has been set.

### GetMathHogql

`func (o *Draftconversiongoal) GetMathHogql() string`

GetMathHogql returns the MathHogql field if non-nil, zero value otherwise.

### GetMathHogqlOk

`func (o *Draftconversiongoal) GetMathHogqlOk() (*string, bool)`

GetMathHogqlOk returns a tuple with the MathHogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathHogql

`func (o *Draftconversiongoal) SetMathHogql(v string)`

SetMathHogql sets MathHogql field to given value.

### HasMathHogql

`func (o *Draftconversiongoal) HasMathHogql() bool`

HasMathHogql returns a boolean if a field has been set.

### SetMathHogqlNil

`func (o *Draftconversiongoal) SetMathHogqlNil(b bool)`

 SetMathHogqlNil sets the value for MathHogql to be an explicit nil

### UnsetMathHogql
`func (o *Draftconversiongoal) UnsetMathHogql()`

UnsetMathHogql ensures that no value is present for MathHogql, not even an explicit nil
### GetMathMultiplier

`func (o *Draftconversiongoal) GetMathMultiplier() float32`

GetMathMultiplier returns the MathMultiplier field if non-nil, zero value otherwise.

### GetMathMultiplierOk

`func (o *Draftconversiongoal) GetMathMultiplierOk() (*float32, bool)`

GetMathMultiplierOk returns a tuple with the MathMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathMultiplier

`func (o *Draftconversiongoal) SetMathMultiplier(v float32)`

SetMathMultiplier sets MathMultiplier field to given value.

### HasMathMultiplier

`func (o *Draftconversiongoal) HasMathMultiplier() bool`

HasMathMultiplier returns a boolean if a field has been set.

### SetMathMultiplierNil

`func (o *Draftconversiongoal) SetMathMultiplierNil(b bool)`

 SetMathMultiplierNil sets the value for MathMultiplier to be an explicit nil

### UnsetMathMultiplier
`func (o *Draftconversiongoal) UnsetMathMultiplier()`

UnsetMathMultiplier ensures that no value is present for MathMultiplier, not even an explicit nil
### GetMathProperty

`func (o *Draftconversiongoal) GetMathProperty() string`

GetMathProperty returns the MathProperty field if non-nil, zero value otherwise.

### GetMathPropertyOk

`func (o *Draftconversiongoal) GetMathPropertyOk() (*string, bool)`

GetMathPropertyOk returns a tuple with the MathProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathProperty

`func (o *Draftconversiongoal) SetMathProperty(v string)`

SetMathProperty sets MathProperty field to given value.

### HasMathProperty

`func (o *Draftconversiongoal) HasMathProperty() bool`

HasMathProperty returns a boolean if a field has been set.

### SetMathPropertyNil

`func (o *Draftconversiongoal) SetMathPropertyNil(b bool)`

 SetMathPropertyNil sets the value for MathProperty to be an explicit nil

### UnsetMathProperty
`func (o *Draftconversiongoal) UnsetMathProperty()`

UnsetMathProperty ensures that no value is present for MathProperty, not even an explicit nil
### GetMathPropertyRevenueCurrency

`func (o *Draftconversiongoal) GetMathPropertyRevenueCurrency() RevenueCurrencyPropertyConfig`

GetMathPropertyRevenueCurrency returns the MathPropertyRevenueCurrency field if non-nil, zero value otherwise.

### GetMathPropertyRevenueCurrencyOk

`func (o *Draftconversiongoal) GetMathPropertyRevenueCurrencyOk() (*RevenueCurrencyPropertyConfig, bool)`

GetMathPropertyRevenueCurrencyOk returns a tuple with the MathPropertyRevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyRevenueCurrency

`func (o *Draftconversiongoal) SetMathPropertyRevenueCurrency(v RevenueCurrencyPropertyConfig)`

SetMathPropertyRevenueCurrency sets MathPropertyRevenueCurrency field to given value.

### HasMathPropertyRevenueCurrency

`func (o *Draftconversiongoal) HasMathPropertyRevenueCurrency() bool`

HasMathPropertyRevenueCurrency returns a boolean if a field has been set.

### GetMathPropertyType

`func (o *Draftconversiongoal) GetMathPropertyType() string`

GetMathPropertyType returns the MathPropertyType field if non-nil, zero value otherwise.

### GetMathPropertyTypeOk

`func (o *Draftconversiongoal) GetMathPropertyTypeOk() (*string, bool)`

GetMathPropertyTypeOk returns a tuple with the MathPropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyType

`func (o *Draftconversiongoal) SetMathPropertyType(v string)`

SetMathPropertyType sets MathPropertyType field to given value.

### HasMathPropertyType

`func (o *Draftconversiongoal) HasMathPropertyType() bool`

HasMathPropertyType returns a boolean if a field has been set.

### SetMathPropertyTypeNil

`func (o *Draftconversiongoal) SetMathPropertyTypeNil(b bool)`

 SetMathPropertyTypeNil sets the value for MathPropertyType to be an explicit nil

### UnsetMathPropertyType
`func (o *Draftconversiongoal) UnsetMathPropertyType()`

UnsetMathPropertyType ensures that no value is present for MathPropertyType, not even an explicit nil
### GetName

`func (o *Draftconversiongoal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Draftconversiongoal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Draftconversiongoal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Draftconversiongoal) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Draftconversiongoal) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Draftconversiongoal) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOptionalInFunnel

`func (o *Draftconversiongoal) GetOptionalInFunnel() bool`

GetOptionalInFunnel returns the OptionalInFunnel field if non-nil, zero value otherwise.

### GetOptionalInFunnelOk

`func (o *Draftconversiongoal) GetOptionalInFunnelOk() (*bool, bool)`

GetOptionalInFunnelOk returns a tuple with the OptionalInFunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalInFunnel

`func (o *Draftconversiongoal) SetOptionalInFunnel(v bool)`

SetOptionalInFunnel sets OptionalInFunnel field to given value.

### HasOptionalInFunnel

`func (o *Draftconversiongoal) HasOptionalInFunnel() bool`

HasOptionalInFunnel returns a boolean if a field has been set.

### SetOptionalInFunnelNil

`func (o *Draftconversiongoal) SetOptionalInFunnelNil(b bool)`

 SetOptionalInFunnelNil sets the value for OptionalInFunnel to be an explicit nil

### UnsetOptionalInFunnel
`func (o *Draftconversiongoal) UnsetOptionalInFunnel()`

UnsetOptionalInFunnel ensures that no value is present for OptionalInFunnel, not even an explicit nil
### GetOrderBy

`func (o *Draftconversiongoal) GetOrderBy() []string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *Draftconversiongoal) GetOrderByOk() (*[]string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *Draftconversiongoal) SetOrderBy(v []string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *Draftconversiongoal) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *Draftconversiongoal) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *Draftconversiongoal) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetProperties

`func (o *Draftconversiongoal) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Draftconversiongoal) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Draftconversiongoal) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Draftconversiongoal) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *Draftconversiongoal) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *Draftconversiongoal) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *Draftconversiongoal) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Draftconversiongoal) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Draftconversiongoal) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Draftconversiongoal) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *Draftconversiongoal) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *Draftconversiongoal) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetSchemaMap

`func (o *Draftconversiongoal) GetSchemaMap() map[string]SchemaMapValue`

GetSchemaMap returns the SchemaMap field if non-nil, zero value otherwise.

### GetSchemaMapOk

`func (o *Draftconversiongoal) GetSchemaMapOk() (*map[string]SchemaMapValue, bool)`

GetSchemaMapOk returns a tuple with the SchemaMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaMap

`func (o *Draftconversiongoal) SetSchemaMap(v map[string]SchemaMapValue)`

SetSchemaMap sets SchemaMap field to given value.


### GetVersion

`func (o *Draftconversiongoal) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Draftconversiongoal) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Draftconversiongoal) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Draftconversiongoal) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Draftconversiongoal) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Draftconversiongoal) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetId

`func (o *Draftconversiongoal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Draftconversiongoal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Draftconversiongoal) SetId(v string)`

SetId sets Id field to given value.


### GetDistinctIdField

`func (o *Draftconversiongoal) GetDistinctIdField() string`

GetDistinctIdField returns the DistinctIdField field if non-nil, zero value otherwise.

### GetDistinctIdFieldOk

`func (o *Draftconversiongoal) GetDistinctIdFieldOk() (*string, bool)`

GetDistinctIdFieldOk returns a tuple with the DistinctIdField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctIdField

`func (o *Draftconversiongoal) SetDistinctIdField(v string)`

SetDistinctIdField sets DistinctIdField field to given value.


### GetDwSourceType

`func (o *Draftconversiongoal) GetDwSourceType() string`

GetDwSourceType returns the DwSourceType field if non-nil, zero value otherwise.

### GetDwSourceTypeOk

`func (o *Draftconversiongoal) GetDwSourceTypeOk() (*string, bool)`

GetDwSourceTypeOk returns a tuple with the DwSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDwSourceType

`func (o *Draftconversiongoal) SetDwSourceType(v string)`

SetDwSourceType sets DwSourceType field to given value.

### HasDwSourceType

`func (o *Draftconversiongoal) HasDwSourceType() bool`

HasDwSourceType returns a boolean if a field has been set.

### SetDwSourceTypeNil

`func (o *Draftconversiongoal) SetDwSourceTypeNil(b bool)`

 SetDwSourceTypeNil sets the value for DwSourceType to be an explicit nil

### UnsetDwSourceType
`func (o *Draftconversiongoal) UnsetDwSourceType()`

UnsetDwSourceType ensures that no value is present for DwSourceType, not even an explicit nil
### GetIdField

`func (o *Draftconversiongoal) GetIdField() string`

GetIdField returns the IdField field if non-nil, zero value otherwise.

### GetIdFieldOk

`func (o *Draftconversiongoal) GetIdFieldOk() (*string, bool)`

GetIdFieldOk returns a tuple with the IdField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdField

`func (o *Draftconversiongoal) SetIdField(v string)`

SetIdField sets IdField field to given value.


### GetTableName

`func (o *Draftconversiongoal) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *Draftconversiongoal) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *Draftconversiongoal) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetTimestampField

`func (o *Draftconversiongoal) GetTimestampField() string`

GetTimestampField returns the TimestampField field if non-nil, zero value otherwise.

### GetTimestampFieldOk

`func (o *Draftconversiongoal) GetTimestampFieldOk() (*string, bool)`

GetTimestampFieldOk returns a tuple with the TimestampField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampField

`func (o *Draftconversiongoal) SetTimestampField(v string)`

SetTimestampField sets TimestampField field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


