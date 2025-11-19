# \ExperimentHoldoutsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExperimentHoldoutsCreate**](ExperimentHoldoutsAPI.md#ExperimentHoldoutsCreate) | **Post** /api/projects/{project_id}/experiment_holdouts/ | 
[**ExperimentHoldoutsDestroy**](ExperimentHoldoutsAPI.md#ExperimentHoldoutsDestroy) | **Delete** /api/projects/{project_id}/experiment_holdouts/{id}/ | 
[**ExperimentHoldoutsList**](ExperimentHoldoutsAPI.md#ExperimentHoldoutsList) | **Get** /api/projects/{project_id}/experiment_holdouts/ | 
[**ExperimentHoldoutsPartialUpdate**](ExperimentHoldoutsAPI.md#ExperimentHoldoutsPartialUpdate) | **Patch** /api/projects/{project_id}/experiment_holdouts/{id}/ | 
[**ExperimentHoldoutsRetrieve**](ExperimentHoldoutsAPI.md#ExperimentHoldoutsRetrieve) | **Get** /api/projects/{project_id}/experiment_holdouts/{id}/ | 
[**ExperimentHoldoutsUpdate**](ExperimentHoldoutsAPI.md#ExperimentHoldoutsUpdate) | **Put** /api/projects/{project_id}/experiment_holdouts/{id}/ | 



## ExperimentHoldoutsCreate

> ExperimentHoldout ExperimentHoldoutsCreate(ctx, projectId).ExperimentHoldout(experimentHoldout).Execute()



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
	experimentHoldout := *openapiclient.NewExperimentHoldout(int32(123), "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now()) // ExperimentHoldout | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentHoldoutsAPI.ExperimentHoldoutsCreate(context.Background(), projectId).ExperimentHoldout(experimentHoldout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentHoldoutsAPI.ExperimentHoldoutsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentHoldoutsCreate`: ExperimentHoldout
	fmt.Fprintf(os.Stdout, "Response from `ExperimentHoldoutsAPI.ExperimentHoldoutsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentHoldoutsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experimentHoldout** | [**ExperimentHoldout**](ExperimentHoldout.md) |  | 

### Return type

[**ExperimentHoldout**](ExperimentHoldout.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExperimentHoldoutsDestroy

> ExperimentHoldoutsDestroy(ctx, id, projectId).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this experiment holdout.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExperimentHoldoutsAPI.ExperimentHoldoutsDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentHoldoutsAPI.ExperimentHoldoutsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment holdout. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentHoldoutsDestroyRequest struct via the builder pattern


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


## ExperimentHoldoutsList

> PaginatedExperimentHoldoutList ExperimentHoldoutsList(ctx, projectId).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.ExperimentHoldoutsAPI.ExperimentHoldoutsList(context.Background(), projectId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentHoldoutsAPI.ExperimentHoldoutsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentHoldoutsList`: PaginatedExperimentHoldoutList
	fmt.Fprintf(os.Stdout, "Response from `ExperimentHoldoutsAPI.ExperimentHoldoutsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentHoldoutsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedExperimentHoldoutList**](PaginatedExperimentHoldoutList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExperimentHoldoutsPartialUpdate

> ExperimentHoldout ExperimentHoldoutsPartialUpdate(ctx, id, projectId).PatchedExperimentHoldout(patchedExperimentHoldout).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this experiment holdout.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedExperimentHoldout := *openapiclient.NewPatchedExperimentHoldout() // PatchedExperimentHoldout |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentHoldoutsAPI.ExperimentHoldoutsPartialUpdate(context.Background(), id, projectId).PatchedExperimentHoldout(patchedExperimentHoldout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentHoldoutsAPI.ExperimentHoldoutsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentHoldoutsPartialUpdate`: ExperimentHoldout
	fmt.Fprintf(os.Stdout, "Response from `ExperimentHoldoutsAPI.ExperimentHoldoutsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment holdout. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentHoldoutsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedExperimentHoldout** | [**PatchedExperimentHoldout**](PatchedExperimentHoldout.md) |  | 

### Return type

[**ExperimentHoldout**](ExperimentHoldout.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExperimentHoldoutsRetrieve

> ExperimentHoldout ExperimentHoldoutsRetrieve(ctx, id, projectId).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this experiment holdout.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentHoldoutsAPI.ExperimentHoldoutsRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentHoldoutsAPI.ExperimentHoldoutsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentHoldoutsRetrieve`: ExperimentHoldout
	fmt.Fprintf(os.Stdout, "Response from `ExperimentHoldoutsAPI.ExperimentHoldoutsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment holdout. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentHoldoutsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExperimentHoldout**](ExperimentHoldout.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExperimentHoldoutsUpdate

> ExperimentHoldout ExperimentHoldoutsUpdate(ctx, id, projectId).ExperimentHoldout(experimentHoldout).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this experiment holdout.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	experimentHoldout := *openapiclient.NewExperimentHoldout(int32(123), "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now()) // ExperimentHoldout | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentHoldoutsAPI.ExperimentHoldoutsUpdate(context.Background(), id, projectId).ExperimentHoldout(experimentHoldout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentHoldoutsAPI.ExperimentHoldoutsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentHoldoutsUpdate`: ExperimentHoldout
	fmt.Fprintf(os.Stdout, "Response from `ExperimentHoldoutsAPI.ExperimentHoldoutsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment holdout. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentHoldoutsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **experimentHoldout** | [**ExperimentHoldout**](ExperimentHoldout.md) |  | 

### Return type

[**ExperimentHoldout**](ExperimentHoldout.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

