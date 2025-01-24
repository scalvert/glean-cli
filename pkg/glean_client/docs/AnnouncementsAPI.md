# \AnnouncementsAPI

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Createannouncement**](AnnouncementsAPI.md#Createannouncement) | **Post** /createannouncement | Create Announcement
[**Createdraftannouncement**](AnnouncementsAPI.md#Createdraftannouncement) | **Post** /createdraftannouncement | Create draft Announcement
[**Deleteannouncement**](AnnouncementsAPI.md#Deleteannouncement) | **Post** /deleteannouncement | Delete Announcement
[**Deletedraftannouncement**](AnnouncementsAPI.md#Deletedraftannouncement) | **Post** /deletedraftannouncement | Delete draft Announcement
[**Getannouncement**](AnnouncementsAPI.md#Getannouncement) | **Post** /getannouncement | Read Announcement
[**Getdraftannouncement**](AnnouncementsAPI.md#Getdraftannouncement) | **Post** /getdraftannouncement | Read draft Announcement
[**Listannouncements**](AnnouncementsAPI.md#Listannouncements) | **Post** /listannouncements | List Announcements
[**Previewannouncement**](AnnouncementsAPI.md#Previewannouncement) | **Post** /previewannouncement | Preview Announcement
[**Previewannouncementdraft**](AnnouncementsAPI.md#Previewannouncementdraft) | **Post** /previewannouncementdraft | Preview draft Announcement
[**Publishdraftannouncement**](AnnouncementsAPI.md#Publishdraftannouncement) | **Post** /publishdraftannouncement | Publish draft Announcement
[**Unpublishannouncement**](AnnouncementsAPI.md#Unpublishannouncement) | **Post** /unpublishannouncement | Unpublish Announcement
[**Updateannouncement**](AnnouncementsAPI.md#Updateannouncement) | **Post** /updateannouncement | Update Announcement
[**Updatedraftannouncement**](AnnouncementsAPI.md#Updatedraftannouncement) | **Post** /updatedraftannouncement | Update draft Announcement



## Createannouncement

> Announcement Createannouncement(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Create Announcement



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	payload := *openapiclient.NewCreateAnnouncementRequest(time.Now(), time.Now(), "Title_example") // CreateAnnouncementRequest | Announcement content
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.Createannouncement(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.Createannouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Createannouncement`: Announcement
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.Createannouncement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateannouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateAnnouncementRequest**](CreateAnnouncementRequest.md) | Announcement content | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**Announcement**](Announcement.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Createdraftannouncement

> Announcement Createdraftannouncement(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Create draft Announcement



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
	payload := *openapiclient.NewCreateDraftAnnouncementRequest() // CreateDraftAnnouncementRequest | Draft announcement content
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.Createdraftannouncement(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.Createdraftannouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Createdraftannouncement`: Announcement
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.Createdraftannouncement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatedraftannouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateDraftAnnouncementRequest**](CreateDraftAnnouncementRequest.md) | Draft announcement content | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**Announcement**](Announcement.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deleteannouncement

> Deleteannouncement(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Delete Announcement



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
	payload := *openapiclient.NewDeleteAnnouncementRequest(int32(123)) // DeleteAnnouncementRequest | Delete announcement request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnnouncementsAPI.Deleteannouncement(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.Deleteannouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteannouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**DeleteAnnouncementRequest**](DeleteAnnouncementRequest.md) | Delete announcement request | 
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


## Deletedraftannouncement

> Deletedraftannouncement(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Delete draft Announcement



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
	payload := *openapiclient.NewDeleteAnnouncementRequest(int32(123)) // DeleteAnnouncementRequest | Delete draft announcement request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnnouncementsAPI.Deletedraftannouncement(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.Deletedraftannouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletedraftannouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**DeleteAnnouncementRequest**](DeleteAnnouncementRequest.md) | Delete draft announcement request | 
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


## Getannouncement

> GetAnnouncementResponse Getannouncement(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Read Announcement



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
	payload := *openapiclient.NewGetAnnouncementRequest(int32(123)) // GetAnnouncementRequest | GetAnnouncement request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.Getannouncement(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.Getannouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Getannouncement`: GetAnnouncementResponse
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.Getannouncement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetannouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**GetAnnouncementRequest**](GetAnnouncementRequest.md) | GetAnnouncement request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**GetAnnouncementResponse**](GetAnnouncementResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Getdraftannouncement

> GetDraftAnnouncementResponse Getdraftannouncement(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Read draft Announcement



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
	payload := *openapiclient.NewGetAnnouncementRequest(int32(123)) // GetAnnouncementRequest | Get draft announcement request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.Getdraftannouncement(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.Getdraftannouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Getdraftannouncement`: GetDraftAnnouncementResponse
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.Getdraftannouncement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetdraftannouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**GetAnnouncementRequest**](GetAnnouncementRequest.md) | Get draft announcement request | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**GetDraftAnnouncementResponse**](GetDraftAnnouncementResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Listannouncements

> ListAnnouncementsResponse Listannouncements(ctx).Payload(payload).XScioActas(xScioActas).Execute()

List Announcements



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
	payload := *openapiclient.NewListAnnouncementsRequest() // ListAnnouncementsRequest | Includes request params for querying announcements.
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.Listannouncements(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.Listannouncements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Listannouncements`: ListAnnouncementsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.Listannouncements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListannouncementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**ListAnnouncementsRequest**](ListAnnouncementsRequest.md) | Includes request params for querying announcements. | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**ListAnnouncementsResponse**](ListAnnouncementsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Previewannouncement

> PreviewStructuredTextResponse Previewannouncement(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Preview Announcement



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
	payload := *openapiclient.NewPreviewStructuredTextRequest("From https://en.wikipedia.org/wiki/Diffuse_sky_radiation, the sky is blue because blue light is more strongly scattered than longer-wavelength light.") // PreviewStructuredTextRequest | preview structured text request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.Previewannouncement(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.Previewannouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Previewannouncement`: PreviewStructuredTextResponse
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.Previewannouncement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewannouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**PreviewStructuredTextRequest**](PreviewStructuredTextRequest.md) | preview structured text request | 
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


## Previewannouncementdraft

> PreviewUgcResponse Previewannouncementdraft(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Preview draft Announcement



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
	payload := *openapiclient.NewPreviewUgcRequest() // PreviewUgcRequest | preview announcement request
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.Previewannouncementdraft(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.Previewannouncementdraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Previewannouncementdraft`: PreviewUgcResponse
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.Previewannouncementdraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewannouncementdraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**PreviewUgcRequest**](PreviewUgcRequest.md) | preview announcement request | 
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


## Publishdraftannouncement

> Publishdraftannouncement(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Publish draft Announcement



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
	payload := *openapiclient.NewPublishDraftAnnouncementRequest(int32(123)) // PublishDraftAnnouncementRequest | Publish draft announcement content.
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnnouncementsAPI.Publishdraftannouncement(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.Publishdraftannouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishdraftannouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**PublishDraftAnnouncementRequest**](PublishDraftAnnouncementRequest.md) | Publish draft announcement content. | 
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


## Unpublishannouncement

> Announcement Unpublishannouncement(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Unpublish Announcement



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
	payload := *openapiclient.NewUnpublishAnnouncementRequest(int32(123)) // UnpublishAnnouncementRequest | Unpublish announcement content.
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.Unpublishannouncement(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.Unpublishannouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Unpublishannouncement`: Announcement
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.Unpublishannouncement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishannouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**UnpublishAnnouncementRequest**](UnpublishAnnouncementRequest.md) | Unpublish announcement content. | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**Announcement**](Announcement.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Updateannouncement

> Announcement Updateannouncement(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Update Announcement



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	payload := *openapiclient.NewUpdateAnnouncementRequest(time.Now(), time.Now(), "Title_example", int32(123)) // UpdateAnnouncementRequest | Announcement content. Id need to be specified for the announcement.
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.Updateannouncement(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.Updateannouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Updateannouncement`: Announcement
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.Updateannouncement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateannouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**UpdateAnnouncementRequest**](UpdateAnnouncementRequest.md) | Announcement content. Id need to be specified for the announcement. | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**Announcement**](Announcement.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Updatedraftannouncement

> Announcement Updatedraftannouncement(ctx).Payload(payload).XScioActas(xScioActas).Execute()

Update draft Announcement



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
	payload := *openapiclient.NewUpdateDraftAnnouncementRequest(int32(123)) // UpdateDraftAnnouncementRequest | Draft announcement content. DraftId needs to be specified.
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.Updatedraftannouncement(context.Background()).Payload(payload).XScioActas(xScioActas).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.Updatedraftannouncement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Updatedraftannouncement`: Announcement
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.Updatedraftannouncement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatedraftannouncementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**UpdateDraftAnnouncementRequest**](UpdateDraftAnnouncementRequest.md) | Draft announcement content. DraftId needs to be specified. | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 

### Return type

[**Announcement**](Announcement.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

