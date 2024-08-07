// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type V3ChallengeAddressEntryRequest struct {
	// Address is the street address of the individual.
	Address *string `json:"address,omitempty"`
	// City of the individual.
	City *string `json:"city,omitempty"`
	// Extended address is the apartment number or other extended address information.
	ExtendedAddress *string `json:"extendedAddress,omitempty"`
	// Postal code is the zip code of the individual. It can be either 5 digits (XXXXX) or ZIP+4 (XXXXX-XXXX).
	PostalCode *string `json:"postalCode,omitempty"`
	// Region is the state or locality of the individual.
	Region *string `json:"region,omitempty"`
}

func (o *V3ChallengeAddressEntryRequest) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *V3ChallengeAddressEntryRequest) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *V3ChallengeAddressEntryRequest) GetExtendedAddress() *string {
	if o == nil {
		return nil
	}
	return o.ExtendedAddress
}

func (o *V3ChallengeAddressEntryRequest) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *V3ChallengeAddressEntryRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}
