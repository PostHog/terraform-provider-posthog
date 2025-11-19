# \SurveysAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SurveysActivityRetrieve**](SurveysAPI.md#SurveysActivityRetrieve) | **Get** /api/projects/{project_id}/surveys/activity/ | 
[**SurveysActivityRetrieve2**](SurveysAPI.md#SurveysActivityRetrieve2) | **Get** /api/projects/{project_id}/surveys/{id}/activity/ | 
[**SurveysCreate**](SurveysAPI.md#SurveysCreate) | **Post** /api/projects/{project_id}/surveys/ | 
[**SurveysDestroy**](SurveysAPI.md#SurveysDestroy) | **Delete** /api/projects/{project_id}/surveys/{id}/ | 
[**SurveysDuplicateToProjectsCreate**](SurveysAPI.md#SurveysDuplicateToProjectsCreate) | **Post** /api/projects/{project_id}/surveys/{id}/duplicate_to_projects/ | 
[**SurveysList**](SurveysAPI.md#SurveysList) | **Get** /api/projects/{project_id}/surveys/ | 
[**SurveysPartialUpdate**](SurveysAPI.md#SurveysPartialUpdate) | **Patch** /api/projects/{project_id}/surveys/{id}/ | 
[**SurveysResponsesCountRetrieve**](SurveysAPI.md#SurveysResponsesCountRetrieve) | **Get** /api/projects/{project_id}/surveys/responses_count/ | 
[**SurveysRetrieve**](SurveysAPI.md#SurveysRetrieve) | **Get** /api/projects/{project_id}/surveys/{id}/ | 
[**SurveysStatsRetrieve**](SurveysAPI.md#SurveysStatsRetrieve) | **Get** /api/projects/{project_id}/surveys/stats/ | 
[**SurveysStatsRetrieve2**](SurveysAPI.md#SurveysStatsRetrieve2) | **Get** /api/projects/{project_id}/surveys/{id}/stats/ | 
[**SurveysSummarizeResponsesCreate**](SurveysAPI.md#SurveysSummarizeResponsesCreate) | **Post** /api/projects/{project_id}/surveys/{id}/summarize_responses/ | 
[**SurveysUpdate**](SurveysAPI.md#SurveysUpdate) | **Put** /api/projects/{project_id}/surveys/{id}/ | 



## SurveysActivityRetrieve

> SurveysActivityRetrieve(ctx, projectId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveysAPI.SurveysActivityRetrieve(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysActivityRetrieve``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSurveysActivityRetrieveRequest struct via the builder pattern


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


## SurveysActivityRetrieve2

> SurveysActivityRetrieve2(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this survey.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveysAPI.SurveysActivityRetrieve2(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysActivityRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this survey. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSurveysActivityRetrieve2Request struct via the builder pattern


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


## SurveysCreate

> SurveySerializerCreateUpdateOnly SurveysCreate(ctx, projectId).SurveySerializerCreateUpdateOnly(surveySerializerCreateUpdateOnly).Execute()



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
	surveySerializerCreateUpdateOnly := *openapiclient.NewSurveySerializerCreateUpdateOnly("Id_example", "Name_example", openapiclient.SurveyType("popover"), *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)})) // SurveySerializerCreateUpdateOnly | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveysAPI.SurveysCreate(context.Background(), projectId).SurveySerializerCreateUpdateOnly(surveySerializerCreateUpdateOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveysCreate`: SurveySerializerCreateUpdateOnly
	fmt.Fprintf(os.Stdout, "Response from `SurveysAPI.SurveysCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSurveysCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **surveySerializerCreateUpdateOnly** | [**SurveySerializerCreateUpdateOnly**](SurveySerializerCreateUpdateOnly.md) |  | 

### Return type

[**SurveySerializerCreateUpdateOnly**](SurveySerializerCreateUpdateOnly.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveysDestroy

> SurveysDestroy(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this survey.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveysAPI.SurveysDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this survey. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSurveysDestroyRequest struct via the builder pattern


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


## SurveysDuplicateToProjectsCreate

> SurveysDuplicateToProjectsCreate(ctx, id, projectId).SurveySerializerCreateUpdateOnly(surveySerializerCreateUpdateOnly).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this survey.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	surveySerializerCreateUpdateOnly := *openapiclient.NewSurveySerializerCreateUpdateOnly("Id_example", "Name_example", openapiclient.SurveyType("popover"), *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)})) // SurveySerializerCreateUpdateOnly | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveysAPI.SurveysDuplicateToProjectsCreate(context.Background(), id, projectId).SurveySerializerCreateUpdateOnly(surveySerializerCreateUpdateOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysDuplicateToProjectsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this survey. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSurveysDuplicateToProjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **surveySerializerCreateUpdateOnly** | [**SurveySerializerCreateUpdateOnly**](SurveySerializerCreateUpdateOnly.md) |  | 

### Return type

 (empty response body)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveysList

> PaginatedSurveyList SurveysList(ctx, projectId).Limit(limit).Offset(offset).Search(search).Execute()



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
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveysAPI.SurveysList(context.Background(), projectId).Limit(limit).Offset(offset).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveysList`: PaginatedSurveyList
	fmt.Fprintf(os.Stdout, "Response from `SurveysAPI.SurveysList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSurveysListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedSurveyList**](PaginatedSurveyList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveysPartialUpdate

> SurveySerializerCreateUpdateOnly SurveysPartialUpdate(ctx, id, projectId).PatchedSurveySerializerCreateUpdateOnly(patchedSurveySerializerCreateUpdateOnly).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this survey.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedSurveySerializerCreateUpdateOnly := *openapiclient.NewPatchedSurveySerializerCreateUpdateOnly() // PatchedSurveySerializerCreateUpdateOnly |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveysAPI.SurveysPartialUpdate(context.Background(), id, projectId).PatchedSurveySerializerCreateUpdateOnly(patchedSurveySerializerCreateUpdateOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveysPartialUpdate`: SurveySerializerCreateUpdateOnly
	fmt.Fprintf(os.Stdout, "Response from `SurveysAPI.SurveysPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this survey. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSurveysPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedSurveySerializerCreateUpdateOnly** | [**PatchedSurveySerializerCreateUpdateOnly**](PatchedSurveySerializerCreateUpdateOnly.md) |  | 

### Return type

[**SurveySerializerCreateUpdateOnly**](SurveySerializerCreateUpdateOnly.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveysResponsesCountRetrieve

> SurveysResponsesCountRetrieve(ctx, projectId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveysAPI.SurveysResponsesCountRetrieve(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysResponsesCountRetrieve``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSurveysResponsesCountRetrieveRequest struct via the builder pattern


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


## SurveysRetrieve

> Survey SurveysRetrieve(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this survey.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveysAPI.SurveysRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveysRetrieve`: Survey
	fmt.Fprintf(os.Stdout, "Response from `SurveysAPI.SurveysRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this survey. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSurveysRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Survey**](Survey.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveysStatsRetrieve

> SurveysStatsRetrieve(ctx, projectId).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveysAPI.SurveysStatsRetrieve(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysStatsRetrieve``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSurveysStatsRetrieveRequest struct via the builder pattern


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


## SurveysStatsRetrieve2

> SurveysStatsRetrieve2(ctx, id, projectId).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this survey.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveysAPI.SurveysStatsRetrieve2(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysStatsRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this survey. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSurveysStatsRetrieve2Request struct via the builder pattern


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


## SurveysSummarizeResponsesCreate

> SurveysSummarizeResponsesCreate(ctx, id, projectId).SurveySerializerCreateUpdateOnly(surveySerializerCreateUpdateOnly).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this survey.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	surveySerializerCreateUpdateOnly := *openapiclient.NewSurveySerializerCreateUpdateOnly("Id_example", "Name_example", openapiclient.SurveyType("popover"), *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)})) // SurveySerializerCreateUpdateOnly | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SurveysAPI.SurveysSummarizeResponsesCreate(context.Background(), id, projectId).SurveySerializerCreateUpdateOnly(surveySerializerCreateUpdateOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysSummarizeResponsesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this survey. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSurveysSummarizeResponsesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **surveySerializerCreateUpdateOnly** | [**SurveySerializerCreateUpdateOnly**](SurveySerializerCreateUpdateOnly.md) |  | 

### Return type

 (empty response body)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SurveysUpdate

> Survey SurveysUpdate(ctx, id, projectId).Survey(survey).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this survey.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	survey := *openapiclient.NewSurvey("Id_example", "Name_example", openapiclient.SurveyType("popover"), *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), "Conditions_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), []interface{}{nil}, "UserAccessLevel_example") // Survey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SurveysAPI.SurveysUpdate(context.Background(), id, projectId).Survey(survey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SurveysAPI.SurveysUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SurveysUpdate`: Survey
	fmt.Fprintf(os.Stdout, "Response from `SurveysAPI.SurveysUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this survey. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSurveysUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **survey** | [**Survey**](Survey.md) |  | 

### Return type

[**Survey**](Survey.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

