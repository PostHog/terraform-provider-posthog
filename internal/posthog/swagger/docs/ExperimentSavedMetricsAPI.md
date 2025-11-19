# \ExperimentSavedMetricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExperimentSavedMetricsCreate**](ExperimentSavedMetricsAPI.md#ExperimentSavedMetricsCreate) | **Post** /api/projects/{project_id}/experiment_saved_metrics/ | 
[**ExperimentSavedMetricsDestroy**](ExperimentSavedMetricsAPI.md#ExperimentSavedMetricsDestroy) | **Delete** /api/projects/{project_id}/experiment_saved_metrics/{id}/ | 
[**ExperimentSavedMetricsList**](ExperimentSavedMetricsAPI.md#ExperimentSavedMetricsList) | **Get** /api/projects/{project_id}/experiment_saved_metrics/ | 
[**ExperimentSavedMetricsPartialUpdate**](ExperimentSavedMetricsAPI.md#ExperimentSavedMetricsPartialUpdate) | **Patch** /api/projects/{project_id}/experiment_saved_metrics/{id}/ | 
[**ExperimentSavedMetricsRetrieve**](ExperimentSavedMetricsAPI.md#ExperimentSavedMetricsRetrieve) | **Get** /api/projects/{project_id}/experiment_saved_metrics/{id}/ | 
[**ExperimentSavedMetricsUpdate**](ExperimentSavedMetricsAPI.md#ExperimentSavedMetricsUpdate) | **Put** /api/projects/{project_id}/experiment_saved_metrics/{id}/ | 



## ExperimentSavedMetricsCreate

> ExperimentSavedMetric ExperimentSavedMetricsCreate(ctx, projectId).ExperimentSavedMetric(experimentSavedMetric).Execute()



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
	experimentSavedMetric := *openapiclient.NewExperimentSavedMetric(int32(123), "Name_example", interface{}(123), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now()) // ExperimentSavedMetric | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentSavedMetricsAPI.ExperimentSavedMetricsCreate(context.Background(), projectId).ExperimentSavedMetric(experimentSavedMetric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentSavedMetricsAPI.ExperimentSavedMetricsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentSavedMetricsCreate`: ExperimentSavedMetric
	fmt.Fprintf(os.Stdout, "Response from `ExperimentSavedMetricsAPI.ExperimentSavedMetricsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentSavedMetricsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experimentSavedMetric** | [**ExperimentSavedMetric**](ExperimentSavedMetric.md) |  | 

### Return type

[**ExperimentSavedMetric**](ExperimentSavedMetric.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExperimentSavedMetricsDestroy

> ExperimentSavedMetricsDestroy(ctx, id, projectId).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this experiment saved metric.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExperimentSavedMetricsAPI.ExperimentSavedMetricsDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentSavedMetricsAPI.ExperimentSavedMetricsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment saved metric. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentSavedMetricsDestroyRequest struct via the builder pattern


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


## ExperimentSavedMetricsList

> PaginatedExperimentSavedMetricList ExperimentSavedMetricsList(ctx, projectId).Limit(limit).Offset(offset).Execute()



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
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentSavedMetricsAPI.ExperimentSavedMetricsList(context.Background(), projectId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentSavedMetricsAPI.ExperimentSavedMetricsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentSavedMetricsList`: PaginatedExperimentSavedMetricList
	fmt.Fprintf(os.Stdout, "Response from `ExperimentSavedMetricsAPI.ExperimentSavedMetricsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentSavedMetricsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedExperimentSavedMetricList**](PaginatedExperimentSavedMetricList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExperimentSavedMetricsPartialUpdate

> ExperimentSavedMetric ExperimentSavedMetricsPartialUpdate(ctx, id, projectId).PatchedExperimentSavedMetric(patchedExperimentSavedMetric).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this experiment saved metric.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedExperimentSavedMetric := *openapiclient.NewPatchedExperimentSavedMetric() // PatchedExperimentSavedMetric |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentSavedMetricsAPI.ExperimentSavedMetricsPartialUpdate(context.Background(), id, projectId).PatchedExperimentSavedMetric(patchedExperimentSavedMetric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentSavedMetricsAPI.ExperimentSavedMetricsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentSavedMetricsPartialUpdate`: ExperimentSavedMetric
	fmt.Fprintf(os.Stdout, "Response from `ExperimentSavedMetricsAPI.ExperimentSavedMetricsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment saved metric. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentSavedMetricsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedExperimentSavedMetric** | [**PatchedExperimentSavedMetric**](PatchedExperimentSavedMetric.md) |  | 

### Return type

[**ExperimentSavedMetric**](ExperimentSavedMetric.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExperimentSavedMetricsRetrieve

> ExperimentSavedMetric ExperimentSavedMetricsRetrieve(ctx, id, projectId).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this experiment saved metric.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentSavedMetricsAPI.ExperimentSavedMetricsRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentSavedMetricsAPI.ExperimentSavedMetricsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentSavedMetricsRetrieve`: ExperimentSavedMetric
	fmt.Fprintf(os.Stdout, "Response from `ExperimentSavedMetricsAPI.ExperimentSavedMetricsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment saved metric. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentSavedMetricsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExperimentSavedMetric**](ExperimentSavedMetric.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExperimentSavedMetricsUpdate

> ExperimentSavedMetric ExperimentSavedMetricsUpdate(ctx, id, projectId).ExperimentSavedMetric(experimentSavedMetric).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this experiment saved metric.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	experimentSavedMetric := *openapiclient.NewExperimentSavedMetric(int32(123), "Name_example", interface{}(123), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now()) // ExperimentSavedMetric | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentSavedMetricsAPI.ExperimentSavedMetricsUpdate(context.Background(), id, projectId).ExperimentSavedMetric(experimentSavedMetric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentSavedMetricsAPI.ExperimentSavedMetricsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentSavedMetricsUpdate`: ExperimentSavedMetric
	fmt.Fprintf(os.Stdout, "Response from `ExperimentSavedMetricsAPI.ExperimentSavedMetricsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment saved metric. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentSavedMetricsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **experimentSavedMetric** | [**ExperimentSavedMetric**](ExperimentSavedMetric.md) |  | 

### Return type

[**ExperimentSavedMetric**](ExperimentSavedMetric.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

