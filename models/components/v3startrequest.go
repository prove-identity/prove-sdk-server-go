// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type V3StartRequest struct {
	// DOB, an optional challenge, is the date of birth in one of these formats: YYYY-MM-DD, YYYY-MM, or MM-DD. Acceptable characters are: numeric with symbol '-'.
	Dob *string `json:"dob,omitempty"`
	// Email is the email address of the customer. Acceptable characters are: alphanumeric with symbols '@.+'.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Final target URL is only required for when flowType=desktop. The final target URL is where the end user will be redirected at the end of Instant Link flow. Acceptable characters are: alphanumeric with symbols '-._+=/:?'.
	FinalTargetURL *string `json:"finalTargetUrl,omitempty"`
	// Flow type is based on the method used - either 'desktop' if using desktop or 'mobile' for iOS/Android native apps and mobile web. Acceptable options are: 'desktop' or 'mobile'.
	FlowType string `json:"flowType"`
	// IP address is the IP address of the device of the customer. Acceptable characters are: numeric with symbols ':.'.
	IPAddress *string `json:"ipAddress,omitempty"`
	// Phone number is the number of the mobile phone. The field is required in the Sandbox environment. In Production, you will likely pass the phone number via the Prove Link client SDK instead of within the Start call depending on how your user experience is implemented. Acceptable characters are: alphanumeric with symbols '+'.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// SMSMessage is an optional field to customize the message body sent in the Instant Link (flowType=desktop) or OTP (on mobile) SMS message.
	// If not provided, the following default messages will be used:
	// 1. For Instant Link: "Complete your verification. If you did not make this request, do not click the link. ####"
	// 2. For OTP: "#### is your temporary code to continue your application. Caution: for your security, don't share this code with anyone."
	// Max length is 160 characters. Only ASCII characters are allowed.
	//
	// The placeholder format varies by flow type:
	// 1. For OTP (mobile flow): Use ####, #####, or ###### to generate 4-6 digit verification codes respectively.
	// 2. For Instant Link (desktop flow): Must use exactly #### which will be replaced with the verification URL.
	SmsMessage *string `json:"smsMessage,omitempty"`
	// SSN, an optional challenge, is either the full or last 4 digits of the social security number. Acceptable characters are: numeric.
	Ssn *string `json:"ssn,omitempty"`
}

func (o *V3StartRequest) GetDob() *string {
	if o == nil {
		return nil
	}
	return o.Dob
}

func (o *V3StartRequest) GetEmailAddress() *string {
	if o == nil {
		return nil
	}
	return o.EmailAddress
}

func (o *V3StartRequest) GetFinalTargetURL() *string {
	if o == nil {
		return nil
	}
	return o.FinalTargetURL
}

func (o *V3StartRequest) GetFlowType() string {
	if o == nil {
		return ""
	}
	return o.FlowType
}

func (o *V3StartRequest) GetIPAddress() *string {
	if o == nil {
		return nil
	}
	return o.IPAddress
}

func (o *V3StartRequest) GetPhoneNumber() *string {
	if o == nil {
		return nil
	}
	return o.PhoneNumber
}

func (o *V3StartRequest) GetSmsMessage() *string {
	if o == nil {
		return nil
	}
	return o.SmsMessage
}

func (o *V3StartRequest) GetSsn() *string {
	if o == nil {
		return nil
	}
	return o.Ssn
}
