# \LanguageProfileSchemaApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3LanguageprofileSchemaGet**](LanguageProfileSchemaApi.md#ApiV3LanguageprofileSchemaGet) | **Get** /api/v3/languageprofile/schema | 



## ApiV3LanguageprofileSchemaGet

> LanguageProfileResource ApiV3LanguageprofileSchemaGet(ctx).Execute()



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
    resp, r, err := apiClient.LanguageProfileSchemaApi.ApiV3LanguageprofileSchemaGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageProfileSchemaApi.ApiV3LanguageprofileSchemaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3LanguageprofileSchemaGet`: LanguageProfileResource
    fmt.Fprintf(os.Stdout, "Response from `LanguageProfileSchemaApi.ApiV3LanguageprofileSchemaGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3LanguageprofileSchemaGetRequest struct via the builder pattern


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

