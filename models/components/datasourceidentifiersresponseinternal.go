// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DataSourceIdentifiersResponseInternal struct {
	Dob                  *bool `json:"dob,omitempty"`
	DriversLicenseNumber *bool `json:"driversLicenseNumber,omitempty"`
	DriversLicenseState  *bool `json:"driversLicenseState,omitempty"`
	Last4                *bool `json:"last4,omitempty"`
	Ssn                  *bool `json:"ssn,omitempty"`
}

func (o *DataSourceIdentifiersResponseInternal) GetDob() *bool {
	if o == nil {
		return nil
	}
	return o.Dob
}

func (o *DataSourceIdentifiersResponseInternal) GetDriversLicenseNumber() *bool {
	if o == nil {
		return nil
	}
	return o.DriversLicenseNumber
}

func (o *DataSourceIdentifiersResponseInternal) GetDriversLicenseState() *bool {
	if o == nil {
		return nil
	}
	return o.DriversLicenseState
}

func (o *DataSourceIdentifiersResponseInternal) GetLast4() *bool {
	if o == nil {
		return nil
	}
	return o.Last4
}

func (o *DataSourceIdentifiersResponseInternal) GetSsn() *bool {
	if o == nil {
		return nil
	}
	return o.Ssn
}
