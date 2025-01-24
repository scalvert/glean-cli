# \ImagesAPI

All URIs are relative to *https://domain-be.glean.com/rest/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Images**](ImagesAPI.md#Images) | **Get** /images | Get image
[**Uploadimage**](ImagesAPI.md#Uploadimage) | **Post** /uploadimage | Upload images



## Images

> *os.File Images(ctx).XScioActas(xScioActas).Key(key).Type_(type_).Id(id).Ds(ds).Cid(cid).Execute()

Get image



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
	key := "key_example" // string | Primary key for the image being asked. The key is returned by the API when an image is uploaded. If key is used, other parameters should not be used. (optional)
	type_ := openapiclient.ImageType("BACKGROUND") // ImageType | The type of image requested. Supported values are listed in ImageMetadata.type enum. (optional)
	id := "id_example" // string | ID, if a specific entity/type is requested. The id may have different meaning for each type. for user, it is user id, for UGC, it is the id of the content, and so on. (optional)
	ds := "ds_example" // string | A specific datasource for which an image is requested for. The ds may have different meaning for each type and can also be empty for some. (optional)
	cid := "cid_example" // string | Content id to differentitate multiple images that can have the same type and datasource e.g. thumnail or image from content of UGC. It can also be empty. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.Images(context.Background()).XScioActas(xScioActas).Key(key).Type_(type_).Id(id).Ds(ds).Cid(cid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.Images``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Images`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.Images`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 
 **key** | **string** | Primary key for the image being asked. The key is returned by the API when an image is uploaded. If key is used, other parameters should not be used. | 
 **type_** | [**ImageType**](ImageType.md) | The type of image requested. Supported values are listed in ImageMetadata.type enum. | 
 **id** | **string** | ID, if a specific entity/type is requested. The id may have different meaning for each type. for user, it is user id, for UGC, it is the id of the content, and so on. | 
 **ds** | **string** | A specific datasource for which an image is requested for. The ds may have different meaning for each type and can also be empty for some. | 
 **cid** | **string** | Content id to differentitate multiple images that can have the same type and datasource e.g. thumnail or image from content of UGC. It can also be empty. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Uploadimage

> UploadImageResponse Uploadimage(ctx).Payload(payload).XScioActas(xScioActas).Type_(type_).Id(id).Ds(ds).Cid(cid).Execute()

Upload images



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
	payload := os.NewFile(1234, "some_file") // *os.File | Content and metadata for the image. Content is in the POST body, metadata is in the URL.
	xScioActas := "xScioActas_example" // string | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). (optional)
	type_ := openapiclient.ImageType("BACKGROUND") // ImageType | The type of image requested. Supported values are listed in ImageMetadata.type enum. (optional)
	id := "id_example" // string | ID, if a specific entity/type is requested. The id may have different meaning for each type. For USER, it is user id For UGC, it is the id of the content For ICON, the doctype. (optional)
	ds := "ds_example" // string | A specific datasource for which an image is requested for. The ds may have different meaning for each type and can also be empty for some. For USER, it is empty or datasource the icon is asked for. For UGC, it is the UGC datasource. For ICON, it is datasource instance the icon is asked for. (optional)
	cid := "cid_example" // string | Content id to differentitate multiple images that can have the same type and datasource e.g. thumnail or image from content of UGC. It can also be empty. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.Uploadimage(context.Background()).Payload(payload).XScioActas(xScioActas).Type_(type_).Id(id).Ds(ds).Cid(cid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.Uploadimage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Uploadimage`: UploadImageResponse
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.Uploadimage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadimageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | ***os.File** | Content and metadata for the image. Content is in the POST body, metadata is in the URL. | 
 **xScioActas** | **string** | Email address of a user on whose behalf the request is intended to be made (should be non-empty only for global tokens). | 
 **type_** | [**ImageType**](ImageType.md) | The type of image requested. Supported values are listed in ImageMetadata.type enum. | 
 **id** | **string** | ID, if a specific entity/type is requested. The id may have different meaning for each type. For USER, it is user id For UGC, it is the id of the content For ICON, the doctype. | 
 **ds** | **string** | A specific datasource for which an image is requested for. The ds may have different meaning for each type and can also be empty for some. For USER, it is empty or datasource the icon is asked for. For UGC, it is the UGC datasource. For ICON, it is datasource instance the icon is asked for. | 
 **cid** | **string** | Content id to differentitate multiple images that can have the same type and datasource e.g. thumnail or image from content of UGC. It can also be empty. | 

### Return type

[**UploadImageResponse**](UploadImageResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: image/*
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

