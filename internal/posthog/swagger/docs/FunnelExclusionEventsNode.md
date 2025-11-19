# FunnelExclusionEventsNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomName** | Pointer to **NullableString** |  | [optional] 
**Event** | Pointer to **NullableString** | The event or &#x60;null&#x60; for all events. | [optional] 
**FixedProperties** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) | Fixed properties in the query, can&#39;t be edited in the interface (e.g. scoping down by person) | [optional] 
**FunnelFromStep** | **int32** |  | 
**FunnelToStep** | **int32** |  | 
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
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewFunnelExclusionEventsNode

`func NewFunnelExclusionEventsNode(funnelFromStep int32, funnelToStep int32, ) *FunnelExclusionEventsNode`

NewFunnelExclusionEventsNode instantiates a new FunnelExclusionEventsNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunnelExclusionEventsNodeWithDefaults

`func NewFunnelExclusionEventsNodeWithDefaults() *FunnelExclusionEventsNode`

NewFunnelExclusionEventsNodeWithDefaults instantiates a new FunnelExclusionEventsNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomName

`func (o *FunnelExclusionEventsNode) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *FunnelExclusionEventsNode) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *FunnelExclusionEventsNode) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *FunnelExclusionEventsNode) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *FunnelExclusionEventsNode) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *FunnelExclusionEventsNode) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetEvent

`func (o *FunnelExclusionEventsNode) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *FunnelExclusionEventsNode) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *FunnelExclusionEventsNode) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *FunnelExclusionEventsNode) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *FunnelExclusionEventsNode) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *FunnelExclusionEventsNode) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetFixedProperties

`func (o *FunnelExclusionEventsNode) GetFixedProperties() []FixedpropertiesInner`

GetFixedProperties returns the FixedProperties field if non-nil, zero value otherwise.

### GetFixedPropertiesOk

`func (o *FunnelExclusionEventsNode) GetFixedPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetFixedPropertiesOk returns a tuple with the FixedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedProperties

`func (o *FunnelExclusionEventsNode) SetFixedProperties(v []FixedpropertiesInner)`

SetFixedProperties sets FixedProperties field to given value.

### HasFixedProperties

`func (o *FunnelExclusionEventsNode) HasFixedProperties() bool`

HasFixedProperties returns a boolean if a field has been set.

### SetFixedPropertiesNil

`func (o *FunnelExclusionEventsNode) SetFixedPropertiesNil(b bool)`

 SetFixedPropertiesNil sets the value for FixedProperties to be an explicit nil

### UnsetFixedProperties
`func (o *FunnelExclusionEventsNode) UnsetFixedProperties()`

UnsetFixedProperties ensures that no value is present for FixedProperties, not even an explicit nil
### GetFunnelFromStep

`func (o *FunnelExclusionEventsNode) GetFunnelFromStep() int32`

GetFunnelFromStep returns the FunnelFromStep field if non-nil, zero value otherwise.

### GetFunnelFromStepOk

`func (o *FunnelExclusionEventsNode) GetFunnelFromStepOk() (*int32, bool)`

GetFunnelFromStepOk returns a tuple with the FunnelFromStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelFromStep

`func (o *FunnelExclusionEventsNode) SetFunnelFromStep(v int32)`

SetFunnelFromStep sets FunnelFromStep field to given value.


### GetFunnelToStep

`func (o *FunnelExclusionEventsNode) GetFunnelToStep() int32`

GetFunnelToStep returns the FunnelToStep field if non-nil, zero value otherwise.

### GetFunnelToStepOk

`func (o *FunnelExclusionEventsNode) GetFunnelToStepOk() (*int32, bool)`

GetFunnelToStepOk returns a tuple with the FunnelToStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelToStep

`func (o *FunnelExclusionEventsNode) SetFunnelToStep(v int32)`

SetFunnelToStep sets FunnelToStep field to given value.


### GetKind

`func (o *FunnelExclusionEventsNode) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FunnelExclusionEventsNode) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FunnelExclusionEventsNode) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FunnelExclusionEventsNode) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *FunnelExclusionEventsNode) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *FunnelExclusionEventsNode) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *FunnelExclusionEventsNode) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *FunnelExclusionEventsNode) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *FunnelExclusionEventsNode) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *FunnelExclusionEventsNode) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetMath

`func (o *FunnelExclusionEventsNode) GetMath() string`

GetMath returns the Math field if non-nil, zero value otherwise.

### GetMathOk

`func (o *FunnelExclusionEventsNode) GetMathOk() (*string, bool)`

GetMathOk returns a tuple with the Math field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMath

`func (o *FunnelExclusionEventsNode) SetMath(v string)`

SetMath sets Math field to given value.

### HasMath

`func (o *FunnelExclusionEventsNode) HasMath() bool`

HasMath returns a boolean if a field has been set.

### SetMathNil

`func (o *FunnelExclusionEventsNode) SetMathNil(b bool)`

 SetMathNil sets the value for Math to be an explicit nil

### UnsetMath
`func (o *FunnelExclusionEventsNode) UnsetMath()`

UnsetMath ensures that no value is present for Math, not even an explicit nil
### GetMathGroupTypeIndex

`func (o *FunnelExclusionEventsNode) GetMathGroupTypeIndex() MathGroupTypeIndex`

GetMathGroupTypeIndex returns the MathGroupTypeIndex field if non-nil, zero value otherwise.

### GetMathGroupTypeIndexOk

`func (o *FunnelExclusionEventsNode) GetMathGroupTypeIndexOk() (*MathGroupTypeIndex, bool)`

GetMathGroupTypeIndexOk returns a tuple with the MathGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathGroupTypeIndex

`func (o *FunnelExclusionEventsNode) SetMathGroupTypeIndex(v MathGroupTypeIndex)`

SetMathGroupTypeIndex sets MathGroupTypeIndex field to given value.

### HasMathGroupTypeIndex

`func (o *FunnelExclusionEventsNode) HasMathGroupTypeIndex() bool`

HasMathGroupTypeIndex returns a boolean if a field has been set.

### GetMathHogql

`func (o *FunnelExclusionEventsNode) GetMathHogql() string`

GetMathHogql returns the MathHogql field if non-nil, zero value otherwise.

### GetMathHogqlOk

`func (o *FunnelExclusionEventsNode) GetMathHogqlOk() (*string, bool)`

GetMathHogqlOk returns a tuple with the MathHogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathHogql

`func (o *FunnelExclusionEventsNode) SetMathHogql(v string)`

SetMathHogql sets MathHogql field to given value.

### HasMathHogql

`func (o *FunnelExclusionEventsNode) HasMathHogql() bool`

HasMathHogql returns a boolean if a field has been set.

### SetMathHogqlNil

`func (o *FunnelExclusionEventsNode) SetMathHogqlNil(b bool)`

 SetMathHogqlNil sets the value for MathHogql to be an explicit nil

### UnsetMathHogql
`func (o *FunnelExclusionEventsNode) UnsetMathHogql()`

UnsetMathHogql ensures that no value is present for MathHogql, not even an explicit nil
### GetMathMultiplier

`func (o *FunnelExclusionEventsNode) GetMathMultiplier() float32`

GetMathMultiplier returns the MathMultiplier field if non-nil, zero value otherwise.

### GetMathMultiplierOk

`func (o *FunnelExclusionEventsNode) GetMathMultiplierOk() (*float32, bool)`

GetMathMultiplierOk returns a tuple with the MathMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathMultiplier

`func (o *FunnelExclusionEventsNode) SetMathMultiplier(v float32)`

SetMathMultiplier sets MathMultiplier field to given value.

### HasMathMultiplier

`func (o *FunnelExclusionEventsNode) HasMathMultiplier() bool`

HasMathMultiplier returns a boolean if a field has been set.

### SetMathMultiplierNil

`func (o *FunnelExclusionEventsNode) SetMathMultiplierNil(b bool)`

 SetMathMultiplierNil sets the value for MathMultiplier to be an explicit nil

### UnsetMathMultiplier
`func (o *FunnelExclusionEventsNode) UnsetMathMultiplier()`

UnsetMathMultiplier ensures that no value is present for MathMultiplier, not even an explicit nil
### GetMathProperty

`func (o *FunnelExclusionEventsNode) GetMathProperty() string`

GetMathProperty returns the MathProperty field if non-nil, zero value otherwise.

### GetMathPropertyOk

`func (o *FunnelExclusionEventsNode) GetMathPropertyOk() (*string, bool)`

GetMathPropertyOk returns a tuple with the MathProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathProperty

`func (o *FunnelExclusionEventsNode) SetMathProperty(v string)`

SetMathProperty sets MathProperty field to given value.

### HasMathProperty

`func (o *FunnelExclusionEventsNode) HasMathProperty() bool`

HasMathProperty returns a boolean if a field has been set.

### SetMathPropertyNil

`func (o *FunnelExclusionEventsNode) SetMathPropertyNil(b bool)`

 SetMathPropertyNil sets the value for MathProperty to be an explicit nil

### UnsetMathProperty
`func (o *FunnelExclusionEventsNode) UnsetMathProperty()`

UnsetMathProperty ensures that no value is present for MathProperty, not even an explicit nil
### GetMathPropertyRevenueCurrency

`func (o *FunnelExclusionEventsNode) GetMathPropertyRevenueCurrency() RevenueCurrencyPropertyConfig`

GetMathPropertyRevenueCurrency returns the MathPropertyRevenueCurrency field if non-nil, zero value otherwise.

### GetMathPropertyRevenueCurrencyOk

`func (o *FunnelExclusionEventsNode) GetMathPropertyRevenueCurrencyOk() (*RevenueCurrencyPropertyConfig, bool)`

GetMathPropertyRevenueCurrencyOk returns a tuple with the MathPropertyRevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyRevenueCurrency

`func (o *FunnelExclusionEventsNode) SetMathPropertyRevenueCurrency(v RevenueCurrencyPropertyConfig)`

SetMathPropertyRevenueCurrency sets MathPropertyRevenueCurrency field to given value.

### HasMathPropertyRevenueCurrency

`func (o *FunnelExclusionEventsNode) HasMathPropertyRevenueCurrency() bool`

HasMathPropertyRevenueCurrency returns a boolean if a field has been set.

### GetMathPropertyType

`func (o *FunnelExclusionEventsNode) GetMathPropertyType() string`

GetMathPropertyType returns the MathPropertyType field if non-nil, zero value otherwise.

### GetMathPropertyTypeOk

`func (o *FunnelExclusionEventsNode) GetMathPropertyTypeOk() (*string, bool)`

GetMathPropertyTypeOk returns a tuple with the MathPropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyType

`func (o *FunnelExclusionEventsNode) SetMathPropertyType(v string)`

SetMathPropertyType sets MathPropertyType field to given value.

### HasMathPropertyType

`func (o *FunnelExclusionEventsNode) HasMathPropertyType() bool`

HasMathPropertyType returns a boolean if a field has been set.

### SetMathPropertyTypeNil

`func (o *FunnelExclusionEventsNode) SetMathPropertyTypeNil(b bool)`

 SetMathPropertyTypeNil sets the value for MathPropertyType to be an explicit nil

### UnsetMathPropertyType
`func (o *FunnelExclusionEventsNode) UnsetMathPropertyType()`

UnsetMathPropertyType ensures that no value is present for MathPropertyType, not even an explicit nil
### GetName

`func (o *FunnelExclusionEventsNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunnelExclusionEventsNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunnelExclusionEventsNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FunnelExclusionEventsNode) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FunnelExclusionEventsNode) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FunnelExclusionEventsNode) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOptionalInFunnel

`func (o *FunnelExclusionEventsNode) GetOptionalInFunnel() bool`

GetOptionalInFunnel returns the OptionalInFunnel field if non-nil, zero value otherwise.

### GetOptionalInFunnelOk

`func (o *FunnelExclusionEventsNode) GetOptionalInFunnelOk() (*bool, bool)`

GetOptionalInFunnelOk returns a tuple with the OptionalInFunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalInFunnel

`func (o *FunnelExclusionEventsNode) SetOptionalInFunnel(v bool)`

SetOptionalInFunnel sets OptionalInFunnel field to given value.

### HasOptionalInFunnel

`func (o *FunnelExclusionEventsNode) HasOptionalInFunnel() bool`

HasOptionalInFunnel returns a boolean if a field has been set.

### SetOptionalInFunnelNil

`func (o *FunnelExclusionEventsNode) SetOptionalInFunnelNil(b bool)`

 SetOptionalInFunnelNil sets the value for OptionalInFunnel to be an explicit nil

### UnsetOptionalInFunnel
`func (o *FunnelExclusionEventsNode) UnsetOptionalInFunnel()`

UnsetOptionalInFunnel ensures that no value is present for OptionalInFunnel, not even an explicit nil
### GetOrderBy

`func (o *FunnelExclusionEventsNode) GetOrderBy() []string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *FunnelExclusionEventsNode) GetOrderByOk() (*[]string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *FunnelExclusionEventsNode) SetOrderBy(v []string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *FunnelExclusionEventsNode) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *FunnelExclusionEventsNode) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *FunnelExclusionEventsNode) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetProperties

`func (o *FunnelExclusionEventsNode) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FunnelExclusionEventsNode) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FunnelExclusionEventsNode) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *FunnelExclusionEventsNode) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *FunnelExclusionEventsNode) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *FunnelExclusionEventsNode) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *FunnelExclusionEventsNode) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *FunnelExclusionEventsNode) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *FunnelExclusionEventsNode) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *FunnelExclusionEventsNode) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *FunnelExclusionEventsNode) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *FunnelExclusionEventsNode) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetVersion

`func (o *FunnelExclusionEventsNode) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FunnelExclusionEventsNode) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FunnelExclusionEventsNode) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FunnelExclusionEventsNode) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *FunnelExclusionEventsNode) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *FunnelExclusionEventsNode) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


