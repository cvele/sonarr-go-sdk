# \HostConfigApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ConfigHostGet**](HostConfigApi.md#ApiV3ConfigHostGet) | **Get** /api/v3/config/host | 
[**ApiV3ConfigHostIdGet**](HostConfigApi.md#ApiV3ConfigHostIdGet) | **Get** /api/v3/config/host/{id} | 
[**ApiV3ConfigHostIdPut**](HostConfigApi.md#ApiV3ConfigHostIdPut) | **Put** /api/v3/config/host/{id} | 



## ApiV3ConfigHostGet

> HostConfigResource ApiV3ConfigHostGet(ctx).Execute()



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
    resp, r, err := apiClient.HostConfigApi.ApiV3ConfigHostGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostConfigApi.ApiV3ConfigHostGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigHostGet`: HostConfigResource
    fmt.Fprintf(os.Stdout, "Response from `HostConfigApi.ApiV3ConfigHostGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigHostGetRequest struct via the builder pattern


### Return type

[**HostConfigResource**](HostConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ConfigHostIdGet

> HostConfigResource ApiV3ConfigHostIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.HostConfigApi.ApiV3ConfigHostIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostConfigApi.ApiV3ConfigHostIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigHostIdGet`: HostConfigResource
    fmt.Fprintf(os.Stdout, "Response from `HostConfigApi.ApiV3ConfigHostIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigHostIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostConfigResource**](HostConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ConfigHostIdPut

> HostConfigResource ApiV3ConfigHostIdPut(ctx, id).HostConfigResource(hostConfigResource).Execute()



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
    hostConfigResource := *openapiclient.NewHostConfigResource() // HostConfigResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HostConfigApi.ApiV3ConfigHostIdPut(context.Background(), id).HostConfigResource(hostConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostConfigApi.ApiV3ConfigHostIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigHostIdPut`: HostConfigResource
    fmt.Fprintf(os.Stdout, "Response from `HostConfigApi.ApiV3ConfigHostIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigHostIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostConfigResource** | [**HostConfigResource**](HostConfigResource.md) |  | 

### Return type

[**HostConfigResource**](HostConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

