# TablesValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**map[string]DatabaseSchemaField**](DatabaseSchemaField.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**RowCount** | Pointer to **NullableFloat32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "posthog"]
**Format** | **string** |  | 
**Schema** | Pointer to [**DatabaseSchemaSchema**](DatabaseSchemaSchema.md) |  | [optional] 
**Source** | Pointer to [**DatabaseSchemaSource**](DatabaseSchemaSource.md) |  | [optional] 
**UrlPattern** | **string** |  | 
**Query** | [**HogQLQuery**](HogQLQuery.md) |  | 
**Kind** | [**DatabaseSchemaManagedViewTableKind**](DatabaseSchemaManagedViewTableKind.md) |  | 
**SourceId** | Pointer to **NullableString** |  | [optional] 
**LastRunAt** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTablesValue

`func NewTablesValue(fields map[string]DatabaseSchemaField, id string, name string, format string, urlPattern string, query HogQLQuery, kind DatabaseSchemaManagedViewTableKind, ) *TablesValue`

NewTablesValue instantiates a new TablesValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTablesValueWithDefaults

`func NewTablesValueWithDefaults() *TablesValue`

NewTablesValueWithDefaults instantiates a new TablesValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *TablesValue) GetFields() map[string]DatabaseSchemaField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TablesValue) GetFieldsOk() (*map[string]DatabaseSchemaField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TablesValue) SetFields(v map[string]DatabaseSchemaField)`

SetFields sets Fields field to given value.


### GetId

`func (o *TablesValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TablesValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TablesValue) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TablesValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TablesValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TablesValue) SetName(v string)`

SetName sets Name field to given value.


### GetRowCount

`func (o *TablesValue) GetRowCount() float32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *TablesValue) GetRowCountOk() (*float32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *TablesValue) SetRowCount(v float32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *TablesValue) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### SetRowCountNil

`func (o *TablesValue) SetRowCountNil(b bool)`

 SetRowCountNil sets the value for RowCount to be an explicit nil

### UnsetRowCount
`func (o *TablesValue) UnsetRowCount()`

UnsetRowCount ensures that no value is present for RowCount, not even an explicit nil
### GetType

`func (o *TablesValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TablesValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TablesValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TablesValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFormat

`func (o *TablesValue) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TablesValue) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TablesValue) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetSchema

`func (o *TablesValue) GetSchema() DatabaseSchemaSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *TablesValue) GetSchemaOk() (*DatabaseSchemaSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *TablesValue) SetSchema(v DatabaseSchemaSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *TablesValue) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSource

`func (o *TablesValue) GetSource() DatabaseSchemaSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TablesValue) GetSourceOk() (*DatabaseSchemaSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TablesValue) SetSource(v DatabaseSchemaSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *TablesValue) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUrlPattern

`func (o *TablesValue) GetUrlPattern() string`

GetUrlPattern returns the UrlPattern field if non-nil, zero value otherwise.

### GetUrlPatternOk

`func (o *TablesValue) GetUrlPatternOk() (*string, bool)`

GetUrlPatternOk returns a tuple with the UrlPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPattern

`func (o *TablesValue) SetUrlPattern(v string)`

SetUrlPattern sets UrlPattern field to given value.


### GetQuery

`func (o *TablesValue) GetQuery() HogQLQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TablesValue) GetQueryOk() (*HogQLQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TablesValue) SetQuery(v HogQLQuery)`

SetQuery sets Query field to given value.


### GetKind

`func (o *TablesValue) GetKind() DatabaseSchemaManagedViewTableKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TablesValue) GetKindOk() (*DatabaseSchemaManagedViewTableKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TablesValue) SetKind(v DatabaseSchemaManagedViewTableKind)`

SetKind sets Kind field to given value.


### GetSourceId

`func (o *TablesValue) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *TablesValue) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *TablesValue) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *TablesValue) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *TablesValue) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *TablesValue) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetLastRunAt

`func (o *TablesValue) GetLastRunAt() string`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *TablesValue) GetLastRunAtOk() (*string, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *TablesValue) SetLastRunAt(v string)`

SetLastRunAt sets LastRunAt field to given value.

### HasLastRunAt

`func (o *TablesValue) HasLastRunAt() bool`

HasLastRunAt returns a boolean if a field has been set.

### SetLastRunAtNil

`func (o *TablesValue) SetLastRunAtNil(b bool)`

 SetLastRunAtNil sets the value for LastRunAt to be an explicit nil

### UnsetLastRunAt
`func (o *TablesValue) UnsetLastRunAt()`

UnsetLastRunAt ensures that no value is present for LastRunAt, not even an explicit nil
### GetStatus

`func (o *TablesValue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TablesValue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TablesValue) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TablesValue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TablesValue) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TablesValue) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


