# ErrorTrackingCorrelatedIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to [**ErrorTrackingIssueAssignee**](ErrorTrackingIssueAssignee.md) |  | [optional] 
**Cohort** | Pointer to [**ErrorTrackingIssueCohort**](ErrorTrackingIssueCohort.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Event** | **string** |  | 
**ExternalIssues** | Pointer to [**[]ErrorTrackingExternalReference**](ErrorTrackingExternalReference.md) |  | [optional] 
**FirstSeen** | **time.Time** |  | 
**Id** | **string** |  | 
**LastSeen** | **time.Time** |  | 
**Library** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OddsRatio** | **float32** |  | 
**Population** | [**Population**](Population.md) |  | 
**Status** | [**Status**](Status.md) |  | 

## Methods

### NewErrorTrackingCorrelatedIssue

`func NewErrorTrackingCorrelatedIssue(event string, firstSeen time.Time, id string, lastSeen time.Time, oddsRatio float32, population Population, status Status, ) *ErrorTrackingCorrelatedIssue`

NewErrorTrackingCorrelatedIssue instantiates a new ErrorTrackingCorrelatedIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTrackingCorrelatedIssueWithDefaults

`func NewErrorTrackingCorrelatedIssueWithDefaults() *ErrorTrackingCorrelatedIssue`

NewErrorTrackingCorrelatedIssueWithDefaults instantiates a new ErrorTrackingCorrelatedIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *ErrorTrackingCorrelatedIssue) GetAssignee() ErrorTrackingIssueAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ErrorTrackingCorrelatedIssue) GetAssigneeOk() (*ErrorTrackingIssueAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ErrorTrackingCorrelatedIssue) SetAssignee(v ErrorTrackingIssueAssignee)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *ErrorTrackingCorrelatedIssue) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCohort

`func (o *ErrorTrackingCorrelatedIssue) GetCohort() ErrorTrackingIssueCohort`

GetCohort returns the Cohort field if non-nil, zero value otherwise.

### GetCohortOk

`func (o *ErrorTrackingCorrelatedIssue) GetCohortOk() (*ErrorTrackingIssueCohort, bool)`

GetCohortOk returns a tuple with the Cohort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohort

`func (o *ErrorTrackingCorrelatedIssue) SetCohort(v ErrorTrackingIssueCohort)`

SetCohort sets Cohort field to given value.

### HasCohort

`func (o *ErrorTrackingCorrelatedIssue) HasCohort() bool`

HasCohort returns a boolean if a field has been set.

### GetDescription

`func (o *ErrorTrackingCorrelatedIssue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorTrackingCorrelatedIssue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorTrackingCorrelatedIssue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorTrackingCorrelatedIssue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ErrorTrackingCorrelatedIssue) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ErrorTrackingCorrelatedIssue) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEvent

`func (o *ErrorTrackingCorrelatedIssue) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ErrorTrackingCorrelatedIssue) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ErrorTrackingCorrelatedIssue) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetExternalIssues

`func (o *ErrorTrackingCorrelatedIssue) GetExternalIssues() []ErrorTrackingExternalReference`

GetExternalIssues returns the ExternalIssues field if non-nil, zero value otherwise.

### GetExternalIssuesOk

`func (o *ErrorTrackingCorrelatedIssue) GetExternalIssuesOk() (*[]ErrorTrackingExternalReference, bool)`

GetExternalIssuesOk returns a tuple with the ExternalIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIssues

`func (o *ErrorTrackingCorrelatedIssue) SetExternalIssues(v []ErrorTrackingExternalReference)`

SetExternalIssues sets ExternalIssues field to given value.

### HasExternalIssues

`func (o *ErrorTrackingCorrelatedIssue) HasExternalIssues() bool`

HasExternalIssues returns a boolean if a field has been set.

### SetExternalIssuesNil

`func (o *ErrorTrackingCorrelatedIssue) SetExternalIssuesNil(b bool)`

 SetExternalIssuesNil sets the value for ExternalIssues to be an explicit nil

### UnsetExternalIssues
`func (o *ErrorTrackingCorrelatedIssue) UnsetExternalIssues()`

UnsetExternalIssues ensures that no value is present for ExternalIssues, not even an explicit nil
### GetFirstSeen

`func (o *ErrorTrackingCorrelatedIssue) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ErrorTrackingCorrelatedIssue) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ErrorTrackingCorrelatedIssue) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.


### GetId

`func (o *ErrorTrackingCorrelatedIssue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorTrackingCorrelatedIssue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorTrackingCorrelatedIssue) SetId(v string)`

SetId sets Id field to given value.


### GetLastSeen

`func (o *ErrorTrackingCorrelatedIssue) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ErrorTrackingCorrelatedIssue) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ErrorTrackingCorrelatedIssue) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.


### GetLibrary

`func (o *ErrorTrackingCorrelatedIssue) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *ErrorTrackingCorrelatedIssue) GetLibraryOk() (*string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *ErrorTrackingCorrelatedIssue) SetLibrary(v string)`

SetLibrary sets Library field to given value.

### HasLibrary

`func (o *ErrorTrackingCorrelatedIssue) HasLibrary() bool`

HasLibrary returns a boolean if a field has been set.

### SetLibraryNil

`func (o *ErrorTrackingCorrelatedIssue) SetLibraryNil(b bool)`

 SetLibraryNil sets the value for Library to be an explicit nil

### UnsetLibrary
`func (o *ErrorTrackingCorrelatedIssue) UnsetLibrary()`

UnsetLibrary ensures that no value is present for Library, not even an explicit nil
### GetName

`func (o *ErrorTrackingCorrelatedIssue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ErrorTrackingCorrelatedIssue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ErrorTrackingCorrelatedIssue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ErrorTrackingCorrelatedIssue) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ErrorTrackingCorrelatedIssue) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ErrorTrackingCorrelatedIssue) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOddsRatio

`func (o *ErrorTrackingCorrelatedIssue) GetOddsRatio() float32`

GetOddsRatio returns the OddsRatio field if non-nil, zero value otherwise.

### GetOddsRatioOk

`func (o *ErrorTrackingCorrelatedIssue) GetOddsRatioOk() (*float32, bool)`

GetOddsRatioOk returns a tuple with the OddsRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOddsRatio

`func (o *ErrorTrackingCorrelatedIssue) SetOddsRatio(v float32)`

SetOddsRatio sets OddsRatio field to given value.


### GetPopulation

`func (o *ErrorTrackingCorrelatedIssue) GetPopulation() Population`

GetPopulation returns the Population field if non-nil, zero value otherwise.

### GetPopulationOk

`func (o *ErrorTrackingCorrelatedIssue) GetPopulationOk() (*Population, bool)`

GetPopulationOk returns a tuple with the Population field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulation

`func (o *ErrorTrackingCorrelatedIssue) SetPopulation(v Population)`

SetPopulation sets Population field to given value.


### GetStatus

`func (o *ErrorTrackingCorrelatedIssue) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorTrackingCorrelatedIssue) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorTrackingCorrelatedIssue) SetStatus(v Status)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


