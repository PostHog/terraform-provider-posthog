# \WebAnalyticsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebAnalyticsBreakdownRetrieve**](WebAnalyticsAPI.md#WebAnalyticsBreakdownRetrieve) | **Get** /api/projects/{project_id}/web_analytics/breakdown/ | 
[**WebAnalyticsOverviewRetrieve**](WebAnalyticsAPI.md#WebAnalyticsOverviewRetrieve) | **Get** /api/projects/{project_id}/web_analytics/overview/ | 



## WebAnalyticsBreakdownRetrieve

> WebAnalyticsBreakdownResponse WebAnalyticsBreakdownRetrieve(ctx, projectId).BreakdownBy(breakdownBy).DateFrom(dateFrom).DateTo(dateTo).ApplyPathCleaning(applyPathCleaning).FilterTestAccounts(filterTestAccounts).Host(host).Limit(limit).Offset(offset).Execute()





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
	breakdownBy := "breakdownBy_example" // string | Property to break down by  * `DeviceType` - DeviceType * `Browser` - Browser * `OS` - OS * `Viewport` - Viewport * `InitialReferringDomain` - InitialReferringDomain * `InitialUTMSource` - InitialUTMSource * `InitialUTMMedium` - InitialUTMMedium * `InitialUTMCampaign` - InitialUTMCampaign * `InitialUTMTerm` - InitialUTMTerm * `InitialUTMContent` - InitialUTMContent * `Country` - Country * `Region` - Region * `City` - City * `InitialPage` - InitialPage * `Page` - Page * `ExitPage` - ExitPage * `InitialChannelType` - InitialChannelType
	dateFrom := time.Now() // string | Start date for the query (format: YYYY-MM-DD)
	dateTo := time.Now() // string | End date for the query (format: YYYY-MM-DD)
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	applyPathCleaning := true // bool | Apply URL path cleaning (optional) (default to true)
	filterTestAccounts := true // bool | Filter out test accounts (optional) (default to true)
	host := "host_example" // string | Host to filter by (e.g. example.com) (optional)
	limit := int32(56) // int32 | Number of results to return (optional) (default to 100)
	offset := int32(56) // int32 | Number of results to skip (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebAnalyticsAPI.WebAnalyticsBreakdownRetrieve(context.Background(), projectId).BreakdownBy(breakdownBy).DateFrom(dateFrom).DateTo(dateTo).ApplyPathCleaning(applyPathCleaning).FilterTestAccounts(filterTestAccounts).Host(host).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebAnalyticsAPI.WebAnalyticsBreakdownRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebAnalyticsBreakdownRetrieve`: WebAnalyticsBreakdownResponse
	fmt.Fprintf(os.Stdout, "Response from `WebAnalyticsAPI.WebAnalyticsBreakdownRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebAnalyticsBreakdownRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **breakdownBy** | **string** | Property to break down by  * &#x60;DeviceType&#x60; - DeviceType * &#x60;Browser&#x60; - Browser * &#x60;OS&#x60; - OS * &#x60;Viewport&#x60; - Viewport * &#x60;InitialReferringDomain&#x60; - InitialReferringDomain * &#x60;InitialUTMSource&#x60; - InitialUTMSource * &#x60;InitialUTMMedium&#x60; - InitialUTMMedium * &#x60;InitialUTMCampaign&#x60; - InitialUTMCampaign * &#x60;InitialUTMTerm&#x60; - InitialUTMTerm * &#x60;InitialUTMContent&#x60; - InitialUTMContent * &#x60;Country&#x60; - Country * &#x60;Region&#x60; - Region * &#x60;City&#x60; - City * &#x60;InitialPage&#x60; - InitialPage * &#x60;Page&#x60; - Page * &#x60;ExitPage&#x60; - ExitPage * &#x60;InitialChannelType&#x60; - InitialChannelType | 
 **dateFrom** | **string** | Start date for the query (format: YYYY-MM-DD) | 
 **dateTo** | **string** | End date for the query (format: YYYY-MM-DD) | 

 **applyPathCleaning** | **bool** | Apply URL path cleaning | [default to true]
 **filterTestAccounts** | **bool** | Filter out test accounts | [default to true]
 **host** | **string** | Host to filter by (e.g. example.com) | 
 **limit** | **int32** | Number of results to return | [default to 100]
 **offset** | **int32** | Number of results to skip | [default to 0]

### Return type

[**WebAnalyticsBreakdownResponse**](WebAnalyticsBreakdownResponse.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebAnalyticsOverviewRetrieve

> WebAnalyticsOverviewResponse WebAnalyticsOverviewRetrieve(ctx, projectId).DateFrom(dateFrom).DateTo(dateTo).FilterTestAccounts(filterTestAccounts).Host(host).Execute()





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
	dateFrom := time.Now() // string | Start date for the query (format: YYYY-MM-DD)
	dateTo := time.Now() // string | End date for the query (format: YYYY-MM-DD)
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	filterTestAccounts := true // bool | Filter out test accounts (optional) (default to true)
	host := "host_example" // string | Host to filter by (e.g. example.com) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebAnalyticsAPI.WebAnalyticsOverviewRetrieve(context.Background(), projectId).DateFrom(dateFrom).DateTo(dateTo).FilterTestAccounts(filterTestAccounts).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebAnalyticsAPI.WebAnalyticsOverviewRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebAnalyticsOverviewRetrieve`: WebAnalyticsOverviewResponse
	fmt.Fprintf(os.Stdout, "Response from `WebAnalyticsAPI.WebAnalyticsOverviewRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebAnalyticsOverviewRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **string** | Start date for the query (format: YYYY-MM-DD) | 
 **dateTo** | **string** | End date for the query (format: YYYY-MM-DD) | 

 **filterTestAccounts** | **bool** | Filter out test accounts | [default to true]
 **host** | **string** | Host to filter by (e.g. example.com) | 

### Return type

[**WebAnalyticsOverviewResponse**](WebAnalyticsOverviewResponse.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

