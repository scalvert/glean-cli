# \DisplayableListsAPI

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Createdisplayablelists**](DisplayableListsAPI.md#Createdisplayablelists) | **Post** /createdisplayablelists | Create displayable lists
[**Deletedisplayablelists**](DisplayableListsAPI.md#Deletedisplayablelists) | **Post** /deletedisplayablelists | Delete displayable lists
[**Getdisplayablelists**](DisplayableListsAPI.md#Getdisplayablelists) | **Post** /getdisplayablelists | Read displayable lists
[**Updatedisplayablelists**](DisplayableListsAPI.md#Updatedisplayablelists) | **Post** /updatedisplayablelists | Update displayable lists



## Createdisplayablelists

> CreateDisplayableListsResponse Createdisplayablelists(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Create displayable lists



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
	payload := *openapiclient.NewCreateDisplayableListsRequest([]openapiclient.DisplayableList{*openapiclient.NewDisplayableList()}) // CreateDisplayableListsRequest | Create new displayable lists
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisplayableListsAPI.Createdisplayablelists(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisplayableListsAPI.Createdisplayablelists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Createdisplayablelists`: CreateDisplayableListsResponse
	fmt.Fprintf(os.Stdout, "Response from `DisplayableListsAPI.Createdisplayablelists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatedisplayablelistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateDisplayableListsRequest**](CreateDisplayableListsRequest.md) | Create new displayable lists | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**CreateDisplayableListsResponse**](CreateDisplayableListsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deletedisplayablelists

> Deletedisplayablelists(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Delete displayable lists



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
	payload := *openapiclient.NewDeleteDisplayableListsRequest([]int32{int32(123)}) // DeleteDisplayableListsRequest | Updated version of the displayable list configs.
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DisplayableListsAPI.Deletedisplayablelists(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisplayableListsAPI.Deletedisplayablelists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletedisplayablelistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**DeleteDisplayableListsRequest**](DeleteDisplayableListsRequest.md) | Updated version of the displayable list configs. | 
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


## Getdisplayablelists

> GetDisplayableListsResponse Getdisplayablelists(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Read displayable lists



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
	payload := *openapiclient.NewGetDisplayableListsRequest([]int32{int32(123)}) // GetDisplayableListsRequest | 
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisplayableListsAPI.Getdisplayablelists(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisplayableListsAPI.Getdisplayablelists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Getdisplayablelists`: GetDisplayableListsResponse
	fmt.Fprintf(os.Stdout, "Response from `DisplayableListsAPI.Getdisplayablelists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetdisplayablelistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**GetDisplayableListsRequest**](GetDisplayableListsRequest.md) |  | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**GetDisplayableListsResponse**](GetDisplayableListsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Updatedisplayablelists

> UpdateDisplayableListsResponse Updatedisplayablelists(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Update displayable lists



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
	payload := *openapiclient.NewUpdateDisplayableListsRequest([]openapiclient.DisplayableList{*openapiclient.NewDisplayableList()}) // UpdateDisplayableListsRequest | Updated version of the displayable list configs.
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisplayableListsAPI.Updatedisplayablelists(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisplayableListsAPI.Updatedisplayablelists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Updatedisplayablelists`: UpdateDisplayableListsResponse
	fmt.Fprintf(os.Stdout, "Response from `DisplayableListsAPI.Updatedisplayablelists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatedisplayablelistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**UpdateDisplayableListsRequest**](UpdateDisplayableListsRequest.md) | Updated version of the displayable list configs. | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**UpdateDisplayableListsResponse**](UpdateDisplayableListsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

