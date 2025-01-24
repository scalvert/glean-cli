# \DocumentsAPI

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getdocpermissions**](DocumentsAPI.md#Getdocpermissions) | **Post** /getdocpermissions | Read document permissions
[**Getdocumentanalytics**](DocumentsAPI.md#Getdocumentanalytics) | **Post** /getdocumentanalytics | Read document analytics
[**Getdocuments**](DocumentsAPI.md#Getdocuments) | **Post** /getdocuments | Read documents



## Getdocpermissions

> GetDocPermissionsResponse Getdocpermissions(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Read document permissions



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
	payload := *openapiclient.NewGetDocPermissionsRequest() // GetDocPermissionsRequest | Document permissions request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.Getdocpermissions(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.Getdocpermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Getdocpermissions`: GetDocPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.Getdocpermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetdocpermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**GetDocPermissionsRequest**](GetDocPermissionsRequest.md) | Document permissions request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**GetDocPermissionsResponse**](GetDocPermissionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getdocumentanalytics

> GetDocumentAnalyticsResponse Getdocumentanalytics(ctx).XScioActas(xScioActas).Payload(payload).Execute()

Read document analytics



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
	payload := *openapiclient.NewGetDocumentAnalyticsRequest([]openapiclient.DocumentSpec{openapiclient.DocumentSpec{DocumentSpecOneOf: openapiclient.NewDocumentSpecOneOf()}}, *openapiclient.NewPeriod()) // GetDocumentAnalyticsRequest | Information about analytics requested. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.Getdocumentanalytics(context.Background()).XScioActas(xScioActas).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.Getdocumentanalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Getdocumentanalytics`: GetDocumentAnalyticsResponse
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.Getdocumentanalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetdocumentanalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 
 **payload** | [**GetDocumentAnalyticsRequest**](GetDocumentAnalyticsRequest.md) | Information about analytics requested. | 

### Return type

[**GetDocumentAnalyticsResponse**](GetDocumentAnalyticsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getdocuments

> GetDocumentsResponse Getdocuments(ctx).XScioActas(xScioActas).Payload(payload).Execute()

Read documents



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
	payload := *openapiclient.NewGetDocumentsRequest([]openapiclient.DocumentSpec{openapiclient.DocumentSpec{DocumentSpecOneOf: openapiclient.NewDocumentSpecOneOf()}}) // GetDocumentsRequest | Information about documents requested. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.Getdocuments(context.Background()).XScioActas(xScioActas).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.Getdocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Getdocuments`: GetDocumentsResponse
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.Getdocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetdocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 
 **payload** | [**GetDocumentsRequest**](GetDocumentsRequest.md) | Information about documents requested. | 

### Return type

[**GetDocumentsResponse**](GetDocumentsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

