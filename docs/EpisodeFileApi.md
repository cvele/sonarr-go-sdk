# \EpisodeFileApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3EpisodefileBulkDelete**](EpisodeFileApi.md#ApiV3EpisodefileBulkDelete) | **Delete** /api/v3/episodefile/bulk | 
[**ApiV3EpisodefileBulkPut**](EpisodeFileApi.md#ApiV3EpisodefileBulkPut) | **Put** /api/v3/episodefile/bulk | 
[**ApiV3EpisodefileEditorPut**](EpisodeFileApi.md#ApiV3EpisodefileEditorPut) | **Put** /api/v3/episodefile/editor | 
[**ApiV3EpisodefileGet**](EpisodeFileApi.md#ApiV3EpisodefileGet) | **Get** /api/v3/episodefile | 
[**ApiV3EpisodefileIdDelete**](EpisodeFileApi.md#ApiV3EpisodefileIdDelete) | **Delete** /api/v3/episodefile/{id} | 
[**ApiV3EpisodefileIdGet**](EpisodeFileApi.md#ApiV3EpisodefileIdGet) | **Get** /api/v3/episodefile/{id} | 
[**ApiV3EpisodefileIdPut**](EpisodeFileApi.md#ApiV3EpisodefileIdPut) | **Put** /api/v3/episodefile/{id} | 



## ApiV3EpisodefileBulkDelete

> ApiV3EpisodefileBulkDelete(ctx).EpisodeFileListResource(episodeFileListResource).Execute()



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
    episodeFileListResource := *openapiclient.NewEpisodeFileListResource() // EpisodeFileListResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodeFileApi.ApiV3EpisodefileBulkDelete(context.Background()).EpisodeFileListResource(episodeFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.ApiV3EpisodefileBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3EpisodefileBulkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **episodeFileListResource** | [**EpisodeFileListResource**](EpisodeFileListResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3EpisodefileBulkPut

> ApiV3EpisodefileBulkPut(ctx).EpisodeFileResource(episodeFileResource).Execute()



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
    episodeFileResource := []openapiclient.EpisodeFileResource{*openapiclient.NewEpisodeFileResource()} // []EpisodeFileResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodeFileApi.ApiV3EpisodefileBulkPut(context.Background()).EpisodeFileResource(episodeFileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.ApiV3EpisodefileBulkPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3EpisodefileBulkPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **episodeFileResource** | [**[]EpisodeFileResource**](EpisodeFileResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3EpisodefileEditorPut

> ApiV3EpisodefileEditorPut(ctx).EpisodeFileListResource(episodeFileListResource).Execute()



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
    episodeFileListResource := *openapiclient.NewEpisodeFileListResource() // EpisodeFileListResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodeFileApi.ApiV3EpisodefileEditorPut(context.Background()).EpisodeFileListResource(episodeFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.ApiV3EpisodefileEditorPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3EpisodefileEditorPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **episodeFileListResource** | [**EpisodeFileListResource**](EpisodeFileListResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3EpisodefileGet

> []EpisodeFileResource ApiV3EpisodefileGet(ctx).SeriesId(seriesId).EpisodeFileIds(episodeFileIds).Execute()



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
    seriesId := int32(56) // int32 |  (optional)
    episodeFileIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodeFileApi.ApiV3EpisodefileGet(context.Background()).SeriesId(seriesId).EpisodeFileIds(episodeFileIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.ApiV3EpisodefileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3EpisodefileGet`: []EpisodeFileResource
    fmt.Fprintf(os.Stdout, "Response from `EpisodeFileApi.ApiV3EpisodefileGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3EpisodefileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesId** | **int32** |  | 
 **episodeFileIds** | **[]int32** |  | 

### Return type

[**[]EpisodeFileResource**](EpisodeFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3EpisodefileIdDelete

> ApiV3EpisodefileIdDelete(ctx, id).Execute()



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
    resp, r, err := apiClient.EpisodeFileApi.ApiV3EpisodefileIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.ApiV3EpisodefileIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV3EpisodefileIdDeleteRequest struct via the builder pattern


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


## ApiV3EpisodefileIdGet

> EpisodeFileResource ApiV3EpisodefileIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.EpisodeFileApi.ApiV3EpisodefileIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.ApiV3EpisodefileIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3EpisodefileIdGet`: EpisodeFileResource
    fmt.Fprintf(os.Stdout, "Response from `EpisodeFileApi.ApiV3EpisodefileIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3EpisodefileIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EpisodeFileResource**](EpisodeFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3EpisodefileIdPut

> EpisodeFileResource ApiV3EpisodefileIdPut(ctx, id).EpisodeFileResource(episodeFileResource).Execute()



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
    episodeFileResource := *openapiclient.NewEpisodeFileResource() // EpisodeFileResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodeFileApi.ApiV3EpisodefileIdPut(context.Background(), id).EpisodeFileResource(episodeFileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodeFileApi.ApiV3EpisodefileIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3EpisodefileIdPut`: EpisodeFileResource
    fmt.Fprintf(os.Stdout, "Response from `EpisodeFileApi.ApiV3EpisodefileIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3EpisodefileIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **episodeFileResource** | [**EpisodeFileResource**](EpisodeFileResource.md) |  | 

### Return type

[**EpisodeFileResource**](EpisodeFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

