# \SeriesEditorApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3SeriesEditorDelete**](SeriesEditorApi.md#ApiV3SeriesEditorDelete) | **Delete** /api/v3/series/editor | 
[**ApiV3SeriesEditorPut**](SeriesEditorApi.md#ApiV3SeriesEditorPut) | **Put** /api/v3/series/editor | 



## ApiV3SeriesEditorDelete

> ApiV3SeriesEditorDelete(ctx).SeriesEditorResource(seriesEditorResource).Execute()



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
    seriesEditorResource := *openapiclient.NewSeriesEditorResource() // SeriesEditorResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesEditorApi.ApiV3SeriesEditorDelete(context.Background()).SeriesEditorResource(seriesEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesEditorApi.ApiV3SeriesEditorDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SeriesEditorDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesEditorResource** | [**SeriesEditorResource**](SeriesEditorResource.md) |  | 

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


## ApiV3SeriesEditorPut

> ApiV3SeriesEditorPut(ctx).SeriesEditorResource(seriesEditorResource).Execute()



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
    seriesEditorResource := *openapiclient.NewSeriesEditorResource() // SeriesEditorResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesEditorApi.ApiV3SeriesEditorPut(context.Background()).SeriesEditorResource(seriesEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesEditorApi.ApiV3SeriesEditorPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SeriesEditorPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesEditorResource** | [**SeriesEditorResource**](SeriesEditorResource.md) |  | 

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

