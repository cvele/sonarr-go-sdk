# \DelayProfileApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3DelayprofileGet**](DelayProfileApi.md#ApiV3DelayprofileGet) | **Get** /api/v3/delayprofile | 
[**ApiV3DelayprofileIdDelete**](DelayProfileApi.md#ApiV3DelayprofileIdDelete) | **Delete** /api/v3/delayprofile/{id} | 
[**ApiV3DelayprofileIdGet**](DelayProfileApi.md#ApiV3DelayprofileIdGet) | **Get** /api/v3/delayprofile/{id} | 
[**ApiV3DelayprofileIdPut**](DelayProfileApi.md#ApiV3DelayprofileIdPut) | **Put** /api/v3/delayprofile/{id} | 
[**ApiV3DelayprofilePost**](DelayProfileApi.md#ApiV3DelayprofilePost) | **Post** /api/v3/delayprofile | 
[**ApiV3DelayprofileReorderIdPut**](DelayProfileApi.md#ApiV3DelayprofileReorderIdPut) | **Put** /api/v3/delayprofile/reorder/{id} | 



## ApiV3DelayprofileGet

> []DelayProfileResource ApiV3DelayprofileGet(ctx).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.ApiV3DelayprofileGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.ApiV3DelayprofileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3DelayprofileGet`: []DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.ApiV3DelayprofileGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DelayprofileGetRequest struct via the builder pattern


### Return type

[**[]DelayProfileResource**](DelayProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3DelayprofileIdDelete

> ApiV3DelayprofileIdDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.ApiV3DelayprofileIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.ApiV3DelayprofileIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3DelayprofileIdDeleteRequest struct via the builder pattern


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


## ApiV3DelayprofileIdGet

> DelayProfileResource ApiV3DelayprofileIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.DelayProfileApi.ApiV3DelayprofileIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.ApiV3DelayprofileIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3DelayprofileIdGet`: DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.ApiV3DelayprofileIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DelayprofileIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DelayProfileResource**](DelayProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3DelayprofileIdPut

> DelayProfileResource ApiV3DelayprofileIdPut(ctx, id).DelayProfileResource(delayProfileResource).Execute()



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
    delayProfileResource := *openapiclient.NewDelayProfileResource() // DelayProfileResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelayProfileApi.ApiV3DelayprofileIdPut(context.Background(), id).DelayProfileResource(delayProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.ApiV3DelayprofileIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3DelayprofileIdPut`: DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.ApiV3DelayprofileIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DelayprofileIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delayProfileResource** | [**DelayProfileResource**](DelayProfileResource.md) |  | 

### Return type

[**DelayProfileResource**](DelayProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3DelayprofilePost

> DelayProfileResource ApiV3DelayprofilePost(ctx).DelayProfileResource(delayProfileResource).Execute()



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
    delayProfileResource := *openapiclient.NewDelayProfileResource() // DelayProfileResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelayProfileApi.ApiV3DelayprofilePost(context.Background()).DelayProfileResource(delayProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.ApiV3DelayprofilePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3DelayprofilePost`: DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.ApiV3DelayprofilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DelayprofilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delayProfileResource** | [**DelayProfileResource**](DelayProfileResource.md) |  | 

### Return type

[**DelayProfileResource**](DelayProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3DelayprofileReorderIdPut

> []DelayProfileResource ApiV3DelayprofileReorderIdPut(ctx, id).After(after).Execute()



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
    after := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DelayProfileApi.ApiV3DelayprofileReorderIdPut(context.Background(), id).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DelayProfileApi.ApiV3DelayprofileReorderIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3DelayprofileReorderIdPut`: []DelayProfileResource
    fmt.Fprintf(os.Stdout, "Response from `DelayProfileApi.ApiV3DelayprofileReorderIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DelayprofileReorderIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **int32** |  | 

### Return type

[**[]DelayProfileResource**](DelayProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

