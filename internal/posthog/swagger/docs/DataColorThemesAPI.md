# \DataColorThemesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataColorThemesCreate**](DataColorThemesAPI.md#DataColorThemesCreate) | **Post** /api/projects/{project_id}/data_color_themes/ | 
[**DataColorThemesDestroy**](DataColorThemesAPI.md#DataColorThemesDestroy) | **Delete** /api/projects/{project_id}/data_color_themes/{id}/ | 
[**DataColorThemesList**](DataColorThemesAPI.md#DataColorThemesList) | **Get** /api/projects/{project_id}/data_color_themes/ | 
[**DataColorThemesPartialUpdate**](DataColorThemesAPI.md#DataColorThemesPartialUpdate) | **Patch** /api/projects/{project_id}/data_color_themes/{id}/ | 
[**DataColorThemesRetrieve**](DataColorThemesAPI.md#DataColorThemesRetrieve) | **Get** /api/projects/{project_id}/data_color_themes/{id}/ | 
[**DataColorThemesUpdate**](DataColorThemesAPI.md#DataColorThemesUpdate) | **Put** /api/projects/{project_id}/data_color_themes/{id}/ | 



## DataColorThemesCreate

> DataColorTheme DataColorThemesCreate(ctx, projectId).DataColorTheme(dataColorTheme).Execute()



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
	dataColorTheme := *openapiclient.NewDataColorTheme(int32(123), "Name_example", "IsGlobal_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)})) // DataColorTheme | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataColorThemesAPI.DataColorThemesCreate(context.Background(), projectId).DataColorTheme(dataColorTheme).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataColorThemesAPI.DataColorThemesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataColorThemesCreate`: DataColorTheme
	fmt.Fprintf(os.Stdout, "Response from `DataColorThemesAPI.DataColorThemesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataColorThemesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataColorTheme** | [**DataColorTheme**](DataColorTheme.md) |  | 

### Return type

[**DataColorTheme**](DataColorTheme.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataColorThemesDestroy

> DataColorThemesDestroy(ctx, id, projectId).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this data color theme.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DataColorThemesAPI.DataColorThemesDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataColorThemesAPI.DataColorThemesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this data color theme. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataColorThemesDestroyRequest struct via the builder pattern


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


## DataColorThemesList

> PaginatedDataColorThemeList DataColorThemesList(ctx, projectId).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.DataColorThemesAPI.DataColorThemesList(context.Background(), projectId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataColorThemesAPI.DataColorThemesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataColorThemesList`: PaginatedDataColorThemeList
	fmt.Fprintf(os.Stdout, "Response from `DataColorThemesAPI.DataColorThemesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataColorThemesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedDataColorThemeList**](PaginatedDataColorThemeList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataColorThemesPartialUpdate

> DataColorTheme DataColorThemesPartialUpdate(ctx, id, projectId).PatchedDataColorTheme(patchedDataColorTheme).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this data color theme.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedDataColorTheme := *openapiclient.NewPatchedDataColorTheme() // PatchedDataColorTheme |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataColorThemesAPI.DataColorThemesPartialUpdate(context.Background(), id, projectId).PatchedDataColorTheme(patchedDataColorTheme).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataColorThemesAPI.DataColorThemesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataColorThemesPartialUpdate`: DataColorTheme
	fmt.Fprintf(os.Stdout, "Response from `DataColorThemesAPI.DataColorThemesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this data color theme. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataColorThemesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedDataColorTheme** | [**PatchedDataColorTheme**](PatchedDataColorTheme.md) |  | 

### Return type

[**DataColorTheme**](DataColorTheme.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataColorThemesRetrieve

> DataColorTheme DataColorThemesRetrieve(ctx, id, projectId).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this data color theme.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataColorThemesAPI.DataColorThemesRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataColorThemesAPI.DataColorThemesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataColorThemesRetrieve`: DataColorTheme
	fmt.Fprintf(os.Stdout, "Response from `DataColorThemesAPI.DataColorThemesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this data color theme. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataColorThemesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataColorTheme**](DataColorTheme.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DataColorThemesUpdate

> DataColorTheme DataColorThemesUpdate(ctx, id, projectId).DataColorTheme(dataColorTheme).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this data color theme.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	dataColorTheme := *openapiclient.NewDataColorTheme(int32(123), "Name_example", "IsGlobal_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)})) // DataColorTheme | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DataColorThemesAPI.DataColorThemesUpdate(context.Background(), id, projectId).DataColorTheme(dataColorTheme).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataColorThemesAPI.DataColorThemesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DataColorThemesUpdate`: DataColorTheme
	fmt.Fprintf(os.Stdout, "Response from `DataColorThemesAPI.DataColorThemesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this data color theme. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataColorThemesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataColorTheme** | [**DataColorTheme**](DataColorTheme.md) |  | 

### Return type

[**DataColorTheme**](DataColorTheme.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

