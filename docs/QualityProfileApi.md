# \QualityProfileApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3QualityprofileGet**](QualityProfileApi.md#ApiV3QualityprofileGet) | **Get** /api/v3/qualityprofile | 
[**ApiV3QualityprofileIdDelete**](QualityProfileApi.md#ApiV3QualityprofileIdDelete) | **Delete** /api/v3/qualityprofile/{id} | 
[**ApiV3QualityprofileIdGet**](QualityProfileApi.md#ApiV3QualityprofileIdGet) | **Get** /api/v3/qualityprofile/{id} | 
[**ApiV3QualityprofileIdPut**](QualityProfileApi.md#ApiV3QualityprofileIdPut) | **Put** /api/v3/qualityprofile/{id} | 
[**ApiV3QualityprofilePost**](QualityProfileApi.md#ApiV3QualityprofilePost) | **Post** /api/v3/qualityprofile | 



## ApiV3QualityprofileGet

> []QualityProfileResource ApiV3QualityprofileGet(ctx).Execute()



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
    resp, r, err := apiClient.QualityProfileApi.ApiV3QualityprofileGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.ApiV3QualityprofileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3QualityprofileGet`: []QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileApi.ApiV3QualityprofileGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QualityprofileGetRequest struct via the builder pattern


### Return type

[**[]QualityProfileResource**](QualityProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3QualityprofileIdDelete

> ApiV3QualityprofileIdDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.QualityProfileApi.ApiV3QualityprofileIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.ApiV3QualityprofileIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3QualityprofileIdDeleteRequest struct via the builder pattern


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


## ApiV3QualityprofileIdGet

> QualityProfileResource ApiV3QualityprofileIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.QualityProfileApi.ApiV3QualityprofileIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.ApiV3QualityprofileIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3QualityprofileIdGet`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileApi.ApiV3QualityprofileIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QualityprofileIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QualityProfileResource**](QualityProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3QualityprofileIdPut

> QualityProfileResource ApiV3QualityprofileIdPut(ctx, id).QualityProfileResource(qualityProfileResource).Execute()



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
    qualityProfileResource := *openapiclient.NewQualityProfileResource() // QualityProfileResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityProfileApi.ApiV3QualityprofileIdPut(context.Background(), id).QualityProfileResource(qualityProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.ApiV3QualityprofileIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3QualityprofileIdPut`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileApi.ApiV3QualityprofileIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QualityprofileIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qualityProfileResource** | [**QualityProfileResource**](QualityProfileResource.md) |  | 

### Return type

[**QualityProfileResource**](QualityProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3QualityprofilePost

> QualityProfileResource ApiV3QualityprofilePost(ctx).QualityProfileResource(qualityProfileResource).Execute()



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
    qualityProfileResource := *openapiclient.NewQualityProfileResource() // QualityProfileResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityProfileApi.ApiV3QualityprofilePost(context.Background()).QualityProfileResource(qualityProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileApi.ApiV3QualityprofilePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3QualityprofilePost`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileApi.ApiV3QualityprofilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QualityprofilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qualityProfileResource** | [**QualityProfileResource**](QualityProfileResource.md) |  | 

### Return type

[**QualityProfileResource**](QualityProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

