# \ProxyRecordsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProxyRecordsCreate**](ProxyRecordsAPI.md#ProxyRecordsCreate) | **Post** /api/organizations/{organization_id}/proxy_records/ | 
[**ProxyRecordsDestroy**](ProxyRecordsAPI.md#ProxyRecordsDestroy) | **Delete** /api/organizations/{organization_id}/proxy_records/{id}/ | 
[**ProxyRecordsList**](ProxyRecordsAPI.md#ProxyRecordsList) | **Get** /api/organizations/{organization_id}/proxy_records/ | 
[**ProxyRecordsPartialUpdate**](ProxyRecordsAPI.md#ProxyRecordsPartialUpdate) | **Patch** /api/organizations/{organization_id}/proxy_records/{id}/ | 
[**ProxyRecordsRetrieve**](ProxyRecordsAPI.md#ProxyRecordsRetrieve) | **Get** /api/organizations/{organization_id}/proxy_records/{id}/ | 
[**ProxyRecordsUpdate**](ProxyRecordsAPI.md#ProxyRecordsUpdate) | **Put** /api/organizations/{organization_id}/proxy_records/{id}/ | 



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
	resp, r, err := apiClient.ProxyRecordsAPI.ProxyRecordsCreate(context.Background(), organizationId).ProxyRecord(proxyRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyRecordsAPI.ProxyRecordsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyRecordsCreate`: ProxyRecord
	fmt.Fprintf(os.Stdout, "Response from `ProxyRecordsAPI.ProxyRecordsCreate`: %v\n", resp)
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
	r, err := apiClient.ProxyRecordsAPI.ProxyRecordsDestroy(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyRecordsAPI.ProxyRecordsDestroy``: %v\n", err)
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
	resp, r, err := apiClient.ProxyRecordsAPI.ProxyRecordsList(context.Background(), organizationId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyRecordsAPI.ProxyRecordsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyRecordsList`: PaginatedProxyRecordList
	fmt.Fprintf(os.Stdout, "Response from `ProxyRecordsAPI.ProxyRecordsList`: %v\n", resp)
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
	resp, r, err := apiClient.ProxyRecordsAPI.ProxyRecordsPartialUpdate(context.Background(), id, organizationId).PatchedProxyRecord(patchedProxyRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyRecordsAPI.ProxyRecordsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyRecordsPartialUpdate`: ProxyRecord
	fmt.Fprintf(os.Stdout, "Response from `ProxyRecordsAPI.ProxyRecordsPartialUpdate`: %v\n", resp)
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
	resp, r, err := apiClient.ProxyRecordsAPI.ProxyRecordsRetrieve(context.Background(), id, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyRecordsAPI.ProxyRecordsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyRecordsRetrieve`: ProxyRecord
	fmt.Fprintf(os.Stdout, "Response from `ProxyRecordsAPI.ProxyRecordsRetrieve`: %v\n", resp)
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
	resp, r, err := apiClient.ProxyRecordsAPI.ProxyRecordsUpdate(context.Background(), id, organizationId).ProxyRecord(proxyRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProxyRecordsAPI.ProxyRecordsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProxyRecordsUpdate`: ProxyRecord
	fmt.Fprintf(os.Stdout, "Response from `ProxyRecordsAPI.ProxyRecordsUpdate`: %v\n", resp)
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

