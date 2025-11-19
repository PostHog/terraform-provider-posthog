# \DashboardsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashboardsCollaboratorsCreate**](DashboardsAPI.md#DashboardsCollaboratorsCreate) | **Post** /api/projects/{project_id}/dashboards/{dashboard_id}/collaborators/ | 
[**DashboardsCollaboratorsDestroy**](DashboardsAPI.md#DashboardsCollaboratorsDestroy) | **Delete** /api/projects/{project_id}/dashboards/{dashboard_id}/collaborators/{user__uuid}/ | 
[**DashboardsCollaboratorsList**](DashboardsAPI.md#DashboardsCollaboratorsList) | **Get** /api/projects/{project_id}/dashboards/{dashboard_id}/collaborators/ | 
[**DashboardsCreate**](DashboardsAPI.md#DashboardsCreate) | **Post** /api/projects/{project_id}/dashboards/ | 
[**DashboardsCreateFromTemplateJsonCreate**](DashboardsAPI.md#DashboardsCreateFromTemplateJsonCreate) | **Post** /api/projects/{project_id}/dashboards/create_from_template_json/ | 
[**DashboardsCreateUnlistedDashboardCreate**](DashboardsAPI.md#DashboardsCreateUnlistedDashboardCreate) | **Post** /api/projects/{project_id}/dashboards/create_unlisted_dashboard/ | 
[**DashboardsDestroy**](DashboardsAPI.md#DashboardsDestroy) | **Delete** /api/projects/{project_id}/dashboards/{id}/ | 
[**DashboardsList**](DashboardsAPI.md#DashboardsList) | **Get** /api/projects/{project_id}/dashboards/ | 
[**DashboardsMoveTilePartialUpdate**](DashboardsAPI.md#DashboardsMoveTilePartialUpdate) | **Patch** /api/projects/{project_id}/dashboards/{id}/move_tile/ | 
[**DashboardsPartialUpdate**](DashboardsAPI.md#DashboardsPartialUpdate) | **Patch** /api/projects/{project_id}/dashboards/{id}/ | 
[**DashboardsRetrieve**](DashboardsAPI.md#DashboardsRetrieve) | **Get** /api/projects/{project_id}/dashboards/{id}/ | 
[**DashboardsSharingList**](DashboardsAPI.md#DashboardsSharingList) | **Get** /api/projects/{project_id}/dashboards/{dashboard_id}/sharing/ | 
[**DashboardsSharingPasswordsCreate**](DashboardsAPI.md#DashboardsSharingPasswordsCreate) | **Post** /api/projects/{project_id}/dashboards/{dashboard_id}/sharing/passwords/ | 
[**DashboardsSharingPasswordsDestroy**](DashboardsAPI.md#DashboardsSharingPasswordsDestroy) | **Delete** /api/projects/{project_id}/dashboards/{dashboard_id}/sharing/passwords/{password_id}/ | 
[**DashboardsSharingRefreshCreate**](DashboardsAPI.md#DashboardsSharingRefreshCreate) | **Post** /api/projects/{project_id}/dashboards/{dashboard_id}/sharing/refresh/ | 
[**DashboardsStreamTilesRetrieve**](DashboardsAPI.md#DashboardsStreamTilesRetrieve) | **Get** /api/projects/{project_id}/dashboards/{id}/stream_tiles/ | 
[**DashboardsUpdate**](DashboardsAPI.md#DashboardsUpdate) | **Put** /api/projects/{project_id}/dashboards/{id}/ | 



## DashboardsCollaboratorsCreate

> DashboardCollaborator DashboardsCollaboratorsCreate(ctx, dashboardId, projectId).DashboardCollaborator(dashboardCollaborator).Execute()



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
	dashboardId := int32(56) // int32 | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	dashboardCollaborator := *openapiclient.NewDashboardCollaborator("Id_example", int32(123), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), openapiclient.DashboardRestrictionLevel(21), time.Now(), time.Now(), "UserUuid_example") // DashboardCollaborator | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.DashboardsCollaboratorsCreate(context.Background(), dashboardId, projectId).DashboardCollaborator(dashboardCollaborator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsCollaboratorsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsCollaboratorsCreate`: DashboardCollaborator
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.DashboardsCollaboratorsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int32** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsCollaboratorsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dashboardCollaborator** | [**DashboardCollaborator**](DashboardCollaborator.md) |  | 

### Return type

[**DashboardCollaborator**](DashboardCollaborator.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsCollaboratorsDestroy

> DashboardsCollaboratorsDestroy(ctx, dashboardId, projectId, userUuid).Execute()



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
	dashboardId := int32(56) // int32 | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DashboardsAPI.DashboardsCollaboratorsDestroy(context.Background(), dashboardId, projectId, userUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsCollaboratorsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int32** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**userUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsCollaboratorsDestroyRequest struct via the builder pattern


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


## DashboardsCollaboratorsList

> []DashboardCollaborator DashboardsCollaboratorsList(ctx, dashboardId, projectId).Execute()



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
	dashboardId := int32(56) // int32 | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.DashboardsCollaboratorsList(context.Background(), dashboardId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsCollaboratorsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsCollaboratorsList`: []DashboardCollaborator
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.DashboardsCollaboratorsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int32** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsCollaboratorsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]DashboardCollaborator**](DashboardCollaborator.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsCreate

> Dashboard DashboardsCreate(ctx, projectId).Format(format).Dashboard(dashboard).Execute()



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
	format := "format_example" // string |  (optional)
	dashboard := *openapiclient.NewDashboard(int32(123), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), false, openapiclient.CreationModeEnum("default"), map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, openapiclient.EffectiveRestrictionLevelEnum(21), openapiclient.EffectivePrivilegeLevelEnum(21), "UserAccessLevel_example", "AccessControlVersion_example", map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, int32(123), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // Dashboard |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.DashboardsCreate(context.Background(), projectId).Format(format).Dashboard(dashboard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsCreate`: Dashboard
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.DashboardsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** |  | 
 **dashboard** | [**Dashboard**](Dashboard.md) |  | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsCreateFromTemplateJsonCreate

> DashboardsCreateFromTemplateJsonCreate(ctx, projectId).Format(format).Dashboard(dashboard).Execute()



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
	format := "format_example" // string |  (optional)
	dashboard := *openapiclient.NewDashboard(int32(123), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), false, openapiclient.CreationModeEnum("default"), map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, openapiclient.EffectiveRestrictionLevelEnum(21), openapiclient.EffectivePrivilegeLevelEnum(21), "UserAccessLevel_example", "AccessControlVersion_example", map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, int32(123), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // Dashboard |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DashboardsAPI.DashboardsCreateFromTemplateJsonCreate(context.Background(), projectId).Format(format).Dashboard(dashboard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsCreateFromTemplateJsonCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDashboardsCreateFromTemplateJsonCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** |  | 
 **dashboard** | [**Dashboard**](Dashboard.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsCreateUnlistedDashboardCreate

> DashboardsCreateUnlistedDashboardCreate(ctx, projectId).Format(format).Dashboard(dashboard).Execute()





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
	format := "format_example" // string |  (optional)
	dashboard := *openapiclient.NewDashboard(int32(123), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), false, openapiclient.CreationModeEnum("default"), map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, openapiclient.EffectiveRestrictionLevelEnum(21), openapiclient.EffectivePrivilegeLevelEnum(21), "UserAccessLevel_example", "AccessControlVersion_example", map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, int32(123), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // Dashboard |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DashboardsAPI.DashboardsCreateUnlistedDashboardCreate(context.Background(), projectId).Format(format).Dashboard(dashboard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsCreateUnlistedDashboardCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDashboardsCreateUnlistedDashboardCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** |  | 
 **dashboard** | [**Dashboard**](Dashboard.md) |  | 

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


## DashboardsDestroy

> DashboardsDestroy(ctx, id, projectId).Format(format).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this dashboard.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DashboardsAPI.DashboardsDestroy(context.Background(), id, projectId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dashboard. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 

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


## DashboardsList

> PaginatedDashboardBasicList DashboardsList(ctx, projectId).Format(format).Limit(limit).Offset(offset).Execute()



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
	format := "format_example" // string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.DashboardsList(context.Background(), projectId).Format(format).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsList`: PaginatedDashboardBasicList
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.DashboardsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedDashboardBasicList**](PaginatedDashboardBasicList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsMoveTilePartialUpdate

> DashboardsMoveTilePartialUpdate(ctx, id, projectId).Format(format).PatchedDashboard(patchedDashboard).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dashboard.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)
	patchedDashboard := *openapiclient.NewPatchedDashboard() // PatchedDashboard |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DashboardsAPI.DashboardsMoveTilePartialUpdate(context.Background(), id, projectId).Format(format).PatchedDashboard(patchedDashboard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsMoveTilePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dashboard. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsMoveTilePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 
 **patchedDashboard** | [**PatchedDashboard**](PatchedDashboard.md) |  | 

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


## DashboardsPartialUpdate

> Dashboard DashboardsPartialUpdate(ctx, id, projectId).Format(format).PatchedDashboard(patchedDashboard).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dashboard.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)
	patchedDashboard := *openapiclient.NewPatchedDashboard() // PatchedDashboard |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.DashboardsPartialUpdate(context.Background(), id, projectId).Format(format).PatchedDashboard(patchedDashboard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsPartialUpdate`: Dashboard
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.DashboardsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dashboard. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 
 **patchedDashboard** | [**PatchedDashboard**](PatchedDashboard.md) |  | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsRetrieve

> Dashboard DashboardsRetrieve(ctx, id, projectId).Format(format).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dashboard.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.DashboardsRetrieve(context.Background(), id, projectId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsRetrieve`: Dashboard
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.DashboardsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dashboard. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsSharingList

> []SharingConfiguration DashboardsSharingList(ctx, dashboardId, projectId).Execute()



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
	dashboardId := int32(56) // int32 | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.DashboardsSharingList(context.Background(), dashboardId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsSharingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsSharingList`: []SharingConfiguration
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.DashboardsSharingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int32** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsSharingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]SharingConfiguration**](SharingConfiguration.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsSharingPasswordsCreate

> SharingConfiguration DashboardsSharingPasswordsCreate(ctx, dashboardId, projectId).SharingConfiguration(sharingConfiguration).Execute()





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
	dashboardId := int32(56) // int32 | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	sharingConfiguration := *openapiclient.NewSharingConfiguration(time.Now(), "AccessToken_example", "SharePasswords_example") // SharingConfiguration |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.DashboardsSharingPasswordsCreate(context.Background(), dashboardId, projectId).SharingConfiguration(sharingConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsSharingPasswordsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsSharingPasswordsCreate`: SharingConfiguration
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.DashboardsSharingPasswordsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int32** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsSharingPasswordsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sharingConfiguration** | [**SharingConfiguration**](SharingConfiguration.md) |  | 

### Return type

[**SharingConfiguration**](SharingConfiguration.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsSharingPasswordsDestroy

> DashboardsSharingPasswordsDestroy(ctx, dashboardId, passwordId, projectId).Execute()





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
	dashboardId := int32(56) // int32 | 
	passwordId := "passwordId_example" // string | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DashboardsAPI.DashboardsSharingPasswordsDestroy(context.Background(), dashboardId, passwordId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsSharingPasswordsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int32** |  | 
**passwordId** | **string** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsSharingPasswordsDestroyRequest struct via the builder pattern


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


## DashboardsSharingRefreshCreate

> SharingConfiguration DashboardsSharingRefreshCreate(ctx, dashboardId, projectId).SharingConfiguration(sharingConfiguration).Execute()



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
	dashboardId := int32(56) // int32 | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	sharingConfiguration := *openapiclient.NewSharingConfiguration(time.Now(), "AccessToken_example", "SharePasswords_example") // SharingConfiguration |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.DashboardsSharingRefreshCreate(context.Background(), dashboardId, projectId).SharingConfiguration(sharingConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsSharingRefreshCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsSharingRefreshCreate`: SharingConfiguration
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.DashboardsSharingRefreshCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int32** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsSharingRefreshCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sharingConfiguration** | [**SharingConfiguration**](SharingConfiguration.md) |  | 

### Return type

[**SharingConfiguration**](SharingConfiguration.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardsStreamTilesRetrieve

> DashboardsStreamTilesRetrieve(ctx, id, projectId).Format(format).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this dashboard.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DashboardsAPI.DashboardsStreamTilesRetrieve(context.Background(), id, projectId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsStreamTilesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dashboard. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsStreamTilesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 

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


## DashboardsUpdate

> Dashboard DashboardsUpdate(ctx, id, projectId).Format(format).Dashboard(dashboard).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dashboard.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)
	dashboard := *openapiclient.NewDashboard(int32(123), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), false, openapiclient.CreationModeEnum("default"), map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, openapiclient.EffectiveRestrictionLevelEnum(21), openapiclient.EffectivePrivilegeLevelEnum(21), "UserAccessLevel_example", "AccessControlVersion_example", map[string]interface{}{"key": interface{}(123)}, map[string]interface{}{"key": interface{}(123)}, int32(123), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // Dashboard |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.DashboardsUpdate(context.Background(), id, projectId).Format(format).Dashboard(dashboard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DashboardsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DashboardsUpdate`: Dashboard
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.DashboardsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dashboard. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDashboardsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 
 **dashboard** | [**Dashboard**](Dashboard.md) |  | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

