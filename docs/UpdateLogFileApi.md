# \UpdateLogFileApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3LogFileUpdateFilenameGet**](UpdateLogFileApi.md#ApiV3LogFileUpdateFilenameGet) | **Get** /api/v3/log/file/update/{filename} | 
[**ApiV3LogFileUpdateGet**](UpdateLogFileApi.md#ApiV3LogFileUpdateGet) | **Get** /api/v3/log/file/update | 



## ApiV3LogFileUpdateFilenameGet

> ApiV3LogFileUpdateFilenameGet(ctx, filename).Execute()



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
    filename := "filename_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdateLogFileApi.ApiV3LogFileUpdateFilenameGet(context.Background(), filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateLogFileApi.ApiV3LogFileUpdateFilenameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3LogFileUpdateFilenameGetRequest struct via the builder pattern


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


## ApiV3LogFileUpdateGet

> []LogFileResource ApiV3LogFileUpdateGet(ctx).Execute()



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
    resp, r, err := apiClient.UpdateLogFileApi.ApiV3LogFileUpdateGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateLogFileApi.ApiV3LogFileUpdateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3LogFileUpdateGet`: []LogFileResource
    fmt.Fprintf(os.Stdout, "Response from `UpdateLogFileApi.ApiV3LogFileUpdateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3LogFileUpdateGetRequest struct via the builder pattern


### Return type

[**[]LogFileResource**](LogFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

