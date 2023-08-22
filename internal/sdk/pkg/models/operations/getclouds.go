// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"MyTest/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetCloudsRequest struct {
	// Morpheus ID of the Object being referenced
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

// GetClouds200ApplicationJSON - Successful Request
type GetClouds200ApplicationJSON struct {
	Zone *shared.Zone `json:"zone,omitempty"`
}

type GetCloudsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Error Codes
	DefaultError *shared.DefaultError
	// Successful Request
	GetClouds200ApplicationJSONObject *GetClouds200ApplicationJSON
}