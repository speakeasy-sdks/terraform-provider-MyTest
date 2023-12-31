// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"MyTest/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateCloudsRequestBody struct {
	Zone shared.Zone `json:"zone"`
}

type UpdateCloudsRequest struct {
	RequestBody *UpdateCloudsRequestBody `request:"mediaType=application/json"`
	// Morpheus ID of the Object being referenced
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

// UpdateClouds200ApplicationJSON - Successful Request
type UpdateClouds200ApplicationJSON struct {
	Zone *shared.Zone `json:"zone,omitempty"`
}

type UpdateCloudsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error Codes
	DefaultError *shared.DefaultError
	// Successful Request
	UpdateClouds200ApplicationJSONObject *UpdateClouds200ApplicationJSON
}
