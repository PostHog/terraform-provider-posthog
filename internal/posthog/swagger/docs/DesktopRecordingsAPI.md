# \DesktopRecordingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsDesktopRecordingsAppendSegmentsCreate**](DesktopRecordingsAPI.md#EnvironmentsDesktopRecordingsAppendSegmentsCreate) | **Post** /api/environments/{project_id}/desktop_recordings/{id}/append_segments/ | 
[**EnvironmentsDesktopRecordingsCreate**](DesktopRecordingsAPI.md#EnvironmentsDesktopRecordingsCreate) | **Post** /api/environments/{project_id}/desktop_recordings/ | 
[**EnvironmentsDesktopRecordingsDestroy**](DesktopRecordingsAPI.md#EnvironmentsDesktopRecordingsDestroy) | **Delete** /api/environments/{project_id}/desktop_recordings/{id}/ | 
[**EnvironmentsDesktopRecordingsList**](DesktopRecordingsAPI.md#EnvironmentsDesktopRecordingsList) | **Get** /api/environments/{project_id}/desktop_recordings/ | 
[**EnvironmentsDesktopRecordingsPartialUpdate**](DesktopRecordingsAPI.md#EnvironmentsDesktopRecordingsPartialUpdate) | **Patch** /api/environments/{project_id}/desktop_recordings/{id}/ | 
[**EnvironmentsDesktopRecordingsRetrieve**](DesktopRecordingsAPI.md#EnvironmentsDesktopRecordingsRetrieve) | **Get** /api/environments/{project_id}/desktop_recordings/{id}/ | 
[**EnvironmentsDesktopRecordingsUpdate**](DesktopRecordingsAPI.md#EnvironmentsDesktopRecordingsUpdate) | **Put** /api/environments/{project_id}/desktop_recordings/{id}/ | 



## EnvironmentsDesktopRecordingsAppendSegmentsCreate

> DesktopRecording EnvironmentsDesktopRecordingsAppendSegmentsCreate(ctx, id, projectId).AppendSegments(appendSegments).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this desktop recording.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	appendSegments := *openapiclient.NewAppendSegments([]openapiclient.TranscriptSegment{*openapiclient.NewTranscriptSegment("Text_example")}) // AppendSegments | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DesktopRecordingsAPI.EnvironmentsDesktopRecordingsAppendSegmentsCreate(context.Background(), id, projectId).AppendSegments(appendSegments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DesktopRecordingsAPI.EnvironmentsDesktopRecordingsAppendSegmentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsDesktopRecordingsAppendSegmentsCreate`: DesktopRecording
	fmt.Fprintf(os.Stdout, "Response from `DesktopRecordingsAPI.EnvironmentsDesktopRecordingsAppendSegmentsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this desktop recording. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsDesktopRecordingsAppendSegmentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appendSegments** | [**AppendSegments**](AppendSegments.md) |  | 

### Return type

[**DesktopRecording**](DesktopRecording.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsDesktopRecordingsCreate

> CreateRecordingResponse EnvironmentsDesktopRecordingsCreate(ctx, projectId).CreateRecordingRequest(createRecordingRequest).Execute()





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
	createRecordingRequest := *openapiclient.NewCreateRecordingRequest() // CreateRecordingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DesktopRecordingsAPI.EnvironmentsDesktopRecordingsCreate(context.Background(), projectId).CreateRecordingRequest(createRecordingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DesktopRecordingsAPI.EnvironmentsDesktopRecordingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsDesktopRecordingsCreate`: CreateRecordingResponse
	fmt.Fprintf(os.Stdout, "Response from `DesktopRecordingsAPI.EnvironmentsDesktopRecordingsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsDesktopRecordingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRecordingRequest** | [**CreateRecordingRequest**](CreateRecordingRequest.md) |  | 

### Return type

[**CreateRecordingResponse**](CreateRecordingResponse.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsDesktopRecordingsDestroy

> EnvironmentsDesktopRecordingsDestroy(ctx, id, projectId).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this desktop recording.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DesktopRecordingsAPI.EnvironmentsDesktopRecordingsDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DesktopRecordingsAPI.EnvironmentsDesktopRecordingsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this desktop recording. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsDesktopRecordingsDestroyRequest struct via the builder pattern


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


## EnvironmentsDesktopRecordingsList

> PaginatedDesktopRecordingList EnvironmentsDesktopRecordingsList(ctx, projectId).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.DesktopRecordingsAPI.EnvironmentsDesktopRecordingsList(context.Background(), projectId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DesktopRecordingsAPI.EnvironmentsDesktopRecordingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsDesktopRecordingsList`: PaginatedDesktopRecordingList
	fmt.Fprintf(os.Stdout, "Response from `DesktopRecordingsAPI.EnvironmentsDesktopRecordingsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsDesktopRecordingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedDesktopRecordingList**](PaginatedDesktopRecordingList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsDesktopRecordingsPartialUpdate

> DesktopRecording EnvironmentsDesktopRecordingsPartialUpdate(ctx, id, projectId).PatchedDesktopRecording(patchedDesktopRecording).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this desktop recording.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedDesktopRecording := *openapiclient.NewPatchedDesktopRecording() // PatchedDesktopRecording |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DesktopRecordingsAPI.EnvironmentsDesktopRecordingsPartialUpdate(context.Background(), id, projectId).PatchedDesktopRecording(patchedDesktopRecording).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DesktopRecordingsAPI.EnvironmentsDesktopRecordingsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsDesktopRecordingsPartialUpdate`: DesktopRecording
	fmt.Fprintf(os.Stdout, "Response from `DesktopRecordingsAPI.EnvironmentsDesktopRecordingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this desktop recording. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsDesktopRecordingsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedDesktopRecording** | [**PatchedDesktopRecording**](PatchedDesktopRecording.md) |  | 

### Return type

[**DesktopRecording**](DesktopRecording.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsDesktopRecordingsRetrieve

> DesktopRecording EnvironmentsDesktopRecordingsRetrieve(ctx, id, projectId).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this desktop recording.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DesktopRecordingsAPI.EnvironmentsDesktopRecordingsRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DesktopRecordingsAPI.EnvironmentsDesktopRecordingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsDesktopRecordingsRetrieve`: DesktopRecording
	fmt.Fprintf(os.Stdout, "Response from `DesktopRecordingsAPI.EnvironmentsDesktopRecordingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this desktop recording. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsDesktopRecordingsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DesktopRecording**](DesktopRecording.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsDesktopRecordingsUpdate

> DesktopRecording EnvironmentsDesktopRecordingsUpdate(ctx, id, projectId).DesktopRecording(desktopRecording).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this desktop recording.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	desktopRecording := *openapiclient.NewDesktopRecording("Id_example", int32(123), NullableInt32(123), "SdkUploadId_example", openapiclient.Platform9aaEnum("zoom"), "TranscriptText_example", time.Now(), time.Now()) // DesktopRecording | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DesktopRecordingsAPI.EnvironmentsDesktopRecordingsUpdate(context.Background(), id, projectId).DesktopRecording(desktopRecording).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DesktopRecordingsAPI.EnvironmentsDesktopRecordingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsDesktopRecordingsUpdate`: DesktopRecording
	fmt.Fprintf(os.Stdout, "Response from `DesktopRecordingsAPI.EnvironmentsDesktopRecordingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this desktop recording. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsDesktopRecordingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **desktopRecording** | [**DesktopRecording**](DesktopRecording.md) |  | 

### Return type

[**DesktopRecording**](DesktopRecording.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

