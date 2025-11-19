# \ProjectsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivityRetrieve**](ProjectsAPI.md#ActivityRetrieve) | **Get** /api/organizations/{organization_id}/projects/{id}/activity/ | 
[**AddProductIntentPartialUpdate**](ProjectsAPI.md#AddProductIntentPartialUpdate) | **Patch** /api/organizations/{organization_id}/projects/{id}/add_product_intent/ | 
[**ChangeOrganizationCreate**](ProjectsAPI.md#ChangeOrganizationCreate) | **Post** /api/organizations/{organization_id}/projects/{id}/change_organization/ | 
[**CompleteProductOnboardingPartialUpdate**](ProjectsAPI.md#CompleteProductOnboardingPartialUpdate) | **Patch** /api/organizations/{organization_id}/projects/{id}/complete_product_onboarding/ | 
[**Create2**](ProjectsAPI.md#Create2) | **Post** /api/organizations/{organization_id}/projects/ | 
[**DeleteSecretTokenBackupPartialUpdate**](ProjectsAPI.md#DeleteSecretTokenBackupPartialUpdate) | **Patch** /api/organizations/{organization_id}/projects/{id}/delete_secret_token_backup/ | 
[**Destroy2**](ProjectsAPI.md#Destroy2) | **Delete** /api/organizations/{organization_id}/projects/{id}/ | 
[**IsGeneratingDemoDataRetrieve**](ProjectsAPI.md#IsGeneratingDemoDataRetrieve) | **Get** /api/organizations/{organization_id}/projects/{id}/is_generating_demo_data/ | 
[**List2**](ProjectsAPI.md#List2) | **Get** /api/organizations/{organization_id}/projects/ | 
[**PartialUpdate2**](ProjectsAPI.md#PartialUpdate2) | **Patch** /api/organizations/{organization_id}/projects/{id}/ | 
[**ResetTokenPartialUpdate**](ProjectsAPI.md#ResetTokenPartialUpdate) | **Patch** /api/organizations/{organization_id}/projects/{id}/reset_token/ | 
[**Retrieve2**](ProjectsAPI.md#Retrieve2) | **Get** /api/organizations/{organization_id}/projects/{id}/ | 
[**RotateSecretTokenPartialUpdate**](ProjectsAPI.md#RotateSecretTokenPartialUpdate) | **Patch** /api/organizations/{organization_id}/projects/{id}/rotate_secret_token/ | 
[**Update2**](ProjectsAPI.md#Update2) | **Put** /api/organizations/{organization_id}/projects/{id}/ | 



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
	resp, r, err := apiClient.ProjectsAPI.ActivityRetrieve(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ActivityRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivityRetrieve`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ActivityRetrieve`: %v\n", resp)
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
	resp, r, err := apiClient.ProjectsAPI.AddProductIntentPartialUpdate(context.Background(), id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.AddProductIntentPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddProductIntentPartialUpdate`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.AddProductIntentPartialUpdate`: %v\n", resp)
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
	resp, r, err := apiClient.ProjectsAPI.ChangeOrganizationCreate(context.Background(), id, organizationId).ProjectBackwardCompat(projectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ChangeOrganizationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeOrganizationCreate`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ChangeOrganizationCreate`: %v\n", resp)
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
	resp, r, err := apiClient.ProjectsAPI.CompleteProductOnboardingPartialUpdate(context.Background(), id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.CompleteProductOnboardingPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteProductOnboardingPartialUpdate`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.CompleteProductOnboardingPartialUpdate`: %v\n", resp)
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
	resp, r, err := apiClient.ProjectsAPI.Create2(context.Background(), organizationId).ProjectBackwardCompat(projectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.Create2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create2`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.Create2`: %v\n", resp)
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
	resp, r, err := apiClient.ProjectsAPI.DeleteSecretTokenBackupPartialUpdate(context.Background(), id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteSecretTokenBackupPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSecretTokenBackupPartialUpdate`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.DeleteSecretTokenBackupPartialUpdate`: %v\n", resp)
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
	r, err := apiClient.ProjectsAPI.Destroy2(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.Destroy2``: %v\n", err)
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
	resp, r, err := apiClient.ProjectsAPI.IsGeneratingDemoDataRetrieve(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.IsGeneratingDemoDataRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsGeneratingDemoDataRetrieve`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.IsGeneratingDemoDataRetrieve`: %v\n", resp)
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
	resp, r, err := apiClient.ProjectsAPI.List2(context.Background(), organizationId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.List2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List2`: PaginatedProjectBackwardCompatBasicList
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.List2`: %v\n", resp)
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
	resp, r, err := apiClient.ProjectsAPI.PartialUpdate2(context.Background(), id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PartialUpdate2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdate2`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PartialUpdate2`: %v\n", resp)
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
	resp, r, err := apiClient.ProjectsAPI.ResetTokenPartialUpdate(context.Background(), id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ResetTokenPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetTokenPartialUpdate`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ResetTokenPartialUpdate`: %v\n", resp)
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
	resp, r, err := apiClient.ProjectsAPI.Retrieve2(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.Retrieve2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Retrieve2`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.Retrieve2`: %v\n", resp)
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
	resp, r, err := apiClient.ProjectsAPI.RotateSecretTokenPartialUpdate(context.Background(), id, organizationId).PatchedProjectBackwardCompat(patchedProjectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.RotateSecretTokenPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateSecretTokenPartialUpdate`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.RotateSecretTokenPartialUpdate`: %v\n", resp)
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
	resp, r, err := apiClient.ProjectsAPI.Update2(context.Background(), id, organizationId).ProjectBackwardCompat(projectBackwardCompat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.Update2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update2`: ProjectBackwardCompat
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.Update2`: %v\n", resp)
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

