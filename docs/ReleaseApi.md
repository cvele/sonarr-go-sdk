# \ReleaseApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ReleaseGet**](ReleaseApi.md#ApiV3ReleaseGet) | **Get** /api/v3/release | 
[**ApiV3ReleaseIdGet**](ReleaseApi.md#ApiV3ReleaseIdGet) | **Get** /api/v3/release/{id} | 
[**ApiV3ReleasePost**](ReleaseApi.md#ApiV3ReleasePost) | **Post** /api/v3/release | 



## ApiV3ReleaseGet

> []ReleaseResource ApiV3ReleaseGet(ctx).SeriesId(seriesId).EpisodeId(episodeId).SeasonNumber(seasonNumber).Execute()



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
    episodeId := int32(56) // int32 |  (optional)
    seasonNumber := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.ApiV3ReleaseGet(context.Background()).SeriesId(seriesId).EpisodeId(episodeId).SeasonNumber(seasonNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.ApiV3ReleaseGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ReleaseGet`: []ReleaseResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.ApiV3ReleaseGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ReleaseGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesId** | **int32** |  | 
 **episodeId** | **int32** |  | 
 **seasonNumber** | **int32** |  | 

### Return type

[**[]ReleaseResource**](ReleaseResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ReleaseIdGet

> ReleaseResource ApiV3ReleaseIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.ReleaseApi.ApiV3ReleaseIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.ApiV3ReleaseIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ReleaseIdGet`: ReleaseResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseApi.ApiV3ReleaseIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ReleaseIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReleaseResource**](ReleaseResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ReleasePost

> ApiV3ReleasePost(ctx).ReleaseResource(releaseResource).Execute()



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
    releaseResource := *openapiclient.NewReleaseResource() // ReleaseResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseApi.ApiV3ReleasePost(context.Background()).ReleaseResource(releaseResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseApi.ApiV3ReleasePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ReleasePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseResource** | [**ReleaseResource**](ReleaseResource.md) |  | 

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

