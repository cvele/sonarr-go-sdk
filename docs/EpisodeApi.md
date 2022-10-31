# \EpisodeApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3EpisodeGet**](EpisodeApi.md#ApiV3EpisodeGet) | **Get** /api/v3/episode | 
[**ApiV3EpisodeIdGet**](EpisodeApi.md#ApiV3EpisodeIdGet) | **Get** /api/v3/episode/{id} | 
[**ApiV3EpisodeIdPut**](EpisodeApi.md#ApiV3EpisodeIdPut) | **Put** /api/v3/episode/{id} | 
[**ApiV3EpisodeMonitorPut**](EpisodeApi.md#ApiV3EpisodeMonitorPut) | **Put** /api/v3/episode/monitor | 



## ApiV3EpisodeGet

> []EpisodeResource ApiV3EpisodeGet(ctx).SeriesId(seriesId).SeasonNumber(seasonNumber).EpisodeIds(episodeIds).EpisodeFileId(episodeFileId).IncludeImages(includeImages).Execute()



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
    seasonNumber := int32(56) // int32 |  (optional)
    episodeIds := []int32{int32(123)} // []int32 |  (optional)
    episodeFileId := int32(56) // int32 |  (optional)
    includeImages := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodeApi.ApiV3EpisodeGet(context.Background()).SeriesId(seriesId).SeasonNumber(seasonNumber).EpisodeIds(episodeIds).EpisodeFileId(episodeFileId).IncludeImages(includeImages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeApi.ApiV3EpisodeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3EpisodeGet`: []EpisodeResource
    fmt.Fprintf(os.Stdout, "Response from `EpisodeApi.ApiV3EpisodeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3EpisodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesId** | **int32** |  | 
 **seasonNumber** | **int32** |  | 
 **episodeIds** | **[]int32** |  | 
 **episodeFileId** | **int32** |  | 
 **includeImages** | **bool** |  | [default to false]

### Return type

[**[]EpisodeResource**](EpisodeResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3EpisodeIdGet

> EpisodeResource ApiV3EpisodeIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.EpisodeApi.ApiV3EpisodeIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeApi.ApiV3EpisodeIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3EpisodeIdGet`: EpisodeResource
    fmt.Fprintf(os.Stdout, "Response from `EpisodeApi.ApiV3EpisodeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3EpisodeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EpisodeResource**](EpisodeResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3EpisodeIdPut

> ApiV3EpisodeIdPut(ctx, id).EpisodeResource(episodeResource).Execute()



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
    episodeResource := *openapiclient.NewEpisodeResource() // EpisodeResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodeApi.ApiV3EpisodeIdPut(context.Background(), id).EpisodeResource(episodeResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeApi.ApiV3EpisodeIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3EpisodeIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **episodeResource** | [**EpisodeResource**](EpisodeResource.md) |  | 

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


## ApiV3EpisodeMonitorPut

> ApiV3EpisodeMonitorPut(ctx).EpisodesMonitoredResource(episodesMonitoredResource).Execute()



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
    episodesMonitoredResource := *openapiclient.NewEpisodesMonitoredResource() // EpisodesMonitoredResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodeApi.ApiV3EpisodeMonitorPut(context.Background()).EpisodesMonitoredResource(episodesMonitoredResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeApi.ApiV3EpisodeMonitorPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3EpisodeMonitorPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **episodesMonitoredResource** | [**EpisodesMonitoredResource**](EpisodesMonitoredResource.md) |  | 

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

