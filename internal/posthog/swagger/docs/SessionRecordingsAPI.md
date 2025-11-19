# \SessionRecordingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionRecordingsDestroy**](SessionRecordingsAPI.md#SessionRecordingsDestroy) | **Delete** /api/projects/{project_id}/session_recordings/{id}/ | 
[**SessionRecordingsList**](SessionRecordingsAPI.md#SessionRecordingsList) | **Get** /api/projects/{project_id}/session_recordings/ | 
[**SessionRecordingsPartialUpdate**](SessionRecordingsAPI.md#SessionRecordingsPartialUpdate) | **Patch** /api/projects/{project_id}/session_recordings/{id}/ | 
[**SessionRecordingsRetrieve**](SessionRecordingsAPI.md#SessionRecordingsRetrieve) | **Get** /api/projects/{project_id}/session_recordings/{id}/ | 
[**SessionRecordingsSharingList**](SessionRecordingsAPI.md#SessionRecordingsSharingList) | **Get** /api/projects/{project_id}/session_recordings/{recording_id}/sharing/ | 
[**SessionRecordingsSharingPasswordsCreate**](SessionRecordingsAPI.md#SessionRecordingsSharingPasswordsCreate) | **Post** /api/projects/{project_id}/session_recordings/{recording_id}/sharing/passwords/ | 
[**SessionRecordingsSharingPasswordsDestroy**](SessionRecordingsAPI.md#SessionRecordingsSharingPasswordsDestroy) | **Delete** /api/projects/{project_id}/session_recordings/{recording_id}/sharing/passwords/{password_id}/ | 
[**SessionRecordingsSharingRefreshCreate**](SessionRecordingsAPI.md#SessionRecordingsSharingRefreshCreate) | **Post** /api/projects/{project_id}/session_recordings/{recording_id}/sharing/refresh/ | 
[**SessionRecordingsUpdate**](SessionRecordingsAPI.md#SessionRecordingsUpdate) | **Put** /api/projects/{project_id}/session_recordings/{id}/ | 



## SessionRecordingsDestroy

> SessionRecordingsDestroy(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this session recording.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionRecordingsAPI.SessionRecordingsDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingsAPI.SessionRecordingsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this session recording. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingsDestroyRequest struct via the builder pattern


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


## SessionRecordingsList

> PaginatedSessionRecordingList SessionRecordingsList(ctx, projectId).Limit(limit).Offset(offset).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionRecordingsAPI.SessionRecordingsList(context.Background(), projectId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingsAPI.SessionRecordingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionRecordingsList`: PaginatedSessionRecordingList
	fmt.Fprintf(os.Stdout, "Response from `SessionRecordingsAPI.SessionRecordingsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedSessionRecordingList**](PaginatedSessionRecordingList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionRecordingsPartialUpdate

> SessionRecording SessionRecordingsPartialUpdate(ctx, id, projectId).PatchedSessionRecording(patchedSessionRecording).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this session recording.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedSessionRecording := *openapiclient.NewPatchedSessionRecording() // PatchedSessionRecording |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionRecordingsAPI.SessionRecordingsPartialUpdate(context.Background(), id, projectId).PatchedSessionRecording(patchedSessionRecording).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingsAPI.SessionRecordingsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionRecordingsPartialUpdate`: SessionRecording
	fmt.Fprintf(os.Stdout, "Response from `SessionRecordingsAPI.SessionRecordingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this session recording. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedSessionRecording** | [**PatchedSessionRecording**](PatchedSessionRecording.md) |  | 

### Return type

[**SessionRecording**](SessionRecording.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionRecordingsRetrieve

> SessionRecording SessionRecordingsRetrieve(ctx, id, projectId).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this session recording.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionRecordingsAPI.SessionRecordingsRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingsAPI.SessionRecordingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionRecordingsRetrieve`: SessionRecording
	fmt.Fprintf(os.Stdout, "Response from `SessionRecordingsAPI.SessionRecordingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this session recording. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SessionRecording**](SessionRecording.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionRecordingsSharingList

> []SharingConfiguration SessionRecordingsSharingList(ctx, projectId, recordingId).Execute()



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
	recordingId := "recordingId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionRecordingsAPI.SessionRecordingsSharingList(context.Background(), projectId, recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingsAPI.SessionRecordingsSharingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionRecordingsSharingList`: []SharingConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SessionRecordingsAPI.SessionRecordingsSharingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**recordingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingsSharingListRequest struct via the builder pattern


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


## SessionRecordingsSharingPasswordsCreate

> SharingConfiguration SessionRecordingsSharingPasswordsCreate(ctx, projectId, recordingId).SharingConfiguration(sharingConfiguration).Execute()





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
	recordingId := "recordingId_example" // string | 
	sharingConfiguration := *openapiclient.NewSharingConfiguration(time.Now(), "AccessToken_example", "SharePasswords_example") // SharingConfiguration |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionRecordingsAPI.SessionRecordingsSharingPasswordsCreate(context.Background(), projectId, recordingId).SharingConfiguration(sharingConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingsAPI.SessionRecordingsSharingPasswordsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionRecordingsSharingPasswordsCreate`: SharingConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SessionRecordingsAPI.SessionRecordingsSharingPasswordsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**recordingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingsSharingPasswordsCreateRequest struct via the builder pattern


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


## SessionRecordingsSharingPasswordsDestroy

> SessionRecordingsSharingPasswordsDestroy(ctx, passwordId, projectId, recordingId).Execute()





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
	passwordId := "passwordId_example" // string | 
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	recordingId := "recordingId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionRecordingsAPI.SessionRecordingsSharingPasswordsDestroy(context.Background(), passwordId, projectId, recordingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingsAPI.SessionRecordingsSharingPasswordsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordId** | **string** |  | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**recordingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingsSharingPasswordsDestroyRequest struct via the builder pattern


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


## SessionRecordingsSharingRefreshCreate

> SharingConfiguration SessionRecordingsSharingRefreshCreate(ctx, projectId, recordingId).SharingConfiguration(sharingConfiguration).Execute()



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
	recordingId := "recordingId_example" // string | 
	sharingConfiguration := *openapiclient.NewSharingConfiguration(time.Now(), "AccessToken_example", "SharePasswords_example") // SharingConfiguration |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionRecordingsAPI.SessionRecordingsSharingRefreshCreate(context.Background(), projectId, recordingId).SharingConfiguration(sharingConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingsAPI.SessionRecordingsSharingRefreshCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionRecordingsSharingRefreshCreate`: SharingConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SessionRecordingsAPI.SessionRecordingsSharingRefreshCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**recordingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingsSharingRefreshCreateRequest struct via the builder pattern


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


## SessionRecordingsUpdate

> SessionRecording SessionRecordingsUpdate(ctx, id, projectId).SessionRecording(sessionRecording).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this session recording.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	sessionRecording := *openapiclient.NewSessionRecording("Id_example", "DistinctId_example", false, []string{"Viewers_example"}, int32(123), NullableInt32(123), NullableInt32(123), time.Now(), time.Now(), NullableInt32(123), NullableInt32(123), NullableInt32(123), NullableInt32(123), NullableInt32(123), NullableInt32(123), "StartUrl_example", "Storage_example", NullableInt32(123), "ExpiryTime_example", "RecordingTtl_example", "SnapshotSource_example", false, NullableFloat64(123)) // SessionRecording |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionRecordingsAPI.SessionRecordingsUpdate(context.Background(), id, projectId).SessionRecording(sessionRecording).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionRecordingsAPI.SessionRecordingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionRecordingsUpdate`: SessionRecording
	fmt.Fprintf(os.Stdout, "Response from `SessionRecordingsAPI.SessionRecordingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this session recording. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionRecordingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sessionRecording** | [**SessionRecording**](SessionRecording.md) |  | 

### Return type

[**SessionRecording**](SessionRecording.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

