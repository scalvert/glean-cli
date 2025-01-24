# \AuthenticationAPI

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Createanonymoustoken**](AuthenticationAPI.md#Createanonymoustoken) | **Post** /createanonymoustoken | Create anonymous token
[**Createauthtoken**](AuthenticationAPI.md#Createauthtoken) | **Post** /createauthtoken | Create authentication token



## Createanonymoustoken

> CreateAuthTokenResponse Createanonymoustoken(ctx).Execute()

Create anonymous token



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Createanonymoustoken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Createanonymoustoken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Createanonymoustoken`: CreateAuthTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Createanonymoustoken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateanonymoustokenRequest struct via the builder pattern


### Return type

[**CreateAuthTokenResponse**](CreateAuthTokenResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Createauthtoken

> CreateAuthTokenResponse Createauthtoken(ctx).XScioActas(xScioActas).Execute()

Create authentication token



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Createauthtoken(context.Background()).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Createauthtoken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Createauthtoken`: CreateAuthTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Createauthtoken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateauthtokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**CreateAuthTokenResponse**](CreateAuthTokenResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

