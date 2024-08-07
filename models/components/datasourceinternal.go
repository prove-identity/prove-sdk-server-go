// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DataSourceInternal struct {
	Address       *DataSourceAddressResponseInternal      `json:"address,omitempty"`
	CipConfidence *string                                 `json:"cipConfidence,omitempty"`
	Email         *DataSourceEmailAddressResponseInternal `json:"email,omitempty"`
	Identifiers   *DataSourceIdentifiersResponseInternal  `json:"identifiers,omitempty"`
	Name          *DataSourceNameResponseInternal         `json:"name,omitempty"`
	ReasonCodes   []string                                `json:"reasonCodes,omitempty"`
	Verified      *bool                                   `json:"verified,omitempty"`
}

func (o *DataSourceInternal) GetAddress() *DataSourceAddressResponseInternal {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *DataSourceInternal) GetCipConfidence() *string {
	if o == nil {
		return nil
	}
	return o.CipConfidence
}

func (o *DataSourceInternal) GetEmail() *DataSourceEmailAddressResponseInternal {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *DataSourceInternal) GetIdentifiers() *DataSourceIdentifiersResponseInternal {
	if o == nil {
		return nil
	}
	return o.Identifiers
}

func (o *DataSourceInternal) GetName() *DataSourceNameResponseInternal {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *DataSourceInternal) GetReasonCodes() []string {
	if o == nil {
		return nil
	}
	return o.ReasonCodes
}

func (o *DataSourceInternal) GetVerified() *bool {
	if o == nil {
		return nil
	}
	return o.Verified
}
