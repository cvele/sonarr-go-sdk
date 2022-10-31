# \QueueDetailsApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3QueueDetailsGet**](QueueDetailsApi.md#ApiV3QueueDetailsGet) | **Get** /api/v3/queue/details | 
[**ApiV3QueueDetailsIdGet**](QueueDetailsApi.md#ApiV3QueueDetailsIdGet) | **Get** /api/v3/queue/details/{id} | 



## ApiV3QueueDetailsGet

> []QueueResource ApiV3QueueDetailsGet(ctx).SeriesId(seriesId).EpisodeIds(episodeIds).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).Execute()



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
    seriesId := int32(56) // int32 |  (optional)
    episodeIds := []int32{int32(123)} // []int32 |  (optional)
    includeSeries := true // bool |  (optional) (default to false)
    includeEpisode := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueDetailsApi.ApiV3QueueDetailsGet(context.Background()).SeriesId(seriesId).EpisodeIds(episodeIds).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueDetailsApi.ApiV3QueueDetailsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3QueueDetailsGet`: []QueueResource
    fmt.Fprintf(os.Stdout, "Response from `QueueDetailsApi.ApiV3QueueDetailsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QueueDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesId** | **int32** |  | 
 **episodeIds** | **[]int32** |  | 
 **includeSeries** | **bool** |  | [default to false]
 **includeEpisode** | **bool** |  | [default to false]

### Return type

[**[]QueueResource**](QueueResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3QueueDetailsIdGet

> QueueResource ApiV3QueueDetailsIdGet(ctx, id).Execute()



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
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueDetailsApi.ApiV3QueueDetailsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueDetailsApi.ApiV3QueueDetailsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3QueueDetailsIdGet`: QueueResource
    fmt.Fprintf(os.Stdout, "Response from `QueueDetailsApi.ApiV3QueueDetailsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QueueDetailsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueueResource**](QueueResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

