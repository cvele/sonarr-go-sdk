# \CalendarApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3CalendarGet**](CalendarApi.md#ApiV3CalendarGet) | **Get** /api/v3/calendar | 
[**ApiV3CalendarIdGet**](CalendarApi.md#ApiV3CalendarIdGet) | **Get** /api/v3/calendar/{id} | 



## ApiV3CalendarGet

> []EpisodeResource ApiV3CalendarGet(ctx).Start(start).End(end).Unmonitored(unmonitored).IncludeSeries(includeSeries).IncludeEpisodeFile(includeEpisodeFile).IncludeEpisodeImages(includeEpisodeImages).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    start := time.Now() // time.Time |  (optional)
    end := time.Now() // time.Time |  (optional)
    unmonitored := true // bool |  (optional) (default to false)
    includeSeries := true // bool |  (optional) (default to false)
    includeEpisodeFile := true // bool |  (optional) (default to false)
    includeEpisodeImages := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalendarApi.ApiV3CalendarGet(context.Background()).Start(start).End(end).Unmonitored(unmonitored).IncludeSeries(includeSeries).IncludeEpisodeFile(includeEpisodeFile).IncludeEpisodeImages(includeEpisodeImages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarApi.ApiV3CalendarGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3CalendarGet`: []EpisodeResource
    fmt.Fprintf(os.Stdout, "Response from `CalendarApi.ApiV3CalendarGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CalendarGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 
 **unmonitored** | **bool** |  | [default to false]
 **includeSeries** | **bool** |  | [default to false]
 **includeEpisodeFile** | **bool** |  | [default to false]
 **includeEpisodeImages** | **bool** |  | [default to false]

### Return type

[**[]EpisodeResource**](EpisodeResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3CalendarIdGet

> EpisodeResource ApiV3CalendarIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.CalendarApi.ApiV3CalendarIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarApi.ApiV3CalendarIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3CalendarIdGet`: EpisodeResource
    fmt.Fprintf(os.Stdout, "Response from `CalendarApi.ApiV3CalendarIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3CalendarIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EpisodeResource**](EpisodeResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

