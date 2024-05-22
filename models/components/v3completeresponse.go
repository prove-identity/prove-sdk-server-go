// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type V3CompleteResponse struct {
	// ChangeDetected returns true if the individual information changed.
	ChangeDetected bool `json:"changeDetected"`
	// Next contains the next set of allowed calls in the same flow.
	Next map[string]string `json:"next"`
	// Success returns true if the individual was verified successfully.
	Success bool `json:"success"`
}

func (o *V3CompleteResponse) GetChangeDetected() bool {
	if o == nil {
		return false
	}
	return o.ChangeDetected
}

func (o *V3CompleteResponse) GetNext() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Next
}

func (o *V3CompleteResponse) GetSuccess() bool {
	if o == nil {
		return false
	}
	return o.Success
}