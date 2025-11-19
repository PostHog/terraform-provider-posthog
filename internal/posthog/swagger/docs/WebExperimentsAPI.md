# \WebExperimentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebExperimentsCreate**](WebExperimentsAPI.md#WebExperimentsCreate) | **Post** /api/projects/{project_id}/web_experiments/ | 
[**WebExperimentsDestroy**](WebExperimentsAPI.md#WebExperimentsDestroy) | **Delete** /api/projects/{project_id}/web_experiments/{id}/ | 
[**WebExperimentsList**](WebExperimentsAPI.md#WebExperimentsList) | **Get** /api/projects/{project_id}/web_experiments/ | 
[**WebExperimentsPartialUpdate**](WebExperimentsAPI.md#WebExperimentsPartialUpdate) | **Patch** /api/projects/{project_id}/web_experiments/{id}/ | 
[**WebExperimentsRetrieve**](WebExperimentsAPI.md#WebExperimentsRetrieve) | **Get** /api/projects/{project_id}/web_experiments/{id}/ | 
[**WebExperimentsUpdate**](WebExperimentsAPI.md#WebExperimentsUpdate) | **Put** /api/projects/{project_id}/web_experiments/{id}/ | 



## WebExperimentsCreate

> WebExperimentsAPI WebExperimentsCreate(ctx, projectId).WebExperimentsAPI(webExperimentsAPI).Execute()



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
	webExperimentsAPI := *openapiclient.NewWebExperimentsAPI(int32(123), "Name_example", "FeatureFlagKey_example", interface{}(123)) // WebExperimentsAPI | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebExperimentsAPI.WebExperimentsCreate(context.Background(), projectId).WebExperimentsAPI(webExperimentsAPI).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebExperimentsAPI.WebExperimentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebExperimentsCreate`: WebExperimentsAPI
	fmt.Fprintf(os.Stdout, "Response from `WebExperimentsAPI.WebExperimentsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebExperimentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webExperimentsAPI** | [**WebExperimentsAPI**](WebExperimentsAPI.md) |  | 

### Return type

[**WebExperimentsAPI**](WebExperimentsAPI.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebExperimentsDestroy

> WebExperimentsDestroy(ctx, id, projectId).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this web experiment.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebExperimentsAPI.WebExperimentsDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebExperimentsAPI.WebExperimentsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this web experiment. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebExperimentsDestroyRequest struct via the builder pattern


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


## WebExperimentsList

> PaginatedWebExperimentsAPIList WebExperimentsList(ctx, projectId).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.WebExperimentsAPI.WebExperimentsList(context.Background(), projectId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebExperimentsAPI.WebExperimentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebExperimentsList`: PaginatedWebExperimentsAPIList
	fmt.Fprintf(os.Stdout, "Response from `WebExperimentsAPI.WebExperimentsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebExperimentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedWebExperimentsAPIList**](PaginatedWebExperimentsAPIList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebExperimentsPartialUpdate

> WebExperimentsAPI WebExperimentsPartialUpdate(ctx, id, projectId).PatchedWebExperimentsAPI(patchedWebExperimentsAPI).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this web experiment.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedWebExperimentsAPI := *openapiclient.NewPatchedWebExperimentsAPI() // PatchedWebExperimentsAPI |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebExperimentsAPI.WebExperimentsPartialUpdate(context.Background(), id, projectId).PatchedWebExperimentsAPI(patchedWebExperimentsAPI).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebExperimentsAPI.WebExperimentsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebExperimentsPartialUpdate`: WebExperimentsAPI
	fmt.Fprintf(os.Stdout, "Response from `WebExperimentsAPI.WebExperimentsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this web experiment. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebExperimentsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedWebExperimentsAPI** | [**PatchedWebExperimentsAPI**](PatchedWebExperimentsAPI.md) |  | 

### Return type

[**WebExperimentsAPI**](WebExperimentsAPI.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebExperimentsRetrieve

> WebExperimentsAPI WebExperimentsRetrieve(ctx, id, projectId).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this web experiment.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebExperimentsAPI.WebExperimentsRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebExperimentsAPI.WebExperimentsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebExperimentsRetrieve`: WebExperimentsAPI
	fmt.Fprintf(os.Stdout, "Response from `WebExperimentsAPI.WebExperimentsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this web experiment. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebExperimentsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WebExperimentsAPI**](WebExperimentsAPI.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebExperimentsUpdate

> WebExperimentsAPI WebExperimentsUpdate(ctx, id, projectId).WebExperimentsAPI(webExperimentsAPI).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this web experiment.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	webExperimentsAPI := *openapiclient.NewWebExperimentsAPI(int32(123), "Name_example", "FeatureFlagKey_example", interface{}(123)) // WebExperimentsAPI | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebExperimentsAPI.WebExperimentsUpdate(context.Background(), id, projectId).WebExperimentsAPI(webExperimentsAPI).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebExperimentsAPI.WebExperimentsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebExperimentsUpdate`: WebExperimentsAPI
	fmt.Fprintf(os.Stdout, "Response from `WebExperimentsAPI.WebExperimentsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this web experiment. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebExperimentsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **webExperimentsAPI** | [**WebExperimentsAPI**](WebExperimentsAPI.md) |  | 

### Return type

[**WebExperimentsAPI**](WebExperimentsAPI.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

