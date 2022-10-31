# \RemotePathMappingApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3RemotepathmappingGet**](RemotePathMappingApi.md#ApiV3RemotepathmappingGet) | **Get** /api/v3/remotepathmapping | 
[**ApiV3RemotepathmappingIdDelete**](RemotePathMappingApi.md#ApiV3RemotepathmappingIdDelete) | **Delete** /api/v3/remotepathmapping/{id} | 
[**ApiV3RemotepathmappingIdGet**](RemotePathMappingApi.md#ApiV3RemotepathmappingIdGet) | **Get** /api/v3/remotepathmapping/{id} | 
[**ApiV3RemotepathmappingIdPut**](RemotePathMappingApi.md#ApiV3RemotepathmappingIdPut) | **Put** /api/v3/remotepathmapping/{id} | 
[**ApiV3RemotepathmappingPost**](RemotePathMappingApi.md#ApiV3RemotepathmappingPost) | **Post** /api/v3/remotepathmapping | 



## ApiV3RemotepathmappingGet

> []RemotePathMappingResource ApiV3RemotepathmappingGet(ctx).Execute()



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
    resp, r, err := apiClient.RemotePathMappingApi.ApiV3RemotepathmappingGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.ApiV3RemotepathmappingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3RemotepathmappingGet`: []RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.ApiV3RemotepathmappingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3RemotepathmappingGetRequest struct via the builder pattern


### Return type

[**[]RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3RemotepathmappingIdDelete

> ApiV3RemotepathmappingIdDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.RemotePathMappingApi.ApiV3RemotepathmappingIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.ApiV3RemotepathmappingIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3RemotepathmappingIdDeleteRequest struct via the builder pattern


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


## ApiV3RemotepathmappingIdGet

> RemotePathMappingResource ApiV3RemotepathmappingIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.RemotePathMappingApi.ApiV3RemotepathmappingIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.ApiV3RemotepathmappingIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3RemotepathmappingIdGet`: RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.ApiV3RemotepathmappingIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3RemotepathmappingIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3RemotepathmappingIdPut

> RemotePathMappingResource ApiV3RemotepathmappingIdPut(ctx, id).RemotePathMappingResource(remotePathMappingResource).Execute()



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
    remotePathMappingResource := *openapiclient.NewRemotePathMappingResource() // RemotePathMappingResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotePathMappingApi.ApiV3RemotepathmappingIdPut(context.Background(), id).RemotePathMappingResource(remotePathMappingResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.ApiV3RemotepathmappingIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3RemotepathmappingIdPut`: RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.ApiV3RemotepathmappingIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3RemotepathmappingIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remotePathMappingResource** | [**RemotePathMappingResource**](RemotePathMappingResource.md) |  | 

### Return type

[**RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3RemotepathmappingPost

> RemotePathMappingResource ApiV3RemotepathmappingPost(ctx).RemotePathMappingResource(remotePathMappingResource).Execute()



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
    remotePathMappingResource := *openapiclient.NewRemotePathMappingResource() // RemotePathMappingResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemotePathMappingApi.ApiV3RemotepathmappingPost(context.Background()).RemotePathMappingResource(remotePathMappingResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemotePathMappingApi.ApiV3RemotepathmappingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3RemotepathmappingPost`: RemotePathMappingResource
    fmt.Fprintf(os.Stdout, "Response from `RemotePathMappingApi.ApiV3RemotepathmappingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3RemotepathmappingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remotePathMappingResource** | [**RemotePathMappingResource**](RemotePathMappingResource.md) |  | 

### Return type

[**RemotePathMappingResource**](RemotePathMappingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

