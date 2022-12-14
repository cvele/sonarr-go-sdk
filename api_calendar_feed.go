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


// CalendarFeedApiService CalendarFeedApi service
type CalendarFeedApiService service

type ApiFeedV3CalendarSonarrIcsGetRequest struct {
	ctx context.Context
	ApiService *CalendarFeedApiService
	pastDays *int32
	futureDays *int32
	tagList *string
	unmonitored *bool
	premieresOnly *bool
	asAllDay *bool
}

func (r ApiFeedV3CalendarSonarrIcsGetRequest) PastDays(pastDays int32) ApiFeedV3CalendarSonarrIcsGetRequest {
	r.pastDays = &pastDays
	return r
}

func (r ApiFeedV3CalendarSonarrIcsGetRequest) FutureDays(futureDays int32) ApiFeedV3CalendarSonarrIcsGetRequest {
	r.futureDays = &futureDays
	return r
}

func (r ApiFeedV3CalendarSonarrIcsGetRequest) TagList(tagList string) ApiFeedV3CalendarSonarrIcsGetRequest {
	r.tagList = &tagList
	return r
}

func (r ApiFeedV3CalendarSonarrIcsGetRequest) Unmonitored(unmonitored bool) ApiFeedV3CalendarSonarrIcsGetRequest {
	r.unmonitored = &unmonitored
	return r
}

func (r ApiFeedV3CalendarSonarrIcsGetRequest) PremieresOnly(premieresOnly bool) ApiFeedV3CalendarSonarrIcsGetRequest {
	r.premieresOnly = &premieresOnly
	return r
}

func (r ApiFeedV3CalendarSonarrIcsGetRequest) AsAllDay(asAllDay bool) ApiFeedV3CalendarSonarrIcsGetRequest {
	r.asAllDay = &asAllDay
	return r
}

func (r ApiFeedV3CalendarSonarrIcsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.FeedV3CalendarSonarrIcsGetExecute(r)
}

/*
FeedV3CalendarSonarrIcsGet Method for FeedV3CalendarSonarrIcsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFeedV3CalendarSonarrIcsGetRequest
*/
func (a *CalendarFeedApiService) FeedV3CalendarSonarrIcsGet(ctx context.Context) ApiFeedV3CalendarSonarrIcsGetRequest {
	return ApiFeedV3CalendarSonarrIcsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CalendarFeedApiService) FeedV3CalendarSonarrIcsGetExecute(r ApiFeedV3CalendarSonarrIcsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarFeedApiService.FeedV3CalendarSonarrIcsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feed/v3/calendar/sonarr.ics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pastDays != nil {
		localVarQueryParams.Add("pastDays", parameterToString(*r.pastDays, ""))
	}
	if r.futureDays != nil {
		localVarQueryParams.Add("futureDays", parameterToString(*r.futureDays, ""))
	}
	if r.tagList != nil {
		localVarQueryParams.Add("tagList", parameterToString(*r.tagList, ""))
	}
	if r.unmonitored != nil {
		localVarQueryParams.Add("unmonitored", parameterToString(*r.unmonitored, ""))
	}
	if r.premieresOnly != nil {
		localVarQueryParams.Add("premieresOnly", parameterToString(*r.premieresOnly, ""))
	}
	if r.asAllDay != nil {
		localVarQueryParams.Add("asAllDay", parameterToString(*r.asAllDay, ""))
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
