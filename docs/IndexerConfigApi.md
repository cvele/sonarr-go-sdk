# \IndexerConfigApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ConfigIndexerGet**](IndexerConfigApi.md#ApiV3ConfigIndexerGet) | **Get** /api/v3/config/indexer | 
[**ApiV3ConfigIndexerIdGet**](IndexerConfigApi.md#ApiV3ConfigIndexerIdGet) | **Get** /api/v3/config/indexer/{id} | 
[**ApiV3ConfigIndexerIdPut**](IndexerConfigApi.md#ApiV3ConfigIndexerIdPut) | **Put** /api/v3/config/indexer/{id} | 



## ApiV3ConfigIndexerGet

> IndexerConfigResource ApiV3ConfigIndexerGet(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerConfigApi.ApiV3ConfigIndexerGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerConfigApi.ApiV3ConfigIndexerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigIndexerGet`: IndexerConfigResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerConfigApi.ApiV3ConfigIndexerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigIndexerGetRequest struct via the builder pattern


### Return type

[**IndexerConfigResource**](IndexerConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ConfigIndexerIdGet

> IndexerConfigResource ApiV3ConfigIndexerIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.IndexerConfigApi.ApiV3ConfigIndexerIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerConfigApi.ApiV3ConfigIndexerIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigIndexerIdGet`: IndexerConfigResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerConfigApi.ApiV3ConfigIndexerIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigIndexerIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndexerConfigResource**](IndexerConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ConfigIndexerIdPut

> IndexerConfigResource ApiV3ConfigIndexerIdPut(ctx, id).IndexerConfigResource(indexerConfigResource).Execute()



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
    id := "id_example" // string | 
    indexerConfigResource := *openapiclient.NewIndexerConfigResource() // IndexerConfigResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndexerConfigApi.ApiV3ConfigIndexerIdPut(context.Background(), id).IndexerConfigResource(indexerConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndexerConfigApi.ApiV3ConfigIndexerIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigIndexerIdPut`: IndexerConfigResource
    fmt.Fprintf(os.Stdout, "Response from `IndexerConfigApi.ApiV3ConfigIndexerIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigIndexerIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **indexerConfigResource** | [**IndexerConfigResource**](IndexerConfigResource.md) |  | 

### Return type

[**IndexerConfigResource**](IndexerConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

