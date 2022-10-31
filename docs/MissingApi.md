# \MissingApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3WantedMissingGet**](MissingApi.md#ApiV3WantedMissingGet) | **Get** /api/v3/wanted/missing | 
[**ApiV3WantedMissingIdGet**](MissingApi.md#ApiV3WantedMissingIdGet) | **Get** /api/v3/wanted/missing/{id} | 



## ApiV3WantedMissingGet

> EpisodeResourcePagingResource ApiV3WantedMissingGet(ctx).IncludeSeries(includeSeries).IncludeImages(includeImages).Execute()



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
    includeSeries := true // bool |  (optional) (default to false)
    includeImages := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MissingApi.ApiV3WantedMissingGet(context.Background()).IncludeSeries(includeSeries).IncludeImages(includeImages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MissingApi.ApiV3WantedMissingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3WantedMissingGet`: EpisodeResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `MissingApi.ApiV3WantedMissingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3WantedMissingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeSeries** | **bool** |  | [default to false]
 **includeImages** | **bool** |  | [default to false]

### Return type

[**EpisodeResourcePagingResource**](EpisodeResourcePagingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3WantedMissingIdGet

> EpisodeResource ApiV3WantedMissingIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.MissingApi.ApiV3WantedMissingIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MissingApi.ApiV3WantedMissingIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3WantedMissingIdGet`: EpisodeResource
    fmt.Fprintf(os.Stdout, "Response from `MissingApi.ApiV3WantedMissingIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3WantedMissingIdGetRequest struct via the builder pattern


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

