# \UiConfigApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ConfigUiGet**](UiConfigApi.md#ApiV3ConfigUiGet) | **Get** /api/v3/config/ui | 
[**ApiV3ConfigUiIdGet**](UiConfigApi.md#ApiV3ConfigUiIdGet) | **Get** /api/v3/config/ui/{id} | 
[**ApiV3ConfigUiIdPut**](UiConfigApi.md#ApiV3ConfigUiIdPut) | **Put** /api/v3/config/ui/{id} | 



## ApiV3ConfigUiGet

> UiConfigResource ApiV3ConfigUiGet(ctx).Execute()



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
    resp, r, err := apiClient.UiConfigApi.ApiV3ConfigUiGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UiConfigApi.ApiV3ConfigUiGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigUiGet`: UiConfigResource
    fmt.Fprintf(os.Stdout, "Response from `UiConfigApi.ApiV3ConfigUiGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigUiGetRequest struct via the builder pattern


### Return type

[**UiConfigResource**](UiConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ConfigUiIdGet

> UiConfigResource ApiV3ConfigUiIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.UiConfigApi.ApiV3ConfigUiIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UiConfigApi.ApiV3ConfigUiIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigUiIdGet`: UiConfigResource
    fmt.Fprintf(os.Stdout, "Response from `UiConfigApi.ApiV3ConfigUiIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigUiIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UiConfigResource**](UiConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ConfigUiIdPut

> UiConfigResource ApiV3ConfigUiIdPut(ctx, id).UiConfigResource(uiConfigResource).Execute()



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
    uiConfigResource := *openapiclient.NewUiConfigResource() // UiConfigResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UiConfigApi.ApiV3ConfigUiIdPut(context.Background(), id).UiConfigResource(uiConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UiConfigApi.ApiV3ConfigUiIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigUiIdPut`: UiConfigResource
    fmt.Fprintf(os.Stdout, "Response from `UiConfigApi.ApiV3ConfigUiIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigUiIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uiConfigResource** | [**UiConfigResource**](UiConfigResource.md) |  | 

### Return type

[**UiConfigResource**](UiConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

