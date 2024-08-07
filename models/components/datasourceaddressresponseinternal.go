// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DataSourceAddressResponseInternal struct {
	AddressScore *int64   `json:"addressScore,omitempty"`
	City         *bool    `json:"city,omitempty"`
	Distance     *float64 `json:"distance,omitempty"`
	PostalCode   *bool    `json:"postalCode,omitempty"`
	Region       *bool    `json:"region,omitempty"`
	Street       *bool    `json:"street,omitempty"`
	StreetNumber *int64   `json:"streetNumber,omitempty"`
}

func (o *DataSourceAddressResponseInternal) GetAddressScore() *int64 {
	if o == nil {
		return nil
	}
	return o.AddressScore
}

func (o *DataSourceAddressResponseInternal) GetCity() *bool {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *DataSourceAddressResponseInternal) GetDistance() *float64 {
	if o == nil {
		return nil
	}
	return o.Distance
}

func (o *DataSourceAddressResponseInternal) GetPostalCode() *bool {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *DataSourceAddressResponseInternal) GetRegion() *bool {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *DataSourceAddressResponseInternal) GetStreet() *bool {
	if o == nil {
		return nil
	}
	return o.Street
}

func (o *DataSourceAddressResponseInternal) GetStreetNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.StreetNumber
}
