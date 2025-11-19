# \OrganizationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivityRetrieve**](OrganizationsAPI.md#ActivityRetrieve) | **Get** /api/organizations/{organization_id}/projects/{id}/activity/ | 
[**AddProductIntentPartialUpdate**](OrganizationsAPI.md#AddProductIntentPartialUpdate) | **Patch** /api/organizations/{organization_id}/projects/{id}/add_product_intent/ | 
[**BatchExportsBackfillCreate**](OrganizationsAPI.md#BatchExportsBackfillCreate) | **Post** /api/organizations/{organization_id}/batch_exports/{id}/backfill/ | 
[**BatchExportsCreate**](OrganizationsAPI.md#BatchExportsCreate) | **Post** /api/organizations/{organization_id}/batch_exports/ | 
[**BatchExportsDestroy**](OrganizationsAPI.md#BatchExportsDestroy) | **Delete** /api/organizations/{organization_id}/batch_exports/{id}/ | 
[**BatchExportsList**](OrganizationsAPI.md#BatchExportsList) | **Get** /api/organizations/{organization_id}/batch_exports/ | 
[**BatchExportsLogsRetrieve**](OrganizationsAPI.md#BatchExportsLogsRetrieve) | **Get** /api/organizations/{organization_id}/batch_exports/{id}/logs/ | 
[**BatchExportsPartialUpdate**](OrganizationsAPI.md#BatchExportsPartialUpdate) | **Patch** /api/organizations/{organization_id}/batch_exports/{id}/ | 
[**BatchExportsPauseCreate**](OrganizationsAPI.md#BatchExportsPauseCreate) | **Post** /api/organizations/{organization_id}/batch_exports/{id}/pause/ | 
[**BatchExportsRetrieve**](OrganizationsAPI.md#BatchExportsRetrieve) | **Get** /api/organizations/{organization_id}/batch_exports/{id}/ | 
[**BatchExportsRunTestStepCreate**](OrganizationsAPI.md#BatchExportsRunTestStepCreate) | **Post** /api/organizations/{organization_id}/batch_exports/{id}/run_test_step/ | 
[**BatchExportsRunTestStepNewCreate**](OrganizationsAPI.md#BatchExportsRunTestStepNewCreate) | **Post** /api/organizations/{organization_id}/batch_exports/run_test_step_new/ | 
[**BatchExportsTestRetrieve**](OrganizationsAPI.md#BatchExportsTestRetrieve) | **Get** /api/organizations/{organization_id}/batch_exports/test/ | 
[**BatchExportsUnpauseCreate**](OrganizationsAPI.md#BatchExportsUnpauseCreate) | **Post** /api/organizations/{organization_id}/batch_exports/{id}/unpause/ | 
[**BatchExportsUpdate**](OrganizationsAPI.md#BatchExportsUpdate) | **Put** /api/organizations/{organization_id}/batch_exports/{id}/ | 
[**ChangeOrganizationCreate**](OrganizationsAPI.md#ChangeOrganizationCreate) | **Post** /api/organizations/{organization_id}/projects/{id}/change_organization/ | 
[**CompleteProductOnboardingPartialUpdate**](OrganizationsAPI.md#CompleteProductOnboardingPartialUpdate) | **Patch** /api/organizations/{organization_id}/projects/{id}/complete_product_onboarding/ | 
[**Create**](OrganizationsAPI.md#Create) | **Post** /api/organizations/ | 
[**Create2**](OrganizationsAPI.md#Create2) | **Post** /api/organizations/{organization_id}/projects/ | 
[**Create_0**](OrganizationsAPI.md#Create_0) | **Post** /api/organizations/ | 
[**DeleteSecretTokenBackupPartialUpdate**](OrganizationsAPI.md#DeleteSecretTokenBackupPartialUpdate) | **Patch** /api/organizations/{organization_id}/projects/{id}/delete_secret_token_backup/ | 
[**Destroy**](OrganizationsAPI.md#Destroy) | **Delete** /api/organizations/{id}/ | 
[**Destroy2**](OrganizationsAPI.md#Destroy2) | **Delete** /api/organizations/{organization_id}/projects/{id}/ | 
[**Destroy_0**](OrganizationsAPI.md#Destroy_0) | **Delete** /api/organizations/{id}/ | 
[**DomainsCreate**](OrganizationsAPI.md#DomainsCreate) | **Post** /api/organizations/{organization_id}/domains/ | 
[**DomainsDestroy**](OrganizationsAPI.md#DomainsDestroy) | **Delete** /api/organizations/{organization_id}/domains/{id}/ | 
[**DomainsList**](OrganizationsAPI.md#DomainsList) | **Get** /api/organizations/{organization_id}/domains/ | 
[**DomainsPartialUpdate**](OrganizationsAPI.md#DomainsPartialUpdate) | **Patch** /api/organizations/{organization_id}/domains/{id}/ | 
[**DomainsRetrieve**](OrganizationsAPI.md#DomainsRetrieve) | **Get** /api/organizations/{organization_id}/domains/{id}/ | 
[**DomainsScimTokenCreate**](OrganizationsAPI.md#DomainsScimTokenCreate) | **Post** /api/organizations/{organization_id}/domains/{id}/scim/token/ | 
[**DomainsUpdate**](OrganizationsAPI.md#DomainsUpdate) | **Put** /api/organizations/{organization_id}/domains/{id}/ | 
[**DomainsVerifyCreate**](OrganizationsAPI.md#DomainsVerifyCreate) | **Post** /api/organizations/{organization_id}/domains/{id}/verify/ | 
[**EnvironmentsRollbackCreate**](OrganizationsAPI.md#EnvironmentsRollbackCreate) | **Post** /api/organizations/{id}/environments_rollback/ | 
[**InvitesBulkCreate**](OrganizationsAPI.md#InvitesBulkCreate) | **Post** /api/organizations/{organization_id}/invites/bulk/ | 
[**InvitesCreate**](OrganizationsAPI.md#InvitesCreate) | **Post** /api/organizations/{organization_id}/invites/ | 
[**InvitesDestroy**](OrganizationsAPI.md#InvitesDestroy) | **Delete** /api/organizations/{organization_id}/invites/{id}/ | 
[**InvitesList**](OrganizationsAPI.md#InvitesList) | **Get** /api/organizations/{organization_id}/invites/ | 
[**IsGeneratingDemoDataRetrieve**](OrganizationsAPI.md#IsGeneratingDemoDataRetrieve) | **Get** /api/organizations/{organization_id}/projects/{id}/is_generating_demo_data/ | 
[**List**](OrganizationsAPI.md#List) | **Get** /api/organizations/ | 
[**List2**](OrganizationsAPI.md#List2) | **Get** /api/organizations/{organization_id}/projects/ | 
[**List_0**](OrganizationsAPI.md#List_0) | **Get** /api/organizations/ | 
[**MembersDestroy**](OrganizationsAPI.md#MembersDestroy) | **Delete** /api/organizations/{organization_id}/members/{user__uuid}/ | 
[**MembersList**](OrganizationsAPI.md#MembersList) | **Get** /api/organizations/{organization_id}/members/ | 
[**MembersPartialUpdate**](OrganizationsAPI.md#MembersPartialUpdate) | **Patch** /api/organizations/{organization_id}/members/{user__uuid}/ | 
[**MembersScopedApiKeysRetrieve**](OrganizationsAPI.md#MembersScopedApiKeysRetrieve) | **Get** /api/organizations/{organization_id}/members/{user__uuid}/scoped_api_keys/ | 
[**MembersUpdate**](OrganizationsAPI.md#MembersUpdate) | **Put** /api/organizations/{organization_id}/members/{user__uuid}/ | 
[**PartialUpdate**](OrganizationsAPI.md#PartialUpdate) | **Patch** /api/organizations/{id}/ | 
[**PartialUpdate2**](OrganizationsAPI.md#PartialUpdate2) | **Patch** /api/organizations/{organization_id}/projects/{id}/ | 
[**PartialUpdate_0**](OrganizationsAPI.md#PartialUpdate_0) | **Patch** /api/organizations/{id}/ | 
[**ProxyRecordsCreate**](OrganizationsAPI.md#ProxyRecordsCreate) | **Post** /api/organizations/{organization_id}/proxy_records/ | 
[**ProxyRecordsDestroy**](OrganizationsAPI.md#ProxyRecordsDestroy) | **Delete** /api/organizations/{organization_id}/proxy_records/{id}/ | 
[**ProxyRecordsList**](OrganizationsAPI.md#ProxyRecordsList) | **Get** /api/organizations/{organization_id}/proxy_records/ | 
[**ProxyRecordsPartialUpdate**](OrganizationsAPI.md#ProxyRecordsPartialUpdate) | **Patch** /api/organizations/{organization_id}/proxy_records/{id}/ | 
[**ProxyRecordsRetrieve**](OrganizationsAPI.md#ProxyRecordsRetrieve) | **Get** /api/organizations/{organization_id}/proxy_records/{id}/ | 
[**ProxyRecordsUpdate**](OrganizationsAPI.md#ProxyRecordsUpdate) | **Put** /api/organizations/{organization_id}/proxy_records/{id}/ | 
[**ResetTokenPartialUpdate**](OrganizationsAPI.md#ResetTokenPartialUpdate) | **Patch** /api/organizations/{organization_id}/projects/{id}/reset_token/ | 
[**Retrieve**](OrganizationsAPI.md#Retrieve) | **Get** /api/organizations/{id}/ | 
[**Retrieve2**](OrganizationsAPI.md#Retrieve2) | **Get** /api/organizations/{organization_id}/projects/{id}/ | 
[**Retrieve_0**](OrganizationsAPI.md#Retrieve_0) | **Get** /api/organizations/{id}/ | 
[**RolesCreate**](OrganizationsAPI.md#RolesCreate) | **Post** /api/organizations/{organization_id}/roles/ | 
[**RolesDestroy**](OrganizationsAPI.md#RolesDestroy) | **Delete** /api/organizations/{organization_id}/roles/{id}/ | 
[**RolesList**](OrganizationsAPI.md#RolesList) | **Get** /api/organizations/{organization_id}/roles/ | 
[**RolesPartialUpdate**](OrganizationsAPI.md#RolesPartialUpdate) | **Patch** /api/organizations/{organization_id}/roles/{id}/ | 
[**RolesRetrieve**](OrganizationsAPI.md#RolesRetrieve) | **Get** /api/organizations/{organization_id}/roles/{id}/ | 
[**RolesRoleMembershipsCreate**](OrganizationsAPI.md#RolesRoleMembershipsCreate) | **Post** /api/organizations/{organization_id}/roles/{role_id}/role_memberships/ | 
[**RolesRoleMembershipsDestroy**](OrganizationsAPI.md#RolesRoleMembershipsDestroy) | **Delete** /api/organizations/{organization_id}/roles/{role_id}/role_memberships/{id}/ | 
[**RolesRoleMembershipsList**](OrganizationsAPI.md#RolesRoleMembershipsList) | **Get** /api/organizations/{organization_id}/roles/{role_id}/role_memberships/ | 
[**RolesUpdate**](OrganizationsAPI.md#RolesUpdate) | **Put** /api/organizations/{organization_id}/roles/{id}/ | 
[**RotateSecretTokenPartialUpdate**](OrganizationsAPI.md#RotateSecretTokenPartialUpdate) | **Patch** /api/organizations/{organization_id}/projects/{id}/rotate_secret_token/ | 
[**Update**](OrganizationsAPI.md#Update) | **Put** /api/organizations/{id}/ | 
[**Update2**](OrganizationsAPI.md#Update2) | **Put** /api/organizations/{organization_id}/projects/{id}/ | 
[**Update_0**](OrganizationsAPI.md#Update_0) | **Put** /api/organizations/{id}/ | 



## ActivityRetrieve

> ProjectBackwardCompat ActivityRetrieve(ctx, id, organizationId).Execute()





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
	id := int64(789) // int64 | A unique value identifying this project.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ActivityRetrieve(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ActivityRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivityRetrieve`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ActivityRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | A unique value identifying this project. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivityRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectBackwardCompat**](ProjectBackwardCompat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddProductIntentPartialUpdate

> ProjectBackwardCompat AddProductIntentPartialUpdate(ctx, id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()





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
	id := int64(789) // int64 | A unique value identifying this project.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchedProjectBackwardCompat := *openapiclient.NewPatchedProjectBackwardCompat() // PatchedProjectBackwardCompat |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AddProductIntentPartialUpdate(context.Background(), id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AddProductIntentPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddProductIntentPartialUpdate`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AddProductIntentPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | A unique value identifying this project. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProductIntentPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedProjectBackwardCompat** | [**PatchedProjectBackwardCompat**](PatchedProjectBackwardCompat.md) |  | 

### Return type

[**ProjectBackwardCompat**](ProjectBackwardCompat.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	r, err := apiClient.OrganizationsAPI.BatchExportsBackfillCreate(context.Background(), id, organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.BatchExportsBackfillCreate``: %v\n", err)
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
	resp, r, err := apiClient.OrganizationsAPI.BatchExportsCreate(context.Background(), organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.BatchExportsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsCreate`: BatchExport
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.BatchExportsCreate`: %v\n", resp)
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
	r, err := apiClient.OrganizationsAPI.BatchExportsDestroy(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.BatchExportsDestroy``: %v\n", err)
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
	resp, r, err := apiClient.OrganizationsAPI.BatchExportsList(context.Background(), organizationId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.BatchExportsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsList`: PaginatedBatchExportList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.BatchExportsList`: %v\n", resp)
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
	r, err := apiClient.OrganizationsAPI.BatchExportsLogsRetrieve(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.BatchExportsLogsRetrieve``: %v\n", err)
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
	resp, r, err := apiClient.OrganizationsAPI.BatchExportsPartialUpdate(context.Background(), id, organizationId).PatchedBatchExport(patchedBatchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.BatchExportsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsPartialUpdate`: BatchExport
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.BatchExportsPartialUpdate`: %v\n", resp)
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
	r, err := apiClient.OrganizationsAPI.BatchExportsPauseCreate(context.Background(), id, organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.BatchExportsPauseCreate``: %v\n", err)
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
	resp, r, err := apiClient.OrganizationsAPI.BatchExportsRetrieve(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.BatchExportsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsRetrieve`: BatchExport
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.BatchExportsRetrieve`: %v\n", resp)
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
	r, err := apiClient.OrganizationsAPI.BatchExportsRunTestStepCreate(context.Background(), id, organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.BatchExportsRunTestStepCreate``: %v\n", err)
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
	r, err := apiClient.OrganizationsAPI.BatchExportsRunTestStepNewCreate(context.Background(), organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.BatchExportsRunTestStepNewCreate``: %v\n", err)
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
	r, err := apiClient.OrganizationsAPI.BatchExportsTestRetrieve(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.BatchExportsTestRetrieve``: %v\n", err)
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
	r, err := apiClient.OrganizationsAPI.BatchExportsUnpauseCreate(context.Background(), id, organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.BatchExportsUnpauseCreate``: %v\n", err)
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
	resp, r, err := apiClient.OrganizationsAPI.BatchExportsUpdate(context.Background(), id, organizationId).BatchExport(batchExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.BatchExportsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportsUpdate`: BatchExport
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.BatchExportsUpdate`: %v\n", resp)
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


## ChangeOrganizationCreate

> ProjectBackwardCompat ChangeOrganizationCreate(ctx, id, organizationId).ProjectBackwardCompat(projectBackwardCompat).Execute()





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
	id := int64(789) // int64 | A unique value identifying this project.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectBackwardCompat := *openapiclient.NewProjectBackwardCompat(int32(123), "Organization_example", time.Now(), "TODO", false, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, "LiveEventsToken_example", time.Now(), "Uuid_example", "ApiToken_example", false, "PersonOnEventsQueryingEnabled_example", "DefaultModifiers_example", "ProductIntents_example", "SecretApiToken_example", "SecretApiTokenBackup_example") // ProjectBackwardCompat |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ChangeOrganizationCreate(context.Background(), id, organizationId).ProjectBackwardCompat(projectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ChangeOrganizationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeOrganizationCreate`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ChangeOrganizationCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | A unique value identifying this project. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeOrganizationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectBackwardCompat** | [**ProjectBackwardCompat**](ProjectBackwardCompat.md) |  | 

### Return type

[**ProjectBackwardCompat**](ProjectBackwardCompat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteProductOnboardingPartialUpdate

> ProjectBackwardCompat CompleteProductOnboardingPartialUpdate(ctx, id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()





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
	id := int64(789) // int64 | A unique value identifying this project.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchedProjectBackwardCompat := *openapiclient.NewPatchedProjectBackwardCompat() // PatchedProjectBackwardCompat |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.CompleteProductOnboardingPartialUpdate(context.Background(), id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.CompleteProductOnboardingPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteProductOnboardingPartialUpdate`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.CompleteProductOnboardingPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | A unique value identifying this project. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteProductOnboardingPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedProjectBackwardCompat** | [**PatchedProjectBackwardCompat**](PatchedProjectBackwardCompat.md) |  | 

### Return type

[**ProjectBackwardCompat**](ProjectBackwardCompat.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> Organization Create(ctx).Organization(organization).Execute()



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
	organization := *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example") // Organization | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.Create(context.Background()).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organization** | [**Organization**](Organization.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create2

> ProjectBackwardCompat Create2(ctx, organizationId).ProjectBackwardCompat(projectBackwardCompat).Execute()





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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectBackwardCompat := *openapiclient.NewProjectBackwardCompat(int32(123), "Organization_example", time.Now(), "TODO", false, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, "LiveEventsToken_example", time.Now(), "Uuid_example", "ApiToken_example", false, "PersonOnEventsQueryingEnabled_example", "DefaultModifiers_example", "ProductIntents_example", "SecretApiToken_example", "SecretApiTokenBackup_example") // ProjectBackwardCompat |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.Create2(context.Background(), organizationId).ProjectBackwardCompat(projectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.Create2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create2`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.Create2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectBackwardCompat** | [**ProjectBackwardCompat**](ProjectBackwardCompat.md) |  | 

### Return type

[**ProjectBackwardCompat**](ProjectBackwardCompat.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create_0

> Organization Create_0(ctx).Organization(organization).Execute()



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
	organization := *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example") // Organization | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.Create_0(context.Background()).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.Create_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create_0`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.Create_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreate_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organization** | [**Organization**](Organization.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecretTokenBackupPartialUpdate

> ProjectBackwardCompat DeleteSecretTokenBackupPartialUpdate(ctx, id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()





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
	id := int64(789) // int64 | A unique value identifying this project.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchedProjectBackwardCompat := *openapiclient.NewPatchedProjectBackwardCompat() // PatchedProjectBackwardCompat |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.DeleteSecretTokenBackupPartialUpdate(context.Background(), id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteSecretTokenBackupPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSecretTokenBackupPartialUpdate`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.DeleteSecretTokenBackupPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | A unique value identifying this project. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretTokenBackupPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedProjectBackwardCompat** | [**PatchedProjectBackwardCompat**](PatchedProjectBackwardCompat.md) |  | 

### Return type

[**ProjectBackwardCompat**](ProjectBackwardCompat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Destroy

> Destroy(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this organization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.Destroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.Destroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyRequest struct via the builder pattern


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


## Destroy2

> Destroy2(ctx, id, organizationId).Execute()





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
	id := int64(789) // int64 | A unique value identifying this project.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.Destroy2(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.Destroy2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | A unique value identifying this project. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroy2Request struct via the builder pattern


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


## Destroy_0

> Destroy_0(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this organization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.Destroy_0(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.Destroy_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroy_2Request struct via the builder pattern


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


## DomainsCreate

> OrganizationDomain DomainsCreate(ctx, organizationId).OrganizationDomain(organizationDomain).Execute()



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	organizationDomain := *openapiclient.NewOrganizationDomain("Id_example", "Domain_example", false, time.Now(), "VerificationChallenge_example", false, false, "ScimBaseUrl_example", "ScimBearerToken_example") // OrganizationDomain | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.DomainsCreate(context.Background(), organizationId).OrganizationDomain(organizationDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DomainsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsCreate`: OrganizationDomain
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.DomainsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationDomain** | [**OrganizationDomain**](OrganizationDomain.md) |  | 

### Return type

[**OrganizationDomain**](OrganizationDomain.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDestroy

> DomainsDestroy(ctx, id, organizationId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this domain.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DomainsDestroy(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DomainsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this domain. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDestroyRequest struct via the builder pattern


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


## DomainsList

> PaginatedOrganizationDomainList DomainsList(ctx, organizationId).Limit(limit).Offset(offset).Execute()



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.DomainsList(context.Background(), organizationId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DomainsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsList`: PaginatedOrganizationDomainList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.DomainsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedOrganizationDomainList**](PaginatedOrganizationDomainList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsPartialUpdate

> OrganizationDomain DomainsPartialUpdate(ctx, id, organizationId).PatchedOrganizationDomain(patchedOrganizationDomain).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this domain.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchedOrganizationDomain := *openapiclient.NewPatchedOrganizationDomain() // PatchedOrganizationDomain |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.DomainsPartialUpdate(context.Background(), id, organizationId).PatchedOrganizationDomain(patchedOrganizationDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DomainsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsPartialUpdate`: OrganizationDomain
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.DomainsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this domain. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedOrganizationDomain** | [**PatchedOrganizationDomain**](PatchedOrganizationDomain.md) |  | 

### Return type

[**OrganizationDomain**](OrganizationDomain.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsRetrieve

> OrganizationDomain DomainsRetrieve(ctx, id, organizationId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this domain.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.DomainsRetrieve(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DomainsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsRetrieve`: OrganizationDomain
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.DomainsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this domain. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationDomain**](OrganizationDomain.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsScimTokenCreate

> DomainsScimTokenCreate(ctx, id, organizationId).OrganizationDomain(organizationDomain).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this domain.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	organizationDomain := *openapiclient.NewOrganizationDomain("Id_example", "Domain_example", false, time.Now(), "VerificationChallenge_example", false, false, "ScimBaseUrl_example", "ScimBearerToken_example") // OrganizationDomain | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DomainsScimTokenCreate(context.Background(), id, organizationId).OrganizationDomain(organizationDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DomainsScimTokenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this domain. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsScimTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **organizationDomain** | [**OrganizationDomain**](OrganizationDomain.md) |  | 

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


## DomainsUpdate

> OrganizationDomain DomainsUpdate(ctx, id, organizationId).OrganizationDomain(organizationDomain).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this domain.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	organizationDomain := *openapiclient.NewOrganizationDomain("Id_example", "Domain_example", false, time.Now(), "VerificationChallenge_example", false, false, "ScimBaseUrl_example", "ScimBearerToken_example") // OrganizationDomain | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.DomainsUpdate(context.Background(), id, organizationId).OrganizationDomain(organizationDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DomainsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DomainsUpdate`: OrganizationDomain
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.DomainsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this domain. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **organizationDomain** | [**OrganizationDomain**](OrganizationDomain.md) |  | 

### Return type

[**OrganizationDomain**](OrganizationDomain.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsVerifyCreate

> DomainsVerifyCreate(ctx, id, organizationId).OrganizationDomain(organizationDomain).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this domain.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	organizationDomain := *openapiclient.NewOrganizationDomain("Id_example", "Domain_example", false, time.Now(), "VerificationChallenge_example", false, false, "ScimBaseUrl_example", "ScimBearerToken_example") // OrganizationDomain | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DomainsVerifyCreate(context.Background(), id, organizationId).OrganizationDomain(organizationDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DomainsVerifyCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this domain. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsVerifyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **organizationDomain** | [**OrganizationDomain**](OrganizationDomain.md) |  | 

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


## EnvironmentsRollbackCreate

> Organization EnvironmentsRollbackCreate(ctx, id).Organization(organization).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this organization.
	organization := *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example") // Organization | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.EnvironmentsRollbackCreate(context.Background(), id).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.EnvironmentsRollbackCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsRollbackCreate`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.EnvironmentsRollbackCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsRollbackCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organization** | [**Organization**](Organization.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitesBulkCreate

> InvitesBulkCreate(ctx, organizationId).OrganizationInvite(organizationInvite).Execute()



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	organizationInvite := *openapiclient.NewOrganizationInvite("Id_example", "TargetEmail_example", false, false, *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now()) // OrganizationInvite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.InvitesBulkCreate(context.Background(), organizationId).OrganizationInvite(organizationInvite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.InvitesBulkCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInvitesBulkCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationInvite** | [**OrganizationInvite**](OrganizationInvite.md) |  | 

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


## InvitesCreate

> OrganizationInvite InvitesCreate(ctx, organizationId).OrganizationInvite(organizationInvite).Execute()



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	organizationInvite := *openapiclient.NewOrganizationInvite("Id_example", "TargetEmail_example", false, false, *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now()) // OrganizationInvite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.InvitesCreate(context.Background(), organizationId).OrganizationInvite(organizationInvite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.InvitesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvitesCreate`: OrganizationInvite
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.InvitesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationInvite** | [**OrganizationInvite**](OrganizationInvite.md) |  | 

### Return type

[**OrganizationInvite**](OrganizationInvite.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvitesDestroy

> InvitesDestroy(ctx, id, organizationId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this organization invite.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.InvitesDestroy(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.InvitesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this organization invite. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitesDestroyRequest struct via the builder pattern


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


## InvitesList

> PaginatedOrganizationInviteList InvitesList(ctx, organizationId).Limit(limit).Offset(offset).Execute()



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.InvitesList(context.Background(), organizationId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.InvitesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvitesList`: PaginatedOrganizationInviteList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.InvitesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvitesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedOrganizationInviteList**](PaginatedOrganizationInviteList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsGeneratingDemoDataRetrieve

> ProjectBackwardCompat IsGeneratingDemoDataRetrieve(ctx, id, organizationId).Execute()





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
	id := int64(789) // int64 | A unique value identifying this project.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.IsGeneratingDemoDataRetrieve(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.IsGeneratingDemoDataRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsGeneratingDemoDataRetrieve`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.IsGeneratingDemoDataRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | A unique value identifying this project. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsGeneratingDemoDataRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectBackwardCompat**](ProjectBackwardCompat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> PaginatedOrganizationList List(ctx).Limit(limit).Offset(offset).Execute()



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
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.List(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: PaginatedOrganizationList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedOrganizationList**](PaginatedOrganizationList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List2

> PaginatedProjectBackwardCompatBasicList List2(ctx, organizationId).Limit(limit).Offset(offset).Execute()





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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.List2(context.Background(), organizationId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.List2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List2`: PaginatedProjectBackwardCompatBasicList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.List2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiList2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedProjectBackwardCompatBasicList**](PaginatedProjectBackwardCompatBasicList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List_0

> PaginatedOrganizationList List_0(ctx).Limit(limit).Offset(offset).Execute()



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
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.List_0(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.List_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List_0`: PaginatedOrganizationList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.List_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiList_3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedOrganizationList**](PaginatedOrganizationList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MembersDestroy

> MembersDestroy(ctx, organizationId, userUuid).Execute()



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.MembersDestroy(context.Background(), organizationId, userUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.MembersDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**userUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMembersDestroyRequest struct via the builder pattern


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


## MembersList

> PaginatedOrganizationMemberList MembersList(ctx, organizationId).Limit(limit).Offset(offset).Execute()



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.MembersList(context.Background(), organizationId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.MembersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MembersList`: PaginatedOrganizationMemberList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.MembersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMembersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedOrganizationMemberList**](PaginatedOrganizationMemberList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MembersPartialUpdate

> OrganizationMember MembersPartialUpdate(ctx, organizationId, userUuid).PatchedOrganizationMember(patchedOrganizationMember).Execute()



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchedOrganizationMember := *openapiclient.NewPatchedOrganizationMember() // PatchedOrganizationMember |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.MembersPartialUpdate(context.Background(), organizationId, userUuid).PatchedOrganizationMember(patchedOrganizationMember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.MembersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MembersPartialUpdate`: OrganizationMember
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.MembersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**userUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMembersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedOrganizationMember** | [**PatchedOrganizationMember**](PatchedOrganizationMember.md) |  | 

### Return type

[**OrganizationMember**](OrganizationMember.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MembersScopedApiKeysRetrieve

> OrganizationMember MembersScopedApiKeysRetrieve(ctx, organizationId, userUuid).Execute()



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.MembersScopedApiKeysRetrieve(context.Background(), organizationId, userUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.MembersScopedApiKeysRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MembersScopedApiKeysRetrieve`: OrganizationMember
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.MembersScopedApiKeysRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**userUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMembersScopedApiKeysRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationMember**](OrganizationMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MembersUpdate

> OrganizationMember MembersUpdate(ctx, organizationId, userUuid).OrganizationMember(organizationMember).Execute()



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	organizationMember := *openapiclient.NewOrganizationMember("Id_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now(), false, false, time.Now()) // OrganizationMember |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.MembersUpdate(context.Background(), organizationId, userUuid).OrganizationMember(organizationMember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.MembersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MembersUpdate`: OrganizationMember
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.MembersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**userUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMembersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **organizationMember** | [**OrganizationMember**](OrganizationMember.md) |  | 

### Return type

[**OrganizationMember**](OrganizationMember.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdate

> Organization PartialUpdate(ctx, id).PatchedOrganization(patchedOrganization).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this organization.
	patchedOrganization := *openapiclient.NewPatchedOrganization() // PatchedOrganization |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PartialUpdate(context.Background(), id).PatchedOrganization(patchedOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdate`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedOrganization** | [**PatchedOrganization**](PatchedOrganization.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdate2

> ProjectBackwardCompat PartialUpdate2(ctx, id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()





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
	id := int64(789) // int64 | A unique value identifying this project.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchedProjectBackwardCompat := *openapiclient.NewPatchedProjectBackwardCompat() // PatchedProjectBackwardCompat |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PartialUpdate2(context.Background(), id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PartialUpdate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdate2`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PartialUpdate2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | A unique value identifying this project. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedProjectBackwardCompat** | [**PatchedProjectBackwardCompat**](PatchedProjectBackwardCompat.md) |  | 

### Return type

[**ProjectBackwardCompat**](ProjectBackwardCompat.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdate_0

> Organization PartialUpdate_0(ctx, id).PatchedOrganization(patchedOrganization).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this organization.
	patchedOrganization := *openapiclient.NewPatchedOrganization() // PatchedOrganization |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.PartialUpdate_0(context.Background(), id).PatchedOrganization(patchedOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.PartialUpdate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdate_0`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.PartialUpdate_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdate_4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedOrganization** | [**PatchedOrganization**](PatchedOrganization.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxyRecordsCreate

> ProxyRecord ProxyRecordsCreate(ctx, organizationId).ProxyRecord(proxyRecord).Execute()



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
	proxyRecord := *openapiclient.NewProxyRecord("Id_example", "Domain_example", "TargetCname_example", openapiclient.ProxyRecordStatusEnum("waiting"), "Message_example", time.Now(), time.Now(), NullableInt32(123)) // ProxyRecord | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ProxyRecordsCreate(context.Background(), organizationId).ProxyRecord(proxyRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ProxyRecordsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyRecordsCreate`: ProxyRecord
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ProxyRecordsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyRecordsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **proxyRecord** | [**ProxyRecord**](ProxyRecord.md) |  | 

### Return type

[**ProxyRecord**](ProxyRecord.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxyRecordsDestroy

> ProxyRecordsDestroy(ctx, id, organizationId).Execute()



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
	id := "id_example" // string | 
	organizationId := "organizationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.ProxyRecordsDestroy(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ProxyRecordsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyRecordsDestroyRequest struct via the builder pattern


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


## ProxyRecordsList

> PaginatedProxyRecordList ProxyRecordsList(ctx, organizationId).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.OrganizationsAPI.ProxyRecordsList(context.Background(), organizationId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ProxyRecordsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyRecordsList`: PaginatedProxyRecordList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ProxyRecordsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyRecordsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedProxyRecordList**](PaginatedProxyRecordList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxyRecordsPartialUpdate

> ProxyRecord ProxyRecordsPartialUpdate(ctx, id, organizationId).PatchedProxyRecord(patchedProxyRecord).Execute()



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
	id := "id_example" // string | 
	organizationId := "organizationId_example" // string | 
	patchedProxyRecord := *openapiclient.NewPatchedProxyRecord() // PatchedProxyRecord |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ProxyRecordsPartialUpdate(context.Background(), id, organizationId).PatchedProxyRecord(patchedProxyRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ProxyRecordsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyRecordsPartialUpdate`: ProxyRecord
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ProxyRecordsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyRecordsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedProxyRecord** | [**PatchedProxyRecord**](PatchedProxyRecord.md) |  | 

### Return type

[**ProxyRecord**](ProxyRecord.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxyRecordsRetrieve

> ProxyRecord ProxyRecordsRetrieve(ctx, id, organizationId).Execute()



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
	id := "id_example" // string | 
	organizationId := "organizationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ProxyRecordsRetrieve(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ProxyRecordsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyRecordsRetrieve`: ProxyRecord
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ProxyRecordsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyRecordsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProxyRecord**](ProxyRecord.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProxyRecordsUpdate

> ProxyRecord ProxyRecordsUpdate(ctx, id, organizationId).ProxyRecord(proxyRecord).Execute()



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
	id := "id_example" // string | 
	organizationId := "organizationId_example" // string | 
	proxyRecord := *openapiclient.NewProxyRecord("Id_example", "Domain_example", "TargetCname_example", openapiclient.ProxyRecordStatusEnum("waiting"), "Message_example", time.Now(), time.Now(), NullableInt32(123)) // ProxyRecord | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ProxyRecordsUpdate(context.Background(), id, organizationId).ProxyRecord(proxyRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ProxyRecordsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyRecordsUpdate`: ProxyRecord
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ProxyRecordsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProxyRecordsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **proxyRecord** | [**ProxyRecord**](ProxyRecord.md) |  | 

### Return type

[**ProxyRecord**](ProxyRecord.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetTokenPartialUpdate

> ProjectBackwardCompat ResetTokenPartialUpdate(ctx, id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()





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
	id := int64(789) // int64 | A unique value identifying this project.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchedProjectBackwardCompat := *openapiclient.NewPatchedProjectBackwardCompat() // PatchedProjectBackwardCompat |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.ResetTokenPartialUpdate(context.Background(), id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.ResetTokenPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetTokenPartialUpdate`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.ResetTokenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | A unique value identifying this project. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetTokenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedProjectBackwardCompat** | [**PatchedProjectBackwardCompat**](PatchedProjectBackwardCompat.md) |  | 

### Return type

[**ProjectBackwardCompat**](ProjectBackwardCompat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retrieve

> Organization Retrieve(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this organization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.Retrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.Retrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Retrieve`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organization**](Organization.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retrieve2

> ProjectBackwardCompat Retrieve2(ctx, id, organizationId).Execute()





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
	id := int64(789) // int64 | A unique value identifying this project.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.Retrieve2(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.Retrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Retrieve2`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.Retrieve2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | A unique value identifying this project. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieve2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectBackwardCompat**](ProjectBackwardCompat.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retrieve_0

> Organization Retrieve_0(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this organization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.Retrieve_0(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.Retrieve_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Retrieve_0`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.Retrieve_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieve_5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organization**](Organization.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesCreate

> Role RolesCreate(ctx, organizationId).Role(role).Execute()



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	role := *openapiclient.NewRole("Id_example", "Name_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), "Members_example", "IsDefault_example") // Role | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.RolesCreate(context.Background(), organizationId).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RolesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolesCreate`: Role
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.RolesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | [**Role**](Role.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesDestroy

> RolesDestroy(ctx, id, organizationId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this role.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.RolesDestroy(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RolesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this role. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesDestroyRequest struct via the builder pattern


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


## RolesList

> PaginatedRoleList RolesList(ctx, organizationId).Limit(limit).Offset(offset).Execute()



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.RolesList(context.Background(), organizationId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RolesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolesList`: PaginatedRoleList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.RolesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedRoleList**](PaginatedRoleList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesPartialUpdate

> Role RolesPartialUpdate(ctx, id, organizationId).PatchedRole(patchedRole).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this role.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchedRole := *openapiclient.NewPatchedRole() // PatchedRole |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.RolesPartialUpdate(context.Background(), id, organizationId).PatchedRole(patchedRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RolesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolesPartialUpdate`: Role
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.RolesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this role. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedRole** | [**PatchedRole**](PatchedRole.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesRetrieve

> Role RolesRetrieve(ctx, id, organizationId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this role.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.RolesRetrieve(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RolesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolesRetrieve`: Role
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.RolesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this role. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Role**](Role.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesRoleMembershipsCreate

> RoleMembership RolesRoleMembershipsCreate(ctx, organizationId, roleId).RoleMembership(roleMembership).Execute()



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
	roleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleMembership := *openapiclient.NewRoleMembership("Id_example", "RoleId_example", *openapiclient.NewOrganizationMember("Id_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now(), false, false, time.Now()), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now(), "UserUuid_example") // RoleMembership | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.RolesRoleMembershipsCreate(context.Background(), organizationId, roleId).RoleMembership(roleMembership).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RolesRoleMembershipsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolesRoleMembershipsCreate`: RoleMembership
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.RolesRoleMembershipsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesRoleMembershipsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleMembership** | [**RoleMembership**](RoleMembership.md) |  | 

### Return type

[**RoleMembership**](RoleMembership.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesRoleMembershipsDestroy

> RolesRoleMembershipsDestroy(ctx, id, organizationId, roleId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this role membership.
	organizationId := "organizationId_example" // string | 
	roleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.RolesRoleMembershipsDestroy(context.Background(), id, organizationId, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RolesRoleMembershipsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this role membership. | 
**organizationId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesRoleMembershipsDestroyRequest struct via the builder pattern


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


## RolesRoleMembershipsList

> PaginatedRoleMembershipList RolesRoleMembershipsList(ctx, organizationId, roleId).Limit(limit).Offset(offset).Execute()



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
	roleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.RolesRoleMembershipsList(context.Background(), organizationId, roleId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RolesRoleMembershipsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolesRoleMembershipsList`: PaginatedRoleMembershipList
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.RolesRoleMembershipsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesRoleMembershipsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedRoleMembershipList**](PaginatedRoleMembershipList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesUpdate

> Role RolesUpdate(ctx, id, organizationId).Role(role).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this role.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	role := *openapiclient.NewRole("Id_example", "Name_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), "Members_example", "IsDefault_example") // Role | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.RolesUpdate(context.Background(), id, organizationId).Role(role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RolesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolesUpdate`: Role
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.RolesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this role. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **role** | [**Role**](Role.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateSecretTokenPartialUpdate

> ProjectBackwardCompat RotateSecretTokenPartialUpdate(ctx, id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()





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
	id := int64(789) // int64 | A unique value identifying this project.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchedProjectBackwardCompat := *openapiclient.NewPatchedProjectBackwardCompat() // PatchedProjectBackwardCompat |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.RotateSecretTokenPartialUpdate(context.Background(), id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.RotateSecretTokenPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateSecretTokenPartialUpdate`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.RotateSecretTokenPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | A unique value identifying this project. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateSecretTokenPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedProjectBackwardCompat** | [**PatchedProjectBackwardCompat**](PatchedProjectBackwardCompat.md) |  | 

### Return type

[**ProjectBackwardCompat**](ProjectBackwardCompat.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Organization Update(ctx, id).Organization(organization).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this organization.
	organization := *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example") // Organization | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.Update(context.Background(), id).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organization** | [**Organization**](Organization.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update2

> ProjectBackwardCompat Update2(ctx, id, organizationId).ProjectBackwardCompat(projectBackwardCompat).Execute()





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
	id := int64(789) // int64 | A unique value identifying this project.
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectBackwardCompat := *openapiclient.NewProjectBackwardCompat(int32(123), "Organization_example", time.Now(), "TODO", false, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, "LiveEventsToken_example", time.Now(), "Uuid_example", "ApiToken_example", false, "PersonOnEventsQueryingEnabled_example", "DefaultModifiers_example", "ProductIntents_example", "SecretApiToken_example", "SecretApiTokenBackup_example") // ProjectBackwardCompat |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.Update2(context.Background(), id, organizationId).ProjectBackwardCompat(projectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.Update2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update2`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.Update2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | A unique value identifying this project. | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectBackwardCompat** | [**ProjectBackwardCompat**](ProjectBackwardCompat.md) |  | 

### Return type

[**ProjectBackwardCompat**](ProjectBackwardCompat.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update_0

> Organization Update_0(ctx, id).Organization(organization).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this organization.
	organization := *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example") // Organization | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.Update_0(context.Background(), id).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.Update_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update_0`: Organization
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.Update_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdate_6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organization** | [**Organization**](Organization.md) |  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

