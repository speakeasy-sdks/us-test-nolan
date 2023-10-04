// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/us-test-nolan/pkg/models/shared"
	"net/http"
)

type Pipeline1GeneralV0031GeneralPostRequest struct {
	PipelineBodyV0031  *shared.PipelineBodyV0031 `request:"mediaType=multipart/form-data"`
	UnstructuredAPIKey *string                   `header:"style=simple,explode=false,name=unstructured-api-key"`
}

func (o *Pipeline1GeneralV0031GeneralPostRequest) GetPipelineBodyV0031() *shared.PipelineBodyV0031 {
	if o == nil {
		return nil
	}
	return o.PipelineBodyV0031
}

func (o *Pipeline1GeneralV0031GeneralPostRequest) GetUnstructuredAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.UnstructuredAPIKey
}

type Pipeline1GeneralV0031GeneralPostResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Validation Error
	HTTPValidationError *shared.HTTPValidationError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *Pipeline1GeneralV0031GeneralPostResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *Pipeline1GeneralV0031GeneralPostResponse) GetHTTPValidationError() *shared.HTTPValidationError {
	if o == nil {
		return nil
	}
	return o.HTTPValidationError
}

func (o *Pipeline1GeneralV0031GeneralPostResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *Pipeline1GeneralV0031GeneralPostResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
