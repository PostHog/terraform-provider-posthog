# \EventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsList**](EventsAPI.md#EventsList) | **Get** /api/projects/{project_id}/events/ | 
[**EventsRetrieve**](EventsAPI.md#EventsRetrieve) | **Get** /api/projects/{project_id}/events/{id}/ | 
[**EventsValuesRetrieve**](EventsAPI.md#EventsValuesRetrieve) | **Get** /api/projects/{project_id}/events/values/ | 



## EventsList

> PaginatedClickhouseEventList EventsList(ctx, projectId).After(after).Before(before).DistinctId(distinctId).Event(event).Format(format).Limit(limit).Offset(offset).PersonId(personId).Properties(properties).Select_(select_).Where(where).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/posthog/terraform-provider"
)

func main() {
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	after := time.Now() // time.Time | Only return events with a timestamp after this time. (optional)
	before := time.Now() // time.Time | Only return events with a timestamp before this time. (optional)
	distinctId := int32(56) // int32 | Filter list by distinct id. (optional)
	event := "event_example" // string | Filter list by event. For example `user sign up` or `$pageview`. (optional)
	format := "format_example" // string |  (optional)
	limit := int32(56) // int32 | The maximum number of results to return (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	personId := int32(56) // int32 | Filter list by person id. (optional)
	properties := []openapiclient.Property{*openapiclient.NewProperty([]openapiclient.PropertyItem{*openapiclient.NewPropertyItem("Key_example", "Value_example")})} // []Property | Filter events by event property, person property, cohort, groups and more. (optional)
	select_ := []string{"Inner_example"} // []string | (Experimental) JSON-serialized array of HogQL expressions to return (optional)
	where := []string{"Inner_example"} // []string | (Experimental) JSON-serialized array of HogQL expressions that must pass (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsList(context.Background(), projectId).After(after).Before(before).DistinctId(distinctId).Event(event).Format(format).Limit(limit).Offset(offset).PersonId(personId).Properties(properties).Select_(select_).Where(where).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsList`: PaginatedClickhouseEventList
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **time.Time** | Only return events with a timestamp after this time. | 
 **before** | **time.Time** | Only return events with a timestamp before this time. | 
 **distinctId** | **int32** | Filter list by distinct id. | 
 **event** | **string** | Filter list by event. For example &#x60;user sign up&#x60; or &#x60;$pageview&#x60;. | 
 **format** | **string** |  | 
 **limit** | **int32** | The maximum number of results to return | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **personId** | **int32** | Filter list by person id. | 
 **properties** | [**[]Property**](Property.md) | Filter events by event property, person property, cohort, groups and more. | 
 **select_** | **[]string** | (Experimental) JSON-serialized array of HogQL expressions to return | 
 **where** | **[]string** | (Experimental) JSON-serialized array of HogQL expressions that must pass | 

### Return type

[**PaginatedClickhouseEventList**](PaginatedClickhouseEventList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsRetrieve

> ClickhouseEvent EventsRetrieve(ctx, id, projectId).Format(format).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/posthog/terraform-provider"
)

func main() {
	id := "id_example" // string | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsRetrieve(context.Background(), id, projectId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsRetrieve`: ClickhouseEvent
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 

### Return type

[**ClickhouseEvent**](ClickhouseEvent.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsValuesRetrieve

> EventsValuesRetrieve(ctx, projectId).Format(format).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/posthog/terraform-provider"
)

func main() {
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.EventsValuesRetrieve(context.Background(), projectId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsValuesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsValuesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

