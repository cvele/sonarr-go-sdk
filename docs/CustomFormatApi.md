# \CustomFormatApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3CustomformatGet**](CustomFormatApi.md#ApiV3CustomformatGet) | **Get** /api/v3/customformat | 
[**ApiV3CustomformatIdDelete**](CustomFormatApi.md#ApiV3CustomformatIdDelete) | **Delete** /api/v3/customformat/{id} | 
[**ApiV3CustomformatIdGet**](CustomFormatApi.md#ApiV3CustomformatIdGet) | **Get** /api/v3/customformat/{id} | 
[**ApiV3CustomformatIdPut**](CustomFormatApi.md#ApiV3CustomformatIdPut) | **Put** /api/v3/customformat/{id} | 
[**ApiV3CustomformatPost**](CustomFormatApi.md#ApiV3CustomformatPost) | **Post** /api/v3/customformat | 
[**ApiV3CustomformatSchemaGet**](CustomFormatApi.md#ApiV3CustomformatSchemaGet) | **Get** /api/v3/customformat/schema | 



## ApiV3CustomformatGet

> []CustomFormatResource ApiV3CustomformatGet(ctx).Execute()



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
    resp, r, err := apiClient.CustomFormatApi.ApiV3CustomformatGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormatApi.ApiV3CustomformatGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3CustomformatGet`: []CustomFormatResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFormatApi.ApiV3CustomformatGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomformatGetRequest struct via the builder pattern


### Return type

[**[]CustomFormatResource**](CustomFormatResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3CustomformatIdDelete

> ApiV3CustomformatIdDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.CustomFormatApi.ApiV3CustomformatIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormatApi.ApiV3CustomformatIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3CustomformatIdDeleteRequest struct via the builder pattern


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


## ApiV3CustomformatIdGet

> CustomFormatResource ApiV3CustomformatIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.CustomFormatApi.ApiV3CustomformatIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormatApi.ApiV3CustomformatIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3CustomformatIdGet`: CustomFormatResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFormatApi.ApiV3CustomformatIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomformatIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomFormatResource**](CustomFormatResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3CustomformatIdPut

> CustomFormatResource ApiV3CustomformatIdPut(ctx, id).CustomFormatResource(customFormatResource).Execute()



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
    customFormatResource := *openapiclient.NewCustomFormatResource() // CustomFormatResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormatApi.ApiV3CustomformatIdPut(context.Background(), id).CustomFormatResource(customFormatResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormatApi.ApiV3CustomformatIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3CustomformatIdPut`: CustomFormatResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFormatApi.ApiV3CustomformatIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomformatIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customFormatResource** | [**CustomFormatResource**](CustomFormatResource.md) |  | 

### Return type

[**CustomFormatResource**](CustomFormatResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3CustomformatPost

> CustomFormatResource ApiV3CustomformatPost(ctx).CustomFormatResource(customFormatResource).Execute()



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
    customFormatResource := *openapiclient.NewCustomFormatResource() // CustomFormatResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormatApi.ApiV3CustomformatPost(context.Background()).CustomFormatResource(customFormatResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormatApi.ApiV3CustomformatPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3CustomformatPost`: CustomFormatResource
    fmt.Fprintf(os.Stdout, "Response from `CustomFormatApi.ApiV3CustomformatPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomformatPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customFormatResource** | [**CustomFormatResource**](CustomFormatResource.md) |  | 

### Return type

[**CustomFormatResource**](CustomFormatResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3CustomformatSchemaGet

> ApiV3CustomformatSchemaGet(ctx).Execute()



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
    resp, r, err := apiClient.CustomFormatApi.ApiV3CustomformatSchemaGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormatApi.ApiV3CustomformatSchemaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CustomformatSchemaGetRequest struct via the builder pattern


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

