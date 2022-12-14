/*
Sonarr

Sonarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// NamingConfigApiService NamingConfigApi service
type NamingConfigApiService service

type ApiApiV3ConfigNamingExamplesGetRequest struct {
	ctx context.Context
	ApiService *NamingConfigApiService
	renameEpisodes *bool
	replaceIllegalCharacters *bool
	multiEpisodeStyle *int32
	standardEpisodeFormat *string
	dailyEpisodeFormat *string
	animeEpisodeFormat *string
	seriesFolderFormat *string
	seasonFolderFormat *string
	specialsFolderFormat *string
	includeSeriesTitle *bool
	includeEpisodeTitle *bool
	includeQuality *bool
	replaceSpaces *bool
	separator *string
	numberStyle *string
	id *int32
	resourceName *string
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) RenameEpisodes(renameEpisodes bool) ApiApiV3ConfigNamingExamplesGetRequest {
	r.renameEpisodes = &renameEpisodes
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) ReplaceIllegalCharacters(replaceIllegalCharacters bool) ApiApiV3ConfigNamingExamplesGetRequest {
	r.replaceIllegalCharacters = &replaceIllegalCharacters
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) MultiEpisodeStyle(multiEpisodeStyle int32) ApiApiV3ConfigNamingExamplesGetRequest {
	r.multiEpisodeStyle = &multiEpisodeStyle
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) StandardEpisodeFormat(standardEpisodeFormat string) ApiApiV3ConfigNamingExamplesGetRequest {
	r.standardEpisodeFormat = &standardEpisodeFormat
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) DailyEpisodeFormat(dailyEpisodeFormat string) ApiApiV3ConfigNamingExamplesGetRequest {
	r.dailyEpisodeFormat = &dailyEpisodeFormat
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) AnimeEpisodeFormat(animeEpisodeFormat string) ApiApiV3ConfigNamingExamplesGetRequest {
	r.animeEpisodeFormat = &animeEpisodeFormat
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) SeriesFolderFormat(seriesFolderFormat string) ApiApiV3ConfigNamingExamplesGetRequest {
	r.seriesFolderFormat = &seriesFolderFormat
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) SeasonFolderFormat(seasonFolderFormat string) ApiApiV3ConfigNamingExamplesGetRequest {
	r.seasonFolderFormat = &seasonFolderFormat
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) SpecialsFolderFormat(specialsFolderFormat string) ApiApiV3ConfigNamingExamplesGetRequest {
	r.specialsFolderFormat = &specialsFolderFormat
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) IncludeSeriesTitle(includeSeriesTitle bool) ApiApiV3ConfigNamingExamplesGetRequest {
	r.includeSeriesTitle = &includeSeriesTitle
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) IncludeEpisodeTitle(includeEpisodeTitle bool) ApiApiV3ConfigNamingExamplesGetRequest {
	r.includeEpisodeTitle = &includeEpisodeTitle
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) IncludeQuality(includeQuality bool) ApiApiV3ConfigNamingExamplesGetRequest {
	r.includeQuality = &includeQuality
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) ReplaceSpaces(replaceSpaces bool) ApiApiV3ConfigNamingExamplesGetRequest {
	r.replaceSpaces = &replaceSpaces
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) Separator(separator string) ApiApiV3ConfigNamingExamplesGetRequest {
	r.separator = &separator
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) NumberStyle(numberStyle string) ApiApiV3ConfigNamingExamplesGetRequest {
	r.numberStyle = &numberStyle
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) Id(id int32) ApiApiV3ConfigNamingExamplesGetRequest {
	r.id = &id
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) ResourceName(resourceName string) ApiApiV3ConfigNamingExamplesGetRequest {
	r.resourceName = &resourceName
	return r
}

func (r ApiApiV3ConfigNamingExamplesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3ConfigNamingExamplesGetExecute(r)
}

/*
ApiV3ConfigNamingExamplesGet Method for ApiV3ConfigNamingExamplesGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV3ConfigNamingExamplesGetRequest
*/
func (a *NamingConfigApiService) ApiV3ConfigNamingExamplesGet(ctx context.Context) ApiApiV3ConfigNamingExamplesGetRequest {
	return ApiApiV3ConfigNamingExamplesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NamingConfigApiService) ApiV3ConfigNamingExamplesGetExecute(r ApiApiV3ConfigNamingExamplesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigApiService.ApiV3ConfigNamingExamplesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming/examples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.renameEpisodes != nil {
		localVarQueryParams.Add("RenameEpisodes", parameterToString(*r.renameEpisodes, ""))
	}
	if r.replaceIllegalCharacters != nil {
		localVarQueryParams.Add("ReplaceIllegalCharacters", parameterToString(*r.replaceIllegalCharacters, ""))
	}
	if r.multiEpisodeStyle != nil {
		localVarQueryParams.Add("MultiEpisodeStyle", parameterToString(*r.multiEpisodeStyle, ""))
	}
	if r.standardEpisodeFormat != nil {
		localVarQueryParams.Add("StandardEpisodeFormat", parameterToString(*r.standardEpisodeFormat, ""))
	}
	if r.dailyEpisodeFormat != nil {
		localVarQueryParams.Add("DailyEpisodeFormat", parameterToString(*r.dailyEpisodeFormat, ""))
	}
	if r.animeEpisodeFormat != nil {
		localVarQueryParams.Add("AnimeEpisodeFormat", parameterToString(*r.animeEpisodeFormat, ""))
	}
	if r.seriesFolderFormat != nil {
		localVarQueryParams.Add("SeriesFolderFormat", parameterToString(*r.seriesFolderFormat, ""))
	}
	if r.seasonFolderFormat != nil {
		localVarQueryParams.Add("SeasonFolderFormat", parameterToString(*r.seasonFolderFormat, ""))
	}
	if r.specialsFolderFormat != nil {
		localVarQueryParams.Add("SpecialsFolderFormat", parameterToString(*r.specialsFolderFormat, ""))
	}
	if r.includeSeriesTitle != nil {
		localVarQueryParams.Add("IncludeSeriesTitle", parameterToString(*r.includeSeriesTitle, ""))
	}
	if r.includeEpisodeTitle != nil {
		localVarQueryParams.Add("IncludeEpisodeTitle", parameterToString(*r.includeEpisodeTitle, ""))
	}
	if r.includeQuality != nil {
		localVarQueryParams.Add("IncludeQuality", parameterToString(*r.includeQuality, ""))
	}
	if r.replaceSpaces != nil {
		localVarQueryParams.Add("ReplaceSpaces", parameterToString(*r.replaceSpaces, ""))
	}
	if r.separator != nil {
		localVarQueryParams.Add("Separator", parameterToString(*r.separator, ""))
	}
	if r.numberStyle != nil {
		localVarQueryParams.Add("NumberStyle", parameterToString(*r.numberStyle, ""))
	}
	if r.id != nil {
		localVarQueryParams.Add("Id", parameterToString(*r.id, ""))
	}
	if r.resourceName != nil {
		localVarQueryParams.Add("ResourceName", parameterToString(*r.resourceName, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV3ConfigNamingGetRequest struct {
	ctx context.Context
	ApiService *NamingConfigApiService
}

func (r ApiApiV3ConfigNamingGetRequest) Execute() (*NamingConfigResource, *http.Response, error) {
	return r.ApiService.ApiV3ConfigNamingGetExecute(r)
}

/*
ApiV3ConfigNamingGet Method for ApiV3ConfigNamingGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV3ConfigNamingGetRequest
*/
func (a *NamingConfigApiService) ApiV3ConfigNamingGet(ctx context.Context) ApiApiV3ConfigNamingGetRequest {
	return ApiApiV3ConfigNamingGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NamingConfigResource
func (a *NamingConfigApiService) ApiV3ConfigNamingGetExecute(r ApiApiV3ConfigNamingGetRequest) (*NamingConfigResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamingConfigResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigApiService.ApiV3ConfigNamingGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV3ConfigNamingIdGetRequest struct {
	ctx context.Context
	ApiService *NamingConfigApiService
	id int32
}

func (r ApiApiV3ConfigNamingIdGetRequest) Execute() (*NamingConfigResource, *http.Response, error) {
	return r.ApiService.ApiV3ConfigNamingIdGetExecute(r)
}

/*
ApiV3ConfigNamingIdGet Method for ApiV3ConfigNamingIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiApiV3ConfigNamingIdGetRequest
*/
func (a *NamingConfigApiService) ApiV3ConfigNamingIdGet(ctx context.Context, id int32) ApiApiV3ConfigNamingIdGetRequest {
	return ApiApiV3ConfigNamingIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NamingConfigResource
func (a *NamingConfigApiService) ApiV3ConfigNamingIdGetExecute(r ApiApiV3ConfigNamingIdGetRequest) (*NamingConfigResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamingConfigResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigApiService.ApiV3ConfigNamingIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV3ConfigNamingIdPutRequest struct {
	ctx context.Context
	ApiService *NamingConfigApiService
	id string
	namingConfigResource *NamingConfigResource
}

func (r ApiApiV3ConfigNamingIdPutRequest) NamingConfigResource(namingConfigResource NamingConfigResource) ApiApiV3ConfigNamingIdPutRequest {
	r.namingConfigResource = &namingConfigResource
	return r
}

func (r ApiApiV3ConfigNamingIdPutRequest) Execute() (*NamingConfigResource, *http.Response, error) {
	return r.ApiService.ApiV3ConfigNamingIdPutExecute(r)
}

/*
ApiV3ConfigNamingIdPut Method for ApiV3ConfigNamingIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiApiV3ConfigNamingIdPutRequest
*/
func (a *NamingConfigApiService) ApiV3ConfigNamingIdPut(ctx context.Context, id string) ApiApiV3ConfigNamingIdPutRequest {
	return ApiApiV3ConfigNamingIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NamingConfigResource
func (a *NamingConfigApiService) ApiV3ConfigNamingIdPutExecute(r ApiApiV3ConfigNamingIdPutRequest) (*NamingConfigResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NamingConfigResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamingConfigApiService.ApiV3ConfigNamingIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/config/naming/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.namingConfigResource
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
