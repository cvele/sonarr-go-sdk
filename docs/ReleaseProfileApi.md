# \ReleaseProfileApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ReleaseprofileGet**](ReleaseProfileApi.md#ApiV3ReleaseprofileGet) | **Get** /api/v3/releaseprofile | 
[**ApiV3ReleaseprofileIdDelete**](ReleaseProfileApi.md#ApiV3ReleaseprofileIdDelete) | **Delete** /api/v3/releaseprofile/{id} | 
[**ApiV3ReleaseprofileIdGet**](ReleaseProfileApi.md#ApiV3ReleaseprofileIdGet) | **Get** /api/v3/releaseprofile/{id} | 
[**ApiV3ReleaseprofileIdPut**](ReleaseProfileApi.md#ApiV3ReleaseprofileIdPut) | **Put** /api/v3/releaseprofile/{id} | 
[**ApiV3ReleaseprofilePost**](ReleaseProfileApi.md#ApiV3ReleaseprofilePost) | **Post** /api/v3/releaseprofile | 



## ApiV3ReleaseprofileGet

> []ReleaseProfileResource ApiV3ReleaseprofileGet(ctx).Execute()



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
    resp, r, err := apiClient.ReleaseProfileApi.ApiV3ReleaseprofileGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.ApiV3ReleaseprofileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ReleaseprofileGet`: []ReleaseProfileResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseProfileApi.ApiV3ReleaseprofileGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ReleaseprofileGetRequest struct via the builder pattern


### Return type

[**[]ReleaseProfileResource**](ReleaseProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ReleaseprofileIdDelete

> ApiV3ReleaseprofileIdDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.ReleaseProfileApi.ApiV3ReleaseprofileIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.ApiV3ReleaseprofileIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3ReleaseprofileIdDeleteRequest struct via the builder pattern


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


## ApiV3ReleaseprofileIdGet

> ReleaseProfileResource ApiV3ReleaseprofileIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.ReleaseProfileApi.ApiV3ReleaseprofileIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.ApiV3ReleaseprofileIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ReleaseprofileIdGet`: ReleaseProfileResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseProfileApi.ApiV3ReleaseprofileIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ReleaseprofileIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReleaseProfileResource**](ReleaseProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ReleaseprofileIdPut

> ReleaseProfileResource ApiV3ReleaseprofileIdPut(ctx, id).ReleaseProfileResource(releaseProfileResource).Execute()



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
    releaseProfileResource := *openapiclient.NewReleaseProfileResource() // ReleaseProfileResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseProfileApi.ApiV3ReleaseprofileIdPut(context.Background(), id).ReleaseProfileResource(releaseProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.ApiV3ReleaseprofileIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ReleaseprofileIdPut`: ReleaseProfileResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseProfileApi.ApiV3ReleaseprofileIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ReleaseprofileIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **releaseProfileResource** | [**ReleaseProfileResource**](ReleaseProfileResource.md) |  | 

### Return type

[**ReleaseProfileResource**](ReleaseProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ReleaseprofilePost

> ReleaseProfileResource ApiV3ReleaseprofilePost(ctx).ReleaseProfileResource(releaseProfileResource).Execute()



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
    releaseProfileResource := *openapiclient.NewReleaseProfileResource() // ReleaseProfileResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleaseProfileApi.ApiV3ReleaseprofilePost(context.Background()).ReleaseProfileResource(releaseProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseProfileApi.ApiV3ReleaseprofilePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ReleaseprofilePost`: ReleaseProfileResource
    fmt.Fprintf(os.Stdout, "Response from `ReleaseProfileApi.ApiV3ReleaseprofilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ReleaseprofilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseProfileResource** | [**ReleaseProfileResource**](ReleaseProfileResource.md) |  | 

### Return type

[**ReleaseProfileResource**](ReleaseProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

