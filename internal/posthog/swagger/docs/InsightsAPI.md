# \InsightsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InsightsActivityRetrieve**](InsightsAPI.md#InsightsActivityRetrieve) | **Get** /api/projects/{project_id}/insights/activity/ | 
[**InsightsActivityRetrieve2**](InsightsAPI.md#InsightsActivityRetrieve2) | **Get** /api/projects/{project_id}/insights/{id}/activity/ | 
[**InsightsCancelCreate**](InsightsAPI.md#InsightsCancelCreate) | **Post** /api/projects/{project_id}/insights/cancel/ | 
[**InsightsCreate**](InsightsAPI.md#InsightsCreate) | **Post** /api/projects/{project_id}/insights/ | 
[**InsightsDestroy**](InsightsAPI.md#InsightsDestroy) | **Delete** /api/projects/{project_id}/insights/{id}/ | 
[**InsightsList**](InsightsAPI.md#InsightsList) | **Get** /api/projects/{project_id}/insights/ | 
[**InsightsMyLastViewedRetrieve**](InsightsAPI.md#InsightsMyLastViewedRetrieve) | **Get** /api/projects/{project_id}/insights/my_last_viewed/ | 
[**InsightsPartialUpdate**](InsightsAPI.md#InsightsPartialUpdate) | **Patch** /api/projects/{project_id}/insights/{id}/ | 
[**InsightsRetrieve**](InsightsAPI.md#InsightsRetrieve) | **Get** /api/projects/{project_id}/insights/{id}/ | 
[**InsightsSharingList**](InsightsAPI.md#InsightsSharingList) | **Get** /api/projects/{project_id}/insights/{insight_id}/sharing/ | 
[**InsightsSharingPasswordsCreate**](InsightsAPI.md#InsightsSharingPasswordsCreate) | **Post** /api/projects/{project_id}/insights/{insight_id}/sharing/passwords/ | 
[**InsightsSharingPasswordsDestroy**](InsightsAPI.md#InsightsSharingPasswordsDestroy) | **Delete** /api/projects/{project_id}/insights/{insight_id}/sharing/passwords/{password_id}/ | 
[**InsightsSharingRefreshCreate**](InsightsAPI.md#InsightsSharingRefreshCreate) | **Post** /api/projects/{project_id}/insights/{insight_id}/sharing/refresh/ | 
[**InsightsUpdate**](InsightsAPI.md#InsightsUpdate) | **Put** /api/projects/{project_id}/insights/{id}/ | 
[**InsightsViewedCreate**](InsightsAPI.md#InsightsViewedCreate) | **Post** /api/projects/{project_id}/insights/viewed/ | 



## InsightsActivityRetrieve

> InsightsActivityRetrieve(ctx, projectId).Format(format).Execute()



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
	r, err := apiClient.InsightsAPI.InsightsActivityRetrieve(context.Background(), projectId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsActivityRetrieve``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInsightsActivityRetrieveRequest struct via the builder pattern


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


## InsightsActivityRetrieve2

> InsightsActivityRetrieve2(ctx, id, projectId).Format(format).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this insight.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InsightsAPI.InsightsActivityRetrieve2(context.Background(), id, projectId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsActivityRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this insight. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightsActivityRetrieve2Request struct via the builder pattern


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


## InsightsCancelCreate

> InsightsCancelCreate(ctx, projectId).Format(format).Insight(insight).Execute()



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
	format := "format_example" // string |  (optional)
	insight := *openapiclient.NewInsight(int32(123), "ShortId_example", []openapiclient.DashboardTileBasic{*openapiclient.NewDashboardTileBasic(int32(123), int32(123))}, "LastRefresh_example", "CacheTargetAge_example", "NextAllowedClientRefresh_example", "Result_example", "HasMore_example", "Columns_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), false, openapiclient.EffectiveRestrictionLevelEnum(21), openapiclient.EffectivePrivilegeLevelEnum(21), "UserAccessLevel_example", "Timezone_example", "IsCached_example", "QueryStatus_example", "Hogql_example", "Types_example", "Alerts_example", "LastViewedAt_example") // Insight |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InsightsAPI.InsightsCancelCreate(context.Background(), projectId).Format(format).Insight(insight).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsCancelCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInsightsCancelCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** |  | 
 **insight** | [**Insight**](Insight.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightsCreate

> Insight InsightsCreate(ctx, projectId).Format(format).Insight(insight).Execute()



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
	format := "format_example" // string |  (optional)
	insight := *openapiclient.NewInsight(int32(123), "ShortId_example", []openapiclient.DashboardTileBasic{*openapiclient.NewDashboardTileBasic(int32(123), int32(123))}, "LastRefresh_example", "CacheTargetAge_example", "NextAllowedClientRefresh_example", "Result_example", "HasMore_example", "Columns_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), false, openapiclient.EffectiveRestrictionLevelEnum(21), openapiclient.EffectivePrivilegeLevelEnum(21), "UserAccessLevel_example", "Timezone_example", "IsCached_example", "QueryStatus_example", "Hogql_example", "Types_example", "Alerts_example", "LastViewedAt_example") // Insight |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightsAPI.InsightsCreate(context.Background(), projectId).Format(format).Insight(insight).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsightsCreate`: Insight
	fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.InsightsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** |  | 
 **insight** | [**Insight**](Insight.md) |  | 

### Return type

[**Insight**](Insight.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightsDestroy

> InsightsDestroy(ctx, id, projectId).Format(format).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this insight.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InsightsAPI.InsightsDestroy(context.Background(), id, projectId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this insight. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightsDestroyRequest struct via the builder pattern


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


## InsightsList

> PaginatedInsightList InsightsList(ctx, projectId).Basic(basic).CreatedBy(createdBy).Format(format).Limit(limit).Offset(offset).Refresh(refresh).ShortId(shortId).Execute()



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
	basic := true // bool | Return basic insight metadata only (no results, faster). (optional)
	createdBy := int32(56) // int32 |  (optional)
	format := "format_example" // string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	refresh := "refresh_example" // string |  Whether to refresh the retrieved insights, how aggresively, and if sync or async: - `'force_cache'` - return cached data or a cache miss; always completes immediately as it never calculates - `'blocking'` - calculate synchronously (returning only when the query is done), UNLESS there are very fresh results in the cache - `'async'` - kick off background calculation (returning immediately with a query status), UNLESS there are very fresh results in the cache - `'lazy_async'` - kick off background calculation, UNLESS there are somewhat fresh results in the cache - `'force_blocking'` - calculate synchronously, even if fresh results are already cached - `'force_async'` - kick off background calculation, even if fresh results are already cached Background calculation can be tracked using the `query_status` response field. (optional) (default to "force_cache")
	shortId := "shortId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightsAPI.InsightsList(context.Background(), projectId).Basic(basic).CreatedBy(createdBy).Format(format).Limit(limit).Offset(offset).Refresh(refresh).ShortId(shortId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsightsList`: PaginatedInsightList
	fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.InsightsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **basic** | **bool** | Return basic insight metadata only (no results, faster). | 
 **createdBy** | **int32** |  | 
 **format** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **refresh** | **string** |  Whether to refresh the retrieved insights, how aggresively, and if sync or async: - &#x60;&#39;force_cache&#39;&#x60; - return cached data or a cache miss; always completes immediately as it never calculates - &#x60;&#39;blocking&#39;&#x60; - calculate synchronously (returning only when the query is done), UNLESS there are very fresh results in the cache - &#x60;&#39;async&#39;&#x60; - kick off background calculation (returning immediately with a query status), UNLESS there are very fresh results in the cache - &#x60;&#39;lazy_async&#39;&#x60; - kick off background calculation, UNLESS there are somewhat fresh results in the cache - &#x60;&#39;force_blocking&#39;&#x60; - calculate synchronously, even if fresh results are already cached - &#x60;&#39;force_async&#39;&#x60; - kick off background calculation, even if fresh results are already cached Background calculation can be tracked using the &#x60;query_status&#x60; response field. | [default to &quot;force_cache&quot;]
 **shortId** | **string** |  | 

### Return type

[**PaginatedInsightList**](PaginatedInsightList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightsMyLastViewedRetrieve

> InsightsMyLastViewedRetrieve(ctx, projectId).Format(format).Execute()





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
	r, err := apiClient.InsightsAPI.InsightsMyLastViewedRetrieve(context.Background(), projectId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsMyLastViewedRetrieve``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInsightsMyLastViewedRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightsPartialUpdate

> Insight InsightsPartialUpdate(ctx, id, projectId).Format(format).PatchedInsight(patchedInsight).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this insight.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)
	patchedInsight := *openapiclient.NewPatchedInsight() // PatchedInsight |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightsAPI.InsightsPartialUpdate(context.Background(), id, projectId).Format(format).PatchedInsight(patchedInsight).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsightsPartialUpdate`: Insight
	fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.InsightsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this insight. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 
 **patchedInsight** | [**PatchedInsight**](PatchedInsight.md) |  | 

### Return type

[**Insight**](Insight.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightsRetrieve

> Insight InsightsRetrieve(ctx, id, projectId).Format(format).FromDashboard(fromDashboard).Refresh(refresh).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this insight.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)
	fromDashboard := int32(56) // int32 |  Only if loading an insight in the context of a dashboard: The relevant dashboard's ID. When set, the specified dashboard's filters and date range override will be applied. (optional)
	refresh := "refresh_example" // string |  Whether to refresh the insight, how aggresively, and if sync or async: - `'force_cache'` - return cached data or a cache miss; always completes immediately as it never calculates - `'blocking'` - calculate synchronously (returning only when the query is done), UNLESS there are very fresh results in the cache - `'async'` - kick off background calculation (returning immediately with a query status), UNLESS there are very fresh results in the cache - `'lazy_async'` - kick off background calculation, UNLESS there are somewhat fresh results in the cache - `'force_blocking'` - calculate synchronously, even if fresh results are already cached - `'force_async'` - kick off background calculation, even if fresh results are already cached Background calculation can be tracked using the `query_status` response field. (optional) (default to "force_cache")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightsAPI.InsightsRetrieve(context.Background(), id, projectId).Format(format).FromDashboard(fromDashboard).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsightsRetrieve`: Insight
	fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.InsightsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this insight. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 
 **fromDashboard** | **int32** |  Only if loading an insight in the context of a dashboard: The relevant dashboard&#39;s ID. When set, the specified dashboard&#39;s filters and date range override will be applied. | 
 **refresh** | **string** |  Whether to refresh the insight, how aggresively, and if sync or async: - &#x60;&#39;force_cache&#39;&#x60; - return cached data or a cache miss; always completes immediately as it never calculates - &#x60;&#39;blocking&#39;&#x60; - calculate synchronously (returning only when the query is done), UNLESS there are very fresh results in the cache - &#x60;&#39;async&#39;&#x60; - kick off background calculation (returning immediately with a query status), UNLESS there are very fresh results in the cache - &#x60;&#39;lazy_async&#39;&#x60; - kick off background calculation, UNLESS there are somewhat fresh results in the cache - &#x60;&#39;force_blocking&#39;&#x60; - calculate synchronously, even if fresh results are already cached - &#x60;&#39;force_async&#39;&#x60; - kick off background calculation, even if fresh results are already cached Background calculation can be tracked using the &#x60;query_status&#x60; response field. | [default to &quot;force_cache&quot;]

### Return type

[**Insight**](Insight.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightsSharingList

> []SharingConfiguration InsightsSharingList(ctx, insightId, projectId).Execute()



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
	insightId := int32(56) // int32 | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightsAPI.InsightsSharingList(context.Background(), insightId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsSharingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsightsSharingList`: []SharingConfiguration
	fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.InsightsSharingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightId** | **int32** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightsSharingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]SharingConfiguration**](SharingConfiguration.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightsSharingPasswordsCreate

> SharingConfiguration InsightsSharingPasswordsCreate(ctx, insightId, projectId).SharingConfiguration(sharingConfiguration).Execute()





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
	insightId := int32(56) // int32 | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	sharingConfiguration := *openapiclient.NewSharingConfiguration(time.Now(), "AccessToken_example", "SharePasswords_example") // SharingConfiguration |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightsAPI.InsightsSharingPasswordsCreate(context.Background(), insightId, projectId).SharingConfiguration(sharingConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsSharingPasswordsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsightsSharingPasswordsCreate`: SharingConfiguration
	fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.InsightsSharingPasswordsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightId** | **int32** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightsSharingPasswordsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sharingConfiguration** | [**SharingConfiguration**](SharingConfiguration.md) |  | 

### Return type

[**SharingConfiguration**](SharingConfiguration.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightsSharingPasswordsDestroy

> InsightsSharingPasswordsDestroy(ctx, insightId, passwordId, projectId).Execute()





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
	insightId := int32(56) // int32 | 
	passwordId := "passwordId_example" // string | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InsightsAPI.InsightsSharingPasswordsDestroy(context.Background(), insightId, passwordId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsSharingPasswordsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightId** | **int32** |  | 
**passwordId** | **string** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightsSharingPasswordsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## InsightsSharingRefreshCreate

> SharingConfiguration InsightsSharingRefreshCreate(ctx, insightId, projectId).SharingConfiguration(sharingConfiguration).Execute()



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
	insightId := int32(56) // int32 | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	sharingConfiguration := *openapiclient.NewSharingConfiguration(time.Now(), "AccessToken_example", "SharePasswords_example") // SharingConfiguration |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightsAPI.InsightsSharingRefreshCreate(context.Background(), insightId, projectId).SharingConfiguration(sharingConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsSharingRefreshCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsightsSharingRefreshCreate`: SharingConfiguration
	fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.InsightsSharingRefreshCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightId** | **int32** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightsSharingRefreshCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sharingConfiguration** | [**SharingConfiguration**](SharingConfiguration.md) |  | 

### Return type

[**SharingConfiguration**](SharingConfiguration.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightsUpdate

> Insight InsightsUpdate(ctx, id, projectId).Format(format).Insight(insight).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this insight.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)
	insight := *openapiclient.NewInsight(int32(123), "ShortId_example", []openapiclient.DashboardTileBasic{*openapiclient.NewDashboardTileBasic(int32(123), int32(123))}, "LastRefresh_example", "CacheTargetAge_example", "NextAllowedClientRefresh_example", "Result_example", "HasMore_example", "Columns_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), false, openapiclient.EffectiveRestrictionLevelEnum(21), openapiclient.EffectivePrivilegeLevelEnum(21), "UserAccessLevel_example", "Timezone_example", "IsCached_example", "QueryStatus_example", "Hogql_example", "Types_example", "Alerts_example", "LastViewedAt_example") // Insight |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightsAPI.InsightsUpdate(context.Background(), id, projectId).Format(format).Insight(insight).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsightsUpdate`: Insight
	fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.InsightsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this insight. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsightsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 
 **insight** | [**Insight**](Insight.md) |  | 

### Return type

[**Insight**](Insight.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightsViewedCreate

> InsightsViewedCreate(ctx, projectId).Format(format).Insight(insight).Execute()





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
	format := "format_example" // string |  (optional)
	insight := *openapiclient.NewInsight(int32(123), "ShortId_example", []openapiclient.DashboardTileBasic{*openapiclient.NewDashboardTileBasic(int32(123), int32(123))}, "LastRefresh_example", "CacheTargetAge_example", "NextAllowedClientRefresh_example", "Result_example", "HasMore_example", "Columns_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), false, openapiclient.EffectiveRestrictionLevelEnum(21), openapiclient.EffectivePrivilegeLevelEnum(21), "UserAccessLevel_example", "Timezone_example", "IsCached_example", "QueryStatus_example", "Hogql_example", "Types_example", "Alerts_example", "LastViewedAt_example") // Insight |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InsightsAPI.InsightsViewedCreate(context.Background(), projectId).Format(format).Insight(insight).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.InsightsViewedCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInsightsViewedCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** |  | 
 **insight** | [**Insight**](Insight.md) |  | 

### Return type

 (empty response body)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

