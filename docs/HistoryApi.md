# \HistoryApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3HistoryFailedIdPost**](HistoryApi.md#ApiV3HistoryFailedIdPost) | **Post** /api/v3/history/failed/{id} | 
[**ApiV3HistoryGet**](HistoryApi.md#ApiV3HistoryGet) | **Get** /api/v3/history | 
[**ApiV3HistorySeriesGet**](HistoryApi.md#ApiV3HistorySeriesGet) | **Get** /api/v3/history/series | 
[**ApiV3HistorySinceGet**](HistoryApi.md#ApiV3HistorySinceGet) | **Get** /api/v3/history/since | 



## ApiV3HistoryFailedIdPost

> ApiV3HistoryFailedIdPost(ctx, id).Execute()



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
    resp, r, err := apiClient.HistoryApi.ApiV3HistoryFailedIdPost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ApiV3HistoryFailedIdPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3HistoryFailedIdPostRequest struct via the builder pattern


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


## ApiV3HistoryGet

> HistoryResourcePagingResource ApiV3HistoryGet(ctx).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).Execute()



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
    includeSeries := true // bool |  (optional)
    includeEpisode := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.ApiV3HistoryGet(context.Background()).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ApiV3HistoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3HistoryGet`: HistoryResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ApiV3HistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3HistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeSeries** | **bool** |  | 
 **includeEpisode** | **bool** |  | 

### Return type

[**HistoryResourcePagingResource**](HistoryResourcePagingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3HistorySeriesGet

> []HistoryResource ApiV3HistorySeriesGet(ctx).SeriesId(seriesId).SeasonNumber(seasonNumber).EventType(eventType).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).Execute()



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
    eventType := openapiclient.EpisodeHistoryEventType("unknown") // EpisodeHistoryEventType |  (optional)
    includeSeries := true // bool |  (optional) (default to false)
    includeEpisode := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.ApiV3HistorySeriesGet(context.Background()).SeriesId(seriesId).SeasonNumber(seasonNumber).EventType(eventType).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ApiV3HistorySeriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3HistorySeriesGet`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ApiV3HistorySeriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3HistorySeriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesId** | **int32** |  | 
 **seasonNumber** | **int32** |  | 
 **eventType** | [**EpisodeHistoryEventType**](EpisodeHistoryEventType.md) |  | 
 **includeSeries** | **bool** |  | [default to false]
 **includeEpisode** | **bool** |  | [default to false]

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3HistorySinceGet

> []HistoryResource ApiV3HistorySinceGet(ctx).Date(date).EventType(eventType).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    date := time.Now() // time.Time |  (optional)
    eventType := openapiclient.EpisodeHistoryEventType("unknown") // EpisodeHistoryEventType |  (optional)
    includeSeries := true // bool |  (optional) (default to false)
    includeEpisode := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.ApiV3HistorySinceGet(context.Background()).Date(date).EventType(eventType).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ApiV3HistorySinceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3HistorySinceGet`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ApiV3HistorySinceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3HistorySinceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** |  | 
 **eventType** | [**EpisodeHistoryEventType**](EpisodeHistoryEventType.md) |  | 
 **includeSeries** | **bool** |  | [default to false]
 **includeEpisode** | **bool** |  | [default to false]

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

