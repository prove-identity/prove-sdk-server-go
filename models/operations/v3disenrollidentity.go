// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/prove-identity/prove-sdk-server-go/models/components"
)

type V3DisenrollIdentityRequest struct {
	// A Prove-generated unique ID for a specific identity.
	IdentityID string `pathParam:"style=simple,explode=false,name=identityId"`
	// A client-generated unique ID for a specific session. This can be used to identify specific requests. The format of this ID is defined by the client - Prove recommends using a GUID, but any format can be accepted. Do not include Personally Identifiable Information (PII) in this field.
	ClientRequestID *string `queryParam:"style=form,explode=true,name=clientRequestId"`
}

func (o *V3DisenrollIdentityRequest) GetIdentityID() string {
	if o == nil {
		return ""
	}
	return o.IdentityID
}

func (o *V3DisenrollIdentityRequest) GetClientRequestID() *string {
	if o == nil {
		return nil
	}
	return o.ClientRequestID
}

type V3DisenrollIdentityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// V3DisenrollIdentityResponse
	V3DisenrollIdentityResponse *components.V3DisenrollIdentityResponse
}

func (o *V3DisenrollIdentityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *V3DisenrollIdentityResponse) GetV3DisenrollIdentityResponse() *components.V3DisenrollIdentityResponse {
	if o == nil {
		return nil
	}
	return o.V3DisenrollIdentityResponse
}
