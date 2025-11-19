# \EnvironmentsRollbackAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsRollbackCreate**](EnvironmentsRollbackAPI.md#EnvironmentsRollbackCreate) | **Post** /api/organizations/{id}/environments_rollback/ | 



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
	resp, r, err := apiClient.EnvironmentsRollbackAPI.EnvironmentsRollbackCreate(context.Background(), id).Organization(organization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsRollbackAPI.EnvironmentsRollbackCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsRollbackCreate`: Organization
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsRollbackAPI.EnvironmentsRollbackCreate`: %v\n", resp)
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

