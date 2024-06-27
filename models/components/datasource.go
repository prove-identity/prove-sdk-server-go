// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type DataSource struct {
	Address       *AddressResponse      `json:"address,omitempty"`
	CipConfidence *string               `json:"cipConfidence,omitempty"`
	Email         *EmailAddressResponse `json:"email,omitempty"`
	Identifiers   *IdentifiersResponse  `json:"identifiers,omitempty"`
	Name          *NameResponse         `json:"name,omitempty"`
	ReasonCodes   []string              `json:"reasonCodes,omitempty"`
	Verified      *bool                 `json:"verified,omitempty"`
}

func (o *DataSource) GetAddress() *AddressResponse {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *DataSource) GetCipConfidence() *string {
	if o == nil {
		return nil
	}
	return o.CipConfidence
}

func (o *DataSource) GetEmail() *EmailAddressResponse {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *DataSource) GetIdentifiers() *IdentifiersResponse {
	if o == nil {
		return nil
	}
	return o.Identifiers
}

func (o *DataSource) GetName() *NameResponse {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *DataSource) GetReasonCodes() []string {
	if o == nil {
		return nil
	}
	return o.ReasonCodes
}

func (o *DataSource) GetVerified() *bool {
	if o == nil {
		return nil
	}
	return o.Verified
}
