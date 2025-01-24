# \SummarizeAPI

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Summarize**](SummarizeAPI.md#Summarize) | **Post** /summarize | Summarize documents



## Summarize

> SummarizeResponse Summarize(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Summarize documents



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
	payload := *openapiclient.NewSummarizeRequest([]openapiclient.DocumentSpec{openapiclient.DocumentSpec{DocumentSpecOneOf: openapiclient.NewDocumentSpecOneOf()}}) // SummarizeRequest | Includes request params such as the query and specs of the documents to summarize.
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SummarizeAPI.Summarize(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SummarizeAPI.Summarize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Summarize`: SummarizeResponse
	fmt.Fprintf(os.Stdout, "Response from `SummarizeAPI.Summarize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummarizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**SummarizeRequest**](SummarizeRequest.md) | Includes request params such as the query and specs of the documents to summarize. | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**SummarizeResponse**](SummarizeResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

