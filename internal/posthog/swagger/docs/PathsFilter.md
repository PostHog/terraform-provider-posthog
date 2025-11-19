# PathsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeLimit** | Pointer to **NullableInt32** |  | [optional] [default to 50]
**EndPoint** | Pointer to **NullableString** |  | [optional] 
**ExcludeEvents** | Pointer to **[]string** |  | [optional] 
**IncludeEventTypes** | Pointer to [**[]PathType**](PathType.md) |  | [optional] 
**LocalPathCleaningFilters** | Pointer to [**[]PathCleaningFilter**](PathCleaningFilter.md) |  | [optional] 
**MaxEdgeWeight** | Pointer to **NullableInt32** |  | [optional] 
**MinEdgeWeight** | Pointer to **NullableInt32** |  | [optional] 
**PathDropoffKey** | Pointer to **NullableString** | Relevant only within actors query | [optional] 
**PathEndKey** | Pointer to **NullableString** | Relevant only within actors query | [optional] 
**PathGroupings** | Pointer to **[]string** |  | [optional] 
**PathReplacements** | Pointer to **NullableBool** |  | [optional] 
**PathStartKey** | Pointer to **NullableString** | Relevant only within actors query | [optional] 
**PathsHogQLExpression** | Pointer to **NullableString** |  | [optional] 
**ShowFullUrls** | Pointer to **NullableBool** |  | [optional] 
**StartPoint** | Pointer to **NullableString** |  | [optional] 
**StepLimit** | Pointer to **NullableInt32** |  | [optional] [default to 5]

## Methods

### NewPathsFilter

`func NewPathsFilter() *PathsFilter`

NewPathsFilter instantiates a new PathsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathsFilterWithDefaults

`func NewPathsFilterWithDefaults() *PathsFilter`

NewPathsFilterWithDefaults instantiates a new PathsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeLimit

`func (o *PathsFilter) GetEdgeLimit() int32`

GetEdgeLimit returns the EdgeLimit field if non-nil, zero value otherwise.

### GetEdgeLimitOk

`func (o *PathsFilter) GetEdgeLimitOk() (*int32, bool)`

GetEdgeLimitOk returns a tuple with the EdgeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeLimit

`func (o *PathsFilter) SetEdgeLimit(v int32)`

SetEdgeLimit sets EdgeLimit field to given value.

### HasEdgeLimit

`func (o *PathsFilter) HasEdgeLimit() bool`

HasEdgeLimit returns a boolean if a field has been set.

### SetEdgeLimitNil

`func (o *PathsFilter) SetEdgeLimitNil(b bool)`

 SetEdgeLimitNil sets the value for EdgeLimit to be an explicit nil

### UnsetEdgeLimit
`func (o *PathsFilter) UnsetEdgeLimit()`

UnsetEdgeLimit ensures that no value is present for EdgeLimit, not even an explicit nil
### GetEndPoint

`func (o *PathsFilter) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *PathsFilter) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *PathsFilter) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.

### HasEndPoint

`func (o *PathsFilter) HasEndPoint() bool`

HasEndPoint returns a boolean if a field has been set.

### SetEndPointNil

`func (o *PathsFilter) SetEndPointNil(b bool)`

 SetEndPointNil sets the value for EndPoint to be an explicit nil

### UnsetEndPoint
`func (o *PathsFilter) UnsetEndPoint()`

UnsetEndPoint ensures that no value is present for EndPoint, not even an explicit nil
### GetExcludeEvents

`func (o *PathsFilter) GetExcludeEvents() []string`

GetExcludeEvents returns the ExcludeEvents field if non-nil, zero value otherwise.

### GetExcludeEventsOk

`func (o *PathsFilter) GetExcludeEventsOk() (*[]string, bool)`

GetExcludeEventsOk returns a tuple with the ExcludeEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeEvents

`func (o *PathsFilter) SetExcludeEvents(v []string)`

SetExcludeEvents sets ExcludeEvents field to given value.

### HasExcludeEvents

`func (o *PathsFilter) HasExcludeEvents() bool`

HasExcludeEvents returns a boolean if a field has been set.

### SetExcludeEventsNil

`func (o *PathsFilter) SetExcludeEventsNil(b bool)`

 SetExcludeEventsNil sets the value for ExcludeEvents to be an explicit nil

### UnsetExcludeEvents
`func (o *PathsFilter) UnsetExcludeEvents()`

UnsetExcludeEvents ensures that no value is present for ExcludeEvents, not even an explicit nil
### GetIncludeEventTypes

`func (o *PathsFilter) GetIncludeEventTypes() []PathType`

GetIncludeEventTypes returns the IncludeEventTypes field if non-nil, zero value otherwise.

### GetIncludeEventTypesOk

`func (o *PathsFilter) GetIncludeEventTypesOk() (*[]PathType, bool)`

GetIncludeEventTypesOk returns a tuple with the IncludeEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEventTypes

`func (o *PathsFilter) SetIncludeEventTypes(v []PathType)`

SetIncludeEventTypes sets IncludeEventTypes field to given value.

### HasIncludeEventTypes

`func (o *PathsFilter) HasIncludeEventTypes() bool`

HasIncludeEventTypes returns a boolean if a field has been set.

### SetIncludeEventTypesNil

`func (o *PathsFilter) SetIncludeEventTypesNil(b bool)`

 SetIncludeEventTypesNil sets the value for IncludeEventTypes to be an explicit nil

### UnsetIncludeEventTypes
`func (o *PathsFilter) UnsetIncludeEventTypes()`

UnsetIncludeEventTypes ensures that no value is present for IncludeEventTypes, not even an explicit nil
### GetLocalPathCleaningFilters

`func (o *PathsFilter) GetLocalPathCleaningFilters() []PathCleaningFilter`

GetLocalPathCleaningFilters returns the LocalPathCleaningFilters field if non-nil, zero value otherwise.

### GetLocalPathCleaningFiltersOk

`func (o *PathsFilter) GetLocalPathCleaningFiltersOk() (*[]PathCleaningFilter, bool)`

GetLocalPathCleaningFiltersOk returns a tuple with the LocalPathCleaningFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPathCleaningFilters

`func (o *PathsFilter) SetLocalPathCleaningFilters(v []PathCleaningFilter)`

SetLocalPathCleaningFilters sets LocalPathCleaningFilters field to given value.

### HasLocalPathCleaningFilters

`func (o *PathsFilter) HasLocalPathCleaningFilters() bool`

HasLocalPathCleaningFilters returns a boolean if a field has been set.

### SetLocalPathCleaningFiltersNil

`func (o *PathsFilter) SetLocalPathCleaningFiltersNil(b bool)`

 SetLocalPathCleaningFiltersNil sets the value for LocalPathCleaningFilters to be an explicit nil

### UnsetLocalPathCleaningFilters
`func (o *PathsFilter) UnsetLocalPathCleaningFilters()`

UnsetLocalPathCleaningFilters ensures that no value is present for LocalPathCleaningFilters, not even an explicit nil
### GetMaxEdgeWeight

`func (o *PathsFilter) GetMaxEdgeWeight() int32`

GetMaxEdgeWeight returns the MaxEdgeWeight field if non-nil, zero value otherwise.

### GetMaxEdgeWeightOk

`func (o *PathsFilter) GetMaxEdgeWeightOk() (*int32, bool)`

GetMaxEdgeWeightOk returns a tuple with the MaxEdgeWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEdgeWeight

`func (o *PathsFilter) SetMaxEdgeWeight(v int32)`

SetMaxEdgeWeight sets MaxEdgeWeight field to given value.

### HasMaxEdgeWeight

`func (o *PathsFilter) HasMaxEdgeWeight() bool`

HasMaxEdgeWeight returns a boolean if a field has been set.

### SetMaxEdgeWeightNil

`func (o *PathsFilter) SetMaxEdgeWeightNil(b bool)`

 SetMaxEdgeWeightNil sets the value for MaxEdgeWeight to be an explicit nil

### UnsetMaxEdgeWeight
`func (o *PathsFilter) UnsetMaxEdgeWeight()`

UnsetMaxEdgeWeight ensures that no value is present for MaxEdgeWeight, not even an explicit nil
### GetMinEdgeWeight

`func (o *PathsFilter) GetMinEdgeWeight() int32`

GetMinEdgeWeight returns the MinEdgeWeight field if non-nil, zero value otherwise.

### GetMinEdgeWeightOk

`func (o *PathsFilter) GetMinEdgeWeightOk() (*int32, bool)`

GetMinEdgeWeightOk returns a tuple with the MinEdgeWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEdgeWeight

`func (o *PathsFilter) SetMinEdgeWeight(v int32)`

SetMinEdgeWeight sets MinEdgeWeight field to given value.

### HasMinEdgeWeight

`func (o *PathsFilter) HasMinEdgeWeight() bool`

HasMinEdgeWeight returns a boolean if a field has been set.

### SetMinEdgeWeightNil

`func (o *PathsFilter) SetMinEdgeWeightNil(b bool)`

 SetMinEdgeWeightNil sets the value for MinEdgeWeight to be an explicit nil

### UnsetMinEdgeWeight
`func (o *PathsFilter) UnsetMinEdgeWeight()`

UnsetMinEdgeWeight ensures that no value is present for MinEdgeWeight, not even an explicit nil
### GetPathDropoffKey

`func (o *PathsFilter) GetPathDropoffKey() string`

GetPathDropoffKey returns the PathDropoffKey field if non-nil, zero value otherwise.

### GetPathDropoffKeyOk

`func (o *PathsFilter) GetPathDropoffKeyOk() (*string, bool)`

GetPathDropoffKeyOk returns a tuple with the PathDropoffKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathDropoffKey

`func (o *PathsFilter) SetPathDropoffKey(v string)`

SetPathDropoffKey sets PathDropoffKey field to given value.

### HasPathDropoffKey

`func (o *PathsFilter) HasPathDropoffKey() bool`

HasPathDropoffKey returns a boolean if a field has been set.

### SetPathDropoffKeyNil

`func (o *PathsFilter) SetPathDropoffKeyNil(b bool)`

 SetPathDropoffKeyNil sets the value for PathDropoffKey to be an explicit nil

### UnsetPathDropoffKey
`func (o *PathsFilter) UnsetPathDropoffKey()`

UnsetPathDropoffKey ensures that no value is present for PathDropoffKey, not even an explicit nil
### GetPathEndKey

`func (o *PathsFilter) GetPathEndKey() string`

GetPathEndKey returns the PathEndKey field if non-nil, zero value otherwise.

### GetPathEndKeyOk

`func (o *PathsFilter) GetPathEndKeyOk() (*string, bool)`

GetPathEndKeyOk returns a tuple with the PathEndKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathEndKey

`func (o *PathsFilter) SetPathEndKey(v string)`

SetPathEndKey sets PathEndKey field to given value.

### HasPathEndKey

`func (o *PathsFilter) HasPathEndKey() bool`

HasPathEndKey returns a boolean if a field has been set.

### SetPathEndKeyNil

`func (o *PathsFilter) SetPathEndKeyNil(b bool)`

 SetPathEndKeyNil sets the value for PathEndKey to be an explicit nil

### UnsetPathEndKey
`func (o *PathsFilter) UnsetPathEndKey()`

UnsetPathEndKey ensures that no value is present for PathEndKey, not even an explicit nil
### GetPathGroupings

`func (o *PathsFilter) GetPathGroupings() []string`

GetPathGroupings returns the PathGroupings field if non-nil, zero value otherwise.

### GetPathGroupingsOk

`func (o *PathsFilter) GetPathGroupingsOk() (*[]string, bool)`

GetPathGroupingsOk returns a tuple with the PathGroupings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathGroupings

`func (o *PathsFilter) SetPathGroupings(v []string)`

SetPathGroupings sets PathGroupings field to given value.

### HasPathGroupings

`func (o *PathsFilter) HasPathGroupings() bool`

HasPathGroupings returns a boolean if a field has been set.

### SetPathGroupingsNil

`func (o *PathsFilter) SetPathGroupingsNil(b bool)`

 SetPathGroupingsNil sets the value for PathGroupings to be an explicit nil

### UnsetPathGroupings
`func (o *PathsFilter) UnsetPathGroupings()`

UnsetPathGroupings ensures that no value is present for PathGroupings, not even an explicit nil
### GetPathReplacements

`func (o *PathsFilter) GetPathReplacements() bool`

GetPathReplacements returns the PathReplacements field if non-nil, zero value otherwise.

### GetPathReplacementsOk

`func (o *PathsFilter) GetPathReplacementsOk() (*bool, bool)`

GetPathReplacementsOk returns a tuple with the PathReplacements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathReplacements

`func (o *PathsFilter) SetPathReplacements(v bool)`

SetPathReplacements sets PathReplacements field to given value.

### HasPathReplacements

`func (o *PathsFilter) HasPathReplacements() bool`

HasPathReplacements returns a boolean if a field has been set.

### SetPathReplacementsNil

`func (o *PathsFilter) SetPathReplacementsNil(b bool)`

 SetPathReplacementsNil sets the value for PathReplacements to be an explicit nil

### UnsetPathReplacements
`func (o *PathsFilter) UnsetPathReplacements()`

UnsetPathReplacements ensures that no value is present for PathReplacements, not even an explicit nil
### GetPathStartKey

`func (o *PathsFilter) GetPathStartKey() string`

GetPathStartKey returns the PathStartKey field if non-nil, zero value otherwise.

### GetPathStartKeyOk

`func (o *PathsFilter) GetPathStartKeyOk() (*string, bool)`

GetPathStartKeyOk returns a tuple with the PathStartKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathStartKey

`func (o *PathsFilter) SetPathStartKey(v string)`

SetPathStartKey sets PathStartKey field to given value.

### HasPathStartKey

`func (o *PathsFilter) HasPathStartKey() bool`

HasPathStartKey returns a boolean if a field has been set.

### SetPathStartKeyNil

`func (o *PathsFilter) SetPathStartKeyNil(b bool)`

 SetPathStartKeyNil sets the value for PathStartKey to be an explicit nil

### UnsetPathStartKey
`func (o *PathsFilter) UnsetPathStartKey()`

UnsetPathStartKey ensures that no value is present for PathStartKey, not even an explicit nil
### GetPathsHogQLExpression

`func (o *PathsFilter) GetPathsHogQLExpression() string`

GetPathsHogQLExpression returns the PathsHogQLExpression field if non-nil, zero value otherwise.

### GetPathsHogQLExpressionOk

`func (o *PathsFilter) GetPathsHogQLExpressionOk() (*string, bool)`

GetPathsHogQLExpressionOk returns a tuple with the PathsHogQLExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathsHogQLExpression

`func (o *PathsFilter) SetPathsHogQLExpression(v string)`

SetPathsHogQLExpression sets PathsHogQLExpression field to given value.

### HasPathsHogQLExpression

`func (o *PathsFilter) HasPathsHogQLExpression() bool`

HasPathsHogQLExpression returns a boolean if a field has been set.

### SetPathsHogQLExpressionNil

`func (o *PathsFilter) SetPathsHogQLExpressionNil(b bool)`

 SetPathsHogQLExpressionNil sets the value for PathsHogQLExpression to be an explicit nil

### UnsetPathsHogQLExpression
`func (o *PathsFilter) UnsetPathsHogQLExpression()`

UnsetPathsHogQLExpression ensures that no value is present for PathsHogQLExpression, not even an explicit nil
### GetShowFullUrls

`func (o *PathsFilter) GetShowFullUrls() bool`

GetShowFullUrls returns the ShowFullUrls field if non-nil, zero value otherwise.

### GetShowFullUrlsOk

`func (o *PathsFilter) GetShowFullUrlsOk() (*bool, bool)`

GetShowFullUrlsOk returns a tuple with the ShowFullUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFullUrls

`func (o *PathsFilter) SetShowFullUrls(v bool)`

SetShowFullUrls sets ShowFullUrls field to given value.

### HasShowFullUrls

`func (o *PathsFilter) HasShowFullUrls() bool`

HasShowFullUrls returns a boolean if a field has been set.

### SetShowFullUrlsNil

`func (o *PathsFilter) SetShowFullUrlsNil(b bool)`

 SetShowFullUrlsNil sets the value for ShowFullUrls to be an explicit nil

### UnsetShowFullUrls
`func (o *PathsFilter) UnsetShowFullUrls()`

UnsetShowFullUrls ensures that no value is present for ShowFullUrls, not even an explicit nil
### GetStartPoint

`func (o *PathsFilter) GetStartPoint() string`

GetStartPoint returns the StartPoint field if non-nil, zero value otherwise.

### GetStartPointOk

`func (o *PathsFilter) GetStartPointOk() (*string, bool)`

GetStartPointOk returns a tuple with the StartPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPoint

`func (o *PathsFilter) SetStartPoint(v string)`

SetStartPoint sets StartPoint field to given value.

### HasStartPoint

`func (o *PathsFilter) HasStartPoint() bool`

HasStartPoint returns a boolean if a field has been set.

### SetStartPointNil

`func (o *PathsFilter) SetStartPointNil(b bool)`

 SetStartPointNil sets the value for StartPoint to be an explicit nil

### UnsetStartPoint
`func (o *PathsFilter) UnsetStartPoint()`

UnsetStartPoint ensures that no value is present for StartPoint, not even an explicit nil
### GetStepLimit

`func (o *PathsFilter) GetStepLimit() int32`

GetStepLimit returns the StepLimit field if non-nil, zero value otherwise.

### GetStepLimitOk

`func (o *PathsFilter) GetStepLimitOk() (*int32, bool)`

GetStepLimitOk returns a tuple with the StepLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepLimit

`func (o *PathsFilter) SetStepLimit(v int32)`

SetStepLimit sets StepLimit field to given value.

### HasStepLimit

`func (o *PathsFilter) HasStepLimit() bool`

HasStepLimit returns a boolean if a field has been set.

### SetStepLimitNil

`func (o *PathsFilter) SetStepLimitNil(b bool)`

 SetStepLimitNil sets the value for StepLimit to be an explicit nil

### UnsetStepLimit
`func (o *PathsFilter) UnsetStepLimit()`

UnsetStepLimit ensures that no value is present for StepLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


