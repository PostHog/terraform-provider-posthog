# StructuredSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Concise title (no longer than 10 words) summarizing the trace/event | 
**FlowDiagram** | **string** | Mermaid flowchart code showing the main flow | 
**SummaryBullets** | [**[]SummaryBullet**](SummaryBullet.md) | Main summary bullets | 
**InterestingNotes** | [**[]InterestingNote**](InterestingNote.md) | Interesting notes (0-2 for minimal, more for detailed) | 

## Methods

### NewStructuredSummary

`func NewStructuredSummary(title string, flowDiagram string, summaryBullets []SummaryBullet, interestingNotes []InterestingNote, ) *StructuredSummary`

NewStructuredSummary instantiates a new StructuredSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuredSummaryWithDefaults

`func NewStructuredSummaryWithDefaults() *StructuredSummary`

NewStructuredSummaryWithDefaults instantiates a new StructuredSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *StructuredSummary) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StructuredSummary) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StructuredSummary) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetFlowDiagram

`func (o *StructuredSummary) GetFlowDiagram() string`

GetFlowDiagram returns the FlowDiagram field if non-nil, zero value otherwise.

### GetFlowDiagramOk

`func (o *StructuredSummary) GetFlowDiagramOk() (*string, bool)`

GetFlowDiagramOk returns a tuple with the FlowDiagram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDiagram

`func (o *StructuredSummary) SetFlowDiagram(v string)`

SetFlowDiagram sets FlowDiagram field to given value.


### GetSummaryBullets

`func (o *StructuredSummary) GetSummaryBullets() []SummaryBullet`

GetSummaryBullets returns the SummaryBullets field if non-nil, zero value otherwise.

### GetSummaryBulletsOk

`func (o *StructuredSummary) GetSummaryBulletsOk() (*[]SummaryBullet, bool)`

GetSummaryBulletsOk returns a tuple with the SummaryBullets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryBullets

`func (o *StructuredSummary) SetSummaryBullets(v []SummaryBullet)`

SetSummaryBullets sets SummaryBullets field to given value.


### GetInterestingNotes

`func (o *StructuredSummary) GetInterestingNotes() []InterestingNote`

GetInterestingNotes returns the InterestingNotes field if non-nil, zero value otherwise.

### GetInterestingNotesOk

`func (o *StructuredSummary) GetInterestingNotesOk() (*[]InterestingNote, bool)`

GetInterestingNotesOk returns a tuple with the InterestingNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestingNotes

`func (o *StructuredSummary) SetInterestingNotes(v []InterestingNote)`

SetInterestingNotes sets InterestingNotes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


