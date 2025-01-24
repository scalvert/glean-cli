# \InsightsAPI

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Insights**](InsightsAPI.md#Insights) | **Post** /insights | Read insights



## Insights

> InsightsResponse Insights(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Read insights



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
	payload := *openapiclient.NewInsightsRequest([]string{"Categories_example"}) // InsightsRequest | Includes request params for insights dashboard data.
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightsAPI.Insights(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsAPI.Insights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Insights`: InsightsResponse
	fmt.Fprintf(os.Stdout, "Response from `InsightsAPI.Insights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**InsightsRequest**](InsightsRequest.md) | Includes request params for insights dashboard data. | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**InsightsResponse**](InsightsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

