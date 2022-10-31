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
)


// ManualImportApiService ManualImportApi service
type ManualImportApiService service

type ApiApiV3ManualimportGetRequest struct {
	ctx context.Context
	ApiService *ManualImportApiService
	folder *string
	downloadId *string
	seriesId *int32
	seasonNumber *int32
	filterExistingFiles *bool
}

func (r ApiApiV3ManualimportGetRequest) Folder(folder string) ApiApiV3ManualimportGetRequest {
	r.folder = &folder
	return r
}

func (r ApiApiV3ManualimportGetRequest) DownloadId(downloadId string) ApiApiV3ManualimportGetRequest {
	r.downloadId = &downloadId
	return r
}

func (r ApiApiV3ManualimportGetRequest) SeriesId(seriesId int32) ApiApiV3ManualimportGetRequest {
	r.seriesId = &seriesId
	return r
}

func (r ApiApiV3ManualimportGetRequest) SeasonNumber(seasonNumber int32) ApiApiV3ManualimportGetRequest {
	r.seasonNumber = &seasonNumber
	return r
}

func (r ApiApiV3ManualimportGetRequest) FilterExistingFiles(filterExistingFiles bool) ApiApiV3ManualimportGetRequest {
	r.filterExistingFiles = &filterExistingFiles
	return r
}

func (r ApiApiV3ManualimportGetRequest) Execute() ([]ManualImportResource, *http.Response, error) {
	return r.ApiService.ApiV3ManualimportGetExecute(r)
}

/*
ApiV3ManualimportGet Method for ApiV3ManualimportGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV3ManualimportGetRequest
*/
func (a *ManualImportApiService) ApiV3ManualimportGet(ctx context.Context) ApiApiV3ManualimportGetRequest {
	return ApiApiV3ManualimportGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ManualImportResource
func (a *ManualImportApiService) ApiV3ManualimportGetExecute(r ApiApiV3ManualimportGetRequest) ([]ManualImportResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ManualImportResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualImportApiService.ApiV3ManualimportGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/manualimport"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.folder != nil {
		localVarQueryParams.Add("folder", parameterToString(*r.folder, ""))
	}
	if r.downloadId != nil {
		localVarQueryParams.Add("downloadId", parameterToString(*r.downloadId, ""))
	}
	if r.seriesId != nil {
		localVarQueryParams.Add("seriesId", parameterToString(*r.seriesId, ""))
	}
	if r.seasonNumber != nil {
		localVarQueryParams.Add("seasonNumber", parameterToString(*r.seasonNumber, ""))
	}
	if r.filterExistingFiles != nil {
		localVarQueryParams.Add("filterExistingFiles", parameterToString(*r.filterExistingFiles, ""))
	}
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

type ApiApiV3ManualimportPostRequest struct {
	ctx context.Context
	ApiService *ManualImportApiService
	manualImportReprocessResource *[]ManualImportReprocessResource
}

func (r ApiApiV3ManualimportPostRequest) ManualImportReprocessResource(manualImportReprocessResource []ManualImportReprocessResource) ApiApiV3ManualimportPostRequest {
	r.manualImportReprocessResource = &manualImportReprocessResource
	return r
}

func (r ApiApiV3ManualimportPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV3ManualimportPostExecute(r)
}

/*
ApiV3ManualimportPost Method for ApiV3ManualimportPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV3ManualimportPostRequest
*/
func (a *ManualImportApiService) ApiV3ManualimportPost(ctx context.Context) ApiApiV3ManualimportPostRequest {
	return ApiApiV3ManualimportPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ManualImportApiService) ApiV3ManualimportPostExecute(r ApiApiV3ManualimportPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualImportApiService.ApiV3ManualimportPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/manualimport"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.manualImportReprocessResource
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
