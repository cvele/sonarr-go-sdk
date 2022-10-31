# \FileSystemApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3FilesystemGet**](FileSystemApi.md#ApiV3FilesystemGet) | **Get** /api/v3/filesystem | 
[**ApiV3FilesystemMediafilesGet**](FileSystemApi.md#ApiV3FilesystemMediafilesGet) | **Get** /api/v3/filesystem/mediafiles | 
[**ApiV3FilesystemTypeGet**](FileSystemApi.md#ApiV3FilesystemTypeGet) | **Get** /api/v3/filesystem/type | 



## ApiV3FilesystemGet

> ApiV3FilesystemGet(ctx).Path(path).IncludeFiles(includeFiles).AllowFoldersWithoutTrailingSlashes(allowFoldersWithoutTrailingSlashes).Execute()



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
    path := "path_example" // string |  (optional)
    includeFiles := true // bool |  (optional) (default to false)
    allowFoldersWithoutTrailingSlashes := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSystemApi.ApiV3FilesystemGet(context.Background()).Path(path).IncludeFiles(includeFiles).AllowFoldersWithoutTrailingSlashes(allowFoldersWithoutTrailingSlashes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemApi.ApiV3FilesystemGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3FilesystemGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 
 **includeFiles** | **bool** |  | [default to false]
 **allowFoldersWithoutTrailingSlashes** | **bool** |  | [default to false]

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


## ApiV3FilesystemMediafilesGet

> ApiV3FilesystemMediafilesGet(ctx).Path(path).Execute()



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
    path := "path_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSystemApi.ApiV3FilesystemMediafilesGet(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemApi.ApiV3FilesystemMediafilesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3FilesystemMediafilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 

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


## ApiV3FilesystemTypeGet

> ApiV3FilesystemTypeGet(ctx).Path(path).Execute()



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
    path := "path_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSystemApi.ApiV3FilesystemTypeGet(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSystemApi.ApiV3FilesystemTypeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3FilesystemTypeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 

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

