# \AnswersAPI

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Createanswer**](AnswersAPI.md#Createanswer) | **Post** /createanswer | Create Answer
[**Createanswerboard**](AnswersAPI.md#Createanswerboard) | **Post** /createanswerboard | Create Answer Board
[**Deleteanswer**](AnswersAPI.md#Deleteanswer) | **Post** /deleteanswer | Delete Answer
[**Deleteanswerboards**](AnswersAPI.md#Deleteanswerboards) | **Post** /deleteanswerboards | Delete Answer Board
[**Editanswer**](AnswersAPI.md#Editanswer) | **Post** /editanswer | Update Answer
[**Editanswerboard**](AnswersAPI.md#Editanswerboard) | **Post** /editanswerboard | Update Answer Board
[**Getanswer**](AnswersAPI.md#Getanswer) | **Post** /getanswer | Read Answer
[**Getanswerboard**](AnswersAPI.md#Getanswerboard) | **Post** /getanswerboard | Read Answer Board
[**Listanswerboards**](AnswersAPI.md#Listanswerboards) | **Post** /listanswerboards | List Answer Boards
[**Listanswers**](AnswersAPI.md#Listanswers) | **Post** /listanswers | List Answers
[**Previewanswer**](AnswersAPI.md#Previewanswer) | **Post** /previewanswer | Preview Answer
[**Previewanswerdraft**](AnswersAPI.md#Previewanswerdraft) | **Post** /previewanswerdraft | Preview draft Answer
[**Updateanswerlikes**](AnswersAPI.md#Updateanswerlikes) | **Post** /updateanswerlikes | Update Answer likes



## Createanswer

> Answer Createanswer(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Create Answer



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
	payload := *openapiclient.NewCreateAnswerRequest(*openapiclient.NewAnswerCreationData()) // CreateAnswerRequest | CreateAnswer request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnswersAPI.Createanswer(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnswersAPI.Createanswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Createanswer`: Answer
	fmt.Fprintf(os.Stdout, "Response from `AnswersAPI.Createanswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateanswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateAnswerRequest**](CreateAnswerRequest.md) | CreateAnswer request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**Answer**](Answer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Createanswerboard

> CreateAnswerBoardResponse Createanswerboard(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Create Answer Board



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
	payload := *openapiclient.NewCreateAnswerBoardRequest("Name_example") // CreateAnswerBoardRequest | Answer Board content plus any additional metadata for the request.
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnswersAPI.Createanswerboard(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnswersAPI.Createanswerboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Createanswerboard`: CreateAnswerBoardResponse
	fmt.Fprintf(os.Stdout, "Response from `AnswersAPI.Createanswerboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateanswerboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateAnswerBoardRequest**](CreateAnswerBoardRequest.md) | Answer Board content plus any additional metadata for the request. | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**CreateAnswerBoardResponse**](CreateAnswerBoardResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deleteanswer

> Deleteanswer(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Delete Answer



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
	payload := *openapiclient.NewDeleteAnswerRequest(int32(3)) // DeleteAnswerRequest | DeleteAnswer request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnswersAPI.Deleteanswer(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnswersAPI.Deleteanswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteanswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**DeleteAnswerRequest**](DeleteAnswerRequest.md) | DeleteAnswer request | 
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


## Deleteanswerboards

> DeleteAnswerBoardsResponse Deleteanswerboards(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Delete Answer Board



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
	payload := *openapiclient.NewDeleteAnswerBoardsRequest([]int32{int32(123)}) // DeleteAnswerBoardsRequest | DeleteAnswerBoards request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnswersAPI.Deleteanswerboards(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnswersAPI.Deleteanswerboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Deleteanswerboards`: DeleteAnswerBoardsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnswersAPI.Deleteanswerboards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteanswerboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**DeleteAnswerBoardsRequest**](DeleteAnswerBoardsRequest.md) | DeleteAnswerBoards request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**DeleteAnswerBoardsResponse**](DeleteAnswerBoardsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Editanswer

> Answer Editanswer(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Update Answer



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
	payload := *openapiclient.NewEditAnswerRequest(int32(3)) // EditAnswerRequest | EditAnswer request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnswersAPI.Editanswer(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnswersAPI.Editanswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Editanswer`: Answer
	fmt.Fprintf(os.Stdout, "Response from `AnswersAPI.Editanswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditanswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**EditAnswerRequest**](EditAnswerRequest.md) | EditAnswer request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**Answer**](Answer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Editanswerboard

> EditAnswerBoardResponse Editanswerboard(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Update Answer Board



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
	payload := *openapiclient.NewEditAnswerBoardRequest("Name_example", int32(123)) // EditAnswerBoardRequest | Answer Board content plus any additional metadata for the request.
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnswersAPI.Editanswerboard(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnswersAPI.Editanswerboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Editanswerboard`: EditAnswerBoardResponse
	fmt.Fprintf(os.Stdout, "Response from `AnswersAPI.Editanswerboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditanswerboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**EditAnswerBoardRequest**](EditAnswerBoardRequest.md) | Answer Board content plus any additional metadata for the request. | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**EditAnswerBoardResponse**](EditAnswerBoardResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getanswer

> GetAnswerResponse Getanswer(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Read Answer



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
	payload := *openapiclient.NewGetAnswerRequest() // GetAnswerRequest | GetAnswer request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnswersAPI.Getanswer(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnswersAPI.Getanswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Getanswer`: GetAnswerResponse
	fmt.Fprintf(os.Stdout, "Response from `AnswersAPI.Getanswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetanswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**GetAnswerRequest**](GetAnswerRequest.md) | GetAnswer request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**GetAnswerResponse**](GetAnswerResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getanswerboard

> GetAnswerBoardResponse Getanswerboard(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Read Answer Board



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
	payload := *openapiclient.NewGetAnswerBoardRequest(int32(123)) // GetAnswerBoardRequest | GetAnswerBoard request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnswersAPI.Getanswerboard(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnswersAPI.Getanswerboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Getanswerboard`: GetAnswerBoardResponse
	fmt.Fprintf(os.Stdout, "Response from `AnswersAPI.Getanswerboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetanswerboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**GetAnswerBoardRequest**](GetAnswerBoardRequest.md) | GetAnswerBoard request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**GetAnswerBoardResponse**](GetAnswerBoardResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Listanswerboards

> ListAnswerBoardsResponse Listanswerboards(ctx).Payload(payload).XScioActas(xScioActas).Execute()

List Answer Boards



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
	payload := *openapiclient.NewListAnswerBoardsRequest() // ListAnswerBoardsRequest | ListAnswerBoards request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnswersAPI.Listanswerboards(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnswersAPI.Listanswerboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Listanswerboards`: ListAnswerBoardsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnswersAPI.Listanswerboards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListanswerboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ListAnswerBoardsRequest**](ListAnswerBoardsRequest.md) | ListAnswerBoards request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**ListAnswerBoardsResponse**](ListAnswerBoardsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Listanswers

> ListAnswersResponse Listanswers(ctx).Payload(payload).XScioActas(xScioActas).Execute()

List Answers



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
	payload := *openapiclient.NewListAnswersRequest() // ListAnswersRequest | ListAnswers request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnswersAPI.Listanswers(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnswersAPI.Listanswers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Listanswers`: ListAnswersResponse
	fmt.Fprintf(os.Stdout, "Response from `AnswersAPI.Listanswers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListanswersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ListAnswersRequest**](ListAnswersRequest.md) | ListAnswers request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**ListAnswersResponse**](ListAnswersResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Previewanswer

> PreviewStructuredTextResponse Previewanswer(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Preview Answer



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
	payload := *openapiclient.NewPreviewStructuredTextRequest("From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light.") // PreviewStructuredTextRequest | PreviewAnswer request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnswersAPI.Previewanswer(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnswersAPI.Previewanswer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Previewanswer`: PreviewStructuredTextResponse
	fmt.Fprintf(os.Stdout, "Response from `AnswersAPI.Previewanswer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewanswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**PreviewStructuredTextRequest**](PreviewStructuredTextRequest.md) | PreviewAnswer request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**PreviewStructuredTextResponse**](PreviewStructuredTextResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Previewanswerdraft

> PreviewUgcResponse Previewanswerdraft(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Preview draft Answer



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
	payload := *openapiclient.NewPreviewUgcRequest() // PreviewUgcRequest | preview answer request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnswersAPI.Previewanswerdraft(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnswersAPI.Previewanswerdraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Previewanswerdraft`: PreviewUgcResponse
	fmt.Fprintf(os.Stdout, "Response from `AnswersAPI.Previewanswerdraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewanswerdraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**PreviewUgcRequest**](PreviewUgcRequest.md) | preview answer request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**PreviewUgcResponse**](PreviewUgcResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Updateanswerlikes

> UpdateAnswerLikesResponse Updateanswerlikes(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Update Answer likes



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
	payload := *openapiclient.NewUpdateAnswerLikesRequest(int32(3), "Action_example") // UpdateAnswerLikesRequest | UpdateAnswerLikes request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnswersAPI.Updateanswerlikes(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnswersAPI.Updateanswerlikes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Updateanswerlikes`: UpdateAnswerLikesResponse
	fmt.Fprintf(os.Stdout, "Response from `AnswersAPI.Updateanswerlikes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateanswerlikesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**UpdateAnswerLikesRequest**](UpdateAnswerLikesRequest.md) | UpdateAnswerLikes request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**UpdateAnswerLikesResponse**](UpdateAnswerLikesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

