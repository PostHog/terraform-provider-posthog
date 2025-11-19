# \NotebooksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotebooksActivityRetrieve**](NotebooksAPI.md#NotebooksActivityRetrieve) | **Get** /api/projects/{project_id}/notebooks/activity/ | 
[**NotebooksActivityRetrieve2**](NotebooksAPI.md#NotebooksActivityRetrieve2) | **Get** /api/projects/{project_id}/notebooks/{short_id}/activity/ | 
[**NotebooksCreate**](NotebooksAPI.md#NotebooksCreate) | **Post** /api/projects/{project_id}/notebooks/ | 
[**NotebooksDestroy**](NotebooksAPI.md#NotebooksDestroy) | **Delete** /api/projects/{project_id}/notebooks/{short_id}/ | 
[**NotebooksList**](NotebooksAPI.md#NotebooksList) | **Get** /api/projects/{project_id}/notebooks/ | 
[**NotebooksPartialUpdate**](NotebooksAPI.md#NotebooksPartialUpdate) | **Patch** /api/projects/{project_id}/notebooks/{short_id}/ | 
[**NotebooksRecordingCommentsRetrieve**](NotebooksAPI.md#NotebooksRecordingCommentsRetrieve) | **Get** /api/projects/{project_id}/notebooks/recording_comments/ | 
[**NotebooksRetrieve**](NotebooksAPI.md#NotebooksRetrieve) | **Get** /api/projects/{project_id}/notebooks/{short_id}/ | 
[**NotebooksUpdate**](NotebooksAPI.md#NotebooksUpdate) | **Put** /api/projects/{project_id}/notebooks/{short_id}/ | 



## NotebooksActivityRetrieve

> NotebooksActivityRetrieve(ctx, projectId).Execute()





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
	r, err := apiClient.NotebooksAPI.NotebooksActivityRetrieve(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksAPI.NotebooksActivityRetrieve``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNotebooksActivityRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## NotebooksActivityRetrieve2

> NotebooksActivityRetrieve2(ctx, projectId, shortId).Execute()





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
	shortId := "shortId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotebooksAPI.NotebooksActivityRetrieve2(context.Background(), projectId, shortId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksAPI.NotebooksActivityRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**shortId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotebooksActivityRetrieve2Request struct via the builder pattern


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


## NotebooksCreate

> Notebook NotebooksCreate(ctx, projectId).Notebook(notebook).Execute()





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
	notebook := *openapiclient.NewNotebook("Id_example", "ShortId_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), "UserAccessLevel_example") // Notebook |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotebooksAPI.NotebooksCreate(context.Background(), projectId).Notebook(notebook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksAPI.NotebooksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotebooksCreate`: Notebook
	fmt.Fprintf(os.Stdout, "Response from `NotebooksAPI.NotebooksCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotebooksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notebook** | [**Notebook**](Notebook.md) |  | 

### Return type

[**Notebook**](Notebook.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotebooksDestroy

> NotebooksDestroy(ctx, projectId, shortId).Execute()





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
	shortId := "shortId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotebooksAPI.NotebooksDestroy(context.Background(), projectId, shortId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksAPI.NotebooksDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**shortId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotebooksDestroyRequest struct via the builder pattern


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


## NotebooksList

> PaginatedNotebookMinimalList NotebooksList(ctx, projectId).Contains(contains).CreatedBy(createdBy).DateFrom(dateFrom).DateTo(dateTo).Limit(limit).Offset(offset).User(user).Execute()





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
	contains := "recording:true" // string | Filter for notebooks that match a provided filter.                 Each match pair is separated by a colon,                 multiple match pairs can be sent separated by a space or a comma (optional)
	createdBy := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The UUID of the Notebook's creator (optional)
	dateFrom := time.Now() // time.Time | Filter for notebooks created after this date & time (optional)
	dateTo := time.Now() // time.Time | Filter for notebooks created before this date & time (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	user := "user_example" // string | If any value is provided for this parameter, return notebooks created by the logged in user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotebooksAPI.NotebooksList(context.Background(), projectId).Contains(contains).CreatedBy(createdBy).DateFrom(dateFrom).DateTo(dateTo).Limit(limit).Offset(offset).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksAPI.NotebooksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotebooksList`: PaginatedNotebookMinimalList
	fmt.Fprintf(os.Stdout, "Response from `NotebooksAPI.NotebooksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotebooksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contains** | **string** | Filter for notebooks that match a provided filter.                 Each match pair is separated by a colon,                 multiple match pairs can be sent separated by a space or a comma | 
 **createdBy** | **string** | The UUID of the Notebook&#39;s creator | 
 **dateFrom** | **time.Time** | Filter for notebooks created after this date &amp; time | 
 **dateTo** | **time.Time** | Filter for notebooks created before this date &amp; time | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **user** | **string** | If any value is provided for this parameter, return notebooks created by the logged in user. | 

### Return type

[**PaginatedNotebookMinimalList**](PaginatedNotebookMinimalList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotebooksPartialUpdate

> Notebook NotebooksPartialUpdate(ctx, projectId, shortId).PatchedNotebook(patchedNotebook).Execute()





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
	shortId := "shortId_example" // string | 
	patchedNotebook := *openapiclient.NewPatchedNotebook() // PatchedNotebook |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotebooksAPI.NotebooksPartialUpdate(context.Background(), projectId, shortId).PatchedNotebook(patchedNotebook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksAPI.NotebooksPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotebooksPartialUpdate`: Notebook
	fmt.Fprintf(os.Stdout, "Response from `NotebooksAPI.NotebooksPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**shortId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotebooksPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedNotebook** | [**PatchedNotebook**](PatchedNotebook.md) |  | 

### Return type

[**Notebook**](Notebook.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotebooksRecordingCommentsRetrieve

> NotebooksRecordingCommentsRetrieve(ctx, projectId).Execute()





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
	r, err := apiClient.NotebooksAPI.NotebooksRecordingCommentsRetrieve(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksAPI.NotebooksRecordingCommentsRetrieve``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNotebooksRecordingCommentsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## NotebooksRetrieve

> Notebook NotebooksRetrieve(ctx, projectId, shortId).Execute()





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
	shortId := "shortId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotebooksAPI.NotebooksRetrieve(context.Background(), projectId, shortId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksAPI.NotebooksRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotebooksRetrieve`: Notebook
	fmt.Fprintf(os.Stdout, "Response from `NotebooksAPI.NotebooksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**shortId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotebooksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Notebook**](Notebook.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotebooksUpdate

> Notebook NotebooksUpdate(ctx, projectId, shortId).Notebook(notebook).Execute()





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
	shortId := "shortId_example" // string | 
	notebook := *openapiclient.NewNotebook("Id_example", "ShortId_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), "UserAccessLevel_example") // Notebook |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotebooksAPI.NotebooksUpdate(context.Background(), projectId, shortId).Notebook(notebook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotebooksAPI.NotebooksUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotebooksUpdate`: Notebook
	fmt.Fprintf(os.Stdout, "Response from `NotebooksAPI.NotebooksUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**shortId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotebooksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **notebook** | [**Notebook**](Notebook.md) |  | 

### Return type

[**Notebook**](Notebook.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

