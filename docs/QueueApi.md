# \QueueApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3QueueBulkDelete**](QueueApi.md#ApiV3QueueBulkDelete) | **Delete** /api/v3/queue/bulk | 
[**ApiV3QueueGet**](QueueApi.md#ApiV3QueueGet) | **Get** /api/v3/queue | 
[**ApiV3QueueIdDelete**](QueueApi.md#ApiV3QueueIdDelete) | **Delete** /api/v3/queue/{id} | 
[**ApiV3QueueIdGet**](QueueApi.md#ApiV3QueueIdGet) | **Get** /api/v3/queue/{id} | 



## ApiV3QueueBulkDelete

> ApiV3QueueBulkDelete(ctx).RemoveFromClient(removeFromClient).Blocklist(blocklist).QueueBulkResource(queueBulkResource).Execute()



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
    removeFromClient := true // bool |  (optional) (default to true)
    blocklist := true // bool |  (optional) (default to false)
    queueBulkResource := *openapiclient.NewQueueBulkResource() // QueueBulkResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueApi.ApiV3QueueBulkDelete(context.Background()).RemoveFromClient(removeFromClient).Blocklist(blocklist).QueueBulkResource(queueBulkResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueApi.ApiV3QueueBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QueueBulkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeFromClient** | **bool** |  | [default to true]
 **blocklist** | **bool** |  | [default to false]
 **queueBulkResource** | [**QueueBulkResource**](QueueBulkResource.md) |  | 

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


## ApiV3QueueGet

> QueueResourcePagingResource ApiV3QueueGet(ctx).IncludeUnknownSeriesItems(includeUnknownSeriesItems).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).Execute()



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
    includeUnknownSeriesItems := true // bool |  (optional) (default to false)
    includeSeries := true // bool |  (optional) (default to false)
    includeEpisode := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueApi.ApiV3QueueGet(context.Background()).IncludeUnknownSeriesItems(includeUnknownSeriesItems).IncludeSeries(includeSeries).IncludeEpisode(includeEpisode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueApi.ApiV3QueueGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3QueueGet`: QueueResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `QueueApi.ApiV3QueueGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QueueGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeUnknownSeriesItems** | **bool** |  | [default to false]
 **includeSeries** | **bool** |  | [default to false]
 **includeEpisode** | **bool** |  | [default to false]

### Return type

[**QueueResourcePagingResource**](QueueResourcePagingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3QueueIdDelete

> ApiV3QueueIdDelete(ctx, id).RemoveFromClient(removeFromClient).Blocklist(blocklist).Execute()



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
    removeFromClient := true // bool |  (optional) (default to true)
    blocklist := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueueApi.ApiV3QueueIdDelete(context.Background(), id).RemoveFromClient(removeFromClient).Blocklist(blocklist).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueApi.ApiV3QueueIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3QueueIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeFromClient** | **bool** |  | [default to true]
 **blocklist** | **bool** |  | [default to false]

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


## ApiV3QueueIdGet

> QueueResource ApiV3QueueIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.QueueApi.ApiV3QueueIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueueApi.ApiV3QueueIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3QueueIdGet`: QueueResource
    fmt.Fprintf(os.Stdout, "Response from `QueueApi.ApiV3QueueIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QueueIdGetRequest struct via the builder pattern


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

