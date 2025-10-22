<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := provesdkservergo.New(
		provesdkservergo.WithSecurity(components.Security{
			ClientID:     provesdkservergo.Pointer("<YOUR_CLIENT_ID_HERE>"),
			ClientSecret: provesdkservergo.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
		}),
	)

	res, err := s.V3.V3StartRequest(ctx, &components.V3StartRequest{
		AllowOTPRetry:  provesdkservergo.Pointer(true),
		Dob:            provesdkservergo.Pointer("1981-01"),
		EmailAddress:   provesdkservergo.Pointer("mpinsonm@dyndns.org"),
		FinalTargetURL: provesdkservergo.Pointer("https://www.example.com/landing-page"),
		FlowType:       "mobile",
		IPAddress:      provesdkservergo.Pointer("10.0.0.1"),
		PhoneNumber:    provesdkservergo.Pointer("2001001695"),
		SmsMessage:     provesdkservergo.Pointer("#### is your temporary code to continue your application. Caution: for your security, don't share this code with anyone."),
		Ssn:            provesdkservergo.Pointer("0596"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.V3StartResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->