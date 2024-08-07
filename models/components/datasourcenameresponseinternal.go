// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DataSourceNameResponseInternal struct {
	FirstName *int64 `json:"firstName,omitempty"`
	LastName  *int64 `json:"lastName,omitempty"`
	NameScore *int64 `json:"nameScore,omitempty"`
}

func (o *DataSourceNameResponseInternal) GetFirstName() *int64 {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *DataSourceNameResponseInternal) GetLastName() *int64 {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *DataSourceNameResponseInternal) GetNameScore() *int64 {
	if o == nil {
		return nil
	}
	return o.NameScore
}
