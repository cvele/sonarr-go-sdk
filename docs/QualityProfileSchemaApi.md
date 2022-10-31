# \QualityProfileSchemaApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3QualityprofileSchemaGet**](QualityProfileSchemaApi.md#ApiV3QualityprofileSchemaGet) | **Get** /api/v3/qualityprofile/schema | 



## ApiV3QualityprofileSchemaGet

> QualityProfileResource ApiV3QualityprofileSchemaGet(ctx).Execute()



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
    resp, r, err := apiClient.QualityProfileSchemaApi.ApiV3QualityprofileSchemaGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityProfileSchemaApi.ApiV3QualityprofileSchemaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3QualityprofileSchemaGet`: QualityProfileResource
    fmt.Fprintf(os.Stdout, "Response from `QualityProfileSchemaApi.ApiV3QualityprofileSchemaGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3QualityprofileSchemaGetRequest struct via the builder pattern


### Return type

[**QualityProfileResource**](QualityProfileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

