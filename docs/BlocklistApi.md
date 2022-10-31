# \BlocklistApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3BlocklistBulkDelete**](BlocklistApi.md#ApiV3BlocklistBulkDelete) | **Delete** /api/v3/blocklist/bulk | 
[**ApiV3BlocklistGet**](BlocklistApi.md#ApiV3BlocklistGet) | **Get** /api/v3/blocklist | 
[**ApiV3BlocklistIdDelete**](BlocklistApi.md#ApiV3BlocklistIdDelete) | **Delete** /api/v3/blocklist/{id} | 



## ApiV3BlocklistBulkDelete

> ApiV3BlocklistBulkDelete(ctx).BlocklistBulkResource(blocklistBulkResource).Execute()



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
    blocklistBulkResource := *openapiclient.NewBlocklistBulkResource() // BlocklistBulkResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlocklistApi.ApiV3BlocklistBulkDelete(context.Background()).BlocklistBulkResource(blocklistBulkResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocklistApi.ApiV3BlocklistBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3BlocklistBulkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blocklistBulkResource** | [**BlocklistBulkResource**](BlocklistBulkResource.md) |  | 

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


## ApiV3BlocklistGet

> BlocklistResourcePagingResource ApiV3BlocklistGet(ctx).Execute()



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
    resp, r, err := apiClient.BlocklistApi.ApiV3BlocklistGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocklistApi.ApiV3BlocklistGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3BlocklistGet`: BlocklistResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `BlocklistApi.ApiV3BlocklistGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3BlocklistGetRequest struct via the builder pattern


### Return type

[**BlocklistResourcePagingResource**](BlocklistResourcePagingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3BlocklistIdDelete

> ApiV3BlocklistIdDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.BlocklistApi.ApiV3BlocklistIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocklistApi.ApiV3BlocklistIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3BlocklistIdDeleteRequest struct via the builder pattern


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

