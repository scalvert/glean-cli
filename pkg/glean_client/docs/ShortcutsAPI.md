# \ShortcutsAPI

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Createshortcut**](ShortcutsAPI.md#Createshortcut) | **Post** /createshortcut | Create shortcut
[**Deleteshortcut**](ShortcutsAPI.md#Deleteshortcut) | **Post** /deleteshortcut | Delete shortcut
[**Getshortcut**](ShortcutsAPI.md#Getshortcut) | **Post** /getshortcut | Read shortcut
[**Getsimilarshortcuts**](ShortcutsAPI.md#Getsimilarshortcuts) | **Post** /getsimilarshortcuts | Get similar shortcuts
[**Listshortcuts**](ShortcutsAPI.md#Listshortcuts) | **Post** /listshortcuts | List shortcuts
[**Previewshortcut**](ShortcutsAPI.md#Previewshortcut) | **Post** /previewshortcut | Preview shortcut
[**Updateshortcut**](ShortcutsAPI.md#Updateshortcut) | **Post** /updateshortcut | Update shortcut



## Createshortcut

> CreateShortcutResponse Createshortcut(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Create shortcut



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
	payload := *openapiclient.NewCreateShortcutRequest(*openapiclient.NewShortcutMutableProperties()) // CreateShortcutRequest | CreateShortcut request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShortcutsAPI.Createshortcut(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShortcutsAPI.Createshortcut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Createshortcut`: CreateShortcutResponse
	fmt.Fprintf(os.Stdout, "Response from `ShortcutsAPI.Createshortcut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateshortcutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateShortcutRequest**](CreateShortcutRequest.md) | CreateShortcut request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**CreateShortcutResponse**](CreateShortcutResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deleteshortcut

> Deleteshortcut(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Delete shortcut



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
	payload := *openapiclient.NewDeleteShortcutRequest(int32(123)) // DeleteShortcutRequest | DeleteShortcut request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShortcutsAPI.Deleteshortcut(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShortcutsAPI.Deleteshortcut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteshortcutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**DeleteShortcutRequest**](DeleteShortcutRequest.md) | DeleteShortcut request | 
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


## Getshortcut

> GetShortcutResponse Getshortcut(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Read shortcut



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
	payload := openapiclient.GetShortcutRequest{GetShortcutRequestOneOf: openapiclient.NewGetShortcutRequestOneOf("Alias_example")} // GetShortcutRequest | GetShortcut request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShortcutsAPI.Getshortcut(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShortcutsAPI.Getshortcut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Getshortcut`: GetShortcutResponse
	fmt.Fprintf(os.Stdout, "Response from `ShortcutsAPI.Getshortcut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetshortcutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**GetShortcutRequest**](GetShortcutRequest.md) | GetShortcut request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**GetShortcutResponse**](GetShortcutResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getsimilarshortcuts

> GetSimilarShortcutsResponse Getsimilarshortcuts(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Get similar shortcuts



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
	payload := *openapiclient.NewGetSimilarShortcutsRequest("Alias_example") // GetSimilarShortcutsRequest | GetSimilarShortcuts request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShortcutsAPI.Getsimilarshortcuts(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShortcutsAPI.Getsimilarshortcuts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Getsimilarshortcuts`: GetSimilarShortcutsResponse
	fmt.Fprintf(os.Stdout, "Response from `ShortcutsAPI.Getsimilarshortcuts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetsimilarshortcutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**GetSimilarShortcutsRequest**](GetSimilarShortcutsRequest.md) | GetSimilarShortcuts request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**GetSimilarShortcutsResponse**](GetSimilarShortcutsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Listshortcuts

> ListShortcutsPaginatedResponse Listshortcuts(ctx).Payload(payload).XScioActas(xScioActas).Execute()

List shortcuts



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
	payload := *openapiclient.NewListShortcutsPaginatedRequest(int32(10)) // ListShortcutsPaginatedRequest | Filters, sorters, paging params required for pagination
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShortcutsAPI.Listshortcuts(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShortcutsAPI.Listshortcuts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Listshortcuts`: ListShortcutsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `ShortcutsAPI.Listshortcuts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListshortcutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ListShortcutsPaginatedRequest**](ListShortcutsPaginatedRequest.md) | Filters, sorters, paging params required for pagination | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**ListShortcutsPaginatedResponse**](ListShortcutsPaginatedResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Previewshortcut

> PreviewShortcutResponse Previewshortcut(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Preview shortcut



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
	payload := *openapiclient.NewShortcutMutableProperties() // ShortcutMutableProperties | CreateShortcut request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShortcutsAPI.Previewshortcut(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShortcutsAPI.Previewshortcut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Previewshortcut`: PreviewShortcutResponse
	fmt.Fprintf(os.Stdout, "Response from `ShortcutsAPI.Previewshortcut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewshortcutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ShortcutMutableProperties**](ShortcutMutableProperties.md) | CreateShortcut request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**PreviewShortcutResponse**](PreviewShortcutResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Updateshortcut

> UpdateShortcutResponse Updateshortcut(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Update shortcut



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
	payload := *openapiclient.NewUpdateShortcutRequest(int32(123)) // UpdateShortcutRequest | Shortcut content. Id need to be specified for the shortcut.
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShortcutsAPI.Updateshortcut(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShortcutsAPI.Updateshortcut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Updateshortcut`: UpdateShortcutResponse
	fmt.Fprintf(os.Stdout, "Response from `ShortcutsAPI.Updateshortcut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateshortcutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**UpdateShortcutRequest**](UpdateShortcutRequest.md) | Shortcut content. Id need to be specified for the shortcut. | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**UpdateShortcutResponse**](UpdateShortcutResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

