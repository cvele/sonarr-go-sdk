# \LanguageProfileApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3LanguageprofileGet**](LanguageProfileApi.md#ApiV3LanguageprofileGet) | **Get** /api/v3/languageprofile | 
[**ApiV3LanguageprofileIdDelete**](LanguageProfileApi.md#ApiV3LanguageprofileIdDelete) | **Delete** /api/v3/languageprofile/{id} | 
[**ApiV3LanguageprofileIdGet**](LanguageProfileApi.md#ApiV3LanguageprofileIdGet) | **Get** /api/v3/languageprofile/{id} | 
[**ApiV3LanguageprofileIdPut**](LanguageProfileApi.md#ApiV3LanguageprofileIdPut) | **Put** /api/v3/languageprofile/{id} | 
[**ApiV3LanguageprofilePost**](LanguageProfileApi.md#ApiV3LanguageprofilePost) | **Post** /api/v3/languageprofile | 



## ApiV3LanguageprofileGet

> []LanguageProfileResource ApiV3LanguageprofileGet(ctx).Execute()



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
    resp, r, err := apiClient.LanguageProfileApi.ApiV3LanguageprofileGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.ApiV3LanguageprofileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3LanguageprofileGet`: []LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileApi.ApiV3LanguageprofileGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3LanguageprofileGetRequest struct via the builder pattern


### Return type

[**[]LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3LanguageprofileIdDelete

> ApiV3LanguageprofileIdDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.LanguageProfileApi.ApiV3LanguageprofileIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.ApiV3LanguageprofileIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3LanguageprofileIdDeleteRequest struct via the builder pattern


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


## ApiV3LanguageprofileIdGet

> LanguageProfileResource ApiV3LanguageprofileIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.LanguageProfileApi.ApiV3LanguageprofileIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.ApiV3LanguageprofileIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3LanguageprofileIdGet`: LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileApi.ApiV3LanguageprofileIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3LanguageprofileIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3LanguageprofileIdPut

> LanguageProfileResource ApiV3LanguageprofileIdPut(ctx, id).LanguageProfileResource(languageProfileResource).Execute()



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
    languageProfileResource := *openapiclient.NewLanguageProfileResource() // LanguageProfileResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageProfileApi.ApiV3LanguageprofileIdPut(context.Background(), id).LanguageProfileResource(languageProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.ApiV3LanguageprofileIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3LanguageprofileIdPut`: LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileApi.ApiV3LanguageprofileIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3LanguageprofileIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **languageProfileResource** | [**LanguageProfileResource**](LanguageProfileResource.md) |  | 

### Return type

[**LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3LanguageprofilePost

> LanguageProfileResource ApiV3LanguageprofilePost(ctx).LanguageProfileResource(languageProfileResource).Execute()



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
    languageProfileResource := *openapiclient.NewLanguageProfileResource() // LanguageProfileResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageProfileApi.ApiV3LanguageprofilePost(context.Background()).LanguageProfileResource(languageProfileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileApi.ApiV3LanguageprofilePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3LanguageprofilePost`: LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileApi.ApiV3LanguageprofilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3LanguageprofilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **languageProfileResource** | [**LanguageProfileResource**](LanguageProfileResource.md) |  | 

### Return type

[**LanguageProfileResource**](LanguageProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

