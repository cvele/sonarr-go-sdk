# \NamingConfigApi

All URIs are relative to *http://localhost:8989*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ConfigNamingExamplesGet**](NamingConfigApi.md#ApiV3ConfigNamingExamplesGet) | **Get** /api/v3/config/naming/examples | 
[**ApiV3ConfigNamingGet**](NamingConfigApi.md#ApiV3ConfigNamingGet) | **Get** /api/v3/config/naming | 
[**ApiV3ConfigNamingIdGet**](NamingConfigApi.md#ApiV3ConfigNamingIdGet) | **Get** /api/v3/config/naming/{id} | 
[**ApiV3ConfigNamingIdPut**](NamingConfigApi.md#ApiV3ConfigNamingIdPut) | **Put** /api/v3/config/naming/{id} | 



## ApiV3ConfigNamingExamplesGet

> ApiV3ConfigNamingExamplesGet(ctx).RenameEpisodes(renameEpisodes).ReplaceIllegalCharacters(replaceIllegalCharacters).MultiEpisodeStyle(multiEpisodeStyle).StandardEpisodeFormat(standardEpisodeFormat).DailyEpisodeFormat(dailyEpisodeFormat).AnimeEpisodeFormat(animeEpisodeFormat).SeriesFolderFormat(seriesFolderFormat).SeasonFolderFormat(seasonFolderFormat).SpecialsFolderFormat(specialsFolderFormat).IncludeSeriesTitle(includeSeriesTitle).IncludeEpisodeTitle(includeEpisodeTitle).IncludeQuality(includeQuality).ReplaceSpaces(replaceSpaces).Separator(separator).NumberStyle(numberStyle).Id(id).ResourceName(resourceName).Execute()



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
    renameEpisodes := true // bool |  (optional)
    replaceIllegalCharacters := true // bool |  (optional)
    multiEpisodeStyle := int32(56) // int32 |  (optional)
    standardEpisodeFormat := "standardEpisodeFormat_example" // string |  (optional)
    dailyEpisodeFormat := "dailyEpisodeFormat_example" // string |  (optional)
    animeEpisodeFormat := "animeEpisodeFormat_example" // string |  (optional)
    seriesFolderFormat := "seriesFolderFormat_example" // string |  (optional)
    seasonFolderFormat := "seasonFolderFormat_example" // string |  (optional)
    specialsFolderFormat := "specialsFolderFormat_example" // string |  (optional)
    includeSeriesTitle := true // bool |  (optional)
    includeEpisodeTitle := true // bool |  (optional)
    includeQuality := true // bool |  (optional)
    replaceSpaces := true // bool |  (optional)
    separator := "separator_example" // string |  (optional)
    numberStyle := "numberStyle_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)
    resourceName := "resourceName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamingConfigApi.ApiV3ConfigNamingExamplesGet(context.Background()).RenameEpisodes(renameEpisodes).ReplaceIllegalCharacters(replaceIllegalCharacters).MultiEpisodeStyle(multiEpisodeStyle).StandardEpisodeFormat(standardEpisodeFormat).DailyEpisodeFormat(dailyEpisodeFormat).AnimeEpisodeFormat(animeEpisodeFormat).SeriesFolderFormat(seriesFolderFormat).SeasonFolderFormat(seasonFolderFormat).SpecialsFolderFormat(specialsFolderFormat).IncludeSeriesTitle(includeSeriesTitle).IncludeEpisodeTitle(includeEpisodeTitle).IncludeQuality(includeQuality).ReplaceSpaces(replaceSpaces).Separator(separator).NumberStyle(numberStyle).Id(id).ResourceName(resourceName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigApi.ApiV3ConfigNamingExamplesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigNamingExamplesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renameEpisodes** | **bool** |  | 
 **replaceIllegalCharacters** | **bool** |  | 
 **multiEpisodeStyle** | **int32** |  | 
 **standardEpisodeFormat** | **string** |  | 
 **dailyEpisodeFormat** | **string** |  | 
 **animeEpisodeFormat** | **string** |  | 
 **seriesFolderFormat** | **string** |  | 
 **seasonFolderFormat** | **string** |  | 
 **specialsFolderFormat** | **string** |  | 
 **includeSeriesTitle** | **bool** |  | 
 **includeEpisodeTitle** | **bool** |  | 
 **includeQuality** | **bool** |  | 
 **replaceSpaces** | **bool** |  | 
 **separator** | **string** |  | 
 **numberStyle** | **string** |  | 
 **id** | **int32** |  | 
 **resourceName** | **string** |  | 

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


## ApiV3ConfigNamingGet

> NamingConfigResource ApiV3ConfigNamingGet(ctx).Execute()



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
    resp, r, err := apiClient.NamingConfigApi.ApiV3ConfigNamingGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigApi.ApiV3ConfigNamingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigNamingGet`: NamingConfigResource
    fmt.Fprintf(os.Stdout, "Response from `NamingConfigApi.ApiV3ConfigNamingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigNamingGetRequest struct via the builder pattern


### Return type

[**NamingConfigResource**](NamingConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ConfigNamingIdGet

> NamingConfigResource ApiV3ConfigNamingIdGet(ctx, id).Execute()



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
    resp, r, err := apiClient.NamingConfigApi.ApiV3ConfigNamingIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigApi.ApiV3ConfigNamingIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigNamingIdGet`: NamingConfigResource
    fmt.Fprintf(os.Stdout, "Response from `NamingConfigApi.ApiV3ConfigNamingIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigNamingIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NamingConfigResource**](NamingConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV3ConfigNamingIdPut

> NamingConfigResource ApiV3ConfigNamingIdPut(ctx, id).NamingConfigResource(namingConfigResource).Execute()



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
    namingConfigResource := *openapiclient.NewNamingConfigResource() // NamingConfigResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamingConfigApi.ApiV3ConfigNamingIdPut(context.Background(), id).NamingConfigResource(namingConfigResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamingConfigApi.ApiV3ConfigNamingIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV3ConfigNamingIdPut`: NamingConfigResource
    fmt.Fprintf(os.Stdout, "Response from `NamingConfigApi.ApiV3ConfigNamingIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ConfigNamingIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namingConfigResource** | [**NamingConfigResource**](NamingConfigResource.md) |  | 

### Return type

[**NamingConfigResource**](NamingConfigResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

