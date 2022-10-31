# \DiskSpaceApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3DiskspaceGet**](DiskSpaceApi.md#ApiV3DiskspaceGet) | **Get** /api/v3/diskspace | 



## ApiV3DiskspaceGet

> []DiskSpaceResource ApiV3DiskspaceGet(ctx).Execute()



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
    resp, r, err := apiClient.DiskSpaceApi.ApiV3DiskspaceGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiskSpaceApi.ApiV3DiskspaceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3DiskspaceGet`: []DiskSpaceResource
    fmt.Fprintf(os.Stdout, "Response from `DiskSpaceApi.ApiV3DiskspaceGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3DiskspaceGetRequest struct via the builder pattern


### Return type

[**[]DiskSpaceResource**](DiskSpaceResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

