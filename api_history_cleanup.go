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


type HistoryCleanupAPI interface {

	/*
	CleanupAsync Clean up history (POST)

	Schedules asynchronous history cleanup (See
[History cleanup](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#history-cleanup)).

**Note:** This endpoint will return at most a single history cleanup job.
Since version `7.9.0` it is possible to configure multiple
[parallel history cleanup jobs](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#parallel-execution). Use
[`GET /history/cleanup/jobs`](https://docs.camunda.org/manual/7.21/reference/rest/history/history-cleanup/get-history-cleanup-jobs)
to find all the available history cleanup jobs.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCleanupAsyncRequest
	*/
	CleanupAsync(ctx context.Context) ApiCleanupAsyncRequest

	// CleanupAsyncExecute executes the request
	//  @return JobDto
	CleanupAsyncExecute(r ApiCleanupAsyncRequest) (*JobDto, *http.Response, error)

	/*
	FindCleanupJob Find clean up history job (GET)

	**Deprecated!** Use `GET /history/cleanup/jobs` instead.

Finds history cleanup job (See
[History cleanup](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#history-cleanup)).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFindCleanupJobRequest

	Deprecated
	*/
	FindCleanupJob(ctx context.Context) ApiFindCleanupJobRequest

	// FindCleanupJobExecute executes the request
	//  @return JobDto
	// Deprecated
	FindCleanupJobExecute(r ApiFindCleanupJobRequest) (*JobDto, *http.Response, error)

	/*
	FindCleanupJobs Find clean up history jobs (GET)

	Finds history cleanup jobs (See
[History cleanup](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#history-cleanup)).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFindCleanupJobsRequest
	*/
	FindCleanupJobs(ctx context.Context) ApiFindCleanupJobsRequest

	// FindCleanupJobsExecute executes the request
	//  @return []JobDto
	FindCleanupJobsExecute(r ApiFindCleanupJobsRequest) ([]JobDto, *http.Response, error)

	/*
	GetHistoryCleanupConfiguration Get History Cleanup Configuration

	Retrieves history cleanup batch window configuration (See
[History cleanup](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#history-cleanup)).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetHistoryCleanupConfigurationRequest
	*/
	GetHistoryCleanupConfiguration(ctx context.Context) ApiGetHistoryCleanupConfigurationRequest

	// GetHistoryCleanupConfigurationExecute executes the request
	//  @return HistoryCleanupConfigurationDto
	GetHistoryCleanupConfigurationExecute(r ApiGetHistoryCleanupConfigurationRequest) (*HistoryCleanupConfigurationDto, *http.Response, error)
}

// HistoryCleanupAPIService HistoryCleanupAPI service
type HistoryCleanupAPIService service

type ApiCleanupAsyncRequest struct {
	ctx context.Context
	ApiService HistoryCleanupAPI
	immediatelyDue *bool
}

// When true the job will be scheduled for nearest future. When &#x60;false&#x60;, the job will be scheduled for next batch window start time. Default is &#x60;true&#x60;.
func (r ApiCleanupAsyncRequest) ImmediatelyDue(immediatelyDue bool) ApiCleanupAsyncRequest {
	r.immediatelyDue = &immediatelyDue
	return r
}

func (r ApiCleanupAsyncRequest) Execute() (*JobDto, *http.Response, error) {
	return r.ApiService.CleanupAsyncExecute(r)
}

/*
CleanupAsync Clean up history (POST)

Schedules asynchronous history cleanup (See
[History cleanup](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#history-cleanup)).

**Note:** This endpoint will return at most a single history cleanup job.
Since version `7.9.0` it is possible to configure multiple
[parallel history cleanup jobs](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#parallel-execution). Use
[`GET /history/cleanup/jobs`](https://docs.camunda.org/manual/7.21/reference/rest/history/history-cleanup/get-history-cleanup-jobs)
to find all the available history cleanup jobs.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCleanupAsyncRequest
*/
func (a *HistoryCleanupAPIService) CleanupAsync(ctx context.Context) ApiCleanupAsyncRequest {
	return ApiCleanupAsyncRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JobDto
func (a *HistoryCleanupAPIService) CleanupAsyncExecute(r ApiCleanupAsyncRequest) (*JobDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JobDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryCleanupAPIService.CleanupAsync")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/history/cleanup"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.immediatelyDue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "immediatelyDue", r.immediatelyDue, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ExceptionDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiFindCleanupJobRequest struct {
	ctx context.Context
	ApiService HistoryCleanupAPI
}

func (r ApiFindCleanupJobRequest) Execute() (*JobDto, *http.Response, error) {
	return r.ApiService.FindCleanupJobExecute(r)
}

/*
FindCleanupJob Find clean up history job (GET)

**Deprecated!** Use `GET /history/cleanup/jobs` instead.

Finds history cleanup job (See
[History cleanup](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#history-cleanup)).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFindCleanupJobRequest

Deprecated
*/
func (a *HistoryCleanupAPIService) FindCleanupJob(ctx context.Context) ApiFindCleanupJobRequest {
	return ApiFindCleanupJobRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JobDto
// Deprecated
func (a *HistoryCleanupAPIService) FindCleanupJobExecute(r ApiFindCleanupJobRequest) (*JobDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JobDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryCleanupAPIService.FindCleanupJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/history/cleanup/job"

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ExceptionDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiFindCleanupJobsRequest struct {
	ctx context.Context
	ApiService HistoryCleanupAPI
}

func (r ApiFindCleanupJobsRequest) Execute() ([]JobDto, *http.Response, error) {
	return r.ApiService.FindCleanupJobsExecute(r)
}

/*
FindCleanupJobs Find clean up history jobs (GET)

Finds history cleanup jobs (See
[History cleanup](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#history-cleanup)).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFindCleanupJobsRequest
*/
func (a *HistoryCleanupAPIService) FindCleanupJobs(ctx context.Context) ApiFindCleanupJobsRequest {
	return ApiFindCleanupJobsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []JobDto
func (a *HistoryCleanupAPIService) FindCleanupJobsExecute(r ApiFindCleanupJobsRequest) ([]JobDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []JobDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryCleanupAPIService.FindCleanupJobs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/history/cleanup/jobs"

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ExceptionDto
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetHistoryCleanupConfigurationRequest struct {
	ctx context.Context
	ApiService HistoryCleanupAPI
}

func (r ApiGetHistoryCleanupConfigurationRequest) Execute() (*HistoryCleanupConfigurationDto, *http.Response, error) {
	return r.ApiService.GetHistoryCleanupConfigurationExecute(r)
}

/*
GetHistoryCleanupConfiguration Get History Cleanup Configuration

Retrieves history cleanup batch window configuration (See
[History cleanup](https://docs.camunda.org/manual/7.21/user-guide/process-engine/history/#history-cleanup)).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetHistoryCleanupConfigurationRequest
*/
func (a *HistoryCleanupAPIService) GetHistoryCleanupConfiguration(ctx context.Context) ApiGetHistoryCleanupConfigurationRequest {
	return ApiGetHistoryCleanupConfigurationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistoryCleanupConfigurationDto
func (a *HistoryCleanupAPIService) GetHistoryCleanupConfigurationExecute(r ApiGetHistoryCleanupConfigurationRequest) (*HistoryCleanupConfigurationDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistoryCleanupConfigurationDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryCleanupAPIService.GetHistoryCleanupConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/history/cleanup/configuration"

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
