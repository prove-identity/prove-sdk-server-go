// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// KnowYourCustomerResponse - KYC response
type KnowYourCustomerResponse struct {
	AdverseMediaList []AdverseMediaResponse `json:"AdverseMediaList,omitempty"`
	AliasList        []string               `json:"AliasList,omitempty"`
	AmlTypeLists     []AmlTypeListResponse  `json:"AmlTypeLists,omitempty"`
	TotalHits        *int64                 `json:"TotalHits,omitempty"`
}

func (o *KnowYourCustomerResponse) GetAdverseMediaList() []AdverseMediaResponse {
	if o == nil {
		return nil
	}
	return o.AdverseMediaList
}

func (o *KnowYourCustomerResponse) GetAliasList() []string {
	if o == nil {
		return nil
	}
	return o.AliasList
}

func (o *KnowYourCustomerResponse) GetAmlTypeLists() []AmlTypeListResponse {
	if o == nil {
		return nil
	}
	return o.AmlTypeLists
}

func (o *KnowYourCustomerResponse) GetTotalHits() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalHits
}
