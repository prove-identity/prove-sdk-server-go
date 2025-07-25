// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// V3DisenrollIdentityResponse - Response body for the V3 Disenroll Identity API.
type V3DisenrollIdentityResponse struct {
	// If true, the disenroll operation was successful.
	Success bool `json:"success"`
}

func (o *V3DisenrollIdentityResponse) GetSuccess() bool {
	if o == nil {
		return false
	}
	return o.Success
}
