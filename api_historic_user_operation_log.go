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
	"strings"
	"time"
)


type HistoricUserOperationLogAPI interface {

	/*
	ClearAnnotationUserOperationLog Clear Annotation of an User Operation Log (Historic)

	Clear the annotation which was previously set for auditing reasons.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param operationId The operation id of the operation log to be updated.
	@return ApiClearAnnotationUserOperationLogRequest
	*/
	ClearAnnotationUserOperationLog(ctx context.Context, operationId string) ApiClearAnnotationUserOperationLogRequest

	// ClearAnnotationUserOperationLogExecute executes the request
	ClearAnnotationUserOperationLogExecute(r ApiClearAnnotationUserOperationLogRequest) (*http.Response, error)

	/*
	QueryUserOperationCount Get User Operation Log Count

	Queries for the number of user operation log entries that fulfill the given parameters.
Takes the same parameters as the
[Get User Operation Log (Historic)](https://docs.camunda.org/manual/7.21/reference/rest/history/user-operation-log/get-user-operation-log-query/)
method.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiQueryUserOperationCountRequest
	*/
	QueryUserOperationCount(ctx context.Context) ApiQueryUserOperationCountRequest

	// QueryUserOperationCountExecute executes the request
	//  @return CountResultDto
	QueryUserOperationCountExecute(r ApiQueryUserOperationCountRequest) (*CountResultDto, *http.Response, error)

	/*
	QueryUserOperationEntries Get User Operation Log (Historic)

	Queries for user operation log entries that fulfill the given parameters.
The size of the result set can be retrieved by using the
[Get User Operation Log Count](https://docs.camunda.org/manual/7.21/reference/rest/history/user-operation-log/get-user-operation-log-query-count/)
method.

Note that the properties of operation log entries are interpreted as
restrictions on the entities they apply to. That means, if a single
process instance is updated, the field `processInstanceId` is
populated. If a single operation updates all process instances of the
same process definition, the field `processInstanceId` is `null` (a
`null` restriction is viewed as a wildcard, i.e., matches a process
instance with any id) and the field `processDefinitionId` is
populated. This way, which entities were changed by a user operation
can easily be reconstructed.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiQueryUserOperationEntriesRequest
	*/
	QueryUserOperationEntries(ctx context.Context) ApiQueryUserOperationEntriesRequest

	// QueryUserOperationEntriesExecute executes the request
	//  @return []UserOperationLogEntryDto
	QueryUserOperationEntriesExecute(r ApiQueryUserOperationEntriesRequest) ([]UserOperationLogEntryDto, *http.Response, error)

	/*
	SetAnnotationUserOperationLog Set Annotation to an User Operation Log (Historic)

	Set an annotation for auditing reasons.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param operationId The operation id of the operation log to be updated.
	@return ApiSetAnnotationUserOperationLogRequest
	*/
	SetAnnotationUserOperationLog(ctx context.Context, operationId string) ApiSetAnnotationUserOperationLogRequest

	// SetAnnotationUserOperationLogExecute executes the request
	SetAnnotationUserOperationLogExecute(r ApiSetAnnotationUserOperationLogRequest) (*http.Response, error)
}

// HistoricUserOperationLogAPIService HistoricUserOperationLogAPI service
type HistoricUserOperationLogAPIService service

type ApiClearAnnotationUserOperationLogRequest struct {
	ctx context.Context
	ApiService HistoricUserOperationLogAPI
	operationId string
}

func (r ApiClearAnnotationUserOperationLogRequest) Execute() (*http.Response, error) {
	return r.ApiService.ClearAnnotationUserOperationLogExecute(r)
}

/*
ClearAnnotationUserOperationLog Clear Annotation of an User Operation Log (Historic)

Clear the annotation which was previously set for auditing reasons.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param operationId The operation id of the operation log to be updated.
 @return ApiClearAnnotationUserOperationLogRequest
*/
func (a *HistoricUserOperationLogAPIService) ClearAnnotationUserOperationLog(ctx context.Context, operationId string) ApiClearAnnotationUserOperationLogRequest {
	return ApiClearAnnotationUserOperationLogRequest{
		ApiService: a,
		ctx: ctx,
		operationId: operationId,
	}
}

// Execute executes the request
func (a *HistoricUserOperationLogAPIService) ClearAnnotationUserOperationLogExecute(r ApiClearAnnotationUserOperationLogRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricUserOperationLogAPIService.ClearAnnotationUserOperationLog")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/history/user-operation/{operationId}/clear-annotation"
	localVarPath = strings.Replace(localVarPath, "{"+"operationId"+"}", url.PathEscape(parameterValueToString(r.operationId, "operationId")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiQueryUserOperationCountRequest struct {
	ctx context.Context
	ApiService HistoricUserOperationLogAPI
	deploymentId *string
	processDefinitionId *string
	processDefinitionKey *string
	processInstanceId *string
	executionId *string
	caseDefinitionId *string
	caseInstanceId *string
	caseExecutionId *string
	taskId *string
	externalTaskId *string
	batchId *string
	jobId *string
	jobDefinitionId *string
	userId *string
	operationId *string
	operationType *string
	entityType *string
	entityTypeIn *string
	category *string
	categoryIn *string
	property *string
	afterTimestamp *time.Time
	beforeTimestamp *time.Time
}

// Filter by deployment id.
func (r ApiQueryUserOperationCountRequest) DeploymentId(deploymentId string) ApiQueryUserOperationCountRequest {
	r.deploymentId = &deploymentId
	return r
}

// Filter by process definition id.
func (r ApiQueryUserOperationCountRequest) ProcessDefinitionId(processDefinitionId string) ApiQueryUserOperationCountRequest {
	r.processDefinitionId = &processDefinitionId
	return r
}

// Filter by process definition key.
func (r ApiQueryUserOperationCountRequest) ProcessDefinitionKey(processDefinitionKey string) ApiQueryUserOperationCountRequest {
	r.processDefinitionKey = &processDefinitionKey
	return r
}

// Filter by process instance id.
func (r ApiQueryUserOperationCountRequest) ProcessInstanceId(processInstanceId string) ApiQueryUserOperationCountRequest {
	r.processInstanceId = &processInstanceId
	return r
}

// Filter by execution id.
func (r ApiQueryUserOperationCountRequest) ExecutionId(executionId string) ApiQueryUserOperationCountRequest {
	r.executionId = &executionId
	return r
}

// Filter by case definition id.
func (r ApiQueryUserOperationCountRequest) CaseDefinitionId(caseDefinitionId string) ApiQueryUserOperationCountRequest {
	r.caseDefinitionId = &caseDefinitionId
	return r
}

// Filter by case instance id.
func (r ApiQueryUserOperationCountRequest) CaseInstanceId(caseInstanceId string) ApiQueryUserOperationCountRequest {
	r.caseInstanceId = &caseInstanceId
	return r
}

// Filter by case execution id.
func (r ApiQueryUserOperationCountRequest) CaseExecutionId(caseExecutionId string) ApiQueryUserOperationCountRequest {
	r.caseExecutionId = &caseExecutionId
	return r
}

// Only include operations on this task.
func (r ApiQueryUserOperationCountRequest) TaskId(taskId string) ApiQueryUserOperationCountRequest {
	r.taskId = &taskId
	return r
}

// Only include operations on this external task.
func (r ApiQueryUserOperationCountRequest) ExternalTaskId(externalTaskId string) ApiQueryUserOperationCountRequest {
	r.externalTaskId = &externalTaskId
	return r
}

// Only include operations on this batch.
func (r ApiQueryUserOperationCountRequest) BatchId(batchId string) ApiQueryUserOperationCountRequest {
	r.batchId = &batchId
	return r
}

// Filter by job id.
func (r ApiQueryUserOperationCountRequest) JobId(jobId string) ApiQueryUserOperationCountRequest {
	r.jobId = &jobId
	return r
}

// Filter by job definition id.
func (r ApiQueryUserOperationCountRequest) JobDefinitionId(jobDefinitionId string) ApiQueryUserOperationCountRequest {
	r.jobDefinitionId = &jobDefinitionId
	return r
}

// Only include operations of this user.
func (r ApiQueryUserOperationCountRequest) UserId(userId string) ApiQueryUserOperationCountRequest {
	r.userId = &userId
	return r
}

// Filter by the id of the operation. This allows fetching of multiple entries which are part of a composite operation.
func (r ApiQueryUserOperationCountRequest) OperationId(operationId string) ApiQueryUserOperationCountRequest {
	r.operationId = &operationId
	return r
}

// Filter by the type of the operation like &#x60;Claim&#x60; or &#x60;Delegate&#x60;. See the [Javadoc](https://docs.camunda.org/manual/7.21/reference/javadoc/?org/camunda/bpm/engine/history/UserOperationLogEntry.html) for a list of available operation types.
func (r ApiQueryUserOperationCountRequest) OperationType(operationType string) ApiQueryUserOperationCountRequest {
	r.operationType = &operationType
	return r
}

// Filter by the type of the entity that was affected by this operation, possible values are &#x60;Task&#x60;, &#x60;Attachment&#x60; or &#x60;IdentityLink&#x60;.
func (r ApiQueryUserOperationCountRequest) EntityType(entityType string) ApiQueryUserOperationCountRequest {
	r.entityType = &entityType
	return r
}

// Filter by a comma-separated list of types of the entities that was affected by this operation, possible values are &#x60;Task&#x60;, &#x60;Attachment&#x60; or &#x60;IdentityLink&#x60;.
func (r ApiQueryUserOperationCountRequest) EntityTypeIn(entityTypeIn string) ApiQueryUserOperationCountRequest {
	r.entityTypeIn = &entityTypeIn
	return r
}

// Filter by the category that this operation is associated with, possible values are &#x60;TaskWorker&#x60;, &#x60;Admin&#x60; or &#x60;Operator&#x60;.
func (r ApiQueryUserOperationCountRequest) Category(category string) ApiQueryUserOperationCountRequest {
	r.category = &category
	return r
}

// Filter by a comma-separated list of categories that this operation is associated with, possible values are &#x60;TaskWorker&#x60;, &#x60;Admin&#x60; or &#x60;Operator&#x60;.
func (r ApiQueryUserOperationCountRequest) CategoryIn(categoryIn string) ApiQueryUserOperationCountRequest {
	r.categoryIn = &categoryIn
	return r
}

// Only include operations that changed this property, e.g., &#x60;owner&#x60; or &#x60;assignee&#x60;.
func (r ApiQueryUserOperationCountRequest) Property(property string) ApiQueryUserOperationCountRequest {
	r.property = &property
	return r
}

// Restrict to entries that were created after the given timestamp. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the timestamp must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200.
func (r ApiQueryUserOperationCountRequest) AfterTimestamp(afterTimestamp time.Time) ApiQueryUserOperationCountRequest {
	r.afterTimestamp = &afterTimestamp
	return r
}

// Restrict to entries that were created before the given timestamp. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the timestamp must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200.
func (r ApiQueryUserOperationCountRequest) BeforeTimestamp(beforeTimestamp time.Time) ApiQueryUserOperationCountRequest {
	r.beforeTimestamp = &beforeTimestamp
	return r
}

func (r ApiQueryUserOperationCountRequest) Execute() (*CountResultDto, *http.Response, error) {
	return r.ApiService.QueryUserOperationCountExecute(r)
}

/*
QueryUserOperationCount Get User Operation Log Count

Queries for the number of user operation log entries that fulfill the given parameters.
Takes the same parameters as the
[Get User Operation Log (Historic)](https://docs.camunda.org/manual/7.21/reference/rest/history/user-operation-log/get-user-operation-log-query/)
method.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryUserOperationCountRequest
*/
func (a *HistoricUserOperationLogAPIService) QueryUserOperationCount(ctx context.Context) ApiQueryUserOperationCountRequest {
	return ApiQueryUserOperationCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CountResultDto
func (a *HistoricUserOperationLogAPIService) QueryUserOperationCountExecute(r ApiQueryUserOperationCountRequest) (*CountResultDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CountResultDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricUserOperationLogAPIService.QueryUserOperationCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/history/user-operation/count"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.deploymentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deploymentId", r.deploymentId, "")
	}
	if r.processDefinitionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "processDefinitionId", r.processDefinitionId, "")
	}
	if r.processDefinitionKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "processDefinitionKey", r.processDefinitionKey, "")
	}
	if r.processInstanceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "processInstanceId", r.processInstanceId, "")
	}
	if r.executionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "executionId", r.executionId, "")
	}
	if r.caseDefinitionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "caseDefinitionId", r.caseDefinitionId, "")
	}
	if r.caseInstanceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "caseInstanceId", r.caseInstanceId, "")
	}
	if r.caseExecutionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "caseExecutionId", r.caseExecutionId, "")
	}
	if r.taskId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "taskId", r.taskId, "")
	}
	if r.externalTaskId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "externalTaskId", r.externalTaskId, "")
	}
	if r.batchId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "batchId", r.batchId, "")
	}
	if r.jobId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "jobId", r.jobId, "")
	}
	if r.jobDefinitionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "jobDefinitionId", r.jobDefinitionId, "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
	}
	if r.operationId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "operationId", r.operationId, "")
	}
	if r.operationType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "operationType", r.operationType, "")
	}
	if r.entityType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entityType", r.entityType, "")
	}
	if r.entityTypeIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entityTypeIn", r.entityTypeIn, "")
	}
	if r.category != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category", r.category, "")
	}
	if r.categoryIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "categoryIn", r.categoryIn, "")
	}
	if r.property != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "property", r.property, "")
	}
	if r.afterTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "afterTimestamp", r.afterTimestamp, "")
	}
	if r.beforeTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "beforeTimestamp", r.beforeTimestamp, "")
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

type ApiQueryUserOperationEntriesRequest struct {
	ctx context.Context
	ApiService HistoricUserOperationLogAPI
	deploymentId *string
	processDefinitionId *string
	processDefinitionKey *string
	processInstanceId *string
	executionId *string
	caseDefinitionId *string
	caseInstanceId *string
	caseExecutionId *string
	taskId *string
	externalTaskId *string
	batchId *string
	jobId *string
	jobDefinitionId *string
	userId *string
	operationId *string
	operationType *string
	entityType *string
	entityTypeIn *string
	category *string
	categoryIn *string
	property *string
	afterTimestamp *time.Time
	beforeTimestamp *time.Time
	sortBy *string
	sortOrder *string
	firstResult *int32
	maxResults *int32
}

// Filter by deployment id.
func (r ApiQueryUserOperationEntriesRequest) DeploymentId(deploymentId string) ApiQueryUserOperationEntriesRequest {
	r.deploymentId = &deploymentId
	return r
}

// Filter by process definition id.
func (r ApiQueryUserOperationEntriesRequest) ProcessDefinitionId(processDefinitionId string) ApiQueryUserOperationEntriesRequest {
	r.processDefinitionId = &processDefinitionId
	return r
}

// Filter by process definition key.
func (r ApiQueryUserOperationEntriesRequest) ProcessDefinitionKey(processDefinitionKey string) ApiQueryUserOperationEntriesRequest {
	r.processDefinitionKey = &processDefinitionKey
	return r
}

// Filter by process instance id.
func (r ApiQueryUserOperationEntriesRequest) ProcessInstanceId(processInstanceId string) ApiQueryUserOperationEntriesRequest {
	r.processInstanceId = &processInstanceId
	return r
}

// Filter by execution id.
func (r ApiQueryUserOperationEntriesRequest) ExecutionId(executionId string) ApiQueryUserOperationEntriesRequest {
	r.executionId = &executionId
	return r
}

// Filter by case definition id.
func (r ApiQueryUserOperationEntriesRequest) CaseDefinitionId(caseDefinitionId string) ApiQueryUserOperationEntriesRequest {
	r.caseDefinitionId = &caseDefinitionId
	return r
}

// Filter by case instance id.
func (r ApiQueryUserOperationEntriesRequest) CaseInstanceId(caseInstanceId string) ApiQueryUserOperationEntriesRequest {
	r.caseInstanceId = &caseInstanceId
	return r
}

// Filter by case execution id.
func (r ApiQueryUserOperationEntriesRequest) CaseExecutionId(caseExecutionId string) ApiQueryUserOperationEntriesRequest {
	r.caseExecutionId = &caseExecutionId
	return r
}

// Only include operations on this task.
func (r ApiQueryUserOperationEntriesRequest) TaskId(taskId string) ApiQueryUserOperationEntriesRequest {
	r.taskId = &taskId
	return r
}

// Only include operations on this external task.
func (r ApiQueryUserOperationEntriesRequest) ExternalTaskId(externalTaskId string) ApiQueryUserOperationEntriesRequest {
	r.externalTaskId = &externalTaskId
	return r
}

// Only include operations on this batch.
func (r ApiQueryUserOperationEntriesRequest) BatchId(batchId string) ApiQueryUserOperationEntriesRequest {
	r.batchId = &batchId
	return r
}

// Filter by job id.
func (r ApiQueryUserOperationEntriesRequest) JobId(jobId string) ApiQueryUserOperationEntriesRequest {
	r.jobId = &jobId
	return r
}

// Filter by job definition id.
func (r ApiQueryUserOperationEntriesRequest) JobDefinitionId(jobDefinitionId string) ApiQueryUserOperationEntriesRequest {
	r.jobDefinitionId = &jobDefinitionId
	return r
}

// Only include operations of this user.
func (r ApiQueryUserOperationEntriesRequest) UserId(userId string) ApiQueryUserOperationEntriesRequest {
	r.userId = &userId
	return r
}

// Filter by the id of the operation. This allows fetching of multiple entries which are part of a composite operation.
func (r ApiQueryUserOperationEntriesRequest) OperationId(operationId string) ApiQueryUserOperationEntriesRequest {
	r.operationId = &operationId
	return r
}

// Filter by the type of the operation like &#x60;Claim&#x60; or &#x60;Delegate&#x60;. See the [Javadoc](https://docs.camunda.org/manual/7.21/reference/javadoc/?org/camunda/bpm/engine/history/UserOperationLogEntry.html) for a list of available operation types.
func (r ApiQueryUserOperationEntriesRequest) OperationType(operationType string) ApiQueryUserOperationEntriesRequest {
	r.operationType = &operationType
	return r
}

// Filter by the type of the entity that was affected by this operation, possible values are &#x60;Task&#x60;, &#x60;Attachment&#x60; or &#x60;IdentityLink&#x60;.
func (r ApiQueryUserOperationEntriesRequest) EntityType(entityType string) ApiQueryUserOperationEntriesRequest {
	r.entityType = &entityType
	return r
}

// Filter by a comma-separated list of types of the entities that was affected by this operation, possible values are &#x60;Task&#x60;, &#x60;Attachment&#x60; or &#x60;IdentityLink&#x60;.
func (r ApiQueryUserOperationEntriesRequest) EntityTypeIn(entityTypeIn string) ApiQueryUserOperationEntriesRequest {
	r.entityTypeIn = &entityTypeIn
	return r
}

// Filter by the category that this operation is associated with, possible values are &#x60;TaskWorker&#x60;, &#x60;Admin&#x60; or &#x60;Operator&#x60;.
func (r ApiQueryUserOperationEntriesRequest) Category(category string) ApiQueryUserOperationEntriesRequest {
	r.category = &category
	return r
}

// Filter by a comma-separated list of categories that this operation is associated with, possible values are &#x60;TaskWorker&#x60;, &#x60;Admin&#x60; or &#x60;Operator&#x60;.
func (r ApiQueryUserOperationEntriesRequest) CategoryIn(categoryIn string) ApiQueryUserOperationEntriesRequest {
	r.categoryIn = &categoryIn
	return r
}

// Only include operations that changed this property, e.g., &#x60;owner&#x60; or &#x60;assignee&#x60;.
func (r ApiQueryUserOperationEntriesRequest) Property(property string) ApiQueryUserOperationEntriesRequest {
	r.property = &property
	return r
}

// Restrict to entries that were created after the given timestamp. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the timestamp must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200.
func (r ApiQueryUserOperationEntriesRequest) AfterTimestamp(afterTimestamp time.Time) ApiQueryUserOperationEntriesRequest {
	r.afterTimestamp = &afterTimestamp
	return r
}

// Restrict to entries that were created before the given timestamp. By [default](https://docs.camunda.org/manual/7.21/reference/rest/overview/date-format/), the timestamp must have the format &#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&#x60;, e.g., 2013-01-23T14:42:45.000+0200.
func (r ApiQueryUserOperationEntriesRequest) BeforeTimestamp(beforeTimestamp time.Time) ApiQueryUserOperationEntriesRequest {
	r.beforeTimestamp = &beforeTimestamp
	return r
}

// Sort the results lexicographically by a given criterion. Must be used in conjunction with the sortOrder parameter.
func (r ApiQueryUserOperationEntriesRequest) SortBy(sortBy string) ApiQueryUserOperationEntriesRequest {
	r.sortBy = &sortBy
	return r
}

// Sort the results in a given order. Values may be asc for ascending order or desc for descending order. Must be used in conjunction with the sortBy parameter.
func (r ApiQueryUserOperationEntriesRequest) SortOrder(sortOrder string) ApiQueryUserOperationEntriesRequest {
	r.sortOrder = &sortOrder
	return r
}

// Pagination of results. Specifies the index of the first result to return.
func (r ApiQueryUserOperationEntriesRequest) FirstResult(firstResult int32) ApiQueryUserOperationEntriesRequest {
	r.firstResult = &firstResult
	return r
}

// Pagination of results. Specifies the maximum number of results to return. Will return less results if there are no more results left.
func (r ApiQueryUserOperationEntriesRequest) MaxResults(maxResults int32) ApiQueryUserOperationEntriesRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiQueryUserOperationEntriesRequest) Execute() ([]UserOperationLogEntryDto, *http.Response, error) {
	return r.ApiService.QueryUserOperationEntriesExecute(r)
}

/*
QueryUserOperationEntries Get User Operation Log (Historic)

Queries for user operation log entries that fulfill the given parameters.
The size of the result set can be retrieved by using the
[Get User Operation Log Count](https://docs.camunda.org/manual/7.21/reference/rest/history/user-operation-log/get-user-operation-log-query-count/)
method.

Note that the properties of operation log entries are interpreted as
restrictions on the entities they apply to. That means, if a single
process instance is updated, the field `processInstanceId` is
populated. If a single operation updates all process instances of the
same process definition, the field `processInstanceId` is `null` (a
`null` restriction is viewed as a wildcard, i.e., matches a process
instance with any id) and the field `processDefinitionId` is
populated. This way, which entities were changed by a user operation
can easily be reconstructed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryUserOperationEntriesRequest
*/
func (a *HistoricUserOperationLogAPIService) QueryUserOperationEntries(ctx context.Context) ApiQueryUserOperationEntriesRequest {
	return ApiQueryUserOperationEntriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []UserOperationLogEntryDto
func (a *HistoricUserOperationLogAPIService) QueryUserOperationEntriesExecute(r ApiQueryUserOperationEntriesRequest) ([]UserOperationLogEntryDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []UserOperationLogEntryDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricUserOperationLogAPIService.QueryUserOperationEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/history/user-operation"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.deploymentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deploymentId", r.deploymentId, "")
	}
	if r.processDefinitionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "processDefinitionId", r.processDefinitionId, "")
	}
	if r.processDefinitionKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "processDefinitionKey", r.processDefinitionKey, "")
	}
	if r.processInstanceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "processInstanceId", r.processInstanceId, "")
	}
	if r.executionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "executionId", r.executionId, "")
	}
	if r.caseDefinitionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "caseDefinitionId", r.caseDefinitionId, "")
	}
	if r.caseInstanceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "caseInstanceId", r.caseInstanceId, "")
	}
	if r.caseExecutionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "caseExecutionId", r.caseExecutionId, "")
	}
	if r.taskId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "taskId", r.taskId, "")
	}
	if r.externalTaskId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "externalTaskId", r.externalTaskId, "")
	}
	if r.batchId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "batchId", r.batchId, "")
	}
	if r.jobId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "jobId", r.jobId, "")
	}
	if r.jobDefinitionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "jobDefinitionId", r.jobDefinitionId, "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
	}
	if r.operationId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "operationId", r.operationId, "")
	}
	if r.operationType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "operationType", r.operationType, "")
	}
	if r.entityType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entityType", r.entityType, "")
	}
	if r.entityTypeIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entityTypeIn", r.entityTypeIn, "")
	}
	if r.category != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category", r.category, "")
	}
	if r.categoryIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "categoryIn", r.categoryIn, "")
	}
	if r.property != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "property", r.property, "")
	}
	if r.afterTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "afterTimestamp", r.afterTimestamp, "")
	}
	if r.beforeTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "beforeTimestamp", r.beforeTimestamp, "")
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

type ApiSetAnnotationUserOperationLogRequest struct {
	ctx context.Context
	ApiService HistoricUserOperationLogAPI
	operationId string
	annotationDto *AnnotationDto
}

func (r ApiSetAnnotationUserOperationLogRequest) AnnotationDto(annotationDto AnnotationDto) ApiSetAnnotationUserOperationLogRequest {
	r.annotationDto = &annotationDto
	return r
}

func (r ApiSetAnnotationUserOperationLogRequest) Execute() (*http.Response, error) {
	return r.ApiService.SetAnnotationUserOperationLogExecute(r)
}

/*
SetAnnotationUserOperationLog Set Annotation to an User Operation Log (Historic)

Set an annotation for auditing reasons.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param operationId The operation id of the operation log to be updated.
 @return ApiSetAnnotationUserOperationLogRequest
*/
func (a *HistoricUserOperationLogAPIService) SetAnnotationUserOperationLog(ctx context.Context, operationId string) ApiSetAnnotationUserOperationLogRequest {
	return ApiSetAnnotationUserOperationLogRequest{
		ApiService: a,
		ctx: ctx,
		operationId: operationId,
	}
}

// Execute executes the request
func (a *HistoricUserOperationLogAPIService) SetAnnotationUserOperationLogExecute(r ApiSetAnnotationUserOperationLogRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricUserOperationLogAPIService.SetAnnotationUserOperationLog")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/history/user-operation/{operationId}/set-annotation"
	localVarPath = strings.Replace(localVarPath, "{"+"operationId"+"}", url.PathEscape(parameterValueToString(r.operationId, "operationId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.annotationDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
