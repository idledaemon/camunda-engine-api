/*
Camunda Platform REST API

OpenApi Spec for Camunda Platform REST API.

API version: 7.21.2-ee
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package camundarestgo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type SchemaLogAPI interface {

	/*
	GetSchemaLog Get List

	Queries for schema log entries that fulfill given parameters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSchemaLogRequest
	*/
	GetSchemaLog(ctx context.Context) ApiGetSchemaLogRequest

	// GetSchemaLogExecute executes the request
	//  @return []SchemaLogEntryDto
	GetSchemaLogExecute(r ApiGetSchemaLogRequest) ([]SchemaLogEntryDto, *http.Response, error)

	/*
	QuerySchemaLog Get List (POST)

	Queries for schema log entries that fulfill given parameters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiQuerySchemaLogRequest
	*/
	QuerySchemaLog(ctx context.Context) ApiQuerySchemaLogRequest

	// QuerySchemaLogExecute executes the request
	//  @return []SchemaLogEntryDto
	QuerySchemaLogExecute(r ApiQuerySchemaLogRequest) ([]SchemaLogEntryDto, *http.Response, error)
}

// SchemaLogAPIService SchemaLogAPI service
type SchemaLogAPIService service

type ApiGetSchemaLogRequest struct {
	ctx context.Context
	ApiService SchemaLogAPI
	version *string
	sortBy *string
	sortOrder *string
	firstResult *int32
	maxResults *int32
}

// Only return schema log entries with a specific version.
func (r ApiGetSchemaLogRequest) Version(version string) ApiGetSchemaLogRequest {
	r.version = &version
	return r
}

// Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter.
func (r ApiGetSchemaLogRequest) SortBy(sortBy string) ApiGetSchemaLogRequest {
	r.sortBy = &sortBy
	return r
}

// Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter.
func (r ApiGetSchemaLogRequest) SortOrder(sortOrder string) ApiGetSchemaLogRequest {
	r.sortOrder = &sortOrder
	return r
}

// Pagination of results. Specifies the index of the first result to return.
func (r ApiGetSchemaLogRequest) FirstResult(firstResult int32) ApiGetSchemaLogRequest {
	r.firstResult = &firstResult
	return r
}

// Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left.
func (r ApiGetSchemaLogRequest) MaxResults(maxResults int32) ApiGetSchemaLogRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiGetSchemaLogRequest) Execute() ([]SchemaLogEntryDto, *http.Response, error) {
	return r.ApiService.GetSchemaLogExecute(r)
}

/*
GetSchemaLog Get List

Queries for schema log entries that fulfill given parameters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSchemaLogRequest
*/
func (a *SchemaLogAPIService) GetSchemaLog(ctx context.Context) ApiGetSchemaLogRequest {
	return ApiGetSchemaLogRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SchemaLogEntryDto
func (a *SchemaLogAPIService) GetSchemaLogExecute(r ApiGetSchemaLogRequest) ([]SchemaLogEntryDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SchemaLogEntryDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SchemaLogAPIService.GetSchemaLog")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/schema/log"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", r.sortBy, "")
	}
	if r.sortOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortOrder", r.sortOrder, "")
	}
	if r.firstResult != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "firstResult", r.firstResult, "")
	}
	if r.maxResults != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxResults", r.maxResults, "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiQuerySchemaLogRequest struct {
	ctx context.Context
	ApiService SchemaLogAPI
	firstResult *int32
	maxResults *int32
	schemaLogQueryDto *SchemaLogQueryDto
}

// Pagination of results. Specifies the index of the first result to return.
func (r ApiQuerySchemaLogRequest) FirstResult(firstResult int32) ApiQuerySchemaLogRequest {
	r.firstResult = &firstResult
	return r
}

// Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left.
func (r ApiQuerySchemaLogRequest) MaxResults(maxResults int32) ApiQuerySchemaLogRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiQuerySchemaLogRequest) SchemaLogQueryDto(schemaLogQueryDto SchemaLogQueryDto) ApiQuerySchemaLogRequest {
	r.schemaLogQueryDto = &schemaLogQueryDto
	return r
}

func (r ApiQuerySchemaLogRequest) Execute() ([]SchemaLogEntryDto, *http.Response, error) {
	return r.ApiService.QuerySchemaLogExecute(r)
}

/*
QuerySchemaLog Get List (POST)

Queries for schema log entries that fulfill given parameters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQuerySchemaLogRequest
*/
func (a *SchemaLogAPIService) QuerySchemaLog(ctx context.Context) ApiQuerySchemaLogRequest {
	return ApiQuerySchemaLogRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SchemaLogEntryDto
func (a *SchemaLogAPIService) QuerySchemaLogExecute(r ApiQuerySchemaLogRequest) ([]SchemaLogEntryDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SchemaLogEntryDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SchemaLogAPIService.QuerySchemaLog")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/schema/log"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.firstResult != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "firstResult", r.firstResult, "")
	}
	if r.maxResults != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxResults", r.maxResults, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.schemaLogQueryDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
