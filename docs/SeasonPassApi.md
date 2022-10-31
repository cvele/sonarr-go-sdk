# \SeasonPassApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3SeasonpassPost**](SeasonPassApi.md#ApiV3SeasonpassPost) | **Post** /api/v3/seasonpass | 



## ApiV3SeasonpassPost

> ApiV3SeasonpassPost(ctx).SeasonPassResource(seasonPassResource).Execute()



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
    seasonPassResource := *openapiclient.NewSeasonPassResource() // SeasonPassResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SeasonPassApi.ApiV3SeasonpassPost(context.Background()).SeasonPassResource(seasonPassResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SeasonPassApi.ApiV3SeasonpassPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3SeasonpassPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seasonPassResource** | [**SeasonPassResource**](SeasonPassResource.md) |  | 

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

