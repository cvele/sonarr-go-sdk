# \DownloadClientConfigApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ConfigDownloadclientGet**](DownloadClientConfigApi.md#ApiV3ConfigDownloadclientGet) | **Get** /api/v3/config/downloadclient | 
[**ApiV3ConfigDownloadclientIdGet**](DownloadClientConfigApi.md#ApiV3ConfigDownloadclientIdGet) | **Get** /api/v3/config/downloadclient/{id} | 
[**ApiV3ConfigDownloadclientIdPut**](DownloadClientConfigApi.md#ApiV3ConfigDownloadclientIdPut) | **Put** /api/v3/config/downloadclient/{id} | 



## ApiV3ConfigDownloadclientGet

> DownloadClientConfigResource ApiV3ConfigDownloadclientGet(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientConfigApi.ApiV3ConfigDownloadclientGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientConfigApi.ApiV3ConfigDownloadclientGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigDownloadclientGet`: DownloadClientConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientConfigApi.ApiV3ConfigDownloadclientGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigDownloadclientGetRequest struct via the builder pattern


### Return type

[**DownloadClientConfigResource**](DownloadClientConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ConfigDownloadclientIdGet

> DownloadClientConfigResource ApiV3ConfigDownloadclientIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.DownloadClientConfigApi.ApiV3ConfigDownloadclientIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientConfigApi.ApiV3ConfigDownloadclientIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigDownloadclientIdGet`: DownloadClientConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientConfigApi.ApiV3ConfigDownloadclientIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigDownloadclientIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DownloadClientConfigResource**](DownloadClientConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ConfigDownloadclientIdPut

> DownloadClientConfigResource ApiV3ConfigDownloadclientIdPut(ctx, id).DownloadClientConfigResource(downloadClientConfigResource).Execute()



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
    downloadClientConfigResource := *openapiclient.NewDownloadClientConfigResource() // DownloadClientConfigResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientConfigApi.ApiV3ConfigDownloadclientIdPut(context.Background(), id).DownloadClientConfigResource(downloadClientConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientConfigApi.ApiV3ConfigDownloadclientIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigDownloadclientIdPut`: DownloadClientConfigResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientConfigApi.ApiV3ConfigDownloadclientIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigDownloadclientIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadClientConfigResource** | [**DownloadClientConfigResource**](DownloadClientConfigResource.md) |  | 

### Return type

[**DownloadClientConfigResource**](DownloadClientConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

