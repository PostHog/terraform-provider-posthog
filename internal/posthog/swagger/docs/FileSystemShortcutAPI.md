# \FileSystemShortcutAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileSystemShortcutCreate**](FileSystemShortcutAPI.md#FileSystemShortcutCreate) | **Post** /api/projects/{project_id}/file_system_shortcut/ | 
[**FileSystemShortcutDestroy**](FileSystemShortcutAPI.md#FileSystemShortcutDestroy) | **Delete** /api/projects/{project_id}/file_system_shortcut/{id}/ | 
[**FileSystemShortcutList**](FileSystemShortcutAPI.md#FileSystemShortcutList) | **Get** /api/projects/{project_id}/file_system_shortcut/ | 
[**FileSystemShortcutPartialUpdate**](FileSystemShortcutAPI.md#FileSystemShortcutPartialUpdate) | **Patch** /api/projects/{project_id}/file_system_shortcut/{id}/ | 
[**FileSystemShortcutRetrieve**](FileSystemShortcutAPI.md#FileSystemShortcutRetrieve) | **Get** /api/projects/{project_id}/file_system_shortcut/{id}/ | 
[**FileSystemShortcutUpdate**](FileSystemShortcutAPI.md#FileSystemShortcutUpdate) | **Put** /api/projects/{project_id}/file_system_shortcut/{id}/ | 



## FileSystemShortcutCreate

> FileSystemShortcut FileSystemShortcutCreate(ctx, projectId).FileSystemShortcut(fileSystemShortcut).Execute()



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
	fileSystemShortcut := *openapiclient.NewFileSystemShortcut("Id_example", "Path_example", time.Now()) // FileSystemShortcut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileSystemShortcutAPI.FileSystemShortcutCreate(context.Background(), projectId).FileSystemShortcut(fileSystemShortcut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileSystemShortcutAPI.FileSystemShortcutCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileSystemShortcutCreate`: FileSystemShortcut
	fmt.Fprintf(os.Stdout, "Response from `FileSystemShortcutAPI.FileSystemShortcutCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFileSystemShortcutCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileSystemShortcut** | [**FileSystemShortcut**](FileSystemShortcut.md) |  | 

### Return type

[**FileSystemShortcut**](FileSystemShortcut.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileSystemShortcutDestroy

> FileSystemShortcutDestroy(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this file system shortcut.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FileSystemShortcutAPI.FileSystemShortcutDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileSystemShortcutAPI.FileSystemShortcutDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this file system shortcut. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFileSystemShortcutDestroyRequest struct via the builder pattern


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


## FileSystemShortcutList

> PaginatedFileSystemShortcutList FileSystemShortcutList(ctx, projectId).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.FileSystemShortcutAPI.FileSystemShortcutList(context.Background(), projectId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileSystemShortcutAPI.FileSystemShortcutList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileSystemShortcutList`: PaginatedFileSystemShortcutList
	fmt.Fprintf(os.Stdout, "Response from `FileSystemShortcutAPI.FileSystemShortcutList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFileSystemShortcutListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedFileSystemShortcutList**](PaginatedFileSystemShortcutList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileSystemShortcutPartialUpdate

> FileSystemShortcut FileSystemShortcutPartialUpdate(ctx, id, projectId).PatchedFileSystemShortcut(patchedFileSystemShortcut).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this file system shortcut.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedFileSystemShortcut := *openapiclient.NewPatchedFileSystemShortcut() // PatchedFileSystemShortcut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileSystemShortcutAPI.FileSystemShortcutPartialUpdate(context.Background(), id, projectId).PatchedFileSystemShortcut(patchedFileSystemShortcut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileSystemShortcutAPI.FileSystemShortcutPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileSystemShortcutPartialUpdate`: FileSystemShortcut
	fmt.Fprintf(os.Stdout, "Response from `FileSystemShortcutAPI.FileSystemShortcutPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this file system shortcut. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFileSystemShortcutPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedFileSystemShortcut** | [**PatchedFileSystemShortcut**](PatchedFileSystemShortcut.md) |  | 

### Return type

[**FileSystemShortcut**](FileSystemShortcut.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileSystemShortcutRetrieve

> FileSystemShortcut FileSystemShortcutRetrieve(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this file system shortcut.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileSystemShortcutAPI.FileSystemShortcutRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileSystemShortcutAPI.FileSystemShortcutRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileSystemShortcutRetrieve`: FileSystemShortcut
	fmt.Fprintf(os.Stdout, "Response from `FileSystemShortcutAPI.FileSystemShortcutRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this file system shortcut. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFileSystemShortcutRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FileSystemShortcut**](FileSystemShortcut.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileSystemShortcutUpdate

> FileSystemShortcut FileSystemShortcutUpdate(ctx, id, projectId).FileSystemShortcut(fileSystemShortcut).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this file system shortcut.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	fileSystemShortcut := *openapiclient.NewFileSystemShortcut("Id_example", "Path_example", time.Now()) // FileSystemShortcut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileSystemShortcutAPI.FileSystemShortcutUpdate(context.Background(), id, projectId).FileSystemShortcut(fileSystemShortcut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileSystemShortcutAPI.FileSystemShortcutUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileSystemShortcutUpdate`: FileSystemShortcut
	fmt.Fprintf(os.Stdout, "Response from `FileSystemShortcutAPI.FileSystemShortcutUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this file system shortcut. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFileSystemShortcutUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fileSystemShortcut** | [**FileSystemShortcut**](FileSystemShortcut.md) |  | 

### Return type

[**FileSystemShortcut**](FileSystemShortcut.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

