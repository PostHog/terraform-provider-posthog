# \BatchExportsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchExportsBackfillCreate**](BatchExportsAPI.md#BatchExportsBackfillCreate) | **Post** /api/organizations/{organization_id}/batch_exports/{id}/backfill/ | 
[**BatchExportsBackfillCreate2**](BatchExportsAPI.md#BatchExportsBackfillCreate2) | **Post** /api/projects/{project_id}/batch_exports/{id}/backfill/ | 
[**BatchExportsBackfillsCancelCreate**](BatchExportsAPI.md#BatchExportsBackfillsCancelCreate) | **Post** /api/projects/{project_id}/batch_exports/{batch_export_id}/backfills/{id}/cancel/ | 
[**BatchExportsBackfillsCreate**](BatchExportsAPI.md#BatchExportsBackfillsCreate) | **Post** /api/projects/{project_id}/batch_exports/{batch_export_id}/backfills/ | 
[**BatchExportsBackfillsList**](BatchExportsAPI.md#BatchExportsBackfillsList) | **Get** /api/projects/{project_id}/batch_exports/{batch_export_id}/backfills/ | 
[**BatchExportsBackfillsRetrieve**](BatchExportsAPI.md#BatchExportsBackfillsRetrieve) | **Get** /api/projects/{project_id}/batch_exports/{batch_export_id}/backfills/{id}/ | 
[**BatchExportsCreate**](BatchExportsAPI.md#BatchExportsCreate) | **Post** /api/organizations/{organization_id}/batch_exports/ | 
[**BatchExportsCreate2**](BatchExportsAPI.md#BatchExportsCreate2) | **Post** /api/projects/{project_id}/batch_exports/ | 
[**BatchExportsDestroy**](BatchExportsAPI.md#BatchExportsDestroy) | **Delete** /api/organizations/{organization_id}/batch_exports/{id}/ | 
[**BatchExportsDestroy2**](BatchExportsAPI.md#BatchExportsDestroy2) | **Delete** /api/projects/{project_id}/batch_exports/{id}/ | 
[**BatchExportsList**](BatchExportsAPI.md#BatchExportsList) | **Get** /api/organizations/{organization_id}/batch_exports/ | 
[**BatchExportsList2**](BatchExportsAPI.md#BatchExportsList2) | **Get** /api/projects/{project_id}/batch_exports/ | 
[**BatchExportsLogsRetrieve**](BatchExportsAPI.md#BatchExportsLogsRetrieve) | **Get** /api/organizations/{organization_id}/batch_exports/{id}/logs/ | 
[**BatchExportsLogsRetrieve2**](BatchExportsAPI.md#BatchExportsLogsRetrieve2) | **Get** /api/projects/{project_id}/batch_exports/{id}/logs/ | 
[**BatchExportsPartialUpdate**](BatchExportsAPI.md#BatchExportsPartialUpdate) | **Patch** /api/organizations/{organization_id}/batch_exports/{id}/ | 
[**BatchExportsPartialUpdate2**](BatchExportsAPI.md#BatchExportsPartialUpdate2) | **Patch** /api/projects/{project_id}/batch_exports/{id}/ | 
[**BatchExportsPauseCreate**](BatchExportsAPI.md#BatchExportsPauseCreate) | **Post** /api/organizations/{organization_id}/batch_exports/{id}/pause/ | 
[**BatchExportsPauseCreate2**](BatchExportsAPI.md#BatchExportsPauseCreate2) | **Post** /api/projects/{project_id}/batch_exports/{id}/pause/ | 
[**BatchExportsRetrieve**](BatchExportsAPI.md#BatchExportsRetrieve) | **Get** /api/organizations/{organization_id}/batch_exports/{id}/ | 
[**BatchExportsRetrieve2**](BatchExportsAPI.md#BatchExportsRetrieve2) | **Get** /api/projects/{project_id}/batch_exports/{id}/ | 
[**BatchExportsRunTestStepCreate**](BatchExportsAPI.md#BatchExportsRunTestStepCreate) | **Post** /api/organizations/{organization_id}/batch_exports/{id}/run_test_step/ | 
[**BatchExportsRunTestStepCreate2**](BatchExportsAPI.md#BatchExportsRunTestStepCreate2) | **Post** /api/projects/{project_id}/batch_exports/{id}/run_test_step/ | 
[**BatchExportsRunTestStepNewCreate**](BatchExportsAPI.md#BatchExportsRunTestStepNewCreate) | **Post** /api/organizations/{organization_id}/batch_exports/run_test_step_new/ | 
[**BatchExportsRunTestStepNewCreate2**](BatchExportsAPI.md#BatchExportsRunTestStepNewCreate2) | **Post** /api/projects/{project_id}/batch_exports/run_test_step_new/ | 
[**BatchExportsRunsCancelCreate**](BatchExportsAPI.md#BatchExportsRunsCancelCreate) | **Post** /api/projects/{project_id}/batch_exports/{batch_export_id}/runs/{id}/cancel/ | 
[**BatchExportsRunsList**](BatchExportsAPI.md#BatchExportsRunsList) | **Get** /api/projects/{project_id}/batch_exports/{batch_export_id}/runs/ | 
[**BatchExportsRunsLogsRetrieve**](BatchExportsAPI.md#BatchExportsRunsLogsRetrieve) | **Get** /api/projects/{project_id}/batch_exports/{batch_export_id}/runs/{id}/logs/ | 
[**BatchExportsRunsRetrieve**](BatchExportsAPI.md#BatchExportsRunsRetrieve) | **Get** /api/projects/{project_id}/batch_exports/{batch_export_id}/runs/{id}/ | 
[**BatchExportsRunsRetryCreate**](BatchExportsAPI.md#BatchExportsRunsRetryCreate) | **Post** /api/projects/{project_id}/batch_exports/{batch_export_id}/runs/{id}/retry/ | 
[**BatchExportsTestRetrieve**](BatchExportsAPI.md#BatchExportsTestRetrieve) | **Get** /api/organizations/{organization_id}/batch_exports/test/ | 
[**BatchExportsTestRetrieve2**](BatchExportsAPI.md#BatchExportsTestRetrieve2) | **Get** /api/projects/{project_id}/batch_exports/test/ | 
[**BatchExportsUnpauseCreate**](BatchExportsAPI.md#BatchExportsUnpauseCreate) | **Post** /api/organizations/{organization_id}/batch_exports/{id}/unpause/ | 
[**BatchExportsUnpauseCreate2**](BatchExportsAPI.md#BatchExportsUnpauseCreate2) | **Post** /api/projects/{project_id}/batch_exports/{id}/unpause/ | 
[**BatchExportsUpdate**](BatchExportsAPI.md#BatchExportsUpdate) | **Put** /api/organizations/{organization_id}/batch_exports/{id}/ | 
[**BatchExportsUpdate2**](BatchExportsAPI.md#BatchExportsUpdate2) | **Put** /api/projects/{project_id}/batch_exports/{id}/ | 



## BatchExportsBackfillCreate

> BatchExportsBackfillCreate(ctx, id, organizationId).BatchExport(batchExport).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	organizationId := "organizationId_example" // string | 
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsBackfillCreate(context.Background(), id, organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsBackfillCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsBackfillCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

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


## BatchExportsBackfillCreate2

> BatchExportsBackfillCreate2(ctx, id, projectId).BatchExport(batchExport).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsBackfillCreate2(context.Background(), id, projectId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsBackfillCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsBackfillCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

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


## BatchExportsBackfillsCancelCreate

> BatchExportsBackfillsCancelCreate(ctx, batchExportId, id, projectId).BatchExportBackfill(batchExportBackfill).Execute()





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
	batchExportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export backfill.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	batchExportBackfill := *openapiclient.NewBatchExportBackfill("Id_example", "Progress_example", openapiclient.BatchExportBackfillStatusEnum("Cancelled"), time.Now(), time.Now(), int32(123), "BatchExport_example") // BatchExportBackfill | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsBackfillsCancelCreate(context.Background(), batchExportId, id, projectId).BatchExportBackfill(batchExportBackfill).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsBackfillsCancelCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchExportId** | **string** |  | 
**id** | **string** | A UUID string identifying this batch export backfill. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsBackfillsCancelCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchExportBackfill** | [**BatchExportBackfill**](BatchExportBackfill.md) |  | 

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


## BatchExportsBackfillsCreate

> BatchExportBackfill BatchExportsBackfillsCreate(ctx, batchExportId, projectId).BatchExportBackfill(batchExportBackfill).Execute()





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
	batchExportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	batchExportBackfill := *openapiclient.NewBatchExportBackfill("Id_example", "Progress_example", openapiclient.BatchExportBackfillStatusEnum("Cancelled"), time.Now(), time.Now(), int32(123), "BatchExport_example") // BatchExportBackfill | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsBackfillsCreate(context.Background(), batchExportId, projectId).BatchExportBackfill(batchExportBackfill).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsBackfillsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsBackfillsCreate`: BatchExportBackfill
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsBackfillsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchExportId** | **string** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsBackfillsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchExportBackfill** | [**BatchExportBackfill**](BatchExportBackfill.md) |  | 

### Return type

[**BatchExportBackfill**](BatchExportBackfill.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsBackfillsList

> PaginatedBatchExportBackfillList BatchExportsBackfillsList(ctx, batchExportId, projectId).Cursor(cursor).Ordering(ordering).Execute()





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
	batchExportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	cursor := "cursor_example" // string | The pagination cursor value. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsBackfillsList(context.Background(), batchExportId, projectId).Cursor(cursor).Ordering(ordering).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsBackfillsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsBackfillsList`: PaginatedBatchExportBackfillList
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsBackfillsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchExportId** | **string** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsBackfillsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** | The pagination cursor value. | 
 **ordering** | **string** | Which field to use when ordering the results. | 

### Return type

[**PaginatedBatchExportBackfillList**](PaginatedBatchExportBackfillList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsBackfillsRetrieve

> BatchExportBackfill BatchExportsBackfillsRetrieve(ctx, batchExportId, id, projectId).Execute()





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
	batchExportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export backfill.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsBackfillsRetrieve(context.Background(), batchExportId, id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsBackfillsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsBackfillsRetrieve`: BatchExportBackfill
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsBackfillsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchExportId** | **string** |  | 
**id** | **string** | A UUID string identifying this batch export backfill. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsBackfillsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BatchExportBackfill**](BatchExportBackfill.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsCreate

> BatchExport BatchExportsCreate(ctx, organizationId).BatchExport(batchExport).Execute()



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
	organizationId := "organizationId_example" // string | 
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsCreate(context.Background(), organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsCreate`: BatchExport
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

### Return type

[**BatchExport**](BatchExport.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsCreate2

> BatchExport BatchExportsCreate2(ctx, projectId).BatchExport(batchExport).Execute()



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
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsCreate2(context.Background(), projectId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsCreate2`: BatchExport
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsCreate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

### Return type

[**BatchExport**](BatchExport.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsDestroy

> BatchExportsDestroy(ctx, id, organizationId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	organizationId := "organizationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsDestroy(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsDestroyRequest struct via the builder pattern


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


## BatchExportsDestroy2

> BatchExportsDestroy2(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsDestroy2(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsDestroy2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsDestroy2Request struct via the builder pattern


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


## BatchExportsList

> PaginatedBatchExportList BatchExportsList(ctx, organizationId).Limit(limit).Offset(offset).Execute()



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
	organizationId := "organizationId_example" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsList(context.Background(), organizationId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsList`: PaginatedBatchExportList
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedBatchExportList**](PaginatedBatchExportList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsList2

> PaginatedBatchExportList BatchExportsList2(ctx, projectId).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsList2(context.Background(), projectId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsList2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsList2`: PaginatedBatchExportList
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsList2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedBatchExportList**](PaginatedBatchExportList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsLogsRetrieve

> BatchExportsLogsRetrieve(ctx, id, organizationId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	organizationId := "organizationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsLogsRetrieve(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsLogsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsLogsRetrieveRequest struct via the builder pattern


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


## BatchExportsLogsRetrieve2

> BatchExportsLogsRetrieve2(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsLogsRetrieve2(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsLogsRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsLogsRetrieve2Request struct via the builder pattern


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


## BatchExportsPartialUpdate

> BatchExport BatchExportsPartialUpdate(ctx, id, organizationId).PatchedBatchExport(patchedBatchExport).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	organizationId := "organizationId_example" // string | 
	patchedBatchExport := *openapiclient.NewPatchedBatchExport() // PatchedBatchExport |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsPartialUpdate(context.Background(), id, organizationId).PatchedBatchExport(patchedBatchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsPartialUpdate`: BatchExport
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedBatchExport** | [**PatchedBatchExport**](PatchedBatchExport.md) |  | 

### Return type

[**BatchExport**](BatchExport.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsPartialUpdate2

> BatchExport BatchExportsPartialUpdate2(ctx, id, projectId).PatchedBatchExport(patchedBatchExport).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedBatchExport := *openapiclient.NewPatchedBatchExport() // PatchedBatchExport |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsPartialUpdate2(context.Background(), id, projectId).PatchedBatchExport(patchedBatchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsPartialUpdate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsPartialUpdate2`: BatchExport
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsPartialUpdate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsPartialUpdate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedBatchExport** | [**PatchedBatchExport**](PatchedBatchExport.md) |  | 

### Return type

[**BatchExport**](BatchExport.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsPauseCreate

> BatchExportsPauseCreate(ctx, id, organizationId).BatchExport(batchExport).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	organizationId := "organizationId_example" // string | 
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsPauseCreate(context.Background(), id, organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsPauseCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsPauseCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

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


## BatchExportsPauseCreate2

> BatchExportsPauseCreate2(ctx, id, projectId).BatchExport(batchExport).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsPauseCreate2(context.Background(), id, projectId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsPauseCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsPauseCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

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


## BatchExportsRetrieve

> BatchExport BatchExportsRetrieve(ctx, id, organizationId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	organizationId := "organizationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsRetrieve(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsRetrieve`: BatchExport
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BatchExport**](BatchExport.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsRetrieve2

> BatchExport BatchExportsRetrieve2(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsRetrieve2(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsRetrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsRetrieve2`: BatchExport
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsRetrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsRetrieve2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BatchExport**](BatchExport.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsRunTestStepCreate

> BatchExportsRunTestStepCreate(ctx, id, organizationId).BatchExport(batchExport).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	organizationId := "organizationId_example" // string | 
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsRunTestStepCreate(context.Background(), id, organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsRunTestStepCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsRunTestStepCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

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


## BatchExportsRunTestStepCreate2

> BatchExportsRunTestStepCreate2(ctx, id, projectId).BatchExport(batchExport).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsRunTestStepCreate2(context.Background(), id, projectId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsRunTestStepCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsRunTestStepCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

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


## BatchExportsRunTestStepNewCreate

> BatchExportsRunTestStepNewCreate(ctx, organizationId).BatchExport(batchExport).Execute()



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
	organizationId := "organizationId_example" // string | 
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsRunTestStepNewCreate(context.Background(), organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsRunTestStepNewCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsRunTestStepNewCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

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


## BatchExportsRunTestStepNewCreate2

> BatchExportsRunTestStepNewCreate2(ctx, projectId).BatchExport(batchExport).Execute()



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
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsRunTestStepNewCreate2(context.Background(), projectId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsRunTestStepNewCreate2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBatchExportsRunTestStepNewCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

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


## BatchExportsRunsCancelCreate

> BatchExportsRunsCancelCreate(ctx, batchExportId, id, projectId).BatchExportRun(batchExportRun).Execute()





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
	batchExportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export run.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	batchExportRun := *openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example") // BatchExportRun | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsRunsCancelCreate(context.Background(), batchExportId, id, projectId).BatchExportRun(batchExportRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsRunsCancelCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchExportId** | **string** |  | 
**id** | **string** | A UUID string identifying this batch export run. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsRunsCancelCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchExportRun** | [**BatchExportRun**](BatchExportRun.md) |  | 

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


## BatchExportsRunsList

> PaginatedBatchExportRunList BatchExportsRunsList(ctx, batchExportId, projectId).Cursor(cursor).Ordering(ordering).Execute()



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
	batchExportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	cursor := "cursor_example" // string | The pagination cursor value. (optional)
	ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsRunsList(context.Background(), batchExportId, projectId).Cursor(cursor).Ordering(ordering).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsRunsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsRunsList`: PaginatedBatchExportRunList
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsRunsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchExportId** | **string** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsRunsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** | The pagination cursor value. | 
 **ordering** | **string** | Which field to use when ordering the results. | 

### Return type

[**PaginatedBatchExportRunList**](PaginatedBatchExportRunList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsRunsLogsRetrieve

> BatchExportsRunsLogsRetrieve(ctx, batchExportId, id, projectId).Execute()



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
	batchExportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export run.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsRunsLogsRetrieve(context.Background(), batchExportId, id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsRunsLogsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchExportId** | **string** |  | 
**id** | **string** | A UUID string identifying this batch export run. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsRunsLogsRetrieveRequest struct via the builder pattern


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


## BatchExportsRunsRetrieve

> BatchExportRun BatchExportsRunsRetrieve(ctx, batchExportId, id, projectId).Execute()



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
	batchExportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export run.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsRunsRetrieve(context.Background(), batchExportId, id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsRunsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsRunsRetrieve`: BatchExportRun
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsRunsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchExportId** | **string** |  | 
**id** | **string** | A UUID string identifying this batch export run. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsRunsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BatchExportRun**](BatchExportRun.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsRunsRetryCreate

> BatchExportsRunsRetryCreate(ctx, batchExportId, id, projectId).BatchExportRun(batchExportRun).Execute()





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
	batchExportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export run.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	batchExportRun := *openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example") // BatchExportRun | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsRunsRetryCreate(context.Background(), batchExportId, id, projectId).BatchExportRun(batchExportRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsRunsRetryCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchExportId** | **string** |  | 
**id** | **string** | A UUID string identifying this batch export run. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsRunsRetryCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **batchExportRun** | [**BatchExportRun**](BatchExportRun.md) |  | 

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


## BatchExportsTestRetrieve

> BatchExportsTestRetrieve(ctx, organizationId).Execute()



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
	organizationId := "organizationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsTestRetrieve(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsTestRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsTestRetrieveRequest struct via the builder pattern


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


## BatchExportsTestRetrieve2

> BatchExportsTestRetrieve2(ctx, projectId).Execute()



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
	r, err := apiClient.BatchExportsAPI.BatchExportsTestRetrieve2(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsTestRetrieve2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBatchExportsTestRetrieve2Request struct via the builder pattern


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


## BatchExportsUnpauseCreate

> BatchExportsUnpauseCreate(ctx, id, organizationId).BatchExport(batchExport).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	organizationId := "organizationId_example" // string | 
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsUnpauseCreate(context.Background(), id, organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsUnpauseCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsUnpauseCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

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


## BatchExportsUnpauseCreate2

> BatchExportsUnpauseCreate2(ctx, id, projectId).BatchExport(batchExport).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchExportsAPI.BatchExportsUnpauseCreate2(context.Background(), id, projectId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsUnpauseCreate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsUnpauseCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

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


## BatchExportsUpdate

> BatchExport BatchExportsUpdate(ctx, id, organizationId).BatchExport(batchExport).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	organizationId := "organizationId_example" // string | 
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsUpdate(context.Background(), id, organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsUpdate`: BatchExport
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

### Return type

[**BatchExport**](BatchExport.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExportsUpdate2

> BatchExport BatchExportsUpdate2(ctx, id, projectId).BatchExport(batchExport).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this batch export.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	batchExport := *openapiclient.NewBatchExport("Id_example", int32(123), "Name_example", *openapiclient.NewBatchExportDestination(openapiclient.BatchExportDestinationTypeEnum("S3")), openapiclient.IntervalEnum("hour"), time.Now(), time.Now(), []openapiclient.BatchExportRun{*openapiclient.NewBatchExportRun("Id_example", openapiclient.BatchExportRunStatusEnum("Cancelled"), time.Now(), time.Now(), time.Now(), "BatchExport_example")}, interface{}(123)) // BatchExport | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchExportsAPI.BatchExportsUpdate2(context.Background(), id, projectId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchExportsAPI.BatchExportsUpdate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsUpdate2`: BatchExport
	fmt.Fprintf(os.Stdout, "Response from `BatchExportsAPI.BatchExportsUpdate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this batch export. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportsUpdate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchExport** | [**BatchExport**](BatchExport.md) |  | 

### Return type

[**BatchExport**](BatchExport.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

