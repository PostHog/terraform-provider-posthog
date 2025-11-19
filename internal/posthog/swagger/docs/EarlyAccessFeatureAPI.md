# \EarlyAccessFeatureAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EarlyAccessFeatureCreate**](EarlyAccessFeatureAPI.md#EarlyAccessFeatureCreate) | **Post** /api/projects/{project_id}/early_access_feature/ | 
[**EarlyAccessFeatureDestroy**](EarlyAccessFeatureAPI.md#EarlyAccessFeatureDestroy) | **Delete** /api/projects/{project_id}/early_access_feature/{id}/ | 
[**EarlyAccessFeatureList**](EarlyAccessFeatureAPI.md#EarlyAccessFeatureList) | **Get** /api/projects/{project_id}/early_access_feature/ | 
[**EarlyAccessFeaturePartialUpdate**](EarlyAccessFeatureAPI.md#EarlyAccessFeaturePartialUpdate) | **Patch** /api/projects/{project_id}/early_access_feature/{id}/ | 
[**EarlyAccessFeatureRetrieve**](EarlyAccessFeatureAPI.md#EarlyAccessFeatureRetrieve) | **Get** /api/projects/{project_id}/early_access_feature/{id}/ | 
[**EarlyAccessFeatureUpdate**](EarlyAccessFeatureAPI.md#EarlyAccessFeatureUpdate) | **Put** /api/projects/{project_id}/early_access_feature/{id}/ | 



## EarlyAccessFeatureCreate

> EarlyAccessFeatureSerializerCreateOnly EarlyAccessFeatureCreate(ctx, projectId).EarlyAccessFeatureSerializerCreateOnly(earlyAccessFeatureSerializerCreateOnly).Execute()



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
	earlyAccessFeatureSerializerCreateOnly := *openapiclient.NewEarlyAccessFeatureSerializerCreateOnly("Id_example", "Name_example", openapiclient.StageEnum("draft"), time.Now(), *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"})) // EarlyAccessFeatureSerializerCreateOnly | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EarlyAccessFeatureAPI.EarlyAccessFeatureCreate(context.Background(), projectId).EarlyAccessFeatureSerializerCreateOnly(earlyAccessFeatureSerializerCreateOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EarlyAccessFeatureAPI.EarlyAccessFeatureCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EarlyAccessFeatureCreate`: EarlyAccessFeatureSerializerCreateOnly
	fmt.Fprintf(os.Stdout, "Response from `EarlyAccessFeatureAPI.EarlyAccessFeatureCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEarlyAccessFeatureCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **earlyAccessFeatureSerializerCreateOnly** | [**EarlyAccessFeatureSerializerCreateOnly**](EarlyAccessFeatureSerializerCreateOnly.md) |  | 

### Return type

[**EarlyAccessFeatureSerializerCreateOnly**](EarlyAccessFeatureSerializerCreateOnly.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EarlyAccessFeatureDestroy

> EarlyAccessFeatureDestroy(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this early access feature.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EarlyAccessFeatureAPI.EarlyAccessFeatureDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EarlyAccessFeatureAPI.EarlyAccessFeatureDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this early access feature. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEarlyAccessFeatureDestroyRequest struct via the builder pattern


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


## EarlyAccessFeatureList

> PaginatedEarlyAccessFeatureList EarlyAccessFeatureList(ctx, projectId).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.EarlyAccessFeatureAPI.EarlyAccessFeatureList(context.Background(), projectId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EarlyAccessFeatureAPI.EarlyAccessFeatureList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EarlyAccessFeatureList`: PaginatedEarlyAccessFeatureList
	fmt.Fprintf(os.Stdout, "Response from `EarlyAccessFeatureAPI.EarlyAccessFeatureList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEarlyAccessFeatureListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedEarlyAccessFeatureList**](PaginatedEarlyAccessFeatureList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EarlyAccessFeaturePartialUpdate

> EarlyAccessFeature EarlyAccessFeaturePartialUpdate(ctx, id, projectId).PatchedEarlyAccessFeature(patchedEarlyAccessFeature).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this early access feature.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedEarlyAccessFeature := *openapiclient.NewPatchedEarlyAccessFeature() // PatchedEarlyAccessFeature |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EarlyAccessFeatureAPI.EarlyAccessFeaturePartialUpdate(context.Background(), id, projectId).PatchedEarlyAccessFeature(patchedEarlyAccessFeature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EarlyAccessFeatureAPI.EarlyAccessFeaturePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EarlyAccessFeaturePartialUpdate`: EarlyAccessFeature
	fmt.Fprintf(os.Stdout, "Response from `EarlyAccessFeatureAPI.EarlyAccessFeaturePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this early access feature. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEarlyAccessFeaturePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedEarlyAccessFeature** | [**PatchedEarlyAccessFeature**](PatchedEarlyAccessFeature.md) |  | 

### Return type

[**EarlyAccessFeature**](EarlyAccessFeature.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EarlyAccessFeatureRetrieve

> EarlyAccessFeature EarlyAccessFeatureRetrieve(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this early access feature.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EarlyAccessFeatureAPI.EarlyAccessFeatureRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EarlyAccessFeatureAPI.EarlyAccessFeatureRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EarlyAccessFeatureRetrieve`: EarlyAccessFeature
	fmt.Fprintf(os.Stdout, "Response from `EarlyAccessFeatureAPI.EarlyAccessFeatureRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this early access feature. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEarlyAccessFeatureRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EarlyAccessFeature**](EarlyAccessFeature.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EarlyAccessFeatureUpdate

> EarlyAccessFeature EarlyAccessFeatureUpdate(ctx, id, projectId).EarlyAccessFeature(earlyAccessFeature).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this early access feature.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	earlyAccessFeature := *openapiclient.NewEarlyAccessFeature("Id_example", *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), "Name_example", openapiclient.StageEnum("draft"), time.Now()) // EarlyAccessFeature | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EarlyAccessFeatureAPI.EarlyAccessFeatureUpdate(context.Background(), id, projectId).EarlyAccessFeature(earlyAccessFeature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EarlyAccessFeatureAPI.EarlyAccessFeatureUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EarlyAccessFeatureUpdate`: EarlyAccessFeature
	fmt.Fprintf(os.Stdout, "Response from `EarlyAccessFeatureAPI.EarlyAccessFeatureUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this early access feature. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEarlyAccessFeatureUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **earlyAccessFeature** | [**EarlyAccessFeature**](EarlyAccessFeature.md) |  | 

### Return type

[**EarlyAccessFeature**](EarlyAccessFeature.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

