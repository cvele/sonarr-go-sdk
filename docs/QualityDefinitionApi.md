# \QualityDefinitionApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3QualitydefinitionGet**](QualityDefinitionApi.md#ApiV3QualitydefinitionGet) | **Get** /api/v3/qualitydefinition | 
[**ApiV3QualitydefinitionIdGet**](QualityDefinitionApi.md#ApiV3QualitydefinitionIdGet) | **Get** /api/v3/qualitydefinition/{id} | 
[**ApiV3QualitydefinitionIdPut**](QualityDefinitionApi.md#ApiV3QualitydefinitionIdPut) | **Put** /api/v3/qualitydefinition/{id} | 
[**ApiV3QualitydefinitionUpdatePut**](QualityDefinitionApi.md#ApiV3QualitydefinitionUpdatePut) | **Put** /api/v3/qualitydefinition/update | 



## ApiV3QualitydefinitionGet

> []QualityDefinitionResource ApiV3QualitydefinitionGet(ctx).Execute()



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
    resp, r, err := apiClient.QualityDefinitionApi.ApiV3QualitydefinitionGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.ApiV3QualitydefinitionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3QualitydefinitionGet`: []QualityDefinitionResource
    fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionApi.ApiV3QualitydefinitionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QualitydefinitionGetRequest struct via the builder pattern


### Return type

[**[]QualityDefinitionResource**](QualityDefinitionResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3QualitydefinitionIdGet

> QualityDefinitionResource ApiV3QualitydefinitionIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.QualityDefinitionApi.ApiV3QualitydefinitionIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.ApiV3QualitydefinitionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3QualitydefinitionIdGet`: QualityDefinitionResource
    fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionApi.ApiV3QualitydefinitionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QualitydefinitionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QualityDefinitionResource**](QualityDefinitionResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3QualitydefinitionIdPut

> QualityDefinitionResource ApiV3QualitydefinitionIdPut(ctx, id).QualityDefinitionResource(qualityDefinitionResource).Execute()



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
    qualityDefinitionResource := *openapiclient.NewQualityDefinitionResource() // QualityDefinitionResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityDefinitionApi.ApiV3QualitydefinitionIdPut(context.Background(), id).QualityDefinitionResource(qualityDefinitionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.ApiV3QualitydefinitionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3QualitydefinitionIdPut`: QualityDefinitionResource
    fmt.Fprintf(os.Stdout, "Response from `QualityDefinitionApi.ApiV3QualitydefinitionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QualitydefinitionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qualityDefinitionResource** | [**QualityDefinitionResource**](QualityDefinitionResource.md) |  | 

### Return type

[**QualityDefinitionResource**](QualityDefinitionResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3QualitydefinitionUpdatePut

> ApiV3QualitydefinitionUpdatePut(ctx).QualityDefinitionResource(qualityDefinitionResource).Execute()



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
    qualityDefinitionResource := []openapiclient.QualityDefinitionResource{*openapiclient.NewQualityDefinitionResource()} // []QualityDefinitionResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityDefinitionApi.ApiV3QualitydefinitionUpdatePut(context.Background()).QualityDefinitionResource(qualityDefinitionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityDefinitionApi.ApiV3QualitydefinitionUpdatePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QualitydefinitionUpdatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **qualityDefinitionResource** | [**[]QualityDefinitionResource**](QualityDefinitionResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

