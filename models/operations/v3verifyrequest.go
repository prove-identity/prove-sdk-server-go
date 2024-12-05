// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/prove-identity/prove-sdk-server-go/models/components"
)

type V3VerifyRequestResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful request.
	V3VerifyResponse *components.V3VerifyResponse
}

func (o *V3VerifyRequestResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *V3VerifyRequestResponse) GetV3VerifyResponse() *components.V3VerifyResponse {
	if o == nil {
		return nil
	}
	return o.V3VerifyResponse
}