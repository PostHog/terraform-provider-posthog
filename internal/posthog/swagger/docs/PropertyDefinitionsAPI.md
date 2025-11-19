# \PropertyDefinitionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PropertyDefinitionsDestroy**](PropertyDefinitionsAPI.md#PropertyDefinitionsDestroy) | **Delete** /api/projects/{project_id}/property_definitions/{id}/ | 
[**PropertyDefinitionsList**](PropertyDefinitionsAPI.md#PropertyDefinitionsList) | **Get** /api/projects/{project_id}/property_definitions/ | 
[**PropertyDefinitionsPartialUpdate**](PropertyDefinitionsAPI.md#PropertyDefinitionsPartialUpdate) | **Patch** /api/projects/{project_id}/property_definitions/{id}/ | 
[**PropertyDefinitionsRetrieve**](PropertyDefinitionsAPI.md#PropertyDefinitionsRetrieve) | **Get** /api/projects/{project_id}/property_definitions/{id}/ | 
[**PropertyDefinitionsSeenTogetherRetrieve**](PropertyDefinitionsAPI.md#PropertyDefinitionsSeenTogetherRetrieve) | **Get** /api/projects/{project_id}/property_definitions/seen_together/ | 
[**PropertyDefinitionsUpdate**](PropertyDefinitionsAPI.md#PropertyDefinitionsUpdate) | **Put** /api/projects/{project_id}/property_definitions/{id}/ | 



## PropertyDefinitionsDestroy

> PropertyDefinitionsDestroy(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this property definition.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PropertyDefinitionsAPI.PropertyDefinitionsDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertyDefinitionsAPI.PropertyDefinitionsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this property definition. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertyDefinitionsDestroyRequest struct via the builder pattern


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


## PropertyDefinitionsList

> PaginatedPropertyDefinitionList PropertyDefinitionsList(ctx, projectId).EventNames(eventNames).ExcludeCoreProperties(excludeCoreProperties).ExcludeHidden(excludeHidden).ExcludedProperties(excludedProperties).FilterByEventNames(filterByEventNames).GroupTypeIndex(groupTypeIndex).IsFeatureFlag(isFeatureFlag).IsNumerical(isNumerical).Limit(limit).Offset(offset).Properties(properties).Search(search).Type_(type_).Execute()



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
	eventNames := "eventNames_example" // string | If sent, response value will have `is_seen_on_filtered_events` populated. JSON-encoded (optional)
	excludeCoreProperties := true // bool | Whether to exclude core properties (optional) (default to false)
	excludeHidden := true // bool | Whether to exclude properties marked as hidden (optional) (default to false)
	excludedProperties := "excludedProperties_example" // string | JSON-encoded list of excluded properties (optional)
	filterByEventNames := true // bool | Whether to return only properties for events in `event_names` (optional)
	groupTypeIndex := int32(56) // int32 | What group type is the property for. Only should be set if `type=group` (optional)
	isFeatureFlag := true // bool | Whether to return only (or excluding) feature flag properties (optional)
	isNumerical := true // bool | Whether to return only (or excluding) numerical property definitions (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	properties := "properties_example" // string | Comma-separated list of properties to filter (optional)
	search := "search_example" // string | Searches properties by name (optional)
	type_ := "type__example" // string | What property definitions to return  * `event` - event * `person` - person * `group` - group * `session` - session (optional) (default to "event")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertyDefinitionsAPI.PropertyDefinitionsList(context.Background(), projectId).EventNames(eventNames).ExcludeCoreProperties(excludeCoreProperties).ExcludeHidden(excludeHidden).ExcludedProperties(excludedProperties).FilterByEventNames(filterByEventNames).GroupTypeIndex(groupTypeIndex).IsFeatureFlag(isFeatureFlag).IsNumerical(isNumerical).Limit(limit).Offset(offset).Properties(properties).Search(search).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertyDefinitionsAPI.PropertyDefinitionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PropertyDefinitionsList`: PaginatedPropertyDefinitionList
	fmt.Fprintf(os.Stdout, "Response from `PropertyDefinitionsAPI.PropertyDefinitionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertyDefinitionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventNames** | **string** | If sent, response value will have &#x60;is_seen_on_filtered_events&#x60; populated. JSON-encoded | 
 **excludeCoreProperties** | **bool** | Whether to exclude core properties | [default to false]
 **excludeHidden** | **bool** | Whether to exclude properties marked as hidden | [default to false]
 **excludedProperties** | **string** | JSON-encoded list of excluded properties | 
 **filterByEventNames** | **bool** | Whether to return only properties for events in &#x60;event_names&#x60; | 
 **groupTypeIndex** | **int32** | What group type is the property for. Only should be set if &#x60;type&#x3D;group&#x60; | 
 **isFeatureFlag** | **bool** | Whether to return only (or excluding) feature flag properties | 
 **isNumerical** | **bool** | Whether to return only (or excluding) numerical property definitions | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **properties** | **string** | Comma-separated list of properties to filter | 
 **search** | **string** | Searches properties by name | 
 **type_** | **string** | What property definitions to return  * &#x60;event&#x60; - event * &#x60;person&#x60; - person * &#x60;group&#x60; - group * &#x60;session&#x60; - session | [default to &quot;event&quot;]

### Return type

[**PaginatedPropertyDefinitionList**](PaginatedPropertyDefinitionList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PropertyDefinitionsPartialUpdate

> PropertyDefinition PropertyDefinitionsPartialUpdate(ctx, id, projectId).PatchedPropertyDefinition(patchedPropertyDefinition).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this property definition.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedPropertyDefinition := *openapiclient.NewPatchedPropertyDefinition() // PatchedPropertyDefinition |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertyDefinitionsAPI.PropertyDefinitionsPartialUpdate(context.Background(), id, projectId).PatchedPropertyDefinition(patchedPropertyDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertyDefinitionsAPI.PropertyDefinitionsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PropertyDefinitionsPartialUpdate`: PropertyDefinition
	fmt.Fprintf(os.Stdout, "Response from `PropertyDefinitionsAPI.PropertyDefinitionsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this property definition. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertyDefinitionsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedPropertyDefinition** | [**PatchedPropertyDefinition**](PatchedPropertyDefinition.md) |  | 

### Return type

[**PropertyDefinition**](PropertyDefinition.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PropertyDefinitionsRetrieve

> PropertyDefinition PropertyDefinitionsRetrieve(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this property definition.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertyDefinitionsAPI.PropertyDefinitionsRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertyDefinitionsAPI.PropertyDefinitionsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PropertyDefinitionsRetrieve`: PropertyDefinition
	fmt.Fprintf(os.Stdout, "Response from `PropertyDefinitionsAPI.PropertyDefinitionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this property definition. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertyDefinitionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PropertyDefinition**](PropertyDefinition.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PropertyDefinitionsSeenTogetherRetrieve

> PropertyDefinitionsSeenTogetherRetrieve(ctx, projectId).Execute()





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
	r, err := apiClient.PropertyDefinitionsAPI.PropertyDefinitionsSeenTogetherRetrieve(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertyDefinitionsAPI.PropertyDefinitionsSeenTogetherRetrieve``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPropertyDefinitionsSeenTogetherRetrieveRequest struct via the builder pattern


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


## PropertyDefinitionsUpdate

> PropertyDefinition PropertyDefinitionsUpdate(ctx, id, projectId).PropertyDefinition(propertyDefinition).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this property definition.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	propertyDefinition := *openapiclient.NewPropertyDefinition("Id_example", "Name_example", "IsSeenOnFilteredEvents_example") // PropertyDefinition | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertyDefinitionsAPI.PropertyDefinitionsUpdate(context.Background(), id, projectId).PropertyDefinition(propertyDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertyDefinitionsAPI.PropertyDefinitionsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PropertyDefinitionsUpdate`: PropertyDefinition
	fmt.Fprintf(os.Stdout, "Response from `PropertyDefinitionsAPI.PropertyDefinitionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this property definition. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertyDefinitionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **propertyDefinition** | [**PropertyDefinition**](PropertyDefinition.md) |  | 

### Return type

[**PropertyDefinition**](PropertyDefinition.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

