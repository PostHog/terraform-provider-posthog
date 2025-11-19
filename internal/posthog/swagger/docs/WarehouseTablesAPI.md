# \WarehouseTablesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WarehouseTablesCreate**](WarehouseTablesAPI.md#WarehouseTablesCreate) | **Post** /api/projects/{project_id}/warehouse_tables/ | 
[**WarehouseTablesDestroy**](WarehouseTablesAPI.md#WarehouseTablesDestroy) | **Delete** /api/projects/{project_id}/warehouse_tables/{id}/ | 
[**WarehouseTablesFileCreate**](WarehouseTablesAPI.md#WarehouseTablesFileCreate) | **Post** /api/projects/{project_id}/warehouse_tables/file/ | 
[**WarehouseTablesList**](WarehouseTablesAPI.md#WarehouseTablesList) | **Get** /api/projects/{project_id}/warehouse_tables/ | 
[**WarehouseTablesPartialUpdate**](WarehouseTablesAPI.md#WarehouseTablesPartialUpdate) | **Patch** /api/projects/{project_id}/warehouse_tables/{id}/ | 
[**WarehouseTablesRefreshSchemaCreate**](WarehouseTablesAPI.md#WarehouseTablesRefreshSchemaCreate) | **Post** /api/projects/{project_id}/warehouse_tables/{id}/refresh_schema/ | 
[**WarehouseTablesRetrieve**](WarehouseTablesAPI.md#WarehouseTablesRetrieve) | **Get** /api/projects/{project_id}/warehouse_tables/{id}/ | 
[**WarehouseTablesUpdate**](WarehouseTablesAPI.md#WarehouseTablesUpdate) | **Put** /api/projects/{project_id}/warehouse_tables/{id}/ | 
[**WarehouseTablesUpdateSchemaCreate**](WarehouseTablesAPI.md#WarehouseTablesUpdateSchemaCreate) | **Post** /api/projects/{project_id}/warehouse_tables/{id}/update_schema/ | 



## WarehouseTablesCreate

> Table WarehouseTablesCreate(ctx, projectId).Table(table).Execute()





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
	table := *openapiclient.NewTable("Id_example", "Name_example", openapiclient.TableFormatEnum("CSV"), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "UrlPattern_example", *openapiclient.NewCredential("Id_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "AccessKey_example", "AccessSecret_example"), "Columns_example", *openapiclient.NewSimpleExternalDataSourceSerializers("Id_example", time.Now(), NullableInt32(123), "Status_example", openapiclient.SourceTypeEnum("CustomerIO")), "ExternalSchema_example") // Table | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseTablesAPI.WarehouseTablesCreate(context.Background(), projectId).Table(table).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseTablesAPI.WarehouseTablesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseTablesCreate`: Table
	fmt.Fprintf(os.Stdout, "Response from `WarehouseTablesAPI.WarehouseTablesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseTablesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **table** | [**Table**](Table.md) |  | 

### Return type

[**Table**](Table.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseTablesDestroy

> WarehouseTablesDestroy(ctx, id, projectId).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse table.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WarehouseTablesAPI.WarehouseTablesDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseTablesAPI.WarehouseTablesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse table. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseTablesDestroyRequest struct via the builder pattern


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


## WarehouseTablesFileCreate

> WarehouseTablesFileCreate(ctx, projectId).Id(id).Name(name).Format(format).CreatedBy(createdBy).CreatedAt(createdAt).UrlPattern(urlPattern).Credential(credential).Columns(columns).ExternalDataSource(externalDataSource).ExternalSchema(externalSchema).Deleted(deleted).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	name := "name_example" // string | 
	format := openapiclient.TableFormatEnum("CSV") // TableFormatEnum | 
	createdBy := *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}) // UserBasic | 
	createdAt := time.Now() // time.Time | 
	urlPattern := "urlPattern_example" // string | 
	credential := *openapiclient.NewCredential("Id_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "AccessKey_example", "AccessSecret_example") // Credential | 
	columns := "columns_example" // string | 
	externalDataSource := *openapiclient.NewSimpleExternalDataSourceSerializers("Id_example", time.Now(), NullableInt32(123), "Status_example", openapiclient.SourceTypeEnum("CustomerIO")) // SimpleExternalDataSourceSerializers | 
	externalSchema := "externalSchema_example" // string | 
	deleted := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WarehouseTablesAPI.WarehouseTablesFileCreate(context.Background(), projectId).Id(id).Name(name).Format(format).CreatedBy(createdBy).CreatedAt(createdAt).UrlPattern(urlPattern).Credential(credential).Columns(columns).ExternalDataSource(externalDataSource).ExternalSchema(externalSchema).Deleted(deleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseTablesAPI.WarehouseTablesFileCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWarehouseTablesFileCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **string** |  | 
 **name** | **string** |  | 
 **format** | [**TableFormatEnum**](TableFormatEnum.md) |  | 
 **createdBy** | [**UserBasic**](UserBasic.md) |  | 
 **createdAt** | **time.Time** |  | 
 **urlPattern** | **string** |  | 
 **credential** | [**Credential**](Credential.md) |  | 
 **columns** | **string** |  | 
 **externalDataSource** | [**SimpleExternalDataSourceSerializers**](SimpleExternalDataSourceSerializers.md) |  | 
 **externalSchema** | **string** |  | 
 **deleted** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseTablesList

> PaginatedTableList WarehouseTablesList(ctx, projectId).Limit(limit).Offset(offset).Search(search).Execute()





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
	search := "search_example" // string | A search term. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseTablesAPI.WarehouseTablesList(context.Background(), projectId).Limit(limit).Offset(offset).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseTablesAPI.WarehouseTablesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseTablesList`: PaginatedTableList
	fmt.Fprintf(os.Stdout, "Response from `WarehouseTablesAPI.WarehouseTablesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseTablesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **search** | **string** | A search term. | 

### Return type

[**PaginatedTableList**](PaginatedTableList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseTablesPartialUpdate

> Table WarehouseTablesPartialUpdate(ctx, id, projectId).PatchedTable(patchedTable).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse table.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedTable := *openapiclient.NewPatchedTable() // PatchedTable |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseTablesAPI.WarehouseTablesPartialUpdate(context.Background(), id, projectId).PatchedTable(patchedTable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseTablesAPI.WarehouseTablesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseTablesPartialUpdate`: Table
	fmt.Fprintf(os.Stdout, "Response from `WarehouseTablesAPI.WarehouseTablesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse table. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseTablesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedTable** | [**PatchedTable**](PatchedTable.md) |  | 

### Return type

[**Table**](Table.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseTablesRefreshSchemaCreate

> WarehouseTablesRefreshSchemaCreate(ctx, id, projectId).Table(table).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse table.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	table := *openapiclient.NewTable("Id_example", "Name_example", openapiclient.TableFormatEnum("CSV"), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "UrlPattern_example", *openapiclient.NewCredential("Id_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "AccessKey_example", "AccessSecret_example"), "Columns_example", *openapiclient.NewSimpleExternalDataSourceSerializers("Id_example", time.Now(), NullableInt32(123), "Status_example", openapiclient.SourceTypeEnum("CustomerIO")), "ExternalSchema_example") // Table | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WarehouseTablesAPI.WarehouseTablesRefreshSchemaCreate(context.Background(), id, projectId).Table(table).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseTablesAPI.WarehouseTablesRefreshSchemaCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse table. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseTablesRefreshSchemaCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **table** | [**Table**](Table.md) |  | 

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


## WarehouseTablesRetrieve

> Table WarehouseTablesRetrieve(ctx, id, projectId).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse table.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseTablesAPI.WarehouseTablesRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseTablesAPI.WarehouseTablesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseTablesRetrieve`: Table
	fmt.Fprintf(os.Stdout, "Response from `WarehouseTablesAPI.WarehouseTablesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse table. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseTablesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Table**](Table.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseTablesUpdate

> Table WarehouseTablesUpdate(ctx, id, projectId).Table(table).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse table.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	table := *openapiclient.NewTable("Id_example", "Name_example", openapiclient.TableFormatEnum("CSV"), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "UrlPattern_example", *openapiclient.NewCredential("Id_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "AccessKey_example", "AccessSecret_example"), "Columns_example", *openapiclient.NewSimpleExternalDataSourceSerializers("Id_example", time.Now(), NullableInt32(123), "Status_example", openapiclient.SourceTypeEnum("CustomerIO")), "ExternalSchema_example") // Table | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehouseTablesAPI.WarehouseTablesUpdate(context.Background(), id, projectId).Table(table).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseTablesAPI.WarehouseTablesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WarehouseTablesUpdate`: Table
	fmt.Fprintf(os.Stdout, "Response from `WarehouseTablesAPI.WarehouseTablesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse table. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseTablesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **table** | [**Table**](Table.md) |  | 

### Return type

[**Table**](Table.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WarehouseTablesUpdateSchemaCreate

> WarehouseTablesUpdateSchemaCreate(ctx, id, projectId).Table(table).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this data warehouse table.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	table := *openapiclient.NewTable("Id_example", "Name_example", openapiclient.TableFormatEnum("CSV"), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "UrlPattern_example", *openapiclient.NewCredential("Id_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), "AccessKey_example", "AccessSecret_example"), "Columns_example", *openapiclient.NewSimpleExternalDataSourceSerializers("Id_example", time.Now(), NullableInt32(123), "Status_example", openapiclient.SourceTypeEnum("CustomerIO")), "ExternalSchema_example") // Table | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WarehouseTablesAPI.WarehouseTablesUpdateSchemaCreate(context.Background(), id, projectId).Table(table).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehouseTablesAPI.WarehouseTablesUpdateSchemaCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this data warehouse table. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWarehouseTablesUpdateSchemaCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **table** | [**Table**](Table.md) |  | 

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

