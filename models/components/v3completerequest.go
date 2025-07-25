// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type V3CompleteRequest struct {
	// The unique ID that Prove generates for the flow. It is returned from the Start endpoint and cannot be reused outside of a single flow.
	CorrelationID string                      `json:"correlationId"`
	Individual    V3CompleteIndividualRequest `json:"individual"`
}

func (o *V3CompleteRequest) GetCorrelationID() string {
	if o == nil {
		return ""
	}
	return o.CorrelationID
}

func (o *V3CompleteRequest) GetIndividual() V3CompleteIndividualRequest {
	if o == nil {
		return V3CompleteIndividualRequest{}
	}
	return o.Individual
}
