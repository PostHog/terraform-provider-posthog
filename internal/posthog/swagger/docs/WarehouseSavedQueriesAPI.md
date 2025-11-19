# \WarehouseSavedQueriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WarehouseSavedQueriesActivityRetrieve**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesActivityRetrieve) | **Get** /api/projects/{project_id}/warehouse_saved_queries/{id}/activity/ | 
[**WarehouseSavedQueriesAncestorsCreate**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesAncestorsCreate) | **Post** /api/projects/{project_id}/warehouse_saved_queries/{id}/ancestors/ | 
[**WarehouseSavedQueriesCancelCreate**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesCancelCreate) | **Post** /api/projects/{project_id}/warehouse_saved_queries/{id}/cancel/ | 
[**WarehouseSavedQueriesCreate**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesCreate) | **Post** /api/projects/{project_id}/warehouse_saved_queries/ | 
[**WarehouseSavedQueriesDependenciesRetrieve**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesDependenciesRetrieve) | **Get** /api/projects/{project_id}/warehouse_saved_queries/{id}/dependencies/ | 
[**WarehouseSavedQueriesDescendantsCreate**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesDescendantsCreate) | **Post** /api/projects/{project_id}/warehouse_saved_queries/{id}/descendants/ | 
[**WarehouseSavedQueriesDestroy**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesDestroy) | **Delete** /api/projects/{project_id}/warehouse_saved_queries/{id}/ | 
[**WarehouseSavedQueriesList**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesList) | **Get** /api/projects/{project_id}/warehouse_saved_queries/ | 
[**WarehouseSavedQueriesPartialUpdate**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesPartialUpdate) | **Patch** /api/projects/{project_id}/warehouse_saved_queries/{id}/ | 
[**WarehouseSavedQueriesRetrieve**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesRetrieve) | **Get** /api/projects/{project_id}/warehouse_saved_queries/{id}/ | 
[**WarehouseSavedQueriesRevertMaterializationCreate**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesRevertMaterializationCreate) | **Post** /api/projects/{project_id}/warehouse_saved_queries/{id}/revert_materialization/ | 
[**WarehouseSavedQueriesRunCreate**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesRunCreate) | **Post** /api/projects/{project_id}/warehouse_saved_queries/{id}/run/ | 
[**WarehouseSavedQueriesRunHistoryRetrieve**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesRunHistoryRetrieve) | **Get** /api/projects/{project_id}/warehouse_saved_queries/{id}/run_history/ | 
[**WarehouseSavedQueriesUpdate**](WarehouseSavedQueriesAPI.md#WarehouseSavedQueriesUpdate) | **Put** /api/projects/{project_id}/warehouse_saved_queries/{id}/ | 



## WarehouseSavedQueriesActivityRetrieve

> DataWarehouseSavedQuery WarehouseSavedQueriesActivityRetrieve(ctx, id, projectId).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse saved query.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesActivityRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesActivityRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseSavedQueriesActivityRetrieve`: DataWarehouseSavedQuery
	fmt.Fprintf(os.Stdout, "Response from `WarehouseSavedQueriesAPI.WarehouseSavedQueriesActivityRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse saved query. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesActivityRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseSavedQueriesAncestorsCreate

> DataWarehouseSavedQuery WarehouseSavedQueriesAncestorsCreate(ctx, id, projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse saved query.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	dataWarehouseSavedQuery := *openapiclient.NewDataWarehouseSavedQuery("Id_example", "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "SyncFrequency_example", "Columns_example", "Status_example", time.Now(), "ManagedViewsetKind_example", "LatestError_example", "LatestHistoryId_example", false) // DataWarehouseSavedQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesAncestorsCreate(context.Background(), id, projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesAncestorsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseSavedQueriesAncestorsCreate`: DataWarehouseSavedQuery
	fmt.Fprintf(os.Stdout, "Response from `WarehouseSavedQueriesAPI.WarehouseSavedQueriesAncestorsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse saved query. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesAncestorsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataWarehouseSavedQuery** | [**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md) |  | 

### Return type

[**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseSavedQueriesCancelCreate

> DataWarehouseSavedQuery WarehouseSavedQueriesCancelCreate(ctx, id, projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse saved query.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	dataWarehouseSavedQuery := *openapiclient.NewDataWarehouseSavedQuery("Id_example", "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "SyncFrequency_example", "Columns_example", "Status_example", time.Now(), "ManagedViewsetKind_example", "LatestError_example", "LatestHistoryId_example", false) // DataWarehouseSavedQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesCancelCreate(context.Background(), id, projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesCancelCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseSavedQueriesCancelCreate`: DataWarehouseSavedQuery
	fmt.Fprintf(os.Stdout, "Response from `WarehouseSavedQueriesAPI.WarehouseSavedQueriesCancelCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse saved query. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesCancelCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataWarehouseSavedQuery** | [**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md) |  | 

### Return type

[**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseSavedQueriesCreate

> DataWarehouseSavedQuery WarehouseSavedQueriesCreate(ctx, projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()





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
	dataWarehouseSavedQuery := *openapiclient.NewDataWarehouseSavedQuery("Id_example", "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "SyncFrequency_example", "Columns_example", "Status_example", time.Now(), "ManagedViewsetKind_example", "LatestError_example", "LatestHistoryId_example", false) // DataWarehouseSavedQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesCreate(context.Background(), projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseSavedQueriesCreate`: DataWarehouseSavedQuery
	fmt.Fprintf(os.Stdout, "Response from `WarehouseSavedQueriesAPI.WarehouseSavedQueriesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataWarehouseSavedQuery** | [**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md) |  | 

### Return type

[**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseSavedQueriesDependenciesRetrieve

> DataWarehouseSavedQuery WarehouseSavedQueriesDependenciesRetrieve(ctx, id, projectId).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse saved query.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesDependenciesRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesDependenciesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseSavedQueriesDependenciesRetrieve`: DataWarehouseSavedQuery
	fmt.Fprintf(os.Stdout, "Response from `WarehouseSavedQueriesAPI.WarehouseSavedQueriesDependenciesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse saved query. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesDependenciesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseSavedQueriesDescendantsCreate

> DataWarehouseSavedQuery WarehouseSavedQueriesDescendantsCreate(ctx, id, projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse saved query.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	dataWarehouseSavedQuery := *openapiclient.NewDataWarehouseSavedQuery("Id_example", "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "SyncFrequency_example", "Columns_example", "Status_example", time.Now(), "ManagedViewsetKind_example", "LatestError_example", "LatestHistoryId_example", false) // DataWarehouseSavedQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesDescendantsCreate(context.Background(), id, projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesDescendantsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseSavedQueriesDescendantsCreate`: DataWarehouseSavedQuery
	fmt.Fprintf(os.Stdout, "Response from `WarehouseSavedQueriesAPI.WarehouseSavedQueriesDescendantsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse saved query. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesDescendantsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataWarehouseSavedQuery** | [**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md) |  | 

### Return type

[**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseSavedQueriesDestroy

> WarehouseSavedQueriesDestroy(ctx, id, projectId).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse saved query.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse saved query. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesDestroyRequest struct via the builder pattern


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


## WarehouseSavedQueriesList

> PaginatedDataWarehouseSavedQueryList WarehouseSavedQueriesList(ctx, projectId).Page(page).Search(search).Execute()





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
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesList(context.Background(), projectId).Page(page).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseSavedQueriesList`: PaginatedDataWarehouseSavedQueryList
	fmt.Fprintf(os.Stdout, "Response from `WarehouseSavedQueriesAPI.WarehouseSavedQueriesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedDataWarehouseSavedQueryList**](PaginatedDataWarehouseSavedQueryList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseSavedQueriesPartialUpdate

> DataWarehouseSavedQuery WarehouseSavedQueriesPartialUpdate(ctx, id, projectId).PatchedDataWarehouseSavedQuery(patchedDataWarehouseSavedQuery).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse saved query.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedDataWarehouseSavedQuery := *openapiclient.NewPatchedDataWarehouseSavedQuery() // PatchedDataWarehouseSavedQuery |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesPartialUpdate(context.Background(), id, projectId).PatchedDataWarehouseSavedQuery(patchedDataWarehouseSavedQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseSavedQueriesPartialUpdate`: DataWarehouseSavedQuery
	fmt.Fprintf(os.Stdout, "Response from `WarehouseSavedQueriesAPI.WarehouseSavedQueriesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse saved query. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedDataWarehouseSavedQuery** | [**PatchedDataWarehouseSavedQuery**](PatchedDataWarehouseSavedQuery.md) |  | 

### Return type

[**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseSavedQueriesRetrieve

> DataWarehouseSavedQuery WarehouseSavedQueriesRetrieve(ctx, id, projectId).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse saved query.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseSavedQueriesRetrieve`: DataWarehouseSavedQuery
	fmt.Fprintf(os.Stdout, "Response from `WarehouseSavedQueriesAPI.WarehouseSavedQueriesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse saved query. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseSavedQueriesRevertMaterializationCreate

> DataWarehouseSavedQuery WarehouseSavedQueriesRevertMaterializationCreate(ctx, id, projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse saved query.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	dataWarehouseSavedQuery := *openapiclient.NewDataWarehouseSavedQuery("Id_example", "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "SyncFrequency_example", "Columns_example", "Status_example", time.Now(), "ManagedViewsetKind_example", "LatestError_example", "LatestHistoryId_example", false) // DataWarehouseSavedQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesRevertMaterializationCreate(context.Background(), id, projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesRevertMaterializationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseSavedQueriesRevertMaterializationCreate`: DataWarehouseSavedQuery
	fmt.Fprintf(os.Stdout, "Response from `WarehouseSavedQueriesAPI.WarehouseSavedQueriesRevertMaterializationCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse saved query. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesRevertMaterializationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataWarehouseSavedQuery** | [**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md) |  | 

### Return type

[**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseSavedQueriesRunCreate

> DataWarehouseSavedQuery WarehouseSavedQueriesRunCreate(ctx, id, projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse saved query.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	dataWarehouseSavedQuery := *openapiclient.NewDataWarehouseSavedQuery("Id_example", "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "SyncFrequency_example", "Columns_example", "Status_example", time.Now(), "ManagedViewsetKind_example", "LatestError_example", "LatestHistoryId_example", false) // DataWarehouseSavedQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesRunCreate(context.Background(), id, projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesRunCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseSavedQueriesRunCreate`: DataWarehouseSavedQuery
	fmt.Fprintf(os.Stdout, "Response from `WarehouseSavedQueriesAPI.WarehouseSavedQueriesRunCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse saved query. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesRunCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataWarehouseSavedQuery** | [**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md) |  | 

### Return type

[**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseSavedQueriesRunHistoryRetrieve

> DataWarehouseSavedQuery WarehouseSavedQueriesRunHistoryRetrieve(ctx, id, projectId).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse saved query.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesRunHistoryRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesRunHistoryRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseSavedQueriesRunHistoryRetrieve`: DataWarehouseSavedQuery
	fmt.Fprintf(os.Stdout, "Response from `WarehouseSavedQueriesAPI.WarehouseSavedQueriesRunHistoryRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse saved query. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesRunHistoryRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseSavedQueriesUpdate

> DataWarehouseSavedQuery WarehouseSavedQueriesUpdate(ctx, id, projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse saved query.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	dataWarehouseSavedQuery := *openapiclient.NewDataWarehouseSavedQuery("Id_example", "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "SyncFrequency_example", "Columns_example", "Status_example", time.Now(), "ManagedViewsetKind_example", "LatestError_example", "LatestHistoryId_example", false) // DataWarehouseSavedQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseSavedQueriesAPI.WarehouseSavedQueriesUpdate(context.Background(), id, projectId).DataWarehouseSavedQuery(dataWarehouseSavedQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseSavedQueriesAPI.WarehouseSavedQueriesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseSavedQueriesUpdate`: DataWarehouseSavedQuery
	fmt.Fprintf(os.Stdout, "Response from `WarehouseSavedQueriesAPI.WarehouseSavedQueriesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse saved query. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseSavedQueriesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dataWarehouseSavedQuery** | [**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md) |  | 

### Return type

[**DataWarehouseSavedQuery**](DataWarehouseSavedQuery.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

