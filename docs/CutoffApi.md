# \CutoffApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3WantedCutoffGet**](CutoffApi.md#ApiV3WantedCutoffGet) | **Get** /api/v3/wanted/cutoff | 
[**ApiV3WantedCutoffIdGet**](CutoffApi.md#ApiV3WantedCutoffIdGet) | **Get** /api/v3/wanted/cutoff/{id} | 



## ApiV3WantedCutoffGet

> EpisodeResourcePagingResource ApiV3WantedCutoffGet(ctx).IncludeSeries(includeSeries).IncludeEpisodeFile(includeEpisodeFile).IncludeImages(includeImages).Execute()



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
    includeEpisodeFile := true // bool |  (optional) (default to false)
    includeImages := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CutoffApi.ApiV3WantedCutoffGet(context.Background()).IncludeSeries(includeSeries).IncludeEpisodeFile(includeEpisodeFile).IncludeImages(includeImages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CutoffApi.ApiV3WantedCutoffGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3WantedCutoffGet`: EpisodeResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `CutoffApi.ApiV3WantedCutoffGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3WantedCutoffGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeSeries** | **bool** |  | [default to false]
 **includeEpisodeFile** | **bool** |  | [default to false]
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


## ApiV3WantedCutoffIdGet

> EpisodeResource ApiV3WantedCutoffIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.CutoffApi.ApiV3WantedCutoffIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CutoffApi.ApiV3WantedCutoffIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3WantedCutoffIdGet`: EpisodeResource
    fmt.Fprintf(os.Stdout, "Response from `CutoffApi.ApiV3WantedCutoffIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3WantedCutoffIdGetRequest struct via the builder pattern


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

