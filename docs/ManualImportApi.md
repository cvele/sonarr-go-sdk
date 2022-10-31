# \ManualImportApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ManualimportGet**](ManualImportApi.md#ApiV3ManualimportGet) | **Get** /api/v3/manualimport | 
[**ApiV3ManualimportPost**](ManualImportApi.md#ApiV3ManualimportPost) | **Post** /api/v3/manualimport | 



## ApiV3ManualimportGet

> []ManualImportResource ApiV3ManualimportGet(ctx).Folder(folder).DownloadId(downloadId).SeriesId(seriesId).SeasonNumber(seasonNumber).FilterExistingFiles(filterExistingFiles).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    folder := "folder_example" // string |  (optional)
    downloadId := "downloadId_example" // string |  (optional)
    seriesId := int32(56) // int32 |  (optional)
    seasonNumber := int32(56) // int32 |  (optional)
    filterExistingFiles := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualImportApi.ApiV3ManualimportGet(context.Background()).Folder(folder).DownloadId(downloadId).SeriesId(seriesId).SeasonNumber(seasonNumber).FilterExistingFiles(filterExistingFiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualImportApi.ApiV3ManualimportGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ManualimportGet`: []ManualImportResource
    fmt.Fprintf(os.Stdout, "Response from `ManualImportApi.ApiV3ManualimportGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ManualimportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folder** | **string** |  | 
 **downloadId** | **string** |  | 
 **seriesId** | **int32** |  | 
 **seasonNumber** | **int32** |  | 
 **filterExistingFiles** | **bool** |  | [default to true]

### Return type

[**[]ManualImportResource**](ManualImportResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ManualimportPost

> ApiV3ManualimportPost(ctx).ManualImportReprocessResource(manualImportReprocessResource).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    manualImportReprocessResource := []openapiclient.ManualImportReprocessResource{*openapiclient.NewManualImportReprocessResource()} // []ManualImportReprocessResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualImportApi.ApiV3ManualimportPost(context.Background()).ManualImportReprocessResource(manualImportReprocessResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualImportApi.ApiV3ManualimportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ManualimportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manualImportReprocessResource** | [**[]ManualImportReprocessResource**](ManualImportReprocessResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

