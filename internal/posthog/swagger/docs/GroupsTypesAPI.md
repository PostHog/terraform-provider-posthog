# \GroupsTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsTypesCreateDetailDashboardUpdate**](GroupsTypesAPI.md#GroupsTypesCreateDetailDashboardUpdate) | **Put** /api/projects/{project_id}/groups_types/create_detail_dashboard/ | 
[**GroupsTypesDestroy**](GroupsTypesAPI.md#GroupsTypesDestroy) | **Delete** /api/projects/{project_id}/groups_types/{group_type_index}/ | 
[**GroupsTypesList**](GroupsTypesAPI.md#GroupsTypesList) | **Get** /api/projects/{project_id}/groups_types/ | 
[**GroupsTypesMetricsCreate**](GroupsTypesAPI.md#GroupsTypesMetricsCreate) | **Post** /api/projects/{project_id}/groups_types/{group_type_index}/metrics/ | 
[**GroupsTypesMetricsDestroy**](GroupsTypesAPI.md#GroupsTypesMetricsDestroy) | **Delete** /api/projects/{project_id}/groups_types/{group_type_index}/metrics/{id}/ | 
[**GroupsTypesMetricsList**](GroupsTypesAPI.md#GroupsTypesMetricsList) | **Get** /api/projects/{project_id}/groups_types/{group_type_index}/metrics/ | 
[**GroupsTypesMetricsPartialUpdate**](GroupsTypesAPI.md#GroupsTypesMetricsPartialUpdate) | **Patch** /api/projects/{project_id}/groups_types/{group_type_index}/metrics/{id}/ | 
[**GroupsTypesMetricsRetrieve**](GroupsTypesAPI.md#GroupsTypesMetricsRetrieve) | **Get** /api/projects/{project_id}/groups_types/{group_type_index}/metrics/{id}/ | 
[**GroupsTypesMetricsUpdate**](GroupsTypesAPI.md#GroupsTypesMetricsUpdate) | **Put** /api/projects/{project_id}/groups_types/{group_type_index}/metrics/{id}/ | 
[**GroupsTypesSetDefaultColumnsUpdate**](GroupsTypesAPI.md#GroupsTypesSetDefaultColumnsUpdate) | **Put** /api/projects/{project_id}/groups_types/set_default_columns/ | 
[**GroupsTypesUpdateMetadataPartialUpdate**](GroupsTypesAPI.md#GroupsTypesUpdateMetadataPartialUpdate) | **Patch** /api/projects/{project_id}/groups_types/update_metadata/ | 



## GroupsTypesCreateDetailDashboardUpdate

> GroupsTypesCreateDetailDashboardUpdate(ctx, projectId).GroupType(groupType).Execute()



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
	groupType := *openapiclient.NewGroupType("GroupType_example", int32(123)) // GroupType |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsTypesAPI.GroupsTypesCreateDetailDashboardUpdate(context.Background(), projectId).GroupType(groupType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsTypesAPI.GroupsTypesCreateDetailDashboardUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGroupsTypesCreateDetailDashboardUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupType** | [**GroupType**](GroupType.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsTypesDestroy

> GroupsTypesDestroy(ctx, groupTypeIndex, projectId).Execute()



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
	groupTypeIndex := int32(56) // int32 | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsTypesAPI.GroupsTypesDestroy(context.Background(), groupTypeIndex, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsTypesAPI.GroupsTypesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupTypeIndex** | **int32** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsTypesDestroyRequest struct via the builder pattern


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


## GroupsTypesList

> []GroupType GroupsTypesList(ctx, projectId).Execute()



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
	resp, r, err := apiClient.GroupsTypesAPI.GroupsTypesList(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsTypesAPI.GroupsTypesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsTypesList`: []GroupType
	fmt.Fprintf(os.Stdout, "Response from `GroupsTypesAPI.GroupsTypesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsTypesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GroupType**](GroupType.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsTypesMetricsCreate

> GroupUsageMetric GroupsTypesMetricsCreate(ctx, groupTypeIndex, projectId).GroupUsageMetric(groupUsageMetric).Execute()



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
	groupTypeIndex := int32(56) // int32 | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	groupUsageMetric := *openapiclient.NewGroupUsageMetric("Id_example", "Name_example", interface{}(123)) // GroupUsageMetric | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsTypesAPI.GroupsTypesMetricsCreate(context.Background(), groupTypeIndex, projectId).GroupUsageMetric(groupUsageMetric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsTypesAPI.GroupsTypesMetricsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsTypesMetricsCreate`: GroupUsageMetric
	fmt.Fprintf(os.Stdout, "Response from `GroupsTypesAPI.GroupsTypesMetricsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupTypeIndex** | **int32** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsTypesMetricsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupUsageMetric** | [**GroupUsageMetric**](GroupUsageMetric.md) |  | 

### Return type

[**GroupUsageMetric**](GroupUsageMetric.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsTypesMetricsDestroy

> GroupsTypesMetricsDestroy(ctx, groupTypeIndex, id, projectId).Execute()



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
	groupTypeIndex := int32(56) // int32 | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this group usage metric.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsTypesAPI.GroupsTypesMetricsDestroy(context.Background(), groupTypeIndex, id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsTypesAPI.GroupsTypesMetricsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupTypeIndex** | **int32** |  | 
**id** | **string** | A UUID string identifying this group usage metric. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsTypesMetricsDestroyRequest struct via the builder pattern


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


## GroupsTypesMetricsList

> PaginatedGroupUsageMetricList GroupsTypesMetricsList(ctx, groupTypeIndex, projectId).Limit(limit).Offset(offset).Execute()



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
	groupTypeIndex := int32(56) // int32 | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsTypesAPI.GroupsTypesMetricsList(context.Background(), groupTypeIndex, projectId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsTypesAPI.GroupsTypesMetricsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsTypesMetricsList`: PaginatedGroupUsageMetricList
	fmt.Fprintf(os.Stdout, "Response from `GroupsTypesAPI.GroupsTypesMetricsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupTypeIndex** | **int32** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsTypesMetricsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedGroupUsageMetricList**](PaginatedGroupUsageMetricList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsTypesMetricsPartialUpdate

> GroupUsageMetric GroupsTypesMetricsPartialUpdate(ctx, groupTypeIndex, id, projectId).PatchedGroupUsageMetric(patchedGroupUsageMetric).Execute()



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
	groupTypeIndex := int32(56) // int32 | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this group usage metric.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedGroupUsageMetric := *openapiclient.NewPatchedGroupUsageMetric() // PatchedGroupUsageMetric |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsTypesAPI.GroupsTypesMetricsPartialUpdate(context.Background(), groupTypeIndex, id, projectId).PatchedGroupUsageMetric(patchedGroupUsageMetric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsTypesAPI.GroupsTypesMetricsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsTypesMetricsPartialUpdate`: GroupUsageMetric
	fmt.Fprintf(os.Stdout, "Response from `GroupsTypesAPI.GroupsTypesMetricsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupTypeIndex** | **int32** |  | 
**id** | **string** | A UUID string identifying this group usage metric. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsTypesMetricsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchedGroupUsageMetric** | [**PatchedGroupUsageMetric**](PatchedGroupUsageMetric.md) |  | 

### Return type

[**GroupUsageMetric**](GroupUsageMetric.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsTypesMetricsRetrieve

> GroupUsageMetric GroupsTypesMetricsRetrieve(ctx, groupTypeIndex, id, projectId).Execute()



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
	groupTypeIndex := int32(56) // int32 | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this group usage metric.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsTypesAPI.GroupsTypesMetricsRetrieve(context.Background(), groupTypeIndex, id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsTypesAPI.GroupsTypesMetricsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsTypesMetricsRetrieve`: GroupUsageMetric
	fmt.Fprintf(os.Stdout, "Response from `GroupsTypesAPI.GroupsTypesMetricsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupTypeIndex** | **int32** |  | 
**id** | **string** | A UUID string identifying this group usage metric. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsTypesMetricsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GroupUsageMetric**](GroupUsageMetric.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsTypesMetricsUpdate

> GroupUsageMetric GroupsTypesMetricsUpdate(ctx, groupTypeIndex, id, projectId).GroupUsageMetric(groupUsageMetric).Execute()



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
	groupTypeIndex := int32(56) // int32 | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this group usage metric.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	groupUsageMetric := *openapiclient.NewGroupUsageMetric("Id_example", "Name_example", interface{}(123)) // GroupUsageMetric | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsTypesAPI.GroupsTypesMetricsUpdate(context.Background(), groupTypeIndex, id, projectId).GroupUsageMetric(groupUsageMetric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsTypesAPI.GroupsTypesMetricsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsTypesMetricsUpdate`: GroupUsageMetric
	fmt.Fprintf(os.Stdout, "Response from `GroupsTypesAPI.GroupsTypesMetricsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupTypeIndex** | **int32** |  | 
**id** | **string** | A UUID string identifying this group usage metric. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsTypesMetricsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **groupUsageMetric** | [**GroupUsageMetric**](GroupUsageMetric.md) |  | 

### Return type

[**GroupUsageMetric**](GroupUsageMetric.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsTypesSetDefaultColumnsUpdate

> GroupsTypesSetDefaultColumnsUpdate(ctx, projectId).GroupType(groupType).Execute()



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
	groupType := *openapiclient.NewGroupType("GroupType_example", int32(123)) // GroupType |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsTypesAPI.GroupsTypesSetDefaultColumnsUpdate(context.Background(), projectId).GroupType(groupType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsTypesAPI.GroupsTypesSetDefaultColumnsUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGroupsTypesSetDefaultColumnsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupType** | [**GroupType**](GroupType.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsTypesUpdateMetadataPartialUpdate

> GroupsTypesUpdateMetadataPartialUpdate(ctx, projectId).PatchedGroupType(patchedGroupType).Execute()



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
	patchedGroupType := *openapiclient.NewPatchedGroupType() // PatchedGroupType |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsTypesAPI.GroupsTypesUpdateMetadataPartialUpdate(context.Background(), projectId).PatchedGroupType(patchedGroupType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsTypesAPI.GroupsTypesUpdateMetadataPartialUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGroupsTypesUpdateMetadataPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedGroupType** | [**PatchedGroupType**](PatchedGroupType.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

