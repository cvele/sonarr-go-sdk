# \CustomFilterApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3CustomfilterGet**](CustomFilterApi.md#ApiV3CustomfilterGet) | **Get** /api/v3/customfilter | 
[**ApiV3CustomfilterIdDelete**](CustomFilterApi.md#ApiV3CustomfilterIdDelete) | **Delete** /api/v3/customfilter/{id} | 
[**ApiV3CustomfilterIdGet**](CustomFilterApi.md#ApiV3CustomfilterIdGet) | **Get** /api/v3/customfilter/{id} | 
[**ApiV3CustomfilterIdPut**](CustomFilterApi.md#ApiV3CustomfilterIdPut) | **Put** /api/v3/customfilter/{id} | 
[**ApiV3CustomfilterPost**](CustomFilterApi.md#ApiV3CustomfilterPost) | **Post** /api/v3/customfilter | 



## ApiV3CustomfilterGet

> []CustomFilterResource ApiV3CustomfilterGet(ctx).Execute()



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
    resp, r, err := apiClient.CustomFilterApi.ApiV3CustomfilterGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterApi.ApiV3CustomfilterGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3CustomfilterGet`: []CustomFilterResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFilterApi.ApiV3CustomfilterGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomfilterGetRequest struct via the builder pattern


### Return type

[**[]CustomFilterResource**](CustomFilterResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3CustomfilterIdDelete

> ApiV3CustomfilterIdDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.CustomFilterApi.ApiV3CustomfilterIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterApi.ApiV3CustomfilterIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3CustomfilterIdDeleteRequest struct via the builder pattern


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


## ApiV3CustomfilterIdGet

> CustomFilterResource ApiV3CustomfilterIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.CustomFilterApi.ApiV3CustomfilterIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterApi.ApiV3CustomfilterIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3CustomfilterIdGet`: CustomFilterResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFilterApi.ApiV3CustomfilterIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomfilterIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomFilterResource**](CustomFilterResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3CustomfilterIdPut

> CustomFilterResource ApiV3CustomfilterIdPut(ctx, id).CustomFilterResource(customFilterResource).Execute()



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
    customFilterResource := *openapiclient.NewCustomFilterResource() // CustomFilterResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFilterApi.ApiV3CustomfilterIdPut(context.Background(), id).CustomFilterResource(customFilterResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterApi.ApiV3CustomfilterIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3CustomfilterIdPut`: CustomFilterResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFilterApi.ApiV3CustomfilterIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomfilterIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customFilterResource** | [**CustomFilterResource**](CustomFilterResource.md) |  | 

### Return type

[**CustomFilterResource**](CustomFilterResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3CustomfilterPost

> CustomFilterResource ApiV3CustomfilterPost(ctx).CustomFilterResource(customFilterResource).Execute()



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
    customFilterResource := *openapiclient.NewCustomFilterResource() // CustomFilterResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFilterApi.ApiV3CustomfilterPost(context.Background()).CustomFilterResource(customFilterResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFilterApi.ApiV3CustomfilterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3CustomfilterPost`: CustomFilterResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFilterApi.ApiV3CustomfilterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomfilterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customFilterResource** | [**CustomFilterResource**](CustomFilterResource.md) |  | 

### Return type

[**CustomFilterResource**](CustomFilterResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

