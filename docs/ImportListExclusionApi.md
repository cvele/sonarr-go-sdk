# \ImportListExclusionApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ImportlistexclusionGet**](ImportListExclusionApi.md#ApiV3ImportlistexclusionGet) | **Get** /api/v3/importlistexclusion | 
[**ApiV3ImportlistexclusionIdDelete**](ImportListExclusionApi.md#ApiV3ImportlistexclusionIdDelete) | **Delete** /api/v3/importlistexclusion/{id} | 
[**ApiV3ImportlistexclusionIdGet**](ImportListExclusionApi.md#ApiV3ImportlistexclusionIdGet) | **Get** /api/v3/importlistexclusion/{id} | 
[**ApiV3ImportlistexclusionIdPut**](ImportListExclusionApi.md#ApiV3ImportlistexclusionIdPut) | **Put** /api/v3/importlistexclusion/{id} | 
[**ApiV3ImportlistexclusionPost**](ImportListExclusionApi.md#ApiV3ImportlistexclusionPost) | **Post** /api/v3/importlistexclusion | 



## ApiV3ImportlistexclusionGet

> []ImportListExclusionResource ApiV3ImportlistexclusionGet(ctx).Execute()



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
    resp, r, err := apiClient.ImportListExclusionApi.ApiV3ImportlistexclusionGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.ApiV3ImportlistexclusionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ImportlistexclusionGet`: []ImportListExclusionResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionApi.ApiV3ImportlistexclusionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ImportlistexclusionGetRequest struct via the builder pattern


### Return type

[**[]ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ImportlistexclusionIdDelete

> ApiV3ImportlistexclusionIdDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.ImportListExclusionApi.ApiV3ImportlistexclusionIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.ApiV3ImportlistexclusionIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3ImportlistexclusionIdDeleteRequest struct via the builder pattern


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


## ApiV3ImportlistexclusionIdGet

> ImportListExclusionResource ApiV3ImportlistexclusionIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.ImportListExclusionApi.ApiV3ImportlistexclusionIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.ApiV3ImportlistexclusionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ImportlistexclusionIdGet`: ImportListExclusionResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionApi.ApiV3ImportlistexclusionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ImportlistexclusionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ImportlistexclusionIdPut

> ImportListExclusionResource ApiV3ImportlistexclusionIdPut(ctx, id).ImportListExclusionResource(importListExclusionResource).Execute()



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
    importListExclusionResource := *openapiclient.NewImportListExclusionResource() // ImportListExclusionResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportListExclusionApi.ApiV3ImportlistexclusionIdPut(context.Background(), id).ImportListExclusionResource(importListExclusionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.ApiV3ImportlistexclusionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ImportlistexclusionIdPut`: ImportListExclusionResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionApi.ApiV3ImportlistexclusionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ImportlistexclusionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importListExclusionResource** | [**ImportListExclusionResource**](ImportListExclusionResource.md) |  | 

### Return type

[**ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ImportlistexclusionPost

> ImportListExclusionResource ApiV3ImportlistexclusionPost(ctx).ImportListExclusionResource(importListExclusionResource).Execute()



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
    importListExclusionResource := *openapiclient.NewImportListExclusionResource() // ImportListExclusionResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportListExclusionApi.ApiV3ImportlistexclusionPost(context.Background()).ImportListExclusionResource(importListExclusionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionApi.ApiV3ImportlistexclusionPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ImportlistexclusionPost`: ImportListExclusionResource
    fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionApi.ApiV3ImportlistexclusionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ImportlistexclusionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importListExclusionResource** | [**ImportListExclusionResource**](ImportListExclusionResource.md) |  | 

### Return type

[**ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

