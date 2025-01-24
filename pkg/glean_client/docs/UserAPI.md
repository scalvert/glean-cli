# \UserAPI

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Addcredential**](UserAPI.md#Addcredential) | **Post** /addcredential | Create credentials
[**Deletequeryhistory**](UserAPI.md#Deletequeryhistory) | **Post** /deletequeryhistory | Delete query history
[**Editpermissions**](UserAPI.md#Editpermissions) | **Post** /editpermissions | Update permissions
[**Invite**](UserAPI.md#Invite) | **Post** /invite | Send invitation
[**Publicconfig**](UserAPI.md#Publicconfig) | **Post** /publicclientconfig | Read public client configuration
[**Removecredential**](UserAPI.md#Removecredential) | **Post** /removecredential | Delete credentials
[**SupportEmail**](UserAPI.md#SupportEmail) | **Post** /support | Send support email



## Addcredential

> Addcredential(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Create credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	payload := *openapiclient.NewAddCredentialRequest() // AddCredentialRequest | Credential content
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.Addcredential(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.Addcredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddcredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**AddCredentialRequest**](AddCredentialRequest.md) | Credential content | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deletequeryhistory

> DeleteQueryHistoryResponse Deletequeryhistory(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Delete query history



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	payload := *openapiclient.NewDeleteQueryHistoryRequest() // DeleteQueryHistoryRequest | Delete query history request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.Deletequeryhistory(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.Deletequeryhistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Deletequeryhistory`: DeleteQueryHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.Deletequeryhistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletequeryhistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**DeleteQueryHistoryRequest**](DeleteQueryHistoryRequest.md) | Delete query history request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**DeleteQueryHistoryResponse**](DeleteQueryHistoryResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Editpermissions

> EditPermissionsResponse Editpermissions(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Update permissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	payload := *openapiclient.NewEditPermissionsRequest(*openapiclient.NewPermissions()) // EditPermissionsRequest | Permissions
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.Editpermissions(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.Editpermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Editpermissions`: EditPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.Editpermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditpermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**EditPermissionsRequest**](EditPermissionsRequest.md) | Permissions | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**EditPermissionsResponse**](EditPermissionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Invite

> Invite(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Send invitation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	payload := *openapiclient.NewInviteRequest() // InviteRequest | Invite request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.Invite(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.Invite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**InviteRequest**](InviteRequest.md) | Invite request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Publicconfig

> ClientConfig Publicconfig(ctx).XScioActas(xScioActas).Payload(payload).Execute()

Read public client configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)
	payload := *openapiclient.NewPublicConfigRequest() // PublicConfigRequest | Public Config request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.Publicconfig(context.Background()).XScioActas(xScioActas).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.Publicconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Publicconfig`: ClientConfig
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.Publicconfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublicconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 
 **payload** | [**PublicConfigRequest**](PublicConfigRequest.md) | Public Config request | 

### Return type

[**ClientConfig**](ClientConfig.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Removecredential

> Removecredential(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Delete credentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	payload := *openapiclient.NewRemoveCredentialRequest() // RemoveCredentialRequest | Credential content
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.Removecredential(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.Removecredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemovecredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**RemoveCredentialRequest**](RemoveCredentialRequest.md) | Credential content | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SupportEmail

> SupportEmail(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Send support email



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	payload := *openapiclient.NewEmailRequest(openapiclient.CommunicationTemplate("ADMIN_ALERT")) // EmailRequest | Support request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.SupportEmail(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.SupportEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSupportEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**EmailRequest**](EmailRequest.md) | Support request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

