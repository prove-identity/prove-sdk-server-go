// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// V3MFABindRequest - Request body for the V3 MFA Bind API
type V3MFABindRequest struct {
	// Client Request ID is a client-generated unique ID for a specific request.
	// This can be used by clients to identify specific requests made to Prove Link.
	// The format of this ID is defined by the client - Prove recommends using a GUID,
	// but any format can be accepted.
	ClientRequestID *string `json:"clientRequestId,omitempty"`
	// Correlation ID is the unique ID that Prove generates for the flow. It is returned
	// from the v3/mfa endpoint and cannot be reused outside of a single flow.
	CorrelationID *string `json:"correlationId,omitempty"`
	// Phone number is only allowed when possessionType=none from the initial MFA request.
	// Required for BYO Possession flow on the third call.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

func (o *V3MFABindRequest) GetClientRequestID() *string {
	if o == nil {
		return nil
	}
	return o.ClientRequestID
}

func (o *V3MFABindRequest) GetCorrelationID() *string {
	if o == nil {
		return nil
	}
	return o.CorrelationID
}

func (o *V3MFABindRequest) GetPhoneNumber() *string {
	if o == nil {
		return nil
	}
	return o.PhoneNumber
}
