# \MediaCoverApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3MediacoverSeriesIdFilenameGet**](MediaCoverApi.md#ApiV3MediacoverSeriesIdFilenameGet) | **Get** /api/v3/mediacover/{seriesId}/{filename} | 



## ApiV3MediacoverSeriesIdFilenameGet

> ApiV3MediacoverSeriesIdFilenameGet(ctx, seriesId, filename).Execute()



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
    seriesId := int32(56) // int32 | 
    filename := "filename_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaCoverApi.ApiV3MediacoverSeriesIdFilenameGet(context.Background(), seriesId, filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaCoverApi.ApiV3MediacoverSeriesIdFilenameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seriesId** | **int32** |  | 
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3MediacoverSeriesIdFilenameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

