// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AmlTypeListResponseInternal struct {
	AmlType  *string                        `json:"amlType,omitempty"`
	Fields   []KYCFieldTypeResponseInternal `json:"fields,omitempty"`
	ListHits *int64                         `json:"listHits,omitempty"`
}

func (o *AmlTypeListResponseInternal) GetAmlType() *string {
	if o == nil {
		return nil
	}
	return o.AmlType
}

func (o *AmlTypeListResponseInternal) GetFields() []KYCFieldTypeResponseInternal {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *AmlTypeListResponseInternal) GetListHits() *int64 {
	if o == nil {
		return nil
	}
	return o.ListHits
}
