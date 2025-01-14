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
)


type HistoricDecisionRequirementsDefinitionAPI interface {

	/*
	GetDecisionStatistics Get DRD Statistics

	Retrieves evaluation statistics of a given decision requirements definition.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The id of the decision requirements definition.
	@return ApiGetDecisionStatisticsRequest
	*/
	GetDecisionStatistics(ctx context.Context, id string) ApiGetDecisionStatisticsRequest

	// GetDecisionStatisticsExecute executes the request
	//  @return []HistoricDecisionInstanceStatisticsDto
	GetDecisionStatisticsExecute(r ApiGetDecisionStatisticsRequest) ([]HistoricDecisionInstanceStatisticsDto, *http.Response, error)
}

// HistoricDecisionRequirementsDefinitionAPIService HistoricDecisionRequirementsDefinitionAPI service
type HistoricDecisionRequirementsDefinitionAPIService service

type ApiGetDecisionStatisticsRequest struct {
	ctx context.Context
	ApiService HistoricDecisionRequirementsDefinitionAPI
	id string
	decisionInstanceId *string
}

// Restrict query results to be based only on specific evaluation instance of a given decision requirements definition.
func (r ApiGetDecisionStatisticsRequest) DecisionInstanceId(decisionInstanceId string) ApiGetDecisionStatisticsRequest {
	r.decisionInstanceId = &decisionInstanceId
	return r
}

func (r ApiGetDecisionStatisticsRequest) Execute() ([]HistoricDecisionInstanceStatisticsDto, *http.Response, error) {
	return r.ApiService.GetDecisionStatisticsExecute(r)
}

/*
GetDecisionStatistics Get DRD Statistics

Retrieves evaluation statistics of a given decision requirements definition.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The id of the decision requirements definition.
 @return ApiGetDecisionStatisticsRequest
*/
func (a *HistoricDecisionRequirementsDefinitionAPIService) GetDecisionStatistics(ctx context.Context, id string) ApiGetDecisionStatisticsRequest {
	return ApiGetDecisionStatisticsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return []HistoricDecisionInstanceStatisticsDto
func (a *HistoricDecisionRequirementsDefinitionAPIService) GetDecisionStatisticsExecute(r ApiGetDecisionStatisticsRequest) ([]HistoricDecisionInstanceStatisticsDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []HistoricDecisionInstanceStatisticsDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoricDecisionRequirementsDefinitionAPIService.GetDecisionStatistics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/history/decision-requirements-definition/{id}/statistics"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.decisionInstanceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "decisionInstanceId", r.decisionInstanceId, "")
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
