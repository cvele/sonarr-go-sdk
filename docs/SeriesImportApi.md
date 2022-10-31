# \SeriesImportApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3SeriesImportPost**](SeriesImportApi.md#ApiV3SeriesImportPost) | **Post** /api/v3/series/import | 



## ApiV3SeriesImportPost

> ApiV3SeriesImportPost(ctx).SeriesResource(seriesResource).Execute()



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
    seriesResource := []openapiclient.SeriesResource{*openapiclient.NewSeriesResource()} // []SeriesResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesImportApi.ApiV3SeriesImportPost(context.Background()).SeriesResource(seriesResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesImportApi.ApiV3SeriesImportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SeriesImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesResource** | [**[]SeriesResource**](SeriesResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

