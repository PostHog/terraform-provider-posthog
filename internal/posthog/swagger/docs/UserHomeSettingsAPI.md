# \UserHomeSettingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserHomeSettingsPartialUpdate**](UserHomeSettingsAPI.md#UserHomeSettingsPartialUpdate) | **Patch** /api/user_home_settings/{uuid}/ | 
[**UserHomeSettingsPartialUpdate_0**](UserHomeSettingsAPI.md#UserHomeSettingsPartialUpdate_0) | **Patch** /api/user_home_settings/{uuid}/ | 
[**UserHomeSettingsRetrieve**](UserHomeSettingsAPI.md#UserHomeSettingsRetrieve) | **Get** /api/user_home_settings/{uuid}/ | 
[**UserHomeSettingsRetrieve_0**](UserHomeSettingsAPI.md#UserHomeSettingsRetrieve_0) | **Get** /api/user_home_settings/{uuid}/ | 



## UserHomeSettingsPartialUpdate

> PinnedSceneTabs UserHomeSettingsPartialUpdate(ctx, uuid).PatchedPinnedSceneTabs(patchedPinnedSceneTabs).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchedPinnedSceneTabs := *openapiclient.NewPatchedPinnedSceneTabs() // PatchedPinnedSceneTabs |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserHomeSettingsAPI.UserHomeSettingsPartialUpdate(context.Background(), uuid).PatchedPinnedSceneTabs(patchedPinnedSceneTabs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserHomeSettingsAPI.UserHomeSettingsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserHomeSettingsPartialUpdate`: PinnedSceneTabs
	fmt.Fprintf(os.Stdout, "Response from `UserHomeSettingsAPI.UserHomeSettingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserHomeSettingsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPinnedSceneTabs** | [**PatchedPinnedSceneTabs**](PatchedPinnedSceneTabs.md) |  | 

### Return type

[**PinnedSceneTabs**](PinnedSceneTabs.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserHomeSettingsPartialUpdate_0

> PinnedSceneTabs UserHomeSettingsPartialUpdate_0(ctx, uuid).PatchedPinnedSceneTabs(patchedPinnedSceneTabs).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchedPinnedSceneTabs := *openapiclient.NewPatchedPinnedSceneTabs() // PatchedPinnedSceneTabs |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserHomeSettingsAPI.UserHomeSettingsPartialUpdate_0(context.Background(), uuid).PatchedPinnedSceneTabs(patchedPinnedSceneTabs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserHomeSettingsAPI.UserHomeSettingsPartialUpdate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserHomeSettingsPartialUpdate_0`: PinnedSceneTabs
	fmt.Fprintf(os.Stdout, "Response from `UserHomeSettingsAPI.UserHomeSettingsPartialUpdate_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserHomeSettingsPartialUpdate_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedPinnedSceneTabs** | [**PatchedPinnedSceneTabs**](PatchedPinnedSceneTabs.md) |  | 

### Return type

[**PinnedSceneTabs**](PinnedSceneTabs.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserHomeSettingsRetrieve

> PinnedSceneTabs UserHomeSettingsRetrieve(ctx, uuid).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserHomeSettingsAPI.UserHomeSettingsRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserHomeSettingsAPI.UserHomeSettingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserHomeSettingsRetrieve`: PinnedSceneTabs
	fmt.Fprintf(os.Stdout, "Response from `UserHomeSettingsAPI.UserHomeSettingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserHomeSettingsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PinnedSceneTabs**](PinnedSceneTabs.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserHomeSettingsRetrieve_0

> PinnedSceneTabs UserHomeSettingsRetrieve_0(ctx, uuid).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserHomeSettingsAPI.UserHomeSettingsRetrieve_0(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserHomeSettingsAPI.UserHomeSettingsRetrieve_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserHomeSettingsRetrieve_0`: PinnedSceneTabs
	fmt.Fprintf(os.Stdout, "Response from `UserHomeSettingsAPI.UserHomeSettingsRetrieve_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserHomeSettingsRetrieve_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PinnedSceneTabs**](PinnedSceneTabs.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

