// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/payfone/prove-sdk-server-go/models/components"
)

type V3ValidateRequestResponse struct {
	HTTPMeta components.HTTPMetadata
	// Successful request.
	V3ValidateResponse *components.V3ValidateResponse
}

func (o *V3ValidateRequestResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *V3ValidateRequestResponse) GetV3ValidateResponse() *components.V3ValidateResponse {
	if o == nil {
		return nil
	}
	return o.V3ValidateResponse
}