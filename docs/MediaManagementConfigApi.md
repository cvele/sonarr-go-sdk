# \MediaManagementConfigApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ConfigMediamanagementGet**](MediaManagementConfigApi.md#ApiV3ConfigMediamanagementGet) | **Get** /api/v3/config/mediamanagement | 
[**ApiV3ConfigMediamanagementIdGet**](MediaManagementConfigApi.md#ApiV3ConfigMediamanagementIdGet) | **Get** /api/v3/config/mediamanagement/{id} | 
[**ApiV3ConfigMediamanagementIdPut**](MediaManagementConfigApi.md#ApiV3ConfigMediamanagementIdPut) | **Put** /api/v3/config/mediamanagement/{id} | 



## ApiV3ConfigMediamanagementGet

> MediaManagementConfigResource ApiV3ConfigMediamanagementGet(ctx).Execute()



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
    resp, r, err := apiClient.MediaManagementConfigApi.ApiV3ConfigMediamanagementGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigApi.ApiV3ConfigMediamanagementGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigMediamanagementGet`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigApi.ApiV3ConfigMediamanagementGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigMediamanagementGetRequest struct via the builder pattern


### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ConfigMediamanagementIdGet

> MediaManagementConfigResource ApiV3ConfigMediamanagementIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.MediaManagementConfigApi.ApiV3ConfigMediamanagementIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigApi.ApiV3ConfigMediamanagementIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigMediamanagementIdGet`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigApi.ApiV3ConfigMediamanagementIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigMediamanagementIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ConfigMediamanagementIdPut

> MediaManagementConfigResource ApiV3ConfigMediamanagementIdPut(ctx, id).MediaManagementConfigResource(mediaManagementConfigResource).Execute()



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
    mediaManagementConfigResource := *openapiclient.NewMediaManagementConfigResource() // MediaManagementConfigResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaManagementConfigApi.ApiV3ConfigMediamanagementIdPut(context.Background(), id).MediaManagementConfigResource(mediaManagementConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaManagementConfigApi.ApiV3ConfigMediamanagementIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigMediamanagementIdPut`: MediaManagementConfigResource
    fmt.Fprintf(os.Stdout, "Response from `MediaManagementConfigApi.ApiV3ConfigMediamanagementIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigMediamanagementIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaManagementConfigResource** | [**MediaManagementConfigResource**](MediaManagementConfigResource.md) |  | 

### Return type

[**MediaManagementConfigResource**](MediaManagementConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

