# \DownloadClientApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3DownloadclientActionNamePost**](DownloadClientApi.md#ApiV3DownloadclientActionNamePost) | **Post** /api/v3/downloadclient/action/{name} | 
[**ApiV3DownloadclientGet**](DownloadClientApi.md#ApiV3DownloadclientGet) | **Get** /api/v3/downloadclient | 
[**ApiV3DownloadclientIdDelete**](DownloadClientApi.md#ApiV3DownloadclientIdDelete) | **Delete** /api/v3/downloadclient/{id} | 
[**ApiV3DownloadclientIdGet**](DownloadClientApi.md#ApiV3DownloadclientIdGet) | **Get** /api/v3/downloadclient/{id} | 
[**ApiV3DownloadclientIdPut**](DownloadClientApi.md#ApiV3DownloadclientIdPut) | **Put** /api/v3/downloadclient/{id} | 
[**ApiV3DownloadclientPost**](DownloadClientApi.md#ApiV3DownloadclientPost) | **Post** /api/v3/downloadclient | 
[**ApiV3DownloadclientSchemaGet**](DownloadClientApi.md#ApiV3DownloadclientSchemaGet) | **Get** /api/v3/downloadclient/schema | 
[**ApiV3DownloadclientTestPost**](DownloadClientApi.md#ApiV3DownloadclientTestPost) | **Post** /api/v3/downloadclient/test | 
[**ApiV3DownloadclientTestallPost**](DownloadClientApi.md#ApiV3DownloadclientTestallPost) | **Post** /api/v3/downloadclient/testall | 



## ApiV3DownloadclientActionNamePost

> ApiV3DownloadclientActionNamePost(ctx, name).DownloadClientResource(downloadClientResource).Execute()



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
    name := "name_example" // string | 
    downloadClientResource := *openapiclient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.ApiV3DownloadclientActionNamePost(context.Background(), name).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.ApiV3DownloadclientActionNamePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DownloadclientActionNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

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


## ApiV3DownloadclientGet

> []DownloadClientResource ApiV3DownloadclientGet(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.ApiV3DownloadclientGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.ApiV3DownloadclientGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3DownloadclientGet`: []DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.ApiV3DownloadclientGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DownloadclientGetRequest struct via the builder pattern


### Return type

[**[]DownloadClientResource**](DownloadClientResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3DownloadclientIdDelete

> ApiV3DownloadclientIdDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.ApiV3DownloadclientIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.ApiV3DownloadclientIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3DownloadclientIdDeleteRequest struct via the builder pattern


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


## ApiV3DownloadclientIdGet

> DownloadClientResource ApiV3DownloadclientIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.ApiV3DownloadclientIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.ApiV3DownloadclientIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3DownloadclientIdGet`: DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.ApiV3DownloadclientIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DownloadclientIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3DownloadclientIdPut

> DownloadClientResource ApiV3DownloadclientIdPut(ctx, id).DownloadClientResource(downloadClientResource).Execute()



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
    downloadClientResource := *openapiclient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.ApiV3DownloadclientIdPut(context.Background(), id).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.ApiV3DownloadclientIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3DownloadclientIdPut`: DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.ApiV3DownloadclientIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DownloadclientIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3DownloadclientPost

> DownloadClientResource ApiV3DownloadclientPost(ctx).DownloadClientResource(downloadClientResource).Execute()



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
    downloadClientResource := *openapiclient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.ApiV3DownloadclientPost(context.Background()).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.ApiV3DownloadclientPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3DownloadclientPost`: DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.ApiV3DownloadclientPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DownloadclientPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

### Return type

[**DownloadClientResource**](DownloadClientResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3DownloadclientSchemaGet

> []DownloadClientResource ApiV3DownloadclientSchemaGet(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.ApiV3DownloadclientSchemaGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.ApiV3DownloadclientSchemaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3DownloadclientSchemaGet`: []DownloadClientResource
    fmt.Fprintf(os.Stdout, "Response from `DownloadClientApi.ApiV3DownloadclientSchemaGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DownloadclientSchemaGetRequest struct via the builder pattern


### Return type

[**[]DownloadClientResource**](DownloadClientResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3DownloadclientTestPost

> ApiV3DownloadclientTestPost(ctx).DownloadClientResource(downloadClientResource).Execute()



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
    downloadClientResource := *openapiclient.NewDownloadClientResource() // DownloadClientResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DownloadClientApi.ApiV3DownloadclientTestPost(context.Background()).DownloadClientResource(downloadClientResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.ApiV3DownloadclientTestPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DownloadclientTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadClientResource** | [**DownloadClientResource**](DownloadClientResource.md) |  | 

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


## ApiV3DownloadclientTestallPost

> ApiV3DownloadclientTestallPost(ctx).Execute()



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
    resp, r, err := apiClient.DownloadClientApi.ApiV3DownloadclientTestallPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DownloadClientApi.ApiV3DownloadclientTestallPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DownloadclientTestallPostRequest struct via the builder pattern


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

