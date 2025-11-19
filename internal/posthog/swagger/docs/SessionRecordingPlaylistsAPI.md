# \SessionRecordingPlaylistsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionRecordingPlaylistsCreate**](SessionRecordingPlaylistsAPI.md#SessionRecordingPlaylistsCreate) | **Post** /api/projects/{project_id}/session_recording_playlists/ | 
[**SessionRecordingPlaylistsDestroy**](SessionRecordingPlaylistsAPI.md#SessionRecordingPlaylistsDestroy) | **Delete** /api/projects/{project_id}/session_recording_playlists/{short_id}/ | 
[**SessionRecordingPlaylistsList**](SessionRecordingPlaylistsAPI.md#SessionRecordingPlaylistsList) | **Get** /api/projects/{project_id}/session_recording_playlists/ | 
[**SessionRecordingPlaylistsPartialUpdate**](SessionRecordingPlaylistsAPI.md#SessionRecordingPlaylistsPartialUpdate) | **Patch** /api/projects/{project_id}/session_recording_playlists/{short_id}/ | 
[**SessionRecordingPlaylistsRecordingsCreate**](SessionRecordingPlaylistsAPI.md#SessionRecordingPlaylistsRecordingsCreate) | **Post** /api/projects/{project_id}/session_recording_playlists/{short_id}/recordings/{session_recording_id}/ | 
[**SessionRecordingPlaylistsRecordingsDestroy**](SessionRecordingPlaylistsAPI.md#SessionRecordingPlaylistsRecordingsDestroy) | **Delete** /api/projects/{project_id}/session_recording_playlists/{short_id}/recordings/{session_recording_id}/ | 
[**SessionRecordingPlaylistsRecordingsRetrieve**](SessionRecordingPlaylistsAPI.md#SessionRecordingPlaylistsRecordingsRetrieve) | **Get** /api/projects/{project_id}/session_recording_playlists/{short_id}/recordings/ | 
[**SessionRecordingPlaylistsRetrieve**](SessionRecordingPlaylistsAPI.md#SessionRecordingPlaylistsRetrieve) | **Get** /api/projects/{project_id}/session_recording_playlists/{short_id}/ | 
[**SessionRecordingPlaylistsUpdate**](SessionRecordingPlaylistsAPI.md#SessionRecordingPlaylistsUpdate) | **Put** /api/projects/{project_id}/session_recording_playlists/{short_id}/ | 



## SessionRecordingPlaylistsCreate

> SessionRecordingPlaylist SessionRecordingPlaylistsCreate(ctx, projectId).SessionRecordingPlaylist(sessionRecordingPlaylist).Execute()



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
	sessionRecordingPlaylist := *openapiclient.NewSessionRecordingPlaylist(int32(123), "ShortId_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), map[string]map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue{"key": map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue{"key": openapiclient.PatchedSessionRecordingPlaylist_recordings_counts_value_value{Bool: new(bool)}}}, "Type_example", false) // SessionRecordingPlaylist |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsCreate(context.Background(), projectId).SessionRecordingPlaylist(sessionRecordingPlaylist).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionRecordingPlaylistsCreate`: SessionRecordingPlaylist
	fmt.Fprintf(os.Stdout, "Response from `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingPlaylistsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionRecordingPlaylist** | [**SessionRecordingPlaylist**](SessionRecordingPlaylist.md) |  | 

### Return type

[**SessionRecordingPlaylist**](SessionRecordingPlaylist.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionRecordingPlaylistsDestroy

> SessionRecordingPlaylistsDestroy(ctx, projectId, shortId).Execute()





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
	shortId := "shortId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsDestroy(context.Background(), projectId, shortId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**shortId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingPlaylistsDestroyRequest struct via the builder pattern


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


## SessionRecordingPlaylistsList

> PaginatedSessionRecordingPlaylistList SessionRecordingPlaylistsList(ctx, projectId).CreatedBy(createdBy).Limit(limit).Offset(offset).ShortId(shortId).Execute()





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
	createdBy := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	shortId := "shortId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsList(context.Background(), projectId).CreatedBy(createdBy).Limit(limit).Offset(offset).ShortId(shortId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionRecordingPlaylistsList`: PaginatedSessionRecordingPlaylistList
	fmt.Fprintf(os.Stdout, "Response from `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingPlaylistsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdBy** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **shortId** | **string** |  | 

### Return type

[**PaginatedSessionRecordingPlaylistList**](PaginatedSessionRecordingPlaylistList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionRecordingPlaylistsPartialUpdate

> SessionRecordingPlaylist SessionRecordingPlaylistsPartialUpdate(ctx, projectId, shortId).PatchedSessionRecordingPlaylist(patchedSessionRecordingPlaylist).Execute()



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
	shortId := "shortId_example" // string | 
	patchedSessionRecordingPlaylist := *openapiclient.NewPatchedSessionRecordingPlaylist() // PatchedSessionRecordingPlaylist |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsPartialUpdate(context.Background(), projectId, shortId).PatchedSessionRecordingPlaylist(patchedSessionRecordingPlaylist).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionRecordingPlaylistsPartialUpdate`: SessionRecordingPlaylist
	fmt.Fprintf(os.Stdout, "Response from `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**shortId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingPlaylistsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedSessionRecordingPlaylist** | [**PatchedSessionRecordingPlaylist**](PatchedSessionRecordingPlaylist.md) |  | 

### Return type

[**SessionRecordingPlaylist**](SessionRecordingPlaylist.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionRecordingPlaylistsRecordingsCreate

> SessionRecordingPlaylistsRecordingsCreate(ctx, projectId, sessionRecordingId, shortId).SessionRecordingPlaylist(sessionRecordingPlaylist).Execute()



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
	sessionRecordingId := "sessionRecordingId_example" // string | 
	shortId := "shortId_example" // string | 
	sessionRecordingPlaylist := *openapiclient.NewSessionRecordingPlaylist(int32(123), "ShortId_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), map[string]map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue{"key": map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue{"key": openapiclient.PatchedSessionRecordingPlaylist_recordings_counts_value_value{Bool: new(bool)}}}, "Type_example", false) // SessionRecordingPlaylist |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsRecordingsCreate(context.Background(), projectId, sessionRecordingId, shortId).SessionRecordingPlaylist(sessionRecordingPlaylist).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsRecordingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**sessionRecordingId** | **string** |  | 
**shortId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingPlaylistsRecordingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sessionRecordingPlaylist** | [**SessionRecordingPlaylist**](SessionRecordingPlaylist.md) |  | 

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


## SessionRecordingPlaylistsRecordingsDestroy

> SessionRecordingPlaylistsRecordingsDestroy(ctx, projectId, sessionRecordingId, shortId).Execute()



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
	sessionRecordingId := "sessionRecordingId_example" // string | 
	shortId := "shortId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsRecordingsDestroy(context.Background(), projectId, sessionRecordingId, shortId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsRecordingsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**sessionRecordingId** | **string** |  | 
**shortId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingPlaylistsRecordingsDestroyRequest struct via the builder pattern


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


## SessionRecordingPlaylistsRecordingsRetrieve

> SessionRecordingPlaylistsRecordingsRetrieve(ctx, projectId, shortId).Execute()



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
	shortId := "shortId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsRecordingsRetrieve(context.Background(), projectId, shortId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsRecordingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**shortId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingPlaylistsRecordingsRetrieveRequest struct via the builder pattern


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


## SessionRecordingPlaylistsRetrieve

> SessionRecordingPlaylist SessionRecordingPlaylistsRetrieve(ctx, projectId, shortId).Execute()



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
	shortId := "shortId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsRetrieve(context.Background(), projectId, shortId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionRecordingPlaylistsRetrieve`: SessionRecordingPlaylist
	fmt.Fprintf(os.Stdout, "Response from `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**shortId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingPlaylistsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SessionRecordingPlaylist**](SessionRecordingPlaylist.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionRecordingPlaylistsUpdate

> SessionRecordingPlaylist SessionRecordingPlaylistsUpdate(ctx, projectId, shortId).SessionRecordingPlaylist(sessionRecordingPlaylist).Execute()



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
	shortId := "shortId_example" // string | 
	sessionRecordingPlaylist := *openapiclient.NewSessionRecordingPlaylist(int32(123), "ShortId_example", time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), map[string]map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue{"key": map[string]PatchedSessionRecordingPlaylistRecordingsCountsValueValue{"key": openapiclient.PatchedSessionRecordingPlaylist_recordings_counts_value_value{Bool: new(bool)}}}, "Type_example", false) // SessionRecordingPlaylist |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsUpdate(context.Background(), projectId, shortId).SessionRecordingPlaylist(sessionRecordingPlaylist).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionRecordingPlaylistsUpdate`: SessionRecordingPlaylist
	fmt.Fprintf(os.Stdout, "Response from `SessionRecordingPlaylistsAPI.SessionRecordingPlaylistsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**shortId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingPlaylistsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sessionRecordingPlaylist** | [**SessionRecordingPlaylist**](SessionRecordingPlaylist.md) |  | 

### Return type

[**SessionRecordingPlaylist**](SessionRecordingPlaylist.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

