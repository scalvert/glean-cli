# \ToolsAPI

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Executeactiontool**](ToolsAPI.md#Executeactiontool) | **Post** /executeactiontool | Execute Action Tool



## Executeactiontool

> ExecuteActionToolResponse Executeactiontool(ctx).Payload(payload).TimezoneOffset(timezoneOffset).Execute()

Execute Action Tool



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
	payload := *openapiclient.NewExecuteActionToolRequest("Name_example") // ExecuteActionToolRequest | Execute Action Tool request
	timezoneOffset := int32(56) // int32 | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsAPI.Executeactiontool(context.Background()).Payload(payload).TimezoneOffset(timezoneOffset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsAPI.Executeactiontool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Executeactiontool`: ExecuteActionToolResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsAPI.Executeactiontool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteactiontoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ExecuteActionToolRequest**](ExecuteActionToolRequest.md) | Execute Action Tool request | 
 **timezoneOffset** | **int32** | The offset of the client&#39;s timezone in minutes from UTC. e.g. PDT is -420 because it&#39;s 7 hours behind UTC. | 

### Return type

[**ExecuteActionToolResponse**](ExecuteActionToolResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

