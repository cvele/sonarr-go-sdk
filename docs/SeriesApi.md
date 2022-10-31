# \SeriesApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3SeriesGet**](SeriesApi.md#ApiV3SeriesGet) | **Get** /api/v3/series | 
[**ApiV3SeriesIdDelete**](SeriesApi.md#ApiV3SeriesIdDelete) | **Delete** /api/v3/series/{id} | 
[**ApiV3SeriesIdGet**](SeriesApi.md#ApiV3SeriesIdGet) | **Get** /api/v3/series/{id} | 
[**ApiV3SeriesIdPut**](SeriesApi.md#ApiV3SeriesIdPut) | **Put** /api/v3/series/{id} | 
[**ApiV3SeriesPost**](SeriesApi.md#ApiV3SeriesPost) | **Post** /api/v3/series | 



## ApiV3SeriesGet

> []SeriesResource ApiV3SeriesGet(ctx).TvdbId(tvdbId).IncludeSeasonImages(includeSeasonImages).Execute()



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
    tvdbId := int32(56) // int32 |  (optional)
    includeSeasonImages := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesApi.ApiV3SeriesGet(context.Background()).TvdbId(tvdbId).IncludeSeasonImages(includeSeasonImages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesApi.ApiV3SeriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3SeriesGet`: []SeriesResource
    fmt.Fprintf(os.Stdout, "Response from `SeriesApi.ApiV3SeriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SeriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tvdbId** | **int32** |  | 
 **includeSeasonImages** | **bool** |  | [default to false]

### Return type

[**[]SeriesResource**](SeriesResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3SeriesIdDelete

> ApiV3SeriesIdDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.SeriesApi.ApiV3SeriesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesApi.ApiV3SeriesIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3SeriesIdDeleteRequest struct via the builder pattern


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


## ApiV3SeriesIdGet

> SeriesResource ApiV3SeriesIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.SeriesApi.ApiV3SeriesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesApi.ApiV3SeriesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3SeriesIdGet`: SeriesResource
    fmt.Fprintf(os.Stdout, "Response from `SeriesApi.ApiV3SeriesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SeriesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SeriesResource**](SeriesResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3SeriesIdPut

> SeriesResource ApiV3SeriesIdPut(ctx, id).SeriesResource(seriesResource).Execute()



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
    seriesResource := *openapiclient.NewSeriesResource() // SeriesResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesApi.ApiV3SeriesIdPut(context.Background(), id).SeriesResource(seriesResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesApi.ApiV3SeriesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3SeriesIdPut`: SeriesResource
    fmt.Fprintf(os.Stdout, "Response from `SeriesApi.ApiV3SeriesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SeriesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seriesResource** | [**SeriesResource**](SeriesResource.md) |  | 

### Return type

[**SeriesResource**](SeriesResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3SeriesPost

> SeriesResource ApiV3SeriesPost(ctx).SeriesResource(seriesResource).Execute()



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
    seriesResource := *openapiclient.NewSeriesResource() // SeriesResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeriesApi.ApiV3SeriesPost(context.Background()).SeriesResource(seriesResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeriesApi.ApiV3SeriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3SeriesPost`: SeriesResource
    fmt.Fprintf(os.Stdout, "Response from `SeriesApi.ApiV3SeriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SeriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesResource** | [**SeriesResource**](SeriesResource.md) |  | 

### Return type

[**SeriesResource**](SeriesResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

