# \ParseApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ParseGet**](ParseApi.md#ApiV3ParseGet) | **Get** /api/v3/parse | 



## ApiV3ParseGet

> ParseResource ApiV3ParseGet(ctx).Title(title).Path(path).Execute()



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
    title := "title_example" // string |  (optional)
    path := "path_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParseApi.ApiV3ParseGet(context.Background()).Title(title).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParseApi.ApiV3ParseGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ParseGet`: ParseResource
    fmt.Fprintf(os.Stdout, "Response from `ParseApi.ApiV3ParseGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ParseGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** |  | 
 **path** | **string** |  | 

### Return type

[**ParseResource**](ParseResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

