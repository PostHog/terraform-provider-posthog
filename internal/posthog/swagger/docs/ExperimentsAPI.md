# \ExperimentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExperimentsCreate**](ExperimentsAPI.md#ExperimentsCreate) | **Post** /api/projects/{project_id}/experiments/ | 
[**ExperimentsCreateExposureCohortForExperimentCreate**](ExperimentsAPI.md#ExperimentsCreateExposureCohortForExperimentCreate) | **Post** /api/projects/{project_id}/experiments/{id}/create_exposure_cohort_for_experiment/ | 
[**ExperimentsDestroy**](ExperimentsAPI.md#ExperimentsDestroy) | **Delete** /api/projects/{project_id}/experiments/{id}/ | 
[**ExperimentsDuplicateCreate**](ExperimentsAPI.md#ExperimentsDuplicateCreate) | **Post** /api/projects/{project_id}/experiments/{id}/duplicate/ | 
[**ExperimentsEligibleFeatureFlagsRetrieve**](ExperimentsAPI.md#ExperimentsEligibleFeatureFlagsRetrieve) | **Get** /api/projects/{project_id}/experiments/eligible_feature_flags/ | 
[**ExperimentsList**](ExperimentsAPI.md#ExperimentsList) | **Get** /api/projects/{project_id}/experiments/ | 
[**ExperimentsPartialUpdate**](ExperimentsAPI.md#ExperimentsPartialUpdate) | **Patch** /api/projects/{project_id}/experiments/{id}/ | 
[**ExperimentsRecalculateTimeseriesCreate**](ExperimentsAPI.md#ExperimentsRecalculateTimeseriesCreate) | **Post** /api/projects/{project_id}/experiments/{id}/recalculate_timeseries/ | 
[**ExperimentsRequiresFlagImplementationRetrieve**](ExperimentsAPI.md#ExperimentsRequiresFlagImplementationRetrieve) | **Get** /api/projects/{project_id}/experiments/requires_flag_implementation/ | 
[**ExperimentsRetrieve**](ExperimentsAPI.md#ExperimentsRetrieve) | **Get** /api/projects/{project_id}/experiments/{id}/ | 
[**ExperimentsStatsRetrieve**](ExperimentsAPI.md#ExperimentsStatsRetrieve) | **Get** /api/projects/{project_id}/experiments/stats/ | 
[**ExperimentsTimeseriesResultsRetrieve**](ExperimentsAPI.md#ExperimentsTimeseriesResultsRetrieve) | **Get** /api/projects/{project_id}/experiments/{id}/timeseries_results/ | 
[**ExperimentsUpdate**](ExperimentsAPI.md#ExperimentsUpdate) | **Put** /api/projects/{project_id}/experiments/{id}/ | 



## ExperimentsCreate

> Experiment ExperimentsCreate(ctx, projectId).Experiment(experiment).Execute()



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
	experiment := *openapiclient.NewExperiment(int32(123), "Name_example", "FeatureFlagKey_example", *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), *openapiclient.NewExperimentHoldout(int32(123), "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now()), NullableInt32(123), []openapiclient.ExperimentToSavedMetric{*openapiclient.NewExperimentToSavedMetric(int32(123), int32(123), int32(123), time.Now(), interface{}(123), "Name_example")}, *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now(), "UserAccessLevel_example") // Experiment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ExperimentsCreate(context.Background(), projectId).Experiment(experiment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ExperimentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentsCreate`: Experiment
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ExperimentsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experiment** | [**Experiment**](Experiment.md) |  | 

### Return type

[**Experiment**](Experiment.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExperimentsCreateExposureCohortForExperimentCreate

> ExperimentsCreateExposureCohortForExperimentCreate(ctx, id, projectId).Experiment(experiment).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this experiment.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	experiment := *openapiclient.NewExperiment(int32(123), "Name_example", "FeatureFlagKey_example", *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), *openapiclient.NewExperimentHoldout(int32(123), "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now()), NullableInt32(123), []openapiclient.ExperimentToSavedMetric{*openapiclient.NewExperimentToSavedMetric(int32(123), int32(123), int32(123), time.Now(), interface{}(123), "Name_example")}, *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now(), "UserAccessLevel_example") // Experiment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExperimentsAPI.ExperimentsCreateExposureCohortForExperimentCreate(context.Background(), id, projectId).Experiment(experiment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ExperimentsCreateExposureCohortForExperimentCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentsCreateExposureCohortForExperimentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **experiment** | [**Experiment**](Experiment.md) |  | 

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


## ExperimentsDestroy

> ExperimentsDestroy(ctx, id, projectId).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this experiment.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExperimentsAPI.ExperimentsDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ExperimentsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentsDestroyRequest struct via the builder pattern


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


## ExperimentsDuplicateCreate

> ExperimentsDuplicateCreate(ctx, id, projectId).Experiment(experiment).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this experiment.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	experiment := *openapiclient.NewExperiment(int32(123), "Name_example", "FeatureFlagKey_example", *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), *openapiclient.NewExperimentHoldout(int32(123), "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now()), NullableInt32(123), []openapiclient.ExperimentToSavedMetric{*openapiclient.NewExperimentToSavedMetric(int32(123), int32(123), int32(123), time.Now(), interface{}(123), "Name_example")}, *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now(), "UserAccessLevel_example") // Experiment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExperimentsAPI.ExperimentsDuplicateCreate(context.Background(), id, projectId).Experiment(experiment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ExperimentsDuplicateCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentsDuplicateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **experiment** | [**Experiment**](Experiment.md) |  | 

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


## ExperimentsEligibleFeatureFlagsRetrieve

> ExperimentsEligibleFeatureFlagsRetrieve(ctx, projectId).Execute()





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
	r, err := apiClient.ExperimentsAPI.ExperimentsEligibleFeatureFlagsRetrieve(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ExperimentsEligibleFeatureFlagsRetrieve``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExperimentsEligibleFeatureFlagsRetrieveRequest struct via the builder pattern


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


## ExperimentsList

> PaginatedExperimentList ExperimentsList(ctx, projectId).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.ExperimentsAPI.ExperimentsList(context.Background(), projectId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ExperimentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentsList`: PaginatedExperimentList
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ExperimentsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedExperimentList**](PaginatedExperimentList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExperimentsPartialUpdate

> Experiment ExperimentsPartialUpdate(ctx, id, projectId).PatchedExperiment(patchedExperiment).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this experiment.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedExperiment := *openapiclient.NewPatchedExperiment() // PatchedExperiment |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ExperimentsPartialUpdate(context.Background(), id, projectId).PatchedExperiment(patchedExperiment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ExperimentsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentsPartialUpdate`: Experiment
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ExperimentsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedExperiment** | [**PatchedExperiment**](PatchedExperiment.md) |  | 

### Return type

[**Experiment**](Experiment.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExperimentsRecalculateTimeseriesCreate

> ExperimentsRecalculateTimeseriesCreate(ctx, id, projectId).Experiment(experiment).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this experiment.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	experiment := *openapiclient.NewExperiment(int32(123), "Name_example", "FeatureFlagKey_example", *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), *openapiclient.NewExperimentHoldout(int32(123), "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now()), NullableInt32(123), []openapiclient.ExperimentToSavedMetric{*openapiclient.NewExperimentToSavedMetric(int32(123), int32(123), int32(123), time.Now(), interface{}(123), "Name_example")}, *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now(), "UserAccessLevel_example") // Experiment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExperimentsAPI.ExperimentsRecalculateTimeseriesCreate(context.Background(), id, projectId).Experiment(experiment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ExperimentsRecalculateTimeseriesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentsRecalculateTimeseriesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **experiment** | [**Experiment**](Experiment.md) |  | 

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


## ExperimentsRequiresFlagImplementationRetrieve

> ExperimentsRequiresFlagImplementationRetrieve(ctx, projectId).Execute()



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
	r, err := apiClient.ExperimentsAPI.ExperimentsRequiresFlagImplementationRetrieve(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ExperimentsRequiresFlagImplementationRetrieve``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExperimentsRequiresFlagImplementationRetrieveRequest struct via the builder pattern


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


## ExperimentsRetrieve

> Experiment ExperimentsRetrieve(ctx, id, projectId).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this experiment.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ExperimentsRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ExperimentsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentsRetrieve`: Experiment
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ExperimentsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Experiment**](Experiment.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExperimentsStatsRetrieve

> ExperimentsStatsRetrieve(ctx, projectId).Execute()





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
	r, err := apiClient.ExperimentsAPI.ExperimentsStatsRetrieve(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ExperimentsStatsRetrieve``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExperimentsStatsRetrieveRequest struct via the builder pattern


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


## ExperimentsTimeseriesResultsRetrieve

> ExperimentsTimeseriesResultsRetrieve(ctx, id, projectId).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this experiment.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExperimentsAPI.ExperimentsTimeseriesResultsRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ExperimentsTimeseriesResultsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentsTimeseriesResultsRetrieveRequest struct via the builder pattern


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


## ExperimentsUpdate

> Experiment ExperimentsUpdate(ctx, id, projectId).Experiment(experiment).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this experiment.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	experiment := *openapiclient.NewExperiment(int32(123), "Name_example", "FeatureFlagKey_example", *openapiclient.NewMinimalFeatureFlag(int32(123), int32(123), "Key_example", []string{"EvaluationTags_example"}), *openapiclient.NewExperimentHoldout(int32(123), "Name_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now()), NullableInt32(123), []openapiclient.ExperimentToSavedMetric{*openapiclient.NewExperimentToSavedMetric(int32(123), int32(123), int32(123), time.Now(), interface{}(123), "Name_example")}, *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), time.Now(), time.Now(), "UserAccessLevel_example") // Experiment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentsAPI.ExperimentsUpdate(context.Background(), id, projectId).Experiment(experiment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsAPI.ExperimentsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExperimentsUpdate`: Experiment
	fmt.Fprintf(os.Stdout, "Response from `ExperimentsAPI.ExperimentsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this experiment. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExperimentsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **experiment** | [**Experiment**](Experiment.md) |  | 

### Return type

[**Experiment**](Experiment.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

