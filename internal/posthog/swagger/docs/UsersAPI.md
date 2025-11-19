# \UsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCancelEmailChangeRequestPartialUpdate**](UsersAPI.md#UsersCancelEmailChangeRequestPartialUpdate) | **Patch** /api/users/cancel_email_change_request/ | 
[**UsersCancelEmailChangeRequestPartialUpdate_0**](UsersAPI.md#UsersCancelEmailChangeRequestPartialUpdate_0) | **Patch** /api/users/cancel_email_change_request/ | 
[**UsersDestroy**](UsersAPI.md#UsersDestroy) | **Delete** /api/users/{uuid}/ | 
[**UsersDestroy_0**](UsersAPI.md#UsersDestroy_0) | **Delete** /api/users/{uuid}/ | 
[**UsersHedgehogConfigPartialUpdate**](UsersAPI.md#UsersHedgehogConfigPartialUpdate) | **Patch** /api/users/{uuid}/hedgehog_config/ | 
[**UsersHedgehogConfigPartialUpdate_0**](UsersAPI.md#UsersHedgehogConfigPartialUpdate_0) | **Patch** /api/users/{uuid}/hedgehog_config/ | 
[**UsersHedgehogConfigRetrieve**](UsersAPI.md#UsersHedgehogConfigRetrieve) | **Get** /api/users/{uuid}/hedgehog_config/ | 
[**UsersHedgehogConfigRetrieve_0**](UsersAPI.md#UsersHedgehogConfigRetrieve_0) | **Get** /api/users/{uuid}/hedgehog_config/ | 
[**UsersList**](UsersAPI.md#UsersList) | **Get** /api/users/ | 
[**UsersList_0**](UsersAPI.md#UsersList_0) | **Get** /api/users/ | 
[**UsersPartialUpdate**](UsersAPI.md#UsersPartialUpdate) | **Patch** /api/users/{uuid}/ | 
[**UsersPartialUpdate_0**](UsersAPI.md#UsersPartialUpdate_0) | **Patch** /api/users/{uuid}/ | 
[**UsersRequestEmailVerificationCreate**](UsersAPI.md#UsersRequestEmailVerificationCreate) | **Post** /api/users/request_email_verification/ | 
[**UsersRequestEmailVerificationCreate_0**](UsersAPI.md#UsersRequestEmailVerificationCreate_0) | **Post** /api/users/request_email_verification/ | 
[**UsersRetrieve**](UsersAPI.md#UsersRetrieve) | **Get** /api/users/{uuid}/ | 
[**UsersRetrieve_0**](UsersAPI.md#UsersRetrieve_0) | **Get** /api/users/{uuid}/ | 
[**UsersScenePersonalisationCreate**](UsersAPI.md#UsersScenePersonalisationCreate) | **Post** /api/users/{uuid}/scene_personalisation/ | 
[**UsersScenePersonalisationCreate_0**](UsersAPI.md#UsersScenePersonalisationCreate_0) | **Post** /api/users/{uuid}/scene_personalisation/ | 
[**UsersStart2faSetupRetrieve**](UsersAPI.md#UsersStart2faSetupRetrieve) | **Get** /api/users/{uuid}/start_2fa_setup/ | 
[**UsersStart2faSetupRetrieve_0**](UsersAPI.md#UsersStart2faSetupRetrieve_0) | **Get** /api/users/{uuid}/start_2fa_setup/ | 
[**UsersTwoFactorBackupCodesCreate**](UsersAPI.md#UsersTwoFactorBackupCodesCreate) | **Post** /api/users/{uuid}/two_factor_backup_codes/ | 
[**UsersTwoFactorBackupCodesCreate_0**](UsersAPI.md#UsersTwoFactorBackupCodesCreate_0) | **Post** /api/users/{uuid}/two_factor_backup_codes/ | 
[**UsersTwoFactorDisableCreate**](UsersAPI.md#UsersTwoFactorDisableCreate) | **Post** /api/users/{uuid}/two_factor_disable/ | 
[**UsersTwoFactorDisableCreate_0**](UsersAPI.md#UsersTwoFactorDisableCreate_0) | **Post** /api/users/{uuid}/two_factor_disable/ | 
[**UsersTwoFactorStartSetupRetrieve**](UsersAPI.md#UsersTwoFactorStartSetupRetrieve) | **Get** /api/users/{uuid}/two_factor_start_setup/ | 
[**UsersTwoFactorStartSetupRetrieve_0**](UsersAPI.md#UsersTwoFactorStartSetupRetrieve_0) | **Get** /api/users/{uuid}/two_factor_start_setup/ | 
[**UsersTwoFactorStatusRetrieve**](UsersAPI.md#UsersTwoFactorStatusRetrieve) | **Get** /api/users/{uuid}/two_factor_status/ | 
[**UsersTwoFactorStatusRetrieve_0**](UsersAPI.md#UsersTwoFactorStatusRetrieve_0) | **Get** /api/users/{uuid}/two_factor_status/ | 
[**UsersTwoFactorValidateCreate**](UsersAPI.md#UsersTwoFactorValidateCreate) | **Post** /api/users/{uuid}/two_factor_validate/ | 
[**UsersTwoFactorValidateCreate_0**](UsersAPI.md#UsersTwoFactorValidateCreate_0) | **Post** /api/users/{uuid}/two_factor_validate/ | 
[**UsersUpdate**](UsersAPI.md#UsersUpdate) | **Put** /api/users/{uuid}/ | 
[**UsersUpdate_0**](UsersAPI.md#UsersUpdate_0) | **Put** /api/users/{uuid}/ | 
[**UsersValidate2faCreate**](UsersAPI.md#UsersValidate2faCreate) | **Post** /api/users/{uuid}/validate_2fa/ | 
[**UsersValidate2faCreate_0**](UsersAPI.md#UsersValidate2faCreate_0) | **Post** /api/users/{uuid}/validate_2fa/ | 
[**UsersVerifyEmailCreate**](UsersAPI.md#UsersVerifyEmailCreate) | **Post** /api/users/verify_email/ | 
[**UsersVerifyEmailCreate_0**](UsersAPI.md#UsersVerifyEmailCreate_0) | **Post** /api/users/verify_email/ | 



## UsersCancelEmailChangeRequestPartialUpdate

> UsersCancelEmailChangeRequestPartialUpdate(ctx).PatchedUser(patchedUser).Execute()



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
	patchedUser := *openapiclient.NewPatchedUser() // PatchedUser |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersCancelEmailChangeRequestPartialUpdate(context.Background()).PatchedUser(patchedUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersCancelEmailChangeRequestPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersCancelEmailChangeRequestPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedUser** | [**PatchedUser**](PatchedUser.md) |  | 

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


## UsersCancelEmailChangeRequestPartialUpdate_0

> UsersCancelEmailChangeRequestPartialUpdate_0(ctx).PatchedUser(patchedUser).Execute()



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
	patchedUser := *openapiclient.NewPatchedUser() // PatchedUser |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersCancelEmailChangeRequestPartialUpdate_0(context.Background()).PatchedUser(patchedUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersCancelEmailChangeRequestPartialUpdate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersCancelEmailChangeRequestPartialUpdate_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedUser** | [**PatchedUser**](PatchedUser.md) |  | 

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


## UsersDestroy

> UsersDestroy(ctx, uuid).Execute()



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
	r, err := apiClient.UsersAPI.UsersDestroy(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDestroyRequest struct via the builder pattern


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


## UsersDestroy_0

> UsersDestroy_0(ctx, uuid).Execute()



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
	r, err := apiClient.UsersAPI.UsersDestroy_0(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersDestroy_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDestroy_2Request struct via the builder pattern


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


## UsersHedgehogConfigPartialUpdate

> UsersHedgehogConfigPartialUpdate(ctx, uuid).PatchedUser(patchedUser).Execute()



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
	patchedUser := *openapiclient.NewPatchedUser() // PatchedUser |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersHedgehogConfigPartialUpdate(context.Background(), uuid).PatchedUser(patchedUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersHedgehogConfigPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersHedgehogConfigPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUser** | [**PatchedUser**](PatchedUser.md) |  | 

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


## UsersHedgehogConfigPartialUpdate_0

> UsersHedgehogConfigPartialUpdate_0(ctx, uuid).PatchedUser(patchedUser).Execute()



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
	patchedUser := *openapiclient.NewPatchedUser() // PatchedUser |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersHedgehogConfigPartialUpdate_0(context.Background(), uuid).PatchedUser(patchedUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersHedgehogConfigPartialUpdate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersHedgehogConfigPartialUpdate_3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUser** | [**PatchedUser**](PatchedUser.md) |  | 

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


## UsersHedgehogConfigRetrieve

> UsersHedgehogConfigRetrieve(ctx, uuid).Execute()



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
	r, err := apiClient.UsersAPI.UsersHedgehogConfigRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersHedgehogConfigRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersHedgehogConfigRetrieveRequest struct via the builder pattern


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


## UsersHedgehogConfigRetrieve_0

> UsersHedgehogConfigRetrieve_0(ctx, uuid).Execute()



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
	r, err := apiClient.UsersAPI.UsersHedgehogConfigRetrieve_0(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersHedgehogConfigRetrieve_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersHedgehogConfigRetrieve_4Request struct via the builder pattern


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


## UsersList

> PaginatedUserList UsersList(ctx).Email(email).IsStaff(isStaff).Limit(limit).Offset(offset).Execute()



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
	email := "email_example" // string |  (optional)
	isStaff := true // bool |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersList(context.Background()).Email(email).IsStaff(isStaff).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersList`: PaginatedUserList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 
 **isStaff** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedUserList**](PaginatedUserList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersList_0

> PaginatedUserList UsersList_0(ctx).Email(email).IsStaff(isStaff).Limit(limit).Offset(offset).Execute()



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
	email := "email_example" // string |  (optional)
	isStaff := true // bool |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersList_0(context.Background()).Email(email).IsStaff(isStaff).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersList_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersList_0`: PaginatedUserList
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersList_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersList_5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 
 **isStaff** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedUserList**](PaginatedUserList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPartialUpdate

> User UsersPartialUpdate(ctx, uuid).PatchedUser(patchedUser).Execute()



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
	patchedUser := *openapiclient.NewPatchedUser() // PatchedUser |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPartialUpdate(context.Background(), uuid).PatchedUser(patchedUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPartialUpdate`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUser** | [**PatchedUser**](PatchedUser.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPartialUpdate_0

> User UsersPartialUpdate_0(ctx, uuid).PatchedUser(patchedUser).Execute()



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
	patchedUser := *openapiclient.NewPatchedUser() // PatchedUser |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPartialUpdate_0(context.Background(), uuid).PatchedUser(patchedUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPartialUpdate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPartialUpdate_0`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPartialUpdate_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPartialUpdate_6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUser** | [**PatchedUser**](PatchedUser.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRequestEmailVerificationCreate

> UsersRequestEmailVerificationCreate(ctx).User(user).Execute()



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
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersRequestEmailVerificationCreate(context.Background()).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersRequestEmailVerificationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersRequestEmailVerificationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) |  | 

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


## UsersRequestEmailVerificationCreate_0

> UsersRequestEmailVerificationCreate_0(ctx).User(user).Execute()



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
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersRequestEmailVerificationCreate_0(context.Background()).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersRequestEmailVerificationCreate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersRequestEmailVerificationCreate_7Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) |  | 

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


## UsersRetrieve

> User UsersRetrieve(ctx, uuid).Execute()



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
	resp, r, err := apiClient.UsersAPI.UsersRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersRetrieve`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRetrieve_0

> User UsersRetrieve_0(ctx, uuid).Execute()



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
	resp, r, err := apiClient.UsersAPI.UsersRetrieve_0(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersRetrieve_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersRetrieve_0`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersRetrieve_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersRetrieve_8Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersScenePersonalisationCreate

> UsersScenePersonalisationCreate(ctx, uuid).User(user).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersScenePersonalisationCreate(context.Background(), uuid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersScenePersonalisationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersScenePersonalisationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

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


## UsersScenePersonalisationCreate_0

> UsersScenePersonalisationCreate_0(ctx, uuid).User(user).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersScenePersonalisationCreate_0(context.Background(), uuid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersScenePersonalisationCreate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersScenePersonalisationCreate_9Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

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


## UsersStart2faSetupRetrieve

> UsersStart2faSetupRetrieve(ctx, uuid).Execute()



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
	r, err := apiClient.UsersAPI.UsersStart2faSetupRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersStart2faSetupRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersStart2faSetupRetrieveRequest struct via the builder pattern


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


## UsersStart2faSetupRetrieve_0

> UsersStart2faSetupRetrieve_0(ctx, uuid).Execute()



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
	r, err := apiClient.UsersAPI.UsersStart2faSetupRetrieve_0(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersStart2faSetupRetrieve_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersStart2faSetupRetrieve_10Request struct via the builder pattern


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


## UsersTwoFactorBackupCodesCreate

> UsersTwoFactorBackupCodesCreate(ctx, uuid).User(user).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersTwoFactorBackupCodesCreate(context.Background(), uuid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTwoFactorBackupCodesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTwoFactorBackupCodesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

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


## UsersTwoFactorBackupCodesCreate_0

> UsersTwoFactorBackupCodesCreate_0(ctx, uuid).User(user).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersTwoFactorBackupCodesCreate_0(context.Background(), uuid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTwoFactorBackupCodesCreate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTwoFactorBackupCodesCreate_11Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

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


## UsersTwoFactorDisableCreate

> UsersTwoFactorDisableCreate(ctx, uuid).User(user).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersTwoFactorDisableCreate(context.Background(), uuid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTwoFactorDisableCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTwoFactorDisableCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

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


## UsersTwoFactorDisableCreate_0

> UsersTwoFactorDisableCreate_0(ctx, uuid).User(user).Execute()





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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersTwoFactorDisableCreate_0(context.Background(), uuid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTwoFactorDisableCreate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTwoFactorDisableCreate_12Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

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


## UsersTwoFactorStartSetupRetrieve

> UsersTwoFactorStartSetupRetrieve(ctx, uuid).Execute()



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
	r, err := apiClient.UsersAPI.UsersTwoFactorStartSetupRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTwoFactorStartSetupRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTwoFactorStartSetupRetrieveRequest struct via the builder pattern


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


## UsersTwoFactorStartSetupRetrieve_0

> UsersTwoFactorStartSetupRetrieve_0(ctx, uuid).Execute()



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
	r, err := apiClient.UsersAPI.UsersTwoFactorStartSetupRetrieve_0(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTwoFactorStartSetupRetrieve_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTwoFactorStartSetupRetrieve_13Request struct via the builder pattern


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


## UsersTwoFactorStatusRetrieve

> UsersTwoFactorStatusRetrieve(ctx, uuid).Execute()





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
	r, err := apiClient.UsersAPI.UsersTwoFactorStatusRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTwoFactorStatusRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTwoFactorStatusRetrieveRequest struct via the builder pattern


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


## UsersTwoFactorStatusRetrieve_0

> UsersTwoFactorStatusRetrieve_0(ctx, uuid).Execute()





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
	r, err := apiClient.UsersAPI.UsersTwoFactorStatusRetrieve_0(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTwoFactorStatusRetrieve_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTwoFactorStatusRetrieve_14Request struct via the builder pattern


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


## UsersTwoFactorValidateCreate

> UsersTwoFactorValidateCreate(ctx, uuid).User(user).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersTwoFactorValidateCreate(context.Background(), uuid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTwoFactorValidateCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTwoFactorValidateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

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


## UsersTwoFactorValidateCreate_0

> UsersTwoFactorValidateCreate_0(ctx, uuid).User(user).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersTwoFactorValidateCreate_0(context.Background(), uuid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersTwoFactorValidateCreate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersTwoFactorValidateCreate_15Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

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


## UsersUpdate

> User UsersUpdate(ctx, uuid).User(user).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUpdate(context.Background(), uuid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUpdate`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdate_0

> User UsersUpdate_0(ctx, uuid).User(user).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUpdate_0(context.Background(), uuid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUpdate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUpdate_0`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUpdate_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdate_16Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersValidate2faCreate

> UsersValidate2faCreate(ctx, uuid).User(user).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersValidate2faCreate(context.Background(), uuid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersValidate2faCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersValidate2faCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

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


## UsersValidate2faCreate_0

> UsersValidate2faCreate_0(ctx, uuid).User(user).Execute()



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
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersValidate2faCreate_0(context.Background(), uuid).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersValidate2faCreate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersValidate2faCreate_17Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

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


## UsersVerifyEmailCreate

> UsersVerifyEmailCreate(ctx).User(user).Execute()



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
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersVerifyEmailCreate(context.Background()).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersVerifyEmailCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersVerifyEmailCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) |  | 

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


## UsersVerifyEmailCreate_0

> UsersVerifyEmailCreate_0(ctx).User(user).Execute()



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
	user := *openapiclient.NewUser(time.Now(), "Uuid_example", "DistinctId_example", "Email_example", "PendingEmail_example", false, false, int32(123), false, "IsImpersonatedUntil_example", "SensitiveSessionExpiresAt_example", *openapiclient.NewTeamBasic(int32(123), "Uuid_example", "Organization_example", int64(123), "ApiToken_example", "Name_example", false, interface{}(123), false, false, openapiclient.TimezoneEnum("Africa/Abidjan"), false), *openapiclient.NewOrganization("Id_example", "Name_example", "Slug_example", time.Now(), time.Now(), "TODO", openapiclient.PluginsAccessLevelEnum(0), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}, []interface{}{nil}, "Metadata_example", "CustomerId_example", "MemberCount_example"), []openapiclient.OrganizationBasic{*openapiclient.NewOrganizationBasic("Id_example", "Name_example", "Slug_example", "LogoMediaId_example", "TODO")}, "Password_example", false, false, false, []openapiclient.ScenePersonalisationBasic{*openapiclient.NewScenePersonalisationBasic("Scene_example")}) // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersVerifyEmailCreate_0(context.Background()).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersVerifyEmailCreate_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersVerifyEmailCreate_18Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) |  | 

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

