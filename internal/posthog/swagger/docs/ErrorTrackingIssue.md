# ErrorTrackingIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | Pointer to [**ErrorTrackingIssueAggregations**](ErrorTrackingIssueAggregations.md) |  | [optional] 
**Assignee** | Pointer to [**ErrorTrackingIssueAssignee**](ErrorTrackingIssueAssignee.md) |  | [optional] 
**Cohort** | Pointer to [**ErrorTrackingIssueCohort**](ErrorTrackingIssueCohort.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExternalIssues** | Pointer to [**[]ErrorTrackingExternalReference**](ErrorTrackingExternalReference.md) |  | [optional] 
**FirstEvent** | Pointer to [**FirstEvent**](FirstEvent.md) |  | [optional] 
**FirstSeen** | **time.Time** |  | 
**Function** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**LastEvent** | Pointer to [**LastEvent**](LastEvent.md) |  | [optional] 
**LastSeen** | **time.Time** |  | 
**Library** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Revenue** | Pointer to **NullableFloat32** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**Status** | [**Status**](Status.md) |  | 

## Methods

### NewErrorTrackingIssue

`func NewErrorTrackingIssue(firstSeen time.Time, id string, lastSeen time.Time, status Status, ) *ErrorTrackingIssue`

NewErrorTrackingIssue instantiates a new ErrorTrackingIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTrackingIssueWithDefaults

`func NewErrorTrackingIssueWithDefaults() *ErrorTrackingIssue`

NewErrorTrackingIssueWithDefaults instantiates a new ErrorTrackingIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *ErrorTrackingIssue) GetAggregations() ErrorTrackingIssueAggregations`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *ErrorTrackingIssue) GetAggregationsOk() (*ErrorTrackingIssueAggregations, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *ErrorTrackingIssue) SetAggregations(v ErrorTrackingIssueAggregations)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *ErrorTrackingIssue) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetAssignee

`func (o *ErrorTrackingIssue) GetAssignee() ErrorTrackingIssueAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ErrorTrackingIssue) GetAssigneeOk() (*ErrorTrackingIssueAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ErrorTrackingIssue) SetAssignee(v ErrorTrackingIssueAssignee)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *ErrorTrackingIssue) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCohort

`func (o *ErrorTrackingIssue) GetCohort() ErrorTrackingIssueCohort`

GetCohort returns the Cohort field if non-nil, zero value otherwise.

### GetCohortOk

`func (o *ErrorTrackingIssue) GetCohortOk() (*ErrorTrackingIssueCohort, bool)`

GetCohortOk returns a tuple with the Cohort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohort

`func (o *ErrorTrackingIssue) SetCohort(v ErrorTrackingIssueCohort)`

SetCohort sets Cohort field to given value.

### HasCohort

`func (o *ErrorTrackingIssue) HasCohort() bool`

HasCohort returns a boolean if a field has been set.

### GetDescription

`func (o *ErrorTrackingIssue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorTrackingIssue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorTrackingIssue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorTrackingIssue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ErrorTrackingIssue) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ErrorTrackingIssue) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalIssues

`func (o *ErrorTrackingIssue) GetExternalIssues() []ErrorTrackingExternalReference`

GetExternalIssues returns the ExternalIssues field if non-nil, zero value otherwise.

### GetExternalIssuesOk

`func (o *ErrorTrackingIssue) GetExternalIssuesOk() (*[]ErrorTrackingExternalReference, bool)`

GetExternalIssuesOk returns a tuple with the ExternalIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIssues

`func (o *ErrorTrackingIssue) SetExternalIssues(v []ErrorTrackingExternalReference)`

SetExternalIssues sets ExternalIssues field to given value.

### HasExternalIssues

`func (o *ErrorTrackingIssue) HasExternalIssues() bool`

HasExternalIssues returns a boolean if a field has been set.

### SetExternalIssuesNil

`func (o *ErrorTrackingIssue) SetExternalIssuesNil(b bool)`

 SetExternalIssuesNil sets the value for ExternalIssues to be an explicit nil

### UnsetExternalIssues
`func (o *ErrorTrackingIssue) UnsetExternalIssues()`

UnsetExternalIssues ensures that no value is present for ExternalIssues, not even an explicit nil
### GetFirstEvent

`func (o *ErrorTrackingIssue) GetFirstEvent() FirstEvent`

GetFirstEvent returns the FirstEvent field if non-nil, zero value otherwise.

### GetFirstEventOk

`func (o *ErrorTrackingIssue) GetFirstEventOk() (*FirstEvent, bool)`

GetFirstEventOk returns a tuple with the FirstEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstEvent

`func (o *ErrorTrackingIssue) SetFirstEvent(v FirstEvent)`

SetFirstEvent sets FirstEvent field to given value.

### HasFirstEvent

`func (o *ErrorTrackingIssue) HasFirstEvent() bool`

HasFirstEvent returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ErrorTrackingIssue) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ErrorTrackingIssue) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ErrorTrackingIssue) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.


### GetFunction

`func (o *ErrorTrackingIssue) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ErrorTrackingIssue) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ErrorTrackingIssue) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *ErrorTrackingIssue) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### SetFunctionNil

`func (o *ErrorTrackingIssue) SetFunctionNil(b bool)`

 SetFunctionNil sets the value for Function to be an explicit nil

### UnsetFunction
`func (o *ErrorTrackingIssue) UnsetFunction()`

UnsetFunction ensures that no value is present for Function, not even an explicit nil
### GetId

`func (o *ErrorTrackingIssue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorTrackingIssue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorTrackingIssue) SetId(v string)`

SetId sets Id field to given value.


### GetLastEvent

`func (o *ErrorTrackingIssue) GetLastEvent() LastEvent`

GetLastEvent returns the LastEvent field if non-nil, zero value otherwise.

### GetLastEventOk

`func (o *ErrorTrackingIssue) GetLastEventOk() (*LastEvent, bool)`

GetLastEventOk returns a tuple with the LastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvent

`func (o *ErrorTrackingIssue) SetLastEvent(v LastEvent)`

SetLastEvent sets LastEvent field to given value.

### HasLastEvent

`func (o *ErrorTrackingIssue) HasLastEvent() bool`

HasLastEvent returns a boolean if a field has been set.

### GetLastSeen

`func (o *ErrorTrackingIssue) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ErrorTrackingIssue) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ErrorTrackingIssue) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.


### GetLibrary

`func (o *ErrorTrackingIssue) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *ErrorTrackingIssue) GetLibraryOk() (*string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *ErrorTrackingIssue) SetLibrary(v string)`

SetLibrary sets Library field to given value.

### HasLibrary

`func (o *ErrorTrackingIssue) HasLibrary() bool`

HasLibrary returns a boolean if a field has been set.

### SetLibraryNil

`func (o *ErrorTrackingIssue) SetLibraryNil(b bool)`

 SetLibraryNil sets the value for Library to be an explicit nil

### UnsetLibrary
`func (o *ErrorTrackingIssue) UnsetLibrary()`

UnsetLibrary ensures that no value is present for Library, not even an explicit nil
### GetName

`func (o *ErrorTrackingIssue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ErrorTrackingIssue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ErrorTrackingIssue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ErrorTrackingIssue) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ErrorTrackingIssue) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ErrorTrackingIssue) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRevenue

`func (o *ErrorTrackingIssue) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *ErrorTrackingIssue) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *ErrorTrackingIssue) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *ErrorTrackingIssue) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### SetRevenueNil

`func (o *ErrorTrackingIssue) SetRevenueNil(b bool)`

 SetRevenueNil sets the value for Revenue to be an explicit nil

### UnsetRevenue
`func (o *ErrorTrackingIssue) UnsetRevenue()`

UnsetRevenue ensures that no value is present for Revenue, not even an explicit nil
### GetSource

`func (o *ErrorTrackingIssue) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ErrorTrackingIssue) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ErrorTrackingIssue) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ErrorTrackingIssue) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *ErrorTrackingIssue) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *ErrorTrackingIssue) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetStatus

`func (o *ErrorTrackingIssue) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorTrackingIssue) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorTrackingIssue) SetStatus(v Status)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


