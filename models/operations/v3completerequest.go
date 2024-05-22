// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/payfone/prove-sdk-server-go/models/components"
)

type V3CompleteRequestResponse struct {
	HTTPMeta components.HTTPMetadata
	// Successful request.
	V3CompleteResponse *components.V3CompleteResponse
}

func (o *V3CompleteRequestResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *V3CompleteRequestResponse) GetV3CompleteResponse() *components.V3CompleteResponse {
	if o == nil {
		return nil
	}
	return o.V3CompleteResponse
}